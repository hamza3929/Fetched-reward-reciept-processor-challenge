package main

import (
	"net/http"
	"unicode"
	"math"
	"strings"
	"strconv"
	"sync"
	"github.com/google/uuid"
	"github.com/gin-gonic/gin"
)

// Define struct for an item
type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

// Define struct for the receipt
type Receipt struct {
	Retailer      string `json:"retailer"`
	PurchaseDate  string `json:"purchaseDate"`
	PurchaseTime  string `json:"purchaseTime"`
	Total         string `json:"total"`
	Items         []Item `json:"items"`
	//points		  int	 `json: points`
}

type ReceiptData struct {
	Data map[string]Receipt
	Sync sync.Mutex
}

var store = ReceiptData{Data: make(map[string]Receipt)}

func Process_Receipt(c *gin.Context){
	var receipt Receipt
	if err := c.ShouldBindJSON(&receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid receipt format"})
		return
	}
	id := uuid.New().String()
	store.Sync.Lock()
	store.Data[id] = receipt
	store.Sync.Unlock()
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func getPoints(c *gin.Context) {
	id := c.Param("id")
	store.Sync.Lock()
	receipt, exists := store.Data[id]
	store.Sync.Unlock()
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Receipt not found"})
		return
	}
	points := calculatePoints(receipt)
	c.JSON(http.StatusOK, gin.H{"points": points})
}

func calculatePoints(receipt Receipt) (int){
	var points int = 0

	points += retailer_name_alphanumeric(receipt.Retailer)
	td, _ :=strconv.ParseFloat(receipt.Total, 64)
	points += total_is_whole_dollar(td)
	points += total_is_multiple_of_25(td)
	points += every_2_items(receipt.Items)
	points += trim_price(receipt.Items)
	points += is_date_odd(receipt.PurchaseDate)
	points += between_2_and_4pm(receipt.PurchaseTime)

	return points
}

func retailer_name_alphanumeric(name string) (int){
	//One point for every alphanumeric character in the retailer name
	var alphanum int = 0
	for i:=0; i < len(name); i+=1 {
		if (unicode.IsLetter(rune(name[i])) || unicode.IsNumber(rune(name[i]))){
			alphanum += 1
		}
	  }
	  
	  return alphanum
}

func total_is_whole_dollar(t float64) (int){
	//50 points if the total is a round dollar amount with no cents
	if (math.Mod(t, 1.0) == 0){
		
		return 50
	}
	
	return 0
}

func total_is_multiple_of_25(t float64) (int){
	//25 points if the total is a multiple of 0.25
	if (math.Mod(t, 0.25) == 0){
		return 25
	}
	
	return 0
}

func every_2_items(it []Item) (int){
	//5 points for every two items on the receipt
	return ((len(it) / 2) * 5)
}

func trim_price(it []Item) (int){
	//If the trimmed length of the item description is a multiple of 3, 
	//multiply the price by 0.2 and round up to the nearest integer. 
	//The result is the number of points earned
	var p int = 0
	for i := 0; i < len(it); i += 1 {

		if (len(it[i].ShortDescription) % 3 == 0) {
			pri, _ := strconv.ParseFloat(it[i].Price, 64)
			p += int(math.Floor(pri * 0.2))
		}
	}
	return p
}

func is_date_odd(d string) (int){
	//6 points if the day in the purchase date is odd
	day, _ :=  strconv.Atoi((strings.Split(d, "-")[2]))

	if (day % 2 != 0) {
		return 6
	}

	return 0
}

func between_2_and_4pm(t string) (int){
	var hm []string =  strings.Split(t, ":")

	var ti [2]int
	ti[0], _ = strconv.Atoi(hm[0])
	ti[1], _ = strconv.Atoi(hm[1])

	if ((ti[0] >= 2 && ti[0] <= 4)) {
		if (ti[0] == 4 && ti[1] == 0) {
			return 10
		}
		return 10
	}
	
	return 0
}

func main() {
	r := gin.Default()
	r.POST("/receipts/process", Process_Receipt)
	r.GET("/receipts/:id/points", getPoints)
	r.Run(":8080")
	
}

