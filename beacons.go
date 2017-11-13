package main

import "github.com/gin-gonic/gin"

func init() {

	r.GET("/search/beacons/:uuid", func(c *gin.Context) {
		var beacons []Beacon
		uuid := c.Param("uuid")
		db.Where("uuid LIKE ?", "%"+uuid+"%").Find(&beacons)
		c.IndentedJSON(200, &beacons)
	})

	r.GET("/beacons", func(c *gin.Context) {
		var beacons []Beacon
		db.Find(&beacons)
		c.IndentedJSON(200, &beacons)
	})

	r.POST("/beacons", func(c *gin.Context) {
		var beacon Beacon
		c.Bind(&beacon)
		db.Create(&beacon)
		c.IndentedJSON(200, &beacon)
	})

	r.GET("/beacon/:id", func(c *gin.Context) {
		var beacon Beacon
		id := c.Param("id")
		db.Find(&beacon, id)
		c.IndentedJSON(200, &beacon)
	})

	r.PUT("/beacon/:id", func(c *gin.Context) {
		var beacon Beacon
		id := c.Param("id")
		db.Find(&beacon, id)
		var newBeacon Beacon
		c.Bind(&newBeacon)
		db.Model(&beacon).Updates(&newBeacon)
		c.IndentedJSON(200, &beacon)

	})

	r.DELETE("/beacon/:id", func(c *gin.Context) {
		var beacon Beacon
		id := c.Param("id")
		if db.First(&beacon, id).RecordNotFound() {
			c.IndentedJSON(500, gin.H{"message": "Didn't find any beacons",
				"status": "failure"})
			return
		}
		db.Delete(&beacon)
		c.IndentedJSON(200, gin.H{"message": "Deleted",
			"status": "succsess"})
	})
}
