# Server for ready-to-do

This directory contains the server code for the task management system

## Configuration Options

The following transport configuration options are available:

| Option                                                                | Type     | Description                                      |
| --------------------------------------------------------------------- | -------- | ------------------------------------------------ |
| **transport.grpc.host**                                               | string   | gRPC server binding address (default: "0.0.0.0") |
| **transport.grpc.port**                                               | int      | gRPC server port (default: 50051)                |
| **transport.grpc.max_connection_age**                                 | duration | Maximum connection lifetime (e.g. "30m")         |
| **transport.grpc.max_connection_age_grace**                           | duration | Grace period for connection shutdown             |
| **transport.grpc.max_concurrent_streams**                             | int      | Maximum concurrent RPC streams                   |
| **transport.grpc.max_recv_msg_size**                                  | int      | Maximum incoming message size in bytes           |
| **transport.grpc.max_send_msg_size**                                  | int      | Maximum outgoing message size in bytes           |****
| **transport.grpc.keepalive.enforcement_policy.min_time**              | duration | Minimum allowed client ping interval             |
| **transport.grpc.keepalive.enforcement_policy.permit_without_stream** | bool     | Allow pings without active streams               |
| **transport.grpc.keepalive.server_parameters.max_connection_idle**    | duration | Max idle time before closing connection          |
| **transport.grpc.keepalive.server_parameters.time**                   | duration | Ping frequency when idle                         |
| **transport.grpc.keepalive.server_parameters.timeout**                | duration | Ping acknowledgement timeout                     |
| **transport.grpc.tls.enabled**                                        | bool     | Enable TLS encryption                            |
| **transport.grpc.tls.cert_file**                                      | string   | Path to server certificate file                  |
| **transport.grpc.tls.key_file**                                       | string   | Path to private key file                         |
| **transport.grpc.tls.client_ca_file**                                 | string   | Path to client CA certificate for mTLS           |
| **transport.grpc.health.enabled**                                     | bool     | Enable health check service                      |
| **transport.grpc.health.service_name**                                | string   | Custom health check service name                 |
| **transport.grpc.reflection.enabled**                                 | bool     | Enable gRPC reflection service                   |
| **transport.grpc.logging.level**                                      | string   | Log level (debug/info/warn/error)                |
| **transport.grpc.logging.format**                                     | string   | Log format (json/text)                           |
| **transport.grpc.logging.output**                                     | string   | Log destination (stdout/file path)               |
| **transport.grpc.shutdown.grace_period**                              | duration | Wait period for active RPCs during shutdown      |
| **transport.grpc.shutdown.timeout**                                   | duration | Force shutdown timeout                           |
| **transport.gateway.host**                                            | string   | HTTP gateway binding address                     |
| **transport.gateway.port**                                            | int      | HTTP gateway port (default: 8080)                |
| **transport.gateway.http.read_timeout**                               | duration | Maximum request read duration                    |
| **transport.gateway.http.write_timeout**                              | duration | Maximum response write duration                  |
| **transport.gateway.http.idle_timeout**                               | duration | Keep-alive timeout                               |
| **transport.gateway.http.max_header_bytes**                           | int      | Maximum header size in bytes                     |
| **transport.gateway.cors.allowed_origins**                            | []string | Allowed CORS origins                             |
| **transport.gateway.cors.allowed_methods**                            | []string | Allowed HTTP methods                             |
| **transport.gateway.cors.allowed_headers**                            | []string | Allowed HTTP headers                             |
| **transport.gateway.cors.allow_credentials**                          | bool     | Allow credentials/cookies                        |
| **transport.gateway.cors.max_age**                                    | duration | CORS preflight cache duration                    |
| **transport.gateway.marshaler.emit_defaults**                         | bool     | Include zero-value fields in JSON                |
| **transport.gateway.marshaler.enums_as_ints**                         | bool     | Serialize enums as integers                      |
| **transport.gateway.marshaler.orig_name**                             | bool     | Use original proto field names                   |
| **transport.gateway.tls.enabled**                                     | bool     | Enable HTTPS                                     |
| **transport.gateway.tls.cert_file**                                   | string   | HTTPS certificate file path                      |
| **transport.gateway.tls.key_file**                                    | string   | HTTPS private key file path                      |
| **transport.gateway.shutdown.grace_period**                           | duration | Wait for active HTTP requests                    |
| **transport.gateway.shutdown.timeout**                                | duration | Force shutdown timeout                           |

Below is the list of available database configuration parameters:

| Option                                | Type     | Description                        |
| ------------------------------------- | -------- | ---------------------------------- |
| **database.host**                     | string   | PostgreSQL server host             |
| **database.port**                     | string   | PostgreSQL server port             |
| **database.user**                     | string   | Database username                  |
| **database.password**                 | string   | Database password                  |
| **database.dbname**                   | string   | Database name                      |
| **database.sslmode**                  | string   | SSL mode (verify-full recommended) |
| **database.pool.max_conns**           | int      | Maximum connection pool size       |
| **database.pool.min_conns**           | int      | Minimum idle connections           |
| **database.pool.max_conn_lifetime**   | duration | Maximum connection lifetime        |
| **database.pool.max_conn_idle_time**  | duration | Maximum idle connection time       |
| **database.pool.health_check_period** | duration | Connection health check interval   |
| **database.timeouts.connect**         | duration | Connection timeout                 |
| **database.timeouts.query**           | duration | Query execution timeout            |
| **database.timeouts.exec**            | duration | Write operation timeout            |

See [example config](./config/server.example.yaml) for details.