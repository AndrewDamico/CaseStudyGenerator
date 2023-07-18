package main

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
)

func createTable() {
	// Create Table

	// Create a new table
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"", "Industry", "Market", "Company Size", "Business Problem"})

	// Add rows to the table
	for i, study := range studies {
		row := []string{
			fmt.Sprintf("Team %d", i+1),
			study.industry,
			study.market,
			study.companySize,
			study.businessProblem,
		}
		table.Append(row)
	}

	// Render the table
	table.Render()
	fmt.Println()

	return
}
