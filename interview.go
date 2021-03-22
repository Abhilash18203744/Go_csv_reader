package customerimporter

import (
	"encoding/csv"
	"errors"
	"io"
	"log"
	"os"
	"strings"
	"sync"
)

func basicDomainCounter(csvFile string) (map[string]int, error) {

	// Map data structure for output
	domainCount := make(map[string]int)

	// Loading csv file
	file, err := os.Open(csvFile)
	if err != nil {
		log.Print("Error: ", err)
		return domainCount, err
	}
	defer file.Close()

	// Creating csv reader
	reader := csv.NewReader(file)
	index := -1

	headers, err := reader.Read()
	for i, header := range headers {
		if header == "email" {
			index = i
			break
		}
	}

	// Checking if email column exists in the csv
	if index == -1 {
		log.Print("Error: email column not present in input file")
		return domainCount, errors.New("Error: email column not present in input file")
	}

	// Parsing data one line at a time and maintaining count of email domain users
	for line, err := reader.Read(); err == nil; line, err = reader.Read() {
		domainName := strings.Split(line[index], "@")
		// Following time sleep added in order to make this operation expensive task
		// time.Sleep(1 * time.Nanosecond)
		if len(domainName) == 2 {
			domainCount[domainName[1]]++
		}
	}

	return domainCount, nil
}

func concWPDomainCounter(csvFile string) (map[string]int, error) {

	// Map data structure for output
	domainCount := make(map[string]int)

	// Loading csv file
	file, err := os.Open(csvFile)
	if err != nil {
		log.Print("Error: ", err)
		return domainCount, err
	}
	defer file.Close()

	// Creating csv reader
	reader := csv.NewReader(file)
	numWps := 200                       // Number of jobs, changed as per requirement
	jobs := make(chan []string, numWps) // Job channel
	result := make(chan string)         // Result channel
	var waitGroup sync.WaitGroup        // Wait Group
	index := -1

	headers, err := reader.Read()
	for i, header := range headers {
		if header == "email" {
			index = i
			break
		}
	}

	// Checking if email column exists in the csv
	if index == -1 {
		log.Print("Error: email column not present in input file")
		return domainCount, errors.New("Error: email column not present in input file")
	}

	worker := func(jobs <-chan []string, results chan<- string) {
		for {
			select {
			case job, ok := <-jobs:
				if !ok {
					return
				}
				results <- parseData(job, index)
			}
		}
	}

	// Init workers
	for w := 0; w < numWps; w++ {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			worker(jobs, result)
		}()
	}

	go func() {
		for {
			data, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Println("ERROR: ", err.Error())
				break
			}
			jobs <- data
		}
		close(jobs) // close jobs to signal workers that no more job are incoming.
	}()

	go func() {
		waitGroup.Wait()
		close(result)
	}()

	for domainName := range result {
		if domainName != "" {
			domainCount[domainName]++
		}
	}

	return domainCount, nil
}

func parseData(dataLine []string, index int) string {
	domainName := strings.Split(dataLine[index], "@")
	// Following time sleep added in order to make this operation expensive task
	// time.Sleep(1 * time.Nanosecond)
	if len(domainName) == 2 {
		return domainName[1]
	}
	return ""
}
