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

var router = gin.Default()
var r = router.Group("/v1", BasicAuth())

func main() {

	router.POST("/login", func(c *gin.Context) {
		var user User
		username := c.PostForm("username")
		password := c.PostForm("password")

		if db.Preload("Groups").Preload("Groups.Owner").Preload("Groups.Users").First(&user, "username = ?", username).RecordNotFound() {
			c.Status(400)
			return
		}
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err != nil {
			c.IndentedJSON(200, &user)
			return
		}
		c.IndentedJSON(200, &user)
		return
	})

	router.POST("/register", func(c *gin.Context) {
		var user User
		username := c.PostForm("username")
		password := c.PostForm("password")
		if db.First(&user, "username = ?", username).RecordNotFound() {
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
			if err != nil {
				panic(err)
			}
			user := &User{Username: username, Password: string(hashedPassword)}
			db.Create(&user)
			c.IndentedJSON(200, &user)
			return
		}
		c.Status(400)
		return
	})
	router.GET("/debug/drop", func(c *gin.Content) {
		db.DropTableIfExists(&User{})
		db.DropTableIfExists(&Groups{})
		db.DropTableIfExists("user_groups")
		db.DropTableIfExists(&Bacons{})
		db.DropTableIfExists(&Stats{})
		db.AutoMigrate(&User{})
		db.AutoMigrate(&Group{})
		db.AutoMigrate(&Beacon{})
		db.AutoMigrate(&Stats{})
	})
	router.Run(":8000")
}
