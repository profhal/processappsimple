package main

import (
	"fmt"
	"os"
	"processappsimple/process"
	"processappsimple/utilities"
)

func main() {

	// The path the folder with the data files.
	//
	const DATA_FOLDER_PATH = "../data/"

	// This is the number of process chains. It should divide the number of product ids evenly
	// or there might be some uncertain behavior.
	//
	// The assumption is that the number of product ids ends in a number wit three zeros.
	//
	// The initial app was built with 1,000,000 product ids so 40 works.
	//
	const CHAIN_COUNT = 40

	// The home base zip code that is used to determine the farthest zip an item
	// was shipped to.
	//
	// Examples: PSNK : 15068
	//           Belle Fourche, SD : 57717 (most centered post office of all 50 US states)
	//           Lebanon, KY : 40033 (most centered post office of the contiguous 48 states)
	//           Honolulu, HI: 96898
	//           Seattle, WA : 98109
	//			 San Diego, CA: 92108
	//
	homeZip := "40033"

	args := os.Args

	if len(args) > 1 {

		homeZip = args[1]

	}

	zipCodeUtil, _ := utilities.GetZipCodeUtilInstance()

	processMaster := process.CreateProcessMaster(DATA_FOLDER_PATH, CHAIN_COUNT, homeZip)

	zipCode, distance := processMaster.FindFarthestZipInMiles()

	farthestCityState, _ := zipCodeUtil.GetCityState(zipCode)
	homeBaseCityState, _ := zipCodeUtil.GetCityState(homeZip)

	fmt.Println("Zip code",
		zipCode,
		"("+farthestCityState+")",
		"is the farthest at",
		distance,
		"miles away from",
		homeZip,
		"("+homeBaseCityState+")")

}
