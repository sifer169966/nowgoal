package handlers

// import (
// 	"context"
// 	"database/sql"
// 	v1 "nowgoal/internal/api/v1"
// 	"nowgoal/internal/core/domain"
// 	"nowgoal/internal/core/ports"

// 	ctxlogrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
// 	"github.com/sirupsen/logrus"
// 	"google.golang.org/grpc/codes"
// 	"google.golang.org/grpc/status"
// )

// const (
// 	apiVersion = "v1"
// )

// type GRPCHandler struct {
// 	db      *sql.DB
// 	toteSrv ports.Service
// }

// func NewGRPCHandler(db *sql.DB, toteSrv ports.Service) *GRPCHandler {
// 	return &GRPCHandler{
// 		db:      db,
// 		toteSrv: toteSrv,
// 	}
// }

// func (hdl *GRPCHandler) GetToteBags(ctx context.Context, in *v1.GetToteBagsRequest) (*v1.GetToteBagsResponse, error) {
// 	err := checkAPI(in.Api)
// 	if err != nil {
// 		return nil, err
// 	}
// 	totesInfomation, err := hdl.toteSrv.GetToteBags(domain.GetToteBagsRequest{
// 		FindBy: domain.GetToteBagsFindBy{
// 			ToteID:     in.FindBy.ToteID,
// 			CustomerID: in.FindBy.CustomerID,
// 		},
// 		FilterBy: domain.GetToteBagsFilterBy{
// 			ToteProductStatus: in.FilterBy.ToteProducStatus,
// 		},
// 		OrderBy: in.OrderBy,
// 		SortBy: domain.GetToteBagsSortBy{
// 			Asc:  in.SortBy.Asc,
// 			Desc: in.SortBy.Desc,
// 		},
// 		Limit: in.Limit,
// 		Page:  in.Page,
// 	})
// 	if err != nil {
// 		grpcLogger(ctx, "error", in.IpClient, err)
// 		return nil, err
// 	}
// 	// Mapping data ...
// 	totesCount := totesInfomation.TotesCount
// 	totesDetail := []*v1.ToteDetail{}
// 	for _, v := range totesInfomation.TotesInformation {
// 		products := []*v1.Product{}
// 		for _, productVal := range v.Products {
// 			products = append(products, &v1.Product{
// 				ToteProductID:     productVal.ToteProductID,
// 				ToteID:            productVal.ToteID,
// 				ProductID:         productVal.ProductID,
// 				SkuID:             productVal.SkuID,
// 				Quantity:          productVal.Quantity,
// 				PricePerUnit:      productVal.PricePerUnit,
// 				ToteProductStatus: productVal.ToteProductStatus,
// 				CreatedAt:         int32(productVal.CreatedAt.Unix()),
// 				UpdatedAt:         int32(productVal.UpdatedAt.Unix()),
// 			})
// 		}
// 		totesDetail = append(totesDetail, &v1.ToteDetail{
// 			Tote: &v1.Tote{
// 				ToteID:     v.Tote.ToteID,
// 				CustomerID: v.Tote.CustomerID,
// 				CreatedAt:  int32(v.Tote.CreatedAt.Unix()),
// 				UpdatedAt:  int32(v.Tote.UpdatedAt.Unix()),
// 			},
// 			AllProductsCount: v.ProductsCount,
// 			Products:         products,
// 		})
// 	}
// 	grpcLogger(ctx, "info", in.IpClient, "success")
// 	return &v1.GetToteBagsResponse{
// 		Api:           apiVersion,
// 		TotesDetail:   totesDetail,
// 		AllTotesCount: totesCount,
// 	}, nil
// }
// func (hdl *GRPCHandler) AddProducts(ctx context.Context, in *v1.AddProductsRequest) (*v1.AddProductsResponse, error) {
// 	err := checkAPI(in.Api)
// 	if err != nil {
// 		return nil, err
// 	}
// 	products := []domain.AddProductDetail{}
// 	for _, v := range in.Products {
// 		products = append(products, domain.AddProductDetail{
// 			ToteID:       v.ToteID,
// 			ProductID:    v.ProductID,
// 			SkuID:        v.SkuID,
// 			Quantity:     v.Quantity,
// 			PricePerUnit: v.PricePerUnit,
// 		})
// 	}
// 	ids, err := hdl.toteSrv.AddProducts(domain.AddProductsRequest{Products: products})
// 	if err != nil {
// 		grpcLogger(ctx, "error", in.IpClient, err)
// 		return nil, err
// 	}
// 	grpcLogger(ctx, "info", in.IpClient, "success")
// 	return &v1.AddProductsResponse{
// 		Api:        apiVersion,
// 		ProductsID: ids,
// 	}, nil
// }

// func (hdl *GRPCHandler) CreateNewTote(ctx context.Context, in *v1.CreateNewToteRequest) (*v1.CreateNewToteReponse, error) {
// 	err := checkAPI(in.Api)
// 	if err != nil {
// 		return nil, err
// 	}
// 	toteID, err := hdl.toteSrv.CreateNewTote(in.CustomerID)
// 	if err != nil {
// 		grpcLogger(ctx, "error", in.IpClient, err)
// 		return nil, err
// 	}
// 	grpcLogger(ctx, "info", in.IpClient, "success")
// 	return &v1.CreateNewToteReponse{
// 		Api:    apiVersion,
// 		ToteID: toteID,
// 	}, nil
// }

// func (hdl *GRPCHandler) DeleteProducts(ctx context.Context, in *v1.DeleteProductsRequest) (*v1.DeleteProductsResponse, error) {
// 	ids, err := hdl.toteSrv.DeletProducts(domain.DeleteProductsRequest{
// 		ToteID:     in.ToteID,
// 		ProductsID: in.ProductsID,
// 	})
// 	if err != nil {
// 		grpcLogger(ctx, "error", in.IpClient, err)
// 		return nil, err
// 	}
// 	return &v1.DeleteProductsResponse{
// 		Api:        apiVersion,
// 		ProductsID: ids,
// 	}, nil
// }
// func (hdl *GRPCHandler) UpdateToteProducts(ctx context.Context, in *v1.UpdateToteProductsRequest) (*v1.UpdateToteProductsResponse, error) {
// 	updateProducts := []domain.UpdateProduct{}
// 	for _, v := range in.UpdateToteProducts {
// 		updateProducts = append(updateProducts, domain.UpdateProduct{
// 			ToteProductID:     v.ToteProductID,
// 			ToteID:            v.ToteID,
// 			Quantity:          v.Quantity,
// 			ToteProductStatus: v.ToteProductStatus,
// 		})
// 	}
// 	ids, err := hdl.toteSrv.UpdateProducts(updateProducts)
// 	if err != nil {
// 		grpcLogger(ctx, "error", in.IpClient, err)
// 		return nil, err
// 	}
// 	grpcLogger(ctx, "info", in.IpClient, "success")
// 	return &v1.UpdateToteProductsResponse{
// 		Api:            apiVersion,
// 		ToteProductIDs: ids,
// 	}, nil
// }
// func checkAPI(api string) error {
// 	// api empty means use current api version ...
// 	if len(api) > 0 && apiVersion != api {
// 		return status.Errorf(
// 			codes.Unimplemented,
// 			"Unsupported API version: service implement API version '%s', but asked for '%s'",
// 			apiVersion,
// 			api,
// 		)
// 	}
// 	return nil
// }
// func (hdl *GRPCHandler) HealthCheckToteService(ctx context.Context, _ *v1.HealthCheckToteRequest) (*v1.HealthCheckToteResponse, error) {
// 	err := hdl.db.Ping()
// 	if err != nil {
// 		return &v1.HealthCheckToteResponse{
// 			Status: v1.HealthCheckToteResponse_NOT_SERVING,
// 		}, nil
// 	}
// 	return &v1.HealthCheckToteResponse{
// 		Status: v1.HealthCheckToteResponse_SERVING,
// 	}, nil
// }

// func grpcLogger(ctx context.Context, logType, ip string, data interface{}) {
// 	logrusFields := logrus.Fields{
// 		"Service":   "TOTE-SERVICE",
// 		"IP Client": ip,
// 		"Message":   data,
// 	}
// 	switch logType {
// 	case "info":
// 		ctxlogrus.Extract(ctx).WithFields(logrusFields).Info()
// 	case "error":
// 		ctxlogrus.Extract(ctx).WithFields(logrusFields)
// 	}
// }
