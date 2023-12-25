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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/saas/armsaas"
	"net/http"
	"net/url"
	"regexp"
)

// OperationServer is a fake server for instances of the armsaas.OperationClient type.
type OperationServer struct {
	// BeginGet is the fake for method OperationClient.BeginGet
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginGet func(ctx context.Context, operationID string, options *armsaas.OperationClientBeginGetOptions) (resp azfake.PollerResponder[armsaas.OperationClientGetResponse], errResp azfake.ErrorResponder)
}

// NewOperationServerTransport creates a new instance of OperationServerTransport with the provided implementation.
// The returned OperationServerTransport instance is connected to an instance of armsaas.OperationClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewOperationServerTransport(srv *OperationServer) *OperationServerTransport {
	return &OperationServerTransport{
		srv:      srv,
		beginGet: newTracker[azfake.PollerResponder[armsaas.OperationClientGetResponse]](),
	}
}

// OperationServerTransport connects instances of armsaas.OperationClient to instances of OperationServer.
// Don't use this type directly, use NewOperationServerTransport instead.
type OperationServerTransport struct {
	srv      *OperationServer
	beginGet *tracker[azfake.PollerResponder[armsaas.OperationClientGetResponse]]
}

// Do implements the policy.Transporter interface for OperationServerTransport.
func (o *OperationServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "OperationClient.BeginGet":
		resp, err = o.dispatchBeginGet(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (o *OperationServerTransport) dispatchBeginGet(req *http.Request) (*http.Response, error) {
	if o.srv.BeginGet == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginGet not implemented")}
	}
	beginGet := o.beginGet.get(req)
	if beginGet == nil {
		const regexStr = `/providers/Microsoft\.SaaS/operationResults/(?P<operationId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		operationIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("operationId")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := o.srv.BeginGet(req.Context(), operationIDParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginGet = &respr
		o.beginGet.add(req, beginGet)
	}

	resp, err := server.PollerResponderNext(beginGet, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		o.beginGet.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginGet) {
		o.beginGet.remove(req)
	}

	return resp, nil
}