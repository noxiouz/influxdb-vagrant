package main

import (
	"encoding/json"
	"log"

	"github.com/influxdb/influxdb/client"
)

const (
	host     = "localhost:8286"
	username = "root"
	password = "root"
	database = "TEST"
)

const (
	asJson = `
{
    "name": "noxiouz_test_as_json",
    "columns": ["a", "b"],
    "points": [[1, 2], [3, 4]]
}
`
)

func main() {
	config := client.ClientConfig{}
	config.Host = host
	config.Password = password
	config.Username = username
	config.Database = database

	cl, err := client.NewClient(&config)
	if err != nil {
		log.Fatal(err)
	}
	// check connection
	log.Print(cl.Ping())

	//Create database
	if err := cl.CreateDatabase(database); err != nil {
		log.Fatalf("unabele to create database %s: %s", database, err)
	}

	shardSpace := client.ShardSpace{
		// required, must be unique within the database
		Name: database,
		// required, a database has many shard spaces and a shard space belongs to a database
		Regex:    "/.*/",
		Database: database,
		// this is optional, if they don't set it, it will default to the storage.dir in the config
		ReplicationFactor: 3,
	}

	if err := cl.CreateShardSpace(database, &shardSpace); err != nil {
		log.Fatalf("unable to create shard space for database %s: %s", database, err)
	}

	// To create continues series. It should be done only once.
	// This example means put into MYAGGNAME value evaluated as 95% of `a` grouped by 1 minute
	// It's possible to get falues as SELECT * FROM MYAGGNAME
	_, err = cl.Query("select percentile(a, 95) from noxiouz_test_as_json group by time(1m) into MYAGGNAME")
	if err != nil {
		log.Fatal(err)
	}

	// One way t write point:
	// desribe one point as go struct
	series := client.Series{
		// name of collection
		Name: "noxiouz_test",
		// name of values
		// SELECT min1 from noxiouz_test where time > '2014-10-12 23:32:01.232'
		Columns: []string{"min1", "min5", "min15"},
		// Two points.
		Points: [][]interface{}{
			[]interface{}{1, 100, 200},
			[]interface{}{2, 4, 6}},
	}

	// The second way:
	// decode point from json
	var series_from_json client.Series
	if err := json.Unmarshal([]byte(asJson), &series_from_json); err != nil {
		log.Fatalf("Unmarshal error: %s", err)
	}

	if err := cl.WriteSeries([]*client.Series{&series, &series_from_json}); err != nil {
		log.Fatalf("WriteSeries: %s", err)
	}
}
