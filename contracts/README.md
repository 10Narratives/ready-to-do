# Protocol Buffers Contracts

This directory contains Protocol Buffer definitions and generated code for service contracts.

## 🛠️ Prerequisites

- Go 1.19+
- Python 3.8+
- Protocol Buffer compiler (`protoc`)


## 🚀 Quick Start

1. Install all dependencies:
```bash
make prepare
```

2. Generate all contracts (default):
```bash
make generate
```

## 🎛 Custom Generation

The Makefile supports flexible generation options:

### Basic commands:
```bash
# Generate only Go code
make generate LANG=go

# Generate with specific features
make generate FEATURE=grpc-gateway

# Process specific proto packages
make generate PKG="proto/tracker"
```

### Advanced examples:
```bash
# Generate Go + Python with Swagger docs
make generate LANG="go python" FEATURE=swagger

# Generate everything for specific package
make generate PKG=proto/tracker/v1
```

## 📚 Protocol Buffers

### Current Services:
- `tracker.v1.*` - Project tracking service definitions

### Adding New Services:
1. Create new `.proto` file in `proto/{service_name}/v*/`
2. The file will be automatically included in next generation
3. For new major versions, create new `vN` directory

## 🔧 Generation Details

### Output Targets:
| Language | Output Directory | Contents                        |
| -------- | ---------------- | ------------------------------- |
| Go       | `gen/go/`        | gRPC server/client + validation |
| Python   | `gen/python/`    | gRPC client bindings            |

| Features     | Output Directory    | Contents                    |
| ------------ | ------------------- | --------------------------- |
| gRPC Gateway | `gen/grpc-gateway/` | HTTP/JSON translation layer |
| Swagger      | `gen/swagger/`      | OpenAPI documentation       |

### Customizing Generation:
Edit `scripts/generate.sh` to:
- Add new output languages
- Modify protoc flags
- Change default output locations

## 🧩 Integration Examples

### Go Applications:
```go
import (
    "github.com/10Narratives/ready-to-go/contracts/gen/go/proto/tracker/v1"
    
    "google.golang.org/grpc"
)

// Usage example
client := v1.NewProjectServiceClient(grpcConn)
```

### Python Applications:
```python
from proto.tracker.v1 import project_service_pb2
from proto.tracker.v1 import project_service_pb2_grpc

# Usage example
channel = grpc.insecure_channel('localhost:50051')
stub = project_service_pb2_grpc.ProjectServiceStub(channel)
```

## 🔄 Regeneration Workflow

After modifying `.proto` files you should to call `make generate` command in this direcroty


## 🧹 Maintenance

| Command         | Description                       |
| --------------- | --------------------------------- |
| `make prepare`  | Install/update all dependencies   |
| `make generate` | Generate all code artifacts       |
| `make help`     | Show available commands and usage |


## 🤝 Contributing Guidelines

1. **Protocol Buffer changes must follow Google AIP standards**:
   - Adhere to [AIP general guidelines](https://aip.dev/general) for all API design
   - Follow resource-oriented design principles ([AIP-121](https://aip.dev/121)) for REST mappings
   - Implement standard methods ([AIP-131](https://aip.dev/131)) where applicable
   - Use consistent naming as specified in [AIP-140](https://aip.dev/140)

2. Protocol Buffer changes require:
   - Approval from API review board for breaking changes
   - Version bump in package name when making breaking changes ([AIP-180](https://aip.dev/180))
   - Backward compatibility analysis ([AIP-181](https://aip.dev/181))
   - Proper change history documentation ([AIP-192](https://aip.dev/192))

3. Never edit generated files manually - always modify `.proto` files and regenerate

4. Follow naming conventions:
   - Services: `{Domain}Service` (e.g. `ProjectService`)
   - RPC methods: `UpperCamelCase` ([AIP-423](https://aip.dev/423))
   - Messages: `UpperCamelCase`
   - Fields: `lower_snake_case`
   - Enums: `UpperCamelCase` with `UPPER_SNAKE_CASE` values ([AIP-140](https://aip.dev/140))

5. **Additional AIP requirements**:
   - Use standard fields ([AIP-142](https://aip.dev/142)) where applicable
   - Follow error handling guidelines ([AIP-193](https://aip.dev/193))
   - Implement pagination per [AIP-158](https://aip.dev/158) when needed
   - Follow resource state guidelines ([AIP-216](https://aip.dev/216))
