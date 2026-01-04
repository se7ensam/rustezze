use std::fmt;

#[derive(Debug)]
pub enum WorkerError {
    // 1. Transient Errors (Network blips, Server down)
    // The worker should probably just sleep and try again.
    NetworkError(String),
    
    // 2. Fatal Errors (Bad JSON, Invalid Data)
    // The worker must report this as a specific FAILED status.
    ParseError(String),
    
    // 3. Logic Errors (The image processing failed)
    ProcessingError(String),
}

// Allow us to print the error nicely (User-friendly string)
impl fmt::Display for WorkerError {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        match self {
            WorkerError::NetworkError(msg) => write!(f, "Network Error: {}", msg),
            WorkerError::ParseError(msg) => write!(f, "Data Parse Error: {}", msg),
            WorkerError::ProcessingError(msg) => write!(f, "Processing Error: {}", msg),
        }
    }
}

// Allow it to be used as a standard Rust error (System-compatible)
impl std::error::Error for WorkerError {}