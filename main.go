package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type CaseStudy struct {
	industry        string
	market          string
	companySize     string
	businessProblem string
}

var industries []string
var market []string
var companySize []string
var businessProblem []string
var teams []CaseStudy
var nTeams uint = 2
var studies []CaseStudy

func generateRandomNumber(min, max int) int {
	rand.Seed(time.Now().UnixNano())  // Initialize random seed based on the current time
	return rand.Intn(max-min+1) + min // Generate a random number within the specified range
}

func addVariable(option []string, name string) []string {
	clearScreen()
	var container = option

	reader := bufio.NewReader(os.Stdin)

	for {
		clearScreen()
		fmt.Println("Please enter options and press enter. Type q to exit.")
		fmt.Println("")
		fmt.Println("Editing:", name)
		fmt.Println("Current Options:", container)
		fmt.Println("")
		fmt.Printf("New %v: \n", name)

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "q" {
			break
		}

		container = append(container, input)
	}

	option = append(option, container...)

	return option
}

func randomAssign(variable *[]string, rem bool) string {

	// Generate random items from each set
	length := len(*variable)
	item := generateRandomNumber(0, length-1)

	// Store the value to return
	value := (*variable)[item]
	if rem {
		// Remove the value from the slice
		(*variable)[item] = (*variable)[length-1]
		*variable = (*variable)[:length-1]
	}

	return value
}

func Generate() {

	//randomNumber := generateRandomNumber(min, max)
	team := CaseStudy{}

	// Populate the case study variables

	if len(industries) == 0 || len(businessProblem) == 0 || len(market) == 0 || len(companySize) == 0 {
		fmt.Println("Please add items first")
	} else {
		team.industry = randomAssign(&industries, true)
		team.businessProblem = randomAssign(&businessProblem, true)
		team.market = randomAssign(&market, false)
		team.companySize = randomAssign(&companySize, false)

		println("Generating Case Study Parameters")
		println("Industry:", team.industry)
		println("Market:", team.market)
		println("Company Size:", team.companySize)
		println("Business Problem:", team.businessProblem)
		studies = append(studies, team)
	}

	return
}

func viewTeams() {
	i := 1
	for _, study := range studies {
		fmt.Println("Teams")
		fmt.Println("")
		fmt.Println("Team:", i)
		fmt.Println(" Industry:", study.industry)
		fmt.Println(" Market:", study.market)
		fmt.Println(" Company Size:", study.companySize)
		fmt.Println(" Business Problem:", study.businessProblem)
		i++
		fmt.Println()
	}
}

func main() {
	fmt.Println("Case Study Generator")
	fmt.Println("Andrew D'Amico")
	fmt.Println()
	mainMenu()
}
