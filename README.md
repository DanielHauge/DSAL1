# 1. Compulsory algorithm Assignment - Experiments and Sorting.

Daniel F. Hauge (cph-dh136)

This is the hand-in paper for Datastructure & Algorithms course. It consists of 3 bullets points (Subjects) which needs attention. This paper will include results, comments to results and the essential code.

I've chossen to make all experiments and programs in the language "Golang".

## Birthday Experiencement
#### Description 
Write a program that takes an integer N from the command line and generate a random sequence of integers between 0 and N – 1. Run experiments to validate the hypothesis that the number of integers generated before the first repeated value is found is ~√(pi)N/2.

#### Results & comments:

My results to this problem is as follows:
```
Expected from formular: 
23.944532972687885
A milion test runs give: 
24.615946
```
This seems very close, the diffence is: ~0.671 and might therefor be true. However i wanted to try see what would happen if you limited towards infinity to see if the expected formular would get closer to actual tests. So new results show from if N was 1000 and not 365.
```
Expected from formular: 
39.633272976060105
A milion test runs give: 
40.27123
```
The difference is ~0.638. This looks like it is getting closer to 0. I wanted to try one last time but with a huge value instead. New tests are conducted with 100000.
```
Expected from formular: 
396.33272976060107
A milion test runs give: 
396.98841
```
The difference is ~0.656. Which gets higher than previus. But it stays about the same offset for any value put in to N. This indicates it is very close to the actual formular.

The big O notation is ofcause the formular. O(√(pi)N/2) But worst case scenario, it could also be O(N). It could potentially, go through the whole available valuespace before it finds a duplicate. Which is N. 

#### Essential Code:
I've calculated this by the following function.

```golang
func NumberDuplicate(n int)int{
	count := 0;
	Chart := map[int]bool{}

	for i := n-1; i>0; i--{
		count++
		ChossenNumber := random(1, n+1)
		if Chart[ChossenNumber] == true{
			break
		}else {
			Chart[ChossenNumber] = true
		}
	}
	return count
}
```
With this, i made another algorithm that would run through the tests N amount of times and take the average. This way i could call and make tests like this:
```golang
fmt.Println(RunTestDuplicate(365, 1000000))            // This one will printout the average of a milion tests with 365 days as N.
```

The Running of many tests and averaging the reults is below:
```golang
func RunTestDuplicate(n int, t int)float64{


	Results := []int{}

	for i:=t; i>0; i--{
		Results = append(Results, NumberDuplicate(n))
	}

	var total float64 = 0
	for _, value := range Results{
		total += float64(value)
	}
	return total/float64(len(Results))
}
```

## Coupon collector problem.
#### Description
Generating random integers as in the previous exercise run experiments to validate the hypothesis that the number of integers generated before all possible values are generated is ~N HN.

#### Results & comments
Basicly the same as before. I ran tests and experiments. First one is with N of 100. How many integer generations does it take to have seen all possible values from 1 to N?. I made the algorithms and got these results shown below:

```golang
fmt.Println(RunTestCoupons(50, 10000)) //-> 225.1315
fmt.Println(RunTestCoupons(100, 10000)) //-> 518.729
fmt.Println(RunTestCoupons(200, 10000)) //-> 1178.5208
fmt.Println(RunTestCoupons(500, 10000)) //-> 3390.9957
fmt.Println(RunTestCoupons(1000, 10000)) //-> 7466.9836
fmt.Println(RunTestCoupons(5000, 100)) //-> 45867.37
fmt.Println(RunTestCoupons(25000, 100)) //-> 264603.78
```
Based purely on the outlook of the result, we can easily say that big o notation is definitly bigger than N. And grows exponentially. This is logically, since before we are done. We need to have seen all the possible values that goes from 0 to N. So the ammount of integers that needs to be generated needs to be bigger than N. Since we have randomized integer generation, it is very likely that an integer will be generated that we have allready seen before. So applying that to higher values of N. We must conclude that, by high N the likelyhood of getting a duplicate gets higher and higher the bigger N becomes. Also the less likely it is to generate that specific last value.

#### Essential Code

The fucntion to run this test is as follows:
```golang
func Couponons(n int)int{
	count := 0
	Remaining := n
	Chart := map[int]bool{}

	for {
		count++
		ChossenNumber := random(1, n+1)

		if Chart[ChossenNumber] == false{
			Chart[ChossenNumber] = true
			Remaining--
		}
		if Remaining == 0{
			break
			log.Println("No Remaning left")
		}
	}
	//log.Println(ThenFinalyCheck(Chart, n))   /// This was temporary to test thouroghly if the algorithm realy worked as inteded.
	return count
}
``` 

## Deck sort
#### Description
Explain how you would put a deck of cards in order by suit (in the
order spades, hearts, clubs, diamonds) and by rank within each suit, with the restriction
that the cards must be laid out face down in a row, and the only allowed operations are
to check the values of two cards and to exchange two cards (keeping them face down).

#### Results & Comments:
This exercise is very open for interpretation and open for many kinds of solutions. When translating the idea from physical cards to a program, this is where you can pick easy solutions or more accurate solutions. But in essense this problem can be solved by many ways. I've chosen this following way:

- 1. Start at the second first card from the left
- 2. Pick up and look at current card Card
- 3. If no additional card is being looked at appart from current card, pickup card to the left of that. Otherwise put down additional card and pick up card left of that.
- 4. If Additional card is higher in suit/rank than current card, go to step 3 swap cards.
- 5. Advance current card and Go to step 2 until current card cannot advance.

This is insertion sort. It could have been done with selection sort and shell sort too. Depending on how much we are allowed to remember. 

When it comes to translating this into a program, it is quite open for interpretation and how you'd go about doing it. I've chosen a little bit of both worlds. I've chosen to have the data structure in the accurate format, Suit and Rank. But when it comes to sorting. I'm transforming it into a format which is alot easier to sort and work with, which is just a number. In essense, the suit is basicly just something that comes fist and then rank afterwards. So you could either map each card to a position into a integer that would span 1 to 52. Or as i have done, is to just treat suits as a much higher value that is being appended to the rank. So Ace of diamonds is basicly, the last suit, but the first rank. You could say that last suit is the 4th suit in the collection, and give the suit the number 4. Then we coukd first sort suits and then ranks afterwards in 5 different sorts. But, as i've done is to append them together. Time suitnumber with 100 and plus it. So ace of diamonds would become the number 401. five of clubs will become 305, queen of hearts will become 212 etc. This way, we can sort these litle larger values easier in one go. Below is shown how the Suit/Rank struct can be transformed into easily sortable int array.

- Card

```golang
type Card struct {
	suit int
	rank int
}
```

- array of cards being transformed into easily sortable array and back again.

```golang
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
```

Furthermore, i've implemented some helper fucntions to better handle the features and such. To see full version, visit github page on https://github.com/Games-of-Threads/DSAL1-DFH/blob/master/DeckSorting.go

```golang
func getSuitFromInt(s int)string
func CreateNewDeck()[]Card
func getRankFromInt(r int)string
func ShuffleIntArray(input []int)[]int
func ShuffleDeck(Deck []Card)[]Card
```



#### Essential Code: 
To Run the deck sorting, any of the functions can be used independendly. But here is a showcase of some code that would have 100% functionality coverage.
```golang
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
```

The Insertion Sorting algorithm is as follows:

```golang
func SortArray(input []int)[]int{
	array := input
	for i := 1; i < len(array); i++ {
		for j := i; j > 0 && array[j] < array[j - 1]; j-- {
			array[j], array[j-1] = array[j-1], array[j]
		}
	}

	return array
}
```

This is the essential part to make it run. It is the sorting of an integer array.
