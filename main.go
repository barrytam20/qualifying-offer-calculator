package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// scrape website for player data
	url := "https://questionnaire-148920.appspot.com/swe/data.html"
	fmt.Println("loading data from", url)
	players, err := getPlayers(url)
	if err != nil {
		log.Fatalf(err.Error())
	}

	// insert player salary into max heap
	salaries := &maxheap{}
	noSalaryData := []*player{}
	for _, p := range players {
		if p.salary == 0 {
			noSalaryData = append(noSalaryData, p)
		} else {
			salaries.push(p.salary)
		}
	}
	salariesProcessed := salaries.size

	// calculate qualifying offer by taking the average of the top 125 salaries
	sum := 0
	for i := 0; i < 125; i++ {
		sum += salaries.pop()
	}
	qo := float64(sum) / 125

	fmt.Println("qualifying offer is:", currencyPrint(qo))
	fmt.Println("number of player salaries processed:", salariesProcessed)
	fmt.Println("number of players missing salary data:", len(noSalaryData))
}

type player struct {
	name   string
	year   string
	level  string
	salary int
}

func (p *player) parseSalary(salaryText string) {
	// https://pythonexamples.org/python-regex-extract-find-all-the-numbers-in-string/#2
	// https://golang.org/pkg/regexp/#example_Regexp_FindAll
	re := regexp.MustCompile(`[0-9]+`)
	parts := re.FindAll([]byte(salaryText), -1)
	digits := ""
	for _, part := range parts {
		digits += string(part)
	}
	salary, err := strconv.Atoi(digits)
	if err == nil && salary > 0 {
		p.salary = salary
	}
}

func getPlayers(url string) ([]*player, error) {
	players := []*player{}

	// https://github.com/PuerkitoBio/goquery
	res, err := http.Get(url)
	if err != nil {
		return players, fmt.Errorf("error retreiving data from %s: %s", url, err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return players, fmt.Errorf("error retreiving data from %s: %d %s", url, res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return players, fmt.Errorf("error reading html %s: %s", url, err)
	}
	doc.Find("#salaries-table tbody tr").Each(func(i int, s *goquery.Selection) {
		p := &player{
			name:  s.Find(".player-name").Text(),
			year:  s.Find(".player-year").Text(),
			level: s.Find(".player-level").Text(),
		}
		p.parseSalary(s.Find(".player-salary").Text())
		players = append(players, p)
	})
	return players, nil
}

type maxheap struct {
	data []int
	size int
}

func getParentIndex(i int) int {
	return (i - 1) / 2
}

func getChildrenIndices(i int) (int, int) {
	return i*2 + 1, i*2 + 2
}

func (m *maxheap) push(v int) {
	m.size++
	if m.size == 1 {
		m.data = []int{v}
		return
	}
	m.data = append(m.data, v)
	newIndex, parentIndex := m.size-1, getParentIndex(m.size-1)
	for newIndex > 0 && m.data[newIndex] > m.data[parentIndex] {
		m.data[newIndex], m.data[parentIndex] = m.data[parentIndex], m.data[newIndex]
		newIndex = parentIndex
		parentIndex = getParentIndex(newIndex)
	}
}

func (m *maxheap) pop() int {
	out := m.data[0]
	m.size--
	m.data[0] = m.data[m.size]
	m.data = m.data[:m.size]
	m.heapify(0)
	return out
}

func (m *maxheap) heapify(i int) {
	if i > m.size/2 {
		return
	}
	top := i
	left, right := getChildrenIndices(i)
	if left < m.size && m.data[left] > m.data[top] {
		top = left
	}
	if right < m.size && m.data[right] > m.data[top] {
		top = right
	}
	if top != i {
		m.data[top], m.data[i] = m.data[i], m.data[top]
		m.heapify(top)
	}
}

func currencyPrint(a float64) string {
	ai := (int)(a * 100)
	dollars, cents := ai/100, ai%100
	dp := []string{}
	for dollars > 0 {
		dp = append([]string{strconv.Itoa(dollars % 1000)}, dp...)
		dollars /= 1000
	}
	if len(dp) == 0 {
		return fmt.Sprintf("$0.%d", cents)
	}
	return fmt.Sprintf("$%s.%d", strings.Join(dp, ","), cents)
}
