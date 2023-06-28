package main

import (
	"fmt"
	"time"
)

func main() {
	monthArr := [12]string{"Январь", "Февраль", "Март", "Апрель", "Май", "Июнь", "Июль", "Август", "Сентябрь", "Октябрь", "Ноябрь", "Декабрь"}

	data := time.Now()
	month := int(data.Month()) - 1
	monthStr := monthArr[month]

	fmt.Println(monthStr)
}
