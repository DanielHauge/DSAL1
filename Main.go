package main

import (
	"math/rand"

	"fmt"
)



func main() {

	/*
	/////// BirthdayParty problem 1.4.44
	fmt.Println("Expected from formular: ")
	fmt.Println(DuplicateExpected(100000))
	fmt.Println("A milion test runs give: ")
	fmt.Println(RunTestDuplicate(100000, 100000))
	*/
	///////// Coupons problem 1.4.45
	fmt.Print("Result: ")
	fmt.Print(RunTestCoupons(25000, 100))
	fmt.Print(" - Expected: ")
	fmt.Print(ExpectedCoupons(25000))



	//////// Deck Sorting problem
	
	/*
	deck := CreateNewDeck()
	CleanArray := DeckToCleanIntArray(deck)
	for _, c := range CleanArray{
		log.Println(c)
	}
	NewDeck := CleanIntArrayToDeck(CleanArray)
	for _, r := range NewDeck{
		log.Println(getRankFromInt(r.rank)+" of "+getSuitFromInt(r.suit))
	}
	shuffled := ShuffleDeck(deck)
	for _, r := range shuffled{
		log.Println(getRankFromInt(r.rank)+" of "+getSuitFromInt(r.suit))
	}
	log.Println("It got shuffled -> Now to sorting \n\n")

	SortedDeck := SortDeck(shuffled)
	log.Println(len(SortedDeck))
	for _, r := range SortedDeck{
		log.Println(getRankFromInt(r.rank)+" of "+getSuitFromInt(r.suit))
	}

	*/








}







func random(min int, max int) int {
	return rand.Intn(max-min) + min
}
