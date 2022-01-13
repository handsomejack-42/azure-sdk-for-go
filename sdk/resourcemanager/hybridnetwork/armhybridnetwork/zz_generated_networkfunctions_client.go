//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhybridnetwork

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

// NetworkFunctionsClient contains the methods for the NetworkFunctions group.
// Don't use this type directly, use NewNetworkFunctionsClient() instead.
type NetworkFunctionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewNetworkFunctionsClient creates a new instance of NetworkFunctionsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewNetworkFunctionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *NetworkFunctionsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &NetworkFunctionsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// BeginCreateOrUpdate - Creates or updates a network function resource. This operation can take up to 6 hours to complete.
// This is expected service behavior.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// networkFunctionName - Resource name for the network function resource.
// parameters - Parameters supplied in the body to the create or update network function operation.
// options - NetworkFunctionsClientBeginCreateOrUpdateOptions contains the optional parameters for the NetworkFunctionsClient.BeginCreateOrUpdate
// method.
func (client *NetworkFunctionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, networkFunctionName string, parameters NetworkFunction, options *NetworkFunctionsClientBeginCreateOrUpdateOptions) (NetworkFunctionsClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, networkFunctionName, parameters, options)
	if err != nil {
		return NetworkFunctionsClientCreateOrUpdatePollerResponse{}, err
	}
	result := NetworkFunctionsClientCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("NetworkFunctionsClient.CreateOrUpdate", "azure-async-operation", resp, client.pl)
	if err != nil {
		return NetworkFunctionsClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &NetworkFunctionsClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a network function resource. This operation can take up to 6 hours to complete. This
// is expected service behavior.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *NetworkFunctionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, networkFunctionName string, parameters NetworkFunction, options *NetworkFunctionsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, networkFunctionName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *NetworkFunctionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, networkFunctionName string, parameters NetworkFunction, options *NetworkFunctionsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridNetwork/networkFunctions/{networkFunctionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkFunctionName == "" {
		return nil, errors.New("parameter networkFunctionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkFunctionName}", url.PathEscape(networkFunctionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes the specified network function resource. This operation can take up to 1 hour to complete. This is
// expected service behavior.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// networkFunctionName - The name of the network function.
// options - NetworkFunctionsClientBeginDeleteOptions contains the optional parameters for the NetworkFunctionsClient.BeginDelete
// method.
func (client *NetworkFunctionsClient) BeginDelete(ctx context.Context, resourceGroupName string, networkFunctionName string, options *NetworkFunctionsClientBeginDeleteOptions) (NetworkFunctionsClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, networkFunctionName, options)
	if err != nil {
		return NetworkFunctionsClientDeletePollerResponse{}, err
	}
	result := NetworkFunctionsClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("NetworkFunctionsClient.Delete", "location", resp, client.pl)
	if err != nil {
		return NetworkFunctionsClientDeletePollerResponse{}, err
	}
	result.Poller = &NetworkFunctionsClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the specified network function resource. This operation can take up to 1 hour to complete. This is expected
// service behavior.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *NetworkFunctionsClient) deleteOperation(ctx context.Context, resourceGroupName string, networkFunctionName string, options *NetworkFunctionsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, networkFunctionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *NetworkFunctionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, networkFunctionName string, options *NetworkFunctionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridNetwork/networkFunctions/{networkFunctionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkFunctionName == "" {
		return nil, errors.New("parameter networkFunctionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkFunctionName}", url.PathEscape(networkFunctionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets information about the specified network function resource.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// networkFunctionName - The name of the network function resource.
// options - NetworkFunctionsClientGetOptions contains the optional parameters for the NetworkFunctionsClient.Get method.
func (client *NetworkFunctionsClient) Get(ctx context.Context, resourceGroupName string, networkFunctionName string, options *NetworkFunctionsClientGetOptions) (NetworkFunctionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, networkFunctionName, options)
	if err != nil {
		return NetworkFunctionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return NetworkFunctionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return NetworkFunctionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *NetworkFunctionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, networkFunctionName string, options *NetworkFunctionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridNetwork/networkFunctions/{networkFunctionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkFunctionName == "" {
		return nil, errors.New("parameter networkFunctionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkFunctionName}", url.PathEscape(networkFunctionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *NetworkFunctionsClient) getHandleResponse(resp *http.Response) (NetworkFunctionsClientGetResponse, error) {
	result := NetworkFunctionsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkFunction); err != nil {
		return NetworkFunctionsClientGetResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - Lists all the network function resources in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - NetworkFunctionsClientListByResourceGroupOptions contains the optional parameters for the NetworkFunctionsClient.ListByResourceGroup
// method.
func (client *NetworkFunctionsClient) ListByResourceGroup(resourceGroupName string, options *NetworkFunctionsClientListByResourceGroupOptions) *NetworkFunctionsClientListByResourceGroupPager {
	return &NetworkFunctionsClientListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp NetworkFunctionsClientListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.NetworkFunctionListResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *NetworkFunctionsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *NetworkFunctionsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridNetwork/networkFunctions"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *NetworkFunctionsClient) listByResourceGroupHandleResponse(resp *http.Response) (NetworkFunctionsClientListByResourceGroupResponse, error) {
	result := NetworkFunctionsClientListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkFunctionListResult); err != nil {
		return NetworkFunctionsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// ListBySubscription - Lists all the network functions in a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - NetworkFunctionsClientListBySubscriptionOptions contains the optional parameters for the NetworkFunctionsClient.ListBySubscription
// method.
func (client *NetworkFunctionsClient) ListBySubscription(options *NetworkFunctionsClientListBySubscriptionOptions) *NetworkFunctionsClientListBySubscriptionPager {
	return &NetworkFunctionsClientListBySubscriptionPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listBySubscriptionCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp NetworkFunctionsClientListBySubscriptionResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.NetworkFunctionListResult.NextLink)
		},
	}
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *NetworkFunctionsClient) listBySubscriptionCreateRequest(ctx context.Context, options *NetworkFunctionsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.HybridNetwork/networkFunctions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *NetworkFunctionsClient) listBySubscriptionHandleResponse(resp *http.Response) (NetworkFunctionsClientListBySubscriptionResponse, error) {
	result := NetworkFunctionsClientListBySubscriptionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkFunctionListResult); err != nil {
		return NetworkFunctionsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// UpdateTags - Updates the tags for the network function resource.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// networkFunctionName - Resource name for the network function resource.
// parameters - Parameters supplied to the update network function tags operation.
// options - NetworkFunctionsClientUpdateTagsOptions contains the optional parameters for the NetworkFunctionsClient.UpdateTags
// method.
func (client *NetworkFunctionsClient) UpdateTags(ctx context.Context, resourceGroupName string, networkFunctionName string, parameters TagsObject, options *NetworkFunctionsClientUpdateTagsOptions) (NetworkFunctionsClientUpdateTagsResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, networkFunctionName, parameters, options)
	if err != nil {
		return NetworkFunctionsClientUpdateTagsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return NetworkFunctionsClientUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return NetworkFunctionsClientUpdateTagsResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *NetworkFunctionsClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, networkFunctionName string, parameters TagsObject, options *NetworkFunctionsClientUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridNetwork/networkFunctions/{networkFunctionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkFunctionName == "" {
		return nil, errors.New("parameter networkFunctionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkFunctionName}", url.PathEscape(networkFunctionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *NetworkFunctionsClient) updateTagsHandleResponse(resp *http.Response) (NetworkFunctionsClientUpdateTagsResponse, error) {
	result := NetworkFunctionsClientUpdateTagsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkFunction); err != nil {
		return NetworkFunctionsClientUpdateTagsResponse{}, err
	}
	return result, nil
}
