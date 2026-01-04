mod config;
mod error;
mod job;
mod worker;

use config::Config;
use worker::Worker;

#[tokio::main]
async fn main() {

    let config = Config::load();
    
    println!("Rust Worker Started!");
    println!("ID: {}", config.worker_id);
    println!("Target: {}", config.server_url);

    let worker = Worker::new(config);
    worker.run().await;
}