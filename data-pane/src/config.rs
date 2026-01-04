use std::env;

#[derive(Debug,Clone)]
pub struct Config{
    pub server_url: String,
    pub worker_id: String
}

impl Config{
    pub fn load() -> Self{
        Self {
            server_url: env::var("SERVER_URL")
                .unwrap_or_else(|_| "http://localhost:8080".to_string()),

            worker_id: env::var("WORKER_ID")
                .unwrap_or_else(|_| "worker-1".to_string()),
        }
    }
}
