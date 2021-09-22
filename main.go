package main

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

func main() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		panic("no cred")
	}

	conn := newConnection("...", cred, nil)
	if conn == nil {
		panic("no conn")
	}

	client := entityClient{con: conn}
	result, err := client.Get(context.Background(), "entity-name", nil)
	if err != nil {
		log.Panicf("service bus error: %v", err)
	}
	log.Print(result)
}
