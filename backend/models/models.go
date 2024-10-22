package models

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
)

// users table
// id, email, number, name, password, expense
type User struct {
	UserID      int    `json:"userid"                      gorm:"column:userid;primaryKey;autoIncrement"`
	Email       string `json:"email"     form:"email"      gorm:"unique;not null"`
	Name        string `json:"name"      form:"name"       gorm:"not null"`
	Number      string `json:"number"    form:"number"     gorm:"not null"`
	Expense     int64  `json:"expense"`
	Password    string `json:"password"  form:"password"   gorm:"not null"`
	ConfirmPass string `json:"confirm"   form:"confirm"    gorm:"-"`
}

// toverify --> vid, username, email, password, timestamp
// helpful for sending verification emails and updating users database
type Verification struct {
	VerificationID int       `json:"vid"                        gorm:"column:vid;primaryKey;autoIncrement"`
	Email          string    `json:"email"    form:"email"      gorm:"unique;not null"`
	Name           string    `json:"name"     form:"name"       gorm:"unique;not null"`
	Password       string    `json:"password" form:"password"   gorm:"not null"`
	CreatedAt      time.Time `                                  gorm:"column:created_at;index"`
}

// struct for credentials
// will be used during login
type Credentials struct {
	Email    string `json:"email"    form:"email"`
	Password string `json:"password" form:"password"`
}

// midlleware token verification struct
type VerifyClaims struct {
	jwt.RegisteredClaims
	Email string  `json:"email"`
}

// transactions table
// id, users, splittype, total
type Transaction struct {
	TransactionID int     `json:"transid"`
	Users         []*User 					 // Each transaction would have a list of users
	SplitType     string  `json:"splittype"` // Type of split taken place in transaction: equal, exact, percentage
	Total         int64   `json:"total"`     // The total amount in
}