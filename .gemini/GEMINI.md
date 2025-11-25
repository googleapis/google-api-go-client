# Google API Go Client (`google-api-go-client`) Context

## Repository Overview
This repository contains the auto-generated Go client libraries for Google APIs, based on the **Discovery API**.

*   **Nature of Code:** Almost every top-level directory (e.g., `calendar/`, `drive/`, `sheets/`) corresponds to a Google API and is **auto-generated**.
*   **Editing:** Avoid manual edits to specific API clients unless absolutely necessary for debugging. Changes should ideally happen in the generator or the source discovery doc.

## Key Directories
While most directories are clients, the following are the core libraries that support them:

*   `googleapi/`: Contains common types, error handling, and JSON helpers used across all clients. **High importance** for cross-cutting concerns.
*   `transport/`: Handles HTTP transport, authentication, and connection pooling. **Critical** for networking behavior.
*   `option/`: Configuration options for clients (e.g., API keys, credentials, endpoints).
*   `idtoken/`: Functionality for validating and generating Google ID tokens.
*   `internal/`: Shared internal logic.

## Architecture & Wiring
*   **Transport Creation (`transport/http/dial.go`):** This is the **critical chokepoint** for all HTTP clients. `NewClient` calls `dial`, which constructs the `http.Client`.
    *   **Instrumentation Point:** This is the exact location to wrap the `http.Transport` (e.g., with OpenTelemetry) or inject middleware.
*   **Configuration:** Users configure clients using `option.ClientOption` (found in `option/`).
    *   **`option.WithTracerProvider`:** This option flows into `dial.go` and determines the tracing provider.

## Context
*   **Difference from `google-cloud-go`:** These clients are based on the JSON Discovery format. They are typically pure REST/HTTP clients. The clients in `google-cloud-go` are often gRPC-based (GAPIC) and considered "idiomatic" for Google Cloud Platform services.
