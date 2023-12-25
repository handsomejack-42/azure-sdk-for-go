//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
	"sync"
)

// ServerFactory is a fake server for instances of the armm365securityandcompliance.ClientFactory type.
type ServerFactory struct {
	OperationResultsServer                                OperationResultsServer
	OperationsServer                                      OperationsServer
	PrivateEndpointConnectionsAdtAPIServer                PrivateEndpointConnectionsAdtAPIServer
	PrivateEndpointConnectionsCompServer                  PrivateEndpointConnectionsCompServer
	PrivateEndpointConnectionsForEDMServer                PrivateEndpointConnectionsForEDMServer
	PrivateEndpointConnectionsForMIPPolicySyncServer      PrivateEndpointConnectionsForMIPPolicySyncServer
	PrivateEndpointConnectionsForSCCPowershellServer      PrivateEndpointConnectionsForSCCPowershellServer
	PrivateEndpointConnectionsSecServer                   PrivateEndpointConnectionsSecServer
	PrivateLinkResourcesAdtAPIServer                      PrivateLinkResourcesAdtAPIServer
	PrivateLinkResourcesServer                            PrivateLinkResourcesServer
	PrivateLinkResourcesCompServer                        PrivateLinkResourcesCompServer
	PrivateLinkResourcesForMIPPolicySyncServer            PrivateLinkResourcesForMIPPolicySyncServer
	PrivateLinkResourcesForSCCPowershellServer            PrivateLinkResourcesForSCCPowershellServer
	PrivateLinkResourcesSecServer                         PrivateLinkResourcesSecServer
	PrivateLinkServicesForEDMUploadServer                 PrivateLinkServicesForEDMUploadServer
	PrivateLinkServicesForM365ComplianceCenterServer      PrivateLinkServicesForM365ComplianceCenterServer
	PrivateLinkServicesForM365SecurityCenterServer        PrivateLinkServicesForM365SecurityCenterServer
	PrivateLinkServicesForMIPPolicySyncServer             PrivateLinkServicesForMIPPolicySyncServer
	PrivateLinkServicesForO365ManagementActivityAPIServer PrivateLinkServicesForO365ManagementActivityAPIServer
	PrivateLinkServicesForSCCPowershellServer             PrivateLinkServicesForSCCPowershellServer
	ServicesServer                                        ServicesServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armm365securityandcompliance.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armm365securityandcompliance.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                                                     *ServerFactory
	trMu                                                    sync.Mutex
	trOperationResultsServer                                *OperationResultsServerTransport
	trOperationsServer                                      *OperationsServerTransport
	trPrivateEndpointConnectionsAdtAPIServer                *PrivateEndpointConnectionsAdtAPIServerTransport
	trPrivateEndpointConnectionsCompServer                  *PrivateEndpointConnectionsCompServerTransport
	trPrivateEndpointConnectionsForEDMServer                *PrivateEndpointConnectionsForEDMServerTransport
	trPrivateEndpointConnectionsForMIPPolicySyncServer      *PrivateEndpointConnectionsForMIPPolicySyncServerTransport
	trPrivateEndpointConnectionsForSCCPowershellServer      *PrivateEndpointConnectionsForSCCPowershellServerTransport
	trPrivateEndpointConnectionsSecServer                   *PrivateEndpointConnectionsSecServerTransport
	trPrivateLinkResourcesAdtAPIServer                      *PrivateLinkResourcesAdtAPIServerTransport
	trPrivateLinkResourcesServer                            *PrivateLinkResourcesServerTransport
	trPrivateLinkResourcesCompServer                        *PrivateLinkResourcesCompServerTransport
	trPrivateLinkResourcesForMIPPolicySyncServer            *PrivateLinkResourcesForMIPPolicySyncServerTransport
	trPrivateLinkResourcesForSCCPowershellServer            *PrivateLinkResourcesForSCCPowershellServerTransport
	trPrivateLinkResourcesSecServer                         *PrivateLinkResourcesSecServerTransport
	trPrivateLinkServicesForEDMUploadServer                 *PrivateLinkServicesForEDMUploadServerTransport
	trPrivateLinkServicesForM365ComplianceCenterServer      *PrivateLinkServicesForM365ComplianceCenterServerTransport
	trPrivateLinkServicesForM365SecurityCenterServer        *PrivateLinkServicesForM365SecurityCenterServerTransport
	trPrivateLinkServicesForMIPPolicySyncServer             *PrivateLinkServicesForMIPPolicySyncServerTransport
	trPrivateLinkServicesForO365ManagementActivityAPIServer *PrivateLinkServicesForO365ManagementActivityAPIServerTransport
	trPrivateLinkServicesForSCCPowershellServer             *PrivateLinkServicesForSCCPowershellServerTransport
	trServicesServer                                        *ServicesServerTransport
}

// Do implements the policy.Transporter interface for ServerFactoryTransport.
func (s *ServerFactoryTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	client := method[:strings.Index(method, ".")]
	var resp *http.Response
	var err error

	switch client {
	case "OperationResultsClient":
		initServer(s, &s.trOperationResultsServer, func() *OperationResultsServerTransport {
			return NewOperationResultsServerTransport(&s.srv.OperationResultsServer)
		})
		resp, err = s.trOperationResultsServer.Do(req)
	case "OperationsClient":
		initServer(s, &s.trOperationsServer, func() *OperationsServerTransport { return NewOperationsServerTransport(&s.srv.OperationsServer) })
		resp, err = s.trOperationsServer.Do(req)
	case "PrivateEndpointConnectionsAdtAPIClient":
		initServer(s, &s.trPrivateEndpointConnectionsAdtAPIServer, func() *PrivateEndpointConnectionsAdtAPIServerTransport {
			return NewPrivateEndpointConnectionsAdtAPIServerTransport(&s.srv.PrivateEndpointConnectionsAdtAPIServer)
		})
		resp, err = s.trPrivateEndpointConnectionsAdtAPIServer.Do(req)
	case "PrivateEndpointConnectionsCompClient":
		initServer(s, &s.trPrivateEndpointConnectionsCompServer, func() *PrivateEndpointConnectionsCompServerTransport {
			return NewPrivateEndpointConnectionsCompServerTransport(&s.srv.PrivateEndpointConnectionsCompServer)
		})
		resp, err = s.trPrivateEndpointConnectionsCompServer.Do(req)
	case "PrivateEndpointConnectionsForEDMClient":
		initServer(s, &s.trPrivateEndpointConnectionsForEDMServer, func() *PrivateEndpointConnectionsForEDMServerTransport {
			return NewPrivateEndpointConnectionsForEDMServerTransport(&s.srv.PrivateEndpointConnectionsForEDMServer)
		})
		resp, err = s.trPrivateEndpointConnectionsForEDMServer.Do(req)
	case "PrivateEndpointConnectionsForMIPPolicySyncClient":
		initServer(s, &s.trPrivateEndpointConnectionsForMIPPolicySyncServer, func() *PrivateEndpointConnectionsForMIPPolicySyncServerTransport {
			return NewPrivateEndpointConnectionsForMIPPolicySyncServerTransport(&s.srv.PrivateEndpointConnectionsForMIPPolicySyncServer)
		})
		resp, err = s.trPrivateEndpointConnectionsForMIPPolicySyncServer.Do(req)
	case "PrivateEndpointConnectionsForSCCPowershellClient":
		initServer(s, &s.trPrivateEndpointConnectionsForSCCPowershellServer, func() *PrivateEndpointConnectionsForSCCPowershellServerTransport {
			return NewPrivateEndpointConnectionsForSCCPowershellServerTransport(&s.srv.PrivateEndpointConnectionsForSCCPowershellServer)
		})
		resp, err = s.trPrivateEndpointConnectionsForSCCPowershellServer.Do(req)
	case "PrivateEndpointConnectionsSecClient":
		initServer(s, &s.trPrivateEndpointConnectionsSecServer, func() *PrivateEndpointConnectionsSecServerTransport {
			return NewPrivateEndpointConnectionsSecServerTransport(&s.srv.PrivateEndpointConnectionsSecServer)
		})
		resp, err = s.trPrivateEndpointConnectionsSecServer.Do(req)
	case "PrivateLinkResourcesAdtAPIClient":
		initServer(s, &s.trPrivateLinkResourcesAdtAPIServer, func() *PrivateLinkResourcesAdtAPIServerTransport {
			return NewPrivateLinkResourcesAdtAPIServerTransport(&s.srv.PrivateLinkResourcesAdtAPIServer)
		})
		resp, err = s.trPrivateLinkResourcesAdtAPIServer.Do(req)
	case "PrivateLinkResourcesClient":
		initServer(s, &s.trPrivateLinkResourcesServer, func() *PrivateLinkResourcesServerTransport {
			return NewPrivateLinkResourcesServerTransport(&s.srv.PrivateLinkResourcesServer)
		})
		resp, err = s.trPrivateLinkResourcesServer.Do(req)
	case "PrivateLinkResourcesCompClient":
		initServer(s, &s.trPrivateLinkResourcesCompServer, func() *PrivateLinkResourcesCompServerTransport {
			return NewPrivateLinkResourcesCompServerTransport(&s.srv.PrivateLinkResourcesCompServer)
		})
		resp, err = s.trPrivateLinkResourcesCompServer.Do(req)
	case "PrivateLinkResourcesForMIPPolicySyncClient":
		initServer(s, &s.trPrivateLinkResourcesForMIPPolicySyncServer, func() *PrivateLinkResourcesForMIPPolicySyncServerTransport {
			return NewPrivateLinkResourcesForMIPPolicySyncServerTransport(&s.srv.PrivateLinkResourcesForMIPPolicySyncServer)
		})
		resp, err = s.trPrivateLinkResourcesForMIPPolicySyncServer.Do(req)
	case "PrivateLinkResourcesForSCCPowershellClient":
		initServer(s, &s.trPrivateLinkResourcesForSCCPowershellServer, func() *PrivateLinkResourcesForSCCPowershellServerTransport {
			return NewPrivateLinkResourcesForSCCPowershellServerTransport(&s.srv.PrivateLinkResourcesForSCCPowershellServer)
		})
		resp, err = s.trPrivateLinkResourcesForSCCPowershellServer.Do(req)
	case "PrivateLinkResourcesSecClient":
		initServer(s, &s.trPrivateLinkResourcesSecServer, func() *PrivateLinkResourcesSecServerTransport {
			return NewPrivateLinkResourcesSecServerTransport(&s.srv.PrivateLinkResourcesSecServer)
		})
		resp, err = s.trPrivateLinkResourcesSecServer.Do(req)
	case "PrivateLinkServicesForEDMUploadClient":
		initServer(s, &s.trPrivateLinkServicesForEDMUploadServer, func() *PrivateLinkServicesForEDMUploadServerTransport {
			return NewPrivateLinkServicesForEDMUploadServerTransport(&s.srv.PrivateLinkServicesForEDMUploadServer)
		})
		resp, err = s.trPrivateLinkServicesForEDMUploadServer.Do(req)
	case "PrivateLinkServicesForM365ComplianceCenterClient":
		initServer(s, &s.trPrivateLinkServicesForM365ComplianceCenterServer, func() *PrivateLinkServicesForM365ComplianceCenterServerTransport {
			return NewPrivateLinkServicesForM365ComplianceCenterServerTransport(&s.srv.PrivateLinkServicesForM365ComplianceCenterServer)
		})
		resp, err = s.trPrivateLinkServicesForM365ComplianceCenterServer.Do(req)
	case "PrivateLinkServicesForM365SecurityCenterClient":
		initServer(s, &s.trPrivateLinkServicesForM365SecurityCenterServer, func() *PrivateLinkServicesForM365SecurityCenterServerTransport {
			return NewPrivateLinkServicesForM365SecurityCenterServerTransport(&s.srv.PrivateLinkServicesForM365SecurityCenterServer)
		})
		resp, err = s.trPrivateLinkServicesForM365SecurityCenterServer.Do(req)
	case "PrivateLinkServicesForMIPPolicySyncClient":
		initServer(s, &s.trPrivateLinkServicesForMIPPolicySyncServer, func() *PrivateLinkServicesForMIPPolicySyncServerTransport {
			return NewPrivateLinkServicesForMIPPolicySyncServerTransport(&s.srv.PrivateLinkServicesForMIPPolicySyncServer)
		})
		resp, err = s.trPrivateLinkServicesForMIPPolicySyncServer.Do(req)
	case "PrivateLinkServicesForO365ManagementActivityAPIClient":
		initServer(s, &s.trPrivateLinkServicesForO365ManagementActivityAPIServer, func() *PrivateLinkServicesForO365ManagementActivityAPIServerTransport {
			return NewPrivateLinkServicesForO365ManagementActivityAPIServerTransport(&s.srv.PrivateLinkServicesForO365ManagementActivityAPIServer)
		})
		resp, err = s.trPrivateLinkServicesForO365ManagementActivityAPIServer.Do(req)
	case "PrivateLinkServicesForSCCPowershellClient":
		initServer(s, &s.trPrivateLinkServicesForSCCPowershellServer, func() *PrivateLinkServicesForSCCPowershellServerTransport {
			return NewPrivateLinkServicesForSCCPowershellServerTransport(&s.srv.PrivateLinkServicesForSCCPowershellServer)
		})
		resp, err = s.trPrivateLinkServicesForSCCPowershellServer.Do(req)
	case "ServicesClient":
		initServer(s, &s.trServicesServer, func() *ServicesServerTransport { return NewServicesServerTransport(&s.srv.ServicesServer) })
		resp, err = s.trServicesServer.Do(req)
	default:
		err = fmt.Errorf("unhandled client %s", client)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func initServer[T any](s *ServerFactoryTransport, dst **T, src func() *T) {
	s.trMu.Lock()
	if *dst == nil {
		*dst = src()
	}
	s.trMu.Unlock()
}