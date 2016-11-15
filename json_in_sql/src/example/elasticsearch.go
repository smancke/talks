package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/olivere/elastic.v3"
	"strconv"
	"time"
)

func elasticsearchDemo() {

	client, err := elastic.NewClient()
	ch(err)

	client.DeleteIndex("demo").Do()

	_, err = client.CreateIndex("demo").Do()
	ch(err)

	i := 0
	bulk := client.Bulk()
	m := StartMeasure("insert cities")
	maxId = IterateCities(func(city *City) {
		bulk = bulk.Add(
			elastic.NewBulkIndexRequest().
				Index("demo").
				Type("demo").
				Id(strconv.Itoa(city.Id)).
				Doc(city),
		)
		m.Action()
		if i%2000 == 0 {
			_, err := bulk.Do()
			ch(err)
			//bulk = client.Bulk()
		}
		i++
	})
	_, err = bulk.Do()
	ch(err)
	m.End()

	// give es some time to index
	time.Sleep(time.Second * 3)

	m = StartMeasure("select by json Id")
	for i := 0; i < 1000; i++ {
		queryId := int(r.Int31n(int32(70000)))

		searchResult, err := client.Search().
			Index("demo").
			Query(elastic.NewTermQuery("Id", queryId)).
			From(0).Size(1).
			Do()
		ch(err)

		if len(searchResult.Hits.Hits) == 0 {
			panic(fmt.Sprintf("no results"))
		}
		hit := searchResult.Hits.Hits[0]
		var city City
		b, err := hit.Source.MarshalJSON()
		ch(err)
		ch(json.Unmarshal(b, &city))

		if city.Id != queryId {
			panic(fmt.Sprintf("received wrong id, expected %v, got %v", queryId, city.Id))
		}
		m.Action()
	}
	m.End()

}
