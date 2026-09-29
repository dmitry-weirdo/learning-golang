package main

import (
	"fmt"
	"slices"
	"strings"
)

const HTTP_PREFIX = "http://"

func crawl(startUrl string, htmlParser HtmlParser) []string {
	host := getHost(startUrl)
	//fmt.Printf("Start url host: \"%v\" \n", host)

	queue := make([]string, 0)

	queue = append(queue, startUrl)

	result := make([]string, 0)

	// track the visited URLs to avoid duplicates
	visited := make(map[string]bool)
	visited[startUrl] = true

	for len(queue) > 0 {
		currentLevelLength := len(queue)

		for range currentLevelLength {
			// pull from queue
			url := queue[0]
			queue = queue[1:]

			// add current url to result
			result = append(result, url)

			// add non-visited neighbors to the queue
			neighbors := htmlParser.GetUrls(url)
			for _, neighbor := range neighbors {
				if visited[neighbor] {
					continue
				}

				if getHost(neighbor) != host { // we only proceed within the same host as the starting url
					continue
				}

				visited[neighbor] = true // mark neighbor as visited
				queue = append(queue, neighbor)
			}

		}
	}

	return result
}

func getHost(url string) string {
	urlWithoutHttp := url[len(HTTP_PREFIX):]

	index := strings.Index(urlWithoutHttp, "/")

	if index == -1 {
		// no next slash after "http://" -> return the whole string
		return url
	}

	return url[:index+len(HTTP_PREFIX)]
}

type HtmlParser struct {
	urls       []string
	adj        [][]int
	urlToIndex map[string]int
}

func (this *HtmlParser) GetUrls(url string) []string {
	index := this.urlToIndex[url]
	neighbors := this.adj[index]

	result := make([]string, len(neighbors))

	for i, neighborIndex := range neighbors {
		result[i] = this.urls[neighborIndex]
	}

	return result
}

func NewHtmlParser(urls []string, edges [][]int) HtmlParser {
	n := len(urls)
	adj := createAdjacencyListDirectedUnweighted(n, edges)

	urlToIndex := make(map[string]int)

	for i, url := range urls {
		urlToIndex[url] = i
	}

	return HtmlParser{
		urls:       urls,
		adj:        adj,
		urlToIndex: urlToIndex,
	}
}

func createAdjacencyListDirectedUnweighted(n int, edges [][]int) [][]int {
	adj := make([][]int, n)

	from := 0
	to := 0

	for _, v := range edges {
		from = v[0]
		to = v[1]

		if from >= n || to >= n { // todo: this is an ugly hack since the test-case 8 / 21 (at least) is incorrect and uses out-of-bound indexes in the edges
			continue
		}

		// add from -> to
		if adj[from] == nil {
			adj[from] = []int{to}
		} else {
			adj[from] = append(adj[from], to)
		}
	}

	return adj
}

func test(urls []string, edges [][]int, startUrl string, expectedResult []string) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("URLs: %v \n", urls)
	fmt.Printf("Edges between URLs: %v \n", edges)
	fmt.Printf("Start URL: %v \n", startUrl)

	parser := NewHtmlParser(urls, edges)

	result := crawl(startUrl, parser)

	// to make same soring as expected result, sort both lists
	slices.Sort(result)
	slices.Sort(expectedResult)

	fmt.Printf("URLs crawled from the start page: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if len(result) != len(expectedResult) {
		fmt.Printf("FAILURE: expected result length = %v, actual result length = %v \n", len(expectedResult), len(result))
		return
	}

	for i, v := range result {
		if v != expectedResult[i] {
			fmt.Printf("FAILURE: expected result[%v] = %v, actual result[%v] = %v \n", i, expectedResult[i], i, v)
			return
		}
	}
}

func test1() {
	urls := []string{
		"http://news.yahoo.com",
		"http://news.yahoo.com/news",
		"http://news.yahoo.com/news/topics/",
		"http://news.google.com",
		"http://news.yahoo.com/us",
	}

	edges := [][]int{
		{2, 0},
		{2, 1},
		{3, 2},
		{3, 1},
		{0, 4},
	}

	startUrl := "http://news.yahoo.com/news/topics/"

	expected := []string{
		"http://news.yahoo.com",
		"http://news.yahoo.com/news",
		"http://news.yahoo.com/news/topics/",
		"http://news.yahoo.com/us",
	}

	test(urls, edges, startUrl, expected)
}

func test2() {
	urls := []string{
		"http://news.yahoo.com",
		"http://news.yahoo.com/news",
		"http://news.yahoo.com/news/topics/",
		"http://news.google.com",
	}

	edges := [][]int{
		{0, 2},
		{2, 1},
		{3, 2},
		{3, 1},
		{3, 0},
	}

	startUrl := "http://news.google.com"

	expected := []string{
		"http://news.google.com",
	}

	test(urls, edges, startUrl, expected)
}

func test3() {
	// Failing test-case 8 / 21
	urls := []string{
		"http://news.yahoo.com/oil-price-fundamental-weekly-forecast-200421659.html",
		"http://news.yahoo.com/latest-israeli-troops-kill-gunman-075217877.html",
		"http://news.yahoo.com/elizabeth-warren-unveils-plan-reduce-193138508.html",
		"http://news.yahoo.com/us-stocks-wall-street-set-125725929.html",
		"http://news.yahoo.com/news/topics/pets",
		"http://news.yahoo.com/m/566613eb-406d-3f03-8007-c764674d10ea/after-facebook-and-twitter%2C.html",
		"http://news.yahoo.com/news/topics/islamic-state",
		"http://news.yahoo.com/miss-usa-almost-didnt-wear-her-hair-natural-after-overhearing-criticism-090000447.html",
		"http://news.yahoo.com/middle-school-teacher-faces-backlash-after-handing-out-genderidentity-worksheet-233129871.html",
		"http://news.yahoo.com/5-influencers-learnt-love-curls-085800700.html",
		"http://news.yahoo.com/liam-hemsworth-doesnt-want-talk-161900831.html",
		"http://news.yahoo.com/explainer-medicare-for-all-single-payer-public-option-090000500.html",
		"http://news.google.com/stories/articles/CAIiEKskZxv044Ls9LYaBydkVOUqGQgEKhAIACoHCAowyNj6CjDyiPICMLv-xAU?hl=en-US&gl=US&ceid=US%3Aen",
		"http://news.google.com/stories/articles/CAIiEARUeBV0VjvOoOi6jCwfzLEqGQgEKhAIACoHCAow4Zn5CjCu8uACMLTRlgY?hl=en-US&gl=US&ceid=US%3Aen",
		"http://news.google.com/publications/articles/CAIiEBAXtnqlnOT-EDDNMkS6Dy8qGAgEKg8IACoHCAowhK-LAjD4ySww-9S0BQ?hl=en-US&gl=US&ceid=US%3Aen",
	}

	edges := [][]int{
		{0, 1},
		{0, 13},
		{0, 0},
		{1, 0},
		{1, 4},
		{1, 9},
		{1, 11},
		{1, 9},
		{1, 4},
		{2, 3},
		{2, 9},
		{2, 8},
		{3, 2},
		{3, 4},
		{3, 13},
		{3, 6},
		{3, 8},
		{4, 2},
		{5, 7},
		{5, 0},
		{5, 8},
		{5, 8},
		{6, 14},
		{6, 8},
		{6, 14},
		{7, 10},
		{7, 12},
		{7, 12},
		{7, 4},
		{8, 12},
		{8, 4},
		{8, 10},
		{8, 13},
		{8, 5},
		{9, 14},
		{9, 9},
		{9, 3},
		{9, 9},
		{10, 5},
		{11, 8},
		{11, 13},
		{11, 7},
		{11, 8},
		{11, 7},
		{11, 14},
		{12, 10},
		{12, 11},
		{12, 4},
		{12, 12},
		{12, 6},
		{12, 7},
		{12, 2},
		{13, 5},
		{13, 6},
		{13, 10},
		{13, 1},
		{15, 4},
		{15, 8},
	}

	startUrl := "http://news.yahoo.com/news/topics/islamic-state"

	expected := []string{
		"http://news.yahoo.com/5-influencers-learnt-love-curls-085800700.html",
		"http://news.yahoo.com/elizabeth-warren-unveils-plan-reduce-193138508.html",
		"http://news.yahoo.com/explainer-medicare-for-all-single-payer-public-option-090000500.html",
		"http://news.yahoo.com/latest-israeli-troops-kill-gunman-075217877.html",
		"http://news.yahoo.com/liam-hemsworth-doesnt-want-talk-161900831.html",
		"http://news.yahoo.com/m/566613eb-406d-3f03-8007-c764674d10ea/after-facebook-and-twitter%2C.html",
		"http://news.yahoo.com/middle-school-teacher-faces-backlash-after-handing-out-genderidentity-worksheet-233129871.html",
		"http://news.yahoo.com/miss-usa-almost-didnt-wear-her-hair-natural-after-overhearing-criticism-090000447.html",
		"http://news.yahoo.com/news/topics/islamic-state",
		"http://news.yahoo.com/news/topics/pets",
		"http://news.yahoo.com/oil-price-fundamental-weekly-forecast-200421659.html",
		"http://news.yahoo.com/us-stocks-wall-street-set-125725929.html",
	}

	test(urls, edges, startUrl, expected)

}

func main() {
	// 1236. Web Crawler
	test1()
	test2()
	test3()
}
