package models

type User struct {
	Email       string `json:"email"`
	Name        string `json:"name"`
	Number      string `json:"number"`
	Expense     int64  `json:"expense"`
	Password    string `json:"password"`
	ConfirmPass string `json:"confirm"`
}

type Transaction struct {
	Users     []*User // Each transaction would have a list of users
	SplitType string  `json:"splittype"` // Type of split taken place in transaction: equal, exact, percentage
	Total     int64   `json:"total"`     // The total amount in
}
