package core

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/module/apmhttp"
)

func GetMoodle(ctx *gin.Context, email string) []byte {

	url := "https://200.14.84.16/api/v1/moodle/getCourses"
	method := "GET"
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	//	client := &http.Client{}
	client := apmhttp.WrapClient(&http.Client{Transport: tr})

	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return []byte("[]")
	}
	// tx := middleware.GetTransaction(ctx.Request)
	// middleware.SetTransaction(req, tx)
	req.Host = "external-udp.xn--ensea-rta.cl"
	req.Header.Add("Email", email)

	//	res, err := client.Do(req)
	res, err := client.Do(req.WithContext(ctx.Request.Context()))

	if err != nil {
		fmt.Println(err)
		return []byte("[]")
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return []byte("[]")
	}

	return body
}
