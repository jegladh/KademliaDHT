package d7024e

import (
	"fmt"
	"testing"
)

func TestCheckAsked(t *testing.T) {
	contact1 := NewRandomKademliaID()
	contact2 := NewRandomKademliaID()
	target := Contact{contact1, "127.0.0.1", contact1}
	nottarget := Contact{contact2, "127.0.0.1", contact2}
	contacts := Contacts{target, nottarget}

	if !checkAskedNodes(target, contacts) {
		t.Errorf("%v was not found in the list, not checked", nottarget)
	}
	if checkAskedNodes(target, contacts) {
		fmt.Printf("%v has been checked, grats", target)
		fmt.Println("")
	}

}

func TestFindNode(t *testing.T) {

}
