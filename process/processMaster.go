package process

import (
	"bufio"
	"fmt"
	"os"
	"processappsimple/utilities"
	"strconv"
	"strings"
	"time"
)

// The process master manages a collection of process nodes orgranized into chainCount
// chains of processCount processes.
//
// The purpose of the process master is to compute the zip code of the farthest location
// from the home base zip code to which a product shipped.
//
// Each process chain works on a subset of the productIds slice by determining the zip
// code of the farthest shippment containined in the subslice. The process master
// determines the overall farthest zip code.
type ProcessMaster struct {
	Master
	chainCount      int
	processCount    int
	nodeCount       int
	nodes           [][]*processNode
	homeBaseZipCode string
	productIds      *[]string
	farthestZips    *[]string
	zipCodeUtil     *utilities.ZipCodeUtil
	inputQ          chan message
}

// Returns a pointer to a stice containing all the product ids found in
// the specified file.
func loadProductIds(filepath string) *[]string {

	productIds := make([]string, 0)

	file, err := os.Open(filepath)

	if err != nil {
		fmt.Println("**ERROR**", err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		productIds = append(productIds, scanner.Text())

	}

	return &productIds

}

// Returns a pointer to a stice containing all the customer-zip code
// pairs in the specified file.
func loadCustomerIdsAndZips(filepath string) *[]customerZip {

	customerZips := make([]customerZip, 0)

	file, err := os.Open(filepath)

	if err != nil {
		fmt.Println("**ERROR**", err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		tokens := strings.Split(scanner.Text(), " ")

		customerZips = append(customerZips, customerZip{tokens[0], tokens[1]})

	}

	return &customerZips

}

// Returns a pointer to a slice containing all the purchse histories
// (a pair of customerId and productId indicating the customer bought)
// the product) found in the specified file.
func loadPurchaseHistory(filepath string) *[]purchase {

	purchaseHistory := make([]purchase, 0)

	file, err := os.Open(filepath)

	if err != nil {
		fmt.Println("**ERROR**", err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		// File lines should be "customerId productId"
		tokens := strings.Split(scanner.Text(), " ")

		purchaseHistory = append(purchaseHistory, purchase{tokens[0], tokens[1]})

	}

	return &purchaseHistory

}

// Returns a pointer to a process master that establishes chainCount process chains (i.e. each chain gets
// 1/chainCount portion of the problem) and will be capable of computing the farthest distance from the
// specified home base.
//
// The data files required by the process master are expected to be in the folder dataFolderPath.
func CreateProcessMaster(dataFolderPath string, chainCount int, homebaseZipCode string) *ProcessMaster {

	dataFolderPath = strings.TrimSpace(dataFolderPath)

	if dataFolderPath != "" && dataFolderPath[len(dataFolderPath)-1] != '/' {

		dataFolderPath += "/"

	}

	PRODUCT_ID_FILEPATH := dataFolderPath + "productIds.txt"
	CUSTOMER_IDS_WITH_ZIPS_FILEPATH := dataFolderPath + "customerIdsWithZips.txt"
	PURCHASE_HISTORY_FILEPATH := dataFolderPath + "purchaseHistory.txt"

	// There are three stages to the data transformation:
	//
	// process node message: "find customer"
	//    takes a set of product ids and finds all customer ids of customers who
	//    purchased the products
	//
	// process node message: "find zipcodes"
	//    takes sets of customer ids and determines all the zip codes in which
	//    they live
	//
	// process node message: "find farthest"
	//    takes sets of zip codes and determines which is the farthes from the
	//    home base zip code
	//
	const PROCESS_COUNT int = 3

	// Create the process master.
	//
	processMaster := new(ProcessMaster)

	// Establish the input and output slices shared between master and nodes and nodes and nodes.
	//
	// productIds          *[]string      // input to process "find customers"
	// purchseHistory      *[]purchase    // input to process "find customers"
	// customerIdLists     *[][]string    // output from process "find customers"/input to "find zips"
	// customerIdsWithZips *[]customerZip // input to process "find zips"
	// zipCodeLists        *[][]string    // output from process "find zips"/input to "find farthest"
	// farthestZips        *[]string      // output from "find farthest"
	//
	fmt.Print("Loading product ids... ")

	start := time.Now()
	productIds := loadProductIds(PRODUCT_ID_FILEPATH)
	elapsed := time.Since(start)

	fmt.Println("Done:", elapsed)

	fmt.Print("Loading purchase history... ")

	start = time.Now()
	purchaseHistory := loadPurchaseHistory(PURCHASE_HISTORY_FILEPATH)
	elapsed = time.Since(start)

	fmt.Println("Done:", elapsed.String()+". Total purchases:", len(*purchaseHistory))

	customerIdLists := make([][]string, len(*productIds))

	for r := range customerIdLists {
		customerIdLists[r] = make([]string, 0)
	}

	fmt.Print("Loading customer zips... ")

	start = time.Now()
	customerIdsWithZips := loadCustomerIdsAndZips(CUSTOMER_IDS_WITH_ZIPS_FILEPATH)
	elapsed = time.Since(start)

	fmt.Println("Done:", elapsed)

	zipCodeLists := make([][]string, 0)

	for r := 0; r < len(customerIdLists); r++ {
		zipCodeLists = append(zipCodeLists, make([]string, 0))
	}

	farthestZips := make([]string, len(*productIds))

	// Determine the numnber of product ids each node will process. This will be used to tell
	// the nodes what section of the slice it is responsible for.
	//
	subsliceCount := (len(*productIds) / chainCount)

	// Prep the nodes
	//
	processMaster.chainCount = chainCount
	processMaster.processCount = PROCESS_COUNT
	processMaster.nodeCount = processMaster.chainCount * processMaster.processCount
	processMaster.productIds = productIds
	processMaster.farthestZips = &farthestZips
	processMaster.homeBaseZipCode = homebaseZipCode

	// Create the zip code utility
	//
	processMaster.zipCodeUtil, _ = utilities.GetZipCodeUtilInstance()

	// There are 40 chains reporting back so set the buffer to 40.
	//
	processMaster.inputQ = make(chan message, processMaster.chainCount)

	processMaster.nodes = make([][]*processNode, 0, processMaster.processCount)

	fmt.Print("Creating and starting the nodes... ")

	start = time.Now()

	for c := 0; c < processMaster.chainCount; c++ {

		processMaster.nodes = append(processMaster.nodes, make([]*processNode, 0, processMaster.processCount))

		for p := 0; p < processMaster.processCount; p++ {

			processMaster.nodes[c] = append(processMaster.nodes[c], new(processNode))

			processMaster.nodes[c][p].id = "(" + strconv.Itoa(c) + ", " + strconv.Itoa(p) + ")"

			processMaster.nodes[c][p].chainId = strconv.Itoa(c)

			processMaster.nodes[c][p].inputQ = make(chan message, 4)

			processMaster.nodes[c][p].productIds = productIds
			processMaster.nodes[c][p].purchaseHistory = purchaseHistory
			processMaster.nodes[c][p].customerIdLists = &customerIdLists
			processMaster.nodes[c][p].customerIdsWithZips = customerIdsWithZips
			processMaster.nodes[c][p].zipCodeLists = &zipCodeLists
			processMaster.nodes[c][p].homeBaseZip = processMaster.homeBaseZipCode
			processMaster.nodes[c][p].farthestZips = &farthestZips

			processMaster.nodes[c][p].startIndex = c * subsliceCount
			processMaster.nodes[c][p].endIndex = processMaster.nodes[c][p].startIndex + subsliceCount - 1

			processMaster.nodes[c][p].zipCodeUtils = processMaster.zipCodeUtil

			processMaster.nodes[c][p].start(processMaster, "Done")

		}

	}

	elapsed = time.Since(start)

	fmt.Println("Done:", elapsed)

	// Wire the grid
	//
	fmt.Print("Configuring the process chains... ")

	start = time.Now()

	for c := range processMaster.nodes {

		for p := range processMaster.nodes[c] {

			// Here's what the processNode chains should look like after the wiring:
			//
			//
			//        ---> O ---> O ---> O --->        \
			//        ---> O ---> O ---> O --->   o    |
			//    i   ---> O ---> O ---> O --->   u    |
			//    n               .               t    |
			//    p               .               p    | CHAIN_COUNT chains
			//    u               .               u    |
			//    t   ---> O ---> O ---> O --->   t    |
			//    s   ---> O ---> O ---> O --->   s    |
			//        ---> O ---> O ---> O --->        /
			//            \_______________/
			//          PROCES_COUNT process levels
			//
			//
			//
			// This means, we only need to wire
			//
			//    - processNodes in col 0 to point to col 1 processNodes
			//    - processNodes in col 1 to point to col 2 processNodes
			//
			if p == 0 {

				processMaster.nodes[c][p].nextProcess = processMaster.nodes[c][p+1]
				processMaster.nodes[c][p].nextProcessMsg = "find zipcodes"

			} else if p == 1 {

				processMaster.nodes[c][p].nextProcess = processMaster.nodes[c][p+1]
				processMaster.nodes[c][p].nextProcessMsg = "find farthest"

			}

		}

	}

	elapsed = time.Since(start)

	fmt.Println("Done:", elapsed)

	return processMaster

}

// Executes the process over the chains and returns the zip code and distance to which the farthest
// shipment was made.
func (pm *ProcessMaster) FindFarthestZipInMiles() (farthestZip string, farthestDistance float64) {

	messageCount := pm.chainCount

	fmt.Println("Determining farthest zip code... ")

	start := time.Now()

	// Send messages to the chains to begin.
	//
	for c := 0; c < pm.chainCount; c++ {

		pm.nodes[c][0].acceptMessage(message{"find customers", NETWORK_MASTER})

	}

	// Wait for all chains to respond.
	//
	for m := 0; m < messageCount; m++ {

		msg := <-pm.inputQ

		fmt.Println("  > Heard from chain:", msg.senderId)

	}

	// Find the farthes zip from the chain results.
	//
	farthestDistance = 0.0
	farthestZip = "UNDEFINED"

	for z := range *pm.farthestZips {

		currentDistance, _ := pm.zipCodeUtil.GetDistanceInMiles(pm.homeBaseZipCode, (*pm.farthestZips)[z])

		if currentDistance > farthestDistance {

			farthestDistance = currentDistance
			farthestZip = (*pm.farthestZips)[z]

		}

	}

	elapsed := time.Since(start)

	fmt.Println("Done determining farthest zip code: ", elapsed)

	return farthestZip, farthestDistance

}

func (pm *ProcessMaster) NodeFinished(msg message) {

	pm.inputQ <- msg

}
