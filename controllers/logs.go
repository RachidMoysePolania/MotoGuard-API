package controllers

import (
	"net/http"

	"github.com/RachidMoysePolania/MotoGuard-API/models"
	"github.com/gin-gonic/gin"
)

func SaveLogs(c *gin.Context) {
	var log models.Road_logs
	if err := c.ShouldBindJSON(&log); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error saving log"})
		return
	}

	Db.Create(&log)
	c.JSON(http.StatusOK, gin.H{"data": "Log creado correctamente!"})
}

func GetLogById(c *gin.Context) {
	logid := c.Param("id")
	var dblog models.Road_logs
	Db.Model(models.Userdata{}).Where("id = ?", logid).Find(&dblog)
	c.JSON(http.StatusOK, gin.H{"data": dblog})
}

func GetallLogs(c *gin.Context) {
	var logs []models.Road_logs
	//Db.Find(&users)
	Db.Model(&models.Road_logs{}).Find(&logs)
	c.JSON(http.StatusOK, gin.H{"data": logs})
}
