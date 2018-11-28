package d7024e

import (
	protoMsg "Kademlia/KademliaDHT/protobuf"
	"fmt"
	protobuf "github.com/golang/protobuf/proto"
	"net"
	"os"
	"sync"
)

type Network struct {
	Contact *Contact
	RT 		*RoutingTable
	Lock	*sync.Mutex
	NetKad	*Kademlia


}

func NewNetwork(contact *Contact, rt *RoutingTable, kademlia *Kademlia) Network{
	network := Network{}
	network.Contact = contact
	network.RT = rt
	network.NetKad = kademlia
	return network
}

func ErrorHandler(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}

//used for packing down and sending messages
func marshaller(address string, message *protoMsg.Message){
	buff, err := protobuf.Marshal(message)
	ErrorHandler(err)
	Conn, err := net.Dial("udp", address)
	ErrorHandler(err)
	_, err = Conn.Write(buff)
	ErrorHandler(err)
	Conn.Close()
}

func Listen(network *Network)(server Contact, port int) {
	serveraddr, err := net.ResolveUDPAddr("udp", server.Address)
	ErrorHandler(err)
	listener, err := net.ListenUDP("udp", serveraddr)
	ErrorHandler(err)
	defer listener.Close()

	fmt.Println("Listening on ", server.Address )
	buf := make([]byte, 1024)
	for {
		n, conn, err := listener.ReadFromUDP(buf)
		fmt.Println("Received ", string(buf[0:n]), " from ", conn)

		ErrorHandler(err)
		recMessage := &protoMsg.Message{}
		error := protobuf.Unmarshal(buf[:n], recMessage)

		switch recMessage.GetType() {
		//när man får en ping
		case "ping":
			fmt.Println("Hello from", recMessage.GetSenderAddress())
			answer := network.SendPongMessage(network.Contact)
			marshaller(recMessage.GetSenderAddress(), answer)
		//när man får en pong
		case "pong":
			fmt.Println("Pong")
		default:
			fmt.Println("Error in switch")
		}

	}
}
func (network *Network) Ping(contact *Contact){

}

func (network *Network) SendPingMessage(contact *Contact){

	msg := &protoMsg.Message{
		SenderID:    	network.Contact.ID.String(),//proto.String(input[0]),
		SenderAddress: 	network.Contact.Address,//proto.String(input[1]),
		Type:     		"ping",//proto.String(input[2]),
		}

	marshaller(contact.Address, msg)
}

func (network *Network) SendPongMessage(contact *Contact) *protoMsg.Message{
	msg := &protoMsg.Message{
		SenderID:		network.Contact.ID.String(),
		SenderAddress:	network.Contact.Address,
		Type:			"pong",

	}
	return msg
}


func (network *Network) SendFindContactMessage(contact *Contact, dest *Contact) {
	// TODO
	//if success
	//kademlia.LookupContact()
}


func (network *Network) SendFindDataMessage(hash string) {
	// TODO
	//if success
	//kademlia.LookupData()
}

func (network *Network) SendStoreMessage(data []byte) {
	// TODO
	//if success
	//kademlia.Store(data)
}
