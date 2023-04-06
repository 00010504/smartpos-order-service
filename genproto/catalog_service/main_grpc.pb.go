// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: main.proto

package catalog_service

import (
	context "context"
	common "genproto/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CatalogServiceClient is the client API for CatalogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CatalogServiceClient interface {
	// measurementUnit
	CreateMeasurementUnit(ctx context.Context, in *CreateMeasurementUnitRequest, opts ...grpc.CallOption) (*common.ResponseID, error)
	GetMeasurementUnitByID(ctx context.Context, in *common.RequestID, opts ...grpc.CallOption) (*MeasurementUnit, error)
	UpdateMeasurementUnit(ctx context.Context, in *UpdateMeasurementUnitRequest, opts ...grpc.CallOption) (*common.ResponseID, error)
	GetAllMeasurementUnits(ctx context.Context, in *GetAllMeasurementUnitsRequest, opts ...grpc.CallOption) (*GetAllMeasurementUnitsResponse, error)
	DeleteMeasurementUnitById(ctx context.Context, in *common.RequestID, opts ...grpc.CallOption) (*common.ResponseID, error)
	GetAllDefaultUnits(ctx context.Context, in *common.SearchRequest, opts ...grpc.CallOption) (*GetAllDefaultUnitsResponse, error)
	// product
	CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*common.ResponseID, error)
	GetProductByID(ctx context.Context, in *common.RequestID, opts ...grpc.CallOption) (*Product, error)
	UpdateProduct(ctx context.Context, in *UpdateProductRequest, opts ...grpc.CallOption) (*common.ResponseID, error)
	GetAllProducts(ctx context.Context, in *GetAllProductsRequest, opts ...grpc.CallOption) (*GetAllProductsResponse, error)
	DeleteProductById(ctx context.Context, in *common.RequestID, opts ...grpc.CallOption) (*common.ResponseID, error)
	DeleteProductsByIds(ctx context.Context, in *common.RequestIDs, opts ...grpc.CallOption) (*common.Empty, error)
	SearchProducts(ctx context.Context, in *GetAllProductsRequest, opts ...grpc.CallOption) (*SearchProductsResponse, error)
	// category
	CreateCategory(ctx context.Context, in *CreateCategoryRequest, opts ...grpc.CallOption) (*common.ResponseID, error)
	GetCategoryByID(ctx context.Context, in *common.RequestID, opts ...grpc.CallOption) (*GetCategoryByIDResponse, error)
	UpdateCategory(ctx context.Context, in *UpdateCategoryRequest, opts ...grpc.CallOption) (*common.ResponseID, error)
	GetAllCategories(ctx context.Context, in *GetAllCategoriesRequest, opts ...grpc.CallOption) (*GetAllCategoriesResponse, error)
	DeleteCategoryById(ctx context.Context, in *common.RequestID, opts ...grpc.CallOption) (*common.ResponseID, error)
}

type catalogServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCatalogServiceClient(cc grpc.ClientConnInterface) CatalogServiceClient {
	return &catalogServiceClient{cc}
}

func (c *catalogServiceClient) CreateMeasurementUnit(ctx context.Context, in *CreateMeasurementUnitRequest, opts ...grpc.CallOption) (*common.ResponseID, error) {
	out := new(common.ResponseID)
	err := c.cc.Invoke(ctx, "/CatalogService/CreateMeasurementUnit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) GetMeasurementUnitByID(ctx context.Context, in *common.RequestID, opts ...grpc.CallOption) (*MeasurementUnit, error) {
	out := new(MeasurementUnit)
	err := c.cc.Invoke(ctx, "/CatalogService/GetMeasurementUnitByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) UpdateMeasurementUnit(ctx context.Context, in *UpdateMeasurementUnitRequest, opts ...grpc.CallOption) (*common.ResponseID, error) {
	out := new(common.ResponseID)
	err := c.cc.Invoke(ctx, "/CatalogService/UpdateMeasurementUnit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) GetAllMeasurementUnits(ctx context.Context, in *GetAllMeasurementUnitsRequest, opts ...grpc.CallOption) (*GetAllMeasurementUnitsResponse, error) {
	out := new(GetAllMeasurementUnitsResponse)
	err := c.cc.Invoke(ctx, "/CatalogService/GetAllMeasurementUnits", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) DeleteMeasurementUnitById(ctx context.Context, in *common.RequestID, opts ...grpc.CallOption) (*common.ResponseID, error) {
	out := new(common.ResponseID)
	err := c.cc.Invoke(ctx, "/CatalogService/DeleteMeasurementUnitById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) GetAllDefaultUnits(ctx context.Context, in *common.SearchRequest, opts ...grpc.CallOption) (*GetAllDefaultUnitsResponse, error) {
	out := new(GetAllDefaultUnitsResponse)
	err := c.cc.Invoke(ctx, "/CatalogService/GetAllDefaultUnits", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*common.ResponseID, error) {
	out := new(common.ResponseID)
	err := c.cc.Invoke(ctx, "/CatalogService/CreateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) GetProductByID(ctx context.Context, in *common.RequestID, opts ...grpc.CallOption) (*Product, error) {
	out := new(Product)
	err := c.cc.Invoke(ctx, "/CatalogService/GetProductByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) UpdateProduct(ctx context.Context, in *UpdateProductRequest, opts ...grpc.CallOption) (*common.ResponseID, error) {
	out := new(common.ResponseID)
	err := c.cc.Invoke(ctx, "/CatalogService/UpdateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) GetAllProducts(ctx context.Context, in *GetAllProductsRequest, opts ...grpc.CallOption) (*GetAllProductsResponse, error) {
	out := new(GetAllProductsResponse)
	err := c.cc.Invoke(ctx, "/CatalogService/GetAllProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) DeleteProductById(ctx context.Context, in *common.RequestID, opts ...grpc.CallOption) (*common.ResponseID, error) {
	out := new(common.ResponseID)
	err := c.cc.Invoke(ctx, "/CatalogService/DeleteProductById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) DeleteProductsByIds(ctx context.Context, in *common.RequestIDs, opts ...grpc.CallOption) (*common.Empty, error) {
	out := new(common.Empty)
	err := c.cc.Invoke(ctx, "/CatalogService/DeleteProductsByIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) SearchProducts(ctx context.Context, in *GetAllProductsRequest, opts ...grpc.CallOption) (*SearchProductsResponse, error) {
	out := new(SearchProductsResponse)
	err := c.cc.Invoke(ctx, "/CatalogService/SearchProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) CreateCategory(ctx context.Context, in *CreateCategoryRequest, opts ...grpc.CallOption) (*common.ResponseID, error) {
	out := new(common.ResponseID)
	err := c.cc.Invoke(ctx, "/CatalogService/CreateCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) GetCategoryByID(ctx context.Context, in *common.RequestID, opts ...grpc.CallOption) (*GetCategoryByIDResponse, error) {
	out := new(GetCategoryByIDResponse)
	err := c.cc.Invoke(ctx, "/CatalogService/GetCategoryByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) UpdateCategory(ctx context.Context, in *UpdateCategoryRequest, opts ...grpc.CallOption) (*common.ResponseID, error) {
	out := new(common.ResponseID)
	err := c.cc.Invoke(ctx, "/CatalogService/UpdateCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) GetAllCategories(ctx context.Context, in *GetAllCategoriesRequest, opts ...grpc.CallOption) (*GetAllCategoriesResponse, error) {
	out := new(GetAllCategoriesResponse)
	err := c.cc.Invoke(ctx, "/CatalogService/GetAllCategories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) DeleteCategoryById(ctx context.Context, in *common.RequestID, opts ...grpc.CallOption) (*common.ResponseID, error) {
	out := new(common.ResponseID)
	err := c.cc.Invoke(ctx, "/CatalogService/DeleteCategoryById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CatalogServiceServer is the server API for CatalogService service.
// All implementations should embed UnimplementedCatalogServiceServer
// for forward compatibility
type CatalogServiceServer interface {
	// measurementUnit
	CreateMeasurementUnit(context.Context, *CreateMeasurementUnitRequest) (*common.ResponseID, error)
	GetMeasurementUnitByID(context.Context, *common.RequestID) (*MeasurementUnit, error)
	UpdateMeasurementUnit(context.Context, *UpdateMeasurementUnitRequest) (*common.ResponseID, error)
	GetAllMeasurementUnits(context.Context, *GetAllMeasurementUnitsRequest) (*GetAllMeasurementUnitsResponse, error)
	DeleteMeasurementUnitById(context.Context, *common.RequestID) (*common.ResponseID, error)
	GetAllDefaultUnits(context.Context, *common.SearchRequest) (*GetAllDefaultUnitsResponse, error)
	// product
	CreateProduct(context.Context, *CreateProductRequest) (*common.ResponseID, error)
	GetProductByID(context.Context, *common.RequestID) (*Product, error)
	UpdateProduct(context.Context, *UpdateProductRequest) (*common.ResponseID, error)
	GetAllProducts(context.Context, *GetAllProductsRequest) (*GetAllProductsResponse, error)
	DeleteProductById(context.Context, *common.RequestID) (*common.ResponseID, error)
	DeleteProductsByIds(context.Context, *common.RequestIDs) (*common.Empty, error)
	SearchProducts(context.Context, *GetAllProductsRequest) (*SearchProductsResponse, error)
	// category
	CreateCategory(context.Context, *CreateCategoryRequest) (*common.ResponseID, error)
	GetCategoryByID(context.Context, *common.RequestID) (*GetCategoryByIDResponse, error)
	UpdateCategory(context.Context, *UpdateCategoryRequest) (*common.ResponseID, error)
	GetAllCategories(context.Context, *GetAllCategoriesRequest) (*GetAllCategoriesResponse, error)
	DeleteCategoryById(context.Context, *common.RequestID) (*common.ResponseID, error)
}

// UnimplementedCatalogServiceServer should be embedded to have forward compatible implementations.
type UnimplementedCatalogServiceServer struct {
}

func (UnimplementedCatalogServiceServer) CreateMeasurementUnit(context.Context, *CreateMeasurementUnitRequest) (*common.ResponseID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMeasurementUnit not implemented")
}
func (UnimplementedCatalogServiceServer) GetMeasurementUnitByID(context.Context, *common.RequestID) (*MeasurementUnit, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMeasurementUnitByID not implemented")
}
func (UnimplementedCatalogServiceServer) UpdateMeasurementUnit(context.Context, *UpdateMeasurementUnitRequest) (*common.ResponseID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMeasurementUnit not implemented")
}
func (UnimplementedCatalogServiceServer) GetAllMeasurementUnits(context.Context, *GetAllMeasurementUnitsRequest) (*GetAllMeasurementUnitsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllMeasurementUnits not implemented")
}
func (UnimplementedCatalogServiceServer) DeleteMeasurementUnitById(context.Context, *common.RequestID) (*common.ResponseID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMeasurementUnitById not implemented")
}
func (UnimplementedCatalogServiceServer) GetAllDefaultUnits(context.Context, *common.SearchRequest) (*GetAllDefaultUnitsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllDefaultUnits not implemented")
}
func (UnimplementedCatalogServiceServer) CreateProduct(context.Context, *CreateProductRequest) (*common.ResponseID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProduct not implemented")
}
func (UnimplementedCatalogServiceServer) GetProductByID(context.Context, *common.RequestID) (*Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductByID not implemented")
}
func (UnimplementedCatalogServiceServer) UpdateProduct(context.Context, *UpdateProductRequest) (*common.ResponseID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProduct not implemented")
}
func (UnimplementedCatalogServiceServer) GetAllProducts(context.Context, *GetAllProductsRequest) (*GetAllProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllProducts not implemented")
}
func (UnimplementedCatalogServiceServer) DeleteProductById(context.Context, *common.RequestID) (*common.ResponseID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProductById not implemented")
}
func (UnimplementedCatalogServiceServer) DeleteProductsByIds(context.Context, *common.RequestIDs) (*common.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProductsByIds not implemented")
}
func (UnimplementedCatalogServiceServer) SearchProducts(context.Context, *GetAllProductsRequest) (*SearchProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchProducts not implemented")
}
func (UnimplementedCatalogServiceServer) CreateCategory(context.Context, *CreateCategoryRequest) (*common.ResponseID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCategory not implemented")
}
func (UnimplementedCatalogServiceServer) GetCategoryByID(context.Context, *common.RequestID) (*GetCategoryByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategoryByID not implemented")
}
func (UnimplementedCatalogServiceServer) UpdateCategory(context.Context, *UpdateCategoryRequest) (*common.ResponseID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCategory not implemented")
}
func (UnimplementedCatalogServiceServer) GetAllCategories(context.Context, *GetAllCategoriesRequest) (*GetAllCategoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCategories not implemented")
}
func (UnimplementedCatalogServiceServer) DeleteCategoryById(context.Context, *common.RequestID) (*common.ResponseID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCategoryById not implemented")
}

// UnsafeCatalogServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CatalogServiceServer will
// result in compilation errors.
type UnsafeCatalogServiceServer interface {
	mustEmbedUnimplementedCatalogServiceServer()
}

func RegisterCatalogServiceServer(s grpc.ServiceRegistrar, srv CatalogServiceServer) {
	s.RegisterService(&CatalogService_ServiceDesc, srv)
}

func _CatalogService_CreateMeasurementUnit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMeasurementUnitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).CreateMeasurementUnit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CatalogService/CreateMeasurementUnit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).CreateMeasurementUnit(ctx, req.(*CreateMeasurementUnitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_GetMeasurementUnitByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.RequestID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).GetMeasurementUnitByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CatalogService/GetMeasurementUnitByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).GetMeasurementUnitByID(ctx, req.(*common.RequestID))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_UpdateMeasurementUnit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMeasurementUnitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).UpdateMeasurementUnit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CatalogService/UpdateMeasurementUnit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).UpdateMeasurementUnit(ctx, req.(*UpdateMeasurementUnitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_GetAllMeasurementUnits_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllMeasurementUnitsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).GetAllMeasurementUnits(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CatalogService/GetAllMeasurementUnits",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).GetAllMeasurementUnits(ctx, req.(*GetAllMeasurementUnitsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_DeleteMeasurementUnitById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.RequestID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).DeleteMeasurementUnitById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CatalogService/DeleteMeasurementUnitById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).DeleteMeasurementUnitById(ctx, req.(*common.RequestID))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_GetAllDefaultUnits_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).GetAllDefaultUnits(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CatalogService/GetAllDefaultUnits",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).GetAllDefaultUnits(ctx, req.(*common.SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_CreateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).CreateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CatalogService/CreateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).CreateProduct(ctx, req.(*CreateProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_GetProductByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.RequestID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).GetProductByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CatalogService/GetProductByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).GetProductByID(ctx, req.(*common.RequestID))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_UpdateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).UpdateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CatalogService/UpdateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).UpdateProduct(ctx, req.(*UpdateProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_GetAllProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).GetAllProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CatalogService/GetAllProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).GetAllProducts(ctx, req.(*GetAllProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_DeleteProductById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.RequestID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).DeleteProductById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CatalogService/DeleteProductById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).DeleteProductById(ctx, req.(*common.RequestID))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_DeleteProductsByIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.RequestIDs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).DeleteProductsByIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CatalogService/DeleteProductsByIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).DeleteProductsByIds(ctx, req.(*common.RequestIDs))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_SearchProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).SearchProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CatalogService/SearchProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).SearchProducts(ctx, req.(*GetAllProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_CreateCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).CreateCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CatalogService/CreateCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).CreateCategory(ctx, req.(*CreateCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_GetCategoryByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.RequestID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).GetCategoryByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CatalogService/GetCategoryByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).GetCategoryByID(ctx, req.(*common.RequestID))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_UpdateCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).UpdateCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CatalogService/UpdateCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).UpdateCategory(ctx, req.(*UpdateCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_GetAllCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllCategoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).GetAllCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CatalogService/GetAllCategories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).GetAllCategories(ctx, req.(*GetAllCategoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_DeleteCategoryById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.RequestID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).DeleteCategoryById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CatalogService/DeleteCategoryById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).DeleteCategoryById(ctx, req.(*common.RequestID))
	}
	return interceptor(ctx, in, info, handler)
}

// CatalogService_ServiceDesc is the grpc.ServiceDesc for CatalogService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CatalogService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "CatalogService",
	HandlerType: (*CatalogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMeasurementUnit",
			Handler:    _CatalogService_CreateMeasurementUnit_Handler,
		},
		{
			MethodName: "GetMeasurementUnitByID",
			Handler:    _CatalogService_GetMeasurementUnitByID_Handler,
		},
		{
			MethodName: "UpdateMeasurementUnit",
			Handler:    _CatalogService_UpdateMeasurementUnit_Handler,
		},
		{
			MethodName: "GetAllMeasurementUnits",
			Handler:    _CatalogService_GetAllMeasurementUnits_Handler,
		},
		{
			MethodName: "DeleteMeasurementUnitById",
			Handler:    _CatalogService_DeleteMeasurementUnitById_Handler,
		},
		{
			MethodName: "GetAllDefaultUnits",
			Handler:    _CatalogService_GetAllDefaultUnits_Handler,
		},
		{
			MethodName: "CreateProduct",
			Handler:    _CatalogService_CreateProduct_Handler,
		},
		{
			MethodName: "GetProductByID",
			Handler:    _CatalogService_GetProductByID_Handler,
		},
		{
			MethodName: "UpdateProduct",
			Handler:    _CatalogService_UpdateProduct_Handler,
		},
		{
			MethodName: "GetAllProducts",
			Handler:    _CatalogService_GetAllProducts_Handler,
		},
		{
			MethodName: "DeleteProductById",
			Handler:    _CatalogService_DeleteProductById_Handler,
		},
		{
			MethodName: "DeleteProductsByIds",
			Handler:    _CatalogService_DeleteProductsByIds_Handler,
		},
		{
			MethodName: "SearchProducts",
			Handler:    _CatalogService_SearchProducts_Handler,
		},
		{
			MethodName: "CreateCategory",
			Handler:    _CatalogService_CreateCategory_Handler,
		},
		{
			MethodName: "GetCategoryByID",
			Handler:    _CatalogService_GetCategoryByID_Handler,
		},
		{
			MethodName: "UpdateCategory",
			Handler:    _CatalogService_UpdateCategory_Handler,
		},
		{
			MethodName: "GetAllCategories",
			Handler:    _CatalogService_GetAllCategories_Handler,
		},
		{
			MethodName: "DeleteCategoryById",
			Handler:    _CatalogService_DeleteCategoryById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "main.proto",
}