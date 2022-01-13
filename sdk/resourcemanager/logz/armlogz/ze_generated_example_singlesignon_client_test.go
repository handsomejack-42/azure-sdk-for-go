//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armlogz_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logz/armlogz"
)

// x-ms-original-file: specification/logz/resource-manager/Microsoft.Logz/stable/2020-10-01/examples/SingleSignOnConfigurations_List.json
func ExampleSingleSignOnClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armlogz.NewSingleSignOnClient("<subscription-id>", cred, nil)
	pager := client.List("<resource-group-name>",
		"<monitor-name>",
		nil)
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}

// x-ms-original-file: specification/logz/resource-manager/Microsoft.Logz/stable/2020-10-01/examples/SingleSignOnConfigurations_CreateOrUpdate.json
func ExampleSingleSignOnClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armlogz.NewSingleSignOnClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<monitor-name>",
		"<configuration-name>",
		&armlogz.SingleSignOnClientBeginCreateOrUpdateOptions{Body: nil})
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.SingleSignOnClientCreateOrUpdateResult)
}

// x-ms-original-file: specification/logz/resource-manager/Microsoft.Logz/stable/2020-10-01/examples/SingleSignOnConfigurations_Get.json
func ExampleSingleSignOnClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armlogz.NewSingleSignOnClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<monitor-name>",
		"<configuration-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.SingleSignOnClientGetResult)
}
