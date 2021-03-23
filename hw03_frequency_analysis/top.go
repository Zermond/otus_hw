package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

type Words struct {
	Name  string
	Count int
}

var re = regexp.MustCompile(`\s`)

func Top10(s string) []string {
	if len(s) == 0 {
		return []string{}
	}

	s = re.ReplaceAllString(s, " ")

	arrayText := strings.Split(s, " ")

	workMap := make(map[string]int)

	for _, val := range arrayText {
		if val == "" {
			continue
		}
		_, ok := workMap[val]
		if !ok {
			workMap[val] = 1
			continue
		}
		valMap, _ := workMap[val] //nolint

		workMap[val] = valMap + 1
	}

	words := []Words{}

	for word := range workMap {
		val := workMap[word]
		insert := Words{Count: val, Name: word}
		words = append(words, insert)
	}

	sort.Slice(words, func(i, j int) bool {
		if words[i].Count == words[j].Count {
			return words[i].Name < words[j].Name
		}
		return words[i].Count > words[j].Count
	})

	result := make([]string, 0, len(words))

	for _, word := range words {
		result = append(result, word.Name)
	}

	if len(result) < 10 {
		return result[0:]
	}
	return result[0:10]
}
