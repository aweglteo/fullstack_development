package main

import (
	"context"
	pb "github.com/aweglteo/fullstack_development/api/scraping"
	"github.com/PuerkitoBio/goquery"
	"github.com/saintfish/chardet"
	"golang.org/x/net/html/charset"
)

type ScrapingService struct {
}

func (s* ScrapingService) GetRestaurantInfo(ctx context.Context, message *pb.GetScrapingRequest) (*pb.ScrapingResponse, error) {
	tabelogUrl := message.TargetLink

	res, _ := http.Get(tabelogUrl)
	defer res.Body.Close()

	buf, _ := ioutil.ReadAll(res.Body)

	det := chardet.NewTextDetector()
	detRslt, _ := det.DetectBest(buf)

	bReader := bytes.NewReader(buf)
	reader, _ := charset.NewReaderLabel(detRslt.Charset, bReader)

	doc, _ := goquery.NewDocumentFromReader(reader)

	tel := doc.Find("p.rstinfo-table__tel-num-wrap > strong")

	return &pb.ScrapingResponse{
		Name: "sample restaurant name",
		Address: "sample Address 1",
		Tell: "sample tell number",
		OpeningHours: "11 hours",
	}, nil
}
