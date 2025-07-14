# Ready-to-Do Task Management Server

The server component for a gRPC/HTTP task management system with PostgreSQL backend.

## Features

- **Dual Protocol Support**:
  - High-performance gRPC interface (50051 by default)
  - RESTful HTTP/JSON Gateway (8080 by default) with full CORS support
- **Production-Ready Security**:
  - TLS 1.2+ encryption for all endpoints
  - Mutual TLS (mTLS) authentication for gRPC
  - Configurable CORS policies with preflight caching
- **Reliable Database Layer**:
  - Intelligent connection pooling (configurable 2-20 connections)
  - Automatic health checks and connection recycling
  - SSL verification with `verify-full` mode
- **Operational Excellence**:
  - Graceful shutdown with configurable timeouts
  - Keepalive policies to detect half-open connections
  - Structured JSON logging at multiple levels
- **Developer Friendly**:
  - Built-in gRPC reflection service
  - Configurable JSON marshaling options
  - Health check endpoints
