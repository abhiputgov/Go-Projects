package main

import (
	"fmt"

	"example.com/b/utils"
)

const accountBalanceFile = "balanceText.txt"

func main() {
	accountBalance, error := utils.GetFloatFromFile(accountBalanceFile, 1000.00)

	if error != nil {
		fmt.Println(error)
		panic("Can't continue. Sorry!")
	}
	for {
		presentOptions()

		var choice int
		fmt.Printf("Enter your choice:\n")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Printf("The current balance is: %.2f\n", accountBalance)
		case 2:
			var deposit float64
			fmt.Printf("How much is the amount to be deposited?")
			fmt.Scan(&deposit)
			if deposit < 0 {
				fmt.Println("You kidding me bro? What did you expect? That you will get free money?")
				continue
			}
			accountBalance += deposit
			fmt.Println("Updating the account balance. Updated account balance is:", accountBalance)
			utils.WriteFloatToFile(accountBalanceFile, accountBalance)
		case 3:
			var withdraw float64
			fmt.Printf("How much are you withdrawing?")
			fmt.Scan(&withdraw)
			if withdraw < 0 {
				fmt.Println("why? f off from here")
				continue
			}
			if withdraw > accountBalance {
				fmt.Println("Whats with you bro? You get only what you have.")
				continue
			}
			accountBalance -= withdraw
			fmt.Println("Updating the account balance. Updated account balance is:", accountBalance)
			utils.WriteFloatToFile(accountBalanceFile, accountBalance)
		default:
			fmt.Println("......Exiting......")
			fmt.Print("Thanks for checking into my bank!")
			return
		}
	}
}
