package main

import "fmt"

func main() {

	//strig
	var nameOne string = "mario"
	var nameTwo = "luigi"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "peach"
	nameThree = "bowser"

	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "yoshi"

	fmt.Println(nameOne, nameTwo, nameThree, nameFour)

	//ints

	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	//bits and num

	var numOne int16 = 215
	var numtWO int16 = -215
	var numThree int8 = 25

	fmt.Println(numOne, numtWO, numThree)

	//float

	var scoreOne float32 = 25.98
	var scoreTwo float64 = 9991231321.31231
	scoreThree := 1.5
	fmt.Println(scoreOne, scoreTwo, scoreThree)

	fmt.Println("my ages is", ageTwo, "and my name is", nameTwo)

	//Printf (formatted strings) %_ = especificador de formato

	fmt.Printf("my age is %v and my name is %v \n", ageOne, nameOne)
	fmt.Printf("my age is %q and my name is %q \n", ageOne, nameOne)
	fmt.Printf("age is of type %T \n", ageOne)
	fmt.Printf("your score is %.2f point!\n", 225.25)

	//Sprintf (save formatted strings)
	var str = fmt.Sprintf("my age is %v and my name is %v \n", ageOne, nameOne)
	fmt.Println("the save string is:", str)

	//arrays

	var ages [3]int = [3]int{20, 26, 30}
	names := [4]string{"yoshi", "Mario", "Peach", "Luigi"}
	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))
	names[3] = "Bowser"
	fmt.Println(names, len(names))

	// slices (uses arrays under the hood)
	var scores = []int{100, 50, 60}
	scores[1] = 20
	scores = append(scores, 85)

	fmt.Println(scores, len(scores))

	//slices ranges

	rangeOne := names[1:]
	rangeTwo := names[1:3]
	rangeThree := names[:3]
	fmt.Println(rangeOne)
	fmt.Println(rangeTwo)
	fmt.Println(rangeThree)
	rangeOne = append(rangeOne, "koopa")
	fmt.Println(rangeOne)
}
