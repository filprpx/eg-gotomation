package main

import (
	"context"
	"fmt"
	"github.com/filprpx/eg-gotomation/internal/netbox"
	_ "github.com/joho/godotenv/autoload"
	"log"
)

func main() {
	client, err := netbox.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	devices, err := client.DCIM.Device.List(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if len(devices) == 0 {
		fmt.Print("No devices, stopping execution\n")
		return
	}

	fmt.Printf("Current number of devices: %d\n", len(devices))

	lastDevice := devices[len(devices)-1]
	fmt.Printf("Selecting last one id:%d, name: %s", lastDevice.Id, lastDevice.Name)

	ok, err := client.DCIM.Device.Delete(ctx, &lastDevice)
	if err != nil {
		log.Fatal(err)
	}

	if ok {
		fmt.Printf("Device deleted successfully")
	}
}
