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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
	"net/http"
	"net/url"
	"regexp"
)

// IntegrationServiceEnvironmentManagedApisServer is a fake server for instances of the armlogic.IntegrationServiceEnvironmentManagedApisClient type.
type IntegrationServiceEnvironmentManagedApisServer struct {
	// BeginDelete is the fake for method IntegrationServiceEnvironmentManagedApisClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, apiName string, options *armlogic.IntegrationServiceEnvironmentManagedApisClientBeginDeleteOptions) (resp azfake.PollerResponder[armlogic.IntegrationServiceEnvironmentManagedApisClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method IntegrationServiceEnvironmentManagedApisClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, apiName string, options *armlogic.IntegrationServiceEnvironmentManagedApisClientGetOptions) (resp azfake.Responder[armlogic.IntegrationServiceEnvironmentManagedApisClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method IntegrationServiceEnvironmentManagedApisClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroup string, integrationServiceEnvironmentName string, options *armlogic.IntegrationServiceEnvironmentManagedApisClientListOptions) (resp azfake.PagerResponder[armlogic.IntegrationServiceEnvironmentManagedApisClientListResponse])

	// BeginPut is the fake for method IntegrationServiceEnvironmentManagedApisClient.BeginPut
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginPut func(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, apiName string, integrationServiceEnvironmentManagedAPI armlogic.IntegrationServiceEnvironmentManagedAPI, options *armlogic.IntegrationServiceEnvironmentManagedApisClientBeginPutOptions) (resp azfake.PollerResponder[armlogic.IntegrationServiceEnvironmentManagedApisClientPutResponse], errResp azfake.ErrorResponder)
}

// NewIntegrationServiceEnvironmentManagedApisServerTransport creates a new instance of IntegrationServiceEnvironmentManagedApisServerTransport with the provided implementation.
// The returned IntegrationServiceEnvironmentManagedApisServerTransport instance is connected to an instance of armlogic.IntegrationServiceEnvironmentManagedApisClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewIntegrationServiceEnvironmentManagedApisServerTransport(srv *IntegrationServiceEnvironmentManagedApisServer) *IntegrationServiceEnvironmentManagedApisServerTransport {
	return &IntegrationServiceEnvironmentManagedApisServerTransport{
		srv:          srv,
		beginDelete:  newTracker[azfake.PollerResponder[armlogic.IntegrationServiceEnvironmentManagedApisClientDeleteResponse]](),
		newListPager: newTracker[azfake.PagerResponder[armlogic.IntegrationServiceEnvironmentManagedApisClientListResponse]](),
		beginPut:     newTracker[azfake.PollerResponder[armlogic.IntegrationServiceEnvironmentManagedApisClientPutResponse]](),
	}
}

// IntegrationServiceEnvironmentManagedApisServerTransport connects instances of armlogic.IntegrationServiceEnvironmentManagedApisClient to instances of IntegrationServiceEnvironmentManagedApisServer.
// Don't use this type directly, use NewIntegrationServiceEnvironmentManagedApisServerTransport instead.
type IntegrationServiceEnvironmentManagedApisServerTransport struct {
	srv          *IntegrationServiceEnvironmentManagedApisServer
	beginDelete  *tracker[azfake.PollerResponder[armlogic.IntegrationServiceEnvironmentManagedApisClientDeleteResponse]]
	newListPager *tracker[azfake.PagerResponder[armlogic.IntegrationServiceEnvironmentManagedApisClientListResponse]]
	beginPut     *tracker[azfake.PollerResponder[armlogic.IntegrationServiceEnvironmentManagedApisClientPutResponse]]
}

// Do implements the policy.Transporter interface for IntegrationServiceEnvironmentManagedApisServerTransport.
func (i *IntegrationServiceEnvironmentManagedApisServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "IntegrationServiceEnvironmentManagedApisClient.BeginDelete":
		resp, err = i.dispatchBeginDelete(req)
	case "IntegrationServiceEnvironmentManagedApisClient.Get":
		resp, err = i.dispatchGet(req)
	case "IntegrationServiceEnvironmentManagedApisClient.NewListPager":
		resp, err = i.dispatchNewListPager(req)
	case "IntegrationServiceEnvironmentManagedApisClient.BeginPut":
		resp, err = i.dispatchBeginPut(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (i *IntegrationServiceEnvironmentManagedApisServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if i.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := i.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroup>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Logic/integrationServiceEnvironments/(?P<integrationServiceEnvironmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/managedApis/(?P<apiName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroup")])
		if err != nil {
			return nil, err
		}
		integrationServiceEnvironmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("integrationServiceEnvironmentName")])
		if err != nil {
			return nil, err
		}
		apiNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("apiName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := i.srv.BeginDelete(req.Context(), resourceGroupParam, integrationServiceEnvironmentNameParam, apiNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		i.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		i.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		i.beginDelete.remove(req)
	}

	return resp, nil
}

func (i *IntegrationServiceEnvironmentManagedApisServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if i.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroup>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Logic/integrationServiceEnvironments/(?P<integrationServiceEnvironmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/managedApis/(?P<apiName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroup")])
	if err != nil {
		return nil, err
	}
	integrationServiceEnvironmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("integrationServiceEnvironmentName")])
	if err != nil {
		return nil, err
	}
	apiNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("apiName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.Get(req.Context(), resourceGroupParam, integrationServiceEnvironmentNameParam, apiNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).IntegrationServiceEnvironmentManagedAPI, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *IntegrationServiceEnvironmentManagedApisServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if i.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := i.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroup>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Logic/integrationServiceEnvironments/(?P<integrationServiceEnvironmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/managedApis`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroup")])
		if err != nil {
			return nil, err
		}
		integrationServiceEnvironmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("integrationServiceEnvironmentName")])
		if err != nil {
			return nil, err
		}
		resp := i.srv.NewListPager(resourceGroupParam, integrationServiceEnvironmentNameParam, nil)
		newListPager = &resp
		i.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armlogic.IntegrationServiceEnvironmentManagedApisClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		i.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		i.newListPager.remove(req)
	}
	return resp, nil
}

func (i *IntegrationServiceEnvironmentManagedApisServerTransport) dispatchBeginPut(req *http.Request) (*http.Response, error) {
	if i.srv.BeginPut == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginPut not implemented")}
	}
	beginPut := i.beginPut.get(req)
	if beginPut == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroup>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Logic/integrationServiceEnvironments/(?P<integrationServiceEnvironmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/managedApis/(?P<apiName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armlogic.IntegrationServiceEnvironmentManagedAPI](req)
		if err != nil {
			return nil, err
		}
		resourceGroupParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroup")])
		if err != nil {
			return nil, err
		}
		integrationServiceEnvironmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("integrationServiceEnvironmentName")])
		if err != nil {
			return nil, err
		}
		apiNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("apiName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := i.srv.BeginPut(req.Context(), resourceGroupParam, integrationServiceEnvironmentNameParam, apiNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginPut = &respr
		i.beginPut.add(req, beginPut)
	}

	resp, err := server.PollerResponderNext(beginPut, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		i.beginPut.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginPut) {
		i.beginPut.remove(req)
	}

	return resp, nil
}