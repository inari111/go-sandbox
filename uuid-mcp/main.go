package main

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	s := server.NewMCPServer(
		"UUID Generator",
		"1.0.0",
		server.WithResourceCapabilities(true, true),
		server.WithLogging(),
		server.WithRecovery(),
	)

	uuidTool := mcp.NewTool("generate-uuid",
		mcp.WithDescription("Generate a UUID"),
	)

	s.AddTool(uuidTool, generateUUID)

	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}

func generateUUID(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	uid := uuid.New()
	return mcp.NewToolResultText(uid.String()), nil
}
