package utilities

import (
	"errors"
	"fmt"
	"sync"

	"github.com/danmichaeli1/zipcodes"
)

const zip_code_filepath = "../data/zipCodes.txt"

type ZipCodeUtil struct {
	zipCodes *zipcodes.Zipcodes
}

var lock = &sync.Mutex{}
var zipCodeUtilInstance *ZipCodeUtil

func GetZipCodeUtilInstance() (*ZipCodeUtil, error) {

	var theUtil *ZipCodeUtil
	var theError error

	if zipCodeUtilInstance == nil {

		lock.Lock()

		defer lock.Unlock()

		if zipCodeUtilInstance == nil {

			zipCodeUtilInstance = new(ZipCodeUtil)

			zipCodes, err := zipcodes.New(zip_code_filepath)

			if err != nil {

				fmt.Println("**ERROR** zipcodeUtil.GetInstance():", err)

				theUtil = nil
				theError = errors.New("unable to load zip codes from " + zip_code_filepath)

			} else {

				zipCodeUtilInstance.zipCodes = zipCodes

				theUtil = zipCodeUtilInstance
				theError = nil

			}

		}

	} else {

		theUtil = zipCodeUtilInstance
		theError = nil

	}

	return theUtil, theError

}

func (z *ZipCodeUtil) GetDistanceInMiles(fromZip string, toZip string) (float64, error) {

	return z.zipCodes.DistanceInMiles(fromZip, toZip)

}

func (z *ZipCodeUtil) GetCityState(zip string) (string, error) {

	location, err := z.zipCodes.Lookup(zip)

	if err == nil {

		return location.PlaceName + ", " + location.State, nil

	} else {

		fmt.Println("**ERROR** utilities.GetCityState():", err)

		return "", err

	}

}
