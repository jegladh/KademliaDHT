package main

import (
	kademlia "Kademlia/KademliaDHT/Lab"
	"fmt"
	"sync"
	"time"

	"strconv"
	/*"net"
	"github.com/tatsushid/go-fastping"
	"os"
	"time"
	*/)

func main() {

	////net := new(kademlia.Network)
	////net.NetKad.Store([]byte{1})
	//randomID := kademlia.NewRandomKademliaID()
	//contact := kademlia.NewContact(randomID, "localhost:"+"8000")
	//contact2 := kademlia.NewContact(kademlia.NewRandomKademliaID(), "localhost"+"8003")
	//rt := kademlia.NewRoutingTable(contact)
	//rt.AddContact(kademlia.NewContact(kademlia.NewRandomKademliaID(), "localhost:8002"))
	//
	//kad := kademlia.NewKademlia(&contact, rt)
	////net.Listen()
	//<-time.After(time.Second * 1)
	//kad.FindNode(&contact2)
	////	net.Listen()
	//// fmt.Println("Starting server")
	//// //net.Listen()

	randomID := kademlia.NewRandomKademliaID()
	contact := kademlia.NewContact(randomID, "localhost:"+"8000")
	//contact2 := kademlia.NewContact(kademlia.NewRandomKademliaID(), "localhost"+"8003")
	rt := kademlia.NewRoutingTable(contact)
	kad := kademlia.NewKademlia(&contact, rt)

	for i:=0 ; i<100; i++  {
		port := 8004 + i
		portstr := strconv.Itoa(port)
		rt.AddContact(kademlia.NewContact(kademlia.NewRandomKademliaID(), "localhost:" + portstr))

	}

	mailBoxbool := false
	net := kademlia.Network{
		Address: "localhost",
		Port:    "8000",
		Mid:     123,
		NetKad:  kad,
		Lock:    &sync.Mutex{},
		AnsList:     nil,
		NetworkLock: nil,
		MailinBox:   &mailBoxbool,
	}
	fmt.Println(kad)
	go net.Listen()
	<-time.After(time.Second * 1)
	fmt.Println(contact.ID)
	go kad.FindNode(&contact)
	<-time.After(time.Second * 2)
	//go net.SendPingMessage(&contact)
	//fmt.Println(hej)

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
