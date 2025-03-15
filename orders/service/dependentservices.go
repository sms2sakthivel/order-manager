package service

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v3/client"
	"github.com/sms2sakthivel/order-manager/orders/model"
)

func GetUserByID(id uint) (*model.UserResponse, error) {
	cli := client.New()
	cli.SetTimeout(10 * time.Second)

	// Send a GET request
	resp, err := cli.Get(fmt.Sprintf("http://127.0.0.1:8001/users/%d", id))
	if err != nil {
		return nil, err
	}

	fmt.Printf("Status: %d\n", resp.StatusCode())
	fmt.Printf("Body: %s\n", string(resp.Body()))
	if resp.StatusCode() != 200 {
		return nil, err
	}
	return nil, nil
}

func GetProductByID(id uint) (*model.ProductResponse, error) {
	cli := client.New()
	cli.SetTimeout(10 * time.Second)

	// Send a GET request
	resp, err := cli.Get(fmt.Sprintf("http://127.0.0.1:8002/products/%d", id))
	if err != nil {
		return nil, err
	}

	fmt.Printf("Status: %d\n", resp.StatusCode())
	fmt.Printf("Body: %s\n", string(resp.Body()))
	if resp.StatusCode() != 200 {
		return nil, err
	}
	var productResponse model.ProductResponse
	err = json.Unmarshal(resp.Body(), &productResponse)
	if err != nil {
		return nil, err
	}
	return &productResponse, nil
}