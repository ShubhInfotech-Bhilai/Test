package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var upd album
var ps = make([]album, 0)
func putAlbums( c *gin.Context) {
    c.IndentedJSON(http.StatusOK,ps )
} 
func main() {
	resp, err := http.Get("http://localhost:8080/albums")
	if err != nil {
		fmt.Println("No response from request")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // response body is []byte
	fmt.Println(string(body))
	// convert to string before print
	//ps := make([]album, 0)
	err = json.Unmarshal(body, &ps)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ps)

	for i:=0; i<3; i++{
		if ps[i].Price >55{
            ps[i].Price=100
		}
		
		}

	
	fmt.Println(ps)

	router := gin.Default()
    router.GET("/modifyalbums", putAlbums)

    router.Run("localhost:9080")
  

	

}
