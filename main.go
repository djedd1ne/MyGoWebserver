package main

import (
    "fmt"
    "log"

    "github/djedd1ne/MyGoWebserver/utils"
)

type Post struct {
    ID    int    `db:"id"`
    Title string `db:"title"`
    Body  string `db:"body"`
}

func main() {
    db := utils.GetConnection()
    defer db.Close()

    var posts []Post
    err := db.Select(&posts, "SELECT * FROM posts")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Posts in database:")
    for _, post := range posts {
        fmt.Printf("ID: %d, Title: %s, Body: %s\n", post.ID, post.Title, post.Body)
    }
}