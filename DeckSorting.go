package main

import "math/rand"

type Card struct {
	suit int
	rank int
}

func ShuffleDeck(Deck []Card)[]Card{
	IntDeck := DeckToCleanIntArray(Deck)
	ShuffledIntDeck := ShuffleIntArray(IntDeck)
	ShuffledDeck := CleanIntArrayToDeck(ShuffledIntDeck)
	return ShuffledDeck
}

func ShuffleIntArray(input []int)[]int{

	array := input

	for i := range array{
		j := rand.Intn(i+1)
		array[i], array[j] = array[j], array[i]
	}

	return array
}

func SortDeck(Deck []Card)[]Card{
	//return CleanIntArrayToDeck(SortArray(DeckToCleanIntArray(Deck)))
	IntDeck := DeckToCleanIntArray(Deck)
	SortedIntDeck := SortArray(IntDeck)
	SortedDeck := CleanIntArrayToDeck(SortedIntDeck)
	return SortedDeck
}

func DeckToCleanIntArray(Deck []Card)[]int{
	CleanIntArray := []int{}

	for _, c := range Deck{
		CleanIntArray = append(CleanIntArray, c.suit*100+c.rank)
	}
	return CleanIntArray
}

func CleanIntArrayToDeck(array []int)[]Card{
	Deck := []Card{}

	for _, c := range array{
		if c > 350{
			Deck = append(Deck, Card{suit:4, rank:c-400})
		} else if c > 250{
			Deck = append(Deck, Card{suit:3, rank:c-300})
		}else if c > 150{
			Deck = append(Deck, Card{suit:2, rank:c-200})
		}else {
			Deck = append(Deck, Card{suit:1, rank:c-100})
		}

	}

	return Deck


}


//Insertion Short btw.
func SortArray(input []int)[]int{
	array := input
	for i := 1; i < len(array); i++ {
		for j := i; j > 0 && array[j] < array[j - 1]; j-- {
			array[j], array[j-1] = array[j-1], array[j]
		}
	}

	return array
}

func CreateNewDeck()[]Card{
	results := []Card{}
	/// For each suit
	for s := 1; s<5; s++{
	/// For each rank
		for r := 1; r<14; r++{
			results = append(results, Card{s,r})
		}
	}
	return results
}


func getSuitFromInt(s int)string{
	switch s {
	case 1:
		return "Spades"
	case 2:
		return "Hearts"
	case 3:
		return "Clubs"
	case 4:
		return "Diamonds"
	default:
		return "Unknown"
	}
}

func getRankFromInt(r int)string{
	switch r {
	case 1:
		return "Ace"
	case 2:
		return "Two"
	case 3:
		return "Three"
	case 4:
		return "Four"
	case 5:
		return "Five"
	case 6:
		return "Six"
	case 7:
		return "Seven"
	case 8:
		return "Eight"
	case 9:
		return "Nine"
	case 10:
		return "Ten"
	case 11:
		return "Jack"
	case 12:
		return "Queen"
	case 13:
		return "King"
	default:
		return "Unknown"
	}
}