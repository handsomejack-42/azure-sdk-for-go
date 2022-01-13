//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armwindowsesu_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/windowsesu/armwindowsesu"
)

// x-ms-original-file: specification/windowsesu/resource-manager/Microsoft.WindowsESU/preview/2019-09-16-preview/examples/ListMultipleActivationKeys.json
func ExampleMultipleActivationKeysClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armwindowsesu.NewMultipleActivationKeysClient("<subscription-id>", cred, nil)
	pager := client.List(nil)
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

// x-ms-original-file: specification/windowsesu/resource-manager/Microsoft.WindowsESU/preview/2019-09-16-preview/examples/ListMultipleActivationKeysByResourceGroup.json
func ExampleMultipleActivationKeysClient_ListByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armwindowsesu.NewMultipleActivationKeysClient("<subscription-id>", cred, nil)
	pager := client.ListByResourceGroup("<resource-group-name>",
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

// x-ms-original-file: specification/windowsesu/resource-manager/Microsoft.WindowsESU/preview/2019-09-16-preview/examples/GetMultipleActivationKey.json
func ExampleMultipleActivationKeysClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armwindowsesu.NewMultipleActivationKeysClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<multiple-activation-key-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.MultipleActivationKeysClientGetResult)
}

// x-ms-original-file: specification/windowsesu/resource-manager/Microsoft.WindowsESU/preview/2019-09-16-preview/examples/CreateMultipleActivationKey.json
func ExampleMultipleActivationKeysClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armwindowsesu.NewMultipleActivationKeysClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<multiple-activation-key-name>",
		armwindowsesu.MultipleActivationKey{
			Location: to.StringPtr("<location>"),
			Properties: &armwindowsesu.MultipleActivationKeyProperties{
				AgreementNumber:       to.StringPtr("<agreement-number>"),
				InstalledServerNumber: to.Int32Ptr(100),
				IsEligible:            to.BoolPtr(true),
				OSType:                armwindowsesu.OsType("WindowsServer2008").ToPtr(),
				SupportType:           armwindowsesu.SupportType("SupplementalServicing").ToPtr(),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.MultipleActivationKeysClientCreateResult)
}

// x-ms-original-file: specification/windowsesu/resource-manager/Microsoft.WindowsESU/preview/2019-09-16-preview/examples/UpdateMultipleActivationKey.json
func ExampleMultipleActivationKeysClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armwindowsesu.NewMultipleActivationKeysClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<multiple-activation-key-name>",
		armwindowsesu.MultipleActivationKeyUpdate{
			Tags: map[string]*string{
				"tag1": to.StringPtr("value1"),
				"tag2": to.StringPtr("value2"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.MultipleActivationKeysClientUpdateResult)
}

// x-ms-original-file: specification/windowsesu/resource-manager/Microsoft.WindowsESU/preview/2019-09-16-preview/examples/DeleteMultipleActivationKey.json
func ExampleMultipleActivationKeysClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armwindowsesu.NewMultipleActivationKeysClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<multiple-activation-key-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
