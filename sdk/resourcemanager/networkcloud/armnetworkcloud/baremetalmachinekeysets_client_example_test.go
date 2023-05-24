//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetworkcloud_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkcloud/armnetworkcloud"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/BareMetalMachineKeySets_ListByResourceGroup.json
func ExampleBareMetalMachineKeySetsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworkcloud.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBareMetalMachineKeySetsClient().NewListByResourceGroupPager("resourceGroupName", "clusterName", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.BareMetalMachineKeySetList = armnetworkcloud.BareMetalMachineKeySetList{
		// 	Value: []*armnetworkcloud.BareMetalMachineKeySet{
		// 		{
		// 			Name: to.Ptr("bareMetalMachineKeySetName"),
		// 			Type: to.Ptr("Microsoft.NetworkCloud/clusters/bareMetalMachineKeySets"),
		// 			ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/clusters/clusterName/bareMetalMachineKeySets/bareMetalMachineKeySetName"),
		// 			SystemData: &armnetworkcloud.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:27:03.008Z"); return t}()),
		// 				CreatedBy: to.Ptr("identityA"),
		// 				CreatedByType: to.Ptr(armnetworkcloud.CreatedByTypeApplication),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:29:03.001Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("identityB"),
		// 				LastModifiedByType: to.Ptr(armnetworkcloud.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("location"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("myvalue1"),
		// 				"key2": to.Ptr("myvalue2"),
		// 			},
		// 			ExtendedLocation: &armnetworkcloud.ExtendedLocation{
		// 				Name: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterExtendedLocationName"),
		// 				Type: to.Ptr("CustomLocation"),
		// 			},
		// 			Properties: &armnetworkcloud.BareMetalMachineKeySetProperties{
		// 				AzureGroupID: to.Ptr("f110271b-XXXX-4163-9b99-214d91660f0e"),
		// 				DetailedStatus: to.Ptr(armnetworkcloud.BareMetalMachineKeySetDetailedStatusSomeInvalid),
		// 				DetailedStatusMessage: to.Ptr("Inalid Azure user(s) were provided: userXYZ"),
		// 				Expiration: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-12-31T23:59:59.008Z"); return t}()),
		// 				JumpHostsAllowed: []*string{
		// 					to.Ptr("192.0.2.1"),
		// 					to.Ptr("192.0.2.5")},
		// 					LastValidation: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-12-12T12:00:00.008Z"); return t}()),
		// 					OSGroupName: to.Ptr("standardAccessGroup"),
		// 					PrivilegeLevel: to.Ptr(armnetworkcloud.BareMetalMachineKeySetPrivilegeLevelStandard),
		// 					ProvisioningState: to.Ptr(armnetworkcloud.BareMetalMachineKeySetProvisioningStateSucceeded),
		// 					UserList: []*armnetworkcloud.KeySetUser{
		// 						{
		// 							Description: to.Ptr("Needs access for troubleshooting as a part of the support team"),
		// 							AzureUserName: to.Ptr("userABC"),
		// 							SSHPublicKey: &armnetworkcloud.SSHPublicKey{
		// 								KeyData: to.Ptr("ssh-rsa AAtsE3njSONzDYRIZv/WLjVuMfrUSByHp+jfaaOLHTIIB4fJvo6dQUZxE20w2iDHV3tEkmnTo84eba97VMueQD6OzJPEyWZMRpz8UYWOd0IXeRqiFu1lawNblZhwNT/ojNZfpB3af/YDzwQCZgTcTRyNNhL4o/blKUmug0daSsSXISTRnIDpcf5qytjs1Xo+yYyJMvzLL59mhAyb3p/cD+Y3/s3WhAx+l0XOKpzXnblrv9d3q4c2tWmm/SyFqthaqd0= admin@vm"),
		// 							},
		// 						},
		// 						{
		// 							Description: to.Ptr("Needs access for troubleshooting as a part of the support team"),
		// 							AzureUserName: to.Ptr("userXYZ"),
		// 							SSHPublicKey: &armnetworkcloud.SSHPublicKey{
		// 								KeyData: to.Ptr("ssh-rsa AAtsE3njSONzDYRIZv/WLjVuMfrUSByHp+jfaaOLHTIIB4fJvo6dQUZxE20w2iDHV3tEkmnTo84eba97VMueQD6OzJPEyWZMRpz8UYWOd0IXeRqiFu1lawNblZhwNT/ojNZfpB3af/YDzwQCZgTcTRyNNhL4o/blKUmug0daSsSXISTRnIDpcf5qytjs1Xo+yYyJMvzLL59mhAyb3p/cD+Y3/s3WhAx+l0XOKpzXnblrv9d3q4c2tWmm/SyFqthaqd0= admin@vm"),
		// 							},
		// 					}},
		// 					UserListStatus: []*armnetworkcloud.KeySetUserStatus{
		// 						{
		// 							AzureUserName: to.Ptr("userABC"),
		// 							Status: to.Ptr(armnetworkcloud.BareMetalMachineKeySetUserSetupStatusActive),
		// 							StatusMessage: to.Ptr("User has been provisioned"),
		// 						},
		// 						{
		// 							AzureUserName: to.Ptr("userXYZ"),
		// 							Status: to.Ptr(armnetworkcloud.BareMetalMachineKeySetUserSetupStatusInvalid),
		// 							StatusMessage: to.Ptr("User is not a valid Azure user"),
		// 					}},
		// 				},
		// 		}},
		// 	}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/BareMetalMachineKeySets_Get.json
func ExampleBareMetalMachineKeySetsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworkcloud.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBareMetalMachineKeySetsClient().Get(ctx, "resourceGroupName", "clusterName", "bareMetalMachineKeySetName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.BareMetalMachineKeySet = armnetworkcloud.BareMetalMachineKeySet{
	// 	Name: to.Ptr("bareMetalMachineKeySetName"),
	// 	Type: to.Ptr("Microsoft.NetworkCloud/clusters/bareMetalMachineKeySets"),
	// 	ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/clusters/clusterName/bareMetalMachineKeySets/bareMetalMachineKeySetName"),
	// 	SystemData: &armnetworkcloud.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:27:03.008Z"); return t}()),
	// 		CreatedBy: to.Ptr("identityA"),
	// 		CreatedByType: to.Ptr(armnetworkcloud.CreatedByTypeApplication),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:29:03.001Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("identityB"),
	// 		LastModifiedByType: to.Ptr(armnetworkcloud.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("location"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("myvalue1"),
	// 		"key2": to.Ptr("myvalue2"),
	// 	},
	// 	ExtendedLocation: &armnetworkcloud.ExtendedLocation{
	// 		Name: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterExtendedLocationName"),
	// 		Type: to.Ptr("CustomLocation"),
	// 	},
	// 	Properties: &armnetworkcloud.BareMetalMachineKeySetProperties{
	// 		AzureGroupID: to.Ptr("f110271b-XXXX-4163-9b99-214d91660f0e"),
	// 		DetailedStatus: to.Ptr(armnetworkcloud.BareMetalMachineKeySetDetailedStatusSomeInvalid),
	// 		DetailedStatusMessage: to.Ptr("Inalid Azure user(s) were provided: userXYZ"),
	// 		Expiration: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-12-31T23:59:59.008Z"); return t}()),
	// 		JumpHostsAllowed: []*string{
	// 			to.Ptr("192.0.2.1"),
	// 			to.Ptr("192.0.2.5")},
	// 			LastValidation: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-12-12T12:00:00.008Z"); return t}()),
	// 			OSGroupName: to.Ptr("standardAccessGroup"),
	// 			PrivilegeLevel: to.Ptr(armnetworkcloud.BareMetalMachineKeySetPrivilegeLevelStandard),
	// 			ProvisioningState: to.Ptr(armnetworkcloud.BareMetalMachineKeySetProvisioningStateSucceeded),
	// 			UserList: []*armnetworkcloud.KeySetUser{
	// 				{
	// 					Description: to.Ptr("Needs access for troubleshooting as a part of the support team"),
	// 					AzureUserName: to.Ptr("userABC"),
	// 					SSHPublicKey: &armnetworkcloud.SSHPublicKey{
	// 						KeyData: to.Ptr("ssh-rsa AAtsE3njSONzDYRIZv/WLjVuMfrUSByHp+jfaaOLHTIIB4fJvo6dQUZxE20w2iDHV3tEkmnTo84eba97VMueQD6OzJPEyWZMRpz8UYWOd0IXeRqiFu1lawNblZhwNT/ojNZfpB3af/YDzwQCZgTcTRyNNhL4o/blKUmug0daSsSXISTRnIDpcf5qytjs1Xo+yYyJMvzLL59mhAyb3p/cD+Y3/s3WhAx+l0XOKpzXnblrv9d3q4c2tWmm/SyFqthaqd0= admin@vm"),
	// 					},
	// 				},
	// 				{
	// 					Description: to.Ptr("Needs access for troubleshooting as a part of the support team"),
	// 					AzureUserName: to.Ptr("userXYZ"),
	// 					SSHPublicKey: &armnetworkcloud.SSHPublicKey{
	// 						KeyData: to.Ptr("ssh-rsa AAtsE3njSONzDYRIZv/WLjVuMfrUSByHp+jfaaOLHTIIB4fJvo6dQUZxE20w2iDHV3tEkmnTo84eba97VMueQD6OzJPEyWZMRpz8UYWOd0IXeRqiFu1lawNblZhwNT/ojNZfpB3af/YDzwQCZgTcTRyNNhL4o/blKUmug0daSsSXISTRnIDpcf5qytjs1Xo+yYyJMvzLL59mhAyb3p/cD+Y3/s3WhAx+l0XOKpzXnblrv9d3q4c2tWmm/SyFqthaqd0= admin@vm"),
	// 					},
	// 			}},
	// 			UserListStatus: []*armnetworkcloud.KeySetUserStatus{
	// 				{
	// 					AzureUserName: to.Ptr("userABC"),
	// 					Status: to.Ptr(armnetworkcloud.BareMetalMachineKeySetUserSetupStatusActive),
	// 					StatusMessage: to.Ptr("User has been provisioned"),
	// 				},
	// 				{
	// 					AzureUserName: to.Ptr("userXYZ"),
	// 					Status: to.Ptr(armnetworkcloud.BareMetalMachineKeySetUserSetupStatusInvalid),
	// 					StatusMessage: to.Ptr("User is not a valid Azure user"),
	// 			}},
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/BareMetalMachineKeySets_Create.json
func ExampleBareMetalMachineKeySetsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworkcloud.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBareMetalMachineKeySetsClient().BeginCreateOrUpdate(ctx, "resourceGroupName", "clusterName", "bareMetalMachineKeySetName", armnetworkcloud.BareMetalMachineKeySet{
		Location: to.Ptr("location"),
		Tags: map[string]*string{
			"key1": to.Ptr("myvalue1"),
			"key2": to.Ptr("myvalue2"),
		},
		ExtendedLocation: &armnetworkcloud.ExtendedLocation{
			Name: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterExtendedLocationName"),
			Type: to.Ptr("CustomLocation"),
		},
		Properties: &armnetworkcloud.BareMetalMachineKeySetProperties{
			AzureGroupID: to.Ptr("f110271b-XXXX-4163-9b99-214d91660f0e"),
			Expiration:   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-12-31T23:59:59.008Z"); return t }()),
			JumpHostsAllowed: []*string{
				to.Ptr("192.0.2.1"),
				to.Ptr("192.0.2.5")},
			OSGroupName:    to.Ptr("standardAccessGroup"),
			PrivilegeLevel: to.Ptr(armnetworkcloud.BareMetalMachineKeySetPrivilegeLevelStandard),
			UserList: []*armnetworkcloud.KeySetUser{
				{
					Description:   to.Ptr("Needs access for troubleshooting as a part of the support team"),
					AzureUserName: to.Ptr("userABC"),
					SSHPublicKey: &armnetworkcloud.SSHPublicKey{
						KeyData: to.Ptr("ssh-rsa AAtsE3njSONzDYRIZv/WLjVuMfrUSByHp+jfaaOLHTIIB4fJvo6dQUZxE20w2iDHV3tEkmnTo84eba97VMueQD6OzJPEyWZMRpz8UYWOd0IXeRqiFu1lawNblZhwNT/ojNZfpB3af/YDzwQCZgTcTRyNNhL4o/blKUmug0daSsSXISTRnIDpcf5qytjs1Xo+yYyJMvzLL59mhAyb3p/cD+Y3/s3WhAx+l0XOKpzXnblrv9d3q4c2tWmm/SyFqthaqd0= admin@vm"),
					},
				},
				{
					Description:   to.Ptr("Needs access for troubleshooting as a part of the support team"),
					AzureUserName: to.Ptr("userXYZ"),
					SSHPublicKey: &armnetworkcloud.SSHPublicKey{
						KeyData: to.Ptr("ssh-rsa AAtsE3njSONzDYRIZv/WLjVuMfrUSByHp+jfaaOLHTIIB4fJvo6dQUZxE20w2iDHV3tEkmnTo84eba97VMueQD6OzJPEyWZMRpz8UYWOd0IXeRqiFu1lawNblZhwNT/ojNZfpB3af/YDzwQCZgTcTRyNNhL4o/blKUmug0daSsSXISTRnIDpcf5qytjs1Xo+yYyJMvzLL59mhAyb3p/cD+Y3/s3WhAx+l0XOKpzXnblrv9d3q4c2tWmm/SyFqthaqd0= admin@vm"),
					},
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.BareMetalMachineKeySet = armnetworkcloud.BareMetalMachineKeySet{
	// 	Name: to.Ptr("bareMetalMachineKeySetName"),
	// 	Type: to.Ptr("Microsoft.NetworkCloud/clusters/bareMetalMachineKeySets"),
	// 	ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/clusters/clusterName/bareMetalMachineKeySets/bareMetalMachineKeySetName"),
	// 	SystemData: &armnetworkcloud.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:27:03.008Z"); return t}()),
	// 		CreatedBy: to.Ptr("identityA"),
	// 		CreatedByType: to.Ptr(armnetworkcloud.CreatedByTypeApplication),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:29:03.001Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("identityB"),
	// 		LastModifiedByType: to.Ptr(armnetworkcloud.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("location"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("myvalue1"),
	// 		"key2": to.Ptr("myvalue2"),
	// 	},
	// 	ExtendedLocation: &armnetworkcloud.ExtendedLocation{
	// 		Name: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterExtendedLocationName"),
	// 		Type: to.Ptr("CustomLocation"),
	// 	},
	// 	Properties: &armnetworkcloud.BareMetalMachineKeySetProperties{
	// 		AzureGroupID: to.Ptr("f110271b-XXXX-4163-9b99-214d91660f0e"),
	// 		DetailedStatus: to.Ptr(armnetworkcloud.BareMetalMachineKeySetDetailedStatusSomeInvalid),
	// 		DetailedStatusMessage: to.Ptr("Inalid Azure user(s) were provided: userXYZ"),
	// 		Expiration: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-12-31T23:59:59.008Z"); return t}()),
	// 		JumpHostsAllowed: []*string{
	// 			to.Ptr("192.0.2.1"),
	// 			to.Ptr("192.0.2.5")},
	// 			LastValidation: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-12-12T12:00:00.008Z"); return t}()),
	// 			OSGroupName: to.Ptr("standardAccessGroup"),
	// 			PrivilegeLevel: to.Ptr(armnetworkcloud.BareMetalMachineKeySetPrivilegeLevelStandard),
	// 			ProvisioningState: to.Ptr(armnetworkcloud.BareMetalMachineKeySetProvisioningStateSucceeded),
	// 			UserList: []*armnetworkcloud.KeySetUser{
	// 				{
	// 					Description: to.Ptr("Needs access for troubleshooting as a part of the support team"),
	// 					AzureUserName: to.Ptr("userABC"),
	// 					SSHPublicKey: &armnetworkcloud.SSHPublicKey{
	// 						KeyData: to.Ptr("ssh-rsa AAtsE3njSONzDYRIZv/WLjVuMfrUSByHp+jfaaOLHTIIB4fJvo6dQUZxE20w2iDHV3tEkmnTo84eba97VMueQD6OzJPEyWZMRpz8UYWOd0IXeRqiFu1lawNblZhwNT/ojNZfpB3af/YDzwQCZgTcTRyNNhL4o/blKUmug0daSsSXISTRnIDpcf5qytjs1Xo+yYyJMvzLL59mhAyb3p/cD+Y3/s3WhAx+l0XOKpzXnblrv9d3q4c2tWmm/SyFqthaqd0= admin@vm"),
	// 					},
	// 				},
	// 				{
	// 					Description: to.Ptr("Needs access for troubleshooting as a part of the support team"),
	// 					AzureUserName: to.Ptr("userXYZ"),
	// 					SSHPublicKey: &armnetworkcloud.SSHPublicKey{
	// 						KeyData: to.Ptr("ssh-rsa AAtsE3njSONzDYRIZv/WLjVuMfrUSByHp+jfaaOLHTIIB4fJvo6dQUZxE20w2iDHV3tEkmnTo84eba97VMueQD6OzJPEyWZMRpz8UYWOd0IXeRqiFu1lawNblZhwNT/ojNZfpB3af/YDzwQCZgTcTRyNNhL4o/blKUmug0daSsSXISTRnIDpcf5qytjs1Xo+yYyJMvzLL59mhAyb3p/cD+Y3/s3WhAx+l0XOKpzXnblrv9d3q4c2tWmm/SyFqthaqd0= admin@vm"),
	// 					},
	// 			}},
	// 			UserListStatus: []*armnetworkcloud.KeySetUserStatus{
	// 				{
	// 					AzureUserName: to.Ptr("userABC"),
	// 					Status: to.Ptr(armnetworkcloud.BareMetalMachineKeySetUserSetupStatusActive),
	// 					StatusMessage: to.Ptr("User has been provisioned"),
	// 				},
	// 				{
	// 					AzureUserName: to.Ptr("userXYZ"),
	// 					Status: to.Ptr(armnetworkcloud.BareMetalMachineKeySetUserSetupStatusInvalid),
	// 					StatusMessage: to.Ptr("User is not a valid Azure user"),
	// 			}},
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/BareMetalMachineKeySets_Delete.json
func ExampleBareMetalMachineKeySetsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworkcloud.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBareMetalMachineKeySetsClient().BeginDelete(ctx, "resourceGroupName", "clusterName", "bareMetalMachineKeySetName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/BareMetalMachineKeySets_Patch.json
func ExampleBareMetalMachineKeySetsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworkcloud.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBareMetalMachineKeySetsClient().BeginUpdate(ctx, "resourceGroupName", "clusterName", "bareMetalMachineKeySetName", armnetworkcloud.BareMetalMachineKeySetPatchParameters{
		Properties: &armnetworkcloud.BareMetalMachineKeySetPatchProperties{
			Expiration: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-12-31T23:59:59.008Z"); return t }()),
			JumpHostsAllowed: []*string{
				to.Ptr("192.0.2.1"),
				to.Ptr("192.0.2.5")},
			UserList: []*armnetworkcloud.KeySetUser{
				{
					Description:   to.Ptr("Needs access for troubleshooting as a part of the support team"),
					AzureUserName: to.Ptr("userABC"),
					SSHPublicKey: &armnetworkcloud.SSHPublicKey{
						KeyData: to.Ptr("ssh-rsa AAtsE3njSONzDYRIZv/WLjVuMfrUSByHp+jfaaOLHTIIB4fJvo6dQUZxE20w2iDHV3tEkmnTo84eba97VMueQD6OzJPEyWZMRpz8UYWOd0IXeRqiFu1lawNblZhwNT/ojNZfpB3af/YDzwQCZgTcTRyNNhL4o/blKUmug0daSsSXISTRnIDpcf5qytjs1Xo+yYyJMvzLL59mhAyb3p/cD+Y3/s3WhAx+l0XOKpzXnblrv9d3q4c2tWmm/SyFqthaqd0= admin@vm"),
					},
				},
				{
					Description:   to.Ptr("Needs access for troubleshooting as a part of the support team"),
					AzureUserName: to.Ptr("userXYZ"),
					SSHPublicKey: &armnetworkcloud.SSHPublicKey{
						KeyData: to.Ptr("ssh-rsa AAtsE3njSONzDYRIZv/WLjVuMfrUSByHp+jfaaOLHTIIB4fJvo6dQUZxE20w2iDHV3tEkmnTo84eba97VMueQD6OzJPEyWZMRpz8UYWOd0IXeRqiFu1lawNblZhwNT/ojNZfpB3af/YDzwQCZgTcTRyNNhL4o/blKUmug0daSsSXISTRnIDpcf5qytjs1Xo+yYyJMvzLL59mhAyb3p/cD+Y3/s3WhAx+l0XOKpzXnblrv9d3q4c2tWmm/SyFqthaqd0= admin@vm"),
					},
				}},
		},
		Tags: map[string]*string{
			"key1": to.Ptr("myvalue1"),
			"key2": to.Ptr("myvalue2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.BareMetalMachineKeySet = armnetworkcloud.BareMetalMachineKeySet{
	// 	Name: to.Ptr("bareMetalMachineKeySetName"),
	// 	Type: to.Ptr("Microsoft.NetworkCloud/clusters/bareMetalMachineKeySets"),
	// 	ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/clusters/clusterName/bareMetalMachineKeySets/bareMetalMachineKeySetName"),
	// 	SystemData: &armnetworkcloud.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:27:03.008Z"); return t}()),
	// 		CreatedBy: to.Ptr("identityA"),
	// 		CreatedByType: to.Ptr(armnetworkcloud.CreatedByTypeApplication),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:29:03.001Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("identityB"),
	// 		LastModifiedByType: to.Ptr(armnetworkcloud.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("location"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("myvalue1"),
	// 		"key2": to.Ptr("myvalue2"),
	// 	},
	// 	ExtendedLocation: &armnetworkcloud.ExtendedLocation{
	// 		Name: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterExtendedLocationName"),
	// 		Type: to.Ptr("CustomLocation"),
	// 	},
	// 	Properties: &armnetworkcloud.BareMetalMachineKeySetProperties{
	// 		AzureGroupID: to.Ptr("f110271b-XXXX-4163-9b99-214d91660f0e"),
	// 		DetailedStatus: to.Ptr(armnetworkcloud.BareMetalMachineKeySetDetailedStatusSomeInvalid),
	// 		DetailedStatusMessage: to.Ptr("Inalid Azure user(s) were provided: userXYZ"),
	// 		Expiration: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-12-31T23:59:59.008Z"); return t}()),
	// 		JumpHostsAllowed: []*string{
	// 			to.Ptr("192.0.2.1"),
	// 			to.Ptr("192.0.2.5")},
	// 			LastValidation: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-12-12T12:00:00.008Z"); return t}()),
	// 			OSGroupName: to.Ptr("standardAccessGroup"),
	// 			PrivilegeLevel: to.Ptr(armnetworkcloud.BareMetalMachineKeySetPrivilegeLevelStandard),
	// 			ProvisioningState: to.Ptr(armnetworkcloud.BareMetalMachineKeySetProvisioningStateSucceeded),
	// 			UserList: []*armnetworkcloud.KeySetUser{
	// 				{
	// 					Description: to.Ptr("Needs access for troubleshooting as a part of the support team"),
	// 					AzureUserName: to.Ptr("userABC"),
	// 					SSHPublicKey: &armnetworkcloud.SSHPublicKey{
	// 						KeyData: to.Ptr("ssh-rsa AAtsE3njSONzDYRIZv/WLjVuMfrUSByHp+jfaaOLHTIIB4fJvo6dQUZxE20w2iDHV3tEkmnTo84eba97VMueQD6OzJPEyWZMRpz8UYWOd0IXeRqiFu1lawNblZhwNT/ojNZfpB3af/YDzwQCZgTcTRyNNhL4o/blKUmug0daSsSXISTRnIDpcf5qytjs1Xo+yYyJMvzLL59mhAyb3p/cD+Y3/s3WhAx+l0XOKpzXnblrv9d3q4c2tWmm/SyFqthaqd0= admin@vm"),
	// 					},
	// 				},
	// 				{
	// 					Description: to.Ptr("Needs access for troubleshooting as a part of the support team"),
	// 					AzureUserName: to.Ptr("userXYZ"),
	// 					SSHPublicKey: &armnetworkcloud.SSHPublicKey{
	// 						KeyData: to.Ptr("ssh-rsa AAtsE3njSONzDYRIZv/WLjVuMfrUSByHp+jfaaOLHTIIB4fJvo6dQUZxE20w2iDHV3tEkmnTo84eba97VMueQD6OzJPEyWZMRpz8UYWOd0IXeRqiFu1lawNblZhwNT/ojNZfpB3af/YDzwQCZgTcTRyNNhL4o/blKUmug0daSsSXISTRnIDpcf5qytjs1Xo+yYyJMvzLL59mhAyb3p/cD+Y3/s3WhAx+l0XOKpzXnblrv9d3q4c2tWmm/SyFqthaqd0= admin@vm"),
	// 					},
	// 			}},
	// 			UserListStatus: []*armnetworkcloud.KeySetUserStatus{
	// 				{
	// 					AzureUserName: to.Ptr("userABC"),
	// 					Status: to.Ptr(armnetworkcloud.BareMetalMachineKeySetUserSetupStatusActive),
	// 					StatusMessage: to.Ptr("User has been provisioned"),
	// 				},
	// 				{
	// 					AzureUserName: to.Ptr("userXYZ"),
	// 					Status: to.Ptr(armnetworkcloud.BareMetalMachineKeySetUserSetupStatusInvalid),
	// 					StatusMessage: to.Ptr("User is not a valid Azure user"),
	// 			}},
	// 		},
	// 	}
}