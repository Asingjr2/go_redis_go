// Practicing use of goloang go-redis client to mirror redis tutorial activities.
package main

import (
	"fmt"
	"time"
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

	//
	//
	// // Setting a key using client.
	// err := client.Set("first_key", "one", 0).Err()
	// err = client.SetNX("first_key", "one thousand", 0).Err()	//  sets value if it does not exist.  Set would update.
	// if err != nil {
	// 	fmt.Println("there was a error", err)
	// } else {
	// 	fmt.Println("no error", err)
	// }
	
	// // Getting stored value
	// first := client.Get("first_key")
	// fmt.Println("value of first =", first.Val())
	// fmt.Println("err in getting first =", first.Err())
	// fmt.Println("args in getting first =", first.Args())

	// // Checking type
	// thisType := client.Type("first_key")
	// fmt.Println("First key is type = ", thisType)

	// // Deleting key
	// client.Del("first_key")		// deleting first
	// first = client.Get("first_key")	
	// fmt.Println("trying to print deleted first_key = ",first.Val())	// Will return blank

	// // Setting value with result function to check for and handles errors.
	// _, err = client.Set("second_key", "two", 0).Result() // first value is response (ok or fail)
	// if err != nil {
	// 	fmt.Println("my error was : ", err)
	// }
	
	// // Checking if a key exists in db.
	// check :=  client.Exists("second_key")
	// fmt.Println(check)

	// // Creating a list in redis must start with lpush.  
	// statusInt := client.LPush("myList", "firstColor")
	// fmt.Println("first push to set list", statusInt)
	// fmt.Println("***")
	// statusInt = client.LLen("myList")
	// fmt.Println("checking length of list", statusInt)
	// fmt.Println("***")
	// statusInt = client.LPush("myList", "red", "blue", "yellow", "green")
	// fmt.Println(statusInt)
	// fmt.Println("***")
	// statusInt = client.LLen("myList")
	// fmt.Println("checking new length of list", statusInt)
	// fmt.Println("***")
	// check = client.Exists("myList")
	// fmt.Println("checking in myList exists with code =", check)
	// fmt.Println("***")
	// thisType = client.Type("myList")
	// fmt.Println(thisType)


	// // ************* FLUSH DATA
	// _ = client.FlushAll()

	//
	//
	//
	
	// Creating channel to send messages to and receive from.
	myChan := make(chan string)

	// Creating PUBSUB type through subscribe methods
	sub := client.Subscribe("myChan")
	defer sub.Close()

	// Checking for receipt of information
	// First handling error than working with messages.

	msg, err := sub.ReceiveTimeout(time.Second * 300)
	if err != nil {
		fmt.Println("error subscribing and here is error ", err, myChan)
		// break
	}
	
	switch msgType := msg.(type) {
	case *redis.Subscription:
		fmt.Println("subscribed to ", msgType.Channel)
	case *redis.Message:
		fmt.Println("got a message saying : ", msgType.Payload)
	default:
		panic("unreached")
	}

	newChan := make(chan string)
	
	
	go listen("newChan")
	
	// fmt.Println(statusInt.Val())
	
	time.Sleep(10 * time.Second)
	
	pubSub := client.Subscribe("newChan")
	fmt.Println(pubSub)
	func() {
		for {
			msg, err := pubSub.ReceiveMessage()
			if err != nil {
				fmt.Println("error listening for a message")
			}
			fmt.Println("Received new message: ", msg.Payload)
		}
	}()
	// _ = client.Publish("newChan", "samus is awesome")
	// fmt.Println(newChan)
	// defer pubSub.Close()

	// myMessage, err := pubSub.ReceiveTimeout(10)
	// if err != nil {
	// 	fmt.Println("this error")
	// } else {
	// 	fmt.Println(myMessage)
	// }

	// thing, err := pubSub.ReceiveMessage()
	// if err != nil {
	// 	fmt.Println("error listening for a message")
	// } else {
	// 	fmt.Println("ssssssssssss", thing.Payload)
	// }



	_ = client.Publish("newChan", "samus is awesome")
	fmt.Println(newChan)






	// } else {

	// 	switch msg.(Type) {
	// 	case redis.Subscription:
	// 		fmt.Println("type was subscription")
	// 	}

	// 	fmt.Println("subscribed easily and here is msg ", msg, myChan)
		
	// 	myMessage, err := sub.ReceiveMessage()
	// 	if err != nil {
	// 		fmt.Println("there was an err", err)
	// 	}
	// 	if myMessage.Payload == "" {
	// 		fmt.Println("no message was sent")
	// 	} 
	// 	if myMessage.Payload != "" {
	// 		fmt.Println("this was the message", myMessage)
	// 	}

	// }
	// 	for {

				
				
	// 			fmt.Println("a \n")
	// 			fmt.Println("b \n")
	// 			fmt.Println("c \n")
	// 	fmt.Println("d \n")
	// 	// break
	// 	switch msg := msg.(type) {
	// 	case *redis.Subscription:
	// 		fmt.Println("subscribed to  ", msg.Channel)
			
	// 		_, err := client.Publish("myChan", "some words").Result()
	// 		if err != nil {
	// 			fmt.Println("error publishing")
	// 		}
	// 		_, err = client.Publish("myChan", "1").Result()
	// 		_, err = client.Publish("myChan", "2").Result()
	// 		_, err = client.Publish("myChan", "3").Result()
	// 		myMessage, err := sub.ReceiveMessage()
	// 		if err != nil {
	// 			fmt.Println("error with receive message function", err)
	// 			} 
	// 			fmt.Println("clean", myMessage)
	// 			fmt.Println("clean", myMessage.Payload)
	// 			myMessage, err = sub.ReceiveMessage()
	// 			if err != nil {
	// 				fmt.Println("error with receive message function", err)
	// 				} 
	// 				fmt.Println("clean", myMessage.Pattern)
	// 				fmt.Println("clean", myMessage.Payload)
	// 				myMessage, err = sub.ReceiveMessage()
	// 				if err != nil {
	// 					fmt.Println("error with receive message function", err)
	// 					} 
	// 					fmt.Println("clean", myMessage)
	// 					fmt.Println("clean", myMessage.Payload)
	// 					break
	// 				default:
	// 					fmt.Println("something went wrong")
	// 					// break
	// 				}
	// 			}
	// 		}
				
	// fmt.Println("left infinite loop")

	// statusInt := client.Publish("myChan", "some cooooooooool words")
	// fmt.Println("status int", statusInt.Val())
	// fmt.Println("status int", statusInt.Args())
	// myMessage, err := sub.ReceiveMessage()
	// if err != nil {
	// 	fmt.Println("error with receive message function", err)
	// } else {
	// 	fmt.Println("clean", myMessage)
	// 	fmt.Println("clean", myMessage.Channel)
	// 	fmt.Println("clean", myMessage.Pattern)
	// 	fmt.Println("clean", myMessage.Payload)
	// }

}

// Need to download vm redis cli in ubuntu to be able to update my password and other information.


func listen(publishedChan string) {
	fmt.Println("starting to listen")

	pubSub := client.Subscribe(publishedChan)
	defer pubSub.Close()

	// myMessage, err := pubSub.ReceiveTimeout(10)
	// if err != nil {
	// 	fmt.Println("this error")
	// } else {
	// 	fmt.Println(myMessage)
	// }

	for {
		msg, err := pubSub.ReceiveMessage()
		if err != nil {
			fmt.Println("error listening for a message")
		}
		fmt.Println("Received new message: ", msg.Payload)
	}

}
