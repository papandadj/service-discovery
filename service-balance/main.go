package main

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"time"

	client "go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/mvcc/mvccpb"
	"google.golang.org/grpc"
)

var (
	mailURL     string
	userlistURL string
	cli         *client.Client
	wg          sync.WaitGroup
	router      map[string]string
)

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
	router = make(map[string]string)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	watchChan := cli.Watch(ctx, "service/", []client.OpOption{client.WithPrefix()}...)

	wg.Add(1)
	go loopWatch(watchChan, &wg)

	http.HandleFunc("/mail", mailHandler)
	http.HandleFunc("/userlist", userHandler)
	err := http.ListenAndServe(":5050", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(10)
	}
	wg.Wait()
}

func checkHandler(resp http.ResponseWriter, req *http.Request, service string) (string, error) {
	value, ok := router[service]
	if !ok {
		resp.WriteHeader(411)
		resp.Write([]byte("Service is not open."))
		return "", errors.New("error")
	}
	return value, nil
}

func mailHandler(resp http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	url, err := checkHandler(resp, req, "service/mail")
	if err != nil {
		return
	}

	response, err := http.Post(url, "application/json", req.Body)
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	resp.WriteHeader(response.StatusCode)
	resp.Write(body)
}

func userHandler(resp http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	url, err := checkHandler(resp, req, "service/userlist")
	if err != nil {
		return
	}

	response, err := http.Post(url, "application/json", req.Body)

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	resp.WriteHeader(response.StatusCode)
	resp.Write(body)
}

func loopWatch(watchChan client.WatchChan, wg *sync.WaitGroup) {
	for {
		select {
		case watchResp := <-watchChan:
			if watchResp.Events[0].Type == mvccpb.PUT {
				router[string(watchResp.Events[0].Kv.Key)] = string(watchResp.Events[0].Kv.Value)
			} else if watchResp.Events[0].Type == mvccpb.DELETE {
				delete(router, string(watchResp.Events[0].Kv.Key))
			}
		}
	}
}
