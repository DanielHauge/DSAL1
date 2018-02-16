# 1. Compulsory algorithm Assignment - Experiments and Sorting.
This is the hand-in paper for Datastructure & Algorithms course. It consists of 3 bullets points (Subjects) which needs attention. This paper will include results, comments to results and the essential code.

I've chossen to make all experiments and programs in the language "Golang".

## Birthday Experiencement
Write a program that takes an integer N from the command line and generate a random sequence of integers between 0 and N – 1. Run experiments to validate the hypothesis that the number of integers generated before the first repeated value is found is ~√(pi)N/2.

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

The big O notation is ofcause the formular. O(√(pi)N/2) But worst case scenario, it could also be O(N). It could potentially, go through the whole available valuespace before it finds a duplicate. Which is N. I've calculated this by the following function.

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
fmt.Println(RunTestDuplicate(365, 1000))            // This one will printout the average of a thousand tests with 365 days as N.
```

## Coupon collector problem.
Generating random integers as in the previous exercise run experiments to validate the hypothesis that the number of integers generated before all possible values are generated is ~N HN.

Basicly the same as before. I ran tests and experiments. First one is with N of 100. How many integer generations does it take to have seen all possible values from 1 to 100?. I made the algorithms and got these results shown below:

```golang
fmt.Println(RunTestCoupons(50, 10000)) -> 225.1315
fmt.Println(RunTestCoupons(100, 10000)) -> 518.729
fmt.Println(RunTestCoupons(200, 10000)) -> 1178.5208
fmt.Println(RunTestCoupons(500, 10000)) -> 3390.9957
fmt.Println(RunTestCoupons(1000, 10000)) -> 7466.9836
fmt.Println(RunTestCoupons(5000, 100)) -> 45867.37
fmt.Println(RunTestCoupons(25000, 100)) -> 264603.78
```
Based purely on the outlook of the result, we can easily say that big o notation is definitly bigger than N. And grows exponentially. This is logically, since before we are done. We need to have seen all the possible values that goes from 0 to N. So the ammount of integers that needs to be generated needs to be bigger than N. Since we have randomized integer generation, it is very likely that an integer will be generated that we have allready seen before. So applying that to higher values of N. We must conclude that, by high N the likelyhood of getting a duplicate gets higher and higher the bigger N becomes. Also the less likely it is to generate that specific last value.

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
Explain how you would put a deck of cards in order by suit (in the
order spades, hearts, clubs, diamonds) and by rank within each suit, with the restriction
that the cards must be laid out face down in a row, and the only allowed operations are
to check the values of two cards and to exchange two cards (keeping them face down).
