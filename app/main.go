package main

import (
	"fmt"
	"processappsimple/process"
)

func main() {

	// The path the folder with the data files.
	//
	const DATA_FOLDER_PATH = "/Volumes/Data Disk/Process App Data/Test/"

	// This is the number of process chains. It should divide the number of product ids evenly
	// or there might be some uncertain behavior.
	//
	// The assumption is that the number of product ids ends in a number wit three zeros.
	//
	// The initial app was built with 1,000,000 product ids so 40 works.
	//
	const CHAIN_COUNT = 40

	const PSNK_ZIP = "15068"
	const SEATTLE_ZIP = "98109"

	homeZip := SEATTLE_ZIP

	processMaster := process.CreateProcessMaster(DATA_FOLDER_PATH, CHAIN_COUNT, homeZip)

	zipCode, distance := processMaster.FindFarthestZipInMiles()

	fmt.Println("Zip code", zipCode, "is the farthest at", distance, "miles away from", homeZip)

}
