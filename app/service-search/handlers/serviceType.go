package handlers

import (
	serviceQuerie "github.com/Bbarbosa7/ineed-full/app/service-search/queries"
	"github.com/gin-gonic/gin"
)

func GetServiceTypesHandler(c *gin.Context)  {
	c.JSON(200, gin.H{"message": serviceQuerie.GetServiceType()})
}