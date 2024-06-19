/* Copyright (c) 2019 Snowflake Inc. All rights reserved.

   Licensed under the Apache License, Version 2.0 (the
   "License"); you may not use this file except in compliance
   with the License.  You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing,
   software distributed under the License is distributed on an
   "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
   KIND, either express or implied.  See the License for the
   specific language governing permissions and limitations
   under the License.
*/

// Package client provides the client interface for 'healthcheck'
package client

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/google/subcommands"

	"github.com/Snowflake-Labs/sansshell/client"
	pb "github.com/Snowflake-Labs/sansshell/services/psql"
	"github.com/Snowflake-Labs/sansshell/services/util"
	"github.com/jedib0t/go-pretty/v6/table"
)

const subPackage = "psql"

func init() {
	subcommands.Register(&psqlCmd{}, subPackage)
}

func (*psqlCmd) GetSubpackage(f *flag.FlagSet) *subcommands.Commander {
	c := client.SetupSubpackage(subPackage, f)
	c.Register(&execCmd{}, "")
	c.Register(&queryCmd{}, "")
	return c
}

type psqlCmd struct{}

func (*psqlCmd) Name() string { return subPackage }
func (p *psqlCmd) Synopsis() string {
	return client.GenerateSynopsis(p.GetSubpackage(flag.NewFlagSet("", flag.ContinueOnError)), 2)
}
func (p *psqlCmd) Usage() string {
	return client.GenerateUsage(subPackage, p.Synopsis())
}
func (*psqlCmd) SetFlags(f *flag.FlagSet) {}

func (p *psqlCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	c := p.GetSubpackage(f)
	return c.Execute(ctx, args...)
}

type execCmd struct{}

func (*execCmd) Name() string     { return "exec" }
func (*execCmd) Synopsis() string { return "Run a sql statement." }
func (*execCmd) Usage() string {
	return `exec "SQL STATEMENT":
	Run a sql statement.
`
}

func (p *execCmd) SetFlags(f *flag.FlagSet) {
	// ideas: read from file, maybe params
}

func (p *execCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	if f.NArg() != 1 {
		fmt.Fprintln(os.Stderr, "Please specify a quoted SQL string.")
		return subcommands.ExitUsageError
	}

	state := args[0].(*util.ExecuteState)
	c := pb.NewPsqlClientProxy(state.Conn)

	resp, err := c.ExecOneMany(ctx, &pb.ExecRequest{Query: f.Args()[0]})
	if err != nil {
		// Emit this to every error file as it's not specific to a given target.
		for _, e := range state.Err {
			fmt.Fprintf(e, "Could not run command: %v\n", err)
		}
		return subcommands.ExitFailure
	}

	retCode := subcommands.ExitSuccess
	for r := range resp {
		if r.Error != nil {
			fmt.Fprintln(state.Err[r.Index], r.Error)
			// If any target had errors it needs to be reported for that target but we still
			// need to process responses off the channel. Final return code though should
			// indicate something failed.
			retCode = subcommands.ExitFailure
			continue
		}
		fmt.Fprintf(state.Out[r.Index], "%v row(s) affected\n", r.Resp.RowsAffected)
	}
	return retCode
}

// renderItem pretty-prints an item for display
func renderItem(item *pb.Row_Item) string {
	switch val := item.Value.(type) {
	case *pb.Row_Item_BoolValue:
		return fmt.Sprint(val.BoolValue)
	case *pb.Row_Item_NumberValue:
		return fmt.Sprint(val.NumberValue)
	case *pb.Row_Item_ByteValue:
		return fmt.Sprint(val.ByteValue)
	case *pb.Row_Item_IntValue:
		return fmt.Sprint(val.IntValue)
	case *pb.Row_Item_StringValue:
		return fmt.Sprint(val.StringValue)
	case *pb.Row_Item_TimeValue:
		return fmt.Sprint(val.TimeValue)
	case nil:
		return "NULL"
	default:
		return fmt.Sprint(val)
	}
}

type queryCmd struct{}

func (*queryCmd) Name() string     { return "query" }
func (*queryCmd) Synopsis() string { return "Run a sql query and print results." }
func (*queryCmd) Usage() string {
	return `query "SQL QUERY:
	Run a sql query and print results.
`
}

func (p *queryCmd) SetFlags(f *flag.FlagSet) {}

func (p *queryCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	if f.NArg() != 1 {
		fmt.Fprintln(os.Stderr, "Please specify a quoted SQL string.")
		return subcommands.ExitUsageError
	}

	state := args[0].(*util.ExecuteState)
	c := pb.NewPsqlClientProxy(state.Conn)

	resp, err := c.QueryOneMany(ctx, &pb.QueryRequest{Query: f.Args()[0]})
	if err != nil {
		// Emit this to every error file as it's not specific to a given target.
		for _, e := range state.Err {
			fmt.Fprintf(e, "Could not run command: %v\n", err)
		}
		return subcommands.ExitFailure
	}

	retCode := subcommands.ExitSuccess
	for r := range resp {
		if r.Error != nil {
			fmt.Fprintln(state.Err[r.Index], r.Error)
			// If any target had errors it needs to be reported for that target but we still
			// need to process responses off the channel. Final return code though should
			// indicate something failed.
			retCode = subcommands.ExitFailure
			continue
		}

		t := table.NewWriter()
		t.SetOutputMirror(state.Out[r.Index])
		var header table.Row
		for _, c := range r.Resp.Columninfo {
			header = append(header, c.Name)
		}
		t.AppendHeader(header)
		for _, r := range r.Resp.Rows {
			var row table.Row
			for _, i := range r.Item {
				row = append(row, renderItem(i))
			}
			t.AppendRow(row)
		}
		t.Render()
	}
	return retCode
}
