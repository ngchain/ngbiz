package jsonrpc

// TODO: add options on machine joining, e.g. encryption
// type runContractParams struct {
//	FakeAccountNum uint64 `json:"fakeAccountNum"`
//	RawContract    []byte `json:"rawContract"`
// }
//
// // runContractFunc typically used to run the long loop task. Can be treated as a deploy
// func (s *Server) runContractFunc(msg *jsonrpc2.JsonRpcMessage) *jsonrpc2.JsonRpcMessage {
//	var params runContractParams
//	err := utils.JSON.Unmarshal(msg.Params, &params)
//	if err != nil {
//		log.Error(err)
//		return jsonrpc2.NewJsonRpcError(msg.ID, jsonrpc2.NewError(0, err))
//	}
//
//	_, err = hive.NewVM(ngtypes.NewAccount(params.FakeAccountNum, nil, params.RawContract, nil))
//	if err != nil {
//		log.Error(err)
//		return jsonrpc2.NewJsonRpcError(msg.ID, jsonrpc2.NewError(0, err))
//	}
//
//	return jsonrpc2.NewJsonRpcSuccess(msg.ID, nil)
// }
