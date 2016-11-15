package main

import (
	"bufio"
	"os"
	"strings"
)

type City struct {
	Id         int
	Country    string
	City       string
	AccentCity string
	Region     string
	Population string
	Latitude   string
	Longitude  string
}

func IterateCities(callback func(city *City)) int {
	file, err := os.Open("german_cities.csv")
	//file, err := os.Open("3000.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	id := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		record := strings.Split(line, ",")
		if len(record) != 7 {
			panic("error rading line: " + line)
		}

		city := &City{}
		city.Id = id
		city.Country = record[0]
		city.City = record[1]
		city.AccentCity = record[2]
		city.Region = record[3]
		city.Population = record[4]
		city.Latitude = record[5]
		city.Longitude = record[6]

		callback(city)
		id++
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return id
}
