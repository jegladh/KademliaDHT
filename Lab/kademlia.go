package d7024e

import (
	"sync"
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
	Asdf            map[int32]chan (ContactCandidates)
	FindNodeCounter int32
}

//type for closest contacts ish
type Contacts []Contact

func NewKademlia(me *Contact, rt *RoutingTable) *Kademlia {
	kademlia := new(Kademlia)
	kademlia.RT = rt
	kademlia.mapdemlia = make(map[string]string)
	kademlia.Asdf = make(map[int32]chan (ContactCandidates))
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
	replyChan := make(chan (ContactCandidates))
	kademlia.Asdf[kademlia.FindNodeCounter] = replyChan
	kademlia.FindNodeCounter++
	outboundRequests := 0

	kademlia.wait.Add(alpha)
	//first alpha requests
	for i := 0; i < alpha; i++ {
		kademlia.Net.SendFindContactMessage(&resultList[i])
		outboundRequests++
		asked.Contacts = append(asked.Contacts, resultList[i])
	}
	kademlia.wait.Wait()
	replyList := <-replyChan
	for _, i := range replyList.Contacts {
		returnList.Contacts = append(returnList.Contacts, i)
	}

	//Send out more requests and check so no dupe msgs
	if outboundRequests < 3 {
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
