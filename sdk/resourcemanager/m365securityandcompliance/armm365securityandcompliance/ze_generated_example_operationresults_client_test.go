//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armm365securityandcompliance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/m365securityandcompliance/armm365securityandcompliance"
)

// x-ms-original-file: specification/m365securityandcompliance/resource-manager/Microsoft.M365SecurityAndCompliance/preview/2021-03-25-preview/examples/OperationResultsGet.json
func ExampleOperationResultsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armm365securityandcompliance.NewOperationResultsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<location-name>",
		"<operation-result-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.OperationResultsClientGetResult)
}
