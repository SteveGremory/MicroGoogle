package main

import (
    "encoding/json"
    "fmt"
    "os"
    "net/http"
    "MicroGoogle/search"
    "github.com/gin-gonic/gin"
)

func check(e error) {
    if e != nil {
        fmt.Println("Error: ", e)
    }
}

type SerpResponse struct {
    Breadcrumb string `json:"breadcrumb"`
    Description string `json:"description"`
    Link string `json:"link"`
    Title string `json:"title"`
}

type SerpRequest struct {
    Pages string `json:"pages"`
    Query string `json:"query"`
}

func QueryGoogle(c *gin.Context) {
    // Validate input
    var input SerpRequest
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    res := search.CrawlGoogle(input.Query, input.Pages, "json")
    c.JSON(http.StatusOK, gin.H{"response": res})
}

func loadResponse() []SerpResponse {
    // Read the file
    data, err := os.ReadFile("./whatagain.json")
    check(err)
    
    // Create a response object
    // Deserialise the JSON
    var res []SerpResponse
    json.Unmarshal([]byte(data), &res)
    
    return res; 
}

func main() {

    router := gin.Default()
    router.POST("/search/google", QueryGoogle)

    router.Run("localhost:8080")
}
