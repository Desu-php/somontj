package main

import (
	"fmt"
	"github.com/Desu-php/somontj/stats"
	"net/http"
	"os"
	"regexp"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

const (
	BaseURL = "https://somon.tj/nedvizhimost/prodazha-kvartir/hudzhand/"
)

func main() {
	stats.Run()
	//pageLength, err := getPageLength()
	//
	//if err != nil {
	//	fmt.Println("PageLength getting err", err)
	//	os.Exit(1)
	//}
	//
	//page := 1
	//
	//for page <= pageLength {
	//	currentPage := strconv.Itoa(page)
	//	fmt.Println("Обработка страницы: " + currentPage)
	//
	//	err := fetchApartmentsByPage(page)
	//
	//	if err != nil {
	//		fmt.Printf("ошибка обработки страницы %d: %w", page, err)
	//		os.Exit(1)
	//	}
	//
	//	page++
	//}
}

func fetchApartmentsByPage(page int) error {
	document, err := request(fmt.Sprintf("%s?page=%d", BaseURL, page))

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	doc := document.Find(".list-announcement-assortiments .list-simple__output .advert")

	urls := getUrls(doc)

	for _, val := range urls {
		id, err := parseIDFromURL(val)

		if err != nil {
			return fmt.Errorf("err when parse id from url %s, %w", val, err)

		}

		parsedId, err := strconv.ParseUint(id, 10, 32)

		if err != nil {
			return fmt.Errorf("err when parse id from string %s %w", val, err)
		}

		fmt.Println(id)

		apartment, err := GetDetailsById(uint(parsedId))

		if err != nil {
			return fmt.Errorf("GetDetailsById: %s %w", id, err)
		}

		err = Save(*apartment, "apartments.json")
		if err != nil {
			return err
		}
	}

	return nil
}

func getPageLength() (int, error) {
	document, err := request(BaseURL)

	if err != nil {
		return 0, err
	}

	page := document.Find(".number-list li a").Last().Text()

	return strconv.Atoi(page)
}

func request(url string) (*goquery.Document, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}

	defer resp.Body.Close()

	return goquery.NewDocumentFromReader(resp.Body)
}

func getUrls(documents *goquery.Selection) []string {
	var urls []string

	documents.Each(func(i int, s *goquery.Selection) {
		a := s.Find(".advert__content-title")
		href, _ := a.Attr("href")
		fmt.Println(href, i)
		urls = append(urls, href)
	})

	return urls
}

func parseIDFromURL(url string) (string, error) {
	// Регулярное выражение для извлечения ID из URL
	re := regexp.MustCompile(`adv/(\d+)_`)
	matches := re.FindStringSubmatch(url)
	if len(matches) < 2 {
		return "", fmt.Errorf("ID не найден в URL")
	}
	return matches[1], nil
}
