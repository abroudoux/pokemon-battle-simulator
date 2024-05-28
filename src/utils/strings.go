package utils

import (
	"net/http"
	"strconv"
	"unicode"

	"github.com/gin-gonic/gin"
)

func CheckString(c *gin.Context) {
	s := c.Param("s")

	if isDigit(s) {
		sTransformed, err := strconv.Atoi(s)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Conversion error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"int": sTransformed})
		return
	}

	if isLetter(s) {
		c.JSON(http.StatusOK, gin.H{"message": "It's a string"})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"error": "The input contains both letters and digits"})
}

func isDigit(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

func isLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}