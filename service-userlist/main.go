package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	configuration mysqlConfig
	client        *sql.DB
)

func init() {
	data, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println("Read config file error.")
		os.Exit(10)
	}

	err = json.Unmarshal(data, &configuration)
	if err != nil {
		fmt.Println("Config.json formation error : ", err)
		os.Exit(10)
	}
	mysqlPath := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", configuration.MysqlUser, configuration.MysqlPassword, configuration.MysqlHost, configuration.MysqlPort, configuration.MysqlDB)
	mysqlPath = mysqlPath + "?parseTime=true&charset=utf8&loc=Asia%2FShanghai"
	db, err := sql.Open("mysql", mysqlPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(10)
	}
	client = db
}

func main() {
	defer client.Close()
	go registry()
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(configuration.Port), nil))
}

func handler(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("into")
	defer req.Body.Close()
	if req.Method != "POST" {
		resp.WriteHeader(410)
		resp.Write([]byte("Method must be post"))
		return
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
	}

	var bodySt requestBody
	json.Unmarshal(body, &bodySt)

	if bodySt.Method == "put" {
		rows, err := client.Query("INSERT INTO user (name, email) VALUES (?, ?);", bodySt.Name, bodySt.Email)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer rows.Close()
		resp.WriteHeader(200)
		return
	} else if bodySt.Method == "del" {
		rows, err := client.Query("DELETE FROM user WHERE name=?;", bodySt.Name)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer rows.Close()
		resp.WriteHeader(200)
		return
	} else if bodySt.Method == "get" {
		rows, err := client.Query("SELECT * FROM user WHERE name=?;", bodySt.Name)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer rows.Close()
		var id int
		var name string
		var email string
		var respBody responseBody
		for rows.Next() {
			err = rows.Scan(&id, &name, &email)
			if err != nil {
				fmt.Println(err)
				return
			}
			respBody.ID = id
			respBody.Name = name
			respBody.Email = email
		}
		resp.WriteHeader(200)
		bodyByte, _ := json.Marshal(&respBody)
		resp.Write(bodyByte)
		return
	}
	resp.WriteHeader(410)
	resp.Write([]byte("Paramater error."))
}

func registry() {
	for {
		time.Sleep(2 * time.Second)
		url := configuration.EtcdRegistryHost + ":" + strconv.Itoa(configuration.EtcdRegistryPort)
		var postBody = []byte(`{"name":"service/userlist", "host":"userlist", "port":6060}`)
		req, _ := http.NewRequest("POST", url, bytes.NewBuffer(postBody))
		req.Header.Set("Content-Type", "application/json")

		reqClient := &http.Client{}
		resp, err := reqClient.Do(req)
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer resp.Body.Close()

		// fmt.Println(resp)

	}
}
