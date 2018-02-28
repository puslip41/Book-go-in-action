package search

import "log"

type Result struct {
	Field   string
	Content string
}

type Searcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

func Match(matcher Searcher, feed *Feed, searchTerm string, results chan<- *Result) {
	searchResults, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println(err)
		return
	}

	for _, result := range searchResults {
		results <- result
	}
}

func Display(results chan *Result) {
	for result := range results {
		log.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}

	log.Println("검색이 완료되었습니다.")
}
