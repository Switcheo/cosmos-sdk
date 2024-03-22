// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cosmos/bank/v1beta1/query.proto

package bankv1beta1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Query_Balance_FullMethodName                    = "/cosmos.bank.v1beta1.Query/Balance"
	Query_AllBalances_FullMethodName                = "/cosmos.bank.v1beta1.Query/AllBalances"
	Query_SpendableBalances_FullMethodName          = "/cosmos.bank.v1beta1.Query/SpendableBalances"
	Query_SpendableBalanceByDenom_FullMethodName    = "/cosmos.bank.v1beta1.Query/SpendableBalanceByDenom"
	Query_TotalSupply_FullMethodName                = "/cosmos.bank.v1beta1.Query/TotalSupply"
	Query_SupplyOf_FullMethodName                   = "/cosmos.bank.v1beta1.Query/SupplyOf"
	Query_Params_FullMethodName                     = "/cosmos.bank.v1beta1.Query/Params"
	Query_DenomMetadata_FullMethodName              = "/cosmos.bank.v1beta1.Query/DenomMetadata"
	Query_DenomMetadataByQueryString_FullMethodName = "/cosmos.bank.v1beta1.Query/DenomMetadataByQueryString"
	Query_DenomsMetadata_FullMethodName             = "/cosmos.bank.v1beta1.Query/DenomsMetadata"
	Query_DenomOwners_FullMethodName                = "/cosmos.bank.v1beta1.Query/DenomOwners"
	Query_DenomOwnersByQuery_FullMethodName         = "/cosmos.bank.v1beta1.Query/DenomOwnersByQuery"
	Query_SendEnabled_FullMethodName                = "/cosmos.bank.v1beta1.Query/SendEnabled"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Balance queries the balance of a single coin for a single account.
	Balance(ctx context.Context, in *QueryBalanceRequest, opts ...grpc.CallOption) (*QueryBalanceResponse, error)
	// AllBalances queries the balance of all coins for a single account.
	//
	// When called from another module, this query might consume a high amount of
	// gas if the pagination field is incorrectly set.
	AllBalances(ctx context.Context, in *QueryAllBalancesRequest, opts ...grpc.CallOption) (*QueryAllBalancesResponse, error)
	// SpendableBalances queries the spendable balance of all coins for a single
	// account.
	//
	// When called from another module, this query might consume a high amount of
	// gas if the pagination field is incorrectly set.
	//
	// Since: cosmos-sdk 0.46
	SpendableBalances(ctx context.Context, in *QuerySpendableBalancesRequest, opts ...grpc.CallOption) (*QuerySpendableBalancesResponse, error)
	// SpendableBalanceByDenom queries the spendable balance of a single denom for
	// a single account.
	//
	// When called from another module, this query might consume a high amount of
	// gas if the pagination field is incorrectly set.
	//
	// Since: cosmos-sdk 0.47
	SpendableBalanceByDenom(ctx context.Context, in *QuerySpendableBalanceByDenomRequest, opts ...grpc.CallOption) (*QuerySpendableBalanceByDenomResponse, error)
	// TotalSupply queries the total supply of all coins.
	//
	// When called from another module, this query might consume a high amount of
	// gas if the pagination field is incorrectly set.
	TotalSupply(ctx context.Context, in *QueryTotalSupplyRequest, opts ...grpc.CallOption) (*QueryTotalSupplyResponse, error)
	// SupplyOf queries the supply of a single coin.
	//
	// When called from another module, this query might consume a high amount of
	// gas if the pagination field is incorrectly set.
	SupplyOf(ctx context.Context, in *QuerySupplyOfRequest, opts ...grpc.CallOption) (*QuerySupplyOfResponse, error)
	// Params queries the parameters of x/bank module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// DenomMetadata queries the client metadata of a given coin denomination.
	DenomMetadata(ctx context.Context, in *QueryDenomMetadataRequest, opts ...grpc.CallOption) (*QueryDenomMetadataResponse, error)
	// DenomMetadataByQueryString queries the client metadata of a given coin denomination.
	DenomMetadataByQueryString(ctx context.Context, in *QueryDenomMetadataByQueryStringRequest, opts ...grpc.CallOption) (*QueryDenomMetadataByQueryStringResponse, error)
	// DenomsMetadata queries the client metadata for all registered coin
	// denominations.
	DenomsMetadata(ctx context.Context, in *QueryDenomsMetadataRequest, opts ...grpc.CallOption) (*QueryDenomsMetadataResponse, error)
	// DenomOwners queries for all account addresses that own a particular token
	// denomination.
	//
	// When called from another module, this query might consume a high amount of
	// gas if the pagination field is incorrectly set.
	//
	// Since: cosmos-sdk 0.46
	DenomOwners(ctx context.Context, in *QueryDenomOwnersRequest, opts ...grpc.CallOption) (*QueryDenomOwnersResponse, error)
	// DenomOwnersByQuery queries for all account addresses that own a particular token
	// denomination.
	//
	// Since: cosmos-sdk 0.50.3
	DenomOwnersByQuery(ctx context.Context, in *QueryDenomOwnersByQueryRequest, opts ...grpc.CallOption) (*QueryDenomOwnersByQueryResponse, error)
	// SendEnabled queries for SendEnabled entries.
	//
	// This query only returns denominations that have specific SendEnabled settings.
	// Any denomination that does not have a specific setting will use the default
	// params.default_send_enabled, and will not be returned by this query.
	//
	// Since: cosmos-sdk 0.47
	SendEnabled(ctx context.Context, in *QuerySendEnabledRequest, opts ...grpc.CallOption) (*QuerySendEnabledResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Balance(ctx context.Context, in *QueryBalanceRequest, opts ...grpc.CallOption) (*QueryBalanceResponse, error) {
	out := new(QueryBalanceResponse)
	err := c.cc.Invoke(ctx, Query_Balance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) AllBalances(ctx context.Context, in *QueryAllBalancesRequest, opts ...grpc.CallOption) (*QueryAllBalancesResponse, error) {
	out := new(QueryAllBalancesResponse)
	err := c.cc.Invoke(ctx, Query_AllBalances_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) SpendableBalances(ctx context.Context, in *QuerySpendableBalancesRequest, opts ...grpc.CallOption) (*QuerySpendableBalancesResponse, error) {
	out := new(QuerySpendableBalancesResponse)
	err := c.cc.Invoke(ctx, Query_SpendableBalances_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) SpendableBalanceByDenom(ctx context.Context, in *QuerySpendableBalanceByDenomRequest, opts ...grpc.CallOption) (*QuerySpendableBalanceByDenomResponse, error) {
	out := new(QuerySpendableBalanceByDenomResponse)
	err := c.cc.Invoke(ctx, Query_SpendableBalanceByDenom_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) TotalSupply(ctx context.Context, in *QueryTotalSupplyRequest, opts ...grpc.CallOption) (*QueryTotalSupplyResponse, error) {
	out := new(QueryTotalSupplyResponse)
	err := c.cc.Invoke(ctx, Query_TotalSupply_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) SupplyOf(ctx context.Context, in *QuerySupplyOfRequest, opts ...grpc.CallOption) (*QuerySupplyOfResponse, error) {
	out := new(QuerySupplyOfResponse)
	err := c.cc.Invoke(ctx, Query_SupplyOf_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, Query_Params_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DenomMetadata(ctx context.Context, in *QueryDenomMetadataRequest, opts ...grpc.CallOption) (*QueryDenomMetadataResponse, error) {
	out := new(QueryDenomMetadataResponse)
	err := c.cc.Invoke(ctx, Query_DenomMetadata_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DenomMetadataByQueryString(ctx context.Context, in *QueryDenomMetadataByQueryStringRequest, opts ...grpc.CallOption) (*QueryDenomMetadataByQueryStringResponse, error) {
	out := new(QueryDenomMetadataByQueryStringResponse)
	err := c.cc.Invoke(ctx, Query_DenomMetadataByQueryString_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DenomsMetadata(ctx context.Context, in *QueryDenomsMetadataRequest, opts ...grpc.CallOption) (*QueryDenomsMetadataResponse, error) {
	out := new(QueryDenomsMetadataResponse)
	err := c.cc.Invoke(ctx, Query_DenomsMetadata_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DenomOwners(ctx context.Context, in *QueryDenomOwnersRequest, opts ...grpc.CallOption) (*QueryDenomOwnersResponse, error) {
	out := new(QueryDenomOwnersResponse)
	err := c.cc.Invoke(ctx, Query_DenomOwners_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DenomOwnersByQuery(ctx context.Context, in *QueryDenomOwnersByQueryRequest, opts ...grpc.CallOption) (*QueryDenomOwnersByQueryResponse, error) {
	out := new(QueryDenomOwnersByQueryResponse)
	err := c.cc.Invoke(ctx, Query_DenomOwnersByQuery_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) SendEnabled(ctx context.Context, in *QuerySendEnabledRequest, opts ...grpc.CallOption) (*QuerySendEnabledResponse, error) {
	out := new(QuerySendEnabledResponse)
	err := c.cc.Invoke(ctx, Query_SendEnabled_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Balance queries the balance of a single coin for a single account.
	Balance(context.Context, *QueryBalanceRequest) (*QueryBalanceResponse, error)
	// AllBalances queries the balance of all coins for a single account.
	//
	// When called from another module, this query might consume a high amount of
	// gas if the pagination field is incorrectly set.
	AllBalances(context.Context, *QueryAllBalancesRequest) (*QueryAllBalancesResponse, error)
	// SpendableBalances queries the spendable balance of all coins for a single
	// account.
	//
	// When called from another module, this query might consume a high amount of
	// gas if the pagination field is incorrectly set.
	//
	// Since: cosmos-sdk 0.46
	SpendableBalances(context.Context, *QuerySpendableBalancesRequest) (*QuerySpendableBalancesResponse, error)
	// SpendableBalanceByDenom queries the spendable balance of a single denom for
	// a single account.
	//
	// When called from another module, this query might consume a high amount of
	// gas if the pagination field is incorrectly set.
	//
	// Since: cosmos-sdk 0.47
	SpendableBalanceByDenom(context.Context, *QuerySpendableBalanceByDenomRequest) (*QuerySpendableBalanceByDenomResponse, error)
	// TotalSupply queries the total supply of all coins.
	//
	// When called from another module, this query might consume a high amount of
	// gas if the pagination field is incorrectly set.
	TotalSupply(context.Context, *QueryTotalSupplyRequest) (*QueryTotalSupplyResponse, error)
	// SupplyOf queries the supply of a single coin.
	//
	// When called from another module, this query might consume a high amount of
	// gas if the pagination field is incorrectly set.
	SupplyOf(context.Context, *QuerySupplyOfRequest) (*QuerySupplyOfResponse, error)
	// Params queries the parameters of x/bank module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// DenomMetadata queries the client metadata of a given coin denomination.
	DenomMetadata(context.Context, *QueryDenomMetadataRequest) (*QueryDenomMetadataResponse, error)
	// DenomMetadataByQueryString queries the client metadata of a given coin denomination.
	DenomMetadataByQueryString(context.Context, *QueryDenomMetadataByQueryStringRequest) (*QueryDenomMetadataByQueryStringResponse, error)
	// DenomsMetadata queries the client metadata for all registered coin
	// denominations.
	DenomsMetadata(context.Context, *QueryDenomsMetadataRequest) (*QueryDenomsMetadataResponse, error)
	// DenomOwners queries for all account addresses that own a particular token
	// denomination.
	//
	// When called from another module, this query might consume a high amount of
	// gas if the pagination field is incorrectly set.
	//
	// Since: cosmos-sdk 0.46
	DenomOwners(context.Context, *QueryDenomOwnersRequest) (*QueryDenomOwnersResponse, error)
	// DenomOwnersByQuery queries for all account addresses that own a particular token
	// denomination.
	//
	// Since: cosmos-sdk 0.50.3
	DenomOwnersByQuery(context.Context, *QueryDenomOwnersByQueryRequest) (*QueryDenomOwnersByQueryResponse, error)
	// SendEnabled queries for SendEnabled entries.
	//
	// This query only returns denominations that have specific SendEnabled settings.
	// Any denomination that does not have a specific setting will use the default
	// params.default_send_enabled, and will not be returned by this query.
	//
	// Since: cosmos-sdk 0.47
	SendEnabled(context.Context, *QuerySendEnabledRequest) (*QuerySendEnabledResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Balance(context.Context, *QueryBalanceRequest) (*QueryBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Balance not implemented")
}
func (UnimplementedQueryServer) AllBalances(context.Context, *QueryAllBalancesRequest) (*QueryAllBalancesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllBalances not implemented")
}
func (UnimplementedQueryServer) SpendableBalances(context.Context, *QuerySpendableBalancesRequest) (*QuerySpendableBalancesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SpendableBalances not implemented")
}
func (UnimplementedQueryServer) SpendableBalanceByDenom(context.Context, *QuerySpendableBalanceByDenomRequest) (*QuerySpendableBalanceByDenomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SpendableBalanceByDenom not implemented")
}
func (UnimplementedQueryServer) TotalSupply(context.Context, *QueryTotalSupplyRequest) (*QueryTotalSupplyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TotalSupply not implemented")
}
func (UnimplementedQueryServer) SupplyOf(context.Context, *QuerySupplyOfRequest) (*QuerySupplyOfResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SupplyOf not implemented")
}
func (UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) DenomMetadata(context.Context, *QueryDenomMetadataRequest) (*QueryDenomMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DenomMetadata not implemented")
}
func (UnimplementedQueryServer) DenomMetadataByQueryString(context.Context, *QueryDenomMetadataByQueryStringRequest) (*QueryDenomMetadataByQueryStringResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DenomMetadataByQueryString not implemented")
}
func (UnimplementedQueryServer) DenomsMetadata(context.Context, *QueryDenomsMetadataRequest) (*QueryDenomsMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DenomsMetadata not implemented")
}
func (UnimplementedQueryServer) DenomOwners(context.Context, *QueryDenomOwnersRequest) (*QueryDenomOwnersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DenomOwners not implemented")
}
func (UnimplementedQueryServer) DenomOwnersByQuery(context.Context, *QueryDenomOwnersByQueryRequest) (*QueryDenomOwnersByQueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DenomOwnersByQuery not implemented")
}
func (UnimplementedQueryServer) SendEnabled(context.Context, *QuerySendEnabledRequest) (*QuerySendEnabledResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendEnabled not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Balance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Balance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Balance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Balance(ctx, req.(*QueryBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_AllBalances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllBalancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AllBalances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_AllBalances_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AllBalances(ctx, req.(*QueryAllBalancesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_SpendableBalances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySpendableBalancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).SpendableBalances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_SpendableBalances_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).SpendableBalances(ctx, req.(*QuerySpendableBalancesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_SpendableBalanceByDenom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySpendableBalanceByDenomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).SpendableBalanceByDenom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_SpendableBalanceByDenom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).SpendableBalanceByDenom(ctx, req.(*QuerySpendableBalanceByDenomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_TotalSupply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTotalSupplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TotalSupply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_TotalSupply_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).TotalSupply(ctx, req.(*QueryTotalSupplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_SupplyOf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySupplyOfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).SupplyOf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_SupplyOf_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).SupplyOf(ctx, req.(*QuerySupplyOfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Params_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DenomMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDenomMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DenomMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_DenomMetadata_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DenomMetadata(ctx, req.(*QueryDenomMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DenomMetadataByQueryString_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDenomMetadataByQueryStringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DenomMetadataByQueryString(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_DenomMetadataByQueryString_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DenomMetadataByQueryString(ctx, req.(*QueryDenomMetadataByQueryStringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DenomsMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDenomsMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DenomsMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_DenomsMetadata_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DenomsMetadata(ctx, req.(*QueryDenomsMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DenomOwners_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDenomOwnersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DenomOwners(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_DenomOwners_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DenomOwners(ctx, req.(*QueryDenomOwnersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DenomOwnersByQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDenomOwnersByQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DenomOwnersByQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_DenomOwnersByQuery_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DenomOwnersByQuery(ctx, req.(*QueryDenomOwnersByQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_SendEnabled_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySendEnabledRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).SendEnabled(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_SendEnabled_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).SendEnabled(ctx, req.(*QuerySendEnabledRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.bank.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Balance",
			Handler:    _Query_Balance_Handler,
		},
		{
			MethodName: "AllBalances",
			Handler:    _Query_AllBalances_Handler,
		},
		{
			MethodName: "SpendableBalances",
			Handler:    _Query_SpendableBalances_Handler,
		},
		{
			MethodName: "SpendableBalanceByDenom",
			Handler:    _Query_SpendableBalanceByDenom_Handler,
		},
		{
			MethodName: "TotalSupply",
			Handler:    _Query_TotalSupply_Handler,
		},
		{
			MethodName: "SupplyOf",
			Handler:    _Query_SupplyOf_Handler,
		},
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "DenomMetadata",
			Handler:    _Query_DenomMetadata_Handler,
		},
		{
			MethodName: "DenomMetadataByQueryString",
			Handler:    _Query_DenomMetadataByQueryString_Handler,
		},
		{
			MethodName: "DenomsMetadata",
			Handler:    _Query_DenomsMetadata_Handler,
		},
		{
			MethodName: "DenomOwners",
			Handler:    _Query_DenomOwners_Handler,
		},
		{
			MethodName: "DenomOwnersByQuery",
			Handler:    _Query_DenomOwnersByQuery_Handler,
		},
		{
			MethodName: "SendEnabled",
			Handler:    _Query_SendEnabled_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/bank/v1beta1/query.proto",
}
