package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	printResult()
}

func writeToDisk() {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	// On every a element which has href attribute call callback
	c.OnHTML(`div[class=exchange-rate]`, func(e *colly.HTMLElement) {
		fmt.Println("Scrapping exchange rate ...!")
		maintext := ""
		mapper := map[string][]string{}

		e.ForEach(`div div:not(:first-child)`, func(_ int, e *colly.HTMLElement) {
			text := e.Text
			// strong := ""
			// span := ""

			var strongArray []string
			e.ForEach(`strong`, func(_ int, e *colly.HTMLElement) {
				strongText := e.Text
				text = strings.Replace(text, strongText, "", -1)
				strongArray = append(strongArray, strongText)
			})

			var spanHeader string
			e.ForEach(`span`, func(_ int, e *colly.HTMLElement) {
				span := e.Text
				text = strings.Replace(text, span, "", -1)
				spanHeader = span
			})

			mapper[spanHeader] = strongArray
			maintext = maintext + text + "\n"
		})

		data := []byte(maintext)
		ioutil.WriteFile("currency.txt", data, 0777)

		fmt.Println("---------------------------------------\n")
		fmt.Println(maintext)
		fmt.Println("\n---------------------------------------")
	})

	c.Visit("https://www.kbzbank.com/en/")
}

func printResult() {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	// On every a element which has href attribute call callback
	c.OnHTML(`div[class=exchange-rate]`, func(e *colly.HTMLElement) {
		fmt.Println("Scrapping exchange rate ...!")
		maintext := ""

		e.ForEach(`div div:not(:first-child)`, func(_ int, e *colly.HTMLElement) {
			text := e.Text
			// strong := ""
			// span := ""

			var strongArray []string
			e.ForEach(`strong`, func(_ int, e *colly.HTMLElement) {
				strongText := e.Text
				text = strings.Replace(text, strongText, "", -1)
				strongArray = append(strongArray, strongText)
			})
			maintext = maintext + text + "\n"
		})

		data := []byte(maintext)
		ioutil.WriteFile("currency.txt", data, 0777)

		fmt.Println("---------------------------------------\n")
		fmt.Println(maintext)
		fmt.Println("\n---------------------------------------")
	})

	c.Visit("https://www.kbzbank.com/en/")
}
