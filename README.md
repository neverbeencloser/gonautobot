# gonautobot

A Go client library for the [Nautobot](https://github.com/nautobot/nautobot) REST API. This library provides a clean Go interface to interact with Nautobot's various endpoints, supporting all major CRUD operations with built-in pagination and error handling.

## Features

- 🔄 **Automatic Pagination**: Built-in pagination handling for large result sets
- 🛡️ **Type Safety**: Leverages Go generics for type-safe operations
- 🔧 **Flexible Configuration**: Environment variables or programmatic configuration
- 📊 **Comprehensive Testing**: Extensive test coverage with mock HTTP responses
- 🎯 **Clean Architecture**: Modular design mirroring Nautobot's API structure

## Installation

```bash
go get github.com/neverbeencloser/gonautobot
```

## Quick Start

### Basic Client Setup

```go
package main

import (
    "fmt"
    "log"

    nautobot "github.com/neverbeencloser/gonautobot"
)

func main() {
    // Create client with environment variables
    client := nautobot.New()

    // Or configure explicitly
    client := nautobot.New(
        nautobot.WithEndpoint("https://your-nautobot.example.com"),
        nautobot.WithToken("your-api-token-here"),
    )

    // Use the client
    devices, err := client.Dcim.DeviceAll()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Found %d devices\n", len(devices))
}
```

### Environment Variables

Set these environment variables for automatic configuration:

```bash
export NAUTOBOT_URL="https://your-nautobot.example.com"
export NAUTOBOT_TOKEN="your-api-token-here"
export DEBUG="true"  # Optional: enable debug logging
```

## API Structure

The client is organized into modules that mirror Nautobot's API structure:

| Module | Description |
|--------|-------------|
| `Circuits` | Circuit and provider management |
| `Dcim` | Data center infrastructure |
| `Ipam` | IP address management |
| `Extras` | Custom fields, tags, status |
| `Tenancy` | Multi-tenancy support |
| `Virtualization` | Virtual machines and clusters |
| `Plugins` | Plugin-specific endpoints |

* The `GraphQL` endpoint is currently housed under the `Extras` module.

## Common Operations

### Creating Resources

```go
// Create a new device
device, err := client.Dcim.DeviceCreate(dcim.NewDevice{
    Name:       "server01",
    Role:       "server",
    Status:     "Active",
    DeviceType: "device-type-uuid",
    Location:   "location-uuid",
})
```

### Reading Resources

```go
// Get a specific location by ID
device, err := client.Dcim.LocationGet(uuid.MustParse(locationID))

// List all devices
devices, err := client.Dcim.DeviceAll()

// Filter devices
filters := &url.Values{
    "status":   {"Active"},
}
devices, err := client.Dcim.DeviceFilter(filters)
```

### Updating Resources

```go
// Update specific attributes using the API's PATCH methods
updatedLocation, err := client.Dcim.LocationUpdate(uuid.MustParse(locationID), map[string]any{
    "name": "location01-updated",
    "comments": "Updated via API",
})
```

### Deleting Resources

```go
// Delete a device
err := client.Dcim.DeviceDelete(deviceID)
```

## Advanced Features

### Pagination

The client automatically handles paginated responses:

```go
// This will fetch ALL devices across multiple pages
allDevices, err := client.Dcim.DeviceAll()
```

### Multipart File Uploads

Specific endpoints support multipart file uploads for Create/Update:
 * `Dcim.DeviceType`

```go
import (
	nautobot "github.com/neverbeencloser/tobot"
	"github.com/neverbeencloser/tobot/dcim"
)
// Create a new device type
payload := dcim.NewDeviceType{
	Model: "server",
	Manufacturer: "Dell",
	UHeight: 2,
}
payload.SetFrontImage("./image.png")

dt, err := c.Dcim.DeviceTypeCreate(payload)
if err != nil {
	log.Fatal(err)
}
```

##### :warning: Image Update Behavior
 * Pass the empty string to delete an existing image
```go
dt, err := c.Dcim.DeviceTypeGet(uuid.MustParse("device-type-uuid"))
if err != nil {
	log.Fatal(err)
}

payload := dcim.NewDeviceType{
	Model: "diff-model",
	Manufacturer: dt.Manufacturer,
}
payload.SetFrontImage("") // delete existing image

updatedDT, err := c.Dcim.DeviceTypeUpdate(dt.ID, payload)
if err != nil {
	log.Fatal(err)
}
```

### Custom HTTP Client

```go
import "time"

httpClient := &http.Client{
    Timeout: 30 * time.Second,
}

client := nautobot.New(
    nautobot.WithHTTPClient(httpClient),
    nautobot.WithToken("your-token"),
)
```

## Development

### Docker Development Environment

The project includes a containerized development Nautobot stack using Docker Compose. This ensures consistent tooling and testing across different development machines.

#### Prerequisites

```bash
# Copy the example environment file
cp develop/example.env develop/.env
```

Edit `develop/.env` to customize Go version, Alpine version, and tool versions as needed.

#### Available Make Commands

##### Controlling the Development Environment

```bash
# Build the development environment
make build

# Start the development environment
make start

# Stop the development environment
make stop

# Restart the development environment
make restart
```

##### Additional Targets
```bash
# Launch interactive shell on nautobot container
make cli

# Launch shell_plus instance on nautobot container
make shell

# Run all tests (linting + unit tests)
make tests

# Run unit tests only
make unittest

# Run linting only
make lint

# Show available commands
make help
```

## Contributing

We welcome contributions! Whether it's bug fixes, new features, or documentation improvements, we'd love to see your pull requests.

1. Fork the repository
2. Create a feature branch
3. Add tests for new functionality
4. Ensure all tests pass
5. Submit a pull request

## API Compatibility

This client is designed to work with Nautobot 2.x. While we strive for backward compatibility, some features may require specific Nautobot versions.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Related Projects

- [Nautobot](https://github.com/nautobot/nautobot) - The source of truth for network automation
- [PyNautobot](https://github.com/nautobot/pynautobot) - Python client for Nautobot

## Support

- 🐛 Issues: [GitHub Issues](https://github.com/neverbeencloser/gonautobot/issues)
- 💬 Discussions: [GitHub Discussions](https://github.com/neverbeencloser/gonautobot/discussions)

## To-Do
- **Full API Coverage**: Support all major Nautobot API endpoints (Circuits, DCIM, IPAM, Extras, Tenancy, Virtualization, Plugins)
- **Documentation**: Create and improve examples for individual resources.
