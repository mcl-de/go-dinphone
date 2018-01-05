package dinphone

import (
	"github.com/glenn-brown/golang-pkg-pcre/src/pkg/pcre"
	"regexp"
	"strings"
	"fmt"
)

var phoneRegex = pcre.MustCompile(`^\s*((?:(?:\+|00)(?:(?!39|43)[1-9][0-9]{1,3}[\s.\/([-]*(?:[([]\s*0\s*[)\]])?\s*[1-9][0-9]{1,}|39[\s.\/([-]*[0-9]{1,}|43[\s.\/([-]*[1-9][0-9]{0,})\s*[)\]]?)|(?:[([]?(?:0)[1-9][0-9]{0,}[)\]]?))[\s.\/([-]*(.*)$`, 0)
var removeBracketsRegex = regexp.MustCompile(`(?:[(\[]\s*0\s*[)\]])`)
var leftPartSplitRegex = regexp.MustCompile(`[\s.\/-]+`)
var rightPartMatchRegex = regexp.MustCompile(`^(.*)[^\d](\d+)$`)
var replaceNonNumbersRegex = regexp.MustCompile(`[^+\d]`)
var startingNumberRegex = regexp.MustCompile(`^[1-9]\d+`)
var someNumberRegex = regexp.MustCompile(`^[789]00|180[1-7]$`)

func Parse(numberToParse string) (din string) {
	matcher := phoneRegex.MatcherString(numberToParse, 0)
	matches := make([]string, matcher.Groups()+1)
	for i := 1; i <= matcher.Groups(); i++ {
		matches[i-1] = strings.TrimSpace(matcher.GroupString(i))
	}

	if len(matches) < 2 {
		return
	}

	matches[0] = removeBracketsRegex.ReplaceAllString(matches[0], " ")
	leftParts := leftPartSplitRegex.Split(matches[0], -1)
	if len(leftParts) == 0 {
		return
	}

	leftParts[0] = replaceNonNumbersRegex.ReplaceAllString(leftParts[0], "")
	if len(leftParts) > 1 && leftParts[1] != "" {
		leftParts[1] = replaceNonNumbersRegex.ReplaceAllString(leftParts[1], "")
	} else {
		if strings.HasPrefix(leftParts[0], "0") {
			leftParts[0] = leftParts[0][1:]
		}
		leftParts = []string{"+49", leftParts[0]}
	}
	if strings.HasPrefix(leftParts[0], "00") {
		leftParts[0] = "+" + leftParts[0][2:]
	}

	rightParts := rightPartMatchRegex.FindStringSubmatch(matches[1])
	if len(rightParts) > 0 {
		rightParts = rightParts[1:]
		for i := range rightParts {
			rightParts[i] = strings.TrimSpace(rightParts[i])
		}
	} else {
		rightParts = []string{matches[1]}
	}

	if startingNumberRegex.MatchString(rightParts[0]) {
		rightParts[0] = replaceNonNumbersRegex.ReplaceAllString(rightParts[0], "")

		if leftParts[0] == "+49" && someNumberRegex.MatchString(leftParts[1]) {
			din = fmt.Sprintf("0%s %s", leftParts[1], rightParts[0])
		} else {
			din = fmt.Sprintf("%s %s %s", leftParts[0], leftParts[1], rightParts[0])
		}

		if len(rightParts) > 1 {
			din += "-"+ rightParts[1]
		}
	}

	return
}