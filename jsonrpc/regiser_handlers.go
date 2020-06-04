package jsonrpc

import (
	"github.com/maoxs2/go-jsonrpc2"
	"github.com/maoxs2/go-jsonrpc2/jsonrpc2http"
)

// newHTTPHandler will create a jsonrpc2http.HTTPHandler struct and register jsonrpc functions onto it.
func newHTTPHandler(s *Server) *jsonrpc2http.HTTPHandler {
	httpHandler := jsonrpc2http.NewHTTPHandler()

	httpHandler.RegisterJsonRpcHandleFunc("test", func(message *jsonrpc2.JsonRpcMessage) *jsonrpc2.JsonRpcMessage {
		return jsonrpc2.NewJsonRpcSuccess(message.ID, []byte("pong"))
	})

	// p2p
	httpHandler.RegisterJsonRpcHandleFunc("addNode", s.addNodeFunc)

	// chain
	httpHandler.RegisterJsonRpcHandleFunc("getBlockByHeight", s.getBlockByHeightFunc)
	httpHandler.RegisterJsonRpcHandleFunc("getBlockByHash", s.getBlockByHeightFunc)

	// state
	httpHandler.RegisterJsonRpcHandleFunc("sendTx", s.sendTxFunc)
	httpHandler.RegisterJsonRpcHandleFunc("signTx", s.signTxFunc)
	httpHandler.RegisterJsonRpcHandleFunc("genRegister", s.genRegisterFunc)
	httpHandler.RegisterJsonRpcHandleFunc("genLogout", s.genLogoutFunc)
	httpHandler.RegisterJsonRpcHandleFunc("genTransaction", s.genTransactionFunc)
	httpHandler.RegisterJsonRpcHandleFunc("genAssign", s.genAssignFunc)
	httpHandler.RegisterJsonRpcHandleFunc("genAppend", s.genAppendFunc)

	httpHandler.RegisterJsonRpcHandleFunc("getAccounts", s.getAccountsFunc)
	httpHandler.RegisterJsonRpcHandleFunc("getBalanceByNum", s.getBalanceByNumFunc)
	httpHandler.RegisterJsonRpcHandleFunc("getBalance", s.getBalanceFunc)

	return httpHandler
}
