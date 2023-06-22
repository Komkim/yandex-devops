package mygrpc

import (
	"context"
	"github.com/golang/protobuf/ptypes/wrappers"
	pb "yandex-devops/proto"
	"yandex-devops/storage"
)

func (r *Router) SaveOrUpdate(ctx context.Context, request *pb.SaveOrUpdateRequest) (*pb.SaveOrUpdateResponse, error) {
	gm := storage.Metrics{
		ID:    request.Id,
		MType: request.Type,
		Hash:  request.Hash,
	}
	if request.Value != nil {
		gm.Value = &request.Value.Value
	}
	if request.Delta != nil {
		gm.Delta = &request.Delta.Value
	}

	if checkHas, err := r.services.StorageService.CheckHash(gm, r.cfg.Key); err != nil || !checkHas {
		if err != nil {
			return nil, err
		}
	}

	result, err := r.services.StorageService.SaveOrUpdateOne(gm, r.cfg.Key)
	if err != nil {
		return nil, err
	}

	rsp := &pb.SaveOrUpdateResponse{
		Id:   result.ID,
		Type: result.MType,
		Hash: result.Hash,
	}
	if result.Value != nil {
		rsp.Value = &wrappers.DoubleValue{Value: *result.Value}
	}
	if result.Delta != nil {
		rsp.Delta = &wrappers.Int64Value{Value: *result.Delta}
	}

	return rsp, nil
}
