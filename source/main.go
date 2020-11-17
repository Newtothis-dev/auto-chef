package main

import (
	"auto-chef/model"
	"fmt"
)

func main() {

	hd := model.HealthData{Carbs: 1, Sugar: 2, Fiber: 3, Fat: 4, SatFat: 6, Protein: 7, Calories: 12}
	fmt.Println(hd)

}
