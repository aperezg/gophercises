package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	filePath := flag.String("csv", "problem.csv", "a csv file in the format of 'question,answer' (default 'problem.csv')")
	limit := flag.Int("limit", 30, "the time limit for quiz in seconds (default 30)")
	shuffle := flag.Bool("shuffle", false, "shuffle the questions")

	flag.Usage = func() {
		flag.PrintDefaults()
	}
	flag.Parse()

	q, err := readFile(*filePath)
	if err != nil {
		log.Fatal(err)
	}

	if *shuffle {
		shuffleQuiz(q)
	}

	fmt.Println("Press enter to start the quiz...")
	fmt.Scanln()
	play(q, *limit)
}

func play(quiz []Quiz, limit int) {
	var score int
	quit := make(chan bool)
	go func() {
		for i, q := range quiz {
			fmt.Printf("Problem #%d: %s =\n", i+1, q.Question)
			var answer string
			fmt.Scanf("%s\n", &answer)
			answer = strings.TrimSpace(answer)
			answer = strings.ToLower(answer)
			if answer == q.Answer {
				score++
			}
		}
		quit <- true
	}()
	select {
	case <-quit:
		printTotal(score, len(quiz))
	case <-time.After(time.Duration(limit) * time.Second):
		printTotal(score, len(quiz))
	}

}

func printTotal(score, total int) {
	fmt.Printf("You scored %d of %d.\n", score, total)
}

func shuffleQuiz(quiz []Quiz) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(quiz) > 0 {
		n := len(quiz)
		index := r.Intn(n)
		quiz[n-1], quiz[index] = quiz[index], quiz[n-1]
		quiz = quiz[:n-1]
	}
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
			Answer:   strings.ToLower(l[1]),
		})
	}

	return quiz, nil
}

// Quiz define format of quiz with questions and answers
type Quiz struct {
	Question string
	Answer   string
}
