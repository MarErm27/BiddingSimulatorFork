package MarErm

import (
	"fmt"
	"strconv"
	"strings"
)

//PriceToString converts int64 to string
func PriceToString(price int64) string {
	s := strconv.FormatInt(price/100, 10) + "." + strconv.FormatInt(price%100, 10)
	return s
}

//PriceToInt64 converts string to int64
func PriceToInt64(price string) int64 {
	s := strings.Replace(price, ".", "", -1)
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		fmt.Printf("%d of type %T", n, n)
	}
	return n
}
