package utilities

import (
	"github.com/danmichaeli1/zipcodes"
)

type ZipCodeUtil struct {
	zipCodes *zipcodes.Zipcodes
}

func CreateZipCodeUtil(zipCodeFilepath string) (zipUtil *ZipCodeUtil) {

	zipUtil = new(ZipCodeUtil)

	zipUtil.zipCodes, _ = zipcodes.New(zipCodeFilepath)

	return zipUtil

}

func (z *ZipCodeUtil) GetDistanceInMiles(fromZip string, toZip string) (float64, error) {

	return z.zipCodes.DistanceInMiles(fromZip, toZip)

}
