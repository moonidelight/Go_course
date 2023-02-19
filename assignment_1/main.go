package main

import (
	"fmt"
)

type User struct {
	UserName string
	Password string
}
type Item struct {
	Name   string
	Price  float64
	Rating int
}
type Data struct {
	Users map[string]User
	Items map[string]Item
}

func (user *User) Login(name, password string) bool {
	if user.UserName == name && user.Password == password {
		return true
	}
	return false
}
func (d *Data) Register(user, password string) {
	_, isPresent := d.Users[user]
	if !isPresent {
		newUser := User{user, password}
		d.Users[user] = newUser
		fmt.Println("You're registered")
		//return true
		return
	}
	fmt.Println("User exists")
	//return false
}
func (d *Data) Authentication(userName, password string) {
	_, present := d.Users[userName]
	if present {
		user := d.Users[userName]
		if user.Login(userName, password) {
			fmt.Println("Correct")

		} else {
			fmt.Println("Not correct password")
		}

	} else {
		fmt.Println("User doesn't exist")
	}
}

func (d *Data) SearchItem(name string) []Item {
	var arr []Item
	if it, found := d.Items[name]; found {
		arr = append(arr, it)
	}
	return arr
}

func (d *Data) FilterItems(price float64, rating int) []Item {
	var filter []Item
	for _, item := range d.Items {
		if item.Price <= price && item.Rating <= rating {
			filter = append(filter, item)
		}
	}
	return filter
}

func (d *Data) GiveRating(name string, rating int) bool {
	if it, isPresent := d.Items[name]; isPresent {
		it.Rating = rating
		d.Items[name] = it
		return true
	}
	return false
}
func main() {
	data := Data{
		make(map[string]User),
		make(map[string]Item),
	}
	data.Register("Gaga", "gaga")
	data.Authentication("Gaga", "ga")
	a := Item{"bag", 16, 0}
	b := Item{"pen", 22, 0}
	//c := Item{"bag1", 66, 5}
	data.Items[a.Name] = a
	data.Items[b.Name] = b
	searchingItem := data.SearchItem("bag1")
	if len(searchingItem) > 0 {
		fmt.Println(searchingItem)
	} else {
		fmt.Println("Not found")
	}
	fmt.Println("GiveRating:", data.GiveRating("bag", 2))
	fmt.Println("Filtered Items:", data.FilterItems(16, 2))
}
