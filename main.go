package main

import (
	"net/http"
	"strconv"
	// "fmt"
	// "golang.org/x/tools/cmd/guru/serial"
	"github.com/gin-gonic/gin"
)

// Waifu represents data about a record Waifu.

type Season int;

const (
    Winter Season = iota + 1
    Spring
    Summer
    Fall
)
type waifu struct {
    ID     int  `json:"id"`
    Name  string  `json:"name"`
    Anime  string  `json:"anime"`
    Season Season  `json:"season"`
    Year  float64 `json:"year"`
}

// func main() {
//     w := waifu{
//             Season: Summer,
//     }
//     fmt.Println(w.Season) // Output: 3
// }
func main() {
    router := gin.Default()
    router.GET("/waifus", getWaifus)
    router.POST("/waifus", postWaifus)
    router.GET("/waifus/:id", getWaifuByID)

    router.Run("localhost:8080")
}


// Waifus slice to seed record Waifu data.
var waifus = []waifu{
    {ID: 1, Name: "Alya Kujo", Anime: "Roshidere", Season: Summer, Year: 2024},
    {ID: 2, Name: "Ayase Seki", Anime: "Gimai Seikatsu", Season: Summer, Year: 2024},
    {ID: 3, Name: "Yamada Anna", Anime: "BokuYaba", Season: Spring, Year: 2023},
    {ID: 4, Name: "Kurumi", Anime: "Date A Live", Season: Winter, Year: 2019},

}
var curId int = 4;

// getWaifus responds with the list of all Waifus as JSON.
func getWaifus(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, waifus)
}

// func NewWaifuStore() *waifuStore {
//     return &waifuStore{
//         users: make(map[int]*User),
//     }
// }

// postWaifus adds an Waifu from JSON received in the request body.
func postWaifus(c *gin.Context) {
    curId++
    newWaifu := waifu{ID: curId};

    // Call BindJSON to bind the received JSON to
    // newWaifu.
    if err := c.BindJSON(&newWaifu); err != nil {
        return
    }

    // Add the new Waifu to the slice.
    waifus = append(waifus, newWaifu)
    c.IndentedJSON(http.StatusCreated, newWaifu)
}


// getWaifuByID locates the Waifu whose ID value matches the id
// parameter sent by the client, then returns that Waifu as a response.
func getWaifuByID(c *gin.Context) {
    id := c.Param("id")

    // Convert id to an integer
    idInt, err := strconv.Atoi(id)
    if err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
        return
    }

    // Loop over the list of Waifus, looking for
    // a Waifu whose ID value matches the parameter.
    for _, a := range waifus {
        if a.ID == idInt {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Waifu not found"})
}

//kyunrious