package main

type mysqlConfig struct {
	MysqlHost        string `json:"MysqlHost"`
	MysqlUser        string `json:"MysqlUser"`
	MysqlPort        int    `json:"MysqlPort"`
	MysqlPassword    string `json:"MysqlPassword"`
	MysqlDB          string `json:"MysqlDB"`
	Port             int    `json:"Port"`
	EtcdRegistryHost string `json:"EtcdRegistryHost"`
	EtcdRegistryPort int    `json:"EtcdRegistryPort"`
}

type requestBody struct {
	Method string `json:"method"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}
type responseBody struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
