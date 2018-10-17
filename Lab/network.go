package d7024e

import (
	"Kademlia/KademliaDHT/protobuf"
	"fmt"
	"net"
	"os"
	"sync"

	"github.com/golang/protobuf/proto"
)

type NetInterface interface {
	SendPingMessage(contact *Contact)
	SendFindContactMessage(contact *Contact, dest *Contact) CloseContacts

	//	SendFindDataMessage(contact *Contact, hash string) (*CloseContacts, *[]byte)
	//	SendStoreMessage(contact *Contact, hash string, data []byte)
}

// type CloseContact struct {
// 	contact  *Contact
// 	distance *KademliaID
// }

//List of the closecontacts fro sendfindnode
type CloseContacts []Contact

//List of responses
type Respons []protobuf.KademliaMessageCallBack

//var Kandidater ContactCandidates

type Network struct {
	Address     string
	Port        string
	Mid         int32
	NetKad      *Kademlia
	Lock        *sync.Mutex
	AnsList     *Respons
	NetworkLock *sync.Cond
	MailinBox   *bool
}
type MockNetwork struct {
}

// type neetwork interface {
// 	SendFindContactMessage(contact *Contact, dest *Contact) []Contact
// 	SendFindDataMessage(hash string, contact *Contact) (*Contacts, []byte)
// 	SendStoreMessage(data []byte)
// 	//LookupContact()
// 	//LookupContactThreads()
// }

/*
func NewNetwork(address string, port string) Network {
}
*/

func (network *Network) getResponse(ID int32) *protobuf.KademliaMessageCallBack {

	//	var start time.Time = time.Now()

	network.NetworkLock.L.Lock()
	for !*network.MailinBox {

		network.NetworkLock.Wait()
	}
	for i := 0; i < len(*network.AnsList); i++ {
		if (*network.AnsList)[i].Id == ID {
			a := (*network.AnsList)[i]
			*network.AnsList = append((*network.AnsList)[:i], (*network.AnsList)[i+1:]...)
			if len(*network.AnsList) == 0 {
				*network.MailinBox = false
			}
			network.NetworkLock.L.Unlock()
			return &a
		}

		network.NetworkLock.L.Unlock()
	}
	return nil
}

func (network *Network) addResponse(response protobuf.KademliaMessageCallBack) {
	//network.responseCond.L.Lock()
	network.NetworkLock.L.Lock()

	//fmt.Println("Twerk. Adding ID ", response.MessageID)
	*network.AnsList = append(*network.AnsList, response)
	*network.MailinBox = true
	network.NetworkLock.Broadcast()
	network.NetworkLock.L.Unlock()
}

func NewNetwork(address string, port string, kad *Kademlia) Network {
	a := Network{Address: address, Port: port, Mid: 1, NetKad: kad, Lock: &sync.Mutex{}}
	go a.Listen()
	return a
}

func (network *Network) getMessageID() int32 {
	network.Lock.Lock()
	ID := network.Mid
	network.Mid++
	network.Lock.Unlock()
	return ID
}

func ErrorHandler(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}

func (network *Network) Listen() {

	serveraddr, err := net.ResolveUDPAddr("udp", network.Address+":"+network.Port)
	ErrorHandler(err)
	listener, err := net.ListenUDP("udp", serveraddr)
	ErrorHandler(err)
	defer listener.Close()
	fmt.Println("Listening on " + network.Address + ":" + network.Port)
	buf := make([]byte, 1024)
	for {

		n, conn, err := listener.ReadFromUDP(buf)
		fmt.Printf("Received %X from %s\n", string(buf[0:n]), conn)

		ErrorHandler(err)

		var receivedMsg *protobuf.KademliaMessageType = &protobuf.KademliaMessageType{}
		//fmt.Println(receivedMsg)
		err = proto.Unmarshal(buf[:n], receivedMsg)
		fmt.Println(receivedMsg.Call.GetMessageString())
		ErrorHandler(err)
		//fmt.Println(receivedMsg)
		fmt.Println("What type am I: ", receivedMsg.Type)
		if receivedMsg.Type == protobuf.KademliaMessageType_CALLBACK {
			switch receivedMsg.Callback.Type {
			case protobuf.KademliaMessageCallBack_FINDC:
				go network.callbackFindReply(receivedMsg)
				break

			default:
				fmt.Println("")
			}
			//fmt.Println(receivedMsg.Callback.GetInfo())
		} else if receivedMsg.Type == protobuf.KademliaMessageType_CALL {

			switch receivedMsg.Call.Type {
			case protobuf.KademliaMessageCall_PING:
				go network.callbackPingMessage(*receivedMsg)
				break
			case protobuf.KademliaMessageCall_FINDC:
				go network.callbackFindContactMessage(*receivedMsg)
			default:
				fmt.Println("Error")
			}
		} else {
			fmt.Println("error")
		}

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

func (network *Network) newMessage(typ protobuf.KademliaMessageType_Type) protobuf.KademliaMessageType {

	var msg protobuf.KademliaMessageType = protobuf.KademliaMessageType{}

	msg.Type = typ
	var me protobuf.Contact = protobuf.Contact{
		ContactID: network.NetKad.RT.me.ID.String(),
		Address:   network.Address + ":" + network.Port}

	msg.SenderC = &me
	return msg
}

func (network *Network) newResponseMessage() protobuf.KademliaMessageType {
	return network.newMessage(protobuf.KademliaMessageType_CALLBACK)
}

func (network *Network) newRequestMessage() protobuf.KademliaMessageType {
	return network.newMessage(protobuf.KademliaMessageType_CALL)
}

func (network *Network) callbackPingMessage(receivedMsg protobuf.KademliaMessageType) {
	ServerAddr, err := net.ResolveUDPAddr("udp", receivedMsg.SenderC.Address)
	ErrorHandler(err)
	Conn, err := net.DialUDP("udp", nil, ServerAddr)
	ErrorHandler(err)
	defer Conn.Close()

	var msg protobuf.KademliaMessageType = network.newResponseMessage()

	var ping protobuf.KademliaMessageCallBack = protobuf.KademliaMessageCallBack{
		Id:       receivedMsg.Call.Id,
		Type:     protobuf.KademliaMessageCallBack_PING,
		Contacts: nil,
		Info:     "Hello",
	}
	msg.Callback = &ping
	fmt.Println(ping)

	//fmt.Println("messege to send ",msg)
	var buff []byte
	buff, err = proto.Marshal(&msg)
	ErrorHandler(err)

	_, err = Conn.Write(buff)
	if err != nil {
		fmt.Println(msg, err)
	}
}
func (network *Network) SendPingMessage(contact *Contact) {
	ServerAddr, err := net.ResolveUDPAddr("udp", contact.Address) //contact.address?
	ErrorHandler(err)
	LocalAddr, err := net.ResolveUDPAddr("udp", "localhost:0")
	ErrorHandler(err)
	Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	ErrorHandler(err)
	defer Conn.Close()

	var msg protobuf.KademliaMessageType = network.newRequestMessage()
	mid := network.getMessageID()
	ping := protobuf.KademliaMessageCall{
		Id:            mid,
		Type:          protobuf.KademliaMessageCall_PING,
		MessageString: "Ping",
		Info:          "",
	}
	msg.Call = &ping

	var buff []byte
	//fmt.Println(&msg)
	buff, err = proto.Marshal(&msg)
	//fmt.Println(buff)

	ErrorHandler(err)
	_, err = Conn.Write(buff)
	if err != nil {
		fmt.Println(msg, err)
	}

}

// func (network *MockNetwork) SendFindContactMessage(contact *Contact, dest *Contact) []Contact {
// var a []Contact
// fmt.Println("I am sending ContatcMsg now")
// for i := 0; i < 5; i++ {
// newcontact := NewContact(NewRandomKademliaID(), "localhost")
// newcontact.CalcDistance(contact.ID)
// a = append(a, newcontact)
// }
// return a
// }

func (network *Network) SendFindContactMessage(contact *Contact) CloseContacts {
	ServerAddr, err := net.ResolveUDPAddr("udp", contact.Address) //contact.Address?
	ErrorHandler(err)
	//	LocalAddr, err := net.ResolveUDPAddr("udp", nil, ServerAddr)
	ErrorHandler(err)
	Conn, err := net.DialUDP("udp", nil, ServerAddr)
	ErrorHandler(err)
	defer Conn.Close()

	var msg protobuf.KademliaMessageType = network.newRequestMessage()
	var mid int32 = network.getMessageID()
	msg.Call = &protobuf.KademliaMessageCall{
		Id:            mid,
		Type:          protobuf.KademliaMessageCall_FINDC,
		MessageString: fmt.Sprint(contact),
		Info:          "",
	}
	var buff []byte
	buff, err = proto.Marshal(&msg)
	ErrorHandler(err)
	_, err = Conn.Write(buff)
	// if err != nil {
	// 	fmt.Println(msg, err)
	// }

	//ÄNDRAS
	var resultList CloseContacts
	// for i := 0; i < len(); i++ {
	// 	resultList = append(resultList[i], Contact{NewKademliaID()})
	// }

	return resultList
}

func (network *Network) callbackFindContactMessage(receievedMsg protobuf.KademliaMessageType) {
	//first we must list all closestcontacts that the node knows about

	var cc CloseContacts = network.NetKad.RT.FindClosestContacts(NewKademliaID(receievedMsg.Call.MessageString), k)
	var msg protobuf.KademliaMessageType = network.newResponseMessage()
	var callback protobuf.KademliaMessageCallBack = protobuf.KademliaMessageCallBack{}
	var clist []*protobuf.Contact = []*protobuf.Contact{}
	callback.Type = protobuf.KademliaMessageCallBack_FINDC

	for i := 0; i < len(cc); i++ {
		var cont protobuf.Contact = protobuf.Contact{
			ContactID: fmt.Sprint(cc[i].ID),
			Address:   cc[i].Address,
			Xor:       fmt.Sprint(cc[i].distance),
		}

		callback.CoontactList = append(clist, &cont)
	}
	callback.Id = receievedMsg.Call.Id
	msg.Callback = &callback
	// var buffer []byte
	// var err error
	buffer, err := proto.Marshal(&msg)
	fmt.Println(buffer)
	ErrorHandler(err)

	ServerAddr, err := net.ResolveUDPAddr("udp", receievedMsg.SenderC.Address) //contact.Address?
	ErrorHandler(err)
	LocalAddr, err := net.ResolveUDPAddr("udp", "localhost:0")
	ErrorHandler(err)
	Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	ErrorHandler(err)
	defer Conn.Close()
}

func (network *Network) callbackFindReply(recMsg *protobuf.KademliaMessageType) {
	replyChan := network.NetKad.Asdf[recMsg.Callback.Thread] // HÄMTA CHANNEL
	afs := recMsg.Callback.CoontactList
	var listan []Contact
	for i := 0; i < len(afs); i++ {
		listan = append(listan, Contact{ID: NewKademliaID(afs[i].ContactID), Address: afs[i].Address, distance: NewKademliaID(afs[i].Xor)})
	}
	whatever := ContactCandidates{listan}
	// GÖR OM CONTACTS TILL TYPEN ContactCandidates
	replyChan <- whatever
}

// func (network *MockNetwork) SendFindDataMessage(hash string, contact *Contact) (*Contacts, []byte) {
// 	fmt.Println("I am sending DataMsg now")
// 	/*if hash == "FFFF" && contact.ID.String() == "rwdfvwsv" {
// 	return _, _
// 	*/
// 	return nil, nil
// }

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

}
func (network *MockNetwork) SendStoreMessage(data []byte) {
	//TODO
}

func (network *Network) SendStoreMessage(data []byte) {
	// TODO
	//if success
	//kademlia.Store(data)
}
