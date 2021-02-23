package main

import (
	"encoding/csv"
	"flag"
	"os"
	"fmt"
	"strings"
	"time"
)

type Problem struct {
	q string
	a string
}

func main()  {
	/*
		Define a flag
	 */
	csvFilename := flag.String("csv", "problems.csv", "Math Problems Quiz")
	timeLimit := flag.Int("limit", 30, "Time limit for the Quiz")

	/*
		Parse the flag
	 */
	flag.Parse()

	/*
		Read the file
	 */
	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV [%s] \n", *csvFilename))
	}

	/*
		Create a reader variable from csv reader lib
	 */
	r := csv.NewReader(file)

	/*
		Read all the lines in the file
	 */
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to read the provided CSV file...")
	}

	/*
		Fetch probelms struct with q's and a's
	 */
	problems := parseLines(lines)

	/*
		Add a timer to a function
	 */
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	/*
		print all of the problems
	 */
	correct := 0
	for i, p := range problems {
		fmt.Printf("Question is [%d]: %s \n", i+1, p.q)
		answerCh := make(chan string)

		go func() {
			var o string
			fmt.Scanf("%s \n", &o)
			answerCh <- o
		}()

		select {
		case <- timer.C:
			fmt.Printf("\nTimeOut: *** You scored %d out of %d ***** \n", correct, len(problems))
			return
		case o := <-answerCh:
			if o == p.a {
				correct++
			}
		}
	}
}

func parseLines(lines [][]string) []Problem {
	qsets := make([]Problem, len(lines))

	for i, line := range lines {
		qsets[i] = Problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}

	return qsets
}

func exit(msg string)  {
	fmt.Println(msg)
	os.Exit(1)
}