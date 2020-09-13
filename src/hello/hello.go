package main

import (
	"fmt"
	"math/rand"
)

const to_long = 62110000

func main() {
	fmt.Printf("%-18v %4v %-13v %5v \n", "Spaceline", "Days", "Trip type", "Price")
	fmt.Println("=========================================== \n")

	for count := 1; count <= 10; count++ {
		// 会社名
		company_name := "undefined"
		switch company_num := rand.Intn(3); company_num {
		case 0:
			company_name = "SpaceX"
		case 1:
			company_name = "Virgin Galactic"
		case 2:
			company_name = "Space Adventures"
		default:
			fmt.Println("company_name: Something wrong.")
		}

		// Trip Type
		trip_type := "undefined"
		if trip_num := rand.Intn(2); trip_num == 0 {
			trip_type = "One-way"
		} else if trip_num == 1 {
			trip_type = "Round-trip"
		} else {
			fmt.Println("trip_type: Something wrong.")
		}

		// 日数
		speed := rand.Intn(24) + 16
		cal_speed := speed * 60 * 60 * 60
		days := 0
		if trip_type == "One-way" {
			days = to_long / cal_speed
		} else if trip_type == "Round-trip" {
			days = to_long * 2 / cal_speed
		}

		// 金額
		price := 0
		cal_price := 20
		if trip_type == "One-way" {
			price = days + cal_price
		} else if trip_type == "Round-trip" {
			price = days + (cal_price * 2)
		}

		// チケット発行
		fmt.Printf("%-18v %4v %-13v $ %-3v \n", company_name, days, trip_type, price)
	}
}
