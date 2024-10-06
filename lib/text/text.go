package text

import (
	"regexp"
	"strings"
)

// Slugify converts a string into a URL-friendly slug.
func Slugify(title string) string {
	// convert to lowercase
	title = strings.ToLower(title)
	// strip weird characters
	puncRegex := regexp.MustCompile("[^\\w\\d\\s\\-]")
	title = puncRegex.ReplaceAllString(title, "")
	// replace whitespace to single dash
	wsRegex := regexp.MustCompile("\\s+")
	title = wsRegex.ReplaceAllString(title, "-")
	dashRegex := regexp.MustCompile("-+")
	title = dashRegex.ReplaceAllString(title, "-")
	return title
}
