package service

import (
	"encoding/hex"
	"fmt"

	"github.com/rockin0098/flash/base/crypto"
	. "github.com/rockin0098/flash/base/logger"
	"github.com/rockin0098/flash/proto/mtproto"
	"github.com/rockin0098/flash/server/session"
)

func (s *LProtoService) TL_req_pq_Process(sess *session.Session, tlobj mtproto.TLObject) (interface{}, error) {
	Log.Infof("entering... sess = %p", sess)

	tl := tlobj.(*mtproto.TL_req_pq)
	nonce := tl.Get_nonce()
	if nonce == nil || len(nonce) != 16 {
		Log.Warnf("invalid nonce = %v", hex.EncodeToString(nonce))
		return nil, fmt.Errorf("req_pq invalid nonce : %v", nonce)
	}

	Log.Infof("nonce = %v", hex.EncodeToString(nonce))

	mtpcryptor := sess.MTProtoCryptor()
	state := sess.MTProtoState()

	// // for debugging
	// tmpnonce, _ := hex.DecodeString("2d91ae9c85bbbd559fc7959341106a7d")
	// // for debugging ===> end

	resPQ := &mtproto.TL_resPQ{
		M_nonce: nonce,
		// M_server_nonce:                   tmpnonce, // crypto.GenerateNonce(16),
		M_server_nonce:                   crypto.GenerateNonce(16),
		M_pq:                             mtpcryptor.PQ,
		M_server_public_key_fingerprints: []int64{int64(mtpcryptor.Fingerprint)},
	}

	// sess 缓存 nonce
	state.Nonce = resPQ.Get_nonce()
	state.ServerNonce = resPQ.Get_server_nonce()
	sess.SetMTProtoState(state)

	return resPQ, nil
}

func (s *LProtoService) TL_req_DH_params_Process(sess *session.Session, tlobj mtproto.TLObject) (interface{}, error) {
	Log.Infof("entering... sess = %p", sess)

	tl := tlobj.(*mtproto.TL_req_DH_params)

	state := sess.MTProtoState()
	cryptor := sess.MTProtoCryptor()

	Log.Infof("tl.nonce = %v, state.nonce = %v",
		hex.EncodeToString(tl.Get_nonce()), hex.EncodeToString(state.Nonce))

	Log.Infof("tl.server_nonce = %v, state.ServerNonce = %v",
		hex.EncodeToString(tl.Get_server_nonce()), hex.EncodeToString(state.ServerNonce))

	Log.Infof("tl.p = %v, state.p = %v",
		hex.EncodeToString([]byte(tl.Get_p())), hex.EncodeToString(cryptor.P))

	Log.Infof("tl.q = %v, state.q = %v",
		hex.EncodeToString([]byte(tl.Get_q())), hex.EncodeToString(cryptor.Q))

	Log.Infof("tl.fingerprint = %v, state.fingerprint = %v",
		tl.Get_public_key_fingerprint(), int64(cryptor.Fingerprint))

	resPQ := &mtproto.TL_resPQ{}

	return resPQ, nil
}
