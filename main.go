package main

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

var EndPoint = "https://wakatime.com/api/v1/"
var DBName = "wakatime.db"
var CompressedDBName = "wakatime.tar.gz"

func getAuthHeader(auth string) http.Header {
	header := http.Header{}
	base64EncodedAuth := base64.StdEncoding.EncodeToString([]byte(auth))
	header.Add("Authorization", "Basic "+base64EncodedAuth)
	return header
}

func requestEndPoint(method string, endPoint string, data url.Values, auth string) (*http.Response, error) {
	link := EndPoint + endPoint + "?" + data.Encode()
	req, _ := http.NewRequest(method, link, nil)
	req.Header = getAuthHeader(auth)
	client := &http.Client{}
	resp, err := client.Do(req)
	return resp, err
}

func main() {
	if len(os.Args) < 2 {
		println("Missing argument: <api-key>")
		os.Exit(1)
	}
	auth := os.Args[1]
	data := url.Values{}
	data.Add("date", time.Now().AddDate(0, 0, -1).Format("2006-01-02"))
	resp, err := requestEndPoint("GET", "users/current/heartbeats", data, auth)
	log.Print("Requesting: ", EndPoint+"users/current/heartbeats?"+data.Encode())
	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode != 200 {
		log.Fatalf("Request failed with status code: %d, data: %s", resp.StatusCode, string(body))
	}

	var respData HeartBeatResp
	err = json.Unmarshal(body, &respData)
	if err != nil {
		panic(err)
	}
	log.Print("Got " + strconv.Itoa(len(respData.Data)) + " heartbeats for yesterday")

	err = resp.Body.Close()
	if err != nil {
		panic(err)
	}

	sqliteMethod(respData)
	log.Print("Finished, have a nice day!")
}

func sqliteMethod(respData HeartBeatResp) {
	if _, err := os.Stat(CompressedDBName); err == nil {
		err = decompressFile(CompressedDBName, DBName)
		if err != nil {
			panic(err)
		}
	} else {
		log.Print("Database file not found, creating new one")
	}

	db := connectToDB(DBName)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	createTableIfNotExists(db)
	insertHeartBeat(db, respData)
	if err := db.Close(); err != nil {
		panic(err)
	}

	err := compressFile(DBName, CompressedDBName)
	if err != nil {
		panic(err)
	}

	err = os.Remove(DBName)
	if err != nil {
		panic(err)
	}
}
