package scrape

import (
	"log"

	"github.com/anaskhan96/soup"
)

func grabHTML(url string) string {
	resp, err := soup.Get(url)
	if err != nil {
		log.Printf("Failed HTTP GET w/ error: %v", err)
		panic(err)
	}

	// doc := soup.HTMLParse(resp)
	// links := doc.Find("div", "id", "comicLinks").FindAll("a")
	// for _, link := range links {
	// 	fmt.Println(link.Text(), "| Link :", link.Attrs()["href"])
	// }
	return resp
}
