package dbchgkinfo

import (
	"errors"
	"net/http"
	"qask/internal/app/model"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//SiteRepository ...
type SiteRepository struct {
	sitename            string
	validAnswerRegexp   *regexp.Regexp
	validCommentRegexp  *regexp.Regexp
	validQuestionRegexp *regexp.Regexp
}

//New ...
func New() *SiteRepository {
	return &SiteRepository{
		sitename:            "db.chgk.info",
		validAnswerRegexp:   regexp.MustCompile("^\n\x20{4}Ответ: (.*)$"),
		validCommentRegexp:  regexp.MustCompile("^\n\x20{4}Комментарий: (.*)$"),
		validQuestionRegexp: regexp.MustCompile("^Вопрос 1:$"),
	}
}

func escape(s string) string {
	str1 := strings.Replace(s, "\n", " ", -1)
	str2 := strings.Replace(str1, "\\", "", -1)
	return str2
}

func parseQuestion(s string, regex *regexp.Regexp) {
}

func parseAnswer(s string, regex *regexp.Regexp) (string, error) {
	if !regex.MatchString(s) {
		return "", errors.New("Not matched")
	}

	answers := regex.FindAllStringSubmatch(s, 1)
	if len(answers[0]) == 2 {
		return answers[0][1], nil
	}

	return "", errors.New("Something went wrong")
}

func parseComment(s string, regex *regexp.Regexp) (string, error) {
	if !regex.MatchString(s) {
		return "", errors.New("Not matched")
	}

	comments := regex.FindAllStringSubmatch(s, 1)
	if len(comments[0]) == 2 {
		return comments[0][1], nil
	}

	return "", errors.New("Something went wrong")
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

	doc.Find(".random_question").Each(func(i int, selection *goquery.Selection) {
		selection.Contents().Each(func(i int, selection *goquery.Selection) {
			if i == 6 {
				q.Question = escape(selection.Text())
			} else if s.validAnswerRegexp.MatchString(selection.Text()) {
				answer, err := parseAnswer(selection.Text(), s.validAnswerRegexp)
				if err == nil {
					q.Answer = escape(answer)
				}
			} else if s.validCommentRegexp.MatchString(selection.Text()) {
				comment, err := parseComment(selection.Text(), s.validCommentRegexp)
				if err == nil {
					q.Comment = escape(comment)
				}
			}
		})
	})

	if q.Question == "" || q.Answer == "" {
		return nil, ErrGetInvalidQuestion
	}

	return q, nil
}
