use tonic::transport::{Certificate, Identity, Server, ServerTlsConfig};
use sansshell::services::healthcheck::HealthCheck;

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let cert = std::fs::read_to_string("/home/stvn/.sansshell/leaf.pem")?;
    let key = std::fs::read_to_string("/home/stvn/.sansshell/leaf.key")?;
    let server_identity = Identity::from_pem(cert, key);

    let client_ca_cert = std::fs::read_to_string("/home/stvn/.sansshell/root.pem")?;
    let client_ca_cert = Certificate::from_pem(client_ca_cert);

    let tls = ServerTlsConfig::new()
        .identity(server_identity)
        .client_ca_root(client_ca_cert);

    let addr = "127.0.0.1:50042".parse().unwrap();
    let healthcheck = HealthCheck::default();

    Server::builder()
        .tls_config(tls)?
        .add_service(healthcheck.to_server())
        .serve(addr)
        .await?;
    Ok(())
}
