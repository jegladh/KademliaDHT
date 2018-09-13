package d7024e

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

type Network struct {
}

func ErrorHandler(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}

func Listen(ip string, port int) {
	/*convert port to string*/
	portstr := strconv.Itoa(port)
	/**/
	serveraddr, err := net.ResolveUDPAddr("udp", ip+":"+portstr)
	ErrorHandler(err)
	/*listen at port*/
	listener, err := net.ListenUDP("udp", serveraddr)
	ErrorHandler(err)
	defer listener.Close()
	fmt.Println("Listening on " + ip + ":" + portstr)
	buf := make([]byte, 1024)
	for {
		n, conn, err := listener.ReadFromUDP(buf)
		fmt.Println("Received ", string(buf[0:n]), " from ", conn)
		ErrorHandler(err)
	}
}

//http://130.240.110.178:8000/
func Ping() {
	ServerAddr, err := net.ResolveUDPAddr("udp", "localhost:9999")
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
}

func (network *Network) SendPingMessage(contact *Contact) {
	// TODO
}

func (network *Network) SendFindContactMessage(contact *Contact) {
	// TODO
	//if succes kademlia.LookupContact()
}

func (network *Network) SendFindDataMessage(hash string) {
	// TODO
}

func (network *Network) SendStoreMessage(data []byte) {
	// TODO
}
