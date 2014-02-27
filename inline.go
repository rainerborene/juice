package juice

import (
	"github.com/PuerkitoBio/goquery"
	"io"
)

func Inline(source io.Reader, rules Rules) (string, error) {
	doc, err := goquery.NewDocumentFromReader(source)

	if err != nil {
		panic(err.Error())
	}

	for index := range rules {
		rule := rules[index]
		query := doc.Find(rule.Selector)

		if query.Length() == 0 {
			continue
		}

		query.Each(func(i int, s *goquery.Selection) {
			var style string

			if styleValue, exists := s.Attr("style"); exists {
				style = rule.Style() + " " + styleValue
			} else {
				style = rule.Style()
			}

			setAttributeValue("style", style, s.Get(i))
		})
	}

	return doc.Html()
}
