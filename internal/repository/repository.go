package repository

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/Reteger/shiki/internal/models"
)

type Repository interface {
	GetForDay(day string) ([]models.Ongoing, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetForDay(day string) ([]models.Ongoing, error) {
	resp, err := http.Get("https://shikimori.one/ongoings")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	upperDay := strings.ToUpper(day)
	var results []models.Ongoing
	inTargetDay := false

	doc.Find("body *").Each(func(i int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Text())
		up := strings.ToUpper(text)

		if !inTargetDay && strings.HasPrefix(up, upperDay) {
			inTargetDay = true
			return
		}

		if inTargetDay {
			if isAnyWeekday(up) && !strings.HasPrefix(up, upperDay) {
				inTargetDay = false
				return
			}

			if s.Is("a") {
				href, ok := s.Attr("href")
				if ok &&
					strings.Contains(href, "/animes/") &&
					!strings.Contains(href, "/animes/kind/") &&
					!strings.Contains(href, "/animes/season/") &&
					!strings.Contains(href, "/animes/genre/") &&
					!strings.Contains(href, "/animes/studio/") {

					rawTitle := strings.TrimSpace(s.Text())
					if rawTitle != "" {
						original, russian := splitTitles(rawTitle)

						link := href
						if !strings.HasPrefix(href, "http") {
							link = "https://shikimori.one" + href
						}

						results = append(results, models.Ongoing{
							OriginalTitle: original,
							RussianTitle:  russian,
							Link:          link,
						})
					}
				}
			}
		}
	})

	return results, nil
}

func isAnyWeekday(up string) bool {
	return strings.Contains(up, "ПОНЕДЕЛЬНИК") ||
		strings.Contains(up, "ВТОРНИК") ||
		strings.Contains(up, "СРЕДА") ||
		strings.Contains(up, "ЧЕТВЕРГ") ||
		strings.Contains(up, "ПЯТНИЦА") ||
		strings.Contains(up, "СУББОТА") ||
		strings.Contains(up, "ВОСКРЕСЕНЬЕ")
}

var cyrillicRe = regexp.MustCompile(`[А-Яа-яЁё]`)

func splitTitles(raw string) (string, string) {
	raw = strings.TrimSpace(raw)

	loc := cyrillicRe.FindStringIndex(raw)
	if loc == nil {

		return raw, ""
	}

	idx := loc[0]
	original := strings.TrimSpace(raw[:idx])
	russian := strings.TrimSpace(raw[idx:])

	return original, russian
}
