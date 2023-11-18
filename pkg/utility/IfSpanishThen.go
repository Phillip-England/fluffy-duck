package utility

import "github.com/gin-gonic/gin"

func IfSpanishThen(c *gin.Context, englishVersion string, spanishVersion string) string {
	language := GetLanguage(c)
	if language == "english" || language == "" {
		return englishVersion
	} else {
		return spanishVersion
	}
}