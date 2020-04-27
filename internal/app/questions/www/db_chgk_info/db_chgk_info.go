package db_chgk_info

import (
	"net/http"
	"qask/internal/app/model"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type SiteRepository struct {
	sitename string
}

func New() *SiteRepository {
	return &SiteRepository{
		sitename: "db.chgk.info",
	}
}

func escape(s string) string {
	str1 := strings.Replace(s, "\n", " ", -1)
	str2 := strings.Replace(str1, "\\", "", -1)
	return str2
}

//GetQuestion ...
func (s *SiteRepository) GetQuestion() (interface{}, error) {
	resp, err := http.Get("https://db.chgk.info/random/answers/limit1")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	q := model.Question{
		Sitename: s.sitename,
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	doc.Find(".random_question").Each(func(i int, s *goquery.Selection) {
		s.Contents().Each(func(i int, s *goquery.Selection) {
			if i == 6 {
				q.Question = escape(s.Text())
			}
			if i == 7 {
				q.Answer = escape(s.Text())
			}
		})
	})

	if q.Question == "" || q.Answer == "" {
		return nil, ErrGetInvalidQuestion
	}

	return q, nil
}
