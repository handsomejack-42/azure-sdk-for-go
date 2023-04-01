//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armpeering

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ConnectionMonitorTestsClient contains the methods for the ConnectionMonitorTests group.
// Don't use this type directly, use NewConnectionMonitorTestsClient() instead.
type ConnectionMonitorTestsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewConnectionMonitorTestsClient creates a new instance of ConnectionMonitorTestsClient with the specified values.
//   - subscriptionID - The Azure subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewConnectionMonitorTestsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ConnectionMonitorTestsClient, error) {
	cl, err := arm.NewClient(moduleName+".ConnectionMonitorTestsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ConnectionMonitorTestsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates a connection monitor test with the specified name under the given subscription, resource
// group and peering service.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-01-01
//   - resourceGroupName - The name of the resource group.
//   - peeringServiceName - The name of the peering service.
//   - connectionMonitorTestName - The name of the connection monitor test
//   - connectionMonitorTest - The properties needed to create a connection monitor test
//   - options - ConnectionMonitorTestsClientCreateOrUpdateOptions contains the optional parameters for the ConnectionMonitorTestsClient.CreateOrUpdate
//     method.
func (client *ConnectionMonitorTestsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, peeringServiceName string, connectionMonitorTestName string, connectionMonitorTest ConnectionMonitorTest, options *ConnectionMonitorTestsClientCreateOrUpdateOptions) (ConnectionMonitorTestsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, peeringServiceName, connectionMonitorTestName, connectionMonitorTest, options)
	if err != nil {
		return ConnectionMonitorTestsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConnectionMonitorTestsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return ConnectionMonitorTestsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ConnectionMonitorTestsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, peeringServiceName string, connectionMonitorTestName string, connectionMonitorTest ConnectionMonitorTest, options *ConnectionMonitorTestsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peeringServices/{peeringServiceName}/connectionMonitorTests/{connectionMonitorTestName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if peeringServiceName == "" {
		return nil, errors.New("parameter peeringServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringServiceName}", url.PathEscape(peeringServiceName))
	if connectionMonitorTestName == "" {
		return nil, errors.New("parameter connectionMonitorTestName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionMonitorTestName}", url.PathEscape(connectionMonitorTestName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, connectionMonitorTest)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ConnectionMonitorTestsClient) createOrUpdateHandleResponse(resp *http.Response) (ConnectionMonitorTestsClientCreateOrUpdateResponse, error) {
	result := ConnectionMonitorTestsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConnectionMonitorTest); err != nil {
		return ConnectionMonitorTestsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes an existing connection monitor test with the specified name under the given subscription, resource group
// and peering service.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-01-01
//   - resourceGroupName - The name of the resource group.
//   - peeringServiceName - The name of the peering service.
//   - connectionMonitorTestName - The name of the connection monitor test
//   - options - ConnectionMonitorTestsClientDeleteOptions contains the optional parameters for the ConnectionMonitorTestsClient.Delete
//     method.
func (client *ConnectionMonitorTestsClient) Delete(ctx context.Context, resourceGroupName string, peeringServiceName string, connectionMonitorTestName string, options *ConnectionMonitorTestsClientDeleteOptions) (ConnectionMonitorTestsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, peeringServiceName, connectionMonitorTestName, options)
	if err != nil {
		return ConnectionMonitorTestsClientDeleteResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConnectionMonitorTestsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ConnectionMonitorTestsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return ConnectionMonitorTestsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ConnectionMonitorTestsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, peeringServiceName string, connectionMonitorTestName string, options *ConnectionMonitorTestsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peeringServices/{peeringServiceName}/connectionMonitorTests/{connectionMonitorTestName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if peeringServiceName == "" {
		return nil, errors.New("parameter peeringServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringServiceName}", url.PathEscape(peeringServiceName))
	if connectionMonitorTestName == "" {
		return nil, errors.New("parameter connectionMonitorTestName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionMonitorTestName}", url.PathEscape(connectionMonitorTestName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets an existing connection monitor test with the specified name under the given subscription, resource group and
// peering service.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-01-01
//   - resourceGroupName - The name of the resource group.
//   - peeringServiceName - The name of the peering service.
//   - connectionMonitorTestName - The name of the connection monitor test
//   - options - ConnectionMonitorTestsClientGetOptions contains the optional parameters for the ConnectionMonitorTestsClient.Get
//     method.
func (client *ConnectionMonitorTestsClient) Get(ctx context.Context, resourceGroupName string, peeringServiceName string, connectionMonitorTestName string, options *ConnectionMonitorTestsClientGetOptions) (ConnectionMonitorTestsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, peeringServiceName, connectionMonitorTestName, options)
	if err != nil {
		return ConnectionMonitorTestsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConnectionMonitorTestsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ConnectionMonitorTestsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ConnectionMonitorTestsClient) getCreateRequest(ctx context.Context, resourceGroupName string, peeringServiceName string, connectionMonitorTestName string, options *ConnectionMonitorTestsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peeringServices/{peeringServiceName}/connectionMonitorTests/{connectionMonitorTestName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if peeringServiceName == "" {
		return nil, errors.New("parameter peeringServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringServiceName}", url.PathEscape(peeringServiceName))
	if connectionMonitorTestName == "" {
		return nil, errors.New("parameter connectionMonitorTestName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionMonitorTestName}", url.PathEscape(connectionMonitorTestName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ConnectionMonitorTestsClient) getHandleResponse(resp *http.Response) (ConnectionMonitorTestsClientGetResponse, error) {
	result := ConnectionMonitorTestsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConnectionMonitorTest); err != nil {
		return ConnectionMonitorTestsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByPeeringServicePager - Lists all connection monitor tests under the given subscription, resource group and peering
// service.
//
// Generated from API version 2022-01-01
//   - resourceGroupName - The name of the resource group.
//   - peeringServiceName - The name of the peering service.
//   - options - ConnectionMonitorTestsClientListByPeeringServiceOptions contains the optional parameters for the ConnectionMonitorTestsClient.NewListByPeeringServicePager
//     method.
func (client *ConnectionMonitorTestsClient) NewListByPeeringServicePager(resourceGroupName string, peeringServiceName string, options *ConnectionMonitorTestsClientListByPeeringServiceOptions) *runtime.Pager[ConnectionMonitorTestsClientListByPeeringServiceResponse] {
	return runtime.NewPager(runtime.PagingHandler[ConnectionMonitorTestsClientListByPeeringServiceResponse]{
		More: func(page ConnectionMonitorTestsClientListByPeeringServiceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ConnectionMonitorTestsClientListByPeeringServiceResponse) (ConnectionMonitorTestsClientListByPeeringServiceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByPeeringServiceCreateRequest(ctx, resourceGroupName, peeringServiceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ConnectionMonitorTestsClientListByPeeringServiceResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ConnectionMonitorTestsClientListByPeeringServiceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ConnectionMonitorTestsClientListByPeeringServiceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByPeeringServiceHandleResponse(resp)
		},
	})
}

// listByPeeringServiceCreateRequest creates the ListByPeeringService request.
func (client *ConnectionMonitorTestsClient) listByPeeringServiceCreateRequest(ctx context.Context, resourceGroupName string, peeringServiceName string, options *ConnectionMonitorTestsClientListByPeeringServiceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peeringServices/{peeringServiceName}/connectionMonitorTests"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if peeringServiceName == "" {
		return nil, errors.New("parameter peeringServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringServiceName}", url.PathEscape(peeringServiceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByPeeringServiceHandleResponse handles the ListByPeeringService response.
func (client *ConnectionMonitorTestsClient) listByPeeringServiceHandleResponse(resp *http.Response) (ConnectionMonitorTestsClientListByPeeringServiceResponse, error) {
	result := ConnectionMonitorTestsClientListByPeeringServiceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConnectionMonitorTestListResult); err != nil {
		return ConnectionMonitorTestsClientListByPeeringServiceResponse{}, err
	}
	return result, nil
}