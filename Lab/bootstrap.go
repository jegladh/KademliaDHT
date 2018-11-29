package d7024e

import (
	"fmt"
)

const (
	bootstrapAddr = "10.0.0.5"
	bootstrapPort = "37680"
	bootstrapID   = "FFFFFFFF00000000000000000000000000000000" //vad ska jag sätta för id?
	port          = "8001"
)

func Bootstrap(kademlia *Kademlia, target *Contact) {
	bsNode := NewContact(NewKademliaID(bootstrapID), bootstrapAddr+":"+bootstrapPort)

	//Connecta restrerande noder till bootstrap nätverket -> GÖRS MED LOOKUPC
	res := kademlia.FindNode(&bsNode) //ger list [0:19] närmaste kontakter, sparar listan i variablen res
	//loopa igenom res för att lägga till alla kontakerna från listan.
	for i := 0; i < len(res); i++ {
		kademlia.RT.AddContact(res[i])
		fmt.Println("Closest contacts: ")
		fmt.Print(res[i])
		kademlia.Net.SendPingMessage(&res[i]) //ändra ping, så att vid mottagandet av en ping updateras bucket med den som kontaktat
	}
	kademlia.RT.AddContact(bsNode) //updaterar bucket med bsNoden

}
