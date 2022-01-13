//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapimanagement

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
	"strconv"
	"strings"
)

// NotificationClient contains the methods for the Notification group.
// Don't use this type directly, use NewNotificationClient() instead.
type NotificationClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewNotificationClient creates a new instance of NotificationClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewNotificationClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *NotificationClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &NotificationClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// CreateOrUpdate - Create or Update API Management publisher notification.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// notificationName - Notification Name Identifier.
// options - NotificationClientCreateOrUpdateOptions contains the optional parameters for the NotificationClient.CreateOrUpdate
// method.
func (client *NotificationClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, notificationName NotificationName, options *NotificationClientCreateOrUpdateOptions) (NotificationClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, notificationName, options)
	if err != nil {
		return NotificationClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return NotificationClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return NotificationClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *NotificationClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, notificationName NotificationName, options *NotificationClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/notifications/{notificationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if notificationName == "" {
		return nil, errors.New("parameter notificationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{notificationName}", url.PathEscape(string(notificationName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header.Set("If-Match", *options.IfMatch)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *NotificationClient) createOrUpdateHandleResponse(resp *http.Response) (NotificationClientCreateOrUpdateResponse, error) {
	result := NotificationClientCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.NotificationContract); err != nil {
		return NotificationClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Get - Gets the details of the Notification specified by its identifier.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// notificationName - Notification Name Identifier.
// options - NotificationClientGetOptions contains the optional parameters for the NotificationClient.Get method.
func (client *NotificationClient) Get(ctx context.Context, resourceGroupName string, serviceName string, notificationName NotificationName, options *NotificationClientGetOptions) (NotificationClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, notificationName, options)
	if err != nil {
		return NotificationClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return NotificationClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return NotificationClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *NotificationClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, notificationName NotificationName, options *NotificationClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/notifications/{notificationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if notificationName == "" {
		return nil, errors.New("parameter notificationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{notificationName}", url.PathEscape(string(notificationName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *NotificationClient) getHandleResponse(resp *http.Response) (NotificationClientGetResponse, error) {
	result := NotificationClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.NotificationContract); err != nil {
		return NotificationClientGetResponse{}, err
	}
	return result, nil
}

// ListByService - Lists a collection of properties defined within a service instance.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// options - NotificationClientListByServiceOptions contains the optional parameters for the NotificationClient.ListByService
// method.
func (client *NotificationClient) ListByService(resourceGroupName string, serviceName string, options *NotificationClientListByServiceOptions) *NotificationClientListByServicePager {
	return &NotificationClientListByServicePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByServiceCreateRequest(ctx, resourceGroupName, serviceName, options)
		},
		advancer: func(ctx context.Context, resp NotificationClientListByServiceResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.NotificationCollection.NextLink)
		},
	}
}

// listByServiceCreateRequest creates the ListByService request.
func (client *NotificationClient) listByServiceCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *NotificationClientListByServiceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/notifications"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByServiceHandleResponse handles the ListByService response.
func (client *NotificationClient) listByServiceHandleResponse(resp *http.Response) (NotificationClientListByServiceResponse, error) {
	result := NotificationClientListByServiceResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.NotificationCollection); err != nil {
		return NotificationClientListByServiceResponse{}, err
	}
	return result, nil
}
