use tonic::{Request, Response, Status};
use health_check::health_check_server::HealthCheckServer;

pub mod health_check {
    tonic::include_proto!("health_check");
}

#[derive(Default)]
pub struct HealthCheck {}

#[tonic::async_trait]
impl health_check::health_check_server::HealthCheck for HealthCheck {
    async fn ok(
        &self,
        _request: Request<()>,
    ) -> Result<Response<()>, Status> {
        Ok(Response::new(()))
    }
}

impl HealthCheck {
    pub fn to_server(self) -> HealthCheckServer<HealthCheck> {
        HealthCheckServer::new(self)
    }
}
