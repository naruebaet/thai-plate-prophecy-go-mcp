package main

import (
	"context"
	"errors"
	"flag"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	tpp "github.com/naruebaet/thai-plate-prophecy-go/v1"
)

func main() {
	// Initialize the server
	sseMode := flag.Bool("sse", false, "Enable SSE mode")
	flag.Parse()

	// Create a new server instance
	mcpServer := server.NewMCPServer(
		"Thai plate prophecy mcp",
		"1.0.0",
		server.WithResourceCapabilities(true, true),
		server.WithLogging(),
		server.WithRecovery(),
	)

	// Create a new tool for getting advice by date of birth
	adviceByDMY := mcp.NewTool(
		"advice_by_dmy",
		mcp.WithDescription("Get advice by date of birth"),
		mcp.WithString("date",
			mcp.Required(),
			mcp.Description("date number of birth date"),
		),
		mcp.WithString("month",
			mcp.Required(),
			mcp.Description("month number of birth month"),
		),
		mcp.WithString("year",
			mcp.Required(),
			mcp.Description("year number of birth year"),
		),
	)

	// Add the tool to the server
	mcpServer.AddTool(adviceByDMY, adviceByDMYHandler)

	// Start the server
	if *sseMode {
		sseServe := server.NewSSEServer(mcpServer)
		fmt.Println("Server started in SSE mode... : http://localhost:8080")
		if err := sseServe.Start(":8080"); err != nil {
			fmt.Printf("Server error: %v\n", err)
		}
	} else {
		fmt.Println("Server started in Stdio mode...")
		if err := server.ServeStdio(mcpServer); err != nil {
			fmt.Printf("Server error: %v\n", err)
		}
	}

}

func adviceByDMYHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// Implement the logic to get advice by date of birth
	// This is a placeholder implementation
	date := request.Params.Arguments["date"].(string)
	month := request.Params.Arguments["month"].(string)
	year := request.Params.Arguments["year"].(string)

	result, err := tpp.AdviceByDMY(date, month, year)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	advice := `
		Your advice based on date of birth: " + date + "/" + month + "/" + year
		---
		Here is your lucky number: %v
		Description: %s
		Here is your avoid number: %v
		Description: %s
		Here is your avoid character: %v
		Description: %s
	`

	return mcp.NewToolResultText(fmt.Sprintf(advice, result.LuckyNum, result.LuckyNumDesc, result.AvoidNum, result.AvoidNumDesc, result.AvoidChar, result.AvoidCharDesc)), nil
}
