package d7024e

import (
	"sync"
	"time"
)

//number of "threads"
const alpha = 3

//maximum amount of lookups
const k = 20

type Kademlia struct {
	RT              *RoutingTable
	mapdemlia       map[string]string
	Net             *Network
	wait            *sync.WaitGroup
	Asdf            map[int32]chan (KadChannel)
	FindNodeCounter int32
}

type KadChannel struct {
	Timestamp time.Time
	kademlia  ContactCandidates
}

//type for closest contacts ish
type Contacts []Contact

func NewKademlia(me *Contact, rt *RoutingTable) *Kademlia {
	kademlia := new(Kademlia)
	kademlia.RT = rt
	kademlia.mapdemlia = make(map[string]string)
	kademlia.Asdf = make(map[int32]chan (KadChannel))

	kademlia.FindNodeCounter = 1
	return kademlia
}

//var neet neetwork = &MockNetwork{}

//Functions for sorting
func (c Contacts) Len() int {
	return len(c)
}

func (c Contacts) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c Contacts) Less(i, j int) bool {
	return c[i].distance.Less(c[j].distance)
}

//function to check if a node already has been asked
func checkAskedNodes(contact Contact, asked []Contact) bool {
	for _, b := range asked {
		if contact == b {
			return true
		}
	}
	return false
}

func (kademlia *Kademlia) FindNode(target *Contact) []Contact {

	resultList := kademlia.RT.FindClosestContacts(target.ID, alpha)
	returnList := ContactCandidates{resultList}
	asked := ContactCandidates{}
	compareList := make([]Contact, 20)
	replyChan := make(chan (KadChannel))
	kademlia.Asdf[kademlia.FindNodeCounter] = replyChan
	kademlia.FindNodeCounter++
	outboundRequests := 0
	//map [contact.Address] timer // kan behöva ändra typen i channeln till (address, []contacts)

	exitCounter := 0
	exitBool := false

	kademlia.wait.Add(alpha)
	//first alpha requests
	for i := 0; i < alpha; i++ {
		kademlia.Net.SendFindContactMessage(&resultList[i])
		outboundRequests++
		asked.Contacts = append(asked.Contacts, resultList[i])
	}
	kademlia.wait.Wait()
	returnList.Sort()
	// for loop123
	//for loop still wrong condition
exit:
	for true {
		// snapshot 20 first
		compareList = returnList.Contacts[:19]
		replyList := <-replyChan
		for _, c := range replyList {
			returnList.Contacts = append(returnList.Contacts, c)
		}
		returnList.Sort()

		compareCounter := 0
		for i := 0; i < len(returnList.Contacts); i++ {
			if compareList[i] == returnList.Contacts[i] {
				compareCounter++
			}
		}
		if compareCounter == k {
			exitCounter++
		} else {
			exitCounter = 0
		}

		if exitCounter == 3 {
			for _, i := range resultList {
				kademlia.Net.SendFindContactMessage(&i)

			}
			break exit
		}
		asked.Sort()
		//Send out more requests and check so no dupe msgs
		if outboundRequests < 3 && exitBool == true {
		loop1:
			for _, i := range resultList {
				if !checkAskedNodes(i, asked.Contacts) {
					kademlia.Net.SendFindContactMessage(&i)
					asked.Contacts = append(asked.Contacts, i)
					outboundRequests++
					break loop1
				}
			}
		}
	}
	// end loop123
	returnList.Sort()
	return returnList.Contacts[0:19]

}

// for i := 0; i < 3; i++ {
// 	network.SendFindContactMessage(resultList[i], target)
// 	outboundRequests++
// 	asked = append(asked, resultList[i])
// }
// WAIT FOR REPLIES
// outboundRequests--
// ADD REPLY CONTACTLIST TO RESULTLIST AND SORT
// for _, i := range responseList {
// 	resultList = append(resultList, i)
// }
// resultList.Sort()

// CHEKC FOR EXIT Cond

// if outboundRequests < 3 {
//loop1:
// for _, i := range (resultList) {
// 	if i not in asked {
// 		// Send request
// 		outboundRequests++
//break loop1
// 	}
// }
// }
//

// After loop SORT then return

func (kademlia *Kademlia) LookupContact(target *Contact) {
	// TODO
}

func (kademlia *Kademlia) LookupData(hash string) {
	// TODO
}

func (kademlia *Kademlia) Store(data []byte) {
	// TODO
}
