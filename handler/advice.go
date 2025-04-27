package handler

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	tpp "github.com/naruebaet/thai-plate-prophecy-go/v1"
)

func AdviceByDMYHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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

func AdviceByDayOfWeekHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// Implement the logic to get advice by day of the week
	dayStr := request.Params.Arguments["day"].(string)

	// Convert string to WeekDay type (integer)
	var day tpp.WeekDay

	switch strings.ToLower(dayStr) {
	case "sunday":
		day = 0
	case "monday":
		day = 1
	case "tuesday":
		day = 2
	case "wednesday":
		day = 3
	case "thursday":
		day = 4
	case "friday":
		day = 5
	case "saturday":
		day = 6
	default:
		return nil, errors.New("invalid day of week: please use Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, or Saturday")
	}

	result, err := tpp.AdviceByWeekDay(day)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	advice := `
		Your advice based on day of the week: ` + dayStr + `
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

func AdviceAllDaysHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// Create a map to store advice for each day
	allAdvice := make(map[string]map[string]interface{})

	// List of all days
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	// Get advice for each day
	for i, dayStr := range days {
		day := tpp.WeekDay(i) // Cast index to WeekDay type

		result, err := tpp.AdviceByWeekDay(day)
		if err != nil {
			return nil, errors.New("error fetching advice for " + dayStr + ": " + err.Error())
		}

		// Store results in a structured format
		dayAdvice := map[string]interface{}{
			"luckyNumber":        result.LuckyNum,
			"luckyNumberDesc":    result.LuckyNumDesc,
			"avoidNumber":        result.AvoidNum,
			"avoidNumberDesc":    result.AvoidNumDesc,
			"avoidCharacter":     result.AvoidChar,
			"avoidCharacterDesc": result.AvoidCharDesc,
		}

		allAdvice[dayStr] = dayAdvice
	}

	// Format output as text
	var outputBuilder strings.Builder
	outputBuilder.WriteString("Advice for all days of the week:\n")
	outputBuilder.WriteString("===========================\n\n")

	for _, dayStr := range days {
		advice := allAdvice[dayStr]

		outputBuilder.WriteString(fmt.Sprintf("## %s\n", dayStr))
		outputBuilder.WriteString(fmt.Sprintf("Lucky numbers: %v\n", advice["luckyNumber"]))
		outputBuilder.WriteString(fmt.Sprintf("Description: %s\n", advice["luckyNumberDesc"]))
		outputBuilder.WriteString(fmt.Sprintf("Numbers to avoid: %v\n", advice["avoidNumber"]))
		outputBuilder.WriteString(fmt.Sprintf("Description: %s\n", advice["avoidNumberDesc"]))
		outputBuilder.WriteString(fmt.Sprintf("Characters to avoid: %v\n", advice["avoidCharacter"]))
		outputBuilder.WriteString(fmt.Sprintf("Description: %s\n\n", advice["avoidCharacterDesc"]))
	}

	return mcp.NewToolResultText(outputBuilder.String()), nil
}

func AdviceByPlateHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// Implement the logic to get advice by plate number
	plate := request.Params.Arguments["plate"].(string)

	// Split the plate data to array explode by space
	plateData := strings.Split(plate, " ")

	result, err := tpp.AdviceByPlateData(plateData[0], plateData[1])
	if err != nil {
		return nil, errors.New(err.Error())
	}

	advice := `
		Your advice based on plate number: %s
		Your first part : %+v
		Your second part : %+v
		---
		Total : %+v
	`

	return mcp.NewToolResultText(fmt.Sprintf(
		advice,
		plate,
		result.FirstPart,
		result.SecondPart,
		result.Total,
	)), nil
}
