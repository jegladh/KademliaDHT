package main

import (
	"strconv"

	kademlia "github.com/KademliaDHT/Lab"
	/*"net"
	"github.com/tatsushid/go-fastping"
	"os"
	"time"
	*/)

var port int

func main() {

	port := 32768
	portstr := strconv.Itoa(port)
	//os.Setenv("VAR", "10.0.0.4")
	//kademlia.Listen(os.Getenv("NODE1"), port)
	kademlia.Listen("10.0.0.5", port)
	kademlia.Ping()
	randomID := kademlia.NewRandomKademliaID()
	contact := kademlia.NewContact(randomID, "10.0.0.5"+portstr)
	newRT := kademlia.NewRoutingTable(contact)
	newKademlia := kademlia.NewKademlia(&contact, newRT)

	//testID := kademlia.NewKademliaID("0f")
	//testContact := kademlia.NewContact(randomID, "localhost:"+portstr)

	// for n := 0; n < 20; n++ {
	// 	portstr := strconv.Itoa(port)
	// 	randomID := kademlia.NewRandomKademliaID()
	// 	newC := kademlia.NewContact(randomID, "172.19.0.3"+portstr)
	// 	fmt.Println(newC)
	// 	newRT.AddContact(newC)
	// 	port++
	//
	// }
	//newKademlia.LookupData(&contact)
	newKademlia.LookupContact(&contact)
	//fmt.Println(&contact)
	//fmt.Println(newRT.GetBucketIndex(kademlia.NewKademliaID("3578")))
	//fmt.Println(newRT.Buckets[2])
	//fmt.Println(newKademlia)

}
