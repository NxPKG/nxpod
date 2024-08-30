// Copyright (c) 2024 Nxpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

// Code generated by protoc-proxy-gen. DO NOT EDIT.

package v1connect

import (
	context "context"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/nxpkg/nxpod/components/public-api/go/v1"
)

var _ AuditLogServiceHandler = (*ProxyAuditLogServiceHandler)(nil)

type ProxyAuditLogServiceHandler struct {
	Client v1.AuditLogServiceClient
	UnimplementedAuditLogServiceHandler
}

func (s *ProxyAuditLogServiceHandler) ListAuditLogs(ctx context.Context, req *connect_go.Request[v1.ListAuditLogsRequest]) (*connect_go.Response[v1.ListAuditLogsResponse], error) {
	resp, err := s.Client.ListAuditLogs(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}
