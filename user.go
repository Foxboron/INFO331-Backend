package main

import "github.com/gin-gonic/gin"

func init() {
	r.GET("/users", func(c *gin.Context) {
		var users []User
		db.Preload("Groups").Preload("Groups.Owner").Preload("Groups.Users").Find(&users)
		c.IndentedJSON(200, users)
	})

	r.POST("/user", func(c *gin.Context) {
		var user User
		c.Bind(&user)
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
		db.Model(&user).Updates(&newUser)
		db.Save(&user)
		c.IndentedJSON(200, &user)
	})
}
