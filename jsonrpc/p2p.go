package jsonrpc

import (
	"context"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/maoxs2/go-jsonrpc2"
	"github.com/multiformats/go-multiaddr"

	"github.com/ngchain/ngcore/ngp2p"
	"github.com/ngchain/ngcore/utils"
)

func (s *Server) addNodeFunc(msg *jsonrpc2.JsonRpcMessage) *jsonrpc2.JsonRpcMessage {
	var params string
	err := utils.JSON.Unmarshal(msg.Params, &params)
	if err != nil {
		return jsonrpc2.NewJsonRpcError(msg.ID, jsonrpc2.NewError(0, err))
	}

	targetAddr, err := multiaddr.NewMultiaddr(params)
	if err != nil {
		return jsonrpc2.NewJsonRpcError(msg.ID, jsonrpc2.NewError(0, err))
	}

	targetInfo, err := peer.AddrInfoFromP2pAddr(targetAddr)
	if err != nil {
		return jsonrpc2.NewJsonRpcError(msg.ID, jsonrpc2.NewError(0, err))
	}

	err = ngp2p.GetLocalNode().Connect(context.Background(), *targetInfo)
	if err != nil {
		return jsonrpc2.NewJsonRpcError(msg.ID, jsonrpc2.NewError(0, err))
	}

	ok, _ := utils.JSON.Marshal(true)
	return jsonrpc2.NewJsonRpcSuccess(msg.ID, ok)
}