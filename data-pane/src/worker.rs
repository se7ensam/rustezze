use std::time::Duration;
use tokio::time::sleep;
use crate::config::Config;
use crate::error::WorkerError;
use crate::job::Job;

pub struct Worker {
    client: reqwest::Client,
    config: Config,
}

impl Worker {
    pub fn new(config: Config) -> Self {
        Self {
            client: reqwest::Client::new(),
            config,
        }
    }

    
    pub async fn run(&self) {
        println!("⚙️  Worker loop started. Polling {}...", self.config.server_url);

        loop {
            match self.fetch_job().await {
                Ok(Some(job)) => {
                    println!("📦 Received Job: {}", job.id);
                    // TODO: Process the job (Phase 3)
                    // TODO: Report success/failure (Phase 3)
                }
                Ok(None) => {
                    // No work available. Sleep to save CPU.
                    println!("No work. Sleeping...");
                    sleep(Duration::from_secs(2)).await;
                }
                Err(e) => {
                    eprintln!("Error fetching job: {}", e);
                    sleep(Duration::from_secs(5)).await;
                }
            }
        }
    }

    // The "Ask" Logic
    async fn fetch_job(&self) -> Result<Option<Job>, WorkerError> {
        // 1. Construct the URL: http://localhost:8080/jobs/poll
        let url = format!("{}/jobs/poll", self.config.server_url);

        // 2. Send the POST request
        // We use the 'client' we stored earlier.
        let response = self.client
            .post(&url)
            .send()
            .await
            .map_err(|e| WorkerError::NetworkError(e.to_string()))?; // Map reqwest error to OUR error

        // 3. Handle the Status Code
        match response.status() {
            reqwest::StatusCode::OK => {
                // 200 OK = We got a job! Parse the JSON.
                let job = response
                    .json::<Job>()
                    .await
                    .map_err(|e| WorkerError::ParseError(e.to_string()))?;
                Ok(Some(job))
            }
            reqwest::StatusCode::NO_CONTENT => {
                // 204 No Content = No work available.
                Ok(None)
            }
            _ => {
                // 500, 404, etc. = Something went wrong.
                Err(WorkerError::NetworkError(format!("Server error: {}", response.status())))
            }
        }
    }
}