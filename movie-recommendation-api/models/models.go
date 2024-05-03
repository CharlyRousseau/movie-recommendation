package models

type User struct {
    ID       int64
    Username string
    Email    string
    Password string
}

type Favorite struct {
    ID     int64
    UserID int64
    ItemID int64
}
