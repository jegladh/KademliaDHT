package d7024e

import {
//  "fmt"
//  "net"
  "github.com/golang/protobuf/proto"
  "github.com/KademliaDHT/protobuf"
}

type protoMessage struct {
  network *Network
}

func sendProtoMessage (input []string) *protobuf.KademliaMessageResponse { //building message to protobuf
  if input[0] == "Ping" || input[0] == "Pong" {
    message := &protobuf.KademliaMessageResponse {
      Label: proto.String(input[0]),
      Senderid: proto.String(input[1]),
      SenderAddress: proto.String(input[2]),
    }
    return message
  }
}

func (this *protoMessage) protoMessageHandler(channel chan []byte, me Contact, network *Network ) { //to unmarshal and handle message from protobuf
  data := <-channel
  message := &protobuf.KademliaMessageResponse{}
  err := proto.Unmarshal(data, message)
  ErrorHandler(err)

  answering := NewContact(NewKademliaID(message.GetSenderid()), message.GetSenderAddress())
  fmt.Println("\n\nListner:", me, "\nSender: ", answering, "\nMessage: ", message)

  //the different protobuf messages here:
  if *message.Label == "Ping" {
    fmt.Println("n", message)
    answer := sendProtoMessage([]string{"Pong", me.ID.String(), me.Address})
    send(message.GetSenderAddress(), answer)
  }

  if *message.Label == "Pong" {
    fmt.Println("\n", message)
    answer := sendProtoMessage([]string{"Pong", me.ID.String(), me.Address})
    send(message.GetSenderid(), answer)
  }
}

func send(Address string, message *protobuf.KademliaMessageResponse) {
    fmt.Println("send to address: ", Address)
    data, err := proto.Marshal(message)
    if err != nil {
      fmt.Println("Marshal Error: ", err)
    }

    Conn, err := net.Dial("udp", Address)
    if err != nil {
      fmt.Println("UDP-Error: ", err)
    }
    defer Conn.Close()
    _, err = Conn.Write(data)
    if err != nil {
      fmt.Println("Write Error: ", err)
    }
}
