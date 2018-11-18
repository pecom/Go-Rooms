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

// Since I'm not using a database and I'm bad we're iterating throughout the list to search for our room. If we change this to an actual project
// we should be using a database to store the room data instead of a list of structs. 
func searchRoom(roomNumber int) *Room {
    for i, _ := range rooms {
        p := rooms[i]
        if p.Number == roomNumber {
            return &p
        }
    }
    return &rooms[0]
}

func returnRooms(c *gin.Context){
    fmt.Println("Returning rooms... ")
    c.JSON(http.StatusOK, gin.H{
        "rooms": rooms,
    })
}

func takeRooms(c *gin.Context){
    fmt.Println("Taking a room...")
    roomNum := c.Query("Number")
    sectionName := c.Query("Section")
    takingRoom := searchRoom(roomNum)
    takingRoom.Occupied = true
    takingRoom.Section = sectionName
    returnRooms(c)
}

func emptyRooms(c *gin.Context){
    fmt.Println("Emptying a room...")
    roomNum := c.Query("Number")
    emptyRoom := searchRoom(roomNum)
    emptyRoom.Occupied = false
    emptyRoom.Section = ""
    returnRooms(c)
}

// our main function
func main() {
    router := gin.Default()
    router.GET("/checkRooms", returnRooms)
    router.POST("/takeRooms", takeRooms)
    router.POST("/emptpyRooms", emptyRooms)
    router.Run(":8080")
}
