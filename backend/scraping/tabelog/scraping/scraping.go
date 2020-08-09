package scraping

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	pb "github.com/aweglteo/fullstack_development/api/scraping"
	"github.com/saintfish/chardet"
	"golang.org/x/net/html/charset"
)

type Server struct {
}

func (s *Server) GetShopInfo(ctx context.Context, message *pb.GetScrapingRequest) (*pb.ScrapingResponse, error) {
	tabelogUrl := message.TargetLink

	res, _ := http.Get(tabelogUrl)
	defer res.Body.Close()

	buf, _ := ioutil.ReadAll(res.Body)

	det := chardet.NewTextDetector()
	detRslt, _ := det.DetectBest(buf)

	bReader := bytes.NewReader(buf)
	reader, _ := charset.NewReaderLabel(detRslt.Charset, bReader)

	doc, _ := goquery.NewDocumentFromReader(reader)

	tel := doc.Find("p.rstinfo-table__tel-num-wrap > strong").Text()

	return &pb.ScrapingResponse{
		Name:         "sample restaurant name",
		Address:      "sample Address 1",
		Tell:         tel,
		OpeningHours: "11 hours",
	}, nil
}
