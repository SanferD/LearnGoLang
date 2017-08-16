package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

const numOfBitsForFloat = 32
type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f\n", float32(d))
}

type database map[string]dollars

func (db database) ListInventory(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) GetPriceForItem(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "The price for %s is %s\n", item, price)	
}

func (db database) UpdatePriceOfItem(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	priceStr := req.URL.Query().Get("price")
	priceFloat, err := strconv.ParseFloat(priceStr, numOfBitsForFloat)
	price := dollars(priceFloat)

	if err!=nil || price < 0.0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid price: %q. Price must be a positive number\n", price)
		return
	}
	
	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Cannot update item %q because it does not exit\n", item)
		return
	}

	db[item] = price
	fmt.Fprintf(w, "Successfully updated price of %s to %s\n", item, price)
}

func (db database) CreateItemInDb(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	priceStr := req.URL.Query().Get("price")
	
	_, ok := db[item]
	if ok {
		w.WriteHeader(http.StatusConflict)
		fmt.Fprintf(w, "The item %s has already been created\n", item)
		return
	}
	db[item] = dollars(0.0)

	if len(priceStr) == 0 {
		fmt.Fprintf(w, "The item %s was created with default price %s\n", item, dollars(0.0))
		return
	}

	priceFloat, err := strconv.ParseFloat(priceStr, numOfBitsForFloat)
	if err != nil || priceFloat < 0.0 {
		w.WriteHeader(http.StatusNotAcceptable)
		fmt.Fprintf(w, "Invalid price: %q. Price must be a positive number\n", priceFloat)
	}
	price := dollars(priceFloat)
	db[item] = price
	fmt.Fprintf(w, "Successfully created item %s with price %s\n", item, price)
}

func (db database) DeleteItem(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")

	if _, ok := db[item]; !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Item %s does not exist inventory. Delete failed.\n", item)
		return
	}

	delete(db, item)
	fmt.Fprintf(w, "Successfully removed item %s from inventory.\n", item)
}

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.ListInventory)
	http.HandleFunc("/create", db.CreateItemInDb)
	http.HandleFunc("/read", db.GetPriceForItem)
	http.HandleFunc("/update", db.UpdatePriceOfItem)
	http.HandleFunc("/delete", db.DeleteItem)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}