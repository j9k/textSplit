/* This program is for splitting up long text files in order to email them to prison inmates.
Jail email systems limit emails to a maximum number of characters and lines*/

/*  done 	1. open the text file and put in a string
	done	2. separate the text into slices made up of strings of sentences
	done	3. get parameters from command line such as the number of characters per file.
4. open the log file and figure out what the last entry was to retrieve parameters for the next set of runes to copy
	done	5. range through the slice of strings to count the runes in each sentence
	done	6. need some way to check if a sentence goes over the limit such as a counter after each period.
7. write the used data to the log file for use the next time
8. logic to create the log file if none resides in the folder
9. write a function to write the log file pass it the struct to append to the end of the file

*/
// https://golang.org/
// go version go1.15.3 windows/amd64
// only uses the standard library
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

type logEntry struct { // struct for holding log entry info
	fileNameInput  string
	fileNameoutput string
	logId          int
	runesRequested int
	runesReturned  int
	startSentence  int
	endSentence    int
}

type sentence struct {
	stringOfSentence string
	runesInSentence  int
}

func main() {
	fmt.Println("Initializing the variables")
	var textFile string = "original.txt"
	var fullText string = "abcde fghij"
	var startSentence int = 0
	var endSentence int = 1
	var runesRequested int = 1467
	var runesReturned int = 0
	var previousLog logEntry // this will be used to track which text has been copied from the original
	var endStatementFlag bool = false

	previousLog.fileNameInput = "original.txt"
	previousLog.fileNameoutput = "previous-output.txt"
	previousLog.logId = 0
	previousLog.runesRequested = 10
	previousLog.runesReturned = 5
	previousLog.startSentence = 0
	previousLog.endSentence = 1

	var futureLog logEntry // this will be used to update a log file and create the copied text file name
	futureLog.fileNameInput = "original.txt"
	futureLog.fileNameoutput = "NewText.txt" // will be combining date and original.txt to
	futureLog.runesRequested = 250           // need to get from flag or os.args
	futureLog.logId = previousLog.logId + 1
	fmt.Println("futureLog.logId = ", futureLog.logId)

	fmt.Println("Aquiring the command line parameters")
	for i, a := range os.Args { // this loop gets arguments from the command line and returns a slice of strings
		fmt.Printf("%v = %v\n", i, a)
	}

	if os.Args[1] == "-getnext" { //compare the terminal input and get the value
		v := os.Args[2]
		if s, err := strconv.Atoi(v); err == nil { // I don't get why this works yet but just copied an example to get it working
			runesRequested = s
		}

	}

	if os.Args[3] == "-from" { // get the name of the text file
		textFile = os.Args[4]
	}

	if os.Args[5] == "-startingat" { //compare the terminal input and get the value
		v := os.Args[6]
		if s, err := strconv.Atoi(v); err == nil { // I don't get why this works yet but just copied an example to get it working
			startSentence = s
		}
	}

	fmt.Println("No more than ", runesRequested, " characters will be returned from ", textFile)

	fmt.Println("Reading the file ", textFile)
	fullTextByte, err := ioutil.ReadFile(textFile) //change to variable from command line at some point
	if err != nil {
		return
	}

	fullText = string(fullTextByte) //converts the byte to string
	fmt.Println("runes in full text = ", utf8.RuneCountInString(fullText))
	fullTextSlice := strings.SplitAfter(fullText, ".") // uses a "." to separate the text into separate strings
	fullTextSliceLength := len(fullTextSlice) - 1
	fmt.Println("fullTextSliceLength = ", fullTextSliceLength)
	sentences := []sentence{} //creates a slice of type sentence
	outputSentencesSlice := []string{}
	runeAccumulator := 0 // keeps a rolling count to compare to runesRequested

	endSentence = startSentence
	//change the loop to start at a specific element in the slice

	fmt.Println("Retrieving the sentences from the text")
CountFind:
	for i := 0; i <= fullTextSliceLength; i++ {

		//fmt.Println("slice position", startSentence)
		//fmt.Println(fullTextSlice[i])

		sentenceRuneCount := utf8.RuneCountInString(fullTextSlice[endSentence]) //
		//fmt.Println("sentenceRuneCount = ", sentenceRuneCount)

		runeAccumulator = sentenceRuneCount + runeAccumulator // a running tally of runes
		//fmt.Println("runeAccumulator = ", runeAccumulator)

		s := sentence{fullTextSlice[endSentence], sentenceRuneCount} //creat a single variable of struct type sentence

		if runeAccumulator <= runesRequested { // compare and stop if over the runes requested amount
			//fmt.Println("runeAccumulator is less than runes Requested")
			//i = fullTextSliceLength - 1
			sentences = append(sentences, s) // add current type sentence to slice of sentences
			//fmt.Println(sentences[i])
			outputSentencesSlice = append(outputSentencesSlice, fullTextSlice[endSentence])
		} else {
			//endSentence = startSentence
			break CountFind
		}

		if endSentence >= fullTextSliceLength {
			endStatementFlag = true
			break CountFind
		}
		endSentence++
	} // result of this for statement is I have my output text inside of a slice of type sentences
	// I need to get the string elements out of the sentences variable now.

	outputTextHeader := fmt.Sprint("sentences ", startSentence, " to ", endSentence-1, " of ", fullTextSliceLength)
	fmt.Println(outputTextHeader)

	outputText := strings.Join(outputSentencesSlice, "") //convert the slice of strings to a single string
	fmt.Println(outputText)                              //print the output string to the terminal

	runesReturned = utf8.RuneCountInString(outputText)

	fmt.Println("runesReturned", runesReturned)

	fmt.Println(endSentence)

	if endStatementFlag == true {
		fmt.Println("YOU HAVE REACHED THE END OF THE STORY. TIME TO FIND ANOTHER ONE FOR DAVE!!!!!!!!!!!!")
	} else {

		//fmt.Println("Collecting data for the next session")

		//fmt.Println("Writing the log file")
	}
	fmt.Println("end of line.......")
}

//func flagSetup() { // function to setup the flag package without crouding the main function
//	var maxCharactersToGet = flag.Int("get", 1234, "get is the maximum number of characters you are trying to get from the text") // flags to initialize -get int -from text.txt

//}
