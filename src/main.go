package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "fmt"
)

type Room struct {
    Number int
    Occupied bool
    Section string
}

var rooms = [4]Room{Room{416, false, ""}, Room{418, false, ""}, Room{420, false, ""}, Room{423, false, ""}}
var sections = [14]string{"CB", "Staff Development", "News", "Photo", "Design", "Opinion", "Sports", "Copy", "The Eye", "A&E", "Product", "Spectrum", "Engagement", "Revenue"}

func returnRooms(c *gin.Context){
    fmt.Println("Returning rooms... ")
    c.JSON(http.StatusOK, gin.H{
        "rooms": rooms,
    })
}

func searchRoom(roomNumber int){
    for i, _ := range rooms {
        p = rooms[i]
        if p.Number == roomNumber {
            return &p
        }
    }
}

func takeRooms(c *ginContext){

}

func emptyRooms(c *ginContext){
    fmt.Println("Emptying a room...")
    roomNum := c.Query("Number")
    emptyRoom := searchRoom(roomNum)
    emptyRoom.Occupied = false
    emptyRoom.Section = false
    
}

// our main function
func main() {
    router := gin.Default()
    router.GET("/checkRooms", returnRooms)
    rounter.POST("/takeRooms", takeRooms)
    router.Run(":8080")
}
