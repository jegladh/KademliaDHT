package main

import (
	kademlia "Kademlia/KademliaDHT/Lab"
)

func main() {
	net := new(kademlia.Network)

	contact := kademlia.NewContact(kademlia.NewRandomKademliaID(), "localhost:8000")
	//contact2 := kademlia.NewContact(kademlia.NewRandomKademliaID(), "localhost:8000")
	//net.Listen(contact, 8000)
	//<-time.After(time.Second * 1)

	net.SendPingMessage(&contact)

}
