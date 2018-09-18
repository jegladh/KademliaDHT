package d7024e

//number of "threads"
const alpha = 3

//maximum amount of lookups
const k = 20

type Kademlia struct {
	RT        *RoutingTable
	mapdemlia map[string]string
	//Net *Network

}

//type for closest contacts ish
type Contacts []Contact

func NewKademlia(me *Contact, rt *RoutingTable) *Kademlia {
	kademlia := new(Kademlia)
	kademlia.RT = rt
	kademlia.mapdemlia = make(map[string]string)
	return kademlia
}

var neet neetwork = &MockNetwork{}

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

func (kademlia *Kademlia) LookupContact(target *Contact) {
	// TODO
}

func (kademlia *Kademlia) LookupData(hash string) {
	// TODO
}

func (kademlia *Kademlia) Store(data []byte) {
	// TODO
}
