package problems_ru

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"qask/internal/app/model"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html/charset"
)

type SiteRepository struct {
	sitename string
}

func New() *SiteRepository {
	return &SiteRepository{
		sitename: "problems.ru",
	}
}

func convertKOI8RToUTF8(r *http.Response) (io.Reader, error) {
	utf8, err := charset.NewReader(r.Body, r.Header.Get("Content-Type"))
	if err != nil {
		return nil, err
	}

	return utf8, nil
}

func getProblemID() (string, error) {
	params := []byte(`mode=1&problems_count=1`)
	resp, err := http.Post("http://www.problems.ru/view_random.php?mode=1", "application/x-www-form-urlencoded", bytes.NewBuffer(params))
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}

	mp_id := ""
	mp_id_exists := false

	doc.Find("[name=id0]").Each(func(i int, s *goquery.Selection) {
		mp_id, mp_id_exists = s.Attr("value")
	})

	if !mp_id_exists {
		return "", ErrCouldNotGetProblemID
	}

	return mp_id, nil
}

//GetQuestion ...
func (s *SiteRepository) GetQuestion() (interface{}, error) {
	q := model.Question{
		Sitename: s.sitename,
	}

	mp_id, err := getProblemID()
	if err != nil {
		if err == ErrCouldNotGetProblemID {
			return nil, ErrCouldNotGetProblem
		}

		return nil, err
	}

	params := fmt.Sprintf("mode=1&problems_count=1&id0=%s&problem_ids_count=1&action_solutions=%%F0%CF%CB%C1%DA%C1%D4%D8+%D2%C5%DB%C5%CE%C9%D1", mp_id)
	resp, err := http.Post("http://www.problems.ru/view_random.php?mode=1", "application/x-www-form-urlencoded", bytes.NewBuffer([]byte(params)))
	if err != nil {
		return nil, err
	}

	body, err := convertKOI8RToUTF8(resp)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return nil, err
	}

	doc.Find(".componentboxcontents").Each(func(i int, s *goquery.Selection) {
		s.Find("p").Each(func(i int, s *goquery.Selection) {
			if i == 1 {
				q.Question = s.Text()
			} else if i == 5 {
				q.Answer = s.Text()
			}
		})
	})

	if q.Question == "" || q.Answer == "" {
		return nil, ErrGetInvalidProblem
	}
	return q, nil
}
