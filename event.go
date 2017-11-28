package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

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
		event.Date = time.Now()
		db.Set("gorm:save_associations", false).Create(&event)
		db.Model(&event).Association("Group").Append(&group)
		db.Model(&event).Association("User").Append(&user)
		db.Save(&event)
		c.IndentedJSON(200, &event)
	})

	r.GET("/events/:userid", func(c *gin.Context) {
		var events []Event
		id := c.Param("userid")
		db.Preload("User").Preload("Group").Where("user_id = ?", id).Find(&events)
		c.IndentedJSON(200, &events)
	})

	r.GET("/stats/:userid", func(c *gin.Context) {
		var events []Event
		id := c.Param("userid")
		db.Preload("User").Preload("Group").Where("user_id = ?", id).Find(&events)
		var score = 0.0
		for k := 0; k < len(events); k += 2 {
			enter := events[k]
			exit := events[k+1]
			b := exit.Date.Sub(enter.Date)
			score += b.Minutes()
		}
		final_score := int(score + 0.5)
		var scoreRet Score
		scoreRet.Score = final_score
		c.IndentedJSON(200, &scoreRet)
	})

	r.GET("/stats/:userid/group/:groupid", func(c *gin.Context) {
		var events []Event
		userid := c.Param("userid")
		groupid := c.Param("groupid")
		db.Preload("User").Preload("Group").Where("user_id = ? and group_id = ?", userid, groupid).Find(&events)
		var score = 0.0
		for k := 0; k < len(events); k += 2 {
			enter := events[k]
			exit := events[k+1]
			b := exit.Date.Sub(enter.Date)
			score += b.Minutes()
		}
		final_score := int(score + 0.5)
		var scoreRet Score
		scoreRet.Score = final_score
		c.IndentedJSON(200, &scoreRet)
	})
}
