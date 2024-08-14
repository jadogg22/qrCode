package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"qrCode/pkg/auth"
	"qrCode/pkg/database"
	"time"
)

func Login(c *gin.Context) {
	// Get the username and password from the POST request
	username := c.PostForm("username")
	password := c.PostForm("password")

	fmt.Println("Username: ", username)
	fmt.Println("Password: ", password)

	// Check if the username and password are correct
	err := auth.CheckPassword(username, password)
	if err != nil {
		c.JSON(401, gin.H{
			"error": "Invalid username or password",
		})
		return
	}

	// create a new jwt token
	token, err := auth.CreateToken(username)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.SetCookie("token", token, 3600, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Logged in successfully"})
}

func Register(c *gin.Context) {
	// Get the username, password and email from the POST request
	username := c.PostForm("username")
	password := c.PostForm("password")
	email := c.PostForm("email")

	fmt.Println("Username: ", username)
	fmt.Println("Password: ", password)
	fmt.Println("Email: ", email)

	// Check if the username and email are unique
	err := database.UserExists(username, email)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Hash the password
	hashedPassword, salt, err := auth.SavePassword(password)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Save the user to the database
	err = database.AddUser(username, hashedPassword, salt, email)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := auth.CreateToken(username)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.SetCookie("token", token, 3600, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Logged in successfully"})

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

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

type MySites struct {
	Site       string `json:"site"`
	Count      int    `json:"count"`
	CountToday int    `json:"count_today"`
}

type User struct {
	Name  string    `json:"name"`
	Sites []MySites `json:"sites"`
}

func GetUser(c *gin.Context) {
	// get the name from the GET request
	user := c.Param("user")
	if user == "" {
		c.JSON(400, gin.H{
			"error": "No user provided",
		})
		return
	}

	// search db for user
	// if user exists, return the user's sites stats
	userStruct, err := database.GetSitesByUser(user)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	// return the user's sites stats
	c.JSON(200, userStruct)
}

func GetSites(c *gin.Context) {
	// Get the username from the context
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Convert username to string
	usernameStr, ok := username.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username"})
		return
	}

	sites, err := database.GetSitesByUser(usernameStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to retrive sites" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, sites)
}
