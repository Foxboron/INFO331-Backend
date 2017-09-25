package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	gorm.Model
	Username string
	Password string
	Groups   []Group
	Photo    string
	Points   int
}

type Group struct {
	gorm.Model
	Name   string
	Owner  User
	Users  []User
	Points int
}

type Beacons struct {
	gorm.Model
	BeaconID string
	Name     string
	Lat      int
	Long     int
}

type Stats struct {
	gorm.Model
	UserID string
	Event  string
	Date   time.Time
	Value  int
}

var db *gorm.DB

func init() {
	init_db, err := gorm.Open("sqlite3", "info331.db")
	if err != nil {
		panic("failed to connect database")
	}
	db = init_db

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Group{})
	db.AutoMigrate(&Beacons{})
	db.AutoMigrate(&Stats{})
	fmt.Println("Woop")
}
