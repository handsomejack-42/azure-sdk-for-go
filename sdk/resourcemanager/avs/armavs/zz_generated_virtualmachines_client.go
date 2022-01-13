//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armavs

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// VirtualMachinesClient contains the methods for the VirtualMachines group.
// Don't use this type directly, use NewVirtualMachinesClient() instead.
type VirtualMachinesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewVirtualMachinesClient creates a new instance of VirtualMachinesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewVirtualMachinesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *VirtualMachinesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &VirtualMachinesClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// Get - Get a virtual machine by id in a private cloud cluster
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// privateCloudName - Name of the private cloud
// clusterName - Name of the cluster in the private cloud
// virtualMachineID - Virtual Machine identifier
// options - VirtualMachinesClientGetOptions contains the optional parameters for the VirtualMachinesClient.Get method.
func (client *VirtualMachinesClient) Get(ctx context.Context, resourceGroupName string, privateCloudName string, clusterName string, virtualMachineID string, options *VirtualMachinesClientGetOptions) (VirtualMachinesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, privateCloudName, clusterName, virtualMachineID, options)
	if err != nil {
		return VirtualMachinesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VirtualMachinesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualMachinesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VirtualMachinesClient) getCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, clusterName string, virtualMachineID string, options *VirtualMachinesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/clusters/{clusterName}/virtualMachines/{virtualMachineId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if virtualMachineID == "" {
		return nil, errors.New("parameter virtualMachineID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualMachineId}", url.PathEscape(virtualMachineID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VirtualMachinesClient) getHandleResponse(resp *http.Response) (VirtualMachinesClientGetResponse, error) {
	result := VirtualMachinesClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachine); err != nil {
		return VirtualMachinesClientGetResponse{}, err
	}
	return result, nil
}

// List - List of virtual machines in a private cloud cluster
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// privateCloudName - Name of the private cloud
// clusterName - Name of the cluster in the private cloud
// options - VirtualMachinesClientListOptions contains the optional parameters for the VirtualMachinesClient.List method.
func (client *VirtualMachinesClient) List(resourceGroupName string, privateCloudName string, clusterName string, options *VirtualMachinesClientListOptions) *VirtualMachinesClientListPager {
	return &VirtualMachinesClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, privateCloudName, clusterName, options)
		},
		advancer: func(ctx context.Context, resp VirtualMachinesClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.VirtualMachinesList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *VirtualMachinesClient) listCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, clusterName string, options *VirtualMachinesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/clusters/{clusterName}/virtualMachines"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *VirtualMachinesClient) listHandleResponse(resp *http.Response) (VirtualMachinesClientListResponse, error) {
	result := VirtualMachinesClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachinesList); err != nil {
		return VirtualMachinesClientListResponse{}, err
	}
	return result, nil
}

// BeginRestrictMovement - Enable or disable DRS-driven VM movement restriction
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// privateCloudName - Name of the private cloud
// clusterName - Name of the cluster in the private cloud
// virtualMachineID - Virtual Machine identifier
// restrictMovement - Whether VM DRS-driven movement is restricted (Enabled) or not (Disabled)
// options - VirtualMachinesClientBeginRestrictMovementOptions contains the optional parameters for the VirtualMachinesClient.BeginRestrictMovement
// method.
func (client *VirtualMachinesClient) BeginRestrictMovement(ctx context.Context, resourceGroupName string, privateCloudName string, clusterName string, virtualMachineID string, restrictMovement VirtualMachineRestrictMovement, options *VirtualMachinesClientBeginRestrictMovementOptions) (VirtualMachinesClientRestrictMovementPollerResponse, error) {
	resp, err := client.restrictMovement(ctx, resourceGroupName, privateCloudName, clusterName, virtualMachineID, restrictMovement, options)
	if err != nil {
		return VirtualMachinesClientRestrictMovementPollerResponse{}, err
	}
	result := VirtualMachinesClientRestrictMovementPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("VirtualMachinesClient.RestrictMovement", "", resp, client.pl)
	if err != nil {
		return VirtualMachinesClientRestrictMovementPollerResponse{}, err
	}
	result.Poller = &VirtualMachinesClientRestrictMovementPoller{
		pt: pt,
	}
	return result, nil
}

// RestrictMovement - Enable or disable DRS-driven VM movement restriction
// If the operation fails it returns an *azcore.ResponseError type.
func (client *VirtualMachinesClient) restrictMovement(ctx context.Context, resourceGroupName string, privateCloudName string, clusterName string, virtualMachineID string, restrictMovement VirtualMachineRestrictMovement, options *VirtualMachinesClientBeginRestrictMovementOptions) (*http.Response, error) {
	req, err := client.restrictMovementCreateRequest(ctx, resourceGroupName, privateCloudName, clusterName, virtualMachineID, restrictMovement, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// restrictMovementCreateRequest creates the RestrictMovement request.
func (client *VirtualMachinesClient) restrictMovementCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, clusterName string, virtualMachineID string, restrictMovement VirtualMachineRestrictMovement, options *VirtualMachinesClientBeginRestrictMovementOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/clusters/{clusterName}/virtualMachines/{virtualMachineId}/restrictMovement"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if virtualMachineID == "" {
		return nil, errors.New("parameter virtualMachineID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualMachineId}", url.PathEscape(virtualMachineID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, restrictMovement)
}
