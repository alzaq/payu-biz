package main

import (
	"crypto/sha512"
	"encoding/hex"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func preparePayuHash(payment map[string]string) string {
	formula := "key|txnid|amount|productinfo|firstname|email|||||||||||salt"

	for key, value := range payment {
		formula = strings.Replace(formula, key, value, -1)
	}

	sha := sha512.New()
	sha.Write([]byte(formula))

	hash := hex.EncodeToString(sha.Sum(nil))

	return hash
}

func main() {
	router := gin.Default()
	router.LoadHTMLFiles("templates/index.tmpl")

	router.GET("/payu", func(c *gin.Context) {
		txnid := c.Query("txnid")
		amount := c.Query("amount")
		firstname := c.Query("firstname")
		email := c.Query("email")
		phone := c.Query("phone")
		surl := c.Query("surl")
		furl := c.Query("furl")

		payment := map[string]string{
			"key":         "gtKFFx",
			"txnid":       txnid,
			"amount":      amount,
			"productinfo": "ZOOMIN",
			"firstname":   firstname,
			"email":       email,
			"phone":       phone,
			"surl":        surl,
			"furl":        furl,
			"salt":        "eCwWELxi",
		}

		hash := preparePayuHash(payment)

		c.HTML(
			http.StatusOK,
			"index.tmpl",
			gin.H{
				"Context": c,
				"Payu":    "https://test.payu.in/_payment",
				"Payment": payment,
				"Hash":    hash,
			},
		)
	})

	router.GET("/payu/success", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"status": "success",
			},
		)
	})

	router.GET("/payu/failed", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"status": "failed",
			},
		)
	})

	router.Run()
}
