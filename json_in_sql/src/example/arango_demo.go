package main

import (
	"fmt"
	ara "github.com/diegogub/aranGO"
)

func aragnoDemo() {

	s, err := ara.Connect("http://localhost:8529", "root", "demo", false)
	if err != nil {
		panic(err)
	}
	s.CreateDB("demo", nil)

	if !s.DB("demo").ColExist("cities") {
		cities := ara.NewCollectionOptions("cities", true)
		ch(s.DB("demo").CreateCollection(cities))
	}

	cities := s.DB("demo").Col("cities")
	m := StartMeasure("insert cities")
	maxId = IterateCities(func(city *City) {
		ch(cities.Save(city))
		m.Action()
	})
	m.End()

	err = s.DB("demo").Col("cities").CreateHash(false, "Id")
	if err != nil {
		panic(err)
	}

	m = StartMeasure("select by json Id")
	for i := 0; i < 1000; i++ {
		queryId := int(r.Int31n(int32(70000)))
		query := fmt.Sprintf(`FOR i in cities FILTER i.Id == %d RETURN i`, queryId)
		q := ara.NewQuery(query)
		c, err := s.DB("demo").Execute(q)
		if err != nil {
			panic(err)
		}
		var city City
		c.FetchOne(&city)
		if city.Id != queryId {
			panic(fmt.Sprintf("received wrong id, expected %v, got %v", city.Id, queryId))
		}
		m.Action()
	}
	m.End()

}
