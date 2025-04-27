# Thai Plate Prophecy MCP

MCP Server for Thai plate prophecy.

## Getting Started

### Prerequisites

- Go 1.16 or higher
- Make (optional, for using the Makefile)
- Git (for cloning the repository)
- Visual Studio Code with Go extension (recommended IDE)
- A compatible MCP client application

### Installation

#### Option 1: Go Install (recommended)

```shell
$ go install github.com/naruebaet/thai-plate-prophecy-go-mcp@latest
```

#### Option 2: Clone and Build

```shell
$ git clone https://github.com/naruebaet/thai-plate-prophecy-go-mcp.git
$ cd thai-plate-prophecy-go-mcp
$ go build
```

#### Option 3: Using Make

```shell
$ git clone https://github.com/naruebaet/thai-plate-prophecy-go-mcp.git
$ cd thai-plate-prophecy-go-mcp
$ make build
```

### Configuration

#### mcp.json config

Create an `mcp.json` file in your project directory with the following content:

```json
{
    "mcpServers": {
        "thai-plate-prophecy-mcp": {
            "type": "stdio",
            "command": "thai-plate-prophecy-mcp",
            "args": []
        }
    }
}
```

#### Environment Variables

The following environment variables can be set to configure the MCP server:

- `PORT`: The port number for the server (default: 8080)
- `DEBUG`: Enable debug logging (set to "true")

### Usage

1. Start the MCP server:
```shell
$ thai-plate-prophecy-mcp
```

2. Connect your client application to the MCP server.

3. Input a Thai license plate number to receive prophecy predictions.

### Development

#### Running Tests

```shell
$ go test ./...
```

#### Building for Different Platforms

```shell
$ GOOS=darwin GOARCH=amd64 go build -o thai-plate-prophecy-mcp-macos
$ GOOS=linux GOARCH=amd64 go build -o thai-plate-prophecy-mcp-linux
$ GOOS=windows GOARCH=amd64 go build -o thai-plate-prophecy-mcp.exe
```

### Example

![Thai Plate Prophecy MCP Example](./Screenshot%202568-04-27%20at%2002.40.58.png)
then
![Thai Plate Prophecy MCP Example](./Screenshot%202568-04-27%20at%2002.41.57.png)

### Support me coffee
![Support me](./bmc_qr.png)