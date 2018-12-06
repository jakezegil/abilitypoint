package main

import (
	"encoding/json"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"io/ioutil"
	"os"
)

func main() {
	var (
		driver  neo4j.Driver
		session neo4j.Session
		result  neo4j.Result
		err     error
	)

	/* CONFIGURATION CODE */
	jsonFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var configuration map[string]interface{}
	json.Unmarshal([]byte(byteValue), &configuration)

	var base_url string
	if str, ok := configuration["base_url"].(string); ok {
		base_url = str
	} else {
		fmt.Println("Can't start server because database url wasn't specified in config!")
		os.Exit(1)
	}
	var user_name string
	if str, ok := configuration["base_url"].(string); ok {
		user_name = str
	} else {
		fmt.Println("Can't start server because database username wasn't specified in config!")
		os.Exit(1)
	}
	var password string
	if str, ok := configuration["base_url"].(string); ok {
		password = str
	} else {
		fmt.Println("Can't start server because database password wasn't specified in config!")
		os.Exit(1)
	}
	/* ******************  */

	if driver, err = neo4j.NewDriver("bolt://"+base_url+":7687", neo4j.BasicAuth(user_name, password, "")); err != nil {
		fmt.Println("ERROR")
	}
	// Used to destroy driver after calls
	defer driver.Close()

	if session, err = driver.Session(neo4j.AccessModeWrite); err != nil {
		fmt.Println("ERROR")
	}
	defer session.Close()

	result, err = session.Run("MATCH (n:Fact) RETURN n.description", nil)
	if err != nil {
		fmt.Println("ERROR")
	}

	for result.Next() {
		fmt.Printf("'%s'", result.Record().GetByIndex(0).(string))
	}
	if err = result.Err(); err != nil {
		fmt.Println("ERROR")
	}
}
