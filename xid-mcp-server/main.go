package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/rs/xid"
)

func main() {
	srv := NewMCPServer()
	if err := server.ServeStdio(srv); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func NewMCPServer() *server.MCPServer {
	srv := server.NewMCPServer(
		"xid",
		"0.0.1",
		server.WithResourceCapabilities(true, true),
		server.WithLogging(),
	)

	srv.AddTool(
		mcp.NewTool(
			"generate_xid",
			mcp.WithDescription("Generate new xid"),
		),
		handlerGenerateXID,
	)
	srv.AddTool(
		mcp.NewTool(
			"parse_xid",
			mcp.WithDescription("Parse xid"),
			mcp.WithString("xid", mcp.Required(), mcp.Description("input xid")),
		),
		handlerParseXID,
	)

	return srv
}

func handlerGenerateXID(_ context.Context, _ mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	newXID := xid.New().String()
	return mcp.NewToolResultText(newXID), nil
}

func handlerParseXID(_ context.Context, reqeust mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	inputXid, err := requiredParam[string](reqeust, "xid")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id, err := xid.FromString(inputXid)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	b, err := json.Marshal(struct {
		Time    string `json:"time"`
		Machine string `json:"machine"`
		PID     uint16 `json:"pid"`
		Counter int32  `json:"counter"`
	}{
		Time:    id.Time().Format("2006-01-02 15:04:05"),
		Machine: fmt.Sprintf("%x", id.Machine()),
		PID:     id.Pid(),
		Counter: id.Counter(),
	})
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return mcp.NewToolResultText(string(b)), nil
}

func requiredParam[T comparable](r mcp.CallToolRequest, p string) (T, error) {
	var zero T

	if _, ok := r.Params.Arguments[p]; !ok {
		return zero, fmt.Errorf("missing required parameter: %s", p)
	}

	if _, ok := r.Params.Arguments[p].(T); !ok {
		return zero, fmt.Errorf("parameter %s is not of type %T", p, zero)
	}

	if r.Params.Arguments[p].(T) == zero {
		return zero, fmt.Errorf("missing required parameter: %s", p)

	}

	return r.Params.Arguments[p].(T), nil
}
