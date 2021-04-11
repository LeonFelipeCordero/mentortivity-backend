package main

import (
	"log"
	"net/http"
	"time"

	"github.com/LeonFelipeCordero/mentortivity-backend/cron"
	"github.com/LeonFelipeCordero/mentortivity-backend/firestore"
	"github.com/LeonFelipeCordero/mentortivity-backend/notebook"
	"github.com/LeonFelipeCordero/mentortivity-backend/report"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "up", "timestamp": time.Now()})
	})

	r.GET("/ready", func(c *gin.Context) {
		client := firestore.NewFirestoreClient()
		client.Close()
		c.JSON(http.StatusOK, gin.H{"status": "ready", "timestamp": time.Now()})
	})

	r.PUT("/notebook/:id/close", func(c *gin.Context) {
		id := c.Param("id")
		client := firestore.NewFirestoreClient()
		reportGenerator := report.DefaultReportGenerator{}
		report := notebook.CloseDay(*client, &reportGenerator, id)
		client.Close()
		c.JSON(http.StatusOK, report)
	})

	return r
}

func main() {
	loadEnv()

	cron.ScheduleJobs()
	r := setupRouter()
	r.Run(":8082")
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading environment variables")
	}
}
