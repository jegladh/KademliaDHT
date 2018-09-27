package main

import (
	kademlia "Kademlia/KademliaDHT/Lab"
	"fmt"
	"sync"
	"time"
	//	"strconv"
	/*"net"
	"github.com/tatsushid/go-fastping"
	"os"
	"time"
	*/)

func main() {

	//net := new(kademlia.Network)

	//	net.Listen()
	fmt.Println("Starting server")
	//net.Listen()

	randomID := kademlia.NewRandomKademliaID()
	contact := kademlia.NewContact(randomID, "localhost:"+"8000")
	rt := kademlia.NewRoutingTable(contact)
	kad := kademlia.NewKademlia(&contact, rt)
	net := kademlia.Network{
		Address: "localhost",
		Port:    "8000",
		Mid:     123,
		NetKad:  kad,
		Lock:    &sync.Mutex{},
	}
	go net.Listen()
	<-time.After(time.Second * 1)
	net.SendPingMessage(&contact)
	<-time.After(time.Second * 1)

	/*port := 8000
	portstr := strconv.Itoa(port)
	kademlia.Listen("localhost", port)
	kademlia.Ping()


	newRT := kademlia.NewRoutingTable(contact)
	newKademlia := kademlia.NewKademlia(&contact, newRT)

	//testID := kademlia.NewKademliaID("0f")
	//testContact := kademlia.NewContact(randomID, "localhost:"+portstr)

	for n := 0; n < 20; n++ {
		portstr := strconv.Itoa(port)
		randomID := kademlia.NewRandomKademliaID()
		newC := kademlia.NewContact(randomID, "localhost:"+portstr)
		fmt.Println(newC)
		newRT.AddContact(newC)
		port++

	}
	//newKademlia.LookupData(&contact)
	newKademlia.LookupContact(&contact)
	//fmt.Println(&contact)
	//fmt.Println(newRT.GetBucketIndex(kademlia.NewKademliaID("3578")))
	//fmt.Println(newRT.Buckets[2])
	//fmt.Println(newKademlia)
	*/

}
