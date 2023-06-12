package gohttp

import "fmt"

type HttpClient interface {
	Get()
	Post()
	Put()
	Delete()
	Patch()
	Options()
}

func New() HttpClient {
	client := &Client{}
	return client
}

type Client struct{}

func (c *Client) Get() {
	fmt.Println("Get method")
}

func (c *Client) Post() {
	fmt.Println("Post method")
}

func (c *Client) Put() {
	fmt.Println("Put method")
}

func (c *Client) Delete() {
	fmt.Println("Delete method")
}
func (c *Client) Patch() {
	fmt.Println("Patch method")
}

func (c *Client) Options() {
	fmt.Println("Option method")
}
