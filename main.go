package main

import (
	_ "github.com/sms2sakthivel/order-manager/docs"
	"github.com/sms2sakthivel/order-manager/orders"
)

func main() {
	// Step 1: Create a New Order Service Application
	app := orders.NewApp()

	// Step 2: Start Server and Listen on the Port 8001
	app.Listen(":8003")
}
