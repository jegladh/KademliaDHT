package d7024e

import (
	protoMsg "Kademlia/KademliaDHT/protobuf"
	"fmt"
	"net"
	"os"
	"sync"

	protobuf "github.com/golang/protobuf/proto"
)

type Network struct {
	Contact *Contact
	RT      *RoutingTable
	Lock    *sync.Mutex
	NetKad  *Kademlia
	KadChan *KademliaChannel
	Threads *ThreadList
}
type KademliaChannel struct {
	MID     int32
	Channel chan *protoMsg.Message
}
type ThreadList struct {
	Threads []*KademliaChannel
}

//adds a thread to the list of threads so we can fetch them later
func (network *Network) AddThread(k *KademliaChannel) {
	network.Lock.Lock()
	network.Threads.Threads = append(network.Threads.Threads, k)
	network.Lock.Unlock()
}

//Fetches a thread/channel for returning messegaes
func (network *Network) GetThread(mid int32) *KademliaChannel {
	for _, c := range network.Threads.Threads {
		if c.MID == mid {
			return c
		}

	}
	return nil
}

func NewNetwork(contact *Contact, rt *RoutingTable, kademlia *Kademlia, kadChan *KademliaChannel) Network {
	network := Network{}
	network.Contact = contact
	network.RT = rt
	network.NetKad = kademlia

	network.KadChan = kadChan
	return network
}

func ErrorHandler(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}

//used for packing down and sending messages
func marshaller(address string, message *protoMsg.Message) {
	fmt.Println(address)
	buff, err := protobuf.Marshal(message)
	ErrorHandler(err)
	Conn, err := net.Dial("udp", address)
	ErrorHandler(err)
	_, err = Conn.Write(buff)
	ErrorHandler(err)
	Conn.Close()
}

func (network *Network) Listen(server Contact, port int) {
	serveraddr, err := net.ResolveUDPAddr("udp", server.Address)
	ErrorHandler(err)
	listener, err := net.ListenUDP("udp", serveraddr)
	ErrorHandler(err)

	defer listener.Close()

	fmt.Println("Listening on ", server.Address)
	buf := make([]byte, 1024)
	channel := make(chan []byte)

	for {
		n, conn, err := listener.ReadFromUDP(buf)
		go MsgHandler(channel, &server, network)
		fmt.Println("Received ", string(buf[0:n]), " from ", conn)

		ErrorHandler(err)
		recMessage := &protoMsg.Message{}
		error := protobuf.Unmarshal(buf[:n], recMessage)
		ErrorHandler(error)

	}
}

func MsgHandler(channel chan []byte, me *Contact, network *Network) {
	data := <-channel
	recMessage := &protoMsg.Message{}

	err := protobuf.Unmarshal(data, recMessage)
	ErrorHandler(err)
	switch recMessage.GetType() {
	//när man får en ping
	case "ping":
		fmt.Println("Hello from", recMessage.GetSenderAddress())
		answer := network.SendPongMessage(me)
		marshaller(recMessage.GetSenderAddress(), answer)

		//när man får en pong
	case "pong":
		fmt.Println("Pong")
		mID := recMessage.GetMID()
		thread := network.GetThread(mID)
		thread.Channel <- recMessage
	default:
		fmt.Println("Error in switch")
	}
}

func (network *Network) SendPingMessage(contact *Contact) {
	fmt.Println(contact.Address)
	msg := &protoMsg.Message{
		SenderID:      network.Contact.ID.String(), //proto.String(input[0]),
		SenderAddress: network.Contact.Address,     //proto.String(input[1]),
		Type:          "ping",                      //proto.String(input[2]),
	}
	fmt.Println(msg)

	marshaller(contact.Address, msg)
}

func (network *Network) SendPongMessage(contact *Contact) *protoMsg.Message {
	msg := &protoMsg.Message{
		SenderID:      network.Contact.ID.String(),
		SenderAddress: network.Contact.Address,
		Type:          "pong",
	}
	return msg
}

func (network *Network) SendFindContactMessage(contact *Contact, dest *Contact) {
	// TODO
	msg := &protoMsg.Message{
		SenderID: network.Contact.ID.String(),

	}
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
