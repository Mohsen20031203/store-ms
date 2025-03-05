package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"

	emailverifier "github.com/AfterShip/email-verifier"
)

func checkEnglishChar(input string) bool {
	for _, r := range input {
		if !unicode.In(r, unicode.Latin) && !('0' <= r && r <= '9') {
			return false
		}
	}
	return true
}

func convertPersianNumber(input string) string {

	persianToEnglishMap := map[rune]rune{
		'۰': '0', '۱': '1', '۲': '2', '۳': '3', '۴': '4', '۵': '5', '۶': '6',
		'۷': '7', '۸': '8', '۹': '9',
	}

	var builder strings.Builder
	builder.Grow(len(input))

	for _, char := range input {

		if englishChar, exists := persianToEnglishMap[char]; exists {
			builder.WriteRune(englishChar)
		} else {
			builder.WriteRune(char)
		}
	}

	return builder.String()
}

func isEmailValid(e string) bool {

	verifier := emailverifier.NewVerifier()

	re := regexp.MustCompile("^[a-zA-Z0-9.]+$")
	result, _ := verifier.Verify(e)
	username := re.MatchString(result.Syntax.Username)

	if !result.Free || !username {
		fmt.Println("email address syntax is invalid")
		return false
	}
	return true

}

func main() {
	fmt.Println(checkEnglishChar("ms;dfj"))

	fmt.Println(convertPersianNumber("ahdfhaنشسیاب۳۱۹۴ اتا"))

	fmt.Println(isEmailValid("mohsssen.sal23ehi2070@fdfd."))
}
