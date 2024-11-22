package main

import (
	"fmt"
	"processappsimple/process"
)

func main() {

	// Main Folder: /Volumes/Data Disk/Process App Data/
	//
	// Sub Folders:
	//    - VO1 : 1M products, 1M customers with zips, 50M purchases
	//    - V02 : in progress
	//
	const DATA_FOLDER_PATH = "/Volumes/Data Disk/Process App Data/Test/"

	const PSNK_ZIP = "15068"
	const SEATTLE_ZIP = "98109"

	homeZip := SEATTLE_ZIP

	processMaster := process.CreateProcessMaster(DATA_FOLDER_PATH, homeZip)

	zipCode, distance := processMaster.FindFarthestZipInMiles()

	fmt.Println("Zip code", zipCode, "is the farthest at", distance, "miles away from", homeZip)

}
