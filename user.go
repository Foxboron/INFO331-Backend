package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func init() {
	r.GET("/users", func(c *gin.Context) {
		var users []User
		db.Preload("Groups").Preload("Groups.Owner").Preload("Groups.Users").Find(&users)
		c.IndentedJSON(200, users)
	})

	r.POST("/user", func(c *gin.Context) {
		var user User
		c.Bind(&user)
		if user.Password != "" {
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
			if err != nil {
				c.IndentedJSON(500, gin.H{"message": "Password error",
					"status": "failure"})
				return
			}
			user.Password = string(hashedPassword)
		}
		db.Create(user)
		c.IndentedJSON(200, &user)
	})

	r.GET("/search/users/:username", func(c *gin.Context) {
		var users []User
		username := c.Param("username")
		db.Preload("Groups").Where("username LIKE ?", "%"+username+"%").Find(&users)
		c.IndentedJSON(200, &users)
	})

	r.GET("/user/:id", func(c *gin.Context) {
		var user User
		id := c.Param("id")
		db.Preload("Groups").Preload("Groups.Owner").Preload("Groups.Users").Find(&user, id)
		c.IndentedJSON(200, user)
	})

	r.PUT("/user/:id", func(c *gin.Context) {
		var user User
		id := c.Param("id")
		db.Find(&user, id)
		var newUser User
		c.Bind(&newUser)
		if user.Password != "" {
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
			if err != nil {
				c.IndentedJSON(500, gin.H{"message": "Password error",
					"status": "failure"})
				return
			}
			newUser.Password = string(hashedPassword)
		}
		db.Model(&user).Updates(&newUser)
		db.Save(&user)
		c.IndentedJSON(200, &user)
	})
}
