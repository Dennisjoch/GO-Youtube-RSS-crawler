package main

import (
    "fmt"
    "strings"
    "github.com/PuerkitoBio/goquery"
)

func main() {
    doc, err := goquery.NewDocument("https://www.youtube.com/channel/UCzH549YlZhdhIqhtvz7XHmQ")
    if err != nil {
        panic("meh")
    }

    doc.Find("link").Each(func(i int, s *goquery.Selection) {
        if name, _ := s.Attr("title"); strings.EqualFold(name, "RSS") {
            description, _ := s.Attr("href")
            fmt.Println(description)
        }
    })
}