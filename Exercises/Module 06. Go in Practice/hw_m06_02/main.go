package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	noData       = "No data"
	dataTemplate = "Overall status is %v, with service_id %v mysql component is %v"
)

func main() {
	urlFlag := flag.String("url", "", "URL of the server")
	flag.Parse()

	client := NewClient(*urlFlag)
	fmt.Println(client.data())
}

type Health struct {
	Status    string      `json:"status"`
	ServiceID string      `json:"service_id"`
	Checks    ServiceList `json:"checks"`
}

type ServiceList struct {
	PingMysql Service `json:"ping_mysql"`
}

type Service struct {
	ComponentID   string `json:"component_id"`
	ComponentType string `json:"component_type"`
	Status        string `json:"status"`
}

type Client struct {
	u string
}

func NewClient(u string) *Client {
	return &Client{u: u}
}

func (c *Client) data() string {
	client := http.DefaultClient

	resp, err := client.Get(c.u)
	if err != nil {
		return noData
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return noData
	}

	var health Health
	err = json.Unmarshal(body, &health)
	if err != nil {
		return noData
	}

	return fmt.Sprintf(dataTemplate, health.Status, health.ServiceID, health.Checks.PingMysql.Status)
}
