package main

import (
	"flag"
	"fmt"
	"thai-plate-prophecy-mcp/handler"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	// Initialize the server
	sseMode := flag.Bool("sse", false, "Enable SSE mode")
	flag.Parse()

	// Create a new server instance
	mcpServer := server.NewMCPServer(
		"Thai plate prophecy mcp",
		"1.0.1",
		server.WithResourceCapabilities(true, true),
		server.WithLogging(),
		server.WithRecovery(),
	)

	// v1.AdviceByDMY()
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

	// v1.AdviceByWeekDay()
	// Create a new tool for getting advice by day of the week
	adviceByDayOfWeek := mcp.NewTool(
		"advice_by_day_of_week",
		mcp.WithDescription("Get advice by day of the week"),
		mcp.WithString("day",
			mcp.Required(),
			mcp.Description("day number of birth day"),
		),
	)

	// Create a new tool for getting advice for all days of the week
	adviceAllDays := mcp.NewTool(
		"advice_all_days",
		mcp.WithDescription("Get advice for all days of the week"),
	)

	// v1.AdviceByPlateData()
	// Create a new tool for getting advice by Plate number
	adviceByPlate := mcp.NewTool("advice_by_plate",
		mcp.WithDescription("Get advice by plate number"),
		mcp.WithString("plate",
			mcp.Required(),
			mcp.Description("plate number, example: 1กก 1234"),
		),
	)

	// Add the tool to the server
	mcpServer.AddTool(adviceByDMY, handler.AdviceByDMYHandler)
	mcpServer.AddTool(adviceByDayOfWeek, handler.AdviceByDayOfWeekHandler)
	mcpServer.AddTool(adviceAllDays, handler.AdviceAllDaysHandler)
	mcpServer.AddTool(adviceByPlate, handler.AdviceByPlateHandler)

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
