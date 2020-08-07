package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/saintfish/chardet"
	"golang.org/x/net/html/charset"
)

// 食べログのURLがぶち込まれたら、スクレイピングして情報をDBに保存する。
func scraping(url string) {
	res, _ := http.Get(url)
	defer res.Body.Close()

	buf, _ := ioutil.ReadAll(res.Body)

	det := chardet.NewTextDetector()
	detRslt, _ := det.DetectBest(buf)

	bReader := bytes.NewReader(buf)
	reader, _ := charset.NewReaderLabel(detRslt.Charset, bReader)

	doc, _ := goquery.NewDocumentFromReader(reader)

	tel := doc.Find("p.rstinfo-table__tel-num-wrap > strong")
	// address := doc.Find("p.rstinfo-table__address > span").Text()
	if 
	fmt.Println(tel)
}

func main() {
	url := "https://tabelog.com/tokyo/A1310/A131002/13181683/"
	scraping(url)
}
