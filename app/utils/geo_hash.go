package utils

import "github.com/mmcloughlin/geohash"

func GenerateGeohash(latitude float64, longitude float64, precision int) string {
	return geohash.EncodeWithPrecision(latitude, longitude, uint(precision))
}
