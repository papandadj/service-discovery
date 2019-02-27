package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	client "go.etcd.io/etcd/clientv3"
	"google.golang.org/grpc"
)

var cli *client.Client

func init() {
	cli1, err := client.New(client.Config{
		Endpoints:   []string{"etcd1:2379", "etcd2:2379", "etcd3:2379"},
		DialTimeout: 5 * time.Second,
		DialOptions: []grpc.DialOption{grpc.WithBlock()},
	})
	cli = cli1
	if err != nil {
		fmt.Println(err)
		os.Exit(10)
	}
}

func main() {
	defer cli.Close()
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(resp http.ResponseWriter, req *http.Request) {
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
	_ = etcdPut(bodySt.Name, "http://"+bodySt.Host+":"+strconv.Itoa(bodySt.Port))
	resp.WriteHeader(200)
}

func etcdPut(key, value string) error {
	leaseGrantRes, err := cli.Grant(context.TODO(), 5)
	if err != nil {
		fmt.Println(err)
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), deadline)
	fmt.Println("Client put key:", key, ",value:", value, " on etcd")
	_, err = cli.Put(ctx, key, value, client.WithLease(leaseGrantRes.ID))
	defer cancel()
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
