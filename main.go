// Practicing use of goloang go-redis client to mirror redis tutorial activities.
package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

type Person struct {
	Name string
	Age int
}

// Creating universal variable
var client *redis.Client

func main() {
	fmt.Println("hey Samus")

	// Instanciate new redis client
	client := redis.NewClient(&redis.Options{
		// Default location port
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	// Checking for successful connection
	_, pingErr := client.Ping().Result()
	if pingErr != nil {
		fmt.Print("there was an error", pingErr)
	} else {
		fmt.Println("connections successful")
	}

	// Setting a key using client.
	err := client.Set("first_key", "one", 0).Err()
	err = client.SetNX("first_key", "one thousand", 0).Err()	//  sets value if it does not exist.  Set would update.
	if err != nil {
		fmt.Println("there was a error", err)
	} else {
		fmt.Println("no error", err)
	}
	
	// Getting stored value
	first := client.Get("first_key")
	fmt.Println("value of first =", first.Val())
	fmt.Println("err in getting first =", first.Err())
	fmt.Println("args in getting first =", first.Args())

	// Checking type
	thisType := client.Type("first_key")
	fmt.Println("First key is type = ", thisType)

	// Deleting key
	client.Del("first_key")		// deleting first
	first = client.Get("first_key")	
	fmt.Println("trying to print deleted first_key = ",first.Val())	// Will return blank

	// Setting value with result function to check for and handles errors.
	_, err = client.Set("second_key", "two", 0).Result() // first value is response (ok or fail)
	if err != nil {
		fmt.Println("my error was : ", err)
	}
	
	// Checking if a key exists in db.
	check :=  client.Exists("second_key")
	fmt.Println(check)

	// Creating a list in redis must start with lpush.  
	statusInt := client.LPush("myList", "firstColor")
	fmt.Println("first push to set list", statusInt)
	fmt.Println("***")
	statusInt = client.LLen("myList")
	fmt.Println("checking length of list", statusInt)
	fmt.Println("***")
	statusInt = client.LPush("myList", "red", "blue", "yellow", "green")
	fmt.Println(statusInt)
	fmt.Println("***")
	statusInt = client.LLen("myList")
	fmt.Println("checking new length of list", statusInt)
	fmt.Println("***")
	check = client.Exists("myList")
	fmt.Println("checking in myList exists with code =", check)
	fmt.Println("***")
	thisType = client.Type("myList")
	fmt.Println(thisType)


	// ************* FLUSH DATA
	_ = client.FlushAll()

}

// Need to download vm redis cli in ubuntu to be able to update my password and other information.
