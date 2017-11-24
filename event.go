package main

import "github.com/gin-gonic/gin"

func init() {

	r.POST("/event/:userid/group/:groupid", func(c *gin.Context) {
		var group Group
		groupid := c.Param("groupid")
		db.Find(&group, groupid)
		var user User
		userid := c.Param("userid")
		if userid == "" {
			c.IndentedJSON(400, gin.H{"message": "Didn't find any user id",
				"status": "failure"})
			return
		}
		db.Find(&user, userid)
		var event Event
		c.Bind(&event)
		db.Model(&event).Association("Group").Append(&group)
		db.Model(&event).Association("User").Append(&user)
		c.IndentedJSON(200, &event)
	})

	r.GET("/events/:userid", func(c *gin.Context) {
		var events []Event
		id := c.Param("userid")
		db.Preload("User").Preload("Group").Find(&events, id)
		c.IndentedJSON(200, &events)
	})

}
