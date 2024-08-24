package main

import (
	"fmt"
	"sort"
)

type report struct {
	links int
	url   string
}

func printReport(pages map[string]int, baseURL string) {
	fmt.Println("=====================")
	fmt.Println("  REPORT for " + baseURL)
	fmt.Println("=====================")

	reports := sortPages(pages)
	for _, report := range reports {
		fmt.Printf("Found %d internal links to %s\n", report.links, report.url)
	}
}

func sortPages(pages map[string]int) []report {
	reports := make([]report, len(pages))

	for key, val := range pages {
		rep := report{
			links: val,
			url:   key,
		}
		reports = append(reports, rep)
	}

	sort.Slice(reports, func(i, j int) bool {
		if reports[i].links != reports[j].links {
			return reports[i].links > reports[j].links
		} else {
			return reports[i].url > reports[j].url
		}
	})

	return reports
}
