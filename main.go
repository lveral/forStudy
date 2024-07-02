package main

import "fmt"

func main() {
	fmt.Print("Test Passed: (2, 89, 102) -> expected: 26 actual: ")
	fmt.Print(monthlyBillIncrease(2, 89, 102))
	fmt.Print("\nTest Passed: (2, 98, 104) -> expected: 12 actual: ")
	fmt.Print(monthlyBillIncrease(2, 98, 104))
	fmt.Print("\nTest Passed: (3, 50, 40) -> expected: -30 actual: ")
	fmt.Print(monthlyBillIncrease(3, 50, 40))
	fmt.Print("\nTest Passed: (3, 60, 60) -> expected: 0 actual: ")
	fmt.Print(monthlyBillIncrease(3, 60, 60))
}

func monthlyBillIncrease(costPerSend, numLastMonth, numThisMonth int) int {
	var lastMonthBill int
	var thisMonthBill int
	lastMonthBill = getBillForMonth(costPerSend, numLastMonth)
	thisMonthBill = getBillForMonth(costPerSend, numThisMonth)
	return thisMonthBill - lastMonthBill
}

func getBillForMonth(costPerSend, messagesSent int) int {
	return costPerSend * messagesSent
}
