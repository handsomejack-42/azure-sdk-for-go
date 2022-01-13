//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstorage_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-08-01/examples/BlobContainersList.json
func ExampleBlobContainersClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armstorage.NewBlobContainersClient("<subscription-id>", cred, nil)
	pager := client.List("<resource-group-name>",
		"<account-name>",
		&armstorage.BlobContainersClientListOptions{Maxpagesize: nil,
			Filter:  nil,
			Include: nil,
		})
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

// x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-08-01/examples/BlobContainersPutDefaultEncryptionScope.json
func ExampleBlobContainersClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armstorage.NewBlobContainersClient("<subscription-id>", cred, nil)
	res, err := client.Create(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<container-name>",
		armstorage.BlobContainer{
			ContainerProperties: &armstorage.ContainerProperties{
				DefaultEncryptionScope:      to.StringPtr("<default-encryption-scope>"),
				DenyEncryptionScopeOverride: to.BoolPtr(true),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.BlobContainersClientCreateResult)
}

// x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-08-01/examples/BlobContainersPatch.json
func ExampleBlobContainersClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armstorage.NewBlobContainersClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<container-name>",
		armstorage.BlobContainer{
			ContainerProperties: &armstorage.ContainerProperties{
				Metadata: map[string]*string{
					"metadata": to.StringPtr("true"),
				},
				PublicAccess: armstorage.PublicAccessContainer.ToPtr(),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.BlobContainersClientUpdateResult)
}

// x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-08-01/examples/BlobContainersGetWithAllowProtectedAppendWritesAll.json
func ExampleBlobContainersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armstorage.NewBlobContainersClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<container-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.BlobContainersClientGetResult)
}

// x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-08-01/examples/BlobContainersDelete.json
func ExampleBlobContainersClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armstorage.NewBlobContainersClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<container-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-08-01/examples/BlobContainersSetLegalHold.json
func ExampleBlobContainersClient_SetLegalHold() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armstorage.NewBlobContainersClient("<subscription-id>", cred, nil)
	res, err := client.SetLegalHold(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<container-name>",
		armstorage.LegalHold{
			Tags: []*string{
				to.StringPtr("tag1"),
				to.StringPtr("tag2"),
				to.StringPtr("tag3")},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.BlobContainersClientSetLegalHoldResult)
}

// x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-08-01/examples/BlobContainersClearLegalHold.json
func ExampleBlobContainersClient_ClearLegalHold() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armstorage.NewBlobContainersClient("<subscription-id>", cred, nil)
	res, err := client.ClearLegalHold(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<container-name>",
		armstorage.LegalHold{
			Tags: []*string{
				to.StringPtr("tag1"),
				to.StringPtr("tag2"),
				to.StringPtr("tag3")},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.BlobContainersClientClearLegalHoldResult)
}

// x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-08-01/examples/BlobContainersPutImmutabilityPolicy.json
func ExampleBlobContainersClient_CreateOrUpdateImmutabilityPolicy() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armstorage.NewBlobContainersClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdateImmutabilityPolicy(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<container-name>",
		&armstorage.BlobContainersClientCreateOrUpdateImmutabilityPolicyOptions{IfMatch: nil,
			Parameters: &armstorage.ImmutabilityPolicy{
				Properties: &armstorage.ImmutabilityPolicyProperty{
					AllowProtectedAppendWrites:            to.BoolPtr(true),
					ImmutabilityPeriodSinceCreationInDays: to.Int32Ptr(3),
				},
			},
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.BlobContainersClientCreateOrUpdateImmutabilityPolicyResult)
}

// x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-08-01/examples/BlobContainersGetImmutabilityPolicy.json
func ExampleBlobContainersClient_GetImmutabilityPolicy() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armstorage.NewBlobContainersClient("<subscription-id>", cred, nil)
	res, err := client.GetImmutabilityPolicy(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<container-name>",
		&armstorage.BlobContainersClientGetImmutabilityPolicyOptions{IfMatch: nil})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.BlobContainersClientGetImmutabilityPolicyResult)
}

// x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-08-01/examples/BlobContainersDeleteImmutabilityPolicy.json
func ExampleBlobContainersClient_DeleteImmutabilityPolicy() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armstorage.NewBlobContainersClient("<subscription-id>", cred, nil)
	res, err := client.DeleteImmutabilityPolicy(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<container-name>",
		"<if-match>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.BlobContainersClientDeleteImmutabilityPolicyResult)
}

// x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-08-01/examples/BlobContainersLockImmutabilityPolicy.json
func ExampleBlobContainersClient_LockImmutabilityPolicy() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armstorage.NewBlobContainersClient("<subscription-id>", cred, nil)
	res, err := client.LockImmutabilityPolicy(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<container-name>",
		"<if-match>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.BlobContainersClientLockImmutabilityPolicyResult)
}

// x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-08-01/examples/BlobContainersExtendImmutabilityPolicy.json
func ExampleBlobContainersClient_ExtendImmutabilityPolicy() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armstorage.NewBlobContainersClient("<subscription-id>", cred, nil)
	res, err := client.ExtendImmutabilityPolicy(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<container-name>",
		"<if-match>",
		&armstorage.BlobContainersClientExtendImmutabilityPolicyOptions{Parameters: &armstorage.ImmutabilityPolicy{
			Properties: &armstorage.ImmutabilityPolicyProperty{
				ImmutabilityPeriodSinceCreationInDays: to.Int32Ptr(100),
			},
		},
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.BlobContainersClientExtendImmutabilityPolicyResult)
}

// x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-08-01/examples/BlobContainersLease_Acquire.json
func ExampleBlobContainersClient_Lease() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armstorage.NewBlobContainersClient("<subscription-id>", cred, nil)
	res, err := client.Lease(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<container-name>",
		&armstorage.BlobContainersClientLeaseOptions{Parameters: &armstorage.LeaseContainerRequest{
			Action:        armstorage.LeaseContainerRequestAction("Acquire").ToPtr(),
			LeaseDuration: to.Int32Ptr(-1),
		},
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.BlobContainersClientLeaseResult)
}

// x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-08-01/examples/ObjectLevelWormContainerMigration.json
func ExampleBlobContainersClient_BeginObjectLevelWorm() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armstorage.NewBlobContainersClient("<subscription-id>", cred, nil)
	poller, err := client.BeginObjectLevelWorm(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<container-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}
