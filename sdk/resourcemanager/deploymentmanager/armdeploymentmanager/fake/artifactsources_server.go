//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deploymentmanager/armdeploymentmanager"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
)

// ArtifactSourcesServer is a fake server for instances of the armdeploymentmanager.ArtifactSourcesClient type.
type ArtifactSourcesServer struct {
	// CreateOrUpdate is the fake for method ArtifactSourcesClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, artifactSourceName string, options *armdeploymentmanager.ArtifactSourcesClientCreateOrUpdateOptions) (resp azfake.Responder[armdeploymentmanager.ArtifactSourcesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method ArtifactSourcesClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, artifactSourceName string, options *armdeploymentmanager.ArtifactSourcesClientDeleteOptions) (resp azfake.Responder[armdeploymentmanager.ArtifactSourcesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ArtifactSourcesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, artifactSourceName string, options *armdeploymentmanager.ArtifactSourcesClientGetOptions) (resp azfake.Responder[armdeploymentmanager.ArtifactSourcesClientGetResponse], errResp azfake.ErrorResponder)

	// List is the fake for method ArtifactSourcesClient.List
	// HTTP status codes to indicate success: http.StatusOK
	List func(ctx context.Context, resourceGroupName string, options *armdeploymentmanager.ArtifactSourcesClientListOptions) (resp azfake.Responder[armdeploymentmanager.ArtifactSourcesClientListResponse], errResp azfake.ErrorResponder)
}

// NewArtifactSourcesServerTransport creates a new instance of ArtifactSourcesServerTransport with the provided implementation.
// The returned ArtifactSourcesServerTransport instance is connected to an instance of armdeploymentmanager.ArtifactSourcesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewArtifactSourcesServerTransport(srv *ArtifactSourcesServer) *ArtifactSourcesServerTransport {
	return &ArtifactSourcesServerTransport{srv: srv}
}

// ArtifactSourcesServerTransport connects instances of armdeploymentmanager.ArtifactSourcesClient to instances of ArtifactSourcesServer.
// Don't use this type directly, use NewArtifactSourcesServerTransport instead.
type ArtifactSourcesServerTransport struct {
	srv *ArtifactSourcesServer
}

// Do implements the policy.Transporter interface for ArtifactSourcesServerTransport.
func (a *ArtifactSourcesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ArtifactSourcesClient.CreateOrUpdate":
		resp, err = a.dispatchCreateOrUpdate(req)
	case "ArtifactSourcesClient.Delete":
		resp, err = a.dispatchDelete(req)
	case "ArtifactSourcesClient.Get":
		resp, err = a.dispatchGet(req)
	case "ArtifactSourcesClient.List":
		resp, err = a.dispatchList(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *ArtifactSourcesServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if a.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DeploymentManager/artifactSources/(?P<artifactSourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armdeploymentmanager.ArtifactSource](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	artifactSourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("artifactSourceName")])
	if err != nil {
		return nil, err
	}
	var options *armdeploymentmanager.ArtifactSourcesClientCreateOrUpdateOptions
	if !reflect.ValueOf(body).IsZero() {
		options = &armdeploymentmanager.ArtifactSourcesClientCreateOrUpdateOptions{
			ArtifactSourceInfo: &body,
		}
	}
	respr, errRespr := a.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, artifactSourceNameParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ArtifactSource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *ArtifactSourcesServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if a.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DeploymentManager/artifactSources/(?P<artifactSourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	artifactSourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("artifactSourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Delete(req.Context(), resourceGroupNameParam, artifactSourceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *ArtifactSourcesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DeploymentManager/artifactSources/(?P<artifactSourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	artifactSourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("artifactSourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Get(req.Context(), resourceGroupNameParam, artifactSourceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ArtifactSource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *ArtifactSourcesServerTransport) dispatchList(req *http.Request) (*http.Response, error) {
	if a.srv.List == nil {
		return nil, &nonRetriableError{errors.New("fake for method List not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DeploymentManager/artifactSources`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.List(req.Context(), resourceGroupNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ArtifactSourceArray, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}