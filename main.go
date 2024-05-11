package yurls

import (
	"fmt"
	"net/url"
	"regexp"
	"slices"
)

var re = regexp.MustCompile(`((https?|ftp|file)://)?[-A-Za-z0-9+&@#%?=~_|!:,;]+(\.[-A-Za-z0-9+&@#%?=~_|!:,;]+)*\.([A-Za-z]){2,}[-A-Za-z0-9+&@#%/=~_|!:,;]*`)

func Parse(content string) []string {
	findUrls := re.FindAllString(content, -1)
	fmt.Println(findUrls)
	invalidUrls := make([]string, 0)
	for _, urlText := range findUrls {
		_, err := url.Parse(urlText)
		if err != nil {
			invalidUrls = append(invalidUrls, urlText)
		}
	}

	validUrls := make([]string, 0)
	for _, urlText := range findUrls {
		if !slices.Contains(invalidUrls, urlText) {
			validUrls = append(validUrls, urlText)
		}
	}
	return validUrls
}
