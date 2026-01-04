use serde::{Deserialize, Serialize};

#[derive(Debug, Deserialize, Serialize,Clone,PartialEq)]

pub enum JobStatus {
    CREATED,
    QUEUED,
    PROCESSING,
    COMPLETED,
    FAILED
}

#[derive(Debug, Deserialize, Serialize,Clone)]

pub struct Job{
    pub id: String,
    pub status: JobStatus,
    pub input_file_url: Option<String>,
    pub output_file_url: Option<String>,
}