package main

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	gorm.Model
	Username string
	Password string
	Groups   []Group `gorm:"many2many:user_groups;"`
	Photo    string
	Points   int
}

type Group struct {
	gorm.Model
	Name     string
	Owner    User
	OwnerID  int
	Beacon   Beacon
	BeaconID int
	Users    []User `gorm:"many2many:user_groups;"`
	// Becaons []Beacon `gorm:"ForeignKey:ID"`
	Points int
}

type Beacon struct {
	gorm.Model
	UUID  string
	Major string
	Minor string
	Name  string
}

type Event struct {
	gorm.Model
	User    User
	UserID  int
	Group   Group
	GroupID int
	Event   string
	Date    time.Time
	Value   int
}

var db *gorm.DB

func init() {
	init_db, err := gorm.Open("sqlite3", "info331.db")
	if err != nil {
		panic("failed to connect database")
	}
	db = init_db
	db.Set("gorm:auto_preload", true)

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Group{})
	db.AutoMigrate(&Beacon{})
	db.AutoMigrate(&Event{})
}
