package controller

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Data   map[string]interface{} `json:"data"`
	Status string                 `json:"status"`
}

func GetRequestData(c *gin.Context) Response {
	requestData := make(map[string]interface{})
	queryParams := c.Request.URL.Query()

	if err := c.BindJSON(&requestData); err != nil {
		log.Println("Error binding request data JSON", err)
	}

	for key, values := range queryParams {
		requestData[key] = values[0]
	}

	var status string

	if requestData["status"] != nil {
		status = requestData["status"].(string)
	} else {
		status = "200"
	}

	return Response{
		Data:   requestData,
		Status: status,
	}
}

func TestHandler(c *gin.Context) {
	response := GetRequestData(c)
	responseStatus, err := strconv.Atoi(response.Status)
	if err != nil {
		log.Println("Error converting status from string to int", err)
		return
	}
	c.JSON(responseStatus, response.Data)
}
