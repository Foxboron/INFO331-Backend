package main

import "github.com/gin-gonic/gin"

func init() {
	r.GET("/users", func(c *gin.Context) {
		var users []User
		db.Find(&users)
		c.IndentedJSON(200, users)
		return
	})

	r.POST("/user", func(c *gin.Context) {
		var user User
		c.Bind(&user)
		db.Create(user)
		c.IndentedJSON(200, &user)
		return
	})

	r.GET("/search/:username", func(c *gin.Context) {
		var users []User
		username := c.Param("username")
		db.Where("username LIKE ?", "%"+username+"%").Find(&users)
		c.IndentedJSON(200, users)
		return
	})

	r.GET("/user/:id", func(c *gin.Context) {
		var user User
		id := c.Param("id")
		db.Find(&user, id)
		c.IndentedJSON(200, user)
		return
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
		return
	})
}
