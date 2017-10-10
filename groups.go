package main

import (
	"github.com/gin-gonic/gin"
)

func init() {

	r.GET("/groups", func(c *gin.Context) {
		var groups []Group
		db.Preload("Owner").Preload("Users").Find(&groups)
		c.IndentedJSON(200, &groups)
		return
	})

	r.GET("/group/:id", func(c *gin.Context) {
		var group Group
		groupid := c.PostForm("id")
		// Preload is used for fields where we use
		// another database
		db.Preload("Owner").Preload("Users").Find(&group, groupid)
		c.IndentedJSON(200, &group)
	})

	r.POST("/group", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(User)
		var group Group
		c.Bind(&group)
		group.Owner = user
		db.Create(&group)
		db.Model(&user).Association("Groups").Append(&group)
		c.IndentedJSON(200, &group)
	})

	r.PUT("/group/:id", func(c *gin.Context) {
		var group Group
		id := c.Param("id")
		db.Find(&group, id)
		var newGroup Group
		c.Bind(&newGroup)
		db.Model(&group).Updates(&newGroup)
		// We dont want to null out associations
		db.Set("gorm:save_associations", false).Save(&group)
		c.IndentedJSON(200, &group)
	})

	r.DELETE("/group/:id", func(c *gin.Context) {
	})

	r.GET("/group/:id/users", func(c *gin.Context) {
	})

	r.POST("/group/:groupid/user/:userid", func(c *gin.Context) {
		authedUser := c.MustGet(gin.AuthUserKey).(User)
		var group Group
		var owner User
		groupid := c.Param("groupid")
		db.Find(&group, groupid)
		db.Model(&group).Association("Owner").Find(&owner)
		if owner.ID != authedUser.ID {
			c.IndentedJSON(400, gin.H{"message": "The owner has to add members to a group",
				"status": "failure"})
			return
		}

		var user User
		userid := c.Param("userid")
		if userid == "" {
			c.IndentedJSON(400, gin.H{"message": "Didn't find any user id",
				"status": "failure"})
			return
		}
		db.Find(&user, userid)
		db.Model(&user).Association("Groups").Append(&group)
	})

	r.DELETE("/group/:groupid/user/:userid", func(c *gin.Context) {
	})
}
