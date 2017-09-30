package main

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func BasicAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user User
		username, password, _ := c.Request.BasicAuth()
		if db.First(&user, "username = ?", username).RecordNotFound() {
			c.Header("WWW-Authenticate", "U HAVE TREAD UPON MY DOMAIN & MUST SUFFER. WHO R U?")
			c.AbortWithStatus(401)
			return
		}
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err != nil {
			c.Header("WWW-Authenticate", "U HAVE TREAD UPON MY DOMAIN & MUST SUFFER. WHO R U?")
			c.AbortWithStatus(401)
			return
		}
		c.Set(gin.AuthUserKey, user)
	}
}

func main() {

	router := gin.Default()

	router.POST("/login", func(c *gin.Context) {
		var user User
		username := c.PostForm("username")
		password := c.PostForm("password")

		if db.First(&user, "username = ?", username).RecordNotFound() {
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
			if err != nil {
				panic(err)
			}
			db.Create(&User{Username: username, Password: string(hashedPassword)})
			c.Status(200)
			return
		}
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err != nil {
			c.Status(400)
			return
		}
		c.Status(200)
		return
	})

	r := router.Group("/v1", BasicAuth())
	r.GET("/users", func(c *gin.Context) {
		var users []User
		db.Find(&users)
		c.JSON(200, users)
		return
	})

	router.Run(":8080")
}
