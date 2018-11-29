package main

import (
	"fmt"

	kademlia "github.com/KademliaDHT/Lab"
	//	"strconv"
	/*"net"
	"github.com/tatsushid/go-fastping"
	"os"
	"time"
	*/)

const (
	bootstrapAddr = "10.0.0.5"
	bootstrapPort = "37680"
	bootstrapID   = "FFFFFFFF00000000000000000000000000000000" //vad ska jag sätta för id?
	port          = "8001"
)

//hjälp funktion för att hitta om ip = bootip
func GetMyIp(ip string, port string) {
	if ip == bootstrapAddr {
		println("I'm the bootstap node: \n ")
		return NewContact(NewKademliaID(bootstrapID), bootstrapAddr+":"+bootstrapPort)
	} else {
		return NewContact(NewRandomKademliaID(), ip+":"+port)
	}
}

func main() {

	contact1 := kademlia.NewRandomKademliaID()
	contact2 := kademlia.NewRandomKademliaID()
	contact3 := kademlia.NewRandomKademliaID()
	contact4 := kademlia.NewRandomKademliaID()
	contact5 := kademlia.NewRandomKademliaID()

	me := GetMyIp(ip, port)
	var bip *kademlia.Kademlia
	if ip != bootstrapAddr {
		bip = kademlia.NewKademlia(me, port)
	} else {
		bip = kademlia.NewKademlia(me, bootstrapPort)
	}

	kademlia.Network.Listen(bip)
	fmt.Println("Listening on: " + bip)
	if ip != bootstrapAddr {
		Bootstrap(bip, me)
	}

	// net := new(kademlia.Network)
	// net.NetKad.Store([]byte{1})
	// randomID := kademlia.NewRandomKademliaID()
	// contact := kademlia.NewContact(randomID, "localhost:"+"8000")
	// contact2 := kademlia.NewContact(kademlia.NewRandomKademliaID(), "localhost"+"8003")
	// rt := kademlia.NewRoutingTable(contact)
	// rt.AddContact(kademlia.NewContact(kademlia.NewRandomKademliaID(), "localhost:8002"))
	//
	// kad := kademlia.NewKademlia(&contact, rt)
	// //net.Listen()
	// kad.FindNode(&contact2)

	//	net.Listen()
	// fmt.Println("Starting server")
	// //net.Listen()
	//
	// randomID := kademlia.NewRandomKademliaID()
	// contact := kademlia.NewContact(randomID, "localhost:"+"8000")
	// rt := kademlia.NewRoutingTable(contact)
	// kad := kademlia.NewKademlia(&contact, rt)
	// net := kademlia.Network{
	// 	Address: "localhost",
	// 	Port:    "8000",
	// 	Mid:     123,
	// 	NetKad:  kad,
	// 	Lock:    &sync.Mutex{},
	// }
	// go net.Listen()
	// <-time.After(time.Second * 1)
	// net.SendFindContactMessage(&contact)
	// <-time.After(time.Second * 1)

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
