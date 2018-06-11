package main

import (
	"fmt"
	"os"
)

const (
	/***Testing iota***/
	iotaTestMessage = "iota test: %d %d\n"
	answer1         = iota //Gets set to 1 by default
	answer2                //Increments by 1 from answer1

	answer3 = iota * 2 //Gets set to 3*2
	answer4            //Increments by 2 from answer3

	/***Answer to life***/
	correctAnswer   = "The answer to life is %d\n"
	incorrectAnswer = "The answer to life is not %d\n"
	answer          = 42
)

var (
	editableAnswer = answer
)

func main() {
	helloWorld()                        //Hello World
	answerToLife()                      //Answer to life
	iotaTest()                          //Testing iota
	floatingPointTest()                 //Testing floating points
	booleanTest()                       //Testing booleans
	byteTest()                          //Testing bytes
	stringTest()                        //Testing strings
	printfReturnTest()                  //Testing Printf return
	switchStatementTest()               //Testing switch cases
	forloopTest()                       //Testing forloops
	writeToFile("Write this to a file") //Testing writing to a file
	arraysTest()                        //Testing arrays
	sliceTest()                         //Testing arrays
}

func helloWorld() {
	/***Hello World***/
	helloMessage := "Hello, World!"
	println(helloMessage)
}

func answerToLife() {
	println("\n*******Answer to life****************")
	/***Answer to life***/
	println(correctAnswer, answer)
	editableAnswer += 1
	println(incorrectAnswer, editableAnswer)
}

func iotaTest() {
	println("\n*******Testing iota*******************")
	/***Testing iota***/
	fmt.Printf(iotaTestMessage, answer1, answer2)
	fmt.Printf(iotaTestMessage, answer3, answer4)
}

func floatingPointTest() {
	println("\n*******Testing floating points********")

	/***Testing floating points***/
	var pi64 float64 = 3.14
	var pi32 = float32(3.14)
	println("Pi64 Value: %f", pi64)
	println("Pi64 Value(2 Decimal): %.2f", pi64)
	println("Pi32 Value: %f", pi32)
	println("Pi32 Value(2 Decimal): %.2f", pi32)
}

func booleanTest() {
	println("\n*******Testing booleans=**************")

	/***Testing booleans***/
	isTrue := true
	println("It is %t that %t is not %t:", isTrue, !isTrue, isTrue)
}

func byteTest() {
	println("\n*******Testing bytes******************")

	/***Testing bytes***/
	b := byte(65)
	println("Byte Value in Hex: %x", b) //Print out in hex
}

func stringTest() {
	println("\n*******Testing strings****************")

	//Testing substrings
	atoz := "the quick brown fox jumps over the lazy dog\n"
	//Everything in `` is treated literally
	atoz2 := `the quick brown fox jumps over the lazy dog\n`

	println(atoz)
	println(atoz2)

	println("%s", atoz[:10]) //Defaults to the beginning of the string
	println("%s", atoz[10:15])
	println("%s", atoz[16:19])
	println("%s", atoz[20:]) //Defaults to the end of the string

	//Testing substrings with forloops
	for i, r := range atoz {
		println("%d: %c", i, r)
	}

	for _, r := range atoz {
		print("%c", r)
	}

	for i := range atoz {
		print("%d ", i)
	}

	println("")
	println("Length: %d", len(atoz))
}

func printfReturnTest() {
	println("\n*******Testing printf****************")
	//n: number of bytes returned (Defined inside the scope of the if statement)
	//err: error returned
	if n, err := println("Hello, World!"); err != nil { //err is nil if nothing went wrong
		os.Exit(1)
	} else {
		println("Printed %d byte(s)", n)
	}
}

func switchStatementTest() {
	println("\n*******Testing Switch Statements*****")

	n, err := println("Hello, World!") //OK!
	//n, err := fmt.Printf("Hello, World!") //Wrong Number of bytes
	//n, err := fmt.Printf("") //No bytes output

	//This looks a lot cleaner than dealing with if else statements
	switch {
	case err != nil:
		os.Exit(1)
		//break (Not necessary as it does not fall through)
		//fallthrough (Allows you the ability to fall through)
	case n == 0:
		println("No bytes output")
	case n != 14:
		println("Wrong number of bytes")
	default:
		println("OK!")
	}

	atoz := "the quick brown fox jumps over the lazy dog"
	vowels := 0
	consonants := 0
	zeds := 0

	for _, r := range atoz {
		switch r {
		case 'a', 'e', 'i', 'o', 'u':
			vowels += 1
		case 'z':
			zeds += 1
			fallthrough //Because z is also a consonant
		default:
			consonants += 1
		}
	}

	println("Vowels: %d\nConsonants: %d\nZ: %d", vowels, consonants, zeds)
}

func forloopTest() {
	println("\n*******Testing forloops**************")

	for i, j := 0, 1; i < 10; i, j = i+1, j*2 {
		println("i: %d, j: %d", i, j)
	}

	var stop bool //Defaults to false
	counter := 0

	for !stop {
		println("counter: %d", counter)
		counter++
		if counter == 10 {
			stop = true
		}
	}
}

func writeToFile(msg string) (err error) {
	f, err := os.Create("helloworld.txt")
	if err != nil {
		return
	}
	defer f.Close() //Defers it to after the function completes

	f.Write([]byte(msg))
	return
}

func arraysTest() {
	println("\n************Testing Arrays***********")
	words := [4]string{"the", "quick", "brown", "fox"}
	//Arrays are passed by copy, not by reference (May impact performance)
	printWords(words)
	//words[2] is still "quick" because only the copy was edited
	printWords(words)
}

func sliceTest() {
	println("\n************Testing Slices**********")
	words := []string{"the", "quick", "brown", "fox", "jumps", "over",
		"the", "lazy", "dog"}
	//Slice is a reference to a subset of the Array
	printWordsSlice(words[0:4])
	//words[2] becomes blue
	printWordsSlice(words[0:4])

	words2 := make([]string, 4) //Contains 4 items
	println("Contains: %d Cap: %d", len(words2), cap(words2))
	words2[0] = "the"
	words2[1] = "quick"
	words2[2] = "brown"
	words2[3] = "fox"
	println("Contains: %d Cap: %d", len(words2), cap(words2))
	printWordsSlice(words2) //Remains Brown

	words3 := make([]string, 0, 4) //Contains 0 items but can hold up to 4 items
	println("Contains: %d Cap: %d", len(words3), cap(words3))
	words3 = append(words3, "the")
	words3 = append(words3, "quick")
	words3 = append(words3, "brown")
	words3 = append(words3, "fox")
	println("Contains: %d Cap: %d", len(words3), cap(words3))
	printWordsSlice(words3)          //Remains Brown
	words3 = append(words3, "jumps") //Cap doubles because there was not enough room
	println("Contains: %d Cap: %d", len(words3), cap(words3))
	printWordsSlice(words3) //Notice that brown is changed to blue
}

func printWords(w [4]string) {
	for _, word := range w {
		print("%s ", word)
	}
	println("")

	w[2] = "blue"
}

func printWordsSlice(w []string) {
	for _, word := range w {
		print("%s ", word)
	}
	println("")

	w[2] = "blue"
}

func println(format string, arguments ...interface{}) (int, error) {
	n, err := fmt.Fprintf(os.Stdout, format+"\n", arguments...)
	return n, err
}

func print(format string, arguments ...interface{}) (int, error) {
	n, err := fmt.Fprintf(os.Stdout, format, arguments...)
	return n, err
}
