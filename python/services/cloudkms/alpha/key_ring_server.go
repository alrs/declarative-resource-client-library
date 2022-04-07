// Copyright 2022 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package server

import (
	"context"
	"errors"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/cloudkms/alpha/cloudkms_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudkms/alpha"
)

// KeyRingServer implements the gRPC interface for KeyRing.
type KeyRingServer struct{}

// ProtoToKeyRing converts a KeyRing resource from its proto representation.
func ProtoToKeyRing(p *alphapb.CloudkmsAlphaKeyRing) *alpha.KeyRing {
	obj := &alpha.KeyRing{
		Name:       dcl.StringOrNil(p.GetName()),
		CreateTime: dcl.StringOrNil(p.GetCreateTime()),
		Project:    dcl.StringOrNil(p.GetProject()),
		Location:   dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// KeyRingToProto converts a KeyRing resource to its proto representation.
func KeyRingToProto(resource *alpha.KeyRing) *alphapb.CloudkmsAlphaKeyRing {
	p := &alphapb.CloudkmsAlphaKeyRing{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))

	return p
}

// applyKeyRing handles the gRPC request by passing it to the underlying KeyRing Apply() method.
func (s *KeyRingServer) applyKeyRing(ctx context.Context, c *alpha.Client, request *alphapb.ApplyCloudkmsAlphaKeyRingRequest) (*alphapb.CloudkmsAlphaKeyRing, error) {
	p := ProtoToKeyRing(request.GetResource())
	res, err := c.ApplyKeyRing(ctx, p)
	if err != nil {
		return nil, err
	}
	r := KeyRingToProto(res)
	return r, nil
}

// applyCloudkmsAlphaKeyRing handles the gRPC request by passing it to the underlying KeyRing Apply() method.
func (s *KeyRingServer) ApplyCloudkmsAlphaKeyRing(ctx context.Context, request *alphapb.ApplyCloudkmsAlphaKeyRingRequest) (*alphapb.CloudkmsAlphaKeyRing, error) {
	cl, err := createConfigKeyRing(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyKeyRing(ctx, cl, request)
}

// DeleteKeyRing handles the gRPC request by passing it to the underlying KeyRing Delete() method.
func (s *KeyRingServer) DeleteCloudkmsAlphaKeyRing(ctx context.Context, request *alphapb.DeleteCloudkmsAlphaKeyRingRequest) (*emptypb.Empty, error) {

	return nil, errors.New("no delete endpoint for KeyRing")

}

// ListCloudkmsAlphaKeyRing handles the gRPC request by passing it to the underlying KeyRingList() method.
func (s *KeyRingServer) ListCloudkmsAlphaKeyRing(ctx context.Context, request *alphapb.ListCloudkmsAlphaKeyRingRequest) (*alphapb.ListCloudkmsAlphaKeyRingResponse, error) {
	cl, err := createConfigKeyRing(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListKeyRing(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.CloudkmsAlphaKeyRing
	for _, r := range resources.Items {
		rp := KeyRingToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListCloudkmsAlphaKeyRingResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigKeyRing(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
