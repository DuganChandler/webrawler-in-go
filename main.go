package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("no args provided")
		os.Exit(1)
	}

	if len(args) > 3 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	url := args[0]

	maxConcurrency, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("unable to convert to int")
		os.Exit(1)
	}
	maxPages, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("unable to convert to int")
		os.Exit(1)
	}

	cfg, err := configure(url, maxConcurrency, maxPages)
	if err != nil {
		fmt.Printf("error - configure: %v", err)
		return
	}

	fmt.Printf("starting crawl of: %s", url)
	fmt.Println("")

	cfg.wg.Add(1)
	go cfg.crawlPage(url)
	cfg.wg.Wait()

	// for normalizedURL, count := range cfg.pages {
	// 	fmt.Printf("%d - %s\n", count, normalizedURL)
	// }

	printReport(cfg.pages, url)

}
