package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

// Quiz define format of quiz with questions and answers
type Quiz struct {
	Question string
	Answer   string
}

func main() {
	filePath := flag.String("csv", "problem.csv", "a csv file in the format of 'question,answer' (default 'problem.csv')")
	flag.Usage = func() {
		flag.PrintDefaults()
	}
	flag.Parse()

	q, err := readFile(*filePath)
	if err != nil {
		log.Fatal(err)
	}

	play(q)
}

func play(quiz []Quiz) {
	problemN := 1
	scann := bufio.NewScanner(os.Stdin)
	score := 0

	for _, q := range quiz {
		fmt.Printf("Problem #%d: %s = ", problemN, q.Question)
		scann.Scan()
		answer := scann.Text()

		if answer == q.Answer {
			score++
		}

		problemN++
	}

	fmt.Printf("You scored %d of %d.\n", score, len(quiz))
}

func readFile(path string) ([]Quiz, error) {
	f, err := os.Open(path)
	reader := csv.NewReader(bufio.NewReader(f))
	if err != nil {
		return []Quiz{}, err
	}

	var quiz []Quiz
	for {
		l, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return []Quiz{}, err
		}

		quiz = append(quiz, Quiz{
			Question: l[0],
			Answer:   l[1],
		})
	}

	return quiz, nil
}
