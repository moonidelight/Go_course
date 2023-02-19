package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Item struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Rating int     `json:"rating"`
}
type Data struct {
	Users []User `json:"users"`
	Items []Item `json:"items"`
}

type System interface {
	Register(name, password string) bool
	Authorization(name, password string) bool
	SearchItem(name string) bool
	FilterItems(MinPrice, MaxPrice float64, MinRating, MaxRating int)
	RatingItem(name string, rating int)
}

func (d *Data) Register(name, password string) bool {
	for _, user := range d.Users {
		if user.Name == name {
			fmt.Println("User exists")
			return false
		}
	}
	id := len(d.Users) + 1
	d.Users = append(d.Users, User{
		ID:       id,
		Name:     name,
		Password: password,
	})
	return true
}

func (d *Data) Authorization(name, password string) bool {
	for _, user := range d.Users {
		if user.Name == name && user.Password == password {
			return true
		}
	}
	return false
}

func (d *Data) SearchItem(name string) []Item {
	var arr []Item
	for _, item := range d.Items {
		if item.Name == name {
			arr = append(arr, item)
		}
	}
	return arr
}
func (d *Data) FilterItems(MinPrice, MaxPrice float64, MinRating, MaxRating int) []Item {
	var arr []Item
	for _, item := range d.Items {
		if item.Price >= MinPrice && item.Price <= MaxPrice && item.Rating >= MinRating && item.Rating <= MaxRating {
			arr = append(arr, item)
		}
	}
	return arr
}
func (d *Data) RatingItem(itemId, userId, rating int) bool {
	indexItem := -1
	for i, item := range d.Items {
		if item.ID == itemId {
			indexItem = i
		}
	}
	if indexItem == -1 {
		return false
	}
	indexUser := -1
	for i, user := range d.Users {
		if user.ID == userId {
			indexUser = i
		}
	}
	if indexUser == -1 {
		return false
	}
	d.Items[indexItem].Rating = rating
	return true
}

func (d *Data) AddItem(name string, price float64) {
	for _, item := range d.Items {
		if item.Name == name {
			return
		}
	}
	id := len(d.Items)

	d.Items = append(d.Items, Item{
		id,
		name,
		price,
		0,
	})

}
func main() {
	data, err := ioutil.ReadFile("data.json")
	if err != nil {
		return
	}
	var d Data
	err = json.Unmarshal(data, &d)
	if err != nil {
		return
	}
	ok := d.Register("User1", "123")
	if ok {
		fmt.Println("Registered")
	}
	ok = d.Authorization("User1", "123")
	if ok {
		fmt.Println("Login")
	}

	d.AddItem("item1", 20)
	d.AddItem("item2", 25)
	d.AddItem("item1", 30)

	fmt.Println(d.SearchItem("item1"))

	fmt.Println(d.FilterItems(0, 35, 0, 6))

	ok = d.RatingItem(2, 1, 5)
	if ok {
		fmt.Println("rated")
	}

	data, err = json.MarshalIndent(d, "", "  ")
	if err != nil {
		return
	}

	err = ioutil.WriteFile("data.json", data, 0644)
	if err != nil {
		return
	}
}
