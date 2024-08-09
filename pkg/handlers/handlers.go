package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"qrCode/pkg/database"
	"time"
)

func Generate(c *gin.Context) {
	// Get the name from the POST request
	name := c.PostForm("name")

	// Log the name or do something with it
	fmt.Println("Name: ", name)

	// validate the name of the site is within parameters
	// if !valid(name)
	// 	c.JSON(400, gin.H{
	// 		"error": "Invalid name",
	// 	})
	// 	return

	// create a new shortened URL and verify it is unique
	newURL := Shrink()

	myURL := fmt.Sprintf("http://localhost:8080/qr/%s", newURL)
	fmt.Println("Generated URL: ", myURL)
	// Create a new URL (as a placeholder for now)
	err := database.InsertSite("user", newURL, name)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	//Generate QR code
	//qrCode, err := qrcodeCreator.GenerateQRCode(myURL)

	// Return the new URL in the response
	c.JSON(200, gin.H{
		"url": myURL,
	})
}

// link shrinker function
func Shrink() string {
	rand.Seed(time.Now().UnixNano())
	// Create a new random string
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, 6)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	// Check if the string is unique
	if !database.IsUnique(string(b)) {
		return Shrink()
	}
	// Return the string
	return string(b)
}

func GO(c *gin.Context) {
	// Get the name from the POST request
	site := c.Param("site")
	if site == "" {
		c.JSON(400, gin.H{
			"error": "No site provided",
		})
		return
	}

	// search db for site
	// if site exists, return the shortened url
	siteStruct, err := database.GetSite(site)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	// increment the site count
	database.IncrementSite(&siteStruct)
	// redirect to the site
	c.Redirect(301, siteStruct.SiteName)

	// check db for the site
	// if site exists, return the shortened url

}
