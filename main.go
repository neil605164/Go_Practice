package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/singleflight"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	var singleFlight singleflight.Group

	r.GET("/ping", func(c *gin.Context) {

		v, err, shared := singleFlight.Do("github", func() (interface{}, error) {
			return githubStatus()
		})
		if err != nil {
			fmt.Println("Error Msg ===>", err)
			return
		}

		status := v.(string)

		c.JSON(200, gin.H{
			"status": status,
			"shared": shared,
		})
	})
	r.Run(":8787")
}

func githubStatus() (string, error) {
	log.Println("Make Request to Github API")
	defer log.Println("Request to GitHub API Complete")

	resp, err := http.Get("https://status.github.com/api/status.json")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("github response: %s", resp.Status)
	}

	status := strconv.Itoa(resp.StatusCode)

	return status, err
}
