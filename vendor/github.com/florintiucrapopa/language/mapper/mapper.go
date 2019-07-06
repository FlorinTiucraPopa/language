package mapper

import "github.com/rylans/getlang"

func Greet(s string) string {
	info := getlang.FromString(s)
	switch info.LanguageCode() {
	case "en":
		return "Hello!"
	case "fr":
		return "Bonjour!"
	case "de":
		return "Guten tag!"
	default:
		return "Don't know the language!"
	}
}