package process

import (
	"processappsimple/utilities"
)

type customerZip struct {
	customerId string
	zipCode    string
}

type purchase struct {
	customerId string
	productId  string
}

// A process node participates in a chain of processing. A node can be at any position in the chain as its
// behvior is determined by the messages it receiveds in inputQ.
//
// If a node has a non-nil nextProcess, it will complete its job and send a nextProcessMsg to nextProcess.
//
// The following messages are handled by a process node, noting that each process chain works on a subset of each
// slice (from startIndex to endIndex, inclusive on both ends):
//
// message: "find customers"
//   - uses productIds and purchaseHistory to create sets of customer ids (what customers bought what products)
//
// message: "find zips"
//   - uses sets of customer ids to create sets of zip codes (the customers live where)
//
// message: "find farthest"
//   - uses sets of zip codes to determine the farthest zip code of the subset (the process chain's answer)
type processNode struct {
	id                  string
	chainId             string
	nextProcess         *processNode
	nextProcessMsg      string
	startIndex          int // inclusive : works on slices from the element start
	endIndex            int // inclusive : works on slices to and including element end
	zipCodeUtils        *utilities.ZipCodeUtil
	productIds          *[]string      // input to process "find customers"
	purchaseHistory     *[]purchase    // input to process "find customers"
	customerIdLists     *[][]string    // output from process "find customers"/input to "find zips"
	customerIdsWithZips *[]customerZip // input to process "find zips"
	zipCodeLists        *[][]string    // output from process "find zips"/input to "find farthest"
	homeBaseZip         string         // the zipcode of the home base to compute distance
	farthestZips        *[]string      // output from "find farthest"
	inputQ              chan message
}

// Places a message in the process node's input queue
func (n *processNode) acceptMessage(msg message) {
	n.inputQ <- msg
}

// Initiates the proces node's processing algorithm (a goroutine). The node will not process messages
// until this function is called.
func (n *processNode) start(master Master, finishedMsgContent string) {

	go func() {

		for {

			select {
			case msg := <-n.inputQ:

				switch msg.content {
				case "find customers":

					index := 0

					for p := n.startIndex; p <= n.endIndex; p++ {

						index = 0

						for ph := n.startIndex; ph <= n.endIndex; ph++ {

							if (*n.productIds)[p] == (*n.purchaseHistory)[ph].productId {

								index = ph
								break

							}

						}

						(*n.customerIdLists)[index] = append((*n.customerIdLists)[index], (*n.purchaseHistory)[index].customerId)

					}

					n.nextProcess.acceptMessage(message{"find zips", n.id})

				case "find zips":

					index := 0

					for p := n.startIndex; p <= n.endIndex; p++ {

						for c := range (*n.customerIdLists)[p] {

							index = 0

							for cz := range *n.customerIdsWithZips {

								if (*n.customerIdsWithZips)[cz].customerId == (*n.customerIdLists)[p][c] {

									index = cz
									break

								}

							}

							// We probably should ensure uniqueness to cutdown on repeat work, but we'll leave that
							// for another day.
							//
							(*n.zipCodeLists)[p] = append((*n.zipCodeLists)[p], (*n.customerIdsWithZips)[index].zipCode)

						}

					}

					n.nextProcess.acceptMessage(message{"compute farthest", n.id})

				case "compute farthest":

					for zl := n.startIndex; zl <= n.endIndex; zl++ {

						farthestDistance := 0.0
						farthestZip := ""

						for z := range (*n.zipCodeLists)[zl] {

							currentDistance, _ := n.zipCodeUtils.GetDistanceInMiles((*n.zipCodeLists)[zl][z], n.homeBaseZip)

							if farthestDistance < currentDistance {

								farthestDistance = currentDistance
								farthestZip = (*n.zipCodeLists)[zl][z]

							}

						}

						(*n.farthestZips)[zl] = farthestZip

					}

					master.NodeFinished(message{finishedMsgContent, n.chainId})

				}

			}

		}

	}()

}
