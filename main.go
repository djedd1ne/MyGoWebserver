package main

import (
    "log"
    "net/http"
    
    "github.com/djedd1ne/MyGoWebserver/router"
    "github.com/djedd1ne/MyGoWebserver/services"
    "github.com/djedd1ne/MyGoWebserver/utils"
)

func main() {
    log.Println("In main app")

    var dbconn = utils.GetConnection()
    defer dbconn.Close()
	services.SetDB(dbconn)
	var appRouter = router.CreateRouter()

	log.Println("Listening on Port 8000")
	log.Fatal(http.ListenAndServe(":8000", appRouter))
}