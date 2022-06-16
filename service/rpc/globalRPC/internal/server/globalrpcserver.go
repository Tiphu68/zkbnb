// Code generated by goctl. DO NOT EDIT!
// Source: globalRPC.proto

package server

import (
	"context"

	"github.com/zecrey-labs/zecrey-legend/service/rpc/globalRPC/globalRPCProto"
	"github.com/zecrey-labs/zecrey-legend/service/rpc/globalRPC/internal/logic"
	"github.com/zecrey-labs/zecrey-legend/service/rpc/globalRPC/internal/svc"
)

type GlobalRPCServer struct {
	svcCtx *svc.ServiceContext
	globalRPCProto.UnimplementedGlobalRPCServer
}

func NewGlobalRPCServer(svcCtx *svc.ServiceContext) *GlobalRPCServer {
	return &GlobalRPCServer{
		svcCtx: svcCtx,
	}
}

//  Asset
func (s *GlobalRPCServer) GetLatestAccountLp(ctx context.Context, in *globalRPCProto.ReqGetLatestAccountLp) (*globalRPCProto.RespGetLatestAccountLp, error) {
	l := logic.NewGetLatestAccountLpLogic(ctx, s.svcCtx)
	return l.GetLatestAccountLp(in)
}

func (s *GlobalRPCServer) GetLatestAssetsListByAccountIndex(ctx context.Context, in *globalRPCProto.ReqGetLatestAssetsListByAccountIndex) (*globalRPCProto.RespGetLatestAssetsListByAccountIndex, error) {
	l := logic.NewGetLatestAssetsListByAccountIndexLogic(ctx, s.svcCtx)
	return l.GetLatestAssetsListByAccountIndex(in)
}

func (s *GlobalRPCServer) GetLatestAssetInfoByAccountIndexAndAssetId(ctx context.Context, in *globalRPCProto.ReqGetLatestAssetInfoByAccountIndexAndAssetId) (*globalRPCProto.RespGetLatestAssetInfoByAccountIndexAndAssetId, error) {
	l := logic.NewGetLatestAssetInfoByAccountIndexAndAssetIdLogic(ctx, s.svcCtx)
	return l.GetLatestAssetInfoByAccountIndexAndAssetId(in)
}

//  Liquidity
func (s *GlobalRPCServer) GetLatestPairInfo(ctx context.Context, in *globalRPCProto.ReqGetLatestPairInfo) (*globalRPCProto.RespGetLatestPairInfo, error) {
	l := logic.NewGetLatestPairInfoLogic(ctx, s.svcCtx)
	return l.GetLatestPairInfo(in)
}

func (s *GlobalRPCServer) GetSwapAmount(ctx context.Context, in *globalRPCProto.ReqGetSwapAmount) (*globalRPCProto.RespGetSwapAmount, error) {
	l := logic.NewGetSwapAmountLogic(ctx, s.svcCtx)
	return l.GetSwapAmount(in)
}

func (s *GlobalRPCServer) GetLpValue(ctx context.Context, in *globalRPCProto.ReqGetLpValue) (*globalRPCProto.RespGetLpValue, error) {
	l := logic.NewGetLpValueLogic(ctx, s.svcCtx)
	return l.GetLpValue(in)
}

//  Transaction
func (s *GlobalRPCServer) GetLatestTxsListByAccountIndex(ctx context.Context, in *globalRPCProto.ReqGetLatestTxsListByAccountIndex) (*globalRPCProto.RespGetLatestTxsListByAccountIndex, error) {
	l := logic.NewGetLatestTxsListByAccountIndexLogic(ctx, s.svcCtx)
	return l.GetLatestTxsListByAccountIndex(in)
}

func (s *GlobalRPCServer) GetLatestTxsListByAccountIndexAndTxType(ctx context.Context, in *globalRPCProto.ReqGetLatestTxsListByAccountIndexAndTxType) (*globalRPCProto.RespGetLatestTxsListByAccountIndexAndTxType, error) {
	l := logic.NewGetLatestTxsListByAccountIndexAndTxTypeLogic(ctx, s.svcCtx)
	return l.GetLatestTxsListByAccountIndexAndTxType(in)
}

func (s *GlobalRPCServer) SendTx(ctx context.Context, in *globalRPCProto.ReqSendTx) (*globalRPCProto.RespSendTx, error) {
	l := logic.NewSendTxLogic(ctx, s.svcCtx)
	return l.SendTx(in)
}

func (s *GlobalRPCServer) GetTransactionCount(ctx context.Context, in *globalRPCProto.ReqGetTransactionCount) (*globalRPCProto.RespGetTransactionCount, error) {
	l := logic.NewGetTransactionCountLogic(ctx, s.svcCtx)
	return l.GetTransactionCount(in)
}

//  NFT
func (s *GlobalRPCServer) GetMaxOfferId(ctx context.Context, in *globalRPCProto.ReqGetMaxOfferId) (*globalRPCProto.RespGetMaxOfferId, error) {
	l := logic.NewGetMaxOfferIdLogic(ctx, s.svcCtx)
	return l.GetMaxOfferId(in)
}
