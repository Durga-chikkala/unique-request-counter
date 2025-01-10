# High-Level Overview

## Introduction

This Go application is a high-performance REST service designed to handle over 10,000 requests per second. It exposes a single endpoint with the following features:

### Endpoint: GET /api/verve/accept

- **Query Parameters:**
    1. **id (int)** - Mandatory: Used for tracking unique requests.
    2. **endpoint (string)** - Optional: Specifies an HTTP endpoint to receive unique request counts in one-minute intervals.

- **Response:**
    - Returns **"ok"** if the request is successfully processed.
    - Returns **"failed"** if any error occurs during processing.

- **Additional Behavior:**
    - Logs the count of unique requests every minute.
    - Sends the unique request count to the specified `endpoint` (if provided).

## Architecture

The application leverages modern design principles and frameworks to ensure scalability, performance, and maintainability.

### Key Components

- **Framework:** Built using the `gorilla/mux` framework for efficient HTTP request and response handling.
- **Scalability:** RESTful design makes the service stateless and easy to scale horizontally.
- **Concurrency:** Utilizes Go’s lightweight goroutines for non-blocking concurrent processing of tasks such as logging and sending HTTP requests.
- **Database:** Redis is used to:
    - Ensure ID deduplication.
    - Store unique request counts.
    - Support distributed system functionality.

## Endpoint Behavior

### GET /api/verve/accept

- **Parameters:**
    - `id` (required): The unique identifier for each request.
    - `endpoint` (optional): Specifies a URL to receive the unique request count.

- **Error Handling:**
    - Responds with HTTP 400 Bad Request if the `id` parameter is missing.

- **Asynchronous Processing:**
    - When the `endpoint` parameter is provided, a goroutine asynchronously sends the unique count to the specified endpoint.

## Design Patterns and Principles

The application follows several design patterns and principles to ensure modularity, testability, and extensibility.

### Layers

1. **Handler Layer:**
    - Extracts data from incoming requests.
    - Handles HTTP-specific concerns such as parameter validation and response formatting.

2. **Service Layer:**
    - Encapsulates business logic.
    - Processes data received from the handler layer.

3. **Store Layer:**
    - Manages interactions with the database (Redis).
    - Handles data persistence and retrieval.

### Design Patterns

1. **Singleton Pattern:**
    - Used for:
        - Initializing Redis connections.
        - Initializing layers (Handler, Service, Store).
        - Initializing writers (logFile, Kafka).
    - Ensures these components are initialized only once during the application lifecycle.

2. **Factory Pattern:**
    - Simplifies the initialization of writers (logFile or Kafka) based on the environment.
    - Dynamically initializes layers based on configuration.

3. **Strategy Pattern:**
    - Provides extensibility for the store and writer layers.
    - New databases or writers can be added by implementing respective interfaces without modifying existing implementations.

## Database

### Redis

Redis is used as a distributed in-memory database with the following purposes:

- **ID Deduplication:** Ensures each `id` is processed only once.
- **Unique Request Count Storage:** Keeps track of the unique request counts for logging and reporting.
- **Locking Mechanism:** Prevents race conditions when calculating unique request counts.

## Summary

This Go application is a scalable, high-performance REST service that processes and tracks unique requests efficiently. It leverages robust design patterns and a well-structured architecture to ensure maintainability and extensibility. Redis’s capabilities further enhance its performance and scalability in distributed environments.
