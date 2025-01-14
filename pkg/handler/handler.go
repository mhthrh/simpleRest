package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simpleAPI/pkg/model"
)

func Handler(c *gin.Context) {
	result := []model.App{
		{
			Id: 0, Name: "John"},
		{
			Id: 1, Name: "Paul",
		}}
	c.IndentedJSON(http.StatusOK, result)
}
