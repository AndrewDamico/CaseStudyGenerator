package scenariogenerator

import (
	"bufio"
	"fmt"
	"os"

	"github.com/inancgumus/screen"
)

func clearScreen() {
	screen.Clear()
	screen.MoveTopLeft()
}
func parameterMenu() {
	clearScreen()
	var choice int64 = -1

	scanner := bufio.NewScanner(os.Stdin)

	menuReset := func() {
		fmt.Println()
		fmt.Println("Press Enter to continue...")
		scanner.Scan()
		clearScreen()
	}

	for {
		fmt.Println("Parameter Selection")
		fmt.Println("Please choose one of the following options:")
		fmt.Println("1. Add Industry:", industries)
		fmt.Println("2. Add Market:", market)
		fmt.Println("3. Add Company Size", companySize)
		fmt.Println("4. Add Business Problem", businessProblem)
		fmt.Println("0. Return to previous Menu")

		var err error

		_, err = fmt.Scanf("%d\n", &choice)
		if err != nil {
			choice = -1
		}

		switch choice {
		case 1:
			clearScreen()
			industries = addVariable(industries, "industry")
			fmt.Println("Industries:", industries)
			menuReset()
		case 2:
			clearScreen()
			market = addVariable(market, "Market")
			fmt.Println("Markets:", market)
			menuReset()
		case 3:
			clearScreen()
			companySize = addVariable(companySize, "Company Size")
			fmt.Println("Company Sizes:", companySize)
			menuReset()
		case 4:
			clearScreen()
			businessProblem = addVariable(businessProblem, "Business Problem")
			fmt.Println("Business Problem:", businessProblem)
			menuReset()
		default:
			fmt.Println("Invalid choice! Please try again.")
		}

		if choice == 0 {
			clearScreen()
			return
		}
	}
}

func mainMenu() {
	clearScreen()
	var choice int64 = -1

	scanner := bufio.NewScanner(os.Stdin)

	menuReset := func() {
		fmt.Println()
		fmt.Println("Press Enter to continue...")
		scanner.Scan()
		clearScreen()
	}

	for {
		fmt.Println("Case Study Generator")
		fmt.Println("Andrew D'Amico")
		fmt.Println("")
		fmt.Println("Main Menu")
		fmt.Println("Please choose one of the following options:")
		fmt.Println("1. Select Parameters")
		fmt.Println("2. Generate Case Study Template")
		fmt.Println("3. View Teams (list)")
		fmt.Println("4. View Teams (table)")
		//fmt.Println("9. Configuration Menu")
		fmt.Println("0. Exit")

		var err error

		_, err = fmt.Scanf("%d\n", &choice)
		if err != nil {
			choice = -1
		}

		switch choice {
		case 0:
			// Exit the program
			fmt.Println("Goodbye...")
		case 1:
			clearScreen()
			parameterMenu()
		case 2:
			clearScreen()
			Generate()
			menuReset()
		case 3:
			clearScreen()
			viewTeams()
			menuReset()
		case 4:
			clearScreen()
			createTable()
			menuReset()
		case 9:
			clearScreen()
			optionsMenu()
		default:
			fmt.Println("Invalid choice! Please try again.")
		}

		if choice == 0 {
			clearScreen()
			return
		}
	}
}

func optionsMenu() {
	clearScreen()
	fmt.Println("Configuration Menu")
	fmt.Println()

	// determine which set to test on

	var choice int64 = -1

	scanner := bufio.NewScanner(os.Stdin)

	menuReset := func() {
		fmt.Println()
		fmt.Println("Press Enter to continue...")
		scanner.Scan()
		clearScreen()
	}

	for {
		fmt.Println("Configuration Menu")
		fmt.Println("Please choose one of the following options:")
		fmt.Println("1. Set number of Teams")
		fmt.Println("Press Enter to Return to Previous Menu")

		var err error

		_, err = fmt.Scanf("%d\n", &choice)
		if err != nil {
			choice = -1
		}

		switch choice {
		case 0:
			clearScreen()
			// Return to Previous Menu
			fmt.Println()
		case 1:
			clearScreen()
			fmt.Println("Runs per test determines number of times each model will be built. Times are then averaged and reported to application.")
			fmt.Println("Current number of teams:", nTeams)
			fmt.Println("Please enter new number and press enter.")
			//#todo if user doesn't enter number use default
			_, err := fmt.Scanf("%d\n", &nTeams)
			if err != nil {
				fmt.Println("Invalid input. Please try again.")
				continue
			}
			fmt.Println("Current number of Teams:", nTeams)
			menuReset()

		default:
			clearScreen()
			return
		}

		if choice == 0 {
			clearScreen()
			return
		}
		//fmt.Println("Press Enter to continue...")
		//scanner.Scan() // Wait for user to press Enter

	}
}
