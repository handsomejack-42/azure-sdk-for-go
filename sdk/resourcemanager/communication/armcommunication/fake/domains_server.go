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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/communication/armcommunication/v2"
	"net/http"
	"net/url"
	"regexp"
)

// DomainsServer is a fake server for instances of the armcommunication.DomainsClient type.
type DomainsServer struct {
	// BeginCancelVerification is the fake for method DomainsClient.BeginCancelVerification
	// HTTP status codes to indicate success: http.StatusAccepted
	BeginCancelVerification func(ctx context.Context, resourceGroupName string, emailServiceName string, domainName string, parameters armcommunication.VerificationParameter, options *armcommunication.DomainsClientBeginCancelVerificationOptions) (resp azfake.PollerResponder[armcommunication.DomainsClientCancelVerificationResponse], errResp azfake.ErrorResponder)

	// BeginCreateOrUpdate is the fake for method DomainsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, emailServiceName string, domainName string, parameters armcommunication.DomainResource, options *armcommunication.DomainsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armcommunication.DomainsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method DomainsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, emailServiceName string, domainName string, options *armcommunication.DomainsClientBeginDeleteOptions) (resp azfake.PollerResponder[armcommunication.DomainsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method DomainsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, emailServiceName string, domainName string, options *armcommunication.DomainsClientGetOptions) (resp azfake.Responder[armcommunication.DomainsClientGetResponse], errResp azfake.ErrorResponder)

	// BeginInitiateVerification is the fake for method DomainsClient.BeginInitiateVerification
	// HTTP status codes to indicate success: http.StatusAccepted
	BeginInitiateVerification func(ctx context.Context, resourceGroupName string, emailServiceName string, domainName string, parameters armcommunication.VerificationParameter, options *armcommunication.DomainsClientBeginInitiateVerificationOptions) (resp azfake.PollerResponder[armcommunication.DomainsClientInitiateVerificationResponse], errResp azfake.ErrorResponder)

	// NewListByEmailServiceResourcePager is the fake for method DomainsClient.NewListByEmailServiceResourcePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByEmailServiceResourcePager func(resourceGroupName string, emailServiceName string, options *armcommunication.DomainsClientListByEmailServiceResourceOptions) (resp azfake.PagerResponder[armcommunication.DomainsClientListByEmailServiceResourceResponse])

	// BeginUpdate is the fake for method DomainsClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginUpdate func(ctx context.Context, resourceGroupName string, emailServiceName string, domainName string, parameters armcommunication.UpdateDomainRequestParameters, options *armcommunication.DomainsClientBeginUpdateOptions) (resp azfake.PollerResponder[armcommunication.DomainsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewDomainsServerTransport creates a new instance of DomainsServerTransport with the provided implementation.
// The returned DomainsServerTransport instance is connected to an instance of armcommunication.DomainsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewDomainsServerTransport(srv *DomainsServer) *DomainsServerTransport {
	return &DomainsServerTransport{
		srv:                                srv,
		beginCancelVerification:            newTracker[azfake.PollerResponder[armcommunication.DomainsClientCancelVerificationResponse]](),
		beginCreateOrUpdate:                newTracker[azfake.PollerResponder[armcommunication.DomainsClientCreateOrUpdateResponse]](),
		beginDelete:                        newTracker[azfake.PollerResponder[armcommunication.DomainsClientDeleteResponse]](),
		beginInitiateVerification:          newTracker[azfake.PollerResponder[armcommunication.DomainsClientInitiateVerificationResponse]](),
		newListByEmailServiceResourcePager: newTracker[azfake.PagerResponder[armcommunication.DomainsClientListByEmailServiceResourceResponse]](),
		beginUpdate:                        newTracker[azfake.PollerResponder[armcommunication.DomainsClientUpdateResponse]](),
	}
}

// DomainsServerTransport connects instances of armcommunication.DomainsClient to instances of DomainsServer.
// Don't use this type directly, use NewDomainsServerTransport instead.
type DomainsServerTransport struct {
	srv                                *DomainsServer
	beginCancelVerification            *tracker[azfake.PollerResponder[armcommunication.DomainsClientCancelVerificationResponse]]
	beginCreateOrUpdate                *tracker[azfake.PollerResponder[armcommunication.DomainsClientCreateOrUpdateResponse]]
	beginDelete                        *tracker[azfake.PollerResponder[armcommunication.DomainsClientDeleteResponse]]
	beginInitiateVerification          *tracker[azfake.PollerResponder[armcommunication.DomainsClientInitiateVerificationResponse]]
	newListByEmailServiceResourcePager *tracker[azfake.PagerResponder[armcommunication.DomainsClientListByEmailServiceResourceResponse]]
	beginUpdate                        *tracker[azfake.PollerResponder[armcommunication.DomainsClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for DomainsServerTransport.
func (d *DomainsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "DomainsClient.BeginCancelVerification":
		resp, err = d.dispatchBeginCancelVerification(req)
	case "DomainsClient.BeginCreateOrUpdate":
		resp, err = d.dispatchBeginCreateOrUpdate(req)
	case "DomainsClient.BeginDelete":
		resp, err = d.dispatchBeginDelete(req)
	case "DomainsClient.Get":
		resp, err = d.dispatchGet(req)
	case "DomainsClient.BeginInitiateVerification":
		resp, err = d.dispatchBeginInitiateVerification(req)
	case "DomainsClient.NewListByEmailServiceResourcePager":
		resp, err = d.dispatchNewListByEmailServiceResourcePager(req)
	case "DomainsClient.BeginUpdate":
		resp, err = d.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (d *DomainsServerTransport) dispatchBeginCancelVerification(req *http.Request) (*http.Response, error) {
	if d.srv.BeginCancelVerification == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCancelVerification not implemented")}
	}
	beginCancelVerification := d.beginCancelVerification.get(req)
	if beginCancelVerification == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Communication/emailServices/(?P<emailServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/domains/(?P<domainName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/cancelVerification`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcommunication.VerificationParameter](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		emailServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("emailServiceName")])
		if err != nil {
			return nil, err
		}
		domainNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("domainName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginCancelVerification(req.Context(), resourceGroupNameParam, emailServiceNameParam, domainNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCancelVerification = &respr
		d.beginCancelVerification.add(req, beginCancelVerification)
	}

	resp, err := server.PollerResponderNext(beginCancelVerification, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted}, resp.StatusCode) {
		d.beginCancelVerification.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCancelVerification) {
		d.beginCancelVerification.remove(req)
	}

	return resp, nil
}

func (d *DomainsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if d.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := d.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Communication/emailServices/(?P<emailServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/domains/(?P<domainName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcommunication.DomainResource](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		emailServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("emailServiceName")])
		if err != nil {
			return nil, err
		}
		domainNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("domainName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, emailServiceNameParam, domainNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		d.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		d.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		d.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (d *DomainsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if d.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := d.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Communication/emailServices/(?P<emailServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/domains/(?P<domainName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		emailServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("emailServiceName")])
		if err != nil {
			return nil, err
		}
		domainNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("domainName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginDelete(req.Context(), resourceGroupNameParam, emailServiceNameParam, domainNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		d.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		d.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		d.beginDelete.remove(req)
	}

	return resp, nil
}

func (d *DomainsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if d.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Communication/emailServices/(?P<emailServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/domains/(?P<domainName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	emailServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("emailServiceName")])
	if err != nil {
		return nil, err
	}
	domainNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("domainName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Get(req.Context(), resourceGroupNameParam, emailServiceNameParam, domainNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DomainResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DomainsServerTransport) dispatchBeginInitiateVerification(req *http.Request) (*http.Response, error) {
	if d.srv.BeginInitiateVerification == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginInitiateVerification not implemented")}
	}
	beginInitiateVerification := d.beginInitiateVerification.get(req)
	if beginInitiateVerification == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Communication/emailServices/(?P<emailServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/domains/(?P<domainName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/initiateVerification`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcommunication.VerificationParameter](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		emailServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("emailServiceName")])
		if err != nil {
			return nil, err
		}
		domainNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("domainName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginInitiateVerification(req.Context(), resourceGroupNameParam, emailServiceNameParam, domainNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginInitiateVerification = &respr
		d.beginInitiateVerification.add(req, beginInitiateVerification)
	}

	resp, err := server.PollerResponderNext(beginInitiateVerification, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted}, resp.StatusCode) {
		d.beginInitiateVerification.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginInitiateVerification) {
		d.beginInitiateVerification.remove(req)
	}

	return resp, nil
}

func (d *DomainsServerTransport) dispatchNewListByEmailServiceResourcePager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListByEmailServiceResourcePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByEmailServiceResourcePager not implemented")}
	}
	newListByEmailServiceResourcePager := d.newListByEmailServiceResourcePager.get(req)
	if newListByEmailServiceResourcePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Communication/emailServices/(?P<emailServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/domains`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		emailServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("emailServiceName")])
		if err != nil {
			return nil, err
		}
		resp := d.srv.NewListByEmailServiceResourcePager(resourceGroupNameParam, emailServiceNameParam, nil)
		newListByEmailServiceResourcePager = &resp
		d.newListByEmailServiceResourcePager.add(req, newListByEmailServiceResourcePager)
		server.PagerResponderInjectNextLinks(newListByEmailServiceResourcePager, req, func(page *armcommunication.DomainsClientListByEmailServiceResourceResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByEmailServiceResourcePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		d.newListByEmailServiceResourcePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByEmailServiceResourcePager) {
		d.newListByEmailServiceResourcePager.remove(req)
	}
	return resp, nil
}

func (d *DomainsServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if d.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := d.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Communication/emailServices/(?P<emailServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/domains/(?P<domainName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcommunication.UpdateDomainRequestParameters](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		emailServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("emailServiceName")])
		if err != nil {
			return nil, err
		}
		domainNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("domainName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginUpdate(req.Context(), resourceGroupNameParam, emailServiceNameParam, domainNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		d.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		d.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		d.beginUpdate.remove(req)
	}

	return resp, nil
}