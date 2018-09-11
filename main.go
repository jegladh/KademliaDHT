package main

import (
	kademlia "Kademlia/KademliaDHT/Lab"
	"fmt"
	"strconv"
	/*"net"
	"github.com/tatsushid/go-fastping"
	"os"
	"time"
	*/)

var port int

func main() {

	port := 8000
	portstr := strconv.Itoa(port)
	//kademlia.Listen("130.240.109.93", port)
	//kademlia.Ping()
	randomID := kademlia.NewRandomKademliaID()
	contact := kademlia.NewContact(randomID, "localhost:"+portstr)
	newRT := kademlia.NewRoutingTable(contact)
	//	newKademlia := kademlia.NewKademlia(&contact, newRT)

	//testID := kademlia.NewKademliaID("0f")
	//testContact := kademlia.NewContact(randomID, "localhost:"+portstr)

	//contact := kademlia.NewContact(randomID, "localhost:"+portstr)

	for n := 0; n < 10; n++ {
		portstr := strconv.Itoa(port)
		randomID := kademlia.NewRandomKademliaID()
		newC := kademlia.NewContact(randomID, "localhost:"+portstr)
		//fmt.Println(newC)
		newRT.AddContact(newC)
		port++

	}
	//newKademlia.LookupData(&contact)
	//newKademlia.LookupContact(&contact)
	fmt.Println(&contact)
	//fmt.Println(newRT.GetBucketIndex(kademlia.NewKademliaID("0f")))
	//fmt.Println(newRT.Buckets[2])
	//fmt.Println(newKademlia)

}
