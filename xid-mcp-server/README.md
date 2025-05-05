# xid-mcp-server

The xid MCP server is a [Model Context Protocol (MCP)](https://modelcontextprotocol.io/introduction) server that generates and parses [xid](https://github.com/rs/xid).

## Installation

### Usage with Visual Studio Code

```json
{
  "mcp": {
    "servers": {
      "xid-mcp-server": {
        "command": "go",
        "args": [
          "run",
          "github.com/candy12t/mcp/xid-mcp-server@latest"
        ]
      }
    }
  }
}
```

### Usage with Claude Desktop

```json
{
  "mcpServers": {
    "xid-mcp-server": {
      "command": "go",
      "args": [
        "run",
        "github.com/candy12t/mcp/xid-mcp-server@latest"
      ]
    }
  }
}
```

### Build from soruce

```bash
go install github.com/candy12t/mcp/xid-mcp-server@latest
```

```json
{
  "mcpServers": {
    "xid-mcp-server": {
      "command": "/path/to/xid-mcp-server"
    }
  }
}
```

## Tools

- `generate_xid`: generate new xid
- `parse_xid`: parse xid
