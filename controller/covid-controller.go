package controller

import (
	"corona-time/entity"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type CovidController interface {
	ShowForToday(ctx *gin.Context)
	ShowHistory(ctx *gin.Context)
}

type covidController struct {
}

func New() CovidController {
	return &covidController{}
}

func (controller *covidController) ShowHistory(ctx *gin.Context) {
	var covids []entity.Covid

	url := "https://api.covidtracking.com/v1/us/daily.json"

	spaceClient := http.Client{
		Timeout: time.Second * 5, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	json.Unmarshal([]byte(body), &covids)

	data := gin.H{
		"covids": covids,
	}
	ctx.HTML(http.StatusOK, "showAll.html", data)
}

func (controller *covidController) ShowForToday(ctx *gin.Context) {

	url := "https://api.covidtracking.com/v1/us/current.json?fbclid=IwAR3NOvX2nr2lnH5ntWxIgQzlBQ_4hMq-EuS0Vd6w3kaPh_9e6aJLVXtIn0A"

	spaceClient := http.Client{
		Timeout: time.Second * 5, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	slic := string(body)[1 : len(body)-2]

	var result entity.Covid

	json.Unmarshal([]byte(slic), &result)

	data := gin.H{
		"covid": result,
	}
	ctx.HTML(http.StatusOK, "index.html", data)
}
