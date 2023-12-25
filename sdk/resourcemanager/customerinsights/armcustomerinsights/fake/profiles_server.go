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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights"
	"net/http"
	"net/url"
	"regexp"
)

// ProfilesServer is a fake server for instances of the armcustomerinsights.ProfilesClient type.
type ProfilesServer struct {
	// BeginCreateOrUpdate is the fake for method ProfilesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, hubName string, profileName string, parameters armcustomerinsights.ProfileResourceFormat, options *armcustomerinsights.ProfilesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armcustomerinsights.ProfilesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method ProfilesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, hubName string, profileName string, options *armcustomerinsights.ProfilesClientBeginDeleteOptions) (resp azfake.PollerResponder[armcustomerinsights.ProfilesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ProfilesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, hubName string, profileName string, options *armcustomerinsights.ProfilesClientGetOptions) (resp azfake.Responder[armcustomerinsights.ProfilesClientGetResponse], errResp azfake.ErrorResponder)

	// GetEnrichingKpis is the fake for method ProfilesClient.GetEnrichingKpis
	// HTTP status codes to indicate success: http.StatusOK
	GetEnrichingKpis func(ctx context.Context, resourceGroupName string, hubName string, profileName string, options *armcustomerinsights.ProfilesClientGetEnrichingKpisOptions) (resp azfake.Responder[armcustomerinsights.ProfilesClientGetEnrichingKpisResponse], errResp azfake.ErrorResponder)

	// NewListByHubPager is the fake for method ProfilesClient.NewListByHubPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByHubPager func(resourceGroupName string, hubName string, options *armcustomerinsights.ProfilesClientListByHubOptions) (resp azfake.PagerResponder[armcustomerinsights.ProfilesClientListByHubResponse])
}

// NewProfilesServerTransport creates a new instance of ProfilesServerTransport with the provided implementation.
// The returned ProfilesServerTransport instance is connected to an instance of armcustomerinsights.ProfilesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewProfilesServerTransport(srv *ProfilesServer) *ProfilesServerTransport {
	return &ProfilesServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armcustomerinsights.ProfilesClientCreateOrUpdateResponse]](),
		beginDelete:         newTracker[azfake.PollerResponder[armcustomerinsights.ProfilesClientDeleteResponse]](),
		newListByHubPager:   newTracker[azfake.PagerResponder[armcustomerinsights.ProfilesClientListByHubResponse]](),
	}
}

// ProfilesServerTransport connects instances of armcustomerinsights.ProfilesClient to instances of ProfilesServer.
// Don't use this type directly, use NewProfilesServerTransport instead.
type ProfilesServerTransport struct {
	srv                 *ProfilesServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armcustomerinsights.ProfilesClientCreateOrUpdateResponse]]
	beginDelete         *tracker[azfake.PollerResponder[armcustomerinsights.ProfilesClientDeleteResponse]]
	newListByHubPager   *tracker[azfake.PagerResponder[armcustomerinsights.ProfilesClientListByHubResponse]]
}

// Do implements the policy.Transporter interface for ProfilesServerTransport.
func (p *ProfilesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ProfilesClient.BeginCreateOrUpdate":
		resp, err = p.dispatchBeginCreateOrUpdate(req)
	case "ProfilesClient.BeginDelete":
		resp, err = p.dispatchBeginDelete(req)
	case "ProfilesClient.Get":
		resp, err = p.dispatchGet(req)
	case "ProfilesClient.GetEnrichingKpis":
		resp, err = p.dispatchGetEnrichingKpis(req)
	case "ProfilesClient.NewListByHubPager":
		resp, err = p.dispatchNewListByHubPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *ProfilesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if p.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := p.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CustomerInsights/hubs/(?P<hubName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/profiles/(?P<profileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcustomerinsights.ProfileResourceFormat](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		hubNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("hubName")])
		if err != nil {
			return nil, err
		}
		profileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("profileName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, hubNameParam, profileNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		p.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		p.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		p.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (p *ProfilesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if p.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := p.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CustomerInsights/hubs/(?P<hubName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/profiles/(?P<profileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		hubNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("hubName")])
		if err != nil {
			return nil, err
		}
		profileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("profileName")])
		if err != nil {
			return nil, err
		}
		localeCodeUnescaped, err := url.QueryUnescape(qp.Get("locale-code"))
		if err != nil {
			return nil, err
		}
		localeCodeParam := getOptional(localeCodeUnescaped)
		var options *armcustomerinsights.ProfilesClientBeginDeleteOptions
		if localeCodeParam != nil {
			options = &armcustomerinsights.ProfilesClientBeginDeleteOptions{
				LocaleCode: localeCodeParam,
			}
		}
		respr, errRespr := p.srv.BeginDelete(req.Context(), resourceGroupNameParam, hubNameParam, profileNameParam, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		p.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		p.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		p.beginDelete.remove(req)
	}

	return resp, nil
}

func (p *ProfilesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CustomerInsights/hubs/(?P<hubName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/profiles/(?P<profileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	hubNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("hubName")])
	if err != nil {
		return nil, err
	}
	profileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("profileName")])
	if err != nil {
		return nil, err
	}
	localeCodeUnescaped, err := url.QueryUnescape(qp.Get("locale-code"))
	if err != nil {
		return nil, err
	}
	localeCodeParam := getOptional(localeCodeUnescaped)
	var options *armcustomerinsights.ProfilesClientGetOptions
	if localeCodeParam != nil {
		options = &armcustomerinsights.ProfilesClientGetOptions{
			LocaleCode: localeCodeParam,
		}
	}
	respr, errRespr := p.srv.Get(req.Context(), resourceGroupNameParam, hubNameParam, profileNameParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ProfileResourceFormat, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *ProfilesServerTransport) dispatchGetEnrichingKpis(req *http.Request) (*http.Response, error) {
	if p.srv.GetEnrichingKpis == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetEnrichingKpis not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CustomerInsights/hubs/(?P<hubName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/profiles/(?P<profileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/getEnrichingKpis`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	hubNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("hubName")])
	if err != nil {
		return nil, err
	}
	profileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("profileName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.GetEnrichingKpis(req.Context(), resourceGroupNameParam, hubNameParam, profileNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).KpiDefinitionArray, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *ProfilesServerTransport) dispatchNewListByHubPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListByHubPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByHubPager not implemented")}
	}
	newListByHubPager := p.newListByHubPager.get(req)
	if newListByHubPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CustomerInsights/hubs/(?P<hubName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/profiles`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		hubNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("hubName")])
		if err != nil {
			return nil, err
		}
		localeCodeUnescaped, err := url.QueryUnescape(qp.Get("locale-code"))
		if err != nil {
			return nil, err
		}
		localeCodeParam := getOptional(localeCodeUnescaped)
		var options *armcustomerinsights.ProfilesClientListByHubOptions
		if localeCodeParam != nil {
			options = &armcustomerinsights.ProfilesClientListByHubOptions{
				LocaleCode: localeCodeParam,
			}
		}
		resp := p.srv.NewListByHubPager(resourceGroupNameParam, hubNameParam, options)
		newListByHubPager = &resp
		p.newListByHubPager.add(req, newListByHubPager)
		server.PagerResponderInjectNextLinks(newListByHubPager, req, func(page *armcustomerinsights.ProfilesClientListByHubResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByHubPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newListByHubPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByHubPager) {
		p.newListByHubPager.remove(req)
	}
	return resp, nil
}