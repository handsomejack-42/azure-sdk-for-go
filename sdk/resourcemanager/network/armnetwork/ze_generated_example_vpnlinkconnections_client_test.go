//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/VpnSiteLinkConnectionReset.json
func ExampleVPNLinkConnectionsClient_BeginResetConnection() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnetwork.NewVPNLinkConnectionsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginResetConnection(ctx,
		"<resource-group-name>",
		"<gateway-name>",
		"<connection-name>",
		"<link-connection-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/VpnSiteLinkConnectionGetIkeSas.json
func ExampleVPNLinkConnectionsClient_BeginGetIkeSas() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnetwork.NewVPNLinkConnectionsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginGetIkeSas(ctx,
		"<resource-group-name>",
		"<gateway-name>",
		"<connection-name>",
		"<link-connection-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.VPNLinkConnectionsClientGetIkeSasResult)
}
