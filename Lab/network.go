package d7024e

import (
	"Kademlia/KademliaDHT/protobuf"
	"fmt"
	"net"
	"os"

	"github.com/golang/protobuf/proto"
)

type Network struct {
	address string
	port    string
}
type MockNetwork struct {
}
type neetwork interface {
	SendFindContactMessage(contact *Contact, dest *Contact) []Contact
	SendFindDataMessage(hash string, contact *Contact) (*Contacts, []byte)
	SendStoreMessage(data []byte)
	//LookupContact()
	//LookupContactThreads()
}

/*
func NewNetwork(address string, port string) Network {

}*/

func ErrorHandler(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}

func (network *Network) Listen() {

	serveraddr, err := net.ResolveUDPAddr("udp", network.address+":"+network.port)
	ErrorHandler(err)
	listener, err := net.ListenUDP("udp", serveraddr)
	ErrorHandler(err)
	defer listener.Close()
	fmt.Println("Listening on " + network.address + ":" + network.port)
	buf := make([]byte, 1024)
	for {

		n, conn, err := listener.ReadFromUDP(buf)
		fmt.Println("Received ", string(buf[0:n]), " from ", conn)
		ErrorHandler(err)

		var receivedMsg *protobuf.KademliaMessageRequest = &protobuf.KademliaMessageRequest{}
		err = proto.Unmarshal(buf[:n], receivedMsg)

		ErrorHandler(err)
		/*
			if receivedMsg.Type == protobuf.KademliaMessageResponse{

			}
			if receivedMsg.Type == protobuf.KademliaMessageRequest{
				go network.newMessage()
			}
		*/

		/*
			if receivedMsg.Type == protobuf.KademliaMessageResponse_PONG {
				go network.SendPongMessage(receivedMsg)
			}
			if receivedMsg.Type == protobuf.KademliaMessageRequest_PING {
				go network.SendPingMessage(contact)
			}
		*/
	}
}

/*//http://130.240.110.178:8000/
func Ping() {
	ServerAddr, err := net.ResolveUDPAddr("udp", "localhost:8000") //contact.address?
	ErrorHandler(err)
	LocalAddr, err := net.ResolveUDPAddr("udp", "localhost:0")
	ErrorHandler(err)
	Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	ErrorHandler(err)
	defer Conn.Close()
	i := 0
	for {
		msg := strconv.Itoa(i)
		i++
		buf := []byte(msg)
		_, err := Conn.Write(buf)
		if err != nil {
			fmt.Println(msg, err)
		}
		time.Sleep(time.Second * 1)
	}
}*/
func (network *Network) SendPongMessage(receivedMsg protobuf.KademliaMessageRequest) {
	ServerAddr, err := net.ResolveUDPAddr("udp", "localhost:8000") //contact.address?
	ErrorHandler(err)
	Conn, err := net.DialUDP("udp", nil, ServerAddr)
	ErrorHandler(err)
	defer Conn.Close()

}

func (network *Network) SendPingMessage(contact *Contact) {
	ServerAddr, err := net.ResolveUDPAddr("udp", contact.Address) //contact.address?
	ErrorHandler(err)
	LocalAddr, err := net.ResolveUDPAddr("udp", "localhost:0")
	ErrorHandler(err)
	Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	ErrorHandler(err)
	defer Conn.Close()

	var msg messages.Message = network.newRequestMessage()
	var mID int64 = network.getMessageID()
	var ping messages.Request = messages.Request{mID, messages.Request_PING, "", nil}
	msg.Request = &ping

	var buff []byte
	buff, err = proto.Marshal(&msg)
	CheckError(err)
	_, err = Conn.Write(buff)
	//fmt.Printf("wrote %d bytes\n", n)
	if err != nil {
		fmt.Println(msg, err)
	}
	var response *messages.Response = network.getResponse(mID)
	if response == nil {
		return false
	}
	//fmt.Printf("ID should be %v, is %v, type %v\n", mID, response.MessageID, response.Type)
	//fmt.Println(response)
	if response.MessageID == mID && response.Type == messages.Response_PING {
		return true
	}
	return false

}

func (network *MockNetwork) SendFindContactMessage(contact *Contact, dest *Contact) []Contact {
	var a []Contact
	fmt.Println("I am sending ContatcMsg now")
	for i := 0; i < 5; i++ {
		newcontact := NewContact(NewRandomKademliaID(), "localhost")
		newcontact.CalcDistance(contact.ID)
		a = append(a, newcontact)
	}
	return a
}

func (network *Network) SendFindContactMessage(contact *Contact, dest *Contact) {
	ServerAddr, err := net.ResolveUDPAddr("udp", "localhost:8000") //contact.Address?
	ErrorHandler(err)
	LocalAddr, err := net.ResolveUDPAddr("udp", "localhost:0")
	ErrorHandler(err)
	Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	ErrorHandler(err)
	defer Conn.Close()
	//if success
	//kademlia.LookupContact()
}

func (network *MockNetwork) SendFindDataMessage(hash string, contact *Contact) (*Contacts, []byte) {
	fmt.Println("I am sending DataMsg now")
	/*if hash == "FFFF" && contact.ID.String() == "rwdfvwsv" {
	return _, _
	*/
	return nil, nil
}

// var s []byte
// fmt.Println(" I am sending DataMsg now")
// for i := 0; i < 5; i++ {
// 	*newdata := NewContact(NewRandomKademliaID(), "localhost")
//
// }
// return &newdata, s

func (network *Network) SendFindDataMessage(hash string) {
	ServerAddr, err := net.ResolveUDPAddr("udp", "localhost:8000")
	ErrorHandler(err)
	LocalAddr, err := net.ResolveUDPAddr("udp", "localhost:0")
	ErrorHandler(err)
	Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	ErrorHandler(err)
	defer Conn.Close()
	//if success
	//kademlia.LookupData()
}
func (network *MockNetwork) SendStoreMessage(data []byte) {
	//TODO
}

func (network *Network) SendStoreMessage(data []byte) {
	// TODO
	//if success
	//kademlia.Store(data)
}
