package d7024e

import (
	"testing"
)

func TestBootstrap(t *testing.T) {
	contact1 := NewRandomKademliaID()
	contact2 := NewRandomKademliaID()
	contact3 := NewRandomKademliaID()
	contact4 := NewRandomKademliaID()
	contact5 := NewRandomKademliaID()
	bsNode := NewContact(NewKademliaID(bootstrapID), bootstrapAddr+":"+bootstrapPort)

	kad := NewKademlia(bsNode, rt)

	//kad := NewKademlia(bsNode, rt)
	bootstrap(kad, contact1)

}

// go test -run nameoftest
