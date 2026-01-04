package main

import (
	"fmt"
	
	// Import your internal packages
	"rusteze/control-pane/internal/jobs"
	
	// ⚠️ IMPORTANT: Alias this import because "http" conflicts with net/http
	transport "rusteze/control-pane/internal/http"
)

func main() {
	fmt.Println("Initializing PixelGrid Control Plane...")

	// 1. Initialize Layers (Dependency Injection)
	store := jobs.NewMemoryStore()       // Data Layer
	service := jobs.NewService(store)    // Business Logic Layer
	handler := transport.NewJobHandler(service) // HTTP Layer

	// 2. Setup Router
	router := transport.NewRouter(handler)

	// 3. Start Server
	fmt.Println("Server running on port 8080")
	router.Run(":8080")
}