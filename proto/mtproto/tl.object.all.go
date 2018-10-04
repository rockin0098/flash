package mtproto

// resPQ#05162463
type TL_resPQ struct {
	_nonce                          []byte
	_server_nonce                   []byte
	_pq                             string
	_server_public_key_fingerprints []byte
}

func (t *TL_resPQ) Set_nonce(_nonce []byte) {
	t._nonce = _nonce
}

func (t *TL_resPQ) Get_nonce() []byte {
	return t._nonce
}

func (t *TL_resPQ) Set_server_nonce(_server_nonce []byte) {
	t._server_nonce = _server_nonce
}

func (t *TL_resPQ) Get_server_nonce() []byte {
	return t._server_nonce
}

func (t *TL_resPQ) Set_pq(_pq string) {
	t._pq = _pq
}

func (t *TL_resPQ) Get_pq() string {
	return t._pq
}

func (t *TL_resPQ) Set_server_public_key_fingerprints(_server_public_key_fingerprints []byte) {
	t._server_public_key_fingerprints = _server_public_key_fingerprints
}

func (t *TL_resPQ) Get_server_public_key_fingerprints() []byte {
	return t._server_public_key_fingerprints
}

func New_TL_resPQ() *TL_resPQ {
	return &TL_resPQ{}
}

func (t *TL_resPQ) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_resPQ))
	ec.Bytes(t.Get_nonce())
	ec.Bytes(t.Get_server_nonce())
	ec.String(t.Get_pq())
	ec.Bytes(t.Get_server_public_key_fingerprints())

	return ec.GetBuffer()
}

func (t *TL_resPQ) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._nonce = dc.Bytes(16)
	t._server_nonce = dc.Bytes(16)
	t._pq = dc.String()
	t._server_public_key_fingerprints = dc.Bytes(16)

}

// p_q_inner_data#83c95aec
type TL_p_q_inner_data struct {
	_pq           string
	_p            string
	_q            string
	_nonce        []byte
	_server_nonce []byte
	_new_nonce    []byte
}

func (t *TL_p_q_inner_data) Set_pq(_pq string) {
	t._pq = _pq
}

func (t *TL_p_q_inner_data) Get_pq() string {
	return t._pq
}

func (t *TL_p_q_inner_data) Set_p(_p string) {
	t._p = _p
}

func (t *TL_p_q_inner_data) Get_p() string {
	return t._p
}

func (t *TL_p_q_inner_data) Set_q(_q string) {
	t._q = _q
}

func (t *TL_p_q_inner_data) Get_q() string {
	return t._q
}

func (t *TL_p_q_inner_data) Set_nonce(_nonce []byte) {
	t._nonce = _nonce
}

func (t *TL_p_q_inner_data) Get_nonce() []byte {
	return t._nonce
}

func (t *TL_p_q_inner_data) Set_server_nonce(_server_nonce []byte) {
	t._server_nonce = _server_nonce
}

func (t *TL_p_q_inner_data) Get_server_nonce() []byte {
	return t._server_nonce
}

func (t *TL_p_q_inner_data) Set_new_nonce(_new_nonce []byte) {
	t._new_nonce = _new_nonce
}

func (t *TL_p_q_inner_data) Get_new_nonce() []byte {
	return t._new_nonce
}

func New_TL_p_q_inner_data() *TL_p_q_inner_data {
	return &TL_p_q_inner_data{}
}

func (t *TL_p_q_inner_data) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_p_q_inner_data))
	ec.String(t.Get_pq())
	ec.String(t.Get_p())
	ec.String(t.Get_q())
	ec.Bytes(t.Get_nonce())
	ec.Bytes(t.Get_server_nonce())
	ec.Bytes(t.Get_new_nonce())

	return ec.GetBuffer()
}

func (t *TL_p_q_inner_data) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._pq = dc.String()
	t._p = dc.String()
	t._q = dc.String()
	t._nonce = dc.Bytes(16)
	t._server_nonce = dc.Bytes(16)
	t._new_nonce = dc.Bytes(32)

}

// server_DH_params_fail#79cb045d
type TL_server_DH_params_fail struct {
	_nonce          []byte
	_server_nonce   []byte
	_new_nonce_hash []byte
}

func (t *TL_server_DH_params_fail) Set_nonce(_nonce []byte) {
	t._nonce = _nonce
}

func (t *TL_server_DH_params_fail) Get_nonce() []byte {
	return t._nonce
}

func (t *TL_server_DH_params_fail) Set_server_nonce(_server_nonce []byte) {
	t._server_nonce = _server_nonce
}

func (t *TL_server_DH_params_fail) Get_server_nonce() []byte {
	return t._server_nonce
}

func (t *TL_server_DH_params_fail) Set_new_nonce_hash(_new_nonce_hash []byte) {
	t._new_nonce_hash = _new_nonce_hash
}

func (t *TL_server_DH_params_fail) Get_new_nonce_hash() []byte {
	return t._new_nonce_hash
}

func New_TL_server_DH_params_fail() *TL_server_DH_params_fail {
	return &TL_server_DH_params_fail{}
}

func (t *TL_server_DH_params_fail) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_server_DH_params_fail))
	ec.Bytes(t.Get_nonce())
	ec.Bytes(t.Get_server_nonce())
	ec.Bytes(t.Get_new_nonce_hash())

	return ec.GetBuffer()
}

func (t *TL_server_DH_params_fail) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._nonce = dc.Bytes(16)
	t._server_nonce = dc.Bytes(16)
	t._new_nonce_hash = dc.Bytes(16)

}

// server_DH_params_ok#d0e8075c
type TL_server_DH_params_ok struct {
	_nonce            []byte
	_server_nonce     []byte
	_encrypted_answer string
}

func (t *TL_server_DH_params_ok) Set_nonce(_nonce []byte) {
	t._nonce = _nonce
}

func (t *TL_server_DH_params_ok) Get_nonce() []byte {
	return t._nonce
}

func (t *TL_server_DH_params_ok) Set_server_nonce(_server_nonce []byte) {
	t._server_nonce = _server_nonce
}

func (t *TL_server_DH_params_ok) Get_server_nonce() []byte {
	return t._server_nonce
}

func (t *TL_server_DH_params_ok) Set_encrypted_answer(_encrypted_answer string) {
	t._encrypted_answer = _encrypted_answer
}

func (t *TL_server_DH_params_ok) Get_encrypted_answer() string {
	return t._encrypted_answer
}

func New_TL_server_DH_params_ok() *TL_server_DH_params_ok {
	return &TL_server_DH_params_ok{}
}

func (t *TL_server_DH_params_ok) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_server_DH_params_ok))
	ec.Bytes(t.Get_nonce())
	ec.Bytes(t.Get_server_nonce())
	ec.String(t.Get_encrypted_answer())

	return ec.GetBuffer()
}

func (t *TL_server_DH_params_ok) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._nonce = dc.Bytes(16)
	t._server_nonce = dc.Bytes(16)
	t._encrypted_answer = dc.String()

}

// server_DH_inner_data#b5890dba
type TL_server_DH_inner_data struct {
	_nonce        []byte
	_server_nonce []byte
	_g            int32
	_dh_prime     string
	_g_a          string
	_server_time  int32
}

func (t *TL_server_DH_inner_data) Set_nonce(_nonce []byte) {
	t._nonce = _nonce
}

func (t *TL_server_DH_inner_data) Get_nonce() []byte {
	return t._nonce
}

func (t *TL_server_DH_inner_data) Set_server_nonce(_server_nonce []byte) {
	t._server_nonce = _server_nonce
}

func (t *TL_server_DH_inner_data) Get_server_nonce() []byte {
	return t._server_nonce
}

func (t *TL_server_DH_inner_data) Set_g(_g int32) {
	t._g = _g
}

func (t *TL_server_DH_inner_data) Get_g() int32 {
	return t._g
}

func (t *TL_server_DH_inner_data) Set_dh_prime(_dh_prime string) {
	t._dh_prime = _dh_prime
}

func (t *TL_server_DH_inner_data) Get_dh_prime() string {
	return t._dh_prime
}

func (t *TL_server_DH_inner_data) Set_g_a(_g_a string) {
	t._g_a = _g_a
}

func (t *TL_server_DH_inner_data) Get_g_a() string {
	return t._g_a
}

func (t *TL_server_DH_inner_data) Set_server_time(_server_time int32) {
	t._server_time = _server_time
}

func (t *TL_server_DH_inner_data) Get_server_time() int32 {
	return t._server_time
}

func New_TL_server_DH_inner_data() *TL_server_DH_inner_data {
	return &TL_server_DH_inner_data{}
}

func (t *TL_server_DH_inner_data) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_server_DH_inner_data))
	ec.Bytes(t.Get_nonce())
	ec.Bytes(t.Get_server_nonce())
	ec.Int(t.Get_g())
	ec.String(t.Get_dh_prime())
	ec.String(t.Get_g_a())
	ec.Int(t.Get_server_time())

	return ec.GetBuffer()
}

func (t *TL_server_DH_inner_data) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._nonce = dc.Bytes(16)
	t._server_nonce = dc.Bytes(16)
	t._g = dc.Int()
	t._dh_prime = dc.String()
	t._g_a = dc.String()
	t._server_time = dc.Int()

}

// client_DH_inner_data#6643b654
type TL_client_DH_inner_data struct {
	_nonce        []byte
	_server_nonce []byte
	_retry_id     int64
	_g_b          string
}

func (t *TL_client_DH_inner_data) Set_nonce(_nonce []byte) {
	t._nonce = _nonce
}

func (t *TL_client_DH_inner_data) Get_nonce() []byte {
	return t._nonce
}

func (t *TL_client_DH_inner_data) Set_server_nonce(_server_nonce []byte) {
	t._server_nonce = _server_nonce
}

func (t *TL_client_DH_inner_data) Get_server_nonce() []byte {
	return t._server_nonce
}

func (t *TL_client_DH_inner_data) Set_retry_id(_retry_id int64) {
	t._retry_id = _retry_id
}

func (t *TL_client_DH_inner_data) Get_retry_id() int64 {
	return t._retry_id
}

func (t *TL_client_DH_inner_data) Set_g_b(_g_b string) {
	t._g_b = _g_b
}

func (t *TL_client_DH_inner_data) Get_g_b() string {
	return t._g_b
}

func New_TL_client_DH_inner_data() *TL_client_DH_inner_data {
	return &TL_client_DH_inner_data{}
}

func (t *TL_client_DH_inner_data) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_client_DH_inner_data))
	ec.Bytes(t.Get_nonce())
	ec.Bytes(t.Get_server_nonce())
	ec.Long(t.Get_retry_id())
	ec.String(t.Get_g_b())

	return ec.GetBuffer()
}

func (t *TL_client_DH_inner_data) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._nonce = dc.Bytes(16)
	t._server_nonce = dc.Bytes(16)
	t._retry_id = dc.Long()
	t._g_b = dc.String()

}

// dh_gen_ok#3bcbf734
type TL_dh_gen_ok struct {
	_nonce           []byte
	_server_nonce    []byte
	_new_nonce_hash1 []byte
}

func (t *TL_dh_gen_ok) Set_nonce(_nonce []byte) {
	t._nonce = _nonce
}

func (t *TL_dh_gen_ok) Get_nonce() []byte {
	return t._nonce
}

func (t *TL_dh_gen_ok) Set_server_nonce(_server_nonce []byte) {
	t._server_nonce = _server_nonce
}

func (t *TL_dh_gen_ok) Get_server_nonce() []byte {
	return t._server_nonce
}

func (t *TL_dh_gen_ok) Set_new_nonce_hash1(_new_nonce_hash1 []byte) {
	t._new_nonce_hash1 = _new_nonce_hash1
}

func (t *TL_dh_gen_ok) Get_new_nonce_hash1() []byte {
	return t._new_nonce_hash1
}

func New_TL_dh_gen_ok() *TL_dh_gen_ok {
	return &TL_dh_gen_ok{}
}

func (t *TL_dh_gen_ok) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_dh_gen_ok))
	ec.Bytes(t.Get_nonce())
	ec.Bytes(t.Get_server_nonce())
	ec.Bytes(t.Get_new_nonce_hash1())

	return ec.GetBuffer()
}

func (t *TL_dh_gen_ok) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._nonce = dc.Bytes(16)
	t._server_nonce = dc.Bytes(16)
	t._new_nonce_hash1 = dc.Bytes(16)

}

// dh_gen_retry#46dc1fb9
type TL_dh_gen_retry struct {
	_nonce           []byte
	_server_nonce    []byte
	_new_nonce_hash2 []byte
}

func (t *TL_dh_gen_retry) Set_nonce(_nonce []byte) {
	t._nonce = _nonce
}

func (t *TL_dh_gen_retry) Get_nonce() []byte {
	return t._nonce
}

func (t *TL_dh_gen_retry) Set_server_nonce(_server_nonce []byte) {
	t._server_nonce = _server_nonce
}

func (t *TL_dh_gen_retry) Get_server_nonce() []byte {
	return t._server_nonce
}

func (t *TL_dh_gen_retry) Set_new_nonce_hash2(_new_nonce_hash2 []byte) {
	t._new_nonce_hash2 = _new_nonce_hash2
}

func (t *TL_dh_gen_retry) Get_new_nonce_hash2() []byte {
	return t._new_nonce_hash2
}

func New_TL_dh_gen_retry() *TL_dh_gen_retry {
	return &TL_dh_gen_retry{}
}

func (t *TL_dh_gen_retry) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_dh_gen_retry))
	ec.Bytes(t.Get_nonce())
	ec.Bytes(t.Get_server_nonce())
	ec.Bytes(t.Get_new_nonce_hash2())

	return ec.GetBuffer()
}

func (t *TL_dh_gen_retry) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._nonce = dc.Bytes(16)
	t._server_nonce = dc.Bytes(16)
	t._new_nonce_hash2 = dc.Bytes(16)

}

// dh_gen_fail#a69dae02
type TL_dh_gen_fail struct {
	_nonce           []byte
	_server_nonce    []byte
	_new_nonce_hash3 []byte
}

func (t *TL_dh_gen_fail) Set_nonce(_nonce []byte) {
	t._nonce = _nonce
}

func (t *TL_dh_gen_fail) Get_nonce() []byte {
	return t._nonce
}

func (t *TL_dh_gen_fail) Set_server_nonce(_server_nonce []byte) {
	t._server_nonce = _server_nonce
}

func (t *TL_dh_gen_fail) Get_server_nonce() []byte {
	return t._server_nonce
}

func (t *TL_dh_gen_fail) Set_new_nonce_hash3(_new_nonce_hash3 []byte) {
	t._new_nonce_hash3 = _new_nonce_hash3
}

func (t *TL_dh_gen_fail) Get_new_nonce_hash3() []byte {
	return t._new_nonce_hash3
}

func New_TL_dh_gen_fail() *TL_dh_gen_fail {
	return &TL_dh_gen_fail{}
}

func (t *TL_dh_gen_fail) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_dh_gen_fail))
	ec.Bytes(t.Get_nonce())
	ec.Bytes(t.Get_server_nonce())
	ec.Bytes(t.Get_new_nonce_hash3())

	return ec.GetBuffer()
}

func (t *TL_dh_gen_fail) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._nonce = dc.Bytes(16)
	t._server_nonce = dc.Bytes(16)
	t._new_nonce_hash3 = dc.Bytes(16)

}

// destroy_auth_key_ok#f660e1d4
type TL_destroy_auth_key_ok struct {
}

func New_TL_destroy_auth_key_ok() *TL_destroy_auth_key_ok {
	return &TL_destroy_auth_key_ok{}
}

func (t *TL_destroy_auth_key_ok) Encode() []byte {
	return nil
}

func (t *TL_destroy_auth_key_ok) Decode(b []byte) {

}

// destroy_auth_key_none#0a9f2259
type TL_destroy_auth_key_none struct {
}

func New_TL_destroy_auth_key_none() *TL_destroy_auth_key_none {
	return &TL_destroy_auth_key_none{}
}

func (t *TL_destroy_auth_key_none) Encode() []byte {
	return nil
}

func (t *TL_destroy_auth_key_none) Decode(b []byte) {

}

// destroy_auth_key_fail#ea109b13
type TL_destroy_auth_key_fail struct {
}

func New_TL_destroy_auth_key_fail() *TL_destroy_auth_key_fail {
	return &TL_destroy_auth_key_fail{}
}

func (t *TL_destroy_auth_key_fail) Encode() []byte {
	return nil
}

func (t *TL_destroy_auth_key_fail) Decode(b []byte) {

}

// req_pq#60469778
type TL_req_pq struct {
	_nonce []byte
}

func (t *TL_req_pq) Set_nonce(_nonce []byte) {
	t._nonce = _nonce
}

func (t *TL_req_pq) Get_nonce() []byte {
	return t._nonce
}

func New_TL_req_pq() *TL_req_pq {
	return &TL_req_pq{}
}

func (t *TL_req_pq) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_req_pq))
	ec.Bytes(t.Get_nonce())

	return ec.GetBuffer()
}

func (t *TL_req_pq) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._nonce = dc.Bytes(16)

}

// req_DH_params#d712e4be
type TL_req_DH_params struct {
	_nonce                  []byte
	_server_nonce           []byte
	_p                      string
	_q                      string
	_public_key_fingerprint int64
	_encrypted_data         string
}

func (t *TL_req_DH_params) Set_nonce(_nonce []byte) {
	t._nonce = _nonce
}

func (t *TL_req_DH_params) Get_nonce() []byte {
	return t._nonce
}

func (t *TL_req_DH_params) Set_server_nonce(_server_nonce []byte) {
	t._server_nonce = _server_nonce
}

func (t *TL_req_DH_params) Get_server_nonce() []byte {
	return t._server_nonce
}

func (t *TL_req_DH_params) Set_p(_p string) {
	t._p = _p
}

func (t *TL_req_DH_params) Get_p() string {
	return t._p
}

func (t *TL_req_DH_params) Set_q(_q string) {
	t._q = _q
}

func (t *TL_req_DH_params) Get_q() string {
	return t._q
}

func (t *TL_req_DH_params) Set_public_key_fingerprint(_public_key_fingerprint int64) {
	t._public_key_fingerprint = _public_key_fingerprint
}

func (t *TL_req_DH_params) Get_public_key_fingerprint() int64 {
	return t._public_key_fingerprint
}

func (t *TL_req_DH_params) Set_encrypted_data(_encrypted_data string) {
	t._encrypted_data = _encrypted_data
}

func (t *TL_req_DH_params) Get_encrypted_data() string {
	return t._encrypted_data
}

func New_TL_req_DH_params() *TL_req_DH_params {
	return &TL_req_DH_params{}
}

func (t *TL_req_DH_params) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_req_DH_params))
	ec.Bytes(t.Get_nonce())
	ec.Bytes(t.Get_server_nonce())
	ec.String(t.Get_p())
	ec.String(t.Get_q())
	ec.Long(t.Get_public_key_fingerprint())
	ec.String(t.Get_encrypted_data())

	return ec.GetBuffer()
}

func (t *TL_req_DH_params) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._nonce = dc.Bytes(16)
	t._server_nonce = dc.Bytes(16)
	t._p = dc.String()
	t._q = dc.String()
	t._public_key_fingerprint = dc.Long()
	t._encrypted_data = dc.String()

}

// set_client_DH_params#f5045f1f
type TL_set_client_DH_params struct {
	_nonce          []byte
	_server_nonce   []byte
	_encrypted_data string
}

func (t *TL_set_client_DH_params) Set_nonce(_nonce []byte) {
	t._nonce = _nonce
}

func (t *TL_set_client_DH_params) Get_nonce() []byte {
	return t._nonce
}

func (t *TL_set_client_DH_params) Set_server_nonce(_server_nonce []byte) {
	t._server_nonce = _server_nonce
}

func (t *TL_set_client_DH_params) Get_server_nonce() []byte {
	return t._server_nonce
}

func (t *TL_set_client_DH_params) Set_encrypted_data(_encrypted_data string) {
	t._encrypted_data = _encrypted_data
}

func (t *TL_set_client_DH_params) Get_encrypted_data() string {
	return t._encrypted_data
}

func New_TL_set_client_DH_params() *TL_set_client_DH_params {
	return &TL_set_client_DH_params{}
}

func (t *TL_set_client_DH_params) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_set_client_DH_params))
	ec.Bytes(t.Get_nonce())
	ec.Bytes(t.Get_server_nonce())
	ec.String(t.Get_encrypted_data())

	return ec.GetBuffer()
}

func (t *TL_set_client_DH_params) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._nonce = dc.Bytes(16)
	t._server_nonce = dc.Bytes(16)
	t._encrypted_data = dc.String()

}

// destroy_auth_key#d1435160
type TL_destroy_auth_key struct {
}

func New_TL_destroy_auth_key() *TL_destroy_auth_key {
	return &TL_destroy_auth_key{}
}

func (t *TL_destroy_auth_key) Encode() []byte {
	return nil
}

func (t *TL_destroy_auth_key) Decode(b []byte) {

}

// msgs_ack#62d6b459
type TL_msgs_ack struct {
	_msg_ids []byte
}

func (t *TL_msgs_ack) Set_msg_ids(_msg_ids []byte) {
	t._msg_ids = _msg_ids
}

func (t *TL_msgs_ack) Get_msg_ids() []byte {
	return t._msg_ids
}

func New_TL_msgs_ack() *TL_msgs_ack {
	return &TL_msgs_ack{}
}

func (t *TL_msgs_ack) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_msgs_ack))
	ec.Bytes(t.Get_msg_ids())

	return ec.GetBuffer()
}

func (t *TL_msgs_ack) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._msg_ids = dc.Bytes(16)

}

// bad_msg_notification#a7eff811
type TL_bad_msg_notification struct {
	_bad_msg_id    int64
	_bad_msg_seqno int32
	_error_code    int32
}

func (t *TL_bad_msg_notification) Set_bad_msg_id(_bad_msg_id int64) {
	t._bad_msg_id = _bad_msg_id
}

func (t *TL_bad_msg_notification) Get_bad_msg_id() int64 {
	return t._bad_msg_id
}

func (t *TL_bad_msg_notification) Set_bad_msg_seqno(_bad_msg_seqno int32) {
	t._bad_msg_seqno = _bad_msg_seqno
}

func (t *TL_bad_msg_notification) Get_bad_msg_seqno() int32 {
	return t._bad_msg_seqno
}

func (t *TL_bad_msg_notification) Set_error_code(_error_code int32) {
	t._error_code = _error_code
}

func (t *TL_bad_msg_notification) Get_error_code() int32 {
	return t._error_code
}

func New_TL_bad_msg_notification() *TL_bad_msg_notification {
	return &TL_bad_msg_notification{}
}

func (t *TL_bad_msg_notification) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_bad_msg_notification))
	ec.Long(t.Get_bad_msg_id())
	ec.Int(t.Get_bad_msg_seqno())
	ec.Int(t.Get_error_code())

	return ec.GetBuffer()
}

func (t *TL_bad_msg_notification) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._bad_msg_id = dc.Long()
	t._bad_msg_seqno = dc.Int()
	t._error_code = dc.Int()

}

// bad_server_salt#edab447b
type TL_bad_server_salt struct {
	_bad_msg_id      int64
	_bad_msg_seqno   int32
	_error_code      int32
	_new_server_salt int64
}

func (t *TL_bad_server_salt) Set_bad_msg_id(_bad_msg_id int64) {
	t._bad_msg_id = _bad_msg_id
}

func (t *TL_bad_server_salt) Get_bad_msg_id() int64 {
	return t._bad_msg_id
}

func (t *TL_bad_server_salt) Set_bad_msg_seqno(_bad_msg_seqno int32) {
	t._bad_msg_seqno = _bad_msg_seqno
}

func (t *TL_bad_server_salt) Get_bad_msg_seqno() int32 {
	return t._bad_msg_seqno
}

func (t *TL_bad_server_salt) Set_error_code(_error_code int32) {
	t._error_code = _error_code
}

func (t *TL_bad_server_salt) Get_error_code() int32 {
	return t._error_code
}

func (t *TL_bad_server_salt) Set_new_server_salt(_new_server_salt int64) {
	t._new_server_salt = _new_server_salt
}

func (t *TL_bad_server_salt) Get_new_server_salt() int64 {
	return t._new_server_salt
}

func New_TL_bad_server_salt() *TL_bad_server_salt {
	return &TL_bad_server_salt{}
}

func (t *TL_bad_server_salt) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_bad_server_salt))
	ec.Long(t.Get_bad_msg_id())
	ec.Int(t.Get_bad_msg_seqno())
	ec.Int(t.Get_error_code())
	ec.Long(t.Get_new_server_salt())

	return ec.GetBuffer()
}

func (t *TL_bad_server_salt) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._bad_msg_id = dc.Long()
	t._bad_msg_seqno = dc.Int()
	t._error_code = dc.Int()
	t._new_server_salt = dc.Long()

}

// msgs_state_req#da69fb52
type TL_msgs_state_req struct {
	_msg_ids []byte
}

func (t *TL_msgs_state_req) Set_msg_ids(_msg_ids []byte) {
	t._msg_ids = _msg_ids
}

func (t *TL_msgs_state_req) Get_msg_ids() []byte {
	return t._msg_ids
}

func New_TL_msgs_state_req() *TL_msgs_state_req {
	return &TL_msgs_state_req{}
}

func (t *TL_msgs_state_req) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_msgs_state_req))
	ec.Bytes(t.Get_msg_ids())

	return ec.GetBuffer()
}

func (t *TL_msgs_state_req) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._msg_ids = dc.Bytes(16)

}

// msgs_state_info#04deb57d
type TL_msgs_state_info struct {
	_req_msg_id int64
	_info       string
}

func (t *TL_msgs_state_info) Set_req_msg_id(_req_msg_id int64) {
	t._req_msg_id = _req_msg_id
}

func (t *TL_msgs_state_info) Get_req_msg_id() int64 {
	return t._req_msg_id
}

func (t *TL_msgs_state_info) Set_info(_info string) {
	t._info = _info
}

func (t *TL_msgs_state_info) Get_info() string {
	return t._info
}

func New_TL_msgs_state_info() *TL_msgs_state_info {
	return &TL_msgs_state_info{}
}

func (t *TL_msgs_state_info) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_msgs_state_info))
	ec.Long(t.Get_req_msg_id())
	ec.String(t.Get_info())

	return ec.GetBuffer()
}

func (t *TL_msgs_state_info) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._req_msg_id = dc.Long()
	t._info = dc.String()

}

// msgs_all_info#8cc0d131
type TL_msgs_all_info struct {
	_msg_ids []byte
	_info    string
}

func (t *TL_msgs_all_info) Set_msg_ids(_msg_ids []byte) {
	t._msg_ids = _msg_ids
}

func (t *TL_msgs_all_info) Get_msg_ids() []byte {
	return t._msg_ids
}

func (t *TL_msgs_all_info) Set_info(_info string) {
	t._info = _info
}

func (t *TL_msgs_all_info) Get_info() string {
	return t._info
}

func New_TL_msgs_all_info() *TL_msgs_all_info {
	return &TL_msgs_all_info{}
}

func (t *TL_msgs_all_info) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_msgs_all_info))
	ec.Bytes(t.Get_msg_ids())
	ec.String(t.Get_info())

	return ec.GetBuffer()
}

func (t *TL_msgs_all_info) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._msg_ids = dc.Bytes(16)
	t._info = dc.String()

}

// msg_detailed_info#276d3ec6
type TL_msg_detailed_info struct {
	_msg_id        int64
	_answer_msg_id int64
	_bytes         int32
	_status        int32
}

func (t *TL_msg_detailed_info) Set_msg_id(_msg_id int64) {
	t._msg_id = _msg_id
}

func (t *TL_msg_detailed_info) Get_msg_id() int64 {
	return t._msg_id
}

func (t *TL_msg_detailed_info) Set_answer_msg_id(_answer_msg_id int64) {
	t._answer_msg_id = _answer_msg_id
}

func (t *TL_msg_detailed_info) Get_answer_msg_id() int64 {
	return t._answer_msg_id
}

func (t *TL_msg_detailed_info) Set_bytes(_bytes int32) {
	t._bytes = _bytes
}

func (t *TL_msg_detailed_info) Get_bytes() int32 {
	return t._bytes
}

func (t *TL_msg_detailed_info) Set_status(_status int32) {
	t._status = _status
}

func (t *TL_msg_detailed_info) Get_status() int32 {
	return t._status
}

func New_TL_msg_detailed_info() *TL_msg_detailed_info {
	return &TL_msg_detailed_info{}
}

func (t *TL_msg_detailed_info) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_msg_detailed_info))
	ec.Long(t.Get_msg_id())
	ec.Long(t.Get_answer_msg_id())
	ec.Int(t.Get_bytes())
	ec.Int(t.Get_status())

	return ec.GetBuffer()
}

func (t *TL_msg_detailed_info) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._msg_id = dc.Long()
	t._answer_msg_id = dc.Long()
	t._bytes = dc.Int()
	t._status = dc.Int()

}

// msg_new_detailed_info#809db6df
type TL_msg_new_detailed_info struct {
	_answer_msg_id int64
	_bytes         int32
	_status        int32
}

func (t *TL_msg_new_detailed_info) Set_answer_msg_id(_answer_msg_id int64) {
	t._answer_msg_id = _answer_msg_id
}

func (t *TL_msg_new_detailed_info) Get_answer_msg_id() int64 {
	return t._answer_msg_id
}

func (t *TL_msg_new_detailed_info) Set_bytes(_bytes int32) {
	t._bytes = _bytes
}

func (t *TL_msg_new_detailed_info) Get_bytes() int32 {
	return t._bytes
}

func (t *TL_msg_new_detailed_info) Set_status(_status int32) {
	t._status = _status
}

func (t *TL_msg_new_detailed_info) Get_status() int32 {
	return t._status
}

func New_TL_msg_new_detailed_info() *TL_msg_new_detailed_info {
	return &TL_msg_new_detailed_info{}
}

func (t *TL_msg_new_detailed_info) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_msg_new_detailed_info))
	ec.Long(t.Get_answer_msg_id())
	ec.Int(t.Get_bytes())
	ec.Int(t.Get_status())

	return ec.GetBuffer()
}

func (t *TL_msg_new_detailed_info) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._answer_msg_id = dc.Long()
	t._bytes = dc.Int()
	t._status = dc.Int()

}

// msg_resend_req#7d861a08
type TL_msg_resend_req struct {
	_msg_ids []byte
}

func (t *TL_msg_resend_req) Set_msg_ids(_msg_ids []byte) {
	t._msg_ids = _msg_ids
}

func (t *TL_msg_resend_req) Get_msg_ids() []byte {
	return t._msg_ids
}

func New_TL_msg_resend_req() *TL_msg_resend_req {
	return &TL_msg_resend_req{}
}

func (t *TL_msg_resend_req) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_msg_resend_req))
	ec.Bytes(t.Get_msg_ids())

	return ec.GetBuffer()
}

func (t *TL_msg_resend_req) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._msg_ids = dc.Bytes(16)

}

// rpc_error#2144ca19
type TL_rpc_error struct {
	_error_code    int32
	_error_message string
}

func (t *TL_rpc_error) Set_error_code(_error_code int32) {
	t._error_code = _error_code
}

func (t *TL_rpc_error) Get_error_code() int32 {
	return t._error_code
}

func (t *TL_rpc_error) Set_error_message(_error_message string) {
	t._error_message = _error_message
}

func (t *TL_rpc_error) Get_error_message() string {
	return t._error_message
}

func New_TL_rpc_error() *TL_rpc_error {
	return &TL_rpc_error{}
}

func (t *TL_rpc_error) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_rpc_error))
	ec.Int(t.Get_error_code())
	ec.String(t.Get_error_message())

	return ec.GetBuffer()
}

func (t *TL_rpc_error) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._error_code = dc.Int()
	t._error_message = dc.String()

}

// rpc_answer_unknown#5e2ad36e
type TL_rpc_answer_unknown struct {
}

func New_TL_rpc_answer_unknown() *TL_rpc_answer_unknown {
	return &TL_rpc_answer_unknown{}
}

func (t *TL_rpc_answer_unknown) Encode() []byte {
	return nil
}

func (t *TL_rpc_answer_unknown) Decode(b []byte) {

}

// rpc_answer_dropped_running#cd78e586
type TL_rpc_answer_dropped_running struct {
}

func New_TL_rpc_answer_dropped_running() *TL_rpc_answer_dropped_running {
	return &TL_rpc_answer_dropped_running{}
}

func (t *TL_rpc_answer_dropped_running) Encode() []byte {
	return nil
}

func (t *TL_rpc_answer_dropped_running) Decode(b []byte) {

}

// rpc_answer_dropped#a43ad8b7
type TL_rpc_answer_dropped struct {
	_msg_id int64
	_seq_no int32
	_bytes  int32
}

func (t *TL_rpc_answer_dropped) Set_msg_id(_msg_id int64) {
	t._msg_id = _msg_id
}

func (t *TL_rpc_answer_dropped) Get_msg_id() int64 {
	return t._msg_id
}

func (t *TL_rpc_answer_dropped) Set_seq_no(_seq_no int32) {
	t._seq_no = _seq_no
}

func (t *TL_rpc_answer_dropped) Get_seq_no() int32 {
	return t._seq_no
}

func (t *TL_rpc_answer_dropped) Set_bytes(_bytes int32) {
	t._bytes = _bytes
}

func (t *TL_rpc_answer_dropped) Get_bytes() int32 {
	return t._bytes
}

func New_TL_rpc_answer_dropped() *TL_rpc_answer_dropped {
	return &TL_rpc_answer_dropped{}
}

func (t *TL_rpc_answer_dropped) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_rpc_answer_dropped))
	ec.Long(t.Get_msg_id())
	ec.Int(t.Get_seq_no())
	ec.Int(t.Get_bytes())

	return ec.GetBuffer()
}

func (t *TL_rpc_answer_dropped) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._msg_id = dc.Long()
	t._seq_no = dc.Int()
	t._bytes = dc.Int()

}

// future_salt#0949d9dc
type TL_future_salt struct {
	_valid_since int32
	_valid_until int32
	_salt        int64
}

func (t *TL_future_salt) Set_valid_since(_valid_since int32) {
	t._valid_since = _valid_since
}

func (t *TL_future_salt) Get_valid_since() int32 {
	return t._valid_since
}

func (t *TL_future_salt) Set_valid_until(_valid_until int32) {
	t._valid_until = _valid_until
}

func (t *TL_future_salt) Get_valid_until() int32 {
	return t._valid_until
}

func (t *TL_future_salt) Set_salt(_salt int64) {
	t._salt = _salt
}

func (t *TL_future_salt) Get_salt() int64 {
	return t._salt
}

func New_TL_future_salt() *TL_future_salt {
	return &TL_future_salt{}
}

func (t *TL_future_salt) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_future_salt))
	ec.Int(t.Get_valid_since())
	ec.Int(t.Get_valid_until())
	ec.Long(t.Get_salt())

	return ec.GetBuffer()
}

func (t *TL_future_salt) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._valid_since = dc.Int()
	t._valid_until = dc.Int()
	t._salt = dc.Long()

}

// future_salts#ae500895
type TL_future_salts struct {
	_req_msg_id int64
	_now        int32
	_salts      []byte
}

func (t *TL_future_salts) Set_req_msg_id(_req_msg_id int64) {
	t._req_msg_id = _req_msg_id
}

func (t *TL_future_salts) Get_req_msg_id() int64 {
	return t._req_msg_id
}

func (t *TL_future_salts) Set_now(_now int32) {
	t._now = _now
}

func (t *TL_future_salts) Get_now() int32 {
	return t._now
}

func (t *TL_future_salts) Set_salts(_salts []byte) {
	t._salts = _salts
}

func (t *TL_future_salts) Get_salts() []byte {
	return t._salts
}

func New_TL_future_salts() *TL_future_salts {
	return &TL_future_salts{}
}

func (t *TL_future_salts) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_future_salts))
	ec.Long(t.Get_req_msg_id())
	ec.Int(t.Get_now())
	ec.Bytes(t.Get_salts())

	return ec.GetBuffer()
}

func (t *TL_future_salts) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._req_msg_id = dc.Long()
	t._now = dc.Int()
	t._salts = dc.Bytes(16)

}

// pong#347773c5
type TL_pong struct {
	_msg_id  int64
	_ping_id int64
}

func (t *TL_pong) Set_msg_id(_msg_id int64) {
	t._msg_id = _msg_id
}

func (t *TL_pong) Get_msg_id() int64 {
	return t._msg_id
}

func (t *TL_pong) Set_ping_id(_ping_id int64) {
	t._ping_id = _ping_id
}

func (t *TL_pong) Get_ping_id() int64 {
	return t._ping_id
}

func New_TL_pong() *TL_pong {
	return &TL_pong{}
}

func (t *TL_pong) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_pong))
	ec.Long(t.Get_msg_id())
	ec.Long(t.Get_ping_id())

	return ec.GetBuffer()
}

func (t *TL_pong) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._msg_id = dc.Long()
	t._ping_id = dc.Long()

}

// destroy_session_ok#e22045fc
type TL_destroy_session_ok struct {
	_session_id int64
}

func (t *TL_destroy_session_ok) Set_session_id(_session_id int64) {
	t._session_id = _session_id
}

func (t *TL_destroy_session_ok) Get_session_id() int64 {
	return t._session_id
}

func New_TL_destroy_session_ok() *TL_destroy_session_ok {
	return &TL_destroy_session_ok{}
}

func (t *TL_destroy_session_ok) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_destroy_session_ok))
	ec.Long(t.Get_session_id())

	return ec.GetBuffer()
}

func (t *TL_destroy_session_ok) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._session_id = dc.Long()

}

// destroy_session_none#62d350c9
type TL_destroy_session_none struct {
	_session_id int64
}

func (t *TL_destroy_session_none) Set_session_id(_session_id int64) {
	t._session_id = _session_id
}

func (t *TL_destroy_session_none) Get_session_id() int64 {
	return t._session_id
}

func New_TL_destroy_session_none() *TL_destroy_session_none {
	return &TL_destroy_session_none{}
}

func (t *TL_destroy_session_none) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_destroy_session_none))
	ec.Long(t.Get_session_id())

	return ec.GetBuffer()
}

func (t *TL_destroy_session_none) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._session_id = dc.Long()

}

// new_session_created#9ec20908
type TL_new_session_created struct {
	_first_msg_id int64
	_unique_id    int64
	_server_salt  int64
}

func (t *TL_new_session_created) Set_first_msg_id(_first_msg_id int64) {
	t._first_msg_id = _first_msg_id
}

func (t *TL_new_session_created) Get_first_msg_id() int64 {
	return t._first_msg_id
}

func (t *TL_new_session_created) Set_unique_id(_unique_id int64) {
	t._unique_id = _unique_id
}

func (t *TL_new_session_created) Get_unique_id() int64 {
	return t._unique_id
}

func (t *TL_new_session_created) Set_server_salt(_server_salt int64) {
	t._server_salt = _server_salt
}

func (t *TL_new_session_created) Get_server_salt() int64 {
	return t._server_salt
}

func New_TL_new_session_created() *TL_new_session_created {
	return &TL_new_session_created{}
}

func (t *TL_new_session_created) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_new_session_created))
	ec.Long(t.Get_first_msg_id())
	ec.Long(t.Get_unique_id())
	ec.Long(t.Get_server_salt())

	return ec.GetBuffer()
}

func (t *TL_new_session_created) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._first_msg_id = dc.Long()
	t._unique_id = dc.Long()
	t._server_salt = dc.Long()

}

// http_wait#9299359f
type TL_http_wait struct {
	_max_delay  int32
	_wait_after int32
	_max_wait   int32
}

func (t *TL_http_wait) Set_max_delay(_max_delay int32) {
	t._max_delay = _max_delay
}

func (t *TL_http_wait) Get_max_delay() int32 {
	return t._max_delay
}

func (t *TL_http_wait) Set_wait_after(_wait_after int32) {
	t._wait_after = _wait_after
}

func (t *TL_http_wait) Get_wait_after() int32 {
	return t._wait_after
}

func (t *TL_http_wait) Set_max_wait(_max_wait int32) {
	t._max_wait = _max_wait
}

func (t *TL_http_wait) Get_max_wait() int32 {
	return t._max_wait
}

func New_TL_http_wait() *TL_http_wait {
	return &TL_http_wait{}
}

func (t *TL_http_wait) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_http_wait))
	ec.Int(t.Get_max_delay())
	ec.Int(t.Get_wait_after())
	ec.Int(t.Get_max_wait())

	return ec.GetBuffer()
}

func (t *TL_http_wait) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._max_delay = dc.Int()
	t._wait_after = dc.Int()
	t._max_wait = dc.Int()

}

// ipPort#
type TL_ipPort struct {
	_ipv4 int32
	_port int32
}

func (t *TL_ipPort) Set_ipv4(_ipv4 int32) {
	t._ipv4 = _ipv4
}

func (t *TL_ipPort) Get_ipv4() int32 {
	return t._ipv4
}

func (t *TL_ipPort) Set_port(_port int32) {
	t._port = _port
}

func (t *TL_ipPort) Get_port() int32 {
	return t._port
}

func New_TL_ipPort() *TL_ipPort {
	return &TL_ipPort{}
}

func (t *TL_ipPort) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_ipPort))
	ec.Int(t.Get_ipv4())
	ec.Int(t.Get_port())

	return ec.GetBuffer()
}

func (t *TL_ipPort) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._ipv4 = dc.Int()
	t._port = dc.Int()

}

// help_configSimple#d997c3c5
type TL_help_configSimple struct {
	_date         int32
	_expires      int32
	_dc_id        int32
	_ip_port_list []byte
}

func (t *TL_help_configSimple) Set_date(_date int32) {
	t._date = _date
}

func (t *TL_help_configSimple) Get_date() int32 {
	return t._date
}

func (t *TL_help_configSimple) Set_expires(_expires int32) {
	t._expires = _expires
}

func (t *TL_help_configSimple) Get_expires() int32 {
	return t._expires
}

func (t *TL_help_configSimple) Set_dc_id(_dc_id int32) {
	t._dc_id = _dc_id
}

func (t *TL_help_configSimple) Get_dc_id() int32 {
	return t._dc_id
}

func (t *TL_help_configSimple) Set_ip_port_list(_ip_port_list []byte) {
	t._ip_port_list = _ip_port_list
}

func (t *TL_help_configSimple) Get_ip_port_list() []byte {
	return t._ip_port_list
}

func New_TL_help_configSimple() *TL_help_configSimple {
	return &TL_help_configSimple{}
}

func (t *TL_help_configSimple) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_help_configSimple))
	ec.Int(t.Get_date())
	ec.Int(t.Get_expires())
	ec.Int(t.Get_dc_id())
	ec.Bytes(t.Get_ip_port_list())

	return ec.GetBuffer()
}

func (t *TL_help_configSimple) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._date = dc.Int()
	t._expires = dc.Int()
	t._dc_id = dc.Int()
	t._ip_port_list = dc.Bytes(16)

}

// rpc_drop_answer#58e4a740
type TL_rpc_drop_answer struct {
	_req_msg_id int64
}

func (t *TL_rpc_drop_answer) Set_req_msg_id(_req_msg_id int64) {
	t._req_msg_id = _req_msg_id
}

func (t *TL_rpc_drop_answer) Get_req_msg_id() int64 {
	return t._req_msg_id
}

func New_TL_rpc_drop_answer() *TL_rpc_drop_answer {
	return &TL_rpc_drop_answer{}
}

func (t *TL_rpc_drop_answer) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_rpc_drop_answer))
	ec.Long(t.Get_req_msg_id())

	return ec.GetBuffer()
}

func (t *TL_rpc_drop_answer) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._req_msg_id = dc.Long()

}

// get_future_salts#b921bd04
type TL_get_future_salts struct {
	_num int32
}

func (t *TL_get_future_salts) Set_num(_num int32) {
	t._num = _num
}

func (t *TL_get_future_salts) Get_num() int32 {
	return t._num
}

func New_TL_get_future_salts() *TL_get_future_salts {
	return &TL_get_future_salts{}
}

func (t *TL_get_future_salts) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_get_future_salts))
	ec.Int(t.Get_num())

	return ec.GetBuffer()
}

func (t *TL_get_future_salts) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._num = dc.Int()

}

// ping#7abe77ec
type TL_ping struct {
	_ping_id int64
}

func (t *TL_ping) Set_ping_id(_ping_id int64) {
	t._ping_id = _ping_id
}

func (t *TL_ping) Get_ping_id() int64 {
	return t._ping_id
}

func New_TL_ping() *TL_ping {
	return &TL_ping{}
}

func (t *TL_ping) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_ping))
	ec.Long(t.Get_ping_id())

	return ec.GetBuffer()
}

func (t *TL_ping) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._ping_id = dc.Long()

}

// ping_delay_disconnect#f3427b8c
type TL_ping_delay_disconnect struct {
	_ping_id          int64
	_disconnect_delay int32
}

func (t *TL_ping_delay_disconnect) Set_ping_id(_ping_id int64) {
	t._ping_id = _ping_id
}

func (t *TL_ping_delay_disconnect) Get_ping_id() int64 {
	return t._ping_id
}

func (t *TL_ping_delay_disconnect) Set_disconnect_delay(_disconnect_delay int32) {
	t._disconnect_delay = _disconnect_delay
}

func (t *TL_ping_delay_disconnect) Get_disconnect_delay() int32 {
	return t._disconnect_delay
}

func New_TL_ping_delay_disconnect() *TL_ping_delay_disconnect {
	return &TL_ping_delay_disconnect{}
}

func (t *TL_ping_delay_disconnect) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_ping_delay_disconnect))
	ec.Long(t.Get_ping_id())
	ec.Int(t.Get_disconnect_delay())

	return ec.GetBuffer()
}

func (t *TL_ping_delay_disconnect) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._ping_id = dc.Long()
	t._disconnect_delay = dc.Int()

}

// destroy_session#e7512126
type TL_destroy_session struct {
	_session_id int64
}

func (t *TL_destroy_session) Set_session_id(_session_id int64) {
	t._session_id = _session_id
}

func (t *TL_destroy_session) Get_session_id() int64 {
	return t._session_id
}

func New_TL_destroy_session() *TL_destroy_session {
	return &TL_destroy_session{}
}

func (t *TL_destroy_session) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_destroy_session))
	ec.Long(t.Get_session_id())

	return ec.GetBuffer()
}

func (t *TL_destroy_session) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._session_id = dc.Long()

}

// contest_saveDeveloperInfo#9a5f6e95
type TL_contest_saveDeveloperInfo struct {
	_vk_id        int32
	_name         string
	_phone_number string
	_age          int32
	_city         string
}

func (t *TL_contest_saveDeveloperInfo) Set_vk_id(_vk_id int32) {
	t._vk_id = _vk_id
}

func (t *TL_contest_saveDeveloperInfo) Get_vk_id() int32 {
	return t._vk_id
}

func (t *TL_contest_saveDeveloperInfo) Set_name(_name string) {
	t._name = _name
}

func (t *TL_contest_saveDeveloperInfo) Get_name() string {
	return t._name
}

func (t *TL_contest_saveDeveloperInfo) Set_phone_number(_phone_number string) {
	t._phone_number = _phone_number
}

func (t *TL_contest_saveDeveloperInfo) Get_phone_number() string {
	return t._phone_number
}

func (t *TL_contest_saveDeveloperInfo) Set_age(_age int32) {
	t._age = _age
}

func (t *TL_contest_saveDeveloperInfo) Get_age() int32 {
	return t._age
}

func (t *TL_contest_saveDeveloperInfo) Set_city(_city string) {
	t._city = _city
}

func (t *TL_contest_saveDeveloperInfo) Get_city() string {
	return t._city
}

func New_TL_contest_saveDeveloperInfo() *TL_contest_saveDeveloperInfo {
	return &TL_contest_saveDeveloperInfo{}
}

func (t *TL_contest_saveDeveloperInfo) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_contest_saveDeveloperInfo))
	ec.Int(t.Get_vk_id())
	ec.String(t.Get_name())
	ec.String(t.Get_phone_number())
	ec.Int(t.Get_age())
	ec.String(t.Get_city())

	return ec.GetBuffer()
}

func (t *TL_contest_saveDeveloperInfo) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._vk_id = dc.Int()
	t._name = dc.String()
	t._phone_number = dc.String()
	t._age = dc.Int()
	t._city = dc.String()

}

// boolFalse#bc799737
type TL_boolFalse struct {
}

func New_TL_boolFalse() *TL_boolFalse {
	return &TL_boolFalse{}
}

func (t *TL_boolFalse) Encode() []byte {
	return nil
}

func (t *TL_boolFalse) Decode(b []byte) {

}

// boolTrue#997275b5
type TL_boolTrue struct {
}

func New_TL_boolTrue() *TL_boolTrue {
	return &TL_boolTrue{}
}

func (t *TL_boolTrue) Encode() []byte {
	return nil
}

func (t *TL_boolTrue) Decode(b []byte) {

}

// true#3fedd339
type TL_true struct {
}

func New_TL_true() *TL_true {
	return &TL_true{}
}

func (t *TL_true) Encode() []byte {
	return nil
}

func (t *TL_true) Decode(b []byte) {

}

// vector#1cb5c415
type TL_vector struct {
}

func New_TL_vector() *TL_vector {
	return &TL_vector{}
}

func (t *TL_vector) Encode() []byte {
	return nil
}

func (t *TL_vector) Decode(b []byte) {

}

// error#c4b9f9bb
type TL_error struct {
	_code int32
	_text string
}

func (t *TL_error) Set_code(_code int32) {
	t._code = _code
}

func (t *TL_error) Get_code() int32 {
	return t._code
}

func (t *TL_error) Set_text(_text string) {
	t._text = _text
}

func (t *TL_error) Get_text() string {
	return t._text
}

func New_TL_error() *TL_error {
	return &TL_error{}
}

func (t *TL_error) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_error))
	ec.Int(t.Get_code())
	ec.String(t.Get_text())

	return ec.GetBuffer()
}

func (t *TL_error) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._code = dc.Int()
	t._text = dc.String()

}

// null#56730bcc
type TL_null struct {
}

func New_TL_null() *TL_null {
	return &TL_null{}
}

func (t *TL_null) Encode() []byte {
	return nil
}

func (t *TL_null) Decode(b []byte) {

}

// inputPeerEmpty#7f3b18ea
type TL_inputPeerEmpty struct {
}

func New_TL_inputPeerEmpty() *TL_inputPeerEmpty {
	return &TL_inputPeerEmpty{}
}

func (t *TL_inputPeerEmpty) Encode() []byte {
	return nil
}

func (t *TL_inputPeerEmpty) Decode(b []byte) {

}

// inputPeerSelf#7da07ec9
type TL_inputPeerSelf struct {
}

func New_TL_inputPeerSelf() *TL_inputPeerSelf {
	return &TL_inputPeerSelf{}
}

func (t *TL_inputPeerSelf) Encode() []byte {
	return nil
}

func (t *TL_inputPeerSelf) Decode(b []byte) {

}

// inputPeerChat#179be863
type TL_inputPeerChat struct {
	_chat_id int32
}

func (t *TL_inputPeerChat) Set_chat_id(_chat_id int32) {
	t._chat_id = _chat_id
}

func (t *TL_inputPeerChat) Get_chat_id() int32 {
	return t._chat_id
}

func New_TL_inputPeerChat() *TL_inputPeerChat {
	return &TL_inputPeerChat{}
}

func (t *TL_inputPeerChat) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputPeerChat))
	ec.Int(t.Get_chat_id())

	return ec.GetBuffer()
}

func (t *TL_inputPeerChat) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._chat_id = dc.Int()

}

// inputPeerUser#7b8e7de6
type TL_inputPeerUser struct {
	_user_id     int32
	_access_hash int64
}

func (t *TL_inputPeerUser) Set_user_id(_user_id int32) {
	t._user_id = _user_id
}

func (t *TL_inputPeerUser) Get_user_id() int32 {
	return t._user_id
}

func (t *TL_inputPeerUser) Set_access_hash(_access_hash int64) {
	t._access_hash = _access_hash
}

func (t *TL_inputPeerUser) Get_access_hash() int64 {
	return t._access_hash
}

func New_TL_inputPeerUser() *TL_inputPeerUser {
	return &TL_inputPeerUser{}
}

func (t *TL_inputPeerUser) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputPeerUser))
	ec.Int(t.Get_user_id())
	ec.Long(t.Get_access_hash())

	return ec.GetBuffer()
}

func (t *TL_inputPeerUser) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._user_id = dc.Int()
	t._access_hash = dc.Long()

}

// inputPeerChannel#20adaef8
type TL_inputPeerChannel struct {
	_channel_id  int32
	_access_hash int64
}

func (t *TL_inputPeerChannel) Set_channel_id(_channel_id int32) {
	t._channel_id = _channel_id
}

func (t *TL_inputPeerChannel) Get_channel_id() int32 {
	return t._channel_id
}

func (t *TL_inputPeerChannel) Set_access_hash(_access_hash int64) {
	t._access_hash = _access_hash
}

func (t *TL_inputPeerChannel) Get_access_hash() int64 {
	return t._access_hash
}

func New_TL_inputPeerChannel() *TL_inputPeerChannel {
	return &TL_inputPeerChannel{}
}

func (t *TL_inputPeerChannel) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputPeerChannel))
	ec.Int(t.Get_channel_id())
	ec.Long(t.Get_access_hash())

	return ec.GetBuffer()
}

func (t *TL_inputPeerChannel) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._channel_id = dc.Int()
	t._access_hash = dc.Long()

}

// inputUserEmpty#b98886cf
type TL_inputUserEmpty struct {
}

func New_TL_inputUserEmpty() *TL_inputUserEmpty {
	return &TL_inputUserEmpty{}
}

func (t *TL_inputUserEmpty) Encode() []byte {
	return nil
}

func (t *TL_inputUserEmpty) Decode(b []byte) {

}

// inputUserSelf#f7c1b13f
type TL_inputUserSelf struct {
}

func New_TL_inputUserSelf() *TL_inputUserSelf {
	return &TL_inputUserSelf{}
}

func (t *TL_inputUserSelf) Encode() []byte {
	return nil
}

func (t *TL_inputUserSelf) Decode(b []byte) {

}

// inputUser#d8292816
type TL_inputUser struct {
	_user_id     int32
	_access_hash int64
}

func (t *TL_inputUser) Set_user_id(_user_id int32) {
	t._user_id = _user_id
}

func (t *TL_inputUser) Get_user_id() int32 {
	return t._user_id
}

func (t *TL_inputUser) Set_access_hash(_access_hash int64) {
	t._access_hash = _access_hash
}

func (t *TL_inputUser) Get_access_hash() int64 {
	return t._access_hash
}

func New_TL_inputUser() *TL_inputUser {
	return &TL_inputUser{}
}

func (t *TL_inputUser) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputUser))
	ec.Int(t.Get_user_id())
	ec.Long(t.Get_access_hash())

	return ec.GetBuffer()
}

func (t *TL_inputUser) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._user_id = dc.Int()
	t._access_hash = dc.Long()

}

// inputPhoneContact#f392b7f4
type TL_inputPhoneContact struct {
	_client_id  int64
	_phone      string
	_first_name string
	_last_name  string
}

func (t *TL_inputPhoneContact) Set_client_id(_client_id int64) {
	t._client_id = _client_id
}

func (t *TL_inputPhoneContact) Get_client_id() int64 {
	return t._client_id
}

func (t *TL_inputPhoneContact) Set_phone(_phone string) {
	t._phone = _phone
}

func (t *TL_inputPhoneContact) Get_phone() string {
	return t._phone
}

func (t *TL_inputPhoneContact) Set_first_name(_first_name string) {
	t._first_name = _first_name
}

func (t *TL_inputPhoneContact) Get_first_name() string {
	return t._first_name
}

func (t *TL_inputPhoneContact) Set_last_name(_last_name string) {
	t._last_name = _last_name
}

func (t *TL_inputPhoneContact) Get_last_name() string {
	return t._last_name
}

func New_TL_inputPhoneContact() *TL_inputPhoneContact {
	return &TL_inputPhoneContact{}
}

func (t *TL_inputPhoneContact) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputPhoneContact))
	ec.Long(t.Get_client_id())
	ec.String(t.Get_phone())
	ec.String(t.Get_first_name())
	ec.String(t.Get_last_name())

	return ec.GetBuffer()
}

func (t *TL_inputPhoneContact) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._client_id = dc.Long()
	t._phone = dc.String()
	t._first_name = dc.String()
	t._last_name = dc.String()

}

// inputFile#f52ff27f
type TL_inputFile struct {
	_id           int64
	_parts        int32
	_name         string
	_md5_checksum string
}

func (t *TL_inputFile) Set_id(_id int64) {
	t._id = _id
}

func (t *TL_inputFile) Get_id() int64 {
	return t._id
}

func (t *TL_inputFile) Set_parts(_parts int32) {
	t._parts = _parts
}

func (t *TL_inputFile) Get_parts() int32 {
	return t._parts
}

func (t *TL_inputFile) Set_name(_name string) {
	t._name = _name
}

func (t *TL_inputFile) Get_name() string {
	return t._name
}

func (t *TL_inputFile) Set_md5_checksum(_md5_checksum string) {
	t._md5_checksum = _md5_checksum
}

func (t *TL_inputFile) Get_md5_checksum() string {
	return t._md5_checksum
}

func New_TL_inputFile() *TL_inputFile {
	return &TL_inputFile{}
}

func (t *TL_inputFile) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputFile))
	ec.Long(t.Get_id())
	ec.Int(t.Get_parts())
	ec.String(t.Get_name())
	ec.String(t.Get_md5_checksum())

	return ec.GetBuffer()
}

func (t *TL_inputFile) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Long()
	t._parts = dc.Int()
	t._name = dc.String()
	t._md5_checksum = dc.String()

}

// inputFileBig#fa4f0bb5
type TL_inputFileBig struct {
	_id    int64
	_parts int32
	_name  string
}

func (t *TL_inputFileBig) Set_id(_id int64) {
	t._id = _id
}

func (t *TL_inputFileBig) Get_id() int64 {
	return t._id
}

func (t *TL_inputFileBig) Set_parts(_parts int32) {
	t._parts = _parts
}

func (t *TL_inputFileBig) Get_parts() int32 {
	return t._parts
}

func (t *TL_inputFileBig) Set_name(_name string) {
	t._name = _name
}

func (t *TL_inputFileBig) Get_name() string {
	return t._name
}

func New_TL_inputFileBig() *TL_inputFileBig {
	return &TL_inputFileBig{}
}

func (t *TL_inputFileBig) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputFileBig))
	ec.Long(t.Get_id())
	ec.Int(t.Get_parts())
	ec.String(t.Get_name())

	return ec.GetBuffer()
}

func (t *TL_inputFileBig) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Long()
	t._parts = dc.Int()
	t._name = dc.String()

}

// inputMediaEmpty#9664f57f
type TL_inputMediaEmpty struct {
}

func New_TL_inputMediaEmpty() *TL_inputMediaEmpty {
	return &TL_inputMediaEmpty{}
}

func (t *TL_inputMediaEmpty) Encode() []byte {
	return nil
}

func (t *TL_inputMediaEmpty) Decode(b []byte) {

}

// inputMediaUploadedPhoto#2f37e231
type TL_inputMediaUploadedPhoto struct {
	_flags       []byte
	_file        []byte
	_caption     string
	_stickers    []byte
	_ttl_seconds []byte
}

func (t *TL_inputMediaUploadedPhoto) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_inputMediaUploadedPhoto) Get_flags() []byte {
	return t._flags
}

func (t *TL_inputMediaUploadedPhoto) Set_file(_file []byte) {
	t._file = _file
}

func (t *TL_inputMediaUploadedPhoto) Get_file() []byte {
	return t._file
}

func (t *TL_inputMediaUploadedPhoto) Set_caption(_caption string) {
	t._caption = _caption
}

func (t *TL_inputMediaUploadedPhoto) Get_caption() string {
	return t._caption
}

func (t *TL_inputMediaUploadedPhoto) Set_stickers(_stickers []byte) {
	t._stickers = _stickers
}

func (t *TL_inputMediaUploadedPhoto) Get_stickers() []byte {
	return t._stickers
}

func (t *TL_inputMediaUploadedPhoto) Set_ttl_seconds(_ttl_seconds []byte) {
	t._ttl_seconds = _ttl_seconds
}

func (t *TL_inputMediaUploadedPhoto) Get_ttl_seconds() []byte {
	return t._ttl_seconds
}

func New_TL_inputMediaUploadedPhoto() *TL_inputMediaUploadedPhoto {
	return &TL_inputMediaUploadedPhoto{}
}

func (t *TL_inputMediaUploadedPhoto) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputMediaUploadedPhoto))
	ec.Bytes(t.Get_file())
	ec.String(t.Get_caption())
	ec.Bytes(t.Get_stickers())
	ec.Bytes(t.Get_ttl_seconds())

	return ec.GetBuffer()
}

func (t *TL_inputMediaUploadedPhoto) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._file = dc.Bytes(16)
	t._caption = dc.String()
	t._stickers = dc.Bytes(16)
	t._ttl_seconds = dc.Bytes(16)

}

// inputMediaPhoto#81fa373a
type TL_inputMediaPhoto struct {
	_flags       []byte
	_id          []byte
	_caption     string
	_ttl_seconds []byte
}

func (t *TL_inputMediaPhoto) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_inputMediaPhoto) Get_flags() []byte {
	return t._flags
}

func (t *TL_inputMediaPhoto) Set_id(_id []byte) {
	t._id = _id
}

func (t *TL_inputMediaPhoto) Get_id() []byte {
	return t._id
}

func (t *TL_inputMediaPhoto) Set_caption(_caption string) {
	t._caption = _caption
}

func (t *TL_inputMediaPhoto) Get_caption() string {
	return t._caption
}

func (t *TL_inputMediaPhoto) Set_ttl_seconds(_ttl_seconds []byte) {
	t._ttl_seconds = _ttl_seconds
}

func (t *TL_inputMediaPhoto) Get_ttl_seconds() []byte {
	return t._ttl_seconds
}

func New_TL_inputMediaPhoto() *TL_inputMediaPhoto {
	return &TL_inputMediaPhoto{}
}

func (t *TL_inputMediaPhoto) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputMediaPhoto))
	ec.Bytes(t.Get_id())
	ec.String(t.Get_caption())
	ec.Bytes(t.Get_ttl_seconds())

	return ec.GetBuffer()
}

func (t *TL_inputMediaPhoto) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Bytes(16)
	t._caption = dc.String()
	t._ttl_seconds = dc.Bytes(16)

}

// inputMediaGeoPoint#f9c44144
type TL_inputMediaGeoPoint struct {
	_geo_point []byte
}

func (t *TL_inputMediaGeoPoint) Set_geo_point(_geo_point []byte) {
	t._geo_point = _geo_point
}

func (t *TL_inputMediaGeoPoint) Get_geo_point() []byte {
	return t._geo_point
}

func New_TL_inputMediaGeoPoint() *TL_inputMediaGeoPoint {
	return &TL_inputMediaGeoPoint{}
}

func (t *TL_inputMediaGeoPoint) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputMediaGeoPoint))
	ec.Bytes(t.Get_geo_point())

	return ec.GetBuffer()
}

func (t *TL_inputMediaGeoPoint) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._geo_point = dc.Bytes(16)

}

// inputMediaContact#a6e45987
type TL_inputMediaContact struct {
	_phone_number string
	_first_name   string
	_last_name    string
}

func (t *TL_inputMediaContact) Set_phone_number(_phone_number string) {
	t._phone_number = _phone_number
}

func (t *TL_inputMediaContact) Get_phone_number() string {
	return t._phone_number
}

func (t *TL_inputMediaContact) Set_first_name(_first_name string) {
	t._first_name = _first_name
}

func (t *TL_inputMediaContact) Get_first_name() string {
	return t._first_name
}

func (t *TL_inputMediaContact) Set_last_name(_last_name string) {
	t._last_name = _last_name
}

func (t *TL_inputMediaContact) Get_last_name() string {
	return t._last_name
}

func New_TL_inputMediaContact() *TL_inputMediaContact {
	return &TL_inputMediaContact{}
}

func (t *TL_inputMediaContact) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputMediaContact))
	ec.String(t.Get_phone_number())
	ec.String(t.Get_first_name())
	ec.String(t.Get_last_name())

	return ec.GetBuffer()
}

func (t *TL_inputMediaContact) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._phone_number = dc.String()
	t._first_name = dc.String()
	t._last_name = dc.String()

}

// inputMediaUploadedDocument#e39621fd
type TL_inputMediaUploadedDocument struct {
	_flags         []byte
	_nosound_video []byte
	_file          []byte
	_thumb         []byte
	_mime_type     string
	_attributes    []byte
	_caption       string
	_stickers      []byte
	_ttl_seconds   []byte
}

func (t *TL_inputMediaUploadedDocument) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_inputMediaUploadedDocument) Get_flags() []byte {
	return t._flags
}

func (t *TL_inputMediaUploadedDocument) Set_nosound_video(_nosound_video []byte) {
	t._nosound_video = _nosound_video
}

func (t *TL_inputMediaUploadedDocument) Get_nosound_video() []byte {
	return t._nosound_video
}

func (t *TL_inputMediaUploadedDocument) Set_file(_file []byte) {
	t._file = _file
}

func (t *TL_inputMediaUploadedDocument) Get_file() []byte {
	return t._file
}

func (t *TL_inputMediaUploadedDocument) Set_thumb(_thumb []byte) {
	t._thumb = _thumb
}

func (t *TL_inputMediaUploadedDocument) Get_thumb() []byte {
	return t._thumb
}

func (t *TL_inputMediaUploadedDocument) Set_mime_type(_mime_type string) {
	t._mime_type = _mime_type
}

func (t *TL_inputMediaUploadedDocument) Get_mime_type() string {
	return t._mime_type
}

func (t *TL_inputMediaUploadedDocument) Set_attributes(_attributes []byte) {
	t._attributes = _attributes
}

func (t *TL_inputMediaUploadedDocument) Get_attributes() []byte {
	return t._attributes
}

func (t *TL_inputMediaUploadedDocument) Set_caption(_caption string) {
	t._caption = _caption
}

func (t *TL_inputMediaUploadedDocument) Get_caption() string {
	return t._caption
}

func (t *TL_inputMediaUploadedDocument) Set_stickers(_stickers []byte) {
	t._stickers = _stickers
}

func (t *TL_inputMediaUploadedDocument) Get_stickers() []byte {
	return t._stickers
}

func (t *TL_inputMediaUploadedDocument) Set_ttl_seconds(_ttl_seconds []byte) {
	t._ttl_seconds = _ttl_seconds
}

func (t *TL_inputMediaUploadedDocument) Get_ttl_seconds() []byte {
	return t._ttl_seconds
}

func New_TL_inputMediaUploadedDocument() *TL_inputMediaUploadedDocument {
	return &TL_inputMediaUploadedDocument{}
}

func (t *TL_inputMediaUploadedDocument) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputMediaUploadedDocument))
	ec.Bytes(t.Get_nosound_video())
	ec.Bytes(t.Get_file())
	ec.Bytes(t.Get_thumb())
	ec.String(t.Get_mime_type())
	ec.Bytes(t.Get_attributes())
	ec.String(t.Get_caption())
	ec.Bytes(t.Get_stickers())
	ec.Bytes(t.Get_ttl_seconds())

	return ec.GetBuffer()
}

func (t *TL_inputMediaUploadedDocument) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._nosound_video = dc.Bytes(16)
	t._file = dc.Bytes(16)
	t._thumb = dc.Bytes(16)
	t._mime_type = dc.String()
	t._attributes = dc.Bytes(16)
	t._caption = dc.String()
	t._stickers = dc.Bytes(16)
	t._ttl_seconds = dc.Bytes(16)

}

// inputMediaDocument#5acb668e
type TL_inputMediaDocument struct {
	_flags       []byte
	_id          []byte
	_caption     string
	_ttl_seconds []byte
}

func (t *TL_inputMediaDocument) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_inputMediaDocument) Get_flags() []byte {
	return t._flags
}

func (t *TL_inputMediaDocument) Set_id(_id []byte) {
	t._id = _id
}

func (t *TL_inputMediaDocument) Get_id() []byte {
	return t._id
}

func (t *TL_inputMediaDocument) Set_caption(_caption string) {
	t._caption = _caption
}

func (t *TL_inputMediaDocument) Get_caption() string {
	return t._caption
}

func (t *TL_inputMediaDocument) Set_ttl_seconds(_ttl_seconds []byte) {
	t._ttl_seconds = _ttl_seconds
}

func (t *TL_inputMediaDocument) Get_ttl_seconds() []byte {
	return t._ttl_seconds
}

func New_TL_inputMediaDocument() *TL_inputMediaDocument {
	return &TL_inputMediaDocument{}
}

func (t *TL_inputMediaDocument) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputMediaDocument))
	ec.Bytes(t.Get_id())
	ec.String(t.Get_caption())
	ec.Bytes(t.Get_ttl_seconds())

	return ec.GetBuffer()
}

func (t *TL_inputMediaDocument) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Bytes(16)
	t._caption = dc.String()
	t._ttl_seconds = dc.Bytes(16)

}

// inputMediaVenue#c13d1c11
type TL_inputMediaVenue struct {
	_geo_point  []byte
	_title      string
	_address    string
	_provider   string
	_venue_id   string
	_venue_type string
}

func (t *TL_inputMediaVenue) Set_geo_point(_geo_point []byte) {
	t._geo_point = _geo_point
}

func (t *TL_inputMediaVenue) Get_geo_point() []byte {
	return t._geo_point
}

func (t *TL_inputMediaVenue) Set_title(_title string) {
	t._title = _title
}

func (t *TL_inputMediaVenue) Get_title() string {
	return t._title
}

func (t *TL_inputMediaVenue) Set_address(_address string) {
	t._address = _address
}

func (t *TL_inputMediaVenue) Get_address() string {
	return t._address
}

func (t *TL_inputMediaVenue) Set_provider(_provider string) {
	t._provider = _provider
}

func (t *TL_inputMediaVenue) Get_provider() string {
	return t._provider
}

func (t *TL_inputMediaVenue) Set_venue_id(_venue_id string) {
	t._venue_id = _venue_id
}

func (t *TL_inputMediaVenue) Get_venue_id() string {
	return t._venue_id
}

func (t *TL_inputMediaVenue) Set_venue_type(_venue_type string) {
	t._venue_type = _venue_type
}

func (t *TL_inputMediaVenue) Get_venue_type() string {
	return t._venue_type
}

func New_TL_inputMediaVenue() *TL_inputMediaVenue {
	return &TL_inputMediaVenue{}
}

func (t *TL_inputMediaVenue) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputMediaVenue))
	ec.Bytes(t.Get_geo_point())
	ec.String(t.Get_title())
	ec.String(t.Get_address())
	ec.String(t.Get_provider())
	ec.String(t.Get_venue_id())
	ec.String(t.Get_venue_type())

	return ec.GetBuffer()
}

func (t *TL_inputMediaVenue) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._geo_point = dc.Bytes(16)
	t._title = dc.String()
	t._address = dc.String()
	t._provider = dc.String()
	t._venue_id = dc.String()
	t._venue_type = dc.String()

}

// inputMediaGifExternal#4843b0fd
type TL_inputMediaGifExternal struct {
	_url string
	_q   string
}

func (t *TL_inputMediaGifExternal) Set_url(_url string) {
	t._url = _url
}

func (t *TL_inputMediaGifExternal) Get_url() string {
	return t._url
}

func (t *TL_inputMediaGifExternal) Set_q(_q string) {
	t._q = _q
}

func (t *TL_inputMediaGifExternal) Get_q() string {
	return t._q
}

func New_TL_inputMediaGifExternal() *TL_inputMediaGifExternal {
	return &TL_inputMediaGifExternal{}
}

func (t *TL_inputMediaGifExternal) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputMediaGifExternal))
	ec.String(t.Get_url())
	ec.String(t.Get_q())

	return ec.GetBuffer()
}

func (t *TL_inputMediaGifExternal) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._url = dc.String()
	t._q = dc.String()

}

// inputMediaPhotoExternal#0922aec1
type TL_inputMediaPhotoExternal struct {
	_flags       []byte
	_url         string
	_caption     string
	_ttl_seconds []byte
}

func (t *TL_inputMediaPhotoExternal) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_inputMediaPhotoExternal) Get_flags() []byte {
	return t._flags
}

func (t *TL_inputMediaPhotoExternal) Set_url(_url string) {
	t._url = _url
}

func (t *TL_inputMediaPhotoExternal) Get_url() string {
	return t._url
}

func (t *TL_inputMediaPhotoExternal) Set_caption(_caption string) {
	t._caption = _caption
}

func (t *TL_inputMediaPhotoExternal) Get_caption() string {
	return t._caption
}

func (t *TL_inputMediaPhotoExternal) Set_ttl_seconds(_ttl_seconds []byte) {
	t._ttl_seconds = _ttl_seconds
}

func (t *TL_inputMediaPhotoExternal) Get_ttl_seconds() []byte {
	return t._ttl_seconds
}

func New_TL_inputMediaPhotoExternal() *TL_inputMediaPhotoExternal {
	return &TL_inputMediaPhotoExternal{}
}

func (t *TL_inputMediaPhotoExternal) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputMediaPhotoExternal))
	ec.String(t.Get_url())
	ec.String(t.Get_caption())
	ec.Bytes(t.Get_ttl_seconds())

	return ec.GetBuffer()
}

func (t *TL_inputMediaPhotoExternal) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._url = dc.String()
	t._caption = dc.String()
	t._ttl_seconds = dc.Bytes(16)

}

// inputMediaDocumentExternal#b6f74335
type TL_inputMediaDocumentExternal struct {
	_flags       []byte
	_url         string
	_caption     string
	_ttl_seconds []byte
}

func (t *TL_inputMediaDocumentExternal) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_inputMediaDocumentExternal) Get_flags() []byte {
	return t._flags
}

func (t *TL_inputMediaDocumentExternal) Set_url(_url string) {
	t._url = _url
}

func (t *TL_inputMediaDocumentExternal) Get_url() string {
	return t._url
}

func (t *TL_inputMediaDocumentExternal) Set_caption(_caption string) {
	t._caption = _caption
}

func (t *TL_inputMediaDocumentExternal) Get_caption() string {
	return t._caption
}

func (t *TL_inputMediaDocumentExternal) Set_ttl_seconds(_ttl_seconds []byte) {
	t._ttl_seconds = _ttl_seconds
}

func (t *TL_inputMediaDocumentExternal) Get_ttl_seconds() []byte {
	return t._ttl_seconds
}

func New_TL_inputMediaDocumentExternal() *TL_inputMediaDocumentExternal {
	return &TL_inputMediaDocumentExternal{}
}

func (t *TL_inputMediaDocumentExternal) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputMediaDocumentExternal))
	ec.String(t.Get_url())
	ec.String(t.Get_caption())
	ec.Bytes(t.Get_ttl_seconds())

	return ec.GetBuffer()
}

func (t *TL_inputMediaDocumentExternal) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._url = dc.String()
	t._caption = dc.String()
	t._ttl_seconds = dc.Bytes(16)

}

// inputMediaGame#d33f43f3
type TL_inputMediaGame struct {
	_id []byte
}

func (t *TL_inputMediaGame) Set_id(_id []byte) {
	t._id = _id
}

func (t *TL_inputMediaGame) Get_id() []byte {
	return t._id
}

func New_TL_inputMediaGame() *TL_inputMediaGame {
	return &TL_inputMediaGame{}
}

func (t *TL_inputMediaGame) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputMediaGame))
	ec.Bytes(t.Get_id())

	return ec.GetBuffer()
}

func (t *TL_inputMediaGame) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Bytes(16)

}

// inputMediaInvoice#f4e096c3
type TL_inputMediaInvoice struct {
	_flags         []byte
	_title         string
	_description   string
	_photo         []byte
	_invoice       []byte
	_payload       []byte
	_provider      string
	_provider_data []byte
	_start_param   string
}

func (t *TL_inputMediaInvoice) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_inputMediaInvoice) Get_flags() []byte {
	return t._flags
}

func (t *TL_inputMediaInvoice) Set_title(_title string) {
	t._title = _title
}

func (t *TL_inputMediaInvoice) Get_title() string {
	return t._title
}

func (t *TL_inputMediaInvoice) Set_description(_description string) {
	t._description = _description
}

func (t *TL_inputMediaInvoice) Get_description() string {
	return t._description
}

func (t *TL_inputMediaInvoice) Set_photo(_photo []byte) {
	t._photo = _photo
}

func (t *TL_inputMediaInvoice) Get_photo() []byte {
	return t._photo
}

func (t *TL_inputMediaInvoice) Set_invoice(_invoice []byte) {
	t._invoice = _invoice
}

func (t *TL_inputMediaInvoice) Get_invoice() []byte {
	return t._invoice
}

func (t *TL_inputMediaInvoice) Set_payload(_payload []byte) {
	t._payload = _payload
}

func (t *TL_inputMediaInvoice) Get_payload() []byte {
	return t._payload
}

func (t *TL_inputMediaInvoice) Set_provider(_provider string) {
	t._provider = _provider
}

func (t *TL_inputMediaInvoice) Get_provider() string {
	return t._provider
}

func (t *TL_inputMediaInvoice) Set_provider_data(_provider_data []byte) {
	t._provider_data = _provider_data
}

func (t *TL_inputMediaInvoice) Get_provider_data() []byte {
	return t._provider_data
}

func (t *TL_inputMediaInvoice) Set_start_param(_start_param string) {
	t._start_param = _start_param
}

func (t *TL_inputMediaInvoice) Get_start_param() string {
	return t._start_param
}

func New_TL_inputMediaInvoice() *TL_inputMediaInvoice {
	return &TL_inputMediaInvoice{}
}

func (t *TL_inputMediaInvoice) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputMediaInvoice))
	ec.String(t.Get_title())
	ec.String(t.Get_description())
	ec.Bytes(t.Get_photo())
	ec.Bytes(t.Get_invoice())
	ec.Bytes(t.Get_payload())
	ec.String(t.Get_provider())
	ec.Bytes(t.Get_provider_data())
	ec.String(t.Get_start_param())

	return ec.GetBuffer()
}

func (t *TL_inputMediaInvoice) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._title = dc.String()
	t._description = dc.String()
	t._photo = dc.Bytes(16)
	t._invoice = dc.Bytes(16)
	t._payload = dc.Bytes(16)
	t._provider = dc.String()
	t._provider_data = dc.Bytes(16)
	t._start_param = dc.String()

}

// inputMediaGeoLive#7b1a118f
type TL_inputMediaGeoLive struct {
	_geo_point []byte
	_period    int32
}

func (t *TL_inputMediaGeoLive) Set_geo_point(_geo_point []byte) {
	t._geo_point = _geo_point
}

func (t *TL_inputMediaGeoLive) Get_geo_point() []byte {
	return t._geo_point
}

func (t *TL_inputMediaGeoLive) Set_period(_period int32) {
	t._period = _period
}

func (t *TL_inputMediaGeoLive) Get_period() int32 {
	return t._period
}

func New_TL_inputMediaGeoLive() *TL_inputMediaGeoLive {
	return &TL_inputMediaGeoLive{}
}

func (t *TL_inputMediaGeoLive) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputMediaGeoLive))
	ec.Bytes(t.Get_geo_point())
	ec.Int(t.Get_period())

	return ec.GetBuffer()
}

func (t *TL_inputMediaGeoLive) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._geo_point = dc.Bytes(16)
	t._period = dc.Int()

}

// inputChatPhotoEmpty#1ca48f57
type TL_inputChatPhotoEmpty struct {
}

func New_TL_inputChatPhotoEmpty() *TL_inputChatPhotoEmpty {
	return &TL_inputChatPhotoEmpty{}
}

func (t *TL_inputChatPhotoEmpty) Encode() []byte {
	return nil
}

func (t *TL_inputChatPhotoEmpty) Decode(b []byte) {

}

// inputChatUploadedPhoto#927c55b4
type TL_inputChatUploadedPhoto struct {
	_file []byte
}

func (t *TL_inputChatUploadedPhoto) Set_file(_file []byte) {
	t._file = _file
}

func (t *TL_inputChatUploadedPhoto) Get_file() []byte {
	return t._file
}

func New_TL_inputChatUploadedPhoto() *TL_inputChatUploadedPhoto {
	return &TL_inputChatUploadedPhoto{}
}

func (t *TL_inputChatUploadedPhoto) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputChatUploadedPhoto))
	ec.Bytes(t.Get_file())

	return ec.GetBuffer()
}

func (t *TL_inputChatUploadedPhoto) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._file = dc.Bytes(16)

}

// inputChatPhoto#8953ad37
type TL_inputChatPhoto struct {
	_id []byte
}

func (t *TL_inputChatPhoto) Set_id(_id []byte) {
	t._id = _id
}

func (t *TL_inputChatPhoto) Get_id() []byte {
	return t._id
}

func New_TL_inputChatPhoto() *TL_inputChatPhoto {
	return &TL_inputChatPhoto{}
}

func (t *TL_inputChatPhoto) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputChatPhoto))
	ec.Bytes(t.Get_id())

	return ec.GetBuffer()
}

func (t *TL_inputChatPhoto) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Bytes(16)

}

// inputGeoPointEmpty#e4c123d6
type TL_inputGeoPointEmpty struct {
}

func New_TL_inputGeoPointEmpty() *TL_inputGeoPointEmpty {
	return &TL_inputGeoPointEmpty{}
}

func (t *TL_inputGeoPointEmpty) Encode() []byte {
	return nil
}

func (t *TL_inputGeoPointEmpty) Decode(b []byte) {

}

// inputGeoPoint#f3b7acc9
type TL_inputGeoPoint struct {
	_lat  float64
	_long float64
}

func (t *TL_inputGeoPoint) Set_lat(_lat float64) {
	t._lat = _lat
}

func (t *TL_inputGeoPoint) Get_lat() float64 {
	return t._lat
}

func (t *TL_inputGeoPoint) Set_long(_long float64) {
	t._long = _long
}

func (t *TL_inputGeoPoint) Get_long() float64 {
	return t._long
}

func New_TL_inputGeoPoint() *TL_inputGeoPoint {
	return &TL_inputGeoPoint{}
}

func (t *TL_inputGeoPoint) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputGeoPoint))
	ec.Double(t.Get_lat())
	ec.Double(t.Get_long())

	return ec.GetBuffer()
}

func (t *TL_inputGeoPoint) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._lat = dc.Double()
	t._long = dc.Double()

}

// inputPhotoEmpty#1cd7bf0d
type TL_inputPhotoEmpty struct {
}

func New_TL_inputPhotoEmpty() *TL_inputPhotoEmpty {
	return &TL_inputPhotoEmpty{}
}

func (t *TL_inputPhotoEmpty) Encode() []byte {
	return nil
}

func (t *TL_inputPhotoEmpty) Decode(b []byte) {

}

// inputPhoto#fb95c6c4
type TL_inputPhoto struct {
	_id          int64
	_access_hash int64
}

func (t *TL_inputPhoto) Set_id(_id int64) {
	t._id = _id
}

func (t *TL_inputPhoto) Get_id() int64 {
	return t._id
}

func (t *TL_inputPhoto) Set_access_hash(_access_hash int64) {
	t._access_hash = _access_hash
}

func (t *TL_inputPhoto) Get_access_hash() int64 {
	return t._access_hash
}

func New_TL_inputPhoto() *TL_inputPhoto {
	return &TL_inputPhoto{}
}

func (t *TL_inputPhoto) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputPhoto))
	ec.Long(t.Get_id())
	ec.Long(t.Get_access_hash())

	return ec.GetBuffer()
}

func (t *TL_inputPhoto) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Long()
	t._access_hash = dc.Long()

}

// inputFileLocation#14637196
type TL_inputFileLocation struct {
	_volume_id int64
	_local_id  int32
	_secret    int64
}

func (t *TL_inputFileLocation) Set_volume_id(_volume_id int64) {
	t._volume_id = _volume_id
}

func (t *TL_inputFileLocation) Get_volume_id() int64 {
	return t._volume_id
}

func (t *TL_inputFileLocation) Set_local_id(_local_id int32) {
	t._local_id = _local_id
}

func (t *TL_inputFileLocation) Get_local_id() int32 {
	return t._local_id
}

func (t *TL_inputFileLocation) Set_secret(_secret int64) {
	t._secret = _secret
}

func (t *TL_inputFileLocation) Get_secret() int64 {
	return t._secret
}

func New_TL_inputFileLocation() *TL_inputFileLocation {
	return &TL_inputFileLocation{}
}

func (t *TL_inputFileLocation) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputFileLocation))
	ec.Long(t.Get_volume_id())
	ec.Int(t.Get_local_id())
	ec.Long(t.Get_secret())

	return ec.GetBuffer()
}

func (t *TL_inputFileLocation) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._volume_id = dc.Long()
	t._local_id = dc.Int()
	t._secret = dc.Long()

}

// inputEncryptedFileLocation#f5235d55
type TL_inputEncryptedFileLocation struct {
	_id          int64
	_access_hash int64
}

func (t *TL_inputEncryptedFileLocation) Set_id(_id int64) {
	t._id = _id
}

func (t *TL_inputEncryptedFileLocation) Get_id() int64 {
	return t._id
}

func (t *TL_inputEncryptedFileLocation) Set_access_hash(_access_hash int64) {
	t._access_hash = _access_hash
}

func (t *TL_inputEncryptedFileLocation) Get_access_hash() int64 {
	return t._access_hash
}

func New_TL_inputEncryptedFileLocation() *TL_inputEncryptedFileLocation {
	return &TL_inputEncryptedFileLocation{}
}

func (t *TL_inputEncryptedFileLocation) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputEncryptedFileLocation))
	ec.Long(t.Get_id())
	ec.Long(t.Get_access_hash())

	return ec.GetBuffer()
}

func (t *TL_inputEncryptedFileLocation) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Long()
	t._access_hash = dc.Long()

}

// inputDocumentFileLocation#430f0724
type TL_inputDocumentFileLocation struct {
	_id          int64
	_access_hash int64
	_version     int32
}

func (t *TL_inputDocumentFileLocation) Set_id(_id int64) {
	t._id = _id
}

func (t *TL_inputDocumentFileLocation) Get_id() int64 {
	return t._id
}

func (t *TL_inputDocumentFileLocation) Set_access_hash(_access_hash int64) {
	t._access_hash = _access_hash
}

func (t *TL_inputDocumentFileLocation) Get_access_hash() int64 {
	return t._access_hash
}

func (t *TL_inputDocumentFileLocation) Set_version(_version int32) {
	t._version = _version
}

func (t *TL_inputDocumentFileLocation) Get_version() int32 {
	return t._version
}

func New_TL_inputDocumentFileLocation() *TL_inputDocumentFileLocation {
	return &TL_inputDocumentFileLocation{}
}

func (t *TL_inputDocumentFileLocation) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputDocumentFileLocation))
	ec.Long(t.Get_id())
	ec.Long(t.Get_access_hash())
	ec.Int(t.Get_version())

	return ec.GetBuffer()
}

func (t *TL_inputDocumentFileLocation) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Long()
	t._access_hash = dc.Long()
	t._version = dc.Int()

}

// inputAppEvent#770656a8
type TL_inputAppEvent struct {
	_time float64
	_type string
	_peer int64
	_data string
}

func (t *TL_inputAppEvent) Set_time(_time float64) {
	t._time = _time
}

func (t *TL_inputAppEvent) Get_time() float64 {
	return t._time
}

func (t *TL_inputAppEvent) Set_type(_type string) {
	t._type = _type
}

func (t *TL_inputAppEvent) Get_type() string {
	return t._type
}

func (t *TL_inputAppEvent) Set_peer(_peer int64) {
	t._peer = _peer
}

func (t *TL_inputAppEvent) Get_peer() int64 {
	return t._peer
}

func (t *TL_inputAppEvent) Set_data(_data string) {
	t._data = _data
}

func (t *TL_inputAppEvent) Get_data() string {
	return t._data
}

func New_TL_inputAppEvent() *TL_inputAppEvent {
	return &TL_inputAppEvent{}
}

func (t *TL_inputAppEvent) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputAppEvent))
	ec.Double(t.Get_time())
	ec.String(t.Get_type())
	ec.Long(t.Get_peer())
	ec.String(t.Get_data())

	return ec.GetBuffer()
}

func (t *TL_inputAppEvent) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._time = dc.Double()
	t._type = dc.String()
	t._peer = dc.Long()
	t._data = dc.String()

}

// peerUser#9db1bc6d
type TL_peerUser struct {
	_user_id int32
}

func (t *TL_peerUser) Set_user_id(_user_id int32) {
	t._user_id = _user_id
}

func (t *TL_peerUser) Get_user_id() int32 {
	return t._user_id
}

func New_TL_peerUser() *TL_peerUser {
	return &TL_peerUser{}
}

func (t *TL_peerUser) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_peerUser))
	ec.Int(t.Get_user_id())

	return ec.GetBuffer()
}

func (t *TL_peerUser) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._user_id = dc.Int()

}

// peerChat#bad0e5bb
type TL_peerChat struct {
	_chat_id int32
}

func (t *TL_peerChat) Set_chat_id(_chat_id int32) {
	t._chat_id = _chat_id
}

func (t *TL_peerChat) Get_chat_id() int32 {
	return t._chat_id
}

func New_TL_peerChat() *TL_peerChat {
	return &TL_peerChat{}
}

func (t *TL_peerChat) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_peerChat))
	ec.Int(t.Get_chat_id())

	return ec.GetBuffer()
}

func (t *TL_peerChat) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._chat_id = dc.Int()

}

// peerChannel#bddde532
type TL_peerChannel struct {
	_channel_id int32
}

func (t *TL_peerChannel) Set_channel_id(_channel_id int32) {
	t._channel_id = _channel_id
}

func (t *TL_peerChannel) Get_channel_id() int32 {
	return t._channel_id
}

func New_TL_peerChannel() *TL_peerChannel {
	return &TL_peerChannel{}
}

func (t *TL_peerChannel) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_peerChannel))
	ec.Int(t.Get_channel_id())

	return ec.GetBuffer()
}

func (t *TL_peerChannel) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._channel_id = dc.Int()

}

// storage_fileUnknown#aa963b05
type TL_storage_fileUnknown struct {
}

func New_TL_storage_fileUnknown() *TL_storage_fileUnknown {
	return &TL_storage_fileUnknown{}
}

func (t *TL_storage_fileUnknown) Encode() []byte {
	return nil
}

func (t *TL_storage_fileUnknown) Decode(b []byte) {

}

// storage_filePartial#40bc6f52
type TL_storage_filePartial struct {
}

func New_TL_storage_filePartial() *TL_storage_filePartial {
	return &TL_storage_filePartial{}
}

func (t *TL_storage_filePartial) Encode() []byte {
	return nil
}

func (t *TL_storage_filePartial) Decode(b []byte) {

}

// storage_fileJpeg#007efe0e
type TL_storage_fileJpeg struct {
}

func New_TL_storage_fileJpeg() *TL_storage_fileJpeg {
	return &TL_storage_fileJpeg{}
}

func (t *TL_storage_fileJpeg) Encode() []byte {
	return nil
}

func (t *TL_storage_fileJpeg) Decode(b []byte) {

}

// storage_fileGif#cae1aadf
type TL_storage_fileGif struct {
}

func New_TL_storage_fileGif() *TL_storage_fileGif {
	return &TL_storage_fileGif{}
}

func (t *TL_storage_fileGif) Encode() []byte {
	return nil
}

func (t *TL_storage_fileGif) Decode(b []byte) {

}

// storage_filePng#0a4f63c0
type TL_storage_filePng struct {
}

func New_TL_storage_filePng() *TL_storage_filePng {
	return &TL_storage_filePng{}
}

func (t *TL_storage_filePng) Encode() []byte {
	return nil
}

func (t *TL_storage_filePng) Decode(b []byte) {

}

// storage_filePdf#ae1e508d
type TL_storage_filePdf struct {
}

func New_TL_storage_filePdf() *TL_storage_filePdf {
	return &TL_storage_filePdf{}
}

func (t *TL_storage_filePdf) Encode() []byte {
	return nil
}

func (t *TL_storage_filePdf) Decode(b []byte) {

}

// storage_fileMp3#528a0677
type TL_storage_fileMp3 struct {
}

func New_TL_storage_fileMp3() *TL_storage_fileMp3 {
	return &TL_storage_fileMp3{}
}

func (t *TL_storage_fileMp3) Encode() []byte {
	return nil
}

func (t *TL_storage_fileMp3) Decode(b []byte) {

}

// storage_fileMov#4b09ebbc
type TL_storage_fileMov struct {
}

func New_TL_storage_fileMov() *TL_storage_fileMov {
	return &TL_storage_fileMov{}
}

func (t *TL_storage_fileMov) Encode() []byte {
	return nil
}

func (t *TL_storage_fileMov) Decode(b []byte) {

}

// storage_fileMp4#b3cea0e4
type TL_storage_fileMp4 struct {
}

func New_TL_storage_fileMp4() *TL_storage_fileMp4 {
	return &TL_storage_fileMp4{}
}

func (t *TL_storage_fileMp4) Encode() []byte {
	return nil
}

func (t *TL_storage_fileMp4) Decode(b []byte) {

}

// storage_fileWebp#1081464c
type TL_storage_fileWebp struct {
}

func New_TL_storage_fileWebp() *TL_storage_fileWebp {
	return &TL_storage_fileWebp{}
}

func (t *TL_storage_fileWebp) Encode() []byte {
	return nil
}

func (t *TL_storage_fileWebp) Decode(b []byte) {

}

// fileLocationUnavailable#7c596b46
type TL_fileLocationUnavailable struct {
	_volume_id int64
	_local_id  int32
	_secret    int64
}

func (t *TL_fileLocationUnavailable) Set_volume_id(_volume_id int64) {
	t._volume_id = _volume_id
}

func (t *TL_fileLocationUnavailable) Get_volume_id() int64 {
	return t._volume_id
}

func (t *TL_fileLocationUnavailable) Set_local_id(_local_id int32) {
	t._local_id = _local_id
}

func (t *TL_fileLocationUnavailable) Get_local_id() int32 {
	return t._local_id
}

func (t *TL_fileLocationUnavailable) Set_secret(_secret int64) {
	t._secret = _secret
}

func (t *TL_fileLocationUnavailable) Get_secret() int64 {
	return t._secret
}

func New_TL_fileLocationUnavailable() *TL_fileLocationUnavailable {
	return &TL_fileLocationUnavailable{}
}

func (t *TL_fileLocationUnavailable) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_fileLocationUnavailable))
	ec.Long(t.Get_volume_id())
	ec.Int(t.Get_local_id())
	ec.Long(t.Get_secret())

	return ec.GetBuffer()
}

func (t *TL_fileLocationUnavailable) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._volume_id = dc.Long()
	t._local_id = dc.Int()
	t._secret = dc.Long()

}

// fileLocation#53d69076
type TL_fileLocation struct {
	_dc_id     int32
	_volume_id int64
	_local_id  int32
	_secret    int64
}

func (t *TL_fileLocation) Set_dc_id(_dc_id int32) {
	t._dc_id = _dc_id
}

func (t *TL_fileLocation) Get_dc_id() int32 {
	return t._dc_id
}

func (t *TL_fileLocation) Set_volume_id(_volume_id int64) {
	t._volume_id = _volume_id
}

func (t *TL_fileLocation) Get_volume_id() int64 {
	return t._volume_id
}

func (t *TL_fileLocation) Set_local_id(_local_id int32) {
	t._local_id = _local_id
}

func (t *TL_fileLocation) Get_local_id() int32 {
	return t._local_id
}

func (t *TL_fileLocation) Set_secret(_secret int64) {
	t._secret = _secret
}

func (t *TL_fileLocation) Get_secret() int64 {
	return t._secret
}

func New_TL_fileLocation() *TL_fileLocation {
	return &TL_fileLocation{}
}

func (t *TL_fileLocation) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_fileLocation))
	ec.Int(t.Get_dc_id())
	ec.Long(t.Get_volume_id())
	ec.Int(t.Get_local_id())
	ec.Long(t.Get_secret())

	return ec.GetBuffer()
}

func (t *TL_fileLocation) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._dc_id = dc.Int()
	t._volume_id = dc.Long()
	t._local_id = dc.Int()
	t._secret = dc.Long()

}

// userEmpty#200250ba
type TL_userEmpty struct {
	_id int32
}

func (t *TL_userEmpty) Set_id(_id int32) {
	t._id = _id
}

func (t *TL_userEmpty) Get_id() int32 {
	return t._id
}

func New_TL_userEmpty() *TL_userEmpty {
	return &TL_userEmpty{}
}

func (t *TL_userEmpty) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_userEmpty))
	ec.Int(t.Get_id())

	return ec.GetBuffer()
}

func (t *TL_userEmpty) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Int()

}

// user#2e13f4c3
type TL_user struct {
	_flags                  []byte
	_self                   []byte
	_contact                []byte
	_mutual_contact         []byte
	_deleted                []byte
	_bot                    []byte
	_bot_chat_history       []byte
	_bot_nochats            []byte
	_verified               []byte
	_restricted             []byte
	_min                    []byte
	_bot_inline_geo         []byte
	_id                     int32
	_access_hash            []byte
	_first_name             []byte
	_last_name              []byte
	_username               []byte
	_phone                  []byte
	_photo                  []byte
	_status                 []byte
	_bot_info_version       []byte
	_restriction_reason     []byte
	_bot_inline_placeholder []byte
	_lang_code              []byte
}

func (t *TL_user) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_user) Get_flags() []byte {
	return t._flags
}

func (t *TL_user) Set_self(_self []byte) {
	t._self = _self
}

func (t *TL_user) Get_self() []byte {
	return t._self
}

func (t *TL_user) Set_contact(_contact []byte) {
	t._contact = _contact
}

func (t *TL_user) Get_contact() []byte {
	return t._contact
}

func (t *TL_user) Set_mutual_contact(_mutual_contact []byte) {
	t._mutual_contact = _mutual_contact
}

func (t *TL_user) Get_mutual_contact() []byte {
	return t._mutual_contact
}

func (t *TL_user) Set_deleted(_deleted []byte) {
	t._deleted = _deleted
}

func (t *TL_user) Get_deleted() []byte {
	return t._deleted
}

func (t *TL_user) Set_bot(_bot []byte) {
	t._bot = _bot
}

func (t *TL_user) Get_bot() []byte {
	return t._bot
}

func (t *TL_user) Set_bot_chat_history(_bot_chat_history []byte) {
	t._bot_chat_history = _bot_chat_history
}

func (t *TL_user) Get_bot_chat_history() []byte {
	return t._bot_chat_history
}

func (t *TL_user) Set_bot_nochats(_bot_nochats []byte) {
	t._bot_nochats = _bot_nochats
}

func (t *TL_user) Get_bot_nochats() []byte {
	return t._bot_nochats
}

func (t *TL_user) Set_verified(_verified []byte) {
	t._verified = _verified
}

func (t *TL_user) Get_verified() []byte {
	return t._verified
}

func (t *TL_user) Set_restricted(_restricted []byte) {
	t._restricted = _restricted
}

func (t *TL_user) Get_restricted() []byte {
	return t._restricted
}

func (t *TL_user) Set_min(_min []byte) {
	t._min = _min
}

func (t *TL_user) Get_min() []byte {
	return t._min
}

func (t *TL_user) Set_bot_inline_geo(_bot_inline_geo []byte) {
	t._bot_inline_geo = _bot_inline_geo
}

func (t *TL_user) Get_bot_inline_geo() []byte {
	return t._bot_inline_geo
}

func (t *TL_user) Set_id(_id int32) {
	t._id = _id
}

func (t *TL_user) Get_id() int32 {
	return t._id
}

func (t *TL_user) Set_access_hash(_access_hash []byte) {
	t._access_hash = _access_hash
}

func (t *TL_user) Get_access_hash() []byte {
	return t._access_hash
}

func (t *TL_user) Set_first_name(_first_name []byte) {
	t._first_name = _first_name
}

func (t *TL_user) Get_first_name() []byte {
	return t._first_name
}

func (t *TL_user) Set_last_name(_last_name []byte) {
	t._last_name = _last_name
}

func (t *TL_user) Get_last_name() []byte {
	return t._last_name
}

func (t *TL_user) Set_username(_username []byte) {
	t._username = _username
}

func (t *TL_user) Get_username() []byte {
	return t._username
}

func (t *TL_user) Set_phone(_phone []byte) {
	t._phone = _phone
}

func (t *TL_user) Get_phone() []byte {
	return t._phone
}

func (t *TL_user) Set_photo(_photo []byte) {
	t._photo = _photo
}

func (t *TL_user) Get_photo() []byte {
	return t._photo
}

func (t *TL_user) Set_status(_status []byte) {
	t._status = _status
}

func (t *TL_user) Get_status() []byte {
	return t._status
}

func (t *TL_user) Set_bot_info_version(_bot_info_version []byte) {
	t._bot_info_version = _bot_info_version
}

func (t *TL_user) Get_bot_info_version() []byte {
	return t._bot_info_version
}

func (t *TL_user) Set_restriction_reason(_restriction_reason []byte) {
	t._restriction_reason = _restriction_reason
}

func (t *TL_user) Get_restriction_reason() []byte {
	return t._restriction_reason
}

func (t *TL_user) Set_bot_inline_placeholder(_bot_inline_placeholder []byte) {
	t._bot_inline_placeholder = _bot_inline_placeholder
}

func (t *TL_user) Get_bot_inline_placeholder() []byte {
	return t._bot_inline_placeholder
}

func (t *TL_user) Set_lang_code(_lang_code []byte) {
	t._lang_code = _lang_code
}

func (t *TL_user) Get_lang_code() []byte {
	return t._lang_code
}

func New_TL_user() *TL_user {
	return &TL_user{}
}

func (t *TL_user) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_user))
	ec.Bytes(t.Get_self())
	ec.Bytes(t.Get_contact())
	ec.Bytes(t.Get_mutual_contact())
	ec.Bytes(t.Get_deleted())
	ec.Bytes(t.Get_bot())
	ec.Bytes(t.Get_bot_chat_history())
	ec.Bytes(t.Get_bot_nochats())
	ec.Bytes(t.Get_verified())
	ec.Bytes(t.Get_restricted())
	ec.Bytes(t.Get_min())
	ec.Bytes(t.Get_bot_inline_geo())
	ec.Int(t.Get_id())
	ec.Bytes(t.Get_access_hash())
	ec.Bytes(t.Get_first_name())
	ec.Bytes(t.Get_last_name())
	ec.Bytes(t.Get_username())
	ec.Bytes(t.Get_phone())
	ec.Bytes(t.Get_photo())
	ec.Bytes(t.Get_status())
	ec.Bytes(t.Get_bot_info_version())
	ec.Bytes(t.Get_restriction_reason())
	ec.Bytes(t.Get_bot_inline_placeholder())
	ec.Bytes(t.Get_lang_code())

	return ec.GetBuffer()
}

func (t *TL_user) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._self = dc.Bytes(16)
	t._contact = dc.Bytes(16)
	t._mutual_contact = dc.Bytes(16)
	t._deleted = dc.Bytes(16)
	t._bot = dc.Bytes(16)
	t._bot_chat_history = dc.Bytes(16)
	t._bot_nochats = dc.Bytes(16)
	t._verified = dc.Bytes(16)
	t._restricted = dc.Bytes(16)
	t._min = dc.Bytes(16)
	t._bot_inline_geo = dc.Bytes(16)
	t._id = dc.Int()
	t._access_hash = dc.Bytes(16)
	t._first_name = dc.Bytes(16)
	t._last_name = dc.Bytes(16)
	t._username = dc.Bytes(16)
	t._phone = dc.Bytes(16)
	t._photo = dc.Bytes(16)
	t._status = dc.Bytes(16)
	t._bot_info_version = dc.Bytes(16)
	t._restriction_reason = dc.Bytes(16)
	t._bot_inline_placeholder = dc.Bytes(16)
	t._lang_code = dc.Bytes(16)

}

// userProfilePhotoEmpty#4f11bae1
type TL_userProfilePhotoEmpty struct {
}

func New_TL_userProfilePhotoEmpty() *TL_userProfilePhotoEmpty {
	return &TL_userProfilePhotoEmpty{}
}

func (t *TL_userProfilePhotoEmpty) Encode() []byte {
	return nil
}

func (t *TL_userProfilePhotoEmpty) Decode(b []byte) {

}

// userProfilePhoto#d559d8c8
type TL_userProfilePhoto struct {
	_photo_id    int64
	_photo_small []byte
	_photo_big   []byte
}

func (t *TL_userProfilePhoto) Set_photo_id(_photo_id int64) {
	t._photo_id = _photo_id
}

func (t *TL_userProfilePhoto) Get_photo_id() int64 {
	return t._photo_id
}

func (t *TL_userProfilePhoto) Set_photo_small(_photo_small []byte) {
	t._photo_small = _photo_small
}

func (t *TL_userProfilePhoto) Get_photo_small() []byte {
	return t._photo_small
}

func (t *TL_userProfilePhoto) Set_photo_big(_photo_big []byte) {
	t._photo_big = _photo_big
}

func (t *TL_userProfilePhoto) Get_photo_big() []byte {
	return t._photo_big
}

func New_TL_userProfilePhoto() *TL_userProfilePhoto {
	return &TL_userProfilePhoto{}
}

func (t *TL_userProfilePhoto) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_userProfilePhoto))
	ec.Long(t.Get_photo_id())
	ec.Bytes(t.Get_photo_small())
	ec.Bytes(t.Get_photo_big())

	return ec.GetBuffer()
}

func (t *TL_userProfilePhoto) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._photo_id = dc.Long()
	t._photo_small = dc.Bytes(16)
	t._photo_big = dc.Bytes(16)

}

// userStatusEmpty#09d05049
type TL_userStatusEmpty struct {
}

func New_TL_userStatusEmpty() *TL_userStatusEmpty {
	return &TL_userStatusEmpty{}
}

func (t *TL_userStatusEmpty) Encode() []byte {
	return nil
}

func (t *TL_userStatusEmpty) Decode(b []byte) {

}

// userStatusOnline#edb93949
type TL_userStatusOnline struct {
	_expires int32
}

func (t *TL_userStatusOnline) Set_expires(_expires int32) {
	t._expires = _expires
}

func (t *TL_userStatusOnline) Get_expires() int32 {
	return t._expires
}

func New_TL_userStatusOnline() *TL_userStatusOnline {
	return &TL_userStatusOnline{}
}

func (t *TL_userStatusOnline) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_userStatusOnline))
	ec.Int(t.Get_expires())

	return ec.GetBuffer()
}

func (t *TL_userStatusOnline) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._expires = dc.Int()

}

// userStatusOffline#008c703f
type TL_userStatusOffline struct {
	_was_online int32
}

func (t *TL_userStatusOffline) Set_was_online(_was_online int32) {
	t._was_online = _was_online
}

func (t *TL_userStatusOffline) Get_was_online() int32 {
	return t._was_online
}

func New_TL_userStatusOffline() *TL_userStatusOffline {
	return &TL_userStatusOffline{}
}

func (t *TL_userStatusOffline) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_userStatusOffline))
	ec.Int(t.Get_was_online())

	return ec.GetBuffer()
}

func (t *TL_userStatusOffline) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._was_online = dc.Int()

}

// userStatusRecently#e26f42f1
type TL_userStatusRecently struct {
}

func New_TL_userStatusRecently() *TL_userStatusRecently {
	return &TL_userStatusRecently{}
}

func (t *TL_userStatusRecently) Encode() []byte {
	return nil
}

func (t *TL_userStatusRecently) Decode(b []byte) {

}

// userStatusLastWeek#07bf09fc
type TL_userStatusLastWeek struct {
}

func New_TL_userStatusLastWeek() *TL_userStatusLastWeek {
	return &TL_userStatusLastWeek{}
}

func (t *TL_userStatusLastWeek) Encode() []byte {
	return nil
}

func (t *TL_userStatusLastWeek) Decode(b []byte) {

}

// userStatusLastMonth#77ebc742
type TL_userStatusLastMonth struct {
}

func New_TL_userStatusLastMonth() *TL_userStatusLastMonth {
	return &TL_userStatusLastMonth{}
}

func (t *TL_userStatusLastMonth) Encode() []byte {
	return nil
}

func (t *TL_userStatusLastMonth) Decode(b []byte) {

}

// chatEmpty#9ba2d800
type TL_chatEmpty struct {
	_id int32
}

func (t *TL_chatEmpty) Set_id(_id int32) {
	t._id = _id
}

func (t *TL_chatEmpty) Get_id() int32 {
	return t._id
}

func New_TL_chatEmpty() *TL_chatEmpty {
	return &TL_chatEmpty{}
}

func (t *TL_chatEmpty) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_chatEmpty))
	ec.Int(t.Get_id())

	return ec.GetBuffer()
}

func (t *TL_chatEmpty) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Int()

}

// chat#d91cdd54
type TL_chat struct {
	_flags              []byte
	_creator            []byte
	_kicked             []byte
	_left               []byte
	_admins_enabled     []byte
	_admin              []byte
	_deactivated        []byte
	_id                 int32
	_title              string
	_photo              []byte
	_participants_count int32
	_date               int32
	_version            int32
	_migrated_to        []byte
}

func (t *TL_chat) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_chat) Get_flags() []byte {
	return t._flags
}

func (t *TL_chat) Set_creator(_creator []byte) {
	t._creator = _creator
}

func (t *TL_chat) Get_creator() []byte {
	return t._creator
}

func (t *TL_chat) Set_kicked(_kicked []byte) {
	t._kicked = _kicked
}

func (t *TL_chat) Get_kicked() []byte {
	return t._kicked
}

func (t *TL_chat) Set_left(_left []byte) {
	t._left = _left
}

func (t *TL_chat) Get_left() []byte {
	return t._left
}

func (t *TL_chat) Set_admins_enabled(_admins_enabled []byte) {
	t._admins_enabled = _admins_enabled
}

func (t *TL_chat) Get_admins_enabled() []byte {
	return t._admins_enabled
}

func (t *TL_chat) Set_admin(_admin []byte) {
	t._admin = _admin
}

func (t *TL_chat) Get_admin() []byte {
	return t._admin
}

func (t *TL_chat) Set_deactivated(_deactivated []byte) {
	t._deactivated = _deactivated
}

func (t *TL_chat) Get_deactivated() []byte {
	return t._deactivated
}

func (t *TL_chat) Set_id(_id int32) {
	t._id = _id
}

func (t *TL_chat) Get_id() int32 {
	return t._id
}

func (t *TL_chat) Set_title(_title string) {
	t._title = _title
}

func (t *TL_chat) Get_title() string {
	return t._title
}

func (t *TL_chat) Set_photo(_photo []byte) {
	t._photo = _photo
}

func (t *TL_chat) Get_photo() []byte {
	return t._photo
}

func (t *TL_chat) Set_participants_count(_participants_count int32) {
	t._participants_count = _participants_count
}

func (t *TL_chat) Get_participants_count() int32 {
	return t._participants_count
}

func (t *TL_chat) Set_date(_date int32) {
	t._date = _date
}

func (t *TL_chat) Get_date() int32 {
	return t._date
}

func (t *TL_chat) Set_version(_version int32) {
	t._version = _version
}

func (t *TL_chat) Get_version() int32 {
	return t._version
}

func (t *TL_chat) Set_migrated_to(_migrated_to []byte) {
	t._migrated_to = _migrated_to
}

func (t *TL_chat) Get_migrated_to() []byte {
	return t._migrated_to
}

func New_TL_chat() *TL_chat {
	return &TL_chat{}
}

func (t *TL_chat) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_chat))
	ec.Bytes(t.Get_creator())
	ec.Bytes(t.Get_kicked())
	ec.Bytes(t.Get_left())
	ec.Bytes(t.Get_admins_enabled())
	ec.Bytes(t.Get_admin())
	ec.Bytes(t.Get_deactivated())
	ec.Int(t.Get_id())
	ec.String(t.Get_title())
	ec.Bytes(t.Get_photo())
	ec.Int(t.Get_participants_count())
	ec.Int(t.Get_date())
	ec.Int(t.Get_version())
	ec.Bytes(t.Get_migrated_to())

	return ec.GetBuffer()
}

func (t *TL_chat) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._creator = dc.Bytes(16)
	t._kicked = dc.Bytes(16)
	t._left = dc.Bytes(16)
	t._admins_enabled = dc.Bytes(16)
	t._admin = dc.Bytes(16)
	t._deactivated = dc.Bytes(16)
	t._id = dc.Int()
	t._title = dc.String()
	t._photo = dc.Bytes(16)
	t._participants_count = dc.Int()
	t._date = dc.Int()
	t._version = dc.Int()
	t._migrated_to = dc.Bytes(16)

}

// chatForbidden#07328bdb
type TL_chatForbidden struct {
	_id    int32
	_title string
}

func (t *TL_chatForbidden) Set_id(_id int32) {
	t._id = _id
}

func (t *TL_chatForbidden) Get_id() int32 {
	return t._id
}

func (t *TL_chatForbidden) Set_title(_title string) {
	t._title = _title
}

func (t *TL_chatForbidden) Get_title() string {
	return t._title
}

func New_TL_chatForbidden() *TL_chatForbidden {
	return &TL_chatForbidden{}
}

func (t *TL_chatForbidden) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_chatForbidden))
	ec.Int(t.Get_id())
	ec.String(t.Get_title())

	return ec.GetBuffer()
}

func (t *TL_chatForbidden) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Int()
	t._title = dc.String()

}

// channel#450b7115
type TL_channel struct {
	_flags              []byte
	_creator            []byte
	_left               []byte
	_editor             []byte
	_broadcast          []byte
	_verified           []byte
	_megagroup          []byte
	_restricted         []byte
	_democracy          []byte
	_signatures         []byte
	_min                []byte
	_id                 int32
	_access_hash        []byte
	_title              string
	_username           []byte
	_photo              []byte
	_date               int32
	_version            int32
	_restriction_reason []byte
	_admin_rights       []byte
	_banned_rights      []byte
	_participants_count []byte
}

func (t *TL_channel) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_channel) Get_flags() []byte {
	return t._flags
}

func (t *TL_channel) Set_creator(_creator []byte) {
	t._creator = _creator
}

func (t *TL_channel) Get_creator() []byte {
	return t._creator
}

func (t *TL_channel) Set_left(_left []byte) {
	t._left = _left
}

func (t *TL_channel) Get_left() []byte {
	return t._left
}

func (t *TL_channel) Set_editor(_editor []byte) {
	t._editor = _editor
}

func (t *TL_channel) Get_editor() []byte {
	return t._editor
}

func (t *TL_channel) Set_broadcast(_broadcast []byte) {
	t._broadcast = _broadcast
}

func (t *TL_channel) Get_broadcast() []byte {
	return t._broadcast
}

func (t *TL_channel) Set_verified(_verified []byte) {
	t._verified = _verified
}

func (t *TL_channel) Get_verified() []byte {
	return t._verified
}

func (t *TL_channel) Set_megagroup(_megagroup []byte) {
	t._megagroup = _megagroup
}

func (t *TL_channel) Get_megagroup() []byte {
	return t._megagroup
}

func (t *TL_channel) Set_restricted(_restricted []byte) {
	t._restricted = _restricted
}

func (t *TL_channel) Get_restricted() []byte {
	return t._restricted
}

func (t *TL_channel) Set_democracy(_democracy []byte) {
	t._democracy = _democracy
}

func (t *TL_channel) Get_democracy() []byte {
	return t._democracy
}

func (t *TL_channel) Set_signatures(_signatures []byte) {
	t._signatures = _signatures
}

func (t *TL_channel) Get_signatures() []byte {
	return t._signatures
}

func (t *TL_channel) Set_min(_min []byte) {
	t._min = _min
}

func (t *TL_channel) Get_min() []byte {
	return t._min
}

func (t *TL_channel) Set_id(_id int32) {
	t._id = _id
}

func (t *TL_channel) Get_id() int32 {
	return t._id
}

func (t *TL_channel) Set_access_hash(_access_hash []byte) {
	t._access_hash = _access_hash
}

func (t *TL_channel) Get_access_hash() []byte {
	return t._access_hash
}

func (t *TL_channel) Set_title(_title string) {
	t._title = _title
}

func (t *TL_channel) Get_title() string {
	return t._title
}

func (t *TL_channel) Set_username(_username []byte) {
	t._username = _username
}

func (t *TL_channel) Get_username() []byte {
	return t._username
}

func (t *TL_channel) Set_photo(_photo []byte) {
	t._photo = _photo
}

func (t *TL_channel) Get_photo() []byte {
	return t._photo
}

func (t *TL_channel) Set_date(_date int32) {
	t._date = _date
}

func (t *TL_channel) Get_date() int32 {
	return t._date
}

func (t *TL_channel) Set_version(_version int32) {
	t._version = _version
}

func (t *TL_channel) Get_version() int32 {
	return t._version
}

func (t *TL_channel) Set_restriction_reason(_restriction_reason []byte) {
	t._restriction_reason = _restriction_reason
}

func (t *TL_channel) Get_restriction_reason() []byte {
	return t._restriction_reason
}

func (t *TL_channel) Set_admin_rights(_admin_rights []byte) {
	t._admin_rights = _admin_rights
}

func (t *TL_channel) Get_admin_rights() []byte {
	return t._admin_rights
}

func (t *TL_channel) Set_banned_rights(_banned_rights []byte) {
	t._banned_rights = _banned_rights
}

func (t *TL_channel) Get_banned_rights() []byte {
	return t._banned_rights
}

func (t *TL_channel) Set_participants_count(_participants_count []byte) {
	t._participants_count = _participants_count
}

func (t *TL_channel) Get_participants_count() []byte {
	return t._participants_count
}

func New_TL_channel() *TL_channel {
	return &TL_channel{}
}

func (t *TL_channel) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channel))
	ec.Bytes(t.Get_creator())
	ec.Bytes(t.Get_left())
	ec.Bytes(t.Get_editor())
	ec.Bytes(t.Get_broadcast())
	ec.Bytes(t.Get_verified())
	ec.Bytes(t.Get_megagroup())
	ec.Bytes(t.Get_restricted())
	ec.Bytes(t.Get_democracy())
	ec.Bytes(t.Get_signatures())
	ec.Bytes(t.Get_min())
	ec.Int(t.Get_id())
	ec.Bytes(t.Get_access_hash())
	ec.String(t.Get_title())
	ec.Bytes(t.Get_username())
	ec.Bytes(t.Get_photo())
	ec.Int(t.Get_date())
	ec.Int(t.Get_version())
	ec.Bytes(t.Get_restriction_reason())
	ec.Bytes(t.Get_admin_rights())
	ec.Bytes(t.Get_banned_rights())
	ec.Bytes(t.Get_participants_count())

	return ec.GetBuffer()
}

func (t *TL_channel) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._creator = dc.Bytes(16)
	t._left = dc.Bytes(16)
	t._editor = dc.Bytes(16)
	t._broadcast = dc.Bytes(16)
	t._verified = dc.Bytes(16)
	t._megagroup = dc.Bytes(16)
	t._restricted = dc.Bytes(16)
	t._democracy = dc.Bytes(16)
	t._signatures = dc.Bytes(16)
	t._min = dc.Bytes(16)
	t._id = dc.Int()
	t._access_hash = dc.Bytes(16)
	t._title = dc.String()
	t._username = dc.Bytes(16)
	t._photo = dc.Bytes(16)
	t._date = dc.Int()
	t._version = dc.Int()
	t._restriction_reason = dc.Bytes(16)
	t._admin_rights = dc.Bytes(16)
	t._banned_rights = dc.Bytes(16)
	t._participants_count = dc.Bytes(16)

}

// channelForbidden#289da732
type TL_channelForbidden struct {
	_flags       []byte
	_broadcast   []byte
	_megagroup   []byte
	_id          int32
	_access_hash int64
	_title       string
	_until_date  []byte
}

func (t *TL_channelForbidden) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_channelForbidden) Get_flags() []byte {
	return t._flags
}

func (t *TL_channelForbidden) Set_broadcast(_broadcast []byte) {
	t._broadcast = _broadcast
}

func (t *TL_channelForbidden) Get_broadcast() []byte {
	return t._broadcast
}

func (t *TL_channelForbidden) Set_megagroup(_megagroup []byte) {
	t._megagroup = _megagroup
}

func (t *TL_channelForbidden) Get_megagroup() []byte {
	return t._megagroup
}

func (t *TL_channelForbidden) Set_id(_id int32) {
	t._id = _id
}

func (t *TL_channelForbidden) Get_id() int32 {
	return t._id
}

func (t *TL_channelForbidden) Set_access_hash(_access_hash int64) {
	t._access_hash = _access_hash
}

func (t *TL_channelForbidden) Get_access_hash() int64 {
	return t._access_hash
}

func (t *TL_channelForbidden) Set_title(_title string) {
	t._title = _title
}

func (t *TL_channelForbidden) Get_title() string {
	return t._title
}

func (t *TL_channelForbidden) Set_until_date(_until_date []byte) {
	t._until_date = _until_date
}

func (t *TL_channelForbidden) Get_until_date() []byte {
	return t._until_date
}

func New_TL_channelForbidden() *TL_channelForbidden {
	return &TL_channelForbidden{}
}

func (t *TL_channelForbidden) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channelForbidden))
	ec.Bytes(t.Get_broadcast())
	ec.Bytes(t.Get_megagroup())
	ec.Int(t.Get_id())
	ec.Long(t.Get_access_hash())
	ec.String(t.Get_title())
	ec.Bytes(t.Get_until_date())

	return ec.GetBuffer()
}

func (t *TL_channelForbidden) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._broadcast = dc.Bytes(16)
	t._megagroup = dc.Bytes(16)
	t._id = dc.Int()
	t._access_hash = dc.Long()
	t._title = dc.String()
	t._until_date = dc.Bytes(16)

}

// chatFull#2e02a614
type TL_chatFull struct {
	_id              int32
	_participants    []byte
	_chat_photo      []byte
	_notify_settings []byte
	_exported_invite []byte
	_bot_info        []byte
}

func (t *TL_chatFull) Set_id(_id int32) {
	t._id = _id
}

func (t *TL_chatFull) Get_id() int32 {
	return t._id
}

func (t *TL_chatFull) Set_participants(_participants []byte) {
	t._participants = _participants
}

func (t *TL_chatFull) Get_participants() []byte {
	return t._participants
}

func (t *TL_chatFull) Set_chat_photo(_chat_photo []byte) {
	t._chat_photo = _chat_photo
}

func (t *TL_chatFull) Get_chat_photo() []byte {
	return t._chat_photo
}

func (t *TL_chatFull) Set_notify_settings(_notify_settings []byte) {
	t._notify_settings = _notify_settings
}

func (t *TL_chatFull) Get_notify_settings() []byte {
	return t._notify_settings
}

func (t *TL_chatFull) Set_exported_invite(_exported_invite []byte) {
	t._exported_invite = _exported_invite
}

func (t *TL_chatFull) Get_exported_invite() []byte {
	return t._exported_invite
}

func (t *TL_chatFull) Set_bot_info(_bot_info []byte) {
	t._bot_info = _bot_info
}

func (t *TL_chatFull) Get_bot_info() []byte {
	return t._bot_info
}

func New_TL_chatFull() *TL_chatFull {
	return &TL_chatFull{}
}

func (t *TL_chatFull) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_chatFull))
	ec.Int(t.Get_id())
	ec.Bytes(t.Get_participants())
	ec.Bytes(t.Get_chat_photo())
	ec.Bytes(t.Get_notify_settings())
	ec.Bytes(t.Get_exported_invite())
	ec.Bytes(t.Get_bot_info())

	return ec.GetBuffer()
}

func (t *TL_chatFull) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Int()
	t._participants = dc.Bytes(16)
	t._chat_photo = dc.Bytes(16)
	t._notify_settings = dc.Bytes(16)
	t._exported_invite = dc.Bytes(16)
	t._bot_info = dc.Bytes(16)

}

// channelFull#76af5481
type TL_channelFull struct {
	_flags                 []byte
	_can_view_participants []byte
	_can_set_username      []byte
	_can_set_stickers      []byte
	_hidden_prehistory     []byte
	_id                    int32
	_about                 string
	_participants_count    []byte
	_admins_count          []byte
	_kicked_count          []byte
	_banned_count          []byte
	_read_inbox_max_id     int32
	_read_outbox_max_id    int32
	_unread_count          int32
	_chat_photo            []byte
	_notify_settings       []byte
	_exported_invite       []byte
	_bot_info              []byte
	_migrated_from_chat_id []byte
	_migrated_from_max_id  []byte
	_pinned_msg_id         []byte
	_stickerset            []byte
	_available_min_id      []byte
}

func (t *TL_channelFull) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_channelFull) Get_flags() []byte {
	return t._flags
}

func (t *TL_channelFull) Set_can_view_participants(_can_view_participants []byte) {
	t._can_view_participants = _can_view_participants
}

func (t *TL_channelFull) Get_can_view_participants() []byte {
	return t._can_view_participants
}

func (t *TL_channelFull) Set_can_set_username(_can_set_username []byte) {
	t._can_set_username = _can_set_username
}

func (t *TL_channelFull) Get_can_set_username() []byte {
	return t._can_set_username
}

func (t *TL_channelFull) Set_can_set_stickers(_can_set_stickers []byte) {
	t._can_set_stickers = _can_set_stickers
}

func (t *TL_channelFull) Get_can_set_stickers() []byte {
	return t._can_set_stickers
}

func (t *TL_channelFull) Set_hidden_prehistory(_hidden_prehistory []byte) {
	t._hidden_prehistory = _hidden_prehistory
}

func (t *TL_channelFull) Get_hidden_prehistory() []byte {
	return t._hidden_prehistory
}

func (t *TL_channelFull) Set_id(_id int32) {
	t._id = _id
}

func (t *TL_channelFull) Get_id() int32 {
	return t._id
}

func (t *TL_channelFull) Set_about(_about string) {
	t._about = _about
}

func (t *TL_channelFull) Get_about() string {
	return t._about
}

func (t *TL_channelFull) Set_participants_count(_participants_count []byte) {
	t._participants_count = _participants_count
}

func (t *TL_channelFull) Get_participants_count() []byte {
	return t._participants_count
}

func (t *TL_channelFull) Set_admins_count(_admins_count []byte) {
	t._admins_count = _admins_count
}

func (t *TL_channelFull) Get_admins_count() []byte {
	return t._admins_count
}

func (t *TL_channelFull) Set_kicked_count(_kicked_count []byte) {
	t._kicked_count = _kicked_count
}

func (t *TL_channelFull) Get_kicked_count() []byte {
	return t._kicked_count
}

func (t *TL_channelFull) Set_banned_count(_banned_count []byte) {
	t._banned_count = _banned_count
}

func (t *TL_channelFull) Get_banned_count() []byte {
	return t._banned_count
}

func (t *TL_channelFull) Set_read_inbox_max_id(_read_inbox_max_id int32) {
	t._read_inbox_max_id = _read_inbox_max_id
}

func (t *TL_channelFull) Get_read_inbox_max_id() int32 {
	return t._read_inbox_max_id
}

func (t *TL_channelFull) Set_read_outbox_max_id(_read_outbox_max_id int32) {
	t._read_outbox_max_id = _read_outbox_max_id
}

func (t *TL_channelFull) Get_read_outbox_max_id() int32 {
	return t._read_outbox_max_id
}

func (t *TL_channelFull) Set_unread_count(_unread_count int32) {
	t._unread_count = _unread_count
}

func (t *TL_channelFull) Get_unread_count() int32 {
	return t._unread_count
}

func (t *TL_channelFull) Set_chat_photo(_chat_photo []byte) {
	t._chat_photo = _chat_photo
}

func (t *TL_channelFull) Get_chat_photo() []byte {
	return t._chat_photo
}

func (t *TL_channelFull) Set_notify_settings(_notify_settings []byte) {
	t._notify_settings = _notify_settings
}

func (t *TL_channelFull) Get_notify_settings() []byte {
	return t._notify_settings
}

func (t *TL_channelFull) Set_exported_invite(_exported_invite []byte) {
	t._exported_invite = _exported_invite
}

func (t *TL_channelFull) Get_exported_invite() []byte {
	return t._exported_invite
}

func (t *TL_channelFull) Set_bot_info(_bot_info []byte) {
	t._bot_info = _bot_info
}

func (t *TL_channelFull) Get_bot_info() []byte {
	return t._bot_info
}

func (t *TL_channelFull) Set_migrated_from_chat_id(_migrated_from_chat_id []byte) {
	t._migrated_from_chat_id = _migrated_from_chat_id
}

func (t *TL_channelFull) Get_migrated_from_chat_id() []byte {
	return t._migrated_from_chat_id
}

func (t *TL_channelFull) Set_migrated_from_max_id(_migrated_from_max_id []byte) {
	t._migrated_from_max_id = _migrated_from_max_id
}

func (t *TL_channelFull) Get_migrated_from_max_id() []byte {
	return t._migrated_from_max_id
}

func (t *TL_channelFull) Set_pinned_msg_id(_pinned_msg_id []byte) {
	t._pinned_msg_id = _pinned_msg_id
}

func (t *TL_channelFull) Get_pinned_msg_id() []byte {
	return t._pinned_msg_id
}

func (t *TL_channelFull) Set_stickerset(_stickerset []byte) {
	t._stickerset = _stickerset
}

func (t *TL_channelFull) Get_stickerset() []byte {
	return t._stickerset
}

func (t *TL_channelFull) Set_available_min_id(_available_min_id []byte) {
	t._available_min_id = _available_min_id
}

func (t *TL_channelFull) Get_available_min_id() []byte {
	return t._available_min_id
}

func New_TL_channelFull() *TL_channelFull {
	return &TL_channelFull{}
}

func (t *TL_channelFull) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channelFull))
	ec.Bytes(t.Get_can_view_participants())
	ec.Bytes(t.Get_can_set_username())
	ec.Bytes(t.Get_can_set_stickers())
	ec.Bytes(t.Get_hidden_prehistory())
	ec.Int(t.Get_id())
	ec.String(t.Get_about())
	ec.Bytes(t.Get_participants_count())
	ec.Bytes(t.Get_admins_count())
	ec.Bytes(t.Get_kicked_count())
	ec.Bytes(t.Get_banned_count())
	ec.Int(t.Get_read_inbox_max_id())
	ec.Int(t.Get_read_outbox_max_id())
	ec.Int(t.Get_unread_count())
	ec.Bytes(t.Get_chat_photo())
	ec.Bytes(t.Get_notify_settings())
	ec.Bytes(t.Get_exported_invite())
	ec.Bytes(t.Get_bot_info())
	ec.Bytes(t.Get_migrated_from_chat_id())
	ec.Bytes(t.Get_migrated_from_max_id())
	ec.Bytes(t.Get_pinned_msg_id())
	ec.Bytes(t.Get_stickerset())
	ec.Bytes(t.Get_available_min_id())

	return ec.GetBuffer()
}

func (t *TL_channelFull) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._can_view_participants = dc.Bytes(16)
	t._can_set_username = dc.Bytes(16)
	t._can_set_stickers = dc.Bytes(16)
	t._hidden_prehistory = dc.Bytes(16)
	t._id = dc.Int()
	t._about = dc.String()
	t._participants_count = dc.Bytes(16)
	t._admins_count = dc.Bytes(16)
	t._kicked_count = dc.Bytes(16)
	t._banned_count = dc.Bytes(16)
	t._read_inbox_max_id = dc.Int()
	t._read_outbox_max_id = dc.Int()
	t._unread_count = dc.Int()
	t._chat_photo = dc.Bytes(16)
	t._notify_settings = dc.Bytes(16)
	t._exported_invite = dc.Bytes(16)
	t._bot_info = dc.Bytes(16)
	t._migrated_from_chat_id = dc.Bytes(16)
	t._migrated_from_max_id = dc.Bytes(16)
	t._pinned_msg_id = dc.Bytes(16)
	t._stickerset = dc.Bytes(16)
	t._available_min_id = dc.Bytes(16)

}

// chatParticipant#c8d7493e
type TL_chatParticipant struct {
	_user_id    int32
	_inviter_id int32
	_date       int32
}

func (t *TL_chatParticipant) Set_user_id(_user_id int32) {
	t._user_id = _user_id
}

func (t *TL_chatParticipant) Get_user_id() int32 {
	return t._user_id
}

func (t *TL_chatParticipant) Set_inviter_id(_inviter_id int32) {
	t._inviter_id = _inviter_id
}

func (t *TL_chatParticipant) Get_inviter_id() int32 {
	return t._inviter_id
}

func (t *TL_chatParticipant) Set_date(_date int32) {
	t._date = _date
}

func (t *TL_chatParticipant) Get_date() int32 {
	return t._date
}

func New_TL_chatParticipant() *TL_chatParticipant {
	return &TL_chatParticipant{}
}

func (t *TL_chatParticipant) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_chatParticipant))
	ec.Int(t.Get_user_id())
	ec.Int(t.Get_inviter_id())
	ec.Int(t.Get_date())

	return ec.GetBuffer()
}

func (t *TL_chatParticipant) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._user_id = dc.Int()
	t._inviter_id = dc.Int()
	t._date = dc.Int()

}

// chatParticipantCreator#da13538a
type TL_chatParticipantCreator struct {
	_user_id int32
}

func (t *TL_chatParticipantCreator) Set_user_id(_user_id int32) {
	t._user_id = _user_id
}

func (t *TL_chatParticipantCreator) Get_user_id() int32 {
	return t._user_id
}

func New_TL_chatParticipantCreator() *TL_chatParticipantCreator {
	return &TL_chatParticipantCreator{}
}

func (t *TL_chatParticipantCreator) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_chatParticipantCreator))
	ec.Int(t.Get_user_id())

	return ec.GetBuffer()
}

func (t *TL_chatParticipantCreator) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._user_id = dc.Int()

}

// chatParticipantAdmin#e2d6e436
type TL_chatParticipantAdmin struct {
	_user_id    int32
	_inviter_id int32
	_date       int32
}

func (t *TL_chatParticipantAdmin) Set_user_id(_user_id int32) {
	t._user_id = _user_id
}

func (t *TL_chatParticipantAdmin) Get_user_id() int32 {
	return t._user_id
}

func (t *TL_chatParticipantAdmin) Set_inviter_id(_inviter_id int32) {
	t._inviter_id = _inviter_id
}

func (t *TL_chatParticipantAdmin) Get_inviter_id() int32 {
	return t._inviter_id
}

func (t *TL_chatParticipantAdmin) Set_date(_date int32) {
	t._date = _date
}

func (t *TL_chatParticipantAdmin) Get_date() int32 {
	return t._date
}

func New_TL_chatParticipantAdmin() *TL_chatParticipantAdmin {
	return &TL_chatParticipantAdmin{}
}

func (t *TL_chatParticipantAdmin) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_chatParticipantAdmin))
	ec.Int(t.Get_user_id())
	ec.Int(t.Get_inviter_id())
	ec.Int(t.Get_date())

	return ec.GetBuffer()
}

func (t *TL_chatParticipantAdmin) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._user_id = dc.Int()
	t._inviter_id = dc.Int()
	t._date = dc.Int()

}

// chatParticipantsForbidden#fc900c2b
type TL_chatParticipantsForbidden struct {
	_flags            []byte
	_chat_id          int32
	_self_participant []byte
}

func (t *TL_chatParticipantsForbidden) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_chatParticipantsForbidden) Get_flags() []byte {
	return t._flags
}

func (t *TL_chatParticipantsForbidden) Set_chat_id(_chat_id int32) {
	t._chat_id = _chat_id
}

func (t *TL_chatParticipantsForbidden) Get_chat_id() int32 {
	return t._chat_id
}

func (t *TL_chatParticipantsForbidden) Set_self_participant(_self_participant []byte) {
	t._self_participant = _self_participant
}

func (t *TL_chatParticipantsForbidden) Get_self_participant() []byte {
	return t._self_participant
}

func New_TL_chatParticipantsForbidden() *TL_chatParticipantsForbidden {
	return &TL_chatParticipantsForbidden{}
}

func (t *TL_chatParticipantsForbidden) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_chatParticipantsForbidden))
	ec.Int(t.Get_chat_id())
	ec.Bytes(t.Get_self_participant())

	return ec.GetBuffer()
}

func (t *TL_chatParticipantsForbidden) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._chat_id = dc.Int()
	t._self_participant = dc.Bytes(16)

}

// chatParticipants#3f460fed
type TL_chatParticipants struct {
	_chat_id      int32
	_participants []byte
	_version      int32
}

func (t *TL_chatParticipants) Set_chat_id(_chat_id int32) {
	t._chat_id = _chat_id
}

func (t *TL_chatParticipants) Get_chat_id() int32 {
	return t._chat_id
}

func (t *TL_chatParticipants) Set_participants(_participants []byte) {
	t._participants = _participants
}

func (t *TL_chatParticipants) Get_participants() []byte {
	return t._participants
}

func (t *TL_chatParticipants) Set_version(_version int32) {
	t._version = _version
}

func (t *TL_chatParticipants) Get_version() int32 {
	return t._version
}

func New_TL_chatParticipants() *TL_chatParticipants {
	return &TL_chatParticipants{}
}

func (t *TL_chatParticipants) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_chatParticipants))
	ec.Int(t.Get_chat_id())
	ec.Bytes(t.Get_participants())
	ec.Int(t.Get_version())

	return ec.GetBuffer()
}

func (t *TL_chatParticipants) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._chat_id = dc.Int()
	t._participants = dc.Bytes(16)
	t._version = dc.Int()

}

// chatPhotoEmpty#37c1011c
type TL_chatPhotoEmpty struct {
}

func New_TL_chatPhotoEmpty() *TL_chatPhotoEmpty {
	return &TL_chatPhotoEmpty{}
}

func (t *TL_chatPhotoEmpty) Encode() []byte {
	return nil
}

func (t *TL_chatPhotoEmpty) Decode(b []byte) {

}

// chatPhoto#6153276a
type TL_chatPhoto struct {
	_photo_small []byte
	_photo_big   []byte
}

func (t *TL_chatPhoto) Set_photo_small(_photo_small []byte) {
	t._photo_small = _photo_small
}

func (t *TL_chatPhoto) Get_photo_small() []byte {
	return t._photo_small
}

func (t *TL_chatPhoto) Set_photo_big(_photo_big []byte) {
	t._photo_big = _photo_big
}

func (t *TL_chatPhoto) Get_photo_big() []byte {
	return t._photo_big
}

func New_TL_chatPhoto() *TL_chatPhoto {
	return &TL_chatPhoto{}
}

func (t *TL_chatPhoto) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_chatPhoto))
	ec.Bytes(t.Get_photo_small())
	ec.Bytes(t.Get_photo_big())

	return ec.GetBuffer()
}

func (t *TL_chatPhoto) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._photo_small = dc.Bytes(16)
	t._photo_big = dc.Bytes(16)

}

// messageEmpty#83e5de54
type TL_messageEmpty struct {
	_id int32
}

func (t *TL_messageEmpty) Set_id(_id int32) {
	t._id = _id
}

func (t *TL_messageEmpty) Get_id() int32 {
	return t._id
}

func New_TL_messageEmpty() *TL_messageEmpty {
	return &TL_messageEmpty{}
}

func (t *TL_messageEmpty) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messageEmpty))
	ec.Int(t.Get_id())

	return ec.GetBuffer()
}

func (t *TL_messageEmpty) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Int()

}

// message#44f9b43d
type TL_message struct {
	_flags           []byte
	_out             []byte
	_mentioned       []byte
	_media_unread    []byte
	_silent          []byte
	_post            []byte
	_id              int32
	_from_id         []byte
	_to_id           []byte
	_fwd_from        []byte
	_via_bot_id      []byte
	_reply_to_msg_id []byte
	_date            int32
	_message         string
	_media           []byte
	_reply_markup    []byte
	_entities        []byte
	_views           []byte
	_edit_date       []byte
	_post_author     []byte
	_grouped_id      []byte
}

func (t *TL_message) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_message) Get_flags() []byte {
	return t._flags
}

func (t *TL_message) Set_out(_out []byte) {
	t._out = _out
}

func (t *TL_message) Get_out() []byte {
	return t._out
}

func (t *TL_message) Set_mentioned(_mentioned []byte) {
	t._mentioned = _mentioned
}

func (t *TL_message) Get_mentioned() []byte {
	return t._mentioned
}

func (t *TL_message) Set_media_unread(_media_unread []byte) {
	t._media_unread = _media_unread
}

func (t *TL_message) Get_media_unread() []byte {
	return t._media_unread
}

func (t *TL_message) Set_silent(_silent []byte) {
	t._silent = _silent
}

func (t *TL_message) Get_silent() []byte {
	return t._silent
}

func (t *TL_message) Set_post(_post []byte) {
	t._post = _post
}

func (t *TL_message) Get_post() []byte {
	return t._post
}

func (t *TL_message) Set_id(_id int32) {
	t._id = _id
}

func (t *TL_message) Get_id() int32 {
	return t._id
}

func (t *TL_message) Set_from_id(_from_id []byte) {
	t._from_id = _from_id
}

func (t *TL_message) Get_from_id() []byte {
	return t._from_id
}

func (t *TL_message) Set_to_id(_to_id []byte) {
	t._to_id = _to_id
}

func (t *TL_message) Get_to_id() []byte {
	return t._to_id
}

func (t *TL_message) Set_fwd_from(_fwd_from []byte) {
	t._fwd_from = _fwd_from
}

func (t *TL_message) Get_fwd_from() []byte {
	return t._fwd_from
}

func (t *TL_message) Set_via_bot_id(_via_bot_id []byte) {
	t._via_bot_id = _via_bot_id
}

func (t *TL_message) Get_via_bot_id() []byte {
	return t._via_bot_id
}

func (t *TL_message) Set_reply_to_msg_id(_reply_to_msg_id []byte) {
	t._reply_to_msg_id = _reply_to_msg_id
}

func (t *TL_message) Get_reply_to_msg_id() []byte {
	return t._reply_to_msg_id
}

func (t *TL_message) Set_date(_date int32) {
	t._date = _date
}

func (t *TL_message) Get_date() int32 {
	return t._date
}

func (t *TL_message) Set_message(_message string) {
	t._message = _message
}

func (t *TL_message) Get_message() string {
	return t._message
}

func (t *TL_message) Set_media(_media []byte) {
	t._media = _media
}

func (t *TL_message) Get_media() []byte {
	return t._media
}

func (t *TL_message) Set_reply_markup(_reply_markup []byte) {
	t._reply_markup = _reply_markup
}

func (t *TL_message) Get_reply_markup() []byte {
	return t._reply_markup
}

func (t *TL_message) Set_entities(_entities []byte) {
	t._entities = _entities
}

func (t *TL_message) Get_entities() []byte {
	return t._entities
}

func (t *TL_message) Set_views(_views []byte) {
	t._views = _views
}

func (t *TL_message) Get_views() []byte {
	return t._views
}

func (t *TL_message) Set_edit_date(_edit_date []byte) {
	t._edit_date = _edit_date
}

func (t *TL_message) Get_edit_date() []byte {
	return t._edit_date
}

func (t *TL_message) Set_post_author(_post_author []byte) {
	t._post_author = _post_author
}

func (t *TL_message) Get_post_author() []byte {
	return t._post_author
}

func (t *TL_message) Set_grouped_id(_grouped_id []byte) {
	t._grouped_id = _grouped_id
}

func (t *TL_message) Get_grouped_id() []byte {
	return t._grouped_id
}

func New_TL_message() *TL_message {
	return &TL_message{}
}

func (t *TL_message) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_message))
	ec.Bytes(t.Get_out())
	ec.Bytes(t.Get_mentioned())
	ec.Bytes(t.Get_media_unread())
	ec.Bytes(t.Get_silent())
	ec.Bytes(t.Get_post())
	ec.Int(t.Get_id())
	ec.Bytes(t.Get_from_id())
	ec.Bytes(t.Get_to_id())
	ec.Bytes(t.Get_fwd_from())
	ec.Bytes(t.Get_via_bot_id())
	ec.Bytes(t.Get_reply_to_msg_id())
	ec.Int(t.Get_date())
	ec.String(t.Get_message())
	ec.Bytes(t.Get_media())
	ec.Bytes(t.Get_reply_markup())
	ec.Bytes(t.Get_entities())
	ec.Bytes(t.Get_views())
	ec.Bytes(t.Get_edit_date())
	ec.Bytes(t.Get_post_author())
	ec.Bytes(t.Get_grouped_id())

	return ec.GetBuffer()
}

func (t *TL_message) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._out = dc.Bytes(16)
	t._mentioned = dc.Bytes(16)
	t._media_unread = dc.Bytes(16)
	t._silent = dc.Bytes(16)
	t._post = dc.Bytes(16)
	t._id = dc.Int()
	t._from_id = dc.Bytes(16)
	t._to_id = dc.Bytes(16)
	t._fwd_from = dc.Bytes(16)
	t._via_bot_id = dc.Bytes(16)
	t._reply_to_msg_id = dc.Bytes(16)
	t._date = dc.Int()
	t._message = dc.String()
	t._media = dc.Bytes(16)
	t._reply_markup = dc.Bytes(16)
	t._entities = dc.Bytes(16)
	t._views = dc.Bytes(16)
	t._edit_date = dc.Bytes(16)
	t._post_author = dc.Bytes(16)
	t._grouped_id = dc.Bytes(16)

}

// messageService#9e19a1f6
type TL_messageService struct {
	_flags           []byte
	_out             []byte
	_mentioned       []byte
	_media_unread    []byte
	_silent          []byte
	_post            []byte
	_id              int32
	_from_id         []byte
	_to_id           []byte
	_reply_to_msg_id []byte
	_date            int32
	_action          []byte
}

func (t *TL_messageService) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_messageService) Get_flags() []byte {
	return t._flags
}

func (t *TL_messageService) Set_out(_out []byte) {
	t._out = _out
}

func (t *TL_messageService) Get_out() []byte {
	return t._out
}

func (t *TL_messageService) Set_mentioned(_mentioned []byte) {
	t._mentioned = _mentioned
}

func (t *TL_messageService) Get_mentioned() []byte {
	return t._mentioned
}

func (t *TL_messageService) Set_media_unread(_media_unread []byte) {
	t._media_unread = _media_unread
}

func (t *TL_messageService) Get_media_unread() []byte {
	return t._media_unread
}

func (t *TL_messageService) Set_silent(_silent []byte) {
	t._silent = _silent
}

func (t *TL_messageService) Get_silent() []byte {
	return t._silent
}

func (t *TL_messageService) Set_post(_post []byte) {
	t._post = _post
}

func (t *TL_messageService) Get_post() []byte {
	return t._post
}

func (t *TL_messageService) Set_id(_id int32) {
	t._id = _id
}

func (t *TL_messageService) Get_id() int32 {
	return t._id
}

func (t *TL_messageService) Set_from_id(_from_id []byte) {
	t._from_id = _from_id
}

func (t *TL_messageService) Get_from_id() []byte {
	return t._from_id
}

func (t *TL_messageService) Set_to_id(_to_id []byte) {
	t._to_id = _to_id
}

func (t *TL_messageService) Get_to_id() []byte {
	return t._to_id
}

func (t *TL_messageService) Set_reply_to_msg_id(_reply_to_msg_id []byte) {
	t._reply_to_msg_id = _reply_to_msg_id
}

func (t *TL_messageService) Get_reply_to_msg_id() []byte {
	return t._reply_to_msg_id
}

func (t *TL_messageService) Set_date(_date int32) {
	t._date = _date
}

func (t *TL_messageService) Get_date() int32 {
	return t._date
}

func (t *TL_messageService) Set_action(_action []byte) {
	t._action = _action
}

func (t *TL_messageService) Get_action() []byte {
	return t._action
}

func New_TL_messageService() *TL_messageService {
	return &TL_messageService{}
}

func (t *TL_messageService) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messageService))
	ec.Bytes(t.Get_out())
	ec.Bytes(t.Get_mentioned())
	ec.Bytes(t.Get_media_unread())
	ec.Bytes(t.Get_silent())
	ec.Bytes(t.Get_post())
	ec.Int(t.Get_id())
	ec.Bytes(t.Get_from_id())
	ec.Bytes(t.Get_to_id())
	ec.Bytes(t.Get_reply_to_msg_id())
	ec.Int(t.Get_date())
	ec.Bytes(t.Get_action())

	return ec.GetBuffer()
}

func (t *TL_messageService) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._out = dc.Bytes(16)
	t._mentioned = dc.Bytes(16)
	t._media_unread = dc.Bytes(16)
	t._silent = dc.Bytes(16)
	t._post = dc.Bytes(16)
	t._id = dc.Int()
	t._from_id = dc.Bytes(16)
	t._to_id = dc.Bytes(16)
	t._reply_to_msg_id = dc.Bytes(16)
	t._date = dc.Int()
	t._action = dc.Bytes(16)

}

// messageMediaEmpty#3ded6320
type TL_messageMediaEmpty struct {
}

func New_TL_messageMediaEmpty() *TL_messageMediaEmpty {
	return &TL_messageMediaEmpty{}
}

func (t *TL_messageMediaEmpty) Encode() []byte {
	return nil
}

func (t *TL_messageMediaEmpty) Decode(b []byte) {

}

// messageMediaPhoto#b5223b0f
type TL_messageMediaPhoto struct {
	_flags       []byte
	_photo       []byte
	_caption     []byte
	_ttl_seconds []byte
}

func (t *TL_messageMediaPhoto) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_messageMediaPhoto) Get_flags() []byte {
	return t._flags
}

func (t *TL_messageMediaPhoto) Set_photo(_photo []byte) {
	t._photo = _photo
}

func (t *TL_messageMediaPhoto) Get_photo() []byte {
	return t._photo
}

func (t *TL_messageMediaPhoto) Set_caption(_caption []byte) {
	t._caption = _caption
}

func (t *TL_messageMediaPhoto) Get_caption() []byte {
	return t._caption
}

func (t *TL_messageMediaPhoto) Set_ttl_seconds(_ttl_seconds []byte) {
	t._ttl_seconds = _ttl_seconds
}

func (t *TL_messageMediaPhoto) Get_ttl_seconds() []byte {
	return t._ttl_seconds
}

func New_TL_messageMediaPhoto() *TL_messageMediaPhoto {
	return &TL_messageMediaPhoto{}
}

func (t *TL_messageMediaPhoto) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messageMediaPhoto))
	ec.Bytes(t.Get_photo())
	ec.Bytes(t.Get_caption())
	ec.Bytes(t.Get_ttl_seconds())

	return ec.GetBuffer()
}

func (t *TL_messageMediaPhoto) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._photo = dc.Bytes(16)
	t._caption = dc.Bytes(16)
	t._ttl_seconds = dc.Bytes(16)

}

// messageMediaGeo#56e0d474
type TL_messageMediaGeo struct {
	_geo []byte
}

func (t *TL_messageMediaGeo) Set_geo(_geo []byte) {
	t._geo = _geo
}

func (t *TL_messageMediaGeo) Get_geo() []byte {
	return t._geo
}

func New_TL_messageMediaGeo() *TL_messageMediaGeo {
	return &TL_messageMediaGeo{}
}

func (t *TL_messageMediaGeo) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messageMediaGeo))
	ec.Bytes(t.Get_geo())

	return ec.GetBuffer()
}

func (t *TL_messageMediaGeo) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._geo = dc.Bytes(16)

}

// messageMediaContact#5e7d2f39
type TL_messageMediaContact struct {
	_phone_number string
	_first_name   string
	_last_name    string
	_user_id      int32
}

func (t *TL_messageMediaContact) Set_phone_number(_phone_number string) {
	t._phone_number = _phone_number
}

func (t *TL_messageMediaContact) Get_phone_number() string {
	return t._phone_number
}

func (t *TL_messageMediaContact) Set_first_name(_first_name string) {
	t._first_name = _first_name
}

func (t *TL_messageMediaContact) Get_first_name() string {
	return t._first_name
}

func (t *TL_messageMediaContact) Set_last_name(_last_name string) {
	t._last_name = _last_name
}

func (t *TL_messageMediaContact) Get_last_name() string {
	return t._last_name
}

func (t *TL_messageMediaContact) Set_user_id(_user_id int32) {
	t._user_id = _user_id
}

func (t *TL_messageMediaContact) Get_user_id() int32 {
	return t._user_id
}

func New_TL_messageMediaContact() *TL_messageMediaContact {
	return &TL_messageMediaContact{}
}

func (t *TL_messageMediaContact) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messageMediaContact))
	ec.String(t.Get_phone_number())
	ec.String(t.Get_first_name())
	ec.String(t.Get_last_name())
	ec.Int(t.Get_user_id())

	return ec.GetBuffer()
}

func (t *TL_messageMediaContact) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._phone_number = dc.String()
	t._first_name = dc.String()
	t._last_name = dc.String()
	t._user_id = dc.Int()

}

// messageMediaUnsupported#9f84f49e
type TL_messageMediaUnsupported struct {
}

func New_TL_messageMediaUnsupported() *TL_messageMediaUnsupported {
	return &TL_messageMediaUnsupported{}
}

func (t *TL_messageMediaUnsupported) Encode() []byte {
	return nil
}

func (t *TL_messageMediaUnsupported) Decode(b []byte) {

}

// messageMediaDocument#7c4414d3
type TL_messageMediaDocument struct {
	_flags       []byte
	_document    []byte
	_caption     []byte
	_ttl_seconds []byte
}

func (t *TL_messageMediaDocument) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_messageMediaDocument) Get_flags() []byte {
	return t._flags
}

func (t *TL_messageMediaDocument) Set_document(_document []byte) {
	t._document = _document
}

func (t *TL_messageMediaDocument) Get_document() []byte {
	return t._document
}

func (t *TL_messageMediaDocument) Set_caption(_caption []byte) {
	t._caption = _caption
}

func (t *TL_messageMediaDocument) Get_caption() []byte {
	return t._caption
}

func (t *TL_messageMediaDocument) Set_ttl_seconds(_ttl_seconds []byte) {
	t._ttl_seconds = _ttl_seconds
}

func (t *TL_messageMediaDocument) Get_ttl_seconds() []byte {
	return t._ttl_seconds
}

func New_TL_messageMediaDocument() *TL_messageMediaDocument {
	return &TL_messageMediaDocument{}
}

func (t *TL_messageMediaDocument) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messageMediaDocument))
	ec.Bytes(t.Get_document())
	ec.Bytes(t.Get_caption())
	ec.Bytes(t.Get_ttl_seconds())

	return ec.GetBuffer()
}

func (t *TL_messageMediaDocument) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._document = dc.Bytes(16)
	t._caption = dc.Bytes(16)
	t._ttl_seconds = dc.Bytes(16)

}

// messageMediaWebPage#a32dd600
type TL_messageMediaWebPage struct {
	_webpage []byte
}

func (t *TL_messageMediaWebPage) Set_webpage(_webpage []byte) {
	t._webpage = _webpage
}

func (t *TL_messageMediaWebPage) Get_webpage() []byte {
	return t._webpage
}

func New_TL_messageMediaWebPage() *TL_messageMediaWebPage {
	return &TL_messageMediaWebPage{}
}

func (t *TL_messageMediaWebPage) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messageMediaWebPage))
	ec.Bytes(t.Get_webpage())

	return ec.GetBuffer()
}

func (t *TL_messageMediaWebPage) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._webpage = dc.Bytes(16)

}

// messageMediaVenue#2ec0533f
type TL_messageMediaVenue struct {
	_geo        []byte
	_title      string
	_address    string
	_provider   string
	_venue_id   string
	_venue_type string
}

func (t *TL_messageMediaVenue) Set_geo(_geo []byte) {
	t._geo = _geo
}

func (t *TL_messageMediaVenue) Get_geo() []byte {
	return t._geo
}

func (t *TL_messageMediaVenue) Set_title(_title string) {
	t._title = _title
}

func (t *TL_messageMediaVenue) Get_title() string {
	return t._title
}

func (t *TL_messageMediaVenue) Set_address(_address string) {
	t._address = _address
}

func (t *TL_messageMediaVenue) Get_address() string {
	return t._address
}

func (t *TL_messageMediaVenue) Set_provider(_provider string) {
	t._provider = _provider
}

func (t *TL_messageMediaVenue) Get_provider() string {
	return t._provider
}

func (t *TL_messageMediaVenue) Set_venue_id(_venue_id string) {
	t._venue_id = _venue_id
}

func (t *TL_messageMediaVenue) Get_venue_id() string {
	return t._venue_id
}

func (t *TL_messageMediaVenue) Set_venue_type(_venue_type string) {
	t._venue_type = _venue_type
}

func (t *TL_messageMediaVenue) Get_venue_type() string {
	return t._venue_type
}

func New_TL_messageMediaVenue() *TL_messageMediaVenue {
	return &TL_messageMediaVenue{}
}

func (t *TL_messageMediaVenue) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messageMediaVenue))
	ec.Bytes(t.Get_geo())
	ec.String(t.Get_title())
	ec.String(t.Get_address())
	ec.String(t.Get_provider())
	ec.String(t.Get_venue_id())
	ec.String(t.Get_venue_type())

	return ec.GetBuffer()
}

func (t *TL_messageMediaVenue) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._geo = dc.Bytes(16)
	t._title = dc.String()
	t._address = dc.String()
	t._provider = dc.String()
	t._venue_id = dc.String()
	t._venue_type = dc.String()

}

// messageMediaGame#fdb19008
type TL_messageMediaGame struct {
	_game []byte
}

func (t *TL_messageMediaGame) Set_game(_game []byte) {
	t._game = _game
}

func (t *TL_messageMediaGame) Get_game() []byte {
	return t._game
}

func New_TL_messageMediaGame() *TL_messageMediaGame {
	return &TL_messageMediaGame{}
}

func (t *TL_messageMediaGame) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messageMediaGame))
	ec.Bytes(t.Get_game())

	return ec.GetBuffer()
}

func (t *TL_messageMediaGame) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._game = dc.Bytes(16)

}

// messageMediaInvoice#84551347
type TL_messageMediaInvoice struct {
	_flags                      []byte
	_shipping_address_requested []byte
	_test                       []byte
	_title                      string
	_description                string
	_photo                      []byte
	_receipt_msg_id             []byte
	_currency                   string
	_total_amount               int64
	_start_param                string
}

func (t *TL_messageMediaInvoice) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_messageMediaInvoice) Get_flags() []byte {
	return t._flags
}

func (t *TL_messageMediaInvoice) Set_shipping_address_requested(_shipping_address_requested []byte) {
	t._shipping_address_requested = _shipping_address_requested
}

func (t *TL_messageMediaInvoice) Get_shipping_address_requested() []byte {
	return t._shipping_address_requested
}

func (t *TL_messageMediaInvoice) Set_test(_test []byte) {
	t._test = _test
}

func (t *TL_messageMediaInvoice) Get_test() []byte {
	return t._test
}

func (t *TL_messageMediaInvoice) Set_title(_title string) {
	t._title = _title
}

func (t *TL_messageMediaInvoice) Get_title() string {
	return t._title
}

func (t *TL_messageMediaInvoice) Set_description(_description string) {
	t._description = _description
}

func (t *TL_messageMediaInvoice) Get_description() string {
	return t._description
}

func (t *TL_messageMediaInvoice) Set_photo(_photo []byte) {
	t._photo = _photo
}

func (t *TL_messageMediaInvoice) Get_photo() []byte {
	return t._photo
}

func (t *TL_messageMediaInvoice) Set_receipt_msg_id(_receipt_msg_id []byte) {
	t._receipt_msg_id = _receipt_msg_id
}

func (t *TL_messageMediaInvoice) Get_receipt_msg_id() []byte {
	return t._receipt_msg_id
}

func (t *TL_messageMediaInvoice) Set_currency(_currency string) {
	t._currency = _currency
}

func (t *TL_messageMediaInvoice) Get_currency() string {
	return t._currency
}

func (t *TL_messageMediaInvoice) Set_total_amount(_total_amount int64) {
	t._total_amount = _total_amount
}

func (t *TL_messageMediaInvoice) Get_total_amount() int64 {
	return t._total_amount
}

func (t *TL_messageMediaInvoice) Set_start_param(_start_param string) {
	t._start_param = _start_param
}

func (t *TL_messageMediaInvoice) Get_start_param() string {
	return t._start_param
}

func New_TL_messageMediaInvoice() *TL_messageMediaInvoice {
	return &TL_messageMediaInvoice{}
}

func (t *TL_messageMediaInvoice) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messageMediaInvoice))
	ec.Bytes(t.Get_shipping_address_requested())
	ec.Bytes(t.Get_test())
	ec.String(t.Get_title())
	ec.String(t.Get_description())
	ec.Bytes(t.Get_photo())
	ec.Bytes(t.Get_receipt_msg_id())
	ec.String(t.Get_currency())
	ec.Long(t.Get_total_amount())
	ec.String(t.Get_start_param())

	return ec.GetBuffer()
}

func (t *TL_messageMediaInvoice) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._shipping_address_requested = dc.Bytes(16)
	t._test = dc.Bytes(16)
	t._title = dc.String()
	t._description = dc.String()
	t._photo = dc.Bytes(16)
	t._receipt_msg_id = dc.Bytes(16)
	t._currency = dc.String()
	t._total_amount = dc.Long()
	t._start_param = dc.String()

}

// messageMediaGeoLive#7c3c2609
type TL_messageMediaGeoLive struct {
	_geo    []byte
	_period int32
}

func (t *TL_messageMediaGeoLive) Set_geo(_geo []byte) {
	t._geo = _geo
}

func (t *TL_messageMediaGeoLive) Get_geo() []byte {
	return t._geo
}

func (t *TL_messageMediaGeoLive) Set_period(_period int32) {
	t._period = _period
}

func (t *TL_messageMediaGeoLive) Get_period() int32 {
	return t._period
}

func New_TL_messageMediaGeoLive() *TL_messageMediaGeoLive {
	return &TL_messageMediaGeoLive{}
}

func (t *TL_messageMediaGeoLive) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messageMediaGeoLive))
	ec.Bytes(t.Get_geo())
	ec.Int(t.Get_period())

	return ec.GetBuffer()
}

func (t *TL_messageMediaGeoLive) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._geo = dc.Bytes(16)
	t._period = dc.Int()

}

// messageActionEmpty#b6aef7b0
type TL_messageActionEmpty struct {
}

func New_TL_messageActionEmpty() *TL_messageActionEmpty {
	return &TL_messageActionEmpty{}
}

func (t *TL_messageActionEmpty) Encode() []byte {
	return nil
}

func (t *TL_messageActionEmpty) Decode(b []byte) {

}

// messageActionChatCreate#a6638b9a
type TL_messageActionChatCreate struct {
	_title string
	_users []byte
}

func (t *TL_messageActionChatCreate) Set_title(_title string) {
	t._title = _title
}

func (t *TL_messageActionChatCreate) Get_title() string {
	return t._title
}

func (t *TL_messageActionChatCreate) Set_users(_users []byte) {
	t._users = _users
}

func (t *TL_messageActionChatCreate) Get_users() []byte {
	return t._users
}

func New_TL_messageActionChatCreate() *TL_messageActionChatCreate {
	return &TL_messageActionChatCreate{}
}

func (t *TL_messageActionChatCreate) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messageActionChatCreate))
	ec.String(t.Get_title())
	ec.Bytes(t.Get_users())

	return ec.GetBuffer()
}

func (t *TL_messageActionChatCreate) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._title = dc.String()
	t._users = dc.Bytes(16)

}

// messageActionChatEditTitle#b5a1ce5a
type TL_messageActionChatEditTitle struct {
	_title string
}

func (t *TL_messageActionChatEditTitle) Set_title(_title string) {
	t._title = _title
}

func (t *TL_messageActionChatEditTitle) Get_title() string {
	return t._title
}

func New_TL_messageActionChatEditTitle() *TL_messageActionChatEditTitle {
	return &TL_messageActionChatEditTitle{}
}

func (t *TL_messageActionChatEditTitle) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messageActionChatEditTitle))
	ec.String(t.Get_title())

	return ec.GetBuffer()
}

func (t *TL_messageActionChatEditTitle) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._title = dc.String()

}

// messageActionChatEditPhoto#7fcb13a8
type TL_messageActionChatEditPhoto struct {
	_photo []byte
}

func (t *TL_messageActionChatEditPhoto) Set_photo(_photo []byte) {
	t._photo = _photo
}

func (t *TL_messageActionChatEditPhoto) Get_photo() []byte {
	return t._photo
}

func New_TL_messageActionChatEditPhoto() *TL_messageActionChatEditPhoto {
	return &TL_messageActionChatEditPhoto{}
}

func (t *TL_messageActionChatEditPhoto) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messageActionChatEditPhoto))
	ec.Bytes(t.Get_photo())

	return ec.GetBuffer()
}

func (t *TL_messageActionChatEditPhoto) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._photo = dc.Bytes(16)

}

// messageActionChatDeletePhoto#95e3fbef
type TL_messageActionChatDeletePhoto struct {
}

func New_TL_messageActionChatDeletePhoto() *TL_messageActionChatDeletePhoto {
	return &TL_messageActionChatDeletePhoto{}
}

func (t *TL_messageActionChatDeletePhoto) Encode() []byte {
	return nil
}

func (t *TL_messageActionChatDeletePhoto) Decode(b []byte) {

}

// messageActionChatAddUser#488a7337
type TL_messageActionChatAddUser struct {
	_users []byte
}

func (t *TL_messageActionChatAddUser) Set_users(_users []byte) {
	t._users = _users
}

func (t *TL_messageActionChatAddUser) Get_users() []byte {
	return t._users
}

func New_TL_messageActionChatAddUser() *TL_messageActionChatAddUser {
	return &TL_messageActionChatAddUser{}
}

func (t *TL_messageActionChatAddUser) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messageActionChatAddUser))
	ec.Bytes(t.Get_users())

	return ec.GetBuffer()
}

func (t *TL_messageActionChatAddUser) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._users = dc.Bytes(16)

}

// messageActionChatDeleteUser#b2ae9b0c
type TL_messageActionChatDeleteUser struct {
	_user_id int32
}

func (t *TL_messageActionChatDeleteUser) Set_user_id(_user_id int32) {
	t._user_id = _user_id
}

func (t *TL_messageActionChatDeleteUser) Get_user_id() int32 {
	return t._user_id
}

func New_TL_messageActionChatDeleteUser() *TL_messageActionChatDeleteUser {
	return &TL_messageActionChatDeleteUser{}
}

func (t *TL_messageActionChatDeleteUser) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messageActionChatDeleteUser))
	ec.Int(t.Get_user_id())

	return ec.GetBuffer()
}

func (t *TL_messageActionChatDeleteUser) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._user_id = dc.Int()

}

// messageActionChatJoinedByLink#f89cf5e8
type TL_messageActionChatJoinedByLink struct {
	_inviter_id int32
}

func (t *TL_messageActionChatJoinedByLink) Set_inviter_id(_inviter_id int32) {
	t._inviter_id = _inviter_id
}

func (t *TL_messageActionChatJoinedByLink) Get_inviter_id() int32 {
	return t._inviter_id
}

func New_TL_messageActionChatJoinedByLink() *TL_messageActionChatJoinedByLink {
	return &TL_messageActionChatJoinedByLink{}
}

func (t *TL_messageActionChatJoinedByLink) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messageActionChatJoinedByLink))
	ec.Int(t.Get_inviter_id())

	return ec.GetBuffer()
}

func (t *TL_messageActionChatJoinedByLink) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._inviter_id = dc.Int()

}

// messageActionChannelCreate#95d2ac92
type TL_messageActionChannelCreate struct {
	_title string
}

func (t *TL_messageActionChannelCreate) Set_title(_title string) {
	t._title = _title
}

func (t *TL_messageActionChannelCreate) Get_title() string {
	return t._title
}

func New_TL_messageActionChannelCreate() *TL_messageActionChannelCreate {
	return &TL_messageActionChannelCreate{}
}

func (t *TL_messageActionChannelCreate) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messageActionChannelCreate))
	ec.String(t.Get_title())

	return ec.GetBuffer()
}

func (t *TL_messageActionChannelCreate) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._title = dc.String()

}

// messageActionChatMigrateTo#51bdb021
type TL_messageActionChatMigrateTo struct {
	_channel_id int32
}

func (t *TL_messageActionChatMigrateTo) Set_channel_id(_channel_id int32) {
	t._channel_id = _channel_id
}

func (t *TL_messageActionChatMigrateTo) Get_channel_id() int32 {
	return t._channel_id
}

func New_TL_messageActionChatMigrateTo() *TL_messageActionChatMigrateTo {
	return &TL_messageActionChatMigrateTo{}
}

func (t *TL_messageActionChatMigrateTo) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messageActionChatMigrateTo))
	ec.Int(t.Get_channel_id())

	return ec.GetBuffer()
}

func (t *TL_messageActionChatMigrateTo) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._channel_id = dc.Int()

}

// messageActionChannelMigrateFrom#b055eaee
type TL_messageActionChannelMigrateFrom struct {
	_title   string
	_chat_id int32
}

func (t *TL_messageActionChannelMigrateFrom) Set_title(_title string) {
	t._title = _title
}

func (t *TL_messageActionChannelMigrateFrom) Get_title() string {
	return t._title
}

func (t *TL_messageActionChannelMigrateFrom) Set_chat_id(_chat_id int32) {
	t._chat_id = _chat_id
}

func (t *TL_messageActionChannelMigrateFrom) Get_chat_id() int32 {
	return t._chat_id
}

func New_TL_messageActionChannelMigrateFrom() *TL_messageActionChannelMigrateFrom {
	return &TL_messageActionChannelMigrateFrom{}
}

func (t *TL_messageActionChannelMigrateFrom) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messageActionChannelMigrateFrom))
	ec.String(t.Get_title())
	ec.Int(t.Get_chat_id())

	return ec.GetBuffer()
}

func (t *TL_messageActionChannelMigrateFrom) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._title = dc.String()
	t._chat_id = dc.Int()

}

// messageActionPinMessage#94bd38ed
type TL_messageActionPinMessage struct {
}

func New_TL_messageActionPinMessage() *TL_messageActionPinMessage {
	return &TL_messageActionPinMessage{}
}

func (t *TL_messageActionPinMessage) Encode() []byte {
	return nil
}

func (t *TL_messageActionPinMessage) Decode(b []byte) {

}

// messageActionHistoryClear#9fbab604
type TL_messageActionHistoryClear struct {
}

func New_TL_messageActionHistoryClear() *TL_messageActionHistoryClear {
	return &TL_messageActionHistoryClear{}
}

func (t *TL_messageActionHistoryClear) Encode() []byte {
	return nil
}

func (t *TL_messageActionHistoryClear) Decode(b []byte) {

}

// messageActionGameScore#92a72876
type TL_messageActionGameScore struct {
	_game_id int64
	_score   int32
}

func (t *TL_messageActionGameScore) Set_game_id(_game_id int64) {
	t._game_id = _game_id
}

func (t *TL_messageActionGameScore) Get_game_id() int64 {
	return t._game_id
}

func (t *TL_messageActionGameScore) Set_score(_score int32) {
	t._score = _score
}

func (t *TL_messageActionGameScore) Get_score() int32 {
	return t._score
}

func New_TL_messageActionGameScore() *TL_messageActionGameScore {
	return &TL_messageActionGameScore{}
}

func (t *TL_messageActionGameScore) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messageActionGameScore))
	ec.Long(t.Get_game_id())
	ec.Int(t.Get_score())

	return ec.GetBuffer()
}

func (t *TL_messageActionGameScore) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._game_id = dc.Long()
	t._score = dc.Int()

}

// messageActionPaymentSentMe#8f31b327
type TL_messageActionPaymentSentMe struct {
	_flags              []byte
	_currency           string
	_total_amount       int64
	_payload            []byte
	_info               []byte
	_shipping_option_id []byte
	_charge             []byte
}

func (t *TL_messageActionPaymentSentMe) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_messageActionPaymentSentMe) Get_flags() []byte {
	return t._flags
}

func (t *TL_messageActionPaymentSentMe) Set_currency(_currency string) {
	t._currency = _currency
}

func (t *TL_messageActionPaymentSentMe) Get_currency() string {
	return t._currency
}

func (t *TL_messageActionPaymentSentMe) Set_total_amount(_total_amount int64) {
	t._total_amount = _total_amount
}

func (t *TL_messageActionPaymentSentMe) Get_total_amount() int64 {
	return t._total_amount
}

func (t *TL_messageActionPaymentSentMe) Set_payload(_payload []byte) {
	t._payload = _payload
}

func (t *TL_messageActionPaymentSentMe) Get_payload() []byte {
	return t._payload
}

func (t *TL_messageActionPaymentSentMe) Set_info(_info []byte) {
	t._info = _info
}

func (t *TL_messageActionPaymentSentMe) Get_info() []byte {
	return t._info
}

func (t *TL_messageActionPaymentSentMe) Set_shipping_option_id(_shipping_option_id []byte) {
	t._shipping_option_id = _shipping_option_id
}

func (t *TL_messageActionPaymentSentMe) Get_shipping_option_id() []byte {
	return t._shipping_option_id
}

func (t *TL_messageActionPaymentSentMe) Set_charge(_charge []byte) {
	t._charge = _charge
}

func (t *TL_messageActionPaymentSentMe) Get_charge() []byte {
	return t._charge
}

func New_TL_messageActionPaymentSentMe() *TL_messageActionPaymentSentMe {
	return &TL_messageActionPaymentSentMe{}
}

func (t *TL_messageActionPaymentSentMe) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messageActionPaymentSentMe))
	ec.String(t.Get_currency())
	ec.Long(t.Get_total_amount())
	ec.Bytes(t.Get_payload())
	ec.Bytes(t.Get_info())
	ec.Bytes(t.Get_shipping_option_id())
	ec.Bytes(t.Get_charge())

	return ec.GetBuffer()
}

func (t *TL_messageActionPaymentSentMe) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._currency = dc.String()
	t._total_amount = dc.Long()
	t._payload = dc.Bytes(16)
	t._info = dc.Bytes(16)
	t._shipping_option_id = dc.Bytes(16)
	t._charge = dc.Bytes(16)

}

// messageActionPaymentSent#40699cd0
type TL_messageActionPaymentSent struct {
	_currency     string
	_total_amount int64
}

func (t *TL_messageActionPaymentSent) Set_currency(_currency string) {
	t._currency = _currency
}

func (t *TL_messageActionPaymentSent) Get_currency() string {
	return t._currency
}

func (t *TL_messageActionPaymentSent) Set_total_amount(_total_amount int64) {
	t._total_amount = _total_amount
}

func (t *TL_messageActionPaymentSent) Get_total_amount() int64 {
	return t._total_amount
}

func New_TL_messageActionPaymentSent() *TL_messageActionPaymentSent {
	return &TL_messageActionPaymentSent{}
}

func (t *TL_messageActionPaymentSent) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messageActionPaymentSent))
	ec.String(t.Get_currency())
	ec.Long(t.Get_total_amount())

	return ec.GetBuffer()
}

func (t *TL_messageActionPaymentSent) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._currency = dc.String()
	t._total_amount = dc.Long()

}

// messageActionPhoneCall#80e11a7f
type TL_messageActionPhoneCall struct {
	_flags    []byte
	_call_id  int64
	_reason   []byte
	_duration []byte
}

func (t *TL_messageActionPhoneCall) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_messageActionPhoneCall) Get_flags() []byte {
	return t._flags
}

func (t *TL_messageActionPhoneCall) Set_call_id(_call_id int64) {
	t._call_id = _call_id
}

func (t *TL_messageActionPhoneCall) Get_call_id() int64 {
	return t._call_id
}

func (t *TL_messageActionPhoneCall) Set_reason(_reason []byte) {
	t._reason = _reason
}

func (t *TL_messageActionPhoneCall) Get_reason() []byte {
	return t._reason
}

func (t *TL_messageActionPhoneCall) Set_duration(_duration []byte) {
	t._duration = _duration
}

func (t *TL_messageActionPhoneCall) Get_duration() []byte {
	return t._duration
}

func New_TL_messageActionPhoneCall() *TL_messageActionPhoneCall {
	return &TL_messageActionPhoneCall{}
}

func (t *TL_messageActionPhoneCall) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messageActionPhoneCall))
	ec.Long(t.Get_call_id())
	ec.Bytes(t.Get_reason())
	ec.Bytes(t.Get_duration())

	return ec.GetBuffer()
}

func (t *TL_messageActionPhoneCall) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._call_id = dc.Long()
	t._reason = dc.Bytes(16)
	t._duration = dc.Bytes(16)

}

// messageActionScreenshotTaken#4792929b
type TL_messageActionScreenshotTaken struct {
}

func New_TL_messageActionScreenshotTaken() *TL_messageActionScreenshotTaken {
	return &TL_messageActionScreenshotTaken{}
}

func (t *TL_messageActionScreenshotTaken) Encode() []byte {
	return nil
}

func (t *TL_messageActionScreenshotTaken) Decode(b []byte) {

}

// messageActionCustomAction#fae69f56
type TL_messageActionCustomAction struct {
	_message string
}

func (t *TL_messageActionCustomAction) Set_message(_message string) {
	t._message = _message
}

func (t *TL_messageActionCustomAction) Get_message() string {
	return t._message
}

func New_TL_messageActionCustomAction() *TL_messageActionCustomAction {
	return &TL_messageActionCustomAction{}
}

func (t *TL_messageActionCustomAction) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messageActionCustomAction))
	ec.String(t.Get_message())

	return ec.GetBuffer()
}

func (t *TL_messageActionCustomAction) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._message = dc.String()

}

// dialog#e4def5db
type TL_dialog struct {
	_flags                 []byte
	_pinned                []byte
	_peer                  []byte
	_top_message           int32
	_read_inbox_max_id     int32
	_read_outbox_max_id    int32
	_unread_count          int32
	_unread_mentions_count int32
	_notify_settings       []byte
	_pts                   []byte
	_draft                 []byte
}

func (t *TL_dialog) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_dialog) Get_flags() []byte {
	return t._flags
}

func (t *TL_dialog) Set_pinned(_pinned []byte) {
	t._pinned = _pinned
}

func (t *TL_dialog) Get_pinned() []byte {
	return t._pinned
}

func (t *TL_dialog) Set_peer(_peer []byte) {
	t._peer = _peer
}

func (t *TL_dialog) Get_peer() []byte {
	return t._peer
}

func (t *TL_dialog) Set_top_message(_top_message int32) {
	t._top_message = _top_message
}

func (t *TL_dialog) Get_top_message() int32 {
	return t._top_message
}

func (t *TL_dialog) Set_read_inbox_max_id(_read_inbox_max_id int32) {
	t._read_inbox_max_id = _read_inbox_max_id
}

func (t *TL_dialog) Get_read_inbox_max_id() int32 {
	return t._read_inbox_max_id
}

func (t *TL_dialog) Set_read_outbox_max_id(_read_outbox_max_id int32) {
	t._read_outbox_max_id = _read_outbox_max_id
}

func (t *TL_dialog) Get_read_outbox_max_id() int32 {
	return t._read_outbox_max_id
}

func (t *TL_dialog) Set_unread_count(_unread_count int32) {
	t._unread_count = _unread_count
}

func (t *TL_dialog) Get_unread_count() int32 {
	return t._unread_count
}

func (t *TL_dialog) Set_unread_mentions_count(_unread_mentions_count int32) {
	t._unread_mentions_count = _unread_mentions_count
}

func (t *TL_dialog) Get_unread_mentions_count() int32 {
	return t._unread_mentions_count
}

func (t *TL_dialog) Set_notify_settings(_notify_settings []byte) {
	t._notify_settings = _notify_settings
}

func (t *TL_dialog) Get_notify_settings() []byte {
	return t._notify_settings
}

func (t *TL_dialog) Set_pts(_pts []byte) {
	t._pts = _pts
}

func (t *TL_dialog) Get_pts() []byte {
	return t._pts
}

func (t *TL_dialog) Set_draft(_draft []byte) {
	t._draft = _draft
}

func (t *TL_dialog) Get_draft() []byte {
	return t._draft
}

func New_TL_dialog() *TL_dialog {
	return &TL_dialog{}
}

func (t *TL_dialog) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_dialog))
	ec.Bytes(t.Get_pinned())
	ec.Bytes(t.Get_peer())
	ec.Int(t.Get_top_message())
	ec.Int(t.Get_read_inbox_max_id())
	ec.Int(t.Get_read_outbox_max_id())
	ec.Int(t.Get_unread_count())
	ec.Int(t.Get_unread_mentions_count())
	ec.Bytes(t.Get_notify_settings())
	ec.Bytes(t.Get_pts())
	ec.Bytes(t.Get_draft())

	return ec.GetBuffer()
}

func (t *TL_dialog) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._pinned = dc.Bytes(16)
	t._peer = dc.Bytes(16)
	t._top_message = dc.Int()
	t._read_inbox_max_id = dc.Int()
	t._read_outbox_max_id = dc.Int()
	t._unread_count = dc.Int()
	t._unread_mentions_count = dc.Int()
	t._notify_settings = dc.Bytes(16)
	t._pts = dc.Bytes(16)
	t._draft = dc.Bytes(16)

}

// photoEmpty#2331b22d
type TL_photoEmpty struct {
	_id int64
}

func (t *TL_photoEmpty) Set_id(_id int64) {
	t._id = _id
}

func (t *TL_photoEmpty) Get_id() int64 {
	return t._id
}

func New_TL_photoEmpty() *TL_photoEmpty {
	return &TL_photoEmpty{}
}

func (t *TL_photoEmpty) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_photoEmpty))
	ec.Long(t.Get_id())

	return ec.GetBuffer()
}

func (t *TL_photoEmpty) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Long()

}

// photo#9288dd29
type TL_photo struct {
	_flags        []byte
	_has_stickers []byte
	_id           int64
	_access_hash  int64
	_date         int32
	_sizes        []byte
}

func (t *TL_photo) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_photo) Get_flags() []byte {
	return t._flags
}

func (t *TL_photo) Set_has_stickers(_has_stickers []byte) {
	t._has_stickers = _has_stickers
}

func (t *TL_photo) Get_has_stickers() []byte {
	return t._has_stickers
}

func (t *TL_photo) Set_id(_id int64) {
	t._id = _id
}

func (t *TL_photo) Get_id() int64 {
	return t._id
}

func (t *TL_photo) Set_access_hash(_access_hash int64) {
	t._access_hash = _access_hash
}

func (t *TL_photo) Get_access_hash() int64 {
	return t._access_hash
}

func (t *TL_photo) Set_date(_date int32) {
	t._date = _date
}

func (t *TL_photo) Get_date() int32 {
	return t._date
}

func (t *TL_photo) Set_sizes(_sizes []byte) {
	t._sizes = _sizes
}

func (t *TL_photo) Get_sizes() []byte {
	return t._sizes
}

func New_TL_photo() *TL_photo {
	return &TL_photo{}
}

func (t *TL_photo) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_photo))
	ec.Bytes(t.Get_has_stickers())
	ec.Long(t.Get_id())
	ec.Long(t.Get_access_hash())
	ec.Int(t.Get_date())
	ec.Bytes(t.Get_sizes())

	return ec.GetBuffer()
}

func (t *TL_photo) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._has_stickers = dc.Bytes(16)
	t._id = dc.Long()
	t._access_hash = dc.Long()
	t._date = dc.Int()
	t._sizes = dc.Bytes(16)

}

// photoSizeEmpty#0e17e23c
type TL_photoSizeEmpty struct {
	_type string
}

func (t *TL_photoSizeEmpty) Set_type(_type string) {
	t._type = _type
}

func (t *TL_photoSizeEmpty) Get_type() string {
	return t._type
}

func New_TL_photoSizeEmpty() *TL_photoSizeEmpty {
	return &TL_photoSizeEmpty{}
}

func (t *TL_photoSizeEmpty) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_photoSizeEmpty))
	ec.String(t.Get_type())

	return ec.GetBuffer()
}

func (t *TL_photoSizeEmpty) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._type = dc.String()

}

// photoSize#77bfb61b
type TL_photoSize struct {
	_type     string
	_location []byte
	_w        int32
	_h        int32
	_size     int32
}

func (t *TL_photoSize) Set_type(_type string) {
	t._type = _type
}

func (t *TL_photoSize) Get_type() string {
	return t._type
}

func (t *TL_photoSize) Set_location(_location []byte) {
	t._location = _location
}

func (t *TL_photoSize) Get_location() []byte {
	return t._location
}

func (t *TL_photoSize) Set_w(_w int32) {
	t._w = _w
}

func (t *TL_photoSize) Get_w() int32 {
	return t._w
}

func (t *TL_photoSize) Set_h(_h int32) {
	t._h = _h
}

func (t *TL_photoSize) Get_h() int32 {
	return t._h
}

func (t *TL_photoSize) Set_size(_size int32) {
	t._size = _size
}

func (t *TL_photoSize) Get_size() int32 {
	return t._size
}

func New_TL_photoSize() *TL_photoSize {
	return &TL_photoSize{}
}

func (t *TL_photoSize) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_photoSize))
	ec.String(t.Get_type())
	ec.Bytes(t.Get_location())
	ec.Int(t.Get_w())
	ec.Int(t.Get_h())
	ec.Int(t.Get_size())

	return ec.GetBuffer()
}

func (t *TL_photoSize) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._type = dc.String()
	t._location = dc.Bytes(16)
	t._w = dc.Int()
	t._h = dc.Int()
	t._size = dc.Int()

}

// photoCachedSize#e9a734fa
type TL_photoCachedSize struct {
	_type     string
	_location []byte
	_w        int32
	_h        int32
	_bytes    []byte
}

func (t *TL_photoCachedSize) Set_type(_type string) {
	t._type = _type
}

func (t *TL_photoCachedSize) Get_type() string {
	return t._type
}

func (t *TL_photoCachedSize) Set_location(_location []byte) {
	t._location = _location
}

func (t *TL_photoCachedSize) Get_location() []byte {
	return t._location
}

func (t *TL_photoCachedSize) Set_w(_w int32) {
	t._w = _w
}

func (t *TL_photoCachedSize) Get_w() int32 {
	return t._w
}

func (t *TL_photoCachedSize) Set_h(_h int32) {
	t._h = _h
}

func (t *TL_photoCachedSize) Get_h() int32 {
	return t._h
}

func (t *TL_photoCachedSize) Set_bytes(_bytes []byte) {
	t._bytes = _bytes
}

func (t *TL_photoCachedSize) Get_bytes() []byte {
	return t._bytes
}

func New_TL_photoCachedSize() *TL_photoCachedSize {
	return &TL_photoCachedSize{}
}

func (t *TL_photoCachedSize) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_photoCachedSize))
	ec.String(t.Get_type())
	ec.Bytes(t.Get_location())
	ec.Int(t.Get_w())
	ec.Int(t.Get_h())
	ec.Bytes(t.Get_bytes())

	return ec.GetBuffer()
}

func (t *TL_photoCachedSize) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._type = dc.String()
	t._location = dc.Bytes(16)
	t._w = dc.Int()
	t._h = dc.Int()
	t._bytes = dc.Bytes(16)

}

// geoPointEmpty#1117dd5f
type TL_geoPointEmpty struct {
}

func New_TL_geoPointEmpty() *TL_geoPointEmpty {
	return &TL_geoPointEmpty{}
}

func (t *TL_geoPointEmpty) Encode() []byte {
	return nil
}

func (t *TL_geoPointEmpty) Decode(b []byte) {

}

// geoPoint#2049d70c
type TL_geoPoint struct {
	_long float64
	_lat  float64
}

func (t *TL_geoPoint) Set_long(_long float64) {
	t._long = _long
}

func (t *TL_geoPoint) Get_long() float64 {
	return t._long
}

func (t *TL_geoPoint) Set_lat(_lat float64) {
	t._lat = _lat
}

func (t *TL_geoPoint) Get_lat() float64 {
	return t._lat
}

func New_TL_geoPoint() *TL_geoPoint {
	return &TL_geoPoint{}
}

func (t *TL_geoPoint) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_geoPoint))
	ec.Double(t.Get_long())
	ec.Double(t.Get_lat())

	return ec.GetBuffer()
}

func (t *TL_geoPoint) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._long = dc.Double()
	t._lat = dc.Double()

}

// auth_checkedPhone#811ea28e
type TL_auth_checkedPhone struct {
	_phone_registered bool
}

func (t *TL_auth_checkedPhone) Set_phone_registered(_phone_registered bool) {
	t._phone_registered = _phone_registered
}

func (t *TL_auth_checkedPhone) Get_phone_registered() bool {
	return t._phone_registered
}

func New_TL_auth_checkedPhone() *TL_auth_checkedPhone {
	return &TL_auth_checkedPhone{}
}

func (t *TL_auth_checkedPhone) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_auth_checkedPhone))
	ec.Bool(t.Get_phone_registered())

	return ec.GetBuffer()
}

func (t *TL_auth_checkedPhone) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._phone_registered = dc.Bool()

}

// auth_sentCode#5e002502
type TL_auth_sentCode struct {
	_flags            []byte
	_phone_registered []byte
	_type             []byte
	_phone_code_hash  string
	_next_type        []byte
	_timeout          []byte
}

func (t *TL_auth_sentCode) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_auth_sentCode) Get_flags() []byte {
	return t._flags
}

func (t *TL_auth_sentCode) Set_phone_registered(_phone_registered []byte) {
	t._phone_registered = _phone_registered
}

func (t *TL_auth_sentCode) Get_phone_registered() []byte {
	return t._phone_registered
}

func (t *TL_auth_sentCode) Set_type(_type []byte) {
	t._type = _type
}

func (t *TL_auth_sentCode) Get_type() []byte {
	return t._type
}

func (t *TL_auth_sentCode) Set_phone_code_hash(_phone_code_hash string) {
	t._phone_code_hash = _phone_code_hash
}

func (t *TL_auth_sentCode) Get_phone_code_hash() string {
	return t._phone_code_hash
}

func (t *TL_auth_sentCode) Set_next_type(_next_type []byte) {
	t._next_type = _next_type
}

func (t *TL_auth_sentCode) Get_next_type() []byte {
	return t._next_type
}

func (t *TL_auth_sentCode) Set_timeout(_timeout []byte) {
	t._timeout = _timeout
}

func (t *TL_auth_sentCode) Get_timeout() []byte {
	return t._timeout
}

func New_TL_auth_sentCode() *TL_auth_sentCode {
	return &TL_auth_sentCode{}
}

func (t *TL_auth_sentCode) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_auth_sentCode))
	ec.Bytes(t.Get_phone_registered())
	ec.Bytes(t.Get_type())
	ec.String(t.Get_phone_code_hash())
	ec.Bytes(t.Get_next_type())
	ec.Bytes(t.Get_timeout())

	return ec.GetBuffer()
}

func (t *TL_auth_sentCode) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._phone_registered = dc.Bytes(16)
	t._type = dc.Bytes(16)
	t._phone_code_hash = dc.String()
	t._next_type = dc.Bytes(16)
	t._timeout = dc.Bytes(16)

}

// auth_authorization#cd050916
type TL_auth_authorization struct {
	_flags        []byte
	_tmp_sessions []byte
	_user         []byte
}

func (t *TL_auth_authorization) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_auth_authorization) Get_flags() []byte {
	return t._flags
}

func (t *TL_auth_authorization) Set_tmp_sessions(_tmp_sessions []byte) {
	t._tmp_sessions = _tmp_sessions
}

func (t *TL_auth_authorization) Get_tmp_sessions() []byte {
	return t._tmp_sessions
}

func (t *TL_auth_authorization) Set_user(_user []byte) {
	t._user = _user
}

func (t *TL_auth_authorization) Get_user() []byte {
	return t._user
}

func New_TL_auth_authorization() *TL_auth_authorization {
	return &TL_auth_authorization{}
}

func (t *TL_auth_authorization) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_auth_authorization))
	ec.Bytes(t.Get_tmp_sessions())
	ec.Bytes(t.Get_user())

	return ec.GetBuffer()
}

func (t *TL_auth_authorization) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._tmp_sessions = dc.Bytes(16)
	t._user = dc.Bytes(16)

}

// auth_exportedAuthorization#df969c2d
type TL_auth_exportedAuthorization struct {
	_id    int32
	_bytes []byte
}

func (t *TL_auth_exportedAuthorization) Set_id(_id int32) {
	t._id = _id
}

func (t *TL_auth_exportedAuthorization) Get_id() int32 {
	return t._id
}

func (t *TL_auth_exportedAuthorization) Set_bytes(_bytes []byte) {
	t._bytes = _bytes
}

func (t *TL_auth_exportedAuthorization) Get_bytes() []byte {
	return t._bytes
}

func New_TL_auth_exportedAuthorization() *TL_auth_exportedAuthorization {
	return &TL_auth_exportedAuthorization{}
}

func (t *TL_auth_exportedAuthorization) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_auth_exportedAuthorization))
	ec.Int(t.Get_id())
	ec.Bytes(t.Get_bytes())

	return ec.GetBuffer()
}

func (t *TL_auth_exportedAuthorization) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Int()
	t._bytes = dc.Bytes(16)

}

// inputNotifyPeer#b8bc5b0c
type TL_inputNotifyPeer struct {
	_peer []byte
}

func (t *TL_inputNotifyPeer) Set_peer(_peer []byte) {
	t._peer = _peer
}

func (t *TL_inputNotifyPeer) Get_peer() []byte {
	return t._peer
}

func New_TL_inputNotifyPeer() *TL_inputNotifyPeer {
	return &TL_inputNotifyPeer{}
}

func (t *TL_inputNotifyPeer) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputNotifyPeer))
	ec.Bytes(t.Get_peer())

	return ec.GetBuffer()
}

func (t *TL_inputNotifyPeer) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._peer = dc.Bytes(16)

}

// inputNotifyUsers#193b4417
type TL_inputNotifyUsers struct {
}

func New_TL_inputNotifyUsers() *TL_inputNotifyUsers {
	return &TL_inputNotifyUsers{}
}

func (t *TL_inputNotifyUsers) Encode() []byte {
	return nil
}

func (t *TL_inputNotifyUsers) Decode(b []byte) {

}

// inputNotifyChats#4a95e84e
type TL_inputNotifyChats struct {
}

func New_TL_inputNotifyChats() *TL_inputNotifyChats {
	return &TL_inputNotifyChats{}
}

func (t *TL_inputNotifyChats) Encode() []byte {
	return nil
}

func (t *TL_inputNotifyChats) Decode(b []byte) {

}

// inputNotifyAll#a429b886
type TL_inputNotifyAll struct {
}

func New_TL_inputNotifyAll() *TL_inputNotifyAll {
	return &TL_inputNotifyAll{}
}

func (t *TL_inputNotifyAll) Encode() []byte {
	return nil
}

func (t *TL_inputNotifyAll) Decode(b []byte) {

}

// inputPeerNotifyEventsEmpty#f03064d8
type TL_inputPeerNotifyEventsEmpty struct {
}

func New_TL_inputPeerNotifyEventsEmpty() *TL_inputPeerNotifyEventsEmpty {
	return &TL_inputPeerNotifyEventsEmpty{}
}

func (t *TL_inputPeerNotifyEventsEmpty) Encode() []byte {
	return nil
}

func (t *TL_inputPeerNotifyEventsEmpty) Decode(b []byte) {

}

// inputPeerNotifyEventsAll#e86a2c74
type TL_inputPeerNotifyEventsAll struct {
}

func New_TL_inputPeerNotifyEventsAll() *TL_inputPeerNotifyEventsAll {
	return &TL_inputPeerNotifyEventsAll{}
}

func (t *TL_inputPeerNotifyEventsAll) Encode() []byte {
	return nil
}

func (t *TL_inputPeerNotifyEventsAll) Decode(b []byte) {

}

// inputPeerNotifySettings#38935eb2
type TL_inputPeerNotifySettings struct {
	_flags         []byte
	_show_previews []byte
	_silent        []byte
	_mute_until    int32
	_sound         string
}

func (t *TL_inputPeerNotifySettings) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_inputPeerNotifySettings) Get_flags() []byte {
	return t._flags
}

func (t *TL_inputPeerNotifySettings) Set_show_previews(_show_previews []byte) {
	t._show_previews = _show_previews
}

func (t *TL_inputPeerNotifySettings) Get_show_previews() []byte {
	return t._show_previews
}

func (t *TL_inputPeerNotifySettings) Set_silent(_silent []byte) {
	t._silent = _silent
}

func (t *TL_inputPeerNotifySettings) Get_silent() []byte {
	return t._silent
}

func (t *TL_inputPeerNotifySettings) Set_mute_until(_mute_until int32) {
	t._mute_until = _mute_until
}

func (t *TL_inputPeerNotifySettings) Get_mute_until() int32 {
	return t._mute_until
}

func (t *TL_inputPeerNotifySettings) Set_sound(_sound string) {
	t._sound = _sound
}

func (t *TL_inputPeerNotifySettings) Get_sound() string {
	return t._sound
}

func New_TL_inputPeerNotifySettings() *TL_inputPeerNotifySettings {
	return &TL_inputPeerNotifySettings{}
}

func (t *TL_inputPeerNotifySettings) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputPeerNotifySettings))
	ec.Bytes(t.Get_show_previews())
	ec.Bytes(t.Get_silent())
	ec.Int(t.Get_mute_until())
	ec.String(t.Get_sound())

	return ec.GetBuffer()
}

func (t *TL_inputPeerNotifySettings) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._show_previews = dc.Bytes(16)
	t._silent = dc.Bytes(16)
	t._mute_until = dc.Int()
	t._sound = dc.String()

}

// peerNotifyEventsEmpty#add53cb3
type TL_peerNotifyEventsEmpty struct {
}

func New_TL_peerNotifyEventsEmpty() *TL_peerNotifyEventsEmpty {
	return &TL_peerNotifyEventsEmpty{}
}

func (t *TL_peerNotifyEventsEmpty) Encode() []byte {
	return nil
}

func (t *TL_peerNotifyEventsEmpty) Decode(b []byte) {

}

// peerNotifyEventsAll#6d1ded88
type TL_peerNotifyEventsAll struct {
}

func New_TL_peerNotifyEventsAll() *TL_peerNotifyEventsAll {
	return &TL_peerNotifyEventsAll{}
}

func (t *TL_peerNotifyEventsAll) Encode() []byte {
	return nil
}

func (t *TL_peerNotifyEventsAll) Decode(b []byte) {

}

// peerNotifySettingsEmpty#70a68512
type TL_peerNotifySettingsEmpty struct {
}

func New_TL_peerNotifySettingsEmpty() *TL_peerNotifySettingsEmpty {
	return &TL_peerNotifySettingsEmpty{}
}

func (t *TL_peerNotifySettingsEmpty) Encode() []byte {
	return nil
}

func (t *TL_peerNotifySettingsEmpty) Decode(b []byte) {

}

// peerNotifySettings#9acda4c0
type TL_peerNotifySettings struct {
	_flags         []byte
	_show_previews []byte
	_silent        []byte
	_mute_until    int32
	_sound         string
}

func (t *TL_peerNotifySettings) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_peerNotifySettings) Get_flags() []byte {
	return t._flags
}

func (t *TL_peerNotifySettings) Set_show_previews(_show_previews []byte) {
	t._show_previews = _show_previews
}

func (t *TL_peerNotifySettings) Get_show_previews() []byte {
	return t._show_previews
}

func (t *TL_peerNotifySettings) Set_silent(_silent []byte) {
	t._silent = _silent
}

func (t *TL_peerNotifySettings) Get_silent() []byte {
	return t._silent
}

func (t *TL_peerNotifySettings) Set_mute_until(_mute_until int32) {
	t._mute_until = _mute_until
}

func (t *TL_peerNotifySettings) Get_mute_until() int32 {
	return t._mute_until
}

func (t *TL_peerNotifySettings) Set_sound(_sound string) {
	t._sound = _sound
}

func (t *TL_peerNotifySettings) Get_sound() string {
	return t._sound
}

func New_TL_peerNotifySettings() *TL_peerNotifySettings {
	return &TL_peerNotifySettings{}
}

func (t *TL_peerNotifySettings) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_peerNotifySettings))
	ec.Bytes(t.Get_show_previews())
	ec.Bytes(t.Get_silent())
	ec.Int(t.Get_mute_until())
	ec.String(t.Get_sound())

	return ec.GetBuffer()
}

func (t *TL_peerNotifySettings) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._show_previews = dc.Bytes(16)
	t._silent = dc.Bytes(16)
	t._mute_until = dc.Int()
	t._sound = dc.String()

}

// peerSettings#818426cd
type TL_peerSettings struct {
	_flags       []byte
	_report_spam []byte
}

func (t *TL_peerSettings) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_peerSettings) Get_flags() []byte {
	return t._flags
}

func (t *TL_peerSettings) Set_report_spam(_report_spam []byte) {
	t._report_spam = _report_spam
}

func (t *TL_peerSettings) Get_report_spam() []byte {
	return t._report_spam
}

func New_TL_peerSettings() *TL_peerSettings {
	return &TL_peerSettings{}
}

func (t *TL_peerSettings) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_peerSettings))
	ec.Bytes(t.Get_report_spam())

	return ec.GetBuffer()
}

func (t *TL_peerSettings) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._report_spam = dc.Bytes(16)

}

// wallPaper#ccb03657
type TL_wallPaper struct {
	_id    int32
	_title string
	_sizes []byte
	_color int32
}

func (t *TL_wallPaper) Set_id(_id int32) {
	t._id = _id
}

func (t *TL_wallPaper) Get_id() int32 {
	return t._id
}

func (t *TL_wallPaper) Set_title(_title string) {
	t._title = _title
}

func (t *TL_wallPaper) Get_title() string {
	return t._title
}

func (t *TL_wallPaper) Set_sizes(_sizes []byte) {
	t._sizes = _sizes
}

func (t *TL_wallPaper) Get_sizes() []byte {
	return t._sizes
}

func (t *TL_wallPaper) Set_color(_color int32) {
	t._color = _color
}

func (t *TL_wallPaper) Get_color() int32 {
	return t._color
}

func New_TL_wallPaper() *TL_wallPaper {
	return &TL_wallPaper{}
}

func (t *TL_wallPaper) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_wallPaper))
	ec.Int(t.Get_id())
	ec.String(t.Get_title())
	ec.Bytes(t.Get_sizes())
	ec.Int(t.Get_color())

	return ec.GetBuffer()
}

func (t *TL_wallPaper) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Int()
	t._title = dc.String()
	t._sizes = dc.Bytes(16)
	t._color = dc.Int()

}

// wallPaperSolid#63117f24
type TL_wallPaperSolid struct {
	_id       int32
	_title    string
	_bg_color int32
	_color    int32
}

func (t *TL_wallPaperSolid) Set_id(_id int32) {
	t._id = _id
}

func (t *TL_wallPaperSolid) Get_id() int32 {
	return t._id
}

func (t *TL_wallPaperSolid) Set_title(_title string) {
	t._title = _title
}

func (t *TL_wallPaperSolid) Get_title() string {
	return t._title
}

func (t *TL_wallPaperSolid) Set_bg_color(_bg_color int32) {
	t._bg_color = _bg_color
}

func (t *TL_wallPaperSolid) Get_bg_color() int32 {
	return t._bg_color
}

func (t *TL_wallPaperSolid) Set_color(_color int32) {
	t._color = _color
}

func (t *TL_wallPaperSolid) Get_color() int32 {
	return t._color
}

func New_TL_wallPaperSolid() *TL_wallPaperSolid {
	return &TL_wallPaperSolid{}
}

func (t *TL_wallPaperSolid) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_wallPaperSolid))
	ec.Int(t.Get_id())
	ec.String(t.Get_title())
	ec.Int(t.Get_bg_color())
	ec.Int(t.Get_color())

	return ec.GetBuffer()
}

func (t *TL_wallPaperSolid) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Int()
	t._title = dc.String()
	t._bg_color = dc.Int()
	t._color = dc.Int()

}

// inputReportReasonSpam#58dbcab8
type TL_inputReportReasonSpam struct {
}

func New_TL_inputReportReasonSpam() *TL_inputReportReasonSpam {
	return &TL_inputReportReasonSpam{}
}

func (t *TL_inputReportReasonSpam) Encode() []byte {
	return nil
}

func (t *TL_inputReportReasonSpam) Decode(b []byte) {

}

// inputReportReasonViolence#1e22c78d
type TL_inputReportReasonViolence struct {
}

func New_TL_inputReportReasonViolence() *TL_inputReportReasonViolence {
	return &TL_inputReportReasonViolence{}
}

func (t *TL_inputReportReasonViolence) Encode() []byte {
	return nil
}

func (t *TL_inputReportReasonViolence) Decode(b []byte) {

}

// inputReportReasonPornography#2e59d922
type TL_inputReportReasonPornography struct {
}

func New_TL_inputReportReasonPornography() *TL_inputReportReasonPornography {
	return &TL_inputReportReasonPornography{}
}

func (t *TL_inputReportReasonPornography) Encode() []byte {
	return nil
}

func (t *TL_inputReportReasonPornography) Decode(b []byte) {

}

// inputReportReasonOther#e1746d0a
type TL_inputReportReasonOther struct {
	_text string
}

func (t *TL_inputReportReasonOther) Set_text(_text string) {
	t._text = _text
}

func (t *TL_inputReportReasonOther) Get_text() string {
	return t._text
}

func New_TL_inputReportReasonOther() *TL_inputReportReasonOther {
	return &TL_inputReportReasonOther{}
}

func (t *TL_inputReportReasonOther) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputReportReasonOther))
	ec.String(t.Get_text())

	return ec.GetBuffer()
}

func (t *TL_inputReportReasonOther) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._text = dc.String()

}

// userFull#0f220f3f
type TL_userFull struct {
	_flags                 []byte
	_blocked               []byte
	_phone_calls_available []byte
	_phone_calls_private   []byte
	_user                  []byte
	_about                 []byte
	_link                  []byte
	_profile_photo         []byte
	_notify_settings       []byte
	_bot_info              []byte
	_common_chats_count    int32
}

func (t *TL_userFull) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_userFull) Get_flags() []byte {
	return t._flags
}

func (t *TL_userFull) Set_blocked(_blocked []byte) {
	t._blocked = _blocked
}

func (t *TL_userFull) Get_blocked() []byte {
	return t._blocked
}

func (t *TL_userFull) Set_phone_calls_available(_phone_calls_available []byte) {
	t._phone_calls_available = _phone_calls_available
}

func (t *TL_userFull) Get_phone_calls_available() []byte {
	return t._phone_calls_available
}

func (t *TL_userFull) Set_phone_calls_private(_phone_calls_private []byte) {
	t._phone_calls_private = _phone_calls_private
}

func (t *TL_userFull) Get_phone_calls_private() []byte {
	return t._phone_calls_private
}

func (t *TL_userFull) Set_user(_user []byte) {
	t._user = _user
}

func (t *TL_userFull) Get_user() []byte {
	return t._user
}

func (t *TL_userFull) Set_about(_about []byte) {
	t._about = _about
}

func (t *TL_userFull) Get_about() []byte {
	return t._about
}

func (t *TL_userFull) Set_link(_link []byte) {
	t._link = _link
}

func (t *TL_userFull) Get_link() []byte {
	return t._link
}

func (t *TL_userFull) Set_profile_photo(_profile_photo []byte) {
	t._profile_photo = _profile_photo
}

func (t *TL_userFull) Get_profile_photo() []byte {
	return t._profile_photo
}

func (t *TL_userFull) Set_notify_settings(_notify_settings []byte) {
	t._notify_settings = _notify_settings
}

func (t *TL_userFull) Get_notify_settings() []byte {
	return t._notify_settings
}

func (t *TL_userFull) Set_bot_info(_bot_info []byte) {
	t._bot_info = _bot_info
}

func (t *TL_userFull) Get_bot_info() []byte {
	return t._bot_info
}

func (t *TL_userFull) Set_common_chats_count(_common_chats_count int32) {
	t._common_chats_count = _common_chats_count
}

func (t *TL_userFull) Get_common_chats_count() int32 {
	return t._common_chats_count
}

func New_TL_userFull() *TL_userFull {
	return &TL_userFull{}
}

func (t *TL_userFull) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_userFull))
	ec.Bytes(t.Get_blocked())
	ec.Bytes(t.Get_phone_calls_available())
	ec.Bytes(t.Get_phone_calls_private())
	ec.Bytes(t.Get_user())
	ec.Bytes(t.Get_about())
	ec.Bytes(t.Get_link())
	ec.Bytes(t.Get_profile_photo())
	ec.Bytes(t.Get_notify_settings())
	ec.Bytes(t.Get_bot_info())
	ec.Int(t.Get_common_chats_count())

	return ec.GetBuffer()
}

func (t *TL_userFull) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._blocked = dc.Bytes(16)
	t._phone_calls_available = dc.Bytes(16)
	t._phone_calls_private = dc.Bytes(16)
	t._user = dc.Bytes(16)
	t._about = dc.Bytes(16)
	t._link = dc.Bytes(16)
	t._profile_photo = dc.Bytes(16)
	t._notify_settings = dc.Bytes(16)
	t._bot_info = dc.Bytes(16)
	t._common_chats_count = dc.Int()

}

// contact#f911c994
type TL_contact struct {
	_user_id int32
	_mutual  bool
}

func (t *TL_contact) Set_user_id(_user_id int32) {
	t._user_id = _user_id
}

func (t *TL_contact) Get_user_id() int32 {
	return t._user_id
}

func (t *TL_contact) Set_mutual(_mutual bool) {
	t._mutual = _mutual
}

func (t *TL_contact) Get_mutual() bool {
	return t._mutual
}

func New_TL_contact() *TL_contact {
	return &TL_contact{}
}

func (t *TL_contact) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_contact))
	ec.Int(t.Get_user_id())
	ec.Bool(t.Get_mutual())

	return ec.GetBuffer()
}

func (t *TL_contact) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._user_id = dc.Int()
	t._mutual = dc.Bool()

}

// importedContact#d0028438
type TL_importedContact struct {
	_user_id   int32
	_client_id int64
}

func (t *TL_importedContact) Set_user_id(_user_id int32) {
	t._user_id = _user_id
}

func (t *TL_importedContact) Get_user_id() int32 {
	return t._user_id
}

func (t *TL_importedContact) Set_client_id(_client_id int64) {
	t._client_id = _client_id
}

func (t *TL_importedContact) Get_client_id() int64 {
	return t._client_id
}

func New_TL_importedContact() *TL_importedContact {
	return &TL_importedContact{}
}

func (t *TL_importedContact) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_importedContact))
	ec.Int(t.Get_user_id())
	ec.Long(t.Get_client_id())

	return ec.GetBuffer()
}

func (t *TL_importedContact) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._user_id = dc.Int()
	t._client_id = dc.Long()

}

// contactBlocked#561bc879
type TL_contactBlocked struct {
	_user_id int32
	_date    int32
}

func (t *TL_contactBlocked) Set_user_id(_user_id int32) {
	t._user_id = _user_id
}

func (t *TL_contactBlocked) Get_user_id() int32 {
	return t._user_id
}

func (t *TL_contactBlocked) Set_date(_date int32) {
	t._date = _date
}

func (t *TL_contactBlocked) Get_date() int32 {
	return t._date
}

func New_TL_contactBlocked() *TL_contactBlocked {
	return &TL_contactBlocked{}
}

func (t *TL_contactBlocked) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_contactBlocked))
	ec.Int(t.Get_user_id())
	ec.Int(t.Get_date())

	return ec.GetBuffer()
}

func (t *TL_contactBlocked) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._user_id = dc.Int()
	t._date = dc.Int()

}

// contactStatus#d3680c61
type TL_contactStatus struct {
	_user_id int32
	_status  []byte
}

func (t *TL_contactStatus) Set_user_id(_user_id int32) {
	t._user_id = _user_id
}

func (t *TL_contactStatus) Get_user_id() int32 {
	return t._user_id
}

func (t *TL_contactStatus) Set_status(_status []byte) {
	t._status = _status
}

func (t *TL_contactStatus) Get_status() []byte {
	return t._status
}

func New_TL_contactStatus() *TL_contactStatus {
	return &TL_contactStatus{}
}

func (t *TL_contactStatus) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_contactStatus))
	ec.Int(t.Get_user_id())
	ec.Bytes(t.Get_status())

	return ec.GetBuffer()
}

func (t *TL_contactStatus) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._user_id = dc.Int()
	t._status = dc.Bytes(16)

}

// contacts_link#3ace484c
type TL_contacts_link struct {
	_my_link      []byte
	_foreign_link []byte
	_user         []byte
}

func (t *TL_contacts_link) Set_my_link(_my_link []byte) {
	t._my_link = _my_link
}

func (t *TL_contacts_link) Get_my_link() []byte {
	return t._my_link
}

func (t *TL_contacts_link) Set_foreign_link(_foreign_link []byte) {
	t._foreign_link = _foreign_link
}

func (t *TL_contacts_link) Get_foreign_link() []byte {
	return t._foreign_link
}

func (t *TL_contacts_link) Set_user(_user []byte) {
	t._user = _user
}

func (t *TL_contacts_link) Get_user() []byte {
	return t._user
}

func New_TL_contacts_link() *TL_contacts_link {
	return &TL_contacts_link{}
}

func (t *TL_contacts_link) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_contacts_link))
	ec.Bytes(t.Get_my_link())
	ec.Bytes(t.Get_foreign_link())
	ec.Bytes(t.Get_user())

	return ec.GetBuffer()
}

func (t *TL_contacts_link) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._my_link = dc.Bytes(16)
	t._foreign_link = dc.Bytes(16)
	t._user = dc.Bytes(16)

}

// contacts_contactsNotModified#b74ba9d2
type TL_contacts_contactsNotModified struct {
}

func New_TL_contacts_contactsNotModified() *TL_contacts_contactsNotModified {
	return &TL_contacts_contactsNotModified{}
}

func (t *TL_contacts_contactsNotModified) Encode() []byte {
	return nil
}

func (t *TL_contacts_contactsNotModified) Decode(b []byte) {

}

// contacts_contacts#eae87e42
type TL_contacts_contacts struct {
	_contacts    []byte
	_saved_count int32
	_users       []byte
}

func (t *TL_contacts_contacts) Set_contacts(_contacts []byte) {
	t._contacts = _contacts
}

func (t *TL_contacts_contacts) Get_contacts() []byte {
	return t._contacts
}

func (t *TL_contacts_contacts) Set_saved_count(_saved_count int32) {
	t._saved_count = _saved_count
}

func (t *TL_contacts_contacts) Get_saved_count() int32 {
	return t._saved_count
}

func (t *TL_contacts_contacts) Set_users(_users []byte) {
	t._users = _users
}

func (t *TL_contacts_contacts) Get_users() []byte {
	return t._users
}

func New_TL_contacts_contacts() *TL_contacts_contacts {
	return &TL_contacts_contacts{}
}

func (t *TL_contacts_contacts) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_contacts_contacts))
	ec.Bytes(t.Get_contacts())
	ec.Int(t.Get_saved_count())
	ec.Bytes(t.Get_users())

	return ec.GetBuffer()
}

func (t *TL_contacts_contacts) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._contacts = dc.Bytes(16)
	t._saved_count = dc.Int()
	t._users = dc.Bytes(16)

}

// contacts_importedContacts#77d01c3b
type TL_contacts_importedContacts struct {
	_imported        []byte
	_popular_invites []byte
	_retry_contacts  []byte
	_users           []byte
}

func (t *TL_contacts_importedContacts) Set_imported(_imported []byte) {
	t._imported = _imported
}

func (t *TL_contacts_importedContacts) Get_imported() []byte {
	return t._imported
}

func (t *TL_contacts_importedContacts) Set_popular_invites(_popular_invites []byte) {
	t._popular_invites = _popular_invites
}

func (t *TL_contacts_importedContacts) Get_popular_invites() []byte {
	return t._popular_invites
}

func (t *TL_contacts_importedContacts) Set_retry_contacts(_retry_contacts []byte) {
	t._retry_contacts = _retry_contacts
}

func (t *TL_contacts_importedContacts) Get_retry_contacts() []byte {
	return t._retry_contacts
}

func (t *TL_contacts_importedContacts) Set_users(_users []byte) {
	t._users = _users
}

func (t *TL_contacts_importedContacts) Get_users() []byte {
	return t._users
}

func New_TL_contacts_importedContacts() *TL_contacts_importedContacts {
	return &TL_contacts_importedContacts{}
}

func (t *TL_contacts_importedContacts) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_contacts_importedContacts))
	ec.Bytes(t.Get_imported())
	ec.Bytes(t.Get_popular_invites())
	ec.Bytes(t.Get_retry_contacts())
	ec.Bytes(t.Get_users())

	return ec.GetBuffer()
}

func (t *TL_contacts_importedContacts) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._imported = dc.Bytes(16)
	t._popular_invites = dc.Bytes(16)
	t._retry_contacts = dc.Bytes(16)
	t._users = dc.Bytes(16)

}

// contacts_blocked#1c138d15
type TL_contacts_blocked struct {
	_blocked []byte
	_users   []byte
}

func (t *TL_contacts_blocked) Set_blocked(_blocked []byte) {
	t._blocked = _blocked
}

func (t *TL_contacts_blocked) Get_blocked() []byte {
	return t._blocked
}

func (t *TL_contacts_blocked) Set_users(_users []byte) {
	t._users = _users
}

func (t *TL_contacts_blocked) Get_users() []byte {
	return t._users
}

func New_TL_contacts_blocked() *TL_contacts_blocked {
	return &TL_contacts_blocked{}
}

func (t *TL_contacts_blocked) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_contacts_blocked))
	ec.Bytes(t.Get_blocked())
	ec.Bytes(t.Get_users())

	return ec.GetBuffer()
}

func (t *TL_contacts_blocked) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._blocked = dc.Bytes(16)
	t._users = dc.Bytes(16)

}

// contacts_blockedSlice#900802a1
type TL_contacts_blockedSlice struct {
	_count   int32
	_blocked []byte
	_users   []byte
}

func (t *TL_contacts_blockedSlice) Set_count(_count int32) {
	t._count = _count
}

func (t *TL_contacts_blockedSlice) Get_count() int32 {
	return t._count
}

func (t *TL_contacts_blockedSlice) Set_blocked(_blocked []byte) {
	t._blocked = _blocked
}

func (t *TL_contacts_blockedSlice) Get_blocked() []byte {
	return t._blocked
}

func (t *TL_contacts_blockedSlice) Set_users(_users []byte) {
	t._users = _users
}

func (t *TL_contacts_blockedSlice) Get_users() []byte {
	return t._users
}

func New_TL_contacts_blockedSlice() *TL_contacts_blockedSlice {
	return &TL_contacts_blockedSlice{}
}

func (t *TL_contacts_blockedSlice) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_contacts_blockedSlice))
	ec.Int(t.Get_count())
	ec.Bytes(t.Get_blocked())
	ec.Bytes(t.Get_users())

	return ec.GetBuffer()
}

func (t *TL_contacts_blockedSlice) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._count = dc.Int()
	t._blocked = dc.Bytes(16)
	t._users = dc.Bytes(16)

}

// messages_dialogs#15ba6c40
type TL_messages_dialogs struct {
	_dialogs  []byte
	_messages []byte
	_chats    []byte
	_users    []byte
}

func (t *TL_messages_dialogs) Set_dialogs(_dialogs []byte) {
	t._dialogs = _dialogs
}

func (t *TL_messages_dialogs) Get_dialogs() []byte {
	return t._dialogs
}

func (t *TL_messages_dialogs) Set_messages(_messages []byte) {
	t._messages = _messages
}

func (t *TL_messages_dialogs) Get_messages() []byte {
	return t._messages
}

func (t *TL_messages_dialogs) Set_chats(_chats []byte) {
	t._chats = _chats
}

func (t *TL_messages_dialogs) Get_chats() []byte {
	return t._chats
}

func (t *TL_messages_dialogs) Set_users(_users []byte) {
	t._users = _users
}

func (t *TL_messages_dialogs) Get_users() []byte {
	return t._users
}

func New_TL_messages_dialogs() *TL_messages_dialogs {
	return &TL_messages_dialogs{}
}

func (t *TL_messages_dialogs) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_dialogs))
	ec.Bytes(t.Get_dialogs())
	ec.Bytes(t.Get_messages())
	ec.Bytes(t.Get_chats())
	ec.Bytes(t.Get_users())

	return ec.GetBuffer()
}

func (t *TL_messages_dialogs) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._dialogs = dc.Bytes(16)
	t._messages = dc.Bytes(16)
	t._chats = dc.Bytes(16)
	t._users = dc.Bytes(16)

}

// messages_dialogsSlice#71e094f3
type TL_messages_dialogsSlice struct {
	_count    int32
	_dialogs  []byte
	_messages []byte
	_chats    []byte
	_users    []byte
}

func (t *TL_messages_dialogsSlice) Set_count(_count int32) {
	t._count = _count
}

func (t *TL_messages_dialogsSlice) Get_count() int32 {
	return t._count
}

func (t *TL_messages_dialogsSlice) Set_dialogs(_dialogs []byte) {
	t._dialogs = _dialogs
}

func (t *TL_messages_dialogsSlice) Get_dialogs() []byte {
	return t._dialogs
}

func (t *TL_messages_dialogsSlice) Set_messages(_messages []byte) {
	t._messages = _messages
}

func (t *TL_messages_dialogsSlice) Get_messages() []byte {
	return t._messages
}

func (t *TL_messages_dialogsSlice) Set_chats(_chats []byte) {
	t._chats = _chats
}

func (t *TL_messages_dialogsSlice) Get_chats() []byte {
	return t._chats
}

func (t *TL_messages_dialogsSlice) Set_users(_users []byte) {
	t._users = _users
}

func (t *TL_messages_dialogsSlice) Get_users() []byte {
	return t._users
}

func New_TL_messages_dialogsSlice() *TL_messages_dialogsSlice {
	return &TL_messages_dialogsSlice{}
}

func (t *TL_messages_dialogsSlice) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_dialogsSlice))
	ec.Int(t.Get_count())
	ec.Bytes(t.Get_dialogs())
	ec.Bytes(t.Get_messages())
	ec.Bytes(t.Get_chats())
	ec.Bytes(t.Get_users())

	return ec.GetBuffer()
}

func (t *TL_messages_dialogsSlice) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._count = dc.Int()
	t._dialogs = dc.Bytes(16)
	t._messages = dc.Bytes(16)
	t._chats = dc.Bytes(16)
	t._users = dc.Bytes(16)

}

// messages_messages#8c718e87
type TL_messages_messages struct {
	_messages []byte
	_chats    []byte
	_users    []byte
}

func (t *TL_messages_messages) Set_messages(_messages []byte) {
	t._messages = _messages
}

func (t *TL_messages_messages) Get_messages() []byte {
	return t._messages
}

func (t *TL_messages_messages) Set_chats(_chats []byte) {
	t._chats = _chats
}

func (t *TL_messages_messages) Get_chats() []byte {
	return t._chats
}

func (t *TL_messages_messages) Set_users(_users []byte) {
	t._users = _users
}

func (t *TL_messages_messages) Get_users() []byte {
	return t._users
}

func New_TL_messages_messages() *TL_messages_messages {
	return &TL_messages_messages{}
}

func (t *TL_messages_messages) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_messages))
	ec.Bytes(t.Get_messages())
	ec.Bytes(t.Get_chats())
	ec.Bytes(t.Get_users())

	return ec.GetBuffer()
}

func (t *TL_messages_messages) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._messages = dc.Bytes(16)
	t._chats = dc.Bytes(16)
	t._users = dc.Bytes(16)

}

// messages_messagesSlice#0b446ae3
type TL_messages_messagesSlice struct {
	_count    int32
	_messages []byte
	_chats    []byte
	_users    []byte
}

func (t *TL_messages_messagesSlice) Set_count(_count int32) {
	t._count = _count
}

func (t *TL_messages_messagesSlice) Get_count() int32 {
	return t._count
}

func (t *TL_messages_messagesSlice) Set_messages(_messages []byte) {
	t._messages = _messages
}

func (t *TL_messages_messagesSlice) Get_messages() []byte {
	return t._messages
}

func (t *TL_messages_messagesSlice) Set_chats(_chats []byte) {
	t._chats = _chats
}

func (t *TL_messages_messagesSlice) Get_chats() []byte {
	return t._chats
}

func (t *TL_messages_messagesSlice) Set_users(_users []byte) {
	t._users = _users
}

func (t *TL_messages_messagesSlice) Get_users() []byte {
	return t._users
}

func New_TL_messages_messagesSlice() *TL_messages_messagesSlice {
	return &TL_messages_messagesSlice{}
}

func (t *TL_messages_messagesSlice) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_messagesSlice))
	ec.Int(t.Get_count())
	ec.Bytes(t.Get_messages())
	ec.Bytes(t.Get_chats())
	ec.Bytes(t.Get_users())

	return ec.GetBuffer()
}

func (t *TL_messages_messagesSlice) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._count = dc.Int()
	t._messages = dc.Bytes(16)
	t._chats = dc.Bytes(16)
	t._users = dc.Bytes(16)

}

// messages_channelMessages#99262e37
type TL_messages_channelMessages struct {
	_flags    []byte
	_pts      int32
	_count    int32
	_messages []byte
	_chats    []byte
	_users    []byte
}

func (t *TL_messages_channelMessages) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_messages_channelMessages) Get_flags() []byte {
	return t._flags
}

func (t *TL_messages_channelMessages) Set_pts(_pts int32) {
	t._pts = _pts
}

func (t *TL_messages_channelMessages) Get_pts() int32 {
	return t._pts
}

func (t *TL_messages_channelMessages) Set_count(_count int32) {
	t._count = _count
}

func (t *TL_messages_channelMessages) Get_count() int32 {
	return t._count
}

func (t *TL_messages_channelMessages) Set_messages(_messages []byte) {
	t._messages = _messages
}

func (t *TL_messages_channelMessages) Get_messages() []byte {
	return t._messages
}

func (t *TL_messages_channelMessages) Set_chats(_chats []byte) {
	t._chats = _chats
}

func (t *TL_messages_channelMessages) Get_chats() []byte {
	return t._chats
}

func (t *TL_messages_channelMessages) Set_users(_users []byte) {
	t._users = _users
}

func (t *TL_messages_channelMessages) Get_users() []byte {
	return t._users
}

func New_TL_messages_channelMessages() *TL_messages_channelMessages {
	return &TL_messages_channelMessages{}
}

func (t *TL_messages_channelMessages) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_channelMessages))
	ec.Int(t.Get_pts())
	ec.Int(t.Get_count())
	ec.Bytes(t.Get_messages())
	ec.Bytes(t.Get_chats())
	ec.Bytes(t.Get_users())

	return ec.GetBuffer()
}

func (t *TL_messages_channelMessages) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._pts = dc.Int()
	t._count = dc.Int()
	t._messages = dc.Bytes(16)
	t._chats = dc.Bytes(16)
	t._users = dc.Bytes(16)

}

// messages_messagesNotModified#74535f21
type TL_messages_messagesNotModified struct {
	_count int32
}

func (t *TL_messages_messagesNotModified) Set_count(_count int32) {
	t._count = _count
}

func (t *TL_messages_messagesNotModified) Get_count() int32 {
	return t._count
}

func New_TL_messages_messagesNotModified() *TL_messages_messagesNotModified {
	return &TL_messages_messagesNotModified{}
}

func (t *TL_messages_messagesNotModified) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_messagesNotModified))
	ec.Int(t.Get_count())

	return ec.GetBuffer()
}

func (t *TL_messages_messagesNotModified) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._count = dc.Int()

}

// messages_chats#64ff9fd5
type TL_messages_chats struct {
	_chats []byte
}

func (t *TL_messages_chats) Set_chats(_chats []byte) {
	t._chats = _chats
}

func (t *TL_messages_chats) Get_chats() []byte {
	return t._chats
}

func New_TL_messages_chats() *TL_messages_chats {
	return &TL_messages_chats{}
}

func (t *TL_messages_chats) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_chats))
	ec.Bytes(t.Get_chats())

	return ec.GetBuffer()
}

func (t *TL_messages_chats) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._chats = dc.Bytes(16)

}

// messages_chatsSlice#9cd81144
type TL_messages_chatsSlice struct {
	_count int32
	_chats []byte
}

func (t *TL_messages_chatsSlice) Set_count(_count int32) {
	t._count = _count
}

func (t *TL_messages_chatsSlice) Get_count() int32 {
	return t._count
}

func (t *TL_messages_chatsSlice) Set_chats(_chats []byte) {
	t._chats = _chats
}

func (t *TL_messages_chatsSlice) Get_chats() []byte {
	return t._chats
}

func New_TL_messages_chatsSlice() *TL_messages_chatsSlice {
	return &TL_messages_chatsSlice{}
}

func (t *TL_messages_chatsSlice) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_chatsSlice))
	ec.Int(t.Get_count())
	ec.Bytes(t.Get_chats())

	return ec.GetBuffer()
}

func (t *TL_messages_chatsSlice) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._count = dc.Int()
	t._chats = dc.Bytes(16)

}

// messages_chatFull#e5d7d19c
type TL_messages_chatFull struct {
	_full_chat []byte
	_chats     []byte
	_users     []byte
}

func (t *TL_messages_chatFull) Set_full_chat(_full_chat []byte) {
	t._full_chat = _full_chat
}

func (t *TL_messages_chatFull) Get_full_chat() []byte {
	return t._full_chat
}

func (t *TL_messages_chatFull) Set_chats(_chats []byte) {
	t._chats = _chats
}

func (t *TL_messages_chatFull) Get_chats() []byte {
	return t._chats
}

func (t *TL_messages_chatFull) Set_users(_users []byte) {
	t._users = _users
}

func (t *TL_messages_chatFull) Get_users() []byte {
	return t._users
}

func New_TL_messages_chatFull() *TL_messages_chatFull {
	return &TL_messages_chatFull{}
}

func (t *TL_messages_chatFull) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_chatFull))
	ec.Bytes(t.Get_full_chat())
	ec.Bytes(t.Get_chats())
	ec.Bytes(t.Get_users())

	return ec.GetBuffer()
}

func (t *TL_messages_chatFull) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._full_chat = dc.Bytes(16)
	t._chats = dc.Bytes(16)
	t._users = dc.Bytes(16)

}

// messages_affectedHistory#b45c69d1
type TL_messages_affectedHistory struct {
	_pts       int32
	_pts_count int32
	_offset    int32
}

func (t *TL_messages_affectedHistory) Set_pts(_pts int32) {
	t._pts = _pts
}

func (t *TL_messages_affectedHistory) Get_pts() int32 {
	return t._pts
}

func (t *TL_messages_affectedHistory) Set_pts_count(_pts_count int32) {
	t._pts_count = _pts_count
}

func (t *TL_messages_affectedHistory) Get_pts_count() int32 {
	return t._pts_count
}

func (t *TL_messages_affectedHistory) Set_offset(_offset int32) {
	t._offset = _offset
}

func (t *TL_messages_affectedHistory) Get_offset() int32 {
	return t._offset
}

func New_TL_messages_affectedHistory() *TL_messages_affectedHistory {
	return &TL_messages_affectedHistory{}
}

func (t *TL_messages_affectedHistory) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_affectedHistory))
	ec.Int(t.Get_pts())
	ec.Int(t.Get_pts_count())
	ec.Int(t.Get_offset())

	return ec.GetBuffer()
}

func (t *TL_messages_affectedHistory) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._pts = dc.Int()
	t._pts_count = dc.Int()
	t._offset = dc.Int()

}

// inputMessagesFilterEmpty#57e2f66c
type TL_inputMessagesFilterEmpty struct {
}

func New_TL_inputMessagesFilterEmpty() *TL_inputMessagesFilterEmpty {
	return &TL_inputMessagesFilterEmpty{}
}

func (t *TL_inputMessagesFilterEmpty) Encode() []byte {
	return nil
}

func (t *TL_inputMessagesFilterEmpty) Decode(b []byte) {

}

// inputMessagesFilterPhotos#9609a51c
type TL_inputMessagesFilterPhotos struct {
}

func New_TL_inputMessagesFilterPhotos() *TL_inputMessagesFilterPhotos {
	return &TL_inputMessagesFilterPhotos{}
}

func (t *TL_inputMessagesFilterPhotos) Encode() []byte {
	return nil
}

func (t *TL_inputMessagesFilterPhotos) Decode(b []byte) {

}

// inputMessagesFilterVideo#9fc00e65
type TL_inputMessagesFilterVideo struct {
}

func New_TL_inputMessagesFilterVideo() *TL_inputMessagesFilterVideo {
	return &TL_inputMessagesFilterVideo{}
}

func (t *TL_inputMessagesFilterVideo) Encode() []byte {
	return nil
}

func (t *TL_inputMessagesFilterVideo) Decode(b []byte) {

}

// inputMessagesFilterPhotoVideo#56e9f0e4
type TL_inputMessagesFilterPhotoVideo struct {
}

func New_TL_inputMessagesFilterPhotoVideo() *TL_inputMessagesFilterPhotoVideo {
	return &TL_inputMessagesFilterPhotoVideo{}
}

func (t *TL_inputMessagesFilterPhotoVideo) Encode() []byte {
	return nil
}

func (t *TL_inputMessagesFilterPhotoVideo) Decode(b []byte) {

}

// inputMessagesFilterDocument#9eddf188
type TL_inputMessagesFilterDocument struct {
}

func New_TL_inputMessagesFilterDocument() *TL_inputMessagesFilterDocument {
	return &TL_inputMessagesFilterDocument{}
}

func (t *TL_inputMessagesFilterDocument) Encode() []byte {
	return nil
}

func (t *TL_inputMessagesFilterDocument) Decode(b []byte) {

}

// inputMessagesFilterUrl#7ef0dd87
type TL_inputMessagesFilterUrl struct {
}

func New_TL_inputMessagesFilterUrl() *TL_inputMessagesFilterUrl {
	return &TL_inputMessagesFilterUrl{}
}

func (t *TL_inputMessagesFilterUrl) Encode() []byte {
	return nil
}

func (t *TL_inputMessagesFilterUrl) Decode(b []byte) {

}

// inputMessagesFilterGif#ffc86587
type TL_inputMessagesFilterGif struct {
}

func New_TL_inputMessagesFilterGif() *TL_inputMessagesFilterGif {
	return &TL_inputMessagesFilterGif{}
}

func (t *TL_inputMessagesFilterGif) Encode() []byte {
	return nil
}

func (t *TL_inputMessagesFilterGif) Decode(b []byte) {

}

// inputMessagesFilterVoice#50f5c392
type TL_inputMessagesFilterVoice struct {
}

func New_TL_inputMessagesFilterVoice() *TL_inputMessagesFilterVoice {
	return &TL_inputMessagesFilterVoice{}
}

func (t *TL_inputMessagesFilterVoice) Encode() []byte {
	return nil
}

func (t *TL_inputMessagesFilterVoice) Decode(b []byte) {

}

// inputMessagesFilterMusic#3751b49e
type TL_inputMessagesFilterMusic struct {
}

func New_TL_inputMessagesFilterMusic() *TL_inputMessagesFilterMusic {
	return &TL_inputMessagesFilterMusic{}
}

func (t *TL_inputMessagesFilterMusic) Encode() []byte {
	return nil
}

func (t *TL_inputMessagesFilterMusic) Decode(b []byte) {

}

// inputMessagesFilterChatPhotos#3a20ecb8
type TL_inputMessagesFilterChatPhotos struct {
}

func New_TL_inputMessagesFilterChatPhotos() *TL_inputMessagesFilterChatPhotos {
	return &TL_inputMessagesFilterChatPhotos{}
}

func (t *TL_inputMessagesFilterChatPhotos) Encode() []byte {
	return nil
}

func (t *TL_inputMessagesFilterChatPhotos) Decode(b []byte) {

}

// inputMessagesFilterPhoneCalls#80c99768
type TL_inputMessagesFilterPhoneCalls struct {
	_flags  []byte
	_missed []byte
}

func (t *TL_inputMessagesFilterPhoneCalls) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_inputMessagesFilterPhoneCalls) Get_flags() []byte {
	return t._flags
}

func (t *TL_inputMessagesFilterPhoneCalls) Set_missed(_missed []byte) {
	t._missed = _missed
}

func (t *TL_inputMessagesFilterPhoneCalls) Get_missed() []byte {
	return t._missed
}

func New_TL_inputMessagesFilterPhoneCalls() *TL_inputMessagesFilterPhoneCalls {
	return &TL_inputMessagesFilterPhoneCalls{}
}

func (t *TL_inputMessagesFilterPhoneCalls) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputMessagesFilterPhoneCalls))
	ec.Bytes(t.Get_missed())

	return ec.GetBuffer()
}

func (t *TL_inputMessagesFilterPhoneCalls) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._missed = dc.Bytes(16)

}

// inputMessagesFilterRoundVoice#7a7c17a4
type TL_inputMessagesFilterRoundVoice struct {
}

func New_TL_inputMessagesFilterRoundVoice() *TL_inputMessagesFilterRoundVoice {
	return &TL_inputMessagesFilterRoundVoice{}
}

func (t *TL_inputMessagesFilterRoundVoice) Encode() []byte {
	return nil
}

func (t *TL_inputMessagesFilterRoundVoice) Decode(b []byte) {

}

// inputMessagesFilterRoundVideo#b549da53
type TL_inputMessagesFilterRoundVideo struct {
}

func New_TL_inputMessagesFilterRoundVideo() *TL_inputMessagesFilterRoundVideo {
	return &TL_inputMessagesFilterRoundVideo{}
}

func (t *TL_inputMessagesFilterRoundVideo) Encode() []byte {
	return nil
}

func (t *TL_inputMessagesFilterRoundVideo) Decode(b []byte) {

}

// inputMessagesFilterMyMentions#c1f8e69a
type TL_inputMessagesFilterMyMentions struct {
}

func New_TL_inputMessagesFilterMyMentions() *TL_inputMessagesFilterMyMentions {
	return &TL_inputMessagesFilterMyMentions{}
}

func (t *TL_inputMessagesFilterMyMentions) Encode() []byte {
	return nil
}

func (t *TL_inputMessagesFilterMyMentions) Decode(b []byte) {

}

// inputMessagesFilterGeo#e7026d0d
type TL_inputMessagesFilterGeo struct {
}

func New_TL_inputMessagesFilterGeo() *TL_inputMessagesFilterGeo {
	return &TL_inputMessagesFilterGeo{}
}

func (t *TL_inputMessagesFilterGeo) Encode() []byte {
	return nil
}

func (t *TL_inputMessagesFilterGeo) Decode(b []byte) {

}

// inputMessagesFilterContacts#e062db83
type TL_inputMessagesFilterContacts struct {
}

func New_TL_inputMessagesFilterContacts() *TL_inputMessagesFilterContacts {
	return &TL_inputMessagesFilterContacts{}
}

func (t *TL_inputMessagesFilterContacts) Encode() []byte {
	return nil
}

func (t *TL_inputMessagesFilterContacts) Decode(b []byte) {

}

// updateNewMessage#1f2b0afd
type TL_updateNewMessage struct {
	_message   []byte
	_pts       int32
	_pts_count int32
}

func (t *TL_updateNewMessage) Set_message(_message []byte) {
	t._message = _message
}

func (t *TL_updateNewMessage) Get_message() []byte {
	return t._message
}

func (t *TL_updateNewMessage) Set_pts(_pts int32) {
	t._pts = _pts
}

func (t *TL_updateNewMessage) Get_pts() int32 {
	return t._pts
}

func (t *TL_updateNewMessage) Set_pts_count(_pts_count int32) {
	t._pts_count = _pts_count
}

func (t *TL_updateNewMessage) Get_pts_count() int32 {
	return t._pts_count
}

func New_TL_updateNewMessage() *TL_updateNewMessage {
	return &TL_updateNewMessage{}
}

func (t *TL_updateNewMessage) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updateNewMessage))
	ec.Bytes(t.Get_message())
	ec.Int(t.Get_pts())
	ec.Int(t.Get_pts_count())

	return ec.GetBuffer()
}

func (t *TL_updateNewMessage) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._message = dc.Bytes(16)
	t._pts = dc.Int()
	t._pts_count = dc.Int()

}

// updateMessageID#4e90bfd6
type TL_updateMessageID struct {
	_id        int32
	_random_id int64
}

func (t *TL_updateMessageID) Set_id(_id int32) {
	t._id = _id
}

func (t *TL_updateMessageID) Get_id() int32 {
	return t._id
}

func (t *TL_updateMessageID) Set_random_id(_random_id int64) {
	t._random_id = _random_id
}

func (t *TL_updateMessageID) Get_random_id() int64 {
	return t._random_id
}

func New_TL_updateMessageID() *TL_updateMessageID {
	return &TL_updateMessageID{}
}

func (t *TL_updateMessageID) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updateMessageID))
	ec.Int(t.Get_id())
	ec.Long(t.Get_random_id())

	return ec.GetBuffer()
}

func (t *TL_updateMessageID) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Int()
	t._random_id = dc.Long()

}

// updateDeleteMessages#a20db0e5
type TL_updateDeleteMessages struct {
	_messages  []byte
	_pts       int32
	_pts_count int32
}

func (t *TL_updateDeleteMessages) Set_messages(_messages []byte) {
	t._messages = _messages
}

func (t *TL_updateDeleteMessages) Get_messages() []byte {
	return t._messages
}

func (t *TL_updateDeleteMessages) Set_pts(_pts int32) {
	t._pts = _pts
}

func (t *TL_updateDeleteMessages) Get_pts() int32 {
	return t._pts
}

func (t *TL_updateDeleteMessages) Set_pts_count(_pts_count int32) {
	t._pts_count = _pts_count
}

func (t *TL_updateDeleteMessages) Get_pts_count() int32 {
	return t._pts_count
}

func New_TL_updateDeleteMessages() *TL_updateDeleteMessages {
	return &TL_updateDeleteMessages{}
}

func (t *TL_updateDeleteMessages) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updateDeleteMessages))
	ec.Bytes(t.Get_messages())
	ec.Int(t.Get_pts())
	ec.Int(t.Get_pts_count())

	return ec.GetBuffer()
}

func (t *TL_updateDeleteMessages) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._messages = dc.Bytes(16)
	t._pts = dc.Int()
	t._pts_count = dc.Int()

}

// updateUserTyping#5c486927
type TL_updateUserTyping struct {
	_user_id int32
	_action  []byte
}

func (t *TL_updateUserTyping) Set_user_id(_user_id int32) {
	t._user_id = _user_id
}

func (t *TL_updateUserTyping) Get_user_id() int32 {
	return t._user_id
}

func (t *TL_updateUserTyping) Set_action(_action []byte) {
	t._action = _action
}

func (t *TL_updateUserTyping) Get_action() []byte {
	return t._action
}

func New_TL_updateUserTyping() *TL_updateUserTyping {
	return &TL_updateUserTyping{}
}

func (t *TL_updateUserTyping) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updateUserTyping))
	ec.Int(t.Get_user_id())
	ec.Bytes(t.Get_action())

	return ec.GetBuffer()
}

func (t *TL_updateUserTyping) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._user_id = dc.Int()
	t._action = dc.Bytes(16)

}

// updateChatUserTyping#9a65ea1f
type TL_updateChatUserTyping struct {
	_chat_id int32
	_user_id int32
	_action  []byte
}

func (t *TL_updateChatUserTyping) Set_chat_id(_chat_id int32) {
	t._chat_id = _chat_id
}

func (t *TL_updateChatUserTyping) Get_chat_id() int32 {
	return t._chat_id
}

func (t *TL_updateChatUserTyping) Set_user_id(_user_id int32) {
	t._user_id = _user_id
}

func (t *TL_updateChatUserTyping) Get_user_id() int32 {
	return t._user_id
}

func (t *TL_updateChatUserTyping) Set_action(_action []byte) {
	t._action = _action
}

func (t *TL_updateChatUserTyping) Get_action() []byte {
	return t._action
}

func New_TL_updateChatUserTyping() *TL_updateChatUserTyping {
	return &TL_updateChatUserTyping{}
}

func (t *TL_updateChatUserTyping) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updateChatUserTyping))
	ec.Int(t.Get_chat_id())
	ec.Int(t.Get_user_id())
	ec.Bytes(t.Get_action())

	return ec.GetBuffer()
}

func (t *TL_updateChatUserTyping) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._chat_id = dc.Int()
	t._user_id = dc.Int()
	t._action = dc.Bytes(16)

}

// updateChatParticipants#07761198
type TL_updateChatParticipants struct {
	_participants []byte
}

func (t *TL_updateChatParticipants) Set_participants(_participants []byte) {
	t._participants = _participants
}

func (t *TL_updateChatParticipants) Get_participants() []byte {
	return t._participants
}

func New_TL_updateChatParticipants() *TL_updateChatParticipants {
	return &TL_updateChatParticipants{}
}

func (t *TL_updateChatParticipants) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updateChatParticipants))
	ec.Bytes(t.Get_participants())

	return ec.GetBuffer()
}

func (t *TL_updateChatParticipants) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._participants = dc.Bytes(16)

}

// updateUserStatus#1bfbd823
type TL_updateUserStatus struct {
	_user_id int32
	_status  []byte
}

func (t *TL_updateUserStatus) Set_user_id(_user_id int32) {
	t._user_id = _user_id
}

func (t *TL_updateUserStatus) Get_user_id() int32 {
	return t._user_id
}

func (t *TL_updateUserStatus) Set_status(_status []byte) {
	t._status = _status
}

func (t *TL_updateUserStatus) Get_status() []byte {
	return t._status
}

func New_TL_updateUserStatus() *TL_updateUserStatus {
	return &TL_updateUserStatus{}
}

func (t *TL_updateUserStatus) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updateUserStatus))
	ec.Int(t.Get_user_id())
	ec.Bytes(t.Get_status())

	return ec.GetBuffer()
}

func (t *TL_updateUserStatus) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._user_id = dc.Int()
	t._status = dc.Bytes(16)

}

// updateUserName#a7332b73
type TL_updateUserName struct {
	_user_id    int32
	_first_name string
	_last_name  string
	_username   string
}

func (t *TL_updateUserName) Set_user_id(_user_id int32) {
	t._user_id = _user_id
}

func (t *TL_updateUserName) Get_user_id() int32 {
	return t._user_id
}

func (t *TL_updateUserName) Set_first_name(_first_name string) {
	t._first_name = _first_name
}

func (t *TL_updateUserName) Get_first_name() string {
	return t._first_name
}

func (t *TL_updateUserName) Set_last_name(_last_name string) {
	t._last_name = _last_name
}

func (t *TL_updateUserName) Get_last_name() string {
	return t._last_name
}

func (t *TL_updateUserName) Set_username(_username string) {
	t._username = _username
}

func (t *TL_updateUserName) Get_username() string {
	return t._username
}

func New_TL_updateUserName() *TL_updateUserName {
	return &TL_updateUserName{}
}

func (t *TL_updateUserName) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updateUserName))
	ec.Int(t.Get_user_id())
	ec.String(t.Get_first_name())
	ec.String(t.Get_last_name())
	ec.String(t.Get_username())

	return ec.GetBuffer()
}

func (t *TL_updateUserName) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._user_id = dc.Int()
	t._first_name = dc.String()
	t._last_name = dc.String()
	t._username = dc.String()

}

// updateUserPhoto#95313b0c
type TL_updateUserPhoto struct {
	_user_id  int32
	_date     int32
	_photo    []byte
	_previous bool
}

func (t *TL_updateUserPhoto) Set_user_id(_user_id int32) {
	t._user_id = _user_id
}

func (t *TL_updateUserPhoto) Get_user_id() int32 {
	return t._user_id
}

func (t *TL_updateUserPhoto) Set_date(_date int32) {
	t._date = _date
}

func (t *TL_updateUserPhoto) Get_date() int32 {
	return t._date
}

func (t *TL_updateUserPhoto) Set_photo(_photo []byte) {
	t._photo = _photo
}

func (t *TL_updateUserPhoto) Get_photo() []byte {
	return t._photo
}

func (t *TL_updateUserPhoto) Set_previous(_previous bool) {
	t._previous = _previous
}

func (t *TL_updateUserPhoto) Get_previous() bool {
	return t._previous
}

func New_TL_updateUserPhoto() *TL_updateUserPhoto {
	return &TL_updateUserPhoto{}
}

func (t *TL_updateUserPhoto) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updateUserPhoto))
	ec.Int(t.Get_user_id())
	ec.Int(t.Get_date())
	ec.Bytes(t.Get_photo())
	ec.Bool(t.Get_previous())

	return ec.GetBuffer()
}

func (t *TL_updateUserPhoto) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._user_id = dc.Int()
	t._date = dc.Int()
	t._photo = dc.Bytes(16)
	t._previous = dc.Bool()

}

// updateContactRegistered#2575bbb9
type TL_updateContactRegistered struct {
	_user_id int32
	_date    int32
}

func (t *TL_updateContactRegistered) Set_user_id(_user_id int32) {
	t._user_id = _user_id
}

func (t *TL_updateContactRegistered) Get_user_id() int32 {
	return t._user_id
}

func (t *TL_updateContactRegistered) Set_date(_date int32) {
	t._date = _date
}

func (t *TL_updateContactRegistered) Get_date() int32 {
	return t._date
}

func New_TL_updateContactRegistered() *TL_updateContactRegistered {
	return &TL_updateContactRegistered{}
}

func (t *TL_updateContactRegistered) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updateContactRegistered))
	ec.Int(t.Get_user_id())
	ec.Int(t.Get_date())

	return ec.GetBuffer()
}

func (t *TL_updateContactRegistered) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._user_id = dc.Int()
	t._date = dc.Int()

}

// updateContactLink#9d2e67c5
type TL_updateContactLink struct {
	_user_id      int32
	_my_link      []byte
	_foreign_link []byte
}

func (t *TL_updateContactLink) Set_user_id(_user_id int32) {
	t._user_id = _user_id
}

func (t *TL_updateContactLink) Get_user_id() int32 {
	return t._user_id
}

func (t *TL_updateContactLink) Set_my_link(_my_link []byte) {
	t._my_link = _my_link
}

func (t *TL_updateContactLink) Get_my_link() []byte {
	return t._my_link
}

func (t *TL_updateContactLink) Set_foreign_link(_foreign_link []byte) {
	t._foreign_link = _foreign_link
}

func (t *TL_updateContactLink) Get_foreign_link() []byte {
	return t._foreign_link
}

func New_TL_updateContactLink() *TL_updateContactLink {
	return &TL_updateContactLink{}
}

func (t *TL_updateContactLink) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updateContactLink))
	ec.Int(t.Get_user_id())
	ec.Bytes(t.Get_my_link())
	ec.Bytes(t.Get_foreign_link())

	return ec.GetBuffer()
}

func (t *TL_updateContactLink) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._user_id = dc.Int()
	t._my_link = dc.Bytes(16)
	t._foreign_link = dc.Bytes(16)

}

// updateNewEncryptedMessage#12bcbd9a
type TL_updateNewEncryptedMessage struct {
	_message []byte
	_qts     int32
}

func (t *TL_updateNewEncryptedMessage) Set_message(_message []byte) {
	t._message = _message
}

func (t *TL_updateNewEncryptedMessage) Get_message() []byte {
	return t._message
}

func (t *TL_updateNewEncryptedMessage) Set_qts(_qts int32) {
	t._qts = _qts
}

func (t *TL_updateNewEncryptedMessage) Get_qts() int32 {
	return t._qts
}

func New_TL_updateNewEncryptedMessage() *TL_updateNewEncryptedMessage {
	return &TL_updateNewEncryptedMessage{}
}

func (t *TL_updateNewEncryptedMessage) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updateNewEncryptedMessage))
	ec.Bytes(t.Get_message())
	ec.Int(t.Get_qts())

	return ec.GetBuffer()
}

func (t *TL_updateNewEncryptedMessage) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._message = dc.Bytes(16)
	t._qts = dc.Int()

}

// updateEncryptedChatTyping#1710f156
type TL_updateEncryptedChatTyping struct {
	_chat_id int32
}

func (t *TL_updateEncryptedChatTyping) Set_chat_id(_chat_id int32) {
	t._chat_id = _chat_id
}

func (t *TL_updateEncryptedChatTyping) Get_chat_id() int32 {
	return t._chat_id
}

func New_TL_updateEncryptedChatTyping() *TL_updateEncryptedChatTyping {
	return &TL_updateEncryptedChatTyping{}
}

func (t *TL_updateEncryptedChatTyping) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updateEncryptedChatTyping))
	ec.Int(t.Get_chat_id())

	return ec.GetBuffer()
}

func (t *TL_updateEncryptedChatTyping) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._chat_id = dc.Int()

}

// updateEncryption#b4a2e88d
type TL_updateEncryption struct {
	_chat []byte
	_date int32
}

func (t *TL_updateEncryption) Set_chat(_chat []byte) {
	t._chat = _chat
}

func (t *TL_updateEncryption) Get_chat() []byte {
	return t._chat
}

func (t *TL_updateEncryption) Set_date(_date int32) {
	t._date = _date
}

func (t *TL_updateEncryption) Get_date() int32 {
	return t._date
}

func New_TL_updateEncryption() *TL_updateEncryption {
	return &TL_updateEncryption{}
}

func (t *TL_updateEncryption) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updateEncryption))
	ec.Bytes(t.Get_chat())
	ec.Int(t.Get_date())

	return ec.GetBuffer()
}

func (t *TL_updateEncryption) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._chat = dc.Bytes(16)
	t._date = dc.Int()

}

// updateEncryptedMessagesRead#38fe25b7
type TL_updateEncryptedMessagesRead struct {
	_chat_id  int32
	_max_date int32
	_date     int32
}

func (t *TL_updateEncryptedMessagesRead) Set_chat_id(_chat_id int32) {
	t._chat_id = _chat_id
}

func (t *TL_updateEncryptedMessagesRead) Get_chat_id() int32 {
	return t._chat_id
}

func (t *TL_updateEncryptedMessagesRead) Set_max_date(_max_date int32) {
	t._max_date = _max_date
}

func (t *TL_updateEncryptedMessagesRead) Get_max_date() int32 {
	return t._max_date
}

func (t *TL_updateEncryptedMessagesRead) Set_date(_date int32) {
	t._date = _date
}

func (t *TL_updateEncryptedMessagesRead) Get_date() int32 {
	return t._date
}

func New_TL_updateEncryptedMessagesRead() *TL_updateEncryptedMessagesRead {
	return &TL_updateEncryptedMessagesRead{}
}

func (t *TL_updateEncryptedMessagesRead) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updateEncryptedMessagesRead))
	ec.Int(t.Get_chat_id())
	ec.Int(t.Get_max_date())
	ec.Int(t.Get_date())

	return ec.GetBuffer()
}

func (t *TL_updateEncryptedMessagesRead) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._chat_id = dc.Int()
	t._max_date = dc.Int()
	t._date = dc.Int()

}

// updateChatParticipantAdd#ea4b0e5c
type TL_updateChatParticipantAdd struct {
	_chat_id    int32
	_user_id    int32
	_inviter_id int32
	_date       int32
	_version    int32
}

func (t *TL_updateChatParticipantAdd) Set_chat_id(_chat_id int32) {
	t._chat_id = _chat_id
}

func (t *TL_updateChatParticipantAdd) Get_chat_id() int32 {
	return t._chat_id
}

func (t *TL_updateChatParticipantAdd) Set_user_id(_user_id int32) {
	t._user_id = _user_id
}

func (t *TL_updateChatParticipantAdd) Get_user_id() int32 {
	return t._user_id
}

func (t *TL_updateChatParticipantAdd) Set_inviter_id(_inviter_id int32) {
	t._inviter_id = _inviter_id
}

func (t *TL_updateChatParticipantAdd) Get_inviter_id() int32 {
	return t._inviter_id
}

func (t *TL_updateChatParticipantAdd) Set_date(_date int32) {
	t._date = _date
}

func (t *TL_updateChatParticipantAdd) Get_date() int32 {
	return t._date
}

func (t *TL_updateChatParticipantAdd) Set_version(_version int32) {
	t._version = _version
}

func (t *TL_updateChatParticipantAdd) Get_version() int32 {
	return t._version
}

func New_TL_updateChatParticipantAdd() *TL_updateChatParticipantAdd {
	return &TL_updateChatParticipantAdd{}
}

func (t *TL_updateChatParticipantAdd) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updateChatParticipantAdd))
	ec.Int(t.Get_chat_id())
	ec.Int(t.Get_user_id())
	ec.Int(t.Get_inviter_id())
	ec.Int(t.Get_date())
	ec.Int(t.Get_version())

	return ec.GetBuffer()
}

func (t *TL_updateChatParticipantAdd) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._chat_id = dc.Int()
	t._user_id = dc.Int()
	t._inviter_id = dc.Int()
	t._date = dc.Int()
	t._version = dc.Int()

}

// updateChatParticipantDelete#6e5f8c22
type TL_updateChatParticipantDelete struct {
	_chat_id int32
	_user_id int32
	_version int32
}

func (t *TL_updateChatParticipantDelete) Set_chat_id(_chat_id int32) {
	t._chat_id = _chat_id
}

func (t *TL_updateChatParticipantDelete) Get_chat_id() int32 {
	return t._chat_id
}

func (t *TL_updateChatParticipantDelete) Set_user_id(_user_id int32) {
	t._user_id = _user_id
}

func (t *TL_updateChatParticipantDelete) Get_user_id() int32 {
	return t._user_id
}

func (t *TL_updateChatParticipantDelete) Set_version(_version int32) {
	t._version = _version
}

func (t *TL_updateChatParticipantDelete) Get_version() int32 {
	return t._version
}

func New_TL_updateChatParticipantDelete() *TL_updateChatParticipantDelete {
	return &TL_updateChatParticipantDelete{}
}

func (t *TL_updateChatParticipantDelete) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updateChatParticipantDelete))
	ec.Int(t.Get_chat_id())
	ec.Int(t.Get_user_id())
	ec.Int(t.Get_version())

	return ec.GetBuffer()
}

func (t *TL_updateChatParticipantDelete) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._chat_id = dc.Int()
	t._user_id = dc.Int()
	t._version = dc.Int()

}

// updateDcOptions#8e5e9873
type TL_updateDcOptions struct {
	_dc_options []byte
}

func (t *TL_updateDcOptions) Set_dc_options(_dc_options []byte) {
	t._dc_options = _dc_options
}

func (t *TL_updateDcOptions) Get_dc_options() []byte {
	return t._dc_options
}

func New_TL_updateDcOptions() *TL_updateDcOptions {
	return &TL_updateDcOptions{}
}

func (t *TL_updateDcOptions) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updateDcOptions))
	ec.Bytes(t.Get_dc_options())

	return ec.GetBuffer()
}

func (t *TL_updateDcOptions) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._dc_options = dc.Bytes(16)

}

// updateUserBlocked#80ece81a
type TL_updateUserBlocked struct {
	_user_id int32
	_blocked bool
}

func (t *TL_updateUserBlocked) Set_user_id(_user_id int32) {
	t._user_id = _user_id
}

func (t *TL_updateUserBlocked) Get_user_id() int32 {
	return t._user_id
}

func (t *TL_updateUserBlocked) Set_blocked(_blocked bool) {
	t._blocked = _blocked
}

func (t *TL_updateUserBlocked) Get_blocked() bool {
	return t._blocked
}

func New_TL_updateUserBlocked() *TL_updateUserBlocked {
	return &TL_updateUserBlocked{}
}

func (t *TL_updateUserBlocked) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updateUserBlocked))
	ec.Int(t.Get_user_id())
	ec.Bool(t.Get_blocked())

	return ec.GetBuffer()
}

func (t *TL_updateUserBlocked) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._user_id = dc.Int()
	t._blocked = dc.Bool()

}

// updateNotifySettings#bec268ef
type TL_updateNotifySettings struct {
	_peer            []byte
	_notify_settings []byte
}

func (t *TL_updateNotifySettings) Set_peer(_peer []byte) {
	t._peer = _peer
}

func (t *TL_updateNotifySettings) Get_peer() []byte {
	return t._peer
}

func (t *TL_updateNotifySettings) Set_notify_settings(_notify_settings []byte) {
	t._notify_settings = _notify_settings
}

func (t *TL_updateNotifySettings) Get_notify_settings() []byte {
	return t._notify_settings
}

func New_TL_updateNotifySettings() *TL_updateNotifySettings {
	return &TL_updateNotifySettings{}
}

func (t *TL_updateNotifySettings) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updateNotifySettings))
	ec.Bytes(t.Get_peer())
	ec.Bytes(t.Get_notify_settings())

	return ec.GetBuffer()
}

func (t *TL_updateNotifySettings) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._peer = dc.Bytes(16)
	t._notify_settings = dc.Bytes(16)

}

// updateServiceNotification#ebe46819
type TL_updateServiceNotification struct {
	_flags      []byte
	_popup      []byte
	_inbox_date []byte
	_type       string
	_message    string
	_media      []byte
	_entities   []byte
}

func (t *TL_updateServiceNotification) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_updateServiceNotification) Get_flags() []byte {
	return t._flags
}

func (t *TL_updateServiceNotification) Set_popup(_popup []byte) {
	t._popup = _popup
}

func (t *TL_updateServiceNotification) Get_popup() []byte {
	return t._popup
}

func (t *TL_updateServiceNotification) Set_inbox_date(_inbox_date []byte) {
	t._inbox_date = _inbox_date
}

func (t *TL_updateServiceNotification) Get_inbox_date() []byte {
	return t._inbox_date
}

func (t *TL_updateServiceNotification) Set_type(_type string) {
	t._type = _type
}

func (t *TL_updateServiceNotification) Get_type() string {
	return t._type
}

func (t *TL_updateServiceNotification) Set_message(_message string) {
	t._message = _message
}

func (t *TL_updateServiceNotification) Get_message() string {
	return t._message
}

func (t *TL_updateServiceNotification) Set_media(_media []byte) {
	t._media = _media
}

func (t *TL_updateServiceNotification) Get_media() []byte {
	return t._media
}

func (t *TL_updateServiceNotification) Set_entities(_entities []byte) {
	t._entities = _entities
}

func (t *TL_updateServiceNotification) Get_entities() []byte {
	return t._entities
}

func New_TL_updateServiceNotification() *TL_updateServiceNotification {
	return &TL_updateServiceNotification{}
}

func (t *TL_updateServiceNotification) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updateServiceNotification))
	ec.Bytes(t.Get_popup())
	ec.Bytes(t.Get_inbox_date())
	ec.String(t.Get_type())
	ec.String(t.Get_message())
	ec.Bytes(t.Get_media())
	ec.Bytes(t.Get_entities())

	return ec.GetBuffer()
}

func (t *TL_updateServiceNotification) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._popup = dc.Bytes(16)
	t._inbox_date = dc.Bytes(16)
	t._type = dc.String()
	t._message = dc.String()
	t._media = dc.Bytes(16)
	t._entities = dc.Bytes(16)

}

// updatePrivacy#ee3b272a
type TL_updatePrivacy struct {
	_key   []byte
	_rules []byte
}

func (t *TL_updatePrivacy) Set_key(_key []byte) {
	t._key = _key
}

func (t *TL_updatePrivacy) Get_key() []byte {
	return t._key
}

func (t *TL_updatePrivacy) Set_rules(_rules []byte) {
	t._rules = _rules
}

func (t *TL_updatePrivacy) Get_rules() []byte {
	return t._rules
}

func New_TL_updatePrivacy() *TL_updatePrivacy {
	return &TL_updatePrivacy{}
}

func (t *TL_updatePrivacy) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updatePrivacy))
	ec.Bytes(t.Get_key())
	ec.Bytes(t.Get_rules())

	return ec.GetBuffer()
}

func (t *TL_updatePrivacy) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._key = dc.Bytes(16)
	t._rules = dc.Bytes(16)

}

// updateUserPhone#12b9417b
type TL_updateUserPhone struct {
	_user_id int32
	_phone   string
}

func (t *TL_updateUserPhone) Set_user_id(_user_id int32) {
	t._user_id = _user_id
}

func (t *TL_updateUserPhone) Get_user_id() int32 {
	return t._user_id
}

func (t *TL_updateUserPhone) Set_phone(_phone string) {
	t._phone = _phone
}

func (t *TL_updateUserPhone) Get_phone() string {
	return t._phone
}

func New_TL_updateUserPhone() *TL_updateUserPhone {
	return &TL_updateUserPhone{}
}

func (t *TL_updateUserPhone) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updateUserPhone))
	ec.Int(t.Get_user_id())
	ec.String(t.Get_phone())

	return ec.GetBuffer()
}

func (t *TL_updateUserPhone) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._user_id = dc.Int()
	t._phone = dc.String()

}

// updateReadHistoryInbox#9961fd5c
type TL_updateReadHistoryInbox struct {
	_peer      []byte
	_max_id    int32
	_pts       int32
	_pts_count int32
}

func (t *TL_updateReadHistoryInbox) Set_peer(_peer []byte) {
	t._peer = _peer
}

func (t *TL_updateReadHistoryInbox) Get_peer() []byte {
	return t._peer
}

func (t *TL_updateReadHistoryInbox) Set_max_id(_max_id int32) {
	t._max_id = _max_id
}

func (t *TL_updateReadHistoryInbox) Get_max_id() int32 {
	return t._max_id
}

func (t *TL_updateReadHistoryInbox) Set_pts(_pts int32) {
	t._pts = _pts
}

func (t *TL_updateReadHistoryInbox) Get_pts() int32 {
	return t._pts
}

func (t *TL_updateReadHistoryInbox) Set_pts_count(_pts_count int32) {
	t._pts_count = _pts_count
}

func (t *TL_updateReadHistoryInbox) Get_pts_count() int32 {
	return t._pts_count
}

func New_TL_updateReadHistoryInbox() *TL_updateReadHistoryInbox {
	return &TL_updateReadHistoryInbox{}
}

func (t *TL_updateReadHistoryInbox) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updateReadHistoryInbox))
	ec.Bytes(t.Get_peer())
	ec.Int(t.Get_max_id())
	ec.Int(t.Get_pts())
	ec.Int(t.Get_pts_count())

	return ec.GetBuffer()
}

func (t *TL_updateReadHistoryInbox) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._peer = dc.Bytes(16)
	t._max_id = dc.Int()
	t._pts = dc.Int()
	t._pts_count = dc.Int()

}

// updateReadHistoryOutbox#2f2f21bf
type TL_updateReadHistoryOutbox struct {
	_peer      []byte
	_max_id    int32
	_pts       int32
	_pts_count int32
}

func (t *TL_updateReadHistoryOutbox) Set_peer(_peer []byte) {
	t._peer = _peer
}

func (t *TL_updateReadHistoryOutbox) Get_peer() []byte {
	return t._peer
}

func (t *TL_updateReadHistoryOutbox) Set_max_id(_max_id int32) {
	t._max_id = _max_id
}

func (t *TL_updateReadHistoryOutbox) Get_max_id() int32 {
	return t._max_id
}

func (t *TL_updateReadHistoryOutbox) Set_pts(_pts int32) {
	t._pts = _pts
}

func (t *TL_updateReadHistoryOutbox) Get_pts() int32 {
	return t._pts
}

func (t *TL_updateReadHistoryOutbox) Set_pts_count(_pts_count int32) {
	t._pts_count = _pts_count
}

func (t *TL_updateReadHistoryOutbox) Get_pts_count() int32 {
	return t._pts_count
}

func New_TL_updateReadHistoryOutbox() *TL_updateReadHistoryOutbox {
	return &TL_updateReadHistoryOutbox{}
}

func (t *TL_updateReadHistoryOutbox) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updateReadHistoryOutbox))
	ec.Bytes(t.Get_peer())
	ec.Int(t.Get_max_id())
	ec.Int(t.Get_pts())
	ec.Int(t.Get_pts_count())

	return ec.GetBuffer()
}

func (t *TL_updateReadHistoryOutbox) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._peer = dc.Bytes(16)
	t._max_id = dc.Int()
	t._pts = dc.Int()
	t._pts_count = dc.Int()

}

// updateWebPage#7f891213
type TL_updateWebPage struct {
	_webpage   []byte
	_pts       int32
	_pts_count int32
}

func (t *TL_updateWebPage) Set_webpage(_webpage []byte) {
	t._webpage = _webpage
}

func (t *TL_updateWebPage) Get_webpage() []byte {
	return t._webpage
}

func (t *TL_updateWebPage) Set_pts(_pts int32) {
	t._pts = _pts
}

func (t *TL_updateWebPage) Get_pts() int32 {
	return t._pts
}

func (t *TL_updateWebPage) Set_pts_count(_pts_count int32) {
	t._pts_count = _pts_count
}

func (t *TL_updateWebPage) Get_pts_count() int32 {
	return t._pts_count
}

func New_TL_updateWebPage() *TL_updateWebPage {
	return &TL_updateWebPage{}
}

func (t *TL_updateWebPage) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updateWebPage))
	ec.Bytes(t.Get_webpage())
	ec.Int(t.Get_pts())
	ec.Int(t.Get_pts_count())

	return ec.GetBuffer()
}

func (t *TL_updateWebPage) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._webpage = dc.Bytes(16)
	t._pts = dc.Int()
	t._pts_count = dc.Int()

}

// updateReadMessagesContents#68c13933
type TL_updateReadMessagesContents struct {
	_messages  []byte
	_pts       int32
	_pts_count int32
}

func (t *TL_updateReadMessagesContents) Set_messages(_messages []byte) {
	t._messages = _messages
}

func (t *TL_updateReadMessagesContents) Get_messages() []byte {
	return t._messages
}

func (t *TL_updateReadMessagesContents) Set_pts(_pts int32) {
	t._pts = _pts
}

func (t *TL_updateReadMessagesContents) Get_pts() int32 {
	return t._pts
}

func (t *TL_updateReadMessagesContents) Set_pts_count(_pts_count int32) {
	t._pts_count = _pts_count
}

func (t *TL_updateReadMessagesContents) Get_pts_count() int32 {
	return t._pts_count
}

func New_TL_updateReadMessagesContents() *TL_updateReadMessagesContents {
	return &TL_updateReadMessagesContents{}
}

func (t *TL_updateReadMessagesContents) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updateReadMessagesContents))
	ec.Bytes(t.Get_messages())
	ec.Int(t.Get_pts())
	ec.Int(t.Get_pts_count())

	return ec.GetBuffer()
}

func (t *TL_updateReadMessagesContents) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._messages = dc.Bytes(16)
	t._pts = dc.Int()
	t._pts_count = dc.Int()

}

// updateChannelTooLong#eb0467fb
type TL_updateChannelTooLong struct {
	_flags      []byte
	_channel_id int32
	_pts        []byte
}

func (t *TL_updateChannelTooLong) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_updateChannelTooLong) Get_flags() []byte {
	return t._flags
}

func (t *TL_updateChannelTooLong) Set_channel_id(_channel_id int32) {
	t._channel_id = _channel_id
}

func (t *TL_updateChannelTooLong) Get_channel_id() int32 {
	return t._channel_id
}

func (t *TL_updateChannelTooLong) Set_pts(_pts []byte) {
	t._pts = _pts
}

func (t *TL_updateChannelTooLong) Get_pts() []byte {
	return t._pts
}

func New_TL_updateChannelTooLong() *TL_updateChannelTooLong {
	return &TL_updateChannelTooLong{}
}

func (t *TL_updateChannelTooLong) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updateChannelTooLong))
	ec.Int(t.Get_channel_id())
	ec.Bytes(t.Get_pts())

	return ec.GetBuffer()
}

func (t *TL_updateChannelTooLong) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._channel_id = dc.Int()
	t._pts = dc.Bytes(16)

}

// updateChannel#b6d45656
type TL_updateChannel struct {
	_channel_id int32
}

func (t *TL_updateChannel) Set_channel_id(_channel_id int32) {
	t._channel_id = _channel_id
}

func (t *TL_updateChannel) Get_channel_id() int32 {
	return t._channel_id
}

func New_TL_updateChannel() *TL_updateChannel {
	return &TL_updateChannel{}
}

func (t *TL_updateChannel) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updateChannel))
	ec.Int(t.Get_channel_id())

	return ec.GetBuffer()
}

func (t *TL_updateChannel) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._channel_id = dc.Int()

}

// updateNewChannelMessage#62ba04d9
type TL_updateNewChannelMessage struct {
	_message   []byte
	_pts       int32
	_pts_count int32
}

func (t *TL_updateNewChannelMessage) Set_message(_message []byte) {
	t._message = _message
}

func (t *TL_updateNewChannelMessage) Get_message() []byte {
	return t._message
}

func (t *TL_updateNewChannelMessage) Set_pts(_pts int32) {
	t._pts = _pts
}

func (t *TL_updateNewChannelMessage) Get_pts() int32 {
	return t._pts
}

func (t *TL_updateNewChannelMessage) Set_pts_count(_pts_count int32) {
	t._pts_count = _pts_count
}

func (t *TL_updateNewChannelMessage) Get_pts_count() int32 {
	return t._pts_count
}

func New_TL_updateNewChannelMessage() *TL_updateNewChannelMessage {
	return &TL_updateNewChannelMessage{}
}

func (t *TL_updateNewChannelMessage) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updateNewChannelMessage))
	ec.Bytes(t.Get_message())
	ec.Int(t.Get_pts())
	ec.Int(t.Get_pts_count())

	return ec.GetBuffer()
}

func (t *TL_updateNewChannelMessage) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._message = dc.Bytes(16)
	t._pts = dc.Int()
	t._pts_count = dc.Int()

}

// updateReadChannelInbox#4214f37f
type TL_updateReadChannelInbox struct {
	_channel_id int32
	_max_id     int32
}

func (t *TL_updateReadChannelInbox) Set_channel_id(_channel_id int32) {
	t._channel_id = _channel_id
}

func (t *TL_updateReadChannelInbox) Get_channel_id() int32 {
	return t._channel_id
}

func (t *TL_updateReadChannelInbox) Set_max_id(_max_id int32) {
	t._max_id = _max_id
}

func (t *TL_updateReadChannelInbox) Get_max_id() int32 {
	return t._max_id
}

func New_TL_updateReadChannelInbox() *TL_updateReadChannelInbox {
	return &TL_updateReadChannelInbox{}
}

func (t *TL_updateReadChannelInbox) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updateReadChannelInbox))
	ec.Int(t.Get_channel_id())
	ec.Int(t.Get_max_id())

	return ec.GetBuffer()
}

func (t *TL_updateReadChannelInbox) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._channel_id = dc.Int()
	t._max_id = dc.Int()

}

// updateDeleteChannelMessages#c37521c9
type TL_updateDeleteChannelMessages struct {
	_channel_id int32
	_messages   []byte
	_pts        int32
	_pts_count  int32
}

func (t *TL_updateDeleteChannelMessages) Set_channel_id(_channel_id int32) {
	t._channel_id = _channel_id
}

func (t *TL_updateDeleteChannelMessages) Get_channel_id() int32 {
	return t._channel_id
}

func (t *TL_updateDeleteChannelMessages) Set_messages(_messages []byte) {
	t._messages = _messages
}

func (t *TL_updateDeleteChannelMessages) Get_messages() []byte {
	return t._messages
}

func (t *TL_updateDeleteChannelMessages) Set_pts(_pts int32) {
	t._pts = _pts
}

func (t *TL_updateDeleteChannelMessages) Get_pts() int32 {
	return t._pts
}

func (t *TL_updateDeleteChannelMessages) Set_pts_count(_pts_count int32) {
	t._pts_count = _pts_count
}

func (t *TL_updateDeleteChannelMessages) Get_pts_count() int32 {
	return t._pts_count
}

func New_TL_updateDeleteChannelMessages() *TL_updateDeleteChannelMessages {
	return &TL_updateDeleteChannelMessages{}
}

func (t *TL_updateDeleteChannelMessages) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updateDeleteChannelMessages))
	ec.Int(t.Get_channel_id())
	ec.Bytes(t.Get_messages())
	ec.Int(t.Get_pts())
	ec.Int(t.Get_pts_count())

	return ec.GetBuffer()
}

func (t *TL_updateDeleteChannelMessages) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._channel_id = dc.Int()
	t._messages = dc.Bytes(16)
	t._pts = dc.Int()
	t._pts_count = dc.Int()

}

// updateChannelMessageViews#98a12b4b
type TL_updateChannelMessageViews struct {
	_channel_id int32
	_id         int32
	_views      int32
}

func (t *TL_updateChannelMessageViews) Set_channel_id(_channel_id int32) {
	t._channel_id = _channel_id
}

func (t *TL_updateChannelMessageViews) Get_channel_id() int32 {
	return t._channel_id
}

func (t *TL_updateChannelMessageViews) Set_id(_id int32) {
	t._id = _id
}

func (t *TL_updateChannelMessageViews) Get_id() int32 {
	return t._id
}

func (t *TL_updateChannelMessageViews) Set_views(_views int32) {
	t._views = _views
}

func (t *TL_updateChannelMessageViews) Get_views() int32 {
	return t._views
}

func New_TL_updateChannelMessageViews() *TL_updateChannelMessageViews {
	return &TL_updateChannelMessageViews{}
}

func (t *TL_updateChannelMessageViews) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updateChannelMessageViews))
	ec.Int(t.Get_channel_id())
	ec.Int(t.Get_id())
	ec.Int(t.Get_views())

	return ec.GetBuffer()
}

func (t *TL_updateChannelMessageViews) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._channel_id = dc.Int()
	t._id = dc.Int()
	t._views = dc.Int()

}

// updateChatAdmins#6e947941
type TL_updateChatAdmins struct {
	_chat_id int32
	_enabled bool
	_version int32
}

func (t *TL_updateChatAdmins) Set_chat_id(_chat_id int32) {
	t._chat_id = _chat_id
}

func (t *TL_updateChatAdmins) Get_chat_id() int32 {
	return t._chat_id
}

func (t *TL_updateChatAdmins) Set_enabled(_enabled bool) {
	t._enabled = _enabled
}

func (t *TL_updateChatAdmins) Get_enabled() bool {
	return t._enabled
}

func (t *TL_updateChatAdmins) Set_version(_version int32) {
	t._version = _version
}

func (t *TL_updateChatAdmins) Get_version() int32 {
	return t._version
}

func New_TL_updateChatAdmins() *TL_updateChatAdmins {
	return &TL_updateChatAdmins{}
}

func (t *TL_updateChatAdmins) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updateChatAdmins))
	ec.Int(t.Get_chat_id())
	ec.Bool(t.Get_enabled())
	ec.Int(t.Get_version())

	return ec.GetBuffer()
}

func (t *TL_updateChatAdmins) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._chat_id = dc.Int()
	t._enabled = dc.Bool()
	t._version = dc.Int()

}

// updateChatParticipantAdmin#b6901959
type TL_updateChatParticipantAdmin struct {
	_chat_id  int32
	_user_id  int32
	_is_admin bool
	_version  int32
}

func (t *TL_updateChatParticipantAdmin) Set_chat_id(_chat_id int32) {
	t._chat_id = _chat_id
}

func (t *TL_updateChatParticipantAdmin) Get_chat_id() int32 {
	return t._chat_id
}

func (t *TL_updateChatParticipantAdmin) Set_user_id(_user_id int32) {
	t._user_id = _user_id
}

func (t *TL_updateChatParticipantAdmin) Get_user_id() int32 {
	return t._user_id
}

func (t *TL_updateChatParticipantAdmin) Set_is_admin(_is_admin bool) {
	t._is_admin = _is_admin
}

func (t *TL_updateChatParticipantAdmin) Get_is_admin() bool {
	return t._is_admin
}

func (t *TL_updateChatParticipantAdmin) Set_version(_version int32) {
	t._version = _version
}

func (t *TL_updateChatParticipantAdmin) Get_version() int32 {
	return t._version
}

func New_TL_updateChatParticipantAdmin() *TL_updateChatParticipantAdmin {
	return &TL_updateChatParticipantAdmin{}
}

func (t *TL_updateChatParticipantAdmin) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updateChatParticipantAdmin))
	ec.Int(t.Get_chat_id())
	ec.Int(t.Get_user_id())
	ec.Bool(t.Get_is_admin())
	ec.Int(t.Get_version())

	return ec.GetBuffer()
}

func (t *TL_updateChatParticipantAdmin) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._chat_id = dc.Int()
	t._user_id = dc.Int()
	t._is_admin = dc.Bool()
	t._version = dc.Int()

}

// updateNewStickerSet#688a30aa
type TL_updateNewStickerSet struct {
	_stickerset []byte
}

func (t *TL_updateNewStickerSet) Set_stickerset(_stickerset []byte) {
	t._stickerset = _stickerset
}

func (t *TL_updateNewStickerSet) Get_stickerset() []byte {
	return t._stickerset
}

func New_TL_updateNewStickerSet() *TL_updateNewStickerSet {
	return &TL_updateNewStickerSet{}
}

func (t *TL_updateNewStickerSet) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updateNewStickerSet))
	ec.Bytes(t.Get_stickerset())

	return ec.GetBuffer()
}

func (t *TL_updateNewStickerSet) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._stickerset = dc.Bytes(16)

}

// updateStickerSetsOrder#0bb2d201
type TL_updateStickerSetsOrder struct {
	_flags []byte
	_masks []byte
	_order []byte
}

func (t *TL_updateStickerSetsOrder) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_updateStickerSetsOrder) Get_flags() []byte {
	return t._flags
}

func (t *TL_updateStickerSetsOrder) Set_masks(_masks []byte) {
	t._masks = _masks
}

func (t *TL_updateStickerSetsOrder) Get_masks() []byte {
	return t._masks
}

func (t *TL_updateStickerSetsOrder) Set_order(_order []byte) {
	t._order = _order
}

func (t *TL_updateStickerSetsOrder) Get_order() []byte {
	return t._order
}

func New_TL_updateStickerSetsOrder() *TL_updateStickerSetsOrder {
	return &TL_updateStickerSetsOrder{}
}

func (t *TL_updateStickerSetsOrder) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updateStickerSetsOrder))
	ec.Bytes(t.Get_masks())
	ec.Bytes(t.Get_order())

	return ec.GetBuffer()
}

func (t *TL_updateStickerSetsOrder) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._masks = dc.Bytes(16)
	t._order = dc.Bytes(16)

}

// updateStickerSets#43ae3dec
type TL_updateStickerSets struct {
}

func New_TL_updateStickerSets() *TL_updateStickerSets {
	return &TL_updateStickerSets{}
}

func (t *TL_updateStickerSets) Encode() []byte {
	return nil
}

func (t *TL_updateStickerSets) Decode(b []byte) {

}

// updateSavedGifs#9375341e
type TL_updateSavedGifs struct {
}

func New_TL_updateSavedGifs() *TL_updateSavedGifs {
	return &TL_updateSavedGifs{}
}

func (t *TL_updateSavedGifs) Encode() []byte {
	return nil
}

func (t *TL_updateSavedGifs) Decode(b []byte) {

}

// updateBotInlineQuery#54826690
type TL_updateBotInlineQuery struct {
	_flags    []byte
	_query_id int64
	_user_id  int32
	_query    string
	_geo      []byte
	_offset   string
}

func (t *TL_updateBotInlineQuery) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_updateBotInlineQuery) Get_flags() []byte {
	return t._flags
}

func (t *TL_updateBotInlineQuery) Set_query_id(_query_id int64) {
	t._query_id = _query_id
}

func (t *TL_updateBotInlineQuery) Get_query_id() int64 {
	return t._query_id
}

func (t *TL_updateBotInlineQuery) Set_user_id(_user_id int32) {
	t._user_id = _user_id
}

func (t *TL_updateBotInlineQuery) Get_user_id() int32 {
	return t._user_id
}

func (t *TL_updateBotInlineQuery) Set_query(_query string) {
	t._query = _query
}

func (t *TL_updateBotInlineQuery) Get_query() string {
	return t._query
}

func (t *TL_updateBotInlineQuery) Set_geo(_geo []byte) {
	t._geo = _geo
}

func (t *TL_updateBotInlineQuery) Get_geo() []byte {
	return t._geo
}

func (t *TL_updateBotInlineQuery) Set_offset(_offset string) {
	t._offset = _offset
}

func (t *TL_updateBotInlineQuery) Get_offset() string {
	return t._offset
}

func New_TL_updateBotInlineQuery() *TL_updateBotInlineQuery {
	return &TL_updateBotInlineQuery{}
}

func (t *TL_updateBotInlineQuery) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updateBotInlineQuery))
	ec.Long(t.Get_query_id())
	ec.Int(t.Get_user_id())
	ec.String(t.Get_query())
	ec.Bytes(t.Get_geo())
	ec.String(t.Get_offset())

	return ec.GetBuffer()
}

func (t *TL_updateBotInlineQuery) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._query_id = dc.Long()
	t._user_id = dc.Int()
	t._query = dc.String()
	t._geo = dc.Bytes(16)
	t._offset = dc.String()

}

// updateBotInlineSend#0e48f964
type TL_updateBotInlineSend struct {
	_flags   []byte
	_user_id int32
	_query   string
	_geo     []byte
	_id      string
	_msg_id  []byte
}

func (t *TL_updateBotInlineSend) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_updateBotInlineSend) Get_flags() []byte {
	return t._flags
}

func (t *TL_updateBotInlineSend) Set_user_id(_user_id int32) {
	t._user_id = _user_id
}

func (t *TL_updateBotInlineSend) Get_user_id() int32 {
	return t._user_id
}

func (t *TL_updateBotInlineSend) Set_query(_query string) {
	t._query = _query
}

func (t *TL_updateBotInlineSend) Get_query() string {
	return t._query
}

func (t *TL_updateBotInlineSend) Set_geo(_geo []byte) {
	t._geo = _geo
}

func (t *TL_updateBotInlineSend) Get_geo() []byte {
	return t._geo
}

func (t *TL_updateBotInlineSend) Set_id(_id string) {
	t._id = _id
}

func (t *TL_updateBotInlineSend) Get_id() string {
	return t._id
}

func (t *TL_updateBotInlineSend) Set_msg_id(_msg_id []byte) {
	t._msg_id = _msg_id
}

func (t *TL_updateBotInlineSend) Get_msg_id() []byte {
	return t._msg_id
}

func New_TL_updateBotInlineSend() *TL_updateBotInlineSend {
	return &TL_updateBotInlineSend{}
}

func (t *TL_updateBotInlineSend) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updateBotInlineSend))
	ec.Int(t.Get_user_id())
	ec.String(t.Get_query())
	ec.Bytes(t.Get_geo())
	ec.String(t.Get_id())
	ec.Bytes(t.Get_msg_id())

	return ec.GetBuffer()
}

func (t *TL_updateBotInlineSend) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._user_id = dc.Int()
	t._query = dc.String()
	t._geo = dc.Bytes(16)
	t._id = dc.String()
	t._msg_id = dc.Bytes(16)

}

// updateEditChannelMessage#1b3f4df7
type TL_updateEditChannelMessage struct {
	_message   []byte
	_pts       int32
	_pts_count int32
}

func (t *TL_updateEditChannelMessage) Set_message(_message []byte) {
	t._message = _message
}

func (t *TL_updateEditChannelMessage) Get_message() []byte {
	return t._message
}

func (t *TL_updateEditChannelMessage) Set_pts(_pts int32) {
	t._pts = _pts
}

func (t *TL_updateEditChannelMessage) Get_pts() int32 {
	return t._pts
}

func (t *TL_updateEditChannelMessage) Set_pts_count(_pts_count int32) {
	t._pts_count = _pts_count
}

func (t *TL_updateEditChannelMessage) Get_pts_count() int32 {
	return t._pts_count
}

func New_TL_updateEditChannelMessage() *TL_updateEditChannelMessage {
	return &TL_updateEditChannelMessage{}
}

func (t *TL_updateEditChannelMessage) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updateEditChannelMessage))
	ec.Bytes(t.Get_message())
	ec.Int(t.Get_pts())
	ec.Int(t.Get_pts_count())

	return ec.GetBuffer()
}

func (t *TL_updateEditChannelMessage) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._message = dc.Bytes(16)
	t._pts = dc.Int()
	t._pts_count = dc.Int()

}

// updateChannelPinnedMessage#98592475
type TL_updateChannelPinnedMessage struct {
	_channel_id int32
	_id         int32
}

func (t *TL_updateChannelPinnedMessage) Set_channel_id(_channel_id int32) {
	t._channel_id = _channel_id
}

func (t *TL_updateChannelPinnedMessage) Get_channel_id() int32 {
	return t._channel_id
}

func (t *TL_updateChannelPinnedMessage) Set_id(_id int32) {
	t._id = _id
}

func (t *TL_updateChannelPinnedMessage) Get_id() int32 {
	return t._id
}

func New_TL_updateChannelPinnedMessage() *TL_updateChannelPinnedMessage {
	return &TL_updateChannelPinnedMessage{}
}

func (t *TL_updateChannelPinnedMessage) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updateChannelPinnedMessage))
	ec.Int(t.Get_channel_id())
	ec.Int(t.Get_id())

	return ec.GetBuffer()
}

func (t *TL_updateChannelPinnedMessage) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._channel_id = dc.Int()
	t._id = dc.Int()

}

// updateBotCallbackQuery#e73547e1
type TL_updateBotCallbackQuery struct {
	_flags           []byte
	_query_id        int64
	_user_id         int32
	_peer            []byte
	_msg_id          int32
	_chat_instance   int64
	_data            []byte
	_game_short_name []byte
}

func (t *TL_updateBotCallbackQuery) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_updateBotCallbackQuery) Get_flags() []byte {
	return t._flags
}

func (t *TL_updateBotCallbackQuery) Set_query_id(_query_id int64) {
	t._query_id = _query_id
}

func (t *TL_updateBotCallbackQuery) Get_query_id() int64 {
	return t._query_id
}

func (t *TL_updateBotCallbackQuery) Set_user_id(_user_id int32) {
	t._user_id = _user_id
}

func (t *TL_updateBotCallbackQuery) Get_user_id() int32 {
	return t._user_id
}

func (t *TL_updateBotCallbackQuery) Set_peer(_peer []byte) {
	t._peer = _peer
}

func (t *TL_updateBotCallbackQuery) Get_peer() []byte {
	return t._peer
}

func (t *TL_updateBotCallbackQuery) Set_msg_id(_msg_id int32) {
	t._msg_id = _msg_id
}

func (t *TL_updateBotCallbackQuery) Get_msg_id() int32 {
	return t._msg_id
}

func (t *TL_updateBotCallbackQuery) Set_chat_instance(_chat_instance int64) {
	t._chat_instance = _chat_instance
}

func (t *TL_updateBotCallbackQuery) Get_chat_instance() int64 {
	return t._chat_instance
}

func (t *TL_updateBotCallbackQuery) Set_data(_data []byte) {
	t._data = _data
}

func (t *TL_updateBotCallbackQuery) Get_data() []byte {
	return t._data
}

func (t *TL_updateBotCallbackQuery) Set_game_short_name(_game_short_name []byte) {
	t._game_short_name = _game_short_name
}

func (t *TL_updateBotCallbackQuery) Get_game_short_name() []byte {
	return t._game_short_name
}

func New_TL_updateBotCallbackQuery() *TL_updateBotCallbackQuery {
	return &TL_updateBotCallbackQuery{}
}

func (t *TL_updateBotCallbackQuery) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updateBotCallbackQuery))
	ec.Long(t.Get_query_id())
	ec.Int(t.Get_user_id())
	ec.Bytes(t.Get_peer())
	ec.Int(t.Get_msg_id())
	ec.Long(t.Get_chat_instance())
	ec.Bytes(t.Get_data())
	ec.Bytes(t.Get_game_short_name())

	return ec.GetBuffer()
}

func (t *TL_updateBotCallbackQuery) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._query_id = dc.Long()
	t._user_id = dc.Int()
	t._peer = dc.Bytes(16)
	t._msg_id = dc.Int()
	t._chat_instance = dc.Long()
	t._data = dc.Bytes(16)
	t._game_short_name = dc.Bytes(16)

}

// updateEditMessage#e40370a3
type TL_updateEditMessage struct {
	_message   []byte
	_pts       int32
	_pts_count int32
}

func (t *TL_updateEditMessage) Set_message(_message []byte) {
	t._message = _message
}

func (t *TL_updateEditMessage) Get_message() []byte {
	return t._message
}

func (t *TL_updateEditMessage) Set_pts(_pts int32) {
	t._pts = _pts
}

func (t *TL_updateEditMessage) Get_pts() int32 {
	return t._pts
}

func (t *TL_updateEditMessage) Set_pts_count(_pts_count int32) {
	t._pts_count = _pts_count
}

func (t *TL_updateEditMessage) Get_pts_count() int32 {
	return t._pts_count
}

func New_TL_updateEditMessage() *TL_updateEditMessage {
	return &TL_updateEditMessage{}
}

func (t *TL_updateEditMessage) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updateEditMessage))
	ec.Bytes(t.Get_message())
	ec.Int(t.Get_pts())
	ec.Int(t.Get_pts_count())

	return ec.GetBuffer()
}

func (t *TL_updateEditMessage) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._message = dc.Bytes(16)
	t._pts = dc.Int()
	t._pts_count = dc.Int()

}

// updateInlineBotCallbackQuery#f9d27a5a
type TL_updateInlineBotCallbackQuery struct {
	_flags           []byte
	_query_id        int64
	_user_id         int32
	_msg_id          []byte
	_chat_instance   int64
	_data            []byte
	_game_short_name []byte
}

func (t *TL_updateInlineBotCallbackQuery) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_updateInlineBotCallbackQuery) Get_flags() []byte {
	return t._flags
}

func (t *TL_updateInlineBotCallbackQuery) Set_query_id(_query_id int64) {
	t._query_id = _query_id
}

func (t *TL_updateInlineBotCallbackQuery) Get_query_id() int64 {
	return t._query_id
}

func (t *TL_updateInlineBotCallbackQuery) Set_user_id(_user_id int32) {
	t._user_id = _user_id
}

func (t *TL_updateInlineBotCallbackQuery) Get_user_id() int32 {
	return t._user_id
}

func (t *TL_updateInlineBotCallbackQuery) Set_msg_id(_msg_id []byte) {
	t._msg_id = _msg_id
}

func (t *TL_updateInlineBotCallbackQuery) Get_msg_id() []byte {
	return t._msg_id
}

func (t *TL_updateInlineBotCallbackQuery) Set_chat_instance(_chat_instance int64) {
	t._chat_instance = _chat_instance
}

func (t *TL_updateInlineBotCallbackQuery) Get_chat_instance() int64 {
	return t._chat_instance
}

func (t *TL_updateInlineBotCallbackQuery) Set_data(_data []byte) {
	t._data = _data
}

func (t *TL_updateInlineBotCallbackQuery) Get_data() []byte {
	return t._data
}

func (t *TL_updateInlineBotCallbackQuery) Set_game_short_name(_game_short_name []byte) {
	t._game_short_name = _game_short_name
}

func (t *TL_updateInlineBotCallbackQuery) Get_game_short_name() []byte {
	return t._game_short_name
}

func New_TL_updateInlineBotCallbackQuery() *TL_updateInlineBotCallbackQuery {
	return &TL_updateInlineBotCallbackQuery{}
}

func (t *TL_updateInlineBotCallbackQuery) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updateInlineBotCallbackQuery))
	ec.Long(t.Get_query_id())
	ec.Int(t.Get_user_id())
	ec.Bytes(t.Get_msg_id())
	ec.Long(t.Get_chat_instance())
	ec.Bytes(t.Get_data())
	ec.Bytes(t.Get_game_short_name())

	return ec.GetBuffer()
}

func (t *TL_updateInlineBotCallbackQuery) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._query_id = dc.Long()
	t._user_id = dc.Int()
	t._msg_id = dc.Bytes(16)
	t._chat_instance = dc.Long()
	t._data = dc.Bytes(16)
	t._game_short_name = dc.Bytes(16)

}

// updateReadChannelOutbox#25d6c9c7
type TL_updateReadChannelOutbox struct {
	_channel_id int32
	_max_id     int32
}

func (t *TL_updateReadChannelOutbox) Set_channel_id(_channel_id int32) {
	t._channel_id = _channel_id
}

func (t *TL_updateReadChannelOutbox) Get_channel_id() int32 {
	return t._channel_id
}

func (t *TL_updateReadChannelOutbox) Set_max_id(_max_id int32) {
	t._max_id = _max_id
}

func (t *TL_updateReadChannelOutbox) Get_max_id() int32 {
	return t._max_id
}

func New_TL_updateReadChannelOutbox() *TL_updateReadChannelOutbox {
	return &TL_updateReadChannelOutbox{}
}

func (t *TL_updateReadChannelOutbox) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updateReadChannelOutbox))
	ec.Int(t.Get_channel_id())
	ec.Int(t.Get_max_id())

	return ec.GetBuffer()
}

func (t *TL_updateReadChannelOutbox) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._channel_id = dc.Int()
	t._max_id = dc.Int()

}

// updateDraftMessage#ee2bb969
type TL_updateDraftMessage struct {
	_peer  []byte
	_draft []byte
}

func (t *TL_updateDraftMessage) Set_peer(_peer []byte) {
	t._peer = _peer
}

func (t *TL_updateDraftMessage) Get_peer() []byte {
	return t._peer
}

func (t *TL_updateDraftMessage) Set_draft(_draft []byte) {
	t._draft = _draft
}

func (t *TL_updateDraftMessage) Get_draft() []byte {
	return t._draft
}

func New_TL_updateDraftMessage() *TL_updateDraftMessage {
	return &TL_updateDraftMessage{}
}

func (t *TL_updateDraftMessage) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updateDraftMessage))
	ec.Bytes(t.Get_peer())
	ec.Bytes(t.Get_draft())

	return ec.GetBuffer()
}

func (t *TL_updateDraftMessage) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._peer = dc.Bytes(16)
	t._draft = dc.Bytes(16)

}

// updateReadFeaturedStickers#571d2742
type TL_updateReadFeaturedStickers struct {
}

func New_TL_updateReadFeaturedStickers() *TL_updateReadFeaturedStickers {
	return &TL_updateReadFeaturedStickers{}
}

func (t *TL_updateReadFeaturedStickers) Encode() []byte {
	return nil
}

func (t *TL_updateReadFeaturedStickers) Decode(b []byte) {

}

// updateRecentStickers#9a422c20
type TL_updateRecentStickers struct {
}

func New_TL_updateRecentStickers() *TL_updateRecentStickers {
	return &TL_updateRecentStickers{}
}

func (t *TL_updateRecentStickers) Encode() []byte {
	return nil
}

func (t *TL_updateRecentStickers) Decode(b []byte) {

}

// updateConfig#a229dd06
type TL_updateConfig struct {
}

func New_TL_updateConfig() *TL_updateConfig {
	return &TL_updateConfig{}
}

func (t *TL_updateConfig) Encode() []byte {
	return nil
}

func (t *TL_updateConfig) Decode(b []byte) {

}

// updatePtsChanged#3354678f
type TL_updatePtsChanged struct {
}

func New_TL_updatePtsChanged() *TL_updatePtsChanged {
	return &TL_updatePtsChanged{}
}

func (t *TL_updatePtsChanged) Encode() []byte {
	return nil
}

func (t *TL_updatePtsChanged) Decode(b []byte) {

}

// updateChannelWebPage#40771900
type TL_updateChannelWebPage struct {
	_channel_id int32
	_webpage    []byte
	_pts        int32
	_pts_count  int32
}

func (t *TL_updateChannelWebPage) Set_channel_id(_channel_id int32) {
	t._channel_id = _channel_id
}

func (t *TL_updateChannelWebPage) Get_channel_id() int32 {
	return t._channel_id
}

func (t *TL_updateChannelWebPage) Set_webpage(_webpage []byte) {
	t._webpage = _webpage
}

func (t *TL_updateChannelWebPage) Get_webpage() []byte {
	return t._webpage
}

func (t *TL_updateChannelWebPage) Set_pts(_pts int32) {
	t._pts = _pts
}

func (t *TL_updateChannelWebPage) Get_pts() int32 {
	return t._pts
}

func (t *TL_updateChannelWebPage) Set_pts_count(_pts_count int32) {
	t._pts_count = _pts_count
}

func (t *TL_updateChannelWebPage) Get_pts_count() int32 {
	return t._pts_count
}

func New_TL_updateChannelWebPage() *TL_updateChannelWebPage {
	return &TL_updateChannelWebPage{}
}

func (t *TL_updateChannelWebPage) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updateChannelWebPage))
	ec.Int(t.Get_channel_id())
	ec.Bytes(t.Get_webpage())
	ec.Int(t.Get_pts())
	ec.Int(t.Get_pts_count())

	return ec.GetBuffer()
}

func (t *TL_updateChannelWebPage) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._channel_id = dc.Int()
	t._webpage = dc.Bytes(16)
	t._pts = dc.Int()
	t._pts_count = dc.Int()

}

// updateDialogPinned#d711a2cc
type TL_updateDialogPinned struct {
	_flags  []byte
	_pinned []byte
	_peer   []byte
}

func (t *TL_updateDialogPinned) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_updateDialogPinned) Get_flags() []byte {
	return t._flags
}

func (t *TL_updateDialogPinned) Set_pinned(_pinned []byte) {
	t._pinned = _pinned
}

func (t *TL_updateDialogPinned) Get_pinned() []byte {
	return t._pinned
}

func (t *TL_updateDialogPinned) Set_peer(_peer []byte) {
	t._peer = _peer
}

func (t *TL_updateDialogPinned) Get_peer() []byte {
	return t._peer
}

func New_TL_updateDialogPinned() *TL_updateDialogPinned {
	return &TL_updateDialogPinned{}
}

func (t *TL_updateDialogPinned) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updateDialogPinned))
	ec.Bytes(t.Get_pinned())
	ec.Bytes(t.Get_peer())

	return ec.GetBuffer()
}

func (t *TL_updateDialogPinned) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._pinned = dc.Bytes(16)
	t._peer = dc.Bytes(16)

}

// updatePinnedDialogs#d8caf68d
type TL_updatePinnedDialogs struct {
	_flags []byte
	_order []byte
}

func (t *TL_updatePinnedDialogs) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_updatePinnedDialogs) Get_flags() []byte {
	return t._flags
}

func (t *TL_updatePinnedDialogs) Set_order(_order []byte) {
	t._order = _order
}

func (t *TL_updatePinnedDialogs) Get_order() []byte {
	return t._order
}

func New_TL_updatePinnedDialogs() *TL_updatePinnedDialogs {
	return &TL_updatePinnedDialogs{}
}

func (t *TL_updatePinnedDialogs) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updatePinnedDialogs))
	ec.Bytes(t.Get_order())

	return ec.GetBuffer()
}

func (t *TL_updatePinnedDialogs) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._order = dc.Bytes(16)

}

// updateBotWebhookJSON#8317c0c3
type TL_updateBotWebhookJSON struct {
	_data []byte
}

func (t *TL_updateBotWebhookJSON) Set_data(_data []byte) {
	t._data = _data
}

func (t *TL_updateBotWebhookJSON) Get_data() []byte {
	return t._data
}

func New_TL_updateBotWebhookJSON() *TL_updateBotWebhookJSON {
	return &TL_updateBotWebhookJSON{}
}

func (t *TL_updateBotWebhookJSON) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updateBotWebhookJSON))
	ec.Bytes(t.Get_data())

	return ec.GetBuffer()
}

func (t *TL_updateBotWebhookJSON) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._data = dc.Bytes(16)

}

// updateBotWebhookJSONQuery#9b9240a6
type TL_updateBotWebhookJSONQuery struct {
	_query_id int64
	_data     []byte
	_timeout  int32
}

func (t *TL_updateBotWebhookJSONQuery) Set_query_id(_query_id int64) {
	t._query_id = _query_id
}

func (t *TL_updateBotWebhookJSONQuery) Get_query_id() int64 {
	return t._query_id
}

func (t *TL_updateBotWebhookJSONQuery) Set_data(_data []byte) {
	t._data = _data
}

func (t *TL_updateBotWebhookJSONQuery) Get_data() []byte {
	return t._data
}

func (t *TL_updateBotWebhookJSONQuery) Set_timeout(_timeout int32) {
	t._timeout = _timeout
}

func (t *TL_updateBotWebhookJSONQuery) Get_timeout() int32 {
	return t._timeout
}

func New_TL_updateBotWebhookJSONQuery() *TL_updateBotWebhookJSONQuery {
	return &TL_updateBotWebhookJSONQuery{}
}

func (t *TL_updateBotWebhookJSONQuery) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updateBotWebhookJSONQuery))
	ec.Long(t.Get_query_id())
	ec.Bytes(t.Get_data())
	ec.Int(t.Get_timeout())

	return ec.GetBuffer()
}

func (t *TL_updateBotWebhookJSONQuery) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._query_id = dc.Long()
	t._data = dc.Bytes(16)
	t._timeout = dc.Int()

}

// updateBotShippingQuery#e0cdc940
type TL_updateBotShippingQuery struct {
	_query_id         int64
	_user_id          int32
	_payload          []byte
	_shipping_address []byte
}

func (t *TL_updateBotShippingQuery) Set_query_id(_query_id int64) {
	t._query_id = _query_id
}

func (t *TL_updateBotShippingQuery) Get_query_id() int64 {
	return t._query_id
}

func (t *TL_updateBotShippingQuery) Set_user_id(_user_id int32) {
	t._user_id = _user_id
}

func (t *TL_updateBotShippingQuery) Get_user_id() int32 {
	return t._user_id
}

func (t *TL_updateBotShippingQuery) Set_payload(_payload []byte) {
	t._payload = _payload
}

func (t *TL_updateBotShippingQuery) Get_payload() []byte {
	return t._payload
}

func (t *TL_updateBotShippingQuery) Set_shipping_address(_shipping_address []byte) {
	t._shipping_address = _shipping_address
}

func (t *TL_updateBotShippingQuery) Get_shipping_address() []byte {
	return t._shipping_address
}

func New_TL_updateBotShippingQuery() *TL_updateBotShippingQuery {
	return &TL_updateBotShippingQuery{}
}

func (t *TL_updateBotShippingQuery) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updateBotShippingQuery))
	ec.Long(t.Get_query_id())
	ec.Int(t.Get_user_id())
	ec.Bytes(t.Get_payload())
	ec.Bytes(t.Get_shipping_address())

	return ec.GetBuffer()
}

func (t *TL_updateBotShippingQuery) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._query_id = dc.Long()
	t._user_id = dc.Int()
	t._payload = dc.Bytes(16)
	t._shipping_address = dc.Bytes(16)

}

// updateBotPrecheckoutQuery#5d2f3aa9
type TL_updateBotPrecheckoutQuery struct {
	_flags              []byte
	_query_id           int64
	_user_id            int32
	_payload            []byte
	_info               []byte
	_shipping_option_id []byte
	_currency           string
	_total_amount       int64
}

func (t *TL_updateBotPrecheckoutQuery) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_updateBotPrecheckoutQuery) Get_flags() []byte {
	return t._flags
}

func (t *TL_updateBotPrecheckoutQuery) Set_query_id(_query_id int64) {
	t._query_id = _query_id
}

func (t *TL_updateBotPrecheckoutQuery) Get_query_id() int64 {
	return t._query_id
}

func (t *TL_updateBotPrecheckoutQuery) Set_user_id(_user_id int32) {
	t._user_id = _user_id
}

func (t *TL_updateBotPrecheckoutQuery) Get_user_id() int32 {
	return t._user_id
}

func (t *TL_updateBotPrecheckoutQuery) Set_payload(_payload []byte) {
	t._payload = _payload
}

func (t *TL_updateBotPrecheckoutQuery) Get_payload() []byte {
	return t._payload
}

func (t *TL_updateBotPrecheckoutQuery) Set_info(_info []byte) {
	t._info = _info
}

func (t *TL_updateBotPrecheckoutQuery) Get_info() []byte {
	return t._info
}

func (t *TL_updateBotPrecheckoutQuery) Set_shipping_option_id(_shipping_option_id []byte) {
	t._shipping_option_id = _shipping_option_id
}

func (t *TL_updateBotPrecheckoutQuery) Get_shipping_option_id() []byte {
	return t._shipping_option_id
}

func (t *TL_updateBotPrecheckoutQuery) Set_currency(_currency string) {
	t._currency = _currency
}

func (t *TL_updateBotPrecheckoutQuery) Get_currency() string {
	return t._currency
}

func (t *TL_updateBotPrecheckoutQuery) Set_total_amount(_total_amount int64) {
	t._total_amount = _total_amount
}

func (t *TL_updateBotPrecheckoutQuery) Get_total_amount() int64 {
	return t._total_amount
}

func New_TL_updateBotPrecheckoutQuery() *TL_updateBotPrecheckoutQuery {
	return &TL_updateBotPrecheckoutQuery{}
}

func (t *TL_updateBotPrecheckoutQuery) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updateBotPrecheckoutQuery))
	ec.Long(t.Get_query_id())
	ec.Int(t.Get_user_id())
	ec.Bytes(t.Get_payload())
	ec.Bytes(t.Get_info())
	ec.Bytes(t.Get_shipping_option_id())
	ec.String(t.Get_currency())
	ec.Long(t.Get_total_amount())

	return ec.GetBuffer()
}

func (t *TL_updateBotPrecheckoutQuery) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._query_id = dc.Long()
	t._user_id = dc.Int()
	t._payload = dc.Bytes(16)
	t._info = dc.Bytes(16)
	t._shipping_option_id = dc.Bytes(16)
	t._currency = dc.String()
	t._total_amount = dc.Long()

}

// updatePhoneCall#ab0f6b1e
type TL_updatePhoneCall struct {
	_phone_call []byte
}

func (t *TL_updatePhoneCall) Set_phone_call(_phone_call []byte) {
	t._phone_call = _phone_call
}

func (t *TL_updatePhoneCall) Get_phone_call() []byte {
	return t._phone_call
}

func New_TL_updatePhoneCall() *TL_updatePhoneCall {
	return &TL_updatePhoneCall{}
}

func (t *TL_updatePhoneCall) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updatePhoneCall))
	ec.Bytes(t.Get_phone_call())

	return ec.GetBuffer()
}

func (t *TL_updatePhoneCall) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._phone_call = dc.Bytes(16)

}

// updateLangPackTooLong#10c2404b
type TL_updateLangPackTooLong struct {
}

func New_TL_updateLangPackTooLong() *TL_updateLangPackTooLong {
	return &TL_updateLangPackTooLong{}
}

func (t *TL_updateLangPackTooLong) Encode() []byte {
	return nil
}

func (t *TL_updateLangPackTooLong) Decode(b []byte) {

}

// updateLangPack#56022f4d
type TL_updateLangPack struct {
	_difference []byte
}

func (t *TL_updateLangPack) Set_difference(_difference []byte) {
	t._difference = _difference
}

func (t *TL_updateLangPack) Get_difference() []byte {
	return t._difference
}

func New_TL_updateLangPack() *TL_updateLangPack {
	return &TL_updateLangPack{}
}

func (t *TL_updateLangPack) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updateLangPack))
	ec.Bytes(t.Get_difference())

	return ec.GetBuffer()
}

func (t *TL_updateLangPack) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._difference = dc.Bytes(16)

}

// updateFavedStickers#e511996d
type TL_updateFavedStickers struct {
}

func New_TL_updateFavedStickers() *TL_updateFavedStickers {
	return &TL_updateFavedStickers{}
}

func (t *TL_updateFavedStickers) Encode() []byte {
	return nil
}

func (t *TL_updateFavedStickers) Decode(b []byte) {

}

// updateChannelReadMessagesContents#89893b45
type TL_updateChannelReadMessagesContents struct {
	_channel_id int32
	_messages   []byte
}

func (t *TL_updateChannelReadMessagesContents) Set_channel_id(_channel_id int32) {
	t._channel_id = _channel_id
}

func (t *TL_updateChannelReadMessagesContents) Get_channel_id() int32 {
	return t._channel_id
}

func (t *TL_updateChannelReadMessagesContents) Set_messages(_messages []byte) {
	t._messages = _messages
}

func (t *TL_updateChannelReadMessagesContents) Get_messages() []byte {
	return t._messages
}

func New_TL_updateChannelReadMessagesContents() *TL_updateChannelReadMessagesContents {
	return &TL_updateChannelReadMessagesContents{}
}

func (t *TL_updateChannelReadMessagesContents) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updateChannelReadMessagesContents))
	ec.Int(t.Get_channel_id())
	ec.Bytes(t.Get_messages())

	return ec.GetBuffer()
}

func (t *TL_updateChannelReadMessagesContents) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._channel_id = dc.Int()
	t._messages = dc.Bytes(16)

}

// updateContactsReset#7084a7be
type TL_updateContactsReset struct {
}

func New_TL_updateContactsReset() *TL_updateContactsReset {
	return &TL_updateContactsReset{}
}

func (t *TL_updateContactsReset) Encode() []byte {
	return nil
}

func (t *TL_updateContactsReset) Decode(b []byte) {

}

// updateChannelAvailableMessages#70db6837
type TL_updateChannelAvailableMessages struct {
	_channel_id       int32
	_available_min_id int32
}

func (t *TL_updateChannelAvailableMessages) Set_channel_id(_channel_id int32) {
	t._channel_id = _channel_id
}

func (t *TL_updateChannelAvailableMessages) Get_channel_id() int32 {
	return t._channel_id
}

func (t *TL_updateChannelAvailableMessages) Set_available_min_id(_available_min_id int32) {
	t._available_min_id = _available_min_id
}

func (t *TL_updateChannelAvailableMessages) Get_available_min_id() int32 {
	return t._available_min_id
}

func New_TL_updateChannelAvailableMessages() *TL_updateChannelAvailableMessages {
	return &TL_updateChannelAvailableMessages{}
}

func (t *TL_updateChannelAvailableMessages) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updateChannelAvailableMessages))
	ec.Int(t.Get_channel_id())
	ec.Int(t.Get_available_min_id())

	return ec.GetBuffer()
}

func (t *TL_updateChannelAvailableMessages) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._channel_id = dc.Int()
	t._available_min_id = dc.Int()

}

// updates_state#a56c2a3e
type TL_updates_state struct {
	_pts          int32
	_qts          int32
	_date         int32
	_seq          int32
	_unread_count int32
}

func (t *TL_updates_state) Set_pts(_pts int32) {
	t._pts = _pts
}

func (t *TL_updates_state) Get_pts() int32 {
	return t._pts
}

func (t *TL_updates_state) Set_qts(_qts int32) {
	t._qts = _qts
}

func (t *TL_updates_state) Get_qts() int32 {
	return t._qts
}

func (t *TL_updates_state) Set_date(_date int32) {
	t._date = _date
}

func (t *TL_updates_state) Get_date() int32 {
	return t._date
}

func (t *TL_updates_state) Set_seq(_seq int32) {
	t._seq = _seq
}

func (t *TL_updates_state) Get_seq() int32 {
	return t._seq
}

func (t *TL_updates_state) Set_unread_count(_unread_count int32) {
	t._unread_count = _unread_count
}

func (t *TL_updates_state) Get_unread_count() int32 {
	return t._unread_count
}

func New_TL_updates_state() *TL_updates_state {
	return &TL_updates_state{}
}

func (t *TL_updates_state) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updates_state))
	ec.Int(t.Get_pts())
	ec.Int(t.Get_qts())
	ec.Int(t.Get_date())
	ec.Int(t.Get_seq())
	ec.Int(t.Get_unread_count())

	return ec.GetBuffer()
}

func (t *TL_updates_state) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._pts = dc.Int()
	t._qts = dc.Int()
	t._date = dc.Int()
	t._seq = dc.Int()
	t._unread_count = dc.Int()

}

// updates_differenceEmpty#5d75a138
type TL_updates_differenceEmpty struct {
	_date int32
	_seq  int32
}

func (t *TL_updates_differenceEmpty) Set_date(_date int32) {
	t._date = _date
}

func (t *TL_updates_differenceEmpty) Get_date() int32 {
	return t._date
}

func (t *TL_updates_differenceEmpty) Set_seq(_seq int32) {
	t._seq = _seq
}

func (t *TL_updates_differenceEmpty) Get_seq() int32 {
	return t._seq
}

func New_TL_updates_differenceEmpty() *TL_updates_differenceEmpty {
	return &TL_updates_differenceEmpty{}
}

func (t *TL_updates_differenceEmpty) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updates_differenceEmpty))
	ec.Int(t.Get_date())
	ec.Int(t.Get_seq())

	return ec.GetBuffer()
}

func (t *TL_updates_differenceEmpty) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._date = dc.Int()
	t._seq = dc.Int()

}

// updates_difference#00f49ca0
type TL_updates_difference struct {
	_new_messages           []byte
	_new_encrypted_messages []byte
	_other_updates          []byte
	_chats                  []byte
	_users                  []byte
	_state                  []byte
}

func (t *TL_updates_difference) Set_new_messages(_new_messages []byte) {
	t._new_messages = _new_messages
}

func (t *TL_updates_difference) Get_new_messages() []byte {
	return t._new_messages
}

func (t *TL_updates_difference) Set_new_encrypted_messages(_new_encrypted_messages []byte) {
	t._new_encrypted_messages = _new_encrypted_messages
}

func (t *TL_updates_difference) Get_new_encrypted_messages() []byte {
	return t._new_encrypted_messages
}

func (t *TL_updates_difference) Set_other_updates(_other_updates []byte) {
	t._other_updates = _other_updates
}

func (t *TL_updates_difference) Get_other_updates() []byte {
	return t._other_updates
}

func (t *TL_updates_difference) Set_chats(_chats []byte) {
	t._chats = _chats
}

func (t *TL_updates_difference) Get_chats() []byte {
	return t._chats
}

func (t *TL_updates_difference) Set_users(_users []byte) {
	t._users = _users
}

func (t *TL_updates_difference) Get_users() []byte {
	return t._users
}

func (t *TL_updates_difference) Set_state(_state []byte) {
	t._state = _state
}

func (t *TL_updates_difference) Get_state() []byte {
	return t._state
}

func New_TL_updates_difference() *TL_updates_difference {
	return &TL_updates_difference{}
}

func (t *TL_updates_difference) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updates_difference))
	ec.Bytes(t.Get_new_messages())
	ec.Bytes(t.Get_new_encrypted_messages())
	ec.Bytes(t.Get_other_updates())
	ec.Bytes(t.Get_chats())
	ec.Bytes(t.Get_users())
	ec.Bytes(t.Get_state())

	return ec.GetBuffer()
}

func (t *TL_updates_difference) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._new_messages = dc.Bytes(16)
	t._new_encrypted_messages = dc.Bytes(16)
	t._other_updates = dc.Bytes(16)
	t._chats = dc.Bytes(16)
	t._users = dc.Bytes(16)
	t._state = dc.Bytes(16)

}

// updates_differenceSlice#a8fb1981
type TL_updates_differenceSlice struct {
	_new_messages           []byte
	_new_encrypted_messages []byte
	_other_updates          []byte
	_chats                  []byte
	_users                  []byte
	_intermediate_state     []byte
}

func (t *TL_updates_differenceSlice) Set_new_messages(_new_messages []byte) {
	t._new_messages = _new_messages
}

func (t *TL_updates_differenceSlice) Get_new_messages() []byte {
	return t._new_messages
}

func (t *TL_updates_differenceSlice) Set_new_encrypted_messages(_new_encrypted_messages []byte) {
	t._new_encrypted_messages = _new_encrypted_messages
}

func (t *TL_updates_differenceSlice) Get_new_encrypted_messages() []byte {
	return t._new_encrypted_messages
}

func (t *TL_updates_differenceSlice) Set_other_updates(_other_updates []byte) {
	t._other_updates = _other_updates
}

func (t *TL_updates_differenceSlice) Get_other_updates() []byte {
	return t._other_updates
}

func (t *TL_updates_differenceSlice) Set_chats(_chats []byte) {
	t._chats = _chats
}

func (t *TL_updates_differenceSlice) Get_chats() []byte {
	return t._chats
}

func (t *TL_updates_differenceSlice) Set_users(_users []byte) {
	t._users = _users
}

func (t *TL_updates_differenceSlice) Get_users() []byte {
	return t._users
}

func (t *TL_updates_differenceSlice) Set_intermediate_state(_intermediate_state []byte) {
	t._intermediate_state = _intermediate_state
}

func (t *TL_updates_differenceSlice) Get_intermediate_state() []byte {
	return t._intermediate_state
}

func New_TL_updates_differenceSlice() *TL_updates_differenceSlice {
	return &TL_updates_differenceSlice{}
}

func (t *TL_updates_differenceSlice) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updates_differenceSlice))
	ec.Bytes(t.Get_new_messages())
	ec.Bytes(t.Get_new_encrypted_messages())
	ec.Bytes(t.Get_other_updates())
	ec.Bytes(t.Get_chats())
	ec.Bytes(t.Get_users())
	ec.Bytes(t.Get_intermediate_state())

	return ec.GetBuffer()
}

func (t *TL_updates_differenceSlice) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._new_messages = dc.Bytes(16)
	t._new_encrypted_messages = dc.Bytes(16)
	t._other_updates = dc.Bytes(16)
	t._chats = dc.Bytes(16)
	t._users = dc.Bytes(16)
	t._intermediate_state = dc.Bytes(16)

}

// updates_differenceTooLong#4afe8f6d
type TL_updates_differenceTooLong struct {
	_pts int32
}

func (t *TL_updates_differenceTooLong) Set_pts(_pts int32) {
	t._pts = _pts
}

func (t *TL_updates_differenceTooLong) Get_pts() int32 {
	return t._pts
}

func New_TL_updates_differenceTooLong() *TL_updates_differenceTooLong {
	return &TL_updates_differenceTooLong{}
}

func (t *TL_updates_differenceTooLong) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updates_differenceTooLong))
	ec.Int(t.Get_pts())

	return ec.GetBuffer()
}

func (t *TL_updates_differenceTooLong) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._pts = dc.Int()

}

// updatesTooLong#e317af7e
type TL_updatesTooLong struct {
}

func New_TL_updatesTooLong() *TL_updatesTooLong {
	return &TL_updatesTooLong{}
}

func (t *TL_updatesTooLong) Encode() []byte {
	return nil
}

func (t *TL_updatesTooLong) Decode(b []byte) {

}

// updateShortMessage#914fbf11
type TL_updateShortMessage struct {
	_flags           []byte
	_out             []byte
	_mentioned       []byte
	_media_unread    []byte
	_silent          []byte
	_id              int32
	_user_id         int32
	_message         string
	_pts             int32
	_pts_count       int32
	_date            int32
	_fwd_from        []byte
	_via_bot_id      []byte
	_reply_to_msg_id []byte
	_entities        []byte
}

func (t *TL_updateShortMessage) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_updateShortMessage) Get_flags() []byte {
	return t._flags
}

func (t *TL_updateShortMessage) Set_out(_out []byte) {
	t._out = _out
}

func (t *TL_updateShortMessage) Get_out() []byte {
	return t._out
}

func (t *TL_updateShortMessage) Set_mentioned(_mentioned []byte) {
	t._mentioned = _mentioned
}

func (t *TL_updateShortMessage) Get_mentioned() []byte {
	return t._mentioned
}

func (t *TL_updateShortMessage) Set_media_unread(_media_unread []byte) {
	t._media_unread = _media_unread
}

func (t *TL_updateShortMessage) Get_media_unread() []byte {
	return t._media_unread
}

func (t *TL_updateShortMessage) Set_silent(_silent []byte) {
	t._silent = _silent
}

func (t *TL_updateShortMessage) Get_silent() []byte {
	return t._silent
}

func (t *TL_updateShortMessage) Set_id(_id int32) {
	t._id = _id
}

func (t *TL_updateShortMessage) Get_id() int32 {
	return t._id
}

func (t *TL_updateShortMessage) Set_user_id(_user_id int32) {
	t._user_id = _user_id
}

func (t *TL_updateShortMessage) Get_user_id() int32 {
	return t._user_id
}

func (t *TL_updateShortMessage) Set_message(_message string) {
	t._message = _message
}

func (t *TL_updateShortMessage) Get_message() string {
	return t._message
}

func (t *TL_updateShortMessage) Set_pts(_pts int32) {
	t._pts = _pts
}

func (t *TL_updateShortMessage) Get_pts() int32 {
	return t._pts
}

func (t *TL_updateShortMessage) Set_pts_count(_pts_count int32) {
	t._pts_count = _pts_count
}

func (t *TL_updateShortMessage) Get_pts_count() int32 {
	return t._pts_count
}

func (t *TL_updateShortMessage) Set_date(_date int32) {
	t._date = _date
}

func (t *TL_updateShortMessage) Get_date() int32 {
	return t._date
}

func (t *TL_updateShortMessage) Set_fwd_from(_fwd_from []byte) {
	t._fwd_from = _fwd_from
}

func (t *TL_updateShortMessage) Get_fwd_from() []byte {
	return t._fwd_from
}

func (t *TL_updateShortMessage) Set_via_bot_id(_via_bot_id []byte) {
	t._via_bot_id = _via_bot_id
}

func (t *TL_updateShortMessage) Get_via_bot_id() []byte {
	return t._via_bot_id
}

func (t *TL_updateShortMessage) Set_reply_to_msg_id(_reply_to_msg_id []byte) {
	t._reply_to_msg_id = _reply_to_msg_id
}

func (t *TL_updateShortMessage) Get_reply_to_msg_id() []byte {
	return t._reply_to_msg_id
}

func (t *TL_updateShortMessage) Set_entities(_entities []byte) {
	t._entities = _entities
}

func (t *TL_updateShortMessage) Get_entities() []byte {
	return t._entities
}

func New_TL_updateShortMessage() *TL_updateShortMessage {
	return &TL_updateShortMessage{}
}

func (t *TL_updateShortMessage) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updateShortMessage))
	ec.Bytes(t.Get_out())
	ec.Bytes(t.Get_mentioned())
	ec.Bytes(t.Get_media_unread())
	ec.Bytes(t.Get_silent())
	ec.Int(t.Get_id())
	ec.Int(t.Get_user_id())
	ec.String(t.Get_message())
	ec.Int(t.Get_pts())
	ec.Int(t.Get_pts_count())
	ec.Int(t.Get_date())
	ec.Bytes(t.Get_fwd_from())
	ec.Bytes(t.Get_via_bot_id())
	ec.Bytes(t.Get_reply_to_msg_id())
	ec.Bytes(t.Get_entities())

	return ec.GetBuffer()
}

func (t *TL_updateShortMessage) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._out = dc.Bytes(16)
	t._mentioned = dc.Bytes(16)
	t._media_unread = dc.Bytes(16)
	t._silent = dc.Bytes(16)
	t._id = dc.Int()
	t._user_id = dc.Int()
	t._message = dc.String()
	t._pts = dc.Int()
	t._pts_count = dc.Int()
	t._date = dc.Int()
	t._fwd_from = dc.Bytes(16)
	t._via_bot_id = dc.Bytes(16)
	t._reply_to_msg_id = dc.Bytes(16)
	t._entities = dc.Bytes(16)

}

// updateShortChatMessage#16812688
type TL_updateShortChatMessage struct {
	_flags           []byte
	_out             []byte
	_mentioned       []byte
	_media_unread    []byte
	_silent          []byte
	_id              int32
	_from_id         int32
	_chat_id         int32
	_message         string
	_pts             int32
	_pts_count       int32
	_date            int32
	_fwd_from        []byte
	_via_bot_id      []byte
	_reply_to_msg_id []byte
	_entities        []byte
}

func (t *TL_updateShortChatMessage) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_updateShortChatMessage) Get_flags() []byte {
	return t._flags
}

func (t *TL_updateShortChatMessage) Set_out(_out []byte) {
	t._out = _out
}

func (t *TL_updateShortChatMessage) Get_out() []byte {
	return t._out
}

func (t *TL_updateShortChatMessage) Set_mentioned(_mentioned []byte) {
	t._mentioned = _mentioned
}

func (t *TL_updateShortChatMessage) Get_mentioned() []byte {
	return t._mentioned
}

func (t *TL_updateShortChatMessage) Set_media_unread(_media_unread []byte) {
	t._media_unread = _media_unread
}

func (t *TL_updateShortChatMessage) Get_media_unread() []byte {
	return t._media_unread
}

func (t *TL_updateShortChatMessage) Set_silent(_silent []byte) {
	t._silent = _silent
}

func (t *TL_updateShortChatMessage) Get_silent() []byte {
	return t._silent
}

func (t *TL_updateShortChatMessage) Set_id(_id int32) {
	t._id = _id
}

func (t *TL_updateShortChatMessage) Get_id() int32 {
	return t._id
}

func (t *TL_updateShortChatMessage) Set_from_id(_from_id int32) {
	t._from_id = _from_id
}

func (t *TL_updateShortChatMessage) Get_from_id() int32 {
	return t._from_id
}

func (t *TL_updateShortChatMessage) Set_chat_id(_chat_id int32) {
	t._chat_id = _chat_id
}

func (t *TL_updateShortChatMessage) Get_chat_id() int32 {
	return t._chat_id
}

func (t *TL_updateShortChatMessage) Set_message(_message string) {
	t._message = _message
}

func (t *TL_updateShortChatMessage) Get_message() string {
	return t._message
}

func (t *TL_updateShortChatMessage) Set_pts(_pts int32) {
	t._pts = _pts
}

func (t *TL_updateShortChatMessage) Get_pts() int32 {
	return t._pts
}

func (t *TL_updateShortChatMessage) Set_pts_count(_pts_count int32) {
	t._pts_count = _pts_count
}

func (t *TL_updateShortChatMessage) Get_pts_count() int32 {
	return t._pts_count
}

func (t *TL_updateShortChatMessage) Set_date(_date int32) {
	t._date = _date
}

func (t *TL_updateShortChatMessage) Get_date() int32 {
	return t._date
}

func (t *TL_updateShortChatMessage) Set_fwd_from(_fwd_from []byte) {
	t._fwd_from = _fwd_from
}

func (t *TL_updateShortChatMessage) Get_fwd_from() []byte {
	return t._fwd_from
}

func (t *TL_updateShortChatMessage) Set_via_bot_id(_via_bot_id []byte) {
	t._via_bot_id = _via_bot_id
}

func (t *TL_updateShortChatMessage) Get_via_bot_id() []byte {
	return t._via_bot_id
}

func (t *TL_updateShortChatMessage) Set_reply_to_msg_id(_reply_to_msg_id []byte) {
	t._reply_to_msg_id = _reply_to_msg_id
}

func (t *TL_updateShortChatMessage) Get_reply_to_msg_id() []byte {
	return t._reply_to_msg_id
}

func (t *TL_updateShortChatMessage) Set_entities(_entities []byte) {
	t._entities = _entities
}

func (t *TL_updateShortChatMessage) Get_entities() []byte {
	return t._entities
}

func New_TL_updateShortChatMessage() *TL_updateShortChatMessage {
	return &TL_updateShortChatMessage{}
}

func (t *TL_updateShortChatMessage) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updateShortChatMessage))
	ec.Bytes(t.Get_out())
	ec.Bytes(t.Get_mentioned())
	ec.Bytes(t.Get_media_unread())
	ec.Bytes(t.Get_silent())
	ec.Int(t.Get_id())
	ec.Int(t.Get_from_id())
	ec.Int(t.Get_chat_id())
	ec.String(t.Get_message())
	ec.Int(t.Get_pts())
	ec.Int(t.Get_pts_count())
	ec.Int(t.Get_date())
	ec.Bytes(t.Get_fwd_from())
	ec.Bytes(t.Get_via_bot_id())
	ec.Bytes(t.Get_reply_to_msg_id())
	ec.Bytes(t.Get_entities())

	return ec.GetBuffer()
}

func (t *TL_updateShortChatMessage) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._out = dc.Bytes(16)
	t._mentioned = dc.Bytes(16)
	t._media_unread = dc.Bytes(16)
	t._silent = dc.Bytes(16)
	t._id = dc.Int()
	t._from_id = dc.Int()
	t._chat_id = dc.Int()
	t._message = dc.String()
	t._pts = dc.Int()
	t._pts_count = dc.Int()
	t._date = dc.Int()
	t._fwd_from = dc.Bytes(16)
	t._via_bot_id = dc.Bytes(16)
	t._reply_to_msg_id = dc.Bytes(16)
	t._entities = dc.Bytes(16)

}

// updateShort#78d4dec1
type TL_updateShort struct {
	_update []byte
	_date   int32
}

func (t *TL_updateShort) Set_update(_update []byte) {
	t._update = _update
}

func (t *TL_updateShort) Get_update() []byte {
	return t._update
}

func (t *TL_updateShort) Set_date(_date int32) {
	t._date = _date
}

func (t *TL_updateShort) Get_date() int32 {
	return t._date
}

func New_TL_updateShort() *TL_updateShort {
	return &TL_updateShort{}
}

func (t *TL_updateShort) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updateShort))
	ec.Bytes(t.Get_update())
	ec.Int(t.Get_date())

	return ec.GetBuffer()
}

func (t *TL_updateShort) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._update = dc.Bytes(16)
	t._date = dc.Int()

}

// updatesCombined#725b04c3
type TL_updatesCombined struct {
	_updates   []byte
	_users     []byte
	_chats     []byte
	_date      int32
	_seq_start int32
	_seq       int32
}

func (t *TL_updatesCombined) Set_updates(_updates []byte) {
	t._updates = _updates
}

func (t *TL_updatesCombined) Get_updates() []byte {
	return t._updates
}

func (t *TL_updatesCombined) Set_users(_users []byte) {
	t._users = _users
}

func (t *TL_updatesCombined) Get_users() []byte {
	return t._users
}

func (t *TL_updatesCombined) Set_chats(_chats []byte) {
	t._chats = _chats
}

func (t *TL_updatesCombined) Get_chats() []byte {
	return t._chats
}

func (t *TL_updatesCombined) Set_date(_date int32) {
	t._date = _date
}

func (t *TL_updatesCombined) Get_date() int32 {
	return t._date
}

func (t *TL_updatesCombined) Set_seq_start(_seq_start int32) {
	t._seq_start = _seq_start
}

func (t *TL_updatesCombined) Get_seq_start() int32 {
	return t._seq_start
}

func (t *TL_updatesCombined) Set_seq(_seq int32) {
	t._seq = _seq
}

func (t *TL_updatesCombined) Get_seq() int32 {
	return t._seq
}

func New_TL_updatesCombined() *TL_updatesCombined {
	return &TL_updatesCombined{}
}

func (t *TL_updatesCombined) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updatesCombined))
	ec.Bytes(t.Get_updates())
	ec.Bytes(t.Get_users())
	ec.Bytes(t.Get_chats())
	ec.Int(t.Get_date())
	ec.Int(t.Get_seq_start())
	ec.Int(t.Get_seq())

	return ec.GetBuffer()
}

func (t *TL_updatesCombined) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._updates = dc.Bytes(16)
	t._users = dc.Bytes(16)
	t._chats = dc.Bytes(16)
	t._date = dc.Int()
	t._seq_start = dc.Int()
	t._seq = dc.Int()

}

// updates#74ae4240
type TL_updates struct {
	_updates []byte
	_users   []byte
	_chats   []byte
	_date    int32
	_seq     int32
}

func (t *TL_updates) Set_updates(_updates []byte) {
	t._updates = _updates
}

func (t *TL_updates) Get_updates() []byte {
	return t._updates
}

func (t *TL_updates) Set_users(_users []byte) {
	t._users = _users
}

func (t *TL_updates) Get_users() []byte {
	return t._users
}

func (t *TL_updates) Set_chats(_chats []byte) {
	t._chats = _chats
}

func (t *TL_updates) Get_chats() []byte {
	return t._chats
}

func (t *TL_updates) Set_date(_date int32) {
	t._date = _date
}

func (t *TL_updates) Get_date() int32 {
	return t._date
}

func (t *TL_updates) Set_seq(_seq int32) {
	t._seq = _seq
}

func (t *TL_updates) Get_seq() int32 {
	return t._seq
}

func New_TL_updates() *TL_updates {
	return &TL_updates{}
}

func (t *TL_updates) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updates))
	ec.Bytes(t.Get_updates())
	ec.Bytes(t.Get_users())
	ec.Bytes(t.Get_chats())
	ec.Int(t.Get_date())
	ec.Int(t.Get_seq())

	return ec.GetBuffer()
}

func (t *TL_updates) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._updates = dc.Bytes(16)
	t._users = dc.Bytes(16)
	t._chats = dc.Bytes(16)
	t._date = dc.Int()
	t._seq = dc.Int()

}

// updateShortSentMessage#11f1331c
type TL_updateShortSentMessage struct {
	_flags     []byte
	_out       []byte
	_id        int32
	_pts       int32
	_pts_count int32
	_date      int32
	_media     []byte
	_entities  []byte
}

func (t *TL_updateShortSentMessage) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_updateShortSentMessage) Get_flags() []byte {
	return t._flags
}

func (t *TL_updateShortSentMessage) Set_out(_out []byte) {
	t._out = _out
}

func (t *TL_updateShortSentMessage) Get_out() []byte {
	return t._out
}

func (t *TL_updateShortSentMessage) Set_id(_id int32) {
	t._id = _id
}

func (t *TL_updateShortSentMessage) Get_id() int32 {
	return t._id
}

func (t *TL_updateShortSentMessage) Set_pts(_pts int32) {
	t._pts = _pts
}

func (t *TL_updateShortSentMessage) Get_pts() int32 {
	return t._pts
}

func (t *TL_updateShortSentMessage) Set_pts_count(_pts_count int32) {
	t._pts_count = _pts_count
}

func (t *TL_updateShortSentMessage) Get_pts_count() int32 {
	return t._pts_count
}

func (t *TL_updateShortSentMessage) Set_date(_date int32) {
	t._date = _date
}

func (t *TL_updateShortSentMessage) Get_date() int32 {
	return t._date
}

func (t *TL_updateShortSentMessage) Set_media(_media []byte) {
	t._media = _media
}

func (t *TL_updateShortSentMessage) Get_media() []byte {
	return t._media
}

func (t *TL_updateShortSentMessage) Set_entities(_entities []byte) {
	t._entities = _entities
}

func (t *TL_updateShortSentMessage) Get_entities() []byte {
	return t._entities
}

func New_TL_updateShortSentMessage() *TL_updateShortSentMessage {
	return &TL_updateShortSentMessage{}
}

func (t *TL_updateShortSentMessage) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updateShortSentMessage))
	ec.Bytes(t.Get_out())
	ec.Int(t.Get_id())
	ec.Int(t.Get_pts())
	ec.Int(t.Get_pts_count())
	ec.Int(t.Get_date())
	ec.Bytes(t.Get_media())
	ec.Bytes(t.Get_entities())

	return ec.GetBuffer()
}

func (t *TL_updateShortSentMessage) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._out = dc.Bytes(16)
	t._id = dc.Int()
	t._pts = dc.Int()
	t._pts_count = dc.Int()
	t._date = dc.Int()
	t._media = dc.Bytes(16)
	t._entities = dc.Bytes(16)

}

// photos_photos#8dca6aa5
type TL_photos_photos struct {
	_photos []byte
	_users  []byte
}

func (t *TL_photos_photos) Set_photos(_photos []byte) {
	t._photos = _photos
}

func (t *TL_photos_photos) Get_photos() []byte {
	return t._photos
}

func (t *TL_photos_photos) Set_users(_users []byte) {
	t._users = _users
}

func (t *TL_photos_photos) Get_users() []byte {
	return t._users
}

func New_TL_photos_photos() *TL_photos_photos {
	return &TL_photos_photos{}
}

func (t *TL_photos_photos) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_photos_photos))
	ec.Bytes(t.Get_photos())
	ec.Bytes(t.Get_users())

	return ec.GetBuffer()
}

func (t *TL_photos_photos) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._photos = dc.Bytes(16)
	t._users = dc.Bytes(16)

}

// photos_photosSlice#15051f54
type TL_photos_photosSlice struct {
	_count  int32
	_photos []byte
	_users  []byte
}

func (t *TL_photos_photosSlice) Set_count(_count int32) {
	t._count = _count
}

func (t *TL_photos_photosSlice) Get_count() int32 {
	return t._count
}

func (t *TL_photos_photosSlice) Set_photos(_photos []byte) {
	t._photos = _photos
}

func (t *TL_photos_photosSlice) Get_photos() []byte {
	return t._photos
}

func (t *TL_photos_photosSlice) Set_users(_users []byte) {
	t._users = _users
}

func (t *TL_photos_photosSlice) Get_users() []byte {
	return t._users
}

func New_TL_photos_photosSlice() *TL_photos_photosSlice {
	return &TL_photos_photosSlice{}
}

func (t *TL_photos_photosSlice) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_photos_photosSlice))
	ec.Int(t.Get_count())
	ec.Bytes(t.Get_photos())
	ec.Bytes(t.Get_users())

	return ec.GetBuffer()
}

func (t *TL_photos_photosSlice) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._count = dc.Int()
	t._photos = dc.Bytes(16)
	t._users = dc.Bytes(16)

}

// photos_photo#20212ca8
type TL_photos_photo struct {
	_photo []byte
	_users []byte
}

func (t *TL_photos_photo) Set_photo(_photo []byte) {
	t._photo = _photo
}

func (t *TL_photos_photo) Get_photo() []byte {
	return t._photo
}

func (t *TL_photos_photo) Set_users(_users []byte) {
	t._users = _users
}

func (t *TL_photos_photo) Get_users() []byte {
	return t._users
}

func New_TL_photos_photo() *TL_photos_photo {
	return &TL_photos_photo{}
}

func (t *TL_photos_photo) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_photos_photo))
	ec.Bytes(t.Get_photo())
	ec.Bytes(t.Get_users())

	return ec.GetBuffer()
}

func (t *TL_photos_photo) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._photo = dc.Bytes(16)
	t._users = dc.Bytes(16)

}

// upload_file#096a18d5
type TL_upload_file struct {
	_type  []byte
	_mtime int32
	_bytes []byte
}

func (t *TL_upload_file) Set_type(_type []byte) {
	t._type = _type
}

func (t *TL_upload_file) Get_type() []byte {
	return t._type
}

func (t *TL_upload_file) Set_mtime(_mtime int32) {
	t._mtime = _mtime
}

func (t *TL_upload_file) Get_mtime() int32 {
	return t._mtime
}

func (t *TL_upload_file) Set_bytes(_bytes []byte) {
	t._bytes = _bytes
}

func (t *TL_upload_file) Get_bytes() []byte {
	return t._bytes
}

func New_TL_upload_file() *TL_upload_file {
	return &TL_upload_file{}
}

func (t *TL_upload_file) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_upload_file))
	ec.Bytes(t.Get_type())
	ec.Int(t.Get_mtime())
	ec.Bytes(t.Get_bytes())

	return ec.GetBuffer()
}

func (t *TL_upload_file) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._type = dc.Bytes(16)
	t._mtime = dc.Int()
	t._bytes = dc.Bytes(16)

}

// upload_fileCdnRedirect#ea52fe5a
type TL_upload_fileCdnRedirect struct {
	_dc_id           int32
	_file_token      []byte
	_encryption_key  []byte
	_encryption_iv   []byte
	_cdn_file_hashes []byte
}

func (t *TL_upload_fileCdnRedirect) Set_dc_id(_dc_id int32) {
	t._dc_id = _dc_id
}

func (t *TL_upload_fileCdnRedirect) Get_dc_id() int32 {
	return t._dc_id
}

func (t *TL_upload_fileCdnRedirect) Set_file_token(_file_token []byte) {
	t._file_token = _file_token
}

func (t *TL_upload_fileCdnRedirect) Get_file_token() []byte {
	return t._file_token
}

func (t *TL_upload_fileCdnRedirect) Set_encryption_key(_encryption_key []byte) {
	t._encryption_key = _encryption_key
}

func (t *TL_upload_fileCdnRedirect) Get_encryption_key() []byte {
	return t._encryption_key
}

func (t *TL_upload_fileCdnRedirect) Set_encryption_iv(_encryption_iv []byte) {
	t._encryption_iv = _encryption_iv
}

func (t *TL_upload_fileCdnRedirect) Get_encryption_iv() []byte {
	return t._encryption_iv
}

func (t *TL_upload_fileCdnRedirect) Set_cdn_file_hashes(_cdn_file_hashes []byte) {
	t._cdn_file_hashes = _cdn_file_hashes
}

func (t *TL_upload_fileCdnRedirect) Get_cdn_file_hashes() []byte {
	return t._cdn_file_hashes
}

func New_TL_upload_fileCdnRedirect() *TL_upload_fileCdnRedirect {
	return &TL_upload_fileCdnRedirect{}
}

func (t *TL_upload_fileCdnRedirect) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_upload_fileCdnRedirect))
	ec.Int(t.Get_dc_id())
	ec.Bytes(t.Get_file_token())
	ec.Bytes(t.Get_encryption_key())
	ec.Bytes(t.Get_encryption_iv())
	ec.Bytes(t.Get_cdn_file_hashes())

	return ec.GetBuffer()
}

func (t *TL_upload_fileCdnRedirect) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._dc_id = dc.Int()
	t._file_token = dc.Bytes(16)
	t._encryption_key = dc.Bytes(16)
	t._encryption_iv = dc.Bytes(16)
	t._cdn_file_hashes = dc.Bytes(16)

}

// dcOption#05d8c6cc
type TL_dcOption struct {
	_flags      []byte
	_ipv6       []byte
	_media_only []byte
	_tcpo_only  []byte
	_cdn        []byte
	_static     []byte
	_id         int32
	_ip_address string
	_port       int32
}

func (t *TL_dcOption) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_dcOption) Get_flags() []byte {
	return t._flags
}

func (t *TL_dcOption) Set_ipv6(_ipv6 []byte) {
	t._ipv6 = _ipv6
}

func (t *TL_dcOption) Get_ipv6() []byte {
	return t._ipv6
}

func (t *TL_dcOption) Set_media_only(_media_only []byte) {
	t._media_only = _media_only
}

func (t *TL_dcOption) Get_media_only() []byte {
	return t._media_only
}

func (t *TL_dcOption) Set_tcpo_only(_tcpo_only []byte) {
	t._tcpo_only = _tcpo_only
}

func (t *TL_dcOption) Get_tcpo_only() []byte {
	return t._tcpo_only
}

func (t *TL_dcOption) Set_cdn(_cdn []byte) {
	t._cdn = _cdn
}

func (t *TL_dcOption) Get_cdn() []byte {
	return t._cdn
}

func (t *TL_dcOption) Set_static(_static []byte) {
	t._static = _static
}

func (t *TL_dcOption) Get_static() []byte {
	return t._static
}

func (t *TL_dcOption) Set_id(_id int32) {
	t._id = _id
}

func (t *TL_dcOption) Get_id() int32 {
	return t._id
}

func (t *TL_dcOption) Set_ip_address(_ip_address string) {
	t._ip_address = _ip_address
}

func (t *TL_dcOption) Get_ip_address() string {
	return t._ip_address
}

func (t *TL_dcOption) Set_port(_port int32) {
	t._port = _port
}

func (t *TL_dcOption) Get_port() int32 {
	return t._port
}

func New_TL_dcOption() *TL_dcOption {
	return &TL_dcOption{}
}

func (t *TL_dcOption) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_dcOption))
	ec.Bytes(t.Get_ipv6())
	ec.Bytes(t.Get_media_only())
	ec.Bytes(t.Get_tcpo_only())
	ec.Bytes(t.Get_cdn())
	ec.Bytes(t.Get_static())
	ec.Int(t.Get_id())
	ec.String(t.Get_ip_address())
	ec.Int(t.Get_port())

	return ec.GetBuffer()
}

func (t *TL_dcOption) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._ipv6 = dc.Bytes(16)
	t._media_only = dc.Bytes(16)
	t._tcpo_only = dc.Bytes(16)
	t._cdn = dc.Bytes(16)
	t._static = dc.Bytes(16)
	t._id = dc.Int()
	t._ip_address = dc.String()
	t._port = dc.Int()

}

// config#9c840964
type TL_config struct {
	_flags                      []byte
	_phonecalls_enabled         []byte
	_default_p2p_contacts       []byte
	_date                       int32
	_expires                    int32
	_test_mode                  bool
	_this_dc                    int32
	_dc_options                 []byte
	_chat_size_max              int32
	_megagroup_size_max         int32
	_forwarded_count_max        int32
	_online_update_period_ms    int32
	_offline_blur_timeout_ms    int32
	_offline_idle_timeout_ms    int32
	_online_cloud_timeout_ms    int32
	_notify_cloud_delay_ms      int32
	_notify_default_delay_ms    int32
	_chat_big_size              int32
	_push_chat_period_ms        int32
	_push_chat_limit            int32
	_saved_gifs_limit           int32
	_edit_time_limit            int32
	_rating_e_decay             int32
	_stickers_recent_limit      int32
	_stickers_faved_limit       int32
	_channels_read_media_period int32
	_tmp_sessions               []byte
	_pinned_dialogs_count_max   int32
	_call_receive_timeout_ms    int32
	_call_ring_timeout_ms       int32
	_call_connect_timeout_ms    int32
	_call_packet_timeout_ms     int32
	_me_url_prefix              string
	_suggested_lang_code        []byte
	_lang_pack_version          []byte
	_disabled_features          []byte
}

func (t *TL_config) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_config) Get_flags() []byte {
	return t._flags
}

func (t *TL_config) Set_phonecalls_enabled(_phonecalls_enabled []byte) {
	t._phonecalls_enabled = _phonecalls_enabled
}

func (t *TL_config) Get_phonecalls_enabled() []byte {
	return t._phonecalls_enabled
}

func (t *TL_config) Set_default_p2p_contacts(_default_p2p_contacts []byte) {
	t._default_p2p_contacts = _default_p2p_contacts
}

func (t *TL_config) Get_default_p2p_contacts() []byte {
	return t._default_p2p_contacts
}

func (t *TL_config) Set_date(_date int32) {
	t._date = _date
}

func (t *TL_config) Get_date() int32 {
	return t._date
}

func (t *TL_config) Set_expires(_expires int32) {
	t._expires = _expires
}

func (t *TL_config) Get_expires() int32 {
	return t._expires
}

func (t *TL_config) Set_test_mode(_test_mode bool) {
	t._test_mode = _test_mode
}

func (t *TL_config) Get_test_mode() bool {
	return t._test_mode
}

func (t *TL_config) Set_this_dc(_this_dc int32) {
	t._this_dc = _this_dc
}

func (t *TL_config) Get_this_dc() int32 {
	return t._this_dc
}

func (t *TL_config) Set_dc_options(_dc_options []byte) {
	t._dc_options = _dc_options
}

func (t *TL_config) Get_dc_options() []byte {
	return t._dc_options
}

func (t *TL_config) Set_chat_size_max(_chat_size_max int32) {
	t._chat_size_max = _chat_size_max
}

func (t *TL_config) Get_chat_size_max() int32 {
	return t._chat_size_max
}

func (t *TL_config) Set_megagroup_size_max(_megagroup_size_max int32) {
	t._megagroup_size_max = _megagroup_size_max
}

func (t *TL_config) Get_megagroup_size_max() int32 {
	return t._megagroup_size_max
}

func (t *TL_config) Set_forwarded_count_max(_forwarded_count_max int32) {
	t._forwarded_count_max = _forwarded_count_max
}

func (t *TL_config) Get_forwarded_count_max() int32 {
	return t._forwarded_count_max
}

func (t *TL_config) Set_online_update_period_ms(_online_update_period_ms int32) {
	t._online_update_period_ms = _online_update_period_ms
}

func (t *TL_config) Get_online_update_period_ms() int32 {
	return t._online_update_period_ms
}

func (t *TL_config) Set_offline_blur_timeout_ms(_offline_blur_timeout_ms int32) {
	t._offline_blur_timeout_ms = _offline_blur_timeout_ms
}

func (t *TL_config) Get_offline_blur_timeout_ms() int32 {
	return t._offline_blur_timeout_ms
}

func (t *TL_config) Set_offline_idle_timeout_ms(_offline_idle_timeout_ms int32) {
	t._offline_idle_timeout_ms = _offline_idle_timeout_ms
}

func (t *TL_config) Get_offline_idle_timeout_ms() int32 {
	return t._offline_idle_timeout_ms
}

func (t *TL_config) Set_online_cloud_timeout_ms(_online_cloud_timeout_ms int32) {
	t._online_cloud_timeout_ms = _online_cloud_timeout_ms
}

func (t *TL_config) Get_online_cloud_timeout_ms() int32 {
	return t._online_cloud_timeout_ms
}

func (t *TL_config) Set_notify_cloud_delay_ms(_notify_cloud_delay_ms int32) {
	t._notify_cloud_delay_ms = _notify_cloud_delay_ms
}

func (t *TL_config) Get_notify_cloud_delay_ms() int32 {
	return t._notify_cloud_delay_ms
}

func (t *TL_config) Set_notify_default_delay_ms(_notify_default_delay_ms int32) {
	t._notify_default_delay_ms = _notify_default_delay_ms
}

func (t *TL_config) Get_notify_default_delay_ms() int32 {
	return t._notify_default_delay_ms
}

func (t *TL_config) Set_chat_big_size(_chat_big_size int32) {
	t._chat_big_size = _chat_big_size
}

func (t *TL_config) Get_chat_big_size() int32 {
	return t._chat_big_size
}

func (t *TL_config) Set_push_chat_period_ms(_push_chat_period_ms int32) {
	t._push_chat_period_ms = _push_chat_period_ms
}

func (t *TL_config) Get_push_chat_period_ms() int32 {
	return t._push_chat_period_ms
}

func (t *TL_config) Set_push_chat_limit(_push_chat_limit int32) {
	t._push_chat_limit = _push_chat_limit
}

func (t *TL_config) Get_push_chat_limit() int32 {
	return t._push_chat_limit
}

func (t *TL_config) Set_saved_gifs_limit(_saved_gifs_limit int32) {
	t._saved_gifs_limit = _saved_gifs_limit
}

func (t *TL_config) Get_saved_gifs_limit() int32 {
	return t._saved_gifs_limit
}

func (t *TL_config) Set_edit_time_limit(_edit_time_limit int32) {
	t._edit_time_limit = _edit_time_limit
}

func (t *TL_config) Get_edit_time_limit() int32 {
	return t._edit_time_limit
}

func (t *TL_config) Set_rating_e_decay(_rating_e_decay int32) {
	t._rating_e_decay = _rating_e_decay
}

func (t *TL_config) Get_rating_e_decay() int32 {
	return t._rating_e_decay
}

func (t *TL_config) Set_stickers_recent_limit(_stickers_recent_limit int32) {
	t._stickers_recent_limit = _stickers_recent_limit
}

func (t *TL_config) Get_stickers_recent_limit() int32 {
	return t._stickers_recent_limit
}

func (t *TL_config) Set_stickers_faved_limit(_stickers_faved_limit int32) {
	t._stickers_faved_limit = _stickers_faved_limit
}

func (t *TL_config) Get_stickers_faved_limit() int32 {
	return t._stickers_faved_limit
}

func (t *TL_config) Set_channels_read_media_period(_channels_read_media_period int32) {
	t._channels_read_media_period = _channels_read_media_period
}

func (t *TL_config) Get_channels_read_media_period() int32 {
	return t._channels_read_media_period
}

func (t *TL_config) Set_tmp_sessions(_tmp_sessions []byte) {
	t._tmp_sessions = _tmp_sessions
}

func (t *TL_config) Get_tmp_sessions() []byte {
	return t._tmp_sessions
}

func (t *TL_config) Set_pinned_dialogs_count_max(_pinned_dialogs_count_max int32) {
	t._pinned_dialogs_count_max = _pinned_dialogs_count_max
}

func (t *TL_config) Get_pinned_dialogs_count_max() int32 {
	return t._pinned_dialogs_count_max
}

func (t *TL_config) Set_call_receive_timeout_ms(_call_receive_timeout_ms int32) {
	t._call_receive_timeout_ms = _call_receive_timeout_ms
}

func (t *TL_config) Get_call_receive_timeout_ms() int32 {
	return t._call_receive_timeout_ms
}

func (t *TL_config) Set_call_ring_timeout_ms(_call_ring_timeout_ms int32) {
	t._call_ring_timeout_ms = _call_ring_timeout_ms
}

func (t *TL_config) Get_call_ring_timeout_ms() int32 {
	return t._call_ring_timeout_ms
}

func (t *TL_config) Set_call_connect_timeout_ms(_call_connect_timeout_ms int32) {
	t._call_connect_timeout_ms = _call_connect_timeout_ms
}

func (t *TL_config) Get_call_connect_timeout_ms() int32 {
	return t._call_connect_timeout_ms
}

func (t *TL_config) Set_call_packet_timeout_ms(_call_packet_timeout_ms int32) {
	t._call_packet_timeout_ms = _call_packet_timeout_ms
}

func (t *TL_config) Get_call_packet_timeout_ms() int32 {
	return t._call_packet_timeout_ms
}

func (t *TL_config) Set_me_url_prefix(_me_url_prefix string) {
	t._me_url_prefix = _me_url_prefix
}

func (t *TL_config) Get_me_url_prefix() string {
	return t._me_url_prefix
}

func (t *TL_config) Set_suggested_lang_code(_suggested_lang_code []byte) {
	t._suggested_lang_code = _suggested_lang_code
}

func (t *TL_config) Get_suggested_lang_code() []byte {
	return t._suggested_lang_code
}

func (t *TL_config) Set_lang_pack_version(_lang_pack_version []byte) {
	t._lang_pack_version = _lang_pack_version
}

func (t *TL_config) Get_lang_pack_version() []byte {
	return t._lang_pack_version
}

func (t *TL_config) Set_disabled_features(_disabled_features []byte) {
	t._disabled_features = _disabled_features
}

func (t *TL_config) Get_disabled_features() []byte {
	return t._disabled_features
}

func New_TL_config() *TL_config {
	return &TL_config{}
}

func (t *TL_config) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_config))
	ec.Bytes(t.Get_phonecalls_enabled())
	ec.Bytes(t.Get_default_p2p_contacts())
	ec.Int(t.Get_date())
	ec.Int(t.Get_expires())
	ec.Bool(t.Get_test_mode())
	ec.Int(t.Get_this_dc())
	ec.Bytes(t.Get_dc_options())
	ec.Int(t.Get_chat_size_max())
	ec.Int(t.Get_megagroup_size_max())
	ec.Int(t.Get_forwarded_count_max())
	ec.Int(t.Get_online_update_period_ms())
	ec.Int(t.Get_offline_blur_timeout_ms())
	ec.Int(t.Get_offline_idle_timeout_ms())
	ec.Int(t.Get_online_cloud_timeout_ms())
	ec.Int(t.Get_notify_cloud_delay_ms())
	ec.Int(t.Get_notify_default_delay_ms())
	ec.Int(t.Get_chat_big_size())
	ec.Int(t.Get_push_chat_period_ms())
	ec.Int(t.Get_push_chat_limit())
	ec.Int(t.Get_saved_gifs_limit())
	ec.Int(t.Get_edit_time_limit())
	ec.Int(t.Get_rating_e_decay())
	ec.Int(t.Get_stickers_recent_limit())
	ec.Int(t.Get_stickers_faved_limit())
	ec.Int(t.Get_channels_read_media_period())
	ec.Bytes(t.Get_tmp_sessions())
	ec.Int(t.Get_pinned_dialogs_count_max())
	ec.Int(t.Get_call_receive_timeout_ms())
	ec.Int(t.Get_call_ring_timeout_ms())
	ec.Int(t.Get_call_connect_timeout_ms())
	ec.Int(t.Get_call_packet_timeout_ms())
	ec.String(t.Get_me_url_prefix())
	ec.Bytes(t.Get_suggested_lang_code())
	ec.Bytes(t.Get_lang_pack_version())
	ec.Bytes(t.Get_disabled_features())

	return ec.GetBuffer()
}

func (t *TL_config) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._phonecalls_enabled = dc.Bytes(16)
	t._default_p2p_contacts = dc.Bytes(16)
	t._date = dc.Int()
	t._expires = dc.Int()
	t._test_mode = dc.Bool()
	t._this_dc = dc.Int()
	t._dc_options = dc.Bytes(16)
	t._chat_size_max = dc.Int()
	t._megagroup_size_max = dc.Int()
	t._forwarded_count_max = dc.Int()
	t._online_update_period_ms = dc.Int()
	t._offline_blur_timeout_ms = dc.Int()
	t._offline_idle_timeout_ms = dc.Int()
	t._online_cloud_timeout_ms = dc.Int()
	t._notify_cloud_delay_ms = dc.Int()
	t._notify_default_delay_ms = dc.Int()
	t._chat_big_size = dc.Int()
	t._push_chat_period_ms = dc.Int()
	t._push_chat_limit = dc.Int()
	t._saved_gifs_limit = dc.Int()
	t._edit_time_limit = dc.Int()
	t._rating_e_decay = dc.Int()
	t._stickers_recent_limit = dc.Int()
	t._stickers_faved_limit = dc.Int()
	t._channels_read_media_period = dc.Int()
	t._tmp_sessions = dc.Bytes(16)
	t._pinned_dialogs_count_max = dc.Int()
	t._call_receive_timeout_ms = dc.Int()
	t._call_ring_timeout_ms = dc.Int()
	t._call_connect_timeout_ms = dc.Int()
	t._call_packet_timeout_ms = dc.Int()
	t._me_url_prefix = dc.String()
	t._suggested_lang_code = dc.Bytes(16)
	t._lang_pack_version = dc.Bytes(16)
	t._disabled_features = dc.Bytes(16)

}

// nearestDc#8e1a1775
type TL_nearestDc struct {
	_country    string
	_this_dc    int32
	_nearest_dc int32
}

func (t *TL_nearestDc) Set_country(_country string) {
	t._country = _country
}

func (t *TL_nearestDc) Get_country() string {
	return t._country
}

func (t *TL_nearestDc) Set_this_dc(_this_dc int32) {
	t._this_dc = _this_dc
}

func (t *TL_nearestDc) Get_this_dc() int32 {
	return t._this_dc
}

func (t *TL_nearestDc) Set_nearest_dc(_nearest_dc int32) {
	t._nearest_dc = _nearest_dc
}

func (t *TL_nearestDc) Get_nearest_dc() int32 {
	return t._nearest_dc
}

func New_TL_nearestDc() *TL_nearestDc {
	return &TL_nearestDc{}
}

func (t *TL_nearestDc) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_nearestDc))
	ec.String(t.Get_country())
	ec.Int(t.Get_this_dc())
	ec.Int(t.Get_nearest_dc())

	return ec.GetBuffer()
}

func (t *TL_nearestDc) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._country = dc.String()
	t._this_dc = dc.Int()
	t._nearest_dc = dc.Int()

}

// help_appUpdate#8987f311
type TL_help_appUpdate struct {
	_id       int32
	_critical bool
	_url      string
	_text     string
}

func (t *TL_help_appUpdate) Set_id(_id int32) {
	t._id = _id
}

func (t *TL_help_appUpdate) Get_id() int32 {
	return t._id
}

func (t *TL_help_appUpdate) Set_critical(_critical bool) {
	t._critical = _critical
}

func (t *TL_help_appUpdate) Get_critical() bool {
	return t._critical
}

func (t *TL_help_appUpdate) Set_url(_url string) {
	t._url = _url
}

func (t *TL_help_appUpdate) Get_url() string {
	return t._url
}

func (t *TL_help_appUpdate) Set_text(_text string) {
	t._text = _text
}

func (t *TL_help_appUpdate) Get_text() string {
	return t._text
}

func New_TL_help_appUpdate() *TL_help_appUpdate {
	return &TL_help_appUpdate{}
}

func (t *TL_help_appUpdate) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_help_appUpdate))
	ec.Int(t.Get_id())
	ec.Bool(t.Get_critical())
	ec.String(t.Get_url())
	ec.String(t.Get_text())

	return ec.GetBuffer()
}

func (t *TL_help_appUpdate) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Int()
	t._critical = dc.Bool()
	t._url = dc.String()
	t._text = dc.String()

}

// help_noAppUpdate#c45a6536
type TL_help_noAppUpdate struct {
}

func New_TL_help_noAppUpdate() *TL_help_noAppUpdate {
	return &TL_help_noAppUpdate{}
}

func (t *TL_help_noAppUpdate) Encode() []byte {
	return nil
}

func (t *TL_help_noAppUpdate) Decode(b []byte) {

}

// help_inviteText#18cb9f78
type TL_help_inviteText struct {
	_message string
}

func (t *TL_help_inviteText) Set_message(_message string) {
	t._message = _message
}

func (t *TL_help_inviteText) Get_message() string {
	return t._message
}

func New_TL_help_inviteText() *TL_help_inviteText {
	return &TL_help_inviteText{}
}

func (t *TL_help_inviteText) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_help_inviteText))
	ec.String(t.Get_message())

	return ec.GetBuffer()
}

func (t *TL_help_inviteText) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._message = dc.String()

}

// encryptedChatEmpty#ab7ec0a0
type TL_encryptedChatEmpty struct {
	_id int32
}

func (t *TL_encryptedChatEmpty) Set_id(_id int32) {
	t._id = _id
}

func (t *TL_encryptedChatEmpty) Get_id() int32 {
	return t._id
}

func New_TL_encryptedChatEmpty() *TL_encryptedChatEmpty {
	return &TL_encryptedChatEmpty{}
}

func (t *TL_encryptedChatEmpty) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_encryptedChatEmpty))
	ec.Int(t.Get_id())

	return ec.GetBuffer()
}

func (t *TL_encryptedChatEmpty) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Int()

}

// encryptedChatWaiting#3bf703dc
type TL_encryptedChatWaiting struct {
	_id             int32
	_access_hash    int64
	_date           int32
	_admin_id       int32
	_participant_id int32
}

func (t *TL_encryptedChatWaiting) Set_id(_id int32) {
	t._id = _id
}

func (t *TL_encryptedChatWaiting) Get_id() int32 {
	return t._id
}

func (t *TL_encryptedChatWaiting) Set_access_hash(_access_hash int64) {
	t._access_hash = _access_hash
}

func (t *TL_encryptedChatWaiting) Get_access_hash() int64 {
	return t._access_hash
}

func (t *TL_encryptedChatWaiting) Set_date(_date int32) {
	t._date = _date
}

func (t *TL_encryptedChatWaiting) Get_date() int32 {
	return t._date
}

func (t *TL_encryptedChatWaiting) Set_admin_id(_admin_id int32) {
	t._admin_id = _admin_id
}

func (t *TL_encryptedChatWaiting) Get_admin_id() int32 {
	return t._admin_id
}

func (t *TL_encryptedChatWaiting) Set_participant_id(_participant_id int32) {
	t._participant_id = _participant_id
}

func (t *TL_encryptedChatWaiting) Get_participant_id() int32 {
	return t._participant_id
}

func New_TL_encryptedChatWaiting() *TL_encryptedChatWaiting {
	return &TL_encryptedChatWaiting{}
}

func (t *TL_encryptedChatWaiting) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_encryptedChatWaiting))
	ec.Int(t.Get_id())
	ec.Long(t.Get_access_hash())
	ec.Int(t.Get_date())
	ec.Int(t.Get_admin_id())
	ec.Int(t.Get_participant_id())

	return ec.GetBuffer()
}

func (t *TL_encryptedChatWaiting) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Int()
	t._access_hash = dc.Long()
	t._date = dc.Int()
	t._admin_id = dc.Int()
	t._participant_id = dc.Int()

}

// encryptedChatRequested#c878527e
type TL_encryptedChatRequested struct {
	_id             int32
	_access_hash    int64
	_date           int32
	_admin_id       int32
	_participant_id int32
	_g_a            []byte
}

func (t *TL_encryptedChatRequested) Set_id(_id int32) {
	t._id = _id
}

func (t *TL_encryptedChatRequested) Get_id() int32 {
	return t._id
}

func (t *TL_encryptedChatRequested) Set_access_hash(_access_hash int64) {
	t._access_hash = _access_hash
}

func (t *TL_encryptedChatRequested) Get_access_hash() int64 {
	return t._access_hash
}

func (t *TL_encryptedChatRequested) Set_date(_date int32) {
	t._date = _date
}

func (t *TL_encryptedChatRequested) Get_date() int32 {
	return t._date
}

func (t *TL_encryptedChatRequested) Set_admin_id(_admin_id int32) {
	t._admin_id = _admin_id
}

func (t *TL_encryptedChatRequested) Get_admin_id() int32 {
	return t._admin_id
}

func (t *TL_encryptedChatRequested) Set_participant_id(_participant_id int32) {
	t._participant_id = _participant_id
}

func (t *TL_encryptedChatRequested) Get_participant_id() int32 {
	return t._participant_id
}

func (t *TL_encryptedChatRequested) Set_g_a(_g_a []byte) {
	t._g_a = _g_a
}

func (t *TL_encryptedChatRequested) Get_g_a() []byte {
	return t._g_a
}

func New_TL_encryptedChatRequested() *TL_encryptedChatRequested {
	return &TL_encryptedChatRequested{}
}

func (t *TL_encryptedChatRequested) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_encryptedChatRequested))
	ec.Int(t.Get_id())
	ec.Long(t.Get_access_hash())
	ec.Int(t.Get_date())
	ec.Int(t.Get_admin_id())
	ec.Int(t.Get_participant_id())
	ec.Bytes(t.Get_g_a())

	return ec.GetBuffer()
}

func (t *TL_encryptedChatRequested) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Int()
	t._access_hash = dc.Long()
	t._date = dc.Int()
	t._admin_id = dc.Int()
	t._participant_id = dc.Int()
	t._g_a = dc.Bytes(16)

}

// encryptedChat#fa56ce36
type TL_encryptedChat struct {
	_id              int32
	_access_hash     int64
	_date            int32
	_admin_id        int32
	_participant_id  int32
	_g_a_or_b        []byte
	_key_fingerprint int64
}

func (t *TL_encryptedChat) Set_id(_id int32) {
	t._id = _id
}

func (t *TL_encryptedChat) Get_id() int32 {
	return t._id
}

func (t *TL_encryptedChat) Set_access_hash(_access_hash int64) {
	t._access_hash = _access_hash
}

func (t *TL_encryptedChat) Get_access_hash() int64 {
	return t._access_hash
}

func (t *TL_encryptedChat) Set_date(_date int32) {
	t._date = _date
}

func (t *TL_encryptedChat) Get_date() int32 {
	return t._date
}

func (t *TL_encryptedChat) Set_admin_id(_admin_id int32) {
	t._admin_id = _admin_id
}

func (t *TL_encryptedChat) Get_admin_id() int32 {
	return t._admin_id
}

func (t *TL_encryptedChat) Set_participant_id(_participant_id int32) {
	t._participant_id = _participant_id
}

func (t *TL_encryptedChat) Get_participant_id() int32 {
	return t._participant_id
}

func (t *TL_encryptedChat) Set_g_a_or_b(_g_a_or_b []byte) {
	t._g_a_or_b = _g_a_or_b
}

func (t *TL_encryptedChat) Get_g_a_or_b() []byte {
	return t._g_a_or_b
}

func (t *TL_encryptedChat) Set_key_fingerprint(_key_fingerprint int64) {
	t._key_fingerprint = _key_fingerprint
}

func (t *TL_encryptedChat) Get_key_fingerprint() int64 {
	return t._key_fingerprint
}

func New_TL_encryptedChat() *TL_encryptedChat {
	return &TL_encryptedChat{}
}

func (t *TL_encryptedChat) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_encryptedChat))
	ec.Int(t.Get_id())
	ec.Long(t.Get_access_hash())
	ec.Int(t.Get_date())
	ec.Int(t.Get_admin_id())
	ec.Int(t.Get_participant_id())
	ec.Bytes(t.Get_g_a_or_b())
	ec.Long(t.Get_key_fingerprint())

	return ec.GetBuffer()
}

func (t *TL_encryptedChat) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Int()
	t._access_hash = dc.Long()
	t._date = dc.Int()
	t._admin_id = dc.Int()
	t._participant_id = dc.Int()
	t._g_a_or_b = dc.Bytes(16)
	t._key_fingerprint = dc.Long()

}

// encryptedChatDiscarded#13d6dd27
type TL_encryptedChatDiscarded struct {
	_id int32
}

func (t *TL_encryptedChatDiscarded) Set_id(_id int32) {
	t._id = _id
}

func (t *TL_encryptedChatDiscarded) Get_id() int32 {
	return t._id
}

func New_TL_encryptedChatDiscarded() *TL_encryptedChatDiscarded {
	return &TL_encryptedChatDiscarded{}
}

func (t *TL_encryptedChatDiscarded) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_encryptedChatDiscarded))
	ec.Int(t.Get_id())

	return ec.GetBuffer()
}

func (t *TL_encryptedChatDiscarded) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Int()

}

// inputEncryptedChat#f141b5e1
type TL_inputEncryptedChat struct {
	_chat_id     int32
	_access_hash int64
}

func (t *TL_inputEncryptedChat) Set_chat_id(_chat_id int32) {
	t._chat_id = _chat_id
}

func (t *TL_inputEncryptedChat) Get_chat_id() int32 {
	return t._chat_id
}

func (t *TL_inputEncryptedChat) Set_access_hash(_access_hash int64) {
	t._access_hash = _access_hash
}

func (t *TL_inputEncryptedChat) Get_access_hash() int64 {
	return t._access_hash
}

func New_TL_inputEncryptedChat() *TL_inputEncryptedChat {
	return &TL_inputEncryptedChat{}
}

func (t *TL_inputEncryptedChat) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputEncryptedChat))
	ec.Int(t.Get_chat_id())
	ec.Long(t.Get_access_hash())

	return ec.GetBuffer()
}

func (t *TL_inputEncryptedChat) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._chat_id = dc.Int()
	t._access_hash = dc.Long()

}

// encryptedFileEmpty#c21f497e
type TL_encryptedFileEmpty struct {
}

func New_TL_encryptedFileEmpty() *TL_encryptedFileEmpty {
	return &TL_encryptedFileEmpty{}
}

func (t *TL_encryptedFileEmpty) Encode() []byte {
	return nil
}

func (t *TL_encryptedFileEmpty) Decode(b []byte) {

}

// encryptedFile#4a70994c
type TL_encryptedFile struct {
	_id              int64
	_access_hash     int64
	_size            int32
	_dc_id           int32
	_key_fingerprint int32
}

func (t *TL_encryptedFile) Set_id(_id int64) {
	t._id = _id
}

func (t *TL_encryptedFile) Get_id() int64 {
	return t._id
}

func (t *TL_encryptedFile) Set_access_hash(_access_hash int64) {
	t._access_hash = _access_hash
}

func (t *TL_encryptedFile) Get_access_hash() int64 {
	return t._access_hash
}

func (t *TL_encryptedFile) Set_size(_size int32) {
	t._size = _size
}

func (t *TL_encryptedFile) Get_size() int32 {
	return t._size
}

func (t *TL_encryptedFile) Set_dc_id(_dc_id int32) {
	t._dc_id = _dc_id
}

func (t *TL_encryptedFile) Get_dc_id() int32 {
	return t._dc_id
}

func (t *TL_encryptedFile) Set_key_fingerprint(_key_fingerprint int32) {
	t._key_fingerprint = _key_fingerprint
}

func (t *TL_encryptedFile) Get_key_fingerprint() int32 {
	return t._key_fingerprint
}

func New_TL_encryptedFile() *TL_encryptedFile {
	return &TL_encryptedFile{}
}

func (t *TL_encryptedFile) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_encryptedFile))
	ec.Long(t.Get_id())
	ec.Long(t.Get_access_hash())
	ec.Int(t.Get_size())
	ec.Int(t.Get_dc_id())
	ec.Int(t.Get_key_fingerprint())

	return ec.GetBuffer()
}

func (t *TL_encryptedFile) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Long()
	t._access_hash = dc.Long()
	t._size = dc.Int()
	t._dc_id = dc.Int()
	t._key_fingerprint = dc.Int()

}

// inputEncryptedFileEmpty#1837c364
type TL_inputEncryptedFileEmpty struct {
}

func New_TL_inputEncryptedFileEmpty() *TL_inputEncryptedFileEmpty {
	return &TL_inputEncryptedFileEmpty{}
}

func (t *TL_inputEncryptedFileEmpty) Encode() []byte {
	return nil
}

func (t *TL_inputEncryptedFileEmpty) Decode(b []byte) {

}

// inputEncryptedFileUploaded#64bd0306
type TL_inputEncryptedFileUploaded struct {
	_id              int64
	_parts           int32
	_md5_checksum    string
	_key_fingerprint int32
}

func (t *TL_inputEncryptedFileUploaded) Set_id(_id int64) {
	t._id = _id
}

func (t *TL_inputEncryptedFileUploaded) Get_id() int64 {
	return t._id
}

func (t *TL_inputEncryptedFileUploaded) Set_parts(_parts int32) {
	t._parts = _parts
}

func (t *TL_inputEncryptedFileUploaded) Get_parts() int32 {
	return t._parts
}

func (t *TL_inputEncryptedFileUploaded) Set_md5_checksum(_md5_checksum string) {
	t._md5_checksum = _md5_checksum
}

func (t *TL_inputEncryptedFileUploaded) Get_md5_checksum() string {
	return t._md5_checksum
}

func (t *TL_inputEncryptedFileUploaded) Set_key_fingerprint(_key_fingerprint int32) {
	t._key_fingerprint = _key_fingerprint
}

func (t *TL_inputEncryptedFileUploaded) Get_key_fingerprint() int32 {
	return t._key_fingerprint
}

func New_TL_inputEncryptedFileUploaded() *TL_inputEncryptedFileUploaded {
	return &TL_inputEncryptedFileUploaded{}
}

func (t *TL_inputEncryptedFileUploaded) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputEncryptedFileUploaded))
	ec.Long(t.Get_id())
	ec.Int(t.Get_parts())
	ec.String(t.Get_md5_checksum())
	ec.Int(t.Get_key_fingerprint())

	return ec.GetBuffer()
}

func (t *TL_inputEncryptedFileUploaded) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Long()
	t._parts = dc.Int()
	t._md5_checksum = dc.String()
	t._key_fingerprint = dc.Int()

}

// inputEncryptedFile#5a17b5e5
type TL_inputEncryptedFile struct {
	_id          int64
	_access_hash int64
}

func (t *TL_inputEncryptedFile) Set_id(_id int64) {
	t._id = _id
}

func (t *TL_inputEncryptedFile) Get_id() int64 {
	return t._id
}

func (t *TL_inputEncryptedFile) Set_access_hash(_access_hash int64) {
	t._access_hash = _access_hash
}

func (t *TL_inputEncryptedFile) Get_access_hash() int64 {
	return t._access_hash
}

func New_TL_inputEncryptedFile() *TL_inputEncryptedFile {
	return &TL_inputEncryptedFile{}
}

func (t *TL_inputEncryptedFile) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputEncryptedFile))
	ec.Long(t.Get_id())
	ec.Long(t.Get_access_hash())

	return ec.GetBuffer()
}

func (t *TL_inputEncryptedFile) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Long()
	t._access_hash = dc.Long()

}

// inputEncryptedFileBigUploaded#2dc173c8
type TL_inputEncryptedFileBigUploaded struct {
	_id              int64
	_parts           int32
	_key_fingerprint int32
}

func (t *TL_inputEncryptedFileBigUploaded) Set_id(_id int64) {
	t._id = _id
}

func (t *TL_inputEncryptedFileBigUploaded) Get_id() int64 {
	return t._id
}

func (t *TL_inputEncryptedFileBigUploaded) Set_parts(_parts int32) {
	t._parts = _parts
}

func (t *TL_inputEncryptedFileBigUploaded) Get_parts() int32 {
	return t._parts
}

func (t *TL_inputEncryptedFileBigUploaded) Set_key_fingerprint(_key_fingerprint int32) {
	t._key_fingerprint = _key_fingerprint
}

func (t *TL_inputEncryptedFileBigUploaded) Get_key_fingerprint() int32 {
	return t._key_fingerprint
}

func New_TL_inputEncryptedFileBigUploaded() *TL_inputEncryptedFileBigUploaded {
	return &TL_inputEncryptedFileBigUploaded{}
}

func (t *TL_inputEncryptedFileBigUploaded) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputEncryptedFileBigUploaded))
	ec.Long(t.Get_id())
	ec.Int(t.Get_parts())
	ec.Int(t.Get_key_fingerprint())

	return ec.GetBuffer()
}

func (t *TL_inputEncryptedFileBigUploaded) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Long()
	t._parts = dc.Int()
	t._key_fingerprint = dc.Int()

}

// encryptedMessage#ed18c118
type TL_encryptedMessage struct {
	_random_id int64
	_chat_id   int32
	_date      int32
	_bytes     []byte
	_file      []byte
}

func (t *TL_encryptedMessage) Set_random_id(_random_id int64) {
	t._random_id = _random_id
}

func (t *TL_encryptedMessage) Get_random_id() int64 {
	return t._random_id
}

func (t *TL_encryptedMessage) Set_chat_id(_chat_id int32) {
	t._chat_id = _chat_id
}

func (t *TL_encryptedMessage) Get_chat_id() int32 {
	return t._chat_id
}

func (t *TL_encryptedMessage) Set_date(_date int32) {
	t._date = _date
}

func (t *TL_encryptedMessage) Get_date() int32 {
	return t._date
}

func (t *TL_encryptedMessage) Set_bytes(_bytes []byte) {
	t._bytes = _bytes
}

func (t *TL_encryptedMessage) Get_bytes() []byte {
	return t._bytes
}

func (t *TL_encryptedMessage) Set_file(_file []byte) {
	t._file = _file
}

func (t *TL_encryptedMessage) Get_file() []byte {
	return t._file
}

func New_TL_encryptedMessage() *TL_encryptedMessage {
	return &TL_encryptedMessage{}
}

func (t *TL_encryptedMessage) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_encryptedMessage))
	ec.Long(t.Get_random_id())
	ec.Int(t.Get_chat_id())
	ec.Int(t.Get_date())
	ec.Bytes(t.Get_bytes())
	ec.Bytes(t.Get_file())

	return ec.GetBuffer()
}

func (t *TL_encryptedMessage) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._random_id = dc.Long()
	t._chat_id = dc.Int()
	t._date = dc.Int()
	t._bytes = dc.Bytes(16)
	t._file = dc.Bytes(16)

}

// encryptedMessageService#23734b06
type TL_encryptedMessageService struct {
	_random_id int64
	_chat_id   int32
	_date      int32
	_bytes     []byte
}

func (t *TL_encryptedMessageService) Set_random_id(_random_id int64) {
	t._random_id = _random_id
}

func (t *TL_encryptedMessageService) Get_random_id() int64 {
	return t._random_id
}

func (t *TL_encryptedMessageService) Set_chat_id(_chat_id int32) {
	t._chat_id = _chat_id
}

func (t *TL_encryptedMessageService) Get_chat_id() int32 {
	return t._chat_id
}

func (t *TL_encryptedMessageService) Set_date(_date int32) {
	t._date = _date
}

func (t *TL_encryptedMessageService) Get_date() int32 {
	return t._date
}

func (t *TL_encryptedMessageService) Set_bytes(_bytes []byte) {
	t._bytes = _bytes
}

func (t *TL_encryptedMessageService) Get_bytes() []byte {
	return t._bytes
}

func New_TL_encryptedMessageService() *TL_encryptedMessageService {
	return &TL_encryptedMessageService{}
}

func (t *TL_encryptedMessageService) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_encryptedMessageService))
	ec.Long(t.Get_random_id())
	ec.Int(t.Get_chat_id())
	ec.Int(t.Get_date())
	ec.Bytes(t.Get_bytes())

	return ec.GetBuffer()
}

func (t *TL_encryptedMessageService) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._random_id = dc.Long()
	t._chat_id = dc.Int()
	t._date = dc.Int()
	t._bytes = dc.Bytes(16)

}

// messages_dhConfigNotModified#c0e24635
type TL_messages_dhConfigNotModified struct {
	_random []byte
}

func (t *TL_messages_dhConfigNotModified) Set_random(_random []byte) {
	t._random = _random
}

func (t *TL_messages_dhConfigNotModified) Get_random() []byte {
	return t._random
}

func New_TL_messages_dhConfigNotModified() *TL_messages_dhConfigNotModified {
	return &TL_messages_dhConfigNotModified{}
}

func (t *TL_messages_dhConfigNotModified) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_dhConfigNotModified))
	ec.Bytes(t.Get_random())

	return ec.GetBuffer()
}

func (t *TL_messages_dhConfigNotModified) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._random = dc.Bytes(16)

}

// messages_dhConfig#2c221edd
type TL_messages_dhConfig struct {
	_g       int32
	_p       []byte
	_version int32
	_random  []byte
}

func (t *TL_messages_dhConfig) Set_g(_g int32) {
	t._g = _g
}

func (t *TL_messages_dhConfig) Get_g() int32 {
	return t._g
}

func (t *TL_messages_dhConfig) Set_p(_p []byte) {
	t._p = _p
}

func (t *TL_messages_dhConfig) Get_p() []byte {
	return t._p
}

func (t *TL_messages_dhConfig) Set_version(_version int32) {
	t._version = _version
}

func (t *TL_messages_dhConfig) Get_version() int32 {
	return t._version
}

func (t *TL_messages_dhConfig) Set_random(_random []byte) {
	t._random = _random
}

func (t *TL_messages_dhConfig) Get_random() []byte {
	return t._random
}

func New_TL_messages_dhConfig() *TL_messages_dhConfig {
	return &TL_messages_dhConfig{}
}

func (t *TL_messages_dhConfig) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_dhConfig))
	ec.Int(t.Get_g())
	ec.Bytes(t.Get_p())
	ec.Int(t.Get_version())
	ec.Bytes(t.Get_random())

	return ec.GetBuffer()
}

func (t *TL_messages_dhConfig) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._g = dc.Int()
	t._p = dc.Bytes(16)
	t._version = dc.Int()
	t._random = dc.Bytes(16)

}

// messages_sentEncryptedMessage#560f8935
type TL_messages_sentEncryptedMessage struct {
	_date int32
}

func (t *TL_messages_sentEncryptedMessage) Set_date(_date int32) {
	t._date = _date
}

func (t *TL_messages_sentEncryptedMessage) Get_date() int32 {
	return t._date
}

func New_TL_messages_sentEncryptedMessage() *TL_messages_sentEncryptedMessage {
	return &TL_messages_sentEncryptedMessage{}
}

func (t *TL_messages_sentEncryptedMessage) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_sentEncryptedMessage))
	ec.Int(t.Get_date())

	return ec.GetBuffer()
}

func (t *TL_messages_sentEncryptedMessage) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._date = dc.Int()

}

// messages_sentEncryptedFile#9493ff32
type TL_messages_sentEncryptedFile struct {
	_date int32
	_file []byte
}

func (t *TL_messages_sentEncryptedFile) Set_date(_date int32) {
	t._date = _date
}

func (t *TL_messages_sentEncryptedFile) Get_date() int32 {
	return t._date
}

func (t *TL_messages_sentEncryptedFile) Set_file(_file []byte) {
	t._file = _file
}

func (t *TL_messages_sentEncryptedFile) Get_file() []byte {
	return t._file
}

func New_TL_messages_sentEncryptedFile() *TL_messages_sentEncryptedFile {
	return &TL_messages_sentEncryptedFile{}
}

func (t *TL_messages_sentEncryptedFile) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_sentEncryptedFile))
	ec.Int(t.Get_date())
	ec.Bytes(t.Get_file())

	return ec.GetBuffer()
}

func (t *TL_messages_sentEncryptedFile) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._date = dc.Int()
	t._file = dc.Bytes(16)

}

// inputDocumentEmpty#72f0eaae
type TL_inputDocumentEmpty struct {
}

func New_TL_inputDocumentEmpty() *TL_inputDocumentEmpty {
	return &TL_inputDocumentEmpty{}
}

func (t *TL_inputDocumentEmpty) Encode() []byte {
	return nil
}

func (t *TL_inputDocumentEmpty) Decode(b []byte) {

}

// inputDocument#18798952
type TL_inputDocument struct {
	_id          int64
	_access_hash int64
}

func (t *TL_inputDocument) Set_id(_id int64) {
	t._id = _id
}

func (t *TL_inputDocument) Get_id() int64 {
	return t._id
}

func (t *TL_inputDocument) Set_access_hash(_access_hash int64) {
	t._access_hash = _access_hash
}

func (t *TL_inputDocument) Get_access_hash() int64 {
	return t._access_hash
}

func New_TL_inputDocument() *TL_inputDocument {
	return &TL_inputDocument{}
}

func (t *TL_inputDocument) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputDocument))
	ec.Long(t.Get_id())
	ec.Long(t.Get_access_hash())

	return ec.GetBuffer()
}

func (t *TL_inputDocument) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Long()
	t._access_hash = dc.Long()

}

// documentEmpty#36f8c871
type TL_documentEmpty struct {
	_id int64
}

func (t *TL_documentEmpty) Set_id(_id int64) {
	t._id = _id
}

func (t *TL_documentEmpty) Get_id() int64 {
	return t._id
}

func New_TL_documentEmpty() *TL_documentEmpty {
	return &TL_documentEmpty{}
}

func (t *TL_documentEmpty) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_documentEmpty))
	ec.Long(t.Get_id())

	return ec.GetBuffer()
}

func (t *TL_documentEmpty) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Long()

}

// document#87232bc7
type TL_document struct {
	_id          int64
	_access_hash int64
	_date        int32
	_mime_type   string
	_size        int32
	_thumb       []byte
	_dc_id       int32
	_version     int32
	_attributes  []byte
}

func (t *TL_document) Set_id(_id int64) {
	t._id = _id
}

func (t *TL_document) Get_id() int64 {
	return t._id
}

func (t *TL_document) Set_access_hash(_access_hash int64) {
	t._access_hash = _access_hash
}

func (t *TL_document) Get_access_hash() int64 {
	return t._access_hash
}

func (t *TL_document) Set_date(_date int32) {
	t._date = _date
}

func (t *TL_document) Get_date() int32 {
	return t._date
}

func (t *TL_document) Set_mime_type(_mime_type string) {
	t._mime_type = _mime_type
}

func (t *TL_document) Get_mime_type() string {
	return t._mime_type
}

func (t *TL_document) Set_size(_size int32) {
	t._size = _size
}

func (t *TL_document) Get_size() int32 {
	return t._size
}

func (t *TL_document) Set_thumb(_thumb []byte) {
	t._thumb = _thumb
}

func (t *TL_document) Get_thumb() []byte {
	return t._thumb
}

func (t *TL_document) Set_dc_id(_dc_id int32) {
	t._dc_id = _dc_id
}

func (t *TL_document) Get_dc_id() int32 {
	return t._dc_id
}

func (t *TL_document) Set_version(_version int32) {
	t._version = _version
}

func (t *TL_document) Get_version() int32 {
	return t._version
}

func (t *TL_document) Set_attributes(_attributes []byte) {
	t._attributes = _attributes
}

func (t *TL_document) Get_attributes() []byte {
	return t._attributes
}

func New_TL_document() *TL_document {
	return &TL_document{}
}

func (t *TL_document) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_document))
	ec.Long(t.Get_id())
	ec.Long(t.Get_access_hash())
	ec.Int(t.Get_date())
	ec.String(t.Get_mime_type())
	ec.Int(t.Get_size())
	ec.Bytes(t.Get_thumb())
	ec.Int(t.Get_dc_id())
	ec.Int(t.Get_version())
	ec.Bytes(t.Get_attributes())

	return ec.GetBuffer()
}

func (t *TL_document) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Long()
	t._access_hash = dc.Long()
	t._date = dc.Int()
	t._mime_type = dc.String()
	t._size = dc.Int()
	t._thumb = dc.Bytes(16)
	t._dc_id = dc.Int()
	t._version = dc.Int()
	t._attributes = dc.Bytes(16)

}

// help_support#17c6b5f6
type TL_help_support struct {
	_phone_number string
	_user         []byte
}

func (t *TL_help_support) Set_phone_number(_phone_number string) {
	t._phone_number = _phone_number
}

func (t *TL_help_support) Get_phone_number() string {
	return t._phone_number
}

func (t *TL_help_support) Set_user(_user []byte) {
	t._user = _user
}

func (t *TL_help_support) Get_user() []byte {
	return t._user
}

func New_TL_help_support() *TL_help_support {
	return &TL_help_support{}
}

func (t *TL_help_support) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_help_support))
	ec.String(t.Get_phone_number())
	ec.Bytes(t.Get_user())

	return ec.GetBuffer()
}

func (t *TL_help_support) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._phone_number = dc.String()
	t._user = dc.Bytes(16)

}

// notifyPeer#9fd40bd8
type TL_notifyPeer struct {
	_peer []byte
}

func (t *TL_notifyPeer) Set_peer(_peer []byte) {
	t._peer = _peer
}

func (t *TL_notifyPeer) Get_peer() []byte {
	return t._peer
}

func New_TL_notifyPeer() *TL_notifyPeer {
	return &TL_notifyPeer{}
}

func (t *TL_notifyPeer) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_notifyPeer))
	ec.Bytes(t.Get_peer())

	return ec.GetBuffer()
}

func (t *TL_notifyPeer) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._peer = dc.Bytes(16)

}

// notifyUsers#b4c83b4c
type TL_notifyUsers struct {
}

func New_TL_notifyUsers() *TL_notifyUsers {
	return &TL_notifyUsers{}
}

func (t *TL_notifyUsers) Encode() []byte {
	return nil
}

func (t *TL_notifyUsers) Decode(b []byte) {

}

// notifyChats#c007cec3
type TL_notifyChats struct {
}

func New_TL_notifyChats() *TL_notifyChats {
	return &TL_notifyChats{}
}

func (t *TL_notifyChats) Encode() []byte {
	return nil
}

func (t *TL_notifyChats) Decode(b []byte) {

}

// notifyAll#74d07c60
type TL_notifyAll struct {
}

func New_TL_notifyAll() *TL_notifyAll {
	return &TL_notifyAll{}
}

func (t *TL_notifyAll) Encode() []byte {
	return nil
}

func (t *TL_notifyAll) Decode(b []byte) {

}

// sendMessageTypingAction#16bf744e
type TL_sendMessageTypingAction struct {
}

func New_TL_sendMessageTypingAction() *TL_sendMessageTypingAction {
	return &TL_sendMessageTypingAction{}
}

func (t *TL_sendMessageTypingAction) Encode() []byte {
	return nil
}

func (t *TL_sendMessageTypingAction) Decode(b []byte) {

}

// sendMessageCancelAction#fd5ec8f5
type TL_sendMessageCancelAction struct {
}

func New_TL_sendMessageCancelAction() *TL_sendMessageCancelAction {
	return &TL_sendMessageCancelAction{}
}

func (t *TL_sendMessageCancelAction) Encode() []byte {
	return nil
}

func (t *TL_sendMessageCancelAction) Decode(b []byte) {

}

// sendMessageRecordVideoAction#a187d66f
type TL_sendMessageRecordVideoAction struct {
}

func New_TL_sendMessageRecordVideoAction() *TL_sendMessageRecordVideoAction {
	return &TL_sendMessageRecordVideoAction{}
}

func (t *TL_sendMessageRecordVideoAction) Encode() []byte {
	return nil
}

func (t *TL_sendMessageRecordVideoAction) Decode(b []byte) {

}

// sendMessageUploadVideoAction#e9763aec
type TL_sendMessageUploadVideoAction struct {
	_progress int32
}

func (t *TL_sendMessageUploadVideoAction) Set_progress(_progress int32) {
	t._progress = _progress
}

func (t *TL_sendMessageUploadVideoAction) Get_progress() int32 {
	return t._progress
}

func New_TL_sendMessageUploadVideoAction() *TL_sendMessageUploadVideoAction {
	return &TL_sendMessageUploadVideoAction{}
}

func (t *TL_sendMessageUploadVideoAction) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_sendMessageUploadVideoAction))
	ec.Int(t.Get_progress())

	return ec.GetBuffer()
}

func (t *TL_sendMessageUploadVideoAction) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._progress = dc.Int()

}

// sendMessageRecordAudioAction#d52f73f7
type TL_sendMessageRecordAudioAction struct {
}

func New_TL_sendMessageRecordAudioAction() *TL_sendMessageRecordAudioAction {
	return &TL_sendMessageRecordAudioAction{}
}

func (t *TL_sendMessageRecordAudioAction) Encode() []byte {
	return nil
}

func (t *TL_sendMessageRecordAudioAction) Decode(b []byte) {

}

// sendMessageUploadAudioAction#f351d7ab
type TL_sendMessageUploadAudioAction struct {
	_progress int32
}

func (t *TL_sendMessageUploadAudioAction) Set_progress(_progress int32) {
	t._progress = _progress
}

func (t *TL_sendMessageUploadAudioAction) Get_progress() int32 {
	return t._progress
}

func New_TL_sendMessageUploadAudioAction() *TL_sendMessageUploadAudioAction {
	return &TL_sendMessageUploadAudioAction{}
}

func (t *TL_sendMessageUploadAudioAction) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_sendMessageUploadAudioAction))
	ec.Int(t.Get_progress())

	return ec.GetBuffer()
}

func (t *TL_sendMessageUploadAudioAction) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._progress = dc.Int()

}

// sendMessageUploadPhotoAction#d1d34a26
type TL_sendMessageUploadPhotoAction struct {
	_progress int32
}

func (t *TL_sendMessageUploadPhotoAction) Set_progress(_progress int32) {
	t._progress = _progress
}

func (t *TL_sendMessageUploadPhotoAction) Get_progress() int32 {
	return t._progress
}

func New_TL_sendMessageUploadPhotoAction() *TL_sendMessageUploadPhotoAction {
	return &TL_sendMessageUploadPhotoAction{}
}

func (t *TL_sendMessageUploadPhotoAction) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_sendMessageUploadPhotoAction))
	ec.Int(t.Get_progress())

	return ec.GetBuffer()
}

func (t *TL_sendMessageUploadPhotoAction) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._progress = dc.Int()

}

// sendMessageUploadDocumentAction#aa0cd9e4
type TL_sendMessageUploadDocumentAction struct {
	_progress int32
}

func (t *TL_sendMessageUploadDocumentAction) Set_progress(_progress int32) {
	t._progress = _progress
}

func (t *TL_sendMessageUploadDocumentAction) Get_progress() int32 {
	return t._progress
}

func New_TL_sendMessageUploadDocumentAction() *TL_sendMessageUploadDocumentAction {
	return &TL_sendMessageUploadDocumentAction{}
}

func (t *TL_sendMessageUploadDocumentAction) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_sendMessageUploadDocumentAction))
	ec.Int(t.Get_progress())

	return ec.GetBuffer()
}

func (t *TL_sendMessageUploadDocumentAction) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._progress = dc.Int()

}

// sendMessageGeoLocationAction#176f8ba1
type TL_sendMessageGeoLocationAction struct {
}

func New_TL_sendMessageGeoLocationAction() *TL_sendMessageGeoLocationAction {
	return &TL_sendMessageGeoLocationAction{}
}

func (t *TL_sendMessageGeoLocationAction) Encode() []byte {
	return nil
}

func (t *TL_sendMessageGeoLocationAction) Decode(b []byte) {

}

// sendMessageChooseContactAction#628cbc6f
type TL_sendMessageChooseContactAction struct {
}

func New_TL_sendMessageChooseContactAction() *TL_sendMessageChooseContactAction {
	return &TL_sendMessageChooseContactAction{}
}

func (t *TL_sendMessageChooseContactAction) Encode() []byte {
	return nil
}

func (t *TL_sendMessageChooseContactAction) Decode(b []byte) {

}

// sendMessageGamePlayAction#dd6a8f48
type TL_sendMessageGamePlayAction struct {
}

func New_TL_sendMessageGamePlayAction() *TL_sendMessageGamePlayAction {
	return &TL_sendMessageGamePlayAction{}
}

func (t *TL_sendMessageGamePlayAction) Encode() []byte {
	return nil
}

func (t *TL_sendMessageGamePlayAction) Decode(b []byte) {

}

// sendMessageRecordRoundAction#88f27fbc
type TL_sendMessageRecordRoundAction struct {
}

func New_TL_sendMessageRecordRoundAction() *TL_sendMessageRecordRoundAction {
	return &TL_sendMessageRecordRoundAction{}
}

func (t *TL_sendMessageRecordRoundAction) Encode() []byte {
	return nil
}

func (t *TL_sendMessageRecordRoundAction) Decode(b []byte) {

}

// sendMessageUploadRoundAction#243e1c66
type TL_sendMessageUploadRoundAction struct {
	_progress int32
}

func (t *TL_sendMessageUploadRoundAction) Set_progress(_progress int32) {
	t._progress = _progress
}

func (t *TL_sendMessageUploadRoundAction) Get_progress() int32 {
	return t._progress
}

func New_TL_sendMessageUploadRoundAction() *TL_sendMessageUploadRoundAction {
	return &TL_sendMessageUploadRoundAction{}
}

func (t *TL_sendMessageUploadRoundAction) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_sendMessageUploadRoundAction))
	ec.Int(t.Get_progress())

	return ec.GetBuffer()
}

func (t *TL_sendMessageUploadRoundAction) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._progress = dc.Int()

}

// contacts_found#1aa1f784
type TL_contacts_found struct {
	_results []byte
	_chats   []byte
	_users   []byte
}

func (t *TL_contacts_found) Set_results(_results []byte) {
	t._results = _results
}

func (t *TL_contacts_found) Get_results() []byte {
	return t._results
}

func (t *TL_contacts_found) Set_chats(_chats []byte) {
	t._chats = _chats
}

func (t *TL_contacts_found) Get_chats() []byte {
	return t._chats
}

func (t *TL_contacts_found) Set_users(_users []byte) {
	t._users = _users
}

func (t *TL_contacts_found) Get_users() []byte {
	return t._users
}

func New_TL_contacts_found() *TL_contacts_found {
	return &TL_contacts_found{}
}

func (t *TL_contacts_found) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_contacts_found))
	ec.Bytes(t.Get_results())
	ec.Bytes(t.Get_chats())
	ec.Bytes(t.Get_users())

	return ec.GetBuffer()
}

func (t *TL_contacts_found) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._results = dc.Bytes(16)
	t._chats = dc.Bytes(16)
	t._users = dc.Bytes(16)

}

// inputPrivacyKeyStatusTimestamp#4f96cb18
type TL_inputPrivacyKeyStatusTimestamp struct {
}

func New_TL_inputPrivacyKeyStatusTimestamp() *TL_inputPrivacyKeyStatusTimestamp {
	return &TL_inputPrivacyKeyStatusTimestamp{}
}

func (t *TL_inputPrivacyKeyStatusTimestamp) Encode() []byte {
	return nil
}

func (t *TL_inputPrivacyKeyStatusTimestamp) Decode(b []byte) {

}

// inputPrivacyKeyChatInvite#bdfb0426
type TL_inputPrivacyKeyChatInvite struct {
}

func New_TL_inputPrivacyKeyChatInvite() *TL_inputPrivacyKeyChatInvite {
	return &TL_inputPrivacyKeyChatInvite{}
}

func (t *TL_inputPrivacyKeyChatInvite) Encode() []byte {
	return nil
}

func (t *TL_inputPrivacyKeyChatInvite) Decode(b []byte) {

}

// inputPrivacyKeyPhoneCall#fabadc5f
type TL_inputPrivacyKeyPhoneCall struct {
}

func New_TL_inputPrivacyKeyPhoneCall() *TL_inputPrivacyKeyPhoneCall {
	return &TL_inputPrivacyKeyPhoneCall{}
}

func (t *TL_inputPrivacyKeyPhoneCall) Encode() []byte {
	return nil
}

func (t *TL_inputPrivacyKeyPhoneCall) Decode(b []byte) {

}

// privacyKeyStatusTimestamp#bc2eab30
type TL_privacyKeyStatusTimestamp struct {
}

func New_TL_privacyKeyStatusTimestamp() *TL_privacyKeyStatusTimestamp {
	return &TL_privacyKeyStatusTimestamp{}
}

func (t *TL_privacyKeyStatusTimestamp) Encode() []byte {
	return nil
}

func (t *TL_privacyKeyStatusTimestamp) Decode(b []byte) {

}

// privacyKeyChatInvite#500e6dfa
type TL_privacyKeyChatInvite struct {
}

func New_TL_privacyKeyChatInvite() *TL_privacyKeyChatInvite {
	return &TL_privacyKeyChatInvite{}
}

func (t *TL_privacyKeyChatInvite) Encode() []byte {
	return nil
}

func (t *TL_privacyKeyChatInvite) Decode(b []byte) {

}

// privacyKeyPhoneCall#3d662b7b
type TL_privacyKeyPhoneCall struct {
}

func New_TL_privacyKeyPhoneCall() *TL_privacyKeyPhoneCall {
	return &TL_privacyKeyPhoneCall{}
}

func (t *TL_privacyKeyPhoneCall) Encode() []byte {
	return nil
}

func (t *TL_privacyKeyPhoneCall) Decode(b []byte) {

}

// inputPrivacyValueAllowContacts#0d09e07b
type TL_inputPrivacyValueAllowContacts struct {
}

func New_TL_inputPrivacyValueAllowContacts() *TL_inputPrivacyValueAllowContacts {
	return &TL_inputPrivacyValueAllowContacts{}
}

func (t *TL_inputPrivacyValueAllowContacts) Encode() []byte {
	return nil
}

func (t *TL_inputPrivacyValueAllowContacts) Decode(b []byte) {

}

// inputPrivacyValueAllowAll#184b35ce
type TL_inputPrivacyValueAllowAll struct {
}

func New_TL_inputPrivacyValueAllowAll() *TL_inputPrivacyValueAllowAll {
	return &TL_inputPrivacyValueAllowAll{}
}

func (t *TL_inputPrivacyValueAllowAll) Encode() []byte {
	return nil
}

func (t *TL_inputPrivacyValueAllowAll) Decode(b []byte) {

}

// inputPrivacyValueAllowUsers#131cc67f
type TL_inputPrivacyValueAllowUsers struct {
	_users []byte
}

func (t *TL_inputPrivacyValueAllowUsers) Set_users(_users []byte) {
	t._users = _users
}

func (t *TL_inputPrivacyValueAllowUsers) Get_users() []byte {
	return t._users
}

func New_TL_inputPrivacyValueAllowUsers() *TL_inputPrivacyValueAllowUsers {
	return &TL_inputPrivacyValueAllowUsers{}
}

func (t *TL_inputPrivacyValueAllowUsers) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputPrivacyValueAllowUsers))
	ec.Bytes(t.Get_users())

	return ec.GetBuffer()
}

func (t *TL_inputPrivacyValueAllowUsers) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._users = dc.Bytes(16)

}

// inputPrivacyValueDisallowContacts#0ba52007
type TL_inputPrivacyValueDisallowContacts struct {
}

func New_TL_inputPrivacyValueDisallowContacts() *TL_inputPrivacyValueDisallowContacts {
	return &TL_inputPrivacyValueDisallowContacts{}
}

func (t *TL_inputPrivacyValueDisallowContacts) Encode() []byte {
	return nil
}

func (t *TL_inputPrivacyValueDisallowContacts) Decode(b []byte) {

}

// inputPrivacyValueDisallowAll#d66b66c9
type TL_inputPrivacyValueDisallowAll struct {
}

func New_TL_inputPrivacyValueDisallowAll() *TL_inputPrivacyValueDisallowAll {
	return &TL_inputPrivacyValueDisallowAll{}
}

func (t *TL_inputPrivacyValueDisallowAll) Encode() []byte {
	return nil
}

func (t *TL_inputPrivacyValueDisallowAll) Decode(b []byte) {

}

// inputPrivacyValueDisallowUsers#90110467
type TL_inputPrivacyValueDisallowUsers struct {
	_users []byte
}

func (t *TL_inputPrivacyValueDisallowUsers) Set_users(_users []byte) {
	t._users = _users
}

func (t *TL_inputPrivacyValueDisallowUsers) Get_users() []byte {
	return t._users
}

func New_TL_inputPrivacyValueDisallowUsers() *TL_inputPrivacyValueDisallowUsers {
	return &TL_inputPrivacyValueDisallowUsers{}
}

func (t *TL_inputPrivacyValueDisallowUsers) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputPrivacyValueDisallowUsers))
	ec.Bytes(t.Get_users())

	return ec.GetBuffer()
}

func (t *TL_inputPrivacyValueDisallowUsers) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._users = dc.Bytes(16)

}

// privacyValueAllowContacts#fffe1bac
type TL_privacyValueAllowContacts struct {
}

func New_TL_privacyValueAllowContacts() *TL_privacyValueAllowContacts {
	return &TL_privacyValueAllowContacts{}
}

func (t *TL_privacyValueAllowContacts) Encode() []byte {
	return nil
}

func (t *TL_privacyValueAllowContacts) Decode(b []byte) {

}

// privacyValueAllowAll#65427b82
type TL_privacyValueAllowAll struct {
}

func New_TL_privacyValueAllowAll() *TL_privacyValueAllowAll {
	return &TL_privacyValueAllowAll{}
}

func (t *TL_privacyValueAllowAll) Encode() []byte {
	return nil
}

func (t *TL_privacyValueAllowAll) Decode(b []byte) {

}

// privacyValueAllowUsers#4d5bbe0c
type TL_privacyValueAllowUsers struct {
	_users []byte
}

func (t *TL_privacyValueAllowUsers) Set_users(_users []byte) {
	t._users = _users
}

func (t *TL_privacyValueAllowUsers) Get_users() []byte {
	return t._users
}

func New_TL_privacyValueAllowUsers() *TL_privacyValueAllowUsers {
	return &TL_privacyValueAllowUsers{}
}

func (t *TL_privacyValueAllowUsers) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_privacyValueAllowUsers))
	ec.Bytes(t.Get_users())

	return ec.GetBuffer()
}

func (t *TL_privacyValueAllowUsers) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._users = dc.Bytes(16)

}

// privacyValueDisallowContacts#f888fa1a
type TL_privacyValueDisallowContacts struct {
}

func New_TL_privacyValueDisallowContacts() *TL_privacyValueDisallowContacts {
	return &TL_privacyValueDisallowContacts{}
}

func (t *TL_privacyValueDisallowContacts) Encode() []byte {
	return nil
}

func (t *TL_privacyValueDisallowContacts) Decode(b []byte) {

}

// privacyValueDisallowAll#8b73e763
type TL_privacyValueDisallowAll struct {
}

func New_TL_privacyValueDisallowAll() *TL_privacyValueDisallowAll {
	return &TL_privacyValueDisallowAll{}
}

func (t *TL_privacyValueDisallowAll) Encode() []byte {
	return nil
}

func (t *TL_privacyValueDisallowAll) Decode(b []byte) {

}

// privacyValueDisallowUsers#0c7f49b7
type TL_privacyValueDisallowUsers struct {
	_users []byte
}

func (t *TL_privacyValueDisallowUsers) Set_users(_users []byte) {
	t._users = _users
}

func (t *TL_privacyValueDisallowUsers) Get_users() []byte {
	return t._users
}

func New_TL_privacyValueDisallowUsers() *TL_privacyValueDisallowUsers {
	return &TL_privacyValueDisallowUsers{}
}

func (t *TL_privacyValueDisallowUsers) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_privacyValueDisallowUsers))
	ec.Bytes(t.Get_users())

	return ec.GetBuffer()
}

func (t *TL_privacyValueDisallowUsers) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._users = dc.Bytes(16)

}

// account_privacyRules#554abb6f
type TL_account_privacyRules struct {
	_rules []byte
	_users []byte
}

func (t *TL_account_privacyRules) Set_rules(_rules []byte) {
	t._rules = _rules
}

func (t *TL_account_privacyRules) Get_rules() []byte {
	return t._rules
}

func (t *TL_account_privacyRules) Set_users(_users []byte) {
	t._users = _users
}

func (t *TL_account_privacyRules) Get_users() []byte {
	return t._users
}

func New_TL_account_privacyRules() *TL_account_privacyRules {
	return &TL_account_privacyRules{}
}

func (t *TL_account_privacyRules) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_account_privacyRules))
	ec.Bytes(t.Get_rules())
	ec.Bytes(t.Get_users())

	return ec.GetBuffer()
}

func (t *TL_account_privacyRules) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._rules = dc.Bytes(16)
	t._users = dc.Bytes(16)

}

// accountDaysTTL#b8d0afdf
type TL_accountDaysTTL struct {
	_days int32
}

func (t *TL_accountDaysTTL) Set_days(_days int32) {
	t._days = _days
}

func (t *TL_accountDaysTTL) Get_days() int32 {
	return t._days
}

func New_TL_accountDaysTTL() *TL_accountDaysTTL {
	return &TL_accountDaysTTL{}
}

func (t *TL_accountDaysTTL) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_accountDaysTTL))
	ec.Int(t.Get_days())

	return ec.GetBuffer()
}

func (t *TL_accountDaysTTL) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._days = dc.Int()

}

// documentAttributeImageSize#6c37c15c
type TL_documentAttributeImageSize struct {
	_w int32
	_h int32
}

func (t *TL_documentAttributeImageSize) Set_w(_w int32) {
	t._w = _w
}

func (t *TL_documentAttributeImageSize) Get_w() int32 {
	return t._w
}

func (t *TL_documentAttributeImageSize) Set_h(_h int32) {
	t._h = _h
}

func (t *TL_documentAttributeImageSize) Get_h() int32 {
	return t._h
}

func New_TL_documentAttributeImageSize() *TL_documentAttributeImageSize {
	return &TL_documentAttributeImageSize{}
}

func (t *TL_documentAttributeImageSize) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_documentAttributeImageSize))
	ec.Int(t.Get_w())
	ec.Int(t.Get_h())

	return ec.GetBuffer()
}

func (t *TL_documentAttributeImageSize) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._w = dc.Int()
	t._h = dc.Int()

}

// documentAttributeAnimated#11b58939
type TL_documentAttributeAnimated struct {
}

func New_TL_documentAttributeAnimated() *TL_documentAttributeAnimated {
	return &TL_documentAttributeAnimated{}
}

func (t *TL_documentAttributeAnimated) Encode() []byte {
	return nil
}

func (t *TL_documentAttributeAnimated) Decode(b []byte) {

}

// documentAttributeSticker#6319d612
type TL_documentAttributeSticker struct {
	_flags       []byte
	_mask        []byte
	_alt         string
	_stickerset  []byte
	_mask_coords []byte
}

func (t *TL_documentAttributeSticker) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_documentAttributeSticker) Get_flags() []byte {
	return t._flags
}

func (t *TL_documentAttributeSticker) Set_mask(_mask []byte) {
	t._mask = _mask
}

func (t *TL_documentAttributeSticker) Get_mask() []byte {
	return t._mask
}

func (t *TL_documentAttributeSticker) Set_alt(_alt string) {
	t._alt = _alt
}

func (t *TL_documentAttributeSticker) Get_alt() string {
	return t._alt
}

func (t *TL_documentAttributeSticker) Set_stickerset(_stickerset []byte) {
	t._stickerset = _stickerset
}

func (t *TL_documentAttributeSticker) Get_stickerset() []byte {
	return t._stickerset
}

func (t *TL_documentAttributeSticker) Set_mask_coords(_mask_coords []byte) {
	t._mask_coords = _mask_coords
}

func (t *TL_documentAttributeSticker) Get_mask_coords() []byte {
	return t._mask_coords
}

func New_TL_documentAttributeSticker() *TL_documentAttributeSticker {
	return &TL_documentAttributeSticker{}
}

func (t *TL_documentAttributeSticker) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_documentAttributeSticker))
	ec.Bytes(t.Get_mask())
	ec.String(t.Get_alt())
	ec.Bytes(t.Get_stickerset())
	ec.Bytes(t.Get_mask_coords())

	return ec.GetBuffer()
}

func (t *TL_documentAttributeSticker) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._mask = dc.Bytes(16)
	t._alt = dc.String()
	t._stickerset = dc.Bytes(16)
	t._mask_coords = dc.Bytes(16)

}

// documentAttributeVideo#0ef02ce6
type TL_documentAttributeVideo struct {
	_flags         []byte
	_round_message []byte
	_duration      int32
	_w             int32
	_h             int32
}

func (t *TL_documentAttributeVideo) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_documentAttributeVideo) Get_flags() []byte {
	return t._flags
}

func (t *TL_documentAttributeVideo) Set_round_message(_round_message []byte) {
	t._round_message = _round_message
}

func (t *TL_documentAttributeVideo) Get_round_message() []byte {
	return t._round_message
}

func (t *TL_documentAttributeVideo) Set_duration(_duration int32) {
	t._duration = _duration
}

func (t *TL_documentAttributeVideo) Get_duration() int32 {
	return t._duration
}

func (t *TL_documentAttributeVideo) Set_w(_w int32) {
	t._w = _w
}

func (t *TL_documentAttributeVideo) Get_w() int32 {
	return t._w
}

func (t *TL_documentAttributeVideo) Set_h(_h int32) {
	t._h = _h
}

func (t *TL_documentAttributeVideo) Get_h() int32 {
	return t._h
}

func New_TL_documentAttributeVideo() *TL_documentAttributeVideo {
	return &TL_documentAttributeVideo{}
}

func (t *TL_documentAttributeVideo) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_documentAttributeVideo))
	ec.Bytes(t.Get_round_message())
	ec.Int(t.Get_duration())
	ec.Int(t.Get_w())
	ec.Int(t.Get_h())

	return ec.GetBuffer()
}

func (t *TL_documentAttributeVideo) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._round_message = dc.Bytes(16)
	t._duration = dc.Int()
	t._w = dc.Int()
	t._h = dc.Int()

}

// documentAttributeAudio#9852f9c6
type TL_documentAttributeAudio struct {
	_flags     []byte
	_voice     []byte
	_duration  int32
	_title     []byte
	_performer []byte
	_waveform  []byte
}

func (t *TL_documentAttributeAudio) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_documentAttributeAudio) Get_flags() []byte {
	return t._flags
}

func (t *TL_documentAttributeAudio) Set_voice(_voice []byte) {
	t._voice = _voice
}

func (t *TL_documentAttributeAudio) Get_voice() []byte {
	return t._voice
}

func (t *TL_documentAttributeAudio) Set_duration(_duration int32) {
	t._duration = _duration
}

func (t *TL_documentAttributeAudio) Get_duration() int32 {
	return t._duration
}

func (t *TL_documentAttributeAudio) Set_title(_title []byte) {
	t._title = _title
}

func (t *TL_documentAttributeAudio) Get_title() []byte {
	return t._title
}

func (t *TL_documentAttributeAudio) Set_performer(_performer []byte) {
	t._performer = _performer
}

func (t *TL_documentAttributeAudio) Get_performer() []byte {
	return t._performer
}

func (t *TL_documentAttributeAudio) Set_waveform(_waveform []byte) {
	t._waveform = _waveform
}

func (t *TL_documentAttributeAudio) Get_waveform() []byte {
	return t._waveform
}

func New_TL_documentAttributeAudio() *TL_documentAttributeAudio {
	return &TL_documentAttributeAudio{}
}

func (t *TL_documentAttributeAudio) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_documentAttributeAudio))
	ec.Bytes(t.Get_voice())
	ec.Int(t.Get_duration())
	ec.Bytes(t.Get_title())
	ec.Bytes(t.Get_performer())
	ec.Bytes(t.Get_waveform())

	return ec.GetBuffer()
}

func (t *TL_documentAttributeAudio) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._voice = dc.Bytes(16)
	t._duration = dc.Int()
	t._title = dc.Bytes(16)
	t._performer = dc.Bytes(16)
	t._waveform = dc.Bytes(16)

}

// documentAttributeFilename#15590068
type TL_documentAttributeFilename struct {
	_file_name string
}

func (t *TL_documentAttributeFilename) Set_file_name(_file_name string) {
	t._file_name = _file_name
}

func (t *TL_documentAttributeFilename) Get_file_name() string {
	return t._file_name
}

func New_TL_documentAttributeFilename() *TL_documentAttributeFilename {
	return &TL_documentAttributeFilename{}
}

func (t *TL_documentAttributeFilename) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_documentAttributeFilename))
	ec.String(t.Get_file_name())

	return ec.GetBuffer()
}

func (t *TL_documentAttributeFilename) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._file_name = dc.String()

}

// documentAttributeHasStickers#9801d2f7
type TL_documentAttributeHasStickers struct {
}

func New_TL_documentAttributeHasStickers() *TL_documentAttributeHasStickers {
	return &TL_documentAttributeHasStickers{}
}

func (t *TL_documentAttributeHasStickers) Encode() []byte {
	return nil
}

func (t *TL_documentAttributeHasStickers) Decode(b []byte) {

}

// messages_stickersNotModified#f1749a22
type TL_messages_stickersNotModified struct {
}

func New_TL_messages_stickersNotModified() *TL_messages_stickersNotModified {
	return &TL_messages_stickersNotModified{}
}

func (t *TL_messages_stickersNotModified) Encode() []byte {
	return nil
}

func (t *TL_messages_stickersNotModified) Decode(b []byte) {

}

// messages_stickers#8a8ecd32
type TL_messages_stickers struct {
	_hash     string
	_stickers []byte
}

func (t *TL_messages_stickers) Set_hash(_hash string) {
	t._hash = _hash
}

func (t *TL_messages_stickers) Get_hash() string {
	return t._hash
}

func (t *TL_messages_stickers) Set_stickers(_stickers []byte) {
	t._stickers = _stickers
}

func (t *TL_messages_stickers) Get_stickers() []byte {
	return t._stickers
}

func New_TL_messages_stickers() *TL_messages_stickers {
	return &TL_messages_stickers{}
}

func (t *TL_messages_stickers) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_stickers))
	ec.String(t.Get_hash())
	ec.Bytes(t.Get_stickers())

	return ec.GetBuffer()
}

func (t *TL_messages_stickers) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._hash = dc.String()
	t._stickers = dc.Bytes(16)

}

// stickerPack#12b299d4
type TL_stickerPack struct {
	_emoticon  string
	_documents []byte
}

func (t *TL_stickerPack) Set_emoticon(_emoticon string) {
	t._emoticon = _emoticon
}

func (t *TL_stickerPack) Get_emoticon() string {
	return t._emoticon
}

func (t *TL_stickerPack) Set_documents(_documents []byte) {
	t._documents = _documents
}

func (t *TL_stickerPack) Get_documents() []byte {
	return t._documents
}

func New_TL_stickerPack() *TL_stickerPack {
	return &TL_stickerPack{}
}

func (t *TL_stickerPack) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_stickerPack))
	ec.String(t.Get_emoticon())
	ec.Bytes(t.Get_documents())

	return ec.GetBuffer()
}

func (t *TL_stickerPack) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._emoticon = dc.String()
	t._documents = dc.Bytes(16)

}

// messages_allStickersNotModified#e86602c3
type TL_messages_allStickersNotModified struct {
}

func New_TL_messages_allStickersNotModified() *TL_messages_allStickersNotModified {
	return &TL_messages_allStickersNotModified{}
}

func (t *TL_messages_allStickersNotModified) Encode() []byte {
	return nil
}

func (t *TL_messages_allStickersNotModified) Decode(b []byte) {

}

// messages_allStickers#edfd405f
type TL_messages_allStickers struct {
	_hash int32
	_sets []byte
}

func (t *TL_messages_allStickers) Set_hash(_hash int32) {
	t._hash = _hash
}

func (t *TL_messages_allStickers) Get_hash() int32 {
	return t._hash
}

func (t *TL_messages_allStickers) Set_sets(_sets []byte) {
	t._sets = _sets
}

func (t *TL_messages_allStickers) Get_sets() []byte {
	return t._sets
}

func New_TL_messages_allStickers() *TL_messages_allStickers {
	return &TL_messages_allStickers{}
}

func (t *TL_messages_allStickers) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_allStickers))
	ec.Int(t.Get_hash())
	ec.Bytes(t.Get_sets())

	return ec.GetBuffer()
}

func (t *TL_messages_allStickers) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._hash = dc.Int()
	t._sets = dc.Bytes(16)

}

// disabledFeature#ae636f24
type TL_disabledFeature struct {
	_feature     string
	_description string
}

func (t *TL_disabledFeature) Set_feature(_feature string) {
	t._feature = _feature
}

func (t *TL_disabledFeature) Get_feature() string {
	return t._feature
}

func (t *TL_disabledFeature) Set_description(_description string) {
	t._description = _description
}

func (t *TL_disabledFeature) Get_description() string {
	return t._description
}

func New_TL_disabledFeature() *TL_disabledFeature {
	return &TL_disabledFeature{}
}

func (t *TL_disabledFeature) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_disabledFeature))
	ec.String(t.Get_feature())
	ec.String(t.Get_description())

	return ec.GetBuffer()
}

func (t *TL_disabledFeature) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._feature = dc.String()
	t._description = dc.String()

}

// messages_affectedMessages#84d19185
type TL_messages_affectedMessages struct {
	_pts       int32
	_pts_count int32
}

func (t *TL_messages_affectedMessages) Set_pts(_pts int32) {
	t._pts = _pts
}

func (t *TL_messages_affectedMessages) Get_pts() int32 {
	return t._pts
}

func (t *TL_messages_affectedMessages) Set_pts_count(_pts_count int32) {
	t._pts_count = _pts_count
}

func (t *TL_messages_affectedMessages) Get_pts_count() int32 {
	return t._pts_count
}

func New_TL_messages_affectedMessages() *TL_messages_affectedMessages {
	return &TL_messages_affectedMessages{}
}

func (t *TL_messages_affectedMessages) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_affectedMessages))
	ec.Int(t.Get_pts())
	ec.Int(t.Get_pts_count())

	return ec.GetBuffer()
}

func (t *TL_messages_affectedMessages) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._pts = dc.Int()
	t._pts_count = dc.Int()

}

// contactLinkUnknown#5f4f9247
type TL_contactLinkUnknown struct {
}

func New_TL_contactLinkUnknown() *TL_contactLinkUnknown {
	return &TL_contactLinkUnknown{}
}

func (t *TL_contactLinkUnknown) Encode() []byte {
	return nil
}

func (t *TL_contactLinkUnknown) Decode(b []byte) {

}

// contactLinkNone#feedd3ad
type TL_contactLinkNone struct {
}

func New_TL_contactLinkNone() *TL_contactLinkNone {
	return &TL_contactLinkNone{}
}

func (t *TL_contactLinkNone) Encode() []byte {
	return nil
}

func (t *TL_contactLinkNone) Decode(b []byte) {

}

// contactLinkHasPhone#268f3f59
type TL_contactLinkHasPhone struct {
}

func New_TL_contactLinkHasPhone() *TL_contactLinkHasPhone {
	return &TL_contactLinkHasPhone{}
}

func (t *TL_contactLinkHasPhone) Encode() []byte {
	return nil
}

func (t *TL_contactLinkHasPhone) Decode(b []byte) {

}

// contactLinkContact#d502c2d0
type TL_contactLinkContact struct {
}

func New_TL_contactLinkContact() *TL_contactLinkContact {
	return &TL_contactLinkContact{}
}

func (t *TL_contactLinkContact) Encode() []byte {
	return nil
}

func (t *TL_contactLinkContact) Decode(b []byte) {

}

// webPageEmpty#eb1477e8
type TL_webPageEmpty struct {
	_id int64
}

func (t *TL_webPageEmpty) Set_id(_id int64) {
	t._id = _id
}

func (t *TL_webPageEmpty) Get_id() int64 {
	return t._id
}

func New_TL_webPageEmpty() *TL_webPageEmpty {
	return &TL_webPageEmpty{}
}

func (t *TL_webPageEmpty) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_webPageEmpty))
	ec.Long(t.Get_id())

	return ec.GetBuffer()
}

func (t *TL_webPageEmpty) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Long()

}

// webPagePending#c586da1c
type TL_webPagePending struct {
	_id   int64
	_date int32
}

func (t *TL_webPagePending) Set_id(_id int64) {
	t._id = _id
}

func (t *TL_webPagePending) Get_id() int64 {
	return t._id
}

func (t *TL_webPagePending) Set_date(_date int32) {
	t._date = _date
}

func (t *TL_webPagePending) Get_date() int32 {
	return t._date
}

func New_TL_webPagePending() *TL_webPagePending {
	return &TL_webPagePending{}
}

func (t *TL_webPagePending) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_webPagePending))
	ec.Long(t.Get_id())
	ec.Int(t.Get_date())

	return ec.GetBuffer()
}

func (t *TL_webPagePending) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Long()
	t._date = dc.Int()

}

// webPage#5f07b4bc
type TL_webPage struct {
	_flags        []byte
	_id           int64
	_url          string
	_display_url  string
	_hash         int32
	_type         []byte
	_site_name    []byte
	_title        []byte
	_description  []byte
	_photo        []byte
	_embed_url    []byte
	_embed_type   []byte
	_embed_width  []byte
	_embed_height []byte
	_duration     []byte
	_author       []byte
	_document     []byte
	_cached_page  []byte
}

func (t *TL_webPage) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_webPage) Get_flags() []byte {
	return t._flags
}

func (t *TL_webPage) Set_id(_id int64) {
	t._id = _id
}

func (t *TL_webPage) Get_id() int64 {
	return t._id
}

func (t *TL_webPage) Set_url(_url string) {
	t._url = _url
}

func (t *TL_webPage) Get_url() string {
	return t._url
}

func (t *TL_webPage) Set_display_url(_display_url string) {
	t._display_url = _display_url
}

func (t *TL_webPage) Get_display_url() string {
	return t._display_url
}

func (t *TL_webPage) Set_hash(_hash int32) {
	t._hash = _hash
}

func (t *TL_webPage) Get_hash() int32 {
	return t._hash
}

func (t *TL_webPage) Set_type(_type []byte) {
	t._type = _type
}

func (t *TL_webPage) Get_type() []byte {
	return t._type
}

func (t *TL_webPage) Set_site_name(_site_name []byte) {
	t._site_name = _site_name
}

func (t *TL_webPage) Get_site_name() []byte {
	return t._site_name
}

func (t *TL_webPage) Set_title(_title []byte) {
	t._title = _title
}

func (t *TL_webPage) Get_title() []byte {
	return t._title
}

func (t *TL_webPage) Set_description(_description []byte) {
	t._description = _description
}

func (t *TL_webPage) Get_description() []byte {
	return t._description
}

func (t *TL_webPage) Set_photo(_photo []byte) {
	t._photo = _photo
}

func (t *TL_webPage) Get_photo() []byte {
	return t._photo
}

func (t *TL_webPage) Set_embed_url(_embed_url []byte) {
	t._embed_url = _embed_url
}

func (t *TL_webPage) Get_embed_url() []byte {
	return t._embed_url
}

func (t *TL_webPage) Set_embed_type(_embed_type []byte) {
	t._embed_type = _embed_type
}

func (t *TL_webPage) Get_embed_type() []byte {
	return t._embed_type
}

func (t *TL_webPage) Set_embed_width(_embed_width []byte) {
	t._embed_width = _embed_width
}

func (t *TL_webPage) Get_embed_width() []byte {
	return t._embed_width
}

func (t *TL_webPage) Set_embed_height(_embed_height []byte) {
	t._embed_height = _embed_height
}

func (t *TL_webPage) Get_embed_height() []byte {
	return t._embed_height
}

func (t *TL_webPage) Set_duration(_duration []byte) {
	t._duration = _duration
}

func (t *TL_webPage) Get_duration() []byte {
	return t._duration
}

func (t *TL_webPage) Set_author(_author []byte) {
	t._author = _author
}

func (t *TL_webPage) Get_author() []byte {
	return t._author
}

func (t *TL_webPage) Set_document(_document []byte) {
	t._document = _document
}

func (t *TL_webPage) Get_document() []byte {
	return t._document
}

func (t *TL_webPage) Set_cached_page(_cached_page []byte) {
	t._cached_page = _cached_page
}

func (t *TL_webPage) Get_cached_page() []byte {
	return t._cached_page
}

func New_TL_webPage() *TL_webPage {
	return &TL_webPage{}
}

func (t *TL_webPage) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_webPage))
	ec.Long(t.Get_id())
	ec.String(t.Get_url())
	ec.String(t.Get_display_url())
	ec.Int(t.Get_hash())
	ec.Bytes(t.Get_type())
	ec.Bytes(t.Get_site_name())
	ec.Bytes(t.Get_title())
	ec.Bytes(t.Get_description())
	ec.Bytes(t.Get_photo())
	ec.Bytes(t.Get_embed_url())
	ec.Bytes(t.Get_embed_type())
	ec.Bytes(t.Get_embed_width())
	ec.Bytes(t.Get_embed_height())
	ec.Bytes(t.Get_duration())
	ec.Bytes(t.Get_author())
	ec.Bytes(t.Get_document())
	ec.Bytes(t.Get_cached_page())

	return ec.GetBuffer()
}

func (t *TL_webPage) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Long()
	t._url = dc.String()
	t._display_url = dc.String()
	t._hash = dc.Int()
	t._type = dc.Bytes(16)
	t._site_name = dc.Bytes(16)
	t._title = dc.Bytes(16)
	t._description = dc.Bytes(16)
	t._photo = dc.Bytes(16)
	t._embed_url = dc.Bytes(16)
	t._embed_type = dc.Bytes(16)
	t._embed_width = dc.Bytes(16)
	t._embed_height = dc.Bytes(16)
	t._duration = dc.Bytes(16)
	t._author = dc.Bytes(16)
	t._document = dc.Bytes(16)
	t._cached_page = dc.Bytes(16)

}

// webPageNotModified#85849473
type TL_webPageNotModified struct {
}

func New_TL_webPageNotModified() *TL_webPageNotModified {
	return &TL_webPageNotModified{}
}

func (t *TL_webPageNotModified) Encode() []byte {
	return nil
}

func (t *TL_webPageNotModified) Decode(b []byte) {

}

// authorization#7bf2e6f6
type TL_authorization struct {
	_hash           int64
	_flags          int32
	_device_model   string
	_platform       string
	_system_version string
	_api_id         int32
	_app_name       string
	_app_version    string
	_date_created   int32
	_date_active    int32
	_ip             string
	_country        string
	_region         string
}

func (t *TL_authorization) Set_hash(_hash int64) {
	t._hash = _hash
}

func (t *TL_authorization) Get_hash() int64 {
	return t._hash
}

func (t *TL_authorization) Set_flags(_flags int32) {
	t._flags = _flags
}

func (t *TL_authorization) Get_flags() int32 {
	return t._flags
}

func (t *TL_authorization) Set_device_model(_device_model string) {
	t._device_model = _device_model
}

func (t *TL_authorization) Get_device_model() string {
	return t._device_model
}

func (t *TL_authorization) Set_platform(_platform string) {
	t._platform = _platform
}

func (t *TL_authorization) Get_platform() string {
	return t._platform
}

func (t *TL_authorization) Set_system_version(_system_version string) {
	t._system_version = _system_version
}

func (t *TL_authorization) Get_system_version() string {
	return t._system_version
}

func (t *TL_authorization) Set_api_id(_api_id int32) {
	t._api_id = _api_id
}

func (t *TL_authorization) Get_api_id() int32 {
	return t._api_id
}

func (t *TL_authorization) Set_app_name(_app_name string) {
	t._app_name = _app_name
}

func (t *TL_authorization) Get_app_name() string {
	return t._app_name
}

func (t *TL_authorization) Set_app_version(_app_version string) {
	t._app_version = _app_version
}

func (t *TL_authorization) Get_app_version() string {
	return t._app_version
}

func (t *TL_authorization) Set_date_created(_date_created int32) {
	t._date_created = _date_created
}

func (t *TL_authorization) Get_date_created() int32 {
	return t._date_created
}

func (t *TL_authorization) Set_date_active(_date_active int32) {
	t._date_active = _date_active
}

func (t *TL_authorization) Get_date_active() int32 {
	return t._date_active
}

func (t *TL_authorization) Set_ip(_ip string) {
	t._ip = _ip
}

func (t *TL_authorization) Get_ip() string {
	return t._ip
}

func (t *TL_authorization) Set_country(_country string) {
	t._country = _country
}

func (t *TL_authorization) Get_country() string {
	return t._country
}

func (t *TL_authorization) Set_region(_region string) {
	t._region = _region
}

func (t *TL_authorization) Get_region() string {
	return t._region
}

func New_TL_authorization() *TL_authorization {
	return &TL_authorization{}
}

func (t *TL_authorization) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_authorization))
	ec.Long(t.Get_hash())
	ec.Int(t.Get_flags())
	ec.String(t.Get_device_model())
	ec.String(t.Get_platform())
	ec.String(t.Get_system_version())
	ec.Int(t.Get_api_id())
	ec.String(t.Get_app_name())
	ec.String(t.Get_app_version())
	ec.Int(t.Get_date_created())
	ec.Int(t.Get_date_active())
	ec.String(t.Get_ip())
	ec.String(t.Get_country())
	ec.String(t.Get_region())

	return ec.GetBuffer()
}

func (t *TL_authorization) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._hash = dc.Long()
	t._flags = dc.Int()
	t._device_model = dc.String()
	t._platform = dc.String()
	t._system_version = dc.String()
	t._api_id = dc.Int()
	t._app_name = dc.String()
	t._app_version = dc.String()
	t._date_created = dc.Int()
	t._date_active = dc.Int()
	t._ip = dc.String()
	t._country = dc.String()
	t._region = dc.String()

}

// account_authorizations#1250abde
type TL_account_authorizations struct {
	_authorizations []byte
}

func (t *TL_account_authorizations) Set_authorizations(_authorizations []byte) {
	t._authorizations = _authorizations
}

func (t *TL_account_authorizations) Get_authorizations() []byte {
	return t._authorizations
}

func New_TL_account_authorizations() *TL_account_authorizations {
	return &TL_account_authorizations{}
}

func (t *TL_account_authorizations) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_account_authorizations))
	ec.Bytes(t.Get_authorizations())

	return ec.GetBuffer()
}

func (t *TL_account_authorizations) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._authorizations = dc.Bytes(16)

}

// account_noPassword#96dabc18
type TL_account_noPassword struct {
	_new_salt                  []byte
	_email_unconfirmed_pattern string
}

func (t *TL_account_noPassword) Set_new_salt(_new_salt []byte) {
	t._new_salt = _new_salt
}

func (t *TL_account_noPassword) Get_new_salt() []byte {
	return t._new_salt
}

func (t *TL_account_noPassword) Set_email_unconfirmed_pattern(_email_unconfirmed_pattern string) {
	t._email_unconfirmed_pattern = _email_unconfirmed_pattern
}

func (t *TL_account_noPassword) Get_email_unconfirmed_pattern() string {
	return t._email_unconfirmed_pattern
}

func New_TL_account_noPassword() *TL_account_noPassword {
	return &TL_account_noPassword{}
}

func (t *TL_account_noPassword) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_account_noPassword))
	ec.Bytes(t.Get_new_salt())
	ec.String(t.Get_email_unconfirmed_pattern())

	return ec.GetBuffer()
}

func (t *TL_account_noPassword) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._new_salt = dc.Bytes(16)
	t._email_unconfirmed_pattern = dc.String()

}

// account_password#7c18141c
type TL_account_password struct {
	_current_salt              []byte
	_new_salt                  []byte
	_hint                      string
	_has_recovery              bool
	_email_unconfirmed_pattern string
}

func (t *TL_account_password) Set_current_salt(_current_salt []byte) {
	t._current_salt = _current_salt
}

func (t *TL_account_password) Get_current_salt() []byte {
	return t._current_salt
}

func (t *TL_account_password) Set_new_salt(_new_salt []byte) {
	t._new_salt = _new_salt
}

func (t *TL_account_password) Get_new_salt() []byte {
	return t._new_salt
}

func (t *TL_account_password) Set_hint(_hint string) {
	t._hint = _hint
}

func (t *TL_account_password) Get_hint() string {
	return t._hint
}

func (t *TL_account_password) Set_has_recovery(_has_recovery bool) {
	t._has_recovery = _has_recovery
}

func (t *TL_account_password) Get_has_recovery() bool {
	return t._has_recovery
}

func (t *TL_account_password) Set_email_unconfirmed_pattern(_email_unconfirmed_pattern string) {
	t._email_unconfirmed_pattern = _email_unconfirmed_pattern
}

func (t *TL_account_password) Get_email_unconfirmed_pattern() string {
	return t._email_unconfirmed_pattern
}

func New_TL_account_password() *TL_account_password {
	return &TL_account_password{}
}

func (t *TL_account_password) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_account_password))
	ec.Bytes(t.Get_current_salt())
	ec.Bytes(t.Get_new_salt())
	ec.String(t.Get_hint())
	ec.Bool(t.Get_has_recovery())
	ec.String(t.Get_email_unconfirmed_pattern())

	return ec.GetBuffer()
}

func (t *TL_account_password) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._current_salt = dc.Bytes(16)
	t._new_salt = dc.Bytes(16)
	t._hint = dc.String()
	t._has_recovery = dc.Bool()
	t._email_unconfirmed_pattern = dc.String()

}

// account_passwordSettings#b7b72ab3
type TL_account_passwordSettings struct {
	_email string
}

func (t *TL_account_passwordSettings) Set_email(_email string) {
	t._email = _email
}

func (t *TL_account_passwordSettings) Get_email() string {
	return t._email
}

func New_TL_account_passwordSettings() *TL_account_passwordSettings {
	return &TL_account_passwordSettings{}
}

func (t *TL_account_passwordSettings) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_account_passwordSettings))
	ec.String(t.Get_email())

	return ec.GetBuffer()
}

func (t *TL_account_passwordSettings) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._email = dc.String()

}

// account_passwordInputSettings#86916deb
type TL_account_passwordInputSettings struct {
	_flags             []byte
	_new_salt          []byte
	_new_password_hash []byte
	_hint              []byte
	_email             []byte
}

func (t *TL_account_passwordInputSettings) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_account_passwordInputSettings) Get_flags() []byte {
	return t._flags
}

func (t *TL_account_passwordInputSettings) Set_new_salt(_new_salt []byte) {
	t._new_salt = _new_salt
}

func (t *TL_account_passwordInputSettings) Get_new_salt() []byte {
	return t._new_salt
}

func (t *TL_account_passwordInputSettings) Set_new_password_hash(_new_password_hash []byte) {
	t._new_password_hash = _new_password_hash
}

func (t *TL_account_passwordInputSettings) Get_new_password_hash() []byte {
	return t._new_password_hash
}

func (t *TL_account_passwordInputSettings) Set_hint(_hint []byte) {
	t._hint = _hint
}

func (t *TL_account_passwordInputSettings) Get_hint() []byte {
	return t._hint
}

func (t *TL_account_passwordInputSettings) Set_email(_email []byte) {
	t._email = _email
}

func (t *TL_account_passwordInputSettings) Get_email() []byte {
	return t._email
}

func New_TL_account_passwordInputSettings() *TL_account_passwordInputSettings {
	return &TL_account_passwordInputSettings{}
}

func (t *TL_account_passwordInputSettings) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_account_passwordInputSettings))
	ec.Bytes(t.Get_new_salt())
	ec.Bytes(t.Get_new_password_hash())
	ec.Bytes(t.Get_hint())
	ec.Bytes(t.Get_email())

	return ec.GetBuffer()
}

func (t *TL_account_passwordInputSettings) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._new_salt = dc.Bytes(16)
	t._new_password_hash = dc.Bytes(16)
	t._hint = dc.Bytes(16)
	t._email = dc.Bytes(16)

}

// auth_passwordRecovery#137948a5
type TL_auth_passwordRecovery struct {
	_email_pattern string
}

func (t *TL_auth_passwordRecovery) Set_email_pattern(_email_pattern string) {
	t._email_pattern = _email_pattern
}

func (t *TL_auth_passwordRecovery) Get_email_pattern() string {
	return t._email_pattern
}

func New_TL_auth_passwordRecovery() *TL_auth_passwordRecovery {
	return &TL_auth_passwordRecovery{}
}

func (t *TL_auth_passwordRecovery) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_auth_passwordRecovery))
	ec.String(t.Get_email_pattern())

	return ec.GetBuffer()
}

func (t *TL_auth_passwordRecovery) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._email_pattern = dc.String()

}

// receivedNotifyMessage#a384b779
type TL_receivedNotifyMessage struct {
	_id    int32
	_flags int32
}

func (t *TL_receivedNotifyMessage) Set_id(_id int32) {
	t._id = _id
}

func (t *TL_receivedNotifyMessage) Get_id() int32 {
	return t._id
}

func (t *TL_receivedNotifyMessage) Set_flags(_flags int32) {
	t._flags = _flags
}

func (t *TL_receivedNotifyMessage) Get_flags() int32 {
	return t._flags
}

func New_TL_receivedNotifyMessage() *TL_receivedNotifyMessage {
	return &TL_receivedNotifyMessage{}
}

func (t *TL_receivedNotifyMessage) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_receivedNotifyMessage))
	ec.Int(t.Get_id())
	ec.Int(t.Get_flags())

	return ec.GetBuffer()
}

func (t *TL_receivedNotifyMessage) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Int()
	t._flags = dc.Int()

}

// chatInviteEmpty#69df3769
type TL_chatInviteEmpty struct {
}

func New_TL_chatInviteEmpty() *TL_chatInviteEmpty {
	return &TL_chatInviteEmpty{}
}

func (t *TL_chatInviteEmpty) Encode() []byte {
	return nil
}

func (t *TL_chatInviteEmpty) Decode(b []byte) {

}

// chatInviteExported#fc2e05bc
type TL_chatInviteExported struct {
	_link string
}

func (t *TL_chatInviteExported) Set_link(_link string) {
	t._link = _link
}

func (t *TL_chatInviteExported) Get_link() string {
	return t._link
}

func New_TL_chatInviteExported() *TL_chatInviteExported {
	return &TL_chatInviteExported{}
}

func (t *TL_chatInviteExported) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_chatInviteExported))
	ec.String(t.Get_link())

	return ec.GetBuffer()
}

func (t *TL_chatInviteExported) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._link = dc.String()

}

// chatInviteAlready#5a686d7c
type TL_chatInviteAlready struct {
	_chat []byte
}

func (t *TL_chatInviteAlready) Set_chat(_chat []byte) {
	t._chat = _chat
}

func (t *TL_chatInviteAlready) Get_chat() []byte {
	return t._chat
}

func New_TL_chatInviteAlready() *TL_chatInviteAlready {
	return &TL_chatInviteAlready{}
}

func (t *TL_chatInviteAlready) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_chatInviteAlready))
	ec.Bytes(t.Get_chat())

	return ec.GetBuffer()
}

func (t *TL_chatInviteAlready) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._chat = dc.Bytes(16)

}

// chatInvite#db74f558
type TL_chatInvite struct {
	_flags              []byte
	_channel            []byte
	_broadcast          []byte
	_public             []byte
	_megagroup          []byte
	_title              string
	_photo              []byte
	_participants_count int32
	_participants       []byte
}

func (t *TL_chatInvite) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_chatInvite) Get_flags() []byte {
	return t._flags
}

func (t *TL_chatInvite) Set_channel(_channel []byte) {
	t._channel = _channel
}

func (t *TL_chatInvite) Get_channel() []byte {
	return t._channel
}

func (t *TL_chatInvite) Set_broadcast(_broadcast []byte) {
	t._broadcast = _broadcast
}

func (t *TL_chatInvite) Get_broadcast() []byte {
	return t._broadcast
}

func (t *TL_chatInvite) Set_public(_public []byte) {
	t._public = _public
}

func (t *TL_chatInvite) Get_public() []byte {
	return t._public
}

func (t *TL_chatInvite) Set_megagroup(_megagroup []byte) {
	t._megagroup = _megagroup
}

func (t *TL_chatInvite) Get_megagroup() []byte {
	return t._megagroup
}

func (t *TL_chatInvite) Set_title(_title string) {
	t._title = _title
}

func (t *TL_chatInvite) Get_title() string {
	return t._title
}

func (t *TL_chatInvite) Set_photo(_photo []byte) {
	t._photo = _photo
}

func (t *TL_chatInvite) Get_photo() []byte {
	return t._photo
}

func (t *TL_chatInvite) Set_participants_count(_participants_count int32) {
	t._participants_count = _participants_count
}

func (t *TL_chatInvite) Get_participants_count() int32 {
	return t._participants_count
}

func (t *TL_chatInvite) Set_participants(_participants []byte) {
	t._participants = _participants
}

func (t *TL_chatInvite) Get_participants() []byte {
	return t._participants
}

func New_TL_chatInvite() *TL_chatInvite {
	return &TL_chatInvite{}
}

func (t *TL_chatInvite) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_chatInvite))
	ec.Bytes(t.Get_channel())
	ec.Bytes(t.Get_broadcast())
	ec.Bytes(t.Get_public())
	ec.Bytes(t.Get_megagroup())
	ec.String(t.Get_title())
	ec.Bytes(t.Get_photo())
	ec.Int(t.Get_participants_count())
	ec.Bytes(t.Get_participants())

	return ec.GetBuffer()
}

func (t *TL_chatInvite) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._channel = dc.Bytes(16)
	t._broadcast = dc.Bytes(16)
	t._public = dc.Bytes(16)
	t._megagroup = dc.Bytes(16)
	t._title = dc.String()
	t._photo = dc.Bytes(16)
	t._participants_count = dc.Int()
	t._participants = dc.Bytes(16)

}

// inputStickerSetEmpty#ffb62b95
type TL_inputStickerSetEmpty struct {
}

func New_TL_inputStickerSetEmpty() *TL_inputStickerSetEmpty {
	return &TL_inputStickerSetEmpty{}
}

func (t *TL_inputStickerSetEmpty) Encode() []byte {
	return nil
}

func (t *TL_inputStickerSetEmpty) Decode(b []byte) {

}

// inputStickerSetID#9de7a269
type TL_inputStickerSetID struct {
	_id          int64
	_access_hash int64
}

func (t *TL_inputStickerSetID) Set_id(_id int64) {
	t._id = _id
}

func (t *TL_inputStickerSetID) Get_id() int64 {
	return t._id
}

func (t *TL_inputStickerSetID) Set_access_hash(_access_hash int64) {
	t._access_hash = _access_hash
}

func (t *TL_inputStickerSetID) Get_access_hash() int64 {
	return t._access_hash
}

func New_TL_inputStickerSetID() *TL_inputStickerSetID {
	return &TL_inputStickerSetID{}
}

func (t *TL_inputStickerSetID) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputStickerSetID))
	ec.Long(t.Get_id())
	ec.Long(t.Get_access_hash())

	return ec.GetBuffer()
}

func (t *TL_inputStickerSetID) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Long()
	t._access_hash = dc.Long()

}

// inputStickerSetShortName#861cc8a0
type TL_inputStickerSetShortName struct {
	_short_name string
}

func (t *TL_inputStickerSetShortName) Set_short_name(_short_name string) {
	t._short_name = _short_name
}

func (t *TL_inputStickerSetShortName) Get_short_name() string {
	return t._short_name
}

func New_TL_inputStickerSetShortName() *TL_inputStickerSetShortName {
	return &TL_inputStickerSetShortName{}
}

func (t *TL_inputStickerSetShortName) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputStickerSetShortName))
	ec.String(t.Get_short_name())

	return ec.GetBuffer()
}

func (t *TL_inputStickerSetShortName) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._short_name = dc.String()

}

// stickerSet#cd303b41
type TL_stickerSet struct {
	_flags       []byte
	_installed   []byte
	_archived    []byte
	_official    []byte
	_masks       []byte
	_id          int64
	_access_hash int64
	_title       string
	_short_name  string
	_count       int32
	_hash        int32
}

func (t *TL_stickerSet) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_stickerSet) Get_flags() []byte {
	return t._flags
}

func (t *TL_stickerSet) Set_installed(_installed []byte) {
	t._installed = _installed
}

func (t *TL_stickerSet) Get_installed() []byte {
	return t._installed
}

func (t *TL_stickerSet) Set_archived(_archived []byte) {
	t._archived = _archived
}

func (t *TL_stickerSet) Get_archived() []byte {
	return t._archived
}

func (t *TL_stickerSet) Set_official(_official []byte) {
	t._official = _official
}

func (t *TL_stickerSet) Get_official() []byte {
	return t._official
}

func (t *TL_stickerSet) Set_masks(_masks []byte) {
	t._masks = _masks
}

func (t *TL_stickerSet) Get_masks() []byte {
	return t._masks
}

func (t *TL_stickerSet) Set_id(_id int64) {
	t._id = _id
}

func (t *TL_stickerSet) Get_id() int64 {
	return t._id
}

func (t *TL_stickerSet) Set_access_hash(_access_hash int64) {
	t._access_hash = _access_hash
}

func (t *TL_stickerSet) Get_access_hash() int64 {
	return t._access_hash
}

func (t *TL_stickerSet) Set_title(_title string) {
	t._title = _title
}

func (t *TL_stickerSet) Get_title() string {
	return t._title
}

func (t *TL_stickerSet) Set_short_name(_short_name string) {
	t._short_name = _short_name
}

func (t *TL_stickerSet) Get_short_name() string {
	return t._short_name
}

func (t *TL_stickerSet) Set_count(_count int32) {
	t._count = _count
}

func (t *TL_stickerSet) Get_count() int32 {
	return t._count
}

func (t *TL_stickerSet) Set_hash(_hash int32) {
	t._hash = _hash
}

func (t *TL_stickerSet) Get_hash() int32 {
	return t._hash
}

func New_TL_stickerSet() *TL_stickerSet {
	return &TL_stickerSet{}
}

func (t *TL_stickerSet) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_stickerSet))
	ec.Bytes(t.Get_installed())
	ec.Bytes(t.Get_archived())
	ec.Bytes(t.Get_official())
	ec.Bytes(t.Get_masks())
	ec.Long(t.Get_id())
	ec.Long(t.Get_access_hash())
	ec.String(t.Get_title())
	ec.String(t.Get_short_name())
	ec.Int(t.Get_count())
	ec.Int(t.Get_hash())

	return ec.GetBuffer()
}

func (t *TL_stickerSet) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._installed = dc.Bytes(16)
	t._archived = dc.Bytes(16)
	t._official = dc.Bytes(16)
	t._masks = dc.Bytes(16)
	t._id = dc.Long()
	t._access_hash = dc.Long()
	t._title = dc.String()
	t._short_name = dc.String()
	t._count = dc.Int()
	t._hash = dc.Int()

}

// messages_stickerSet#b60a24a6
type TL_messages_stickerSet struct {
	_set       []byte
	_packs     []byte
	_documents []byte
}

func (t *TL_messages_stickerSet) Set_set(_set []byte) {
	t._set = _set
}

func (t *TL_messages_stickerSet) Get_set() []byte {
	return t._set
}

func (t *TL_messages_stickerSet) Set_packs(_packs []byte) {
	t._packs = _packs
}

func (t *TL_messages_stickerSet) Get_packs() []byte {
	return t._packs
}

func (t *TL_messages_stickerSet) Set_documents(_documents []byte) {
	t._documents = _documents
}

func (t *TL_messages_stickerSet) Get_documents() []byte {
	return t._documents
}

func New_TL_messages_stickerSet() *TL_messages_stickerSet {
	return &TL_messages_stickerSet{}
}

func (t *TL_messages_stickerSet) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_stickerSet))
	ec.Bytes(t.Get_set())
	ec.Bytes(t.Get_packs())
	ec.Bytes(t.Get_documents())

	return ec.GetBuffer()
}

func (t *TL_messages_stickerSet) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._set = dc.Bytes(16)
	t._packs = dc.Bytes(16)
	t._documents = dc.Bytes(16)

}

// botCommand#c27ac8c7
type TL_botCommand struct {
	_command     string
	_description string
}

func (t *TL_botCommand) Set_command(_command string) {
	t._command = _command
}

func (t *TL_botCommand) Get_command() string {
	return t._command
}

func (t *TL_botCommand) Set_description(_description string) {
	t._description = _description
}

func (t *TL_botCommand) Get_description() string {
	return t._description
}

func New_TL_botCommand() *TL_botCommand {
	return &TL_botCommand{}
}

func (t *TL_botCommand) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_botCommand))
	ec.String(t.Get_command())
	ec.String(t.Get_description())

	return ec.GetBuffer()
}

func (t *TL_botCommand) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._command = dc.String()
	t._description = dc.String()

}

// botInfo#98e81d3a
type TL_botInfo struct {
	_user_id     int32
	_description string
	_commands    []byte
}

func (t *TL_botInfo) Set_user_id(_user_id int32) {
	t._user_id = _user_id
}

func (t *TL_botInfo) Get_user_id() int32 {
	return t._user_id
}

func (t *TL_botInfo) Set_description(_description string) {
	t._description = _description
}

func (t *TL_botInfo) Get_description() string {
	return t._description
}

func (t *TL_botInfo) Set_commands(_commands []byte) {
	t._commands = _commands
}

func (t *TL_botInfo) Get_commands() []byte {
	return t._commands
}

func New_TL_botInfo() *TL_botInfo {
	return &TL_botInfo{}
}

func (t *TL_botInfo) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_botInfo))
	ec.Int(t.Get_user_id())
	ec.String(t.Get_description())
	ec.Bytes(t.Get_commands())

	return ec.GetBuffer()
}

func (t *TL_botInfo) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._user_id = dc.Int()
	t._description = dc.String()
	t._commands = dc.Bytes(16)

}

// keyboardButton#a2fa4880
type TL_keyboardButton struct {
	_text string
}

func (t *TL_keyboardButton) Set_text(_text string) {
	t._text = _text
}

func (t *TL_keyboardButton) Get_text() string {
	return t._text
}

func New_TL_keyboardButton() *TL_keyboardButton {
	return &TL_keyboardButton{}
}

func (t *TL_keyboardButton) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_keyboardButton))
	ec.String(t.Get_text())

	return ec.GetBuffer()
}

func (t *TL_keyboardButton) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._text = dc.String()

}

// keyboardButtonUrl#258aff05
type TL_keyboardButtonUrl struct {
	_text string
	_url  string
}

func (t *TL_keyboardButtonUrl) Set_text(_text string) {
	t._text = _text
}

func (t *TL_keyboardButtonUrl) Get_text() string {
	return t._text
}

func (t *TL_keyboardButtonUrl) Set_url(_url string) {
	t._url = _url
}

func (t *TL_keyboardButtonUrl) Get_url() string {
	return t._url
}

func New_TL_keyboardButtonUrl() *TL_keyboardButtonUrl {
	return &TL_keyboardButtonUrl{}
}

func (t *TL_keyboardButtonUrl) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_keyboardButtonUrl))
	ec.String(t.Get_text())
	ec.String(t.Get_url())

	return ec.GetBuffer()
}

func (t *TL_keyboardButtonUrl) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._text = dc.String()
	t._url = dc.String()

}

// keyboardButtonCallback#683a5e46
type TL_keyboardButtonCallback struct {
	_text string
	_data []byte
}

func (t *TL_keyboardButtonCallback) Set_text(_text string) {
	t._text = _text
}

func (t *TL_keyboardButtonCallback) Get_text() string {
	return t._text
}

func (t *TL_keyboardButtonCallback) Set_data(_data []byte) {
	t._data = _data
}

func (t *TL_keyboardButtonCallback) Get_data() []byte {
	return t._data
}

func New_TL_keyboardButtonCallback() *TL_keyboardButtonCallback {
	return &TL_keyboardButtonCallback{}
}

func (t *TL_keyboardButtonCallback) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_keyboardButtonCallback))
	ec.String(t.Get_text())
	ec.Bytes(t.Get_data())

	return ec.GetBuffer()
}

func (t *TL_keyboardButtonCallback) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._text = dc.String()
	t._data = dc.Bytes(16)

}

// keyboardButtonRequestPhone#b16a6c29
type TL_keyboardButtonRequestPhone struct {
	_text string
}

func (t *TL_keyboardButtonRequestPhone) Set_text(_text string) {
	t._text = _text
}

func (t *TL_keyboardButtonRequestPhone) Get_text() string {
	return t._text
}

func New_TL_keyboardButtonRequestPhone() *TL_keyboardButtonRequestPhone {
	return &TL_keyboardButtonRequestPhone{}
}

func (t *TL_keyboardButtonRequestPhone) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_keyboardButtonRequestPhone))
	ec.String(t.Get_text())

	return ec.GetBuffer()
}

func (t *TL_keyboardButtonRequestPhone) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._text = dc.String()

}

// keyboardButtonRequestGeoLocation#fc796b3f
type TL_keyboardButtonRequestGeoLocation struct {
	_text string
}

func (t *TL_keyboardButtonRequestGeoLocation) Set_text(_text string) {
	t._text = _text
}

func (t *TL_keyboardButtonRequestGeoLocation) Get_text() string {
	return t._text
}

func New_TL_keyboardButtonRequestGeoLocation() *TL_keyboardButtonRequestGeoLocation {
	return &TL_keyboardButtonRequestGeoLocation{}
}

func (t *TL_keyboardButtonRequestGeoLocation) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_keyboardButtonRequestGeoLocation))
	ec.String(t.Get_text())

	return ec.GetBuffer()
}

func (t *TL_keyboardButtonRequestGeoLocation) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._text = dc.String()

}

// keyboardButtonSwitchInline#0568a748
type TL_keyboardButtonSwitchInline struct {
	_flags     []byte
	_same_peer []byte
	_text      string
	_query     string
}

func (t *TL_keyboardButtonSwitchInline) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_keyboardButtonSwitchInline) Get_flags() []byte {
	return t._flags
}

func (t *TL_keyboardButtonSwitchInline) Set_same_peer(_same_peer []byte) {
	t._same_peer = _same_peer
}

func (t *TL_keyboardButtonSwitchInline) Get_same_peer() []byte {
	return t._same_peer
}

func (t *TL_keyboardButtonSwitchInline) Set_text(_text string) {
	t._text = _text
}

func (t *TL_keyboardButtonSwitchInline) Get_text() string {
	return t._text
}

func (t *TL_keyboardButtonSwitchInline) Set_query(_query string) {
	t._query = _query
}

func (t *TL_keyboardButtonSwitchInline) Get_query() string {
	return t._query
}

func New_TL_keyboardButtonSwitchInline() *TL_keyboardButtonSwitchInline {
	return &TL_keyboardButtonSwitchInline{}
}

func (t *TL_keyboardButtonSwitchInline) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_keyboardButtonSwitchInline))
	ec.Bytes(t.Get_same_peer())
	ec.String(t.Get_text())
	ec.String(t.Get_query())

	return ec.GetBuffer()
}

func (t *TL_keyboardButtonSwitchInline) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._same_peer = dc.Bytes(16)
	t._text = dc.String()
	t._query = dc.String()

}

// keyboardButtonGame#50f41ccf
type TL_keyboardButtonGame struct {
	_text string
}

func (t *TL_keyboardButtonGame) Set_text(_text string) {
	t._text = _text
}

func (t *TL_keyboardButtonGame) Get_text() string {
	return t._text
}

func New_TL_keyboardButtonGame() *TL_keyboardButtonGame {
	return &TL_keyboardButtonGame{}
}

func (t *TL_keyboardButtonGame) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_keyboardButtonGame))
	ec.String(t.Get_text())

	return ec.GetBuffer()
}

func (t *TL_keyboardButtonGame) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._text = dc.String()

}

// keyboardButtonBuy#afd93fbb
type TL_keyboardButtonBuy struct {
	_text string
}

func (t *TL_keyboardButtonBuy) Set_text(_text string) {
	t._text = _text
}

func (t *TL_keyboardButtonBuy) Get_text() string {
	return t._text
}

func New_TL_keyboardButtonBuy() *TL_keyboardButtonBuy {
	return &TL_keyboardButtonBuy{}
}

func (t *TL_keyboardButtonBuy) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_keyboardButtonBuy))
	ec.String(t.Get_text())

	return ec.GetBuffer()
}

func (t *TL_keyboardButtonBuy) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._text = dc.String()

}

// keyboardButtonRow#77608b83
type TL_keyboardButtonRow struct {
	_buttons []byte
}

func (t *TL_keyboardButtonRow) Set_buttons(_buttons []byte) {
	t._buttons = _buttons
}

func (t *TL_keyboardButtonRow) Get_buttons() []byte {
	return t._buttons
}

func New_TL_keyboardButtonRow() *TL_keyboardButtonRow {
	return &TL_keyboardButtonRow{}
}

func (t *TL_keyboardButtonRow) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_keyboardButtonRow))
	ec.Bytes(t.Get_buttons())

	return ec.GetBuffer()
}

func (t *TL_keyboardButtonRow) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._buttons = dc.Bytes(16)

}

// replyKeyboardHide#a03e5b85
type TL_replyKeyboardHide struct {
	_flags     []byte
	_selective []byte
}

func (t *TL_replyKeyboardHide) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_replyKeyboardHide) Get_flags() []byte {
	return t._flags
}

func (t *TL_replyKeyboardHide) Set_selective(_selective []byte) {
	t._selective = _selective
}

func (t *TL_replyKeyboardHide) Get_selective() []byte {
	return t._selective
}

func New_TL_replyKeyboardHide() *TL_replyKeyboardHide {
	return &TL_replyKeyboardHide{}
}

func (t *TL_replyKeyboardHide) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_replyKeyboardHide))
	ec.Bytes(t.Get_selective())

	return ec.GetBuffer()
}

func (t *TL_replyKeyboardHide) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._selective = dc.Bytes(16)

}

// replyKeyboardForceReply#f4108aa0
type TL_replyKeyboardForceReply struct {
	_flags      []byte
	_single_use []byte
	_selective  []byte
}

func (t *TL_replyKeyboardForceReply) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_replyKeyboardForceReply) Get_flags() []byte {
	return t._flags
}

func (t *TL_replyKeyboardForceReply) Set_single_use(_single_use []byte) {
	t._single_use = _single_use
}

func (t *TL_replyKeyboardForceReply) Get_single_use() []byte {
	return t._single_use
}

func (t *TL_replyKeyboardForceReply) Set_selective(_selective []byte) {
	t._selective = _selective
}

func (t *TL_replyKeyboardForceReply) Get_selective() []byte {
	return t._selective
}

func New_TL_replyKeyboardForceReply() *TL_replyKeyboardForceReply {
	return &TL_replyKeyboardForceReply{}
}

func (t *TL_replyKeyboardForceReply) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_replyKeyboardForceReply))
	ec.Bytes(t.Get_single_use())
	ec.Bytes(t.Get_selective())

	return ec.GetBuffer()
}

func (t *TL_replyKeyboardForceReply) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._single_use = dc.Bytes(16)
	t._selective = dc.Bytes(16)

}

// replyKeyboardMarkup#3502758c
type TL_replyKeyboardMarkup struct {
	_flags      []byte
	_resize     []byte
	_single_use []byte
	_selective  []byte
	_rows       []byte
}

func (t *TL_replyKeyboardMarkup) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_replyKeyboardMarkup) Get_flags() []byte {
	return t._flags
}

func (t *TL_replyKeyboardMarkup) Set_resize(_resize []byte) {
	t._resize = _resize
}

func (t *TL_replyKeyboardMarkup) Get_resize() []byte {
	return t._resize
}

func (t *TL_replyKeyboardMarkup) Set_single_use(_single_use []byte) {
	t._single_use = _single_use
}

func (t *TL_replyKeyboardMarkup) Get_single_use() []byte {
	return t._single_use
}

func (t *TL_replyKeyboardMarkup) Set_selective(_selective []byte) {
	t._selective = _selective
}

func (t *TL_replyKeyboardMarkup) Get_selective() []byte {
	return t._selective
}

func (t *TL_replyKeyboardMarkup) Set_rows(_rows []byte) {
	t._rows = _rows
}

func (t *TL_replyKeyboardMarkup) Get_rows() []byte {
	return t._rows
}

func New_TL_replyKeyboardMarkup() *TL_replyKeyboardMarkup {
	return &TL_replyKeyboardMarkup{}
}

func (t *TL_replyKeyboardMarkup) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_replyKeyboardMarkup))
	ec.Bytes(t.Get_resize())
	ec.Bytes(t.Get_single_use())
	ec.Bytes(t.Get_selective())
	ec.Bytes(t.Get_rows())

	return ec.GetBuffer()
}

func (t *TL_replyKeyboardMarkup) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._resize = dc.Bytes(16)
	t._single_use = dc.Bytes(16)
	t._selective = dc.Bytes(16)
	t._rows = dc.Bytes(16)

}

// replyInlineMarkup#48a30254
type TL_replyInlineMarkup struct {
	_rows []byte
}

func (t *TL_replyInlineMarkup) Set_rows(_rows []byte) {
	t._rows = _rows
}

func (t *TL_replyInlineMarkup) Get_rows() []byte {
	return t._rows
}

func New_TL_replyInlineMarkup() *TL_replyInlineMarkup {
	return &TL_replyInlineMarkup{}
}

func (t *TL_replyInlineMarkup) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_replyInlineMarkup))
	ec.Bytes(t.Get_rows())

	return ec.GetBuffer()
}

func (t *TL_replyInlineMarkup) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._rows = dc.Bytes(16)

}

// messageEntityUnknown#bb92ba95
type TL_messageEntityUnknown struct {
	_offset int32
	_length int32
}

func (t *TL_messageEntityUnknown) Set_offset(_offset int32) {
	t._offset = _offset
}

func (t *TL_messageEntityUnknown) Get_offset() int32 {
	return t._offset
}

func (t *TL_messageEntityUnknown) Set_length(_length int32) {
	t._length = _length
}

func (t *TL_messageEntityUnknown) Get_length() int32 {
	return t._length
}

func New_TL_messageEntityUnknown() *TL_messageEntityUnknown {
	return &TL_messageEntityUnknown{}
}

func (t *TL_messageEntityUnknown) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messageEntityUnknown))
	ec.Int(t.Get_offset())
	ec.Int(t.Get_length())

	return ec.GetBuffer()
}

func (t *TL_messageEntityUnknown) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._offset = dc.Int()
	t._length = dc.Int()

}

// messageEntityMention#fa04579d
type TL_messageEntityMention struct {
	_offset int32
	_length int32
}

func (t *TL_messageEntityMention) Set_offset(_offset int32) {
	t._offset = _offset
}

func (t *TL_messageEntityMention) Get_offset() int32 {
	return t._offset
}

func (t *TL_messageEntityMention) Set_length(_length int32) {
	t._length = _length
}

func (t *TL_messageEntityMention) Get_length() int32 {
	return t._length
}

func New_TL_messageEntityMention() *TL_messageEntityMention {
	return &TL_messageEntityMention{}
}

func (t *TL_messageEntityMention) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messageEntityMention))
	ec.Int(t.Get_offset())
	ec.Int(t.Get_length())

	return ec.GetBuffer()
}

func (t *TL_messageEntityMention) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._offset = dc.Int()
	t._length = dc.Int()

}

// messageEntityHashtag#6f635b0d
type TL_messageEntityHashtag struct {
	_offset int32
	_length int32
}

func (t *TL_messageEntityHashtag) Set_offset(_offset int32) {
	t._offset = _offset
}

func (t *TL_messageEntityHashtag) Get_offset() int32 {
	return t._offset
}

func (t *TL_messageEntityHashtag) Set_length(_length int32) {
	t._length = _length
}

func (t *TL_messageEntityHashtag) Get_length() int32 {
	return t._length
}

func New_TL_messageEntityHashtag() *TL_messageEntityHashtag {
	return &TL_messageEntityHashtag{}
}

func (t *TL_messageEntityHashtag) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messageEntityHashtag))
	ec.Int(t.Get_offset())
	ec.Int(t.Get_length())

	return ec.GetBuffer()
}

func (t *TL_messageEntityHashtag) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._offset = dc.Int()
	t._length = dc.Int()

}

// messageEntityBotCommand#6cef8ac7
type TL_messageEntityBotCommand struct {
	_offset int32
	_length int32
}

func (t *TL_messageEntityBotCommand) Set_offset(_offset int32) {
	t._offset = _offset
}

func (t *TL_messageEntityBotCommand) Get_offset() int32 {
	return t._offset
}

func (t *TL_messageEntityBotCommand) Set_length(_length int32) {
	t._length = _length
}

func (t *TL_messageEntityBotCommand) Get_length() int32 {
	return t._length
}

func New_TL_messageEntityBotCommand() *TL_messageEntityBotCommand {
	return &TL_messageEntityBotCommand{}
}

func (t *TL_messageEntityBotCommand) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messageEntityBotCommand))
	ec.Int(t.Get_offset())
	ec.Int(t.Get_length())

	return ec.GetBuffer()
}

func (t *TL_messageEntityBotCommand) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._offset = dc.Int()
	t._length = dc.Int()

}

// messageEntityUrl#6ed02538
type TL_messageEntityUrl struct {
	_offset int32
	_length int32
}

func (t *TL_messageEntityUrl) Set_offset(_offset int32) {
	t._offset = _offset
}

func (t *TL_messageEntityUrl) Get_offset() int32 {
	return t._offset
}

func (t *TL_messageEntityUrl) Set_length(_length int32) {
	t._length = _length
}

func (t *TL_messageEntityUrl) Get_length() int32 {
	return t._length
}

func New_TL_messageEntityUrl() *TL_messageEntityUrl {
	return &TL_messageEntityUrl{}
}

func (t *TL_messageEntityUrl) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messageEntityUrl))
	ec.Int(t.Get_offset())
	ec.Int(t.Get_length())

	return ec.GetBuffer()
}

func (t *TL_messageEntityUrl) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._offset = dc.Int()
	t._length = dc.Int()

}

// messageEntityEmail#64e475c2
type TL_messageEntityEmail struct {
	_offset int32
	_length int32
}

func (t *TL_messageEntityEmail) Set_offset(_offset int32) {
	t._offset = _offset
}

func (t *TL_messageEntityEmail) Get_offset() int32 {
	return t._offset
}

func (t *TL_messageEntityEmail) Set_length(_length int32) {
	t._length = _length
}

func (t *TL_messageEntityEmail) Get_length() int32 {
	return t._length
}

func New_TL_messageEntityEmail() *TL_messageEntityEmail {
	return &TL_messageEntityEmail{}
}

func (t *TL_messageEntityEmail) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messageEntityEmail))
	ec.Int(t.Get_offset())
	ec.Int(t.Get_length())

	return ec.GetBuffer()
}

func (t *TL_messageEntityEmail) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._offset = dc.Int()
	t._length = dc.Int()

}

// messageEntityBold#bd610bc9
type TL_messageEntityBold struct {
	_offset int32
	_length int32
}

func (t *TL_messageEntityBold) Set_offset(_offset int32) {
	t._offset = _offset
}

func (t *TL_messageEntityBold) Get_offset() int32 {
	return t._offset
}

func (t *TL_messageEntityBold) Set_length(_length int32) {
	t._length = _length
}

func (t *TL_messageEntityBold) Get_length() int32 {
	return t._length
}

func New_TL_messageEntityBold() *TL_messageEntityBold {
	return &TL_messageEntityBold{}
}

func (t *TL_messageEntityBold) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messageEntityBold))
	ec.Int(t.Get_offset())
	ec.Int(t.Get_length())

	return ec.GetBuffer()
}

func (t *TL_messageEntityBold) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._offset = dc.Int()
	t._length = dc.Int()

}

// messageEntityItalic#826f8b60
type TL_messageEntityItalic struct {
	_offset int32
	_length int32
}

func (t *TL_messageEntityItalic) Set_offset(_offset int32) {
	t._offset = _offset
}

func (t *TL_messageEntityItalic) Get_offset() int32 {
	return t._offset
}

func (t *TL_messageEntityItalic) Set_length(_length int32) {
	t._length = _length
}

func (t *TL_messageEntityItalic) Get_length() int32 {
	return t._length
}

func New_TL_messageEntityItalic() *TL_messageEntityItalic {
	return &TL_messageEntityItalic{}
}

func (t *TL_messageEntityItalic) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messageEntityItalic))
	ec.Int(t.Get_offset())
	ec.Int(t.Get_length())

	return ec.GetBuffer()
}

func (t *TL_messageEntityItalic) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._offset = dc.Int()
	t._length = dc.Int()

}

// messageEntityCode#28a20571
type TL_messageEntityCode struct {
	_offset int32
	_length int32
}

func (t *TL_messageEntityCode) Set_offset(_offset int32) {
	t._offset = _offset
}

func (t *TL_messageEntityCode) Get_offset() int32 {
	return t._offset
}

func (t *TL_messageEntityCode) Set_length(_length int32) {
	t._length = _length
}

func (t *TL_messageEntityCode) Get_length() int32 {
	return t._length
}

func New_TL_messageEntityCode() *TL_messageEntityCode {
	return &TL_messageEntityCode{}
}

func (t *TL_messageEntityCode) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messageEntityCode))
	ec.Int(t.Get_offset())
	ec.Int(t.Get_length())

	return ec.GetBuffer()
}

func (t *TL_messageEntityCode) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._offset = dc.Int()
	t._length = dc.Int()

}

// messageEntityPre#73924be0
type TL_messageEntityPre struct {
	_offset   int32
	_length   int32
	_language string
}

func (t *TL_messageEntityPre) Set_offset(_offset int32) {
	t._offset = _offset
}

func (t *TL_messageEntityPre) Get_offset() int32 {
	return t._offset
}

func (t *TL_messageEntityPre) Set_length(_length int32) {
	t._length = _length
}

func (t *TL_messageEntityPre) Get_length() int32 {
	return t._length
}

func (t *TL_messageEntityPre) Set_language(_language string) {
	t._language = _language
}

func (t *TL_messageEntityPre) Get_language() string {
	return t._language
}

func New_TL_messageEntityPre() *TL_messageEntityPre {
	return &TL_messageEntityPre{}
}

func (t *TL_messageEntityPre) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messageEntityPre))
	ec.Int(t.Get_offset())
	ec.Int(t.Get_length())
	ec.String(t.Get_language())

	return ec.GetBuffer()
}

func (t *TL_messageEntityPre) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._offset = dc.Int()
	t._length = dc.Int()
	t._language = dc.String()

}

// messageEntityTextUrl#76a6d327
type TL_messageEntityTextUrl struct {
	_offset int32
	_length int32
	_url    string
}

func (t *TL_messageEntityTextUrl) Set_offset(_offset int32) {
	t._offset = _offset
}

func (t *TL_messageEntityTextUrl) Get_offset() int32 {
	return t._offset
}

func (t *TL_messageEntityTextUrl) Set_length(_length int32) {
	t._length = _length
}

func (t *TL_messageEntityTextUrl) Get_length() int32 {
	return t._length
}

func (t *TL_messageEntityTextUrl) Set_url(_url string) {
	t._url = _url
}

func (t *TL_messageEntityTextUrl) Get_url() string {
	return t._url
}

func New_TL_messageEntityTextUrl() *TL_messageEntityTextUrl {
	return &TL_messageEntityTextUrl{}
}

func (t *TL_messageEntityTextUrl) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messageEntityTextUrl))
	ec.Int(t.Get_offset())
	ec.Int(t.Get_length())
	ec.String(t.Get_url())

	return ec.GetBuffer()
}

func (t *TL_messageEntityTextUrl) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._offset = dc.Int()
	t._length = dc.Int()
	t._url = dc.String()

}

// messageEntityMentionName#352dca58
type TL_messageEntityMentionName struct {
	_offset  int32
	_length  int32
	_user_id int32
}

func (t *TL_messageEntityMentionName) Set_offset(_offset int32) {
	t._offset = _offset
}

func (t *TL_messageEntityMentionName) Get_offset() int32 {
	return t._offset
}

func (t *TL_messageEntityMentionName) Set_length(_length int32) {
	t._length = _length
}

func (t *TL_messageEntityMentionName) Get_length() int32 {
	return t._length
}

func (t *TL_messageEntityMentionName) Set_user_id(_user_id int32) {
	t._user_id = _user_id
}

func (t *TL_messageEntityMentionName) Get_user_id() int32 {
	return t._user_id
}

func New_TL_messageEntityMentionName() *TL_messageEntityMentionName {
	return &TL_messageEntityMentionName{}
}

func (t *TL_messageEntityMentionName) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messageEntityMentionName))
	ec.Int(t.Get_offset())
	ec.Int(t.Get_length())
	ec.Int(t.Get_user_id())

	return ec.GetBuffer()
}

func (t *TL_messageEntityMentionName) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._offset = dc.Int()
	t._length = dc.Int()
	t._user_id = dc.Int()

}

// inputMessageEntityMentionName#208e68c9
type TL_inputMessageEntityMentionName struct {
	_offset  int32
	_length  int32
	_user_id []byte
}

func (t *TL_inputMessageEntityMentionName) Set_offset(_offset int32) {
	t._offset = _offset
}

func (t *TL_inputMessageEntityMentionName) Get_offset() int32 {
	return t._offset
}

func (t *TL_inputMessageEntityMentionName) Set_length(_length int32) {
	t._length = _length
}

func (t *TL_inputMessageEntityMentionName) Get_length() int32 {
	return t._length
}

func (t *TL_inputMessageEntityMentionName) Set_user_id(_user_id []byte) {
	t._user_id = _user_id
}

func (t *TL_inputMessageEntityMentionName) Get_user_id() []byte {
	return t._user_id
}

func New_TL_inputMessageEntityMentionName() *TL_inputMessageEntityMentionName {
	return &TL_inputMessageEntityMentionName{}
}

func (t *TL_inputMessageEntityMentionName) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputMessageEntityMentionName))
	ec.Int(t.Get_offset())
	ec.Int(t.Get_length())
	ec.Bytes(t.Get_user_id())

	return ec.GetBuffer()
}

func (t *TL_inputMessageEntityMentionName) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._offset = dc.Int()
	t._length = dc.Int()
	t._user_id = dc.Bytes(16)

}

// inputChannelEmpty#ee8c1e86
type TL_inputChannelEmpty struct {
}

func New_TL_inputChannelEmpty() *TL_inputChannelEmpty {
	return &TL_inputChannelEmpty{}
}

func (t *TL_inputChannelEmpty) Encode() []byte {
	return nil
}

func (t *TL_inputChannelEmpty) Decode(b []byte) {

}

// inputChannel#afeb712e
type TL_inputChannel struct {
	_channel_id  int32
	_access_hash int64
}

func (t *TL_inputChannel) Set_channel_id(_channel_id int32) {
	t._channel_id = _channel_id
}

func (t *TL_inputChannel) Get_channel_id() int32 {
	return t._channel_id
}

func (t *TL_inputChannel) Set_access_hash(_access_hash int64) {
	t._access_hash = _access_hash
}

func (t *TL_inputChannel) Get_access_hash() int64 {
	return t._access_hash
}

func New_TL_inputChannel() *TL_inputChannel {
	return &TL_inputChannel{}
}

func (t *TL_inputChannel) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputChannel))
	ec.Int(t.Get_channel_id())
	ec.Long(t.Get_access_hash())

	return ec.GetBuffer()
}

func (t *TL_inputChannel) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._channel_id = dc.Int()
	t._access_hash = dc.Long()

}

// contacts_resolvedPeer#7f077ad9
type TL_contacts_resolvedPeer struct {
	_peer  []byte
	_chats []byte
	_users []byte
}

func (t *TL_contacts_resolvedPeer) Set_peer(_peer []byte) {
	t._peer = _peer
}

func (t *TL_contacts_resolvedPeer) Get_peer() []byte {
	return t._peer
}

func (t *TL_contacts_resolvedPeer) Set_chats(_chats []byte) {
	t._chats = _chats
}

func (t *TL_contacts_resolvedPeer) Get_chats() []byte {
	return t._chats
}

func (t *TL_contacts_resolvedPeer) Set_users(_users []byte) {
	t._users = _users
}

func (t *TL_contacts_resolvedPeer) Get_users() []byte {
	return t._users
}

func New_TL_contacts_resolvedPeer() *TL_contacts_resolvedPeer {
	return &TL_contacts_resolvedPeer{}
}

func (t *TL_contacts_resolvedPeer) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_contacts_resolvedPeer))
	ec.Bytes(t.Get_peer())
	ec.Bytes(t.Get_chats())
	ec.Bytes(t.Get_users())

	return ec.GetBuffer()
}

func (t *TL_contacts_resolvedPeer) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._peer = dc.Bytes(16)
	t._chats = dc.Bytes(16)
	t._users = dc.Bytes(16)

}

// messageRange#0ae30253
type TL_messageRange struct {
	_min_id int32
	_max_id int32
}

func (t *TL_messageRange) Set_min_id(_min_id int32) {
	t._min_id = _min_id
}

func (t *TL_messageRange) Get_min_id() int32 {
	return t._min_id
}

func (t *TL_messageRange) Set_max_id(_max_id int32) {
	t._max_id = _max_id
}

func (t *TL_messageRange) Get_max_id() int32 {
	return t._max_id
}

func New_TL_messageRange() *TL_messageRange {
	return &TL_messageRange{}
}

func (t *TL_messageRange) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messageRange))
	ec.Int(t.Get_min_id())
	ec.Int(t.Get_max_id())

	return ec.GetBuffer()
}

func (t *TL_messageRange) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._min_id = dc.Int()
	t._max_id = dc.Int()

}

// updates_channelDifferenceEmpty#3e11affb
type TL_updates_channelDifferenceEmpty struct {
	_flags   []byte
	_final   []byte
	_pts     int32
	_timeout []byte
}

func (t *TL_updates_channelDifferenceEmpty) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_updates_channelDifferenceEmpty) Get_flags() []byte {
	return t._flags
}

func (t *TL_updates_channelDifferenceEmpty) Set_final(_final []byte) {
	t._final = _final
}

func (t *TL_updates_channelDifferenceEmpty) Get_final() []byte {
	return t._final
}

func (t *TL_updates_channelDifferenceEmpty) Set_pts(_pts int32) {
	t._pts = _pts
}

func (t *TL_updates_channelDifferenceEmpty) Get_pts() int32 {
	return t._pts
}

func (t *TL_updates_channelDifferenceEmpty) Set_timeout(_timeout []byte) {
	t._timeout = _timeout
}

func (t *TL_updates_channelDifferenceEmpty) Get_timeout() []byte {
	return t._timeout
}

func New_TL_updates_channelDifferenceEmpty() *TL_updates_channelDifferenceEmpty {
	return &TL_updates_channelDifferenceEmpty{}
}

func (t *TL_updates_channelDifferenceEmpty) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updates_channelDifferenceEmpty))
	ec.Bytes(t.Get_final())
	ec.Int(t.Get_pts())
	ec.Bytes(t.Get_timeout())

	return ec.GetBuffer()
}

func (t *TL_updates_channelDifferenceEmpty) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._final = dc.Bytes(16)
	t._pts = dc.Int()
	t._timeout = dc.Bytes(16)

}

// updates_channelDifferenceTooLong#6a9d7b35
type TL_updates_channelDifferenceTooLong struct {
	_flags                 []byte
	_final                 []byte
	_pts                   int32
	_timeout               []byte
	_top_message           int32
	_read_inbox_max_id     int32
	_read_outbox_max_id    int32
	_unread_count          int32
	_unread_mentions_count int32
	_messages              []byte
	_chats                 []byte
	_users                 []byte
}

func (t *TL_updates_channelDifferenceTooLong) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_updates_channelDifferenceTooLong) Get_flags() []byte {
	return t._flags
}

func (t *TL_updates_channelDifferenceTooLong) Set_final(_final []byte) {
	t._final = _final
}

func (t *TL_updates_channelDifferenceTooLong) Get_final() []byte {
	return t._final
}

func (t *TL_updates_channelDifferenceTooLong) Set_pts(_pts int32) {
	t._pts = _pts
}

func (t *TL_updates_channelDifferenceTooLong) Get_pts() int32 {
	return t._pts
}

func (t *TL_updates_channelDifferenceTooLong) Set_timeout(_timeout []byte) {
	t._timeout = _timeout
}

func (t *TL_updates_channelDifferenceTooLong) Get_timeout() []byte {
	return t._timeout
}

func (t *TL_updates_channelDifferenceTooLong) Set_top_message(_top_message int32) {
	t._top_message = _top_message
}

func (t *TL_updates_channelDifferenceTooLong) Get_top_message() int32 {
	return t._top_message
}

func (t *TL_updates_channelDifferenceTooLong) Set_read_inbox_max_id(_read_inbox_max_id int32) {
	t._read_inbox_max_id = _read_inbox_max_id
}

func (t *TL_updates_channelDifferenceTooLong) Get_read_inbox_max_id() int32 {
	return t._read_inbox_max_id
}

func (t *TL_updates_channelDifferenceTooLong) Set_read_outbox_max_id(_read_outbox_max_id int32) {
	t._read_outbox_max_id = _read_outbox_max_id
}

func (t *TL_updates_channelDifferenceTooLong) Get_read_outbox_max_id() int32 {
	return t._read_outbox_max_id
}

func (t *TL_updates_channelDifferenceTooLong) Set_unread_count(_unread_count int32) {
	t._unread_count = _unread_count
}

func (t *TL_updates_channelDifferenceTooLong) Get_unread_count() int32 {
	return t._unread_count
}

func (t *TL_updates_channelDifferenceTooLong) Set_unread_mentions_count(_unread_mentions_count int32) {
	t._unread_mentions_count = _unread_mentions_count
}

func (t *TL_updates_channelDifferenceTooLong) Get_unread_mentions_count() int32 {
	return t._unread_mentions_count
}

func (t *TL_updates_channelDifferenceTooLong) Set_messages(_messages []byte) {
	t._messages = _messages
}

func (t *TL_updates_channelDifferenceTooLong) Get_messages() []byte {
	return t._messages
}

func (t *TL_updates_channelDifferenceTooLong) Set_chats(_chats []byte) {
	t._chats = _chats
}

func (t *TL_updates_channelDifferenceTooLong) Get_chats() []byte {
	return t._chats
}

func (t *TL_updates_channelDifferenceTooLong) Set_users(_users []byte) {
	t._users = _users
}

func (t *TL_updates_channelDifferenceTooLong) Get_users() []byte {
	return t._users
}

func New_TL_updates_channelDifferenceTooLong() *TL_updates_channelDifferenceTooLong {
	return &TL_updates_channelDifferenceTooLong{}
}

func (t *TL_updates_channelDifferenceTooLong) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updates_channelDifferenceTooLong))
	ec.Bytes(t.Get_final())
	ec.Int(t.Get_pts())
	ec.Bytes(t.Get_timeout())
	ec.Int(t.Get_top_message())
	ec.Int(t.Get_read_inbox_max_id())
	ec.Int(t.Get_read_outbox_max_id())
	ec.Int(t.Get_unread_count())
	ec.Int(t.Get_unread_mentions_count())
	ec.Bytes(t.Get_messages())
	ec.Bytes(t.Get_chats())
	ec.Bytes(t.Get_users())

	return ec.GetBuffer()
}

func (t *TL_updates_channelDifferenceTooLong) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._final = dc.Bytes(16)
	t._pts = dc.Int()
	t._timeout = dc.Bytes(16)
	t._top_message = dc.Int()
	t._read_inbox_max_id = dc.Int()
	t._read_outbox_max_id = dc.Int()
	t._unread_count = dc.Int()
	t._unread_mentions_count = dc.Int()
	t._messages = dc.Bytes(16)
	t._chats = dc.Bytes(16)
	t._users = dc.Bytes(16)

}

// updates_channelDifference#2064674e
type TL_updates_channelDifference struct {
	_flags         []byte
	_final         []byte
	_pts           int32
	_timeout       []byte
	_new_messages  []byte
	_other_updates []byte
	_chats         []byte
	_users         []byte
}

func (t *TL_updates_channelDifference) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_updates_channelDifference) Get_flags() []byte {
	return t._flags
}

func (t *TL_updates_channelDifference) Set_final(_final []byte) {
	t._final = _final
}

func (t *TL_updates_channelDifference) Get_final() []byte {
	return t._final
}

func (t *TL_updates_channelDifference) Set_pts(_pts int32) {
	t._pts = _pts
}

func (t *TL_updates_channelDifference) Get_pts() int32 {
	return t._pts
}

func (t *TL_updates_channelDifference) Set_timeout(_timeout []byte) {
	t._timeout = _timeout
}

func (t *TL_updates_channelDifference) Get_timeout() []byte {
	return t._timeout
}

func (t *TL_updates_channelDifference) Set_new_messages(_new_messages []byte) {
	t._new_messages = _new_messages
}

func (t *TL_updates_channelDifference) Get_new_messages() []byte {
	return t._new_messages
}

func (t *TL_updates_channelDifference) Set_other_updates(_other_updates []byte) {
	t._other_updates = _other_updates
}

func (t *TL_updates_channelDifference) Get_other_updates() []byte {
	return t._other_updates
}

func (t *TL_updates_channelDifference) Set_chats(_chats []byte) {
	t._chats = _chats
}

func (t *TL_updates_channelDifference) Get_chats() []byte {
	return t._chats
}

func (t *TL_updates_channelDifference) Set_users(_users []byte) {
	t._users = _users
}

func (t *TL_updates_channelDifference) Get_users() []byte {
	return t._users
}

func New_TL_updates_channelDifference() *TL_updates_channelDifference {
	return &TL_updates_channelDifference{}
}

func (t *TL_updates_channelDifference) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updates_channelDifference))
	ec.Bytes(t.Get_final())
	ec.Int(t.Get_pts())
	ec.Bytes(t.Get_timeout())
	ec.Bytes(t.Get_new_messages())
	ec.Bytes(t.Get_other_updates())
	ec.Bytes(t.Get_chats())
	ec.Bytes(t.Get_users())

	return ec.GetBuffer()
}

func (t *TL_updates_channelDifference) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._final = dc.Bytes(16)
	t._pts = dc.Int()
	t._timeout = dc.Bytes(16)
	t._new_messages = dc.Bytes(16)
	t._other_updates = dc.Bytes(16)
	t._chats = dc.Bytes(16)
	t._users = dc.Bytes(16)

}

// channelMessagesFilterEmpty#94d42ee7
type TL_channelMessagesFilterEmpty struct {
}

func New_TL_channelMessagesFilterEmpty() *TL_channelMessagesFilterEmpty {
	return &TL_channelMessagesFilterEmpty{}
}

func (t *TL_channelMessagesFilterEmpty) Encode() []byte {
	return nil
}

func (t *TL_channelMessagesFilterEmpty) Decode(b []byte) {

}

// channelMessagesFilter#cd77d957
type TL_channelMessagesFilter struct {
	_flags                []byte
	_exclude_new_messages []byte
	_ranges               []byte
}

func (t *TL_channelMessagesFilter) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_channelMessagesFilter) Get_flags() []byte {
	return t._flags
}

func (t *TL_channelMessagesFilter) Set_exclude_new_messages(_exclude_new_messages []byte) {
	t._exclude_new_messages = _exclude_new_messages
}

func (t *TL_channelMessagesFilter) Get_exclude_new_messages() []byte {
	return t._exclude_new_messages
}

func (t *TL_channelMessagesFilter) Set_ranges(_ranges []byte) {
	t._ranges = _ranges
}

func (t *TL_channelMessagesFilter) Get_ranges() []byte {
	return t._ranges
}

func New_TL_channelMessagesFilter() *TL_channelMessagesFilter {
	return &TL_channelMessagesFilter{}
}

func (t *TL_channelMessagesFilter) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channelMessagesFilter))
	ec.Bytes(t.Get_exclude_new_messages())
	ec.Bytes(t.Get_ranges())

	return ec.GetBuffer()
}

func (t *TL_channelMessagesFilter) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._exclude_new_messages = dc.Bytes(16)
	t._ranges = dc.Bytes(16)

}

// channelParticipant#15ebac1d
type TL_channelParticipant struct {
	_user_id int32
	_date    int32
}

func (t *TL_channelParticipant) Set_user_id(_user_id int32) {
	t._user_id = _user_id
}

func (t *TL_channelParticipant) Get_user_id() int32 {
	return t._user_id
}

func (t *TL_channelParticipant) Set_date(_date int32) {
	t._date = _date
}

func (t *TL_channelParticipant) Get_date() int32 {
	return t._date
}

func New_TL_channelParticipant() *TL_channelParticipant {
	return &TL_channelParticipant{}
}

func (t *TL_channelParticipant) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channelParticipant))
	ec.Int(t.Get_user_id())
	ec.Int(t.Get_date())

	return ec.GetBuffer()
}

func (t *TL_channelParticipant) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._user_id = dc.Int()
	t._date = dc.Int()

}

// channelParticipantSelf#a3289a6d
type TL_channelParticipantSelf struct {
	_user_id    int32
	_inviter_id int32
	_date       int32
}

func (t *TL_channelParticipantSelf) Set_user_id(_user_id int32) {
	t._user_id = _user_id
}

func (t *TL_channelParticipantSelf) Get_user_id() int32 {
	return t._user_id
}

func (t *TL_channelParticipantSelf) Set_inviter_id(_inviter_id int32) {
	t._inviter_id = _inviter_id
}

func (t *TL_channelParticipantSelf) Get_inviter_id() int32 {
	return t._inviter_id
}

func (t *TL_channelParticipantSelf) Set_date(_date int32) {
	t._date = _date
}

func (t *TL_channelParticipantSelf) Get_date() int32 {
	return t._date
}

func New_TL_channelParticipantSelf() *TL_channelParticipantSelf {
	return &TL_channelParticipantSelf{}
}

func (t *TL_channelParticipantSelf) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channelParticipantSelf))
	ec.Int(t.Get_user_id())
	ec.Int(t.Get_inviter_id())
	ec.Int(t.Get_date())

	return ec.GetBuffer()
}

func (t *TL_channelParticipantSelf) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._user_id = dc.Int()
	t._inviter_id = dc.Int()
	t._date = dc.Int()

}

// channelParticipantCreator#e3e2e1f9
type TL_channelParticipantCreator struct {
	_user_id int32
}

func (t *TL_channelParticipantCreator) Set_user_id(_user_id int32) {
	t._user_id = _user_id
}

func (t *TL_channelParticipantCreator) Get_user_id() int32 {
	return t._user_id
}

func New_TL_channelParticipantCreator() *TL_channelParticipantCreator {
	return &TL_channelParticipantCreator{}
}

func (t *TL_channelParticipantCreator) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channelParticipantCreator))
	ec.Int(t.Get_user_id())

	return ec.GetBuffer()
}

func (t *TL_channelParticipantCreator) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._user_id = dc.Int()

}

// channelParticipantAdmin#a82fa898
type TL_channelParticipantAdmin struct {
	_flags        []byte
	_can_edit     []byte
	_user_id      int32
	_inviter_id   int32
	_promoted_by  int32
	_date         int32
	_admin_rights []byte
}

func (t *TL_channelParticipantAdmin) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_channelParticipantAdmin) Get_flags() []byte {
	return t._flags
}

func (t *TL_channelParticipantAdmin) Set_can_edit(_can_edit []byte) {
	t._can_edit = _can_edit
}

func (t *TL_channelParticipantAdmin) Get_can_edit() []byte {
	return t._can_edit
}

func (t *TL_channelParticipantAdmin) Set_user_id(_user_id int32) {
	t._user_id = _user_id
}

func (t *TL_channelParticipantAdmin) Get_user_id() int32 {
	return t._user_id
}

func (t *TL_channelParticipantAdmin) Set_inviter_id(_inviter_id int32) {
	t._inviter_id = _inviter_id
}

func (t *TL_channelParticipantAdmin) Get_inviter_id() int32 {
	return t._inviter_id
}

func (t *TL_channelParticipantAdmin) Set_promoted_by(_promoted_by int32) {
	t._promoted_by = _promoted_by
}

func (t *TL_channelParticipantAdmin) Get_promoted_by() int32 {
	return t._promoted_by
}

func (t *TL_channelParticipantAdmin) Set_date(_date int32) {
	t._date = _date
}

func (t *TL_channelParticipantAdmin) Get_date() int32 {
	return t._date
}

func (t *TL_channelParticipantAdmin) Set_admin_rights(_admin_rights []byte) {
	t._admin_rights = _admin_rights
}

func (t *TL_channelParticipantAdmin) Get_admin_rights() []byte {
	return t._admin_rights
}

func New_TL_channelParticipantAdmin() *TL_channelParticipantAdmin {
	return &TL_channelParticipantAdmin{}
}

func (t *TL_channelParticipantAdmin) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channelParticipantAdmin))
	ec.Bytes(t.Get_can_edit())
	ec.Int(t.Get_user_id())
	ec.Int(t.Get_inviter_id())
	ec.Int(t.Get_promoted_by())
	ec.Int(t.Get_date())
	ec.Bytes(t.Get_admin_rights())

	return ec.GetBuffer()
}

func (t *TL_channelParticipantAdmin) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._can_edit = dc.Bytes(16)
	t._user_id = dc.Int()
	t._inviter_id = dc.Int()
	t._promoted_by = dc.Int()
	t._date = dc.Int()
	t._admin_rights = dc.Bytes(16)

}

// channelParticipantBanned#222c1886
type TL_channelParticipantBanned struct {
	_flags         []byte
	_left          []byte
	_user_id       int32
	_kicked_by     int32
	_date          int32
	_banned_rights []byte
}

func (t *TL_channelParticipantBanned) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_channelParticipantBanned) Get_flags() []byte {
	return t._flags
}

func (t *TL_channelParticipantBanned) Set_left(_left []byte) {
	t._left = _left
}

func (t *TL_channelParticipantBanned) Get_left() []byte {
	return t._left
}

func (t *TL_channelParticipantBanned) Set_user_id(_user_id int32) {
	t._user_id = _user_id
}

func (t *TL_channelParticipantBanned) Get_user_id() int32 {
	return t._user_id
}

func (t *TL_channelParticipantBanned) Set_kicked_by(_kicked_by int32) {
	t._kicked_by = _kicked_by
}

func (t *TL_channelParticipantBanned) Get_kicked_by() int32 {
	return t._kicked_by
}

func (t *TL_channelParticipantBanned) Set_date(_date int32) {
	t._date = _date
}

func (t *TL_channelParticipantBanned) Get_date() int32 {
	return t._date
}

func (t *TL_channelParticipantBanned) Set_banned_rights(_banned_rights []byte) {
	t._banned_rights = _banned_rights
}

func (t *TL_channelParticipantBanned) Get_banned_rights() []byte {
	return t._banned_rights
}

func New_TL_channelParticipantBanned() *TL_channelParticipantBanned {
	return &TL_channelParticipantBanned{}
}

func (t *TL_channelParticipantBanned) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channelParticipantBanned))
	ec.Bytes(t.Get_left())
	ec.Int(t.Get_user_id())
	ec.Int(t.Get_kicked_by())
	ec.Int(t.Get_date())
	ec.Bytes(t.Get_banned_rights())

	return ec.GetBuffer()
}

func (t *TL_channelParticipantBanned) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._left = dc.Bytes(16)
	t._user_id = dc.Int()
	t._kicked_by = dc.Int()
	t._date = dc.Int()
	t._banned_rights = dc.Bytes(16)

}

// channelParticipantsRecent#de3f3c79
type TL_channelParticipantsRecent struct {
}

func New_TL_channelParticipantsRecent() *TL_channelParticipantsRecent {
	return &TL_channelParticipantsRecent{}
}

func (t *TL_channelParticipantsRecent) Encode() []byte {
	return nil
}

func (t *TL_channelParticipantsRecent) Decode(b []byte) {

}

// channelParticipantsAdmins#b4608969
type TL_channelParticipantsAdmins struct {
}

func New_TL_channelParticipantsAdmins() *TL_channelParticipantsAdmins {
	return &TL_channelParticipantsAdmins{}
}

func (t *TL_channelParticipantsAdmins) Encode() []byte {
	return nil
}

func (t *TL_channelParticipantsAdmins) Decode(b []byte) {

}

// channelParticipantsKicked#a3b54985
type TL_channelParticipantsKicked struct {
	_q string
}

func (t *TL_channelParticipantsKicked) Set_q(_q string) {
	t._q = _q
}

func (t *TL_channelParticipantsKicked) Get_q() string {
	return t._q
}

func New_TL_channelParticipantsKicked() *TL_channelParticipantsKicked {
	return &TL_channelParticipantsKicked{}
}

func (t *TL_channelParticipantsKicked) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channelParticipantsKicked))
	ec.String(t.Get_q())

	return ec.GetBuffer()
}

func (t *TL_channelParticipantsKicked) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._q = dc.String()

}

// channelParticipantsBots#b0d1865b
type TL_channelParticipantsBots struct {
}

func New_TL_channelParticipantsBots() *TL_channelParticipantsBots {
	return &TL_channelParticipantsBots{}
}

func (t *TL_channelParticipantsBots) Encode() []byte {
	return nil
}

func (t *TL_channelParticipantsBots) Decode(b []byte) {

}

// channelParticipantsBanned#1427a5e1
type TL_channelParticipantsBanned struct {
	_q string
}

func (t *TL_channelParticipantsBanned) Set_q(_q string) {
	t._q = _q
}

func (t *TL_channelParticipantsBanned) Get_q() string {
	return t._q
}

func New_TL_channelParticipantsBanned() *TL_channelParticipantsBanned {
	return &TL_channelParticipantsBanned{}
}

func (t *TL_channelParticipantsBanned) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channelParticipantsBanned))
	ec.String(t.Get_q())

	return ec.GetBuffer()
}

func (t *TL_channelParticipantsBanned) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._q = dc.String()

}

// channelParticipantsSearch#0656ac4b
type TL_channelParticipantsSearch struct {
	_q string
}

func (t *TL_channelParticipantsSearch) Set_q(_q string) {
	t._q = _q
}

func (t *TL_channelParticipantsSearch) Get_q() string {
	return t._q
}

func New_TL_channelParticipantsSearch() *TL_channelParticipantsSearch {
	return &TL_channelParticipantsSearch{}
}

func (t *TL_channelParticipantsSearch) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channelParticipantsSearch))
	ec.String(t.Get_q())

	return ec.GetBuffer()
}

func (t *TL_channelParticipantsSearch) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._q = dc.String()

}

// channels_channelParticipants#f56ee2a8
type TL_channels_channelParticipants struct {
	_count        int32
	_participants []byte
	_users        []byte
}

func (t *TL_channels_channelParticipants) Set_count(_count int32) {
	t._count = _count
}

func (t *TL_channels_channelParticipants) Get_count() int32 {
	return t._count
}

func (t *TL_channels_channelParticipants) Set_participants(_participants []byte) {
	t._participants = _participants
}

func (t *TL_channels_channelParticipants) Get_participants() []byte {
	return t._participants
}

func (t *TL_channels_channelParticipants) Set_users(_users []byte) {
	t._users = _users
}

func (t *TL_channels_channelParticipants) Get_users() []byte {
	return t._users
}

func New_TL_channels_channelParticipants() *TL_channels_channelParticipants {
	return &TL_channels_channelParticipants{}
}

func (t *TL_channels_channelParticipants) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channels_channelParticipants))
	ec.Int(t.Get_count())
	ec.Bytes(t.Get_participants())
	ec.Bytes(t.Get_users())

	return ec.GetBuffer()
}

func (t *TL_channels_channelParticipants) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._count = dc.Int()
	t._participants = dc.Bytes(16)
	t._users = dc.Bytes(16)

}

// channels_channelParticipantsNotModified#f0173fe9
type TL_channels_channelParticipantsNotModified struct {
}

func New_TL_channels_channelParticipantsNotModified() *TL_channels_channelParticipantsNotModified {
	return &TL_channels_channelParticipantsNotModified{}
}

func (t *TL_channels_channelParticipantsNotModified) Encode() []byte {
	return nil
}

func (t *TL_channels_channelParticipantsNotModified) Decode(b []byte) {

}

// channels_channelParticipant#d0d9b163
type TL_channels_channelParticipant struct {
	_participant []byte
	_users       []byte
}

func (t *TL_channels_channelParticipant) Set_participant(_participant []byte) {
	t._participant = _participant
}

func (t *TL_channels_channelParticipant) Get_participant() []byte {
	return t._participant
}

func (t *TL_channels_channelParticipant) Set_users(_users []byte) {
	t._users = _users
}

func (t *TL_channels_channelParticipant) Get_users() []byte {
	return t._users
}

func New_TL_channels_channelParticipant() *TL_channels_channelParticipant {
	return &TL_channels_channelParticipant{}
}

func (t *TL_channels_channelParticipant) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channels_channelParticipant))
	ec.Bytes(t.Get_participant())
	ec.Bytes(t.Get_users())

	return ec.GetBuffer()
}

func (t *TL_channels_channelParticipant) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._participant = dc.Bytes(16)
	t._users = dc.Bytes(16)

}

// help_termsOfService#f1ee3e90
type TL_help_termsOfService struct {
	_text string
}

func (t *TL_help_termsOfService) Set_text(_text string) {
	t._text = _text
}

func (t *TL_help_termsOfService) Get_text() string {
	return t._text
}

func New_TL_help_termsOfService() *TL_help_termsOfService {
	return &TL_help_termsOfService{}
}

func (t *TL_help_termsOfService) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_help_termsOfService))
	ec.String(t.Get_text())

	return ec.GetBuffer()
}

func (t *TL_help_termsOfService) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._text = dc.String()

}

// foundGif#162ecc1f
type TL_foundGif struct {
	_url          string
	_thumb_url    string
	_content_url  string
	_content_type string
	_w            int32
	_h            int32
}

func (t *TL_foundGif) Set_url(_url string) {
	t._url = _url
}

func (t *TL_foundGif) Get_url() string {
	return t._url
}

func (t *TL_foundGif) Set_thumb_url(_thumb_url string) {
	t._thumb_url = _thumb_url
}

func (t *TL_foundGif) Get_thumb_url() string {
	return t._thumb_url
}

func (t *TL_foundGif) Set_content_url(_content_url string) {
	t._content_url = _content_url
}

func (t *TL_foundGif) Get_content_url() string {
	return t._content_url
}

func (t *TL_foundGif) Set_content_type(_content_type string) {
	t._content_type = _content_type
}

func (t *TL_foundGif) Get_content_type() string {
	return t._content_type
}

func (t *TL_foundGif) Set_w(_w int32) {
	t._w = _w
}

func (t *TL_foundGif) Get_w() int32 {
	return t._w
}

func (t *TL_foundGif) Set_h(_h int32) {
	t._h = _h
}

func (t *TL_foundGif) Get_h() int32 {
	return t._h
}

func New_TL_foundGif() *TL_foundGif {
	return &TL_foundGif{}
}

func (t *TL_foundGif) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_foundGif))
	ec.String(t.Get_url())
	ec.String(t.Get_thumb_url())
	ec.String(t.Get_content_url())
	ec.String(t.Get_content_type())
	ec.Int(t.Get_w())
	ec.Int(t.Get_h())

	return ec.GetBuffer()
}

func (t *TL_foundGif) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._url = dc.String()
	t._thumb_url = dc.String()
	t._content_url = dc.String()
	t._content_type = dc.String()
	t._w = dc.Int()
	t._h = dc.Int()

}

// foundGifCached#9c750409
type TL_foundGifCached struct {
	_url      string
	_photo    []byte
	_document []byte
}

func (t *TL_foundGifCached) Set_url(_url string) {
	t._url = _url
}

func (t *TL_foundGifCached) Get_url() string {
	return t._url
}

func (t *TL_foundGifCached) Set_photo(_photo []byte) {
	t._photo = _photo
}

func (t *TL_foundGifCached) Get_photo() []byte {
	return t._photo
}

func (t *TL_foundGifCached) Set_document(_document []byte) {
	t._document = _document
}

func (t *TL_foundGifCached) Get_document() []byte {
	return t._document
}

func New_TL_foundGifCached() *TL_foundGifCached {
	return &TL_foundGifCached{}
}

func (t *TL_foundGifCached) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_foundGifCached))
	ec.String(t.Get_url())
	ec.Bytes(t.Get_photo())
	ec.Bytes(t.Get_document())

	return ec.GetBuffer()
}

func (t *TL_foundGifCached) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._url = dc.String()
	t._photo = dc.Bytes(16)
	t._document = dc.Bytes(16)

}

// messages_foundGifs#450a1c0a
type TL_messages_foundGifs struct {
	_next_offset int32
	_results     []byte
}

func (t *TL_messages_foundGifs) Set_next_offset(_next_offset int32) {
	t._next_offset = _next_offset
}

func (t *TL_messages_foundGifs) Get_next_offset() int32 {
	return t._next_offset
}

func (t *TL_messages_foundGifs) Set_results(_results []byte) {
	t._results = _results
}

func (t *TL_messages_foundGifs) Get_results() []byte {
	return t._results
}

func New_TL_messages_foundGifs() *TL_messages_foundGifs {
	return &TL_messages_foundGifs{}
}

func (t *TL_messages_foundGifs) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_foundGifs))
	ec.Int(t.Get_next_offset())
	ec.Bytes(t.Get_results())

	return ec.GetBuffer()
}

func (t *TL_messages_foundGifs) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._next_offset = dc.Int()
	t._results = dc.Bytes(16)

}

// messages_savedGifsNotModified#e8025ca2
type TL_messages_savedGifsNotModified struct {
}

func New_TL_messages_savedGifsNotModified() *TL_messages_savedGifsNotModified {
	return &TL_messages_savedGifsNotModified{}
}

func (t *TL_messages_savedGifsNotModified) Encode() []byte {
	return nil
}

func (t *TL_messages_savedGifsNotModified) Decode(b []byte) {

}

// messages_savedGifs#2e0709a5
type TL_messages_savedGifs struct {
	_hash int32
	_gifs []byte
}

func (t *TL_messages_savedGifs) Set_hash(_hash int32) {
	t._hash = _hash
}

func (t *TL_messages_savedGifs) Get_hash() int32 {
	return t._hash
}

func (t *TL_messages_savedGifs) Set_gifs(_gifs []byte) {
	t._gifs = _gifs
}

func (t *TL_messages_savedGifs) Get_gifs() []byte {
	return t._gifs
}

func New_TL_messages_savedGifs() *TL_messages_savedGifs {
	return &TL_messages_savedGifs{}
}

func (t *TL_messages_savedGifs) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_savedGifs))
	ec.Int(t.Get_hash())
	ec.Bytes(t.Get_gifs())

	return ec.GetBuffer()
}

func (t *TL_messages_savedGifs) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._hash = dc.Int()
	t._gifs = dc.Bytes(16)

}

// inputBotInlineMessageMediaAuto#292fed13
type TL_inputBotInlineMessageMediaAuto struct {
	_flags        []byte
	_caption      string
	_reply_markup []byte
}

func (t *TL_inputBotInlineMessageMediaAuto) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_inputBotInlineMessageMediaAuto) Get_flags() []byte {
	return t._flags
}

func (t *TL_inputBotInlineMessageMediaAuto) Set_caption(_caption string) {
	t._caption = _caption
}

func (t *TL_inputBotInlineMessageMediaAuto) Get_caption() string {
	return t._caption
}

func (t *TL_inputBotInlineMessageMediaAuto) Set_reply_markup(_reply_markup []byte) {
	t._reply_markup = _reply_markup
}

func (t *TL_inputBotInlineMessageMediaAuto) Get_reply_markup() []byte {
	return t._reply_markup
}

func New_TL_inputBotInlineMessageMediaAuto() *TL_inputBotInlineMessageMediaAuto {
	return &TL_inputBotInlineMessageMediaAuto{}
}

func (t *TL_inputBotInlineMessageMediaAuto) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputBotInlineMessageMediaAuto))
	ec.String(t.Get_caption())
	ec.Bytes(t.Get_reply_markup())

	return ec.GetBuffer()
}

func (t *TL_inputBotInlineMessageMediaAuto) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._caption = dc.String()
	t._reply_markup = dc.Bytes(16)

}

// inputBotInlineMessageText#3dcd7a87
type TL_inputBotInlineMessageText struct {
	_flags        []byte
	_no_webpage   []byte
	_message      string
	_entities     []byte
	_reply_markup []byte
}

func (t *TL_inputBotInlineMessageText) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_inputBotInlineMessageText) Get_flags() []byte {
	return t._flags
}

func (t *TL_inputBotInlineMessageText) Set_no_webpage(_no_webpage []byte) {
	t._no_webpage = _no_webpage
}

func (t *TL_inputBotInlineMessageText) Get_no_webpage() []byte {
	return t._no_webpage
}

func (t *TL_inputBotInlineMessageText) Set_message(_message string) {
	t._message = _message
}

func (t *TL_inputBotInlineMessageText) Get_message() string {
	return t._message
}

func (t *TL_inputBotInlineMessageText) Set_entities(_entities []byte) {
	t._entities = _entities
}

func (t *TL_inputBotInlineMessageText) Get_entities() []byte {
	return t._entities
}

func (t *TL_inputBotInlineMessageText) Set_reply_markup(_reply_markup []byte) {
	t._reply_markup = _reply_markup
}

func (t *TL_inputBotInlineMessageText) Get_reply_markup() []byte {
	return t._reply_markup
}

func New_TL_inputBotInlineMessageText() *TL_inputBotInlineMessageText {
	return &TL_inputBotInlineMessageText{}
}

func (t *TL_inputBotInlineMessageText) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputBotInlineMessageText))
	ec.Bytes(t.Get_no_webpage())
	ec.String(t.Get_message())
	ec.Bytes(t.Get_entities())
	ec.Bytes(t.Get_reply_markup())

	return ec.GetBuffer()
}

func (t *TL_inputBotInlineMessageText) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._no_webpage = dc.Bytes(16)
	t._message = dc.String()
	t._entities = dc.Bytes(16)
	t._reply_markup = dc.Bytes(16)

}

// inputBotInlineMessageMediaGeo#c1b15d65
type TL_inputBotInlineMessageMediaGeo struct {
	_flags        []byte
	_geo_point    []byte
	_period       int32
	_reply_markup []byte
}

func (t *TL_inputBotInlineMessageMediaGeo) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_inputBotInlineMessageMediaGeo) Get_flags() []byte {
	return t._flags
}

func (t *TL_inputBotInlineMessageMediaGeo) Set_geo_point(_geo_point []byte) {
	t._geo_point = _geo_point
}

func (t *TL_inputBotInlineMessageMediaGeo) Get_geo_point() []byte {
	return t._geo_point
}

func (t *TL_inputBotInlineMessageMediaGeo) Set_period(_period int32) {
	t._period = _period
}

func (t *TL_inputBotInlineMessageMediaGeo) Get_period() int32 {
	return t._period
}

func (t *TL_inputBotInlineMessageMediaGeo) Set_reply_markup(_reply_markup []byte) {
	t._reply_markup = _reply_markup
}

func (t *TL_inputBotInlineMessageMediaGeo) Get_reply_markup() []byte {
	return t._reply_markup
}

func New_TL_inputBotInlineMessageMediaGeo() *TL_inputBotInlineMessageMediaGeo {
	return &TL_inputBotInlineMessageMediaGeo{}
}

func (t *TL_inputBotInlineMessageMediaGeo) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputBotInlineMessageMediaGeo))
	ec.Bytes(t.Get_geo_point())
	ec.Int(t.Get_period())
	ec.Bytes(t.Get_reply_markup())

	return ec.GetBuffer()
}

func (t *TL_inputBotInlineMessageMediaGeo) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._geo_point = dc.Bytes(16)
	t._period = dc.Int()
	t._reply_markup = dc.Bytes(16)

}

// inputBotInlineMessageMediaVenue#aaafadc8
type TL_inputBotInlineMessageMediaVenue struct {
	_flags        []byte
	_geo_point    []byte
	_title        string
	_address      string
	_provider     string
	_venue_id     string
	_reply_markup []byte
}

func (t *TL_inputBotInlineMessageMediaVenue) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_inputBotInlineMessageMediaVenue) Get_flags() []byte {
	return t._flags
}

func (t *TL_inputBotInlineMessageMediaVenue) Set_geo_point(_geo_point []byte) {
	t._geo_point = _geo_point
}

func (t *TL_inputBotInlineMessageMediaVenue) Get_geo_point() []byte {
	return t._geo_point
}

func (t *TL_inputBotInlineMessageMediaVenue) Set_title(_title string) {
	t._title = _title
}

func (t *TL_inputBotInlineMessageMediaVenue) Get_title() string {
	return t._title
}

func (t *TL_inputBotInlineMessageMediaVenue) Set_address(_address string) {
	t._address = _address
}

func (t *TL_inputBotInlineMessageMediaVenue) Get_address() string {
	return t._address
}

func (t *TL_inputBotInlineMessageMediaVenue) Set_provider(_provider string) {
	t._provider = _provider
}

func (t *TL_inputBotInlineMessageMediaVenue) Get_provider() string {
	return t._provider
}

func (t *TL_inputBotInlineMessageMediaVenue) Set_venue_id(_venue_id string) {
	t._venue_id = _venue_id
}

func (t *TL_inputBotInlineMessageMediaVenue) Get_venue_id() string {
	return t._venue_id
}

func (t *TL_inputBotInlineMessageMediaVenue) Set_reply_markup(_reply_markup []byte) {
	t._reply_markup = _reply_markup
}

func (t *TL_inputBotInlineMessageMediaVenue) Get_reply_markup() []byte {
	return t._reply_markup
}

func New_TL_inputBotInlineMessageMediaVenue() *TL_inputBotInlineMessageMediaVenue {
	return &TL_inputBotInlineMessageMediaVenue{}
}

func (t *TL_inputBotInlineMessageMediaVenue) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputBotInlineMessageMediaVenue))
	ec.Bytes(t.Get_geo_point())
	ec.String(t.Get_title())
	ec.String(t.Get_address())
	ec.String(t.Get_provider())
	ec.String(t.Get_venue_id())
	ec.Bytes(t.Get_reply_markup())

	return ec.GetBuffer()
}

func (t *TL_inputBotInlineMessageMediaVenue) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._geo_point = dc.Bytes(16)
	t._title = dc.String()
	t._address = dc.String()
	t._provider = dc.String()
	t._venue_id = dc.String()
	t._reply_markup = dc.Bytes(16)

}

// inputBotInlineMessageMediaContact#2daf01a7
type TL_inputBotInlineMessageMediaContact struct {
	_flags        []byte
	_phone_number string
	_first_name   string
	_last_name    string
	_reply_markup []byte
}

func (t *TL_inputBotInlineMessageMediaContact) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_inputBotInlineMessageMediaContact) Get_flags() []byte {
	return t._flags
}

func (t *TL_inputBotInlineMessageMediaContact) Set_phone_number(_phone_number string) {
	t._phone_number = _phone_number
}

func (t *TL_inputBotInlineMessageMediaContact) Get_phone_number() string {
	return t._phone_number
}

func (t *TL_inputBotInlineMessageMediaContact) Set_first_name(_first_name string) {
	t._first_name = _first_name
}

func (t *TL_inputBotInlineMessageMediaContact) Get_first_name() string {
	return t._first_name
}

func (t *TL_inputBotInlineMessageMediaContact) Set_last_name(_last_name string) {
	t._last_name = _last_name
}

func (t *TL_inputBotInlineMessageMediaContact) Get_last_name() string {
	return t._last_name
}

func (t *TL_inputBotInlineMessageMediaContact) Set_reply_markup(_reply_markup []byte) {
	t._reply_markup = _reply_markup
}

func (t *TL_inputBotInlineMessageMediaContact) Get_reply_markup() []byte {
	return t._reply_markup
}

func New_TL_inputBotInlineMessageMediaContact() *TL_inputBotInlineMessageMediaContact {
	return &TL_inputBotInlineMessageMediaContact{}
}

func (t *TL_inputBotInlineMessageMediaContact) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputBotInlineMessageMediaContact))
	ec.String(t.Get_phone_number())
	ec.String(t.Get_first_name())
	ec.String(t.Get_last_name())
	ec.Bytes(t.Get_reply_markup())

	return ec.GetBuffer()
}

func (t *TL_inputBotInlineMessageMediaContact) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._phone_number = dc.String()
	t._first_name = dc.String()
	t._last_name = dc.String()
	t._reply_markup = dc.Bytes(16)

}

// inputBotInlineMessageGame#4b425864
type TL_inputBotInlineMessageGame struct {
	_flags        []byte
	_reply_markup []byte
}

func (t *TL_inputBotInlineMessageGame) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_inputBotInlineMessageGame) Get_flags() []byte {
	return t._flags
}

func (t *TL_inputBotInlineMessageGame) Set_reply_markup(_reply_markup []byte) {
	t._reply_markup = _reply_markup
}

func (t *TL_inputBotInlineMessageGame) Get_reply_markup() []byte {
	return t._reply_markup
}

func New_TL_inputBotInlineMessageGame() *TL_inputBotInlineMessageGame {
	return &TL_inputBotInlineMessageGame{}
}

func (t *TL_inputBotInlineMessageGame) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputBotInlineMessageGame))
	ec.Bytes(t.Get_reply_markup())

	return ec.GetBuffer()
}

func (t *TL_inputBotInlineMessageGame) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._reply_markup = dc.Bytes(16)

}

// inputBotInlineResult#2cbbe15a
type TL_inputBotInlineResult struct {
	_flags        []byte
	_id           string
	_type         string
	_title        []byte
	_description  []byte
	_url          []byte
	_thumb_url    []byte
	_content_url  []byte
	_content_type []byte
	_w            []byte
	_h            []byte
	_duration     []byte
	_send_message []byte
}

func (t *TL_inputBotInlineResult) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_inputBotInlineResult) Get_flags() []byte {
	return t._flags
}

func (t *TL_inputBotInlineResult) Set_id(_id string) {
	t._id = _id
}

func (t *TL_inputBotInlineResult) Get_id() string {
	return t._id
}

func (t *TL_inputBotInlineResult) Set_type(_type string) {
	t._type = _type
}

func (t *TL_inputBotInlineResult) Get_type() string {
	return t._type
}

func (t *TL_inputBotInlineResult) Set_title(_title []byte) {
	t._title = _title
}

func (t *TL_inputBotInlineResult) Get_title() []byte {
	return t._title
}

func (t *TL_inputBotInlineResult) Set_description(_description []byte) {
	t._description = _description
}

func (t *TL_inputBotInlineResult) Get_description() []byte {
	return t._description
}

func (t *TL_inputBotInlineResult) Set_url(_url []byte) {
	t._url = _url
}

func (t *TL_inputBotInlineResult) Get_url() []byte {
	return t._url
}

func (t *TL_inputBotInlineResult) Set_thumb_url(_thumb_url []byte) {
	t._thumb_url = _thumb_url
}

func (t *TL_inputBotInlineResult) Get_thumb_url() []byte {
	return t._thumb_url
}

func (t *TL_inputBotInlineResult) Set_content_url(_content_url []byte) {
	t._content_url = _content_url
}

func (t *TL_inputBotInlineResult) Get_content_url() []byte {
	return t._content_url
}

func (t *TL_inputBotInlineResult) Set_content_type(_content_type []byte) {
	t._content_type = _content_type
}

func (t *TL_inputBotInlineResult) Get_content_type() []byte {
	return t._content_type
}

func (t *TL_inputBotInlineResult) Set_w(_w []byte) {
	t._w = _w
}

func (t *TL_inputBotInlineResult) Get_w() []byte {
	return t._w
}

func (t *TL_inputBotInlineResult) Set_h(_h []byte) {
	t._h = _h
}

func (t *TL_inputBotInlineResult) Get_h() []byte {
	return t._h
}

func (t *TL_inputBotInlineResult) Set_duration(_duration []byte) {
	t._duration = _duration
}

func (t *TL_inputBotInlineResult) Get_duration() []byte {
	return t._duration
}

func (t *TL_inputBotInlineResult) Set_send_message(_send_message []byte) {
	t._send_message = _send_message
}

func (t *TL_inputBotInlineResult) Get_send_message() []byte {
	return t._send_message
}

func New_TL_inputBotInlineResult() *TL_inputBotInlineResult {
	return &TL_inputBotInlineResult{}
}

func (t *TL_inputBotInlineResult) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputBotInlineResult))
	ec.String(t.Get_id())
	ec.String(t.Get_type())
	ec.Bytes(t.Get_title())
	ec.Bytes(t.Get_description())
	ec.Bytes(t.Get_url())
	ec.Bytes(t.Get_thumb_url())
	ec.Bytes(t.Get_content_url())
	ec.Bytes(t.Get_content_type())
	ec.Bytes(t.Get_w())
	ec.Bytes(t.Get_h())
	ec.Bytes(t.Get_duration())
	ec.Bytes(t.Get_send_message())

	return ec.GetBuffer()
}

func (t *TL_inputBotInlineResult) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.String()
	t._type = dc.String()
	t._title = dc.Bytes(16)
	t._description = dc.Bytes(16)
	t._url = dc.Bytes(16)
	t._thumb_url = dc.Bytes(16)
	t._content_url = dc.Bytes(16)
	t._content_type = dc.Bytes(16)
	t._w = dc.Bytes(16)
	t._h = dc.Bytes(16)
	t._duration = dc.Bytes(16)
	t._send_message = dc.Bytes(16)

}

// inputBotInlineResultPhoto#a8d864a7
type TL_inputBotInlineResultPhoto struct {
	_id           string
	_type         string
	_photo        []byte
	_send_message []byte
}

func (t *TL_inputBotInlineResultPhoto) Set_id(_id string) {
	t._id = _id
}

func (t *TL_inputBotInlineResultPhoto) Get_id() string {
	return t._id
}

func (t *TL_inputBotInlineResultPhoto) Set_type(_type string) {
	t._type = _type
}

func (t *TL_inputBotInlineResultPhoto) Get_type() string {
	return t._type
}

func (t *TL_inputBotInlineResultPhoto) Set_photo(_photo []byte) {
	t._photo = _photo
}

func (t *TL_inputBotInlineResultPhoto) Get_photo() []byte {
	return t._photo
}

func (t *TL_inputBotInlineResultPhoto) Set_send_message(_send_message []byte) {
	t._send_message = _send_message
}

func (t *TL_inputBotInlineResultPhoto) Get_send_message() []byte {
	return t._send_message
}

func New_TL_inputBotInlineResultPhoto() *TL_inputBotInlineResultPhoto {
	return &TL_inputBotInlineResultPhoto{}
}

func (t *TL_inputBotInlineResultPhoto) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputBotInlineResultPhoto))
	ec.String(t.Get_id())
	ec.String(t.Get_type())
	ec.Bytes(t.Get_photo())
	ec.Bytes(t.Get_send_message())

	return ec.GetBuffer()
}

func (t *TL_inputBotInlineResultPhoto) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.String()
	t._type = dc.String()
	t._photo = dc.Bytes(16)
	t._send_message = dc.Bytes(16)

}

// inputBotInlineResultDocument#fff8fdc4
type TL_inputBotInlineResultDocument struct {
	_flags        []byte
	_id           string
	_type         string
	_title        []byte
	_description  []byte
	_document     []byte
	_send_message []byte
}

func (t *TL_inputBotInlineResultDocument) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_inputBotInlineResultDocument) Get_flags() []byte {
	return t._flags
}

func (t *TL_inputBotInlineResultDocument) Set_id(_id string) {
	t._id = _id
}

func (t *TL_inputBotInlineResultDocument) Get_id() string {
	return t._id
}

func (t *TL_inputBotInlineResultDocument) Set_type(_type string) {
	t._type = _type
}

func (t *TL_inputBotInlineResultDocument) Get_type() string {
	return t._type
}

func (t *TL_inputBotInlineResultDocument) Set_title(_title []byte) {
	t._title = _title
}

func (t *TL_inputBotInlineResultDocument) Get_title() []byte {
	return t._title
}

func (t *TL_inputBotInlineResultDocument) Set_description(_description []byte) {
	t._description = _description
}

func (t *TL_inputBotInlineResultDocument) Get_description() []byte {
	return t._description
}

func (t *TL_inputBotInlineResultDocument) Set_document(_document []byte) {
	t._document = _document
}

func (t *TL_inputBotInlineResultDocument) Get_document() []byte {
	return t._document
}

func (t *TL_inputBotInlineResultDocument) Set_send_message(_send_message []byte) {
	t._send_message = _send_message
}

func (t *TL_inputBotInlineResultDocument) Get_send_message() []byte {
	return t._send_message
}

func New_TL_inputBotInlineResultDocument() *TL_inputBotInlineResultDocument {
	return &TL_inputBotInlineResultDocument{}
}

func (t *TL_inputBotInlineResultDocument) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputBotInlineResultDocument))
	ec.String(t.Get_id())
	ec.String(t.Get_type())
	ec.Bytes(t.Get_title())
	ec.Bytes(t.Get_description())
	ec.Bytes(t.Get_document())
	ec.Bytes(t.Get_send_message())

	return ec.GetBuffer()
}

func (t *TL_inputBotInlineResultDocument) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.String()
	t._type = dc.String()
	t._title = dc.Bytes(16)
	t._description = dc.Bytes(16)
	t._document = dc.Bytes(16)
	t._send_message = dc.Bytes(16)

}

// inputBotInlineResultGame#4fa417f2
type TL_inputBotInlineResultGame struct {
	_id           string
	_short_name   string
	_send_message []byte
}

func (t *TL_inputBotInlineResultGame) Set_id(_id string) {
	t._id = _id
}

func (t *TL_inputBotInlineResultGame) Get_id() string {
	return t._id
}

func (t *TL_inputBotInlineResultGame) Set_short_name(_short_name string) {
	t._short_name = _short_name
}

func (t *TL_inputBotInlineResultGame) Get_short_name() string {
	return t._short_name
}

func (t *TL_inputBotInlineResultGame) Set_send_message(_send_message []byte) {
	t._send_message = _send_message
}

func (t *TL_inputBotInlineResultGame) Get_send_message() []byte {
	return t._send_message
}

func New_TL_inputBotInlineResultGame() *TL_inputBotInlineResultGame {
	return &TL_inputBotInlineResultGame{}
}

func (t *TL_inputBotInlineResultGame) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputBotInlineResultGame))
	ec.String(t.Get_id())
	ec.String(t.Get_short_name())
	ec.Bytes(t.Get_send_message())

	return ec.GetBuffer()
}

func (t *TL_inputBotInlineResultGame) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.String()
	t._short_name = dc.String()
	t._send_message = dc.Bytes(16)

}

// botInlineMessageMediaAuto#0a74b15b
type TL_botInlineMessageMediaAuto struct {
	_flags        []byte
	_caption      string
	_reply_markup []byte
}

func (t *TL_botInlineMessageMediaAuto) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_botInlineMessageMediaAuto) Get_flags() []byte {
	return t._flags
}

func (t *TL_botInlineMessageMediaAuto) Set_caption(_caption string) {
	t._caption = _caption
}

func (t *TL_botInlineMessageMediaAuto) Get_caption() string {
	return t._caption
}

func (t *TL_botInlineMessageMediaAuto) Set_reply_markup(_reply_markup []byte) {
	t._reply_markup = _reply_markup
}

func (t *TL_botInlineMessageMediaAuto) Get_reply_markup() []byte {
	return t._reply_markup
}

func New_TL_botInlineMessageMediaAuto() *TL_botInlineMessageMediaAuto {
	return &TL_botInlineMessageMediaAuto{}
}

func (t *TL_botInlineMessageMediaAuto) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_botInlineMessageMediaAuto))
	ec.String(t.Get_caption())
	ec.Bytes(t.Get_reply_markup())

	return ec.GetBuffer()
}

func (t *TL_botInlineMessageMediaAuto) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._caption = dc.String()
	t._reply_markup = dc.Bytes(16)

}

// botInlineMessageText#8c7f65e2
type TL_botInlineMessageText struct {
	_flags        []byte
	_no_webpage   []byte
	_message      string
	_entities     []byte
	_reply_markup []byte
}

func (t *TL_botInlineMessageText) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_botInlineMessageText) Get_flags() []byte {
	return t._flags
}

func (t *TL_botInlineMessageText) Set_no_webpage(_no_webpage []byte) {
	t._no_webpage = _no_webpage
}

func (t *TL_botInlineMessageText) Get_no_webpage() []byte {
	return t._no_webpage
}

func (t *TL_botInlineMessageText) Set_message(_message string) {
	t._message = _message
}

func (t *TL_botInlineMessageText) Get_message() string {
	return t._message
}

func (t *TL_botInlineMessageText) Set_entities(_entities []byte) {
	t._entities = _entities
}

func (t *TL_botInlineMessageText) Get_entities() []byte {
	return t._entities
}

func (t *TL_botInlineMessageText) Set_reply_markup(_reply_markup []byte) {
	t._reply_markup = _reply_markup
}

func (t *TL_botInlineMessageText) Get_reply_markup() []byte {
	return t._reply_markup
}

func New_TL_botInlineMessageText() *TL_botInlineMessageText {
	return &TL_botInlineMessageText{}
}

func (t *TL_botInlineMessageText) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_botInlineMessageText))
	ec.Bytes(t.Get_no_webpage())
	ec.String(t.Get_message())
	ec.Bytes(t.Get_entities())
	ec.Bytes(t.Get_reply_markup())

	return ec.GetBuffer()
}

func (t *TL_botInlineMessageText) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._no_webpage = dc.Bytes(16)
	t._message = dc.String()
	t._entities = dc.Bytes(16)
	t._reply_markup = dc.Bytes(16)

}

// botInlineMessageMediaGeo#b722de65
type TL_botInlineMessageMediaGeo struct {
	_flags        []byte
	_geo          []byte
	_period       int32
	_reply_markup []byte
}

func (t *TL_botInlineMessageMediaGeo) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_botInlineMessageMediaGeo) Get_flags() []byte {
	return t._flags
}

func (t *TL_botInlineMessageMediaGeo) Set_geo(_geo []byte) {
	t._geo = _geo
}

func (t *TL_botInlineMessageMediaGeo) Get_geo() []byte {
	return t._geo
}

func (t *TL_botInlineMessageMediaGeo) Set_period(_period int32) {
	t._period = _period
}

func (t *TL_botInlineMessageMediaGeo) Get_period() int32 {
	return t._period
}

func (t *TL_botInlineMessageMediaGeo) Set_reply_markup(_reply_markup []byte) {
	t._reply_markup = _reply_markup
}

func (t *TL_botInlineMessageMediaGeo) Get_reply_markup() []byte {
	return t._reply_markup
}

func New_TL_botInlineMessageMediaGeo() *TL_botInlineMessageMediaGeo {
	return &TL_botInlineMessageMediaGeo{}
}

func (t *TL_botInlineMessageMediaGeo) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_botInlineMessageMediaGeo))
	ec.Bytes(t.Get_geo())
	ec.Int(t.Get_period())
	ec.Bytes(t.Get_reply_markup())

	return ec.GetBuffer()
}

func (t *TL_botInlineMessageMediaGeo) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._geo = dc.Bytes(16)
	t._period = dc.Int()
	t._reply_markup = dc.Bytes(16)

}

// botInlineMessageMediaVenue#4366232e
type TL_botInlineMessageMediaVenue struct {
	_flags        []byte
	_geo          []byte
	_title        string
	_address      string
	_provider     string
	_venue_id     string
	_reply_markup []byte
}

func (t *TL_botInlineMessageMediaVenue) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_botInlineMessageMediaVenue) Get_flags() []byte {
	return t._flags
}

func (t *TL_botInlineMessageMediaVenue) Set_geo(_geo []byte) {
	t._geo = _geo
}

func (t *TL_botInlineMessageMediaVenue) Get_geo() []byte {
	return t._geo
}

func (t *TL_botInlineMessageMediaVenue) Set_title(_title string) {
	t._title = _title
}

func (t *TL_botInlineMessageMediaVenue) Get_title() string {
	return t._title
}

func (t *TL_botInlineMessageMediaVenue) Set_address(_address string) {
	t._address = _address
}

func (t *TL_botInlineMessageMediaVenue) Get_address() string {
	return t._address
}

func (t *TL_botInlineMessageMediaVenue) Set_provider(_provider string) {
	t._provider = _provider
}

func (t *TL_botInlineMessageMediaVenue) Get_provider() string {
	return t._provider
}

func (t *TL_botInlineMessageMediaVenue) Set_venue_id(_venue_id string) {
	t._venue_id = _venue_id
}

func (t *TL_botInlineMessageMediaVenue) Get_venue_id() string {
	return t._venue_id
}

func (t *TL_botInlineMessageMediaVenue) Set_reply_markup(_reply_markup []byte) {
	t._reply_markup = _reply_markup
}

func (t *TL_botInlineMessageMediaVenue) Get_reply_markup() []byte {
	return t._reply_markup
}

func New_TL_botInlineMessageMediaVenue() *TL_botInlineMessageMediaVenue {
	return &TL_botInlineMessageMediaVenue{}
}

func (t *TL_botInlineMessageMediaVenue) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_botInlineMessageMediaVenue))
	ec.Bytes(t.Get_geo())
	ec.String(t.Get_title())
	ec.String(t.Get_address())
	ec.String(t.Get_provider())
	ec.String(t.Get_venue_id())
	ec.Bytes(t.Get_reply_markup())

	return ec.GetBuffer()
}

func (t *TL_botInlineMessageMediaVenue) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._geo = dc.Bytes(16)
	t._title = dc.String()
	t._address = dc.String()
	t._provider = dc.String()
	t._venue_id = dc.String()
	t._reply_markup = dc.Bytes(16)

}

// botInlineMessageMediaContact#35edb4d4
type TL_botInlineMessageMediaContact struct {
	_flags        []byte
	_phone_number string
	_first_name   string
	_last_name    string
	_reply_markup []byte
}

func (t *TL_botInlineMessageMediaContact) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_botInlineMessageMediaContact) Get_flags() []byte {
	return t._flags
}

func (t *TL_botInlineMessageMediaContact) Set_phone_number(_phone_number string) {
	t._phone_number = _phone_number
}

func (t *TL_botInlineMessageMediaContact) Get_phone_number() string {
	return t._phone_number
}

func (t *TL_botInlineMessageMediaContact) Set_first_name(_first_name string) {
	t._first_name = _first_name
}

func (t *TL_botInlineMessageMediaContact) Get_first_name() string {
	return t._first_name
}

func (t *TL_botInlineMessageMediaContact) Set_last_name(_last_name string) {
	t._last_name = _last_name
}

func (t *TL_botInlineMessageMediaContact) Get_last_name() string {
	return t._last_name
}

func (t *TL_botInlineMessageMediaContact) Set_reply_markup(_reply_markup []byte) {
	t._reply_markup = _reply_markup
}

func (t *TL_botInlineMessageMediaContact) Get_reply_markup() []byte {
	return t._reply_markup
}

func New_TL_botInlineMessageMediaContact() *TL_botInlineMessageMediaContact {
	return &TL_botInlineMessageMediaContact{}
}

func (t *TL_botInlineMessageMediaContact) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_botInlineMessageMediaContact))
	ec.String(t.Get_phone_number())
	ec.String(t.Get_first_name())
	ec.String(t.Get_last_name())
	ec.Bytes(t.Get_reply_markup())

	return ec.GetBuffer()
}

func (t *TL_botInlineMessageMediaContact) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._phone_number = dc.String()
	t._first_name = dc.String()
	t._last_name = dc.String()
	t._reply_markup = dc.Bytes(16)

}

// botInlineResult#9bebaeb9
type TL_botInlineResult struct {
	_flags        []byte
	_id           string
	_type         string
	_title        []byte
	_description  []byte
	_url          []byte
	_thumb_url    []byte
	_content_url  []byte
	_content_type []byte
	_w            []byte
	_h            []byte
	_duration     []byte
	_send_message []byte
}

func (t *TL_botInlineResult) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_botInlineResult) Get_flags() []byte {
	return t._flags
}

func (t *TL_botInlineResult) Set_id(_id string) {
	t._id = _id
}

func (t *TL_botInlineResult) Get_id() string {
	return t._id
}

func (t *TL_botInlineResult) Set_type(_type string) {
	t._type = _type
}

func (t *TL_botInlineResult) Get_type() string {
	return t._type
}

func (t *TL_botInlineResult) Set_title(_title []byte) {
	t._title = _title
}

func (t *TL_botInlineResult) Get_title() []byte {
	return t._title
}

func (t *TL_botInlineResult) Set_description(_description []byte) {
	t._description = _description
}

func (t *TL_botInlineResult) Get_description() []byte {
	return t._description
}

func (t *TL_botInlineResult) Set_url(_url []byte) {
	t._url = _url
}

func (t *TL_botInlineResult) Get_url() []byte {
	return t._url
}

func (t *TL_botInlineResult) Set_thumb_url(_thumb_url []byte) {
	t._thumb_url = _thumb_url
}

func (t *TL_botInlineResult) Get_thumb_url() []byte {
	return t._thumb_url
}

func (t *TL_botInlineResult) Set_content_url(_content_url []byte) {
	t._content_url = _content_url
}

func (t *TL_botInlineResult) Get_content_url() []byte {
	return t._content_url
}

func (t *TL_botInlineResult) Set_content_type(_content_type []byte) {
	t._content_type = _content_type
}

func (t *TL_botInlineResult) Get_content_type() []byte {
	return t._content_type
}

func (t *TL_botInlineResult) Set_w(_w []byte) {
	t._w = _w
}

func (t *TL_botInlineResult) Get_w() []byte {
	return t._w
}

func (t *TL_botInlineResult) Set_h(_h []byte) {
	t._h = _h
}

func (t *TL_botInlineResult) Get_h() []byte {
	return t._h
}

func (t *TL_botInlineResult) Set_duration(_duration []byte) {
	t._duration = _duration
}

func (t *TL_botInlineResult) Get_duration() []byte {
	return t._duration
}

func (t *TL_botInlineResult) Set_send_message(_send_message []byte) {
	t._send_message = _send_message
}

func (t *TL_botInlineResult) Get_send_message() []byte {
	return t._send_message
}

func New_TL_botInlineResult() *TL_botInlineResult {
	return &TL_botInlineResult{}
}

func (t *TL_botInlineResult) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_botInlineResult))
	ec.String(t.Get_id())
	ec.String(t.Get_type())
	ec.Bytes(t.Get_title())
	ec.Bytes(t.Get_description())
	ec.Bytes(t.Get_url())
	ec.Bytes(t.Get_thumb_url())
	ec.Bytes(t.Get_content_url())
	ec.Bytes(t.Get_content_type())
	ec.Bytes(t.Get_w())
	ec.Bytes(t.Get_h())
	ec.Bytes(t.Get_duration())
	ec.Bytes(t.Get_send_message())

	return ec.GetBuffer()
}

func (t *TL_botInlineResult) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.String()
	t._type = dc.String()
	t._title = dc.Bytes(16)
	t._description = dc.Bytes(16)
	t._url = dc.Bytes(16)
	t._thumb_url = dc.Bytes(16)
	t._content_url = dc.Bytes(16)
	t._content_type = dc.Bytes(16)
	t._w = dc.Bytes(16)
	t._h = dc.Bytes(16)
	t._duration = dc.Bytes(16)
	t._send_message = dc.Bytes(16)

}

// botInlineMediaResult#17db940b
type TL_botInlineMediaResult struct {
	_flags        []byte
	_id           string
	_type         string
	_photo        []byte
	_document     []byte
	_title        []byte
	_description  []byte
	_send_message []byte
}

func (t *TL_botInlineMediaResult) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_botInlineMediaResult) Get_flags() []byte {
	return t._flags
}

func (t *TL_botInlineMediaResult) Set_id(_id string) {
	t._id = _id
}

func (t *TL_botInlineMediaResult) Get_id() string {
	return t._id
}

func (t *TL_botInlineMediaResult) Set_type(_type string) {
	t._type = _type
}

func (t *TL_botInlineMediaResult) Get_type() string {
	return t._type
}

func (t *TL_botInlineMediaResult) Set_photo(_photo []byte) {
	t._photo = _photo
}

func (t *TL_botInlineMediaResult) Get_photo() []byte {
	return t._photo
}

func (t *TL_botInlineMediaResult) Set_document(_document []byte) {
	t._document = _document
}

func (t *TL_botInlineMediaResult) Get_document() []byte {
	return t._document
}

func (t *TL_botInlineMediaResult) Set_title(_title []byte) {
	t._title = _title
}

func (t *TL_botInlineMediaResult) Get_title() []byte {
	return t._title
}

func (t *TL_botInlineMediaResult) Set_description(_description []byte) {
	t._description = _description
}

func (t *TL_botInlineMediaResult) Get_description() []byte {
	return t._description
}

func (t *TL_botInlineMediaResult) Set_send_message(_send_message []byte) {
	t._send_message = _send_message
}

func (t *TL_botInlineMediaResult) Get_send_message() []byte {
	return t._send_message
}

func New_TL_botInlineMediaResult() *TL_botInlineMediaResult {
	return &TL_botInlineMediaResult{}
}

func (t *TL_botInlineMediaResult) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_botInlineMediaResult))
	ec.String(t.Get_id())
	ec.String(t.Get_type())
	ec.Bytes(t.Get_photo())
	ec.Bytes(t.Get_document())
	ec.Bytes(t.Get_title())
	ec.Bytes(t.Get_description())
	ec.Bytes(t.Get_send_message())

	return ec.GetBuffer()
}

func (t *TL_botInlineMediaResult) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.String()
	t._type = dc.String()
	t._photo = dc.Bytes(16)
	t._document = dc.Bytes(16)
	t._title = dc.Bytes(16)
	t._description = dc.Bytes(16)
	t._send_message = dc.Bytes(16)

}

// messages_botResults#947ca848
type TL_messages_botResults struct {
	_flags       []byte
	_gallery     []byte
	_query_id    int64
	_next_offset []byte
	_switch_pm   []byte
	_results     []byte
	_cache_time  int32
	_users       []byte
}

func (t *TL_messages_botResults) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_messages_botResults) Get_flags() []byte {
	return t._flags
}

func (t *TL_messages_botResults) Set_gallery(_gallery []byte) {
	t._gallery = _gallery
}

func (t *TL_messages_botResults) Get_gallery() []byte {
	return t._gallery
}

func (t *TL_messages_botResults) Set_query_id(_query_id int64) {
	t._query_id = _query_id
}

func (t *TL_messages_botResults) Get_query_id() int64 {
	return t._query_id
}

func (t *TL_messages_botResults) Set_next_offset(_next_offset []byte) {
	t._next_offset = _next_offset
}

func (t *TL_messages_botResults) Get_next_offset() []byte {
	return t._next_offset
}

func (t *TL_messages_botResults) Set_switch_pm(_switch_pm []byte) {
	t._switch_pm = _switch_pm
}

func (t *TL_messages_botResults) Get_switch_pm() []byte {
	return t._switch_pm
}

func (t *TL_messages_botResults) Set_results(_results []byte) {
	t._results = _results
}

func (t *TL_messages_botResults) Get_results() []byte {
	return t._results
}

func (t *TL_messages_botResults) Set_cache_time(_cache_time int32) {
	t._cache_time = _cache_time
}

func (t *TL_messages_botResults) Get_cache_time() int32 {
	return t._cache_time
}

func (t *TL_messages_botResults) Set_users(_users []byte) {
	t._users = _users
}

func (t *TL_messages_botResults) Get_users() []byte {
	return t._users
}

func New_TL_messages_botResults() *TL_messages_botResults {
	return &TL_messages_botResults{}
}

func (t *TL_messages_botResults) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_botResults))
	ec.Bytes(t.Get_gallery())
	ec.Long(t.Get_query_id())
	ec.Bytes(t.Get_next_offset())
	ec.Bytes(t.Get_switch_pm())
	ec.Bytes(t.Get_results())
	ec.Int(t.Get_cache_time())
	ec.Bytes(t.Get_users())

	return ec.GetBuffer()
}

func (t *TL_messages_botResults) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._gallery = dc.Bytes(16)
	t._query_id = dc.Long()
	t._next_offset = dc.Bytes(16)
	t._switch_pm = dc.Bytes(16)
	t._results = dc.Bytes(16)
	t._cache_time = dc.Int()
	t._users = dc.Bytes(16)

}

// exportedMessageLink#1f486803
type TL_exportedMessageLink struct {
	_link string
}

func (t *TL_exportedMessageLink) Set_link(_link string) {
	t._link = _link
}

func (t *TL_exportedMessageLink) Get_link() string {
	return t._link
}

func New_TL_exportedMessageLink() *TL_exportedMessageLink {
	return &TL_exportedMessageLink{}
}

func (t *TL_exportedMessageLink) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_exportedMessageLink))
	ec.String(t.Get_link())

	return ec.GetBuffer()
}

func (t *TL_exportedMessageLink) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._link = dc.String()

}

// messageFwdHeader#559ebe6d
type TL_messageFwdHeader struct {
	_flags             []byte
	_from_id           []byte
	_date              int32
	_channel_id        []byte
	_channel_post      []byte
	_post_author       []byte
	_saved_from_peer   []byte
	_saved_from_msg_id []byte
}

func (t *TL_messageFwdHeader) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_messageFwdHeader) Get_flags() []byte {
	return t._flags
}

func (t *TL_messageFwdHeader) Set_from_id(_from_id []byte) {
	t._from_id = _from_id
}

func (t *TL_messageFwdHeader) Get_from_id() []byte {
	return t._from_id
}

func (t *TL_messageFwdHeader) Set_date(_date int32) {
	t._date = _date
}

func (t *TL_messageFwdHeader) Get_date() int32 {
	return t._date
}

func (t *TL_messageFwdHeader) Set_channel_id(_channel_id []byte) {
	t._channel_id = _channel_id
}

func (t *TL_messageFwdHeader) Get_channel_id() []byte {
	return t._channel_id
}

func (t *TL_messageFwdHeader) Set_channel_post(_channel_post []byte) {
	t._channel_post = _channel_post
}

func (t *TL_messageFwdHeader) Get_channel_post() []byte {
	return t._channel_post
}

func (t *TL_messageFwdHeader) Set_post_author(_post_author []byte) {
	t._post_author = _post_author
}

func (t *TL_messageFwdHeader) Get_post_author() []byte {
	return t._post_author
}

func (t *TL_messageFwdHeader) Set_saved_from_peer(_saved_from_peer []byte) {
	t._saved_from_peer = _saved_from_peer
}

func (t *TL_messageFwdHeader) Get_saved_from_peer() []byte {
	return t._saved_from_peer
}

func (t *TL_messageFwdHeader) Set_saved_from_msg_id(_saved_from_msg_id []byte) {
	t._saved_from_msg_id = _saved_from_msg_id
}

func (t *TL_messageFwdHeader) Get_saved_from_msg_id() []byte {
	return t._saved_from_msg_id
}

func New_TL_messageFwdHeader() *TL_messageFwdHeader {
	return &TL_messageFwdHeader{}
}

func (t *TL_messageFwdHeader) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messageFwdHeader))
	ec.Bytes(t.Get_from_id())
	ec.Int(t.Get_date())
	ec.Bytes(t.Get_channel_id())
	ec.Bytes(t.Get_channel_post())
	ec.Bytes(t.Get_post_author())
	ec.Bytes(t.Get_saved_from_peer())
	ec.Bytes(t.Get_saved_from_msg_id())

	return ec.GetBuffer()
}

func (t *TL_messageFwdHeader) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._from_id = dc.Bytes(16)
	t._date = dc.Int()
	t._channel_id = dc.Bytes(16)
	t._channel_post = dc.Bytes(16)
	t._post_author = dc.Bytes(16)
	t._saved_from_peer = dc.Bytes(16)
	t._saved_from_msg_id = dc.Bytes(16)

}

// auth_codeTypeSms#72a3158c
type TL_auth_codeTypeSms struct {
}

func New_TL_auth_codeTypeSms() *TL_auth_codeTypeSms {
	return &TL_auth_codeTypeSms{}
}

func (t *TL_auth_codeTypeSms) Encode() []byte {
	return nil
}

func (t *TL_auth_codeTypeSms) Decode(b []byte) {

}

// auth_codeTypeCall#741cd3e3
type TL_auth_codeTypeCall struct {
}

func New_TL_auth_codeTypeCall() *TL_auth_codeTypeCall {
	return &TL_auth_codeTypeCall{}
}

func (t *TL_auth_codeTypeCall) Encode() []byte {
	return nil
}

func (t *TL_auth_codeTypeCall) Decode(b []byte) {

}

// auth_codeTypeFlashCall#226ccefb
type TL_auth_codeTypeFlashCall struct {
}

func New_TL_auth_codeTypeFlashCall() *TL_auth_codeTypeFlashCall {
	return &TL_auth_codeTypeFlashCall{}
}

func (t *TL_auth_codeTypeFlashCall) Encode() []byte {
	return nil
}

func (t *TL_auth_codeTypeFlashCall) Decode(b []byte) {

}

// auth_sentCodeTypeApp#3dbb5986
type TL_auth_sentCodeTypeApp struct {
	_length int32
}

func (t *TL_auth_sentCodeTypeApp) Set_length(_length int32) {
	t._length = _length
}

func (t *TL_auth_sentCodeTypeApp) Get_length() int32 {
	return t._length
}

func New_TL_auth_sentCodeTypeApp() *TL_auth_sentCodeTypeApp {
	return &TL_auth_sentCodeTypeApp{}
}

func (t *TL_auth_sentCodeTypeApp) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_auth_sentCodeTypeApp))
	ec.Int(t.Get_length())

	return ec.GetBuffer()
}

func (t *TL_auth_sentCodeTypeApp) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._length = dc.Int()

}

// auth_sentCodeTypeSms#c000bba2
type TL_auth_sentCodeTypeSms struct {
	_length int32
}

func (t *TL_auth_sentCodeTypeSms) Set_length(_length int32) {
	t._length = _length
}

func (t *TL_auth_sentCodeTypeSms) Get_length() int32 {
	return t._length
}

func New_TL_auth_sentCodeTypeSms() *TL_auth_sentCodeTypeSms {
	return &TL_auth_sentCodeTypeSms{}
}

func (t *TL_auth_sentCodeTypeSms) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_auth_sentCodeTypeSms))
	ec.Int(t.Get_length())

	return ec.GetBuffer()
}

func (t *TL_auth_sentCodeTypeSms) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._length = dc.Int()

}

// auth_sentCodeTypeCall#5353e5a7
type TL_auth_sentCodeTypeCall struct {
	_length int32
}

func (t *TL_auth_sentCodeTypeCall) Set_length(_length int32) {
	t._length = _length
}

func (t *TL_auth_sentCodeTypeCall) Get_length() int32 {
	return t._length
}

func New_TL_auth_sentCodeTypeCall() *TL_auth_sentCodeTypeCall {
	return &TL_auth_sentCodeTypeCall{}
}

func (t *TL_auth_sentCodeTypeCall) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_auth_sentCodeTypeCall))
	ec.Int(t.Get_length())

	return ec.GetBuffer()
}

func (t *TL_auth_sentCodeTypeCall) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._length = dc.Int()

}

// auth_sentCodeTypeFlashCall#ab03c6d9
type TL_auth_sentCodeTypeFlashCall struct {
	_pattern string
}

func (t *TL_auth_sentCodeTypeFlashCall) Set_pattern(_pattern string) {
	t._pattern = _pattern
}

func (t *TL_auth_sentCodeTypeFlashCall) Get_pattern() string {
	return t._pattern
}

func New_TL_auth_sentCodeTypeFlashCall() *TL_auth_sentCodeTypeFlashCall {
	return &TL_auth_sentCodeTypeFlashCall{}
}

func (t *TL_auth_sentCodeTypeFlashCall) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_auth_sentCodeTypeFlashCall))
	ec.String(t.Get_pattern())

	return ec.GetBuffer()
}

func (t *TL_auth_sentCodeTypeFlashCall) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._pattern = dc.String()

}

// messages_botCallbackAnswer#36585ea4
type TL_messages_botCallbackAnswer struct {
	_flags      []byte
	_alert      []byte
	_has_url    []byte
	_native_ui  []byte
	_message    []byte
	_url        []byte
	_cache_time int32
}

func (t *TL_messages_botCallbackAnswer) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_messages_botCallbackAnswer) Get_flags() []byte {
	return t._flags
}

func (t *TL_messages_botCallbackAnswer) Set_alert(_alert []byte) {
	t._alert = _alert
}

func (t *TL_messages_botCallbackAnswer) Get_alert() []byte {
	return t._alert
}

func (t *TL_messages_botCallbackAnswer) Set_has_url(_has_url []byte) {
	t._has_url = _has_url
}

func (t *TL_messages_botCallbackAnswer) Get_has_url() []byte {
	return t._has_url
}

func (t *TL_messages_botCallbackAnswer) Set_native_ui(_native_ui []byte) {
	t._native_ui = _native_ui
}

func (t *TL_messages_botCallbackAnswer) Get_native_ui() []byte {
	return t._native_ui
}

func (t *TL_messages_botCallbackAnswer) Set_message(_message []byte) {
	t._message = _message
}

func (t *TL_messages_botCallbackAnswer) Get_message() []byte {
	return t._message
}

func (t *TL_messages_botCallbackAnswer) Set_url(_url []byte) {
	t._url = _url
}

func (t *TL_messages_botCallbackAnswer) Get_url() []byte {
	return t._url
}

func (t *TL_messages_botCallbackAnswer) Set_cache_time(_cache_time int32) {
	t._cache_time = _cache_time
}

func (t *TL_messages_botCallbackAnswer) Get_cache_time() int32 {
	return t._cache_time
}

func New_TL_messages_botCallbackAnswer() *TL_messages_botCallbackAnswer {
	return &TL_messages_botCallbackAnswer{}
}

func (t *TL_messages_botCallbackAnswer) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_botCallbackAnswer))
	ec.Bytes(t.Get_alert())
	ec.Bytes(t.Get_has_url())
	ec.Bytes(t.Get_native_ui())
	ec.Bytes(t.Get_message())
	ec.Bytes(t.Get_url())
	ec.Int(t.Get_cache_time())

	return ec.GetBuffer()
}

func (t *TL_messages_botCallbackAnswer) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._alert = dc.Bytes(16)
	t._has_url = dc.Bytes(16)
	t._native_ui = dc.Bytes(16)
	t._message = dc.Bytes(16)
	t._url = dc.Bytes(16)
	t._cache_time = dc.Int()

}

// messages_messageEditData#26b5dde6
type TL_messages_messageEditData struct {
	_flags   []byte
	_caption []byte
}

func (t *TL_messages_messageEditData) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_messages_messageEditData) Get_flags() []byte {
	return t._flags
}

func (t *TL_messages_messageEditData) Set_caption(_caption []byte) {
	t._caption = _caption
}

func (t *TL_messages_messageEditData) Get_caption() []byte {
	return t._caption
}

func New_TL_messages_messageEditData() *TL_messages_messageEditData {
	return &TL_messages_messageEditData{}
}

func (t *TL_messages_messageEditData) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_messageEditData))
	ec.Bytes(t.Get_caption())

	return ec.GetBuffer()
}

func (t *TL_messages_messageEditData) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._caption = dc.Bytes(16)

}

// inputBotInlineMessageID#890c3d89
type TL_inputBotInlineMessageID struct {
	_dc_id       int32
	_id          int64
	_access_hash int64
}

func (t *TL_inputBotInlineMessageID) Set_dc_id(_dc_id int32) {
	t._dc_id = _dc_id
}

func (t *TL_inputBotInlineMessageID) Get_dc_id() int32 {
	return t._dc_id
}

func (t *TL_inputBotInlineMessageID) Set_id(_id int64) {
	t._id = _id
}

func (t *TL_inputBotInlineMessageID) Get_id() int64 {
	return t._id
}

func (t *TL_inputBotInlineMessageID) Set_access_hash(_access_hash int64) {
	t._access_hash = _access_hash
}

func (t *TL_inputBotInlineMessageID) Get_access_hash() int64 {
	return t._access_hash
}

func New_TL_inputBotInlineMessageID() *TL_inputBotInlineMessageID {
	return &TL_inputBotInlineMessageID{}
}

func (t *TL_inputBotInlineMessageID) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputBotInlineMessageID))
	ec.Int(t.Get_dc_id())
	ec.Long(t.Get_id())
	ec.Long(t.Get_access_hash())

	return ec.GetBuffer()
}

func (t *TL_inputBotInlineMessageID) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._dc_id = dc.Int()
	t._id = dc.Long()
	t._access_hash = dc.Long()

}

// inlineBotSwitchPM#3c20629f
type TL_inlineBotSwitchPM struct {
	_text        string
	_start_param string
}

func (t *TL_inlineBotSwitchPM) Set_text(_text string) {
	t._text = _text
}

func (t *TL_inlineBotSwitchPM) Get_text() string {
	return t._text
}

func (t *TL_inlineBotSwitchPM) Set_start_param(_start_param string) {
	t._start_param = _start_param
}

func (t *TL_inlineBotSwitchPM) Get_start_param() string {
	return t._start_param
}

func New_TL_inlineBotSwitchPM() *TL_inlineBotSwitchPM {
	return &TL_inlineBotSwitchPM{}
}

func (t *TL_inlineBotSwitchPM) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inlineBotSwitchPM))
	ec.String(t.Get_text())
	ec.String(t.Get_start_param())

	return ec.GetBuffer()
}

func (t *TL_inlineBotSwitchPM) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._text = dc.String()
	t._start_param = dc.String()

}

// messages_peerDialogs#3371c354
type TL_messages_peerDialogs struct {
	_dialogs  []byte
	_messages []byte
	_chats    []byte
	_users    []byte
	_state    []byte
}

func (t *TL_messages_peerDialogs) Set_dialogs(_dialogs []byte) {
	t._dialogs = _dialogs
}

func (t *TL_messages_peerDialogs) Get_dialogs() []byte {
	return t._dialogs
}

func (t *TL_messages_peerDialogs) Set_messages(_messages []byte) {
	t._messages = _messages
}

func (t *TL_messages_peerDialogs) Get_messages() []byte {
	return t._messages
}

func (t *TL_messages_peerDialogs) Set_chats(_chats []byte) {
	t._chats = _chats
}

func (t *TL_messages_peerDialogs) Get_chats() []byte {
	return t._chats
}

func (t *TL_messages_peerDialogs) Set_users(_users []byte) {
	t._users = _users
}

func (t *TL_messages_peerDialogs) Get_users() []byte {
	return t._users
}

func (t *TL_messages_peerDialogs) Set_state(_state []byte) {
	t._state = _state
}

func (t *TL_messages_peerDialogs) Get_state() []byte {
	return t._state
}

func New_TL_messages_peerDialogs() *TL_messages_peerDialogs {
	return &TL_messages_peerDialogs{}
}

func (t *TL_messages_peerDialogs) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_peerDialogs))
	ec.Bytes(t.Get_dialogs())
	ec.Bytes(t.Get_messages())
	ec.Bytes(t.Get_chats())
	ec.Bytes(t.Get_users())
	ec.Bytes(t.Get_state())

	return ec.GetBuffer()
}

func (t *TL_messages_peerDialogs) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._dialogs = dc.Bytes(16)
	t._messages = dc.Bytes(16)
	t._chats = dc.Bytes(16)
	t._users = dc.Bytes(16)
	t._state = dc.Bytes(16)

}

// topPeer#edcdc05b
type TL_topPeer struct {
	_peer   []byte
	_rating float64
}

func (t *TL_topPeer) Set_peer(_peer []byte) {
	t._peer = _peer
}

func (t *TL_topPeer) Get_peer() []byte {
	return t._peer
}

func (t *TL_topPeer) Set_rating(_rating float64) {
	t._rating = _rating
}

func (t *TL_topPeer) Get_rating() float64 {
	return t._rating
}

func New_TL_topPeer() *TL_topPeer {
	return &TL_topPeer{}
}

func (t *TL_topPeer) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_topPeer))
	ec.Bytes(t.Get_peer())
	ec.Double(t.Get_rating())

	return ec.GetBuffer()
}

func (t *TL_topPeer) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._peer = dc.Bytes(16)
	t._rating = dc.Double()

}

// topPeerCategoryBotsPM#ab661b5b
type TL_topPeerCategoryBotsPM struct {
}

func New_TL_topPeerCategoryBotsPM() *TL_topPeerCategoryBotsPM {
	return &TL_topPeerCategoryBotsPM{}
}

func (t *TL_topPeerCategoryBotsPM) Encode() []byte {
	return nil
}

func (t *TL_topPeerCategoryBotsPM) Decode(b []byte) {

}

// topPeerCategoryBotsInline#148677e2
type TL_topPeerCategoryBotsInline struct {
}

func New_TL_topPeerCategoryBotsInline() *TL_topPeerCategoryBotsInline {
	return &TL_topPeerCategoryBotsInline{}
}

func (t *TL_topPeerCategoryBotsInline) Encode() []byte {
	return nil
}

func (t *TL_topPeerCategoryBotsInline) Decode(b []byte) {

}

// topPeerCategoryCorrespondents#0637b7ed
type TL_topPeerCategoryCorrespondents struct {
}

func New_TL_topPeerCategoryCorrespondents() *TL_topPeerCategoryCorrespondents {
	return &TL_topPeerCategoryCorrespondents{}
}

func (t *TL_topPeerCategoryCorrespondents) Encode() []byte {
	return nil
}

func (t *TL_topPeerCategoryCorrespondents) Decode(b []byte) {

}

// topPeerCategoryGroups#bd17a14a
type TL_topPeerCategoryGroups struct {
}

func New_TL_topPeerCategoryGroups() *TL_topPeerCategoryGroups {
	return &TL_topPeerCategoryGroups{}
}

func (t *TL_topPeerCategoryGroups) Encode() []byte {
	return nil
}

func (t *TL_topPeerCategoryGroups) Decode(b []byte) {

}

// topPeerCategoryChannels#161d9628
type TL_topPeerCategoryChannels struct {
}

func New_TL_topPeerCategoryChannels() *TL_topPeerCategoryChannels {
	return &TL_topPeerCategoryChannels{}
}

func (t *TL_topPeerCategoryChannels) Encode() []byte {
	return nil
}

func (t *TL_topPeerCategoryChannels) Decode(b []byte) {

}

// topPeerCategoryPhoneCalls#1e76a78c
type TL_topPeerCategoryPhoneCalls struct {
}

func New_TL_topPeerCategoryPhoneCalls() *TL_topPeerCategoryPhoneCalls {
	return &TL_topPeerCategoryPhoneCalls{}
}

func (t *TL_topPeerCategoryPhoneCalls) Encode() []byte {
	return nil
}

func (t *TL_topPeerCategoryPhoneCalls) Decode(b []byte) {

}

// topPeerCategoryPeers#fb834291
type TL_topPeerCategoryPeers struct {
	_category []byte
	_count    int32
	_peers    []byte
}

func (t *TL_topPeerCategoryPeers) Set_category(_category []byte) {
	t._category = _category
}

func (t *TL_topPeerCategoryPeers) Get_category() []byte {
	return t._category
}

func (t *TL_topPeerCategoryPeers) Set_count(_count int32) {
	t._count = _count
}

func (t *TL_topPeerCategoryPeers) Get_count() int32 {
	return t._count
}

func (t *TL_topPeerCategoryPeers) Set_peers(_peers []byte) {
	t._peers = _peers
}

func (t *TL_topPeerCategoryPeers) Get_peers() []byte {
	return t._peers
}

func New_TL_topPeerCategoryPeers() *TL_topPeerCategoryPeers {
	return &TL_topPeerCategoryPeers{}
}

func (t *TL_topPeerCategoryPeers) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_topPeerCategoryPeers))
	ec.Bytes(t.Get_category())
	ec.Int(t.Get_count())
	ec.Bytes(t.Get_peers())

	return ec.GetBuffer()
}

func (t *TL_topPeerCategoryPeers) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._category = dc.Bytes(16)
	t._count = dc.Int()
	t._peers = dc.Bytes(16)

}

// contacts_topPeersNotModified#de266ef5
type TL_contacts_topPeersNotModified struct {
}

func New_TL_contacts_topPeersNotModified() *TL_contacts_topPeersNotModified {
	return &TL_contacts_topPeersNotModified{}
}

func (t *TL_contacts_topPeersNotModified) Encode() []byte {
	return nil
}

func (t *TL_contacts_topPeersNotModified) Decode(b []byte) {

}

// contacts_topPeers#70b772a8
type TL_contacts_topPeers struct {
	_categories []byte
	_chats      []byte
	_users      []byte
}

func (t *TL_contacts_topPeers) Set_categories(_categories []byte) {
	t._categories = _categories
}

func (t *TL_contacts_topPeers) Get_categories() []byte {
	return t._categories
}

func (t *TL_contacts_topPeers) Set_chats(_chats []byte) {
	t._chats = _chats
}

func (t *TL_contacts_topPeers) Get_chats() []byte {
	return t._chats
}

func (t *TL_contacts_topPeers) Set_users(_users []byte) {
	t._users = _users
}

func (t *TL_contacts_topPeers) Get_users() []byte {
	return t._users
}

func New_TL_contacts_topPeers() *TL_contacts_topPeers {
	return &TL_contacts_topPeers{}
}

func (t *TL_contacts_topPeers) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_contacts_topPeers))
	ec.Bytes(t.Get_categories())
	ec.Bytes(t.Get_chats())
	ec.Bytes(t.Get_users())

	return ec.GetBuffer()
}

func (t *TL_contacts_topPeers) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._categories = dc.Bytes(16)
	t._chats = dc.Bytes(16)
	t._users = dc.Bytes(16)

}

// draftMessageEmpty#ba4baec5
type TL_draftMessageEmpty struct {
}

func New_TL_draftMessageEmpty() *TL_draftMessageEmpty {
	return &TL_draftMessageEmpty{}
}

func (t *TL_draftMessageEmpty) Encode() []byte {
	return nil
}

func (t *TL_draftMessageEmpty) Decode(b []byte) {

}

// draftMessage#fd8e711f
type TL_draftMessage struct {
	_flags           []byte
	_no_webpage      []byte
	_reply_to_msg_id []byte
	_message         string
	_entities        []byte
	_date            int32
}

func (t *TL_draftMessage) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_draftMessage) Get_flags() []byte {
	return t._flags
}

func (t *TL_draftMessage) Set_no_webpage(_no_webpage []byte) {
	t._no_webpage = _no_webpage
}

func (t *TL_draftMessage) Get_no_webpage() []byte {
	return t._no_webpage
}

func (t *TL_draftMessage) Set_reply_to_msg_id(_reply_to_msg_id []byte) {
	t._reply_to_msg_id = _reply_to_msg_id
}

func (t *TL_draftMessage) Get_reply_to_msg_id() []byte {
	return t._reply_to_msg_id
}

func (t *TL_draftMessage) Set_message(_message string) {
	t._message = _message
}

func (t *TL_draftMessage) Get_message() string {
	return t._message
}

func (t *TL_draftMessage) Set_entities(_entities []byte) {
	t._entities = _entities
}

func (t *TL_draftMessage) Get_entities() []byte {
	return t._entities
}

func (t *TL_draftMessage) Set_date(_date int32) {
	t._date = _date
}

func (t *TL_draftMessage) Get_date() int32 {
	return t._date
}

func New_TL_draftMessage() *TL_draftMessage {
	return &TL_draftMessage{}
}

func (t *TL_draftMessage) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_draftMessage))
	ec.Bytes(t.Get_no_webpage())
	ec.Bytes(t.Get_reply_to_msg_id())
	ec.String(t.Get_message())
	ec.Bytes(t.Get_entities())
	ec.Int(t.Get_date())

	return ec.GetBuffer()
}

func (t *TL_draftMessage) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._no_webpage = dc.Bytes(16)
	t._reply_to_msg_id = dc.Bytes(16)
	t._message = dc.String()
	t._entities = dc.Bytes(16)
	t._date = dc.Int()

}

// messages_featuredStickersNotModified#04ede3cf
type TL_messages_featuredStickersNotModified struct {
}

func New_TL_messages_featuredStickersNotModified() *TL_messages_featuredStickersNotModified {
	return &TL_messages_featuredStickersNotModified{}
}

func (t *TL_messages_featuredStickersNotModified) Encode() []byte {
	return nil
}

func (t *TL_messages_featuredStickersNotModified) Decode(b []byte) {

}

// messages_featuredStickers#f89d88e5
type TL_messages_featuredStickers struct {
	_hash   int32
	_sets   []byte
	_unread []byte
}

func (t *TL_messages_featuredStickers) Set_hash(_hash int32) {
	t._hash = _hash
}

func (t *TL_messages_featuredStickers) Get_hash() int32 {
	return t._hash
}

func (t *TL_messages_featuredStickers) Set_sets(_sets []byte) {
	t._sets = _sets
}

func (t *TL_messages_featuredStickers) Get_sets() []byte {
	return t._sets
}

func (t *TL_messages_featuredStickers) Set_unread(_unread []byte) {
	t._unread = _unread
}

func (t *TL_messages_featuredStickers) Get_unread() []byte {
	return t._unread
}

func New_TL_messages_featuredStickers() *TL_messages_featuredStickers {
	return &TL_messages_featuredStickers{}
}

func (t *TL_messages_featuredStickers) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_featuredStickers))
	ec.Int(t.Get_hash())
	ec.Bytes(t.Get_sets())
	ec.Bytes(t.Get_unread())

	return ec.GetBuffer()
}

func (t *TL_messages_featuredStickers) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._hash = dc.Int()
	t._sets = dc.Bytes(16)
	t._unread = dc.Bytes(16)

}

// messages_recentStickersNotModified#0b17f890
type TL_messages_recentStickersNotModified struct {
}

func New_TL_messages_recentStickersNotModified() *TL_messages_recentStickersNotModified {
	return &TL_messages_recentStickersNotModified{}
}

func (t *TL_messages_recentStickersNotModified) Encode() []byte {
	return nil
}

func (t *TL_messages_recentStickersNotModified) Decode(b []byte) {

}

// messages_recentStickers#5ce20970
type TL_messages_recentStickers struct {
	_hash     int32
	_stickers []byte
}

func (t *TL_messages_recentStickers) Set_hash(_hash int32) {
	t._hash = _hash
}

func (t *TL_messages_recentStickers) Get_hash() int32 {
	return t._hash
}

func (t *TL_messages_recentStickers) Set_stickers(_stickers []byte) {
	t._stickers = _stickers
}

func (t *TL_messages_recentStickers) Get_stickers() []byte {
	return t._stickers
}

func New_TL_messages_recentStickers() *TL_messages_recentStickers {
	return &TL_messages_recentStickers{}
}

func (t *TL_messages_recentStickers) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_recentStickers))
	ec.Int(t.Get_hash())
	ec.Bytes(t.Get_stickers())

	return ec.GetBuffer()
}

func (t *TL_messages_recentStickers) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._hash = dc.Int()
	t._stickers = dc.Bytes(16)

}

// messages_archivedStickers#4fcba9c8
type TL_messages_archivedStickers struct {
	_count int32
	_sets  []byte
}

func (t *TL_messages_archivedStickers) Set_count(_count int32) {
	t._count = _count
}

func (t *TL_messages_archivedStickers) Get_count() int32 {
	return t._count
}

func (t *TL_messages_archivedStickers) Set_sets(_sets []byte) {
	t._sets = _sets
}

func (t *TL_messages_archivedStickers) Get_sets() []byte {
	return t._sets
}

func New_TL_messages_archivedStickers() *TL_messages_archivedStickers {
	return &TL_messages_archivedStickers{}
}

func (t *TL_messages_archivedStickers) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_archivedStickers))
	ec.Int(t.Get_count())
	ec.Bytes(t.Get_sets())

	return ec.GetBuffer()
}

func (t *TL_messages_archivedStickers) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._count = dc.Int()
	t._sets = dc.Bytes(16)

}

// messages_stickerSetInstallResultSuccess#38641628
type TL_messages_stickerSetInstallResultSuccess struct {
}

func New_TL_messages_stickerSetInstallResultSuccess() *TL_messages_stickerSetInstallResultSuccess {
	return &TL_messages_stickerSetInstallResultSuccess{}
}

func (t *TL_messages_stickerSetInstallResultSuccess) Encode() []byte {
	return nil
}

func (t *TL_messages_stickerSetInstallResultSuccess) Decode(b []byte) {

}

// messages_stickerSetInstallResultArchive#35e410a8
type TL_messages_stickerSetInstallResultArchive struct {
	_sets []byte
}

func (t *TL_messages_stickerSetInstallResultArchive) Set_sets(_sets []byte) {
	t._sets = _sets
}

func (t *TL_messages_stickerSetInstallResultArchive) Get_sets() []byte {
	return t._sets
}

func New_TL_messages_stickerSetInstallResultArchive() *TL_messages_stickerSetInstallResultArchive {
	return &TL_messages_stickerSetInstallResultArchive{}
}

func (t *TL_messages_stickerSetInstallResultArchive) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_stickerSetInstallResultArchive))
	ec.Bytes(t.Get_sets())

	return ec.GetBuffer()
}

func (t *TL_messages_stickerSetInstallResultArchive) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._sets = dc.Bytes(16)

}

// stickerSetCovered#6410a5d2
type TL_stickerSetCovered struct {
	_set   []byte
	_cover []byte
}

func (t *TL_stickerSetCovered) Set_set(_set []byte) {
	t._set = _set
}

func (t *TL_stickerSetCovered) Get_set() []byte {
	return t._set
}

func (t *TL_stickerSetCovered) Set_cover(_cover []byte) {
	t._cover = _cover
}

func (t *TL_stickerSetCovered) Get_cover() []byte {
	return t._cover
}

func New_TL_stickerSetCovered() *TL_stickerSetCovered {
	return &TL_stickerSetCovered{}
}

func (t *TL_stickerSetCovered) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_stickerSetCovered))
	ec.Bytes(t.Get_set())
	ec.Bytes(t.Get_cover())

	return ec.GetBuffer()
}

func (t *TL_stickerSetCovered) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._set = dc.Bytes(16)
	t._cover = dc.Bytes(16)

}

// stickerSetMultiCovered#3407e51b
type TL_stickerSetMultiCovered struct {
	_set    []byte
	_covers []byte
}

func (t *TL_stickerSetMultiCovered) Set_set(_set []byte) {
	t._set = _set
}

func (t *TL_stickerSetMultiCovered) Get_set() []byte {
	return t._set
}

func (t *TL_stickerSetMultiCovered) Set_covers(_covers []byte) {
	t._covers = _covers
}

func (t *TL_stickerSetMultiCovered) Get_covers() []byte {
	return t._covers
}

func New_TL_stickerSetMultiCovered() *TL_stickerSetMultiCovered {
	return &TL_stickerSetMultiCovered{}
}

func (t *TL_stickerSetMultiCovered) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_stickerSetMultiCovered))
	ec.Bytes(t.Get_set())
	ec.Bytes(t.Get_covers())

	return ec.GetBuffer()
}

func (t *TL_stickerSetMultiCovered) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._set = dc.Bytes(16)
	t._covers = dc.Bytes(16)

}

// maskCoords#aed6dbb2
type TL_maskCoords struct {
	_n    int32
	_x    float64
	_y    float64
	_zoom float64
}

func (t *TL_maskCoords) Set_n(_n int32) {
	t._n = _n
}

func (t *TL_maskCoords) Get_n() int32 {
	return t._n
}

func (t *TL_maskCoords) Set_x(_x float64) {
	t._x = _x
}

func (t *TL_maskCoords) Get_x() float64 {
	return t._x
}

func (t *TL_maskCoords) Set_y(_y float64) {
	t._y = _y
}

func (t *TL_maskCoords) Get_y() float64 {
	return t._y
}

func (t *TL_maskCoords) Set_zoom(_zoom float64) {
	t._zoom = _zoom
}

func (t *TL_maskCoords) Get_zoom() float64 {
	return t._zoom
}

func New_TL_maskCoords() *TL_maskCoords {
	return &TL_maskCoords{}
}

func (t *TL_maskCoords) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_maskCoords))
	ec.Int(t.Get_n())
	ec.Double(t.Get_x())
	ec.Double(t.Get_y())
	ec.Double(t.Get_zoom())

	return ec.GetBuffer()
}

func (t *TL_maskCoords) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._n = dc.Int()
	t._x = dc.Double()
	t._y = dc.Double()
	t._zoom = dc.Double()

}

// inputStickeredMediaPhoto#4a992157
type TL_inputStickeredMediaPhoto struct {
	_id []byte
}

func (t *TL_inputStickeredMediaPhoto) Set_id(_id []byte) {
	t._id = _id
}

func (t *TL_inputStickeredMediaPhoto) Get_id() []byte {
	return t._id
}

func New_TL_inputStickeredMediaPhoto() *TL_inputStickeredMediaPhoto {
	return &TL_inputStickeredMediaPhoto{}
}

func (t *TL_inputStickeredMediaPhoto) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputStickeredMediaPhoto))
	ec.Bytes(t.Get_id())

	return ec.GetBuffer()
}

func (t *TL_inputStickeredMediaPhoto) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Bytes(16)

}

// inputStickeredMediaDocument#0438865b
type TL_inputStickeredMediaDocument struct {
	_id []byte
}

func (t *TL_inputStickeredMediaDocument) Set_id(_id []byte) {
	t._id = _id
}

func (t *TL_inputStickeredMediaDocument) Get_id() []byte {
	return t._id
}

func New_TL_inputStickeredMediaDocument() *TL_inputStickeredMediaDocument {
	return &TL_inputStickeredMediaDocument{}
}

func (t *TL_inputStickeredMediaDocument) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputStickeredMediaDocument))
	ec.Bytes(t.Get_id())

	return ec.GetBuffer()
}

func (t *TL_inputStickeredMediaDocument) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Bytes(16)

}

// game#bdf9653b
type TL_game struct {
	_flags       []byte
	_id          int64
	_access_hash int64
	_short_name  string
	_title       string
	_description string
	_photo       []byte
	_document    []byte
}

func (t *TL_game) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_game) Get_flags() []byte {
	return t._flags
}

func (t *TL_game) Set_id(_id int64) {
	t._id = _id
}

func (t *TL_game) Get_id() int64 {
	return t._id
}

func (t *TL_game) Set_access_hash(_access_hash int64) {
	t._access_hash = _access_hash
}

func (t *TL_game) Get_access_hash() int64 {
	return t._access_hash
}

func (t *TL_game) Set_short_name(_short_name string) {
	t._short_name = _short_name
}

func (t *TL_game) Get_short_name() string {
	return t._short_name
}

func (t *TL_game) Set_title(_title string) {
	t._title = _title
}

func (t *TL_game) Get_title() string {
	return t._title
}

func (t *TL_game) Set_description(_description string) {
	t._description = _description
}

func (t *TL_game) Get_description() string {
	return t._description
}

func (t *TL_game) Set_photo(_photo []byte) {
	t._photo = _photo
}

func (t *TL_game) Get_photo() []byte {
	return t._photo
}

func (t *TL_game) Set_document(_document []byte) {
	t._document = _document
}

func (t *TL_game) Get_document() []byte {
	return t._document
}

func New_TL_game() *TL_game {
	return &TL_game{}
}

func (t *TL_game) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_game))
	ec.Long(t.Get_id())
	ec.Long(t.Get_access_hash())
	ec.String(t.Get_short_name())
	ec.String(t.Get_title())
	ec.String(t.Get_description())
	ec.Bytes(t.Get_photo())
	ec.Bytes(t.Get_document())

	return ec.GetBuffer()
}

func (t *TL_game) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Long()
	t._access_hash = dc.Long()
	t._short_name = dc.String()
	t._title = dc.String()
	t._description = dc.String()
	t._photo = dc.Bytes(16)
	t._document = dc.Bytes(16)

}

// inputGameID#032c3e77
type TL_inputGameID struct {
	_id          int64
	_access_hash int64
}

func (t *TL_inputGameID) Set_id(_id int64) {
	t._id = _id
}

func (t *TL_inputGameID) Get_id() int64 {
	return t._id
}

func (t *TL_inputGameID) Set_access_hash(_access_hash int64) {
	t._access_hash = _access_hash
}

func (t *TL_inputGameID) Get_access_hash() int64 {
	return t._access_hash
}

func New_TL_inputGameID() *TL_inputGameID {
	return &TL_inputGameID{}
}

func (t *TL_inputGameID) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputGameID))
	ec.Long(t.Get_id())
	ec.Long(t.Get_access_hash())

	return ec.GetBuffer()
}

func (t *TL_inputGameID) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Long()
	t._access_hash = dc.Long()

}

// inputGameShortName#c331e80a
type TL_inputGameShortName struct {
	_bot_id     []byte
	_short_name string
}

func (t *TL_inputGameShortName) Set_bot_id(_bot_id []byte) {
	t._bot_id = _bot_id
}

func (t *TL_inputGameShortName) Get_bot_id() []byte {
	return t._bot_id
}

func (t *TL_inputGameShortName) Set_short_name(_short_name string) {
	t._short_name = _short_name
}

func (t *TL_inputGameShortName) Get_short_name() string {
	return t._short_name
}

func New_TL_inputGameShortName() *TL_inputGameShortName {
	return &TL_inputGameShortName{}
}

func (t *TL_inputGameShortName) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputGameShortName))
	ec.Bytes(t.Get_bot_id())
	ec.String(t.Get_short_name())

	return ec.GetBuffer()
}

func (t *TL_inputGameShortName) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._bot_id = dc.Bytes(16)
	t._short_name = dc.String()

}

// highScore#58fffcd0
type TL_highScore struct {
	_pos     int32
	_user_id int32
	_score   int32
}

func (t *TL_highScore) Set_pos(_pos int32) {
	t._pos = _pos
}

func (t *TL_highScore) Get_pos() int32 {
	return t._pos
}

func (t *TL_highScore) Set_user_id(_user_id int32) {
	t._user_id = _user_id
}

func (t *TL_highScore) Get_user_id() int32 {
	return t._user_id
}

func (t *TL_highScore) Set_score(_score int32) {
	t._score = _score
}

func (t *TL_highScore) Get_score() int32 {
	return t._score
}

func New_TL_highScore() *TL_highScore {
	return &TL_highScore{}
}

func (t *TL_highScore) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_highScore))
	ec.Int(t.Get_pos())
	ec.Int(t.Get_user_id())
	ec.Int(t.Get_score())

	return ec.GetBuffer()
}

func (t *TL_highScore) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._pos = dc.Int()
	t._user_id = dc.Int()
	t._score = dc.Int()

}

// messages_highScores#9a3bfd99
type TL_messages_highScores struct {
	_scores []byte
	_users  []byte
}

func (t *TL_messages_highScores) Set_scores(_scores []byte) {
	t._scores = _scores
}

func (t *TL_messages_highScores) Get_scores() []byte {
	return t._scores
}

func (t *TL_messages_highScores) Set_users(_users []byte) {
	t._users = _users
}

func (t *TL_messages_highScores) Get_users() []byte {
	return t._users
}

func New_TL_messages_highScores() *TL_messages_highScores {
	return &TL_messages_highScores{}
}

func (t *TL_messages_highScores) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_highScores))
	ec.Bytes(t.Get_scores())
	ec.Bytes(t.Get_users())

	return ec.GetBuffer()
}

func (t *TL_messages_highScores) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._scores = dc.Bytes(16)
	t._users = dc.Bytes(16)

}

// textEmpty#dc3d824f
type TL_textEmpty struct {
}

func New_TL_textEmpty() *TL_textEmpty {
	return &TL_textEmpty{}
}

func (t *TL_textEmpty) Encode() []byte {
	return nil
}

func (t *TL_textEmpty) Decode(b []byte) {

}

// textPlain#744694e0
type TL_textPlain struct {
	_text string
}

func (t *TL_textPlain) Set_text(_text string) {
	t._text = _text
}

func (t *TL_textPlain) Get_text() string {
	return t._text
}

func New_TL_textPlain() *TL_textPlain {
	return &TL_textPlain{}
}

func (t *TL_textPlain) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_textPlain))
	ec.String(t.Get_text())

	return ec.GetBuffer()
}

func (t *TL_textPlain) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._text = dc.String()

}

// textBold#6724abc4
type TL_textBold struct {
	_text []byte
}

func (t *TL_textBold) Set_text(_text []byte) {
	t._text = _text
}

func (t *TL_textBold) Get_text() []byte {
	return t._text
}

func New_TL_textBold() *TL_textBold {
	return &TL_textBold{}
}

func (t *TL_textBold) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_textBold))
	ec.Bytes(t.Get_text())

	return ec.GetBuffer()
}

func (t *TL_textBold) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._text = dc.Bytes(16)

}

// textItalic#d912a59c
type TL_textItalic struct {
	_text []byte
}

func (t *TL_textItalic) Set_text(_text []byte) {
	t._text = _text
}

func (t *TL_textItalic) Get_text() []byte {
	return t._text
}

func New_TL_textItalic() *TL_textItalic {
	return &TL_textItalic{}
}

func (t *TL_textItalic) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_textItalic))
	ec.Bytes(t.Get_text())

	return ec.GetBuffer()
}

func (t *TL_textItalic) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._text = dc.Bytes(16)

}

// textUnderline#c12622c4
type TL_textUnderline struct {
	_text []byte
}

func (t *TL_textUnderline) Set_text(_text []byte) {
	t._text = _text
}

func (t *TL_textUnderline) Get_text() []byte {
	return t._text
}

func New_TL_textUnderline() *TL_textUnderline {
	return &TL_textUnderline{}
}

func (t *TL_textUnderline) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_textUnderline))
	ec.Bytes(t.Get_text())

	return ec.GetBuffer()
}

func (t *TL_textUnderline) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._text = dc.Bytes(16)

}

// textStrike#9bf8bb95
type TL_textStrike struct {
	_text []byte
}

func (t *TL_textStrike) Set_text(_text []byte) {
	t._text = _text
}

func (t *TL_textStrike) Get_text() []byte {
	return t._text
}

func New_TL_textStrike() *TL_textStrike {
	return &TL_textStrike{}
}

func (t *TL_textStrike) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_textStrike))
	ec.Bytes(t.Get_text())

	return ec.GetBuffer()
}

func (t *TL_textStrike) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._text = dc.Bytes(16)

}

// textFixed#6c3f19b9
type TL_textFixed struct {
	_text []byte
}

func (t *TL_textFixed) Set_text(_text []byte) {
	t._text = _text
}

func (t *TL_textFixed) Get_text() []byte {
	return t._text
}

func New_TL_textFixed() *TL_textFixed {
	return &TL_textFixed{}
}

func (t *TL_textFixed) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_textFixed))
	ec.Bytes(t.Get_text())

	return ec.GetBuffer()
}

func (t *TL_textFixed) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._text = dc.Bytes(16)

}

// textUrl#3c2884c1
type TL_textUrl struct {
	_text       []byte
	_url        string
	_webpage_id int64
}

func (t *TL_textUrl) Set_text(_text []byte) {
	t._text = _text
}

func (t *TL_textUrl) Get_text() []byte {
	return t._text
}

func (t *TL_textUrl) Set_url(_url string) {
	t._url = _url
}

func (t *TL_textUrl) Get_url() string {
	return t._url
}

func (t *TL_textUrl) Set_webpage_id(_webpage_id int64) {
	t._webpage_id = _webpage_id
}

func (t *TL_textUrl) Get_webpage_id() int64 {
	return t._webpage_id
}

func New_TL_textUrl() *TL_textUrl {
	return &TL_textUrl{}
}

func (t *TL_textUrl) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_textUrl))
	ec.Bytes(t.Get_text())
	ec.String(t.Get_url())
	ec.Long(t.Get_webpage_id())

	return ec.GetBuffer()
}

func (t *TL_textUrl) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._text = dc.Bytes(16)
	t._url = dc.String()
	t._webpage_id = dc.Long()

}

// textEmail#de5a0dd6
type TL_textEmail struct {
	_text  []byte
	_email string
}

func (t *TL_textEmail) Set_text(_text []byte) {
	t._text = _text
}

func (t *TL_textEmail) Get_text() []byte {
	return t._text
}

func (t *TL_textEmail) Set_email(_email string) {
	t._email = _email
}

func (t *TL_textEmail) Get_email() string {
	return t._email
}

func New_TL_textEmail() *TL_textEmail {
	return &TL_textEmail{}
}

func (t *TL_textEmail) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_textEmail))
	ec.Bytes(t.Get_text())
	ec.String(t.Get_email())

	return ec.GetBuffer()
}

func (t *TL_textEmail) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._text = dc.Bytes(16)
	t._email = dc.String()

}

// textConcat#7e6260d7
type TL_textConcat struct {
	_texts []byte
}

func (t *TL_textConcat) Set_texts(_texts []byte) {
	t._texts = _texts
}

func (t *TL_textConcat) Get_texts() []byte {
	return t._texts
}

func New_TL_textConcat() *TL_textConcat {
	return &TL_textConcat{}
}

func (t *TL_textConcat) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_textConcat))
	ec.Bytes(t.Get_texts())

	return ec.GetBuffer()
}

func (t *TL_textConcat) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._texts = dc.Bytes(16)

}

// pageBlockUnsupported#13567e8a
type TL_pageBlockUnsupported struct {
}

func New_TL_pageBlockUnsupported() *TL_pageBlockUnsupported {
	return &TL_pageBlockUnsupported{}
}

func (t *TL_pageBlockUnsupported) Encode() []byte {
	return nil
}

func (t *TL_pageBlockUnsupported) Decode(b []byte) {

}

// pageBlockTitle#70abc3fd
type TL_pageBlockTitle struct {
	_text []byte
}

func (t *TL_pageBlockTitle) Set_text(_text []byte) {
	t._text = _text
}

func (t *TL_pageBlockTitle) Get_text() []byte {
	return t._text
}

func New_TL_pageBlockTitle() *TL_pageBlockTitle {
	return &TL_pageBlockTitle{}
}

func (t *TL_pageBlockTitle) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_pageBlockTitle))
	ec.Bytes(t.Get_text())

	return ec.GetBuffer()
}

func (t *TL_pageBlockTitle) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._text = dc.Bytes(16)

}

// pageBlockSubtitle#8ffa9a1f
type TL_pageBlockSubtitle struct {
	_text []byte
}

func (t *TL_pageBlockSubtitle) Set_text(_text []byte) {
	t._text = _text
}

func (t *TL_pageBlockSubtitle) Get_text() []byte {
	return t._text
}

func New_TL_pageBlockSubtitle() *TL_pageBlockSubtitle {
	return &TL_pageBlockSubtitle{}
}

func (t *TL_pageBlockSubtitle) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_pageBlockSubtitle))
	ec.Bytes(t.Get_text())

	return ec.GetBuffer()
}

func (t *TL_pageBlockSubtitle) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._text = dc.Bytes(16)

}

// pageBlockAuthorDate#baafe5e0
type TL_pageBlockAuthorDate struct {
	_author         []byte
	_published_date int32
}

func (t *TL_pageBlockAuthorDate) Set_author(_author []byte) {
	t._author = _author
}

func (t *TL_pageBlockAuthorDate) Get_author() []byte {
	return t._author
}

func (t *TL_pageBlockAuthorDate) Set_published_date(_published_date int32) {
	t._published_date = _published_date
}

func (t *TL_pageBlockAuthorDate) Get_published_date() int32 {
	return t._published_date
}

func New_TL_pageBlockAuthorDate() *TL_pageBlockAuthorDate {
	return &TL_pageBlockAuthorDate{}
}

func (t *TL_pageBlockAuthorDate) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_pageBlockAuthorDate))
	ec.Bytes(t.Get_author())
	ec.Int(t.Get_published_date())

	return ec.GetBuffer()
}

func (t *TL_pageBlockAuthorDate) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._author = dc.Bytes(16)
	t._published_date = dc.Int()

}

// pageBlockHeader#bfd064ec
type TL_pageBlockHeader struct {
	_text []byte
}

func (t *TL_pageBlockHeader) Set_text(_text []byte) {
	t._text = _text
}

func (t *TL_pageBlockHeader) Get_text() []byte {
	return t._text
}

func New_TL_pageBlockHeader() *TL_pageBlockHeader {
	return &TL_pageBlockHeader{}
}

func (t *TL_pageBlockHeader) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_pageBlockHeader))
	ec.Bytes(t.Get_text())

	return ec.GetBuffer()
}

func (t *TL_pageBlockHeader) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._text = dc.Bytes(16)

}

// pageBlockSubheader#f12bb6e1
type TL_pageBlockSubheader struct {
	_text []byte
}

func (t *TL_pageBlockSubheader) Set_text(_text []byte) {
	t._text = _text
}

func (t *TL_pageBlockSubheader) Get_text() []byte {
	return t._text
}

func New_TL_pageBlockSubheader() *TL_pageBlockSubheader {
	return &TL_pageBlockSubheader{}
}

func (t *TL_pageBlockSubheader) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_pageBlockSubheader))
	ec.Bytes(t.Get_text())

	return ec.GetBuffer()
}

func (t *TL_pageBlockSubheader) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._text = dc.Bytes(16)

}

// pageBlockParagraph#467a0766
type TL_pageBlockParagraph struct {
	_text []byte
}

func (t *TL_pageBlockParagraph) Set_text(_text []byte) {
	t._text = _text
}

func (t *TL_pageBlockParagraph) Get_text() []byte {
	return t._text
}

func New_TL_pageBlockParagraph() *TL_pageBlockParagraph {
	return &TL_pageBlockParagraph{}
}

func (t *TL_pageBlockParagraph) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_pageBlockParagraph))
	ec.Bytes(t.Get_text())

	return ec.GetBuffer()
}

func (t *TL_pageBlockParagraph) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._text = dc.Bytes(16)

}

// pageBlockPreformatted#c070d93e
type TL_pageBlockPreformatted struct {
	_text     []byte
	_language string
}

func (t *TL_pageBlockPreformatted) Set_text(_text []byte) {
	t._text = _text
}

func (t *TL_pageBlockPreformatted) Get_text() []byte {
	return t._text
}

func (t *TL_pageBlockPreformatted) Set_language(_language string) {
	t._language = _language
}

func (t *TL_pageBlockPreformatted) Get_language() string {
	return t._language
}

func New_TL_pageBlockPreformatted() *TL_pageBlockPreformatted {
	return &TL_pageBlockPreformatted{}
}

func (t *TL_pageBlockPreformatted) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_pageBlockPreformatted))
	ec.Bytes(t.Get_text())
	ec.String(t.Get_language())

	return ec.GetBuffer()
}

func (t *TL_pageBlockPreformatted) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._text = dc.Bytes(16)
	t._language = dc.String()

}

// pageBlockFooter#48870999
type TL_pageBlockFooter struct {
	_text []byte
}

func (t *TL_pageBlockFooter) Set_text(_text []byte) {
	t._text = _text
}

func (t *TL_pageBlockFooter) Get_text() []byte {
	return t._text
}

func New_TL_pageBlockFooter() *TL_pageBlockFooter {
	return &TL_pageBlockFooter{}
}

func (t *TL_pageBlockFooter) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_pageBlockFooter))
	ec.Bytes(t.Get_text())

	return ec.GetBuffer()
}

func (t *TL_pageBlockFooter) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._text = dc.Bytes(16)

}

// pageBlockDivider#db20b188
type TL_pageBlockDivider struct {
}

func New_TL_pageBlockDivider() *TL_pageBlockDivider {
	return &TL_pageBlockDivider{}
}

func (t *TL_pageBlockDivider) Encode() []byte {
	return nil
}

func (t *TL_pageBlockDivider) Decode(b []byte) {

}

// pageBlockAnchor#ce0d37b0
type TL_pageBlockAnchor struct {
	_name string
}

func (t *TL_pageBlockAnchor) Set_name(_name string) {
	t._name = _name
}

func (t *TL_pageBlockAnchor) Get_name() string {
	return t._name
}

func New_TL_pageBlockAnchor() *TL_pageBlockAnchor {
	return &TL_pageBlockAnchor{}
}

func (t *TL_pageBlockAnchor) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_pageBlockAnchor))
	ec.String(t.Get_name())

	return ec.GetBuffer()
}

func (t *TL_pageBlockAnchor) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._name = dc.String()

}

// pageBlockList#3a58c7f4
type TL_pageBlockList struct {
	_ordered bool
	_items   []byte
}

func (t *TL_pageBlockList) Set_ordered(_ordered bool) {
	t._ordered = _ordered
}

func (t *TL_pageBlockList) Get_ordered() bool {
	return t._ordered
}

func (t *TL_pageBlockList) Set_items(_items []byte) {
	t._items = _items
}

func (t *TL_pageBlockList) Get_items() []byte {
	return t._items
}

func New_TL_pageBlockList() *TL_pageBlockList {
	return &TL_pageBlockList{}
}

func (t *TL_pageBlockList) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_pageBlockList))
	ec.Bool(t.Get_ordered())
	ec.Bytes(t.Get_items())

	return ec.GetBuffer()
}

func (t *TL_pageBlockList) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._ordered = dc.Bool()
	t._items = dc.Bytes(16)

}

// pageBlockBlockquote#263d7c26
type TL_pageBlockBlockquote struct {
	_text    []byte
	_caption []byte
}

func (t *TL_pageBlockBlockquote) Set_text(_text []byte) {
	t._text = _text
}

func (t *TL_pageBlockBlockquote) Get_text() []byte {
	return t._text
}

func (t *TL_pageBlockBlockquote) Set_caption(_caption []byte) {
	t._caption = _caption
}

func (t *TL_pageBlockBlockquote) Get_caption() []byte {
	return t._caption
}

func New_TL_pageBlockBlockquote() *TL_pageBlockBlockquote {
	return &TL_pageBlockBlockquote{}
}

func (t *TL_pageBlockBlockquote) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_pageBlockBlockquote))
	ec.Bytes(t.Get_text())
	ec.Bytes(t.Get_caption())

	return ec.GetBuffer()
}

func (t *TL_pageBlockBlockquote) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._text = dc.Bytes(16)
	t._caption = dc.Bytes(16)

}

// pageBlockPullquote#4f4456d3
type TL_pageBlockPullquote struct {
	_text    []byte
	_caption []byte
}

func (t *TL_pageBlockPullquote) Set_text(_text []byte) {
	t._text = _text
}

func (t *TL_pageBlockPullquote) Get_text() []byte {
	return t._text
}

func (t *TL_pageBlockPullquote) Set_caption(_caption []byte) {
	t._caption = _caption
}

func (t *TL_pageBlockPullquote) Get_caption() []byte {
	return t._caption
}

func New_TL_pageBlockPullquote() *TL_pageBlockPullquote {
	return &TL_pageBlockPullquote{}
}

func (t *TL_pageBlockPullquote) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_pageBlockPullquote))
	ec.Bytes(t.Get_text())
	ec.Bytes(t.Get_caption())

	return ec.GetBuffer()
}

func (t *TL_pageBlockPullquote) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._text = dc.Bytes(16)
	t._caption = dc.Bytes(16)

}

// pageBlockPhoto#e9c69982
type TL_pageBlockPhoto struct {
	_photo_id int64
	_caption  []byte
}

func (t *TL_pageBlockPhoto) Set_photo_id(_photo_id int64) {
	t._photo_id = _photo_id
}

func (t *TL_pageBlockPhoto) Get_photo_id() int64 {
	return t._photo_id
}

func (t *TL_pageBlockPhoto) Set_caption(_caption []byte) {
	t._caption = _caption
}

func (t *TL_pageBlockPhoto) Get_caption() []byte {
	return t._caption
}

func New_TL_pageBlockPhoto() *TL_pageBlockPhoto {
	return &TL_pageBlockPhoto{}
}

func (t *TL_pageBlockPhoto) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_pageBlockPhoto))
	ec.Long(t.Get_photo_id())
	ec.Bytes(t.Get_caption())

	return ec.GetBuffer()
}

func (t *TL_pageBlockPhoto) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._photo_id = dc.Long()
	t._caption = dc.Bytes(16)

}

// pageBlockVideo#d9d71866
type TL_pageBlockVideo struct {
	_flags    []byte
	_autoplay []byte
	_loop     []byte
	_video_id int64
	_caption  []byte
}

func (t *TL_pageBlockVideo) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_pageBlockVideo) Get_flags() []byte {
	return t._flags
}

func (t *TL_pageBlockVideo) Set_autoplay(_autoplay []byte) {
	t._autoplay = _autoplay
}

func (t *TL_pageBlockVideo) Get_autoplay() []byte {
	return t._autoplay
}

func (t *TL_pageBlockVideo) Set_loop(_loop []byte) {
	t._loop = _loop
}

func (t *TL_pageBlockVideo) Get_loop() []byte {
	return t._loop
}

func (t *TL_pageBlockVideo) Set_video_id(_video_id int64) {
	t._video_id = _video_id
}

func (t *TL_pageBlockVideo) Get_video_id() int64 {
	return t._video_id
}

func (t *TL_pageBlockVideo) Set_caption(_caption []byte) {
	t._caption = _caption
}

func (t *TL_pageBlockVideo) Get_caption() []byte {
	return t._caption
}

func New_TL_pageBlockVideo() *TL_pageBlockVideo {
	return &TL_pageBlockVideo{}
}

func (t *TL_pageBlockVideo) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_pageBlockVideo))
	ec.Bytes(t.Get_autoplay())
	ec.Bytes(t.Get_loop())
	ec.Long(t.Get_video_id())
	ec.Bytes(t.Get_caption())

	return ec.GetBuffer()
}

func (t *TL_pageBlockVideo) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._autoplay = dc.Bytes(16)
	t._loop = dc.Bytes(16)
	t._video_id = dc.Long()
	t._caption = dc.Bytes(16)

}

// pageBlockCover#39f23300
type TL_pageBlockCover struct {
	_cover []byte
}

func (t *TL_pageBlockCover) Set_cover(_cover []byte) {
	t._cover = _cover
}

func (t *TL_pageBlockCover) Get_cover() []byte {
	return t._cover
}

func New_TL_pageBlockCover() *TL_pageBlockCover {
	return &TL_pageBlockCover{}
}

func (t *TL_pageBlockCover) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_pageBlockCover))
	ec.Bytes(t.Get_cover())

	return ec.GetBuffer()
}

func (t *TL_pageBlockCover) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._cover = dc.Bytes(16)

}

// pageBlockEmbed#cde200d1
type TL_pageBlockEmbed struct {
	_flags           []byte
	_full_width      []byte
	_allow_scrolling []byte
	_url             []byte
	_html            []byte
	_poster_photo_id []byte
	_w               int32
	_h               int32
	_caption         []byte
}

func (t *TL_pageBlockEmbed) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_pageBlockEmbed) Get_flags() []byte {
	return t._flags
}

func (t *TL_pageBlockEmbed) Set_full_width(_full_width []byte) {
	t._full_width = _full_width
}

func (t *TL_pageBlockEmbed) Get_full_width() []byte {
	return t._full_width
}

func (t *TL_pageBlockEmbed) Set_allow_scrolling(_allow_scrolling []byte) {
	t._allow_scrolling = _allow_scrolling
}

func (t *TL_pageBlockEmbed) Get_allow_scrolling() []byte {
	return t._allow_scrolling
}

func (t *TL_pageBlockEmbed) Set_url(_url []byte) {
	t._url = _url
}

func (t *TL_pageBlockEmbed) Get_url() []byte {
	return t._url
}

func (t *TL_pageBlockEmbed) Set_html(_html []byte) {
	t._html = _html
}

func (t *TL_pageBlockEmbed) Get_html() []byte {
	return t._html
}

func (t *TL_pageBlockEmbed) Set_poster_photo_id(_poster_photo_id []byte) {
	t._poster_photo_id = _poster_photo_id
}

func (t *TL_pageBlockEmbed) Get_poster_photo_id() []byte {
	return t._poster_photo_id
}

func (t *TL_pageBlockEmbed) Set_w(_w int32) {
	t._w = _w
}

func (t *TL_pageBlockEmbed) Get_w() int32 {
	return t._w
}

func (t *TL_pageBlockEmbed) Set_h(_h int32) {
	t._h = _h
}

func (t *TL_pageBlockEmbed) Get_h() int32 {
	return t._h
}

func (t *TL_pageBlockEmbed) Set_caption(_caption []byte) {
	t._caption = _caption
}

func (t *TL_pageBlockEmbed) Get_caption() []byte {
	return t._caption
}

func New_TL_pageBlockEmbed() *TL_pageBlockEmbed {
	return &TL_pageBlockEmbed{}
}

func (t *TL_pageBlockEmbed) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_pageBlockEmbed))
	ec.Bytes(t.Get_full_width())
	ec.Bytes(t.Get_allow_scrolling())
	ec.Bytes(t.Get_url())
	ec.Bytes(t.Get_html())
	ec.Bytes(t.Get_poster_photo_id())
	ec.Int(t.Get_w())
	ec.Int(t.Get_h())
	ec.Bytes(t.Get_caption())

	return ec.GetBuffer()
}

func (t *TL_pageBlockEmbed) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._full_width = dc.Bytes(16)
	t._allow_scrolling = dc.Bytes(16)
	t._url = dc.Bytes(16)
	t._html = dc.Bytes(16)
	t._poster_photo_id = dc.Bytes(16)
	t._w = dc.Int()
	t._h = dc.Int()
	t._caption = dc.Bytes(16)

}

// pageBlockEmbedPost#292c7be9
type TL_pageBlockEmbedPost struct {
	_url             string
	_webpage_id      int64
	_author_photo_id int64
	_author          string
	_date            int32
	_blocks          []byte
	_caption         []byte
}

func (t *TL_pageBlockEmbedPost) Set_url(_url string) {
	t._url = _url
}

func (t *TL_pageBlockEmbedPost) Get_url() string {
	return t._url
}

func (t *TL_pageBlockEmbedPost) Set_webpage_id(_webpage_id int64) {
	t._webpage_id = _webpage_id
}

func (t *TL_pageBlockEmbedPost) Get_webpage_id() int64 {
	return t._webpage_id
}

func (t *TL_pageBlockEmbedPost) Set_author_photo_id(_author_photo_id int64) {
	t._author_photo_id = _author_photo_id
}

func (t *TL_pageBlockEmbedPost) Get_author_photo_id() int64 {
	return t._author_photo_id
}

func (t *TL_pageBlockEmbedPost) Set_author(_author string) {
	t._author = _author
}

func (t *TL_pageBlockEmbedPost) Get_author() string {
	return t._author
}

func (t *TL_pageBlockEmbedPost) Set_date(_date int32) {
	t._date = _date
}

func (t *TL_pageBlockEmbedPost) Get_date() int32 {
	return t._date
}

func (t *TL_pageBlockEmbedPost) Set_blocks(_blocks []byte) {
	t._blocks = _blocks
}

func (t *TL_pageBlockEmbedPost) Get_blocks() []byte {
	return t._blocks
}

func (t *TL_pageBlockEmbedPost) Set_caption(_caption []byte) {
	t._caption = _caption
}

func (t *TL_pageBlockEmbedPost) Get_caption() []byte {
	return t._caption
}

func New_TL_pageBlockEmbedPost() *TL_pageBlockEmbedPost {
	return &TL_pageBlockEmbedPost{}
}

func (t *TL_pageBlockEmbedPost) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_pageBlockEmbedPost))
	ec.String(t.Get_url())
	ec.Long(t.Get_webpage_id())
	ec.Long(t.Get_author_photo_id())
	ec.String(t.Get_author())
	ec.Int(t.Get_date())
	ec.Bytes(t.Get_blocks())
	ec.Bytes(t.Get_caption())

	return ec.GetBuffer()
}

func (t *TL_pageBlockEmbedPost) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._url = dc.String()
	t._webpage_id = dc.Long()
	t._author_photo_id = dc.Long()
	t._author = dc.String()
	t._date = dc.Int()
	t._blocks = dc.Bytes(16)
	t._caption = dc.Bytes(16)

}

// pageBlockCollage#08b31c4f
type TL_pageBlockCollage struct {
	_items   []byte
	_caption []byte
}

func (t *TL_pageBlockCollage) Set_items(_items []byte) {
	t._items = _items
}

func (t *TL_pageBlockCollage) Get_items() []byte {
	return t._items
}

func (t *TL_pageBlockCollage) Set_caption(_caption []byte) {
	t._caption = _caption
}

func (t *TL_pageBlockCollage) Get_caption() []byte {
	return t._caption
}

func New_TL_pageBlockCollage() *TL_pageBlockCollage {
	return &TL_pageBlockCollage{}
}

func (t *TL_pageBlockCollage) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_pageBlockCollage))
	ec.Bytes(t.Get_items())
	ec.Bytes(t.Get_caption())

	return ec.GetBuffer()
}

func (t *TL_pageBlockCollage) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._items = dc.Bytes(16)
	t._caption = dc.Bytes(16)

}

// pageBlockSlideshow#130c8963
type TL_pageBlockSlideshow struct {
	_items   []byte
	_caption []byte
}

func (t *TL_pageBlockSlideshow) Set_items(_items []byte) {
	t._items = _items
}

func (t *TL_pageBlockSlideshow) Get_items() []byte {
	return t._items
}

func (t *TL_pageBlockSlideshow) Set_caption(_caption []byte) {
	t._caption = _caption
}

func (t *TL_pageBlockSlideshow) Get_caption() []byte {
	return t._caption
}

func New_TL_pageBlockSlideshow() *TL_pageBlockSlideshow {
	return &TL_pageBlockSlideshow{}
}

func (t *TL_pageBlockSlideshow) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_pageBlockSlideshow))
	ec.Bytes(t.Get_items())
	ec.Bytes(t.Get_caption())

	return ec.GetBuffer()
}

func (t *TL_pageBlockSlideshow) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._items = dc.Bytes(16)
	t._caption = dc.Bytes(16)

}

// pageBlockChannel#ef1751b5
type TL_pageBlockChannel struct {
	_channel []byte
}

func (t *TL_pageBlockChannel) Set_channel(_channel []byte) {
	t._channel = _channel
}

func (t *TL_pageBlockChannel) Get_channel() []byte {
	return t._channel
}

func New_TL_pageBlockChannel() *TL_pageBlockChannel {
	return &TL_pageBlockChannel{}
}

func (t *TL_pageBlockChannel) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_pageBlockChannel))
	ec.Bytes(t.Get_channel())

	return ec.GetBuffer()
}

func (t *TL_pageBlockChannel) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._channel = dc.Bytes(16)

}

// pageBlockAudio#31b81a7f
type TL_pageBlockAudio struct {
	_audio_id int64
	_caption  []byte
}

func (t *TL_pageBlockAudio) Set_audio_id(_audio_id int64) {
	t._audio_id = _audio_id
}

func (t *TL_pageBlockAudio) Get_audio_id() int64 {
	return t._audio_id
}

func (t *TL_pageBlockAudio) Set_caption(_caption []byte) {
	t._caption = _caption
}

func (t *TL_pageBlockAudio) Get_caption() []byte {
	return t._caption
}

func New_TL_pageBlockAudio() *TL_pageBlockAudio {
	return &TL_pageBlockAudio{}
}

func (t *TL_pageBlockAudio) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_pageBlockAudio))
	ec.Long(t.Get_audio_id())
	ec.Bytes(t.Get_caption())

	return ec.GetBuffer()
}

func (t *TL_pageBlockAudio) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._audio_id = dc.Long()
	t._caption = dc.Bytes(16)

}

// pagePart#8e3f9ebe
type TL_pagePart struct {
	_blocks    []byte
	_photos    []byte
	_documents []byte
}

func (t *TL_pagePart) Set_blocks(_blocks []byte) {
	t._blocks = _blocks
}

func (t *TL_pagePart) Get_blocks() []byte {
	return t._blocks
}

func (t *TL_pagePart) Set_photos(_photos []byte) {
	t._photos = _photos
}

func (t *TL_pagePart) Get_photos() []byte {
	return t._photos
}

func (t *TL_pagePart) Set_documents(_documents []byte) {
	t._documents = _documents
}

func (t *TL_pagePart) Get_documents() []byte {
	return t._documents
}

func New_TL_pagePart() *TL_pagePart {
	return &TL_pagePart{}
}

func (t *TL_pagePart) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_pagePart))
	ec.Bytes(t.Get_blocks())
	ec.Bytes(t.Get_photos())
	ec.Bytes(t.Get_documents())

	return ec.GetBuffer()
}

func (t *TL_pagePart) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._blocks = dc.Bytes(16)
	t._photos = dc.Bytes(16)
	t._documents = dc.Bytes(16)

}

// pageFull#556ec7aa
type TL_pageFull struct {
	_blocks    []byte
	_photos    []byte
	_documents []byte
}

func (t *TL_pageFull) Set_blocks(_blocks []byte) {
	t._blocks = _blocks
}

func (t *TL_pageFull) Get_blocks() []byte {
	return t._blocks
}

func (t *TL_pageFull) Set_photos(_photos []byte) {
	t._photos = _photos
}

func (t *TL_pageFull) Get_photos() []byte {
	return t._photos
}

func (t *TL_pageFull) Set_documents(_documents []byte) {
	t._documents = _documents
}

func (t *TL_pageFull) Get_documents() []byte {
	return t._documents
}

func New_TL_pageFull() *TL_pageFull {
	return &TL_pageFull{}
}

func (t *TL_pageFull) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_pageFull))
	ec.Bytes(t.Get_blocks())
	ec.Bytes(t.Get_photos())
	ec.Bytes(t.Get_documents())

	return ec.GetBuffer()
}

func (t *TL_pageFull) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._blocks = dc.Bytes(16)
	t._photos = dc.Bytes(16)
	t._documents = dc.Bytes(16)

}

// phoneCallDiscardReasonMissed#85e42301
type TL_phoneCallDiscardReasonMissed struct {
}

func New_TL_phoneCallDiscardReasonMissed() *TL_phoneCallDiscardReasonMissed {
	return &TL_phoneCallDiscardReasonMissed{}
}

func (t *TL_phoneCallDiscardReasonMissed) Encode() []byte {
	return nil
}

func (t *TL_phoneCallDiscardReasonMissed) Decode(b []byte) {

}

// phoneCallDiscardReasonDisconnect#e095c1a0
type TL_phoneCallDiscardReasonDisconnect struct {
}

func New_TL_phoneCallDiscardReasonDisconnect() *TL_phoneCallDiscardReasonDisconnect {
	return &TL_phoneCallDiscardReasonDisconnect{}
}

func (t *TL_phoneCallDiscardReasonDisconnect) Encode() []byte {
	return nil
}

func (t *TL_phoneCallDiscardReasonDisconnect) Decode(b []byte) {

}

// phoneCallDiscardReasonHangup#57adc690
type TL_phoneCallDiscardReasonHangup struct {
}

func New_TL_phoneCallDiscardReasonHangup() *TL_phoneCallDiscardReasonHangup {
	return &TL_phoneCallDiscardReasonHangup{}
}

func (t *TL_phoneCallDiscardReasonHangup) Encode() []byte {
	return nil
}

func (t *TL_phoneCallDiscardReasonHangup) Decode(b []byte) {

}

// phoneCallDiscardReasonBusy#faf7e8c9
type TL_phoneCallDiscardReasonBusy struct {
}

func New_TL_phoneCallDiscardReasonBusy() *TL_phoneCallDiscardReasonBusy {
	return &TL_phoneCallDiscardReasonBusy{}
}

func (t *TL_phoneCallDiscardReasonBusy) Encode() []byte {
	return nil
}

func (t *TL_phoneCallDiscardReasonBusy) Decode(b []byte) {

}

// dataJSON#7d748d04
type TL_dataJSON struct {
	_data string
}

func (t *TL_dataJSON) Set_data(_data string) {
	t._data = _data
}

func (t *TL_dataJSON) Get_data() string {
	return t._data
}

func New_TL_dataJSON() *TL_dataJSON {
	return &TL_dataJSON{}
}

func (t *TL_dataJSON) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_dataJSON))
	ec.String(t.Get_data())

	return ec.GetBuffer()
}

func (t *TL_dataJSON) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._data = dc.String()

}

// labeledPrice#cb296bf8
type TL_labeledPrice struct {
	_label  string
	_amount int64
}

func (t *TL_labeledPrice) Set_label(_label string) {
	t._label = _label
}

func (t *TL_labeledPrice) Get_label() string {
	return t._label
}

func (t *TL_labeledPrice) Set_amount(_amount int64) {
	t._amount = _amount
}

func (t *TL_labeledPrice) Get_amount() int64 {
	return t._amount
}

func New_TL_labeledPrice() *TL_labeledPrice {
	return &TL_labeledPrice{}
}

func (t *TL_labeledPrice) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_labeledPrice))
	ec.String(t.Get_label())
	ec.Long(t.Get_amount())

	return ec.GetBuffer()
}

func (t *TL_labeledPrice) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._label = dc.String()
	t._amount = dc.Long()

}

// invoice#c30aa358
type TL_invoice struct {
	_flags                      []byte
	_test                       []byte
	_name_requested             []byte
	_phone_requested            []byte
	_email_requested            []byte
	_shipping_address_requested []byte
	_flexible                   []byte
	_phone_to_provider          []byte
	_email_to_provider          []byte
	_currency                   string
	_prices                     []byte
}

func (t *TL_invoice) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_invoice) Get_flags() []byte {
	return t._flags
}

func (t *TL_invoice) Set_test(_test []byte) {
	t._test = _test
}

func (t *TL_invoice) Get_test() []byte {
	return t._test
}

func (t *TL_invoice) Set_name_requested(_name_requested []byte) {
	t._name_requested = _name_requested
}

func (t *TL_invoice) Get_name_requested() []byte {
	return t._name_requested
}

func (t *TL_invoice) Set_phone_requested(_phone_requested []byte) {
	t._phone_requested = _phone_requested
}

func (t *TL_invoice) Get_phone_requested() []byte {
	return t._phone_requested
}

func (t *TL_invoice) Set_email_requested(_email_requested []byte) {
	t._email_requested = _email_requested
}

func (t *TL_invoice) Get_email_requested() []byte {
	return t._email_requested
}

func (t *TL_invoice) Set_shipping_address_requested(_shipping_address_requested []byte) {
	t._shipping_address_requested = _shipping_address_requested
}

func (t *TL_invoice) Get_shipping_address_requested() []byte {
	return t._shipping_address_requested
}

func (t *TL_invoice) Set_flexible(_flexible []byte) {
	t._flexible = _flexible
}

func (t *TL_invoice) Get_flexible() []byte {
	return t._flexible
}

func (t *TL_invoice) Set_phone_to_provider(_phone_to_provider []byte) {
	t._phone_to_provider = _phone_to_provider
}

func (t *TL_invoice) Get_phone_to_provider() []byte {
	return t._phone_to_provider
}

func (t *TL_invoice) Set_email_to_provider(_email_to_provider []byte) {
	t._email_to_provider = _email_to_provider
}

func (t *TL_invoice) Get_email_to_provider() []byte {
	return t._email_to_provider
}

func (t *TL_invoice) Set_currency(_currency string) {
	t._currency = _currency
}

func (t *TL_invoice) Get_currency() string {
	return t._currency
}

func (t *TL_invoice) Set_prices(_prices []byte) {
	t._prices = _prices
}

func (t *TL_invoice) Get_prices() []byte {
	return t._prices
}

func New_TL_invoice() *TL_invoice {
	return &TL_invoice{}
}

func (t *TL_invoice) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_invoice))
	ec.Bytes(t.Get_test())
	ec.Bytes(t.Get_name_requested())
	ec.Bytes(t.Get_phone_requested())
	ec.Bytes(t.Get_email_requested())
	ec.Bytes(t.Get_shipping_address_requested())
	ec.Bytes(t.Get_flexible())
	ec.Bytes(t.Get_phone_to_provider())
	ec.Bytes(t.Get_email_to_provider())
	ec.String(t.Get_currency())
	ec.Bytes(t.Get_prices())

	return ec.GetBuffer()
}

func (t *TL_invoice) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._test = dc.Bytes(16)
	t._name_requested = dc.Bytes(16)
	t._phone_requested = dc.Bytes(16)
	t._email_requested = dc.Bytes(16)
	t._shipping_address_requested = dc.Bytes(16)
	t._flexible = dc.Bytes(16)
	t._phone_to_provider = dc.Bytes(16)
	t._email_to_provider = dc.Bytes(16)
	t._currency = dc.String()
	t._prices = dc.Bytes(16)

}

// paymentCharge#ea02c27e
type TL_paymentCharge struct {
	_id                 string
	_provider_charge_id string
}

func (t *TL_paymentCharge) Set_id(_id string) {
	t._id = _id
}

func (t *TL_paymentCharge) Get_id() string {
	return t._id
}

func (t *TL_paymentCharge) Set_provider_charge_id(_provider_charge_id string) {
	t._provider_charge_id = _provider_charge_id
}

func (t *TL_paymentCharge) Get_provider_charge_id() string {
	return t._provider_charge_id
}

func New_TL_paymentCharge() *TL_paymentCharge {
	return &TL_paymentCharge{}
}

func (t *TL_paymentCharge) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_paymentCharge))
	ec.String(t.Get_id())
	ec.String(t.Get_provider_charge_id())

	return ec.GetBuffer()
}

func (t *TL_paymentCharge) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.String()
	t._provider_charge_id = dc.String()

}

// postAddress#1e8caaeb
type TL_postAddress struct {
	_street_line1 string
	_street_line2 string
	_city         string
	_state        string
	_country_iso2 string
	_post_code    string
}

func (t *TL_postAddress) Set_street_line1(_street_line1 string) {
	t._street_line1 = _street_line1
}

func (t *TL_postAddress) Get_street_line1() string {
	return t._street_line1
}

func (t *TL_postAddress) Set_street_line2(_street_line2 string) {
	t._street_line2 = _street_line2
}

func (t *TL_postAddress) Get_street_line2() string {
	return t._street_line2
}

func (t *TL_postAddress) Set_city(_city string) {
	t._city = _city
}

func (t *TL_postAddress) Get_city() string {
	return t._city
}

func (t *TL_postAddress) Set_state(_state string) {
	t._state = _state
}

func (t *TL_postAddress) Get_state() string {
	return t._state
}

func (t *TL_postAddress) Set_country_iso2(_country_iso2 string) {
	t._country_iso2 = _country_iso2
}

func (t *TL_postAddress) Get_country_iso2() string {
	return t._country_iso2
}

func (t *TL_postAddress) Set_post_code(_post_code string) {
	t._post_code = _post_code
}

func (t *TL_postAddress) Get_post_code() string {
	return t._post_code
}

func New_TL_postAddress() *TL_postAddress {
	return &TL_postAddress{}
}

func (t *TL_postAddress) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_postAddress))
	ec.String(t.Get_street_line1())
	ec.String(t.Get_street_line2())
	ec.String(t.Get_city())
	ec.String(t.Get_state())
	ec.String(t.Get_country_iso2())
	ec.String(t.Get_post_code())

	return ec.GetBuffer()
}

func (t *TL_postAddress) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._street_line1 = dc.String()
	t._street_line2 = dc.String()
	t._city = dc.String()
	t._state = dc.String()
	t._country_iso2 = dc.String()
	t._post_code = dc.String()

}

// paymentRequestedInfo#909c3f94
type TL_paymentRequestedInfo struct {
	_flags            []byte
	_name             []byte
	_phone            []byte
	_email            []byte
	_shipping_address []byte
}

func (t *TL_paymentRequestedInfo) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_paymentRequestedInfo) Get_flags() []byte {
	return t._flags
}

func (t *TL_paymentRequestedInfo) Set_name(_name []byte) {
	t._name = _name
}

func (t *TL_paymentRequestedInfo) Get_name() []byte {
	return t._name
}

func (t *TL_paymentRequestedInfo) Set_phone(_phone []byte) {
	t._phone = _phone
}

func (t *TL_paymentRequestedInfo) Get_phone() []byte {
	return t._phone
}

func (t *TL_paymentRequestedInfo) Set_email(_email []byte) {
	t._email = _email
}

func (t *TL_paymentRequestedInfo) Get_email() []byte {
	return t._email
}

func (t *TL_paymentRequestedInfo) Set_shipping_address(_shipping_address []byte) {
	t._shipping_address = _shipping_address
}

func (t *TL_paymentRequestedInfo) Get_shipping_address() []byte {
	return t._shipping_address
}

func New_TL_paymentRequestedInfo() *TL_paymentRequestedInfo {
	return &TL_paymentRequestedInfo{}
}

func (t *TL_paymentRequestedInfo) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_paymentRequestedInfo))
	ec.Bytes(t.Get_name())
	ec.Bytes(t.Get_phone())
	ec.Bytes(t.Get_email())
	ec.Bytes(t.Get_shipping_address())

	return ec.GetBuffer()
}

func (t *TL_paymentRequestedInfo) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._name = dc.Bytes(16)
	t._phone = dc.Bytes(16)
	t._email = dc.Bytes(16)
	t._shipping_address = dc.Bytes(16)

}

// paymentSavedCredentialsCard#cdc27a1f
type TL_paymentSavedCredentialsCard struct {
	_id    string
	_title string
}

func (t *TL_paymentSavedCredentialsCard) Set_id(_id string) {
	t._id = _id
}

func (t *TL_paymentSavedCredentialsCard) Get_id() string {
	return t._id
}

func (t *TL_paymentSavedCredentialsCard) Set_title(_title string) {
	t._title = _title
}

func (t *TL_paymentSavedCredentialsCard) Get_title() string {
	return t._title
}

func New_TL_paymentSavedCredentialsCard() *TL_paymentSavedCredentialsCard {
	return &TL_paymentSavedCredentialsCard{}
}

func (t *TL_paymentSavedCredentialsCard) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_paymentSavedCredentialsCard))
	ec.String(t.Get_id())
	ec.String(t.Get_title())

	return ec.GetBuffer()
}

func (t *TL_paymentSavedCredentialsCard) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.String()
	t._title = dc.String()

}

// webDocument#c61acbd8
type TL_webDocument struct {
	_url         string
	_access_hash int64
	_size        int32
	_mime_type   string
	_attributes  []byte
	_dc_id       int32
}

func (t *TL_webDocument) Set_url(_url string) {
	t._url = _url
}

func (t *TL_webDocument) Get_url() string {
	return t._url
}

func (t *TL_webDocument) Set_access_hash(_access_hash int64) {
	t._access_hash = _access_hash
}

func (t *TL_webDocument) Get_access_hash() int64 {
	return t._access_hash
}

func (t *TL_webDocument) Set_size(_size int32) {
	t._size = _size
}

func (t *TL_webDocument) Get_size() int32 {
	return t._size
}

func (t *TL_webDocument) Set_mime_type(_mime_type string) {
	t._mime_type = _mime_type
}

func (t *TL_webDocument) Get_mime_type() string {
	return t._mime_type
}

func (t *TL_webDocument) Set_attributes(_attributes []byte) {
	t._attributes = _attributes
}

func (t *TL_webDocument) Get_attributes() []byte {
	return t._attributes
}

func (t *TL_webDocument) Set_dc_id(_dc_id int32) {
	t._dc_id = _dc_id
}

func (t *TL_webDocument) Get_dc_id() int32 {
	return t._dc_id
}

func New_TL_webDocument() *TL_webDocument {
	return &TL_webDocument{}
}

func (t *TL_webDocument) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_webDocument))
	ec.String(t.Get_url())
	ec.Long(t.Get_access_hash())
	ec.Int(t.Get_size())
	ec.String(t.Get_mime_type())
	ec.Bytes(t.Get_attributes())
	ec.Int(t.Get_dc_id())

	return ec.GetBuffer()
}

func (t *TL_webDocument) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._url = dc.String()
	t._access_hash = dc.Long()
	t._size = dc.Int()
	t._mime_type = dc.String()
	t._attributes = dc.Bytes(16)
	t._dc_id = dc.Int()

}

// inputWebDocument#9bed434d
type TL_inputWebDocument struct {
	_url        string
	_size       int32
	_mime_type  string
	_attributes []byte
}

func (t *TL_inputWebDocument) Set_url(_url string) {
	t._url = _url
}

func (t *TL_inputWebDocument) Get_url() string {
	return t._url
}

func (t *TL_inputWebDocument) Set_size(_size int32) {
	t._size = _size
}

func (t *TL_inputWebDocument) Get_size() int32 {
	return t._size
}

func (t *TL_inputWebDocument) Set_mime_type(_mime_type string) {
	t._mime_type = _mime_type
}

func (t *TL_inputWebDocument) Get_mime_type() string {
	return t._mime_type
}

func (t *TL_inputWebDocument) Set_attributes(_attributes []byte) {
	t._attributes = _attributes
}

func (t *TL_inputWebDocument) Get_attributes() []byte {
	return t._attributes
}

func New_TL_inputWebDocument() *TL_inputWebDocument {
	return &TL_inputWebDocument{}
}

func (t *TL_inputWebDocument) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputWebDocument))
	ec.String(t.Get_url())
	ec.Int(t.Get_size())
	ec.String(t.Get_mime_type())
	ec.Bytes(t.Get_attributes())

	return ec.GetBuffer()
}

func (t *TL_inputWebDocument) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._url = dc.String()
	t._size = dc.Int()
	t._mime_type = dc.String()
	t._attributes = dc.Bytes(16)

}

// inputWebFileLocation#c239d686
type TL_inputWebFileLocation struct {
	_url         string
	_access_hash int64
}

func (t *TL_inputWebFileLocation) Set_url(_url string) {
	t._url = _url
}

func (t *TL_inputWebFileLocation) Get_url() string {
	return t._url
}

func (t *TL_inputWebFileLocation) Set_access_hash(_access_hash int64) {
	t._access_hash = _access_hash
}

func (t *TL_inputWebFileLocation) Get_access_hash() int64 {
	return t._access_hash
}

func New_TL_inputWebFileLocation() *TL_inputWebFileLocation {
	return &TL_inputWebFileLocation{}
}

func (t *TL_inputWebFileLocation) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputWebFileLocation))
	ec.String(t.Get_url())
	ec.Long(t.Get_access_hash())

	return ec.GetBuffer()
}

func (t *TL_inputWebFileLocation) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._url = dc.String()
	t._access_hash = dc.Long()

}

// upload_webFile#21e753bc
type TL_upload_webFile struct {
	_size      int32
	_mime_type string
	_file_type []byte
	_mtime     int32
	_bytes     []byte
}

func (t *TL_upload_webFile) Set_size(_size int32) {
	t._size = _size
}

func (t *TL_upload_webFile) Get_size() int32 {
	return t._size
}

func (t *TL_upload_webFile) Set_mime_type(_mime_type string) {
	t._mime_type = _mime_type
}

func (t *TL_upload_webFile) Get_mime_type() string {
	return t._mime_type
}

func (t *TL_upload_webFile) Set_file_type(_file_type []byte) {
	t._file_type = _file_type
}

func (t *TL_upload_webFile) Get_file_type() []byte {
	return t._file_type
}

func (t *TL_upload_webFile) Set_mtime(_mtime int32) {
	t._mtime = _mtime
}

func (t *TL_upload_webFile) Get_mtime() int32 {
	return t._mtime
}

func (t *TL_upload_webFile) Set_bytes(_bytes []byte) {
	t._bytes = _bytes
}

func (t *TL_upload_webFile) Get_bytes() []byte {
	return t._bytes
}

func New_TL_upload_webFile() *TL_upload_webFile {
	return &TL_upload_webFile{}
}

func (t *TL_upload_webFile) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_upload_webFile))
	ec.Int(t.Get_size())
	ec.String(t.Get_mime_type())
	ec.Bytes(t.Get_file_type())
	ec.Int(t.Get_mtime())
	ec.Bytes(t.Get_bytes())

	return ec.GetBuffer()
}

func (t *TL_upload_webFile) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._size = dc.Int()
	t._mime_type = dc.String()
	t._file_type = dc.Bytes(16)
	t._mtime = dc.Int()
	t._bytes = dc.Bytes(16)

}

// payments_paymentForm#3f56aea3
type TL_payments_paymentForm struct {
	_flags                []byte
	_can_save_credentials []byte
	_password_missing     []byte
	_bot_id               int32
	_invoice              []byte
	_provider_id          int32
	_url                  string
	_native_provider      []byte
	_native_params        []byte
	_saved_info           []byte
	_saved_credentials    []byte
	_users                []byte
}

func (t *TL_payments_paymentForm) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_payments_paymentForm) Get_flags() []byte {
	return t._flags
}

func (t *TL_payments_paymentForm) Set_can_save_credentials(_can_save_credentials []byte) {
	t._can_save_credentials = _can_save_credentials
}

func (t *TL_payments_paymentForm) Get_can_save_credentials() []byte {
	return t._can_save_credentials
}

func (t *TL_payments_paymentForm) Set_password_missing(_password_missing []byte) {
	t._password_missing = _password_missing
}

func (t *TL_payments_paymentForm) Get_password_missing() []byte {
	return t._password_missing
}

func (t *TL_payments_paymentForm) Set_bot_id(_bot_id int32) {
	t._bot_id = _bot_id
}

func (t *TL_payments_paymentForm) Get_bot_id() int32 {
	return t._bot_id
}

func (t *TL_payments_paymentForm) Set_invoice(_invoice []byte) {
	t._invoice = _invoice
}

func (t *TL_payments_paymentForm) Get_invoice() []byte {
	return t._invoice
}

func (t *TL_payments_paymentForm) Set_provider_id(_provider_id int32) {
	t._provider_id = _provider_id
}

func (t *TL_payments_paymentForm) Get_provider_id() int32 {
	return t._provider_id
}

func (t *TL_payments_paymentForm) Set_url(_url string) {
	t._url = _url
}

func (t *TL_payments_paymentForm) Get_url() string {
	return t._url
}

func (t *TL_payments_paymentForm) Set_native_provider(_native_provider []byte) {
	t._native_provider = _native_provider
}

func (t *TL_payments_paymentForm) Get_native_provider() []byte {
	return t._native_provider
}

func (t *TL_payments_paymentForm) Set_native_params(_native_params []byte) {
	t._native_params = _native_params
}

func (t *TL_payments_paymentForm) Get_native_params() []byte {
	return t._native_params
}

func (t *TL_payments_paymentForm) Set_saved_info(_saved_info []byte) {
	t._saved_info = _saved_info
}

func (t *TL_payments_paymentForm) Get_saved_info() []byte {
	return t._saved_info
}

func (t *TL_payments_paymentForm) Set_saved_credentials(_saved_credentials []byte) {
	t._saved_credentials = _saved_credentials
}

func (t *TL_payments_paymentForm) Get_saved_credentials() []byte {
	return t._saved_credentials
}

func (t *TL_payments_paymentForm) Set_users(_users []byte) {
	t._users = _users
}

func (t *TL_payments_paymentForm) Get_users() []byte {
	return t._users
}

func New_TL_payments_paymentForm() *TL_payments_paymentForm {
	return &TL_payments_paymentForm{}
}

func (t *TL_payments_paymentForm) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_payments_paymentForm))
	ec.Bytes(t.Get_can_save_credentials())
	ec.Bytes(t.Get_password_missing())
	ec.Int(t.Get_bot_id())
	ec.Bytes(t.Get_invoice())
	ec.Int(t.Get_provider_id())
	ec.String(t.Get_url())
	ec.Bytes(t.Get_native_provider())
	ec.Bytes(t.Get_native_params())
	ec.Bytes(t.Get_saved_info())
	ec.Bytes(t.Get_saved_credentials())
	ec.Bytes(t.Get_users())

	return ec.GetBuffer()
}

func (t *TL_payments_paymentForm) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._can_save_credentials = dc.Bytes(16)
	t._password_missing = dc.Bytes(16)
	t._bot_id = dc.Int()
	t._invoice = dc.Bytes(16)
	t._provider_id = dc.Int()
	t._url = dc.String()
	t._native_provider = dc.Bytes(16)
	t._native_params = dc.Bytes(16)
	t._saved_info = dc.Bytes(16)
	t._saved_credentials = dc.Bytes(16)
	t._users = dc.Bytes(16)

}

// payments_validatedRequestedInfo#d1451883
type TL_payments_validatedRequestedInfo struct {
	_flags            []byte
	_id               []byte
	_shipping_options []byte
}

func (t *TL_payments_validatedRequestedInfo) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_payments_validatedRequestedInfo) Get_flags() []byte {
	return t._flags
}

func (t *TL_payments_validatedRequestedInfo) Set_id(_id []byte) {
	t._id = _id
}

func (t *TL_payments_validatedRequestedInfo) Get_id() []byte {
	return t._id
}

func (t *TL_payments_validatedRequestedInfo) Set_shipping_options(_shipping_options []byte) {
	t._shipping_options = _shipping_options
}

func (t *TL_payments_validatedRequestedInfo) Get_shipping_options() []byte {
	return t._shipping_options
}

func New_TL_payments_validatedRequestedInfo() *TL_payments_validatedRequestedInfo {
	return &TL_payments_validatedRequestedInfo{}
}

func (t *TL_payments_validatedRequestedInfo) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_payments_validatedRequestedInfo))
	ec.Bytes(t.Get_id())
	ec.Bytes(t.Get_shipping_options())

	return ec.GetBuffer()
}

func (t *TL_payments_validatedRequestedInfo) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Bytes(16)
	t._shipping_options = dc.Bytes(16)

}

// payments_paymentResult#4e5f810d
type TL_payments_paymentResult struct {
	_updates []byte
}

func (t *TL_payments_paymentResult) Set_updates(_updates []byte) {
	t._updates = _updates
}

func (t *TL_payments_paymentResult) Get_updates() []byte {
	return t._updates
}

func New_TL_payments_paymentResult() *TL_payments_paymentResult {
	return &TL_payments_paymentResult{}
}

func (t *TL_payments_paymentResult) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_payments_paymentResult))
	ec.Bytes(t.Get_updates())

	return ec.GetBuffer()
}

func (t *TL_payments_paymentResult) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._updates = dc.Bytes(16)

}

// payments_paymentVerficationNeeded#6b56b921
type TL_payments_paymentVerficationNeeded struct {
	_url string
}

func (t *TL_payments_paymentVerficationNeeded) Set_url(_url string) {
	t._url = _url
}

func (t *TL_payments_paymentVerficationNeeded) Get_url() string {
	return t._url
}

func New_TL_payments_paymentVerficationNeeded() *TL_payments_paymentVerficationNeeded {
	return &TL_payments_paymentVerficationNeeded{}
}

func (t *TL_payments_paymentVerficationNeeded) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_payments_paymentVerficationNeeded))
	ec.String(t.Get_url())

	return ec.GetBuffer()
}

func (t *TL_payments_paymentVerficationNeeded) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._url = dc.String()

}

// payments_paymentReceipt#500911e1
type TL_payments_paymentReceipt struct {
	_flags             []byte
	_date              int32
	_bot_id            int32
	_invoice           []byte
	_provider_id       int32
	_info              []byte
	_shipping          []byte
	_currency          string
	_total_amount      int64
	_credentials_title string
	_users             []byte
}

func (t *TL_payments_paymentReceipt) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_payments_paymentReceipt) Get_flags() []byte {
	return t._flags
}

func (t *TL_payments_paymentReceipt) Set_date(_date int32) {
	t._date = _date
}

func (t *TL_payments_paymentReceipt) Get_date() int32 {
	return t._date
}

func (t *TL_payments_paymentReceipt) Set_bot_id(_bot_id int32) {
	t._bot_id = _bot_id
}

func (t *TL_payments_paymentReceipt) Get_bot_id() int32 {
	return t._bot_id
}

func (t *TL_payments_paymentReceipt) Set_invoice(_invoice []byte) {
	t._invoice = _invoice
}

func (t *TL_payments_paymentReceipt) Get_invoice() []byte {
	return t._invoice
}

func (t *TL_payments_paymentReceipt) Set_provider_id(_provider_id int32) {
	t._provider_id = _provider_id
}

func (t *TL_payments_paymentReceipt) Get_provider_id() int32 {
	return t._provider_id
}

func (t *TL_payments_paymentReceipt) Set_info(_info []byte) {
	t._info = _info
}

func (t *TL_payments_paymentReceipt) Get_info() []byte {
	return t._info
}

func (t *TL_payments_paymentReceipt) Set_shipping(_shipping []byte) {
	t._shipping = _shipping
}

func (t *TL_payments_paymentReceipt) Get_shipping() []byte {
	return t._shipping
}

func (t *TL_payments_paymentReceipt) Set_currency(_currency string) {
	t._currency = _currency
}

func (t *TL_payments_paymentReceipt) Get_currency() string {
	return t._currency
}

func (t *TL_payments_paymentReceipt) Set_total_amount(_total_amount int64) {
	t._total_amount = _total_amount
}

func (t *TL_payments_paymentReceipt) Get_total_amount() int64 {
	return t._total_amount
}

func (t *TL_payments_paymentReceipt) Set_credentials_title(_credentials_title string) {
	t._credentials_title = _credentials_title
}

func (t *TL_payments_paymentReceipt) Get_credentials_title() string {
	return t._credentials_title
}

func (t *TL_payments_paymentReceipt) Set_users(_users []byte) {
	t._users = _users
}

func (t *TL_payments_paymentReceipt) Get_users() []byte {
	return t._users
}

func New_TL_payments_paymentReceipt() *TL_payments_paymentReceipt {
	return &TL_payments_paymentReceipt{}
}

func (t *TL_payments_paymentReceipt) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_payments_paymentReceipt))
	ec.Int(t.Get_date())
	ec.Int(t.Get_bot_id())
	ec.Bytes(t.Get_invoice())
	ec.Int(t.Get_provider_id())
	ec.Bytes(t.Get_info())
	ec.Bytes(t.Get_shipping())
	ec.String(t.Get_currency())
	ec.Long(t.Get_total_amount())
	ec.String(t.Get_credentials_title())
	ec.Bytes(t.Get_users())

	return ec.GetBuffer()
}

func (t *TL_payments_paymentReceipt) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._date = dc.Int()
	t._bot_id = dc.Int()
	t._invoice = dc.Bytes(16)
	t._provider_id = dc.Int()
	t._info = dc.Bytes(16)
	t._shipping = dc.Bytes(16)
	t._currency = dc.String()
	t._total_amount = dc.Long()
	t._credentials_title = dc.String()
	t._users = dc.Bytes(16)

}

// payments_savedInfo#fb8fe43c
type TL_payments_savedInfo struct {
	_flags                 []byte
	_has_saved_credentials []byte
	_saved_info            []byte
}

func (t *TL_payments_savedInfo) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_payments_savedInfo) Get_flags() []byte {
	return t._flags
}

func (t *TL_payments_savedInfo) Set_has_saved_credentials(_has_saved_credentials []byte) {
	t._has_saved_credentials = _has_saved_credentials
}

func (t *TL_payments_savedInfo) Get_has_saved_credentials() []byte {
	return t._has_saved_credentials
}

func (t *TL_payments_savedInfo) Set_saved_info(_saved_info []byte) {
	t._saved_info = _saved_info
}

func (t *TL_payments_savedInfo) Get_saved_info() []byte {
	return t._saved_info
}

func New_TL_payments_savedInfo() *TL_payments_savedInfo {
	return &TL_payments_savedInfo{}
}

func (t *TL_payments_savedInfo) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_payments_savedInfo))
	ec.Bytes(t.Get_has_saved_credentials())
	ec.Bytes(t.Get_saved_info())

	return ec.GetBuffer()
}

func (t *TL_payments_savedInfo) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._has_saved_credentials = dc.Bytes(16)
	t._saved_info = dc.Bytes(16)

}

// inputPaymentCredentialsSaved#c10eb2cf
type TL_inputPaymentCredentialsSaved struct {
	_id           string
	_tmp_password []byte
}

func (t *TL_inputPaymentCredentialsSaved) Set_id(_id string) {
	t._id = _id
}

func (t *TL_inputPaymentCredentialsSaved) Get_id() string {
	return t._id
}

func (t *TL_inputPaymentCredentialsSaved) Set_tmp_password(_tmp_password []byte) {
	t._tmp_password = _tmp_password
}

func (t *TL_inputPaymentCredentialsSaved) Get_tmp_password() []byte {
	return t._tmp_password
}

func New_TL_inputPaymentCredentialsSaved() *TL_inputPaymentCredentialsSaved {
	return &TL_inputPaymentCredentialsSaved{}
}

func (t *TL_inputPaymentCredentialsSaved) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputPaymentCredentialsSaved))
	ec.String(t.Get_id())
	ec.Bytes(t.Get_tmp_password())

	return ec.GetBuffer()
}

func (t *TL_inputPaymentCredentialsSaved) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.String()
	t._tmp_password = dc.Bytes(16)

}

// inputPaymentCredentials#3417d728
type TL_inputPaymentCredentials struct {
	_flags []byte
	_save  []byte
	_data  []byte
}

func (t *TL_inputPaymentCredentials) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_inputPaymentCredentials) Get_flags() []byte {
	return t._flags
}

func (t *TL_inputPaymentCredentials) Set_save(_save []byte) {
	t._save = _save
}

func (t *TL_inputPaymentCredentials) Get_save() []byte {
	return t._save
}

func (t *TL_inputPaymentCredentials) Set_data(_data []byte) {
	t._data = _data
}

func (t *TL_inputPaymentCredentials) Get_data() []byte {
	return t._data
}

func New_TL_inputPaymentCredentials() *TL_inputPaymentCredentials {
	return &TL_inputPaymentCredentials{}
}

func (t *TL_inputPaymentCredentials) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputPaymentCredentials))
	ec.Bytes(t.Get_save())
	ec.Bytes(t.Get_data())

	return ec.GetBuffer()
}

func (t *TL_inputPaymentCredentials) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._save = dc.Bytes(16)
	t._data = dc.Bytes(16)

}

// inputPaymentCredentialsApplePay#0aa1c39f
type TL_inputPaymentCredentialsApplePay struct {
	_payment_data []byte
}

func (t *TL_inputPaymentCredentialsApplePay) Set_payment_data(_payment_data []byte) {
	t._payment_data = _payment_data
}

func (t *TL_inputPaymentCredentialsApplePay) Get_payment_data() []byte {
	return t._payment_data
}

func New_TL_inputPaymentCredentialsApplePay() *TL_inputPaymentCredentialsApplePay {
	return &TL_inputPaymentCredentialsApplePay{}
}

func (t *TL_inputPaymentCredentialsApplePay) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputPaymentCredentialsApplePay))
	ec.Bytes(t.Get_payment_data())

	return ec.GetBuffer()
}

func (t *TL_inputPaymentCredentialsApplePay) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._payment_data = dc.Bytes(16)

}

// inputPaymentCredentialsAndroidPay#795667a6
type TL_inputPaymentCredentialsAndroidPay struct {
	_payment_token []byte
}

func (t *TL_inputPaymentCredentialsAndroidPay) Set_payment_token(_payment_token []byte) {
	t._payment_token = _payment_token
}

func (t *TL_inputPaymentCredentialsAndroidPay) Get_payment_token() []byte {
	return t._payment_token
}

func New_TL_inputPaymentCredentialsAndroidPay() *TL_inputPaymentCredentialsAndroidPay {
	return &TL_inputPaymentCredentialsAndroidPay{}
}

func (t *TL_inputPaymentCredentialsAndroidPay) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputPaymentCredentialsAndroidPay))
	ec.Bytes(t.Get_payment_token())

	return ec.GetBuffer()
}

func (t *TL_inputPaymentCredentialsAndroidPay) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._payment_token = dc.Bytes(16)

}

// account_tmpPassword#db64fd34
type TL_account_tmpPassword struct {
	_tmp_password []byte
	_valid_until  int32
}

func (t *TL_account_tmpPassword) Set_tmp_password(_tmp_password []byte) {
	t._tmp_password = _tmp_password
}

func (t *TL_account_tmpPassword) Get_tmp_password() []byte {
	return t._tmp_password
}

func (t *TL_account_tmpPassword) Set_valid_until(_valid_until int32) {
	t._valid_until = _valid_until
}

func (t *TL_account_tmpPassword) Get_valid_until() int32 {
	return t._valid_until
}

func New_TL_account_tmpPassword() *TL_account_tmpPassword {
	return &TL_account_tmpPassword{}
}

func (t *TL_account_tmpPassword) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_account_tmpPassword))
	ec.Bytes(t.Get_tmp_password())
	ec.Int(t.Get_valid_until())

	return ec.GetBuffer()
}

func (t *TL_account_tmpPassword) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._tmp_password = dc.Bytes(16)
	t._valid_until = dc.Int()

}

// shippingOption#b6213cdf
type TL_shippingOption struct {
	_id     string
	_title  string
	_prices []byte
}

func (t *TL_shippingOption) Set_id(_id string) {
	t._id = _id
}

func (t *TL_shippingOption) Get_id() string {
	return t._id
}

func (t *TL_shippingOption) Set_title(_title string) {
	t._title = _title
}

func (t *TL_shippingOption) Get_title() string {
	return t._title
}

func (t *TL_shippingOption) Set_prices(_prices []byte) {
	t._prices = _prices
}

func (t *TL_shippingOption) Get_prices() []byte {
	return t._prices
}

func New_TL_shippingOption() *TL_shippingOption {
	return &TL_shippingOption{}
}

func (t *TL_shippingOption) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_shippingOption))
	ec.String(t.Get_id())
	ec.String(t.Get_title())
	ec.Bytes(t.Get_prices())

	return ec.GetBuffer()
}

func (t *TL_shippingOption) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.String()
	t._title = dc.String()
	t._prices = dc.Bytes(16)

}

// inputStickerSetItem#ffa0a496
type TL_inputStickerSetItem struct {
	_flags       []byte
	_document    []byte
	_emoji       string
	_mask_coords []byte
}

func (t *TL_inputStickerSetItem) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_inputStickerSetItem) Get_flags() []byte {
	return t._flags
}

func (t *TL_inputStickerSetItem) Set_document(_document []byte) {
	t._document = _document
}

func (t *TL_inputStickerSetItem) Get_document() []byte {
	return t._document
}

func (t *TL_inputStickerSetItem) Set_emoji(_emoji string) {
	t._emoji = _emoji
}

func (t *TL_inputStickerSetItem) Get_emoji() string {
	return t._emoji
}

func (t *TL_inputStickerSetItem) Set_mask_coords(_mask_coords []byte) {
	t._mask_coords = _mask_coords
}

func (t *TL_inputStickerSetItem) Get_mask_coords() []byte {
	return t._mask_coords
}

func New_TL_inputStickerSetItem() *TL_inputStickerSetItem {
	return &TL_inputStickerSetItem{}
}

func (t *TL_inputStickerSetItem) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputStickerSetItem))
	ec.Bytes(t.Get_document())
	ec.String(t.Get_emoji())
	ec.Bytes(t.Get_mask_coords())

	return ec.GetBuffer()
}

func (t *TL_inputStickerSetItem) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._document = dc.Bytes(16)
	t._emoji = dc.String()
	t._mask_coords = dc.Bytes(16)

}

// inputPhoneCall#1e36fded
type TL_inputPhoneCall struct {
	_id          int64
	_access_hash int64
}

func (t *TL_inputPhoneCall) Set_id(_id int64) {
	t._id = _id
}

func (t *TL_inputPhoneCall) Get_id() int64 {
	return t._id
}

func (t *TL_inputPhoneCall) Set_access_hash(_access_hash int64) {
	t._access_hash = _access_hash
}

func (t *TL_inputPhoneCall) Get_access_hash() int64 {
	return t._access_hash
}

func New_TL_inputPhoneCall() *TL_inputPhoneCall {
	return &TL_inputPhoneCall{}
}

func (t *TL_inputPhoneCall) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputPhoneCall))
	ec.Long(t.Get_id())
	ec.Long(t.Get_access_hash())

	return ec.GetBuffer()
}

func (t *TL_inputPhoneCall) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Long()
	t._access_hash = dc.Long()

}

// phoneCallEmpty#5366c915
type TL_phoneCallEmpty struct {
	_id int64
}

func (t *TL_phoneCallEmpty) Set_id(_id int64) {
	t._id = _id
}

func (t *TL_phoneCallEmpty) Get_id() int64 {
	return t._id
}

func New_TL_phoneCallEmpty() *TL_phoneCallEmpty {
	return &TL_phoneCallEmpty{}
}

func (t *TL_phoneCallEmpty) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_phoneCallEmpty))
	ec.Long(t.Get_id())

	return ec.GetBuffer()
}

func (t *TL_phoneCallEmpty) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Long()

}

// phoneCallWaiting#1b8f4ad1
type TL_phoneCallWaiting struct {
	_flags          []byte
	_id             int64
	_access_hash    int64
	_date           int32
	_admin_id       int32
	_participant_id int32
	_protocol       []byte
	_receive_date   []byte
}

func (t *TL_phoneCallWaiting) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_phoneCallWaiting) Get_flags() []byte {
	return t._flags
}

func (t *TL_phoneCallWaiting) Set_id(_id int64) {
	t._id = _id
}

func (t *TL_phoneCallWaiting) Get_id() int64 {
	return t._id
}

func (t *TL_phoneCallWaiting) Set_access_hash(_access_hash int64) {
	t._access_hash = _access_hash
}

func (t *TL_phoneCallWaiting) Get_access_hash() int64 {
	return t._access_hash
}

func (t *TL_phoneCallWaiting) Set_date(_date int32) {
	t._date = _date
}

func (t *TL_phoneCallWaiting) Get_date() int32 {
	return t._date
}

func (t *TL_phoneCallWaiting) Set_admin_id(_admin_id int32) {
	t._admin_id = _admin_id
}

func (t *TL_phoneCallWaiting) Get_admin_id() int32 {
	return t._admin_id
}

func (t *TL_phoneCallWaiting) Set_participant_id(_participant_id int32) {
	t._participant_id = _participant_id
}

func (t *TL_phoneCallWaiting) Get_participant_id() int32 {
	return t._participant_id
}

func (t *TL_phoneCallWaiting) Set_protocol(_protocol []byte) {
	t._protocol = _protocol
}

func (t *TL_phoneCallWaiting) Get_protocol() []byte {
	return t._protocol
}

func (t *TL_phoneCallWaiting) Set_receive_date(_receive_date []byte) {
	t._receive_date = _receive_date
}

func (t *TL_phoneCallWaiting) Get_receive_date() []byte {
	return t._receive_date
}

func New_TL_phoneCallWaiting() *TL_phoneCallWaiting {
	return &TL_phoneCallWaiting{}
}

func (t *TL_phoneCallWaiting) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_phoneCallWaiting))
	ec.Long(t.Get_id())
	ec.Long(t.Get_access_hash())
	ec.Int(t.Get_date())
	ec.Int(t.Get_admin_id())
	ec.Int(t.Get_participant_id())
	ec.Bytes(t.Get_protocol())
	ec.Bytes(t.Get_receive_date())

	return ec.GetBuffer()
}

func (t *TL_phoneCallWaiting) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Long()
	t._access_hash = dc.Long()
	t._date = dc.Int()
	t._admin_id = dc.Int()
	t._participant_id = dc.Int()
	t._protocol = dc.Bytes(16)
	t._receive_date = dc.Bytes(16)

}

// phoneCallRequested#83761ce4
type TL_phoneCallRequested struct {
	_id             int64
	_access_hash    int64
	_date           int32
	_admin_id       int32
	_participant_id int32
	_g_a_hash       []byte
	_protocol       []byte
}

func (t *TL_phoneCallRequested) Set_id(_id int64) {
	t._id = _id
}

func (t *TL_phoneCallRequested) Get_id() int64 {
	return t._id
}

func (t *TL_phoneCallRequested) Set_access_hash(_access_hash int64) {
	t._access_hash = _access_hash
}

func (t *TL_phoneCallRequested) Get_access_hash() int64 {
	return t._access_hash
}

func (t *TL_phoneCallRequested) Set_date(_date int32) {
	t._date = _date
}

func (t *TL_phoneCallRequested) Get_date() int32 {
	return t._date
}

func (t *TL_phoneCallRequested) Set_admin_id(_admin_id int32) {
	t._admin_id = _admin_id
}

func (t *TL_phoneCallRequested) Get_admin_id() int32 {
	return t._admin_id
}

func (t *TL_phoneCallRequested) Set_participant_id(_participant_id int32) {
	t._participant_id = _participant_id
}

func (t *TL_phoneCallRequested) Get_participant_id() int32 {
	return t._participant_id
}

func (t *TL_phoneCallRequested) Set_g_a_hash(_g_a_hash []byte) {
	t._g_a_hash = _g_a_hash
}

func (t *TL_phoneCallRequested) Get_g_a_hash() []byte {
	return t._g_a_hash
}

func (t *TL_phoneCallRequested) Set_protocol(_protocol []byte) {
	t._protocol = _protocol
}

func (t *TL_phoneCallRequested) Get_protocol() []byte {
	return t._protocol
}

func New_TL_phoneCallRequested() *TL_phoneCallRequested {
	return &TL_phoneCallRequested{}
}

func (t *TL_phoneCallRequested) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_phoneCallRequested))
	ec.Long(t.Get_id())
	ec.Long(t.Get_access_hash())
	ec.Int(t.Get_date())
	ec.Int(t.Get_admin_id())
	ec.Int(t.Get_participant_id())
	ec.Bytes(t.Get_g_a_hash())
	ec.Bytes(t.Get_protocol())

	return ec.GetBuffer()
}

func (t *TL_phoneCallRequested) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Long()
	t._access_hash = dc.Long()
	t._date = dc.Int()
	t._admin_id = dc.Int()
	t._participant_id = dc.Int()
	t._g_a_hash = dc.Bytes(16)
	t._protocol = dc.Bytes(16)

}

// phoneCallAccepted#6d003d3f
type TL_phoneCallAccepted struct {
	_id             int64
	_access_hash    int64
	_date           int32
	_admin_id       int32
	_participant_id int32
	_g_b            []byte
	_protocol       []byte
}

func (t *TL_phoneCallAccepted) Set_id(_id int64) {
	t._id = _id
}

func (t *TL_phoneCallAccepted) Get_id() int64 {
	return t._id
}

func (t *TL_phoneCallAccepted) Set_access_hash(_access_hash int64) {
	t._access_hash = _access_hash
}

func (t *TL_phoneCallAccepted) Get_access_hash() int64 {
	return t._access_hash
}

func (t *TL_phoneCallAccepted) Set_date(_date int32) {
	t._date = _date
}

func (t *TL_phoneCallAccepted) Get_date() int32 {
	return t._date
}

func (t *TL_phoneCallAccepted) Set_admin_id(_admin_id int32) {
	t._admin_id = _admin_id
}

func (t *TL_phoneCallAccepted) Get_admin_id() int32 {
	return t._admin_id
}

func (t *TL_phoneCallAccepted) Set_participant_id(_participant_id int32) {
	t._participant_id = _participant_id
}

func (t *TL_phoneCallAccepted) Get_participant_id() int32 {
	return t._participant_id
}

func (t *TL_phoneCallAccepted) Set_g_b(_g_b []byte) {
	t._g_b = _g_b
}

func (t *TL_phoneCallAccepted) Get_g_b() []byte {
	return t._g_b
}

func (t *TL_phoneCallAccepted) Set_protocol(_protocol []byte) {
	t._protocol = _protocol
}

func (t *TL_phoneCallAccepted) Get_protocol() []byte {
	return t._protocol
}

func New_TL_phoneCallAccepted() *TL_phoneCallAccepted {
	return &TL_phoneCallAccepted{}
}

func (t *TL_phoneCallAccepted) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_phoneCallAccepted))
	ec.Long(t.Get_id())
	ec.Long(t.Get_access_hash())
	ec.Int(t.Get_date())
	ec.Int(t.Get_admin_id())
	ec.Int(t.Get_participant_id())
	ec.Bytes(t.Get_g_b())
	ec.Bytes(t.Get_protocol())

	return ec.GetBuffer()
}

func (t *TL_phoneCallAccepted) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Long()
	t._access_hash = dc.Long()
	t._date = dc.Int()
	t._admin_id = dc.Int()
	t._participant_id = dc.Int()
	t._g_b = dc.Bytes(16)
	t._protocol = dc.Bytes(16)

}

// phoneCall#ffe6ab67
type TL_phoneCall struct {
	_id                      int64
	_access_hash             int64
	_date                    int32
	_admin_id                int32
	_participant_id          int32
	_g_a_or_b                []byte
	_key_fingerprint         int64
	_protocol                []byte
	_connection              []byte
	_alternative_connections []byte
	_start_date              int32
}

func (t *TL_phoneCall) Set_id(_id int64) {
	t._id = _id
}

func (t *TL_phoneCall) Get_id() int64 {
	return t._id
}

func (t *TL_phoneCall) Set_access_hash(_access_hash int64) {
	t._access_hash = _access_hash
}

func (t *TL_phoneCall) Get_access_hash() int64 {
	return t._access_hash
}

func (t *TL_phoneCall) Set_date(_date int32) {
	t._date = _date
}

func (t *TL_phoneCall) Get_date() int32 {
	return t._date
}

func (t *TL_phoneCall) Set_admin_id(_admin_id int32) {
	t._admin_id = _admin_id
}

func (t *TL_phoneCall) Get_admin_id() int32 {
	return t._admin_id
}

func (t *TL_phoneCall) Set_participant_id(_participant_id int32) {
	t._participant_id = _participant_id
}

func (t *TL_phoneCall) Get_participant_id() int32 {
	return t._participant_id
}

func (t *TL_phoneCall) Set_g_a_or_b(_g_a_or_b []byte) {
	t._g_a_or_b = _g_a_or_b
}

func (t *TL_phoneCall) Get_g_a_or_b() []byte {
	return t._g_a_or_b
}

func (t *TL_phoneCall) Set_key_fingerprint(_key_fingerprint int64) {
	t._key_fingerprint = _key_fingerprint
}

func (t *TL_phoneCall) Get_key_fingerprint() int64 {
	return t._key_fingerprint
}

func (t *TL_phoneCall) Set_protocol(_protocol []byte) {
	t._protocol = _protocol
}

func (t *TL_phoneCall) Get_protocol() []byte {
	return t._protocol
}

func (t *TL_phoneCall) Set_connection(_connection []byte) {
	t._connection = _connection
}

func (t *TL_phoneCall) Get_connection() []byte {
	return t._connection
}

func (t *TL_phoneCall) Set_alternative_connections(_alternative_connections []byte) {
	t._alternative_connections = _alternative_connections
}

func (t *TL_phoneCall) Get_alternative_connections() []byte {
	return t._alternative_connections
}

func (t *TL_phoneCall) Set_start_date(_start_date int32) {
	t._start_date = _start_date
}

func (t *TL_phoneCall) Get_start_date() int32 {
	return t._start_date
}

func New_TL_phoneCall() *TL_phoneCall {
	return &TL_phoneCall{}
}

func (t *TL_phoneCall) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_phoneCall))
	ec.Long(t.Get_id())
	ec.Long(t.Get_access_hash())
	ec.Int(t.Get_date())
	ec.Int(t.Get_admin_id())
	ec.Int(t.Get_participant_id())
	ec.Bytes(t.Get_g_a_or_b())
	ec.Long(t.Get_key_fingerprint())
	ec.Bytes(t.Get_protocol())
	ec.Bytes(t.Get_connection())
	ec.Bytes(t.Get_alternative_connections())
	ec.Int(t.Get_start_date())

	return ec.GetBuffer()
}

func (t *TL_phoneCall) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Long()
	t._access_hash = dc.Long()
	t._date = dc.Int()
	t._admin_id = dc.Int()
	t._participant_id = dc.Int()
	t._g_a_or_b = dc.Bytes(16)
	t._key_fingerprint = dc.Long()
	t._protocol = dc.Bytes(16)
	t._connection = dc.Bytes(16)
	t._alternative_connections = dc.Bytes(16)
	t._start_date = dc.Int()

}

// phoneCallDiscarded#50ca4de1
type TL_phoneCallDiscarded struct {
	_flags       []byte
	_need_rating []byte
	_need_debug  []byte
	_id          int64
	_reason      []byte
	_duration    []byte
}

func (t *TL_phoneCallDiscarded) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_phoneCallDiscarded) Get_flags() []byte {
	return t._flags
}

func (t *TL_phoneCallDiscarded) Set_need_rating(_need_rating []byte) {
	t._need_rating = _need_rating
}

func (t *TL_phoneCallDiscarded) Get_need_rating() []byte {
	return t._need_rating
}

func (t *TL_phoneCallDiscarded) Set_need_debug(_need_debug []byte) {
	t._need_debug = _need_debug
}

func (t *TL_phoneCallDiscarded) Get_need_debug() []byte {
	return t._need_debug
}

func (t *TL_phoneCallDiscarded) Set_id(_id int64) {
	t._id = _id
}

func (t *TL_phoneCallDiscarded) Get_id() int64 {
	return t._id
}

func (t *TL_phoneCallDiscarded) Set_reason(_reason []byte) {
	t._reason = _reason
}

func (t *TL_phoneCallDiscarded) Get_reason() []byte {
	return t._reason
}

func (t *TL_phoneCallDiscarded) Set_duration(_duration []byte) {
	t._duration = _duration
}

func (t *TL_phoneCallDiscarded) Get_duration() []byte {
	return t._duration
}

func New_TL_phoneCallDiscarded() *TL_phoneCallDiscarded {
	return &TL_phoneCallDiscarded{}
}

func (t *TL_phoneCallDiscarded) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_phoneCallDiscarded))
	ec.Bytes(t.Get_need_rating())
	ec.Bytes(t.Get_need_debug())
	ec.Long(t.Get_id())
	ec.Bytes(t.Get_reason())
	ec.Bytes(t.Get_duration())

	return ec.GetBuffer()
}

func (t *TL_phoneCallDiscarded) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._need_rating = dc.Bytes(16)
	t._need_debug = dc.Bytes(16)
	t._id = dc.Long()
	t._reason = dc.Bytes(16)
	t._duration = dc.Bytes(16)

}

// phoneConnection#9d4c17c0
type TL_phoneConnection struct {
	_id       int64
	_ip       string
	_ipv6     string
	_port     int32
	_peer_tag []byte
}

func (t *TL_phoneConnection) Set_id(_id int64) {
	t._id = _id
}

func (t *TL_phoneConnection) Get_id() int64 {
	return t._id
}

func (t *TL_phoneConnection) Set_ip(_ip string) {
	t._ip = _ip
}

func (t *TL_phoneConnection) Get_ip() string {
	return t._ip
}

func (t *TL_phoneConnection) Set_ipv6(_ipv6 string) {
	t._ipv6 = _ipv6
}

func (t *TL_phoneConnection) Get_ipv6() string {
	return t._ipv6
}

func (t *TL_phoneConnection) Set_port(_port int32) {
	t._port = _port
}

func (t *TL_phoneConnection) Get_port() int32 {
	return t._port
}

func (t *TL_phoneConnection) Set_peer_tag(_peer_tag []byte) {
	t._peer_tag = _peer_tag
}

func (t *TL_phoneConnection) Get_peer_tag() []byte {
	return t._peer_tag
}

func New_TL_phoneConnection() *TL_phoneConnection {
	return &TL_phoneConnection{}
}

func (t *TL_phoneConnection) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_phoneConnection))
	ec.Long(t.Get_id())
	ec.String(t.Get_ip())
	ec.String(t.Get_ipv6())
	ec.Int(t.Get_port())
	ec.Bytes(t.Get_peer_tag())

	return ec.GetBuffer()
}

func (t *TL_phoneConnection) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Long()
	t._ip = dc.String()
	t._ipv6 = dc.String()
	t._port = dc.Int()
	t._peer_tag = dc.Bytes(16)

}

// phoneCallProtocol#a2bb35cb
type TL_phoneCallProtocol struct {
	_flags         []byte
	_udp_p2p       []byte
	_udp_reflector []byte
	_min_layer     int32
	_max_layer     int32
}

func (t *TL_phoneCallProtocol) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_phoneCallProtocol) Get_flags() []byte {
	return t._flags
}

func (t *TL_phoneCallProtocol) Set_udp_p2p(_udp_p2p []byte) {
	t._udp_p2p = _udp_p2p
}

func (t *TL_phoneCallProtocol) Get_udp_p2p() []byte {
	return t._udp_p2p
}

func (t *TL_phoneCallProtocol) Set_udp_reflector(_udp_reflector []byte) {
	t._udp_reflector = _udp_reflector
}

func (t *TL_phoneCallProtocol) Get_udp_reflector() []byte {
	return t._udp_reflector
}

func (t *TL_phoneCallProtocol) Set_min_layer(_min_layer int32) {
	t._min_layer = _min_layer
}

func (t *TL_phoneCallProtocol) Get_min_layer() int32 {
	return t._min_layer
}

func (t *TL_phoneCallProtocol) Set_max_layer(_max_layer int32) {
	t._max_layer = _max_layer
}

func (t *TL_phoneCallProtocol) Get_max_layer() int32 {
	return t._max_layer
}

func New_TL_phoneCallProtocol() *TL_phoneCallProtocol {
	return &TL_phoneCallProtocol{}
}

func (t *TL_phoneCallProtocol) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_phoneCallProtocol))
	ec.Bytes(t.Get_udp_p2p())
	ec.Bytes(t.Get_udp_reflector())
	ec.Int(t.Get_min_layer())
	ec.Int(t.Get_max_layer())

	return ec.GetBuffer()
}

func (t *TL_phoneCallProtocol) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._udp_p2p = dc.Bytes(16)
	t._udp_reflector = dc.Bytes(16)
	t._min_layer = dc.Int()
	t._max_layer = dc.Int()

}

// phone_phoneCall#ec82e140
type TL_phone_phoneCall struct {
	_phone_call []byte
	_users      []byte
}

func (t *TL_phone_phoneCall) Set_phone_call(_phone_call []byte) {
	t._phone_call = _phone_call
}

func (t *TL_phone_phoneCall) Get_phone_call() []byte {
	return t._phone_call
}

func (t *TL_phone_phoneCall) Set_users(_users []byte) {
	t._users = _users
}

func (t *TL_phone_phoneCall) Get_users() []byte {
	return t._users
}

func New_TL_phone_phoneCall() *TL_phone_phoneCall {
	return &TL_phone_phoneCall{}
}

func (t *TL_phone_phoneCall) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_phone_phoneCall))
	ec.Bytes(t.Get_phone_call())
	ec.Bytes(t.Get_users())

	return ec.GetBuffer()
}

func (t *TL_phone_phoneCall) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._phone_call = dc.Bytes(16)
	t._users = dc.Bytes(16)

}

// upload_cdnFileReuploadNeeded#eea8e46e
type TL_upload_cdnFileReuploadNeeded struct {
	_request_token []byte
}

func (t *TL_upload_cdnFileReuploadNeeded) Set_request_token(_request_token []byte) {
	t._request_token = _request_token
}

func (t *TL_upload_cdnFileReuploadNeeded) Get_request_token() []byte {
	return t._request_token
}

func New_TL_upload_cdnFileReuploadNeeded() *TL_upload_cdnFileReuploadNeeded {
	return &TL_upload_cdnFileReuploadNeeded{}
}

func (t *TL_upload_cdnFileReuploadNeeded) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_upload_cdnFileReuploadNeeded))
	ec.Bytes(t.Get_request_token())

	return ec.GetBuffer()
}

func (t *TL_upload_cdnFileReuploadNeeded) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._request_token = dc.Bytes(16)

}

// upload_cdnFile#a99fca4f
type TL_upload_cdnFile struct {
	_bytes []byte
}

func (t *TL_upload_cdnFile) Set_bytes(_bytes []byte) {
	t._bytes = _bytes
}

func (t *TL_upload_cdnFile) Get_bytes() []byte {
	return t._bytes
}

func New_TL_upload_cdnFile() *TL_upload_cdnFile {
	return &TL_upload_cdnFile{}
}

func (t *TL_upload_cdnFile) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_upload_cdnFile))
	ec.Bytes(t.Get_bytes())

	return ec.GetBuffer()
}

func (t *TL_upload_cdnFile) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._bytes = dc.Bytes(16)

}

// cdnPublicKey#c982eaba
type TL_cdnPublicKey struct {
	_dc_id      int32
	_public_key string
}

func (t *TL_cdnPublicKey) Set_dc_id(_dc_id int32) {
	t._dc_id = _dc_id
}

func (t *TL_cdnPublicKey) Get_dc_id() int32 {
	return t._dc_id
}

func (t *TL_cdnPublicKey) Set_public_key(_public_key string) {
	t._public_key = _public_key
}

func (t *TL_cdnPublicKey) Get_public_key() string {
	return t._public_key
}

func New_TL_cdnPublicKey() *TL_cdnPublicKey {
	return &TL_cdnPublicKey{}
}

func (t *TL_cdnPublicKey) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_cdnPublicKey))
	ec.Int(t.Get_dc_id())
	ec.String(t.Get_public_key())

	return ec.GetBuffer()
}

func (t *TL_cdnPublicKey) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._dc_id = dc.Int()
	t._public_key = dc.String()

}

// cdnConfig#5725e40a
type TL_cdnConfig struct {
	_public_keys []byte
}

func (t *TL_cdnConfig) Set_public_keys(_public_keys []byte) {
	t._public_keys = _public_keys
}

func (t *TL_cdnConfig) Get_public_keys() []byte {
	return t._public_keys
}

func New_TL_cdnConfig() *TL_cdnConfig {
	return &TL_cdnConfig{}
}

func (t *TL_cdnConfig) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_cdnConfig))
	ec.Bytes(t.Get_public_keys())

	return ec.GetBuffer()
}

func (t *TL_cdnConfig) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._public_keys = dc.Bytes(16)

}

// langPackString#cad181f6
type TL_langPackString struct {
	_key   string
	_value string
}

func (t *TL_langPackString) Set_key(_key string) {
	t._key = _key
}

func (t *TL_langPackString) Get_key() string {
	return t._key
}

func (t *TL_langPackString) Set_value(_value string) {
	t._value = _value
}

func (t *TL_langPackString) Get_value() string {
	return t._value
}

func New_TL_langPackString() *TL_langPackString {
	return &TL_langPackString{}
}

func (t *TL_langPackString) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_langPackString))
	ec.String(t.Get_key())
	ec.String(t.Get_value())

	return ec.GetBuffer()
}

func (t *TL_langPackString) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._key = dc.String()
	t._value = dc.String()

}

// langPackStringPluralized#6c47ac9f
type TL_langPackStringPluralized struct {
	_flags       []byte
	_key         string
	_zero_value  []byte
	_one_value   []byte
	_two_value   []byte
	_few_value   []byte
	_many_value  []byte
	_other_value string
}

func (t *TL_langPackStringPluralized) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_langPackStringPluralized) Get_flags() []byte {
	return t._flags
}

func (t *TL_langPackStringPluralized) Set_key(_key string) {
	t._key = _key
}

func (t *TL_langPackStringPluralized) Get_key() string {
	return t._key
}

func (t *TL_langPackStringPluralized) Set_zero_value(_zero_value []byte) {
	t._zero_value = _zero_value
}

func (t *TL_langPackStringPluralized) Get_zero_value() []byte {
	return t._zero_value
}

func (t *TL_langPackStringPluralized) Set_one_value(_one_value []byte) {
	t._one_value = _one_value
}

func (t *TL_langPackStringPluralized) Get_one_value() []byte {
	return t._one_value
}

func (t *TL_langPackStringPluralized) Set_two_value(_two_value []byte) {
	t._two_value = _two_value
}

func (t *TL_langPackStringPluralized) Get_two_value() []byte {
	return t._two_value
}

func (t *TL_langPackStringPluralized) Set_few_value(_few_value []byte) {
	t._few_value = _few_value
}

func (t *TL_langPackStringPluralized) Get_few_value() []byte {
	return t._few_value
}

func (t *TL_langPackStringPluralized) Set_many_value(_many_value []byte) {
	t._many_value = _many_value
}

func (t *TL_langPackStringPluralized) Get_many_value() []byte {
	return t._many_value
}

func (t *TL_langPackStringPluralized) Set_other_value(_other_value string) {
	t._other_value = _other_value
}

func (t *TL_langPackStringPluralized) Get_other_value() string {
	return t._other_value
}

func New_TL_langPackStringPluralized() *TL_langPackStringPluralized {
	return &TL_langPackStringPluralized{}
}

func (t *TL_langPackStringPluralized) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_langPackStringPluralized))
	ec.String(t.Get_key())
	ec.Bytes(t.Get_zero_value())
	ec.Bytes(t.Get_one_value())
	ec.Bytes(t.Get_two_value())
	ec.Bytes(t.Get_few_value())
	ec.Bytes(t.Get_many_value())
	ec.String(t.Get_other_value())

	return ec.GetBuffer()
}

func (t *TL_langPackStringPluralized) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._key = dc.String()
	t._zero_value = dc.Bytes(16)
	t._one_value = dc.Bytes(16)
	t._two_value = dc.Bytes(16)
	t._few_value = dc.Bytes(16)
	t._many_value = dc.Bytes(16)
	t._other_value = dc.String()

}

// langPackStringDeleted#2979eeb2
type TL_langPackStringDeleted struct {
	_key string
}

func (t *TL_langPackStringDeleted) Set_key(_key string) {
	t._key = _key
}

func (t *TL_langPackStringDeleted) Get_key() string {
	return t._key
}

func New_TL_langPackStringDeleted() *TL_langPackStringDeleted {
	return &TL_langPackStringDeleted{}
}

func (t *TL_langPackStringDeleted) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_langPackStringDeleted))
	ec.String(t.Get_key())

	return ec.GetBuffer()
}

func (t *TL_langPackStringDeleted) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._key = dc.String()

}

// langPackDifference#f385c1f6
type TL_langPackDifference struct {
	_lang_code    string
	_from_version int32
	_version      int32
	_strings      []byte
}

func (t *TL_langPackDifference) Set_lang_code(_lang_code string) {
	t._lang_code = _lang_code
}

func (t *TL_langPackDifference) Get_lang_code() string {
	return t._lang_code
}

func (t *TL_langPackDifference) Set_from_version(_from_version int32) {
	t._from_version = _from_version
}

func (t *TL_langPackDifference) Get_from_version() int32 {
	return t._from_version
}

func (t *TL_langPackDifference) Set_version(_version int32) {
	t._version = _version
}

func (t *TL_langPackDifference) Get_version() int32 {
	return t._version
}

func (t *TL_langPackDifference) Set_strings(_strings []byte) {
	t._strings = _strings
}

func (t *TL_langPackDifference) Get_strings() []byte {
	return t._strings
}

func New_TL_langPackDifference() *TL_langPackDifference {
	return &TL_langPackDifference{}
}

func (t *TL_langPackDifference) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_langPackDifference))
	ec.String(t.Get_lang_code())
	ec.Int(t.Get_from_version())
	ec.Int(t.Get_version())
	ec.Bytes(t.Get_strings())

	return ec.GetBuffer()
}

func (t *TL_langPackDifference) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._lang_code = dc.String()
	t._from_version = dc.Int()
	t._version = dc.Int()
	t._strings = dc.Bytes(16)

}

// langPackLanguage#117698f1
type TL_langPackLanguage struct {
	_name        string
	_native_name string
	_lang_code   string
}

func (t *TL_langPackLanguage) Set_name(_name string) {
	t._name = _name
}

func (t *TL_langPackLanguage) Get_name() string {
	return t._name
}

func (t *TL_langPackLanguage) Set_native_name(_native_name string) {
	t._native_name = _native_name
}

func (t *TL_langPackLanguage) Get_native_name() string {
	return t._native_name
}

func (t *TL_langPackLanguage) Set_lang_code(_lang_code string) {
	t._lang_code = _lang_code
}

func (t *TL_langPackLanguage) Get_lang_code() string {
	return t._lang_code
}

func New_TL_langPackLanguage() *TL_langPackLanguage {
	return &TL_langPackLanguage{}
}

func (t *TL_langPackLanguage) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_langPackLanguage))
	ec.String(t.Get_name())
	ec.String(t.Get_native_name())
	ec.String(t.Get_lang_code())

	return ec.GetBuffer()
}

func (t *TL_langPackLanguage) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._name = dc.String()
	t._native_name = dc.String()
	t._lang_code = dc.String()

}

// channelAdminRights#5d7ceba5
type TL_channelAdminRights struct {
	_flags           []byte
	_change_info     []byte
	_post_messages   []byte
	_edit_messages   []byte
	_delete_messages []byte
	_ban_users       []byte
	_invite_users    []byte
	_invite_link     []byte
	_pin_messages    []byte
	_add_admins      []byte
}

func (t *TL_channelAdminRights) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_channelAdminRights) Get_flags() []byte {
	return t._flags
}

func (t *TL_channelAdminRights) Set_change_info(_change_info []byte) {
	t._change_info = _change_info
}

func (t *TL_channelAdminRights) Get_change_info() []byte {
	return t._change_info
}

func (t *TL_channelAdminRights) Set_post_messages(_post_messages []byte) {
	t._post_messages = _post_messages
}

func (t *TL_channelAdminRights) Get_post_messages() []byte {
	return t._post_messages
}

func (t *TL_channelAdminRights) Set_edit_messages(_edit_messages []byte) {
	t._edit_messages = _edit_messages
}

func (t *TL_channelAdminRights) Get_edit_messages() []byte {
	return t._edit_messages
}

func (t *TL_channelAdminRights) Set_delete_messages(_delete_messages []byte) {
	t._delete_messages = _delete_messages
}

func (t *TL_channelAdminRights) Get_delete_messages() []byte {
	return t._delete_messages
}

func (t *TL_channelAdminRights) Set_ban_users(_ban_users []byte) {
	t._ban_users = _ban_users
}

func (t *TL_channelAdminRights) Get_ban_users() []byte {
	return t._ban_users
}

func (t *TL_channelAdminRights) Set_invite_users(_invite_users []byte) {
	t._invite_users = _invite_users
}

func (t *TL_channelAdminRights) Get_invite_users() []byte {
	return t._invite_users
}

func (t *TL_channelAdminRights) Set_invite_link(_invite_link []byte) {
	t._invite_link = _invite_link
}

func (t *TL_channelAdminRights) Get_invite_link() []byte {
	return t._invite_link
}

func (t *TL_channelAdminRights) Set_pin_messages(_pin_messages []byte) {
	t._pin_messages = _pin_messages
}

func (t *TL_channelAdminRights) Get_pin_messages() []byte {
	return t._pin_messages
}

func (t *TL_channelAdminRights) Set_add_admins(_add_admins []byte) {
	t._add_admins = _add_admins
}

func (t *TL_channelAdminRights) Get_add_admins() []byte {
	return t._add_admins
}

func New_TL_channelAdminRights() *TL_channelAdminRights {
	return &TL_channelAdminRights{}
}

func (t *TL_channelAdminRights) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channelAdminRights))
	ec.Bytes(t.Get_change_info())
	ec.Bytes(t.Get_post_messages())
	ec.Bytes(t.Get_edit_messages())
	ec.Bytes(t.Get_delete_messages())
	ec.Bytes(t.Get_ban_users())
	ec.Bytes(t.Get_invite_users())
	ec.Bytes(t.Get_invite_link())
	ec.Bytes(t.Get_pin_messages())
	ec.Bytes(t.Get_add_admins())

	return ec.GetBuffer()
}

func (t *TL_channelAdminRights) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._change_info = dc.Bytes(16)
	t._post_messages = dc.Bytes(16)
	t._edit_messages = dc.Bytes(16)
	t._delete_messages = dc.Bytes(16)
	t._ban_users = dc.Bytes(16)
	t._invite_users = dc.Bytes(16)
	t._invite_link = dc.Bytes(16)
	t._pin_messages = dc.Bytes(16)
	t._add_admins = dc.Bytes(16)

}

// channelBannedRights#58cf4249
type TL_channelBannedRights struct {
	_flags         []byte
	_view_messages []byte
	_send_messages []byte
	_send_media    []byte
	_send_stickers []byte
	_send_gifs     []byte
	_send_games    []byte
	_send_inline   []byte
	_embed_links   []byte
	_until_date    int32
}

func (t *TL_channelBannedRights) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_channelBannedRights) Get_flags() []byte {
	return t._flags
}

func (t *TL_channelBannedRights) Set_view_messages(_view_messages []byte) {
	t._view_messages = _view_messages
}

func (t *TL_channelBannedRights) Get_view_messages() []byte {
	return t._view_messages
}

func (t *TL_channelBannedRights) Set_send_messages(_send_messages []byte) {
	t._send_messages = _send_messages
}

func (t *TL_channelBannedRights) Get_send_messages() []byte {
	return t._send_messages
}

func (t *TL_channelBannedRights) Set_send_media(_send_media []byte) {
	t._send_media = _send_media
}

func (t *TL_channelBannedRights) Get_send_media() []byte {
	return t._send_media
}

func (t *TL_channelBannedRights) Set_send_stickers(_send_stickers []byte) {
	t._send_stickers = _send_stickers
}

func (t *TL_channelBannedRights) Get_send_stickers() []byte {
	return t._send_stickers
}

func (t *TL_channelBannedRights) Set_send_gifs(_send_gifs []byte) {
	t._send_gifs = _send_gifs
}

func (t *TL_channelBannedRights) Get_send_gifs() []byte {
	return t._send_gifs
}

func (t *TL_channelBannedRights) Set_send_games(_send_games []byte) {
	t._send_games = _send_games
}

func (t *TL_channelBannedRights) Get_send_games() []byte {
	return t._send_games
}

func (t *TL_channelBannedRights) Set_send_inline(_send_inline []byte) {
	t._send_inline = _send_inline
}

func (t *TL_channelBannedRights) Get_send_inline() []byte {
	return t._send_inline
}

func (t *TL_channelBannedRights) Set_embed_links(_embed_links []byte) {
	t._embed_links = _embed_links
}

func (t *TL_channelBannedRights) Get_embed_links() []byte {
	return t._embed_links
}

func (t *TL_channelBannedRights) Set_until_date(_until_date int32) {
	t._until_date = _until_date
}

func (t *TL_channelBannedRights) Get_until_date() int32 {
	return t._until_date
}

func New_TL_channelBannedRights() *TL_channelBannedRights {
	return &TL_channelBannedRights{}
}

func (t *TL_channelBannedRights) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channelBannedRights))
	ec.Bytes(t.Get_view_messages())
	ec.Bytes(t.Get_send_messages())
	ec.Bytes(t.Get_send_media())
	ec.Bytes(t.Get_send_stickers())
	ec.Bytes(t.Get_send_gifs())
	ec.Bytes(t.Get_send_games())
	ec.Bytes(t.Get_send_inline())
	ec.Bytes(t.Get_embed_links())
	ec.Int(t.Get_until_date())

	return ec.GetBuffer()
}

func (t *TL_channelBannedRights) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._view_messages = dc.Bytes(16)
	t._send_messages = dc.Bytes(16)
	t._send_media = dc.Bytes(16)
	t._send_stickers = dc.Bytes(16)
	t._send_gifs = dc.Bytes(16)
	t._send_games = dc.Bytes(16)
	t._send_inline = dc.Bytes(16)
	t._embed_links = dc.Bytes(16)
	t._until_date = dc.Int()

}

// channelAdminLogEventActionChangeTitle#e6dfb825
type TL_channelAdminLogEventActionChangeTitle struct {
	_prev_value string
	_new_value  string
}

func (t *TL_channelAdminLogEventActionChangeTitle) Set_prev_value(_prev_value string) {
	t._prev_value = _prev_value
}

func (t *TL_channelAdminLogEventActionChangeTitle) Get_prev_value() string {
	return t._prev_value
}

func (t *TL_channelAdminLogEventActionChangeTitle) Set_new_value(_new_value string) {
	t._new_value = _new_value
}

func (t *TL_channelAdminLogEventActionChangeTitle) Get_new_value() string {
	return t._new_value
}

func New_TL_channelAdminLogEventActionChangeTitle() *TL_channelAdminLogEventActionChangeTitle {
	return &TL_channelAdminLogEventActionChangeTitle{}
}

func (t *TL_channelAdminLogEventActionChangeTitle) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channelAdminLogEventActionChangeTitle))
	ec.String(t.Get_prev_value())
	ec.String(t.Get_new_value())

	return ec.GetBuffer()
}

func (t *TL_channelAdminLogEventActionChangeTitle) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._prev_value = dc.String()
	t._new_value = dc.String()

}

// channelAdminLogEventActionChangeAbout#55188a2e
type TL_channelAdminLogEventActionChangeAbout struct {
	_prev_value string
	_new_value  string
}

func (t *TL_channelAdminLogEventActionChangeAbout) Set_prev_value(_prev_value string) {
	t._prev_value = _prev_value
}

func (t *TL_channelAdminLogEventActionChangeAbout) Get_prev_value() string {
	return t._prev_value
}

func (t *TL_channelAdminLogEventActionChangeAbout) Set_new_value(_new_value string) {
	t._new_value = _new_value
}

func (t *TL_channelAdminLogEventActionChangeAbout) Get_new_value() string {
	return t._new_value
}

func New_TL_channelAdminLogEventActionChangeAbout() *TL_channelAdminLogEventActionChangeAbout {
	return &TL_channelAdminLogEventActionChangeAbout{}
}

func (t *TL_channelAdminLogEventActionChangeAbout) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channelAdminLogEventActionChangeAbout))
	ec.String(t.Get_prev_value())
	ec.String(t.Get_new_value())

	return ec.GetBuffer()
}

func (t *TL_channelAdminLogEventActionChangeAbout) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._prev_value = dc.String()
	t._new_value = dc.String()

}

// channelAdminLogEventActionChangeUsername#6a4afc38
type TL_channelAdminLogEventActionChangeUsername struct {
	_prev_value string
	_new_value  string
}

func (t *TL_channelAdminLogEventActionChangeUsername) Set_prev_value(_prev_value string) {
	t._prev_value = _prev_value
}

func (t *TL_channelAdminLogEventActionChangeUsername) Get_prev_value() string {
	return t._prev_value
}

func (t *TL_channelAdminLogEventActionChangeUsername) Set_new_value(_new_value string) {
	t._new_value = _new_value
}

func (t *TL_channelAdminLogEventActionChangeUsername) Get_new_value() string {
	return t._new_value
}

func New_TL_channelAdminLogEventActionChangeUsername() *TL_channelAdminLogEventActionChangeUsername {
	return &TL_channelAdminLogEventActionChangeUsername{}
}

func (t *TL_channelAdminLogEventActionChangeUsername) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channelAdminLogEventActionChangeUsername))
	ec.String(t.Get_prev_value())
	ec.String(t.Get_new_value())

	return ec.GetBuffer()
}

func (t *TL_channelAdminLogEventActionChangeUsername) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._prev_value = dc.String()
	t._new_value = dc.String()

}

// channelAdminLogEventActionChangePhoto#b82f55c3
type TL_channelAdminLogEventActionChangePhoto struct {
	_prev_photo []byte
	_new_photo  []byte
}

func (t *TL_channelAdminLogEventActionChangePhoto) Set_prev_photo(_prev_photo []byte) {
	t._prev_photo = _prev_photo
}

func (t *TL_channelAdminLogEventActionChangePhoto) Get_prev_photo() []byte {
	return t._prev_photo
}

func (t *TL_channelAdminLogEventActionChangePhoto) Set_new_photo(_new_photo []byte) {
	t._new_photo = _new_photo
}

func (t *TL_channelAdminLogEventActionChangePhoto) Get_new_photo() []byte {
	return t._new_photo
}

func New_TL_channelAdminLogEventActionChangePhoto() *TL_channelAdminLogEventActionChangePhoto {
	return &TL_channelAdminLogEventActionChangePhoto{}
}

func (t *TL_channelAdminLogEventActionChangePhoto) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channelAdminLogEventActionChangePhoto))
	ec.Bytes(t.Get_prev_photo())
	ec.Bytes(t.Get_new_photo())

	return ec.GetBuffer()
}

func (t *TL_channelAdminLogEventActionChangePhoto) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._prev_photo = dc.Bytes(16)
	t._new_photo = dc.Bytes(16)

}

// channelAdminLogEventActionToggleInvites#1b7907ae
type TL_channelAdminLogEventActionToggleInvites struct {
	_new_value bool
}

func (t *TL_channelAdminLogEventActionToggleInvites) Set_new_value(_new_value bool) {
	t._new_value = _new_value
}

func (t *TL_channelAdminLogEventActionToggleInvites) Get_new_value() bool {
	return t._new_value
}

func New_TL_channelAdminLogEventActionToggleInvites() *TL_channelAdminLogEventActionToggleInvites {
	return &TL_channelAdminLogEventActionToggleInvites{}
}

func (t *TL_channelAdminLogEventActionToggleInvites) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channelAdminLogEventActionToggleInvites))
	ec.Bool(t.Get_new_value())

	return ec.GetBuffer()
}

func (t *TL_channelAdminLogEventActionToggleInvites) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._new_value = dc.Bool()

}

// channelAdminLogEventActionToggleSignatures#26ae0971
type TL_channelAdminLogEventActionToggleSignatures struct {
	_new_value bool
}

func (t *TL_channelAdminLogEventActionToggleSignatures) Set_new_value(_new_value bool) {
	t._new_value = _new_value
}

func (t *TL_channelAdminLogEventActionToggleSignatures) Get_new_value() bool {
	return t._new_value
}

func New_TL_channelAdminLogEventActionToggleSignatures() *TL_channelAdminLogEventActionToggleSignatures {
	return &TL_channelAdminLogEventActionToggleSignatures{}
}

func (t *TL_channelAdminLogEventActionToggleSignatures) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channelAdminLogEventActionToggleSignatures))
	ec.Bool(t.Get_new_value())

	return ec.GetBuffer()
}

func (t *TL_channelAdminLogEventActionToggleSignatures) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._new_value = dc.Bool()

}

// channelAdminLogEventActionUpdatePinned#e9e82c18
type TL_channelAdminLogEventActionUpdatePinned struct {
	_message []byte
}

func (t *TL_channelAdminLogEventActionUpdatePinned) Set_message(_message []byte) {
	t._message = _message
}

func (t *TL_channelAdminLogEventActionUpdatePinned) Get_message() []byte {
	return t._message
}

func New_TL_channelAdminLogEventActionUpdatePinned() *TL_channelAdminLogEventActionUpdatePinned {
	return &TL_channelAdminLogEventActionUpdatePinned{}
}

func (t *TL_channelAdminLogEventActionUpdatePinned) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channelAdminLogEventActionUpdatePinned))
	ec.Bytes(t.Get_message())

	return ec.GetBuffer()
}

func (t *TL_channelAdminLogEventActionUpdatePinned) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._message = dc.Bytes(16)

}

// channelAdminLogEventActionEditMessage#709b2405
type TL_channelAdminLogEventActionEditMessage struct {
	_prev_message []byte
	_new_message  []byte
}

func (t *TL_channelAdminLogEventActionEditMessage) Set_prev_message(_prev_message []byte) {
	t._prev_message = _prev_message
}

func (t *TL_channelAdminLogEventActionEditMessage) Get_prev_message() []byte {
	return t._prev_message
}

func (t *TL_channelAdminLogEventActionEditMessage) Set_new_message(_new_message []byte) {
	t._new_message = _new_message
}

func (t *TL_channelAdminLogEventActionEditMessage) Get_new_message() []byte {
	return t._new_message
}

func New_TL_channelAdminLogEventActionEditMessage() *TL_channelAdminLogEventActionEditMessage {
	return &TL_channelAdminLogEventActionEditMessage{}
}

func (t *TL_channelAdminLogEventActionEditMessage) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channelAdminLogEventActionEditMessage))
	ec.Bytes(t.Get_prev_message())
	ec.Bytes(t.Get_new_message())

	return ec.GetBuffer()
}

func (t *TL_channelAdminLogEventActionEditMessage) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._prev_message = dc.Bytes(16)
	t._new_message = dc.Bytes(16)

}

// channelAdminLogEventActionDeleteMessage#42e047bb
type TL_channelAdminLogEventActionDeleteMessage struct {
	_message []byte
}

func (t *TL_channelAdminLogEventActionDeleteMessage) Set_message(_message []byte) {
	t._message = _message
}

func (t *TL_channelAdminLogEventActionDeleteMessage) Get_message() []byte {
	return t._message
}

func New_TL_channelAdminLogEventActionDeleteMessage() *TL_channelAdminLogEventActionDeleteMessage {
	return &TL_channelAdminLogEventActionDeleteMessage{}
}

func (t *TL_channelAdminLogEventActionDeleteMessage) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channelAdminLogEventActionDeleteMessage))
	ec.Bytes(t.Get_message())

	return ec.GetBuffer()
}

func (t *TL_channelAdminLogEventActionDeleteMessage) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._message = dc.Bytes(16)

}

// channelAdminLogEventActionParticipantJoin#183040d3
type TL_channelAdminLogEventActionParticipantJoin struct {
}

func New_TL_channelAdminLogEventActionParticipantJoin() *TL_channelAdminLogEventActionParticipantJoin {
	return &TL_channelAdminLogEventActionParticipantJoin{}
}

func (t *TL_channelAdminLogEventActionParticipantJoin) Encode() []byte {
	return nil
}

func (t *TL_channelAdminLogEventActionParticipantJoin) Decode(b []byte) {

}

// channelAdminLogEventActionParticipantLeave#f89777f2
type TL_channelAdminLogEventActionParticipantLeave struct {
}

func New_TL_channelAdminLogEventActionParticipantLeave() *TL_channelAdminLogEventActionParticipantLeave {
	return &TL_channelAdminLogEventActionParticipantLeave{}
}

func (t *TL_channelAdminLogEventActionParticipantLeave) Encode() []byte {
	return nil
}

func (t *TL_channelAdminLogEventActionParticipantLeave) Decode(b []byte) {

}

// channelAdminLogEventActionParticipantInvite#e31c34d8
type TL_channelAdminLogEventActionParticipantInvite struct {
	_participant []byte
}

func (t *TL_channelAdminLogEventActionParticipantInvite) Set_participant(_participant []byte) {
	t._participant = _participant
}

func (t *TL_channelAdminLogEventActionParticipantInvite) Get_participant() []byte {
	return t._participant
}

func New_TL_channelAdminLogEventActionParticipantInvite() *TL_channelAdminLogEventActionParticipantInvite {
	return &TL_channelAdminLogEventActionParticipantInvite{}
}

func (t *TL_channelAdminLogEventActionParticipantInvite) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channelAdminLogEventActionParticipantInvite))
	ec.Bytes(t.Get_participant())

	return ec.GetBuffer()
}

func (t *TL_channelAdminLogEventActionParticipantInvite) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._participant = dc.Bytes(16)

}

// channelAdminLogEventActionParticipantToggleBan#e6d83d7e
type TL_channelAdminLogEventActionParticipantToggleBan struct {
	_prev_participant []byte
	_new_participant  []byte
}

func (t *TL_channelAdminLogEventActionParticipantToggleBan) Set_prev_participant(_prev_participant []byte) {
	t._prev_participant = _prev_participant
}

func (t *TL_channelAdminLogEventActionParticipantToggleBan) Get_prev_participant() []byte {
	return t._prev_participant
}

func (t *TL_channelAdminLogEventActionParticipantToggleBan) Set_new_participant(_new_participant []byte) {
	t._new_participant = _new_participant
}

func (t *TL_channelAdminLogEventActionParticipantToggleBan) Get_new_participant() []byte {
	return t._new_participant
}

func New_TL_channelAdminLogEventActionParticipantToggleBan() *TL_channelAdminLogEventActionParticipantToggleBan {
	return &TL_channelAdminLogEventActionParticipantToggleBan{}
}

func (t *TL_channelAdminLogEventActionParticipantToggleBan) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channelAdminLogEventActionParticipantToggleBan))
	ec.Bytes(t.Get_prev_participant())
	ec.Bytes(t.Get_new_participant())

	return ec.GetBuffer()
}

func (t *TL_channelAdminLogEventActionParticipantToggleBan) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._prev_participant = dc.Bytes(16)
	t._new_participant = dc.Bytes(16)

}

// channelAdminLogEventActionParticipantToggleAdmin#d5676710
type TL_channelAdminLogEventActionParticipantToggleAdmin struct {
	_prev_participant []byte
	_new_participant  []byte
}

func (t *TL_channelAdminLogEventActionParticipantToggleAdmin) Set_prev_participant(_prev_participant []byte) {
	t._prev_participant = _prev_participant
}

func (t *TL_channelAdminLogEventActionParticipantToggleAdmin) Get_prev_participant() []byte {
	return t._prev_participant
}

func (t *TL_channelAdminLogEventActionParticipantToggleAdmin) Set_new_participant(_new_participant []byte) {
	t._new_participant = _new_participant
}

func (t *TL_channelAdminLogEventActionParticipantToggleAdmin) Get_new_participant() []byte {
	return t._new_participant
}

func New_TL_channelAdminLogEventActionParticipantToggleAdmin() *TL_channelAdminLogEventActionParticipantToggleAdmin {
	return &TL_channelAdminLogEventActionParticipantToggleAdmin{}
}

func (t *TL_channelAdminLogEventActionParticipantToggleAdmin) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channelAdminLogEventActionParticipantToggleAdmin))
	ec.Bytes(t.Get_prev_participant())
	ec.Bytes(t.Get_new_participant())

	return ec.GetBuffer()
}

func (t *TL_channelAdminLogEventActionParticipantToggleAdmin) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._prev_participant = dc.Bytes(16)
	t._new_participant = dc.Bytes(16)

}

// channelAdminLogEventActionChangeStickerSet#b1c3caa7
type TL_channelAdminLogEventActionChangeStickerSet struct {
	_prev_stickerset []byte
	_new_stickerset  []byte
}

func (t *TL_channelAdminLogEventActionChangeStickerSet) Set_prev_stickerset(_prev_stickerset []byte) {
	t._prev_stickerset = _prev_stickerset
}

func (t *TL_channelAdminLogEventActionChangeStickerSet) Get_prev_stickerset() []byte {
	return t._prev_stickerset
}

func (t *TL_channelAdminLogEventActionChangeStickerSet) Set_new_stickerset(_new_stickerset []byte) {
	t._new_stickerset = _new_stickerset
}

func (t *TL_channelAdminLogEventActionChangeStickerSet) Get_new_stickerset() []byte {
	return t._new_stickerset
}

func New_TL_channelAdminLogEventActionChangeStickerSet() *TL_channelAdminLogEventActionChangeStickerSet {
	return &TL_channelAdminLogEventActionChangeStickerSet{}
}

func (t *TL_channelAdminLogEventActionChangeStickerSet) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channelAdminLogEventActionChangeStickerSet))
	ec.Bytes(t.Get_prev_stickerset())
	ec.Bytes(t.Get_new_stickerset())

	return ec.GetBuffer()
}

func (t *TL_channelAdminLogEventActionChangeStickerSet) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._prev_stickerset = dc.Bytes(16)
	t._new_stickerset = dc.Bytes(16)

}

// channelAdminLogEventActionTogglePreHistoryHidden#5f5c95f1
type TL_channelAdminLogEventActionTogglePreHistoryHidden struct {
	_new_value bool
}

func (t *TL_channelAdminLogEventActionTogglePreHistoryHidden) Set_new_value(_new_value bool) {
	t._new_value = _new_value
}

func (t *TL_channelAdminLogEventActionTogglePreHistoryHidden) Get_new_value() bool {
	return t._new_value
}

func New_TL_channelAdminLogEventActionTogglePreHistoryHidden() *TL_channelAdminLogEventActionTogglePreHistoryHidden {
	return &TL_channelAdminLogEventActionTogglePreHistoryHidden{}
}

func (t *TL_channelAdminLogEventActionTogglePreHistoryHidden) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channelAdminLogEventActionTogglePreHistoryHidden))
	ec.Bool(t.Get_new_value())

	return ec.GetBuffer()
}

func (t *TL_channelAdminLogEventActionTogglePreHistoryHidden) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._new_value = dc.Bool()

}

// channelAdminLogEvent#3b5a3e40
type TL_channelAdminLogEvent struct {
	_id      int64
	_date    int32
	_user_id int32
	_action  []byte
}

func (t *TL_channelAdminLogEvent) Set_id(_id int64) {
	t._id = _id
}

func (t *TL_channelAdminLogEvent) Get_id() int64 {
	return t._id
}

func (t *TL_channelAdminLogEvent) Set_date(_date int32) {
	t._date = _date
}

func (t *TL_channelAdminLogEvent) Get_date() int32 {
	return t._date
}

func (t *TL_channelAdminLogEvent) Set_user_id(_user_id int32) {
	t._user_id = _user_id
}

func (t *TL_channelAdminLogEvent) Get_user_id() int32 {
	return t._user_id
}

func (t *TL_channelAdminLogEvent) Set_action(_action []byte) {
	t._action = _action
}

func (t *TL_channelAdminLogEvent) Get_action() []byte {
	return t._action
}

func New_TL_channelAdminLogEvent() *TL_channelAdminLogEvent {
	return &TL_channelAdminLogEvent{}
}

func (t *TL_channelAdminLogEvent) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channelAdminLogEvent))
	ec.Long(t.Get_id())
	ec.Int(t.Get_date())
	ec.Int(t.Get_user_id())
	ec.Bytes(t.Get_action())

	return ec.GetBuffer()
}

func (t *TL_channelAdminLogEvent) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Long()
	t._date = dc.Int()
	t._user_id = dc.Int()
	t._action = dc.Bytes(16)

}

// channels_adminLogResults#ed8af74d
type TL_channels_adminLogResults struct {
	_events []byte
	_chats  []byte
	_users  []byte
}

func (t *TL_channels_adminLogResults) Set_events(_events []byte) {
	t._events = _events
}

func (t *TL_channels_adminLogResults) Get_events() []byte {
	return t._events
}

func (t *TL_channels_adminLogResults) Set_chats(_chats []byte) {
	t._chats = _chats
}

func (t *TL_channels_adminLogResults) Get_chats() []byte {
	return t._chats
}

func (t *TL_channels_adminLogResults) Set_users(_users []byte) {
	t._users = _users
}

func (t *TL_channels_adminLogResults) Get_users() []byte {
	return t._users
}

func New_TL_channels_adminLogResults() *TL_channels_adminLogResults {
	return &TL_channels_adminLogResults{}
}

func (t *TL_channels_adminLogResults) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channels_adminLogResults))
	ec.Bytes(t.Get_events())
	ec.Bytes(t.Get_chats())
	ec.Bytes(t.Get_users())

	return ec.GetBuffer()
}

func (t *TL_channels_adminLogResults) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._events = dc.Bytes(16)
	t._chats = dc.Bytes(16)
	t._users = dc.Bytes(16)

}

// channelAdminLogEventsFilter#ea107ae4
type TL_channelAdminLogEventsFilter struct {
	_flags    []byte
	_join     []byte
	_leave    []byte
	_invite   []byte
	_ban      []byte
	_unban    []byte
	_kick     []byte
	_unkick   []byte
	_promote  []byte
	_demote   []byte
	_info     []byte
	_settings []byte
	_pinned   []byte
	_edit     []byte
	_delete   []byte
}

func (t *TL_channelAdminLogEventsFilter) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_channelAdminLogEventsFilter) Get_flags() []byte {
	return t._flags
}

func (t *TL_channelAdminLogEventsFilter) Set_join(_join []byte) {
	t._join = _join
}

func (t *TL_channelAdminLogEventsFilter) Get_join() []byte {
	return t._join
}

func (t *TL_channelAdminLogEventsFilter) Set_leave(_leave []byte) {
	t._leave = _leave
}

func (t *TL_channelAdminLogEventsFilter) Get_leave() []byte {
	return t._leave
}

func (t *TL_channelAdminLogEventsFilter) Set_invite(_invite []byte) {
	t._invite = _invite
}

func (t *TL_channelAdminLogEventsFilter) Get_invite() []byte {
	return t._invite
}

func (t *TL_channelAdminLogEventsFilter) Set_ban(_ban []byte) {
	t._ban = _ban
}

func (t *TL_channelAdminLogEventsFilter) Get_ban() []byte {
	return t._ban
}

func (t *TL_channelAdminLogEventsFilter) Set_unban(_unban []byte) {
	t._unban = _unban
}

func (t *TL_channelAdminLogEventsFilter) Get_unban() []byte {
	return t._unban
}

func (t *TL_channelAdminLogEventsFilter) Set_kick(_kick []byte) {
	t._kick = _kick
}

func (t *TL_channelAdminLogEventsFilter) Get_kick() []byte {
	return t._kick
}

func (t *TL_channelAdminLogEventsFilter) Set_unkick(_unkick []byte) {
	t._unkick = _unkick
}

func (t *TL_channelAdminLogEventsFilter) Get_unkick() []byte {
	return t._unkick
}

func (t *TL_channelAdminLogEventsFilter) Set_promote(_promote []byte) {
	t._promote = _promote
}

func (t *TL_channelAdminLogEventsFilter) Get_promote() []byte {
	return t._promote
}

func (t *TL_channelAdminLogEventsFilter) Set_demote(_demote []byte) {
	t._demote = _demote
}

func (t *TL_channelAdminLogEventsFilter) Get_demote() []byte {
	return t._demote
}

func (t *TL_channelAdminLogEventsFilter) Set_info(_info []byte) {
	t._info = _info
}

func (t *TL_channelAdminLogEventsFilter) Get_info() []byte {
	return t._info
}

func (t *TL_channelAdminLogEventsFilter) Set_settings(_settings []byte) {
	t._settings = _settings
}

func (t *TL_channelAdminLogEventsFilter) Get_settings() []byte {
	return t._settings
}

func (t *TL_channelAdminLogEventsFilter) Set_pinned(_pinned []byte) {
	t._pinned = _pinned
}

func (t *TL_channelAdminLogEventsFilter) Get_pinned() []byte {
	return t._pinned
}

func (t *TL_channelAdminLogEventsFilter) Set_edit(_edit []byte) {
	t._edit = _edit
}

func (t *TL_channelAdminLogEventsFilter) Get_edit() []byte {
	return t._edit
}

func (t *TL_channelAdminLogEventsFilter) Set_delete(_delete []byte) {
	t._delete = _delete
}

func (t *TL_channelAdminLogEventsFilter) Get_delete() []byte {
	return t._delete
}

func New_TL_channelAdminLogEventsFilter() *TL_channelAdminLogEventsFilter {
	return &TL_channelAdminLogEventsFilter{}
}

func (t *TL_channelAdminLogEventsFilter) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channelAdminLogEventsFilter))
	ec.Bytes(t.Get_join())
	ec.Bytes(t.Get_leave())
	ec.Bytes(t.Get_invite())
	ec.Bytes(t.Get_ban())
	ec.Bytes(t.Get_unban())
	ec.Bytes(t.Get_kick())
	ec.Bytes(t.Get_unkick())
	ec.Bytes(t.Get_promote())
	ec.Bytes(t.Get_demote())
	ec.Bytes(t.Get_info())
	ec.Bytes(t.Get_settings())
	ec.Bytes(t.Get_pinned())
	ec.Bytes(t.Get_edit())
	ec.Bytes(t.Get_delete())

	return ec.GetBuffer()
}

func (t *TL_channelAdminLogEventsFilter) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._join = dc.Bytes(16)
	t._leave = dc.Bytes(16)
	t._invite = dc.Bytes(16)
	t._ban = dc.Bytes(16)
	t._unban = dc.Bytes(16)
	t._kick = dc.Bytes(16)
	t._unkick = dc.Bytes(16)
	t._promote = dc.Bytes(16)
	t._demote = dc.Bytes(16)
	t._info = dc.Bytes(16)
	t._settings = dc.Bytes(16)
	t._pinned = dc.Bytes(16)
	t._edit = dc.Bytes(16)
	t._delete = dc.Bytes(16)

}

// popularContact#5ce14175
type TL_popularContact struct {
	_client_id int64
	_importers int32
}

func (t *TL_popularContact) Set_client_id(_client_id int64) {
	t._client_id = _client_id
}

func (t *TL_popularContact) Get_client_id() int64 {
	return t._client_id
}

func (t *TL_popularContact) Set_importers(_importers int32) {
	t._importers = _importers
}

func (t *TL_popularContact) Get_importers() int32 {
	return t._importers
}

func New_TL_popularContact() *TL_popularContact {
	return &TL_popularContact{}
}

func (t *TL_popularContact) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_popularContact))
	ec.Long(t.Get_client_id())
	ec.Int(t.Get_importers())

	return ec.GetBuffer()
}

func (t *TL_popularContact) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._client_id = dc.Long()
	t._importers = dc.Int()

}

// cdnFileHash#77eec38f
type TL_cdnFileHash struct {
	_offset int32
	_limit  int32
	_hash   []byte
}

func (t *TL_cdnFileHash) Set_offset(_offset int32) {
	t._offset = _offset
}

func (t *TL_cdnFileHash) Get_offset() int32 {
	return t._offset
}

func (t *TL_cdnFileHash) Set_limit(_limit int32) {
	t._limit = _limit
}

func (t *TL_cdnFileHash) Get_limit() int32 {
	return t._limit
}

func (t *TL_cdnFileHash) Set_hash(_hash []byte) {
	t._hash = _hash
}

func (t *TL_cdnFileHash) Get_hash() []byte {
	return t._hash
}

func New_TL_cdnFileHash() *TL_cdnFileHash {
	return &TL_cdnFileHash{}
}

func (t *TL_cdnFileHash) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_cdnFileHash))
	ec.Int(t.Get_offset())
	ec.Int(t.Get_limit())
	ec.Bytes(t.Get_hash())

	return ec.GetBuffer()
}

func (t *TL_cdnFileHash) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._offset = dc.Int()
	t._limit = dc.Int()
	t._hash = dc.Bytes(16)

}

// messages_favedStickersNotModified#9e8fa6d3
type TL_messages_favedStickersNotModified struct {
}

func New_TL_messages_favedStickersNotModified() *TL_messages_favedStickersNotModified {
	return &TL_messages_favedStickersNotModified{}
}

func (t *TL_messages_favedStickersNotModified) Encode() []byte {
	return nil
}

func (t *TL_messages_favedStickersNotModified) Decode(b []byte) {

}

// messages_favedStickers#f37f2f16
type TL_messages_favedStickers struct {
	_hash     int32
	_packs    []byte
	_stickers []byte
}

func (t *TL_messages_favedStickers) Set_hash(_hash int32) {
	t._hash = _hash
}

func (t *TL_messages_favedStickers) Get_hash() int32 {
	return t._hash
}

func (t *TL_messages_favedStickers) Set_packs(_packs []byte) {
	t._packs = _packs
}

func (t *TL_messages_favedStickers) Get_packs() []byte {
	return t._packs
}

func (t *TL_messages_favedStickers) Set_stickers(_stickers []byte) {
	t._stickers = _stickers
}

func (t *TL_messages_favedStickers) Get_stickers() []byte {
	return t._stickers
}

func New_TL_messages_favedStickers() *TL_messages_favedStickers {
	return &TL_messages_favedStickers{}
}

func (t *TL_messages_favedStickers) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_favedStickers))
	ec.Int(t.Get_hash())
	ec.Bytes(t.Get_packs())
	ec.Bytes(t.Get_stickers())

	return ec.GetBuffer()
}

func (t *TL_messages_favedStickers) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._hash = dc.Int()
	t._packs = dc.Bytes(16)
	t._stickers = dc.Bytes(16)

}

// recentMeUrlUnknown#46e1d13d
type TL_recentMeUrlUnknown struct {
	_url string
}

func (t *TL_recentMeUrlUnknown) Set_url(_url string) {
	t._url = _url
}

func (t *TL_recentMeUrlUnknown) Get_url() string {
	return t._url
}

func New_TL_recentMeUrlUnknown() *TL_recentMeUrlUnknown {
	return &TL_recentMeUrlUnknown{}
}

func (t *TL_recentMeUrlUnknown) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_recentMeUrlUnknown))
	ec.String(t.Get_url())

	return ec.GetBuffer()
}

func (t *TL_recentMeUrlUnknown) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._url = dc.String()

}

// recentMeUrlUser#8dbc3336
type TL_recentMeUrlUser struct {
	_url     string
	_user_id int32
}

func (t *TL_recentMeUrlUser) Set_url(_url string) {
	t._url = _url
}

func (t *TL_recentMeUrlUser) Get_url() string {
	return t._url
}

func (t *TL_recentMeUrlUser) Set_user_id(_user_id int32) {
	t._user_id = _user_id
}

func (t *TL_recentMeUrlUser) Get_user_id() int32 {
	return t._user_id
}

func New_TL_recentMeUrlUser() *TL_recentMeUrlUser {
	return &TL_recentMeUrlUser{}
}

func (t *TL_recentMeUrlUser) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_recentMeUrlUser))
	ec.String(t.Get_url())
	ec.Int(t.Get_user_id())

	return ec.GetBuffer()
}

func (t *TL_recentMeUrlUser) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._url = dc.String()
	t._user_id = dc.Int()

}

// recentMeUrlChat#a01b22f9
type TL_recentMeUrlChat struct {
	_url     string
	_chat_id int32
}

func (t *TL_recentMeUrlChat) Set_url(_url string) {
	t._url = _url
}

func (t *TL_recentMeUrlChat) Get_url() string {
	return t._url
}

func (t *TL_recentMeUrlChat) Set_chat_id(_chat_id int32) {
	t._chat_id = _chat_id
}

func (t *TL_recentMeUrlChat) Get_chat_id() int32 {
	return t._chat_id
}

func New_TL_recentMeUrlChat() *TL_recentMeUrlChat {
	return &TL_recentMeUrlChat{}
}

func (t *TL_recentMeUrlChat) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_recentMeUrlChat))
	ec.String(t.Get_url())
	ec.Int(t.Get_chat_id())

	return ec.GetBuffer()
}

func (t *TL_recentMeUrlChat) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._url = dc.String()
	t._chat_id = dc.Int()

}

// recentMeUrlChatInvite#eb49081d
type TL_recentMeUrlChatInvite struct {
	_url         string
	_chat_invite []byte
}

func (t *TL_recentMeUrlChatInvite) Set_url(_url string) {
	t._url = _url
}

func (t *TL_recentMeUrlChatInvite) Get_url() string {
	return t._url
}

func (t *TL_recentMeUrlChatInvite) Set_chat_invite(_chat_invite []byte) {
	t._chat_invite = _chat_invite
}

func (t *TL_recentMeUrlChatInvite) Get_chat_invite() []byte {
	return t._chat_invite
}

func New_TL_recentMeUrlChatInvite() *TL_recentMeUrlChatInvite {
	return &TL_recentMeUrlChatInvite{}
}

func (t *TL_recentMeUrlChatInvite) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_recentMeUrlChatInvite))
	ec.String(t.Get_url())
	ec.Bytes(t.Get_chat_invite())

	return ec.GetBuffer()
}

func (t *TL_recentMeUrlChatInvite) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._url = dc.String()
	t._chat_invite = dc.Bytes(16)

}

// recentMeUrlStickerSet#bc0a57dc
type TL_recentMeUrlStickerSet struct {
	_url string
	_set []byte
}

func (t *TL_recentMeUrlStickerSet) Set_url(_url string) {
	t._url = _url
}

func (t *TL_recentMeUrlStickerSet) Get_url() string {
	return t._url
}

func (t *TL_recentMeUrlStickerSet) Set_set(_set []byte) {
	t._set = _set
}

func (t *TL_recentMeUrlStickerSet) Get_set() []byte {
	return t._set
}

func New_TL_recentMeUrlStickerSet() *TL_recentMeUrlStickerSet {
	return &TL_recentMeUrlStickerSet{}
}

func (t *TL_recentMeUrlStickerSet) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_recentMeUrlStickerSet))
	ec.String(t.Get_url())
	ec.Bytes(t.Get_set())

	return ec.GetBuffer()
}

func (t *TL_recentMeUrlStickerSet) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._url = dc.String()
	t._set = dc.Bytes(16)

}

// help_recentMeUrls#0e0310d7
type TL_help_recentMeUrls struct {
	_urls  []byte
	_chats []byte
	_users []byte
}

func (t *TL_help_recentMeUrls) Set_urls(_urls []byte) {
	t._urls = _urls
}

func (t *TL_help_recentMeUrls) Get_urls() []byte {
	return t._urls
}

func (t *TL_help_recentMeUrls) Set_chats(_chats []byte) {
	t._chats = _chats
}

func (t *TL_help_recentMeUrls) Get_chats() []byte {
	return t._chats
}

func (t *TL_help_recentMeUrls) Set_users(_users []byte) {
	t._users = _users
}

func (t *TL_help_recentMeUrls) Get_users() []byte {
	return t._users
}

func New_TL_help_recentMeUrls() *TL_help_recentMeUrls {
	return &TL_help_recentMeUrls{}
}

func (t *TL_help_recentMeUrls) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_help_recentMeUrls))
	ec.Bytes(t.Get_urls())
	ec.Bytes(t.Get_chats())
	ec.Bytes(t.Get_users())

	return ec.GetBuffer()
}

func (t *TL_help_recentMeUrls) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._urls = dc.Bytes(16)
	t._chats = dc.Bytes(16)
	t._users = dc.Bytes(16)

}

// inputSingleMedia#5eaa7809
type TL_inputSingleMedia struct {
	_media     []byte
	_random_id int64
}

func (t *TL_inputSingleMedia) Set_media(_media []byte) {
	t._media = _media
}

func (t *TL_inputSingleMedia) Get_media() []byte {
	return t._media
}

func (t *TL_inputSingleMedia) Set_random_id(_random_id int64) {
	t._random_id = _random_id
}

func (t *TL_inputSingleMedia) Get_random_id() int64 {
	return t._random_id
}

func New_TL_inputSingleMedia() *TL_inputSingleMedia {
	return &TL_inputSingleMedia{}
}

func (t *TL_inputSingleMedia) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_inputSingleMedia))
	ec.Bytes(t.Get_media())
	ec.Long(t.Get_random_id())

	return ec.GetBuffer()
}

func (t *TL_inputSingleMedia) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._media = dc.Bytes(16)
	t._random_id = dc.Long()

}

// invokeAfterMsg#cb9f372d
type TL_invokeAfterMsg struct {
	_msg_id int64
	_query  []byte
}

func (t *TL_invokeAfterMsg) Set_msg_id(_msg_id int64) {
	t._msg_id = _msg_id
}

func (t *TL_invokeAfterMsg) Get_msg_id() int64 {
	return t._msg_id
}

func (t *TL_invokeAfterMsg) Set_query(_query []byte) {
	t._query = _query
}

func (t *TL_invokeAfterMsg) Get_query() []byte {
	return t._query
}

func New_TL_invokeAfterMsg() *TL_invokeAfterMsg {
	return &TL_invokeAfterMsg{}
}

func (t *TL_invokeAfterMsg) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_invokeAfterMsg))
	ec.Long(t.Get_msg_id())
	ec.Bytes(t.Get_query())

	return ec.GetBuffer()
}

func (t *TL_invokeAfterMsg) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._msg_id = dc.Long()
	t._query = dc.Bytes(16)

}

// invokeAfterMsgs#3dc4b4f0
type TL_invokeAfterMsgs struct {
	_msg_ids []byte
	_query   []byte
}

func (t *TL_invokeAfterMsgs) Set_msg_ids(_msg_ids []byte) {
	t._msg_ids = _msg_ids
}

func (t *TL_invokeAfterMsgs) Get_msg_ids() []byte {
	return t._msg_ids
}

func (t *TL_invokeAfterMsgs) Set_query(_query []byte) {
	t._query = _query
}

func (t *TL_invokeAfterMsgs) Get_query() []byte {
	return t._query
}

func New_TL_invokeAfterMsgs() *TL_invokeAfterMsgs {
	return &TL_invokeAfterMsgs{}
}

func (t *TL_invokeAfterMsgs) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_invokeAfterMsgs))
	ec.Bytes(t.Get_msg_ids())
	ec.Bytes(t.Get_query())

	return ec.GetBuffer()
}

func (t *TL_invokeAfterMsgs) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._msg_ids = dc.Bytes(16)
	t._query = dc.Bytes(16)

}

// initConnection#c7481da6
type TL_initConnection struct {
	_api_id           int32
	_device_model     string
	_system_version   string
	_app_version      string
	_system_lang_code string
	_lang_pack        string
	_lang_code        string
	_query            []byte
}

func (t *TL_initConnection) Set_api_id(_api_id int32) {
	t._api_id = _api_id
}

func (t *TL_initConnection) Get_api_id() int32 {
	return t._api_id
}

func (t *TL_initConnection) Set_device_model(_device_model string) {
	t._device_model = _device_model
}

func (t *TL_initConnection) Get_device_model() string {
	return t._device_model
}

func (t *TL_initConnection) Set_system_version(_system_version string) {
	t._system_version = _system_version
}

func (t *TL_initConnection) Get_system_version() string {
	return t._system_version
}

func (t *TL_initConnection) Set_app_version(_app_version string) {
	t._app_version = _app_version
}

func (t *TL_initConnection) Get_app_version() string {
	return t._app_version
}

func (t *TL_initConnection) Set_system_lang_code(_system_lang_code string) {
	t._system_lang_code = _system_lang_code
}

func (t *TL_initConnection) Get_system_lang_code() string {
	return t._system_lang_code
}

func (t *TL_initConnection) Set_lang_pack(_lang_pack string) {
	t._lang_pack = _lang_pack
}

func (t *TL_initConnection) Get_lang_pack() string {
	return t._lang_pack
}

func (t *TL_initConnection) Set_lang_code(_lang_code string) {
	t._lang_code = _lang_code
}

func (t *TL_initConnection) Get_lang_code() string {
	return t._lang_code
}

func (t *TL_initConnection) Set_query(_query []byte) {
	t._query = _query
}

func (t *TL_initConnection) Get_query() []byte {
	return t._query
}

func New_TL_initConnection() *TL_initConnection {
	return &TL_initConnection{}
}

func (t *TL_initConnection) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_initConnection))
	ec.Int(t.Get_api_id())
	ec.String(t.Get_device_model())
	ec.String(t.Get_system_version())
	ec.String(t.Get_app_version())
	ec.String(t.Get_system_lang_code())
	ec.String(t.Get_lang_pack())
	ec.String(t.Get_lang_code())
	ec.Bytes(t.Get_query())

	return ec.GetBuffer()
}

func (t *TL_initConnection) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._api_id = dc.Int()
	t._device_model = dc.String()
	t._system_version = dc.String()
	t._app_version = dc.String()
	t._system_lang_code = dc.String()
	t._lang_pack = dc.String()
	t._lang_code = dc.String()
	t._query = dc.Bytes(16)

}

// invokeWithLayer#da9b0d0d
type TL_invokeWithLayer struct {
	_layer int32
	_query []byte
}

func (t *TL_invokeWithLayer) Set_layer(_layer int32) {
	t._layer = _layer
}

func (t *TL_invokeWithLayer) Get_layer() int32 {
	return t._layer
}

func (t *TL_invokeWithLayer) Set_query(_query []byte) {
	t._query = _query
}

func (t *TL_invokeWithLayer) Get_query() []byte {
	return t._query
}

func New_TL_invokeWithLayer() *TL_invokeWithLayer {
	return &TL_invokeWithLayer{}
}

func (t *TL_invokeWithLayer) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_invokeWithLayer))
	ec.Int(t.Get_layer())
	ec.Bytes(t.Get_query())

	return ec.GetBuffer()
}

func (t *TL_invokeWithLayer) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._layer = dc.Int()
	t._query = dc.Bytes(16)

}

// invokeWithoutUpdates#bf9459b7
type TL_invokeWithoutUpdates struct {
	_query []byte
}

func (t *TL_invokeWithoutUpdates) Set_query(_query []byte) {
	t._query = _query
}

func (t *TL_invokeWithoutUpdates) Get_query() []byte {
	return t._query
}

func New_TL_invokeWithoutUpdates() *TL_invokeWithoutUpdates {
	return &TL_invokeWithoutUpdates{}
}

func (t *TL_invokeWithoutUpdates) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_invokeWithoutUpdates))
	ec.Bytes(t.Get_query())

	return ec.GetBuffer()
}

func (t *TL_invokeWithoutUpdates) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._query = dc.Bytes(16)

}

// auth_checkPhone#6fe51dfb
type TL_auth_checkPhone struct {
	_phone_number string
}

func (t *TL_auth_checkPhone) Set_phone_number(_phone_number string) {
	t._phone_number = _phone_number
}

func (t *TL_auth_checkPhone) Get_phone_number() string {
	return t._phone_number
}

func New_TL_auth_checkPhone() *TL_auth_checkPhone {
	return &TL_auth_checkPhone{}
}

func (t *TL_auth_checkPhone) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_auth_checkPhone))
	ec.String(t.Get_phone_number())

	return ec.GetBuffer()
}

func (t *TL_auth_checkPhone) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._phone_number = dc.String()

}

// auth_sendCode#86aef0ec
type TL_auth_sendCode struct {
	_flags           []byte
	_allow_flashcall []byte
	_phone_number    string
	_current_number  []byte
	_api_id          int32
	_api_hash        string
}

func (t *TL_auth_sendCode) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_auth_sendCode) Get_flags() []byte {
	return t._flags
}

func (t *TL_auth_sendCode) Set_allow_flashcall(_allow_flashcall []byte) {
	t._allow_flashcall = _allow_flashcall
}

func (t *TL_auth_sendCode) Get_allow_flashcall() []byte {
	return t._allow_flashcall
}

func (t *TL_auth_sendCode) Set_phone_number(_phone_number string) {
	t._phone_number = _phone_number
}

func (t *TL_auth_sendCode) Get_phone_number() string {
	return t._phone_number
}

func (t *TL_auth_sendCode) Set_current_number(_current_number []byte) {
	t._current_number = _current_number
}

func (t *TL_auth_sendCode) Get_current_number() []byte {
	return t._current_number
}

func (t *TL_auth_sendCode) Set_api_id(_api_id int32) {
	t._api_id = _api_id
}

func (t *TL_auth_sendCode) Get_api_id() int32 {
	return t._api_id
}

func (t *TL_auth_sendCode) Set_api_hash(_api_hash string) {
	t._api_hash = _api_hash
}

func (t *TL_auth_sendCode) Get_api_hash() string {
	return t._api_hash
}

func New_TL_auth_sendCode() *TL_auth_sendCode {
	return &TL_auth_sendCode{}
}

func (t *TL_auth_sendCode) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_auth_sendCode))
	ec.Bytes(t.Get_allow_flashcall())
	ec.String(t.Get_phone_number())
	ec.Bytes(t.Get_current_number())
	ec.Int(t.Get_api_id())
	ec.String(t.Get_api_hash())

	return ec.GetBuffer()
}

func (t *TL_auth_sendCode) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._allow_flashcall = dc.Bytes(16)
	t._phone_number = dc.String()
	t._current_number = dc.Bytes(16)
	t._api_id = dc.Int()
	t._api_hash = dc.String()

}

// auth_signUp#1b067634
type TL_auth_signUp struct {
	_phone_number    string
	_phone_code_hash string
	_phone_code      string
	_first_name      string
	_last_name       string
}

func (t *TL_auth_signUp) Set_phone_number(_phone_number string) {
	t._phone_number = _phone_number
}

func (t *TL_auth_signUp) Get_phone_number() string {
	return t._phone_number
}

func (t *TL_auth_signUp) Set_phone_code_hash(_phone_code_hash string) {
	t._phone_code_hash = _phone_code_hash
}

func (t *TL_auth_signUp) Get_phone_code_hash() string {
	return t._phone_code_hash
}

func (t *TL_auth_signUp) Set_phone_code(_phone_code string) {
	t._phone_code = _phone_code
}

func (t *TL_auth_signUp) Get_phone_code() string {
	return t._phone_code
}

func (t *TL_auth_signUp) Set_first_name(_first_name string) {
	t._first_name = _first_name
}

func (t *TL_auth_signUp) Get_first_name() string {
	return t._first_name
}

func (t *TL_auth_signUp) Set_last_name(_last_name string) {
	t._last_name = _last_name
}

func (t *TL_auth_signUp) Get_last_name() string {
	return t._last_name
}

func New_TL_auth_signUp() *TL_auth_signUp {
	return &TL_auth_signUp{}
}

func (t *TL_auth_signUp) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_auth_signUp))
	ec.String(t.Get_phone_number())
	ec.String(t.Get_phone_code_hash())
	ec.String(t.Get_phone_code())
	ec.String(t.Get_first_name())
	ec.String(t.Get_last_name())

	return ec.GetBuffer()
}

func (t *TL_auth_signUp) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._phone_number = dc.String()
	t._phone_code_hash = dc.String()
	t._phone_code = dc.String()
	t._first_name = dc.String()
	t._last_name = dc.String()

}

// auth_signIn#bcd51581
type TL_auth_signIn struct {
	_phone_number    string
	_phone_code_hash string
	_phone_code      string
}

func (t *TL_auth_signIn) Set_phone_number(_phone_number string) {
	t._phone_number = _phone_number
}

func (t *TL_auth_signIn) Get_phone_number() string {
	return t._phone_number
}

func (t *TL_auth_signIn) Set_phone_code_hash(_phone_code_hash string) {
	t._phone_code_hash = _phone_code_hash
}

func (t *TL_auth_signIn) Get_phone_code_hash() string {
	return t._phone_code_hash
}

func (t *TL_auth_signIn) Set_phone_code(_phone_code string) {
	t._phone_code = _phone_code
}

func (t *TL_auth_signIn) Get_phone_code() string {
	return t._phone_code
}

func New_TL_auth_signIn() *TL_auth_signIn {
	return &TL_auth_signIn{}
}

func (t *TL_auth_signIn) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_auth_signIn))
	ec.String(t.Get_phone_number())
	ec.String(t.Get_phone_code_hash())
	ec.String(t.Get_phone_code())

	return ec.GetBuffer()
}

func (t *TL_auth_signIn) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._phone_number = dc.String()
	t._phone_code_hash = dc.String()
	t._phone_code = dc.String()

}

// auth_logOut#5717da40
type TL_auth_logOut struct {
}

func New_TL_auth_logOut() *TL_auth_logOut {
	return &TL_auth_logOut{}
}

func (t *TL_auth_logOut) Encode() []byte {
	return nil
}

func (t *TL_auth_logOut) Decode(b []byte) {

}

// auth_resetAuthorizations#9fab0d1a
type TL_auth_resetAuthorizations struct {
}

func New_TL_auth_resetAuthorizations() *TL_auth_resetAuthorizations {
	return &TL_auth_resetAuthorizations{}
}

func (t *TL_auth_resetAuthorizations) Encode() []byte {
	return nil
}

func (t *TL_auth_resetAuthorizations) Decode(b []byte) {

}

// auth_sendInvites#771c1d97
type TL_auth_sendInvites struct {
	_phone_numbers []byte
	_message       string
}

func (t *TL_auth_sendInvites) Set_phone_numbers(_phone_numbers []byte) {
	t._phone_numbers = _phone_numbers
}

func (t *TL_auth_sendInvites) Get_phone_numbers() []byte {
	return t._phone_numbers
}

func (t *TL_auth_sendInvites) Set_message(_message string) {
	t._message = _message
}

func (t *TL_auth_sendInvites) Get_message() string {
	return t._message
}

func New_TL_auth_sendInvites() *TL_auth_sendInvites {
	return &TL_auth_sendInvites{}
}

func (t *TL_auth_sendInvites) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_auth_sendInvites))
	ec.Bytes(t.Get_phone_numbers())
	ec.String(t.Get_message())

	return ec.GetBuffer()
}

func (t *TL_auth_sendInvites) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._phone_numbers = dc.Bytes(16)
	t._message = dc.String()

}

// auth_exportAuthorization#e5bfffcd
type TL_auth_exportAuthorization struct {
	_dc_id int32
}

func (t *TL_auth_exportAuthorization) Set_dc_id(_dc_id int32) {
	t._dc_id = _dc_id
}

func (t *TL_auth_exportAuthorization) Get_dc_id() int32 {
	return t._dc_id
}

func New_TL_auth_exportAuthorization() *TL_auth_exportAuthorization {
	return &TL_auth_exportAuthorization{}
}

func (t *TL_auth_exportAuthorization) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_auth_exportAuthorization))
	ec.Int(t.Get_dc_id())

	return ec.GetBuffer()
}

func (t *TL_auth_exportAuthorization) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._dc_id = dc.Int()

}

// auth_importAuthorization#e3ef9613
type TL_auth_importAuthorization struct {
	_id    int32
	_bytes []byte
}

func (t *TL_auth_importAuthorization) Set_id(_id int32) {
	t._id = _id
}

func (t *TL_auth_importAuthorization) Get_id() int32 {
	return t._id
}

func (t *TL_auth_importAuthorization) Set_bytes(_bytes []byte) {
	t._bytes = _bytes
}

func (t *TL_auth_importAuthorization) Get_bytes() []byte {
	return t._bytes
}

func New_TL_auth_importAuthorization() *TL_auth_importAuthorization {
	return &TL_auth_importAuthorization{}
}

func (t *TL_auth_importAuthorization) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_auth_importAuthorization))
	ec.Int(t.Get_id())
	ec.Bytes(t.Get_bytes())

	return ec.GetBuffer()
}

func (t *TL_auth_importAuthorization) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Int()
	t._bytes = dc.Bytes(16)

}

// auth_bindTempAuthKey#cdd42a05
type TL_auth_bindTempAuthKey struct {
	_perm_auth_key_id  int64
	_nonce             int64
	_expires_at        int32
	_encrypted_message []byte
}

func (t *TL_auth_bindTempAuthKey) Set_perm_auth_key_id(_perm_auth_key_id int64) {
	t._perm_auth_key_id = _perm_auth_key_id
}

func (t *TL_auth_bindTempAuthKey) Get_perm_auth_key_id() int64 {
	return t._perm_auth_key_id
}

func (t *TL_auth_bindTempAuthKey) Set_nonce(_nonce int64) {
	t._nonce = _nonce
}

func (t *TL_auth_bindTempAuthKey) Get_nonce() int64 {
	return t._nonce
}

func (t *TL_auth_bindTempAuthKey) Set_expires_at(_expires_at int32) {
	t._expires_at = _expires_at
}

func (t *TL_auth_bindTempAuthKey) Get_expires_at() int32 {
	return t._expires_at
}

func (t *TL_auth_bindTempAuthKey) Set_encrypted_message(_encrypted_message []byte) {
	t._encrypted_message = _encrypted_message
}

func (t *TL_auth_bindTempAuthKey) Get_encrypted_message() []byte {
	return t._encrypted_message
}

func New_TL_auth_bindTempAuthKey() *TL_auth_bindTempAuthKey {
	return &TL_auth_bindTempAuthKey{}
}

func (t *TL_auth_bindTempAuthKey) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_auth_bindTempAuthKey))
	ec.Long(t.Get_perm_auth_key_id())
	ec.Long(t.Get_nonce())
	ec.Int(t.Get_expires_at())
	ec.Bytes(t.Get_encrypted_message())

	return ec.GetBuffer()
}

func (t *TL_auth_bindTempAuthKey) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._perm_auth_key_id = dc.Long()
	t._nonce = dc.Long()
	t._expires_at = dc.Int()
	t._encrypted_message = dc.Bytes(16)

}

// auth_importBotAuthorization#67a3ff2c
type TL_auth_importBotAuthorization struct {
	_flags          int32
	_api_id         int32
	_api_hash       string
	_bot_auth_token string
}

func (t *TL_auth_importBotAuthorization) Set_flags(_flags int32) {
	t._flags = _flags
}

func (t *TL_auth_importBotAuthorization) Get_flags() int32 {
	return t._flags
}

func (t *TL_auth_importBotAuthorization) Set_api_id(_api_id int32) {
	t._api_id = _api_id
}

func (t *TL_auth_importBotAuthorization) Get_api_id() int32 {
	return t._api_id
}

func (t *TL_auth_importBotAuthorization) Set_api_hash(_api_hash string) {
	t._api_hash = _api_hash
}

func (t *TL_auth_importBotAuthorization) Get_api_hash() string {
	return t._api_hash
}

func (t *TL_auth_importBotAuthorization) Set_bot_auth_token(_bot_auth_token string) {
	t._bot_auth_token = _bot_auth_token
}

func (t *TL_auth_importBotAuthorization) Get_bot_auth_token() string {
	return t._bot_auth_token
}

func New_TL_auth_importBotAuthorization() *TL_auth_importBotAuthorization {
	return &TL_auth_importBotAuthorization{}
}

func (t *TL_auth_importBotAuthorization) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_auth_importBotAuthorization))
	ec.Int(t.Get_flags())
	ec.Int(t.Get_api_id())
	ec.String(t.Get_api_hash())
	ec.String(t.Get_bot_auth_token())

	return ec.GetBuffer()
}

func (t *TL_auth_importBotAuthorization) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._flags = dc.Int()
	t._api_id = dc.Int()
	t._api_hash = dc.String()
	t._bot_auth_token = dc.String()

}

// auth_checkPassword#0a63011e
type TL_auth_checkPassword struct {
	_password_hash []byte
}

func (t *TL_auth_checkPassword) Set_password_hash(_password_hash []byte) {
	t._password_hash = _password_hash
}

func (t *TL_auth_checkPassword) Get_password_hash() []byte {
	return t._password_hash
}

func New_TL_auth_checkPassword() *TL_auth_checkPassword {
	return &TL_auth_checkPassword{}
}

func (t *TL_auth_checkPassword) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_auth_checkPassword))
	ec.Bytes(t.Get_password_hash())

	return ec.GetBuffer()
}

func (t *TL_auth_checkPassword) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._password_hash = dc.Bytes(16)

}

// auth_requestPasswordRecovery#d897bc66
type TL_auth_requestPasswordRecovery struct {
}

func New_TL_auth_requestPasswordRecovery() *TL_auth_requestPasswordRecovery {
	return &TL_auth_requestPasswordRecovery{}
}

func (t *TL_auth_requestPasswordRecovery) Encode() []byte {
	return nil
}

func (t *TL_auth_requestPasswordRecovery) Decode(b []byte) {

}

// auth_recoverPassword#4ea56e92
type TL_auth_recoverPassword struct {
	_code string
}

func (t *TL_auth_recoverPassword) Set_code(_code string) {
	t._code = _code
}

func (t *TL_auth_recoverPassword) Get_code() string {
	return t._code
}

func New_TL_auth_recoverPassword() *TL_auth_recoverPassword {
	return &TL_auth_recoverPassword{}
}

func (t *TL_auth_recoverPassword) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_auth_recoverPassword))
	ec.String(t.Get_code())

	return ec.GetBuffer()
}

func (t *TL_auth_recoverPassword) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._code = dc.String()

}

// auth_resendCode#3ef1a9bf
type TL_auth_resendCode struct {
	_phone_number    string
	_phone_code_hash string
}

func (t *TL_auth_resendCode) Set_phone_number(_phone_number string) {
	t._phone_number = _phone_number
}

func (t *TL_auth_resendCode) Get_phone_number() string {
	return t._phone_number
}

func (t *TL_auth_resendCode) Set_phone_code_hash(_phone_code_hash string) {
	t._phone_code_hash = _phone_code_hash
}

func (t *TL_auth_resendCode) Get_phone_code_hash() string {
	return t._phone_code_hash
}

func New_TL_auth_resendCode() *TL_auth_resendCode {
	return &TL_auth_resendCode{}
}

func (t *TL_auth_resendCode) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_auth_resendCode))
	ec.String(t.Get_phone_number())
	ec.String(t.Get_phone_code_hash())

	return ec.GetBuffer()
}

func (t *TL_auth_resendCode) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._phone_number = dc.String()
	t._phone_code_hash = dc.String()

}

// auth_cancelCode#1f040578
type TL_auth_cancelCode struct {
	_phone_number    string
	_phone_code_hash string
}

func (t *TL_auth_cancelCode) Set_phone_number(_phone_number string) {
	t._phone_number = _phone_number
}

func (t *TL_auth_cancelCode) Get_phone_number() string {
	return t._phone_number
}

func (t *TL_auth_cancelCode) Set_phone_code_hash(_phone_code_hash string) {
	t._phone_code_hash = _phone_code_hash
}

func (t *TL_auth_cancelCode) Get_phone_code_hash() string {
	return t._phone_code_hash
}

func New_TL_auth_cancelCode() *TL_auth_cancelCode {
	return &TL_auth_cancelCode{}
}

func (t *TL_auth_cancelCode) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_auth_cancelCode))
	ec.String(t.Get_phone_number())
	ec.String(t.Get_phone_code_hash())

	return ec.GetBuffer()
}

func (t *TL_auth_cancelCode) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._phone_number = dc.String()
	t._phone_code_hash = dc.String()

}

// auth_dropTempAuthKeys#8e48a188
type TL_auth_dropTempAuthKeys struct {
	_except_auth_keys []byte
}

func (t *TL_auth_dropTempAuthKeys) Set_except_auth_keys(_except_auth_keys []byte) {
	t._except_auth_keys = _except_auth_keys
}

func (t *TL_auth_dropTempAuthKeys) Get_except_auth_keys() []byte {
	return t._except_auth_keys
}

func New_TL_auth_dropTempAuthKeys() *TL_auth_dropTempAuthKeys {
	return &TL_auth_dropTempAuthKeys{}
}

func (t *TL_auth_dropTempAuthKeys) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_auth_dropTempAuthKeys))
	ec.Bytes(t.Get_except_auth_keys())

	return ec.GetBuffer()
}

func (t *TL_auth_dropTempAuthKeys) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._except_auth_keys = dc.Bytes(16)

}

// account_registerDevice#637ea878
type TL_account_registerDevice struct {
	_token_type int32
	_token      string
}

func (t *TL_account_registerDevice) Set_token_type(_token_type int32) {
	t._token_type = _token_type
}

func (t *TL_account_registerDevice) Get_token_type() int32 {
	return t._token_type
}

func (t *TL_account_registerDevice) Set_token(_token string) {
	t._token = _token
}

func (t *TL_account_registerDevice) Get_token() string {
	return t._token
}

func New_TL_account_registerDevice() *TL_account_registerDevice {
	return &TL_account_registerDevice{}
}

func (t *TL_account_registerDevice) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_account_registerDevice))
	ec.Int(t.Get_token_type())
	ec.String(t.Get_token())

	return ec.GetBuffer()
}

func (t *TL_account_registerDevice) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._token_type = dc.Int()
	t._token = dc.String()

}

// account_unregisterDevice#65c55b40
type TL_account_unregisterDevice struct {
	_token_type int32
	_token      string
}

func (t *TL_account_unregisterDevice) Set_token_type(_token_type int32) {
	t._token_type = _token_type
}

func (t *TL_account_unregisterDevice) Get_token_type() int32 {
	return t._token_type
}

func (t *TL_account_unregisterDevice) Set_token(_token string) {
	t._token = _token
}

func (t *TL_account_unregisterDevice) Get_token() string {
	return t._token
}

func New_TL_account_unregisterDevice() *TL_account_unregisterDevice {
	return &TL_account_unregisterDevice{}
}

func (t *TL_account_unregisterDevice) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_account_unregisterDevice))
	ec.Int(t.Get_token_type())
	ec.String(t.Get_token())

	return ec.GetBuffer()
}

func (t *TL_account_unregisterDevice) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._token_type = dc.Int()
	t._token = dc.String()

}

// account_updateNotifySettings#84be5b93
type TL_account_updateNotifySettings struct {
	_peer     []byte
	_settings []byte
}

func (t *TL_account_updateNotifySettings) Set_peer(_peer []byte) {
	t._peer = _peer
}

func (t *TL_account_updateNotifySettings) Get_peer() []byte {
	return t._peer
}

func (t *TL_account_updateNotifySettings) Set_settings(_settings []byte) {
	t._settings = _settings
}

func (t *TL_account_updateNotifySettings) Get_settings() []byte {
	return t._settings
}

func New_TL_account_updateNotifySettings() *TL_account_updateNotifySettings {
	return &TL_account_updateNotifySettings{}
}

func (t *TL_account_updateNotifySettings) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_account_updateNotifySettings))
	ec.Bytes(t.Get_peer())
	ec.Bytes(t.Get_settings())

	return ec.GetBuffer()
}

func (t *TL_account_updateNotifySettings) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._peer = dc.Bytes(16)
	t._settings = dc.Bytes(16)

}

// account_getNotifySettings#12b3ad31
type TL_account_getNotifySettings struct {
	_peer []byte
}

func (t *TL_account_getNotifySettings) Set_peer(_peer []byte) {
	t._peer = _peer
}

func (t *TL_account_getNotifySettings) Get_peer() []byte {
	return t._peer
}

func New_TL_account_getNotifySettings() *TL_account_getNotifySettings {
	return &TL_account_getNotifySettings{}
}

func (t *TL_account_getNotifySettings) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_account_getNotifySettings))
	ec.Bytes(t.Get_peer())

	return ec.GetBuffer()
}

func (t *TL_account_getNotifySettings) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._peer = dc.Bytes(16)

}

// account_resetNotifySettings#db7e1747
type TL_account_resetNotifySettings struct {
}

func New_TL_account_resetNotifySettings() *TL_account_resetNotifySettings {
	return &TL_account_resetNotifySettings{}
}

func (t *TL_account_resetNotifySettings) Encode() []byte {
	return nil
}

func (t *TL_account_resetNotifySettings) Decode(b []byte) {

}

// account_updateProfile#78515775
type TL_account_updateProfile struct {
	_flags      []byte
	_first_name []byte
	_last_name  []byte
	_about      []byte
}

func (t *TL_account_updateProfile) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_account_updateProfile) Get_flags() []byte {
	return t._flags
}

func (t *TL_account_updateProfile) Set_first_name(_first_name []byte) {
	t._first_name = _first_name
}

func (t *TL_account_updateProfile) Get_first_name() []byte {
	return t._first_name
}

func (t *TL_account_updateProfile) Set_last_name(_last_name []byte) {
	t._last_name = _last_name
}

func (t *TL_account_updateProfile) Get_last_name() []byte {
	return t._last_name
}

func (t *TL_account_updateProfile) Set_about(_about []byte) {
	t._about = _about
}

func (t *TL_account_updateProfile) Get_about() []byte {
	return t._about
}

func New_TL_account_updateProfile() *TL_account_updateProfile {
	return &TL_account_updateProfile{}
}

func (t *TL_account_updateProfile) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_account_updateProfile))
	ec.Bytes(t.Get_first_name())
	ec.Bytes(t.Get_last_name())
	ec.Bytes(t.Get_about())

	return ec.GetBuffer()
}

func (t *TL_account_updateProfile) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._first_name = dc.Bytes(16)
	t._last_name = dc.Bytes(16)
	t._about = dc.Bytes(16)

}

// account_updateStatus#6628562c
type TL_account_updateStatus struct {
	_offline bool
}

func (t *TL_account_updateStatus) Set_offline(_offline bool) {
	t._offline = _offline
}

func (t *TL_account_updateStatus) Get_offline() bool {
	return t._offline
}

func New_TL_account_updateStatus() *TL_account_updateStatus {
	return &TL_account_updateStatus{}
}

func (t *TL_account_updateStatus) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_account_updateStatus))
	ec.Bool(t.Get_offline())

	return ec.GetBuffer()
}

func (t *TL_account_updateStatus) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._offline = dc.Bool()

}

// account_getWallPapers#c04cfac2
type TL_account_getWallPapers struct {
}

func New_TL_account_getWallPapers() *TL_account_getWallPapers {
	return &TL_account_getWallPapers{}
}

func (t *TL_account_getWallPapers) Encode() []byte {
	return nil
}

func (t *TL_account_getWallPapers) Decode(b []byte) {

}

// account_reportPeer#ae189d5f
type TL_account_reportPeer struct {
	_peer   []byte
	_reason []byte
}

func (t *TL_account_reportPeer) Set_peer(_peer []byte) {
	t._peer = _peer
}

func (t *TL_account_reportPeer) Get_peer() []byte {
	return t._peer
}

func (t *TL_account_reportPeer) Set_reason(_reason []byte) {
	t._reason = _reason
}

func (t *TL_account_reportPeer) Get_reason() []byte {
	return t._reason
}

func New_TL_account_reportPeer() *TL_account_reportPeer {
	return &TL_account_reportPeer{}
}

func (t *TL_account_reportPeer) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_account_reportPeer))
	ec.Bytes(t.Get_peer())
	ec.Bytes(t.Get_reason())

	return ec.GetBuffer()
}

func (t *TL_account_reportPeer) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._peer = dc.Bytes(16)
	t._reason = dc.Bytes(16)

}

// account_checkUsername#2714d86c
type TL_account_checkUsername struct {
	_username string
}

func (t *TL_account_checkUsername) Set_username(_username string) {
	t._username = _username
}

func (t *TL_account_checkUsername) Get_username() string {
	return t._username
}

func New_TL_account_checkUsername() *TL_account_checkUsername {
	return &TL_account_checkUsername{}
}

func (t *TL_account_checkUsername) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_account_checkUsername))
	ec.String(t.Get_username())

	return ec.GetBuffer()
}

func (t *TL_account_checkUsername) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._username = dc.String()

}

// account_updateUsername#3e0bdd7c
type TL_account_updateUsername struct {
	_username string
}

func (t *TL_account_updateUsername) Set_username(_username string) {
	t._username = _username
}

func (t *TL_account_updateUsername) Get_username() string {
	return t._username
}

func New_TL_account_updateUsername() *TL_account_updateUsername {
	return &TL_account_updateUsername{}
}

func (t *TL_account_updateUsername) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_account_updateUsername))
	ec.String(t.Get_username())

	return ec.GetBuffer()
}

func (t *TL_account_updateUsername) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._username = dc.String()

}

// account_getPrivacy#dadbc950
type TL_account_getPrivacy struct {
	_key []byte
}

func (t *TL_account_getPrivacy) Set_key(_key []byte) {
	t._key = _key
}

func (t *TL_account_getPrivacy) Get_key() []byte {
	return t._key
}

func New_TL_account_getPrivacy() *TL_account_getPrivacy {
	return &TL_account_getPrivacy{}
}

func (t *TL_account_getPrivacy) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_account_getPrivacy))
	ec.Bytes(t.Get_key())

	return ec.GetBuffer()
}

func (t *TL_account_getPrivacy) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._key = dc.Bytes(16)

}

// account_setPrivacy#c9f81ce8
type TL_account_setPrivacy struct {
	_key   []byte
	_rules []byte
}

func (t *TL_account_setPrivacy) Set_key(_key []byte) {
	t._key = _key
}

func (t *TL_account_setPrivacy) Get_key() []byte {
	return t._key
}

func (t *TL_account_setPrivacy) Set_rules(_rules []byte) {
	t._rules = _rules
}

func (t *TL_account_setPrivacy) Get_rules() []byte {
	return t._rules
}

func New_TL_account_setPrivacy() *TL_account_setPrivacy {
	return &TL_account_setPrivacy{}
}

func (t *TL_account_setPrivacy) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_account_setPrivacy))
	ec.Bytes(t.Get_key())
	ec.Bytes(t.Get_rules())

	return ec.GetBuffer()
}

func (t *TL_account_setPrivacy) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._key = dc.Bytes(16)
	t._rules = dc.Bytes(16)

}

// account_deleteAccount#418d4e0b
type TL_account_deleteAccount struct {
	_reason string
}

func (t *TL_account_deleteAccount) Set_reason(_reason string) {
	t._reason = _reason
}

func (t *TL_account_deleteAccount) Get_reason() string {
	return t._reason
}

func New_TL_account_deleteAccount() *TL_account_deleteAccount {
	return &TL_account_deleteAccount{}
}

func (t *TL_account_deleteAccount) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_account_deleteAccount))
	ec.String(t.Get_reason())

	return ec.GetBuffer()
}

func (t *TL_account_deleteAccount) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._reason = dc.String()

}

// account_getAccountTTL#08fc711d
type TL_account_getAccountTTL struct {
}

func New_TL_account_getAccountTTL() *TL_account_getAccountTTL {
	return &TL_account_getAccountTTL{}
}

func (t *TL_account_getAccountTTL) Encode() []byte {
	return nil
}

func (t *TL_account_getAccountTTL) Decode(b []byte) {

}

// account_setAccountTTL#2442485e
type TL_account_setAccountTTL struct {
	_ttl []byte
}

func (t *TL_account_setAccountTTL) Set_ttl(_ttl []byte) {
	t._ttl = _ttl
}

func (t *TL_account_setAccountTTL) Get_ttl() []byte {
	return t._ttl
}

func New_TL_account_setAccountTTL() *TL_account_setAccountTTL {
	return &TL_account_setAccountTTL{}
}

func (t *TL_account_setAccountTTL) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_account_setAccountTTL))
	ec.Bytes(t.Get_ttl())

	return ec.GetBuffer()
}

func (t *TL_account_setAccountTTL) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._ttl = dc.Bytes(16)

}

// account_sendChangePhoneCode#08e57deb
type TL_account_sendChangePhoneCode struct {
	_flags           []byte
	_allow_flashcall []byte
	_phone_number    string
	_current_number  []byte
}

func (t *TL_account_sendChangePhoneCode) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_account_sendChangePhoneCode) Get_flags() []byte {
	return t._flags
}

func (t *TL_account_sendChangePhoneCode) Set_allow_flashcall(_allow_flashcall []byte) {
	t._allow_flashcall = _allow_flashcall
}

func (t *TL_account_sendChangePhoneCode) Get_allow_flashcall() []byte {
	return t._allow_flashcall
}

func (t *TL_account_sendChangePhoneCode) Set_phone_number(_phone_number string) {
	t._phone_number = _phone_number
}

func (t *TL_account_sendChangePhoneCode) Get_phone_number() string {
	return t._phone_number
}

func (t *TL_account_sendChangePhoneCode) Set_current_number(_current_number []byte) {
	t._current_number = _current_number
}

func (t *TL_account_sendChangePhoneCode) Get_current_number() []byte {
	return t._current_number
}

func New_TL_account_sendChangePhoneCode() *TL_account_sendChangePhoneCode {
	return &TL_account_sendChangePhoneCode{}
}

func (t *TL_account_sendChangePhoneCode) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_account_sendChangePhoneCode))
	ec.Bytes(t.Get_allow_flashcall())
	ec.String(t.Get_phone_number())
	ec.Bytes(t.Get_current_number())

	return ec.GetBuffer()
}

func (t *TL_account_sendChangePhoneCode) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._allow_flashcall = dc.Bytes(16)
	t._phone_number = dc.String()
	t._current_number = dc.Bytes(16)

}

// account_changePhone#70c32edb
type TL_account_changePhone struct {
	_phone_number    string
	_phone_code_hash string
	_phone_code      string
}

func (t *TL_account_changePhone) Set_phone_number(_phone_number string) {
	t._phone_number = _phone_number
}

func (t *TL_account_changePhone) Get_phone_number() string {
	return t._phone_number
}

func (t *TL_account_changePhone) Set_phone_code_hash(_phone_code_hash string) {
	t._phone_code_hash = _phone_code_hash
}

func (t *TL_account_changePhone) Get_phone_code_hash() string {
	return t._phone_code_hash
}

func (t *TL_account_changePhone) Set_phone_code(_phone_code string) {
	t._phone_code = _phone_code
}

func (t *TL_account_changePhone) Get_phone_code() string {
	return t._phone_code
}

func New_TL_account_changePhone() *TL_account_changePhone {
	return &TL_account_changePhone{}
}

func (t *TL_account_changePhone) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_account_changePhone))
	ec.String(t.Get_phone_number())
	ec.String(t.Get_phone_code_hash())
	ec.String(t.Get_phone_code())

	return ec.GetBuffer()
}

func (t *TL_account_changePhone) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._phone_number = dc.String()
	t._phone_code_hash = dc.String()
	t._phone_code = dc.String()

}

// account_updateDeviceLocked#38df3532
type TL_account_updateDeviceLocked struct {
	_period int32
}

func (t *TL_account_updateDeviceLocked) Set_period(_period int32) {
	t._period = _period
}

func (t *TL_account_updateDeviceLocked) Get_period() int32 {
	return t._period
}

func New_TL_account_updateDeviceLocked() *TL_account_updateDeviceLocked {
	return &TL_account_updateDeviceLocked{}
}

func (t *TL_account_updateDeviceLocked) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_account_updateDeviceLocked))
	ec.Int(t.Get_period())

	return ec.GetBuffer()
}

func (t *TL_account_updateDeviceLocked) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._period = dc.Int()

}

// account_getAuthorizations#e320c158
type TL_account_getAuthorizations struct {
}

func New_TL_account_getAuthorizations() *TL_account_getAuthorizations {
	return &TL_account_getAuthorizations{}
}

func (t *TL_account_getAuthorizations) Encode() []byte {
	return nil
}

func (t *TL_account_getAuthorizations) Decode(b []byte) {

}

// account_resetAuthorization#df77f3bc
type TL_account_resetAuthorization struct {
	_hash int64
}

func (t *TL_account_resetAuthorization) Set_hash(_hash int64) {
	t._hash = _hash
}

func (t *TL_account_resetAuthorization) Get_hash() int64 {
	return t._hash
}

func New_TL_account_resetAuthorization() *TL_account_resetAuthorization {
	return &TL_account_resetAuthorization{}
}

func (t *TL_account_resetAuthorization) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_account_resetAuthorization))
	ec.Long(t.Get_hash())

	return ec.GetBuffer()
}

func (t *TL_account_resetAuthorization) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._hash = dc.Long()

}

// account_getPassword#548a30f5
type TL_account_getPassword struct {
}

func New_TL_account_getPassword() *TL_account_getPassword {
	return &TL_account_getPassword{}
}

func (t *TL_account_getPassword) Encode() []byte {
	return nil
}

func (t *TL_account_getPassword) Decode(b []byte) {

}

// account_getPasswordSettings#bc8d11bb
type TL_account_getPasswordSettings struct {
	_current_password_hash []byte
}

func (t *TL_account_getPasswordSettings) Set_current_password_hash(_current_password_hash []byte) {
	t._current_password_hash = _current_password_hash
}

func (t *TL_account_getPasswordSettings) Get_current_password_hash() []byte {
	return t._current_password_hash
}

func New_TL_account_getPasswordSettings() *TL_account_getPasswordSettings {
	return &TL_account_getPasswordSettings{}
}

func (t *TL_account_getPasswordSettings) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_account_getPasswordSettings))
	ec.Bytes(t.Get_current_password_hash())

	return ec.GetBuffer()
}

func (t *TL_account_getPasswordSettings) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._current_password_hash = dc.Bytes(16)

}

// account_updatePasswordSettings#fa7c4b86
type TL_account_updatePasswordSettings struct {
	_current_password_hash []byte
	_new_settings          []byte
}

func (t *TL_account_updatePasswordSettings) Set_current_password_hash(_current_password_hash []byte) {
	t._current_password_hash = _current_password_hash
}

func (t *TL_account_updatePasswordSettings) Get_current_password_hash() []byte {
	return t._current_password_hash
}

func (t *TL_account_updatePasswordSettings) Set_new_settings(_new_settings []byte) {
	t._new_settings = _new_settings
}

func (t *TL_account_updatePasswordSettings) Get_new_settings() []byte {
	return t._new_settings
}

func New_TL_account_updatePasswordSettings() *TL_account_updatePasswordSettings {
	return &TL_account_updatePasswordSettings{}
}

func (t *TL_account_updatePasswordSettings) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_account_updatePasswordSettings))
	ec.Bytes(t.Get_current_password_hash())
	ec.Bytes(t.Get_new_settings())

	return ec.GetBuffer()
}

func (t *TL_account_updatePasswordSettings) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._current_password_hash = dc.Bytes(16)
	t._new_settings = dc.Bytes(16)

}

// account_sendConfirmPhoneCode#1516d7bd
type TL_account_sendConfirmPhoneCode struct {
	_flags           []byte
	_allow_flashcall []byte
	_hash            string
	_current_number  []byte
}

func (t *TL_account_sendConfirmPhoneCode) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_account_sendConfirmPhoneCode) Get_flags() []byte {
	return t._flags
}

func (t *TL_account_sendConfirmPhoneCode) Set_allow_flashcall(_allow_flashcall []byte) {
	t._allow_flashcall = _allow_flashcall
}

func (t *TL_account_sendConfirmPhoneCode) Get_allow_flashcall() []byte {
	return t._allow_flashcall
}

func (t *TL_account_sendConfirmPhoneCode) Set_hash(_hash string) {
	t._hash = _hash
}

func (t *TL_account_sendConfirmPhoneCode) Get_hash() string {
	return t._hash
}

func (t *TL_account_sendConfirmPhoneCode) Set_current_number(_current_number []byte) {
	t._current_number = _current_number
}

func (t *TL_account_sendConfirmPhoneCode) Get_current_number() []byte {
	return t._current_number
}

func New_TL_account_sendConfirmPhoneCode() *TL_account_sendConfirmPhoneCode {
	return &TL_account_sendConfirmPhoneCode{}
}

func (t *TL_account_sendConfirmPhoneCode) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_account_sendConfirmPhoneCode))
	ec.Bytes(t.Get_allow_flashcall())
	ec.String(t.Get_hash())
	ec.Bytes(t.Get_current_number())

	return ec.GetBuffer()
}

func (t *TL_account_sendConfirmPhoneCode) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._allow_flashcall = dc.Bytes(16)
	t._hash = dc.String()
	t._current_number = dc.Bytes(16)

}

// account_confirmPhone#5f2178c3
type TL_account_confirmPhone struct {
	_phone_code_hash string
	_phone_code      string
}

func (t *TL_account_confirmPhone) Set_phone_code_hash(_phone_code_hash string) {
	t._phone_code_hash = _phone_code_hash
}

func (t *TL_account_confirmPhone) Get_phone_code_hash() string {
	return t._phone_code_hash
}

func (t *TL_account_confirmPhone) Set_phone_code(_phone_code string) {
	t._phone_code = _phone_code
}

func (t *TL_account_confirmPhone) Get_phone_code() string {
	return t._phone_code
}

func New_TL_account_confirmPhone() *TL_account_confirmPhone {
	return &TL_account_confirmPhone{}
}

func (t *TL_account_confirmPhone) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_account_confirmPhone))
	ec.String(t.Get_phone_code_hash())
	ec.String(t.Get_phone_code())

	return ec.GetBuffer()
}

func (t *TL_account_confirmPhone) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._phone_code_hash = dc.String()
	t._phone_code = dc.String()

}

// account_getTmpPassword#4a82327e
type TL_account_getTmpPassword struct {
	_password_hash []byte
	_period        int32
}

func (t *TL_account_getTmpPassword) Set_password_hash(_password_hash []byte) {
	t._password_hash = _password_hash
}

func (t *TL_account_getTmpPassword) Get_password_hash() []byte {
	return t._password_hash
}

func (t *TL_account_getTmpPassword) Set_period(_period int32) {
	t._period = _period
}

func (t *TL_account_getTmpPassword) Get_period() int32 {
	return t._period
}

func New_TL_account_getTmpPassword() *TL_account_getTmpPassword {
	return &TL_account_getTmpPassword{}
}

func (t *TL_account_getTmpPassword) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_account_getTmpPassword))
	ec.Bytes(t.Get_password_hash())
	ec.Int(t.Get_period())

	return ec.GetBuffer()
}

func (t *TL_account_getTmpPassword) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._password_hash = dc.Bytes(16)
	t._period = dc.Int()

}

// users_getUsers#0d91a548
type TL_users_getUsers struct {
	_id []byte
}

func (t *TL_users_getUsers) Set_id(_id []byte) {
	t._id = _id
}

func (t *TL_users_getUsers) Get_id() []byte {
	return t._id
}

func New_TL_users_getUsers() *TL_users_getUsers {
	return &TL_users_getUsers{}
}

func (t *TL_users_getUsers) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_users_getUsers))
	ec.Bytes(t.Get_id())

	return ec.GetBuffer()
}

func (t *TL_users_getUsers) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Bytes(16)

}

// users_getFullUser#ca30a5b1
type TL_users_getFullUser struct {
	_id []byte
}

func (t *TL_users_getFullUser) Set_id(_id []byte) {
	t._id = _id
}

func (t *TL_users_getFullUser) Get_id() []byte {
	return t._id
}

func New_TL_users_getFullUser() *TL_users_getFullUser {
	return &TL_users_getFullUser{}
}

func (t *TL_users_getFullUser) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_users_getFullUser))
	ec.Bytes(t.Get_id())

	return ec.GetBuffer()
}

func (t *TL_users_getFullUser) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Bytes(16)

}

// contacts_getStatuses#c4a353ee
type TL_contacts_getStatuses struct {
}

func New_TL_contacts_getStatuses() *TL_contacts_getStatuses {
	return &TL_contacts_getStatuses{}
}

func (t *TL_contacts_getStatuses) Encode() []byte {
	return nil
}

func (t *TL_contacts_getStatuses) Decode(b []byte) {

}

// contacts_getContacts#c023849f
type TL_contacts_getContacts struct {
	_hash int32
}

func (t *TL_contacts_getContacts) Set_hash(_hash int32) {
	t._hash = _hash
}

func (t *TL_contacts_getContacts) Get_hash() int32 {
	return t._hash
}

func New_TL_contacts_getContacts() *TL_contacts_getContacts {
	return &TL_contacts_getContacts{}
}

func (t *TL_contacts_getContacts) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_contacts_getContacts))
	ec.Int(t.Get_hash())

	return ec.GetBuffer()
}

func (t *TL_contacts_getContacts) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._hash = dc.Int()

}

// contacts_importContacts#2c800be5
type TL_contacts_importContacts struct {
	_contacts []byte
}

func (t *TL_contacts_importContacts) Set_contacts(_contacts []byte) {
	t._contacts = _contacts
}

func (t *TL_contacts_importContacts) Get_contacts() []byte {
	return t._contacts
}

func New_TL_contacts_importContacts() *TL_contacts_importContacts {
	return &TL_contacts_importContacts{}
}

func (t *TL_contacts_importContacts) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_contacts_importContacts))
	ec.Bytes(t.Get_contacts())

	return ec.GetBuffer()
}

func (t *TL_contacts_importContacts) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._contacts = dc.Bytes(16)

}

// contacts_deleteContact#8e953744
type TL_contacts_deleteContact struct {
	_id []byte
}

func (t *TL_contacts_deleteContact) Set_id(_id []byte) {
	t._id = _id
}

func (t *TL_contacts_deleteContact) Get_id() []byte {
	return t._id
}

func New_TL_contacts_deleteContact() *TL_contacts_deleteContact {
	return &TL_contacts_deleteContact{}
}

func (t *TL_contacts_deleteContact) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_contacts_deleteContact))
	ec.Bytes(t.Get_id())

	return ec.GetBuffer()
}

func (t *TL_contacts_deleteContact) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Bytes(16)

}

// contacts_deleteContacts#59ab389e
type TL_contacts_deleteContacts struct {
	_id []byte
}

func (t *TL_contacts_deleteContacts) Set_id(_id []byte) {
	t._id = _id
}

func (t *TL_contacts_deleteContacts) Get_id() []byte {
	return t._id
}

func New_TL_contacts_deleteContacts() *TL_contacts_deleteContacts {
	return &TL_contacts_deleteContacts{}
}

func (t *TL_contacts_deleteContacts) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_contacts_deleteContacts))
	ec.Bytes(t.Get_id())

	return ec.GetBuffer()
}

func (t *TL_contacts_deleteContacts) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Bytes(16)

}

// contacts_block#332b49fc
type TL_contacts_block struct {
	_id []byte
}

func (t *TL_contacts_block) Set_id(_id []byte) {
	t._id = _id
}

func (t *TL_contacts_block) Get_id() []byte {
	return t._id
}

func New_TL_contacts_block() *TL_contacts_block {
	return &TL_contacts_block{}
}

func (t *TL_contacts_block) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_contacts_block))
	ec.Bytes(t.Get_id())

	return ec.GetBuffer()
}

func (t *TL_contacts_block) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Bytes(16)

}

// contacts_unblock#e54100bd
type TL_contacts_unblock struct {
	_id []byte
}

func (t *TL_contacts_unblock) Set_id(_id []byte) {
	t._id = _id
}

func (t *TL_contacts_unblock) Get_id() []byte {
	return t._id
}

func New_TL_contacts_unblock() *TL_contacts_unblock {
	return &TL_contacts_unblock{}
}

func (t *TL_contacts_unblock) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_contacts_unblock))
	ec.Bytes(t.Get_id())

	return ec.GetBuffer()
}

func (t *TL_contacts_unblock) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Bytes(16)

}

// contacts_getBlocked#f57c350f
type TL_contacts_getBlocked struct {
	_offset int32
	_limit  int32
}

func (t *TL_contacts_getBlocked) Set_offset(_offset int32) {
	t._offset = _offset
}

func (t *TL_contacts_getBlocked) Get_offset() int32 {
	return t._offset
}

func (t *TL_contacts_getBlocked) Set_limit(_limit int32) {
	t._limit = _limit
}

func (t *TL_contacts_getBlocked) Get_limit() int32 {
	return t._limit
}

func New_TL_contacts_getBlocked() *TL_contacts_getBlocked {
	return &TL_contacts_getBlocked{}
}

func (t *TL_contacts_getBlocked) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_contacts_getBlocked))
	ec.Int(t.Get_offset())
	ec.Int(t.Get_limit())

	return ec.GetBuffer()
}

func (t *TL_contacts_getBlocked) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._offset = dc.Int()
	t._limit = dc.Int()

}

// contacts_exportCard#84e53737
type TL_contacts_exportCard struct {
}

func New_TL_contacts_exportCard() *TL_contacts_exportCard {
	return &TL_contacts_exportCard{}
}

func (t *TL_contacts_exportCard) Encode() []byte {
	return nil
}

func (t *TL_contacts_exportCard) Decode(b []byte) {

}

// contacts_importCard#4fe196fe
type TL_contacts_importCard struct {
	_export_card []byte
}

func (t *TL_contacts_importCard) Set_export_card(_export_card []byte) {
	t._export_card = _export_card
}

func (t *TL_contacts_importCard) Get_export_card() []byte {
	return t._export_card
}

func New_TL_contacts_importCard() *TL_contacts_importCard {
	return &TL_contacts_importCard{}
}

func (t *TL_contacts_importCard) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_contacts_importCard))
	ec.Bytes(t.Get_export_card())

	return ec.GetBuffer()
}

func (t *TL_contacts_importCard) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._export_card = dc.Bytes(16)

}

// contacts_search#11f812d8
type TL_contacts_search struct {
	_q     string
	_limit int32
}

func (t *TL_contacts_search) Set_q(_q string) {
	t._q = _q
}

func (t *TL_contacts_search) Get_q() string {
	return t._q
}

func (t *TL_contacts_search) Set_limit(_limit int32) {
	t._limit = _limit
}

func (t *TL_contacts_search) Get_limit() int32 {
	return t._limit
}

func New_TL_contacts_search() *TL_contacts_search {
	return &TL_contacts_search{}
}

func (t *TL_contacts_search) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_contacts_search))
	ec.String(t.Get_q())
	ec.Int(t.Get_limit())

	return ec.GetBuffer()
}

func (t *TL_contacts_search) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._q = dc.String()
	t._limit = dc.Int()

}

// contacts_resolveUsername#f93ccba3
type TL_contacts_resolveUsername struct {
	_username string
}

func (t *TL_contacts_resolveUsername) Set_username(_username string) {
	t._username = _username
}

func (t *TL_contacts_resolveUsername) Get_username() string {
	return t._username
}

func New_TL_contacts_resolveUsername() *TL_contacts_resolveUsername {
	return &TL_contacts_resolveUsername{}
}

func (t *TL_contacts_resolveUsername) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_contacts_resolveUsername))
	ec.String(t.Get_username())

	return ec.GetBuffer()
}

func (t *TL_contacts_resolveUsername) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._username = dc.String()

}

// contacts_getTopPeers#d4982db5
type TL_contacts_getTopPeers struct {
	_flags          []byte
	_correspondents []byte
	_bots_pm        []byte
	_bots_inline    []byte
	_phone_calls    []byte
	_groups         []byte
	_channels       []byte
	_offset         int32
	_limit          int32
	_hash           int32
}

func (t *TL_contacts_getTopPeers) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_contacts_getTopPeers) Get_flags() []byte {
	return t._flags
}

func (t *TL_contacts_getTopPeers) Set_correspondents(_correspondents []byte) {
	t._correspondents = _correspondents
}

func (t *TL_contacts_getTopPeers) Get_correspondents() []byte {
	return t._correspondents
}

func (t *TL_contacts_getTopPeers) Set_bots_pm(_bots_pm []byte) {
	t._bots_pm = _bots_pm
}

func (t *TL_contacts_getTopPeers) Get_bots_pm() []byte {
	return t._bots_pm
}

func (t *TL_contacts_getTopPeers) Set_bots_inline(_bots_inline []byte) {
	t._bots_inline = _bots_inline
}

func (t *TL_contacts_getTopPeers) Get_bots_inline() []byte {
	return t._bots_inline
}

func (t *TL_contacts_getTopPeers) Set_phone_calls(_phone_calls []byte) {
	t._phone_calls = _phone_calls
}

func (t *TL_contacts_getTopPeers) Get_phone_calls() []byte {
	return t._phone_calls
}

func (t *TL_contacts_getTopPeers) Set_groups(_groups []byte) {
	t._groups = _groups
}

func (t *TL_contacts_getTopPeers) Get_groups() []byte {
	return t._groups
}

func (t *TL_contacts_getTopPeers) Set_channels(_channels []byte) {
	t._channels = _channels
}

func (t *TL_contacts_getTopPeers) Get_channels() []byte {
	return t._channels
}

func (t *TL_contacts_getTopPeers) Set_offset(_offset int32) {
	t._offset = _offset
}

func (t *TL_contacts_getTopPeers) Get_offset() int32 {
	return t._offset
}

func (t *TL_contacts_getTopPeers) Set_limit(_limit int32) {
	t._limit = _limit
}

func (t *TL_contacts_getTopPeers) Get_limit() int32 {
	return t._limit
}

func (t *TL_contacts_getTopPeers) Set_hash(_hash int32) {
	t._hash = _hash
}

func (t *TL_contacts_getTopPeers) Get_hash() int32 {
	return t._hash
}

func New_TL_contacts_getTopPeers() *TL_contacts_getTopPeers {
	return &TL_contacts_getTopPeers{}
}

func (t *TL_contacts_getTopPeers) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_contacts_getTopPeers))
	ec.Bytes(t.Get_correspondents())
	ec.Bytes(t.Get_bots_pm())
	ec.Bytes(t.Get_bots_inline())
	ec.Bytes(t.Get_phone_calls())
	ec.Bytes(t.Get_groups())
	ec.Bytes(t.Get_channels())
	ec.Int(t.Get_offset())
	ec.Int(t.Get_limit())
	ec.Int(t.Get_hash())

	return ec.GetBuffer()
}

func (t *TL_contacts_getTopPeers) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._correspondents = dc.Bytes(16)
	t._bots_pm = dc.Bytes(16)
	t._bots_inline = dc.Bytes(16)
	t._phone_calls = dc.Bytes(16)
	t._groups = dc.Bytes(16)
	t._channels = dc.Bytes(16)
	t._offset = dc.Int()
	t._limit = dc.Int()
	t._hash = dc.Int()

}

// contacts_resetTopPeerRating#1ae373ac
type TL_contacts_resetTopPeerRating struct {
	_category []byte
	_peer     []byte
}

func (t *TL_contacts_resetTopPeerRating) Set_category(_category []byte) {
	t._category = _category
}

func (t *TL_contacts_resetTopPeerRating) Get_category() []byte {
	return t._category
}

func (t *TL_contacts_resetTopPeerRating) Set_peer(_peer []byte) {
	t._peer = _peer
}

func (t *TL_contacts_resetTopPeerRating) Get_peer() []byte {
	return t._peer
}

func New_TL_contacts_resetTopPeerRating() *TL_contacts_resetTopPeerRating {
	return &TL_contacts_resetTopPeerRating{}
}

func (t *TL_contacts_resetTopPeerRating) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_contacts_resetTopPeerRating))
	ec.Bytes(t.Get_category())
	ec.Bytes(t.Get_peer())

	return ec.GetBuffer()
}

func (t *TL_contacts_resetTopPeerRating) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._category = dc.Bytes(16)
	t._peer = dc.Bytes(16)

}

// contacts_resetSaved#879537f1
type TL_contacts_resetSaved struct {
}

func New_TL_contacts_resetSaved() *TL_contacts_resetSaved {
	return &TL_contacts_resetSaved{}
}

func (t *TL_contacts_resetSaved) Encode() []byte {
	return nil
}

func (t *TL_contacts_resetSaved) Decode(b []byte) {

}

// messages_getMessages#4222fa74
type TL_messages_getMessages struct {
	_id []byte
}

func (t *TL_messages_getMessages) Set_id(_id []byte) {
	t._id = _id
}

func (t *TL_messages_getMessages) Get_id() []byte {
	return t._id
}

func New_TL_messages_getMessages() *TL_messages_getMessages {
	return &TL_messages_getMessages{}
}

func (t *TL_messages_getMessages) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_getMessages))
	ec.Bytes(t.Get_id())

	return ec.GetBuffer()
}

func (t *TL_messages_getMessages) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Bytes(16)

}

// messages_getDialogs#191ba9c5
type TL_messages_getDialogs struct {
	_flags          []byte
	_exclude_pinned []byte
	_offset_date    int32
	_offset_id      int32
	_offset_peer    []byte
	_limit          int32
}

func (t *TL_messages_getDialogs) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_messages_getDialogs) Get_flags() []byte {
	return t._flags
}

func (t *TL_messages_getDialogs) Set_exclude_pinned(_exclude_pinned []byte) {
	t._exclude_pinned = _exclude_pinned
}

func (t *TL_messages_getDialogs) Get_exclude_pinned() []byte {
	return t._exclude_pinned
}

func (t *TL_messages_getDialogs) Set_offset_date(_offset_date int32) {
	t._offset_date = _offset_date
}

func (t *TL_messages_getDialogs) Get_offset_date() int32 {
	return t._offset_date
}

func (t *TL_messages_getDialogs) Set_offset_id(_offset_id int32) {
	t._offset_id = _offset_id
}

func (t *TL_messages_getDialogs) Get_offset_id() int32 {
	return t._offset_id
}

func (t *TL_messages_getDialogs) Set_offset_peer(_offset_peer []byte) {
	t._offset_peer = _offset_peer
}

func (t *TL_messages_getDialogs) Get_offset_peer() []byte {
	return t._offset_peer
}

func (t *TL_messages_getDialogs) Set_limit(_limit int32) {
	t._limit = _limit
}

func (t *TL_messages_getDialogs) Get_limit() int32 {
	return t._limit
}

func New_TL_messages_getDialogs() *TL_messages_getDialogs {
	return &TL_messages_getDialogs{}
}

func (t *TL_messages_getDialogs) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_getDialogs))
	ec.Bytes(t.Get_exclude_pinned())
	ec.Int(t.Get_offset_date())
	ec.Int(t.Get_offset_id())
	ec.Bytes(t.Get_offset_peer())
	ec.Int(t.Get_limit())

	return ec.GetBuffer()
}

func (t *TL_messages_getDialogs) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._exclude_pinned = dc.Bytes(16)
	t._offset_date = dc.Int()
	t._offset_id = dc.Int()
	t._offset_peer = dc.Bytes(16)
	t._limit = dc.Int()

}

// messages_getHistory#dcbb8260
type TL_messages_getHistory struct {
	_peer        []byte
	_offset_id   int32
	_offset_date int32
	_add_offset  int32
	_limit       int32
	_max_id      int32
	_min_id      int32
	_hash        int32
}

func (t *TL_messages_getHistory) Set_peer(_peer []byte) {
	t._peer = _peer
}

func (t *TL_messages_getHistory) Get_peer() []byte {
	return t._peer
}

func (t *TL_messages_getHistory) Set_offset_id(_offset_id int32) {
	t._offset_id = _offset_id
}

func (t *TL_messages_getHistory) Get_offset_id() int32 {
	return t._offset_id
}

func (t *TL_messages_getHistory) Set_offset_date(_offset_date int32) {
	t._offset_date = _offset_date
}

func (t *TL_messages_getHistory) Get_offset_date() int32 {
	return t._offset_date
}

func (t *TL_messages_getHistory) Set_add_offset(_add_offset int32) {
	t._add_offset = _add_offset
}

func (t *TL_messages_getHistory) Get_add_offset() int32 {
	return t._add_offset
}

func (t *TL_messages_getHistory) Set_limit(_limit int32) {
	t._limit = _limit
}

func (t *TL_messages_getHistory) Get_limit() int32 {
	return t._limit
}

func (t *TL_messages_getHistory) Set_max_id(_max_id int32) {
	t._max_id = _max_id
}

func (t *TL_messages_getHistory) Get_max_id() int32 {
	return t._max_id
}

func (t *TL_messages_getHistory) Set_min_id(_min_id int32) {
	t._min_id = _min_id
}

func (t *TL_messages_getHistory) Get_min_id() int32 {
	return t._min_id
}

func (t *TL_messages_getHistory) Set_hash(_hash int32) {
	t._hash = _hash
}

func (t *TL_messages_getHistory) Get_hash() int32 {
	return t._hash
}

func New_TL_messages_getHistory() *TL_messages_getHistory {
	return &TL_messages_getHistory{}
}

func (t *TL_messages_getHistory) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_getHistory))
	ec.Bytes(t.Get_peer())
	ec.Int(t.Get_offset_id())
	ec.Int(t.Get_offset_date())
	ec.Int(t.Get_add_offset())
	ec.Int(t.Get_limit())
	ec.Int(t.Get_max_id())
	ec.Int(t.Get_min_id())
	ec.Int(t.Get_hash())

	return ec.GetBuffer()
}

func (t *TL_messages_getHistory) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._peer = dc.Bytes(16)
	t._offset_id = dc.Int()
	t._offset_date = dc.Int()
	t._add_offset = dc.Int()
	t._limit = dc.Int()
	t._max_id = dc.Int()
	t._min_id = dc.Int()
	t._hash = dc.Int()

}

// messages_search#039e9ea0
type TL_messages_search struct {
	_flags      []byte
	_peer       []byte
	_q          string
	_from_id    []byte
	_filter     []byte
	_min_date   int32
	_max_date   int32
	_offset_id  int32
	_add_offset int32
	_limit      int32
	_max_id     int32
	_min_id     int32
}

func (t *TL_messages_search) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_messages_search) Get_flags() []byte {
	return t._flags
}

func (t *TL_messages_search) Set_peer(_peer []byte) {
	t._peer = _peer
}

func (t *TL_messages_search) Get_peer() []byte {
	return t._peer
}

func (t *TL_messages_search) Set_q(_q string) {
	t._q = _q
}

func (t *TL_messages_search) Get_q() string {
	return t._q
}

func (t *TL_messages_search) Set_from_id(_from_id []byte) {
	t._from_id = _from_id
}

func (t *TL_messages_search) Get_from_id() []byte {
	return t._from_id
}

func (t *TL_messages_search) Set_filter(_filter []byte) {
	t._filter = _filter
}

func (t *TL_messages_search) Get_filter() []byte {
	return t._filter
}

func (t *TL_messages_search) Set_min_date(_min_date int32) {
	t._min_date = _min_date
}

func (t *TL_messages_search) Get_min_date() int32 {
	return t._min_date
}

func (t *TL_messages_search) Set_max_date(_max_date int32) {
	t._max_date = _max_date
}

func (t *TL_messages_search) Get_max_date() int32 {
	return t._max_date
}

func (t *TL_messages_search) Set_offset_id(_offset_id int32) {
	t._offset_id = _offset_id
}

func (t *TL_messages_search) Get_offset_id() int32 {
	return t._offset_id
}

func (t *TL_messages_search) Set_add_offset(_add_offset int32) {
	t._add_offset = _add_offset
}

func (t *TL_messages_search) Get_add_offset() int32 {
	return t._add_offset
}

func (t *TL_messages_search) Set_limit(_limit int32) {
	t._limit = _limit
}

func (t *TL_messages_search) Get_limit() int32 {
	return t._limit
}

func (t *TL_messages_search) Set_max_id(_max_id int32) {
	t._max_id = _max_id
}

func (t *TL_messages_search) Get_max_id() int32 {
	return t._max_id
}

func (t *TL_messages_search) Set_min_id(_min_id int32) {
	t._min_id = _min_id
}

func (t *TL_messages_search) Get_min_id() int32 {
	return t._min_id
}

func New_TL_messages_search() *TL_messages_search {
	return &TL_messages_search{}
}

func (t *TL_messages_search) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_search))
	ec.Bytes(t.Get_peer())
	ec.String(t.Get_q())
	ec.Bytes(t.Get_from_id())
	ec.Bytes(t.Get_filter())
	ec.Int(t.Get_min_date())
	ec.Int(t.Get_max_date())
	ec.Int(t.Get_offset_id())
	ec.Int(t.Get_add_offset())
	ec.Int(t.Get_limit())
	ec.Int(t.Get_max_id())
	ec.Int(t.Get_min_id())

	return ec.GetBuffer()
}

func (t *TL_messages_search) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._peer = dc.Bytes(16)
	t._q = dc.String()
	t._from_id = dc.Bytes(16)
	t._filter = dc.Bytes(16)
	t._min_date = dc.Int()
	t._max_date = dc.Int()
	t._offset_id = dc.Int()
	t._add_offset = dc.Int()
	t._limit = dc.Int()
	t._max_id = dc.Int()
	t._min_id = dc.Int()

}

// messages_readHistory#0e306d3a
type TL_messages_readHistory struct {
	_peer   []byte
	_max_id int32
}

func (t *TL_messages_readHistory) Set_peer(_peer []byte) {
	t._peer = _peer
}

func (t *TL_messages_readHistory) Get_peer() []byte {
	return t._peer
}

func (t *TL_messages_readHistory) Set_max_id(_max_id int32) {
	t._max_id = _max_id
}

func (t *TL_messages_readHistory) Get_max_id() int32 {
	return t._max_id
}

func New_TL_messages_readHistory() *TL_messages_readHistory {
	return &TL_messages_readHistory{}
}

func (t *TL_messages_readHistory) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_readHistory))
	ec.Bytes(t.Get_peer())
	ec.Int(t.Get_max_id())

	return ec.GetBuffer()
}

func (t *TL_messages_readHistory) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._peer = dc.Bytes(16)
	t._max_id = dc.Int()

}

// messages_deleteHistory#1c015b09
type TL_messages_deleteHistory struct {
	_flags      []byte
	_just_clear []byte
	_peer       []byte
	_max_id     int32
}

func (t *TL_messages_deleteHistory) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_messages_deleteHistory) Get_flags() []byte {
	return t._flags
}

func (t *TL_messages_deleteHistory) Set_just_clear(_just_clear []byte) {
	t._just_clear = _just_clear
}

func (t *TL_messages_deleteHistory) Get_just_clear() []byte {
	return t._just_clear
}

func (t *TL_messages_deleteHistory) Set_peer(_peer []byte) {
	t._peer = _peer
}

func (t *TL_messages_deleteHistory) Get_peer() []byte {
	return t._peer
}

func (t *TL_messages_deleteHistory) Set_max_id(_max_id int32) {
	t._max_id = _max_id
}

func (t *TL_messages_deleteHistory) Get_max_id() int32 {
	return t._max_id
}

func New_TL_messages_deleteHistory() *TL_messages_deleteHistory {
	return &TL_messages_deleteHistory{}
}

func (t *TL_messages_deleteHistory) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_deleteHistory))
	ec.Bytes(t.Get_just_clear())
	ec.Bytes(t.Get_peer())
	ec.Int(t.Get_max_id())

	return ec.GetBuffer()
}

func (t *TL_messages_deleteHistory) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._just_clear = dc.Bytes(16)
	t._peer = dc.Bytes(16)
	t._max_id = dc.Int()

}

// messages_deleteMessages#e58e95d2
type TL_messages_deleteMessages struct {
	_flags  []byte
	_revoke []byte
	_id     []byte
}

func (t *TL_messages_deleteMessages) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_messages_deleteMessages) Get_flags() []byte {
	return t._flags
}

func (t *TL_messages_deleteMessages) Set_revoke(_revoke []byte) {
	t._revoke = _revoke
}

func (t *TL_messages_deleteMessages) Get_revoke() []byte {
	return t._revoke
}

func (t *TL_messages_deleteMessages) Set_id(_id []byte) {
	t._id = _id
}

func (t *TL_messages_deleteMessages) Get_id() []byte {
	return t._id
}

func New_TL_messages_deleteMessages() *TL_messages_deleteMessages {
	return &TL_messages_deleteMessages{}
}

func (t *TL_messages_deleteMessages) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_deleteMessages))
	ec.Bytes(t.Get_revoke())
	ec.Bytes(t.Get_id())

	return ec.GetBuffer()
}

func (t *TL_messages_deleteMessages) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._revoke = dc.Bytes(16)
	t._id = dc.Bytes(16)

}

// messages_receivedMessages#05a954c0
type TL_messages_receivedMessages struct {
	_max_id int32
}

func (t *TL_messages_receivedMessages) Set_max_id(_max_id int32) {
	t._max_id = _max_id
}

func (t *TL_messages_receivedMessages) Get_max_id() int32 {
	return t._max_id
}

func New_TL_messages_receivedMessages() *TL_messages_receivedMessages {
	return &TL_messages_receivedMessages{}
}

func (t *TL_messages_receivedMessages) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_receivedMessages))
	ec.Int(t.Get_max_id())

	return ec.GetBuffer()
}

func (t *TL_messages_receivedMessages) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._max_id = dc.Int()

}

// messages_setTyping#a3825e50
type TL_messages_setTyping struct {
	_peer   []byte
	_action []byte
}

func (t *TL_messages_setTyping) Set_peer(_peer []byte) {
	t._peer = _peer
}

func (t *TL_messages_setTyping) Get_peer() []byte {
	return t._peer
}

func (t *TL_messages_setTyping) Set_action(_action []byte) {
	t._action = _action
}

func (t *TL_messages_setTyping) Get_action() []byte {
	return t._action
}

func New_TL_messages_setTyping() *TL_messages_setTyping {
	return &TL_messages_setTyping{}
}

func (t *TL_messages_setTyping) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_setTyping))
	ec.Bytes(t.Get_peer())
	ec.Bytes(t.Get_action())

	return ec.GetBuffer()
}

func (t *TL_messages_setTyping) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._peer = dc.Bytes(16)
	t._action = dc.Bytes(16)

}

// messages_sendMessage#fa88427a
type TL_messages_sendMessage struct {
	_flags           []byte
	_no_webpage      []byte
	_silent          []byte
	_background      []byte
	_clear_draft     []byte
	_peer            []byte
	_reply_to_msg_id []byte
	_message         string
	_random_id       int64
	_reply_markup    []byte
	_entities        []byte
}

func (t *TL_messages_sendMessage) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_messages_sendMessage) Get_flags() []byte {
	return t._flags
}

func (t *TL_messages_sendMessage) Set_no_webpage(_no_webpage []byte) {
	t._no_webpage = _no_webpage
}

func (t *TL_messages_sendMessage) Get_no_webpage() []byte {
	return t._no_webpage
}

func (t *TL_messages_sendMessage) Set_silent(_silent []byte) {
	t._silent = _silent
}

func (t *TL_messages_sendMessage) Get_silent() []byte {
	return t._silent
}

func (t *TL_messages_sendMessage) Set_background(_background []byte) {
	t._background = _background
}

func (t *TL_messages_sendMessage) Get_background() []byte {
	return t._background
}

func (t *TL_messages_sendMessage) Set_clear_draft(_clear_draft []byte) {
	t._clear_draft = _clear_draft
}

func (t *TL_messages_sendMessage) Get_clear_draft() []byte {
	return t._clear_draft
}

func (t *TL_messages_sendMessage) Set_peer(_peer []byte) {
	t._peer = _peer
}

func (t *TL_messages_sendMessage) Get_peer() []byte {
	return t._peer
}

func (t *TL_messages_sendMessage) Set_reply_to_msg_id(_reply_to_msg_id []byte) {
	t._reply_to_msg_id = _reply_to_msg_id
}

func (t *TL_messages_sendMessage) Get_reply_to_msg_id() []byte {
	return t._reply_to_msg_id
}

func (t *TL_messages_sendMessage) Set_message(_message string) {
	t._message = _message
}

func (t *TL_messages_sendMessage) Get_message() string {
	return t._message
}

func (t *TL_messages_sendMessage) Set_random_id(_random_id int64) {
	t._random_id = _random_id
}

func (t *TL_messages_sendMessage) Get_random_id() int64 {
	return t._random_id
}

func (t *TL_messages_sendMessage) Set_reply_markup(_reply_markup []byte) {
	t._reply_markup = _reply_markup
}

func (t *TL_messages_sendMessage) Get_reply_markup() []byte {
	return t._reply_markup
}

func (t *TL_messages_sendMessage) Set_entities(_entities []byte) {
	t._entities = _entities
}

func (t *TL_messages_sendMessage) Get_entities() []byte {
	return t._entities
}

func New_TL_messages_sendMessage() *TL_messages_sendMessage {
	return &TL_messages_sendMessage{}
}

func (t *TL_messages_sendMessage) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_sendMessage))
	ec.Bytes(t.Get_no_webpage())
	ec.Bytes(t.Get_silent())
	ec.Bytes(t.Get_background())
	ec.Bytes(t.Get_clear_draft())
	ec.Bytes(t.Get_peer())
	ec.Bytes(t.Get_reply_to_msg_id())
	ec.String(t.Get_message())
	ec.Long(t.Get_random_id())
	ec.Bytes(t.Get_reply_markup())
	ec.Bytes(t.Get_entities())

	return ec.GetBuffer()
}

func (t *TL_messages_sendMessage) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._no_webpage = dc.Bytes(16)
	t._silent = dc.Bytes(16)
	t._background = dc.Bytes(16)
	t._clear_draft = dc.Bytes(16)
	t._peer = dc.Bytes(16)
	t._reply_to_msg_id = dc.Bytes(16)
	t._message = dc.String()
	t._random_id = dc.Long()
	t._reply_markup = dc.Bytes(16)
	t._entities = dc.Bytes(16)

}

// messages_sendMedia#c8f16791
type TL_messages_sendMedia struct {
	_flags           []byte
	_silent          []byte
	_background      []byte
	_clear_draft     []byte
	_peer            []byte
	_reply_to_msg_id []byte
	_media           []byte
	_random_id       int64
	_reply_markup    []byte
}

func (t *TL_messages_sendMedia) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_messages_sendMedia) Get_flags() []byte {
	return t._flags
}

func (t *TL_messages_sendMedia) Set_silent(_silent []byte) {
	t._silent = _silent
}

func (t *TL_messages_sendMedia) Get_silent() []byte {
	return t._silent
}

func (t *TL_messages_sendMedia) Set_background(_background []byte) {
	t._background = _background
}

func (t *TL_messages_sendMedia) Get_background() []byte {
	return t._background
}

func (t *TL_messages_sendMedia) Set_clear_draft(_clear_draft []byte) {
	t._clear_draft = _clear_draft
}

func (t *TL_messages_sendMedia) Get_clear_draft() []byte {
	return t._clear_draft
}

func (t *TL_messages_sendMedia) Set_peer(_peer []byte) {
	t._peer = _peer
}

func (t *TL_messages_sendMedia) Get_peer() []byte {
	return t._peer
}

func (t *TL_messages_sendMedia) Set_reply_to_msg_id(_reply_to_msg_id []byte) {
	t._reply_to_msg_id = _reply_to_msg_id
}

func (t *TL_messages_sendMedia) Get_reply_to_msg_id() []byte {
	return t._reply_to_msg_id
}

func (t *TL_messages_sendMedia) Set_media(_media []byte) {
	t._media = _media
}

func (t *TL_messages_sendMedia) Get_media() []byte {
	return t._media
}

func (t *TL_messages_sendMedia) Set_random_id(_random_id int64) {
	t._random_id = _random_id
}

func (t *TL_messages_sendMedia) Get_random_id() int64 {
	return t._random_id
}

func (t *TL_messages_sendMedia) Set_reply_markup(_reply_markup []byte) {
	t._reply_markup = _reply_markup
}

func (t *TL_messages_sendMedia) Get_reply_markup() []byte {
	return t._reply_markup
}

func New_TL_messages_sendMedia() *TL_messages_sendMedia {
	return &TL_messages_sendMedia{}
}

func (t *TL_messages_sendMedia) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_sendMedia))
	ec.Bytes(t.Get_silent())
	ec.Bytes(t.Get_background())
	ec.Bytes(t.Get_clear_draft())
	ec.Bytes(t.Get_peer())
	ec.Bytes(t.Get_reply_to_msg_id())
	ec.Bytes(t.Get_media())
	ec.Long(t.Get_random_id())
	ec.Bytes(t.Get_reply_markup())

	return ec.GetBuffer()
}

func (t *TL_messages_sendMedia) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._silent = dc.Bytes(16)
	t._background = dc.Bytes(16)
	t._clear_draft = dc.Bytes(16)
	t._peer = dc.Bytes(16)
	t._reply_to_msg_id = dc.Bytes(16)
	t._media = dc.Bytes(16)
	t._random_id = dc.Long()
	t._reply_markup = dc.Bytes(16)

}

// messages_forwardMessages#708e0195
type TL_messages_forwardMessages struct {
	_flags         []byte
	_silent        []byte
	_background    []byte
	_with_my_score []byte
	_grouped       []byte
	_from_peer     []byte
	_id            []byte
	_random_id     []byte
	_to_peer       []byte
}

func (t *TL_messages_forwardMessages) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_messages_forwardMessages) Get_flags() []byte {
	return t._flags
}

func (t *TL_messages_forwardMessages) Set_silent(_silent []byte) {
	t._silent = _silent
}

func (t *TL_messages_forwardMessages) Get_silent() []byte {
	return t._silent
}

func (t *TL_messages_forwardMessages) Set_background(_background []byte) {
	t._background = _background
}

func (t *TL_messages_forwardMessages) Get_background() []byte {
	return t._background
}

func (t *TL_messages_forwardMessages) Set_with_my_score(_with_my_score []byte) {
	t._with_my_score = _with_my_score
}

func (t *TL_messages_forwardMessages) Get_with_my_score() []byte {
	return t._with_my_score
}

func (t *TL_messages_forwardMessages) Set_grouped(_grouped []byte) {
	t._grouped = _grouped
}

func (t *TL_messages_forwardMessages) Get_grouped() []byte {
	return t._grouped
}

func (t *TL_messages_forwardMessages) Set_from_peer(_from_peer []byte) {
	t._from_peer = _from_peer
}

func (t *TL_messages_forwardMessages) Get_from_peer() []byte {
	return t._from_peer
}

func (t *TL_messages_forwardMessages) Set_id(_id []byte) {
	t._id = _id
}

func (t *TL_messages_forwardMessages) Get_id() []byte {
	return t._id
}

func (t *TL_messages_forwardMessages) Set_random_id(_random_id []byte) {
	t._random_id = _random_id
}

func (t *TL_messages_forwardMessages) Get_random_id() []byte {
	return t._random_id
}

func (t *TL_messages_forwardMessages) Set_to_peer(_to_peer []byte) {
	t._to_peer = _to_peer
}

func (t *TL_messages_forwardMessages) Get_to_peer() []byte {
	return t._to_peer
}

func New_TL_messages_forwardMessages() *TL_messages_forwardMessages {
	return &TL_messages_forwardMessages{}
}

func (t *TL_messages_forwardMessages) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_forwardMessages))
	ec.Bytes(t.Get_silent())
	ec.Bytes(t.Get_background())
	ec.Bytes(t.Get_with_my_score())
	ec.Bytes(t.Get_grouped())
	ec.Bytes(t.Get_from_peer())
	ec.Bytes(t.Get_id())
	ec.Bytes(t.Get_random_id())
	ec.Bytes(t.Get_to_peer())

	return ec.GetBuffer()
}

func (t *TL_messages_forwardMessages) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._silent = dc.Bytes(16)
	t._background = dc.Bytes(16)
	t._with_my_score = dc.Bytes(16)
	t._grouped = dc.Bytes(16)
	t._from_peer = dc.Bytes(16)
	t._id = dc.Bytes(16)
	t._random_id = dc.Bytes(16)
	t._to_peer = dc.Bytes(16)

}

// messages_reportSpam#cf1592db
type TL_messages_reportSpam struct {
	_peer []byte
}

func (t *TL_messages_reportSpam) Set_peer(_peer []byte) {
	t._peer = _peer
}

func (t *TL_messages_reportSpam) Get_peer() []byte {
	return t._peer
}

func New_TL_messages_reportSpam() *TL_messages_reportSpam {
	return &TL_messages_reportSpam{}
}

func (t *TL_messages_reportSpam) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_reportSpam))
	ec.Bytes(t.Get_peer())

	return ec.GetBuffer()
}

func (t *TL_messages_reportSpam) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._peer = dc.Bytes(16)

}

// messages_hideReportSpam#a8f1709b
type TL_messages_hideReportSpam struct {
	_peer []byte
}

func (t *TL_messages_hideReportSpam) Set_peer(_peer []byte) {
	t._peer = _peer
}

func (t *TL_messages_hideReportSpam) Get_peer() []byte {
	return t._peer
}

func New_TL_messages_hideReportSpam() *TL_messages_hideReportSpam {
	return &TL_messages_hideReportSpam{}
}

func (t *TL_messages_hideReportSpam) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_hideReportSpam))
	ec.Bytes(t.Get_peer())

	return ec.GetBuffer()
}

func (t *TL_messages_hideReportSpam) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._peer = dc.Bytes(16)

}

// messages_getPeerSettings#3672e09c
type TL_messages_getPeerSettings struct {
	_peer []byte
}

func (t *TL_messages_getPeerSettings) Set_peer(_peer []byte) {
	t._peer = _peer
}

func (t *TL_messages_getPeerSettings) Get_peer() []byte {
	return t._peer
}

func New_TL_messages_getPeerSettings() *TL_messages_getPeerSettings {
	return &TL_messages_getPeerSettings{}
}

func (t *TL_messages_getPeerSettings) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_getPeerSettings))
	ec.Bytes(t.Get_peer())

	return ec.GetBuffer()
}

func (t *TL_messages_getPeerSettings) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._peer = dc.Bytes(16)

}

// messages_getChats#3c6aa187
type TL_messages_getChats struct {
	_id []byte
}

func (t *TL_messages_getChats) Set_id(_id []byte) {
	t._id = _id
}

func (t *TL_messages_getChats) Get_id() []byte {
	return t._id
}

func New_TL_messages_getChats() *TL_messages_getChats {
	return &TL_messages_getChats{}
}

func (t *TL_messages_getChats) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_getChats))
	ec.Bytes(t.Get_id())

	return ec.GetBuffer()
}

func (t *TL_messages_getChats) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Bytes(16)

}

// messages_getFullChat#3b831c66
type TL_messages_getFullChat struct {
	_chat_id int32
}

func (t *TL_messages_getFullChat) Set_chat_id(_chat_id int32) {
	t._chat_id = _chat_id
}

func (t *TL_messages_getFullChat) Get_chat_id() int32 {
	return t._chat_id
}

func New_TL_messages_getFullChat() *TL_messages_getFullChat {
	return &TL_messages_getFullChat{}
}

func (t *TL_messages_getFullChat) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_getFullChat))
	ec.Int(t.Get_chat_id())

	return ec.GetBuffer()
}

func (t *TL_messages_getFullChat) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._chat_id = dc.Int()

}

// messages_editChatTitle#dc452855
type TL_messages_editChatTitle struct {
	_chat_id int32
	_title   string
}

func (t *TL_messages_editChatTitle) Set_chat_id(_chat_id int32) {
	t._chat_id = _chat_id
}

func (t *TL_messages_editChatTitle) Get_chat_id() int32 {
	return t._chat_id
}

func (t *TL_messages_editChatTitle) Set_title(_title string) {
	t._title = _title
}

func (t *TL_messages_editChatTitle) Get_title() string {
	return t._title
}

func New_TL_messages_editChatTitle() *TL_messages_editChatTitle {
	return &TL_messages_editChatTitle{}
}

func (t *TL_messages_editChatTitle) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_editChatTitle))
	ec.Int(t.Get_chat_id())
	ec.String(t.Get_title())

	return ec.GetBuffer()
}

func (t *TL_messages_editChatTitle) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._chat_id = dc.Int()
	t._title = dc.String()

}

// messages_editChatPhoto#ca4c79d8
type TL_messages_editChatPhoto struct {
	_chat_id int32
	_photo   []byte
}

func (t *TL_messages_editChatPhoto) Set_chat_id(_chat_id int32) {
	t._chat_id = _chat_id
}

func (t *TL_messages_editChatPhoto) Get_chat_id() int32 {
	return t._chat_id
}

func (t *TL_messages_editChatPhoto) Set_photo(_photo []byte) {
	t._photo = _photo
}

func (t *TL_messages_editChatPhoto) Get_photo() []byte {
	return t._photo
}

func New_TL_messages_editChatPhoto() *TL_messages_editChatPhoto {
	return &TL_messages_editChatPhoto{}
}

func (t *TL_messages_editChatPhoto) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_editChatPhoto))
	ec.Int(t.Get_chat_id())
	ec.Bytes(t.Get_photo())

	return ec.GetBuffer()
}

func (t *TL_messages_editChatPhoto) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._chat_id = dc.Int()
	t._photo = dc.Bytes(16)

}

// messages_addChatUser#f9a0aa09
type TL_messages_addChatUser struct {
	_chat_id   int32
	_user_id   []byte
	_fwd_limit int32
}

func (t *TL_messages_addChatUser) Set_chat_id(_chat_id int32) {
	t._chat_id = _chat_id
}

func (t *TL_messages_addChatUser) Get_chat_id() int32 {
	return t._chat_id
}

func (t *TL_messages_addChatUser) Set_user_id(_user_id []byte) {
	t._user_id = _user_id
}

func (t *TL_messages_addChatUser) Get_user_id() []byte {
	return t._user_id
}

func (t *TL_messages_addChatUser) Set_fwd_limit(_fwd_limit int32) {
	t._fwd_limit = _fwd_limit
}

func (t *TL_messages_addChatUser) Get_fwd_limit() int32 {
	return t._fwd_limit
}

func New_TL_messages_addChatUser() *TL_messages_addChatUser {
	return &TL_messages_addChatUser{}
}

func (t *TL_messages_addChatUser) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_addChatUser))
	ec.Int(t.Get_chat_id())
	ec.Bytes(t.Get_user_id())
	ec.Int(t.Get_fwd_limit())

	return ec.GetBuffer()
}

func (t *TL_messages_addChatUser) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._chat_id = dc.Int()
	t._user_id = dc.Bytes(16)
	t._fwd_limit = dc.Int()

}

// messages_deleteChatUser#e0611f16
type TL_messages_deleteChatUser struct {
	_chat_id int32
	_user_id []byte
}

func (t *TL_messages_deleteChatUser) Set_chat_id(_chat_id int32) {
	t._chat_id = _chat_id
}

func (t *TL_messages_deleteChatUser) Get_chat_id() int32 {
	return t._chat_id
}

func (t *TL_messages_deleteChatUser) Set_user_id(_user_id []byte) {
	t._user_id = _user_id
}

func (t *TL_messages_deleteChatUser) Get_user_id() []byte {
	return t._user_id
}

func New_TL_messages_deleteChatUser() *TL_messages_deleteChatUser {
	return &TL_messages_deleteChatUser{}
}

func (t *TL_messages_deleteChatUser) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_deleteChatUser))
	ec.Int(t.Get_chat_id())
	ec.Bytes(t.Get_user_id())

	return ec.GetBuffer()
}

func (t *TL_messages_deleteChatUser) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._chat_id = dc.Int()
	t._user_id = dc.Bytes(16)

}

// messages_createChat#09cb126e
type TL_messages_createChat struct {
	_users []byte
	_title string
}

func (t *TL_messages_createChat) Set_users(_users []byte) {
	t._users = _users
}

func (t *TL_messages_createChat) Get_users() []byte {
	return t._users
}

func (t *TL_messages_createChat) Set_title(_title string) {
	t._title = _title
}

func (t *TL_messages_createChat) Get_title() string {
	return t._title
}

func New_TL_messages_createChat() *TL_messages_createChat {
	return &TL_messages_createChat{}
}

func (t *TL_messages_createChat) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_createChat))
	ec.Bytes(t.Get_users())
	ec.String(t.Get_title())

	return ec.GetBuffer()
}

func (t *TL_messages_createChat) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._users = dc.Bytes(16)
	t._title = dc.String()

}

// messages_forwardMessage#33963bf9
type TL_messages_forwardMessage struct {
	_peer      []byte
	_id        int32
	_random_id int64
}

func (t *TL_messages_forwardMessage) Set_peer(_peer []byte) {
	t._peer = _peer
}

func (t *TL_messages_forwardMessage) Get_peer() []byte {
	return t._peer
}

func (t *TL_messages_forwardMessage) Set_id(_id int32) {
	t._id = _id
}

func (t *TL_messages_forwardMessage) Get_id() int32 {
	return t._id
}

func (t *TL_messages_forwardMessage) Set_random_id(_random_id int64) {
	t._random_id = _random_id
}

func (t *TL_messages_forwardMessage) Get_random_id() int64 {
	return t._random_id
}

func New_TL_messages_forwardMessage() *TL_messages_forwardMessage {
	return &TL_messages_forwardMessage{}
}

func (t *TL_messages_forwardMessage) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_forwardMessage))
	ec.Bytes(t.Get_peer())
	ec.Int(t.Get_id())
	ec.Long(t.Get_random_id())

	return ec.GetBuffer()
}

func (t *TL_messages_forwardMessage) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._peer = dc.Bytes(16)
	t._id = dc.Int()
	t._random_id = dc.Long()

}

// messages_getDhConfig#26cf8950
type TL_messages_getDhConfig struct {
	_version       int32
	_random_length int32
}

func (t *TL_messages_getDhConfig) Set_version(_version int32) {
	t._version = _version
}

func (t *TL_messages_getDhConfig) Get_version() int32 {
	return t._version
}

func (t *TL_messages_getDhConfig) Set_random_length(_random_length int32) {
	t._random_length = _random_length
}

func (t *TL_messages_getDhConfig) Get_random_length() int32 {
	return t._random_length
}

func New_TL_messages_getDhConfig() *TL_messages_getDhConfig {
	return &TL_messages_getDhConfig{}
}

func (t *TL_messages_getDhConfig) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_getDhConfig))
	ec.Int(t.Get_version())
	ec.Int(t.Get_random_length())

	return ec.GetBuffer()
}

func (t *TL_messages_getDhConfig) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._version = dc.Int()
	t._random_length = dc.Int()

}

// messages_requestEncryption#f64daf43
type TL_messages_requestEncryption struct {
	_user_id   []byte
	_random_id int32
	_g_a       []byte
}

func (t *TL_messages_requestEncryption) Set_user_id(_user_id []byte) {
	t._user_id = _user_id
}

func (t *TL_messages_requestEncryption) Get_user_id() []byte {
	return t._user_id
}

func (t *TL_messages_requestEncryption) Set_random_id(_random_id int32) {
	t._random_id = _random_id
}

func (t *TL_messages_requestEncryption) Get_random_id() int32 {
	return t._random_id
}

func (t *TL_messages_requestEncryption) Set_g_a(_g_a []byte) {
	t._g_a = _g_a
}

func (t *TL_messages_requestEncryption) Get_g_a() []byte {
	return t._g_a
}

func New_TL_messages_requestEncryption() *TL_messages_requestEncryption {
	return &TL_messages_requestEncryption{}
}

func (t *TL_messages_requestEncryption) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_requestEncryption))
	ec.Bytes(t.Get_user_id())
	ec.Int(t.Get_random_id())
	ec.Bytes(t.Get_g_a())

	return ec.GetBuffer()
}

func (t *TL_messages_requestEncryption) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._user_id = dc.Bytes(16)
	t._random_id = dc.Int()
	t._g_a = dc.Bytes(16)

}

// messages_acceptEncryption#3dbc0415
type TL_messages_acceptEncryption struct {
	_peer            []byte
	_g_b             []byte
	_key_fingerprint int64
}

func (t *TL_messages_acceptEncryption) Set_peer(_peer []byte) {
	t._peer = _peer
}

func (t *TL_messages_acceptEncryption) Get_peer() []byte {
	return t._peer
}

func (t *TL_messages_acceptEncryption) Set_g_b(_g_b []byte) {
	t._g_b = _g_b
}

func (t *TL_messages_acceptEncryption) Get_g_b() []byte {
	return t._g_b
}

func (t *TL_messages_acceptEncryption) Set_key_fingerprint(_key_fingerprint int64) {
	t._key_fingerprint = _key_fingerprint
}

func (t *TL_messages_acceptEncryption) Get_key_fingerprint() int64 {
	return t._key_fingerprint
}

func New_TL_messages_acceptEncryption() *TL_messages_acceptEncryption {
	return &TL_messages_acceptEncryption{}
}

func (t *TL_messages_acceptEncryption) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_acceptEncryption))
	ec.Bytes(t.Get_peer())
	ec.Bytes(t.Get_g_b())
	ec.Long(t.Get_key_fingerprint())

	return ec.GetBuffer()
}

func (t *TL_messages_acceptEncryption) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._peer = dc.Bytes(16)
	t._g_b = dc.Bytes(16)
	t._key_fingerprint = dc.Long()

}

// messages_discardEncryption#edd923c5
type TL_messages_discardEncryption struct {
	_chat_id int32
}

func (t *TL_messages_discardEncryption) Set_chat_id(_chat_id int32) {
	t._chat_id = _chat_id
}

func (t *TL_messages_discardEncryption) Get_chat_id() int32 {
	return t._chat_id
}

func New_TL_messages_discardEncryption() *TL_messages_discardEncryption {
	return &TL_messages_discardEncryption{}
}

func (t *TL_messages_discardEncryption) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_discardEncryption))
	ec.Int(t.Get_chat_id())

	return ec.GetBuffer()
}

func (t *TL_messages_discardEncryption) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._chat_id = dc.Int()

}

// messages_setEncryptedTyping#791451ed
type TL_messages_setEncryptedTyping struct {
	_peer   []byte
	_typing bool
}

func (t *TL_messages_setEncryptedTyping) Set_peer(_peer []byte) {
	t._peer = _peer
}

func (t *TL_messages_setEncryptedTyping) Get_peer() []byte {
	return t._peer
}

func (t *TL_messages_setEncryptedTyping) Set_typing(_typing bool) {
	t._typing = _typing
}

func (t *TL_messages_setEncryptedTyping) Get_typing() bool {
	return t._typing
}

func New_TL_messages_setEncryptedTyping() *TL_messages_setEncryptedTyping {
	return &TL_messages_setEncryptedTyping{}
}

func (t *TL_messages_setEncryptedTyping) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_setEncryptedTyping))
	ec.Bytes(t.Get_peer())
	ec.Bool(t.Get_typing())

	return ec.GetBuffer()
}

func (t *TL_messages_setEncryptedTyping) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._peer = dc.Bytes(16)
	t._typing = dc.Bool()

}

// messages_readEncryptedHistory#7f4b690a
type TL_messages_readEncryptedHistory struct {
	_peer     []byte
	_max_date int32
}

func (t *TL_messages_readEncryptedHistory) Set_peer(_peer []byte) {
	t._peer = _peer
}

func (t *TL_messages_readEncryptedHistory) Get_peer() []byte {
	return t._peer
}

func (t *TL_messages_readEncryptedHistory) Set_max_date(_max_date int32) {
	t._max_date = _max_date
}

func (t *TL_messages_readEncryptedHistory) Get_max_date() int32 {
	return t._max_date
}

func New_TL_messages_readEncryptedHistory() *TL_messages_readEncryptedHistory {
	return &TL_messages_readEncryptedHistory{}
}

func (t *TL_messages_readEncryptedHistory) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_readEncryptedHistory))
	ec.Bytes(t.Get_peer())
	ec.Int(t.Get_max_date())

	return ec.GetBuffer()
}

func (t *TL_messages_readEncryptedHistory) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._peer = dc.Bytes(16)
	t._max_date = dc.Int()

}

// messages_sendEncrypted#a9776773
type TL_messages_sendEncrypted struct {
	_peer      []byte
	_random_id int64
	_data      []byte
}

func (t *TL_messages_sendEncrypted) Set_peer(_peer []byte) {
	t._peer = _peer
}

func (t *TL_messages_sendEncrypted) Get_peer() []byte {
	return t._peer
}

func (t *TL_messages_sendEncrypted) Set_random_id(_random_id int64) {
	t._random_id = _random_id
}

func (t *TL_messages_sendEncrypted) Get_random_id() int64 {
	return t._random_id
}

func (t *TL_messages_sendEncrypted) Set_data(_data []byte) {
	t._data = _data
}

func (t *TL_messages_sendEncrypted) Get_data() []byte {
	return t._data
}

func New_TL_messages_sendEncrypted() *TL_messages_sendEncrypted {
	return &TL_messages_sendEncrypted{}
}

func (t *TL_messages_sendEncrypted) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_sendEncrypted))
	ec.Bytes(t.Get_peer())
	ec.Long(t.Get_random_id())
	ec.Bytes(t.Get_data())

	return ec.GetBuffer()
}

func (t *TL_messages_sendEncrypted) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._peer = dc.Bytes(16)
	t._random_id = dc.Long()
	t._data = dc.Bytes(16)

}

// messages_sendEncryptedFile#9a901b66
type TL_messages_sendEncryptedFile struct {
	_peer      []byte
	_random_id int64
	_data      []byte
	_file      []byte
}

func (t *TL_messages_sendEncryptedFile) Set_peer(_peer []byte) {
	t._peer = _peer
}

func (t *TL_messages_sendEncryptedFile) Get_peer() []byte {
	return t._peer
}

func (t *TL_messages_sendEncryptedFile) Set_random_id(_random_id int64) {
	t._random_id = _random_id
}

func (t *TL_messages_sendEncryptedFile) Get_random_id() int64 {
	return t._random_id
}

func (t *TL_messages_sendEncryptedFile) Set_data(_data []byte) {
	t._data = _data
}

func (t *TL_messages_sendEncryptedFile) Get_data() []byte {
	return t._data
}

func (t *TL_messages_sendEncryptedFile) Set_file(_file []byte) {
	t._file = _file
}

func (t *TL_messages_sendEncryptedFile) Get_file() []byte {
	return t._file
}

func New_TL_messages_sendEncryptedFile() *TL_messages_sendEncryptedFile {
	return &TL_messages_sendEncryptedFile{}
}

func (t *TL_messages_sendEncryptedFile) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_sendEncryptedFile))
	ec.Bytes(t.Get_peer())
	ec.Long(t.Get_random_id())
	ec.Bytes(t.Get_data())
	ec.Bytes(t.Get_file())

	return ec.GetBuffer()
}

func (t *TL_messages_sendEncryptedFile) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._peer = dc.Bytes(16)
	t._random_id = dc.Long()
	t._data = dc.Bytes(16)
	t._file = dc.Bytes(16)

}

// messages_sendEncryptedService#32d439a4
type TL_messages_sendEncryptedService struct {
	_peer      []byte
	_random_id int64
	_data      []byte
}

func (t *TL_messages_sendEncryptedService) Set_peer(_peer []byte) {
	t._peer = _peer
}

func (t *TL_messages_sendEncryptedService) Get_peer() []byte {
	return t._peer
}

func (t *TL_messages_sendEncryptedService) Set_random_id(_random_id int64) {
	t._random_id = _random_id
}

func (t *TL_messages_sendEncryptedService) Get_random_id() int64 {
	return t._random_id
}

func (t *TL_messages_sendEncryptedService) Set_data(_data []byte) {
	t._data = _data
}

func (t *TL_messages_sendEncryptedService) Get_data() []byte {
	return t._data
}

func New_TL_messages_sendEncryptedService() *TL_messages_sendEncryptedService {
	return &TL_messages_sendEncryptedService{}
}

func (t *TL_messages_sendEncryptedService) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_sendEncryptedService))
	ec.Bytes(t.Get_peer())
	ec.Long(t.Get_random_id())
	ec.Bytes(t.Get_data())

	return ec.GetBuffer()
}

func (t *TL_messages_sendEncryptedService) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._peer = dc.Bytes(16)
	t._random_id = dc.Long()
	t._data = dc.Bytes(16)

}

// messages_receivedQueue#55a5bb66
type TL_messages_receivedQueue struct {
	_max_qts int32
}

func (t *TL_messages_receivedQueue) Set_max_qts(_max_qts int32) {
	t._max_qts = _max_qts
}

func (t *TL_messages_receivedQueue) Get_max_qts() int32 {
	return t._max_qts
}

func New_TL_messages_receivedQueue() *TL_messages_receivedQueue {
	return &TL_messages_receivedQueue{}
}

func (t *TL_messages_receivedQueue) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_receivedQueue))
	ec.Int(t.Get_max_qts())

	return ec.GetBuffer()
}

func (t *TL_messages_receivedQueue) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._max_qts = dc.Int()

}

// messages_reportEncryptedSpam#4b0c8c0f
type TL_messages_reportEncryptedSpam struct {
	_peer []byte
}

func (t *TL_messages_reportEncryptedSpam) Set_peer(_peer []byte) {
	t._peer = _peer
}

func (t *TL_messages_reportEncryptedSpam) Get_peer() []byte {
	return t._peer
}

func New_TL_messages_reportEncryptedSpam() *TL_messages_reportEncryptedSpam {
	return &TL_messages_reportEncryptedSpam{}
}

func (t *TL_messages_reportEncryptedSpam) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_reportEncryptedSpam))
	ec.Bytes(t.Get_peer())

	return ec.GetBuffer()
}

func (t *TL_messages_reportEncryptedSpam) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._peer = dc.Bytes(16)

}

// messages_readMessageContents#36a73f77
type TL_messages_readMessageContents struct {
	_id []byte
}

func (t *TL_messages_readMessageContents) Set_id(_id []byte) {
	t._id = _id
}

func (t *TL_messages_readMessageContents) Get_id() []byte {
	return t._id
}

func New_TL_messages_readMessageContents() *TL_messages_readMessageContents {
	return &TL_messages_readMessageContents{}
}

func (t *TL_messages_readMessageContents) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_readMessageContents))
	ec.Bytes(t.Get_id())

	return ec.GetBuffer()
}

func (t *TL_messages_readMessageContents) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Bytes(16)

}

// messages_getAllStickers#1c9618b1
type TL_messages_getAllStickers struct {
	_hash int32
}

func (t *TL_messages_getAllStickers) Set_hash(_hash int32) {
	t._hash = _hash
}

func (t *TL_messages_getAllStickers) Get_hash() int32 {
	return t._hash
}

func New_TL_messages_getAllStickers() *TL_messages_getAllStickers {
	return &TL_messages_getAllStickers{}
}

func (t *TL_messages_getAllStickers) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_getAllStickers))
	ec.Int(t.Get_hash())

	return ec.GetBuffer()
}

func (t *TL_messages_getAllStickers) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._hash = dc.Int()

}

// messages_getWebPagePreview#25223e24
type TL_messages_getWebPagePreview struct {
	_message string
}

func (t *TL_messages_getWebPagePreview) Set_message(_message string) {
	t._message = _message
}

func (t *TL_messages_getWebPagePreview) Get_message() string {
	return t._message
}

func New_TL_messages_getWebPagePreview() *TL_messages_getWebPagePreview {
	return &TL_messages_getWebPagePreview{}
}

func (t *TL_messages_getWebPagePreview) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_getWebPagePreview))
	ec.String(t.Get_message())

	return ec.GetBuffer()
}

func (t *TL_messages_getWebPagePreview) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._message = dc.String()

}

// messages_exportChatInvite#7d885289
type TL_messages_exportChatInvite struct {
	_chat_id int32
}

func (t *TL_messages_exportChatInvite) Set_chat_id(_chat_id int32) {
	t._chat_id = _chat_id
}

func (t *TL_messages_exportChatInvite) Get_chat_id() int32 {
	return t._chat_id
}

func New_TL_messages_exportChatInvite() *TL_messages_exportChatInvite {
	return &TL_messages_exportChatInvite{}
}

func (t *TL_messages_exportChatInvite) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_exportChatInvite))
	ec.Int(t.Get_chat_id())

	return ec.GetBuffer()
}

func (t *TL_messages_exportChatInvite) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._chat_id = dc.Int()

}

// messages_checkChatInvite#3eadb1bb
type TL_messages_checkChatInvite struct {
	_hash string
}

func (t *TL_messages_checkChatInvite) Set_hash(_hash string) {
	t._hash = _hash
}

func (t *TL_messages_checkChatInvite) Get_hash() string {
	return t._hash
}

func New_TL_messages_checkChatInvite() *TL_messages_checkChatInvite {
	return &TL_messages_checkChatInvite{}
}

func (t *TL_messages_checkChatInvite) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_checkChatInvite))
	ec.String(t.Get_hash())

	return ec.GetBuffer()
}

func (t *TL_messages_checkChatInvite) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._hash = dc.String()

}

// messages_importChatInvite#6c50051c
type TL_messages_importChatInvite struct {
	_hash string
}

func (t *TL_messages_importChatInvite) Set_hash(_hash string) {
	t._hash = _hash
}

func (t *TL_messages_importChatInvite) Get_hash() string {
	return t._hash
}

func New_TL_messages_importChatInvite() *TL_messages_importChatInvite {
	return &TL_messages_importChatInvite{}
}

func (t *TL_messages_importChatInvite) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_importChatInvite))
	ec.String(t.Get_hash())

	return ec.GetBuffer()
}

func (t *TL_messages_importChatInvite) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._hash = dc.String()

}

// messages_getStickerSet#2619a90e
type TL_messages_getStickerSet struct {
	_stickerset []byte
}

func (t *TL_messages_getStickerSet) Set_stickerset(_stickerset []byte) {
	t._stickerset = _stickerset
}

func (t *TL_messages_getStickerSet) Get_stickerset() []byte {
	return t._stickerset
}

func New_TL_messages_getStickerSet() *TL_messages_getStickerSet {
	return &TL_messages_getStickerSet{}
}

func (t *TL_messages_getStickerSet) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_getStickerSet))
	ec.Bytes(t.Get_stickerset())

	return ec.GetBuffer()
}

func (t *TL_messages_getStickerSet) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._stickerset = dc.Bytes(16)

}

// messages_installStickerSet#c78fe460
type TL_messages_installStickerSet struct {
	_stickerset []byte
	_archived   bool
}

func (t *TL_messages_installStickerSet) Set_stickerset(_stickerset []byte) {
	t._stickerset = _stickerset
}

func (t *TL_messages_installStickerSet) Get_stickerset() []byte {
	return t._stickerset
}

func (t *TL_messages_installStickerSet) Set_archived(_archived bool) {
	t._archived = _archived
}

func (t *TL_messages_installStickerSet) Get_archived() bool {
	return t._archived
}

func New_TL_messages_installStickerSet() *TL_messages_installStickerSet {
	return &TL_messages_installStickerSet{}
}

func (t *TL_messages_installStickerSet) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_installStickerSet))
	ec.Bytes(t.Get_stickerset())
	ec.Bool(t.Get_archived())

	return ec.GetBuffer()
}

func (t *TL_messages_installStickerSet) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._stickerset = dc.Bytes(16)
	t._archived = dc.Bool()

}

// messages_uninstallStickerSet#f96e55de
type TL_messages_uninstallStickerSet struct {
	_stickerset []byte
}

func (t *TL_messages_uninstallStickerSet) Set_stickerset(_stickerset []byte) {
	t._stickerset = _stickerset
}

func (t *TL_messages_uninstallStickerSet) Get_stickerset() []byte {
	return t._stickerset
}

func New_TL_messages_uninstallStickerSet() *TL_messages_uninstallStickerSet {
	return &TL_messages_uninstallStickerSet{}
}

func (t *TL_messages_uninstallStickerSet) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_uninstallStickerSet))
	ec.Bytes(t.Get_stickerset())

	return ec.GetBuffer()
}

func (t *TL_messages_uninstallStickerSet) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._stickerset = dc.Bytes(16)

}

// messages_startBot#e6df7378
type TL_messages_startBot struct {
	_bot         []byte
	_peer        []byte
	_random_id   int64
	_start_param string
}

func (t *TL_messages_startBot) Set_bot(_bot []byte) {
	t._bot = _bot
}

func (t *TL_messages_startBot) Get_bot() []byte {
	return t._bot
}

func (t *TL_messages_startBot) Set_peer(_peer []byte) {
	t._peer = _peer
}

func (t *TL_messages_startBot) Get_peer() []byte {
	return t._peer
}

func (t *TL_messages_startBot) Set_random_id(_random_id int64) {
	t._random_id = _random_id
}

func (t *TL_messages_startBot) Get_random_id() int64 {
	return t._random_id
}

func (t *TL_messages_startBot) Set_start_param(_start_param string) {
	t._start_param = _start_param
}

func (t *TL_messages_startBot) Get_start_param() string {
	return t._start_param
}

func New_TL_messages_startBot() *TL_messages_startBot {
	return &TL_messages_startBot{}
}

func (t *TL_messages_startBot) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_startBot))
	ec.Bytes(t.Get_bot())
	ec.Bytes(t.Get_peer())
	ec.Long(t.Get_random_id())
	ec.String(t.Get_start_param())

	return ec.GetBuffer()
}

func (t *TL_messages_startBot) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._bot = dc.Bytes(16)
	t._peer = dc.Bytes(16)
	t._random_id = dc.Long()
	t._start_param = dc.String()

}

// messages_getMessagesViews#c4c8a55d
type TL_messages_getMessagesViews struct {
	_peer      []byte
	_id        []byte
	_increment bool
}

func (t *TL_messages_getMessagesViews) Set_peer(_peer []byte) {
	t._peer = _peer
}

func (t *TL_messages_getMessagesViews) Get_peer() []byte {
	return t._peer
}

func (t *TL_messages_getMessagesViews) Set_id(_id []byte) {
	t._id = _id
}

func (t *TL_messages_getMessagesViews) Get_id() []byte {
	return t._id
}

func (t *TL_messages_getMessagesViews) Set_increment(_increment bool) {
	t._increment = _increment
}

func (t *TL_messages_getMessagesViews) Get_increment() bool {
	return t._increment
}

func New_TL_messages_getMessagesViews() *TL_messages_getMessagesViews {
	return &TL_messages_getMessagesViews{}
}

func (t *TL_messages_getMessagesViews) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_getMessagesViews))
	ec.Bytes(t.Get_peer())
	ec.Bytes(t.Get_id())
	ec.Bool(t.Get_increment())

	return ec.GetBuffer()
}

func (t *TL_messages_getMessagesViews) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._peer = dc.Bytes(16)
	t._id = dc.Bytes(16)
	t._increment = dc.Bool()

}

// messages_toggleChatAdmins#ec8bd9e1
type TL_messages_toggleChatAdmins struct {
	_chat_id int32
	_enabled bool
}

func (t *TL_messages_toggleChatAdmins) Set_chat_id(_chat_id int32) {
	t._chat_id = _chat_id
}

func (t *TL_messages_toggleChatAdmins) Get_chat_id() int32 {
	return t._chat_id
}

func (t *TL_messages_toggleChatAdmins) Set_enabled(_enabled bool) {
	t._enabled = _enabled
}

func (t *TL_messages_toggleChatAdmins) Get_enabled() bool {
	return t._enabled
}

func New_TL_messages_toggleChatAdmins() *TL_messages_toggleChatAdmins {
	return &TL_messages_toggleChatAdmins{}
}

func (t *TL_messages_toggleChatAdmins) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_toggleChatAdmins))
	ec.Int(t.Get_chat_id())
	ec.Bool(t.Get_enabled())

	return ec.GetBuffer()
}

func (t *TL_messages_toggleChatAdmins) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._chat_id = dc.Int()
	t._enabled = dc.Bool()

}

// messages_editChatAdmin#a9e69f2e
type TL_messages_editChatAdmin struct {
	_chat_id  int32
	_user_id  []byte
	_is_admin bool
}

func (t *TL_messages_editChatAdmin) Set_chat_id(_chat_id int32) {
	t._chat_id = _chat_id
}

func (t *TL_messages_editChatAdmin) Get_chat_id() int32 {
	return t._chat_id
}

func (t *TL_messages_editChatAdmin) Set_user_id(_user_id []byte) {
	t._user_id = _user_id
}

func (t *TL_messages_editChatAdmin) Get_user_id() []byte {
	return t._user_id
}

func (t *TL_messages_editChatAdmin) Set_is_admin(_is_admin bool) {
	t._is_admin = _is_admin
}

func (t *TL_messages_editChatAdmin) Get_is_admin() bool {
	return t._is_admin
}

func New_TL_messages_editChatAdmin() *TL_messages_editChatAdmin {
	return &TL_messages_editChatAdmin{}
}

func (t *TL_messages_editChatAdmin) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_editChatAdmin))
	ec.Int(t.Get_chat_id())
	ec.Bytes(t.Get_user_id())
	ec.Bool(t.Get_is_admin())

	return ec.GetBuffer()
}

func (t *TL_messages_editChatAdmin) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._chat_id = dc.Int()
	t._user_id = dc.Bytes(16)
	t._is_admin = dc.Bool()

}

// messages_migrateChat#15a3b8e3
type TL_messages_migrateChat struct {
	_chat_id int32
}

func (t *TL_messages_migrateChat) Set_chat_id(_chat_id int32) {
	t._chat_id = _chat_id
}

func (t *TL_messages_migrateChat) Get_chat_id() int32 {
	return t._chat_id
}

func New_TL_messages_migrateChat() *TL_messages_migrateChat {
	return &TL_messages_migrateChat{}
}

func (t *TL_messages_migrateChat) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_migrateChat))
	ec.Int(t.Get_chat_id())

	return ec.GetBuffer()
}

func (t *TL_messages_migrateChat) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._chat_id = dc.Int()

}

// messages_searchGlobal#9e3cacb0
type TL_messages_searchGlobal struct {
	_q           string
	_offset_date int32
	_offset_peer []byte
	_offset_id   int32
	_limit       int32
}

func (t *TL_messages_searchGlobal) Set_q(_q string) {
	t._q = _q
}

func (t *TL_messages_searchGlobal) Get_q() string {
	return t._q
}

func (t *TL_messages_searchGlobal) Set_offset_date(_offset_date int32) {
	t._offset_date = _offset_date
}

func (t *TL_messages_searchGlobal) Get_offset_date() int32 {
	return t._offset_date
}

func (t *TL_messages_searchGlobal) Set_offset_peer(_offset_peer []byte) {
	t._offset_peer = _offset_peer
}

func (t *TL_messages_searchGlobal) Get_offset_peer() []byte {
	return t._offset_peer
}

func (t *TL_messages_searchGlobal) Set_offset_id(_offset_id int32) {
	t._offset_id = _offset_id
}

func (t *TL_messages_searchGlobal) Get_offset_id() int32 {
	return t._offset_id
}

func (t *TL_messages_searchGlobal) Set_limit(_limit int32) {
	t._limit = _limit
}

func (t *TL_messages_searchGlobal) Get_limit() int32 {
	return t._limit
}

func New_TL_messages_searchGlobal() *TL_messages_searchGlobal {
	return &TL_messages_searchGlobal{}
}

func (t *TL_messages_searchGlobal) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_searchGlobal))
	ec.String(t.Get_q())
	ec.Int(t.Get_offset_date())
	ec.Bytes(t.Get_offset_peer())
	ec.Int(t.Get_offset_id())
	ec.Int(t.Get_limit())

	return ec.GetBuffer()
}

func (t *TL_messages_searchGlobal) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._q = dc.String()
	t._offset_date = dc.Int()
	t._offset_peer = dc.Bytes(16)
	t._offset_id = dc.Int()
	t._limit = dc.Int()

}

// messages_reorderStickerSets#78337739
type TL_messages_reorderStickerSets struct {
	_flags []byte
	_masks []byte
	_order []byte
}

func (t *TL_messages_reorderStickerSets) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_messages_reorderStickerSets) Get_flags() []byte {
	return t._flags
}

func (t *TL_messages_reorderStickerSets) Set_masks(_masks []byte) {
	t._masks = _masks
}

func (t *TL_messages_reorderStickerSets) Get_masks() []byte {
	return t._masks
}

func (t *TL_messages_reorderStickerSets) Set_order(_order []byte) {
	t._order = _order
}

func (t *TL_messages_reorderStickerSets) Get_order() []byte {
	return t._order
}

func New_TL_messages_reorderStickerSets() *TL_messages_reorderStickerSets {
	return &TL_messages_reorderStickerSets{}
}

func (t *TL_messages_reorderStickerSets) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_reorderStickerSets))
	ec.Bytes(t.Get_masks())
	ec.Bytes(t.Get_order())

	return ec.GetBuffer()
}

func (t *TL_messages_reorderStickerSets) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._masks = dc.Bytes(16)
	t._order = dc.Bytes(16)

}

// messages_getDocumentByHash#338e2464
type TL_messages_getDocumentByHash struct {
	_sha256    []byte
	_size      int32
	_mime_type string
}

func (t *TL_messages_getDocumentByHash) Set_sha256(_sha256 []byte) {
	t._sha256 = _sha256
}

func (t *TL_messages_getDocumentByHash) Get_sha256() []byte {
	return t._sha256
}

func (t *TL_messages_getDocumentByHash) Set_size(_size int32) {
	t._size = _size
}

func (t *TL_messages_getDocumentByHash) Get_size() int32 {
	return t._size
}

func (t *TL_messages_getDocumentByHash) Set_mime_type(_mime_type string) {
	t._mime_type = _mime_type
}

func (t *TL_messages_getDocumentByHash) Get_mime_type() string {
	return t._mime_type
}

func New_TL_messages_getDocumentByHash() *TL_messages_getDocumentByHash {
	return &TL_messages_getDocumentByHash{}
}

func (t *TL_messages_getDocumentByHash) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_getDocumentByHash))
	ec.Bytes(t.Get_sha256())
	ec.Int(t.Get_size())
	ec.String(t.Get_mime_type())

	return ec.GetBuffer()
}

func (t *TL_messages_getDocumentByHash) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._sha256 = dc.Bytes(16)
	t._size = dc.Int()
	t._mime_type = dc.String()

}

// messages_searchGifs#bf9a776b
type TL_messages_searchGifs struct {
	_q      string
	_offset int32
}

func (t *TL_messages_searchGifs) Set_q(_q string) {
	t._q = _q
}

func (t *TL_messages_searchGifs) Get_q() string {
	return t._q
}

func (t *TL_messages_searchGifs) Set_offset(_offset int32) {
	t._offset = _offset
}

func (t *TL_messages_searchGifs) Get_offset() int32 {
	return t._offset
}

func New_TL_messages_searchGifs() *TL_messages_searchGifs {
	return &TL_messages_searchGifs{}
}

func (t *TL_messages_searchGifs) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_searchGifs))
	ec.String(t.Get_q())
	ec.Int(t.Get_offset())

	return ec.GetBuffer()
}

func (t *TL_messages_searchGifs) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._q = dc.String()
	t._offset = dc.Int()

}

// messages_getSavedGifs#83bf3d52
type TL_messages_getSavedGifs struct {
	_hash int32
}

func (t *TL_messages_getSavedGifs) Set_hash(_hash int32) {
	t._hash = _hash
}

func (t *TL_messages_getSavedGifs) Get_hash() int32 {
	return t._hash
}

func New_TL_messages_getSavedGifs() *TL_messages_getSavedGifs {
	return &TL_messages_getSavedGifs{}
}

func (t *TL_messages_getSavedGifs) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_getSavedGifs))
	ec.Int(t.Get_hash())

	return ec.GetBuffer()
}

func (t *TL_messages_getSavedGifs) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._hash = dc.Int()

}

// messages_saveGif#327a30cb
type TL_messages_saveGif struct {
	_id     []byte
	_unsave bool
}

func (t *TL_messages_saveGif) Set_id(_id []byte) {
	t._id = _id
}

func (t *TL_messages_saveGif) Get_id() []byte {
	return t._id
}

func (t *TL_messages_saveGif) Set_unsave(_unsave bool) {
	t._unsave = _unsave
}

func (t *TL_messages_saveGif) Get_unsave() bool {
	return t._unsave
}

func New_TL_messages_saveGif() *TL_messages_saveGif {
	return &TL_messages_saveGif{}
}

func (t *TL_messages_saveGif) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_saveGif))
	ec.Bytes(t.Get_id())
	ec.Bool(t.Get_unsave())

	return ec.GetBuffer()
}

func (t *TL_messages_saveGif) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Bytes(16)
	t._unsave = dc.Bool()

}

// messages_getInlineBotResults#514e999d
type TL_messages_getInlineBotResults struct {
	_flags     []byte
	_bot       []byte
	_peer      []byte
	_geo_point []byte
	_query     string
	_offset    string
}

func (t *TL_messages_getInlineBotResults) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_messages_getInlineBotResults) Get_flags() []byte {
	return t._flags
}

func (t *TL_messages_getInlineBotResults) Set_bot(_bot []byte) {
	t._bot = _bot
}

func (t *TL_messages_getInlineBotResults) Get_bot() []byte {
	return t._bot
}

func (t *TL_messages_getInlineBotResults) Set_peer(_peer []byte) {
	t._peer = _peer
}

func (t *TL_messages_getInlineBotResults) Get_peer() []byte {
	return t._peer
}

func (t *TL_messages_getInlineBotResults) Set_geo_point(_geo_point []byte) {
	t._geo_point = _geo_point
}

func (t *TL_messages_getInlineBotResults) Get_geo_point() []byte {
	return t._geo_point
}

func (t *TL_messages_getInlineBotResults) Set_query(_query string) {
	t._query = _query
}

func (t *TL_messages_getInlineBotResults) Get_query() string {
	return t._query
}

func (t *TL_messages_getInlineBotResults) Set_offset(_offset string) {
	t._offset = _offset
}

func (t *TL_messages_getInlineBotResults) Get_offset() string {
	return t._offset
}

func New_TL_messages_getInlineBotResults() *TL_messages_getInlineBotResults {
	return &TL_messages_getInlineBotResults{}
}

func (t *TL_messages_getInlineBotResults) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_getInlineBotResults))
	ec.Bytes(t.Get_bot())
	ec.Bytes(t.Get_peer())
	ec.Bytes(t.Get_geo_point())
	ec.String(t.Get_query())
	ec.String(t.Get_offset())

	return ec.GetBuffer()
}

func (t *TL_messages_getInlineBotResults) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._bot = dc.Bytes(16)
	t._peer = dc.Bytes(16)
	t._geo_point = dc.Bytes(16)
	t._query = dc.String()
	t._offset = dc.String()

}

// messages_setInlineBotResults#eb5ea206
type TL_messages_setInlineBotResults struct {
	_flags       []byte
	_gallery     []byte
	_private     []byte
	_query_id    int64
	_results     []byte
	_cache_time  int32
	_next_offset []byte
	_switch_pm   []byte
}

func (t *TL_messages_setInlineBotResults) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_messages_setInlineBotResults) Get_flags() []byte {
	return t._flags
}

func (t *TL_messages_setInlineBotResults) Set_gallery(_gallery []byte) {
	t._gallery = _gallery
}

func (t *TL_messages_setInlineBotResults) Get_gallery() []byte {
	return t._gallery
}

func (t *TL_messages_setInlineBotResults) Set_private(_private []byte) {
	t._private = _private
}

func (t *TL_messages_setInlineBotResults) Get_private() []byte {
	return t._private
}

func (t *TL_messages_setInlineBotResults) Set_query_id(_query_id int64) {
	t._query_id = _query_id
}

func (t *TL_messages_setInlineBotResults) Get_query_id() int64 {
	return t._query_id
}

func (t *TL_messages_setInlineBotResults) Set_results(_results []byte) {
	t._results = _results
}

func (t *TL_messages_setInlineBotResults) Get_results() []byte {
	return t._results
}

func (t *TL_messages_setInlineBotResults) Set_cache_time(_cache_time int32) {
	t._cache_time = _cache_time
}

func (t *TL_messages_setInlineBotResults) Get_cache_time() int32 {
	return t._cache_time
}

func (t *TL_messages_setInlineBotResults) Set_next_offset(_next_offset []byte) {
	t._next_offset = _next_offset
}

func (t *TL_messages_setInlineBotResults) Get_next_offset() []byte {
	return t._next_offset
}

func (t *TL_messages_setInlineBotResults) Set_switch_pm(_switch_pm []byte) {
	t._switch_pm = _switch_pm
}

func (t *TL_messages_setInlineBotResults) Get_switch_pm() []byte {
	return t._switch_pm
}

func New_TL_messages_setInlineBotResults() *TL_messages_setInlineBotResults {
	return &TL_messages_setInlineBotResults{}
}

func (t *TL_messages_setInlineBotResults) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_setInlineBotResults))
	ec.Bytes(t.Get_gallery())
	ec.Bytes(t.Get_private())
	ec.Long(t.Get_query_id())
	ec.Bytes(t.Get_results())
	ec.Int(t.Get_cache_time())
	ec.Bytes(t.Get_next_offset())
	ec.Bytes(t.Get_switch_pm())

	return ec.GetBuffer()
}

func (t *TL_messages_setInlineBotResults) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._gallery = dc.Bytes(16)
	t._private = dc.Bytes(16)
	t._query_id = dc.Long()
	t._results = dc.Bytes(16)
	t._cache_time = dc.Int()
	t._next_offset = dc.Bytes(16)
	t._switch_pm = dc.Bytes(16)

}

// messages_sendInlineBotResult#b16e06fe
type TL_messages_sendInlineBotResult struct {
	_flags           []byte
	_silent          []byte
	_background      []byte
	_clear_draft     []byte
	_peer            []byte
	_reply_to_msg_id []byte
	_random_id       int64
	_query_id        int64
	_id              string
}

func (t *TL_messages_sendInlineBotResult) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_messages_sendInlineBotResult) Get_flags() []byte {
	return t._flags
}

func (t *TL_messages_sendInlineBotResult) Set_silent(_silent []byte) {
	t._silent = _silent
}

func (t *TL_messages_sendInlineBotResult) Get_silent() []byte {
	return t._silent
}

func (t *TL_messages_sendInlineBotResult) Set_background(_background []byte) {
	t._background = _background
}

func (t *TL_messages_sendInlineBotResult) Get_background() []byte {
	return t._background
}

func (t *TL_messages_sendInlineBotResult) Set_clear_draft(_clear_draft []byte) {
	t._clear_draft = _clear_draft
}

func (t *TL_messages_sendInlineBotResult) Get_clear_draft() []byte {
	return t._clear_draft
}

func (t *TL_messages_sendInlineBotResult) Set_peer(_peer []byte) {
	t._peer = _peer
}

func (t *TL_messages_sendInlineBotResult) Get_peer() []byte {
	return t._peer
}

func (t *TL_messages_sendInlineBotResult) Set_reply_to_msg_id(_reply_to_msg_id []byte) {
	t._reply_to_msg_id = _reply_to_msg_id
}

func (t *TL_messages_sendInlineBotResult) Get_reply_to_msg_id() []byte {
	return t._reply_to_msg_id
}

func (t *TL_messages_sendInlineBotResult) Set_random_id(_random_id int64) {
	t._random_id = _random_id
}

func (t *TL_messages_sendInlineBotResult) Get_random_id() int64 {
	return t._random_id
}

func (t *TL_messages_sendInlineBotResult) Set_query_id(_query_id int64) {
	t._query_id = _query_id
}

func (t *TL_messages_sendInlineBotResult) Get_query_id() int64 {
	return t._query_id
}

func (t *TL_messages_sendInlineBotResult) Set_id(_id string) {
	t._id = _id
}

func (t *TL_messages_sendInlineBotResult) Get_id() string {
	return t._id
}

func New_TL_messages_sendInlineBotResult() *TL_messages_sendInlineBotResult {
	return &TL_messages_sendInlineBotResult{}
}

func (t *TL_messages_sendInlineBotResult) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_sendInlineBotResult))
	ec.Bytes(t.Get_silent())
	ec.Bytes(t.Get_background())
	ec.Bytes(t.Get_clear_draft())
	ec.Bytes(t.Get_peer())
	ec.Bytes(t.Get_reply_to_msg_id())
	ec.Long(t.Get_random_id())
	ec.Long(t.Get_query_id())
	ec.String(t.Get_id())

	return ec.GetBuffer()
}

func (t *TL_messages_sendInlineBotResult) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._silent = dc.Bytes(16)
	t._background = dc.Bytes(16)
	t._clear_draft = dc.Bytes(16)
	t._peer = dc.Bytes(16)
	t._reply_to_msg_id = dc.Bytes(16)
	t._random_id = dc.Long()
	t._query_id = dc.Long()
	t._id = dc.String()

}

// messages_getMessageEditData#fda68d36
type TL_messages_getMessageEditData struct {
	_peer []byte
	_id   int32
}

func (t *TL_messages_getMessageEditData) Set_peer(_peer []byte) {
	t._peer = _peer
}

func (t *TL_messages_getMessageEditData) Get_peer() []byte {
	return t._peer
}

func (t *TL_messages_getMessageEditData) Set_id(_id int32) {
	t._id = _id
}

func (t *TL_messages_getMessageEditData) Get_id() int32 {
	return t._id
}

func New_TL_messages_getMessageEditData() *TL_messages_getMessageEditData {
	return &TL_messages_getMessageEditData{}
}

func (t *TL_messages_getMessageEditData) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_getMessageEditData))
	ec.Bytes(t.Get_peer())
	ec.Int(t.Get_id())

	return ec.GetBuffer()
}

func (t *TL_messages_getMessageEditData) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._peer = dc.Bytes(16)
	t._id = dc.Int()

}

// messages_editMessage#05d1b8dd
type TL_messages_editMessage struct {
	_flags         []byte
	_no_webpage    []byte
	_stop_geo_live []byte
	_peer          []byte
	_id            int32
	_message       []byte
	_reply_markup  []byte
	_entities      []byte
	_geo_point     []byte
}

func (t *TL_messages_editMessage) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_messages_editMessage) Get_flags() []byte {
	return t._flags
}

func (t *TL_messages_editMessage) Set_no_webpage(_no_webpage []byte) {
	t._no_webpage = _no_webpage
}

func (t *TL_messages_editMessage) Get_no_webpage() []byte {
	return t._no_webpage
}

func (t *TL_messages_editMessage) Set_stop_geo_live(_stop_geo_live []byte) {
	t._stop_geo_live = _stop_geo_live
}

func (t *TL_messages_editMessage) Get_stop_geo_live() []byte {
	return t._stop_geo_live
}

func (t *TL_messages_editMessage) Set_peer(_peer []byte) {
	t._peer = _peer
}

func (t *TL_messages_editMessage) Get_peer() []byte {
	return t._peer
}

func (t *TL_messages_editMessage) Set_id(_id int32) {
	t._id = _id
}

func (t *TL_messages_editMessage) Get_id() int32 {
	return t._id
}

func (t *TL_messages_editMessage) Set_message(_message []byte) {
	t._message = _message
}

func (t *TL_messages_editMessage) Get_message() []byte {
	return t._message
}

func (t *TL_messages_editMessage) Set_reply_markup(_reply_markup []byte) {
	t._reply_markup = _reply_markup
}

func (t *TL_messages_editMessage) Get_reply_markup() []byte {
	return t._reply_markup
}

func (t *TL_messages_editMessage) Set_entities(_entities []byte) {
	t._entities = _entities
}

func (t *TL_messages_editMessage) Get_entities() []byte {
	return t._entities
}

func (t *TL_messages_editMessage) Set_geo_point(_geo_point []byte) {
	t._geo_point = _geo_point
}

func (t *TL_messages_editMessage) Get_geo_point() []byte {
	return t._geo_point
}

func New_TL_messages_editMessage() *TL_messages_editMessage {
	return &TL_messages_editMessage{}
}

func (t *TL_messages_editMessage) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_editMessage))
	ec.Bytes(t.Get_no_webpage())
	ec.Bytes(t.Get_stop_geo_live())
	ec.Bytes(t.Get_peer())
	ec.Int(t.Get_id())
	ec.Bytes(t.Get_message())
	ec.Bytes(t.Get_reply_markup())
	ec.Bytes(t.Get_entities())
	ec.Bytes(t.Get_geo_point())

	return ec.GetBuffer()
}

func (t *TL_messages_editMessage) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._no_webpage = dc.Bytes(16)
	t._stop_geo_live = dc.Bytes(16)
	t._peer = dc.Bytes(16)
	t._id = dc.Int()
	t._message = dc.Bytes(16)
	t._reply_markup = dc.Bytes(16)
	t._entities = dc.Bytes(16)
	t._geo_point = dc.Bytes(16)

}

// messages_editInlineBotMessage#b0e08243
type TL_messages_editInlineBotMessage struct {
	_flags         []byte
	_no_webpage    []byte
	_stop_geo_live []byte
	_id            []byte
	_message       []byte
	_reply_markup  []byte
	_entities      []byte
	_geo_point     []byte
}

func (t *TL_messages_editInlineBotMessage) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_messages_editInlineBotMessage) Get_flags() []byte {
	return t._flags
}

func (t *TL_messages_editInlineBotMessage) Set_no_webpage(_no_webpage []byte) {
	t._no_webpage = _no_webpage
}

func (t *TL_messages_editInlineBotMessage) Get_no_webpage() []byte {
	return t._no_webpage
}

func (t *TL_messages_editInlineBotMessage) Set_stop_geo_live(_stop_geo_live []byte) {
	t._stop_geo_live = _stop_geo_live
}

func (t *TL_messages_editInlineBotMessage) Get_stop_geo_live() []byte {
	return t._stop_geo_live
}

func (t *TL_messages_editInlineBotMessage) Set_id(_id []byte) {
	t._id = _id
}

func (t *TL_messages_editInlineBotMessage) Get_id() []byte {
	return t._id
}

func (t *TL_messages_editInlineBotMessage) Set_message(_message []byte) {
	t._message = _message
}

func (t *TL_messages_editInlineBotMessage) Get_message() []byte {
	return t._message
}

func (t *TL_messages_editInlineBotMessage) Set_reply_markup(_reply_markup []byte) {
	t._reply_markup = _reply_markup
}

func (t *TL_messages_editInlineBotMessage) Get_reply_markup() []byte {
	return t._reply_markup
}

func (t *TL_messages_editInlineBotMessage) Set_entities(_entities []byte) {
	t._entities = _entities
}

func (t *TL_messages_editInlineBotMessage) Get_entities() []byte {
	return t._entities
}

func (t *TL_messages_editInlineBotMessage) Set_geo_point(_geo_point []byte) {
	t._geo_point = _geo_point
}

func (t *TL_messages_editInlineBotMessage) Get_geo_point() []byte {
	return t._geo_point
}

func New_TL_messages_editInlineBotMessage() *TL_messages_editInlineBotMessage {
	return &TL_messages_editInlineBotMessage{}
}

func (t *TL_messages_editInlineBotMessage) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_editInlineBotMessage))
	ec.Bytes(t.Get_no_webpage())
	ec.Bytes(t.Get_stop_geo_live())
	ec.Bytes(t.Get_id())
	ec.Bytes(t.Get_message())
	ec.Bytes(t.Get_reply_markup())
	ec.Bytes(t.Get_entities())
	ec.Bytes(t.Get_geo_point())

	return ec.GetBuffer()
}

func (t *TL_messages_editInlineBotMessage) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._no_webpage = dc.Bytes(16)
	t._stop_geo_live = dc.Bytes(16)
	t._id = dc.Bytes(16)
	t._message = dc.Bytes(16)
	t._reply_markup = dc.Bytes(16)
	t._entities = dc.Bytes(16)
	t._geo_point = dc.Bytes(16)

}

// messages_getBotCallbackAnswer#810a9fec
type TL_messages_getBotCallbackAnswer struct {
	_flags  []byte
	_game   []byte
	_peer   []byte
	_msg_id int32
	_data   []byte
}

func (t *TL_messages_getBotCallbackAnswer) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_messages_getBotCallbackAnswer) Get_flags() []byte {
	return t._flags
}

func (t *TL_messages_getBotCallbackAnswer) Set_game(_game []byte) {
	t._game = _game
}

func (t *TL_messages_getBotCallbackAnswer) Get_game() []byte {
	return t._game
}

func (t *TL_messages_getBotCallbackAnswer) Set_peer(_peer []byte) {
	t._peer = _peer
}

func (t *TL_messages_getBotCallbackAnswer) Get_peer() []byte {
	return t._peer
}

func (t *TL_messages_getBotCallbackAnswer) Set_msg_id(_msg_id int32) {
	t._msg_id = _msg_id
}

func (t *TL_messages_getBotCallbackAnswer) Get_msg_id() int32 {
	return t._msg_id
}

func (t *TL_messages_getBotCallbackAnswer) Set_data(_data []byte) {
	t._data = _data
}

func (t *TL_messages_getBotCallbackAnswer) Get_data() []byte {
	return t._data
}

func New_TL_messages_getBotCallbackAnswer() *TL_messages_getBotCallbackAnswer {
	return &TL_messages_getBotCallbackAnswer{}
}

func (t *TL_messages_getBotCallbackAnswer) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_getBotCallbackAnswer))
	ec.Bytes(t.Get_game())
	ec.Bytes(t.Get_peer())
	ec.Int(t.Get_msg_id())
	ec.Bytes(t.Get_data())

	return ec.GetBuffer()
}

func (t *TL_messages_getBotCallbackAnswer) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._game = dc.Bytes(16)
	t._peer = dc.Bytes(16)
	t._msg_id = dc.Int()
	t._data = dc.Bytes(16)

}

// messages_setBotCallbackAnswer#d58f130a
type TL_messages_setBotCallbackAnswer struct {
	_flags      []byte
	_alert      []byte
	_query_id   int64
	_message    []byte
	_url        []byte
	_cache_time int32
}

func (t *TL_messages_setBotCallbackAnswer) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_messages_setBotCallbackAnswer) Get_flags() []byte {
	return t._flags
}

func (t *TL_messages_setBotCallbackAnswer) Set_alert(_alert []byte) {
	t._alert = _alert
}

func (t *TL_messages_setBotCallbackAnswer) Get_alert() []byte {
	return t._alert
}

func (t *TL_messages_setBotCallbackAnswer) Set_query_id(_query_id int64) {
	t._query_id = _query_id
}

func (t *TL_messages_setBotCallbackAnswer) Get_query_id() int64 {
	return t._query_id
}

func (t *TL_messages_setBotCallbackAnswer) Set_message(_message []byte) {
	t._message = _message
}

func (t *TL_messages_setBotCallbackAnswer) Get_message() []byte {
	return t._message
}

func (t *TL_messages_setBotCallbackAnswer) Set_url(_url []byte) {
	t._url = _url
}

func (t *TL_messages_setBotCallbackAnswer) Get_url() []byte {
	return t._url
}

func (t *TL_messages_setBotCallbackAnswer) Set_cache_time(_cache_time int32) {
	t._cache_time = _cache_time
}

func (t *TL_messages_setBotCallbackAnswer) Get_cache_time() int32 {
	return t._cache_time
}

func New_TL_messages_setBotCallbackAnswer() *TL_messages_setBotCallbackAnswer {
	return &TL_messages_setBotCallbackAnswer{}
}

func (t *TL_messages_setBotCallbackAnswer) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_setBotCallbackAnswer))
	ec.Bytes(t.Get_alert())
	ec.Long(t.Get_query_id())
	ec.Bytes(t.Get_message())
	ec.Bytes(t.Get_url())
	ec.Int(t.Get_cache_time())

	return ec.GetBuffer()
}

func (t *TL_messages_setBotCallbackAnswer) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._alert = dc.Bytes(16)
	t._query_id = dc.Long()
	t._message = dc.Bytes(16)
	t._url = dc.Bytes(16)
	t._cache_time = dc.Int()

}

// messages_getPeerDialogs#2d9776b9
type TL_messages_getPeerDialogs struct {
	_peers []byte
}

func (t *TL_messages_getPeerDialogs) Set_peers(_peers []byte) {
	t._peers = _peers
}

func (t *TL_messages_getPeerDialogs) Get_peers() []byte {
	return t._peers
}

func New_TL_messages_getPeerDialogs() *TL_messages_getPeerDialogs {
	return &TL_messages_getPeerDialogs{}
}

func (t *TL_messages_getPeerDialogs) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_getPeerDialogs))
	ec.Bytes(t.Get_peers())

	return ec.GetBuffer()
}

func (t *TL_messages_getPeerDialogs) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._peers = dc.Bytes(16)

}

// messages_saveDraft#bc39e14b
type TL_messages_saveDraft struct {
	_flags           []byte
	_no_webpage      []byte
	_reply_to_msg_id []byte
	_peer            []byte
	_message         string
	_entities        []byte
}

func (t *TL_messages_saveDraft) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_messages_saveDraft) Get_flags() []byte {
	return t._flags
}

func (t *TL_messages_saveDraft) Set_no_webpage(_no_webpage []byte) {
	t._no_webpage = _no_webpage
}

func (t *TL_messages_saveDraft) Get_no_webpage() []byte {
	return t._no_webpage
}

func (t *TL_messages_saveDraft) Set_reply_to_msg_id(_reply_to_msg_id []byte) {
	t._reply_to_msg_id = _reply_to_msg_id
}

func (t *TL_messages_saveDraft) Get_reply_to_msg_id() []byte {
	return t._reply_to_msg_id
}

func (t *TL_messages_saveDraft) Set_peer(_peer []byte) {
	t._peer = _peer
}

func (t *TL_messages_saveDraft) Get_peer() []byte {
	return t._peer
}

func (t *TL_messages_saveDraft) Set_message(_message string) {
	t._message = _message
}

func (t *TL_messages_saveDraft) Get_message() string {
	return t._message
}

func (t *TL_messages_saveDraft) Set_entities(_entities []byte) {
	t._entities = _entities
}

func (t *TL_messages_saveDraft) Get_entities() []byte {
	return t._entities
}

func New_TL_messages_saveDraft() *TL_messages_saveDraft {
	return &TL_messages_saveDraft{}
}

func (t *TL_messages_saveDraft) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_saveDraft))
	ec.Bytes(t.Get_no_webpage())
	ec.Bytes(t.Get_reply_to_msg_id())
	ec.Bytes(t.Get_peer())
	ec.String(t.Get_message())
	ec.Bytes(t.Get_entities())

	return ec.GetBuffer()
}

func (t *TL_messages_saveDraft) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._no_webpage = dc.Bytes(16)
	t._reply_to_msg_id = dc.Bytes(16)
	t._peer = dc.Bytes(16)
	t._message = dc.String()
	t._entities = dc.Bytes(16)

}

// messages_getAllDrafts#6a3f8d65
type TL_messages_getAllDrafts struct {
}

func New_TL_messages_getAllDrafts() *TL_messages_getAllDrafts {
	return &TL_messages_getAllDrafts{}
}

func (t *TL_messages_getAllDrafts) Encode() []byte {
	return nil
}

func (t *TL_messages_getAllDrafts) Decode(b []byte) {

}

// messages_getFeaturedStickers#2dacca4f
type TL_messages_getFeaturedStickers struct {
	_hash int32
}

func (t *TL_messages_getFeaturedStickers) Set_hash(_hash int32) {
	t._hash = _hash
}

func (t *TL_messages_getFeaturedStickers) Get_hash() int32 {
	return t._hash
}

func New_TL_messages_getFeaturedStickers() *TL_messages_getFeaturedStickers {
	return &TL_messages_getFeaturedStickers{}
}

func (t *TL_messages_getFeaturedStickers) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_getFeaturedStickers))
	ec.Int(t.Get_hash())

	return ec.GetBuffer()
}

func (t *TL_messages_getFeaturedStickers) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._hash = dc.Int()

}

// messages_readFeaturedStickers#5b118126
type TL_messages_readFeaturedStickers struct {
	_id []byte
}

func (t *TL_messages_readFeaturedStickers) Set_id(_id []byte) {
	t._id = _id
}

func (t *TL_messages_readFeaturedStickers) Get_id() []byte {
	return t._id
}

func New_TL_messages_readFeaturedStickers() *TL_messages_readFeaturedStickers {
	return &TL_messages_readFeaturedStickers{}
}

func (t *TL_messages_readFeaturedStickers) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_readFeaturedStickers))
	ec.Bytes(t.Get_id())

	return ec.GetBuffer()
}

func (t *TL_messages_readFeaturedStickers) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Bytes(16)

}

// messages_getRecentStickers#5ea192c9
type TL_messages_getRecentStickers struct {
	_flags    []byte
	_attached []byte
	_hash     int32
}

func (t *TL_messages_getRecentStickers) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_messages_getRecentStickers) Get_flags() []byte {
	return t._flags
}

func (t *TL_messages_getRecentStickers) Set_attached(_attached []byte) {
	t._attached = _attached
}

func (t *TL_messages_getRecentStickers) Get_attached() []byte {
	return t._attached
}

func (t *TL_messages_getRecentStickers) Set_hash(_hash int32) {
	t._hash = _hash
}

func (t *TL_messages_getRecentStickers) Get_hash() int32 {
	return t._hash
}

func New_TL_messages_getRecentStickers() *TL_messages_getRecentStickers {
	return &TL_messages_getRecentStickers{}
}

func (t *TL_messages_getRecentStickers) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_getRecentStickers))
	ec.Bytes(t.Get_attached())
	ec.Int(t.Get_hash())

	return ec.GetBuffer()
}

func (t *TL_messages_getRecentStickers) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._attached = dc.Bytes(16)
	t._hash = dc.Int()

}

// messages_saveRecentSticker#392718f8
type TL_messages_saveRecentSticker struct {
	_flags    []byte
	_attached []byte
	_id       []byte
	_unsave   bool
}

func (t *TL_messages_saveRecentSticker) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_messages_saveRecentSticker) Get_flags() []byte {
	return t._flags
}

func (t *TL_messages_saveRecentSticker) Set_attached(_attached []byte) {
	t._attached = _attached
}

func (t *TL_messages_saveRecentSticker) Get_attached() []byte {
	return t._attached
}

func (t *TL_messages_saveRecentSticker) Set_id(_id []byte) {
	t._id = _id
}

func (t *TL_messages_saveRecentSticker) Get_id() []byte {
	return t._id
}

func (t *TL_messages_saveRecentSticker) Set_unsave(_unsave bool) {
	t._unsave = _unsave
}

func (t *TL_messages_saveRecentSticker) Get_unsave() bool {
	return t._unsave
}

func New_TL_messages_saveRecentSticker() *TL_messages_saveRecentSticker {
	return &TL_messages_saveRecentSticker{}
}

func (t *TL_messages_saveRecentSticker) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_saveRecentSticker))
	ec.Bytes(t.Get_attached())
	ec.Bytes(t.Get_id())
	ec.Bool(t.Get_unsave())

	return ec.GetBuffer()
}

func (t *TL_messages_saveRecentSticker) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._attached = dc.Bytes(16)
	t._id = dc.Bytes(16)
	t._unsave = dc.Bool()

}

// messages_clearRecentStickers#8999602d
type TL_messages_clearRecentStickers struct {
	_flags    []byte
	_attached []byte
}

func (t *TL_messages_clearRecentStickers) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_messages_clearRecentStickers) Get_flags() []byte {
	return t._flags
}

func (t *TL_messages_clearRecentStickers) Set_attached(_attached []byte) {
	t._attached = _attached
}

func (t *TL_messages_clearRecentStickers) Get_attached() []byte {
	return t._attached
}

func New_TL_messages_clearRecentStickers() *TL_messages_clearRecentStickers {
	return &TL_messages_clearRecentStickers{}
}

func (t *TL_messages_clearRecentStickers) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_clearRecentStickers))
	ec.Bytes(t.Get_attached())

	return ec.GetBuffer()
}

func (t *TL_messages_clearRecentStickers) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._attached = dc.Bytes(16)

}

// messages_getArchivedStickers#57f17692
type TL_messages_getArchivedStickers struct {
	_flags     []byte
	_masks     []byte
	_offset_id int64
	_limit     int32
}

func (t *TL_messages_getArchivedStickers) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_messages_getArchivedStickers) Get_flags() []byte {
	return t._flags
}

func (t *TL_messages_getArchivedStickers) Set_masks(_masks []byte) {
	t._masks = _masks
}

func (t *TL_messages_getArchivedStickers) Get_masks() []byte {
	return t._masks
}

func (t *TL_messages_getArchivedStickers) Set_offset_id(_offset_id int64) {
	t._offset_id = _offset_id
}

func (t *TL_messages_getArchivedStickers) Get_offset_id() int64 {
	return t._offset_id
}

func (t *TL_messages_getArchivedStickers) Set_limit(_limit int32) {
	t._limit = _limit
}

func (t *TL_messages_getArchivedStickers) Get_limit() int32 {
	return t._limit
}

func New_TL_messages_getArchivedStickers() *TL_messages_getArchivedStickers {
	return &TL_messages_getArchivedStickers{}
}

func (t *TL_messages_getArchivedStickers) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_getArchivedStickers))
	ec.Bytes(t.Get_masks())
	ec.Long(t.Get_offset_id())
	ec.Int(t.Get_limit())

	return ec.GetBuffer()
}

func (t *TL_messages_getArchivedStickers) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._masks = dc.Bytes(16)
	t._offset_id = dc.Long()
	t._limit = dc.Int()

}

// messages_getMaskStickers#65b8c79f
type TL_messages_getMaskStickers struct {
	_hash int32
}

func (t *TL_messages_getMaskStickers) Set_hash(_hash int32) {
	t._hash = _hash
}

func (t *TL_messages_getMaskStickers) Get_hash() int32 {
	return t._hash
}

func New_TL_messages_getMaskStickers() *TL_messages_getMaskStickers {
	return &TL_messages_getMaskStickers{}
}

func (t *TL_messages_getMaskStickers) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_getMaskStickers))
	ec.Int(t.Get_hash())

	return ec.GetBuffer()
}

func (t *TL_messages_getMaskStickers) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._hash = dc.Int()

}

// messages_getAttachedStickers#cc5b67cc
type TL_messages_getAttachedStickers struct {
	_media []byte
}

func (t *TL_messages_getAttachedStickers) Set_media(_media []byte) {
	t._media = _media
}

func (t *TL_messages_getAttachedStickers) Get_media() []byte {
	return t._media
}

func New_TL_messages_getAttachedStickers() *TL_messages_getAttachedStickers {
	return &TL_messages_getAttachedStickers{}
}

func (t *TL_messages_getAttachedStickers) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_getAttachedStickers))
	ec.Bytes(t.Get_media())

	return ec.GetBuffer()
}

func (t *TL_messages_getAttachedStickers) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._media = dc.Bytes(16)

}

// messages_setGameScore#8ef8ecc0
type TL_messages_setGameScore struct {
	_flags        []byte
	_edit_message []byte
	_force        []byte
	_peer         []byte
	_id           int32
	_user_id      []byte
	_score        int32
}

func (t *TL_messages_setGameScore) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_messages_setGameScore) Get_flags() []byte {
	return t._flags
}

func (t *TL_messages_setGameScore) Set_edit_message(_edit_message []byte) {
	t._edit_message = _edit_message
}

func (t *TL_messages_setGameScore) Get_edit_message() []byte {
	return t._edit_message
}

func (t *TL_messages_setGameScore) Set_force(_force []byte) {
	t._force = _force
}

func (t *TL_messages_setGameScore) Get_force() []byte {
	return t._force
}

func (t *TL_messages_setGameScore) Set_peer(_peer []byte) {
	t._peer = _peer
}

func (t *TL_messages_setGameScore) Get_peer() []byte {
	return t._peer
}

func (t *TL_messages_setGameScore) Set_id(_id int32) {
	t._id = _id
}

func (t *TL_messages_setGameScore) Get_id() int32 {
	return t._id
}

func (t *TL_messages_setGameScore) Set_user_id(_user_id []byte) {
	t._user_id = _user_id
}

func (t *TL_messages_setGameScore) Get_user_id() []byte {
	return t._user_id
}

func (t *TL_messages_setGameScore) Set_score(_score int32) {
	t._score = _score
}

func (t *TL_messages_setGameScore) Get_score() int32 {
	return t._score
}

func New_TL_messages_setGameScore() *TL_messages_setGameScore {
	return &TL_messages_setGameScore{}
}

func (t *TL_messages_setGameScore) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_setGameScore))
	ec.Bytes(t.Get_edit_message())
	ec.Bytes(t.Get_force())
	ec.Bytes(t.Get_peer())
	ec.Int(t.Get_id())
	ec.Bytes(t.Get_user_id())
	ec.Int(t.Get_score())

	return ec.GetBuffer()
}

func (t *TL_messages_setGameScore) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._edit_message = dc.Bytes(16)
	t._force = dc.Bytes(16)
	t._peer = dc.Bytes(16)
	t._id = dc.Int()
	t._user_id = dc.Bytes(16)
	t._score = dc.Int()

}

// messages_setInlineGameScore#15ad9f64
type TL_messages_setInlineGameScore struct {
	_flags        []byte
	_edit_message []byte
	_force        []byte
	_id           []byte
	_user_id      []byte
	_score        int32
}

func (t *TL_messages_setInlineGameScore) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_messages_setInlineGameScore) Get_flags() []byte {
	return t._flags
}

func (t *TL_messages_setInlineGameScore) Set_edit_message(_edit_message []byte) {
	t._edit_message = _edit_message
}

func (t *TL_messages_setInlineGameScore) Get_edit_message() []byte {
	return t._edit_message
}

func (t *TL_messages_setInlineGameScore) Set_force(_force []byte) {
	t._force = _force
}

func (t *TL_messages_setInlineGameScore) Get_force() []byte {
	return t._force
}

func (t *TL_messages_setInlineGameScore) Set_id(_id []byte) {
	t._id = _id
}

func (t *TL_messages_setInlineGameScore) Get_id() []byte {
	return t._id
}

func (t *TL_messages_setInlineGameScore) Set_user_id(_user_id []byte) {
	t._user_id = _user_id
}

func (t *TL_messages_setInlineGameScore) Get_user_id() []byte {
	return t._user_id
}

func (t *TL_messages_setInlineGameScore) Set_score(_score int32) {
	t._score = _score
}

func (t *TL_messages_setInlineGameScore) Get_score() int32 {
	return t._score
}

func New_TL_messages_setInlineGameScore() *TL_messages_setInlineGameScore {
	return &TL_messages_setInlineGameScore{}
}

func (t *TL_messages_setInlineGameScore) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_setInlineGameScore))
	ec.Bytes(t.Get_edit_message())
	ec.Bytes(t.Get_force())
	ec.Bytes(t.Get_id())
	ec.Bytes(t.Get_user_id())
	ec.Int(t.Get_score())

	return ec.GetBuffer()
}

func (t *TL_messages_setInlineGameScore) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._edit_message = dc.Bytes(16)
	t._force = dc.Bytes(16)
	t._id = dc.Bytes(16)
	t._user_id = dc.Bytes(16)
	t._score = dc.Int()

}

// messages_getGameHighScores#e822649d
type TL_messages_getGameHighScores struct {
	_peer    []byte
	_id      int32
	_user_id []byte
}

func (t *TL_messages_getGameHighScores) Set_peer(_peer []byte) {
	t._peer = _peer
}

func (t *TL_messages_getGameHighScores) Get_peer() []byte {
	return t._peer
}

func (t *TL_messages_getGameHighScores) Set_id(_id int32) {
	t._id = _id
}

func (t *TL_messages_getGameHighScores) Get_id() int32 {
	return t._id
}

func (t *TL_messages_getGameHighScores) Set_user_id(_user_id []byte) {
	t._user_id = _user_id
}

func (t *TL_messages_getGameHighScores) Get_user_id() []byte {
	return t._user_id
}

func New_TL_messages_getGameHighScores() *TL_messages_getGameHighScores {
	return &TL_messages_getGameHighScores{}
}

func (t *TL_messages_getGameHighScores) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_getGameHighScores))
	ec.Bytes(t.Get_peer())
	ec.Int(t.Get_id())
	ec.Bytes(t.Get_user_id())

	return ec.GetBuffer()
}

func (t *TL_messages_getGameHighScores) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._peer = dc.Bytes(16)
	t._id = dc.Int()
	t._user_id = dc.Bytes(16)

}

// messages_getInlineGameHighScores#0f635e1b
type TL_messages_getInlineGameHighScores struct {
	_id      []byte
	_user_id []byte
}

func (t *TL_messages_getInlineGameHighScores) Set_id(_id []byte) {
	t._id = _id
}

func (t *TL_messages_getInlineGameHighScores) Get_id() []byte {
	return t._id
}

func (t *TL_messages_getInlineGameHighScores) Set_user_id(_user_id []byte) {
	t._user_id = _user_id
}

func (t *TL_messages_getInlineGameHighScores) Get_user_id() []byte {
	return t._user_id
}

func New_TL_messages_getInlineGameHighScores() *TL_messages_getInlineGameHighScores {
	return &TL_messages_getInlineGameHighScores{}
}

func (t *TL_messages_getInlineGameHighScores) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_getInlineGameHighScores))
	ec.Bytes(t.Get_id())
	ec.Bytes(t.Get_user_id())

	return ec.GetBuffer()
}

func (t *TL_messages_getInlineGameHighScores) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Bytes(16)
	t._user_id = dc.Bytes(16)

}

// messages_getCommonChats#0d0a48c4
type TL_messages_getCommonChats struct {
	_user_id []byte
	_max_id  int32
	_limit   int32
}

func (t *TL_messages_getCommonChats) Set_user_id(_user_id []byte) {
	t._user_id = _user_id
}

func (t *TL_messages_getCommonChats) Get_user_id() []byte {
	return t._user_id
}

func (t *TL_messages_getCommonChats) Set_max_id(_max_id int32) {
	t._max_id = _max_id
}

func (t *TL_messages_getCommonChats) Get_max_id() int32 {
	return t._max_id
}

func (t *TL_messages_getCommonChats) Set_limit(_limit int32) {
	t._limit = _limit
}

func (t *TL_messages_getCommonChats) Get_limit() int32 {
	return t._limit
}

func New_TL_messages_getCommonChats() *TL_messages_getCommonChats {
	return &TL_messages_getCommonChats{}
}

func (t *TL_messages_getCommonChats) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_getCommonChats))
	ec.Bytes(t.Get_user_id())
	ec.Int(t.Get_max_id())
	ec.Int(t.Get_limit())

	return ec.GetBuffer()
}

func (t *TL_messages_getCommonChats) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._user_id = dc.Bytes(16)
	t._max_id = dc.Int()
	t._limit = dc.Int()

}

// messages_getAllChats#eba80ff0
type TL_messages_getAllChats struct {
	_except_ids []byte
}

func (t *TL_messages_getAllChats) Set_except_ids(_except_ids []byte) {
	t._except_ids = _except_ids
}

func (t *TL_messages_getAllChats) Get_except_ids() []byte {
	return t._except_ids
}

func New_TL_messages_getAllChats() *TL_messages_getAllChats {
	return &TL_messages_getAllChats{}
}

func (t *TL_messages_getAllChats) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_getAllChats))
	ec.Bytes(t.Get_except_ids())

	return ec.GetBuffer()
}

func (t *TL_messages_getAllChats) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._except_ids = dc.Bytes(16)

}

// messages_getWebPage#32ca8f91
type TL_messages_getWebPage struct {
	_url  string
	_hash int32
}

func (t *TL_messages_getWebPage) Set_url(_url string) {
	t._url = _url
}

func (t *TL_messages_getWebPage) Get_url() string {
	return t._url
}

func (t *TL_messages_getWebPage) Set_hash(_hash int32) {
	t._hash = _hash
}

func (t *TL_messages_getWebPage) Get_hash() int32 {
	return t._hash
}

func New_TL_messages_getWebPage() *TL_messages_getWebPage {
	return &TL_messages_getWebPage{}
}

func (t *TL_messages_getWebPage) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_getWebPage))
	ec.String(t.Get_url())
	ec.Int(t.Get_hash())

	return ec.GetBuffer()
}

func (t *TL_messages_getWebPage) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._url = dc.String()
	t._hash = dc.Int()

}

// messages_toggleDialogPin#3289be6a
type TL_messages_toggleDialogPin struct {
	_flags  []byte
	_pinned []byte
	_peer   []byte
}

func (t *TL_messages_toggleDialogPin) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_messages_toggleDialogPin) Get_flags() []byte {
	return t._flags
}

func (t *TL_messages_toggleDialogPin) Set_pinned(_pinned []byte) {
	t._pinned = _pinned
}

func (t *TL_messages_toggleDialogPin) Get_pinned() []byte {
	return t._pinned
}

func (t *TL_messages_toggleDialogPin) Set_peer(_peer []byte) {
	t._peer = _peer
}

func (t *TL_messages_toggleDialogPin) Get_peer() []byte {
	return t._peer
}

func New_TL_messages_toggleDialogPin() *TL_messages_toggleDialogPin {
	return &TL_messages_toggleDialogPin{}
}

func (t *TL_messages_toggleDialogPin) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_toggleDialogPin))
	ec.Bytes(t.Get_pinned())
	ec.Bytes(t.Get_peer())

	return ec.GetBuffer()
}

func (t *TL_messages_toggleDialogPin) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._pinned = dc.Bytes(16)
	t._peer = dc.Bytes(16)

}

// messages_reorderPinnedDialogs#959ff644
type TL_messages_reorderPinnedDialogs struct {
	_flags []byte
	_force []byte
	_order []byte
}

func (t *TL_messages_reorderPinnedDialogs) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_messages_reorderPinnedDialogs) Get_flags() []byte {
	return t._flags
}

func (t *TL_messages_reorderPinnedDialogs) Set_force(_force []byte) {
	t._force = _force
}

func (t *TL_messages_reorderPinnedDialogs) Get_force() []byte {
	return t._force
}

func (t *TL_messages_reorderPinnedDialogs) Set_order(_order []byte) {
	t._order = _order
}

func (t *TL_messages_reorderPinnedDialogs) Get_order() []byte {
	return t._order
}

func New_TL_messages_reorderPinnedDialogs() *TL_messages_reorderPinnedDialogs {
	return &TL_messages_reorderPinnedDialogs{}
}

func (t *TL_messages_reorderPinnedDialogs) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_reorderPinnedDialogs))
	ec.Bytes(t.Get_force())
	ec.Bytes(t.Get_order())

	return ec.GetBuffer()
}

func (t *TL_messages_reorderPinnedDialogs) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._force = dc.Bytes(16)
	t._order = dc.Bytes(16)

}

// messages_getPinnedDialogs#e254d64e
type TL_messages_getPinnedDialogs struct {
}

func New_TL_messages_getPinnedDialogs() *TL_messages_getPinnedDialogs {
	return &TL_messages_getPinnedDialogs{}
}

func (t *TL_messages_getPinnedDialogs) Encode() []byte {
	return nil
}

func (t *TL_messages_getPinnedDialogs) Decode(b []byte) {

}

// messages_setBotShippingResults#e5f672fa
type TL_messages_setBotShippingResults struct {
	_flags            []byte
	_query_id         int64
	_error            []byte
	_shipping_options []byte
}

func (t *TL_messages_setBotShippingResults) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_messages_setBotShippingResults) Get_flags() []byte {
	return t._flags
}

func (t *TL_messages_setBotShippingResults) Set_query_id(_query_id int64) {
	t._query_id = _query_id
}

func (t *TL_messages_setBotShippingResults) Get_query_id() int64 {
	return t._query_id
}

func (t *TL_messages_setBotShippingResults) Set_error(_error []byte) {
	t._error = _error
}

func (t *TL_messages_setBotShippingResults) Get_error() []byte {
	return t._error
}

func (t *TL_messages_setBotShippingResults) Set_shipping_options(_shipping_options []byte) {
	t._shipping_options = _shipping_options
}

func (t *TL_messages_setBotShippingResults) Get_shipping_options() []byte {
	return t._shipping_options
}

func New_TL_messages_setBotShippingResults() *TL_messages_setBotShippingResults {
	return &TL_messages_setBotShippingResults{}
}

func (t *TL_messages_setBotShippingResults) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_setBotShippingResults))
	ec.Long(t.Get_query_id())
	ec.Bytes(t.Get_error())
	ec.Bytes(t.Get_shipping_options())

	return ec.GetBuffer()
}

func (t *TL_messages_setBotShippingResults) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._query_id = dc.Long()
	t._error = dc.Bytes(16)
	t._shipping_options = dc.Bytes(16)

}

// messages_setBotPrecheckoutResults#09c2dd95
type TL_messages_setBotPrecheckoutResults struct {
	_flags    []byte
	_success  []byte
	_query_id int64
	_error    []byte
}

func (t *TL_messages_setBotPrecheckoutResults) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_messages_setBotPrecheckoutResults) Get_flags() []byte {
	return t._flags
}

func (t *TL_messages_setBotPrecheckoutResults) Set_success(_success []byte) {
	t._success = _success
}

func (t *TL_messages_setBotPrecheckoutResults) Get_success() []byte {
	return t._success
}

func (t *TL_messages_setBotPrecheckoutResults) Set_query_id(_query_id int64) {
	t._query_id = _query_id
}

func (t *TL_messages_setBotPrecheckoutResults) Get_query_id() int64 {
	return t._query_id
}

func (t *TL_messages_setBotPrecheckoutResults) Set_error(_error []byte) {
	t._error = _error
}

func (t *TL_messages_setBotPrecheckoutResults) Get_error() []byte {
	return t._error
}

func New_TL_messages_setBotPrecheckoutResults() *TL_messages_setBotPrecheckoutResults {
	return &TL_messages_setBotPrecheckoutResults{}
}

func (t *TL_messages_setBotPrecheckoutResults) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_setBotPrecheckoutResults))
	ec.Bytes(t.Get_success())
	ec.Long(t.Get_query_id())
	ec.Bytes(t.Get_error())

	return ec.GetBuffer()
}

func (t *TL_messages_setBotPrecheckoutResults) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._success = dc.Bytes(16)
	t._query_id = dc.Long()
	t._error = dc.Bytes(16)

}

// messages_uploadMedia#519bc2b1
type TL_messages_uploadMedia struct {
	_peer  []byte
	_media []byte
}

func (t *TL_messages_uploadMedia) Set_peer(_peer []byte) {
	t._peer = _peer
}

func (t *TL_messages_uploadMedia) Get_peer() []byte {
	return t._peer
}

func (t *TL_messages_uploadMedia) Set_media(_media []byte) {
	t._media = _media
}

func (t *TL_messages_uploadMedia) Get_media() []byte {
	return t._media
}

func New_TL_messages_uploadMedia() *TL_messages_uploadMedia {
	return &TL_messages_uploadMedia{}
}

func (t *TL_messages_uploadMedia) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_uploadMedia))
	ec.Bytes(t.Get_peer())
	ec.Bytes(t.Get_media())

	return ec.GetBuffer()
}

func (t *TL_messages_uploadMedia) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._peer = dc.Bytes(16)
	t._media = dc.Bytes(16)

}

// messages_sendScreenshotNotification#c97df020
type TL_messages_sendScreenshotNotification struct {
	_peer            []byte
	_reply_to_msg_id int32
	_random_id       int64
}

func (t *TL_messages_sendScreenshotNotification) Set_peer(_peer []byte) {
	t._peer = _peer
}

func (t *TL_messages_sendScreenshotNotification) Get_peer() []byte {
	return t._peer
}

func (t *TL_messages_sendScreenshotNotification) Set_reply_to_msg_id(_reply_to_msg_id int32) {
	t._reply_to_msg_id = _reply_to_msg_id
}

func (t *TL_messages_sendScreenshotNotification) Get_reply_to_msg_id() int32 {
	return t._reply_to_msg_id
}

func (t *TL_messages_sendScreenshotNotification) Set_random_id(_random_id int64) {
	t._random_id = _random_id
}

func (t *TL_messages_sendScreenshotNotification) Get_random_id() int64 {
	return t._random_id
}

func New_TL_messages_sendScreenshotNotification() *TL_messages_sendScreenshotNotification {
	return &TL_messages_sendScreenshotNotification{}
}

func (t *TL_messages_sendScreenshotNotification) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_sendScreenshotNotification))
	ec.Bytes(t.Get_peer())
	ec.Int(t.Get_reply_to_msg_id())
	ec.Long(t.Get_random_id())

	return ec.GetBuffer()
}

func (t *TL_messages_sendScreenshotNotification) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._peer = dc.Bytes(16)
	t._reply_to_msg_id = dc.Int()
	t._random_id = dc.Long()

}

// messages_getFavedStickers#21ce0b0e
type TL_messages_getFavedStickers struct {
	_hash int32
}

func (t *TL_messages_getFavedStickers) Set_hash(_hash int32) {
	t._hash = _hash
}

func (t *TL_messages_getFavedStickers) Get_hash() int32 {
	return t._hash
}

func New_TL_messages_getFavedStickers() *TL_messages_getFavedStickers {
	return &TL_messages_getFavedStickers{}
}

func (t *TL_messages_getFavedStickers) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_getFavedStickers))
	ec.Int(t.Get_hash())

	return ec.GetBuffer()
}

func (t *TL_messages_getFavedStickers) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._hash = dc.Int()

}

// messages_faveSticker#b9ffc55b
type TL_messages_faveSticker struct {
	_id     []byte
	_unfave bool
}

func (t *TL_messages_faveSticker) Set_id(_id []byte) {
	t._id = _id
}

func (t *TL_messages_faveSticker) Get_id() []byte {
	return t._id
}

func (t *TL_messages_faveSticker) Set_unfave(_unfave bool) {
	t._unfave = _unfave
}

func (t *TL_messages_faveSticker) Get_unfave() bool {
	return t._unfave
}

func New_TL_messages_faveSticker() *TL_messages_faveSticker {
	return &TL_messages_faveSticker{}
}

func (t *TL_messages_faveSticker) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_faveSticker))
	ec.Bytes(t.Get_id())
	ec.Bool(t.Get_unfave())

	return ec.GetBuffer()
}

func (t *TL_messages_faveSticker) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Bytes(16)
	t._unfave = dc.Bool()

}

// messages_getUnreadMentions#46578472
type TL_messages_getUnreadMentions struct {
	_peer       []byte
	_offset_id  int32
	_add_offset int32
	_limit      int32
	_max_id     int32
	_min_id     int32
}

func (t *TL_messages_getUnreadMentions) Set_peer(_peer []byte) {
	t._peer = _peer
}

func (t *TL_messages_getUnreadMentions) Get_peer() []byte {
	return t._peer
}

func (t *TL_messages_getUnreadMentions) Set_offset_id(_offset_id int32) {
	t._offset_id = _offset_id
}

func (t *TL_messages_getUnreadMentions) Get_offset_id() int32 {
	return t._offset_id
}

func (t *TL_messages_getUnreadMentions) Set_add_offset(_add_offset int32) {
	t._add_offset = _add_offset
}

func (t *TL_messages_getUnreadMentions) Get_add_offset() int32 {
	return t._add_offset
}

func (t *TL_messages_getUnreadMentions) Set_limit(_limit int32) {
	t._limit = _limit
}

func (t *TL_messages_getUnreadMentions) Get_limit() int32 {
	return t._limit
}

func (t *TL_messages_getUnreadMentions) Set_max_id(_max_id int32) {
	t._max_id = _max_id
}

func (t *TL_messages_getUnreadMentions) Get_max_id() int32 {
	return t._max_id
}

func (t *TL_messages_getUnreadMentions) Set_min_id(_min_id int32) {
	t._min_id = _min_id
}

func (t *TL_messages_getUnreadMentions) Get_min_id() int32 {
	return t._min_id
}

func New_TL_messages_getUnreadMentions() *TL_messages_getUnreadMentions {
	return &TL_messages_getUnreadMentions{}
}

func (t *TL_messages_getUnreadMentions) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_getUnreadMentions))
	ec.Bytes(t.Get_peer())
	ec.Int(t.Get_offset_id())
	ec.Int(t.Get_add_offset())
	ec.Int(t.Get_limit())
	ec.Int(t.Get_max_id())
	ec.Int(t.Get_min_id())

	return ec.GetBuffer()
}

func (t *TL_messages_getUnreadMentions) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._peer = dc.Bytes(16)
	t._offset_id = dc.Int()
	t._add_offset = dc.Int()
	t._limit = dc.Int()
	t._max_id = dc.Int()
	t._min_id = dc.Int()

}

// messages_readMentions#0f0189d3
type TL_messages_readMentions struct {
	_peer []byte
}

func (t *TL_messages_readMentions) Set_peer(_peer []byte) {
	t._peer = _peer
}

func (t *TL_messages_readMentions) Get_peer() []byte {
	return t._peer
}

func New_TL_messages_readMentions() *TL_messages_readMentions {
	return &TL_messages_readMentions{}
}

func (t *TL_messages_readMentions) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_readMentions))
	ec.Bytes(t.Get_peer())

	return ec.GetBuffer()
}

func (t *TL_messages_readMentions) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._peer = dc.Bytes(16)

}

// messages_getRecentLocations#249431e2
type TL_messages_getRecentLocations struct {
	_peer  []byte
	_limit int32
}

func (t *TL_messages_getRecentLocations) Set_peer(_peer []byte) {
	t._peer = _peer
}

func (t *TL_messages_getRecentLocations) Get_peer() []byte {
	return t._peer
}

func (t *TL_messages_getRecentLocations) Set_limit(_limit int32) {
	t._limit = _limit
}

func (t *TL_messages_getRecentLocations) Get_limit() int32 {
	return t._limit
}

func New_TL_messages_getRecentLocations() *TL_messages_getRecentLocations {
	return &TL_messages_getRecentLocations{}
}

func (t *TL_messages_getRecentLocations) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_getRecentLocations))
	ec.Bytes(t.Get_peer())
	ec.Int(t.Get_limit())

	return ec.GetBuffer()
}

func (t *TL_messages_getRecentLocations) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._peer = dc.Bytes(16)
	t._limit = dc.Int()

}

// messages_sendMultiMedia#2095512f
type TL_messages_sendMultiMedia struct {
	_flags           []byte
	_silent          []byte
	_background      []byte
	_clear_draft     []byte
	_peer            []byte
	_reply_to_msg_id []byte
	_multi_media     []byte
}

func (t *TL_messages_sendMultiMedia) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_messages_sendMultiMedia) Get_flags() []byte {
	return t._flags
}

func (t *TL_messages_sendMultiMedia) Set_silent(_silent []byte) {
	t._silent = _silent
}

func (t *TL_messages_sendMultiMedia) Get_silent() []byte {
	return t._silent
}

func (t *TL_messages_sendMultiMedia) Set_background(_background []byte) {
	t._background = _background
}

func (t *TL_messages_sendMultiMedia) Get_background() []byte {
	return t._background
}

func (t *TL_messages_sendMultiMedia) Set_clear_draft(_clear_draft []byte) {
	t._clear_draft = _clear_draft
}

func (t *TL_messages_sendMultiMedia) Get_clear_draft() []byte {
	return t._clear_draft
}

func (t *TL_messages_sendMultiMedia) Set_peer(_peer []byte) {
	t._peer = _peer
}

func (t *TL_messages_sendMultiMedia) Get_peer() []byte {
	return t._peer
}

func (t *TL_messages_sendMultiMedia) Set_reply_to_msg_id(_reply_to_msg_id []byte) {
	t._reply_to_msg_id = _reply_to_msg_id
}

func (t *TL_messages_sendMultiMedia) Get_reply_to_msg_id() []byte {
	return t._reply_to_msg_id
}

func (t *TL_messages_sendMultiMedia) Set_multi_media(_multi_media []byte) {
	t._multi_media = _multi_media
}

func (t *TL_messages_sendMultiMedia) Get_multi_media() []byte {
	return t._multi_media
}

func New_TL_messages_sendMultiMedia() *TL_messages_sendMultiMedia {
	return &TL_messages_sendMultiMedia{}
}

func (t *TL_messages_sendMultiMedia) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_sendMultiMedia))
	ec.Bytes(t.Get_silent())
	ec.Bytes(t.Get_background())
	ec.Bytes(t.Get_clear_draft())
	ec.Bytes(t.Get_peer())
	ec.Bytes(t.Get_reply_to_msg_id())
	ec.Bytes(t.Get_multi_media())

	return ec.GetBuffer()
}

func (t *TL_messages_sendMultiMedia) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._silent = dc.Bytes(16)
	t._background = dc.Bytes(16)
	t._clear_draft = dc.Bytes(16)
	t._peer = dc.Bytes(16)
	t._reply_to_msg_id = dc.Bytes(16)
	t._multi_media = dc.Bytes(16)

}

// messages_uploadEncryptedFile#5057c497
type TL_messages_uploadEncryptedFile struct {
	_peer []byte
	_file []byte
}

func (t *TL_messages_uploadEncryptedFile) Set_peer(_peer []byte) {
	t._peer = _peer
}

func (t *TL_messages_uploadEncryptedFile) Get_peer() []byte {
	return t._peer
}

func (t *TL_messages_uploadEncryptedFile) Set_file(_file []byte) {
	t._file = _file
}

func (t *TL_messages_uploadEncryptedFile) Get_file() []byte {
	return t._file
}

func New_TL_messages_uploadEncryptedFile() *TL_messages_uploadEncryptedFile {
	return &TL_messages_uploadEncryptedFile{}
}

func (t *TL_messages_uploadEncryptedFile) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_messages_uploadEncryptedFile))
	ec.Bytes(t.Get_peer())
	ec.Bytes(t.Get_file())

	return ec.GetBuffer()
}

func (t *TL_messages_uploadEncryptedFile) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._peer = dc.Bytes(16)
	t._file = dc.Bytes(16)

}

// updates_getState#edd4882a
type TL_updates_getState struct {
}

func New_TL_updates_getState() *TL_updates_getState {
	return &TL_updates_getState{}
}

func (t *TL_updates_getState) Encode() []byte {
	return nil
}

func (t *TL_updates_getState) Decode(b []byte) {

}

// updates_getDifference#25939651
type TL_updates_getDifference struct {
	_flags           []byte
	_pts             int32
	_pts_total_limit []byte
	_date            int32
	_qts             int32
}

func (t *TL_updates_getDifference) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_updates_getDifference) Get_flags() []byte {
	return t._flags
}

func (t *TL_updates_getDifference) Set_pts(_pts int32) {
	t._pts = _pts
}

func (t *TL_updates_getDifference) Get_pts() int32 {
	return t._pts
}

func (t *TL_updates_getDifference) Set_pts_total_limit(_pts_total_limit []byte) {
	t._pts_total_limit = _pts_total_limit
}

func (t *TL_updates_getDifference) Get_pts_total_limit() []byte {
	return t._pts_total_limit
}

func (t *TL_updates_getDifference) Set_date(_date int32) {
	t._date = _date
}

func (t *TL_updates_getDifference) Get_date() int32 {
	return t._date
}

func (t *TL_updates_getDifference) Set_qts(_qts int32) {
	t._qts = _qts
}

func (t *TL_updates_getDifference) Get_qts() int32 {
	return t._qts
}

func New_TL_updates_getDifference() *TL_updates_getDifference {
	return &TL_updates_getDifference{}
}

func (t *TL_updates_getDifference) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updates_getDifference))
	ec.Int(t.Get_pts())
	ec.Bytes(t.Get_pts_total_limit())
	ec.Int(t.Get_date())
	ec.Int(t.Get_qts())

	return ec.GetBuffer()
}

func (t *TL_updates_getDifference) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._pts = dc.Int()
	t._pts_total_limit = dc.Bytes(16)
	t._date = dc.Int()
	t._qts = dc.Int()

}

// updates_getChannelDifference#03173d78
type TL_updates_getChannelDifference struct {
	_flags   []byte
	_force   []byte
	_channel []byte
	_filter  []byte
	_pts     int32
	_limit   int32
}

func (t *TL_updates_getChannelDifference) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_updates_getChannelDifference) Get_flags() []byte {
	return t._flags
}

func (t *TL_updates_getChannelDifference) Set_force(_force []byte) {
	t._force = _force
}

func (t *TL_updates_getChannelDifference) Get_force() []byte {
	return t._force
}

func (t *TL_updates_getChannelDifference) Set_channel(_channel []byte) {
	t._channel = _channel
}

func (t *TL_updates_getChannelDifference) Get_channel() []byte {
	return t._channel
}

func (t *TL_updates_getChannelDifference) Set_filter(_filter []byte) {
	t._filter = _filter
}

func (t *TL_updates_getChannelDifference) Get_filter() []byte {
	return t._filter
}

func (t *TL_updates_getChannelDifference) Set_pts(_pts int32) {
	t._pts = _pts
}

func (t *TL_updates_getChannelDifference) Get_pts() int32 {
	return t._pts
}

func (t *TL_updates_getChannelDifference) Set_limit(_limit int32) {
	t._limit = _limit
}

func (t *TL_updates_getChannelDifference) Get_limit() int32 {
	return t._limit
}

func New_TL_updates_getChannelDifference() *TL_updates_getChannelDifference {
	return &TL_updates_getChannelDifference{}
}

func (t *TL_updates_getChannelDifference) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_updates_getChannelDifference))
	ec.Bytes(t.Get_force())
	ec.Bytes(t.Get_channel())
	ec.Bytes(t.Get_filter())
	ec.Int(t.Get_pts())
	ec.Int(t.Get_limit())

	return ec.GetBuffer()
}

func (t *TL_updates_getChannelDifference) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._force = dc.Bytes(16)
	t._channel = dc.Bytes(16)
	t._filter = dc.Bytes(16)
	t._pts = dc.Int()
	t._limit = dc.Int()

}

// photos_updateProfilePhoto#f0bb5152
type TL_photos_updateProfilePhoto struct {
	_id []byte
}

func (t *TL_photos_updateProfilePhoto) Set_id(_id []byte) {
	t._id = _id
}

func (t *TL_photos_updateProfilePhoto) Get_id() []byte {
	return t._id
}

func New_TL_photos_updateProfilePhoto() *TL_photos_updateProfilePhoto {
	return &TL_photos_updateProfilePhoto{}
}

func (t *TL_photos_updateProfilePhoto) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_photos_updateProfilePhoto))
	ec.Bytes(t.Get_id())

	return ec.GetBuffer()
}

func (t *TL_photos_updateProfilePhoto) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Bytes(16)

}

// photos_uploadProfilePhoto#4f32c098
type TL_photos_uploadProfilePhoto struct {
	_file []byte
}

func (t *TL_photos_uploadProfilePhoto) Set_file(_file []byte) {
	t._file = _file
}

func (t *TL_photos_uploadProfilePhoto) Get_file() []byte {
	return t._file
}

func New_TL_photos_uploadProfilePhoto() *TL_photos_uploadProfilePhoto {
	return &TL_photos_uploadProfilePhoto{}
}

func (t *TL_photos_uploadProfilePhoto) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_photos_uploadProfilePhoto))
	ec.Bytes(t.Get_file())

	return ec.GetBuffer()
}

func (t *TL_photos_uploadProfilePhoto) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._file = dc.Bytes(16)

}

// photos_deletePhotos#87cf7f2f
type TL_photos_deletePhotos struct {
	_id []byte
}

func (t *TL_photos_deletePhotos) Set_id(_id []byte) {
	t._id = _id
}

func (t *TL_photos_deletePhotos) Get_id() []byte {
	return t._id
}

func New_TL_photos_deletePhotos() *TL_photos_deletePhotos {
	return &TL_photos_deletePhotos{}
}

func (t *TL_photos_deletePhotos) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_photos_deletePhotos))
	ec.Bytes(t.Get_id())

	return ec.GetBuffer()
}

func (t *TL_photos_deletePhotos) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Bytes(16)

}

// photos_getUserPhotos#91cd32a8
type TL_photos_getUserPhotos struct {
	_user_id []byte
	_offset  int32
	_max_id  int64
	_limit   int32
}

func (t *TL_photos_getUserPhotos) Set_user_id(_user_id []byte) {
	t._user_id = _user_id
}

func (t *TL_photos_getUserPhotos) Get_user_id() []byte {
	return t._user_id
}

func (t *TL_photos_getUserPhotos) Set_offset(_offset int32) {
	t._offset = _offset
}

func (t *TL_photos_getUserPhotos) Get_offset() int32 {
	return t._offset
}

func (t *TL_photos_getUserPhotos) Set_max_id(_max_id int64) {
	t._max_id = _max_id
}

func (t *TL_photos_getUserPhotos) Get_max_id() int64 {
	return t._max_id
}

func (t *TL_photos_getUserPhotos) Set_limit(_limit int32) {
	t._limit = _limit
}

func (t *TL_photos_getUserPhotos) Get_limit() int32 {
	return t._limit
}

func New_TL_photos_getUserPhotos() *TL_photos_getUserPhotos {
	return &TL_photos_getUserPhotos{}
}

func (t *TL_photos_getUserPhotos) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_photos_getUserPhotos))
	ec.Bytes(t.Get_user_id())
	ec.Int(t.Get_offset())
	ec.Long(t.Get_max_id())
	ec.Int(t.Get_limit())

	return ec.GetBuffer()
}

func (t *TL_photos_getUserPhotos) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._user_id = dc.Bytes(16)
	t._offset = dc.Int()
	t._max_id = dc.Long()
	t._limit = dc.Int()

}

// upload_saveFilePart#b304a621
type TL_upload_saveFilePart struct {
	_file_id   int64
	_file_part int32
	_bytes     []byte
}

func (t *TL_upload_saveFilePart) Set_file_id(_file_id int64) {
	t._file_id = _file_id
}

func (t *TL_upload_saveFilePart) Get_file_id() int64 {
	return t._file_id
}

func (t *TL_upload_saveFilePart) Set_file_part(_file_part int32) {
	t._file_part = _file_part
}

func (t *TL_upload_saveFilePart) Get_file_part() int32 {
	return t._file_part
}

func (t *TL_upload_saveFilePart) Set_bytes(_bytes []byte) {
	t._bytes = _bytes
}

func (t *TL_upload_saveFilePart) Get_bytes() []byte {
	return t._bytes
}

func New_TL_upload_saveFilePart() *TL_upload_saveFilePart {
	return &TL_upload_saveFilePart{}
}

func (t *TL_upload_saveFilePart) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_upload_saveFilePart))
	ec.Long(t.Get_file_id())
	ec.Int(t.Get_file_part())
	ec.Bytes(t.Get_bytes())

	return ec.GetBuffer()
}

func (t *TL_upload_saveFilePart) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._file_id = dc.Long()
	t._file_part = dc.Int()
	t._bytes = dc.Bytes(16)

}

// upload_getFile#e3a6cfb5
type TL_upload_getFile struct {
	_location []byte
	_offset   int32
	_limit    int32
}

func (t *TL_upload_getFile) Set_location(_location []byte) {
	t._location = _location
}

func (t *TL_upload_getFile) Get_location() []byte {
	return t._location
}

func (t *TL_upload_getFile) Set_offset(_offset int32) {
	t._offset = _offset
}

func (t *TL_upload_getFile) Get_offset() int32 {
	return t._offset
}

func (t *TL_upload_getFile) Set_limit(_limit int32) {
	t._limit = _limit
}

func (t *TL_upload_getFile) Get_limit() int32 {
	return t._limit
}

func New_TL_upload_getFile() *TL_upload_getFile {
	return &TL_upload_getFile{}
}

func (t *TL_upload_getFile) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_upload_getFile))
	ec.Bytes(t.Get_location())
	ec.Int(t.Get_offset())
	ec.Int(t.Get_limit())

	return ec.GetBuffer()
}

func (t *TL_upload_getFile) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._location = dc.Bytes(16)
	t._offset = dc.Int()
	t._limit = dc.Int()

}

// upload_saveBigFilePart#de7b673d
type TL_upload_saveBigFilePart struct {
	_file_id          int64
	_file_part        int32
	_file_total_parts int32
	_bytes            []byte
}

func (t *TL_upload_saveBigFilePart) Set_file_id(_file_id int64) {
	t._file_id = _file_id
}

func (t *TL_upload_saveBigFilePart) Get_file_id() int64 {
	return t._file_id
}

func (t *TL_upload_saveBigFilePart) Set_file_part(_file_part int32) {
	t._file_part = _file_part
}

func (t *TL_upload_saveBigFilePart) Get_file_part() int32 {
	return t._file_part
}

func (t *TL_upload_saveBigFilePart) Set_file_total_parts(_file_total_parts int32) {
	t._file_total_parts = _file_total_parts
}

func (t *TL_upload_saveBigFilePart) Get_file_total_parts() int32 {
	return t._file_total_parts
}

func (t *TL_upload_saveBigFilePart) Set_bytes(_bytes []byte) {
	t._bytes = _bytes
}

func (t *TL_upload_saveBigFilePart) Get_bytes() []byte {
	return t._bytes
}

func New_TL_upload_saveBigFilePart() *TL_upload_saveBigFilePart {
	return &TL_upload_saveBigFilePart{}
}

func (t *TL_upload_saveBigFilePart) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_upload_saveBigFilePart))
	ec.Long(t.Get_file_id())
	ec.Int(t.Get_file_part())
	ec.Int(t.Get_file_total_parts())
	ec.Bytes(t.Get_bytes())

	return ec.GetBuffer()
}

func (t *TL_upload_saveBigFilePart) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._file_id = dc.Long()
	t._file_part = dc.Int()
	t._file_total_parts = dc.Int()
	t._bytes = dc.Bytes(16)

}

// upload_getWebFile#24e6818d
type TL_upload_getWebFile struct {
	_location []byte
	_offset   int32
	_limit    int32
}

func (t *TL_upload_getWebFile) Set_location(_location []byte) {
	t._location = _location
}

func (t *TL_upload_getWebFile) Get_location() []byte {
	return t._location
}

func (t *TL_upload_getWebFile) Set_offset(_offset int32) {
	t._offset = _offset
}

func (t *TL_upload_getWebFile) Get_offset() int32 {
	return t._offset
}

func (t *TL_upload_getWebFile) Set_limit(_limit int32) {
	t._limit = _limit
}

func (t *TL_upload_getWebFile) Get_limit() int32 {
	return t._limit
}

func New_TL_upload_getWebFile() *TL_upload_getWebFile {
	return &TL_upload_getWebFile{}
}

func (t *TL_upload_getWebFile) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_upload_getWebFile))
	ec.Bytes(t.Get_location())
	ec.Int(t.Get_offset())
	ec.Int(t.Get_limit())

	return ec.GetBuffer()
}

func (t *TL_upload_getWebFile) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._location = dc.Bytes(16)
	t._offset = dc.Int()
	t._limit = dc.Int()

}

// upload_getCdnFile#2000bcc3
type TL_upload_getCdnFile struct {
	_file_token []byte
	_offset     int32
	_limit      int32
}

func (t *TL_upload_getCdnFile) Set_file_token(_file_token []byte) {
	t._file_token = _file_token
}

func (t *TL_upload_getCdnFile) Get_file_token() []byte {
	return t._file_token
}

func (t *TL_upload_getCdnFile) Set_offset(_offset int32) {
	t._offset = _offset
}

func (t *TL_upload_getCdnFile) Get_offset() int32 {
	return t._offset
}

func (t *TL_upload_getCdnFile) Set_limit(_limit int32) {
	t._limit = _limit
}

func (t *TL_upload_getCdnFile) Get_limit() int32 {
	return t._limit
}

func New_TL_upload_getCdnFile() *TL_upload_getCdnFile {
	return &TL_upload_getCdnFile{}
}

func (t *TL_upload_getCdnFile) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_upload_getCdnFile))
	ec.Bytes(t.Get_file_token())
	ec.Int(t.Get_offset())
	ec.Int(t.Get_limit())

	return ec.GetBuffer()
}

func (t *TL_upload_getCdnFile) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._file_token = dc.Bytes(16)
	t._offset = dc.Int()
	t._limit = dc.Int()

}

// upload_reuploadCdnFile#1af91c09
type TL_upload_reuploadCdnFile struct {
	_file_token    []byte
	_request_token []byte
}

func (t *TL_upload_reuploadCdnFile) Set_file_token(_file_token []byte) {
	t._file_token = _file_token
}

func (t *TL_upload_reuploadCdnFile) Get_file_token() []byte {
	return t._file_token
}

func (t *TL_upload_reuploadCdnFile) Set_request_token(_request_token []byte) {
	t._request_token = _request_token
}

func (t *TL_upload_reuploadCdnFile) Get_request_token() []byte {
	return t._request_token
}

func New_TL_upload_reuploadCdnFile() *TL_upload_reuploadCdnFile {
	return &TL_upload_reuploadCdnFile{}
}

func (t *TL_upload_reuploadCdnFile) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_upload_reuploadCdnFile))
	ec.Bytes(t.Get_file_token())
	ec.Bytes(t.Get_request_token())

	return ec.GetBuffer()
}

func (t *TL_upload_reuploadCdnFile) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._file_token = dc.Bytes(16)
	t._request_token = dc.Bytes(16)

}

// upload_getCdnFileHashes#f715c87b
type TL_upload_getCdnFileHashes struct {
	_file_token []byte
	_offset     int32
}

func (t *TL_upload_getCdnFileHashes) Set_file_token(_file_token []byte) {
	t._file_token = _file_token
}

func (t *TL_upload_getCdnFileHashes) Get_file_token() []byte {
	return t._file_token
}

func (t *TL_upload_getCdnFileHashes) Set_offset(_offset int32) {
	t._offset = _offset
}

func (t *TL_upload_getCdnFileHashes) Get_offset() int32 {
	return t._offset
}

func New_TL_upload_getCdnFileHashes() *TL_upload_getCdnFileHashes {
	return &TL_upload_getCdnFileHashes{}
}

func (t *TL_upload_getCdnFileHashes) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_upload_getCdnFileHashes))
	ec.Bytes(t.Get_file_token())
	ec.Int(t.Get_offset())

	return ec.GetBuffer()
}

func (t *TL_upload_getCdnFileHashes) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._file_token = dc.Bytes(16)
	t._offset = dc.Int()

}

// help_getConfig#c4f9186b
type TL_help_getConfig struct {
}

func New_TL_help_getConfig() *TL_help_getConfig {
	return &TL_help_getConfig{}
}

func (t *TL_help_getConfig) Encode() []byte {
	return nil
}

func (t *TL_help_getConfig) Decode(b []byte) {

}

// help_getNearestDc#1fb33026
type TL_help_getNearestDc struct {
}

func New_TL_help_getNearestDc() *TL_help_getNearestDc {
	return &TL_help_getNearestDc{}
}

func (t *TL_help_getNearestDc) Encode() []byte {
	return nil
}

func (t *TL_help_getNearestDc) Decode(b []byte) {

}

// help_getAppUpdate#ae2de196
type TL_help_getAppUpdate struct {
}

func New_TL_help_getAppUpdate() *TL_help_getAppUpdate {
	return &TL_help_getAppUpdate{}
}

func (t *TL_help_getAppUpdate) Encode() []byte {
	return nil
}

func (t *TL_help_getAppUpdate) Decode(b []byte) {

}

// help_saveAppLog#6f02f748
type TL_help_saveAppLog struct {
	_events []byte
}

func (t *TL_help_saveAppLog) Set_events(_events []byte) {
	t._events = _events
}

func (t *TL_help_saveAppLog) Get_events() []byte {
	return t._events
}

func New_TL_help_saveAppLog() *TL_help_saveAppLog {
	return &TL_help_saveAppLog{}
}

func (t *TL_help_saveAppLog) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_help_saveAppLog))
	ec.Bytes(t.Get_events())

	return ec.GetBuffer()
}

func (t *TL_help_saveAppLog) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._events = dc.Bytes(16)

}

// help_getInviteText#4d392343
type TL_help_getInviteText struct {
}

func New_TL_help_getInviteText() *TL_help_getInviteText {
	return &TL_help_getInviteText{}
}

func (t *TL_help_getInviteText) Encode() []byte {
	return nil
}

func (t *TL_help_getInviteText) Decode(b []byte) {

}

// help_getSupport#9cdf08cd
type TL_help_getSupport struct {
}

func New_TL_help_getSupport() *TL_help_getSupport {
	return &TL_help_getSupport{}
}

func (t *TL_help_getSupport) Encode() []byte {
	return nil
}

func (t *TL_help_getSupport) Decode(b []byte) {

}

// help_getAppChangelog#9010ef6f
type TL_help_getAppChangelog struct {
	_prev_app_version string
}

func (t *TL_help_getAppChangelog) Set_prev_app_version(_prev_app_version string) {
	t._prev_app_version = _prev_app_version
}

func (t *TL_help_getAppChangelog) Get_prev_app_version() string {
	return t._prev_app_version
}

func New_TL_help_getAppChangelog() *TL_help_getAppChangelog {
	return &TL_help_getAppChangelog{}
}

func (t *TL_help_getAppChangelog) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_help_getAppChangelog))
	ec.String(t.Get_prev_app_version())

	return ec.GetBuffer()
}

func (t *TL_help_getAppChangelog) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._prev_app_version = dc.String()

}

// help_getTermsOfService#350170f3
type TL_help_getTermsOfService struct {
}

func New_TL_help_getTermsOfService() *TL_help_getTermsOfService {
	return &TL_help_getTermsOfService{}
}

func (t *TL_help_getTermsOfService) Encode() []byte {
	return nil
}

func (t *TL_help_getTermsOfService) Decode(b []byte) {

}

// help_setBotUpdatesStatus#ec22cfcd
type TL_help_setBotUpdatesStatus struct {
	_pending_updates_count int32
	_message               string
}

func (t *TL_help_setBotUpdatesStatus) Set_pending_updates_count(_pending_updates_count int32) {
	t._pending_updates_count = _pending_updates_count
}

func (t *TL_help_setBotUpdatesStatus) Get_pending_updates_count() int32 {
	return t._pending_updates_count
}

func (t *TL_help_setBotUpdatesStatus) Set_message(_message string) {
	t._message = _message
}

func (t *TL_help_setBotUpdatesStatus) Get_message() string {
	return t._message
}

func New_TL_help_setBotUpdatesStatus() *TL_help_setBotUpdatesStatus {
	return &TL_help_setBotUpdatesStatus{}
}

func (t *TL_help_setBotUpdatesStatus) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_help_setBotUpdatesStatus))
	ec.Int(t.Get_pending_updates_count())
	ec.String(t.Get_message())

	return ec.GetBuffer()
}

func (t *TL_help_setBotUpdatesStatus) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._pending_updates_count = dc.Int()
	t._message = dc.String()

}

// help_getCdnConfig#52029342
type TL_help_getCdnConfig struct {
}

func New_TL_help_getCdnConfig() *TL_help_getCdnConfig {
	return &TL_help_getCdnConfig{}
}

func (t *TL_help_getCdnConfig) Encode() []byte {
	return nil
}

func (t *TL_help_getCdnConfig) Decode(b []byte) {

}

// help_getRecentMeUrls#3dc0f114
type TL_help_getRecentMeUrls struct {
	_referer string
}

func (t *TL_help_getRecentMeUrls) Set_referer(_referer string) {
	t._referer = _referer
}

func (t *TL_help_getRecentMeUrls) Get_referer() string {
	return t._referer
}

func New_TL_help_getRecentMeUrls() *TL_help_getRecentMeUrls {
	return &TL_help_getRecentMeUrls{}
}

func (t *TL_help_getRecentMeUrls) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_help_getRecentMeUrls))
	ec.String(t.Get_referer())

	return ec.GetBuffer()
}

func (t *TL_help_getRecentMeUrls) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._referer = dc.String()

}

// channels_readHistory#cc104937
type TL_channels_readHistory struct {
	_channel []byte
	_max_id  int32
}

func (t *TL_channels_readHistory) Set_channel(_channel []byte) {
	t._channel = _channel
}

func (t *TL_channels_readHistory) Get_channel() []byte {
	return t._channel
}

func (t *TL_channels_readHistory) Set_max_id(_max_id int32) {
	t._max_id = _max_id
}

func (t *TL_channels_readHistory) Get_max_id() int32 {
	return t._max_id
}

func New_TL_channels_readHistory() *TL_channels_readHistory {
	return &TL_channels_readHistory{}
}

func (t *TL_channels_readHistory) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channels_readHistory))
	ec.Bytes(t.Get_channel())
	ec.Int(t.Get_max_id())

	return ec.GetBuffer()
}

func (t *TL_channels_readHistory) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._channel = dc.Bytes(16)
	t._max_id = dc.Int()

}

// channels_deleteMessages#84c1fd4e
type TL_channels_deleteMessages struct {
	_channel []byte
	_id      []byte
}

func (t *TL_channels_deleteMessages) Set_channel(_channel []byte) {
	t._channel = _channel
}

func (t *TL_channels_deleteMessages) Get_channel() []byte {
	return t._channel
}

func (t *TL_channels_deleteMessages) Set_id(_id []byte) {
	t._id = _id
}

func (t *TL_channels_deleteMessages) Get_id() []byte {
	return t._id
}

func New_TL_channels_deleteMessages() *TL_channels_deleteMessages {
	return &TL_channels_deleteMessages{}
}

func (t *TL_channels_deleteMessages) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channels_deleteMessages))
	ec.Bytes(t.Get_channel())
	ec.Bytes(t.Get_id())

	return ec.GetBuffer()
}

func (t *TL_channels_deleteMessages) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._channel = dc.Bytes(16)
	t._id = dc.Bytes(16)

}

// channels_deleteUserHistory#d10dd71b
type TL_channels_deleteUserHistory struct {
	_channel []byte
	_user_id []byte
}

func (t *TL_channels_deleteUserHistory) Set_channel(_channel []byte) {
	t._channel = _channel
}

func (t *TL_channels_deleteUserHistory) Get_channel() []byte {
	return t._channel
}

func (t *TL_channels_deleteUserHistory) Set_user_id(_user_id []byte) {
	t._user_id = _user_id
}

func (t *TL_channels_deleteUserHistory) Get_user_id() []byte {
	return t._user_id
}

func New_TL_channels_deleteUserHistory() *TL_channels_deleteUserHistory {
	return &TL_channels_deleteUserHistory{}
}

func (t *TL_channels_deleteUserHistory) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channels_deleteUserHistory))
	ec.Bytes(t.Get_channel())
	ec.Bytes(t.Get_user_id())

	return ec.GetBuffer()
}

func (t *TL_channels_deleteUserHistory) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._channel = dc.Bytes(16)
	t._user_id = dc.Bytes(16)

}

// channels_reportSpam#fe087810
type TL_channels_reportSpam struct {
	_channel []byte
	_user_id []byte
	_id      []byte
}

func (t *TL_channels_reportSpam) Set_channel(_channel []byte) {
	t._channel = _channel
}

func (t *TL_channels_reportSpam) Get_channel() []byte {
	return t._channel
}

func (t *TL_channels_reportSpam) Set_user_id(_user_id []byte) {
	t._user_id = _user_id
}

func (t *TL_channels_reportSpam) Get_user_id() []byte {
	return t._user_id
}

func (t *TL_channels_reportSpam) Set_id(_id []byte) {
	t._id = _id
}

func (t *TL_channels_reportSpam) Get_id() []byte {
	return t._id
}

func New_TL_channels_reportSpam() *TL_channels_reportSpam {
	return &TL_channels_reportSpam{}
}

func (t *TL_channels_reportSpam) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channels_reportSpam))
	ec.Bytes(t.Get_channel())
	ec.Bytes(t.Get_user_id())
	ec.Bytes(t.Get_id())

	return ec.GetBuffer()
}

func (t *TL_channels_reportSpam) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._channel = dc.Bytes(16)
	t._user_id = dc.Bytes(16)
	t._id = dc.Bytes(16)

}

// channels_getMessages#93d7b347
type TL_channels_getMessages struct {
	_channel []byte
	_id      []byte
}

func (t *TL_channels_getMessages) Set_channel(_channel []byte) {
	t._channel = _channel
}

func (t *TL_channels_getMessages) Get_channel() []byte {
	return t._channel
}

func (t *TL_channels_getMessages) Set_id(_id []byte) {
	t._id = _id
}

func (t *TL_channels_getMessages) Get_id() []byte {
	return t._id
}

func New_TL_channels_getMessages() *TL_channels_getMessages {
	return &TL_channels_getMessages{}
}

func (t *TL_channels_getMessages) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channels_getMessages))
	ec.Bytes(t.Get_channel())
	ec.Bytes(t.Get_id())

	return ec.GetBuffer()
}

func (t *TL_channels_getMessages) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._channel = dc.Bytes(16)
	t._id = dc.Bytes(16)

}

// channels_getParticipants#123e05e9
type TL_channels_getParticipants struct {
	_channel []byte
	_filter  []byte
	_offset  int32
	_limit   int32
	_hash    int32
}

func (t *TL_channels_getParticipants) Set_channel(_channel []byte) {
	t._channel = _channel
}

func (t *TL_channels_getParticipants) Get_channel() []byte {
	return t._channel
}

func (t *TL_channels_getParticipants) Set_filter(_filter []byte) {
	t._filter = _filter
}

func (t *TL_channels_getParticipants) Get_filter() []byte {
	return t._filter
}

func (t *TL_channels_getParticipants) Set_offset(_offset int32) {
	t._offset = _offset
}

func (t *TL_channels_getParticipants) Get_offset() int32 {
	return t._offset
}

func (t *TL_channels_getParticipants) Set_limit(_limit int32) {
	t._limit = _limit
}

func (t *TL_channels_getParticipants) Get_limit() int32 {
	return t._limit
}

func (t *TL_channels_getParticipants) Set_hash(_hash int32) {
	t._hash = _hash
}

func (t *TL_channels_getParticipants) Get_hash() int32 {
	return t._hash
}

func New_TL_channels_getParticipants() *TL_channels_getParticipants {
	return &TL_channels_getParticipants{}
}

func (t *TL_channels_getParticipants) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channels_getParticipants))
	ec.Bytes(t.Get_channel())
	ec.Bytes(t.Get_filter())
	ec.Int(t.Get_offset())
	ec.Int(t.Get_limit())
	ec.Int(t.Get_hash())

	return ec.GetBuffer()
}

func (t *TL_channels_getParticipants) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._channel = dc.Bytes(16)
	t._filter = dc.Bytes(16)
	t._offset = dc.Int()
	t._limit = dc.Int()
	t._hash = dc.Int()

}

// channels_getParticipant#546dd7a6
type TL_channels_getParticipant struct {
	_channel []byte
	_user_id []byte
}

func (t *TL_channels_getParticipant) Set_channel(_channel []byte) {
	t._channel = _channel
}

func (t *TL_channels_getParticipant) Get_channel() []byte {
	return t._channel
}

func (t *TL_channels_getParticipant) Set_user_id(_user_id []byte) {
	t._user_id = _user_id
}

func (t *TL_channels_getParticipant) Get_user_id() []byte {
	return t._user_id
}

func New_TL_channels_getParticipant() *TL_channels_getParticipant {
	return &TL_channels_getParticipant{}
}

func (t *TL_channels_getParticipant) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channels_getParticipant))
	ec.Bytes(t.Get_channel())
	ec.Bytes(t.Get_user_id())

	return ec.GetBuffer()
}

func (t *TL_channels_getParticipant) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._channel = dc.Bytes(16)
	t._user_id = dc.Bytes(16)

}

// channels_getChannels#0a7f6bbb
type TL_channels_getChannels struct {
	_id []byte
}

func (t *TL_channels_getChannels) Set_id(_id []byte) {
	t._id = _id
}

func (t *TL_channels_getChannels) Get_id() []byte {
	return t._id
}

func New_TL_channels_getChannels() *TL_channels_getChannels {
	return &TL_channels_getChannels{}
}

func (t *TL_channels_getChannels) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channels_getChannels))
	ec.Bytes(t.Get_id())

	return ec.GetBuffer()
}

func (t *TL_channels_getChannels) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._id = dc.Bytes(16)

}

// channels_getFullChannel#08736a09
type TL_channels_getFullChannel struct {
	_channel []byte
}

func (t *TL_channels_getFullChannel) Set_channel(_channel []byte) {
	t._channel = _channel
}

func (t *TL_channels_getFullChannel) Get_channel() []byte {
	return t._channel
}

func New_TL_channels_getFullChannel() *TL_channels_getFullChannel {
	return &TL_channels_getFullChannel{}
}

func (t *TL_channels_getFullChannel) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channels_getFullChannel))
	ec.Bytes(t.Get_channel())

	return ec.GetBuffer()
}

func (t *TL_channels_getFullChannel) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._channel = dc.Bytes(16)

}

// channels_createChannel#f4893d7f
type TL_channels_createChannel struct {
	_flags     []byte
	_broadcast []byte
	_megagroup []byte
	_title     string
	_about     string
}

func (t *TL_channels_createChannel) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_channels_createChannel) Get_flags() []byte {
	return t._flags
}

func (t *TL_channels_createChannel) Set_broadcast(_broadcast []byte) {
	t._broadcast = _broadcast
}

func (t *TL_channels_createChannel) Get_broadcast() []byte {
	return t._broadcast
}

func (t *TL_channels_createChannel) Set_megagroup(_megagroup []byte) {
	t._megagroup = _megagroup
}

func (t *TL_channels_createChannel) Get_megagroup() []byte {
	return t._megagroup
}

func (t *TL_channels_createChannel) Set_title(_title string) {
	t._title = _title
}

func (t *TL_channels_createChannel) Get_title() string {
	return t._title
}

func (t *TL_channels_createChannel) Set_about(_about string) {
	t._about = _about
}

func (t *TL_channels_createChannel) Get_about() string {
	return t._about
}

func New_TL_channels_createChannel() *TL_channels_createChannel {
	return &TL_channels_createChannel{}
}

func (t *TL_channels_createChannel) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channels_createChannel))
	ec.Bytes(t.Get_broadcast())
	ec.Bytes(t.Get_megagroup())
	ec.String(t.Get_title())
	ec.String(t.Get_about())

	return ec.GetBuffer()
}

func (t *TL_channels_createChannel) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._broadcast = dc.Bytes(16)
	t._megagroup = dc.Bytes(16)
	t._title = dc.String()
	t._about = dc.String()

}

// channels_editAbout#13e27f1e
type TL_channels_editAbout struct {
	_channel []byte
	_about   string
}

func (t *TL_channels_editAbout) Set_channel(_channel []byte) {
	t._channel = _channel
}

func (t *TL_channels_editAbout) Get_channel() []byte {
	return t._channel
}

func (t *TL_channels_editAbout) Set_about(_about string) {
	t._about = _about
}

func (t *TL_channels_editAbout) Get_about() string {
	return t._about
}

func New_TL_channels_editAbout() *TL_channels_editAbout {
	return &TL_channels_editAbout{}
}

func (t *TL_channels_editAbout) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channels_editAbout))
	ec.Bytes(t.Get_channel())
	ec.String(t.Get_about())

	return ec.GetBuffer()
}

func (t *TL_channels_editAbout) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._channel = dc.Bytes(16)
	t._about = dc.String()

}

// channels_editAdmin#20b88214
type TL_channels_editAdmin struct {
	_channel      []byte
	_user_id      []byte
	_admin_rights []byte
}

func (t *TL_channels_editAdmin) Set_channel(_channel []byte) {
	t._channel = _channel
}

func (t *TL_channels_editAdmin) Get_channel() []byte {
	return t._channel
}

func (t *TL_channels_editAdmin) Set_user_id(_user_id []byte) {
	t._user_id = _user_id
}

func (t *TL_channels_editAdmin) Get_user_id() []byte {
	return t._user_id
}

func (t *TL_channels_editAdmin) Set_admin_rights(_admin_rights []byte) {
	t._admin_rights = _admin_rights
}

func (t *TL_channels_editAdmin) Get_admin_rights() []byte {
	return t._admin_rights
}

func New_TL_channels_editAdmin() *TL_channels_editAdmin {
	return &TL_channels_editAdmin{}
}

func (t *TL_channels_editAdmin) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channels_editAdmin))
	ec.Bytes(t.Get_channel())
	ec.Bytes(t.Get_user_id())
	ec.Bytes(t.Get_admin_rights())

	return ec.GetBuffer()
}

func (t *TL_channels_editAdmin) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._channel = dc.Bytes(16)
	t._user_id = dc.Bytes(16)
	t._admin_rights = dc.Bytes(16)

}

// channels_editTitle#566decd0
type TL_channels_editTitle struct {
	_channel []byte
	_title   string
}

func (t *TL_channels_editTitle) Set_channel(_channel []byte) {
	t._channel = _channel
}

func (t *TL_channels_editTitle) Get_channel() []byte {
	return t._channel
}

func (t *TL_channels_editTitle) Set_title(_title string) {
	t._title = _title
}

func (t *TL_channels_editTitle) Get_title() string {
	return t._title
}

func New_TL_channels_editTitle() *TL_channels_editTitle {
	return &TL_channels_editTitle{}
}

func (t *TL_channels_editTitle) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channels_editTitle))
	ec.Bytes(t.Get_channel())
	ec.String(t.Get_title())

	return ec.GetBuffer()
}

func (t *TL_channels_editTitle) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._channel = dc.Bytes(16)
	t._title = dc.String()

}

// channels_editPhoto#f12e57c9
type TL_channels_editPhoto struct {
	_channel []byte
	_photo   []byte
}

func (t *TL_channels_editPhoto) Set_channel(_channel []byte) {
	t._channel = _channel
}

func (t *TL_channels_editPhoto) Get_channel() []byte {
	return t._channel
}

func (t *TL_channels_editPhoto) Set_photo(_photo []byte) {
	t._photo = _photo
}

func (t *TL_channels_editPhoto) Get_photo() []byte {
	return t._photo
}

func New_TL_channels_editPhoto() *TL_channels_editPhoto {
	return &TL_channels_editPhoto{}
}

func (t *TL_channels_editPhoto) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channels_editPhoto))
	ec.Bytes(t.Get_channel())
	ec.Bytes(t.Get_photo())

	return ec.GetBuffer()
}

func (t *TL_channels_editPhoto) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._channel = dc.Bytes(16)
	t._photo = dc.Bytes(16)

}

// channels_checkUsername#10e6bd2c
type TL_channels_checkUsername struct {
	_channel  []byte
	_username string
}

func (t *TL_channels_checkUsername) Set_channel(_channel []byte) {
	t._channel = _channel
}

func (t *TL_channels_checkUsername) Get_channel() []byte {
	return t._channel
}

func (t *TL_channels_checkUsername) Set_username(_username string) {
	t._username = _username
}

func (t *TL_channels_checkUsername) Get_username() string {
	return t._username
}

func New_TL_channels_checkUsername() *TL_channels_checkUsername {
	return &TL_channels_checkUsername{}
}

func (t *TL_channels_checkUsername) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channels_checkUsername))
	ec.Bytes(t.Get_channel())
	ec.String(t.Get_username())

	return ec.GetBuffer()
}

func (t *TL_channels_checkUsername) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._channel = dc.Bytes(16)
	t._username = dc.String()

}

// channels_updateUsername#3514b3de
type TL_channels_updateUsername struct {
	_channel  []byte
	_username string
}

func (t *TL_channels_updateUsername) Set_channel(_channel []byte) {
	t._channel = _channel
}

func (t *TL_channels_updateUsername) Get_channel() []byte {
	return t._channel
}

func (t *TL_channels_updateUsername) Set_username(_username string) {
	t._username = _username
}

func (t *TL_channels_updateUsername) Get_username() string {
	return t._username
}

func New_TL_channels_updateUsername() *TL_channels_updateUsername {
	return &TL_channels_updateUsername{}
}

func (t *TL_channels_updateUsername) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channels_updateUsername))
	ec.Bytes(t.Get_channel())
	ec.String(t.Get_username())

	return ec.GetBuffer()
}

func (t *TL_channels_updateUsername) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._channel = dc.Bytes(16)
	t._username = dc.String()

}

// channels_joinChannel#24b524c5
type TL_channels_joinChannel struct {
	_channel []byte
}

func (t *TL_channels_joinChannel) Set_channel(_channel []byte) {
	t._channel = _channel
}

func (t *TL_channels_joinChannel) Get_channel() []byte {
	return t._channel
}

func New_TL_channels_joinChannel() *TL_channels_joinChannel {
	return &TL_channels_joinChannel{}
}

func (t *TL_channels_joinChannel) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channels_joinChannel))
	ec.Bytes(t.Get_channel())

	return ec.GetBuffer()
}

func (t *TL_channels_joinChannel) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._channel = dc.Bytes(16)

}

// channels_leaveChannel#f836aa95
type TL_channels_leaveChannel struct {
	_channel []byte
}

func (t *TL_channels_leaveChannel) Set_channel(_channel []byte) {
	t._channel = _channel
}

func (t *TL_channels_leaveChannel) Get_channel() []byte {
	return t._channel
}

func New_TL_channels_leaveChannel() *TL_channels_leaveChannel {
	return &TL_channels_leaveChannel{}
}

func (t *TL_channels_leaveChannel) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channels_leaveChannel))
	ec.Bytes(t.Get_channel())

	return ec.GetBuffer()
}

func (t *TL_channels_leaveChannel) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._channel = dc.Bytes(16)

}

// channels_inviteToChannel#199f3a6c
type TL_channels_inviteToChannel struct {
	_channel []byte
	_users   []byte
}

func (t *TL_channels_inviteToChannel) Set_channel(_channel []byte) {
	t._channel = _channel
}

func (t *TL_channels_inviteToChannel) Get_channel() []byte {
	return t._channel
}

func (t *TL_channels_inviteToChannel) Set_users(_users []byte) {
	t._users = _users
}

func (t *TL_channels_inviteToChannel) Get_users() []byte {
	return t._users
}

func New_TL_channels_inviteToChannel() *TL_channels_inviteToChannel {
	return &TL_channels_inviteToChannel{}
}

func (t *TL_channels_inviteToChannel) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channels_inviteToChannel))
	ec.Bytes(t.Get_channel())
	ec.Bytes(t.Get_users())

	return ec.GetBuffer()
}

func (t *TL_channels_inviteToChannel) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._channel = dc.Bytes(16)
	t._users = dc.Bytes(16)

}

// channels_exportInvite#c7560885
type TL_channels_exportInvite struct {
	_channel []byte
}

func (t *TL_channels_exportInvite) Set_channel(_channel []byte) {
	t._channel = _channel
}

func (t *TL_channels_exportInvite) Get_channel() []byte {
	return t._channel
}

func New_TL_channels_exportInvite() *TL_channels_exportInvite {
	return &TL_channels_exportInvite{}
}

func (t *TL_channels_exportInvite) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channels_exportInvite))
	ec.Bytes(t.Get_channel())

	return ec.GetBuffer()
}

func (t *TL_channels_exportInvite) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._channel = dc.Bytes(16)

}

// channels_deleteChannel#c0111fe3
type TL_channels_deleteChannel struct {
	_channel []byte
}

func (t *TL_channels_deleteChannel) Set_channel(_channel []byte) {
	t._channel = _channel
}

func (t *TL_channels_deleteChannel) Get_channel() []byte {
	return t._channel
}

func New_TL_channels_deleteChannel() *TL_channels_deleteChannel {
	return &TL_channels_deleteChannel{}
}

func (t *TL_channels_deleteChannel) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channels_deleteChannel))
	ec.Bytes(t.Get_channel())

	return ec.GetBuffer()
}

func (t *TL_channels_deleteChannel) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._channel = dc.Bytes(16)

}

// channels_toggleInvites#49609307
type TL_channels_toggleInvites struct {
	_channel []byte
	_enabled bool
}

func (t *TL_channels_toggleInvites) Set_channel(_channel []byte) {
	t._channel = _channel
}

func (t *TL_channels_toggleInvites) Get_channel() []byte {
	return t._channel
}

func (t *TL_channels_toggleInvites) Set_enabled(_enabled bool) {
	t._enabled = _enabled
}

func (t *TL_channels_toggleInvites) Get_enabled() bool {
	return t._enabled
}

func New_TL_channels_toggleInvites() *TL_channels_toggleInvites {
	return &TL_channels_toggleInvites{}
}

func (t *TL_channels_toggleInvites) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channels_toggleInvites))
	ec.Bytes(t.Get_channel())
	ec.Bool(t.Get_enabled())

	return ec.GetBuffer()
}

func (t *TL_channels_toggleInvites) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._channel = dc.Bytes(16)
	t._enabled = dc.Bool()

}

// channels_exportMessageLink#c846d22d
type TL_channels_exportMessageLink struct {
	_channel []byte
	_id      int32
}

func (t *TL_channels_exportMessageLink) Set_channel(_channel []byte) {
	t._channel = _channel
}

func (t *TL_channels_exportMessageLink) Get_channel() []byte {
	return t._channel
}

func (t *TL_channels_exportMessageLink) Set_id(_id int32) {
	t._id = _id
}

func (t *TL_channels_exportMessageLink) Get_id() int32 {
	return t._id
}

func New_TL_channels_exportMessageLink() *TL_channels_exportMessageLink {
	return &TL_channels_exportMessageLink{}
}

func (t *TL_channels_exportMessageLink) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channels_exportMessageLink))
	ec.Bytes(t.Get_channel())
	ec.Int(t.Get_id())

	return ec.GetBuffer()
}

func (t *TL_channels_exportMessageLink) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._channel = dc.Bytes(16)
	t._id = dc.Int()

}

// channels_toggleSignatures#1f69b606
type TL_channels_toggleSignatures struct {
	_channel []byte
	_enabled bool
}

func (t *TL_channels_toggleSignatures) Set_channel(_channel []byte) {
	t._channel = _channel
}

func (t *TL_channels_toggleSignatures) Get_channel() []byte {
	return t._channel
}

func (t *TL_channels_toggleSignatures) Set_enabled(_enabled bool) {
	t._enabled = _enabled
}

func (t *TL_channels_toggleSignatures) Get_enabled() bool {
	return t._enabled
}

func New_TL_channels_toggleSignatures() *TL_channels_toggleSignatures {
	return &TL_channels_toggleSignatures{}
}

func (t *TL_channels_toggleSignatures) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channels_toggleSignatures))
	ec.Bytes(t.Get_channel())
	ec.Bool(t.Get_enabled())

	return ec.GetBuffer()
}

func (t *TL_channels_toggleSignatures) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._channel = dc.Bytes(16)
	t._enabled = dc.Bool()

}

// channels_updatePinnedMessage#a72ded52
type TL_channels_updatePinnedMessage struct {
	_flags   []byte
	_silent  []byte
	_channel []byte
	_id      int32
}

func (t *TL_channels_updatePinnedMessage) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_channels_updatePinnedMessage) Get_flags() []byte {
	return t._flags
}

func (t *TL_channels_updatePinnedMessage) Set_silent(_silent []byte) {
	t._silent = _silent
}

func (t *TL_channels_updatePinnedMessage) Get_silent() []byte {
	return t._silent
}

func (t *TL_channels_updatePinnedMessage) Set_channel(_channel []byte) {
	t._channel = _channel
}

func (t *TL_channels_updatePinnedMessage) Get_channel() []byte {
	return t._channel
}

func (t *TL_channels_updatePinnedMessage) Set_id(_id int32) {
	t._id = _id
}

func (t *TL_channels_updatePinnedMessage) Get_id() int32 {
	return t._id
}

func New_TL_channels_updatePinnedMessage() *TL_channels_updatePinnedMessage {
	return &TL_channels_updatePinnedMessage{}
}

func (t *TL_channels_updatePinnedMessage) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channels_updatePinnedMessage))
	ec.Bytes(t.Get_silent())
	ec.Bytes(t.Get_channel())
	ec.Int(t.Get_id())

	return ec.GetBuffer()
}

func (t *TL_channels_updatePinnedMessage) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._silent = dc.Bytes(16)
	t._channel = dc.Bytes(16)
	t._id = dc.Int()

}

// channels_getAdminedPublicChannels#8d8d82d7
type TL_channels_getAdminedPublicChannels struct {
}

func New_TL_channels_getAdminedPublicChannels() *TL_channels_getAdminedPublicChannels {
	return &TL_channels_getAdminedPublicChannels{}
}

func (t *TL_channels_getAdminedPublicChannels) Encode() []byte {
	return nil
}

func (t *TL_channels_getAdminedPublicChannels) Decode(b []byte) {

}

// channels_editBanned#bfd915cd
type TL_channels_editBanned struct {
	_channel       []byte
	_user_id       []byte
	_banned_rights []byte
}

func (t *TL_channels_editBanned) Set_channel(_channel []byte) {
	t._channel = _channel
}

func (t *TL_channels_editBanned) Get_channel() []byte {
	return t._channel
}

func (t *TL_channels_editBanned) Set_user_id(_user_id []byte) {
	t._user_id = _user_id
}

func (t *TL_channels_editBanned) Get_user_id() []byte {
	return t._user_id
}

func (t *TL_channels_editBanned) Set_banned_rights(_banned_rights []byte) {
	t._banned_rights = _banned_rights
}

func (t *TL_channels_editBanned) Get_banned_rights() []byte {
	return t._banned_rights
}

func New_TL_channels_editBanned() *TL_channels_editBanned {
	return &TL_channels_editBanned{}
}

func (t *TL_channels_editBanned) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channels_editBanned))
	ec.Bytes(t.Get_channel())
	ec.Bytes(t.Get_user_id())
	ec.Bytes(t.Get_banned_rights())

	return ec.GetBuffer()
}

func (t *TL_channels_editBanned) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._channel = dc.Bytes(16)
	t._user_id = dc.Bytes(16)
	t._banned_rights = dc.Bytes(16)

}

// channels_getAdminLog#33ddf480
type TL_channels_getAdminLog struct {
	_flags         []byte
	_channel       []byte
	_q             string
	_events_filter []byte
	_admins        []byte
	_max_id        int64
	_min_id        int64
	_limit         int32
}

func (t *TL_channels_getAdminLog) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_channels_getAdminLog) Get_flags() []byte {
	return t._flags
}

func (t *TL_channels_getAdminLog) Set_channel(_channel []byte) {
	t._channel = _channel
}

func (t *TL_channels_getAdminLog) Get_channel() []byte {
	return t._channel
}

func (t *TL_channels_getAdminLog) Set_q(_q string) {
	t._q = _q
}

func (t *TL_channels_getAdminLog) Get_q() string {
	return t._q
}

func (t *TL_channels_getAdminLog) Set_events_filter(_events_filter []byte) {
	t._events_filter = _events_filter
}

func (t *TL_channels_getAdminLog) Get_events_filter() []byte {
	return t._events_filter
}

func (t *TL_channels_getAdminLog) Set_admins(_admins []byte) {
	t._admins = _admins
}

func (t *TL_channels_getAdminLog) Get_admins() []byte {
	return t._admins
}

func (t *TL_channels_getAdminLog) Set_max_id(_max_id int64) {
	t._max_id = _max_id
}

func (t *TL_channels_getAdminLog) Get_max_id() int64 {
	return t._max_id
}

func (t *TL_channels_getAdminLog) Set_min_id(_min_id int64) {
	t._min_id = _min_id
}

func (t *TL_channels_getAdminLog) Get_min_id() int64 {
	return t._min_id
}

func (t *TL_channels_getAdminLog) Set_limit(_limit int32) {
	t._limit = _limit
}

func (t *TL_channels_getAdminLog) Get_limit() int32 {
	return t._limit
}

func New_TL_channels_getAdminLog() *TL_channels_getAdminLog {
	return &TL_channels_getAdminLog{}
}

func (t *TL_channels_getAdminLog) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channels_getAdminLog))
	ec.Bytes(t.Get_channel())
	ec.String(t.Get_q())
	ec.Bytes(t.Get_events_filter())
	ec.Bytes(t.Get_admins())
	ec.Long(t.Get_max_id())
	ec.Long(t.Get_min_id())
	ec.Int(t.Get_limit())

	return ec.GetBuffer()
}

func (t *TL_channels_getAdminLog) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._channel = dc.Bytes(16)
	t._q = dc.String()
	t._events_filter = dc.Bytes(16)
	t._admins = dc.Bytes(16)
	t._max_id = dc.Long()
	t._min_id = dc.Long()
	t._limit = dc.Int()

}

// channels_setStickers#ea8ca4f9
type TL_channels_setStickers struct {
	_channel    []byte
	_stickerset []byte
}

func (t *TL_channels_setStickers) Set_channel(_channel []byte) {
	t._channel = _channel
}

func (t *TL_channels_setStickers) Get_channel() []byte {
	return t._channel
}

func (t *TL_channels_setStickers) Set_stickerset(_stickerset []byte) {
	t._stickerset = _stickerset
}

func (t *TL_channels_setStickers) Get_stickerset() []byte {
	return t._stickerset
}

func New_TL_channels_setStickers() *TL_channels_setStickers {
	return &TL_channels_setStickers{}
}

func (t *TL_channels_setStickers) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channels_setStickers))
	ec.Bytes(t.Get_channel())
	ec.Bytes(t.Get_stickerset())

	return ec.GetBuffer()
}

func (t *TL_channels_setStickers) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._channel = dc.Bytes(16)
	t._stickerset = dc.Bytes(16)

}

// channels_readMessageContents#eab5dc38
type TL_channels_readMessageContents struct {
	_channel []byte
	_id      []byte
}

func (t *TL_channels_readMessageContents) Set_channel(_channel []byte) {
	t._channel = _channel
}

func (t *TL_channels_readMessageContents) Get_channel() []byte {
	return t._channel
}

func (t *TL_channels_readMessageContents) Set_id(_id []byte) {
	t._id = _id
}

func (t *TL_channels_readMessageContents) Get_id() []byte {
	return t._id
}

func New_TL_channels_readMessageContents() *TL_channels_readMessageContents {
	return &TL_channels_readMessageContents{}
}

func (t *TL_channels_readMessageContents) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channels_readMessageContents))
	ec.Bytes(t.Get_channel())
	ec.Bytes(t.Get_id())

	return ec.GetBuffer()
}

func (t *TL_channels_readMessageContents) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._channel = dc.Bytes(16)
	t._id = dc.Bytes(16)

}

// channels_deleteHistory#af369d42
type TL_channels_deleteHistory struct {
	_channel []byte
	_max_id  int32
}

func (t *TL_channels_deleteHistory) Set_channel(_channel []byte) {
	t._channel = _channel
}

func (t *TL_channels_deleteHistory) Get_channel() []byte {
	return t._channel
}

func (t *TL_channels_deleteHistory) Set_max_id(_max_id int32) {
	t._max_id = _max_id
}

func (t *TL_channels_deleteHistory) Get_max_id() int32 {
	return t._max_id
}

func New_TL_channels_deleteHistory() *TL_channels_deleteHistory {
	return &TL_channels_deleteHistory{}
}

func (t *TL_channels_deleteHistory) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channels_deleteHistory))
	ec.Bytes(t.Get_channel())
	ec.Int(t.Get_max_id())

	return ec.GetBuffer()
}

func (t *TL_channels_deleteHistory) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._channel = dc.Bytes(16)
	t._max_id = dc.Int()

}

// channels_togglePreHistoryHidden#eabbb94c
type TL_channels_togglePreHistoryHidden struct {
	_channel []byte
	_enabled bool
}

func (t *TL_channels_togglePreHistoryHidden) Set_channel(_channel []byte) {
	t._channel = _channel
}

func (t *TL_channels_togglePreHistoryHidden) Get_channel() []byte {
	return t._channel
}

func (t *TL_channels_togglePreHistoryHidden) Set_enabled(_enabled bool) {
	t._enabled = _enabled
}

func (t *TL_channels_togglePreHistoryHidden) Get_enabled() bool {
	return t._enabled
}

func New_TL_channels_togglePreHistoryHidden() *TL_channels_togglePreHistoryHidden {
	return &TL_channels_togglePreHistoryHidden{}
}

func (t *TL_channels_togglePreHistoryHidden) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_channels_togglePreHistoryHidden))
	ec.Bytes(t.Get_channel())
	ec.Bool(t.Get_enabled())

	return ec.GetBuffer()
}

func (t *TL_channels_togglePreHistoryHidden) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._channel = dc.Bytes(16)
	t._enabled = dc.Bool()

}

// bots_sendCustomRequest#aa2769ed
type TL_bots_sendCustomRequest struct {
	_custom_method string
	_params        []byte
}

func (t *TL_bots_sendCustomRequest) Set_custom_method(_custom_method string) {
	t._custom_method = _custom_method
}

func (t *TL_bots_sendCustomRequest) Get_custom_method() string {
	return t._custom_method
}

func (t *TL_bots_sendCustomRequest) Set_params(_params []byte) {
	t._params = _params
}

func (t *TL_bots_sendCustomRequest) Get_params() []byte {
	return t._params
}

func New_TL_bots_sendCustomRequest() *TL_bots_sendCustomRequest {
	return &TL_bots_sendCustomRequest{}
}

func (t *TL_bots_sendCustomRequest) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_bots_sendCustomRequest))
	ec.String(t.Get_custom_method())
	ec.Bytes(t.Get_params())

	return ec.GetBuffer()
}

func (t *TL_bots_sendCustomRequest) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._custom_method = dc.String()
	t._params = dc.Bytes(16)

}

// bots_answerWebhookJSONQuery#e6213f4d
type TL_bots_answerWebhookJSONQuery struct {
	_query_id int64
	_data     []byte
}

func (t *TL_bots_answerWebhookJSONQuery) Set_query_id(_query_id int64) {
	t._query_id = _query_id
}

func (t *TL_bots_answerWebhookJSONQuery) Get_query_id() int64 {
	return t._query_id
}

func (t *TL_bots_answerWebhookJSONQuery) Set_data(_data []byte) {
	t._data = _data
}

func (t *TL_bots_answerWebhookJSONQuery) Get_data() []byte {
	return t._data
}

func New_TL_bots_answerWebhookJSONQuery() *TL_bots_answerWebhookJSONQuery {
	return &TL_bots_answerWebhookJSONQuery{}
}

func (t *TL_bots_answerWebhookJSONQuery) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_bots_answerWebhookJSONQuery))
	ec.Long(t.Get_query_id())
	ec.Bytes(t.Get_data())

	return ec.GetBuffer()
}

func (t *TL_bots_answerWebhookJSONQuery) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._query_id = dc.Long()
	t._data = dc.Bytes(16)

}

// payments_getPaymentForm#99f09745
type TL_payments_getPaymentForm struct {
	_msg_id int32
}

func (t *TL_payments_getPaymentForm) Set_msg_id(_msg_id int32) {
	t._msg_id = _msg_id
}

func (t *TL_payments_getPaymentForm) Get_msg_id() int32 {
	return t._msg_id
}

func New_TL_payments_getPaymentForm() *TL_payments_getPaymentForm {
	return &TL_payments_getPaymentForm{}
}

func (t *TL_payments_getPaymentForm) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_payments_getPaymentForm))
	ec.Int(t.Get_msg_id())

	return ec.GetBuffer()
}

func (t *TL_payments_getPaymentForm) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._msg_id = dc.Int()

}

// payments_getPaymentReceipt#a092a980
type TL_payments_getPaymentReceipt struct {
	_msg_id int32
}

func (t *TL_payments_getPaymentReceipt) Set_msg_id(_msg_id int32) {
	t._msg_id = _msg_id
}

func (t *TL_payments_getPaymentReceipt) Get_msg_id() int32 {
	return t._msg_id
}

func New_TL_payments_getPaymentReceipt() *TL_payments_getPaymentReceipt {
	return &TL_payments_getPaymentReceipt{}
}

func (t *TL_payments_getPaymentReceipt) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_payments_getPaymentReceipt))
	ec.Int(t.Get_msg_id())

	return ec.GetBuffer()
}

func (t *TL_payments_getPaymentReceipt) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._msg_id = dc.Int()

}

// payments_validateRequestedInfo#770a8e74
type TL_payments_validateRequestedInfo struct {
	_flags  []byte
	_save   []byte
	_msg_id int32
	_info   []byte
}

func (t *TL_payments_validateRequestedInfo) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_payments_validateRequestedInfo) Get_flags() []byte {
	return t._flags
}

func (t *TL_payments_validateRequestedInfo) Set_save(_save []byte) {
	t._save = _save
}

func (t *TL_payments_validateRequestedInfo) Get_save() []byte {
	return t._save
}

func (t *TL_payments_validateRequestedInfo) Set_msg_id(_msg_id int32) {
	t._msg_id = _msg_id
}

func (t *TL_payments_validateRequestedInfo) Get_msg_id() int32 {
	return t._msg_id
}

func (t *TL_payments_validateRequestedInfo) Set_info(_info []byte) {
	t._info = _info
}

func (t *TL_payments_validateRequestedInfo) Get_info() []byte {
	return t._info
}

func New_TL_payments_validateRequestedInfo() *TL_payments_validateRequestedInfo {
	return &TL_payments_validateRequestedInfo{}
}

func (t *TL_payments_validateRequestedInfo) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_payments_validateRequestedInfo))
	ec.Bytes(t.Get_save())
	ec.Int(t.Get_msg_id())
	ec.Bytes(t.Get_info())

	return ec.GetBuffer()
}

func (t *TL_payments_validateRequestedInfo) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._save = dc.Bytes(16)
	t._msg_id = dc.Int()
	t._info = dc.Bytes(16)

}

// payments_sendPaymentForm#2b8879b3
type TL_payments_sendPaymentForm struct {
	_flags              []byte
	_msg_id             int32
	_requested_info_id  []byte
	_shipping_option_id []byte
	_credentials        []byte
}

func (t *TL_payments_sendPaymentForm) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_payments_sendPaymentForm) Get_flags() []byte {
	return t._flags
}

func (t *TL_payments_sendPaymentForm) Set_msg_id(_msg_id int32) {
	t._msg_id = _msg_id
}

func (t *TL_payments_sendPaymentForm) Get_msg_id() int32 {
	return t._msg_id
}

func (t *TL_payments_sendPaymentForm) Set_requested_info_id(_requested_info_id []byte) {
	t._requested_info_id = _requested_info_id
}

func (t *TL_payments_sendPaymentForm) Get_requested_info_id() []byte {
	return t._requested_info_id
}

func (t *TL_payments_sendPaymentForm) Set_shipping_option_id(_shipping_option_id []byte) {
	t._shipping_option_id = _shipping_option_id
}

func (t *TL_payments_sendPaymentForm) Get_shipping_option_id() []byte {
	return t._shipping_option_id
}

func (t *TL_payments_sendPaymentForm) Set_credentials(_credentials []byte) {
	t._credentials = _credentials
}

func (t *TL_payments_sendPaymentForm) Get_credentials() []byte {
	return t._credentials
}

func New_TL_payments_sendPaymentForm() *TL_payments_sendPaymentForm {
	return &TL_payments_sendPaymentForm{}
}

func (t *TL_payments_sendPaymentForm) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_payments_sendPaymentForm))
	ec.Int(t.Get_msg_id())
	ec.Bytes(t.Get_requested_info_id())
	ec.Bytes(t.Get_shipping_option_id())
	ec.Bytes(t.Get_credentials())

	return ec.GetBuffer()
}

func (t *TL_payments_sendPaymentForm) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._msg_id = dc.Int()
	t._requested_info_id = dc.Bytes(16)
	t._shipping_option_id = dc.Bytes(16)
	t._credentials = dc.Bytes(16)

}

// payments_getSavedInfo#227d824b
type TL_payments_getSavedInfo struct {
}

func New_TL_payments_getSavedInfo() *TL_payments_getSavedInfo {
	return &TL_payments_getSavedInfo{}
}

func (t *TL_payments_getSavedInfo) Encode() []byte {
	return nil
}

func (t *TL_payments_getSavedInfo) Decode(b []byte) {

}

// payments_clearSavedInfo#d83d70c1
type TL_payments_clearSavedInfo struct {
	_flags       []byte
	_credentials []byte
	_info        []byte
}

func (t *TL_payments_clearSavedInfo) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_payments_clearSavedInfo) Get_flags() []byte {
	return t._flags
}

func (t *TL_payments_clearSavedInfo) Set_credentials(_credentials []byte) {
	t._credentials = _credentials
}

func (t *TL_payments_clearSavedInfo) Get_credentials() []byte {
	return t._credentials
}

func (t *TL_payments_clearSavedInfo) Set_info(_info []byte) {
	t._info = _info
}

func (t *TL_payments_clearSavedInfo) Get_info() []byte {
	return t._info
}

func New_TL_payments_clearSavedInfo() *TL_payments_clearSavedInfo {
	return &TL_payments_clearSavedInfo{}
}

func (t *TL_payments_clearSavedInfo) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_payments_clearSavedInfo))
	ec.Bytes(t.Get_credentials())
	ec.Bytes(t.Get_info())

	return ec.GetBuffer()
}

func (t *TL_payments_clearSavedInfo) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._credentials = dc.Bytes(16)
	t._info = dc.Bytes(16)

}

// stickers_createStickerSet#9bd86e6a
type TL_stickers_createStickerSet struct {
	_flags      []byte
	_masks      []byte
	_user_id    []byte
	_title      string
	_short_name string
	_stickers   []byte
}

func (t *TL_stickers_createStickerSet) Set_flags(_flags []byte) {
	t._flags = _flags
}

func (t *TL_stickers_createStickerSet) Get_flags() []byte {
	return t._flags
}

func (t *TL_stickers_createStickerSet) Set_masks(_masks []byte) {
	t._masks = _masks
}

func (t *TL_stickers_createStickerSet) Get_masks() []byte {
	return t._masks
}

func (t *TL_stickers_createStickerSet) Set_user_id(_user_id []byte) {
	t._user_id = _user_id
}

func (t *TL_stickers_createStickerSet) Get_user_id() []byte {
	return t._user_id
}

func (t *TL_stickers_createStickerSet) Set_title(_title string) {
	t._title = _title
}

func (t *TL_stickers_createStickerSet) Get_title() string {
	return t._title
}

func (t *TL_stickers_createStickerSet) Set_short_name(_short_name string) {
	t._short_name = _short_name
}

func (t *TL_stickers_createStickerSet) Get_short_name() string {
	return t._short_name
}

func (t *TL_stickers_createStickerSet) Set_stickers(_stickers []byte) {
	t._stickers = _stickers
}

func (t *TL_stickers_createStickerSet) Get_stickers() []byte {
	return t._stickers
}

func New_TL_stickers_createStickerSet() *TL_stickers_createStickerSet {
	return &TL_stickers_createStickerSet{}
}

func (t *TL_stickers_createStickerSet) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_stickers_createStickerSet))
	ec.Bytes(t.Get_masks())
	ec.Bytes(t.Get_user_id())
	ec.String(t.Get_title())
	ec.String(t.Get_short_name())
	ec.Bytes(t.Get_stickers())

	return ec.GetBuffer()
}

func (t *TL_stickers_createStickerSet) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._masks = dc.Bytes(16)
	t._user_id = dc.Bytes(16)
	t._title = dc.String()
	t._short_name = dc.String()
	t._stickers = dc.Bytes(16)

}

// stickers_removeStickerFromSet#f7760f51
type TL_stickers_removeStickerFromSet struct {
	_sticker []byte
}

func (t *TL_stickers_removeStickerFromSet) Set_sticker(_sticker []byte) {
	t._sticker = _sticker
}

func (t *TL_stickers_removeStickerFromSet) Get_sticker() []byte {
	return t._sticker
}

func New_TL_stickers_removeStickerFromSet() *TL_stickers_removeStickerFromSet {
	return &TL_stickers_removeStickerFromSet{}
}

func (t *TL_stickers_removeStickerFromSet) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_stickers_removeStickerFromSet))
	ec.Bytes(t.Get_sticker())

	return ec.GetBuffer()
}

func (t *TL_stickers_removeStickerFromSet) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._sticker = dc.Bytes(16)

}

// stickers_changeStickerPosition#ffb6d4ca
type TL_stickers_changeStickerPosition struct {
	_sticker  []byte
	_position int32
}

func (t *TL_stickers_changeStickerPosition) Set_sticker(_sticker []byte) {
	t._sticker = _sticker
}

func (t *TL_stickers_changeStickerPosition) Get_sticker() []byte {
	return t._sticker
}

func (t *TL_stickers_changeStickerPosition) Set_position(_position int32) {
	t._position = _position
}

func (t *TL_stickers_changeStickerPosition) Get_position() int32 {
	return t._position
}

func New_TL_stickers_changeStickerPosition() *TL_stickers_changeStickerPosition {
	return &TL_stickers_changeStickerPosition{}
}

func (t *TL_stickers_changeStickerPosition) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_stickers_changeStickerPosition))
	ec.Bytes(t.Get_sticker())
	ec.Int(t.Get_position())

	return ec.GetBuffer()
}

func (t *TL_stickers_changeStickerPosition) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._sticker = dc.Bytes(16)
	t._position = dc.Int()

}

// stickers_addStickerToSet#8653febe
type TL_stickers_addStickerToSet struct {
	_stickerset []byte
	_sticker    []byte
}

func (t *TL_stickers_addStickerToSet) Set_stickerset(_stickerset []byte) {
	t._stickerset = _stickerset
}

func (t *TL_stickers_addStickerToSet) Get_stickerset() []byte {
	return t._stickerset
}

func (t *TL_stickers_addStickerToSet) Set_sticker(_sticker []byte) {
	t._sticker = _sticker
}

func (t *TL_stickers_addStickerToSet) Get_sticker() []byte {
	return t._sticker
}

func New_TL_stickers_addStickerToSet() *TL_stickers_addStickerToSet {
	return &TL_stickers_addStickerToSet{}
}

func (t *TL_stickers_addStickerToSet) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_stickers_addStickerToSet))
	ec.Bytes(t.Get_stickerset())
	ec.Bytes(t.Get_sticker())

	return ec.GetBuffer()
}

func (t *TL_stickers_addStickerToSet) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._stickerset = dc.Bytes(16)
	t._sticker = dc.Bytes(16)

}

// phone_getCallConfig#55451fa9
type TL_phone_getCallConfig struct {
}

func New_TL_phone_getCallConfig() *TL_phone_getCallConfig {
	return &TL_phone_getCallConfig{}
}

func (t *TL_phone_getCallConfig) Encode() []byte {
	return nil
}

func (t *TL_phone_getCallConfig) Decode(b []byte) {

}

// phone_requestCall#5b95b3d4
type TL_phone_requestCall struct {
	_user_id   []byte
	_random_id int32
	_g_a_hash  []byte
	_protocol  []byte
}

func (t *TL_phone_requestCall) Set_user_id(_user_id []byte) {
	t._user_id = _user_id
}

func (t *TL_phone_requestCall) Get_user_id() []byte {
	return t._user_id
}

func (t *TL_phone_requestCall) Set_random_id(_random_id int32) {
	t._random_id = _random_id
}

func (t *TL_phone_requestCall) Get_random_id() int32 {
	return t._random_id
}

func (t *TL_phone_requestCall) Set_g_a_hash(_g_a_hash []byte) {
	t._g_a_hash = _g_a_hash
}

func (t *TL_phone_requestCall) Get_g_a_hash() []byte {
	return t._g_a_hash
}

func (t *TL_phone_requestCall) Set_protocol(_protocol []byte) {
	t._protocol = _protocol
}

func (t *TL_phone_requestCall) Get_protocol() []byte {
	return t._protocol
}

func New_TL_phone_requestCall() *TL_phone_requestCall {
	return &TL_phone_requestCall{}
}

func (t *TL_phone_requestCall) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_phone_requestCall))
	ec.Bytes(t.Get_user_id())
	ec.Int(t.Get_random_id())
	ec.Bytes(t.Get_g_a_hash())
	ec.Bytes(t.Get_protocol())

	return ec.GetBuffer()
}

func (t *TL_phone_requestCall) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._user_id = dc.Bytes(16)
	t._random_id = dc.Int()
	t._g_a_hash = dc.Bytes(16)
	t._protocol = dc.Bytes(16)

}

// phone_acceptCall#3bd2b4a0
type TL_phone_acceptCall struct {
	_peer     []byte
	_g_b      []byte
	_protocol []byte
}

func (t *TL_phone_acceptCall) Set_peer(_peer []byte) {
	t._peer = _peer
}

func (t *TL_phone_acceptCall) Get_peer() []byte {
	return t._peer
}

func (t *TL_phone_acceptCall) Set_g_b(_g_b []byte) {
	t._g_b = _g_b
}

func (t *TL_phone_acceptCall) Get_g_b() []byte {
	return t._g_b
}

func (t *TL_phone_acceptCall) Set_protocol(_protocol []byte) {
	t._protocol = _protocol
}

func (t *TL_phone_acceptCall) Get_protocol() []byte {
	return t._protocol
}

func New_TL_phone_acceptCall() *TL_phone_acceptCall {
	return &TL_phone_acceptCall{}
}

func (t *TL_phone_acceptCall) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_phone_acceptCall))
	ec.Bytes(t.Get_peer())
	ec.Bytes(t.Get_g_b())
	ec.Bytes(t.Get_protocol())

	return ec.GetBuffer()
}

func (t *TL_phone_acceptCall) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._peer = dc.Bytes(16)
	t._g_b = dc.Bytes(16)
	t._protocol = dc.Bytes(16)

}

// phone_confirmCall#2efe1722
type TL_phone_confirmCall struct {
	_peer            []byte
	_g_a             []byte
	_key_fingerprint int64
	_protocol        []byte
}

func (t *TL_phone_confirmCall) Set_peer(_peer []byte) {
	t._peer = _peer
}

func (t *TL_phone_confirmCall) Get_peer() []byte {
	return t._peer
}

func (t *TL_phone_confirmCall) Set_g_a(_g_a []byte) {
	t._g_a = _g_a
}

func (t *TL_phone_confirmCall) Get_g_a() []byte {
	return t._g_a
}

func (t *TL_phone_confirmCall) Set_key_fingerprint(_key_fingerprint int64) {
	t._key_fingerprint = _key_fingerprint
}

func (t *TL_phone_confirmCall) Get_key_fingerprint() int64 {
	return t._key_fingerprint
}

func (t *TL_phone_confirmCall) Set_protocol(_protocol []byte) {
	t._protocol = _protocol
}

func (t *TL_phone_confirmCall) Get_protocol() []byte {
	return t._protocol
}

func New_TL_phone_confirmCall() *TL_phone_confirmCall {
	return &TL_phone_confirmCall{}
}

func (t *TL_phone_confirmCall) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_phone_confirmCall))
	ec.Bytes(t.Get_peer())
	ec.Bytes(t.Get_g_a())
	ec.Long(t.Get_key_fingerprint())
	ec.Bytes(t.Get_protocol())

	return ec.GetBuffer()
}

func (t *TL_phone_confirmCall) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._peer = dc.Bytes(16)
	t._g_a = dc.Bytes(16)
	t._key_fingerprint = dc.Long()
	t._protocol = dc.Bytes(16)

}

// phone_receivedCall#17d54f61
type TL_phone_receivedCall struct {
	_peer []byte
}

func (t *TL_phone_receivedCall) Set_peer(_peer []byte) {
	t._peer = _peer
}

func (t *TL_phone_receivedCall) Get_peer() []byte {
	return t._peer
}

func New_TL_phone_receivedCall() *TL_phone_receivedCall {
	return &TL_phone_receivedCall{}
}

func (t *TL_phone_receivedCall) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_phone_receivedCall))
	ec.Bytes(t.Get_peer())

	return ec.GetBuffer()
}

func (t *TL_phone_receivedCall) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._peer = dc.Bytes(16)

}

// phone_discardCall#78d413a6
type TL_phone_discardCall struct {
	_peer          []byte
	_duration      int32
	_reason        []byte
	_connection_id int64
}

func (t *TL_phone_discardCall) Set_peer(_peer []byte) {
	t._peer = _peer
}

func (t *TL_phone_discardCall) Get_peer() []byte {
	return t._peer
}

func (t *TL_phone_discardCall) Set_duration(_duration int32) {
	t._duration = _duration
}

func (t *TL_phone_discardCall) Get_duration() int32 {
	return t._duration
}

func (t *TL_phone_discardCall) Set_reason(_reason []byte) {
	t._reason = _reason
}

func (t *TL_phone_discardCall) Get_reason() []byte {
	return t._reason
}

func (t *TL_phone_discardCall) Set_connection_id(_connection_id int64) {
	t._connection_id = _connection_id
}

func (t *TL_phone_discardCall) Get_connection_id() int64 {
	return t._connection_id
}

func New_TL_phone_discardCall() *TL_phone_discardCall {
	return &TL_phone_discardCall{}
}

func (t *TL_phone_discardCall) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_phone_discardCall))
	ec.Bytes(t.Get_peer())
	ec.Int(t.Get_duration())
	ec.Bytes(t.Get_reason())
	ec.Long(t.Get_connection_id())

	return ec.GetBuffer()
}

func (t *TL_phone_discardCall) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._peer = dc.Bytes(16)
	t._duration = dc.Int()
	t._reason = dc.Bytes(16)
	t._connection_id = dc.Long()

}

// phone_setCallRating#1c536a34
type TL_phone_setCallRating struct {
	_peer    []byte
	_rating  int32
	_comment string
}

func (t *TL_phone_setCallRating) Set_peer(_peer []byte) {
	t._peer = _peer
}

func (t *TL_phone_setCallRating) Get_peer() []byte {
	return t._peer
}

func (t *TL_phone_setCallRating) Set_rating(_rating int32) {
	t._rating = _rating
}

func (t *TL_phone_setCallRating) Get_rating() int32 {
	return t._rating
}

func (t *TL_phone_setCallRating) Set_comment(_comment string) {
	t._comment = _comment
}

func (t *TL_phone_setCallRating) Get_comment() string {
	return t._comment
}

func New_TL_phone_setCallRating() *TL_phone_setCallRating {
	return &TL_phone_setCallRating{}
}

func (t *TL_phone_setCallRating) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_phone_setCallRating))
	ec.Bytes(t.Get_peer())
	ec.Int(t.Get_rating())
	ec.String(t.Get_comment())

	return ec.GetBuffer()
}

func (t *TL_phone_setCallRating) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._peer = dc.Bytes(16)
	t._rating = dc.Int()
	t._comment = dc.String()

}

// phone_saveCallDebug#277add7e
type TL_phone_saveCallDebug struct {
	_peer  []byte
	_debug []byte
}

func (t *TL_phone_saveCallDebug) Set_peer(_peer []byte) {
	t._peer = _peer
}

func (t *TL_phone_saveCallDebug) Get_peer() []byte {
	return t._peer
}

func (t *TL_phone_saveCallDebug) Set_debug(_debug []byte) {
	t._debug = _debug
}

func (t *TL_phone_saveCallDebug) Get_debug() []byte {
	return t._debug
}

func New_TL_phone_saveCallDebug() *TL_phone_saveCallDebug {
	return &TL_phone_saveCallDebug{}
}

func (t *TL_phone_saveCallDebug) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_phone_saveCallDebug))
	ec.Bytes(t.Get_peer())
	ec.Bytes(t.Get_debug())

	return ec.GetBuffer()
}

func (t *TL_phone_saveCallDebug) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._peer = dc.Bytes(16)
	t._debug = dc.Bytes(16)

}

// langpack_getLangPack#9ab5c58e
type TL_langpack_getLangPack struct {
	_lang_code string
}

func (t *TL_langpack_getLangPack) Set_lang_code(_lang_code string) {
	t._lang_code = _lang_code
}

func (t *TL_langpack_getLangPack) Get_lang_code() string {
	return t._lang_code
}

func New_TL_langpack_getLangPack() *TL_langpack_getLangPack {
	return &TL_langpack_getLangPack{}
}

func (t *TL_langpack_getLangPack) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_langpack_getLangPack))
	ec.String(t.Get_lang_code())

	return ec.GetBuffer()
}

func (t *TL_langpack_getLangPack) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._lang_code = dc.String()

}

// langpack_getStrings#2e1ee318
type TL_langpack_getStrings struct {
	_lang_code string
	_keys      []byte
}

func (t *TL_langpack_getStrings) Set_lang_code(_lang_code string) {
	t._lang_code = _lang_code
}

func (t *TL_langpack_getStrings) Get_lang_code() string {
	return t._lang_code
}

func (t *TL_langpack_getStrings) Set_keys(_keys []byte) {
	t._keys = _keys
}

func (t *TL_langpack_getStrings) Get_keys() []byte {
	return t._keys
}

func New_TL_langpack_getStrings() *TL_langpack_getStrings {
	return &TL_langpack_getStrings{}
}

func (t *TL_langpack_getStrings) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_langpack_getStrings))
	ec.String(t.Get_lang_code())
	ec.Bytes(t.Get_keys())

	return ec.GetBuffer()
}

func (t *TL_langpack_getStrings) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._lang_code = dc.String()
	t._keys = dc.Bytes(16)

}

// langpack_getDifference#0b2e4d7d
type TL_langpack_getDifference struct {
	_from_version int32
}

func (t *TL_langpack_getDifference) Set_from_version(_from_version int32) {
	t._from_version = _from_version
}

func (t *TL_langpack_getDifference) Get_from_version() int32 {
	return t._from_version
}

func New_TL_langpack_getDifference() *TL_langpack_getDifference {
	return &TL_langpack_getDifference{}
}

func (t *TL_langpack_getDifference) Encode() []byte {
	ec := NewMTPEncodeBuffer()

	ec.Int(int32(TL_CLASS_langpack_getDifference))
	ec.Int(t.Get_from_version())

	return ec.GetBuffer()
}

func (t *TL_langpack_getDifference) Decode(b []byte) {
	dc := NewMTPDecodeBuffer(b)

	t._from_version = dc.Int()

}

// langpack_getLanguages#800fd57d
type TL_langpack_getLanguages struct {
}

func New_TL_langpack_getLanguages() *TL_langpack_getLanguages {
	return &TL_langpack_getLanguages{}
}

func (t *TL_langpack_getLanguages) Encode() []byte {
	return nil
}

func (t *TL_langpack_getLanguages) Decode(b []byte) {

}
