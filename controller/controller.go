package controller

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Data   map[string]interface{} `json:"data"`
	Status string                 `json:"status"`
}

func GetRequestData(c *gin.Context) ResponseData {
	requestData := make(map[string]interface{})
	queryParams := c.Request.URL.Query()

	if err := c.BindJSON(&requestData); err != nil {
		log.Println("Error binding request data JSON", err)
	}

	for key, values := range queryParams {
		requestData[key] = values[0]
	}

	status := requestData["status"]

	if status != nil {
		return ResponseData{
			Data:   requestData,
			Status: status.(string),
		}
	} else {
		return ResponseData{
			Data:   requestData,
			Status: "200",
		}
	}
}

func TestHandler(c *gin.Context) {
	responseData := GetRequestData(c)
	responseStatus, err := strconv.Atoi(responseData.Status)
	if err != nil {
		log.Println("Error converting status from string to int", err)
		return
	}
	c.JSON(responseStatus, responseData.Data)
}
