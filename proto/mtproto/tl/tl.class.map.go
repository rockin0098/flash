package tl

type newTLObjectFunc func() TLObject

var tlObjectClassMap = map[int32]newTLObjectFunc{
	// int32(TLConstructor_CRC32_resPQ):                 func() TLObject { return NewTLResPQ() },
	// int32(TLConstructor_CRC32_p_q_inner_data):        func() TLObject { return NewTLPQInnerData() },
	// int32(TLConstructor_CRC32_server_DH_params_fail): func() TLObject { return NewTLServer_DHParamsFail() },
	// int32(TLConstructor_CRC32_server_DH_params_ok):   func() TLObject { return NewTLServer_DHParamsOk() },
	// int32(TLConstructor_CRC32_server_DH_inner_data):  func() TLObject { return NewTLServer_DHInnerData() },
	// int32(TLConstructor_CRC32_client_DH_inner_data):  func() TLObject { return NewTLClient_DHInnerData() },
}
