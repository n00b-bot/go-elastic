package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go-elastic/model"
	"go-elastic/utils"
	"log"
	"time"

	"github.com/olivere/elastic/v7"
)

var urlElastic = "http://10.14.140.226:9200"
var indexElastic = "wase-burp"

func main() {
	client, err := elastic.NewClient(elastic.SetURL(urlElastic))
	if err != nil {
		log.Println(err)
		return
	}
	ctx := context.Background()
	_, _, err = client.Ping(urlElastic).Do(ctx)
	if err != nil {
		log.Println(err)
		return
	}
	searchSource := elastic.NewSearchSource()
	//searchSource.Query(elastic.NewMatchQuery("port", "443"))
	searchSource.Query(elastic.NewBoolQuery().Filter(elastic.NewRangeQuery("timestamp").From(time.Now().Add(-time.Hour * 24)).To(time.Now())))
	searchService := client.Search().Index(indexElastic).SearchSource(searchSource)
	searchResult, err := searchService.Do(ctx)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(searchResult.TotalHits())
	searchService.Size(int(searchResult.TotalHits()))
	searchResult, err = searchService.Do(ctx)
	if err != nil {
		log.Println(err)
		return
	}
	//fmt.Println(searchResult.TotalHits())
	//fmt.Println(len(searchResult.Hits.Hits))

	newPath := make(map[string]bool)
	for _, hit := range searchResult.Hits.Hits {
		var http model.Http
		//fmt.Println(string(hit.Source))
		err := json.Unmarshal(hit.Source, &http)
		if err != nil {
			log.Println(err)
			return
		}
		//fmt.Println("line" + http.Request.RequestLine)
		if url := utils.ParseUrl(http.Request.RequestLine); url != nil {
			for _, u := range url {
				if u != "" {
					if !utils.CheckUrlExists(u, "file.txt") {
						newPath[u] = true
					}

				}
			}
		}

	}
	for k, _ := range newPath {
		fmt.Println(k)
	}

}
