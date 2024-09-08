fn main() {
    tonic_build::compile_protos("services/healthcheck/healthcheck.proto").unwrap();
}
