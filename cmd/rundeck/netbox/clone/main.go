package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"github.com/filprpx/eg-gotomation/internal/netbox"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	client, err := netbox.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	devices, err := client.Device.List(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if len(devices) == 0 {
		fmt.Print("No devices, stopping execution\n")
		return
	}

	fmt.Printf("Current number of devices: %d\n", len(devices))

	firstDevice := devices[0]
	fmt.Printf("Selecting first one id:%d, name: %s\n", firstDevice.Id, firstDevice.Name)

	firstDevice.Name = fmt.Sprintf("%s-%d", firstDevice.Name, rand.Intn(100000))
	firstDevice.Id = 0
	clonedDevice, err := client.Device.Create(ctx, &firstDevice)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Device cloned successfully, new device id: %d", clonedDevice.Id)
}
