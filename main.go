package main

import (
	"bytes"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"
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

func putRequest(url string, data io.Reader) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, url, data)
	req.Header.Add("Content-Type", "binary/octet-stream")

	if err != nil {
		panic(err)
	}

	return client.Do(req)
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

		payment := map[string]string{
			"key":         "gtKFFx",
			"txnid":       txnid,
			"amount":      amount,
			"productinfo": "ZOOMIN",
			"firstname":   firstname,
			"email":       email,
			"phone":       phone,
			"surl":        "https://toalety.herokuapp.com/payu/success",
			"furl":        "https://toalety.herokuapp.com/payu/failed",
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

	router.POST("/upload", func(c *gin.Context) {
		r := c.Request

		url := r.FormValue("url")

		file, _, _ := r.FormFile("file")
		defer file.Close()

		buf := bytes.NewBuffer(nil)
		if _, err := io.Copy(buf, file); err != nil {
			panic(err)
		}

		fmt.Println(url)

		resp, err := putRequest(url, buf)

		if err != nil {
			panic(err)
		}

		c.JSON(
			resp.StatusCode,
			gin.H{
				"status": "OK",
			},
		)
	})

	router.Run()
}
