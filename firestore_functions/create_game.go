package create_game

import (
	"context"
	"firebase.google.com/go"
	"google.golang.org/api/option"
	"log"
  )


func main() {
	_, err := client.Collection("cities").Doc("LA").Set(context, map[string]interface{}{
        "name":    "Los Angeles",
        "state":   "CA",
        "country": "USA",
})
if err != nil {
        // Handle any errors in an appropriate way, such as returning them.
        log.Printf("An error has occurred: %s", err)
}
}