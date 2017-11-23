package main

import (
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	uuid2 "github.com/hashicorp/go-uuid"
	"math/rand"
	"time"
)

func main() {
	config := consulapi.DefaultConfig()
	go register(config)
	time.Sleep(1000 * time.Second)

}

func register(config *consulapi.Config) {
	consul, _ := consulapi.NewClient(config)
	tags := []string{"foo", "test"}
	uuid, _ := uuid2.GenerateUUID()

	ip := fmt.Sprintf("%d.%d.%d,%d", rand.Intn(254), rand.Intn(254), rand.Intn(254), rand.Intn(254))
	register := consulapi.CatalogRegistration{
		ID: uuid,
		Service: &consulapi.AgentService{
			ID:      "consul",
			Tags:    tags,
			Port:    80,
			Address: ip,
			Service: "consul",
		},
		Check: &consulapi.AgentCheck{
			CheckID:     "3",
			Status:      "passing",
			ServiceID:   "consul",
			ServiceName: "consul",
		},

		Node:    fmt.Sprintf("test%d", rand.Intn(25400)),
		Address: "127.0.0.2",
	}
	catalog := consul.Catalog()
	for {
		println("passing")
		register.Check.Status = "passing"
		catalog.Register(&register, &consulapi.WriteOptions{})
		time.Sleep(10 * time.Second)
		register.Check.Status = "failing"
		println("failing")
		catalog.Register(&register, &consulapi.WriteOptions{})
		time.Sleep(20 * time.Second)

	}

}
