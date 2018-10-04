package tl

//====resPQ#05162463====

type TL_resPQ struct {
	_nonce                          []byte
	_server_nonce                   []byte
	_pq                             string
	_server_public_key_fingerprints []byte
}

func New_TL_resPQ() *TL_resPQ {
	return &TL_resPQ{}
}

func (t *TL_resPQ) Encode() []byte {
	return nil
}

func (t *TL_resPQ) Decode(b []byte) {}

//====p_q_inner_data#83c95aec====

type TL_p_q_inner_data struct {
	_pq           string
	_p            string
	_q            string
	_nonce        []byte
	_server_nonce []byte
	_new_nonce    []byte
}

func New_TL_p_q_inner_data() *TL_p_q_inner_data {
	return &TL_p_q_inner_data{}
}

func (t *TL_p_q_inner_data) Encode() []byte {
	return nil
}

func (t *TL_p_q_inner_data) Decode(b []byte) {}

//====server_DH_params_fail#79cb045d====

type TL_server_DH_params_fail struct {
	_nonce          []byte
	_server_nonce   []byte
	_new_nonce_hash []byte
}

func New_TL_server_DH_params_fail() *TL_server_DH_params_fail {
	return &TL_server_DH_params_fail{}
}

func (t *TL_server_DH_params_fail) Encode() []byte {
	return nil
}

func (t *TL_server_DH_params_fail) Decode(b []byte) {}

//====server_DH_params_ok#d0e8075c====

type TL_server_DH_params_ok struct {
	_nonce            []byte
	_server_nonce     []byte
	_encrypted_answer string
}

func New_TL_server_DH_params_ok() *TL_server_DH_params_ok {
	return &TL_server_DH_params_ok{}
}

func (t *TL_server_DH_params_ok) Encode() []byte {
	return nil
}

func (t *TL_server_DH_params_ok) Decode(b []byte) {}

//====server_DH_inner_data#b5890dba====

type TL_server_DH_inner_data struct {
	_nonce        []byte
	_server_nonce []byte
	_g            []byte
	_dh_prime     string
	_g_a          string
	_server_time  []byte
}

func New_TL_server_DH_inner_data() *TL_server_DH_inner_data {
	return &TL_server_DH_inner_data{}
}

func (t *TL_server_DH_inner_data) Encode() []byte {
	return nil
}

func (t *TL_server_DH_inner_data) Decode(b []byte) {}

//====client_DH_inner_data#6643b654====

type TL_client_DH_inner_data struct {
	_nonce        []byte
	_server_nonce []byte
	_retry_id     int64
	_g_b          string
}

func New_TL_client_DH_inner_data() *TL_client_DH_inner_data {
	return &TL_client_DH_inner_data{}
}

func (t *TL_client_DH_inner_data) Encode() []byte {
	return nil
}

func (t *TL_client_DH_inner_data) Decode(b []byte) {}

//====dh_gen_ok#3bcbf734====

type TL_dh_gen_ok struct {
	_nonce           []byte
	_server_nonce    []byte
	_new_nonce_hash1 []byte
}

func New_TL_dh_gen_ok() *TL_dh_gen_ok {
	return &TL_dh_gen_ok{}
}

func (t *TL_dh_gen_ok) Encode() []byte {
	return nil
}

func (t *TL_dh_gen_ok) Decode(b []byte) {}

//====dh_gen_retry#46dc1fb9====

type TL_dh_gen_retry struct {
	_nonce           []byte
	_server_nonce    []byte
	_new_nonce_hash2 []byte
}

func New_TL_dh_gen_retry() *TL_dh_gen_retry {
	return &TL_dh_gen_retry{}
}

func (t *TL_dh_gen_retry) Encode() []byte {
	return nil
}

func (t *TL_dh_gen_retry) Decode(b []byte) {}

//====dh_gen_fail#a69dae02====

type TL_dh_gen_fail struct {
	_nonce           []byte
	_server_nonce    []byte
	_new_nonce_hash3 []byte
}

func New_TL_dh_gen_fail() *TL_dh_gen_fail {
	return &TL_dh_gen_fail{}
}

func (t *TL_dh_gen_fail) Encode() []byte {
	return nil
}

func (t *TL_dh_gen_fail) Decode(b []byte) {}

//====destroy_auth_key_ok#f660e1d4====

type TL_destroy_auth_key_ok struct {
}

func New_TL_destroy_auth_key_ok() *TL_destroy_auth_key_ok {
	return &TL_destroy_auth_key_ok{}
}

func (t *TL_destroy_auth_key_ok) Encode() []byte {
	return nil
}

func (t *TL_destroy_auth_key_ok) Decode(b []byte) {}

//====destroy_auth_key_none#0a9f2259====

type TL_destroy_auth_key_none struct {
}

func New_TL_destroy_auth_key_none() *TL_destroy_auth_key_none {
	return &TL_destroy_auth_key_none{}
}

func (t *TL_destroy_auth_key_none) Encode() []byte {
	return nil
}

func (t *TL_destroy_auth_key_none) Decode(b []byte) {}

//====destroy_auth_key_fail#ea109b13====

type TL_destroy_auth_key_fail struct {
}

func New_TL_destroy_auth_key_fail() *TL_destroy_auth_key_fail {
	return &TL_destroy_auth_key_fail{}
}

func (t *TL_destroy_auth_key_fail) Encode() []byte {
	return nil
}

func (t *TL_destroy_auth_key_fail) Decode(b []byte) {}

//====req_pq#60469778====

type TL_req_pq struct {
	_nonce []byte
}

func New_TL_req_pq() *TL_req_pq {
	return &TL_req_pq{}
}

func (t *TL_req_pq) Encode() []byte {
	return nil
}

func (t *TL_req_pq) Decode(b []byte) {}

//====req_DH_params#d712e4be====

type TL_req_DH_params struct {
	_nonce                  []byte
	_server_nonce           []byte
	_p                      string
	_q                      string
	_public_key_fingerprint int64
	_encrypted_data         string
}

func New_TL_req_DH_params() *TL_req_DH_params {
	return &TL_req_DH_params{}
}

func (t *TL_req_DH_params) Encode() []byte {
	return nil
}

func (t *TL_req_DH_params) Decode(b []byte) {}

//====set_client_DH_params#f5045f1f====

type TL_set_client_DH_params struct {
	_nonce          []byte
	_server_nonce   []byte
	_encrypted_data string
}

func New_TL_set_client_DH_params() *TL_set_client_DH_params {
	return &TL_set_client_DH_params{}
}

func (t *TL_set_client_DH_params) Encode() []byte {
	return nil
}

func (t *TL_set_client_DH_params) Decode(b []byte) {}

//====destroy_auth_key#d1435160====

type TL_destroy_auth_key struct {
}

func New_TL_destroy_auth_key() *TL_destroy_auth_key {
	return &TL_destroy_auth_key{}
}

func (t *TL_destroy_auth_key) Encode() []byte {
	return nil
}

func (t *TL_destroy_auth_key) Decode(b []byte) {}

//====msgs_ack#62d6b459====

type TL_msgs_ack struct {
	_msg_ids []byte
}

func New_TL_msgs_ack() *TL_msgs_ack {
	return &TL_msgs_ack{}
}

func (t *TL_msgs_ack) Encode() []byte {
	return nil
}

func (t *TL_msgs_ack) Decode(b []byte) {}

//====bad_msg_notification#a7eff811====

type TL_bad_msg_notification struct {
	_bad_msg_id    int64
	_bad_msg_seqno []byte
	_error_code    []byte
}

func New_TL_bad_msg_notification() *TL_bad_msg_notification {
	return &TL_bad_msg_notification{}
}

func (t *TL_bad_msg_notification) Encode() []byte {
	return nil
}

func (t *TL_bad_msg_notification) Decode(b []byte) {}

//====bad_server_salt#edab447b====

type TL_bad_server_salt struct {
	_bad_msg_id      int64
	_bad_msg_seqno   []byte
	_error_code      []byte
	_new_server_salt int64
}

func New_TL_bad_server_salt() *TL_bad_server_salt {
	return &TL_bad_server_salt{}
}

func (t *TL_bad_server_salt) Encode() []byte {
	return nil
}

func (t *TL_bad_server_salt) Decode(b []byte) {}

//====msgs_state_req#da69fb52====

type TL_msgs_state_req struct {
	_msg_ids []byte
}

func New_TL_msgs_state_req() *TL_msgs_state_req {
	return &TL_msgs_state_req{}
}

func (t *TL_msgs_state_req) Encode() []byte {
	return nil
}

func (t *TL_msgs_state_req) Decode(b []byte) {}

//====msgs_state_info#04deb57d====

type TL_msgs_state_info struct {
	_req_msg_id int64
	_info       string
}

func New_TL_msgs_state_info() *TL_msgs_state_info {
	return &TL_msgs_state_info{}
}

func (t *TL_msgs_state_info) Encode() []byte {
	return nil
}

func (t *TL_msgs_state_info) Decode(b []byte) {}

//====msgs_all_info#8cc0d131====

type TL_msgs_all_info struct {
	_msg_ids []byte
	_info    string
}

func New_TL_msgs_all_info() *TL_msgs_all_info {
	return &TL_msgs_all_info{}
}

func (t *TL_msgs_all_info) Encode() []byte {
	return nil
}

func (t *TL_msgs_all_info) Decode(b []byte) {}

//====msg_detailed_info#276d3ec6====

type TL_msg_detailed_info struct {
	_msg_id        int64
	_answer_msg_id int64
	_bytes         []byte
	_status        []byte
}

func New_TL_msg_detailed_info() *TL_msg_detailed_info {
	return &TL_msg_detailed_info{}
}

func (t *TL_msg_detailed_info) Encode() []byte {
	return nil
}

func (t *TL_msg_detailed_info) Decode(b []byte) {}

//====msg_new_detailed_info#809db6df====

type TL_msg_new_detailed_info struct {
	_answer_msg_id int64
	_bytes         []byte
	_status        []byte
}

func New_TL_msg_new_detailed_info() *TL_msg_new_detailed_info {
	return &TL_msg_new_detailed_info{}
}

func (t *TL_msg_new_detailed_info) Encode() []byte {
	return nil
}

func (t *TL_msg_new_detailed_info) Decode(b []byte) {}

//====msg_resend_req#7d861a08====

type TL_msg_resend_req struct {
	_msg_ids []byte
}

func New_TL_msg_resend_req() *TL_msg_resend_req {
	return &TL_msg_resend_req{}
}

func (t *TL_msg_resend_req) Encode() []byte {
	return nil
}

func (t *TL_msg_resend_req) Decode(b []byte) {}

//====rpc_error#2144ca19====

type TL_rpc_error struct {
	_error_code    []byte
	_error_message string
}

func New_TL_rpc_error() *TL_rpc_error {
	return &TL_rpc_error{}
}

func (t *TL_rpc_error) Encode() []byte {
	return nil
}

func (t *TL_rpc_error) Decode(b []byte) {}

//====rpc_answer_unknown#5e2ad36e====

type TL_rpc_answer_unknown struct {
}

func New_TL_rpc_answer_unknown() *TL_rpc_answer_unknown {
	return &TL_rpc_answer_unknown{}
}

func (t *TL_rpc_answer_unknown) Encode() []byte {
	return nil
}

func (t *TL_rpc_answer_unknown) Decode(b []byte) {}

//====rpc_answer_dropped_running#cd78e586====

type TL_rpc_answer_dropped_running struct {
}

func New_TL_rpc_answer_dropped_running() *TL_rpc_answer_dropped_running {
	return &TL_rpc_answer_dropped_running{}
}

func (t *TL_rpc_answer_dropped_running) Encode() []byte {
	return nil
}

func (t *TL_rpc_answer_dropped_running) Decode(b []byte) {}

//====rpc_answer_dropped#a43ad8b7====

type TL_rpc_answer_dropped struct {
	_msg_id int64
	_seq_no []byte
	_bytes  []byte
}

func New_TL_rpc_answer_dropped() *TL_rpc_answer_dropped {
	return &TL_rpc_answer_dropped{}
}

func (t *TL_rpc_answer_dropped) Encode() []byte {
	return nil
}

func (t *TL_rpc_answer_dropped) Decode(b []byte) {}

//====future_salt#0949d9dc====

type TL_future_salt struct {
	_valid_since []byte
	_valid_until []byte
	_salt        int64
}

func New_TL_future_salt() *TL_future_salt {
	return &TL_future_salt{}
}

func (t *TL_future_salt) Encode() []byte {
	return nil
}

func (t *TL_future_salt) Decode(b []byte) {}

//====future_salts#ae500895====

type TL_future_salts struct {
	_req_msg_id int64
	_now        []byte
	_salts      []byte
}

func New_TL_future_salts() *TL_future_salts {
	return &TL_future_salts{}
}

func (t *TL_future_salts) Encode() []byte {
	return nil
}

func (t *TL_future_salts) Decode(b []byte) {}

//====pong#347773c5====

type TL_pong struct {
	_msg_id  int64
	_ping_id int64
}

func New_TL_pong() *TL_pong {
	return &TL_pong{}
}

func (t *TL_pong) Encode() []byte {
	return nil
}

func (t *TL_pong) Decode(b []byte) {}

//====destroy_session_ok#e22045fc====

type TL_destroy_session_ok struct {
	_session_id int64
}

func New_TL_destroy_session_ok() *TL_destroy_session_ok {
	return &TL_destroy_session_ok{}
}

func (t *TL_destroy_session_ok) Encode() []byte {
	return nil
}

func (t *TL_destroy_session_ok) Decode(b []byte) {}

//====destroy_session_none#62d350c9====

type TL_destroy_session_none struct {
	_session_id int64
}

func New_TL_destroy_session_none() *TL_destroy_session_none {
	return &TL_destroy_session_none{}
}

func (t *TL_destroy_session_none) Encode() []byte {
	return nil
}

func (t *TL_destroy_session_none) Decode(b []byte) {}

//====new_session_created#9ec20908====

type TL_new_session_created struct {
	_first_msg_id int64
	_unique_id    int64
	_server_salt  int64
}

func New_TL_new_session_created() *TL_new_session_created {
	return &TL_new_session_created{}
}

func (t *TL_new_session_created) Encode() []byte {
	return nil
}

func (t *TL_new_session_created) Decode(b []byte) {}

//====http_wait#9299359f====

type TL_http_wait struct {
	_max_delay  []byte
	_wait_after []byte
	_max_wait   []byte
}

func New_TL_http_wait() *TL_http_wait {
	return &TL_http_wait{}
}

func (t *TL_http_wait) Encode() []byte {
	return nil
}

func (t *TL_http_wait) Decode(b []byte) {}

//====ipPort#====

type TL_ipPort struct {
	_ipv4 []byte
	_port []byte
}

func New_TL_ipPort() *TL_ipPort {
	return &TL_ipPort{}
}

func (t *TL_ipPort) Encode() []byte {
	return nil
}

func (t *TL_ipPort) Decode(b []byte) {}

//====help_configSimple#d997c3c5====

type TL_help_configSimple struct {
	_date         []byte
	_expires      []byte
	_dc_id        []byte
	_ip_port_list []byte
}

func New_TL_help_configSimple() *TL_help_configSimple {
	return &TL_help_configSimple{}
}

func (t *TL_help_configSimple) Encode() []byte {
	return nil
}

func (t *TL_help_configSimple) Decode(b []byte) {}

//====rpc_drop_answer#58e4a740====

type TL_rpc_drop_answer struct {
	_req_msg_id int64
}

func New_TL_rpc_drop_answer() *TL_rpc_drop_answer {
	return &TL_rpc_drop_answer{}
}

func (t *TL_rpc_drop_answer) Encode() []byte {
	return nil
}

func (t *TL_rpc_drop_answer) Decode(b []byte) {}

//====get_future_salts#b921bd04====

type TL_get_future_salts struct {
	_num []byte
}

func New_TL_get_future_salts() *TL_get_future_salts {
	return &TL_get_future_salts{}
}

func (t *TL_get_future_salts) Encode() []byte {
	return nil
}

func (t *TL_get_future_salts) Decode(b []byte) {}

//====ping#7abe77ec====

type TL_ping struct {
	_ping_id int64
}

func New_TL_ping() *TL_ping {
	return &TL_ping{}
}

func (t *TL_ping) Encode() []byte {
	return nil
}

func (t *TL_ping) Decode(b []byte) {}

//====ping_delay_disconnect#f3427b8c====

type TL_ping_delay_disconnect struct {
	_ping_id          int64
	_disconnect_delay []byte
}

func New_TL_ping_delay_disconnect() *TL_ping_delay_disconnect {
	return &TL_ping_delay_disconnect{}
}

func (t *TL_ping_delay_disconnect) Encode() []byte {
	return nil
}

func (t *TL_ping_delay_disconnect) Decode(b []byte) {}

//====destroy_session#e7512126====

type TL_destroy_session struct {
	_session_id int64
}

func New_TL_destroy_session() *TL_destroy_session {
	return &TL_destroy_session{}
}

func (t *TL_destroy_session) Encode() []byte {
	return nil
}

func (t *TL_destroy_session) Decode(b []byte) {}

//====contest_saveDeveloperInfo#9a5f6e95====

type TL_contest_saveDeveloperInfo struct {
	_vk_id        []byte
	_name         string
	_phone_number string
	_age          []byte
	_city         string
}

func New_TL_contest_saveDeveloperInfo() *TL_contest_saveDeveloperInfo {
	return &TL_contest_saveDeveloperInfo{}
}

func (t *TL_contest_saveDeveloperInfo) Encode() []byte {
	return nil
}

func (t *TL_contest_saveDeveloperInfo) Decode(b []byte) {}

//====boolFalse#bc799737====

type TL_boolFalse struct {
}

func New_TL_boolFalse() *TL_boolFalse {
	return &TL_boolFalse{}
}

func (t *TL_boolFalse) Encode() []byte {
	return nil
}

func (t *TL_boolFalse) Decode(b []byte) {}

//====boolTrue#997275b5====

type TL_boolTrue struct {
}

func New_TL_boolTrue() *TL_boolTrue {
	return &TL_boolTrue{}
}

func (t *TL_boolTrue) Encode() []byte {
	return nil
}

func (t *TL_boolTrue) Decode(b []byte) {}

//====true#3fedd339====

type TL_true struct {
}

func New_TL_true() *TL_true {
	return &TL_true{}
}

func (t *TL_true) Encode() []byte {
	return nil
}

func (t *TL_true) Decode(b []byte) {}

//====vector#1cb5c415====

type TL_vector struct {
}

func New_TL_vector() *TL_vector {
	return &TL_vector{}
}

func (t *TL_vector) Encode() []byte {
	return nil
}

func (t *TL_vector) Decode(b []byte) {}

//====error#c4b9f9bb====

type TL_error struct {
	_code []byte
	_text string
}

func New_TL_error() *TL_error {
	return &TL_error{}
}

func (t *TL_error) Encode() []byte {
	return nil
}

func (t *TL_error) Decode(b []byte) {}

//====null#56730bcc====

type TL_null struct {
}

func New_TL_null() *TL_null {
	return &TL_null{}
}

func (t *TL_null) Encode() []byte {
	return nil
}

func (t *TL_null) Decode(b []byte) {}

//====inputPeerEmpty#7f3b18ea====

type TL_inputPeerEmpty struct {
}

func New_TL_inputPeerEmpty() *TL_inputPeerEmpty {
	return &TL_inputPeerEmpty{}
}

func (t *TL_inputPeerEmpty) Encode() []byte {
	return nil
}

func (t *TL_inputPeerEmpty) Decode(b []byte) {}

//====inputPeerSelf#7da07ec9====

type TL_inputPeerSelf struct {
}

func New_TL_inputPeerSelf() *TL_inputPeerSelf {
	return &TL_inputPeerSelf{}
}

func (t *TL_inputPeerSelf) Encode() []byte {
	return nil
}

func (t *TL_inputPeerSelf) Decode(b []byte) {}

//====inputPeerChat#179be863====

type TL_inputPeerChat struct {
	_chat_id []byte
}

func New_TL_inputPeerChat() *TL_inputPeerChat {
	return &TL_inputPeerChat{}
}

func (t *TL_inputPeerChat) Encode() []byte {
	return nil
}

func (t *TL_inputPeerChat) Decode(b []byte) {}

//====inputPeerUser#7b8e7de6====

type TL_inputPeerUser struct {
	_user_id     []byte
	_access_hash int64
}

func New_TL_inputPeerUser() *TL_inputPeerUser {
	return &TL_inputPeerUser{}
}

func (t *TL_inputPeerUser) Encode() []byte {
	return nil
}

func (t *TL_inputPeerUser) Decode(b []byte) {}

//====inputPeerChannel#20adaef8====

type TL_inputPeerChannel struct {
	_channel_id  []byte
	_access_hash int64
}

func New_TL_inputPeerChannel() *TL_inputPeerChannel {
	return &TL_inputPeerChannel{}
}

func (t *TL_inputPeerChannel) Encode() []byte {
	return nil
}

func (t *TL_inputPeerChannel) Decode(b []byte) {}

//====inputUserEmpty#b98886cf====

type TL_inputUserEmpty struct {
}

func New_TL_inputUserEmpty() *TL_inputUserEmpty {
	return &TL_inputUserEmpty{}
}

func (t *TL_inputUserEmpty) Encode() []byte {
	return nil
}

func (t *TL_inputUserEmpty) Decode(b []byte) {}

//====inputUserSelf#f7c1b13f====

type TL_inputUserSelf struct {
}

func New_TL_inputUserSelf() *TL_inputUserSelf {
	return &TL_inputUserSelf{}
}

func (t *TL_inputUserSelf) Encode() []byte {
	return nil
}

func (t *TL_inputUserSelf) Decode(b []byte) {}

//====inputUser#d8292816====

type TL_inputUser struct {
	_user_id     []byte
	_access_hash int64
}

func New_TL_inputUser() *TL_inputUser {
	return &TL_inputUser{}
}

func (t *TL_inputUser) Encode() []byte {
	return nil
}

func (t *TL_inputUser) Decode(b []byte) {}

//====inputPhoneContact#f392b7f4====

type TL_inputPhoneContact struct {
	_client_id  int64
	_phone      string
	_first_name string
	_last_name  string
}

func New_TL_inputPhoneContact() *TL_inputPhoneContact {
	return &TL_inputPhoneContact{}
}

func (t *TL_inputPhoneContact) Encode() []byte {
	return nil
}

func (t *TL_inputPhoneContact) Decode(b []byte) {}

//====inputFile#f52ff27f====

type TL_inputFile struct {
	_id           int64
	_parts        []byte
	_name         string
	_md5_checksum string
}

func New_TL_inputFile() *TL_inputFile {
	return &TL_inputFile{}
}

func (t *TL_inputFile) Encode() []byte {
	return nil
}

func (t *TL_inputFile) Decode(b []byte) {}

//====inputFileBig#fa4f0bb5====

type TL_inputFileBig struct {
	_id    int64
	_parts []byte
	_name  string
}

func New_TL_inputFileBig() *TL_inputFileBig {
	return &TL_inputFileBig{}
}

func (t *TL_inputFileBig) Encode() []byte {
	return nil
}

func (t *TL_inputFileBig) Decode(b []byte) {}

//====inputMediaEmpty#9664f57f====

type TL_inputMediaEmpty struct {
}

func New_TL_inputMediaEmpty() *TL_inputMediaEmpty {
	return &TL_inputMediaEmpty{}
}

func (t *TL_inputMediaEmpty) Encode() []byte {
	return nil
}

func (t *TL_inputMediaEmpty) Decode(b []byte) {}

//====inputMediaUploadedPhoto#2f37e231====

type TL_inputMediaUploadedPhoto struct {
	_flags       []byte
	_file        []byte
	_caption     string
	_stickers    []byte
	_ttl_seconds []byte
}

func New_TL_inputMediaUploadedPhoto() *TL_inputMediaUploadedPhoto {
	return &TL_inputMediaUploadedPhoto{}
}

func (t *TL_inputMediaUploadedPhoto) Encode() []byte {
	return nil
}

func (t *TL_inputMediaUploadedPhoto) Decode(b []byte) {}

//====inputMediaPhoto#81fa373a====

type TL_inputMediaPhoto struct {
	_flags       []byte
	_id          []byte
	_caption     string
	_ttl_seconds []byte
}

func New_TL_inputMediaPhoto() *TL_inputMediaPhoto {
	return &TL_inputMediaPhoto{}
}

func (t *TL_inputMediaPhoto) Encode() []byte {
	return nil
}

func (t *TL_inputMediaPhoto) Decode(b []byte) {}

//====inputMediaGeoPoint#f9c44144====

type TL_inputMediaGeoPoint struct {
	_geo_point []byte
}

func New_TL_inputMediaGeoPoint() *TL_inputMediaGeoPoint {
	return &TL_inputMediaGeoPoint{}
}

func (t *TL_inputMediaGeoPoint) Encode() []byte {
	return nil
}

func (t *TL_inputMediaGeoPoint) Decode(b []byte) {}

//====inputMediaContact#a6e45987====

type TL_inputMediaContact struct {
	_phone_number string
	_first_name   string
	_last_name    string
}

func New_TL_inputMediaContact() *TL_inputMediaContact {
	return &TL_inputMediaContact{}
}

func (t *TL_inputMediaContact) Encode() []byte {
	return nil
}

func (t *TL_inputMediaContact) Decode(b []byte) {}

//====inputMediaUploadedDocument#e39621fd====

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

func New_TL_inputMediaUploadedDocument() *TL_inputMediaUploadedDocument {
	return &TL_inputMediaUploadedDocument{}
}

func (t *TL_inputMediaUploadedDocument) Encode() []byte {
	return nil
}

func (t *TL_inputMediaUploadedDocument) Decode(b []byte) {}

//====inputMediaDocument#5acb668e====

type TL_inputMediaDocument struct {
	_flags       []byte
	_id          []byte
	_caption     string
	_ttl_seconds []byte
}

func New_TL_inputMediaDocument() *TL_inputMediaDocument {
	return &TL_inputMediaDocument{}
}

func (t *TL_inputMediaDocument) Encode() []byte {
	return nil
}

func (t *TL_inputMediaDocument) Decode(b []byte) {}

//====inputMediaVenue#c13d1c11====

type TL_inputMediaVenue struct {
	_geo_point  []byte
	_title      string
	_address    string
	_provider   string
	_venue_id   string
	_venue_type string
}

func New_TL_inputMediaVenue() *TL_inputMediaVenue {
	return &TL_inputMediaVenue{}
}

func (t *TL_inputMediaVenue) Encode() []byte {
	return nil
}

func (t *TL_inputMediaVenue) Decode(b []byte) {}

//====inputMediaGifExternal#4843b0fd====

type TL_inputMediaGifExternal struct {
	_url string
	_q   string
}

func New_TL_inputMediaGifExternal() *TL_inputMediaGifExternal {
	return &TL_inputMediaGifExternal{}
}

func (t *TL_inputMediaGifExternal) Encode() []byte {
	return nil
}

func (t *TL_inputMediaGifExternal) Decode(b []byte) {}

//====inputMediaPhotoExternal#0922aec1====

type TL_inputMediaPhotoExternal struct {
	_flags       []byte
	_url         string
	_caption     string
	_ttl_seconds []byte
}

func New_TL_inputMediaPhotoExternal() *TL_inputMediaPhotoExternal {
	return &TL_inputMediaPhotoExternal{}
}

func (t *TL_inputMediaPhotoExternal) Encode() []byte {
	return nil
}

func (t *TL_inputMediaPhotoExternal) Decode(b []byte) {}

//====inputMediaDocumentExternal#b6f74335====

type TL_inputMediaDocumentExternal struct {
	_flags       []byte
	_url         string
	_caption     string
	_ttl_seconds []byte
}

func New_TL_inputMediaDocumentExternal() *TL_inputMediaDocumentExternal {
	return &TL_inputMediaDocumentExternal{}
}

func (t *TL_inputMediaDocumentExternal) Encode() []byte {
	return nil
}

func (t *TL_inputMediaDocumentExternal) Decode(b []byte) {}

//====inputMediaGame#d33f43f3====

type TL_inputMediaGame struct {
	_id []byte
}

func New_TL_inputMediaGame() *TL_inputMediaGame {
	return &TL_inputMediaGame{}
}

func (t *TL_inputMediaGame) Encode() []byte {
	return nil
}

func (t *TL_inputMediaGame) Decode(b []byte) {}

//====inputMediaInvoice#f4e096c3====

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

func New_TL_inputMediaInvoice() *TL_inputMediaInvoice {
	return &TL_inputMediaInvoice{}
}

func (t *TL_inputMediaInvoice) Encode() []byte {
	return nil
}

func (t *TL_inputMediaInvoice) Decode(b []byte) {}

//====inputMediaGeoLive#7b1a118f====

type TL_inputMediaGeoLive struct {
	_geo_point []byte
	_period    []byte
}

func New_TL_inputMediaGeoLive() *TL_inputMediaGeoLive {
	return &TL_inputMediaGeoLive{}
}

func (t *TL_inputMediaGeoLive) Encode() []byte {
	return nil
}

func (t *TL_inputMediaGeoLive) Decode(b []byte) {}

//====inputChatPhotoEmpty#1ca48f57====

type TL_inputChatPhotoEmpty struct {
}

func New_TL_inputChatPhotoEmpty() *TL_inputChatPhotoEmpty {
	return &TL_inputChatPhotoEmpty{}
}

func (t *TL_inputChatPhotoEmpty) Encode() []byte {
	return nil
}

func (t *TL_inputChatPhotoEmpty) Decode(b []byte) {}

//====inputChatUploadedPhoto#927c55b4====

type TL_inputChatUploadedPhoto struct {
	_file []byte
}

func New_TL_inputChatUploadedPhoto() *TL_inputChatUploadedPhoto {
	return &TL_inputChatUploadedPhoto{}
}

func (t *TL_inputChatUploadedPhoto) Encode() []byte {
	return nil
}

func (t *TL_inputChatUploadedPhoto) Decode(b []byte) {}

//====inputChatPhoto#8953ad37====

type TL_inputChatPhoto struct {
	_id []byte
}

func New_TL_inputChatPhoto() *TL_inputChatPhoto {
	return &TL_inputChatPhoto{}
}

func (t *TL_inputChatPhoto) Encode() []byte {
	return nil
}

func (t *TL_inputChatPhoto) Decode(b []byte) {}

//====inputGeoPointEmpty#e4c123d6====

type TL_inputGeoPointEmpty struct {
}

func New_TL_inputGeoPointEmpty() *TL_inputGeoPointEmpty {
	return &TL_inputGeoPointEmpty{}
}

func (t *TL_inputGeoPointEmpty) Encode() []byte {
	return nil
}

func (t *TL_inputGeoPointEmpty) Decode(b []byte) {}

//====inputGeoPoint#f3b7acc9====

type TL_inputGeoPoint struct {
	_lat  []byte
	_long []byte
}

func New_TL_inputGeoPoint() *TL_inputGeoPoint {
	return &TL_inputGeoPoint{}
}

func (t *TL_inputGeoPoint) Encode() []byte {
	return nil
}

func (t *TL_inputGeoPoint) Decode(b []byte) {}

//====inputPhotoEmpty#1cd7bf0d====

type TL_inputPhotoEmpty struct {
}

func New_TL_inputPhotoEmpty() *TL_inputPhotoEmpty {
	return &TL_inputPhotoEmpty{}
}

func (t *TL_inputPhotoEmpty) Encode() []byte {
	return nil
}

func (t *TL_inputPhotoEmpty) Decode(b []byte) {}

//====inputPhoto#fb95c6c4====

type TL_inputPhoto struct {
	_id          int64
	_access_hash int64
}

func New_TL_inputPhoto() *TL_inputPhoto {
	return &TL_inputPhoto{}
}

func (t *TL_inputPhoto) Encode() []byte {
	return nil
}

func (t *TL_inputPhoto) Decode(b []byte) {}

//====inputFileLocation#14637196====

type TL_inputFileLocation struct {
	_volume_id int64
	_local_id  []byte
	_secret    int64
}

func New_TL_inputFileLocation() *TL_inputFileLocation {
	return &TL_inputFileLocation{}
}

func (t *TL_inputFileLocation) Encode() []byte {
	return nil
}

func (t *TL_inputFileLocation) Decode(b []byte) {}

//====inputEncryptedFileLocation#f5235d55====

type TL_inputEncryptedFileLocation struct {
	_id          int64
	_access_hash int64
}

func New_TL_inputEncryptedFileLocation() *TL_inputEncryptedFileLocation {
	return &TL_inputEncryptedFileLocation{}
}

func (t *TL_inputEncryptedFileLocation) Encode() []byte {
	return nil
}

func (t *TL_inputEncryptedFileLocation) Decode(b []byte) {}

//====inputDocumentFileLocation#430f0724====

type TL_inputDocumentFileLocation struct {
	_id          int64
	_access_hash int64
	_version     []byte
}

func New_TL_inputDocumentFileLocation() *TL_inputDocumentFileLocation {
	return &TL_inputDocumentFileLocation{}
}

func (t *TL_inputDocumentFileLocation) Encode() []byte {
	return nil
}

func (t *TL_inputDocumentFileLocation) Decode(b []byte) {}

//====inputAppEvent#770656a8====

type TL_inputAppEvent struct {
	_time []byte
	_type string
	_peer int64
	_data string
}

func New_TL_inputAppEvent() *TL_inputAppEvent {
	return &TL_inputAppEvent{}
}

func (t *TL_inputAppEvent) Encode() []byte {
	return nil
}

func (t *TL_inputAppEvent) Decode(b []byte) {}

//====peerUser#9db1bc6d====

type TL_peerUser struct {
	_user_id []byte
}

func New_TL_peerUser() *TL_peerUser {
	return &TL_peerUser{}
}

func (t *TL_peerUser) Encode() []byte {
	return nil
}

func (t *TL_peerUser) Decode(b []byte) {}

//====peerChat#bad0e5bb====

type TL_peerChat struct {
	_chat_id []byte
}

func New_TL_peerChat() *TL_peerChat {
	return &TL_peerChat{}
}

func (t *TL_peerChat) Encode() []byte {
	return nil
}

func (t *TL_peerChat) Decode(b []byte) {}

//====peerChannel#bddde532====

type TL_peerChannel struct {
	_channel_id []byte
}

func New_TL_peerChannel() *TL_peerChannel {
	return &TL_peerChannel{}
}

func (t *TL_peerChannel) Encode() []byte {
	return nil
}

func (t *TL_peerChannel) Decode(b []byte) {}

//====storage_fileUnknown#aa963b05====

type TL_storage_fileUnknown struct {
}

func New_TL_storage_fileUnknown() *TL_storage_fileUnknown {
	return &TL_storage_fileUnknown{}
}

func (t *TL_storage_fileUnknown) Encode() []byte {
	return nil
}

func (t *TL_storage_fileUnknown) Decode(b []byte) {}

//====storage_filePartial#40bc6f52====

type TL_storage_filePartial struct {
}

func New_TL_storage_filePartial() *TL_storage_filePartial {
	return &TL_storage_filePartial{}
}

func (t *TL_storage_filePartial) Encode() []byte {
	return nil
}

func (t *TL_storage_filePartial) Decode(b []byte) {}

//====storage_fileJpeg#007efe0e====

type TL_storage_fileJpeg struct {
}

func New_TL_storage_fileJpeg() *TL_storage_fileJpeg {
	return &TL_storage_fileJpeg{}
}

func (t *TL_storage_fileJpeg) Encode() []byte {
	return nil
}

func (t *TL_storage_fileJpeg) Decode(b []byte) {}

//====storage_fileGif#cae1aadf====

type TL_storage_fileGif struct {
}

func New_TL_storage_fileGif() *TL_storage_fileGif {
	return &TL_storage_fileGif{}
}

func (t *TL_storage_fileGif) Encode() []byte {
	return nil
}

func (t *TL_storage_fileGif) Decode(b []byte) {}

//====storage_filePng#0a4f63c0====

type TL_storage_filePng struct {
}

func New_TL_storage_filePng() *TL_storage_filePng {
	return &TL_storage_filePng{}
}

func (t *TL_storage_filePng) Encode() []byte {
	return nil
}

func (t *TL_storage_filePng) Decode(b []byte) {}

//====storage_filePdf#ae1e508d====

type TL_storage_filePdf struct {
}

func New_TL_storage_filePdf() *TL_storage_filePdf {
	return &TL_storage_filePdf{}
}

func (t *TL_storage_filePdf) Encode() []byte {
	return nil
}

func (t *TL_storage_filePdf) Decode(b []byte) {}

//====storage_fileMp3#528a0677====

type TL_storage_fileMp3 struct {
}

func New_TL_storage_fileMp3() *TL_storage_fileMp3 {
	return &TL_storage_fileMp3{}
}

func (t *TL_storage_fileMp3) Encode() []byte {
	return nil
}

func (t *TL_storage_fileMp3) Decode(b []byte) {}

//====storage_fileMov#4b09ebbc====

type TL_storage_fileMov struct {
}

func New_TL_storage_fileMov() *TL_storage_fileMov {
	return &TL_storage_fileMov{}
}

func (t *TL_storage_fileMov) Encode() []byte {
	return nil
}

func (t *TL_storage_fileMov) Decode(b []byte) {}

//====storage_fileMp4#b3cea0e4====

type TL_storage_fileMp4 struct {
}

func New_TL_storage_fileMp4() *TL_storage_fileMp4 {
	return &TL_storage_fileMp4{}
}

func (t *TL_storage_fileMp4) Encode() []byte {
	return nil
}

func (t *TL_storage_fileMp4) Decode(b []byte) {}

//====storage_fileWebp#1081464c====

type TL_storage_fileWebp struct {
}

func New_TL_storage_fileWebp() *TL_storage_fileWebp {
	return &TL_storage_fileWebp{}
}

func (t *TL_storage_fileWebp) Encode() []byte {
	return nil
}

func (t *TL_storage_fileWebp) Decode(b []byte) {}

//====fileLocationUnavailable#7c596b46====

type TL_fileLocationUnavailable struct {
	_volume_id int64
	_local_id  []byte
	_secret    int64
}

func New_TL_fileLocationUnavailable() *TL_fileLocationUnavailable {
	return &TL_fileLocationUnavailable{}
}

func (t *TL_fileLocationUnavailable) Encode() []byte {
	return nil
}

func (t *TL_fileLocationUnavailable) Decode(b []byte) {}

//====fileLocation#53d69076====

type TL_fileLocation struct {
	_dc_id     []byte
	_volume_id int64
	_local_id  []byte
	_secret    int64
}

func New_TL_fileLocation() *TL_fileLocation {
	return &TL_fileLocation{}
}

func (t *TL_fileLocation) Encode() []byte {
	return nil
}

func (t *TL_fileLocation) Decode(b []byte) {}

//====userEmpty#200250ba====

type TL_userEmpty struct {
	_id []byte
}

func New_TL_userEmpty() *TL_userEmpty {
	return &TL_userEmpty{}
}

func (t *TL_userEmpty) Encode() []byte {
	return nil
}

func (t *TL_userEmpty) Decode(b []byte) {}

//====user#2e13f4c3====

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
	_id                     []byte
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

func New_TL_user() *TL_user {
	return &TL_user{}
}

func (t *TL_user) Encode() []byte {
	return nil
}

func (t *TL_user) Decode(b []byte) {}

//====userProfilePhotoEmpty#4f11bae1====

type TL_userProfilePhotoEmpty struct {
}

func New_TL_userProfilePhotoEmpty() *TL_userProfilePhotoEmpty {
	return &TL_userProfilePhotoEmpty{}
}

func (t *TL_userProfilePhotoEmpty) Encode() []byte {
	return nil
}

func (t *TL_userProfilePhotoEmpty) Decode(b []byte) {}

//====userProfilePhoto#d559d8c8====

type TL_userProfilePhoto struct {
	_photo_id    int64
	_photo_small []byte
	_photo_big   []byte
}

func New_TL_userProfilePhoto() *TL_userProfilePhoto {
	return &TL_userProfilePhoto{}
}

func (t *TL_userProfilePhoto) Encode() []byte {
	return nil
}

func (t *TL_userProfilePhoto) Decode(b []byte) {}

//====userStatusEmpty#09d05049====

type TL_userStatusEmpty struct {
}

func New_TL_userStatusEmpty() *TL_userStatusEmpty {
	return &TL_userStatusEmpty{}
}

func (t *TL_userStatusEmpty) Encode() []byte {
	return nil
}

func (t *TL_userStatusEmpty) Decode(b []byte) {}

//====userStatusOnline#edb93949====

type TL_userStatusOnline struct {
	_expires []byte
}

func New_TL_userStatusOnline() *TL_userStatusOnline {
	return &TL_userStatusOnline{}
}

func (t *TL_userStatusOnline) Encode() []byte {
	return nil
}

func (t *TL_userStatusOnline) Decode(b []byte) {}

//====userStatusOffline#008c703f====

type TL_userStatusOffline struct {
	_was_online []byte
}

func New_TL_userStatusOffline() *TL_userStatusOffline {
	return &TL_userStatusOffline{}
}

func (t *TL_userStatusOffline) Encode() []byte {
	return nil
}

func (t *TL_userStatusOffline) Decode(b []byte) {}

//====userStatusRecently#e26f42f1====

type TL_userStatusRecently struct {
}

func New_TL_userStatusRecently() *TL_userStatusRecently {
	return &TL_userStatusRecently{}
}

func (t *TL_userStatusRecently) Encode() []byte {
	return nil
}

func (t *TL_userStatusRecently) Decode(b []byte) {}

//====userStatusLastWeek#07bf09fc====

type TL_userStatusLastWeek struct {
}

func New_TL_userStatusLastWeek() *TL_userStatusLastWeek {
	return &TL_userStatusLastWeek{}
}

func (t *TL_userStatusLastWeek) Encode() []byte {
	return nil
}

func (t *TL_userStatusLastWeek) Decode(b []byte) {}

//====userStatusLastMonth#77ebc742====

type TL_userStatusLastMonth struct {
}

func New_TL_userStatusLastMonth() *TL_userStatusLastMonth {
	return &TL_userStatusLastMonth{}
}

func (t *TL_userStatusLastMonth) Encode() []byte {
	return nil
}

func (t *TL_userStatusLastMonth) Decode(b []byte) {}

//====chatEmpty#9ba2d800====

type TL_chatEmpty struct {
	_id []byte
}

func New_TL_chatEmpty() *TL_chatEmpty {
	return &TL_chatEmpty{}
}

func (t *TL_chatEmpty) Encode() []byte {
	return nil
}

func (t *TL_chatEmpty) Decode(b []byte) {}

//====chat#d91cdd54====

type TL_chat struct {
	_flags              []byte
	_creator            []byte
	_kicked             []byte
	_left               []byte
	_admins_enabled     []byte
	_admin              []byte
	_deactivated        []byte
	_id                 []byte
	_title              string
	_photo              []byte
	_participants_count []byte
	_date               []byte
	_version            []byte
	_migrated_to        []byte
}

func New_TL_chat() *TL_chat {
	return &TL_chat{}
}

func (t *TL_chat) Encode() []byte {
	return nil
}

func (t *TL_chat) Decode(b []byte) {}

//====chatForbidden#07328bdb====

type TL_chatForbidden struct {
	_id    []byte
	_title string
}

func New_TL_chatForbidden() *TL_chatForbidden {
	return &TL_chatForbidden{}
}

func (t *TL_chatForbidden) Encode() []byte {
	return nil
}

func (t *TL_chatForbidden) Decode(b []byte) {}

//====channel#450b7115====

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
	_id                 []byte
	_access_hash        []byte
	_title              string
	_username           []byte
	_photo              []byte
	_date               []byte
	_version            []byte
	_restriction_reason []byte
	_admin_rights       []byte
	_banned_rights      []byte
	_participants_count []byte
}

func New_TL_channel() *TL_channel {
	return &TL_channel{}
}

func (t *TL_channel) Encode() []byte {
	return nil
}

func (t *TL_channel) Decode(b []byte) {}

//====channelForbidden#289da732====

type TL_channelForbidden struct {
	_flags       []byte
	_broadcast   []byte
	_megagroup   []byte
	_id          []byte
	_access_hash int64
	_title       string
	_until_date  []byte
}

func New_TL_channelForbidden() *TL_channelForbidden {
	return &TL_channelForbidden{}
}

func (t *TL_channelForbidden) Encode() []byte {
	return nil
}

func (t *TL_channelForbidden) Decode(b []byte) {}

//====chatFull#2e02a614====

type TL_chatFull struct {
	_id              []byte
	_participants    []byte
	_chat_photo      []byte
	_notify_settings []byte
	_exported_invite []byte
	_bot_info        []byte
}

func New_TL_chatFull() *TL_chatFull {
	return &TL_chatFull{}
}

func (t *TL_chatFull) Encode() []byte {
	return nil
}

func (t *TL_chatFull) Decode(b []byte) {}

//====channelFull#76af5481====

type TL_channelFull struct {
	_flags                 []byte
	_can_view_participants []byte
	_can_set_username      []byte
	_can_set_stickers      []byte
	_hidden_prehistory     []byte
	_id                    []byte
	_about                 string
	_participants_count    []byte
	_admins_count          []byte
	_kicked_count          []byte
	_banned_count          []byte
	_read_inbox_max_id     []byte
	_read_outbox_max_id    []byte
	_unread_count          []byte
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

func New_TL_channelFull() *TL_channelFull {
	return &TL_channelFull{}
}

func (t *TL_channelFull) Encode() []byte {
	return nil
}

func (t *TL_channelFull) Decode(b []byte) {}

//====chatParticipant#c8d7493e====

type TL_chatParticipant struct {
	_user_id    []byte
	_inviter_id []byte
	_date       []byte
}

func New_TL_chatParticipant() *TL_chatParticipant {
	return &TL_chatParticipant{}
}

func (t *TL_chatParticipant) Encode() []byte {
	return nil
}

func (t *TL_chatParticipant) Decode(b []byte) {}

//====chatParticipantCreator#da13538a====

type TL_chatParticipantCreator struct {
	_user_id []byte
}

func New_TL_chatParticipantCreator() *TL_chatParticipantCreator {
	return &TL_chatParticipantCreator{}
}

func (t *TL_chatParticipantCreator) Encode() []byte {
	return nil
}

func (t *TL_chatParticipantCreator) Decode(b []byte) {}

//====chatParticipantAdmin#e2d6e436====

type TL_chatParticipantAdmin struct {
	_user_id    []byte
	_inviter_id []byte
	_date       []byte
}

func New_TL_chatParticipantAdmin() *TL_chatParticipantAdmin {
	return &TL_chatParticipantAdmin{}
}

func (t *TL_chatParticipantAdmin) Encode() []byte {
	return nil
}

func (t *TL_chatParticipantAdmin) Decode(b []byte) {}

//====chatParticipantsForbidden#fc900c2b====

type TL_chatParticipantsForbidden struct {
	_flags            []byte
	_chat_id          []byte
	_self_participant []byte
}

func New_TL_chatParticipantsForbidden() *TL_chatParticipantsForbidden {
	return &TL_chatParticipantsForbidden{}
}

func (t *TL_chatParticipantsForbidden) Encode() []byte {
	return nil
}

func (t *TL_chatParticipantsForbidden) Decode(b []byte) {}

//====chatParticipants#3f460fed====

type TL_chatParticipants struct {
	_chat_id      []byte
	_participants []byte
	_version      []byte
}

func New_TL_chatParticipants() *TL_chatParticipants {
	return &TL_chatParticipants{}
}

func (t *TL_chatParticipants) Encode() []byte {
	return nil
}

func (t *TL_chatParticipants) Decode(b []byte) {}

//====chatPhotoEmpty#37c1011c====

type TL_chatPhotoEmpty struct {
}

func New_TL_chatPhotoEmpty() *TL_chatPhotoEmpty {
	return &TL_chatPhotoEmpty{}
}

func (t *TL_chatPhotoEmpty) Encode() []byte {
	return nil
}

func (t *TL_chatPhotoEmpty) Decode(b []byte) {}

//====chatPhoto#6153276a====

type TL_chatPhoto struct {
	_photo_small []byte
	_photo_big   []byte
}

func New_TL_chatPhoto() *TL_chatPhoto {
	return &TL_chatPhoto{}
}

func (t *TL_chatPhoto) Encode() []byte {
	return nil
}

func (t *TL_chatPhoto) Decode(b []byte) {}

//====messageEmpty#83e5de54====

type TL_messageEmpty struct {
	_id []byte
}

func New_TL_messageEmpty() *TL_messageEmpty {
	return &TL_messageEmpty{}
}

func (t *TL_messageEmpty) Encode() []byte {
	return nil
}

func (t *TL_messageEmpty) Decode(b []byte) {}

//====message#44f9b43d====

type TL_message struct {
	_flags           []byte
	_out             []byte
	_mentioned       []byte
	_media_unread    []byte
	_silent          []byte
	_post            []byte
	_id              []byte
	_from_id         []byte
	_to_id           []byte
	_fwd_from        []byte
	_via_bot_id      []byte
	_reply_to_msg_id []byte
	_date            []byte
	_message         string
	_media           []byte
	_reply_markup    []byte
	_entities        []byte
	_views           []byte
	_edit_date       []byte
	_post_author     []byte
	_grouped_id      []byte
}

func New_TL_message() *TL_message {
	return &TL_message{}
}

func (t *TL_message) Encode() []byte {
	return nil
}

func (t *TL_message) Decode(b []byte) {}

//====messageService#9e19a1f6====

type TL_messageService struct {
	_flags           []byte
	_out             []byte
	_mentioned       []byte
	_media_unread    []byte
	_silent          []byte
	_post            []byte
	_id              []byte
	_from_id         []byte
	_to_id           []byte
	_reply_to_msg_id []byte
	_date            []byte
	_action          []byte
}

func New_TL_messageService() *TL_messageService {
	return &TL_messageService{}
}

func (t *TL_messageService) Encode() []byte {
	return nil
}

func (t *TL_messageService) Decode(b []byte) {}

//====messageMediaEmpty#3ded6320====

type TL_messageMediaEmpty struct {
}

func New_TL_messageMediaEmpty() *TL_messageMediaEmpty {
	return &TL_messageMediaEmpty{}
}

func (t *TL_messageMediaEmpty) Encode() []byte {
	return nil
}

func (t *TL_messageMediaEmpty) Decode(b []byte) {}

//====messageMediaPhoto#b5223b0f====

type TL_messageMediaPhoto struct {
	_flags       []byte
	_photo       []byte
	_caption     []byte
	_ttl_seconds []byte
}

func New_TL_messageMediaPhoto() *TL_messageMediaPhoto {
	return &TL_messageMediaPhoto{}
}

func (t *TL_messageMediaPhoto) Encode() []byte {
	return nil
}

func (t *TL_messageMediaPhoto) Decode(b []byte) {}

//====messageMediaGeo#56e0d474====

type TL_messageMediaGeo struct {
	_geo []byte
}

func New_TL_messageMediaGeo() *TL_messageMediaGeo {
	return &TL_messageMediaGeo{}
}

func (t *TL_messageMediaGeo) Encode() []byte {
	return nil
}

func (t *TL_messageMediaGeo) Decode(b []byte) {}

//====messageMediaContact#5e7d2f39====

type TL_messageMediaContact struct {
	_phone_number string
	_first_name   string
	_last_name    string
	_user_id      []byte
}

func New_TL_messageMediaContact() *TL_messageMediaContact {
	return &TL_messageMediaContact{}
}

func (t *TL_messageMediaContact) Encode() []byte {
	return nil
}

func (t *TL_messageMediaContact) Decode(b []byte) {}

//====messageMediaUnsupported#9f84f49e====

type TL_messageMediaUnsupported struct {
}

func New_TL_messageMediaUnsupported() *TL_messageMediaUnsupported {
	return &TL_messageMediaUnsupported{}
}

func (t *TL_messageMediaUnsupported) Encode() []byte {
	return nil
}

func (t *TL_messageMediaUnsupported) Decode(b []byte) {}

//====messageMediaDocument#7c4414d3====

type TL_messageMediaDocument struct {
	_flags       []byte
	_document    []byte
	_caption     []byte
	_ttl_seconds []byte
}

func New_TL_messageMediaDocument() *TL_messageMediaDocument {
	return &TL_messageMediaDocument{}
}

func (t *TL_messageMediaDocument) Encode() []byte {
	return nil
}

func (t *TL_messageMediaDocument) Decode(b []byte) {}

//====messageMediaWebPage#a32dd600====

type TL_messageMediaWebPage struct {
	_webpage []byte
}

func New_TL_messageMediaWebPage() *TL_messageMediaWebPage {
	return &TL_messageMediaWebPage{}
}

func (t *TL_messageMediaWebPage) Encode() []byte {
	return nil
}

func (t *TL_messageMediaWebPage) Decode(b []byte) {}

//====messageMediaVenue#2ec0533f====

type TL_messageMediaVenue struct {
	_geo        []byte
	_title      string
	_address    string
	_provider   string
	_venue_id   string
	_venue_type string
}

func New_TL_messageMediaVenue() *TL_messageMediaVenue {
	return &TL_messageMediaVenue{}
}

func (t *TL_messageMediaVenue) Encode() []byte {
	return nil
}

func (t *TL_messageMediaVenue) Decode(b []byte) {}

//====messageMediaGame#fdb19008====

type TL_messageMediaGame struct {
	_game []byte
}

func New_TL_messageMediaGame() *TL_messageMediaGame {
	return &TL_messageMediaGame{}
}

func (t *TL_messageMediaGame) Encode() []byte {
	return nil
}

func (t *TL_messageMediaGame) Decode(b []byte) {}

//====messageMediaInvoice#84551347====

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

func New_TL_messageMediaInvoice() *TL_messageMediaInvoice {
	return &TL_messageMediaInvoice{}
}

func (t *TL_messageMediaInvoice) Encode() []byte {
	return nil
}

func (t *TL_messageMediaInvoice) Decode(b []byte) {}

//====messageMediaGeoLive#7c3c2609====

type TL_messageMediaGeoLive struct {
	_geo    []byte
	_period []byte
}

func New_TL_messageMediaGeoLive() *TL_messageMediaGeoLive {
	return &TL_messageMediaGeoLive{}
}

func (t *TL_messageMediaGeoLive) Encode() []byte {
	return nil
}

func (t *TL_messageMediaGeoLive) Decode(b []byte) {}

//====messageActionEmpty#b6aef7b0====

type TL_messageActionEmpty struct {
}

func New_TL_messageActionEmpty() *TL_messageActionEmpty {
	return &TL_messageActionEmpty{}
}

func (t *TL_messageActionEmpty) Encode() []byte {
	return nil
}

func (t *TL_messageActionEmpty) Decode(b []byte) {}

//====messageActionChatCreate#a6638b9a====

type TL_messageActionChatCreate struct {
	_title string
	_users []byte
}

func New_TL_messageActionChatCreate() *TL_messageActionChatCreate {
	return &TL_messageActionChatCreate{}
}

func (t *TL_messageActionChatCreate) Encode() []byte {
	return nil
}

func (t *TL_messageActionChatCreate) Decode(b []byte) {}

//====messageActionChatEditTitle#b5a1ce5a====

type TL_messageActionChatEditTitle struct {
	_title string
}

func New_TL_messageActionChatEditTitle() *TL_messageActionChatEditTitle {
	return &TL_messageActionChatEditTitle{}
}

func (t *TL_messageActionChatEditTitle) Encode() []byte {
	return nil
}

func (t *TL_messageActionChatEditTitle) Decode(b []byte) {}

//====messageActionChatEditPhoto#7fcb13a8====

type TL_messageActionChatEditPhoto struct {
	_photo []byte
}

func New_TL_messageActionChatEditPhoto() *TL_messageActionChatEditPhoto {
	return &TL_messageActionChatEditPhoto{}
}

func (t *TL_messageActionChatEditPhoto) Encode() []byte {
	return nil
}

func (t *TL_messageActionChatEditPhoto) Decode(b []byte) {}

//====messageActionChatDeletePhoto#95e3fbef====

type TL_messageActionChatDeletePhoto struct {
}

func New_TL_messageActionChatDeletePhoto() *TL_messageActionChatDeletePhoto {
	return &TL_messageActionChatDeletePhoto{}
}

func (t *TL_messageActionChatDeletePhoto) Encode() []byte {
	return nil
}

func (t *TL_messageActionChatDeletePhoto) Decode(b []byte) {}

//====messageActionChatAddUser#488a7337====

type TL_messageActionChatAddUser struct {
	_users []byte
}

func New_TL_messageActionChatAddUser() *TL_messageActionChatAddUser {
	return &TL_messageActionChatAddUser{}
}

func (t *TL_messageActionChatAddUser) Encode() []byte {
	return nil
}

func (t *TL_messageActionChatAddUser) Decode(b []byte) {}

//====messageActionChatDeleteUser#b2ae9b0c====

type TL_messageActionChatDeleteUser struct {
	_user_id []byte
}

func New_TL_messageActionChatDeleteUser() *TL_messageActionChatDeleteUser {
	return &TL_messageActionChatDeleteUser{}
}

func (t *TL_messageActionChatDeleteUser) Encode() []byte {
	return nil
}

func (t *TL_messageActionChatDeleteUser) Decode(b []byte) {}

//====messageActionChatJoinedByLink#f89cf5e8====

type TL_messageActionChatJoinedByLink struct {
	_inviter_id []byte
}

func New_TL_messageActionChatJoinedByLink() *TL_messageActionChatJoinedByLink {
	return &TL_messageActionChatJoinedByLink{}
}

func (t *TL_messageActionChatJoinedByLink) Encode() []byte {
	return nil
}

func (t *TL_messageActionChatJoinedByLink) Decode(b []byte) {}

//====messageActionChannelCreate#95d2ac92====

type TL_messageActionChannelCreate struct {
	_title string
}

func New_TL_messageActionChannelCreate() *TL_messageActionChannelCreate {
	return &TL_messageActionChannelCreate{}
}

func (t *TL_messageActionChannelCreate) Encode() []byte {
	return nil
}

func (t *TL_messageActionChannelCreate) Decode(b []byte) {}

//====messageActionChatMigrateTo#51bdb021====

type TL_messageActionChatMigrateTo struct {
	_channel_id []byte
}

func New_TL_messageActionChatMigrateTo() *TL_messageActionChatMigrateTo {
	return &TL_messageActionChatMigrateTo{}
}

func (t *TL_messageActionChatMigrateTo) Encode() []byte {
	return nil
}

func (t *TL_messageActionChatMigrateTo) Decode(b []byte) {}

//====messageActionChannelMigrateFrom#b055eaee====

type TL_messageActionChannelMigrateFrom struct {
	_title   string
	_chat_id []byte
}

func New_TL_messageActionChannelMigrateFrom() *TL_messageActionChannelMigrateFrom {
	return &TL_messageActionChannelMigrateFrom{}
}

func (t *TL_messageActionChannelMigrateFrom) Encode() []byte {
	return nil
}

func (t *TL_messageActionChannelMigrateFrom) Decode(b []byte) {}

//====messageActionPinMessage#94bd38ed====

type TL_messageActionPinMessage struct {
}

func New_TL_messageActionPinMessage() *TL_messageActionPinMessage {
	return &TL_messageActionPinMessage{}
}

func (t *TL_messageActionPinMessage) Encode() []byte {
	return nil
}

func (t *TL_messageActionPinMessage) Decode(b []byte) {}

//====messageActionHistoryClear#9fbab604====

type TL_messageActionHistoryClear struct {
}

func New_TL_messageActionHistoryClear() *TL_messageActionHistoryClear {
	return &TL_messageActionHistoryClear{}
}

func (t *TL_messageActionHistoryClear) Encode() []byte {
	return nil
}

func (t *TL_messageActionHistoryClear) Decode(b []byte) {}

//====messageActionGameScore#92a72876====

type TL_messageActionGameScore struct {
	_game_id int64
	_score   []byte
}

func New_TL_messageActionGameScore() *TL_messageActionGameScore {
	return &TL_messageActionGameScore{}
}

func (t *TL_messageActionGameScore) Encode() []byte {
	return nil
}

func (t *TL_messageActionGameScore) Decode(b []byte) {}

//====messageActionPaymentSentMe#8f31b327====

type TL_messageActionPaymentSentMe struct {
	_flags              []byte
	_currency           string
	_total_amount       int64
	_payload            []byte
	_info               []byte
	_shipping_option_id []byte
	_charge             []byte
}

func New_TL_messageActionPaymentSentMe() *TL_messageActionPaymentSentMe {
	return &TL_messageActionPaymentSentMe{}
}

func (t *TL_messageActionPaymentSentMe) Encode() []byte {
	return nil
}

func (t *TL_messageActionPaymentSentMe) Decode(b []byte) {}

//====messageActionPaymentSent#40699cd0====

type TL_messageActionPaymentSent struct {
	_currency     string
	_total_amount int64
}

func New_TL_messageActionPaymentSent() *TL_messageActionPaymentSent {
	return &TL_messageActionPaymentSent{}
}

func (t *TL_messageActionPaymentSent) Encode() []byte {
	return nil
}

func (t *TL_messageActionPaymentSent) Decode(b []byte) {}

//====messageActionPhoneCall#80e11a7f====

type TL_messageActionPhoneCall struct {
	_flags    []byte
	_call_id  int64
	_reason   []byte
	_duration []byte
}

func New_TL_messageActionPhoneCall() *TL_messageActionPhoneCall {
	return &TL_messageActionPhoneCall{}
}

func (t *TL_messageActionPhoneCall) Encode() []byte {
	return nil
}

func (t *TL_messageActionPhoneCall) Decode(b []byte) {}

//====messageActionScreenshotTaken#4792929b====

type TL_messageActionScreenshotTaken struct {
}

func New_TL_messageActionScreenshotTaken() *TL_messageActionScreenshotTaken {
	return &TL_messageActionScreenshotTaken{}
}

func (t *TL_messageActionScreenshotTaken) Encode() []byte {
	return nil
}

func (t *TL_messageActionScreenshotTaken) Decode(b []byte) {}

//====messageActionCustomAction#fae69f56====

type TL_messageActionCustomAction struct {
	_message string
}

func New_TL_messageActionCustomAction() *TL_messageActionCustomAction {
	return &TL_messageActionCustomAction{}
}

func (t *TL_messageActionCustomAction) Encode() []byte {
	return nil
}

func (t *TL_messageActionCustomAction) Decode(b []byte) {}

//====dialog#e4def5db====

type TL_dialog struct {
	_flags                 []byte
	_pinned                []byte
	_peer                  []byte
	_top_message           []byte
	_read_inbox_max_id     []byte
	_read_outbox_max_id    []byte
	_unread_count          []byte
	_unread_mentions_count []byte
	_notify_settings       []byte
	_pts                   []byte
	_draft                 []byte
}

func New_TL_dialog() *TL_dialog {
	return &TL_dialog{}
}

func (t *TL_dialog) Encode() []byte {
	return nil
}

func (t *TL_dialog) Decode(b []byte) {}

//====photoEmpty#2331b22d====

type TL_photoEmpty struct {
	_id int64
}

func New_TL_photoEmpty() *TL_photoEmpty {
	return &TL_photoEmpty{}
}

func (t *TL_photoEmpty) Encode() []byte {
	return nil
}

func (t *TL_photoEmpty) Decode(b []byte) {}

//====photo#9288dd29====

type TL_photo struct {
	_flags        []byte
	_has_stickers []byte
	_id           int64
	_access_hash  int64
	_date         []byte
	_sizes        []byte
}

func New_TL_photo() *TL_photo {
	return &TL_photo{}
}

func (t *TL_photo) Encode() []byte {
	return nil
}

func (t *TL_photo) Decode(b []byte) {}

//====photoSizeEmpty#0e17e23c====

type TL_photoSizeEmpty struct {
	_type string
}

func New_TL_photoSizeEmpty() *TL_photoSizeEmpty {
	return &TL_photoSizeEmpty{}
}

func (t *TL_photoSizeEmpty) Encode() []byte {
	return nil
}

func (t *TL_photoSizeEmpty) Decode(b []byte) {}

//====photoSize#77bfb61b====

type TL_photoSize struct {
	_type     string
	_location []byte
	_w        []byte
	_h        []byte
	_size     []byte
}

func New_TL_photoSize() *TL_photoSize {
	return &TL_photoSize{}
}

func (t *TL_photoSize) Encode() []byte {
	return nil
}

func (t *TL_photoSize) Decode(b []byte) {}

//====photoCachedSize#e9a734fa====

type TL_photoCachedSize struct {
	_type     string
	_location []byte
	_w        []byte
	_h        []byte
	_bytes    []byte
}

func New_TL_photoCachedSize() *TL_photoCachedSize {
	return &TL_photoCachedSize{}
}

func (t *TL_photoCachedSize) Encode() []byte {
	return nil
}

func (t *TL_photoCachedSize) Decode(b []byte) {}

//====geoPointEmpty#1117dd5f====

type TL_geoPointEmpty struct {
}

func New_TL_geoPointEmpty() *TL_geoPointEmpty {
	return &TL_geoPointEmpty{}
}

func (t *TL_geoPointEmpty) Encode() []byte {
	return nil
}

func (t *TL_geoPointEmpty) Decode(b []byte) {}

//====geoPoint#2049d70c====

type TL_geoPoint struct {
	_long []byte
	_lat  []byte
}

func New_TL_geoPoint() *TL_geoPoint {
	return &TL_geoPoint{}
}

func (t *TL_geoPoint) Encode() []byte {
	return nil
}

func (t *TL_geoPoint) Decode(b []byte) {}

//====auth_checkedPhone#811ea28e====

type TL_auth_checkedPhone struct {
	_phone_registered []byte
}

func New_TL_auth_checkedPhone() *TL_auth_checkedPhone {
	return &TL_auth_checkedPhone{}
}

func (t *TL_auth_checkedPhone) Encode() []byte {
	return nil
}

func (t *TL_auth_checkedPhone) Decode(b []byte) {}

//====auth_sentCode#5e002502====

type TL_auth_sentCode struct {
	_flags            []byte
	_phone_registered []byte
	_type             []byte
	_phone_code_hash  string
	_next_type        []byte
	_timeout          []byte
}

func New_TL_auth_sentCode() *TL_auth_sentCode {
	return &TL_auth_sentCode{}
}

func (t *TL_auth_sentCode) Encode() []byte {
	return nil
}

func (t *TL_auth_sentCode) Decode(b []byte) {}

//====auth_authorization#cd050916====

type TL_auth_authorization struct {
	_flags        []byte
	_tmp_sessions []byte
	_user         []byte
}

func New_TL_auth_authorization() *TL_auth_authorization {
	return &TL_auth_authorization{}
}

func (t *TL_auth_authorization) Encode() []byte {
	return nil
}

func (t *TL_auth_authorization) Decode(b []byte) {}

//====auth_exportedAuthorization#df969c2d====

type TL_auth_exportedAuthorization struct {
	_id    []byte
	_bytes []byte
}

func New_TL_auth_exportedAuthorization() *TL_auth_exportedAuthorization {
	return &TL_auth_exportedAuthorization{}
}

func (t *TL_auth_exportedAuthorization) Encode() []byte {
	return nil
}

func (t *TL_auth_exportedAuthorization) Decode(b []byte) {}

//====inputNotifyPeer#b8bc5b0c====

type TL_inputNotifyPeer struct {
	_peer []byte
}

func New_TL_inputNotifyPeer() *TL_inputNotifyPeer {
	return &TL_inputNotifyPeer{}
}

func (t *TL_inputNotifyPeer) Encode() []byte {
	return nil
}

func (t *TL_inputNotifyPeer) Decode(b []byte) {}

//====inputNotifyUsers#193b4417====

type TL_inputNotifyUsers struct {
}

func New_TL_inputNotifyUsers() *TL_inputNotifyUsers {
	return &TL_inputNotifyUsers{}
}

func (t *TL_inputNotifyUsers) Encode() []byte {
	return nil
}

func (t *TL_inputNotifyUsers) Decode(b []byte) {}

//====inputNotifyChats#4a95e84e====

type TL_inputNotifyChats struct {
}

func New_TL_inputNotifyChats() *TL_inputNotifyChats {
	return &TL_inputNotifyChats{}
}

func (t *TL_inputNotifyChats) Encode() []byte {
	return nil
}

func (t *TL_inputNotifyChats) Decode(b []byte) {}

//====inputNotifyAll#a429b886====

type TL_inputNotifyAll struct {
}

func New_TL_inputNotifyAll() *TL_inputNotifyAll {
	return &TL_inputNotifyAll{}
}

func (t *TL_inputNotifyAll) Encode() []byte {
	return nil
}

func (t *TL_inputNotifyAll) Decode(b []byte) {}

//====inputPeerNotifyEventsEmpty#f03064d8====

type TL_inputPeerNotifyEventsEmpty struct {
}

func New_TL_inputPeerNotifyEventsEmpty() *TL_inputPeerNotifyEventsEmpty {
	return &TL_inputPeerNotifyEventsEmpty{}
}

func (t *TL_inputPeerNotifyEventsEmpty) Encode() []byte {
	return nil
}

func (t *TL_inputPeerNotifyEventsEmpty) Decode(b []byte) {}

//====inputPeerNotifyEventsAll#e86a2c74====

type TL_inputPeerNotifyEventsAll struct {
}

func New_TL_inputPeerNotifyEventsAll() *TL_inputPeerNotifyEventsAll {
	return &TL_inputPeerNotifyEventsAll{}
}

func (t *TL_inputPeerNotifyEventsAll) Encode() []byte {
	return nil
}

func (t *TL_inputPeerNotifyEventsAll) Decode(b []byte) {}

//====inputPeerNotifySettings#38935eb2====

type TL_inputPeerNotifySettings struct {
	_flags         []byte
	_show_previews []byte
	_silent        []byte
	_mute_until    []byte
	_sound         string
}

func New_TL_inputPeerNotifySettings() *TL_inputPeerNotifySettings {
	return &TL_inputPeerNotifySettings{}
}

func (t *TL_inputPeerNotifySettings) Encode() []byte {
	return nil
}

func (t *TL_inputPeerNotifySettings) Decode(b []byte) {}

//====peerNotifyEventsEmpty#add53cb3====

type TL_peerNotifyEventsEmpty struct {
}

func New_TL_peerNotifyEventsEmpty() *TL_peerNotifyEventsEmpty {
	return &TL_peerNotifyEventsEmpty{}
}

func (t *TL_peerNotifyEventsEmpty) Encode() []byte {
	return nil
}

func (t *TL_peerNotifyEventsEmpty) Decode(b []byte) {}

//====peerNotifyEventsAll#6d1ded88====

type TL_peerNotifyEventsAll struct {
}

func New_TL_peerNotifyEventsAll() *TL_peerNotifyEventsAll {
	return &TL_peerNotifyEventsAll{}
}

func (t *TL_peerNotifyEventsAll) Encode() []byte {
	return nil
}

func (t *TL_peerNotifyEventsAll) Decode(b []byte) {}

//====peerNotifySettingsEmpty#70a68512====

type TL_peerNotifySettingsEmpty struct {
}

func New_TL_peerNotifySettingsEmpty() *TL_peerNotifySettingsEmpty {
	return &TL_peerNotifySettingsEmpty{}
}

func (t *TL_peerNotifySettingsEmpty) Encode() []byte {
	return nil
}

func (t *TL_peerNotifySettingsEmpty) Decode(b []byte) {}

//====peerNotifySettings#9acda4c0====

type TL_peerNotifySettings struct {
	_flags         []byte
	_show_previews []byte
	_silent        []byte
	_mute_until    []byte
	_sound         string
}

func New_TL_peerNotifySettings() *TL_peerNotifySettings {
	return &TL_peerNotifySettings{}
}

func (t *TL_peerNotifySettings) Encode() []byte {
	return nil
}

func (t *TL_peerNotifySettings) Decode(b []byte) {}

//====peerSettings#818426cd====

type TL_peerSettings struct {
	_flags       []byte
	_report_spam []byte
}

func New_TL_peerSettings() *TL_peerSettings {
	return &TL_peerSettings{}
}

func (t *TL_peerSettings) Encode() []byte {
	return nil
}

func (t *TL_peerSettings) Decode(b []byte) {}

//====wallPaper#ccb03657====

type TL_wallPaper struct {
	_id    []byte
	_title string
	_sizes []byte
	_color []byte
}

func New_TL_wallPaper() *TL_wallPaper {
	return &TL_wallPaper{}
}

func (t *TL_wallPaper) Encode() []byte {
	return nil
}

func (t *TL_wallPaper) Decode(b []byte) {}

//====wallPaperSolid#63117f24====

type TL_wallPaperSolid struct {
	_id       []byte
	_title    string
	_bg_color []byte
	_color    []byte
}

func New_TL_wallPaperSolid() *TL_wallPaperSolid {
	return &TL_wallPaperSolid{}
}

func (t *TL_wallPaperSolid) Encode() []byte {
	return nil
}

func (t *TL_wallPaperSolid) Decode(b []byte) {}

//====inputReportReasonSpam#58dbcab8====

type TL_inputReportReasonSpam struct {
}

func New_TL_inputReportReasonSpam() *TL_inputReportReasonSpam {
	return &TL_inputReportReasonSpam{}
}

func (t *TL_inputReportReasonSpam) Encode() []byte {
	return nil
}

func (t *TL_inputReportReasonSpam) Decode(b []byte) {}

//====inputReportReasonViolence#1e22c78d====

type TL_inputReportReasonViolence struct {
}

func New_TL_inputReportReasonViolence() *TL_inputReportReasonViolence {
	return &TL_inputReportReasonViolence{}
}

func (t *TL_inputReportReasonViolence) Encode() []byte {
	return nil
}

func (t *TL_inputReportReasonViolence) Decode(b []byte) {}

//====inputReportReasonPornography#2e59d922====

type TL_inputReportReasonPornography struct {
}

func New_TL_inputReportReasonPornography() *TL_inputReportReasonPornography {
	return &TL_inputReportReasonPornography{}
}

func (t *TL_inputReportReasonPornography) Encode() []byte {
	return nil
}

func (t *TL_inputReportReasonPornography) Decode(b []byte) {}

//====inputReportReasonOther#e1746d0a====

type TL_inputReportReasonOther struct {
	_text string
}

func New_TL_inputReportReasonOther() *TL_inputReportReasonOther {
	return &TL_inputReportReasonOther{}
}

func (t *TL_inputReportReasonOther) Encode() []byte {
	return nil
}

func (t *TL_inputReportReasonOther) Decode(b []byte) {}

//====userFull#0f220f3f====

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
	_common_chats_count    []byte
}

func New_TL_userFull() *TL_userFull {
	return &TL_userFull{}
}

func (t *TL_userFull) Encode() []byte {
	return nil
}

func (t *TL_userFull) Decode(b []byte) {}

//====contact#f911c994====

type TL_contact struct {
	_user_id []byte
	_mutual  []byte
}

func New_TL_contact() *TL_contact {
	return &TL_contact{}
}

func (t *TL_contact) Encode() []byte {
	return nil
}

func (t *TL_contact) Decode(b []byte) {}

//====importedContact#d0028438====

type TL_importedContact struct {
	_user_id   []byte
	_client_id int64
}

func New_TL_importedContact() *TL_importedContact {
	return &TL_importedContact{}
}

func (t *TL_importedContact) Encode() []byte {
	return nil
}

func (t *TL_importedContact) Decode(b []byte) {}

//====contactBlocked#561bc879====

type TL_contactBlocked struct {
	_user_id []byte
	_date    []byte
}

func New_TL_contactBlocked() *TL_contactBlocked {
	return &TL_contactBlocked{}
}

func (t *TL_contactBlocked) Encode() []byte {
	return nil
}

func (t *TL_contactBlocked) Decode(b []byte) {}

//====contactStatus#d3680c61====

type TL_contactStatus struct {
	_user_id []byte
	_status  []byte
}

func New_TL_contactStatus() *TL_contactStatus {
	return &TL_contactStatus{}
}

func (t *TL_contactStatus) Encode() []byte {
	return nil
}

func (t *TL_contactStatus) Decode(b []byte) {}

//====contacts_link#3ace484c====

type TL_contacts_link struct {
	_my_link      []byte
	_foreign_link []byte
	_user         []byte
}

func New_TL_contacts_link() *TL_contacts_link {
	return &TL_contacts_link{}
}

func (t *TL_contacts_link) Encode() []byte {
	return nil
}

func (t *TL_contacts_link) Decode(b []byte) {}

//====contacts_contactsNotModified#b74ba9d2====

type TL_contacts_contactsNotModified struct {
}

func New_TL_contacts_contactsNotModified() *TL_contacts_contactsNotModified {
	return &TL_contacts_contactsNotModified{}
}

func (t *TL_contacts_contactsNotModified) Encode() []byte {
	return nil
}

func (t *TL_contacts_contactsNotModified) Decode(b []byte) {}

//====contacts_contacts#eae87e42====

type TL_contacts_contacts struct {
	_contacts    []byte
	_saved_count []byte
	_users       []byte
}

func New_TL_contacts_contacts() *TL_contacts_contacts {
	return &TL_contacts_contacts{}
}

func (t *TL_contacts_contacts) Encode() []byte {
	return nil
}

func (t *TL_contacts_contacts) Decode(b []byte) {}

//====contacts_importedContacts#77d01c3b====

type TL_contacts_importedContacts struct {
	_imported        []byte
	_popular_invites []byte
	_retry_contacts  []byte
	_users           []byte
}

func New_TL_contacts_importedContacts() *TL_contacts_importedContacts {
	return &TL_contacts_importedContacts{}
}

func (t *TL_contacts_importedContacts) Encode() []byte {
	return nil
}

func (t *TL_contacts_importedContacts) Decode(b []byte) {}

//====contacts_blocked#1c138d15====

type TL_contacts_blocked struct {
	_blocked []byte
	_users   []byte
}

func New_TL_contacts_blocked() *TL_contacts_blocked {
	return &TL_contacts_blocked{}
}

func (t *TL_contacts_blocked) Encode() []byte {
	return nil
}

func (t *TL_contacts_blocked) Decode(b []byte) {}

//====contacts_blockedSlice#900802a1====

type TL_contacts_blockedSlice struct {
	_count   []byte
	_blocked []byte
	_users   []byte
}

func New_TL_contacts_blockedSlice() *TL_contacts_blockedSlice {
	return &TL_contacts_blockedSlice{}
}

func (t *TL_contacts_blockedSlice) Encode() []byte {
	return nil
}

func (t *TL_contacts_blockedSlice) Decode(b []byte) {}

//====messages_dialogs#15ba6c40====

type TL_messages_dialogs struct {
	_dialogs  []byte
	_messages []byte
	_chats    []byte
	_users    []byte
}

func New_TL_messages_dialogs() *TL_messages_dialogs {
	return &TL_messages_dialogs{}
}

func (t *TL_messages_dialogs) Encode() []byte {
	return nil
}

func (t *TL_messages_dialogs) Decode(b []byte) {}

//====messages_dialogsSlice#71e094f3====

type TL_messages_dialogsSlice struct {
	_count    []byte
	_dialogs  []byte
	_messages []byte
	_chats    []byte
	_users    []byte
}

func New_TL_messages_dialogsSlice() *TL_messages_dialogsSlice {
	return &TL_messages_dialogsSlice{}
}

func (t *TL_messages_dialogsSlice) Encode() []byte {
	return nil
}

func (t *TL_messages_dialogsSlice) Decode(b []byte) {}

//====messages_messages#8c718e87====

type TL_messages_messages struct {
	_messages []byte
	_chats    []byte
	_users    []byte
}

func New_TL_messages_messages() *TL_messages_messages {
	return &TL_messages_messages{}
}

func (t *TL_messages_messages) Encode() []byte {
	return nil
}

func (t *TL_messages_messages) Decode(b []byte) {}

//====messages_messagesSlice#0b446ae3====

type TL_messages_messagesSlice struct {
	_count    []byte
	_messages []byte
	_chats    []byte
	_users    []byte
}

func New_TL_messages_messagesSlice() *TL_messages_messagesSlice {
	return &TL_messages_messagesSlice{}
}

func (t *TL_messages_messagesSlice) Encode() []byte {
	return nil
}

func (t *TL_messages_messagesSlice) Decode(b []byte) {}

//====messages_channelMessages#99262e37====

type TL_messages_channelMessages struct {
	_flags    []byte
	_pts      []byte
	_count    []byte
	_messages []byte
	_chats    []byte
	_users    []byte
}

func New_TL_messages_channelMessages() *TL_messages_channelMessages {
	return &TL_messages_channelMessages{}
}

func (t *TL_messages_channelMessages) Encode() []byte {
	return nil
}

func (t *TL_messages_channelMessages) Decode(b []byte) {}

//====messages_messagesNotModified#74535f21====

type TL_messages_messagesNotModified struct {
	_count []byte
}

func New_TL_messages_messagesNotModified() *TL_messages_messagesNotModified {
	return &TL_messages_messagesNotModified{}
}

func (t *TL_messages_messagesNotModified) Encode() []byte {
	return nil
}

func (t *TL_messages_messagesNotModified) Decode(b []byte) {}

//====messages_chats#64ff9fd5====

type TL_messages_chats struct {
	_chats []byte
}

func New_TL_messages_chats() *TL_messages_chats {
	return &TL_messages_chats{}
}

func (t *TL_messages_chats) Encode() []byte {
	return nil
}

func (t *TL_messages_chats) Decode(b []byte) {}

//====messages_chatsSlice#9cd81144====

type TL_messages_chatsSlice struct {
	_count []byte
	_chats []byte
}

func New_TL_messages_chatsSlice() *TL_messages_chatsSlice {
	return &TL_messages_chatsSlice{}
}

func (t *TL_messages_chatsSlice) Encode() []byte {
	return nil
}

func (t *TL_messages_chatsSlice) Decode(b []byte) {}

//====messages_chatFull#e5d7d19c====

type TL_messages_chatFull struct {
	_full_chat []byte
	_chats     []byte
	_users     []byte
}

func New_TL_messages_chatFull() *TL_messages_chatFull {
	return &TL_messages_chatFull{}
}

func (t *TL_messages_chatFull) Encode() []byte {
	return nil
}

func (t *TL_messages_chatFull) Decode(b []byte) {}

//====messages_affectedHistory#b45c69d1====

type TL_messages_affectedHistory struct {
	_pts       []byte
	_pts_count []byte
	_offset    []byte
}

func New_TL_messages_affectedHistory() *TL_messages_affectedHistory {
	return &TL_messages_affectedHistory{}
}

func (t *TL_messages_affectedHistory) Encode() []byte {
	return nil
}

func (t *TL_messages_affectedHistory) Decode(b []byte) {}

//====inputMessagesFilterEmpty#57e2f66c====

type TL_inputMessagesFilterEmpty struct {
}

func New_TL_inputMessagesFilterEmpty() *TL_inputMessagesFilterEmpty {
	return &TL_inputMessagesFilterEmpty{}
}

func (t *TL_inputMessagesFilterEmpty) Encode() []byte {
	return nil
}

func (t *TL_inputMessagesFilterEmpty) Decode(b []byte) {}

//====inputMessagesFilterPhotos#9609a51c====

type TL_inputMessagesFilterPhotos struct {
}

func New_TL_inputMessagesFilterPhotos() *TL_inputMessagesFilterPhotos {
	return &TL_inputMessagesFilterPhotos{}
}

func (t *TL_inputMessagesFilterPhotos) Encode() []byte {
	return nil
}

func (t *TL_inputMessagesFilterPhotos) Decode(b []byte) {}

//====inputMessagesFilterVideo#9fc00e65====

type TL_inputMessagesFilterVideo struct {
}

func New_TL_inputMessagesFilterVideo() *TL_inputMessagesFilterVideo {
	return &TL_inputMessagesFilterVideo{}
}

func (t *TL_inputMessagesFilterVideo) Encode() []byte {
	return nil
}

func (t *TL_inputMessagesFilterVideo) Decode(b []byte) {}

//====inputMessagesFilterPhotoVideo#56e9f0e4====

type TL_inputMessagesFilterPhotoVideo struct {
}

func New_TL_inputMessagesFilterPhotoVideo() *TL_inputMessagesFilterPhotoVideo {
	return &TL_inputMessagesFilterPhotoVideo{}
}

func (t *TL_inputMessagesFilterPhotoVideo) Encode() []byte {
	return nil
}

func (t *TL_inputMessagesFilterPhotoVideo) Decode(b []byte) {}

//====inputMessagesFilterDocument#9eddf188====

type TL_inputMessagesFilterDocument struct {
}

func New_TL_inputMessagesFilterDocument() *TL_inputMessagesFilterDocument {
	return &TL_inputMessagesFilterDocument{}
}

func (t *TL_inputMessagesFilterDocument) Encode() []byte {
	return nil
}

func (t *TL_inputMessagesFilterDocument) Decode(b []byte) {}

//====inputMessagesFilterUrl#7ef0dd87====

type TL_inputMessagesFilterUrl struct {
}

func New_TL_inputMessagesFilterUrl() *TL_inputMessagesFilterUrl {
	return &TL_inputMessagesFilterUrl{}
}

func (t *TL_inputMessagesFilterUrl) Encode() []byte {
	return nil
}

func (t *TL_inputMessagesFilterUrl) Decode(b []byte) {}

//====inputMessagesFilterGif#ffc86587====

type TL_inputMessagesFilterGif struct {
}

func New_TL_inputMessagesFilterGif() *TL_inputMessagesFilterGif {
	return &TL_inputMessagesFilterGif{}
}

func (t *TL_inputMessagesFilterGif) Encode() []byte {
	return nil
}

func (t *TL_inputMessagesFilterGif) Decode(b []byte) {}

//====inputMessagesFilterVoice#50f5c392====

type TL_inputMessagesFilterVoice struct {
}

func New_TL_inputMessagesFilterVoice() *TL_inputMessagesFilterVoice {
	return &TL_inputMessagesFilterVoice{}
}

func (t *TL_inputMessagesFilterVoice) Encode() []byte {
	return nil
}

func (t *TL_inputMessagesFilterVoice) Decode(b []byte) {}

//====inputMessagesFilterMusic#3751b49e====

type TL_inputMessagesFilterMusic struct {
}

func New_TL_inputMessagesFilterMusic() *TL_inputMessagesFilterMusic {
	return &TL_inputMessagesFilterMusic{}
}

func (t *TL_inputMessagesFilterMusic) Encode() []byte {
	return nil
}

func (t *TL_inputMessagesFilterMusic) Decode(b []byte) {}

//====inputMessagesFilterChatPhotos#3a20ecb8====

type TL_inputMessagesFilterChatPhotos struct {
}

func New_TL_inputMessagesFilterChatPhotos() *TL_inputMessagesFilterChatPhotos {
	return &TL_inputMessagesFilterChatPhotos{}
}

func (t *TL_inputMessagesFilterChatPhotos) Encode() []byte {
	return nil
}

func (t *TL_inputMessagesFilterChatPhotos) Decode(b []byte) {}

//====inputMessagesFilterPhoneCalls#80c99768====

type TL_inputMessagesFilterPhoneCalls struct {
	_flags  []byte
	_missed []byte
}

func New_TL_inputMessagesFilterPhoneCalls() *TL_inputMessagesFilterPhoneCalls {
	return &TL_inputMessagesFilterPhoneCalls{}
}

func (t *TL_inputMessagesFilterPhoneCalls) Encode() []byte {
	return nil
}

func (t *TL_inputMessagesFilterPhoneCalls) Decode(b []byte) {}

//====inputMessagesFilterRoundVoice#7a7c17a4====

type TL_inputMessagesFilterRoundVoice struct {
}

func New_TL_inputMessagesFilterRoundVoice() *TL_inputMessagesFilterRoundVoice {
	return &TL_inputMessagesFilterRoundVoice{}
}

func (t *TL_inputMessagesFilterRoundVoice) Encode() []byte {
	return nil
}

func (t *TL_inputMessagesFilterRoundVoice) Decode(b []byte) {}

//====inputMessagesFilterRoundVideo#b549da53====

type TL_inputMessagesFilterRoundVideo struct {
}

func New_TL_inputMessagesFilterRoundVideo() *TL_inputMessagesFilterRoundVideo {
	return &TL_inputMessagesFilterRoundVideo{}
}

func (t *TL_inputMessagesFilterRoundVideo) Encode() []byte {
	return nil
}

func (t *TL_inputMessagesFilterRoundVideo) Decode(b []byte) {}

//====inputMessagesFilterMyMentions#c1f8e69a====

type TL_inputMessagesFilterMyMentions struct {
}

func New_TL_inputMessagesFilterMyMentions() *TL_inputMessagesFilterMyMentions {
	return &TL_inputMessagesFilterMyMentions{}
}

func (t *TL_inputMessagesFilterMyMentions) Encode() []byte {
	return nil
}

func (t *TL_inputMessagesFilterMyMentions) Decode(b []byte) {}

//====inputMessagesFilterGeo#e7026d0d====

type TL_inputMessagesFilterGeo struct {
}

func New_TL_inputMessagesFilterGeo() *TL_inputMessagesFilterGeo {
	return &TL_inputMessagesFilterGeo{}
}

func (t *TL_inputMessagesFilterGeo) Encode() []byte {
	return nil
}

func (t *TL_inputMessagesFilterGeo) Decode(b []byte) {}

//====inputMessagesFilterContacts#e062db83====

type TL_inputMessagesFilterContacts struct {
}

func New_TL_inputMessagesFilterContacts() *TL_inputMessagesFilterContacts {
	return &TL_inputMessagesFilterContacts{}
}

func (t *TL_inputMessagesFilterContacts) Encode() []byte {
	return nil
}

func (t *TL_inputMessagesFilterContacts) Decode(b []byte) {}

//====updateNewMessage#1f2b0afd====

type TL_updateNewMessage struct {
	_message   []byte
	_pts       []byte
	_pts_count []byte
}

func New_TL_updateNewMessage() *TL_updateNewMessage {
	return &TL_updateNewMessage{}
}

func (t *TL_updateNewMessage) Encode() []byte {
	return nil
}

func (t *TL_updateNewMessage) Decode(b []byte) {}

//====updateMessageID#4e90bfd6====

type TL_updateMessageID struct {
	_id        []byte
	_random_id int64
}

func New_TL_updateMessageID() *TL_updateMessageID {
	return &TL_updateMessageID{}
}

func (t *TL_updateMessageID) Encode() []byte {
	return nil
}

func (t *TL_updateMessageID) Decode(b []byte) {}

//====updateDeleteMessages#a20db0e5====

type TL_updateDeleteMessages struct {
	_messages  []byte
	_pts       []byte
	_pts_count []byte
}

func New_TL_updateDeleteMessages() *TL_updateDeleteMessages {
	return &TL_updateDeleteMessages{}
}

func (t *TL_updateDeleteMessages) Encode() []byte {
	return nil
}

func (t *TL_updateDeleteMessages) Decode(b []byte) {}

//====updateUserTyping#5c486927====

type TL_updateUserTyping struct {
	_user_id []byte
	_action  []byte
}

func New_TL_updateUserTyping() *TL_updateUserTyping {
	return &TL_updateUserTyping{}
}

func (t *TL_updateUserTyping) Encode() []byte {
	return nil
}

func (t *TL_updateUserTyping) Decode(b []byte) {}

//====updateChatUserTyping#9a65ea1f====

type TL_updateChatUserTyping struct {
	_chat_id []byte
	_user_id []byte
	_action  []byte
}

func New_TL_updateChatUserTyping() *TL_updateChatUserTyping {
	return &TL_updateChatUserTyping{}
}

func (t *TL_updateChatUserTyping) Encode() []byte {
	return nil
}

func (t *TL_updateChatUserTyping) Decode(b []byte) {}

//====updateChatParticipants#07761198====

type TL_updateChatParticipants struct {
	_participants []byte
}

func New_TL_updateChatParticipants() *TL_updateChatParticipants {
	return &TL_updateChatParticipants{}
}

func (t *TL_updateChatParticipants) Encode() []byte {
	return nil
}

func (t *TL_updateChatParticipants) Decode(b []byte) {}

//====updateUserStatus#1bfbd823====

type TL_updateUserStatus struct {
	_user_id []byte
	_status  []byte
}

func New_TL_updateUserStatus() *TL_updateUserStatus {
	return &TL_updateUserStatus{}
}

func (t *TL_updateUserStatus) Encode() []byte {
	return nil
}

func (t *TL_updateUserStatus) Decode(b []byte) {}

//====updateUserName#a7332b73====

type TL_updateUserName struct {
	_user_id    []byte
	_first_name string
	_last_name  string
	_username   string
}

func New_TL_updateUserName() *TL_updateUserName {
	return &TL_updateUserName{}
}

func (t *TL_updateUserName) Encode() []byte {
	return nil
}

func (t *TL_updateUserName) Decode(b []byte) {}

//====updateUserPhoto#95313b0c====

type TL_updateUserPhoto struct {
	_user_id  []byte
	_date     []byte
	_photo    []byte
	_previous []byte
}

func New_TL_updateUserPhoto() *TL_updateUserPhoto {
	return &TL_updateUserPhoto{}
}

func (t *TL_updateUserPhoto) Encode() []byte {
	return nil
}

func (t *TL_updateUserPhoto) Decode(b []byte) {}

//====updateContactRegistered#2575bbb9====

type TL_updateContactRegistered struct {
	_user_id []byte
	_date    []byte
}

func New_TL_updateContactRegistered() *TL_updateContactRegistered {
	return &TL_updateContactRegistered{}
}

func (t *TL_updateContactRegistered) Encode() []byte {
	return nil
}

func (t *TL_updateContactRegistered) Decode(b []byte) {}

//====updateContactLink#9d2e67c5====

type TL_updateContactLink struct {
	_user_id      []byte
	_my_link      []byte
	_foreign_link []byte
}

func New_TL_updateContactLink() *TL_updateContactLink {
	return &TL_updateContactLink{}
}

func (t *TL_updateContactLink) Encode() []byte {
	return nil
}

func (t *TL_updateContactLink) Decode(b []byte) {}

//====updateNewEncryptedMessage#12bcbd9a====

type TL_updateNewEncryptedMessage struct {
	_message []byte
	_qts     []byte
}

func New_TL_updateNewEncryptedMessage() *TL_updateNewEncryptedMessage {
	return &TL_updateNewEncryptedMessage{}
}

func (t *TL_updateNewEncryptedMessage) Encode() []byte {
	return nil
}

func (t *TL_updateNewEncryptedMessage) Decode(b []byte) {}

//====updateEncryptedChatTyping#1710f156====

type TL_updateEncryptedChatTyping struct {
	_chat_id []byte
}

func New_TL_updateEncryptedChatTyping() *TL_updateEncryptedChatTyping {
	return &TL_updateEncryptedChatTyping{}
}

func (t *TL_updateEncryptedChatTyping) Encode() []byte {
	return nil
}

func (t *TL_updateEncryptedChatTyping) Decode(b []byte) {}

//====updateEncryption#b4a2e88d====

type TL_updateEncryption struct {
	_chat []byte
	_date []byte
}

func New_TL_updateEncryption() *TL_updateEncryption {
	return &TL_updateEncryption{}
}

func (t *TL_updateEncryption) Encode() []byte {
	return nil
}

func (t *TL_updateEncryption) Decode(b []byte) {}

//====updateEncryptedMessagesRead#38fe25b7====

type TL_updateEncryptedMessagesRead struct {
	_chat_id  []byte
	_max_date []byte
	_date     []byte
}

func New_TL_updateEncryptedMessagesRead() *TL_updateEncryptedMessagesRead {
	return &TL_updateEncryptedMessagesRead{}
}

func (t *TL_updateEncryptedMessagesRead) Encode() []byte {
	return nil
}

func (t *TL_updateEncryptedMessagesRead) Decode(b []byte) {}

//====updateChatParticipantAdd#ea4b0e5c====

type TL_updateChatParticipantAdd struct {
	_chat_id    []byte
	_user_id    []byte
	_inviter_id []byte
	_date       []byte
	_version    []byte
}

func New_TL_updateChatParticipantAdd() *TL_updateChatParticipantAdd {
	return &TL_updateChatParticipantAdd{}
}

func (t *TL_updateChatParticipantAdd) Encode() []byte {
	return nil
}

func (t *TL_updateChatParticipantAdd) Decode(b []byte) {}

//====updateChatParticipantDelete#6e5f8c22====

type TL_updateChatParticipantDelete struct {
	_chat_id []byte
	_user_id []byte
	_version []byte
}

func New_TL_updateChatParticipantDelete() *TL_updateChatParticipantDelete {
	return &TL_updateChatParticipantDelete{}
}

func (t *TL_updateChatParticipantDelete) Encode() []byte {
	return nil
}

func (t *TL_updateChatParticipantDelete) Decode(b []byte) {}

//====updateDcOptions#8e5e9873====

type TL_updateDcOptions struct {
	_dc_options []byte
}

func New_TL_updateDcOptions() *TL_updateDcOptions {
	return &TL_updateDcOptions{}
}

func (t *TL_updateDcOptions) Encode() []byte {
	return nil
}

func (t *TL_updateDcOptions) Decode(b []byte) {}

//====updateUserBlocked#80ece81a====

type TL_updateUserBlocked struct {
	_user_id []byte
	_blocked []byte
}

func New_TL_updateUserBlocked() *TL_updateUserBlocked {
	return &TL_updateUserBlocked{}
}

func (t *TL_updateUserBlocked) Encode() []byte {
	return nil
}

func (t *TL_updateUserBlocked) Decode(b []byte) {}

//====updateNotifySettings#bec268ef====

type TL_updateNotifySettings struct {
	_peer            []byte
	_notify_settings []byte
}

func New_TL_updateNotifySettings() *TL_updateNotifySettings {
	return &TL_updateNotifySettings{}
}

func (t *TL_updateNotifySettings) Encode() []byte {
	return nil
}

func (t *TL_updateNotifySettings) Decode(b []byte) {}

//====updateServiceNotification#ebe46819====

type TL_updateServiceNotification struct {
	_flags      []byte
	_popup      []byte
	_inbox_date []byte
	_type       string
	_message    string
	_media      []byte
	_entities   []byte
}

func New_TL_updateServiceNotification() *TL_updateServiceNotification {
	return &TL_updateServiceNotification{}
}

func (t *TL_updateServiceNotification) Encode() []byte {
	return nil
}

func (t *TL_updateServiceNotification) Decode(b []byte) {}

//====updatePrivacy#ee3b272a====

type TL_updatePrivacy struct {
	_key   []byte
	_rules []byte
}

func New_TL_updatePrivacy() *TL_updatePrivacy {
	return &TL_updatePrivacy{}
}

func (t *TL_updatePrivacy) Encode() []byte {
	return nil
}

func (t *TL_updatePrivacy) Decode(b []byte) {}

//====updateUserPhone#12b9417b====

type TL_updateUserPhone struct {
	_user_id []byte
	_phone   string
}

func New_TL_updateUserPhone() *TL_updateUserPhone {
	return &TL_updateUserPhone{}
}

func (t *TL_updateUserPhone) Encode() []byte {
	return nil
}

func (t *TL_updateUserPhone) Decode(b []byte) {}

//====updateReadHistoryInbox#9961fd5c====

type TL_updateReadHistoryInbox struct {
	_peer      []byte
	_max_id    []byte
	_pts       []byte
	_pts_count []byte
}

func New_TL_updateReadHistoryInbox() *TL_updateReadHistoryInbox {
	return &TL_updateReadHistoryInbox{}
}

func (t *TL_updateReadHistoryInbox) Encode() []byte {
	return nil
}

func (t *TL_updateReadHistoryInbox) Decode(b []byte) {}

//====updateReadHistoryOutbox#2f2f21bf====

type TL_updateReadHistoryOutbox struct {
	_peer      []byte
	_max_id    []byte
	_pts       []byte
	_pts_count []byte
}

func New_TL_updateReadHistoryOutbox() *TL_updateReadHistoryOutbox {
	return &TL_updateReadHistoryOutbox{}
}

func (t *TL_updateReadHistoryOutbox) Encode() []byte {
	return nil
}

func (t *TL_updateReadHistoryOutbox) Decode(b []byte) {}

//====updateWebPage#7f891213====

type TL_updateWebPage struct {
	_webpage   []byte
	_pts       []byte
	_pts_count []byte
}

func New_TL_updateWebPage() *TL_updateWebPage {
	return &TL_updateWebPage{}
}

func (t *TL_updateWebPage) Encode() []byte {
	return nil
}

func (t *TL_updateWebPage) Decode(b []byte) {}

//====updateReadMessagesContents#68c13933====

type TL_updateReadMessagesContents struct {
	_messages  []byte
	_pts       []byte
	_pts_count []byte
}

func New_TL_updateReadMessagesContents() *TL_updateReadMessagesContents {
	return &TL_updateReadMessagesContents{}
}

func (t *TL_updateReadMessagesContents) Encode() []byte {
	return nil
}

func (t *TL_updateReadMessagesContents) Decode(b []byte) {}

//====updateChannelTooLong#eb0467fb====

type TL_updateChannelTooLong struct {
	_flags      []byte
	_channel_id []byte
	_pts        []byte
}

func New_TL_updateChannelTooLong() *TL_updateChannelTooLong {
	return &TL_updateChannelTooLong{}
}

func (t *TL_updateChannelTooLong) Encode() []byte {
	return nil
}

func (t *TL_updateChannelTooLong) Decode(b []byte) {}

//====updateChannel#b6d45656====

type TL_updateChannel struct {
	_channel_id []byte
}

func New_TL_updateChannel() *TL_updateChannel {
	return &TL_updateChannel{}
}

func (t *TL_updateChannel) Encode() []byte {
	return nil
}

func (t *TL_updateChannel) Decode(b []byte) {}

//====updateNewChannelMessage#62ba04d9====

type TL_updateNewChannelMessage struct {
	_message   []byte
	_pts       []byte
	_pts_count []byte
}

func New_TL_updateNewChannelMessage() *TL_updateNewChannelMessage {
	return &TL_updateNewChannelMessage{}
}

func (t *TL_updateNewChannelMessage) Encode() []byte {
	return nil
}

func (t *TL_updateNewChannelMessage) Decode(b []byte) {}

//====updateReadChannelInbox#4214f37f====

type TL_updateReadChannelInbox struct {
	_channel_id []byte
	_max_id     []byte
}

func New_TL_updateReadChannelInbox() *TL_updateReadChannelInbox {
	return &TL_updateReadChannelInbox{}
}

func (t *TL_updateReadChannelInbox) Encode() []byte {
	return nil
}

func (t *TL_updateReadChannelInbox) Decode(b []byte) {}

//====updateDeleteChannelMessages#c37521c9====

type TL_updateDeleteChannelMessages struct {
	_channel_id []byte
	_messages   []byte
	_pts        []byte
	_pts_count  []byte
}

func New_TL_updateDeleteChannelMessages() *TL_updateDeleteChannelMessages {
	return &TL_updateDeleteChannelMessages{}
}

func (t *TL_updateDeleteChannelMessages) Encode() []byte {
	return nil
}

func (t *TL_updateDeleteChannelMessages) Decode(b []byte) {}

//====updateChannelMessageViews#98a12b4b====

type TL_updateChannelMessageViews struct {
	_channel_id []byte
	_id         []byte
	_views      []byte
}

func New_TL_updateChannelMessageViews() *TL_updateChannelMessageViews {
	return &TL_updateChannelMessageViews{}
}

func (t *TL_updateChannelMessageViews) Encode() []byte {
	return nil
}

func (t *TL_updateChannelMessageViews) Decode(b []byte) {}

//====updateChatAdmins#6e947941====

type TL_updateChatAdmins struct {
	_chat_id []byte
	_enabled []byte
	_version []byte
}

func New_TL_updateChatAdmins() *TL_updateChatAdmins {
	return &TL_updateChatAdmins{}
}

func (t *TL_updateChatAdmins) Encode() []byte {
	return nil
}

func (t *TL_updateChatAdmins) Decode(b []byte) {}

//====updateChatParticipantAdmin#b6901959====

type TL_updateChatParticipantAdmin struct {
	_chat_id  []byte
	_user_id  []byte
	_is_admin []byte
	_version  []byte
}

func New_TL_updateChatParticipantAdmin() *TL_updateChatParticipantAdmin {
	return &TL_updateChatParticipantAdmin{}
}

func (t *TL_updateChatParticipantAdmin) Encode() []byte {
	return nil
}

func (t *TL_updateChatParticipantAdmin) Decode(b []byte) {}

//====updateNewStickerSet#688a30aa====

type TL_updateNewStickerSet struct {
	_stickerset []byte
}

func New_TL_updateNewStickerSet() *TL_updateNewStickerSet {
	return &TL_updateNewStickerSet{}
}

func (t *TL_updateNewStickerSet) Encode() []byte {
	return nil
}

func (t *TL_updateNewStickerSet) Decode(b []byte) {}

//====updateStickerSetsOrder#0bb2d201====

type TL_updateStickerSetsOrder struct {
	_flags []byte
	_masks []byte
	_order []byte
}

func New_TL_updateStickerSetsOrder() *TL_updateStickerSetsOrder {
	return &TL_updateStickerSetsOrder{}
}

func (t *TL_updateStickerSetsOrder) Encode() []byte {
	return nil
}

func (t *TL_updateStickerSetsOrder) Decode(b []byte) {}

//====updateStickerSets#43ae3dec====

type TL_updateStickerSets struct {
}

func New_TL_updateStickerSets() *TL_updateStickerSets {
	return &TL_updateStickerSets{}
}

func (t *TL_updateStickerSets) Encode() []byte {
	return nil
}

func (t *TL_updateStickerSets) Decode(b []byte) {}

//====updateSavedGifs#9375341e====

type TL_updateSavedGifs struct {
}

func New_TL_updateSavedGifs() *TL_updateSavedGifs {
	return &TL_updateSavedGifs{}
}

func (t *TL_updateSavedGifs) Encode() []byte {
	return nil
}

func (t *TL_updateSavedGifs) Decode(b []byte) {}

//====updateBotInlineQuery#54826690====

type TL_updateBotInlineQuery struct {
	_flags    []byte
	_query_id int64
	_user_id  []byte
	_query    string
	_geo      []byte
	_offset   string
}

func New_TL_updateBotInlineQuery() *TL_updateBotInlineQuery {
	return &TL_updateBotInlineQuery{}
}

func (t *TL_updateBotInlineQuery) Encode() []byte {
	return nil
}

func (t *TL_updateBotInlineQuery) Decode(b []byte) {}

//====updateBotInlineSend#0e48f964====

type TL_updateBotInlineSend struct {
	_flags   []byte
	_user_id []byte
	_query   string
	_geo     []byte
	_id      string
	_msg_id  []byte
}

func New_TL_updateBotInlineSend() *TL_updateBotInlineSend {
	return &TL_updateBotInlineSend{}
}

func (t *TL_updateBotInlineSend) Encode() []byte {
	return nil
}

func (t *TL_updateBotInlineSend) Decode(b []byte) {}

//====updateEditChannelMessage#1b3f4df7====

type TL_updateEditChannelMessage struct {
	_message   []byte
	_pts       []byte
	_pts_count []byte
}

func New_TL_updateEditChannelMessage() *TL_updateEditChannelMessage {
	return &TL_updateEditChannelMessage{}
}

func (t *TL_updateEditChannelMessage) Encode() []byte {
	return nil
}

func (t *TL_updateEditChannelMessage) Decode(b []byte) {}

//====updateChannelPinnedMessage#98592475====

type TL_updateChannelPinnedMessage struct {
	_channel_id []byte
	_id         []byte
}

func New_TL_updateChannelPinnedMessage() *TL_updateChannelPinnedMessage {
	return &TL_updateChannelPinnedMessage{}
}

func (t *TL_updateChannelPinnedMessage) Encode() []byte {
	return nil
}

func (t *TL_updateChannelPinnedMessage) Decode(b []byte) {}

//====updateBotCallbackQuery#e73547e1====

type TL_updateBotCallbackQuery struct {
	_flags           []byte
	_query_id        int64
	_user_id         []byte
	_peer            []byte
	_msg_id          []byte
	_chat_instance   int64
	_data            []byte
	_game_short_name []byte
}

func New_TL_updateBotCallbackQuery() *TL_updateBotCallbackQuery {
	return &TL_updateBotCallbackQuery{}
}

func (t *TL_updateBotCallbackQuery) Encode() []byte {
	return nil
}

func (t *TL_updateBotCallbackQuery) Decode(b []byte) {}

//====updateEditMessage#e40370a3====

type TL_updateEditMessage struct {
	_message   []byte
	_pts       []byte
	_pts_count []byte
}

func New_TL_updateEditMessage() *TL_updateEditMessage {
	return &TL_updateEditMessage{}
}

func (t *TL_updateEditMessage) Encode() []byte {
	return nil
}

func (t *TL_updateEditMessage) Decode(b []byte) {}

//====updateInlineBotCallbackQuery#f9d27a5a====

type TL_updateInlineBotCallbackQuery struct {
	_flags           []byte
	_query_id        int64
	_user_id         []byte
	_msg_id          []byte
	_chat_instance   int64
	_data            []byte
	_game_short_name []byte
}

func New_TL_updateInlineBotCallbackQuery() *TL_updateInlineBotCallbackQuery {
	return &TL_updateInlineBotCallbackQuery{}
}

func (t *TL_updateInlineBotCallbackQuery) Encode() []byte {
	return nil
}

func (t *TL_updateInlineBotCallbackQuery) Decode(b []byte) {}

//====updateReadChannelOutbox#25d6c9c7====

type TL_updateReadChannelOutbox struct {
	_channel_id []byte
	_max_id     []byte
}

func New_TL_updateReadChannelOutbox() *TL_updateReadChannelOutbox {
	return &TL_updateReadChannelOutbox{}
}

func (t *TL_updateReadChannelOutbox) Encode() []byte {
	return nil
}

func (t *TL_updateReadChannelOutbox) Decode(b []byte) {}

//====updateDraftMessage#ee2bb969====

type TL_updateDraftMessage struct {
	_peer  []byte
	_draft []byte
}

func New_TL_updateDraftMessage() *TL_updateDraftMessage {
	return &TL_updateDraftMessage{}
}

func (t *TL_updateDraftMessage) Encode() []byte {
	return nil
}

func (t *TL_updateDraftMessage) Decode(b []byte) {}

//====updateReadFeaturedStickers#571d2742====

type TL_updateReadFeaturedStickers struct {
}

func New_TL_updateReadFeaturedStickers() *TL_updateReadFeaturedStickers {
	return &TL_updateReadFeaturedStickers{}
}

func (t *TL_updateReadFeaturedStickers) Encode() []byte {
	return nil
}

func (t *TL_updateReadFeaturedStickers) Decode(b []byte) {}

//====updateRecentStickers#9a422c20====

type TL_updateRecentStickers struct {
}

func New_TL_updateRecentStickers() *TL_updateRecentStickers {
	return &TL_updateRecentStickers{}
}

func (t *TL_updateRecentStickers) Encode() []byte {
	return nil
}

func (t *TL_updateRecentStickers) Decode(b []byte) {}

//====updateConfig#a229dd06====

type TL_updateConfig struct {
}

func New_TL_updateConfig() *TL_updateConfig {
	return &TL_updateConfig{}
}

func (t *TL_updateConfig) Encode() []byte {
	return nil
}

func (t *TL_updateConfig) Decode(b []byte) {}

//====updatePtsChanged#3354678f====

type TL_updatePtsChanged struct {
}

func New_TL_updatePtsChanged() *TL_updatePtsChanged {
	return &TL_updatePtsChanged{}
}

func (t *TL_updatePtsChanged) Encode() []byte {
	return nil
}

func (t *TL_updatePtsChanged) Decode(b []byte) {}

//====updateChannelWebPage#40771900====

type TL_updateChannelWebPage struct {
	_channel_id []byte
	_webpage    []byte
	_pts        []byte
	_pts_count  []byte
}

func New_TL_updateChannelWebPage() *TL_updateChannelWebPage {
	return &TL_updateChannelWebPage{}
}

func (t *TL_updateChannelWebPage) Encode() []byte {
	return nil
}

func (t *TL_updateChannelWebPage) Decode(b []byte) {}

//====updateDialogPinned#d711a2cc====

type TL_updateDialogPinned struct {
	_flags  []byte
	_pinned []byte
	_peer   []byte
}

func New_TL_updateDialogPinned() *TL_updateDialogPinned {
	return &TL_updateDialogPinned{}
}

func (t *TL_updateDialogPinned) Encode() []byte {
	return nil
}

func (t *TL_updateDialogPinned) Decode(b []byte) {}

//====updatePinnedDialogs#d8caf68d====

type TL_updatePinnedDialogs struct {
	_flags []byte
	_order []byte
}

func New_TL_updatePinnedDialogs() *TL_updatePinnedDialogs {
	return &TL_updatePinnedDialogs{}
}

func (t *TL_updatePinnedDialogs) Encode() []byte {
	return nil
}

func (t *TL_updatePinnedDialogs) Decode(b []byte) {}

//====updateBotWebhookJSON#8317c0c3====

type TL_updateBotWebhookJSON struct {
	_data []byte
}

func New_TL_updateBotWebhookJSON() *TL_updateBotWebhookJSON {
	return &TL_updateBotWebhookJSON{}
}

func (t *TL_updateBotWebhookJSON) Encode() []byte {
	return nil
}

func (t *TL_updateBotWebhookJSON) Decode(b []byte) {}

//====updateBotWebhookJSONQuery#9b9240a6====

type TL_updateBotWebhookJSONQuery struct {
	_query_id int64
	_data     []byte
	_timeout  []byte
}

func New_TL_updateBotWebhookJSONQuery() *TL_updateBotWebhookJSONQuery {
	return &TL_updateBotWebhookJSONQuery{}
}

func (t *TL_updateBotWebhookJSONQuery) Encode() []byte {
	return nil
}

func (t *TL_updateBotWebhookJSONQuery) Decode(b []byte) {}

//====updateBotShippingQuery#e0cdc940====

type TL_updateBotShippingQuery struct {
	_query_id         int64
	_user_id          []byte
	_payload          []byte
	_shipping_address []byte
}

func New_TL_updateBotShippingQuery() *TL_updateBotShippingQuery {
	return &TL_updateBotShippingQuery{}
}

func (t *TL_updateBotShippingQuery) Encode() []byte {
	return nil
}

func (t *TL_updateBotShippingQuery) Decode(b []byte) {}

//====updateBotPrecheckoutQuery#5d2f3aa9====

type TL_updateBotPrecheckoutQuery struct {
	_flags              []byte
	_query_id           int64
	_user_id            []byte
	_payload            []byte
	_info               []byte
	_shipping_option_id []byte
	_currency           string
	_total_amount       int64
}

func New_TL_updateBotPrecheckoutQuery() *TL_updateBotPrecheckoutQuery {
	return &TL_updateBotPrecheckoutQuery{}
}

func (t *TL_updateBotPrecheckoutQuery) Encode() []byte {
	return nil
}

func (t *TL_updateBotPrecheckoutQuery) Decode(b []byte) {}

//====updatePhoneCall#ab0f6b1e====

type TL_updatePhoneCall struct {
	_phone_call []byte
}

func New_TL_updatePhoneCall() *TL_updatePhoneCall {
	return &TL_updatePhoneCall{}
}

func (t *TL_updatePhoneCall) Encode() []byte {
	return nil
}

func (t *TL_updatePhoneCall) Decode(b []byte) {}

//====updateLangPackTooLong#10c2404b====

type TL_updateLangPackTooLong struct {
}

func New_TL_updateLangPackTooLong() *TL_updateLangPackTooLong {
	return &TL_updateLangPackTooLong{}
}

func (t *TL_updateLangPackTooLong) Encode() []byte {
	return nil
}

func (t *TL_updateLangPackTooLong) Decode(b []byte) {}

//====updateLangPack#56022f4d====

type TL_updateLangPack struct {
	_difference []byte
}

func New_TL_updateLangPack() *TL_updateLangPack {
	return &TL_updateLangPack{}
}

func (t *TL_updateLangPack) Encode() []byte {
	return nil
}

func (t *TL_updateLangPack) Decode(b []byte) {}

//====updateFavedStickers#e511996d====

type TL_updateFavedStickers struct {
}

func New_TL_updateFavedStickers() *TL_updateFavedStickers {
	return &TL_updateFavedStickers{}
}

func (t *TL_updateFavedStickers) Encode() []byte {
	return nil
}

func (t *TL_updateFavedStickers) Decode(b []byte) {}

//====updateChannelReadMessagesContents#89893b45====

type TL_updateChannelReadMessagesContents struct {
	_channel_id []byte
	_messages   []byte
}

func New_TL_updateChannelReadMessagesContents() *TL_updateChannelReadMessagesContents {
	return &TL_updateChannelReadMessagesContents{}
}

func (t *TL_updateChannelReadMessagesContents) Encode() []byte {
	return nil
}

func (t *TL_updateChannelReadMessagesContents) Decode(b []byte) {}

//====updateContactsReset#7084a7be====

type TL_updateContactsReset struct {
}

func New_TL_updateContactsReset() *TL_updateContactsReset {
	return &TL_updateContactsReset{}
}

func (t *TL_updateContactsReset) Encode() []byte {
	return nil
}

func (t *TL_updateContactsReset) Decode(b []byte) {}

//====updateChannelAvailableMessages#70db6837====

type TL_updateChannelAvailableMessages struct {
	_channel_id       []byte
	_available_min_id []byte
}

func New_TL_updateChannelAvailableMessages() *TL_updateChannelAvailableMessages {
	return &TL_updateChannelAvailableMessages{}
}

func (t *TL_updateChannelAvailableMessages) Encode() []byte {
	return nil
}

func (t *TL_updateChannelAvailableMessages) Decode(b []byte) {}

//====updates_state#a56c2a3e====

type TL_updates_state struct {
	_pts          []byte
	_qts          []byte
	_date         []byte
	_seq          []byte
	_unread_count []byte
}

func New_TL_updates_state() *TL_updates_state {
	return &TL_updates_state{}
}

func (t *TL_updates_state) Encode() []byte {
	return nil
}

func (t *TL_updates_state) Decode(b []byte) {}

//====updates_differenceEmpty#5d75a138====

type TL_updates_differenceEmpty struct {
	_date []byte
	_seq  []byte
}

func New_TL_updates_differenceEmpty() *TL_updates_differenceEmpty {
	return &TL_updates_differenceEmpty{}
}

func (t *TL_updates_differenceEmpty) Encode() []byte {
	return nil
}

func (t *TL_updates_differenceEmpty) Decode(b []byte) {}

//====updates_difference#00f49ca0====

type TL_updates_difference struct {
	_new_messages           []byte
	_new_encrypted_messages []byte
	_other_updates          []byte
	_chats                  []byte
	_users                  []byte
	_state                  []byte
}

func New_TL_updates_difference() *TL_updates_difference {
	return &TL_updates_difference{}
}

func (t *TL_updates_difference) Encode() []byte {
	return nil
}

func (t *TL_updates_difference) Decode(b []byte) {}

//====updates_differenceSlice#a8fb1981====

type TL_updates_differenceSlice struct {
	_new_messages           []byte
	_new_encrypted_messages []byte
	_other_updates          []byte
	_chats                  []byte
	_users                  []byte
	_intermediate_state     []byte
}

func New_TL_updates_differenceSlice() *TL_updates_differenceSlice {
	return &TL_updates_differenceSlice{}
}

func (t *TL_updates_differenceSlice) Encode() []byte {
	return nil
}

func (t *TL_updates_differenceSlice) Decode(b []byte) {}

//====updates_differenceTooLong#4afe8f6d====

type TL_updates_differenceTooLong struct {
	_pts []byte
}

func New_TL_updates_differenceTooLong() *TL_updates_differenceTooLong {
	return &TL_updates_differenceTooLong{}
}

func (t *TL_updates_differenceTooLong) Encode() []byte {
	return nil
}

func (t *TL_updates_differenceTooLong) Decode(b []byte) {}

//====updatesTooLong#e317af7e====

type TL_updatesTooLong struct {
}

func New_TL_updatesTooLong() *TL_updatesTooLong {
	return &TL_updatesTooLong{}
}

func (t *TL_updatesTooLong) Encode() []byte {
	return nil
}

func (t *TL_updatesTooLong) Decode(b []byte) {}

//====updateShortMessage#914fbf11====

type TL_updateShortMessage struct {
	_flags           []byte
	_out             []byte
	_mentioned       []byte
	_media_unread    []byte
	_silent          []byte
	_id              []byte
	_user_id         []byte
	_message         string
	_pts             []byte
	_pts_count       []byte
	_date            []byte
	_fwd_from        []byte
	_via_bot_id      []byte
	_reply_to_msg_id []byte
	_entities        []byte
}

func New_TL_updateShortMessage() *TL_updateShortMessage {
	return &TL_updateShortMessage{}
}

func (t *TL_updateShortMessage) Encode() []byte {
	return nil
}

func (t *TL_updateShortMessage) Decode(b []byte) {}

//====updateShortChatMessage#16812688====

type TL_updateShortChatMessage struct {
	_flags           []byte
	_out             []byte
	_mentioned       []byte
	_media_unread    []byte
	_silent          []byte
	_id              []byte
	_from_id         []byte
	_chat_id         []byte
	_message         string
	_pts             []byte
	_pts_count       []byte
	_date            []byte
	_fwd_from        []byte
	_via_bot_id      []byte
	_reply_to_msg_id []byte
	_entities        []byte
}

func New_TL_updateShortChatMessage() *TL_updateShortChatMessage {
	return &TL_updateShortChatMessage{}
}

func (t *TL_updateShortChatMessage) Encode() []byte {
	return nil
}

func (t *TL_updateShortChatMessage) Decode(b []byte) {}

//====updateShort#78d4dec1====

type TL_updateShort struct {
	_update []byte
	_date   []byte
}

func New_TL_updateShort() *TL_updateShort {
	return &TL_updateShort{}
}

func (t *TL_updateShort) Encode() []byte {
	return nil
}

func (t *TL_updateShort) Decode(b []byte) {}

//====updatesCombined#725b04c3====

type TL_updatesCombined struct {
	_updates   []byte
	_users     []byte
	_chats     []byte
	_date      []byte
	_seq_start []byte
	_seq       []byte
}

func New_TL_updatesCombined() *TL_updatesCombined {
	return &TL_updatesCombined{}
}

func (t *TL_updatesCombined) Encode() []byte {
	return nil
}

func (t *TL_updatesCombined) Decode(b []byte) {}

//====updates#74ae4240====

type TL_updates struct {
	_updates []byte
	_users   []byte
	_chats   []byte
	_date    []byte
	_seq     []byte
}

func New_TL_updates() *TL_updates {
	return &TL_updates{}
}

func (t *TL_updates) Encode() []byte {
	return nil
}

func (t *TL_updates) Decode(b []byte) {}

//====updateShortSentMessage#11f1331c====

type TL_updateShortSentMessage struct {
	_flags     []byte
	_out       []byte
	_id        []byte
	_pts       []byte
	_pts_count []byte
	_date      []byte
	_media     []byte
	_entities  []byte
}

func New_TL_updateShortSentMessage() *TL_updateShortSentMessage {
	return &TL_updateShortSentMessage{}
}

func (t *TL_updateShortSentMessage) Encode() []byte {
	return nil
}

func (t *TL_updateShortSentMessage) Decode(b []byte) {}

//====photos_photos#8dca6aa5====

type TL_photos_photos struct {
	_photos []byte
	_users  []byte
}

func New_TL_photos_photos() *TL_photos_photos {
	return &TL_photos_photos{}
}

func (t *TL_photos_photos) Encode() []byte {
	return nil
}

func (t *TL_photos_photos) Decode(b []byte) {}

//====photos_photosSlice#15051f54====

type TL_photos_photosSlice struct {
	_count  []byte
	_photos []byte
	_users  []byte
}

func New_TL_photos_photosSlice() *TL_photos_photosSlice {
	return &TL_photos_photosSlice{}
}

func (t *TL_photos_photosSlice) Encode() []byte {
	return nil
}

func (t *TL_photos_photosSlice) Decode(b []byte) {}

//====photos_photo#20212ca8====

type TL_photos_photo struct {
	_photo []byte
	_users []byte
}

func New_TL_photos_photo() *TL_photos_photo {
	return &TL_photos_photo{}
}

func (t *TL_photos_photo) Encode() []byte {
	return nil
}

func (t *TL_photos_photo) Decode(b []byte) {}

//====upload_file#096a18d5====

type TL_upload_file struct {
	_type  []byte
	_mtime []byte
	_bytes []byte
}

func New_TL_upload_file() *TL_upload_file {
	return &TL_upload_file{}
}

func (t *TL_upload_file) Encode() []byte {
	return nil
}

func (t *TL_upload_file) Decode(b []byte) {}

//====upload_fileCdnRedirect#ea52fe5a====

type TL_upload_fileCdnRedirect struct {
	_dc_id           []byte
	_file_token      []byte
	_encryption_key  []byte
	_encryption_iv   []byte
	_cdn_file_hashes []byte
}

func New_TL_upload_fileCdnRedirect() *TL_upload_fileCdnRedirect {
	return &TL_upload_fileCdnRedirect{}
}

func (t *TL_upload_fileCdnRedirect) Encode() []byte {
	return nil
}

func (t *TL_upload_fileCdnRedirect) Decode(b []byte) {}

//====dcOption#05d8c6cc====

type TL_dcOption struct {
	_flags      []byte
	_ipv6       []byte
	_media_only []byte
	_tcpo_only  []byte
	_cdn        []byte
	_static     []byte
	_id         []byte
	_ip_address string
	_port       []byte
}

func New_TL_dcOption() *TL_dcOption {
	return &TL_dcOption{}
}

func (t *TL_dcOption) Encode() []byte {
	return nil
}

func (t *TL_dcOption) Decode(b []byte) {}

//====config#9c840964====

type TL_config struct {
	_flags                      []byte
	_phonecalls_enabled         []byte
	_default_p2p_contacts       []byte
	_date                       []byte
	_expires                    []byte
	_test_mode                  []byte
	_this_dc                    []byte
	_dc_options                 []byte
	_chat_size_max              []byte
	_megagroup_size_max         []byte
	_forwarded_count_max        []byte
	_online_update_period_ms    []byte
	_offline_blur_timeout_ms    []byte
	_offline_idle_timeout_ms    []byte
	_online_cloud_timeout_ms    []byte
	_notify_cloud_delay_ms      []byte
	_notify_default_delay_ms    []byte
	_chat_big_size              []byte
	_push_chat_period_ms        []byte
	_push_chat_limit            []byte
	_saved_gifs_limit           []byte
	_edit_time_limit            []byte
	_rating_e_decay             []byte
	_stickers_recent_limit      []byte
	_stickers_faved_limit       []byte
	_channels_read_media_period []byte
	_tmp_sessions               []byte
	_pinned_dialogs_count_max   []byte
	_call_receive_timeout_ms    []byte
	_call_ring_timeout_ms       []byte
	_call_connect_timeout_ms    []byte
	_call_packet_timeout_ms     []byte
	_me_url_prefix              string
	_suggested_lang_code        []byte
	_lang_pack_version          []byte
	_disabled_features          []byte
}

func New_TL_config() *TL_config {
	return &TL_config{}
}

func (t *TL_config) Encode() []byte {
	return nil
}

func (t *TL_config) Decode(b []byte) {}

//====nearestDc#8e1a1775====

type TL_nearestDc struct {
	_country    string
	_this_dc    []byte
	_nearest_dc []byte
}

func New_TL_nearestDc() *TL_nearestDc {
	return &TL_nearestDc{}
}

func (t *TL_nearestDc) Encode() []byte {
	return nil
}

func (t *TL_nearestDc) Decode(b []byte) {}

//====help_appUpdate#8987f311====

type TL_help_appUpdate struct {
	_id       []byte
	_critical []byte
	_url      string
	_text     string
}

func New_TL_help_appUpdate() *TL_help_appUpdate {
	return &TL_help_appUpdate{}
}

func (t *TL_help_appUpdate) Encode() []byte {
	return nil
}

func (t *TL_help_appUpdate) Decode(b []byte) {}

//====help_noAppUpdate#c45a6536====

type TL_help_noAppUpdate struct {
}

func New_TL_help_noAppUpdate() *TL_help_noAppUpdate {
	return &TL_help_noAppUpdate{}
}

func (t *TL_help_noAppUpdate) Encode() []byte {
	return nil
}

func (t *TL_help_noAppUpdate) Decode(b []byte) {}

//====help_inviteText#18cb9f78====

type TL_help_inviteText struct {
	_message string
}

func New_TL_help_inviteText() *TL_help_inviteText {
	return &TL_help_inviteText{}
}

func (t *TL_help_inviteText) Encode() []byte {
	return nil
}

func (t *TL_help_inviteText) Decode(b []byte) {}

//====encryptedChatEmpty#ab7ec0a0====

type TL_encryptedChatEmpty struct {
	_id []byte
}

func New_TL_encryptedChatEmpty() *TL_encryptedChatEmpty {
	return &TL_encryptedChatEmpty{}
}

func (t *TL_encryptedChatEmpty) Encode() []byte {
	return nil
}

func (t *TL_encryptedChatEmpty) Decode(b []byte) {}

//====encryptedChatWaiting#3bf703dc====

type TL_encryptedChatWaiting struct {
	_id             []byte
	_access_hash    int64
	_date           []byte
	_admin_id       []byte
	_participant_id []byte
}

func New_TL_encryptedChatWaiting() *TL_encryptedChatWaiting {
	return &TL_encryptedChatWaiting{}
}

func (t *TL_encryptedChatWaiting) Encode() []byte {
	return nil
}

func (t *TL_encryptedChatWaiting) Decode(b []byte) {}

//====encryptedChatRequested#c878527e====

type TL_encryptedChatRequested struct {
	_id             []byte
	_access_hash    int64
	_date           []byte
	_admin_id       []byte
	_participant_id []byte
	_g_a            []byte
}

func New_TL_encryptedChatRequested() *TL_encryptedChatRequested {
	return &TL_encryptedChatRequested{}
}

func (t *TL_encryptedChatRequested) Encode() []byte {
	return nil
}

func (t *TL_encryptedChatRequested) Decode(b []byte) {}

//====encryptedChat#fa56ce36====

type TL_encryptedChat struct {
	_id              []byte
	_access_hash     int64
	_date            []byte
	_admin_id        []byte
	_participant_id  []byte
	_g_a_or_b        []byte
	_key_fingerprint int64
}

func New_TL_encryptedChat() *TL_encryptedChat {
	return &TL_encryptedChat{}
}

func (t *TL_encryptedChat) Encode() []byte {
	return nil
}

func (t *TL_encryptedChat) Decode(b []byte) {}

//====encryptedChatDiscarded#13d6dd27====

type TL_encryptedChatDiscarded struct {
	_id []byte
}

func New_TL_encryptedChatDiscarded() *TL_encryptedChatDiscarded {
	return &TL_encryptedChatDiscarded{}
}

func (t *TL_encryptedChatDiscarded) Encode() []byte {
	return nil
}

func (t *TL_encryptedChatDiscarded) Decode(b []byte) {}

//====inputEncryptedChat#f141b5e1====

type TL_inputEncryptedChat struct {
	_chat_id     []byte
	_access_hash int64
}

func New_TL_inputEncryptedChat() *TL_inputEncryptedChat {
	return &TL_inputEncryptedChat{}
}

func (t *TL_inputEncryptedChat) Encode() []byte {
	return nil
}

func (t *TL_inputEncryptedChat) Decode(b []byte) {}

//====encryptedFileEmpty#c21f497e====

type TL_encryptedFileEmpty struct {
}

func New_TL_encryptedFileEmpty() *TL_encryptedFileEmpty {
	return &TL_encryptedFileEmpty{}
}

func (t *TL_encryptedFileEmpty) Encode() []byte {
	return nil
}

func (t *TL_encryptedFileEmpty) Decode(b []byte) {}

//====encryptedFile#4a70994c====

type TL_encryptedFile struct {
	_id              int64
	_access_hash     int64
	_size            []byte
	_dc_id           []byte
	_key_fingerprint []byte
}

func New_TL_encryptedFile() *TL_encryptedFile {
	return &TL_encryptedFile{}
}

func (t *TL_encryptedFile) Encode() []byte {
	return nil
}

func (t *TL_encryptedFile) Decode(b []byte) {}

//====inputEncryptedFileEmpty#1837c364====

type TL_inputEncryptedFileEmpty struct {
}

func New_TL_inputEncryptedFileEmpty() *TL_inputEncryptedFileEmpty {
	return &TL_inputEncryptedFileEmpty{}
}

func (t *TL_inputEncryptedFileEmpty) Encode() []byte {
	return nil
}

func (t *TL_inputEncryptedFileEmpty) Decode(b []byte) {}

//====inputEncryptedFileUploaded#64bd0306====

type TL_inputEncryptedFileUploaded struct {
	_id              int64
	_parts           []byte
	_md5_checksum    string
	_key_fingerprint []byte
}

func New_TL_inputEncryptedFileUploaded() *TL_inputEncryptedFileUploaded {
	return &TL_inputEncryptedFileUploaded{}
}

func (t *TL_inputEncryptedFileUploaded) Encode() []byte {
	return nil
}

func (t *TL_inputEncryptedFileUploaded) Decode(b []byte) {}

//====inputEncryptedFile#5a17b5e5====

type TL_inputEncryptedFile struct {
	_id          int64
	_access_hash int64
}

func New_TL_inputEncryptedFile() *TL_inputEncryptedFile {
	return &TL_inputEncryptedFile{}
}

func (t *TL_inputEncryptedFile) Encode() []byte {
	return nil
}

func (t *TL_inputEncryptedFile) Decode(b []byte) {}

//====inputEncryptedFileBigUploaded#2dc173c8====

type TL_inputEncryptedFileBigUploaded struct {
	_id              int64
	_parts           []byte
	_key_fingerprint []byte
}

func New_TL_inputEncryptedFileBigUploaded() *TL_inputEncryptedFileBigUploaded {
	return &TL_inputEncryptedFileBigUploaded{}
}

func (t *TL_inputEncryptedFileBigUploaded) Encode() []byte {
	return nil
}

func (t *TL_inputEncryptedFileBigUploaded) Decode(b []byte) {}

//====encryptedMessage#ed18c118====

type TL_encryptedMessage struct {
	_random_id int64
	_chat_id   []byte
	_date      []byte
	_bytes     []byte
	_file      []byte
}

func New_TL_encryptedMessage() *TL_encryptedMessage {
	return &TL_encryptedMessage{}
}

func (t *TL_encryptedMessage) Encode() []byte {
	return nil
}

func (t *TL_encryptedMessage) Decode(b []byte) {}

//====encryptedMessageService#23734b06====

type TL_encryptedMessageService struct {
	_random_id int64
	_chat_id   []byte
	_date      []byte
	_bytes     []byte
}

func New_TL_encryptedMessageService() *TL_encryptedMessageService {
	return &TL_encryptedMessageService{}
}

func (t *TL_encryptedMessageService) Encode() []byte {
	return nil
}

func (t *TL_encryptedMessageService) Decode(b []byte) {}

//====messages_dhConfigNotModified#c0e24635====

type TL_messages_dhConfigNotModified struct {
	_random []byte
}

func New_TL_messages_dhConfigNotModified() *TL_messages_dhConfigNotModified {
	return &TL_messages_dhConfigNotModified{}
}

func (t *TL_messages_dhConfigNotModified) Encode() []byte {
	return nil
}

func (t *TL_messages_dhConfigNotModified) Decode(b []byte) {}

//====messages_dhConfig#2c221edd====

type TL_messages_dhConfig struct {
	_g       []byte
	_p       []byte
	_version []byte
	_random  []byte
}

func New_TL_messages_dhConfig() *TL_messages_dhConfig {
	return &TL_messages_dhConfig{}
}

func (t *TL_messages_dhConfig) Encode() []byte {
	return nil
}

func (t *TL_messages_dhConfig) Decode(b []byte) {}

//====messages_sentEncryptedMessage#560f8935====

type TL_messages_sentEncryptedMessage struct {
	_date []byte
}

func New_TL_messages_sentEncryptedMessage() *TL_messages_sentEncryptedMessage {
	return &TL_messages_sentEncryptedMessage{}
}

func (t *TL_messages_sentEncryptedMessage) Encode() []byte {
	return nil
}

func (t *TL_messages_sentEncryptedMessage) Decode(b []byte) {}

//====messages_sentEncryptedFile#9493ff32====

type TL_messages_sentEncryptedFile struct {
	_date []byte
	_file []byte
}

func New_TL_messages_sentEncryptedFile() *TL_messages_sentEncryptedFile {
	return &TL_messages_sentEncryptedFile{}
}

func (t *TL_messages_sentEncryptedFile) Encode() []byte {
	return nil
}

func (t *TL_messages_sentEncryptedFile) Decode(b []byte) {}

//====inputDocumentEmpty#72f0eaae====

type TL_inputDocumentEmpty struct {
}

func New_TL_inputDocumentEmpty() *TL_inputDocumentEmpty {
	return &TL_inputDocumentEmpty{}
}

func (t *TL_inputDocumentEmpty) Encode() []byte {
	return nil
}

func (t *TL_inputDocumentEmpty) Decode(b []byte) {}

//====inputDocument#18798952====

type TL_inputDocument struct {
	_id          int64
	_access_hash int64
}

func New_TL_inputDocument() *TL_inputDocument {
	return &TL_inputDocument{}
}

func (t *TL_inputDocument) Encode() []byte {
	return nil
}

func (t *TL_inputDocument) Decode(b []byte) {}

//====documentEmpty#36f8c871====

type TL_documentEmpty struct {
	_id int64
}

func New_TL_documentEmpty() *TL_documentEmpty {
	return &TL_documentEmpty{}
}

func (t *TL_documentEmpty) Encode() []byte {
	return nil
}

func (t *TL_documentEmpty) Decode(b []byte) {}

//====document#87232bc7====

type TL_document struct {
	_id          int64
	_access_hash int64
	_date        []byte
	_mime_type   string
	_size        []byte
	_thumb       []byte
	_dc_id       []byte
	_version     []byte
	_attributes  []byte
}

func New_TL_document() *TL_document {
	return &TL_document{}
}

func (t *TL_document) Encode() []byte {
	return nil
}

func (t *TL_document) Decode(b []byte) {}

//====help_support#17c6b5f6====

type TL_help_support struct {
	_phone_number string
	_user         []byte
}

func New_TL_help_support() *TL_help_support {
	return &TL_help_support{}
}

func (t *TL_help_support) Encode() []byte {
	return nil
}

func (t *TL_help_support) Decode(b []byte) {}

//====notifyPeer#9fd40bd8====

type TL_notifyPeer struct {
	_peer []byte
}

func New_TL_notifyPeer() *TL_notifyPeer {
	return &TL_notifyPeer{}
}

func (t *TL_notifyPeer) Encode() []byte {
	return nil
}

func (t *TL_notifyPeer) Decode(b []byte) {}

//====notifyUsers#b4c83b4c====

type TL_notifyUsers struct {
}

func New_TL_notifyUsers() *TL_notifyUsers {
	return &TL_notifyUsers{}
}

func (t *TL_notifyUsers) Encode() []byte {
	return nil
}

func (t *TL_notifyUsers) Decode(b []byte) {}

//====notifyChats#c007cec3====

type TL_notifyChats struct {
}

func New_TL_notifyChats() *TL_notifyChats {
	return &TL_notifyChats{}
}

func (t *TL_notifyChats) Encode() []byte {
	return nil
}

func (t *TL_notifyChats) Decode(b []byte) {}

//====notifyAll#74d07c60====

type TL_notifyAll struct {
}

func New_TL_notifyAll() *TL_notifyAll {
	return &TL_notifyAll{}
}

func (t *TL_notifyAll) Encode() []byte {
	return nil
}

func (t *TL_notifyAll) Decode(b []byte) {}

//====sendMessageTypingAction#16bf744e====

type TL_sendMessageTypingAction struct {
}

func New_TL_sendMessageTypingAction() *TL_sendMessageTypingAction {
	return &TL_sendMessageTypingAction{}
}

func (t *TL_sendMessageTypingAction) Encode() []byte {
	return nil
}

func (t *TL_sendMessageTypingAction) Decode(b []byte) {}

//====sendMessageCancelAction#fd5ec8f5====

type TL_sendMessageCancelAction struct {
}

func New_TL_sendMessageCancelAction() *TL_sendMessageCancelAction {
	return &TL_sendMessageCancelAction{}
}

func (t *TL_sendMessageCancelAction) Encode() []byte {
	return nil
}

func (t *TL_sendMessageCancelAction) Decode(b []byte) {}

//====sendMessageRecordVideoAction#a187d66f====

type TL_sendMessageRecordVideoAction struct {
}

func New_TL_sendMessageRecordVideoAction() *TL_sendMessageRecordVideoAction {
	return &TL_sendMessageRecordVideoAction{}
}

func (t *TL_sendMessageRecordVideoAction) Encode() []byte {
	return nil
}

func (t *TL_sendMessageRecordVideoAction) Decode(b []byte) {}

//====sendMessageUploadVideoAction#e9763aec====

type TL_sendMessageUploadVideoAction struct {
	_progress []byte
}

func New_TL_sendMessageUploadVideoAction() *TL_sendMessageUploadVideoAction {
	return &TL_sendMessageUploadVideoAction{}
}

func (t *TL_sendMessageUploadVideoAction) Encode() []byte {
	return nil
}

func (t *TL_sendMessageUploadVideoAction) Decode(b []byte) {}

//====sendMessageRecordAudioAction#d52f73f7====

type TL_sendMessageRecordAudioAction struct {
}

func New_TL_sendMessageRecordAudioAction() *TL_sendMessageRecordAudioAction {
	return &TL_sendMessageRecordAudioAction{}
}

func (t *TL_sendMessageRecordAudioAction) Encode() []byte {
	return nil
}

func (t *TL_sendMessageRecordAudioAction) Decode(b []byte) {}

//====sendMessageUploadAudioAction#f351d7ab====

type TL_sendMessageUploadAudioAction struct {
	_progress []byte
}

func New_TL_sendMessageUploadAudioAction() *TL_sendMessageUploadAudioAction {
	return &TL_sendMessageUploadAudioAction{}
}

func (t *TL_sendMessageUploadAudioAction) Encode() []byte {
	return nil
}

func (t *TL_sendMessageUploadAudioAction) Decode(b []byte) {}

//====sendMessageUploadPhotoAction#d1d34a26====

type TL_sendMessageUploadPhotoAction struct {
	_progress []byte
}

func New_TL_sendMessageUploadPhotoAction() *TL_sendMessageUploadPhotoAction {
	return &TL_sendMessageUploadPhotoAction{}
}

func (t *TL_sendMessageUploadPhotoAction) Encode() []byte {
	return nil
}

func (t *TL_sendMessageUploadPhotoAction) Decode(b []byte) {}

//====sendMessageUploadDocumentAction#aa0cd9e4====

type TL_sendMessageUploadDocumentAction struct {
	_progress []byte
}

func New_TL_sendMessageUploadDocumentAction() *TL_sendMessageUploadDocumentAction {
	return &TL_sendMessageUploadDocumentAction{}
}

func (t *TL_sendMessageUploadDocumentAction) Encode() []byte {
	return nil
}

func (t *TL_sendMessageUploadDocumentAction) Decode(b []byte) {}

//====sendMessageGeoLocationAction#176f8ba1====

type TL_sendMessageGeoLocationAction struct {
}

func New_TL_sendMessageGeoLocationAction() *TL_sendMessageGeoLocationAction {
	return &TL_sendMessageGeoLocationAction{}
}

func (t *TL_sendMessageGeoLocationAction) Encode() []byte {
	return nil
}

func (t *TL_sendMessageGeoLocationAction) Decode(b []byte) {}

//====sendMessageChooseContactAction#628cbc6f====

type TL_sendMessageChooseContactAction struct {
}

func New_TL_sendMessageChooseContactAction() *TL_sendMessageChooseContactAction {
	return &TL_sendMessageChooseContactAction{}
}

func (t *TL_sendMessageChooseContactAction) Encode() []byte {
	return nil
}

func (t *TL_sendMessageChooseContactAction) Decode(b []byte) {}

//====sendMessageGamePlayAction#dd6a8f48====

type TL_sendMessageGamePlayAction struct {
}

func New_TL_sendMessageGamePlayAction() *TL_sendMessageGamePlayAction {
	return &TL_sendMessageGamePlayAction{}
}

func (t *TL_sendMessageGamePlayAction) Encode() []byte {
	return nil
}

func (t *TL_sendMessageGamePlayAction) Decode(b []byte) {}

//====sendMessageRecordRoundAction#88f27fbc====

type TL_sendMessageRecordRoundAction struct {
}

func New_TL_sendMessageRecordRoundAction() *TL_sendMessageRecordRoundAction {
	return &TL_sendMessageRecordRoundAction{}
}

func (t *TL_sendMessageRecordRoundAction) Encode() []byte {
	return nil
}

func (t *TL_sendMessageRecordRoundAction) Decode(b []byte) {}

//====sendMessageUploadRoundAction#243e1c66====

type TL_sendMessageUploadRoundAction struct {
	_progress []byte
}

func New_TL_sendMessageUploadRoundAction() *TL_sendMessageUploadRoundAction {
	return &TL_sendMessageUploadRoundAction{}
}

func (t *TL_sendMessageUploadRoundAction) Encode() []byte {
	return nil
}

func (t *TL_sendMessageUploadRoundAction) Decode(b []byte) {}

//====contacts_found#1aa1f784====

type TL_contacts_found struct {
	_results []byte
	_chats   []byte
	_users   []byte
}

func New_TL_contacts_found() *TL_contacts_found {
	return &TL_contacts_found{}
}

func (t *TL_contacts_found) Encode() []byte {
	return nil
}

func (t *TL_contacts_found) Decode(b []byte) {}

//====inputPrivacyKeyStatusTimestamp#4f96cb18====

type TL_inputPrivacyKeyStatusTimestamp struct {
}

func New_TL_inputPrivacyKeyStatusTimestamp() *TL_inputPrivacyKeyStatusTimestamp {
	return &TL_inputPrivacyKeyStatusTimestamp{}
}

func (t *TL_inputPrivacyKeyStatusTimestamp) Encode() []byte {
	return nil
}

func (t *TL_inputPrivacyKeyStatusTimestamp) Decode(b []byte) {}

//====inputPrivacyKeyChatInvite#bdfb0426====

type TL_inputPrivacyKeyChatInvite struct {
}

func New_TL_inputPrivacyKeyChatInvite() *TL_inputPrivacyKeyChatInvite {
	return &TL_inputPrivacyKeyChatInvite{}
}

func (t *TL_inputPrivacyKeyChatInvite) Encode() []byte {
	return nil
}

func (t *TL_inputPrivacyKeyChatInvite) Decode(b []byte) {}

//====inputPrivacyKeyPhoneCall#fabadc5f====

type TL_inputPrivacyKeyPhoneCall struct {
}

func New_TL_inputPrivacyKeyPhoneCall() *TL_inputPrivacyKeyPhoneCall {
	return &TL_inputPrivacyKeyPhoneCall{}
}

func (t *TL_inputPrivacyKeyPhoneCall) Encode() []byte {
	return nil
}

func (t *TL_inputPrivacyKeyPhoneCall) Decode(b []byte) {}

//====privacyKeyStatusTimestamp#bc2eab30====

type TL_privacyKeyStatusTimestamp struct {
}

func New_TL_privacyKeyStatusTimestamp() *TL_privacyKeyStatusTimestamp {
	return &TL_privacyKeyStatusTimestamp{}
}

func (t *TL_privacyKeyStatusTimestamp) Encode() []byte {
	return nil
}

func (t *TL_privacyKeyStatusTimestamp) Decode(b []byte) {}

//====privacyKeyChatInvite#500e6dfa====

type TL_privacyKeyChatInvite struct {
}

func New_TL_privacyKeyChatInvite() *TL_privacyKeyChatInvite {
	return &TL_privacyKeyChatInvite{}
}

func (t *TL_privacyKeyChatInvite) Encode() []byte {
	return nil
}

func (t *TL_privacyKeyChatInvite) Decode(b []byte) {}

//====privacyKeyPhoneCall#3d662b7b====

type TL_privacyKeyPhoneCall struct {
}

func New_TL_privacyKeyPhoneCall() *TL_privacyKeyPhoneCall {
	return &TL_privacyKeyPhoneCall{}
}

func (t *TL_privacyKeyPhoneCall) Encode() []byte {
	return nil
}

func (t *TL_privacyKeyPhoneCall) Decode(b []byte) {}

//====inputPrivacyValueAllowContacts#0d09e07b====

type TL_inputPrivacyValueAllowContacts struct {
}

func New_TL_inputPrivacyValueAllowContacts() *TL_inputPrivacyValueAllowContacts {
	return &TL_inputPrivacyValueAllowContacts{}
}

func (t *TL_inputPrivacyValueAllowContacts) Encode() []byte {
	return nil
}

func (t *TL_inputPrivacyValueAllowContacts) Decode(b []byte) {}

//====inputPrivacyValueAllowAll#184b35ce====

type TL_inputPrivacyValueAllowAll struct {
}

func New_TL_inputPrivacyValueAllowAll() *TL_inputPrivacyValueAllowAll {
	return &TL_inputPrivacyValueAllowAll{}
}

func (t *TL_inputPrivacyValueAllowAll) Encode() []byte {
	return nil
}

func (t *TL_inputPrivacyValueAllowAll) Decode(b []byte) {}

//====inputPrivacyValueAllowUsers#131cc67f====

type TL_inputPrivacyValueAllowUsers struct {
	_users []byte
}

func New_TL_inputPrivacyValueAllowUsers() *TL_inputPrivacyValueAllowUsers {
	return &TL_inputPrivacyValueAllowUsers{}
}

func (t *TL_inputPrivacyValueAllowUsers) Encode() []byte {
	return nil
}

func (t *TL_inputPrivacyValueAllowUsers) Decode(b []byte) {}

//====inputPrivacyValueDisallowContacts#0ba52007====

type TL_inputPrivacyValueDisallowContacts struct {
}

func New_TL_inputPrivacyValueDisallowContacts() *TL_inputPrivacyValueDisallowContacts {
	return &TL_inputPrivacyValueDisallowContacts{}
}

func (t *TL_inputPrivacyValueDisallowContacts) Encode() []byte {
	return nil
}

func (t *TL_inputPrivacyValueDisallowContacts) Decode(b []byte) {}

//====inputPrivacyValueDisallowAll#d66b66c9====

type TL_inputPrivacyValueDisallowAll struct {
}

func New_TL_inputPrivacyValueDisallowAll() *TL_inputPrivacyValueDisallowAll {
	return &TL_inputPrivacyValueDisallowAll{}
}

func (t *TL_inputPrivacyValueDisallowAll) Encode() []byte {
	return nil
}

func (t *TL_inputPrivacyValueDisallowAll) Decode(b []byte) {}

//====inputPrivacyValueDisallowUsers#90110467====

type TL_inputPrivacyValueDisallowUsers struct {
	_users []byte
}

func New_TL_inputPrivacyValueDisallowUsers() *TL_inputPrivacyValueDisallowUsers {
	return &TL_inputPrivacyValueDisallowUsers{}
}

func (t *TL_inputPrivacyValueDisallowUsers) Encode() []byte {
	return nil
}

func (t *TL_inputPrivacyValueDisallowUsers) Decode(b []byte) {}

//====privacyValueAllowContacts#fffe1bac====

type TL_privacyValueAllowContacts struct {
}

func New_TL_privacyValueAllowContacts() *TL_privacyValueAllowContacts {
	return &TL_privacyValueAllowContacts{}
}

func (t *TL_privacyValueAllowContacts) Encode() []byte {
	return nil
}

func (t *TL_privacyValueAllowContacts) Decode(b []byte) {}

//====privacyValueAllowAll#65427b82====

type TL_privacyValueAllowAll struct {
}

func New_TL_privacyValueAllowAll() *TL_privacyValueAllowAll {
	return &TL_privacyValueAllowAll{}
}

func (t *TL_privacyValueAllowAll) Encode() []byte {
	return nil
}

func (t *TL_privacyValueAllowAll) Decode(b []byte) {}

//====privacyValueAllowUsers#4d5bbe0c====

type TL_privacyValueAllowUsers struct {
	_users []byte
}

func New_TL_privacyValueAllowUsers() *TL_privacyValueAllowUsers {
	return &TL_privacyValueAllowUsers{}
}

func (t *TL_privacyValueAllowUsers) Encode() []byte {
	return nil
}

func (t *TL_privacyValueAllowUsers) Decode(b []byte) {}

//====privacyValueDisallowContacts#f888fa1a====

type TL_privacyValueDisallowContacts struct {
}

func New_TL_privacyValueDisallowContacts() *TL_privacyValueDisallowContacts {
	return &TL_privacyValueDisallowContacts{}
}

func (t *TL_privacyValueDisallowContacts) Encode() []byte {
	return nil
}

func (t *TL_privacyValueDisallowContacts) Decode(b []byte) {}

//====privacyValueDisallowAll#8b73e763====

type TL_privacyValueDisallowAll struct {
}

func New_TL_privacyValueDisallowAll() *TL_privacyValueDisallowAll {
	return &TL_privacyValueDisallowAll{}
}

func (t *TL_privacyValueDisallowAll) Encode() []byte {
	return nil
}

func (t *TL_privacyValueDisallowAll) Decode(b []byte) {}

//====privacyValueDisallowUsers#0c7f49b7====

type TL_privacyValueDisallowUsers struct {
	_users []byte
}

func New_TL_privacyValueDisallowUsers() *TL_privacyValueDisallowUsers {
	return &TL_privacyValueDisallowUsers{}
}

func (t *TL_privacyValueDisallowUsers) Encode() []byte {
	return nil
}

func (t *TL_privacyValueDisallowUsers) Decode(b []byte) {}

//====account_privacyRules#554abb6f====

type TL_account_privacyRules struct {
	_rules []byte
	_users []byte
}

func New_TL_account_privacyRules() *TL_account_privacyRules {
	return &TL_account_privacyRules{}
}

func (t *TL_account_privacyRules) Encode() []byte {
	return nil
}

func (t *TL_account_privacyRules) Decode(b []byte) {}

//====accountDaysTTL#b8d0afdf====

type TL_accountDaysTTL struct {
	_days []byte
}

func New_TL_accountDaysTTL() *TL_accountDaysTTL {
	return &TL_accountDaysTTL{}
}

func (t *TL_accountDaysTTL) Encode() []byte {
	return nil
}

func (t *TL_accountDaysTTL) Decode(b []byte) {}

//====documentAttributeImageSize#6c37c15c====

type TL_documentAttributeImageSize struct {
	_w []byte
	_h []byte
}

func New_TL_documentAttributeImageSize() *TL_documentAttributeImageSize {
	return &TL_documentAttributeImageSize{}
}

func (t *TL_documentAttributeImageSize) Encode() []byte {
	return nil
}

func (t *TL_documentAttributeImageSize) Decode(b []byte) {}

//====documentAttributeAnimated#11b58939====

type TL_documentAttributeAnimated struct {
}

func New_TL_documentAttributeAnimated() *TL_documentAttributeAnimated {
	return &TL_documentAttributeAnimated{}
}

func (t *TL_documentAttributeAnimated) Encode() []byte {
	return nil
}

func (t *TL_documentAttributeAnimated) Decode(b []byte) {}

//====documentAttributeSticker#6319d612====

type TL_documentAttributeSticker struct {
	_flags       []byte
	_mask        []byte
	_alt         string
	_stickerset  []byte
	_mask_coords []byte
}

func New_TL_documentAttributeSticker() *TL_documentAttributeSticker {
	return &TL_documentAttributeSticker{}
}

func (t *TL_documentAttributeSticker) Encode() []byte {
	return nil
}

func (t *TL_documentAttributeSticker) Decode(b []byte) {}

//====documentAttributeVideo#0ef02ce6====

type TL_documentAttributeVideo struct {
	_flags         []byte
	_round_message []byte
	_duration      []byte
	_w             []byte
	_h             []byte
}

func New_TL_documentAttributeVideo() *TL_documentAttributeVideo {
	return &TL_documentAttributeVideo{}
}

func (t *TL_documentAttributeVideo) Encode() []byte {
	return nil
}

func (t *TL_documentAttributeVideo) Decode(b []byte) {}

//====documentAttributeAudio#9852f9c6====

type TL_documentAttributeAudio struct {
	_flags     []byte
	_voice     []byte
	_duration  []byte
	_title     []byte
	_performer []byte
	_waveform  []byte
}

func New_TL_documentAttributeAudio() *TL_documentAttributeAudio {
	return &TL_documentAttributeAudio{}
}

func (t *TL_documentAttributeAudio) Encode() []byte {
	return nil
}

func (t *TL_documentAttributeAudio) Decode(b []byte) {}

//====documentAttributeFilename#15590068====

type TL_documentAttributeFilename struct {
	_file_name string
}

func New_TL_documentAttributeFilename() *TL_documentAttributeFilename {
	return &TL_documentAttributeFilename{}
}

func (t *TL_documentAttributeFilename) Encode() []byte {
	return nil
}

func (t *TL_documentAttributeFilename) Decode(b []byte) {}

//====documentAttributeHasStickers#9801d2f7====

type TL_documentAttributeHasStickers struct {
}

func New_TL_documentAttributeHasStickers() *TL_documentAttributeHasStickers {
	return &TL_documentAttributeHasStickers{}
}

func (t *TL_documentAttributeHasStickers) Encode() []byte {
	return nil
}

func (t *TL_documentAttributeHasStickers) Decode(b []byte) {}

//====messages_stickersNotModified#f1749a22====

type TL_messages_stickersNotModified struct {
}

func New_TL_messages_stickersNotModified() *TL_messages_stickersNotModified {
	return &TL_messages_stickersNotModified{}
}

func (t *TL_messages_stickersNotModified) Encode() []byte {
	return nil
}

func (t *TL_messages_stickersNotModified) Decode(b []byte) {}

//====messages_stickers#8a8ecd32====

type TL_messages_stickers struct {
	_hash     string
	_stickers []byte
}

func New_TL_messages_stickers() *TL_messages_stickers {
	return &TL_messages_stickers{}
}

func (t *TL_messages_stickers) Encode() []byte {
	return nil
}

func (t *TL_messages_stickers) Decode(b []byte) {}

//====stickerPack#12b299d4====

type TL_stickerPack struct {
	_emoticon  string
	_documents []byte
}

func New_TL_stickerPack() *TL_stickerPack {
	return &TL_stickerPack{}
}

func (t *TL_stickerPack) Encode() []byte {
	return nil
}

func (t *TL_stickerPack) Decode(b []byte) {}

//====messages_allStickersNotModified#e86602c3====

type TL_messages_allStickersNotModified struct {
}

func New_TL_messages_allStickersNotModified() *TL_messages_allStickersNotModified {
	return &TL_messages_allStickersNotModified{}
}

func (t *TL_messages_allStickersNotModified) Encode() []byte {
	return nil
}

func (t *TL_messages_allStickersNotModified) Decode(b []byte) {}

//====messages_allStickers#edfd405f====

type TL_messages_allStickers struct {
	_hash []byte
	_sets []byte
}

func New_TL_messages_allStickers() *TL_messages_allStickers {
	return &TL_messages_allStickers{}
}

func (t *TL_messages_allStickers) Encode() []byte {
	return nil
}

func (t *TL_messages_allStickers) Decode(b []byte) {}

//====disabledFeature#ae636f24====

type TL_disabledFeature struct {
	_feature     string
	_description string
}

func New_TL_disabledFeature() *TL_disabledFeature {
	return &TL_disabledFeature{}
}

func (t *TL_disabledFeature) Encode() []byte {
	return nil
}

func (t *TL_disabledFeature) Decode(b []byte) {}

//====messages_affectedMessages#84d19185====

type TL_messages_affectedMessages struct {
	_pts       []byte
	_pts_count []byte
}

func New_TL_messages_affectedMessages() *TL_messages_affectedMessages {
	return &TL_messages_affectedMessages{}
}

func (t *TL_messages_affectedMessages) Encode() []byte {
	return nil
}

func (t *TL_messages_affectedMessages) Decode(b []byte) {}

//====contactLinkUnknown#5f4f9247====

type TL_contactLinkUnknown struct {
}

func New_TL_contactLinkUnknown() *TL_contactLinkUnknown {
	return &TL_contactLinkUnknown{}
}

func (t *TL_contactLinkUnknown) Encode() []byte {
	return nil
}

func (t *TL_contactLinkUnknown) Decode(b []byte) {}

//====contactLinkNone#feedd3ad====

type TL_contactLinkNone struct {
}

func New_TL_contactLinkNone() *TL_contactLinkNone {
	return &TL_contactLinkNone{}
}

func (t *TL_contactLinkNone) Encode() []byte {
	return nil
}

func (t *TL_contactLinkNone) Decode(b []byte) {}

//====contactLinkHasPhone#268f3f59====

type TL_contactLinkHasPhone struct {
}

func New_TL_contactLinkHasPhone() *TL_contactLinkHasPhone {
	return &TL_contactLinkHasPhone{}
}

func (t *TL_contactLinkHasPhone) Encode() []byte {
	return nil
}

func (t *TL_contactLinkHasPhone) Decode(b []byte) {}

//====contactLinkContact#d502c2d0====

type TL_contactLinkContact struct {
}

func New_TL_contactLinkContact() *TL_contactLinkContact {
	return &TL_contactLinkContact{}
}

func (t *TL_contactLinkContact) Encode() []byte {
	return nil
}

func (t *TL_contactLinkContact) Decode(b []byte) {}

//====webPageEmpty#eb1477e8====

type TL_webPageEmpty struct {
	_id int64
}

func New_TL_webPageEmpty() *TL_webPageEmpty {
	return &TL_webPageEmpty{}
}

func (t *TL_webPageEmpty) Encode() []byte {
	return nil
}

func (t *TL_webPageEmpty) Decode(b []byte) {}

//====webPagePending#c586da1c====

type TL_webPagePending struct {
	_id   int64
	_date []byte
}

func New_TL_webPagePending() *TL_webPagePending {
	return &TL_webPagePending{}
}

func (t *TL_webPagePending) Encode() []byte {
	return nil
}

func (t *TL_webPagePending) Decode(b []byte) {}

//====webPage#5f07b4bc====

type TL_webPage struct {
	_flags        []byte
	_id           int64
	_url          string
	_display_url  string
	_hash         []byte
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

func New_TL_webPage() *TL_webPage {
	return &TL_webPage{}
}

func (t *TL_webPage) Encode() []byte {
	return nil
}

func (t *TL_webPage) Decode(b []byte) {}

//====webPageNotModified#85849473====

type TL_webPageNotModified struct {
}

func New_TL_webPageNotModified() *TL_webPageNotModified {
	return &TL_webPageNotModified{}
}

func (t *TL_webPageNotModified) Encode() []byte {
	return nil
}

func (t *TL_webPageNotModified) Decode(b []byte) {}

//====authorization#7bf2e6f6====

type TL_authorization struct {
	_hash           int64
	_flags          []byte
	_device_model   string
	_platform       string
	_system_version string
	_api_id         []byte
	_app_name       string
	_app_version    string
	_date_created   []byte
	_date_active    []byte
	_ip             string
	_country        string
	_region         string
}

func New_TL_authorization() *TL_authorization {
	return &TL_authorization{}
}

func (t *TL_authorization) Encode() []byte {
	return nil
}

func (t *TL_authorization) Decode(b []byte) {}

//====account_authorizations#1250abde====

type TL_account_authorizations struct {
	_authorizations []byte
}

func New_TL_account_authorizations() *TL_account_authorizations {
	return &TL_account_authorizations{}
}

func (t *TL_account_authorizations) Encode() []byte {
	return nil
}

func (t *TL_account_authorizations) Decode(b []byte) {}

//====account_noPassword#96dabc18====

type TL_account_noPassword struct {
	_new_salt                  []byte
	_email_unconfirmed_pattern string
}

func New_TL_account_noPassword() *TL_account_noPassword {
	return &TL_account_noPassword{}
}

func (t *TL_account_noPassword) Encode() []byte {
	return nil
}

func (t *TL_account_noPassword) Decode(b []byte) {}

//====account_password#7c18141c====

type TL_account_password struct {
	_current_salt              []byte
	_new_salt                  []byte
	_hint                      string
	_has_recovery              []byte
	_email_unconfirmed_pattern string
}

func New_TL_account_password() *TL_account_password {
	return &TL_account_password{}
}

func (t *TL_account_password) Encode() []byte {
	return nil
}

func (t *TL_account_password) Decode(b []byte) {}

//====account_passwordSettings#b7b72ab3====

type TL_account_passwordSettings struct {
	_email string
}

func New_TL_account_passwordSettings() *TL_account_passwordSettings {
	return &TL_account_passwordSettings{}
}

func (t *TL_account_passwordSettings) Encode() []byte {
	return nil
}

func (t *TL_account_passwordSettings) Decode(b []byte) {}

//====account_passwordInputSettings#86916deb====

type TL_account_passwordInputSettings struct {
	_flags             []byte
	_new_salt          []byte
	_new_password_hash []byte
	_hint              []byte
	_email             []byte
}

func New_TL_account_passwordInputSettings() *TL_account_passwordInputSettings {
	return &TL_account_passwordInputSettings{}
}

func (t *TL_account_passwordInputSettings) Encode() []byte {
	return nil
}

func (t *TL_account_passwordInputSettings) Decode(b []byte) {}

//====auth_passwordRecovery#137948a5====

type TL_auth_passwordRecovery struct {
	_email_pattern string
}

func New_TL_auth_passwordRecovery() *TL_auth_passwordRecovery {
	return &TL_auth_passwordRecovery{}
}

func (t *TL_auth_passwordRecovery) Encode() []byte {
	return nil
}

func (t *TL_auth_passwordRecovery) Decode(b []byte) {}

//====receivedNotifyMessage#a384b779====

type TL_receivedNotifyMessage struct {
	_id    []byte
	_flags []byte
}

func New_TL_receivedNotifyMessage() *TL_receivedNotifyMessage {
	return &TL_receivedNotifyMessage{}
}

func (t *TL_receivedNotifyMessage) Encode() []byte {
	return nil
}

func (t *TL_receivedNotifyMessage) Decode(b []byte) {}

//====chatInviteEmpty#69df3769====

type TL_chatInviteEmpty struct {
}

func New_TL_chatInviteEmpty() *TL_chatInviteEmpty {
	return &TL_chatInviteEmpty{}
}

func (t *TL_chatInviteEmpty) Encode() []byte {
	return nil
}

func (t *TL_chatInviteEmpty) Decode(b []byte) {}

//====chatInviteExported#fc2e05bc====

type TL_chatInviteExported struct {
	_link string
}

func New_TL_chatInviteExported() *TL_chatInviteExported {
	return &TL_chatInviteExported{}
}

func (t *TL_chatInviteExported) Encode() []byte {
	return nil
}

func (t *TL_chatInviteExported) Decode(b []byte) {}

//====chatInviteAlready#5a686d7c====

type TL_chatInviteAlready struct {
	_chat []byte
}

func New_TL_chatInviteAlready() *TL_chatInviteAlready {
	return &TL_chatInviteAlready{}
}

func (t *TL_chatInviteAlready) Encode() []byte {
	return nil
}

func (t *TL_chatInviteAlready) Decode(b []byte) {}

//====chatInvite#db74f558====

type TL_chatInvite struct {
	_flags              []byte
	_channel            []byte
	_broadcast          []byte
	_public             []byte
	_megagroup          []byte
	_title              string
	_photo              []byte
	_participants_count []byte
	_participants       []byte
}

func New_TL_chatInvite() *TL_chatInvite {
	return &TL_chatInvite{}
}

func (t *TL_chatInvite) Encode() []byte {
	return nil
}

func (t *TL_chatInvite) Decode(b []byte) {}

//====inputStickerSetEmpty#ffb62b95====

type TL_inputStickerSetEmpty struct {
}

func New_TL_inputStickerSetEmpty() *TL_inputStickerSetEmpty {
	return &TL_inputStickerSetEmpty{}
}

func (t *TL_inputStickerSetEmpty) Encode() []byte {
	return nil
}

func (t *TL_inputStickerSetEmpty) Decode(b []byte) {}

//====inputStickerSetID#9de7a269====

type TL_inputStickerSetID struct {
	_id          int64
	_access_hash int64
}

func New_TL_inputStickerSetID() *TL_inputStickerSetID {
	return &TL_inputStickerSetID{}
}

func (t *TL_inputStickerSetID) Encode() []byte {
	return nil
}

func (t *TL_inputStickerSetID) Decode(b []byte) {}

//====inputStickerSetShortName#861cc8a0====

type TL_inputStickerSetShortName struct {
	_short_name string
}

func New_TL_inputStickerSetShortName() *TL_inputStickerSetShortName {
	return &TL_inputStickerSetShortName{}
}

func (t *TL_inputStickerSetShortName) Encode() []byte {
	return nil
}

func (t *TL_inputStickerSetShortName) Decode(b []byte) {}

//====stickerSet#cd303b41====

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
	_count       []byte
	_hash        []byte
}

func New_TL_stickerSet() *TL_stickerSet {
	return &TL_stickerSet{}
}

func (t *TL_stickerSet) Encode() []byte {
	return nil
}

func (t *TL_stickerSet) Decode(b []byte) {}

//====messages_stickerSet#b60a24a6====

type TL_messages_stickerSet struct {
	_set       []byte
	_packs     []byte
	_documents []byte
}

func New_TL_messages_stickerSet() *TL_messages_stickerSet {
	return &TL_messages_stickerSet{}
}

func (t *TL_messages_stickerSet) Encode() []byte {
	return nil
}

func (t *TL_messages_stickerSet) Decode(b []byte) {}

//====botCommand#c27ac8c7====

type TL_botCommand struct {
	_command     string
	_description string
}

func New_TL_botCommand() *TL_botCommand {
	return &TL_botCommand{}
}

func (t *TL_botCommand) Encode() []byte {
	return nil
}

func (t *TL_botCommand) Decode(b []byte) {}

//====botInfo#98e81d3a====

type TL_botInfo struct {
	_user_id     []byte
	_description string
	_commands    []byte
}

func New_TL_botInfo() *TL_botInfo {
	return &TL_botInfo{}
}

func (t *TL_botInfo) Encode() []byte {
	return nil
}

func (t *TL_botInfo) Decode(b []byte) {}

//====keyboardButton#a2fa4880====

type TL_keyboardButton struct {
	_text string
}

func New_TL_keyboardButton() *TL_keyboardButton {
	return &TL_keyboardButton{}
}

func (t *TL_keyboardButton) Encode() []byte {
	return nil
}

func (t *TL_keyboardButton) Decode(b []byte) {}

//====keyboardButtonUrl#258aff05====

type TL_keyboardButtonUrl struct {
	_text string
	_url  string
}

func New_TL_keyboardButtonUrl() *TL_keyboardButtonUrl {
	return &TL_keyboardButtonUrl{}
}

func (t *TL_keyboardButtonUrl) Encode() []byte {
	return nil
}

func (t *TL_keyboardButtonUrl) Decode(b []byte) {}

//====keyboardButtonCallback#683a5e46====

type TL_keyboardButtonCallback struct {
	_text string
	_data []byte
}

func New_TL_keyboardButtonCallback() *TL_keyboardButtonCallback {
	return &TL_keyboardButtonCallback{}
}

func (t *TL_keyboardButtonCallback) Encode() []byte {
	return nil
}

func (t *TL_keyboardButtonCallback) Decode(b []byte) {}

//====keyboardButtonRequestPhone#b16a6c29====

type TL_keyboardButtonRequestPhone struct {
	_text string
}

func New_TL_keyboardButtonRequestPhone() *TL_keyboardButtonRequestPhone {
	return &TL_keyboardButtonRequestPhone{}
}

func (t *TL_keyboardButtonRequestPhone) Encode() []byte {
	return nil
}

func (t *TL_keyboardButtonRequestPhone) Decode(b []byte) {}

//====keyboardButtonRequestGeoLocation#fc796b3f====

type TL_keyboardButtonRequestGeoLocation struct {
	_text string
}

func New_TL_keyboardButtonRequestGeoLocation() *TL_keyboardButtonRequestGeoLocation {
	return &TL_keyboardButtonRequestGeoLocation{}
}

func (t *TL_keyboardButtonRequestGeoLocation) Encode() []byte {
	return nil
}

func (t *TL_keyboardButtonRequestGeoLocation) Decode(b []byte) {}

//====keyboardButtonSwitchInline#0568a748====

type TL_keyboardButtonSwitchInline struct {
	_flags     []byte
	_same_peer []byte
	_text      string
	_query     string
}

func New_TL_keyboardButtonSwitchInline() *TL_keyboardButtonSwitchInline {
	return &TL_keyboardButtonSwitchInline{}
}

func (t *TL_keyboardButtonSwitchInline) Encode() []byte {
	return nil
}

func (t *TL_keyboardButtonSwitchInline) Decode(b []byte) {}

//====keyboardButtonGame#50f41ccf====

type TL_keyboardButtonGame struct {
	_text string
}

func New_TL_keyboardButtonGame() *TL_keyboardButtonGame {
	return &TL_keyboardButtonGame{}
}

func (t *TL_keyboardButtonGame) Encode() []byte {
	return nil
}

func (t *TL_keyboardButtonGame) Decode(b []byte) {}

//====keyboardButtonBuy#afd93fbb====

type TL_keyboardButtonBuy struct {
	_text string
}

func New_TL_keyboardButtonBuy() *TL_keyboardButtonBuy {
	return &TL_keyboardButtonBuy{}
}

func (t *TL_keyboardButtonBuy) Encode() []byte {
	return nil
}

func (t *TL_keyboardButtonBuy) Decode(b []byte) {}

//====keyboardButtonRow#77608b83====

type TL_keyboardButtonRow struct {
	_buttons []byte
}

func New_TL_keyboardButtonRow() *TL_keyboardButtonRow {
	return &TL_keyboardButtonRow{}
}

func (t *TL_keyboardButtonRow) Encode() []byte {
	return nil
}

func (t *TL_keyboardButtonRow) Decode(b []byte) {}

//====replyKeyboardHide#a03e5b85====

type TL_replyKeyboardHide struct {
	_flags     []byte
	_selective []byte
}

func New_TL_replyKeyboardHide() *TL_replyKeyboardHide {
	return &TL_replyKeyboardHide{}
}

func (t *TL_replyKeyboardHide) Encode() []byte {
	return nil
}

func (t *TL_replyKeyboardHide) Decode(b []byte) {}

//====replyKeyboardForceReply#f4108aa0====

type TL_replyKeyboardForceReply struct {
	_flags      []byte
	_single_use []byte
	_selective  []byte
}

func New_TL_replyKeyboardForceReply() *TL_replyKeyboardForceReply {
	return &TL_replyKeyboardForceReply{}
}

func (t *TL_replyKeyboardForceReply) Encode() []byte {
	return nil
}

func (t *TL_replyKeyboardForceReply) Decode(b []byte) {}

//====replyKeyboardMarkup#3502758c====

type TL_replyKeyboardMarkup struct {
	_flags      []byte
	_resize     []byte
	_single_use []byte
	_selective  []byte
	_rows       []byte
}

func New_TL_replyKeyboardMarkup() *TL_replyKeyboardMarkup {
	return &TL_replyKeyboardMarkup{}
}

func (t *TL_replyKeyboardMarkup) Encode() []byte {
	return nil
}

func (t *TL_replyKeyboardMarkup) Decode(b []byte) {}

//====replyInlineMarkup#48a30254====

type TL_replyInlineMarkup struct {
	_rows []byte
}

func New_TL_replyInlineMarkup() *TL_replyInlineMarkup {
	return &TL_replyInlineMarkup{}
}

func (t *TL_replyInlineMarkup) Encode() []byte {
	return nil
}

func (t *TL_replyInlineMarkup) Decode(b []byte) {}

//====messageEntityUnknown#bb92ba95====

type TL_messageEntityUnknown struct {
	_offset []byte
	_length []byte
}

func New_TL_messageEntityUnknown() *TL_messageEntityUnknown {
	return &TL_messageEntityUnknown{}
}

func (t *TL_messageEntityUnknown) Encode() []byte {
	return nil
}

func (t *TL_messageEntityUnknown) Decode(b []byte) {}

//====messageEntityMention#fa04579d====

type TL_messageEntityMention struct {
	_offset []byte
	_length []byte
}

func New_TL_messageEntityMention() *TL_messageEntityMention {
	return &TL_messageEntityMention{}
}

func (t *TL_messageEntityMention) Encode() []byte {
	return nil
}

func (t *TL_messageEntityMention) Decode(b []byte) {}

//====messageEntityHashtag#6f635b0d====

type TL_messageEntityHashtag struct {
	_offset []byte
	_length []byte
}

func New_TL_messageEntityHashtag() *TL_messageEntityHashtag {
	return &TL_messageEntityHashtag{}
}

func (t *TL_messageEntityHashtag) Encode() []byte {
	return nil
}

func (t *TL_messageEntityHashtag) Decode(b []byte) {}

//====messageEntityBotCommand#6cef8ac7====

type TL_messageEntityBotCommand struct {
	_offset []byte
	_length []byte
}

func New_TL_messageEntityBotCommand() *TL_messageEntityBotCommand {
	return &TL_messageEntityBotCommand{}
}

func (t *TL_messageEntityBotCommand) Encode() []byte {
	return nil
}

func (t *TL_messageEntityBotCommand) Decode(b []byte) {}

//====messageEntityUrl#6ed02538====

type TL_messageEntityUrl struct {
	_offset []byte
	_length []byte
}

func New_TL_messageEntityUrl() *TL_messageEntityUrl {
	return &TL_messageEntityUrl{}
}

func (t *TL_messageEntityUrl) Encode() []byte {
	return nil
}

func (t *TL_messageEntityUrl) Decode(b []byte) {}

//====messageEntityEmail#64e475c2====

type TL_messageEntityEmail struct {
	_offset []byte
	_length []byte
}

func New_TL_messageEntityEmail() *TL_messageEntityEmail {
	return &TL_messageEntityEmail{}
}

func (t *TL_messageEntityEmail) Encode() []byte {
	return nil
}

func (t *TL_messageEntityEmail) Decode(b []byte) {}

//====messageEntityBold#bd610bc9====

type TL_messageEntityBold struct {
	_offset []byte
	_length []byte
}

func New_TL_messageEntityBold() *TL_messageEntityBold {
	return &TL_messageEntityBold{}
}

func (t *TL_messageEntityBold) Encode() []byte {
	return nil
}

func (t *TL_messageEntityBold) Decode(b []byte) {}

//====messageEntityItalic#826f8b60====

type TL_messageEntityItalic struct {
	_offset []byte
	_length []byte
}

func New_TL_messageEntityItalic() *TL_messageEntityItalic {
	return &TL_messageEntityItalic{}
}

func (t *TL_messageEntityItalic) Encode() []byte {
	return nil
}

func (t *TL_messageEntityItalic) Decode(b []byte) {}

//====messageEntityCode#28a20571====

type TL_messageEntityCode struct {
	_offset []byte
	_length []byte
}

func New_TL_messageEntityCode() *TL_messageEntityCode {
	return &TL_messageEntityCode{}
}

func (t *TL_messageEntityCode) Encode() []byte {
	return nil
}

func (t *TL_messageEntityCode) Decode(b []byte) {}

//====messageEntityPre#73924be0====

type TL_messageEntityPre struct {
	_offset   []byte
	_length   []byte
	_language string
}

func New_TL_messageEntityPre() *TL_messageEntityPre {
	return &TL_messageEntityPre{}
}

func (t *TL_messageEntityPre) Encode() []byte {
	return nil
}

func (t *TL_messageEntityPre) Decode(b []byte) {}

//====messageEntityTextUrl#76a6d327====

type TL_messageEntityTextUrl struct {
	_offset []byte
	_length []byte
	_url    string
}

func New_TL_messageEntityTextUrl() *TL_messageEntityTextUrl {
	return &TL_messageEntityTextUrl{}
}

func (t *TL_messageEntityTextUrl) Encode() []byte {
	return nil
}

func (t *TL_messageEntityTextUrl) Decode(b []byte) {}

//====messageEntityMentionName#352dca58====

type TL_messageEntityMentionName struct {
	_offset  []byte
	_length  []byte
	_user_id []byte
}

func New_TL_messageEntityMentionName() *TL_messageEntityMentionName {
	return &TL_messageEntityMentionName{}
}

func (t *TL_messageEntityMentionName) Encode() []byte {
	return nil
}

func (t *TL_messageEntityMentionName) Decode(b []byte) {}

//====inputMessageEntityMentionName#208e68c9====

type TL_inputMessageEntityMentionName struct {
	_offset  []byte
	_length  []byte
	_user_id []byte
}

func New_TL_inputMessageEntityMentionName() *TL_inputMessageEntityMentionName {
	return &TL_inputMessageEntityMentionName{}
}

func (t *TL_inputMessageEntityMentionName) Encode() []byte {
	return nil
}

func (t *TL_inputMessageEntityMentionName) Decode(b []byte) {}

//====inputChannelEmpty#ee8c1e86====

type TL_inputChannelEmpty struct {
}

func New_TL_inputChannelEmpty() *TL_inputChannelEmpty {
	return &TL_inputChannelEmpty{}
}

func (t *TL_inputChannelEmpty) Encode() []byte {
	return nil
}

func (t *TL_inputChannelEmpty) Decode(b []byte) {}

//====inputChannel#afeb712e====

type TL_inputChannel struct {
	_channel_id  []byte
	_access_hash int64
}

func New_TL_inputChannel() *TL_inputChannel {
	return &TL_inputChannel{}
}

func (t *TL_inputChannel) Encode() []byte {
	return nil
}

func (t *TL_inputChannel) Decode(b []byte) {}

//====contacts_resolvedPeer#7f077ad9====

type TL_contacts_resolvedPeer struct {
	_peer  []byte
	_chats []byte
	_users []byte
}

func New_TL_contacts_resolvedPeer() *TL_contacts_resolvedPeer {
	return &TL_contacts_resolvedPeer{}
}

func (t *TL_contacts_resolvedPeer) Encode() []byte {
	return nil
}

func (t *TL_contacts_resolvedPeer) Decode(b []byte) {}

//====messageRange#0ae30253====

type TL_messageRange struct {
	_min_id []byte
	_max_id []byte
}

func New_TL_messageRange() *TL_messageRange {
	return &TL_messageRange{}
}

func (t *TL_messageRange) Encode() []byte {
	return nil
}

func (t *TL_messageRange) Decode(b []byte) {}

//====updates_channelDifferenceEmpty#3e11affb====

type TL_updates_channelDifferenceEmpty struct {
	_flags   []byte
	_final   []byte
	_pts     []byte
	_timeout []byte
}

func New_TL_updates_channelDifferenceEmpty() *TL_updates_channelDifferenceEmpty {
	return &TL_updates_channelDifferenceEmpty{}
}

func (t *TL_updates_channelDifferenceEmpty) Encode() []byte {
	return nil
}

func (t *TL_updates_channelDifferenceEmpty) Decode(b []byte) {}

//====updates_channelDifferenceTooLong#6a9d7b35====

type TL_updates_channelDifferenceTooLong struct {
	_flags                 []byte
	_final                 []byte
	_pts                   []byte
	_timeout               []byte
	_top_message           []byte
	_read_inbox_max_id     []byte
	_read_outbox_max_id    []byte
	_unread_count          []byte
	_unread_mentions_count []byte
	_messages              []byte
	_chats                 []byte
	_users                 []byte
}

func New_TL_updates_channelDifferenceTooLong() *TL_updates_channelDifferenceTooLong {
	return &TL_updates_channelDifferenceTooLong{}
}

func (t *TL_updates_channelDifferenceTooLong) Encode() []byte {
	return nil
}

func (t *TL_updates_channelDifferenceTooLong) Decode(b []byte) {}

//====updates_channelDifference#2064674e====

type TL_updates_channelDifference struct {
	_flags         []byte
	_final         []byte
	_pts           []byte
	_timeout       []byte
	_new_messages  []byte
	_other_updates []byte
	_chats         []byte
	_users         []byte
}

func New_TL_updates_channelDifference() *TL_updates_channelDifference {
	return &TL_updates_channelDifference{}
}

func (t *TL_updates_channelDifference) Encode() []byte {
	return nil
}

func (t *TL_updates_channelDifference) Decode(b []byte) {}

//====channelMessagesFilterEmpty#94d42ee7====

type TL_channelMessagesFilterEmpty struct {
}

func New_TL_channelMessagesFilterEmpty() *TL_channelMessagesFilterEmpty {
	return &TL_channelMessagesFilterEmpty{}
}

func (t *TL_channelMessagesFilterEmpty) Encode() []byte {
	return nil
}

func (t *TL_channelMessagesFilterEmpty) Decode(b []byte) {}

//====channelMessagesFilter#cd77d957====

type TL_channelMessagesFilter struct {
	_flags                []byte
	_exclude_new_messages []byte
	_ranges               []byte
}

func New_TL_channelMessagesFilter() *TL_channelMessagesFilter {
	return &TL_channelMessagesFilter{}
}

func (t *TL_channelMessagesFilter) Encode() []byte {
	return nil
}

func (t *TL_channelMessagesFilter) Decode(b []byte) {}

//====channelParticipant#15ebac1d====

type TL_channelParticipant struct {
	_user_id []byte
	_date    []byte
}

func New_TL_channelParticipant() *TL_channelParticipant {
	return &TL_channelParticipant{}
}

func (t *TL_channelParticipant) Encode() []byte {
	return nil
}

func (t *TL_channelParticipant) Decode(b []byte) {}

//====channelParticipantSelf#a3289a6d====

type TL_channelParticipantSelf struct {
	_user_id    []byte
	_inviter_id []byte
	_date       []byte
}

func New_TL_channelParticipantSelf() *TL_channelParticipantSelf {
	return &TL_channelParticipantSelf{}
}

func (t *TL_channelParticipantSelf) Encode() []byte {
	return nil
}

func (t *TL_channelParticipantSelf) Decode(b []byte) {}

//====channelParticipantCreator#e3e2e1f9====

type TL_channelParticipantCreator struct {
	_user_id []byte
}

func New_TL_channelParticipantCreator() *TL_channelParticipantCreator {
	return &TL_channelParticipantCreator{}
}

func (t *TL_channelParticipantCreator) Encode() []byte {
	return nil
}

func (t *TL_channelParticipantCreator) Decode(b []byte) {}

//====channelParticipantAdmin#a82fa898====

type TL_channelParticipantAdmin struct {
	_flags        []byte
	_can_edit     []byte
	_user_id      []byte
	_inviter_id   []byte
	_promoted_by  []byte
	_date         []byte
	_admin_rights []byte
}

func New_TL_channelParticipantAdmin() *TL_channelParticipantAdmin {
	return &TL_channelParticipantAdmin{}
}

func (t *TL_channelParticipantAdmin) Encode() []byte {
	return nil
}

func (t *TL_channelParticipantAdmin) Decode(b []byte) {}

//====channelParticipantBanned#222c1886====

type TL_channelParticipantBanned struct {
	_flags         []byte
	_left          []byte
	_user_id       []byte
	_kicked_by     []byte
	_date          []byte
	_banned_rights []byte
}

func New_TL_channelParticipantBanned() *TL_channelParticipantBanned {
	return &TL_channelParticipantBanned{}
}

func (t *TL_channelParticipantBanned) Encode() []byte {
	return nil
}

func (t *TL_channelParticipantBanned) Decode(b []byte) {}

//====channelParticipantsRecent#de3f3c79====

type TL_channelParticipantsRecent struct {
}

func New_TL_channelParticipantsRecent() *TL_channelParticipantsRecent {
	return &TL_channelParticipantsRecent{}
}

func (t *TL_channelParticipantsRecent) Encode() []byte {
	return nil
}

func (t *TL_channelParticipantsRecent) Decode(b []byte) {}

//====channelParticipantsAdmins#b4608969====

type TL_channelParticipantsAdmins struct {
}

func New_TL_channelParticipantsAdmins() *TL_channelParticipantsAdmins {
	return &TL_channelParticipantsAdmins{}
}

func (t *TL_channelParticipantsAdmins) Encode() []byte {
	return nil
}

func (t *TL_channelParticipantsAdmins) Decode(b []byte) {}

//====channelParticipantsKicked#a3b54985====

type TL_channelParticipantsKicked struct {
	_q string
}

func New_TL_channelParticipantsKicked() *TL_channelParticipantsKicked {
	return &TL_channelParticipantsKicked{}
}

func (t *TL_channelParticipantsKicked) Encode() []byte {
	return nil
}

func (t *TL_channelParticipantsKicked) Decode(b []byte) {}

//====channelParticipantsBots#b0d1865b====

type TL_channelParticipantsBots struct {
}

func New_TL_channelParticipantsBots() *TL_channelParticipantsBots {
	return &TL_channelParticipantsBots{}
}

func (t *TL_channelParticipantsBots) Encode() []byte {
	return nil
}

func (t *TL_channelParticipantsBots) Decode(b []byte) {}

//====channelParticipantsBanned#1427a5e1====

type TL_channelParticipantsBanned struct {
	_q string
}

func New_TL_channelParticipantsBanned() *TL_channelParticipantsBanned {
	return &TL_channelParticipantsBanned{}
}

func (t *TL_channelParticipantsBanned) Encode() []byte {
	return nil
}

func (t *TL_channelParticipantsBanned) Decode(b []byte) {}

//====channelParticipantsSearch#0656ac4b====

type TL_channelParticipantsSearch struct {
	_q string
}

func New_TL_channelParticipantsSearch() *TL_channelParticipantsSearch {
	return &TL_channelParticipantsSearch{}
}

func (t *TL_channelParticipantsSearch) Encode() []byte {
	return nil
}

func (t *TL_channelParticipantsSearch) Decode(b []byte) {}

//====channels_channelParticipants#f56ee2a8====

type TL_channels_channelParticipants struct {
	_count        []byte
	_participants []byte
	_users        []byte
}

func New_TL_channels_channelParticipants() *TL_channels_channelParticipants {
	return &TL_channels_channelParticipants{}
}

func (t *TL_channels_channelParticipants) Encode() []byte {
	return nil
}

func (t *TL_channels_channelParticipants) Decode(b []byte) {}

//====channels_channelParticipantsNotModified#f0173fe9====

type TL_channels_channelParticipantsNotModified struct {
}

func New_TL_channels_channelParticipantsNotModified() *TL_channels_channelParticipantsNotModified {
	return &TL_channels_channelParticipantsNotModified{}
}

func (t *TL_channels_channelParticipantsNotModified) Encode() []byte {
	return nil
}

func (t *TL_channels_channelParticipantsNotModified) Decode(b []byte) {}

//====channels_channelParticipant#d0d9b163====

type TL_channels_channelParticipant struct {
	_participant []byte
	_users       []byte
}

func New_TL_channels_channelParticipant() *TL_channels_channelParticipant {
	return &TL_channels_channelParticipant{}
}

func (t *TL_channels_channelParticipant) Encode() []byte {
	return nil
}

func (t *TL_channels_channelParticipant) Decode(b []byte) {}

//====help_termsOfService#f1ee3e90====

type TL_help_termsOfService struct {
	_text string
}

func New_TL_help_termsOfService() *TL_help_termsOfService {
	return &TL_help_termsOfService{}
}

func (t *TL_help_termsOfService) Encode() []byte {
	return nil
}

func (t *TL_help_termsOfService) Decode(b []byte) {}

//====foundGif#162ecc1f====

type TL_foundGif struct {
	_url          string
	_thumb_url    string
	_content_url  string
	_content_type string
	_w            []byte
	_h            []byte
}

func New_TL_foundGif() *TL_foundGif {
	return &TL_foundGif{}
}

func (t *TL_foundGif) Encode() []byte {
	return nil
}

func (t *TL_foundGif) Decode(b []byte) {}

//====foundGifCached#9c750409====

type TL_foundGifCached struct {
	_url      string
	_photo    []byte
	_document []byte
}

func New_TL_foundGifCached() *TL_foundGifCached {
	return &TL_foundGifCached{}
}

func (t *TL_foundGifCached) Encode() []byte {
	return nil
}

func (t *TL_foundGifCached) Decode(b []byte) {}

//====messages_foundGifs#450a1c0a====

type TL_messages_foundGifs struct {
	_next_offset []byte
	_results     []byte
}

func New_TL_messages_foundGifs() *TL_messages_foundGifs {
	return &TL_messages_foundGifs{}
}

func (t *TL_messages_foundGifs) Encode() []byte {
	return nil
}

func (t *TL_messages_foundGifs) Decode(b []byte) {}

//====messages_savedGifsNotModified#e8025ca2====

type TL_messages_savedGifsNotModified struct {
}

func New_TL_messages_savedGifsNotModified() *TL_messages_savedGifsNotModified {
	return &TL_messages_savedGifsNotModified{}
}

func (t *TL_messages_savedGifsNotModified) Encode() []byte {
	return nil
}

func (t *TL_messages_savedGifsNotModified) Decode(b []byte) {}

//====messages_savedGifs#2e0709a5====

type TL_messages_savedGifs struct {
	_hash []byte
	_gifs []byte
}

func New_TL_messages_savedGifs() *TL_messages_savedGifs {
	return &TL_messages_savedGifs{}
}

func (t *TL_messages_savedGifs) Encode() []byte {
	return nil
}

func (t *TL_messages_savedGifs) Decode(b []byte) {}

//====inputBotInlineMessageMediaAuto#292fed13====

type TL_inputBotInlineMessageMediaAuto struct {
	_flags        []byte
	_caption      string
	_reply_markup []byte
}

func New_TL_inputBotInlineMessageMediaAuto() *TL_inputBotInlineMessageMediaAuto {
	return &TL_inputBotInlineMessageMediaAuto{}
}

func (t *TL_inputBotInlineMessageMediaAuto) Encode() []byte {
	return nil
}

func (t *TL_inputBotInlineMessageMediaAuto) Decode(b []byte) {}

//====inputBotInlineMessageText#3dcd7a87====

type TL_inputBotInlineMessageText struct {
	_flags        []byte
	_no_webpage   []byte
	_message      string
	_entities     []byte
	_reply_markup []byte
}

func New_TL_inputBotInlineMessageText() *TL_inputBotInlineMessageText {
	return &TL_inputBotInlineMessageText{}
}

func (t *TL_inputBotInlineMessageText) Encode() []byte {
	return nil
}

func (t *TL_inputBotInlineMessageText) Decode(b []byte) {}

//====inputBotInlineMessageMediaGeo#c1b15d65====

type TL_inputBotInlineMessageMediaGeo struct {
	_flags        []byte
	_geo_point    []byte
	_period       []byte
	_reply_markup []byte
}

func New_TL_inputBotInlineMessageMediaGeo() *TL_inputBotInlineMessageMediaGeo {
	return &TL_inputBotInlineMessageMediaGeo{}
}

func (t *TL_inputBotInlineMessageMediaGeo) Encode() []byte {
	return nil
}

func (t *TL_inputBotInlineMessageMediaGeo) Decode(b []byte) {}

//====inputBotInlineMessageMediaVenue#aaafadc8====

type TL_inputBotInlineMessageMediaVenue struct {
	_flags        []byte
	_geo_point    []byte
	_title        string
	_address      string
	_provider     string
	_venue_id     string
	_reply_markup []byte
}

func New_TL_inputBotInlineMessageMediaVenue() *TL_inputBotInlineMessageMediaVenue {
	return &TL_inputBotInlineMessageMediaVenue{}
}

func (t *TL_inputBotInlineMessageMediaVenue) Encode() []byte {
	return nil
}

func (t *TL_inputBotInlineMessageMediaVenue) Decode(b []byte) {}

//====inputBotInlineMessageMediaContact#2daf01a7====

type TL_inputBotInlineMessageMediaContact struct {
	_flags        []byte
	_phone_number string
	_first_name   string
	_last_name    string
	_reply_markup []byte
}

func New_TL_inputBotInlineMessageMediaContact() *TL_inputBotInlineMessageMediaContact {
	return &TL_inputBotInlineMessageMediaContact{}
}

func (t *TL_inputBotInlineMessageMediaContact) Encode() []byte {
	return nil
}

func (t *TL_inputBotInlineMessageMediaContact) Decode(b []byte) {}

//====inputBotInlineMessageGame#4b425864====

type TL_inputBotInlineMessageGame struct {
	_flags        []byte
	_reply_markup []byte
}

func New_TL_inputBotInlineMessageGame() *TL_inputBotInlineMessageGame {
	return &TL_inputBotInlineMessageGame{}
}

func (t *TL_inputBotInlineMessageGame) Encode() []byte {
	return nil
}

func (t *TL_inputBotInlineMessageGame) Decode(b []byte) {}

//====inputBotInlineResult#2cbbe15a====

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

func New_TL_inputBotInlineResult() *TL_inputBotInlineResult {
	return &TL_inputBotInlineResult{}
}

func (t *TL_inputBotInlineResult) Encode() []byte {
	return nil
}

func (t *TL_inputBotInlineResult) Decode(b []byte) {}

//====inputBotInlineResultPhoto#a8d864a7====

type TL_inputBotInlineResultPhoto struct {
	_id           string
	_type         string
	_photo        []byte
	_send_message []byte
}

func New_TL_inputBotInlineResultPhoto() *TL_inputBotInlineResultPhoto {
	return &TL_inputBotInlineResultPhoto{}
}

func (t *TL_inputBotInlineResultPhoto) Encode() []byte {
	return nil
}

func (t *TL_inputBotInlineResultPhoto) Decode(b []byte) {}

//====inputBotInlineResultDocument#fff8fdc4====

type TL_inputBotInlineResultDocument struct {
	_flags        []byte
	_id           string
	_type         string
	_title        []byte
	_description  []byte
	_document     []byte
	_send_message []byte
}

func New_TL_inputBotInlineResultDocument() *TL_inputBotInlineResultDocument {
	return &TL_inputBotInlineResultDocument{}
}

func (t *TL_inputBotInlineResultDocument) Encode() []byte {
	return nil
}

func (t *TL_inputBotInlineResultDocument) Decode(b []byte) {}

//====inputBotInlineResultGame#4fa417f2====

type TL_inputBotInlineResultGame struct {
	_id           string
	_short_name   string
	_send_message []byte
}

func New_TL_inputBotInlineResultGame() *TL_inputBotInlineResultGame {
	return &TL_inputBotInlineResultGame{}
}

func (t *TL_inputBotInlineResultGame) Encode() []byte {
	return nil
}

func (t *TL_inputBotInlineResultGame) Decode(b []byte) {}

//====botInlineMessageMediaAuto#0a74b15b====

type TL_botInlineMessageMediaAuto struct {
	_flags        []byte
	_caption      string
	_reply_markup []byte
}

func New_TL_botInlineMessageMediaAuto() *TL_botInlineMessageMediaAuto {
	return &TL_botInlineMessageMediaAuto{}
}

func (t *TL_botInlineMessageMediaAuto) Encode() []byte {
	return nil
}

func (t *TL_botInlineMessageMediaAuto) Decode(b []byte) {}

//====botInlineMessageText#8c7f65e2====

type TL_botInlineMessageText struct {
	_flags        []byte
	_no_webpage   []byte
	_message      string
	_entities     []byte
	_reply_markup []byte
}

func New_TL_botInlineMessageText() *TL_botInlineMessageText {
	return &TL_botInlineMessageText{}
}

func (t *TL_botInlineMessageText) Encode() []byte {
	return nil
}

func (t *TL_botInlineMessageText) Decode(b []byte) {}

//====botInlineMessageMediaGeo#b722de65====

type TL_botInlineMessageMediaGeo struct {
	_flags        []byte
	_geo          []byte
	_period       []byte
	_reply_markup []byte
}

func New_TL_botInlineMessageMediaGeo() *TL_botInlineMessageMediaGeo {
	return &TL_botInlineMessageMediaGeo{}
}

func (t *TL_botInlineMessageMediaGeo) Encode() []byte {
	return nil
}

func (t *TL_botInlineMessageMediaGeo) Decode(b []byte) {}

//====botInlineMessageMediaVenue#4366232e====

type TL_botInlineMessageMediaVenue struct {
	_flags        []byte
	_geo          []byte
	_title        string
	_address      string
	_provider     string
	_venue_id     string
	_reply_markup []byte
}

func New_TL_botInlineMessageMediaVenue() *TL_botInlineMessageMediaVenue {
	return &TL_botInlineMessageMediaVenue{}
}

func (t *TL_botInlineMessageMediaVenue) Encode() []byte {
	return nil
}

func (t *TL_botInlineMessageMediaVenue) Decode(b []byte) {}

//====botInlineMessageMediaContact#35edb4d4====

type TL_botInlineMessageMediaContact struct {
	_flags        []byte
	_phone_number string
	_first_name   string
	_last_name    string
	_reply_markup []byte
}

func New_TL_botInlineMessageMediaContact() *TL_botInlineMessageMediaContact {
	return &TL_botInlineMessageMediaContact{}
}

func (t *TL_botInlineMessageMediaContact) Encode() []byte {
	return nil
}

func (t *TL_botInlineMessageMediaContact) Decode(b []byte) {}

//====botInlineResult#9bebaeb9====

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

func New_TL_botInlineResult() *TL_botInlineResult {
	return &TL_botInlineResult{}
}

func (t *TL_botInlineResult) Encode() []byte {
	return nil
}

func (t *TL_botInlineResult) Decode(b []byte) {}

//====botInlineMediaResult#17db940b====

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

func New_TL_botInlineMediaResult() *TL_botInlineMediaResult {
	return &TL_botInlineMediaResult{}
}

func (t *TL_botInlineMediaResult) Encode() []byte {
	return nil
}

func (t *TL_botInlineMediaResult) Decode(b []byte) {}

//====messages_botResults#947ca848====

type TL_messages_botResults struct {
	_flags       []byte
	_gallery     []byte
	_query_id    int64
	_next_offset []byte
	_switch_pm   []byte
	_results     []byte
	_cache_time  []byte
	_users       []byte
}

func New_TL_messages_botResults() *TL_messages_botResults {
	return &TL_messages_botResults{}
}

func (t *TL_messages_botResults) Encode() []byte {
	return nil
}

func (t *TL_messages_botResults) Decode(b []byte) {}

//====exportedMessageLink#1f486803====

type TL_exportedMessageLink struct {
	_link string
}

func New_TL_exportedMessageLink() *TL_exportedMessageLink {
	return &TL_exportedMessageLink{}
}

func (t *TL_exportedMessageLink) Encode() []byte {
	return nil
}

func (t *TL_exportedMessageLink) Decode(b []byte) {}

//====messageFwdHeader#559ebe6d====

type TL_messageFwdHeader struct {
	_flags             []byte
	_from_id           []byte
	_date              []byte
	_channel_id        []byte
	_channel_post      []byte
	_post_author       []byte
	_saved_from_peer   []byte
	_saved_from_msg_id []byte
}

func New_TL_messageFwdHeader() *TL_messageFwdHeader {
	return &TL_messageFwdHeader{}
}

func (t *TL_messageFwdHeader) Encode() []byte {
	return nil
}

func (t *TL_messageFwdHeader) Decode(b []byte) {}

//====auth_codeTypeSms#72a3158c====

type TL_auth_codeTypeSms struct {
}

func New_TL_auth_codeTypeSms() *TL_auth_codeTypeSms {
	return &TL_auth_codeTypeSms{}
}

func (t *TL_auth_codeTypeSms) Encode() []byte {
	return nil
}

func (t *TL_auth_codeTypeSms) Decode(b []byte) {}

//====auth_codeTypeCall#741cd3e3====

type TL_auth_codeTypeCall struct {
}

func New_TL_auth_codeTypeCall() *TL_auth_codeTypeCall {
	return &TL_auth_codeTypeCall{}
}

func (t *TL_auth_codeTypeCall) Encode() []byte {
	return nil
}

func (t *TL_auth_codeTypeCall) Decode(b []byte) {}

//====auth_codeTypeFlashCall#226ccefb====

type TL_auth_codeTypeFlashCall struct {
}

func New_TL_auth_codeTypeFlashCall() *TL_auth_codeTypeFlashCall {
	return &TL_auth_codeTypeFlashCall{}
}

func (t *TL_auth_codeTypeFlashCall) Encode() []byte {
	return nil
}

func (t *TL_auth_codeTypeFlashCall) Decode(b []byte) {}

//====auth_sentCodeTypeApp#3dbb5986====

type TL_auth_sentCodeTypeApp struct {
	_length []byte
}

func New_TL_auth_sentCodeTypeApp() *TL_auth_sentCodeTypeApp {
	return &TL_auth_sentCodeTypeApp{}
}

func (t *TL_auth_sentCodeTypeApp) Encode() []byte {
	return nil
}

func (t *TL_auth_sentCodeTypeApp) Decode(b []byte) {}

//====auth_sentCodeTypeSms#c000bba2====

type TL_auth_sentCodeTypeSms struct {
	_length []byte
}

func New_TL_auth_sentCodeTypeSms() *TL_auth_sentCodeTypeSms {
	return &TL_auth_sentCodeTypeSms{}
}

func (t *TL_auth_sentCodeTypeSms) Encode() []byte {
	return nil
}

func (t *TL_auth_sentCodeTypeSms) Decode(b []byte) {}

//====auth_sentCodeTypeCall#5353e5a7====

type TL_auth_sentCodeTypeCall struct {
	_length []byte
}

func New_TL_auth_sentCodeTypeCall() *TL_auth_sentCodeTypeCall {
	return &TL_auth_sentCodeTypeCall{}
}

func (t *TL_auth_sentCodeTypeCall) Encode() []byte {
	return nil
}

func (t *TL_auth_sentCodeTypeCall) Decode(b []byte) {}

//====auth_sentCodeTypeFlashCall#ab03c6d9====

type TL_auth_sentCodeTypeFlashCall struct {
	_pattern string
}

func New_TL_auth_sentCodeTypeFlashCall() *TL_auth_sentCodeTypeFlashCall {
	return &TL_auth_sentCodeTypeFlashCall{}
}

func (t *TL_auth_sentCodeTypeFlashCall) Encode() []byte {
	return nil
}

func (t *TL_auth_sentCodeTypeFlashCall) Decode(b []byte) {}

//====messages_botCallbackAnswer#36585ea4====

type TL_messages_botCallbackAnswer struct {
	_flags      []byte
	_alert      []byte
	_has_url    []byte
	_native_ui  []byte
	_message    []byte
	_url        []byte
	_cache_time []byte
}

func New_TL_messages_botCallbackAnswer() *TL_messages_botCallbackAnswer {
	return &TL_messages_botCallbackAnswer{}
}

func (t *TL_messages_botCallbackAnswer) Encode() []byte {
	return nil
}

func (t *TL_messages_botCallbackAnswer) Decode(b []byte) {}

//====messages_messageEditData#26b5dde6====

type TL_messages_messageEditData struct {
	_flags   []byte
	_caption []byte
}

func New_TL_messages_messageEditData() *TL_messages_messageEditData {
	return &TL_messages_messageEditData{}
}

func (t *TL_messages_messageEditData) Encode() []byte {
	return nil
}

func (t *TL_messages_messageEditData) Decode(b []byte) {}

//====inputBotInlineMessageID#890c3d89====

type TL_inputBotInlineMessageID struct {
	_dc_id       []byte
	_id          int64
	_access_hash int64
}

func New_TL_inputBotInlineMessageID() *TL_inputBotInlineMessageID {
	return &TL_inputBotInlineMessageID{}
}

func (t *TL_inputBotInlineMessageID) Encode() []byte {
	return nil
}

func (t *TL_inputBotInlineMessageID) Decode(b []byte) {}

//====inlineBotSwitchPM#3c20629f====

type TL_inlineBotSwitchPM struct {
	_text        string
	_start_param string
}

func New_TL_inlineBotSwitchPM() *TL_inlineBotSwitchPM {
	return &TL_inlineBotSwitchPM{}
}

func (t *TL_inlineBotSwitchPM) Encode() []byte {
	return nil
}

func (t *TL_inlineBotSwitchPM) Decode(b []byte) {}

//====messages_peerDialogs#3371c354====

type TL_messages_peerDialogs struct {
	_dialogs  []byte
	_messages []byte
	_chats    []byte
	_users    []byte
	_state    []byte
}

func New_TL_messages_peerDialogs() *TL_messages_peerDialogs {
	return &TL_messages_peerDialogs{}
}

func (t *TL_messages_peerDialogs) Encode() []byte {
	return nil
}

func (t *TL_messages_peerDialogs) Decode(b []byte) {}

//====topPeer#edcdc05b====

type TL_topPeer struct {
	_peer   []byte
	_rating []byte
}

func New_TL_topPeer() *TL_topPeer {
	return &TL_topPeer{}
}

func (t *TL_topPeer) Encode() []byte {
	return nil
}

func (t *TL_topPeer) Decode(b []byte) {}

//====topPeerCategoryBotsPM#ab661b5b====

type TL_topPeerCategoryBotsPM struct {
}

func New_TL_topPeerCategoryBotsPM() *TL_topPeerCategoryBotsPM {
	return &TL_topPeerCategoryBotsPM{}
}

func (t *TL_topPeerCategoryBotsPM) Encode() []byte {
	return nil
}

func (t *TL_topPeerCategoryBotsPM) Decode(b []byte) {}

//====topPeerCategoryBotsInline#148677e2====

type TL_topPeerCategoryBotsInline struct {
}

func New_TL_topPeerCategoryBotsInline() *TL_topPeerCategoryBotsInline {
	return &TL_topPeerCategoryBotsInline{}
}

func (t *TL_topPeerCategoryBotsInline) Encode() []byte {
	return nil
}

func (t *TL_topPeerCategoryBotsInline) Decode(b []byte) {}

//====topPeerCategoryCorrespondents#0637b7ed====

type TL_topPeerCategoryCorrespondents struct {
}

func New_TL_topPeerCategoryCorrespondents() *TL_topPeerCategoryCorrespondents {
	return &TL_topPeerCategoryCorrespondents{}
}

func (t *TL_topPeerCategoryCorrespondents) Encode() []byte {
	return nil
}

func (t *TL_topPeerCategoryCorrespondents) Decode(b []byte) {}

//====topPeerCategoryGroups#bd17a14a====

type TL_topPeerCategoryGroups struct {
}

func New_TL_topPeerCategoryGroups() *TL_topPeerCategoryGroups {
	return &TL_topPeerCategoryGroups{}
}

func (t *TL_topPeerCategoryGroups) Encode() []byte {
	return nil
}

func (t *TL_topPeerCategoryGroups) Decode(b []byte) {}

//====topPeerCategoryChannels#161d9628====

type TL_topPeerCategoryChannels struct {
}

func New_TL_topPeerCategoryChannels() *TL_topPeerCategoryChannels {
	return &TL_topPeerCategoryChannels{}
}

func (t *TL_topPeerCategoryChannels) Encode() []byte {
	return nil
}

func (t *TL_topPeerCategoryChannels) Decode(b []byte) {}

//====topPeerCategoryPhoneCalls#1e76a78c====

type TL_topPeerCategoryPhoneCalls struct {
}

func New_TL_topPeerCategoryPhoneCalls() *TL_topPeerCategoryPhoneCalls {
	return &TL_topPeerCategoryPhoneCalls{}
}

func (t *TL_topPeerCategoryPhoneCalls) Encode() []byte {
	return nil
}

func (t *TL_topPeerCategoryPhoneCalls) Decode(b []byte) {}

//====topPeerCategoryPeers#fb834291====

type TL_topPeerCategoryPeers struct {
	_category []byte
	_count    []byte
	_peers    []byte
}

func New_TL_topPeerCategoryPeers() *TL_topPeerCategoryPeers {
	return &TL_topPeerCategoryPeers{}
}

func (t *TL_topPeerCategoryPeers) Encode() []byte {
	return nil
}

func (t *TL_topPeerCategoryPeers) Decode(b []byte) {}

//====contacts_topPeersNotModified#de266ef5====

type TL_contacts_topPeersNotModified struct {
}

func New_TL_contacts_topPeersNotModified() *TL_contacts_topPeersNotModified {
	return &TL_contacts_topPeersNotModified{}
}

func (t *TL_contacts_topPeersNotModified) Encode() []byte {
	return nil
}

func (t *TL_contacts_topPeersNotModified) Decode(b []byte) {}

//====contacts_topPeers#70b772a8====

type TL_contacts_topPeers struct {
	_categories []byte
	_chats      []byte
	_users      []byte
}

func New_TL_contacts_topPeers() *TL_contacts_topPeers {
	return &TL_contacts_topPeers{}
}

func (t *TL_contacts_topPeers) Encode() []byte {
	return nil
}

func (t *TL_contacts_topPeers) Decode(b []byte) {}

//====draftMessageEmpty#ba4baec5====

type TL_draftMessageEmpty struct {
}

func New_TL_draftMessageEmpty() *TL_draftMessageEmpty {
	return &TL_draftMessageEmpty{}
}

func (t *TL_draftMessageEmpty) Encode() []byte {
	return nil
}

func (t *TL_draftMessageEmpty) Decode(b []byte) {}

//====draftMessage#fd8e711f====

type TL_draftMessage struct {
	_flags           []byte
	_no_webpage      []byte
	_reply_to_msg_id []byte
	_message         string
	_entities        []byte
	_date            []byte
}

func New_TL_draftMessage() *TL_draftMessage {
	return &TL_draftMessage{}
}

func (t *TL_draftMessage) Encode() []byte {
	return nil
}

func (t *TL_draftMessage) Decode(b []byte) {}

//====messages_featuredStickersNotModified#04ede3cf====

type TL_messages_featuredStickersNotModified struct {
}

func New_TL_messages_featuredStickersNotModified() *TL_messages_featuredStickersNotModified {
	return &TL_messages_featuredStickersNotModified{}
}

func (t *TL_messages_featuredStickersNotModified) Encode() []byte {
	return nil
}

func (t *TL_messages_featuredStickersNotModified) Decode(b []byte) {}

//====messages_featuredStickers#f89d88e5====

type TL_messages_featuredStickers struct {
	_hash   []byte
	_sets   []byte
	_unread []byte
}

func New_TL_messages_featuredStickers() *TL_messages_featuredStickers {
	return &TL_messages_featuredStickers{}
}

func (t *TL_messages_featuredStickers) Encode() []byte {
	return nil
}

func (t *TL_messages_featuredStickers) Decode(b []byte) {}

//====messages_recentStickersNotModified#0b17f890====

type TL_messages_recentStickersNotModified struct {
}

func New_TL_messages_recentStickersNotModified() *TL_messages_recentStickersNotModified {
	return &TL_messages_recentStickersNotModified{}
}

func (t *TL_messages_recentStickersNotModified) Encode() []byte {
	return nil
}

func (t *TL_messages_recentStickersNotModified) Decode(b []byte) {}

//====messages_recentStickers#5ce20970====

type TL_messages_recentStickers struct {
	_hash     []byte
	_stickers []byte
}

func New_TL_messages_recentStickers() *TL_messages_recentStickers {
	return &TL_messages_recentStickers{}
}

func (t *TL_messages_recentStickers) Encode() []byte {
	return nil
}

func (t *TL_messages_recentStickers) Decode(b []byte) {}

//====messages_archivedStickers#4fcba9c8====

type TL_messages_archivedStickers struct {
	_count []byte
	_sets  []byte
}

func New_TL_messages_archivedStickers() *TL_messages_archivedStickers {
	return &TL_messages_archivedStickers{}
}

func (t *TL_messages_archivedStickers) Encode() []byte {
	return nil
}

func (t *TL_messages_archivedStickers) Decode(b []byte) {}

//====messages_stickerSetInstallResultSuccess#38641628====

type TL_messages_stickerSetInstallResultSuccess struct {
}

func New_TL_messages_stickerSetInstallResultSuccess() *TL_messages_stickerSetInstallResultSuccess {
	return &TL_messages_stickerSetInstallResultSuccess{}
}

func (t *TL_messages_stickerSetInstallResultSuccess) Encode() []byte {
	return nil
}

func (t *TL_messages_stickerSetInstallResultSuccess) Decode(b []byte) {}

//====messages_stickerSetInstallResultArchive#35e410a8====

type TL_messages_stickerSetInstallResultArchive struct {
	_sets []byte
}

func New_TL_messages_stickerSetInstallResultArchive() *TL_messages_stickerSetInstallResultArchive {
	return &TL_messages_stickerSetInstallResultArchive{}
}

func (t *TL_messages_stickerSetInstallResultArchive) Encode() []byte {
	return nil
}

func (t *TL_messages_stickerSetInstallResultArchive) Decode(b []byte) {}

//====stickerSetCovered#6410a5d2====

type TL_stickerSetCovered struct {
	_set   []byte
	_cover []byte
}

func New_TL_stickerSetCovered() *TL_stickerSetCovered {
	return &TL_stickerSetCovered{}
}

func (t *TL_stickerSetCovered) Encode() []byte {
	return nil
}

func (t *TL_stickerSetCovered) Decode(b []byte) {}

//====stickerSetMultiCovered#3407e51b====

type TL_stickerSetMultiCovered struct {
	_set    []byte
	_covers []byte
}

func New_TL_stickerSetMultiCovered() *TL_stickerSetMultiCovered {
	return &TL_stickerSetMultiCovered{}
}

func (t *TL_stickerSetMultiCovered) Encode() []byte {
	return nil
}

func (t *TL_stickerSetMultiCovered) Decode(b []byte) {}

//====maskCoords#aed6dbb2====

type TL_maskCoords struct {
	_n    []byte
	_x    []byte
	_y    []byte
	_zoom []byte
}

func New_TL_maskCoords() *TL_maskCoords {
	return &TL_maskCoords{}
}

func (t *TL_maskCoords) Encode() []byte {
	return nil
}

func (t *TL_maskCoords) Decode(b []byte) {}

//====inputStickeredMediaPhoto#4a992157====

type TL_inputStickeredMediaPhoto struct {
	_id []byte
}

func New_TL_inputStickeredMediaPhoto() *TL_inputStickeredMediaPhoto {
	return &TL_inputStickeredMediaPhoto{}
}

func (t *TL_inputStickeredMediaPhoto) Encode() []byte {
	return nil
}

func (t *TL_inputStickeredMediaPhoto) Decode(b []byte) {}

//====inputStickeredMediaDocument#0438865b====

type TL_inputStickeredMediaDocument struct {
	_id []byte
}

func New_TL_inputStickeredMediaDocument() *TL_inputStickeredMediaDocument {
	return &TL_inputStickeredMediaDocument{}
}

func (t *TL_inputStickeredMediaDocument) Encode() []byte {
	return nil
}

func (t *TL_inputStickeredMediaDocument) Decode(b []byte) {}

//====game#bdf9653b====

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

func New_TL_game() *TL_game {
	return &TL_game{}
}

func (t *TL_game) Encode() []byte {
	return nil
}

func (t *TL_game) Decode(b []byte) {}

//====inputGameID#032c3e77====

type TL_inputGameID struct {
	_id          int64
	_access_hash int64
}

func New_TL_inputGameID() *TL_inputGameID {
	return &TL_inputGameID{}
}

func (t *TL_inputGameID) Encode() []byte {
	return nil
}

func (t *TL_inputGameID) Decode(b []byte) {}

//====inputGameShortName#c331e80a====

type TL_inputGameShortName struct {
	_bot_id     []byte
	_short_name string
}

func New_TL_inputGameShortName() *TL_inputGameShortName {
	return &TL_inputGameShortName{}
}

func (t *TL_inputGameShortName) Encode() []byte {
	return nil
}

func (t *TL_inputGameShortName) Decode(b []byte) {}

//====highScore#58fffcd0====

type TL_highScore struct {
	_pos     []byte
	_user_id []byte
	_score   []byte
}

func New_TL_highScore() *TL_highScore {
	return &TL_highScore{}
}

func (t *TL_highScore) Encode() []byte {
	return nil
}

func (t *TL_highScore) Decode(b []byte) {}

//====messages_highScores#9a3bfd99====

type TL_messages_highScores struct {
	_scores []byte
	_users  []byte
}

func New_TL_messages_highScores() *TL_messages_highScores {
	return &TL_messages_highScores{}
}

func (t *TL_messages_highScores) Encode() []byte {
	return nil
}

func (t *TL_messages_highScores) Decode(b []byte) {}

//====textEmpty#dc3d824f====

type TL_textEmpty struct {
}

func New_TL_textEmpty() *TL_textEmpty {
	return &TL_textEmpty{}
}

func (t *TL_textEmpty) Encode() []byte {
	return nil
}

func (t *TL_textEmpty) Decode(b []byte) {}

//====textPlain#744694e0====

type TL_textPlain struct {
	_text string
}

func New_TL_textPlain() *TL_textPlain {
	return &TL_textPlain{}
}

func (t *TL_textPlain) Encode() []byte {
	return nil
}

func (t *TL_textPlain) Decode(b []byte) {}

//====textBold#6724abc4====

type TL_textBold struct {
	_text []byte
}

func New_TL_textBold() *TL_textBold {
	return &TL_textBold{}
}

func (t *TL_textBold) Encode() []byte {
	return nil
}

func (t *TL_textBold) Decode(b []byte) {}

//====textItalic#d912a59c====

type TL_textItalic struct {
	_text []byte
}

func New_TL_textItalic() *TL_textItalic {
	return &TL_textItalic{}
}

func (t *TL_textItalic) Encode() []byte {
	return nil
}

func (t *TL_textItalic) Decode(b []byte) {}

//====textUnderline#c12622c4====

type TL_textUnderline struct {
	_text []byte
}

func New_TL_textUnderline() *TL_textUnderline {
	return &TL_textUnderline{}
}

func (t *TL_textUnderline) Encode() []byte {
	return nil
}

func (t *TL_textUnderline) Decode(b []byte) {}

//====textStrike#9bf8bb95====

type TL_textStrike struct {
	_text []byte
}

func New_TL_textStrike() *TL_textStrike {
	return &TL_textStrike{}
}

func (t *TL_textStrike) Encode() []byte {
	return nil
}

func (t *TL_textStrike) Decode(b []byte) {}

//====textFixed#6c3f19b9====

type TL_textFixed struct {
	_text []byte
}

func New_TL_textFixed() *TL_textFixed {
	return &TL_textFixed{}
}

func (t *TL_textFixed) Encode() []byte {
	return nil
}

func (t *TL_textFixed) Decode(b []byte) {}

//====textUrl#3c2884c1====

type TL_textUrl struct {
	_text       []byte
	_url        string
	_webpage_id int64
}

func New_TL_textUrl() *TL_textUrl {
	return &TL_textUrl{}
}

func (t *TL_textUrl) Encode() []byte {
	return nil
}

func (t *TL_textUrl) Decode(b []byte) {}

//====textEmail#de5a0dd6====

type TL_textEmail struct {
	_text  []byte
	_email string
}

func New_TL_textEmail() *TL_textEmail {
	return &TL_textEmail{}
}

func (t *TL_textEmail) Encode() []byte {
	return nil
}

func (t *TL_textEmail) Decode(b []byte) {}

//====textConcat#7e6260d7====

type TL_textConcat struct {
	_texts []byte
}

func New_TL_textConcat() *TL_textConcat {
	return &TL_textConcat{}
}

func (t *TL_textConcat) Encode() []byte {
	return nil
}

func (t *TL_textConcat) Decode(b []byte) {}

//====pageBlockUnsupported#13567e8a====

type TL_pageBlockUnsupported struct {
}

func New_TL_pageBlockUnsupported() *TL_pageBlockUnsupported {
	return &TL_pageBlockUnsupported{}
}

func (t *TL_pageBlockUnsupported) Encode() []byte {
	return nil
}

func (t *TL_pageBlockUnsupported) Decode(b []byte) {}

//====pageBlockTitle#70abc3fd====

type TL_pageBlockTitle struct {
	_text []byte
}

func New_TL_pageBlockTitle() *TL_pageBlockTitle {
	return &TL_pageBlockTitle{}
}

func (t *TL_pageBlockTitle) Encode() []byte {
	return nil
}

func (t *TL_pageBlockTitle) Decode(b []byte) {}

//====pageBlockSubtitle#8ffa9a1f====

type TL_pageBlockSubtitle struct {
	_text []byte
}

func New_TL_pageBlockSubtitle() *TL_pageBlockSubtitle {
	return &TL_pageBlockSubtitle{}
}

func (t *TL_pageBlockSubtitle) Encode() []byte {
	return nil
}

func (t *TL_pageBlockSubtitle) Decode(b []byte) {}

//====pageBlockAuthorDate#baafe5e0====

type TL_pageBlockAuthorDate struct {
	_author         []byte
	_published_date []byte
}

func New_TL_pageBlockAuthorDate() *TL_pageBlockAuthorDate {
	return &TL_pageBlockAuthorDate{}
}

func (t *TL_pageBlockAuthorDate) Encode() []byte {
	return nil
}

func (t *TL_pageBlockAuthorDate) Decode(b []byte) {}

//====pageBlockHeader#bfd064ec====

type TL_pageBlockHeader struct {
	_text []byte
}

func New_TL_pageBlockHeader() *TL_pageBlockHeader {
	return &TL_pageBlockHeader{}
}

func (t *TL_pageBlockHeader) Encode() []byte {
	return nil
}

func (t *TL_pageBlockHeader) Decode(b []byte) {}

//====pageBlockSubheader#f12bb6e1====

type TL_pageBlockSubheader struct {
	_text []byte
}

func New_TL_pageBlockSubheader() *TL_pageBlockSubheader {
	return &TL_pageBlockSubheader{}
}

func (t *TL_pageBlockSubheader) Encode() []byte {
	return nil
}

func (t *TL_pageBlockSubheader) Decode(b []byte) {}

//====pageBlockParagraph#467a0766====

type TL_pageBlockParagraph struct {
	_text []byte
}

func New_TL_pageBlockParagraph() *TL_pageBlockParagraph {
	return &TL_pageBlockParagraph{}
}

func (t *TL_pageBlockParagraph) Encode() []byte {
	return nil
}

func (t *TL_pageBlockParagraph) Decode(b []byte) {}

//====pageBlockPreformatted#c070d93e====

type TL_pageBlockPreformatted struct {
	_text     []byte
	_language string
}

func New_TL_pageBlockPreformatted() *TL_pageBlockPreformatted {
	return &TL_pageBlockPreformatted{}
}

func (t *TL_pageBlockPreformatted) Encode() []byte {
	return nil
}

func (t *TL_pageBlockPreformatted) Decode(b []byte) {}

//====pageBlockFooter#48870999====

type TL_pageBlockFooter struct {
	_text []byte
}

func New_TL_pageBlockFooter() *TL_pageBlockFooter {
	return &TL_pageBlockFooter{}
}

func (t *TL_pageBlockFooter) Encode() []byte {
	return nil
}

func (t *TL_pageBlockFooter) Decode(b []byte) {}

//====pageBlockDivider#db20b188====

type TL_pageBlockDivider struct {
}

func New_TL_pageBlockDivider() *TL_pageBlockDivider {
	return &TL_pageBlockDivider{}
}

func (t *TL_pageBlockDivider) Encode() []byte {
	return nil
}

func (t *TL_pageBlockDivider) Decode(b []byte) {}

//====pageBlockAnchor#ce0d37b0====

type TL_pageBlockAnchor struct {
	_name string
}

func New_TL_pageBlockAnchor() *TL_pageBlockAnchor {
	return &TL_pageBlockAnchor{}
}

func (t *TL_pageBlockAnchor) Encode() []byte {
	return nil
}

func (t *TL_pageBlockAnchor) Decode(b []byte) {}

//====pageBlockList#3a58c7f4====

type TL_pageBlockList struct {
	_ordered []byte
	_items   []byte
}

func New_TL_pageBlockList() *TL_pageBlockList {
	return &TL_pageBlockList{}
}

func (t *TL_pageBlockList) Encode() []byte {
	return nil
}

func (t *TL_pageBlockList) Decode(b []byte) {}

//====pageBlockBlockquote#263d7c26====

type TL_pageBlockBlockquote struct {
	_text    []byte
	_caption []byte
}

func New_TL_pageBlockBlockquote() *TL_pageBlockBlockquote {
	return &TL_pageBlockBlockquote{}
}

func (t *TL_pageBlockBlockquote) Encode() []byte {
	return nil
}

func (t *TL_pageBlockBlockquote) Decode(b []byte) {}

//====pageBlockPullquote#4f4456d3====

type TL_pageBlockPullquote struct {
	_text    []byte
	_caption []byte
}

func New_TL_pageBlockPullquote() *TL_pageBlockPullquote {
	return &TL_pageBlockPullquote{}
}

func (t *TL_pageBlockPullquote) Encode() []byte {
	return nil
}

func (t *TL_pageBlockPullquote) Decode(b []byte) {}

//====pageBlockPhoto#e9c69982====

type TL_pageBlockPhoto struct {
	_photo_id int64
	_caption  []byte
}

func New_TL_pageBlockPhoto() *TL_pageBlockPhoto {
	return &TL_pageBlockPhoto{}
}

func (t *TL_pageBlockPhoto) Encode() []byte {
	return nil
}

func (t *TL_pageBlockPhoto) Decode(b []byte) {}

//====pageBlockVideo#d9d71866====

type TL_pageBlockVideo struct {
	_flags    []byte
	_autoplay []byte
	_loop     []byte
	_video_id int64
	_caption  []byte
}

func New_TL_pageBlockVideo() *TL_pageBlockVideo {
	return &TL_pageBlockVideo{}
}

func (t *TL_pageBlockVideo) Encode() []byte {
	return nil
}

func (t *TL_pageBlockVideo) Decode(b []byte) {}

//====pageBlockCover#39f23300====

type TL_pageBlockCover struct {
	_cover []byte
}

func New_TL_pageBlockCover() *TL_pageBlockCover {
	return &TL_pageBlockCover{}
}

func (t *TL_pageBlockCover) Encode() []byte {
	return nil
}

func (t *TL_pageBlockCover) Decode(b []byte) {}

//====pageBlockEmbed#cde200d1====

type TL_pageBlockEmbed struct {
	_flags           []byte
	_full_width      []byte
	_allow_scrolling []byte
	_url             []byte
	_html            []byte
	_poster_photo_id []byte
	_w               []byte
	_h               []byte
	_caption         []byte
}

func New_TL_pageBlockEmbed() *TL_pageBlockEmbed {
	return &TL_pageBlockEmbed{}
}

func (t *TL_pageBlockEmbed) Encode() []byte {
	return nil
}

func (t *TL_pageBlockEmbed) Decode(b []byte) {}

//====pageBlockEmbedPost#292c7be9====

type TL_pageBlockEmbedPost struct {
	_url             string
	_webpage_id      int64
	_author_photo_id int64
	_author          string
	_date            []byte
	_blocks          []byte
	_caption         []byte
}

func New_TL_pageBlockEmbedPost() *TL_pageBlockEmbedPost {
	return &TL_pageBlockEmbedPost{}
}

func (t *TL_pageBlockEmbedPost) Encode() []byte {
	return nil
}

func (t *TL_pageBlockEmbedPost) Decode(b []byte) {}

//====pageBlockCollage#08b31c4f====

type TL_pageBlockCollage struct {
	_items   []byte
	_caption []byte
}

func New_TL_pageBlockCollage() *TL_pageBlockCollage {
	return &TL_pageBlockCollage{}
}

func (t *TL_pageBlockCollage) Encode() []byte {
	return nil
}

func (t *TL_pageBlockCollage) Decode(b []byte) {}

//====pageBlockSlideshow#130c8963====

type TL_pageBlockSlideshow struct {
	_items   []byte
	_caption []byte
}

func New_TL_pageBlockSlideshow() *TL_pageBlockSlideshow {
	return &TL_pageBlockSlideshow{}
}

func (t *TL_pageBlockSlideshow) Encode() []byte {
	return nil
}

func (t *TL_pageBlockSlideshow) Decode(b []byte) {}

//====pageBlockChannel#ef1751b5====

type TL_pageBlockChannel struct {
	_channel []byte
}

func New_TL_pageBlockChannel() *TL_pageBlockChannel {
	return &TL_pageBlockChannel{}
}

func (t *TL_pageBlockChannel) Encode() []byte {
	return nil
}

func (t *TL_pageBlockChannel) Decode(b []byte) {}

//====pageBlockAudio#31b81a7f====

type TL_pageBlockAudio struct {
	_audio_id int64
	_caption  []byte
}

func New_TL_pageBlockAudio() *TL_pageBlockAudio {
	return &TL_pageBlockAudio{}
}

func (t *TL_pageBlockAudio) Encode() []byte {
	return nil
}

func (t *TL_pageBlockAudio) Decode(b []byte) {}

//====pagePart#8e3f9ebe====

type TL_pagePart struct {
	_blocks    []byte
	_photos    []byte
	_documents []byte
}

func New_TL_pagePart() *TL_pagePart {
	return &TL_pagePart{}
}

func (t *TL_pagePart) Encode() []byte {
	return nil
}

func (t *TL_pagePart) Decode(b []byte) {}

//====pageFull#556ec7aa====

type TL_pageFull struct {
	_blocks    []byte
	_photos    []byte
	_documents []byte
}

func New_TL_pageFull() *TL_pageFull {
	return &TL_pageFull{}
}

func (t *TL_pageFull) Encode() []byte {
	return nil
}

func (t *TL_pageFull) Decode(b []byte) {}

//====phoneCallDiscardReasonMissed#85e42301====

type TL_phoneCallDiscardReasonMissed struct {
}

func New_TL_phoneCallDiscardReasonMissed() *TL_phoneCallDiscardReasonMissed {
	return &TL_phoneCallDiscardReasonMissed{}
}

func (t *TL_phoneCallDiscardReasonMissed) Encode() []byte {
	return nil
}

func (t *TL_phoneCallDiscardReasonMissed) Decode(b []byte) {}

//====phoneCallDiscardReasonDisconnect#e095c1a0====

type TL_phoneCallDiscardReasonDisconnect struct {
}

func New_TL_phoneCallDiscardReasonDisconnect() *TL_phoneCallDiscardReasonDisconnect {
	return &TL_phoneCallDiscardReasonDisconnect{}
}

func (t *TL_phoneCallDiscardReasonDisconnect) Encode() []byte {
	return nil
}

func (t *TL_phoneCallDiscardReasonDisconnect) Decode(b []byte) {}

//====phoneCallDiscardReasonHangup#57adc690====

type TL_phoneCallDiscardReasonHangup struct {
}

func New_TL_phoneCallDiscardReasonHangup() *TL_phoneCallDiscardReasonHangup {
	return &TL_phoneCallDiscardReasonHangup{}
}

func (t *TL_phoneCallDiscardReasonHangup) Encode() []byte {
	return nil
}

func (t *TL_phoneCallDiscardReasonHangup) Decode(b []byte) {}

//====phoneCallDiscardReasonBusy#faf7e8c9====

type TL_phoneCallDiscardReasonBusy struct {
}

func New_TL_phoneCallDiscardReasonBusy() *TL_phoneCallDiscardReasonBusy {
	return &TL_phoneCallDiscardReasonBusy{}
}

func (t *TL_phoneCallDiscardReasonBusy) Encode() []byte {
	return nil
}

func (t *TL_phoneCallDiscardReasonBusy) Decode(b []byte) {}

//====dataJSON#7d748d04====

type TL_dataJSON struct {
	_data string
}

func New_TL_dataJSON() *TL_dataJSON {
	return &TL_dataJSON{}
}

func (t *TL_dataJSON) Encode() []byte {
	return nil
}

func (t *TL_dataJSON) Decode(b []byte) {}

//====labeledPrice#cb296bf8====

type TL_labeledPrice struct {
	_label  string
	_amount int64
}

func New_TL_labeledPrice() *TL_labeledPrice {
	return &TL_labeledPrice{}
}

func (t *TL_labeledPrice) Encode() []byte {
	return nil
}

func (t *TL_labeledPrice) Decode(b []byte) {}

//====invoice#c30aa358====

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

func New_TL_invoice() *TL_invoice {
	return &TL_invoice{}
}

func (t *TL_invoice) Encode() []byte {
	return nil
}

func (t *TL_invoice) Decode(b []byte) {}

//====paymentCharge#ea02c27e====

type TL_paymentCharge struct {
	_id                 string
	_provider_charge_id string
}

func New_TL_paymentCharge() *TL_paymentCharge {
	return &TL_paymentCharge{}
}

func (t *TL_paymentCharge) Encode() []byte {
	return nil
}

func (t *TL_paymentCharge) Decode(b []byte) {}

//====postAddress#1e8caaeb====

type TL_postAddress struct {
	_street_line1 string
	_street_line2 string
	_city         string
	_state        string
	_country_iso2 string
	_post_code    string
}

func New_TL_postAddress() *TL_postAddress {
	return &TL_postAddress{}
}

func (t *TL_postAddress) Encode() []byte {
	return nil
}

func (t *TL_postAddress) Decode(b []byte) {}

//====paymentRequestedInfo#909c3f94====

type TL_paymentRequestedInfo struct {
	_flags            []byte
	_name             []byte
	_phone            []byte
	_email            []byte
	_shipping_address []byte
}

func New_TL_paymentRequestedInfo() *TL_paymentRequestedInfo {
	return &TL_paymentRequestedInfo{}
}

func (t *TL_paymentRequestedInfo) Encode() []byte {
	return nil
}

func (t *TL_paymentRequestedInfo) Decode(b []byte) {}

//====paymentSavedCredentialsCard#cdc27a1f====

type TL_paymentSavedCredentialsCard struct {
	_id    string
	_title string
}

func New_TL_paymentSavedCredentialsCard() *TL_paymentSavedCredentialsCard {
	return &TL_paymentSavedCredentialsCard{}
}

func (t *TL_paymentSavedCredentialsCard) Encode() []byte {
	return nil
}

func (t *TL_paymentSavedCredentialsCard) Decode(b []byte) {}

//====webDocument#c61acbd8====

type TL_webDocument struct {
	_url         string
	_access_hash int64
	_size        []byte
	_mime_type   string
	_attributes  []byte
	_dc_id       []byte
}

func New_TL_webDocument() *TL_webDocument {
	return &TL_webDocument{}
}

func (t *TL_webDocument) Encode() []byte {
	return nil
}

func (t *TL_webDocument) Decode(b []byte) {}

//====inputWebDocument#9bed434d====

type TL_inputWebDocument struct {
	_url        string
	_size       []byte
	_mime_type  string
	_attributes []byte
}

func New_TL_inputWebDocument() *TL_inputWebDocument {
	return &TL_inputWebDocument{}
}

func (t *TL_inputWebDocument) Encode() []byte {
	return nil
}

func (t *TL_inputWebDocument) Decode(b []byte) {}

//====inputWebFileLocation#c239d686====

type TL_inputWebFileLocation struct {
	_url         string
	_access_hash int64
}

func New_TL_inputWebFileLocation() *TL_inputWebFileLocation {
	return &TL_inputWebFileLocation{}
}

func (t *TL_inputWebFileLocation) Encode() []byte {
	return nil
}

func (t *TL_inputWebFileLocation) Decode(b []byte) {}

//====upload_webFile#21e753bc====

type TL_upload_webFile struct {
	_size      []byte
	_mime_type string
	_file_type []byte
	_mtime     []byte
	_bytes     []byte
}

func New_TL_upload_webFile() *TL_upload_webFile {
	return &TL_upload_webFile{}
}

func (t *TL_upload_webFile) Encode() []byte {
	return nil
}

func (t *TL_upload_webFile) Decode(b []byte) {}

//====payments_paymentForm#3f56aea3====

type TL_payments_paymentForm struct {
	_flags                []byte
	_can_save_credentials []byte
	_password_missing     []byte
	_bot_id               []byte
	_invoice              []byte
	_provider_id          []byte
	_url                  string
	_native_provider      []byte
	_native_params        []byte
	_saved_info           []byte
	_saved_credentials    []byte
	_users                []byte
}

func New_TL_payments_paymentForm() *TL_payments_paymentForm {
	return &TL_payments_paymentForm{}
}

func (t *TL_payments_paymentForm) Encode() []byte {
	return nil
}

func (t *TL_payments_paymentForm) Decode(b []byte) {}

//====payments_validatedRequestedInfo#d1451883====

type TL_payments_validatedRequestedInfo struct {
	_flags            []byte
	_id               []byte
	_shipping_options []byte
}

func New_TL_payments_validatedRequestedInfo() *TL_payments_validatedRequestedInfo {
	return &TL_payments_validatedRequestedInfo{}
}

func (t *TL_payments_validatedRequestedInfo) Encode() []byte {
	return nil
}

func (t *TL_payments_validatedRequestedInfo) Decode(b []byte) {}

//====payments_paymentResult#4e5f810d====

type TL_payments_paymentResult struct {
	_updates []byte
}

func New_TL_payments_paymentResult() *TL_payments_paymentResult {
	return &TL_payments_paymentResult{}
}

func (t *TL_payments_paymentResult) Encode() []byte {
	return nil
}

func (t *TL_payments_paymentResult) Decode(b []byte) {}

//====payments_paymentVerficationNeeded#6b56b921====

type TL_payments_paymentVerficationNeeded struct {
	_url string
}

func New_TL_payments_paymentVerficationNeeded() *TL_payments_paymentVerficationNeeded {
	return &TL_payments_paymentVerficationNeeded{}
}

func (t *TL_payments_paymentVerficationNeeded) Encode() []byte {
	return nil
}

func (t *TL_payments_paymentVerficationNeeded) Decode(b []byte) {}

//====payments_paymentReceipt#500911e1====

type TL_payments_paymentReceipt struct {
	_flags             []byte
	_date              []byte
	_bot_id            []byte
	_invoice           []byte
	_provider_id       []byte
	_info              []byte
	_shipping          []byte
	_currency          string
	_total_amount      int64
	_credentials_title string
	_users             []byte
}

func New_TL_payments_paymentReceipt() *TL_payments_paymentReceipt {
	return &TL_payments_paymentReceipt{}
}

func (t *TL_payments_paymentReceipt) Encode() []byte {
	return nil
}

func (t *TL_payments_paymentReceipt) Decode(b []byte) {}

//====payments_savedInfo#fb8fe43c====

type TL_payments_savedInfo struct {
	_flags                 []byte
	_has_saved_credentials []byte
	_saved_info            []byte
}

func New_TL_payments_savedInfo() *TL_payments_savedInfo {
	return &TL_payments_savedInfo{}
}

func (t *TL_payments_savedInfo) Encode() []byte {
	return nil
}

func (t *TL_payments_savedInfo) Decode(b []byte) {}

//====inputPaymentCredentialsSaved#c10eb2cf====

type TL_inputPaymentCredentialsSaved struct {
	_id           string
	_tmp_password []byte
}

func New_TL_inputPaymentCredentialsSaved() *TL_inputPaymentCredentialsSaved {
	return &TL_inputPaymentCredentialsSaved{}
}

func (t *TL_inputPaymentCredentialsSaved) Encode() []byte {
	return nil
}

func (t *TL_inputPaymentCredentialsSaved) Decode(b []byte) {}

//====inputPaymentCredentials#3417d728====

type TL_inputPaymentCredentials struct {
	_flags []byte
	_save  []byte
	_data  []byte
}

func New_TL_inputPaymentCredentials() *TL_inputPaymentCredentials {
	return &TL_inputPaymentCredentials{}
}

func (t *TL_inputPaymentCredentials) Encode() []byte {
	return nil
}

func (t *TL_inputPaymentCredentials) Decode(b []byte) {}

//====inputPaymentCredentialsApplePay#0aa1c39f====

type TL_inputPaymentCredentialsApplePay struct {
	_payment_data []byte
}

func New_TL_inputPaymentCredentialsApplePay() *TL_inputPaymentCredentialsApplePay {
	return &TL_inputPaymentCredentialsApplePay{}
}

func (t *TL_inputPaymentCredentialsApplePay) Encode() []byte {
	return nil
}

func (t *TL_inputPaymentCredentialsApplePay) Decode(b []byte) {}

//====inputPaymentCredentialsAndroidPay#795667a6====

type TL_inputPaymentCredentialsAndroidPay struct {
	_payment_token []byte
}

func New_TL_inputPaymentCredentialsAndroidPay() *TL_inputPaymentCredentialsAndroidPay {
	return &TL_inputPaymentCredentialsAndroidPay{}
}

func (t *TL_inputPaymentCredentialsAndroidPay) Encode() []byte {
	return nil
}

func (t *TL_inputPaymentCredentialsAndroidPay) Decode(b []byte) {}

//====account_tmpPassword#db64fd34====

type TL_account_tmpPassword struct {
	_tmp_password []byte
	_valid_until  []byte
}

func New_TL_account_tmpPassword() *TL_account_tmpPassword {
	return &TL_account_tmpPassword{}
}

func (t *TL_account_tmpPassword) Encode() []byte {
	return nil
}

func (t *TL_account_tmpPassword) Decode(b []byte) {}

//====shippingOption#b6213cdf====

type TL_shippingOption struct {
	_id     string
	_title  string
	_prices []byte
}

func New_TL_shippingOption() *TL_shippingOption {
	return &TL_shippingOption{}
}

func (t *TL_shippingOption) Encode() []byte {
	return nil
}

func (t *TL_shippingOption) Decode(b []byte) {}

//====inputStickerSetItem#ffa0a496====

type TL_inputStickerSetItem struct {
	_flags       []byte
	_document    []byte
	_emoji       string
	_mask_coords []byte
}

func New_TL_inputStickerSetItem() *TL_inputStickerSetItem {
	return &TL_inputStickerSetItem{}
}

func (t *TL_inputStickerSetItem) Encode() []byte {
	return nil
}

func (t *TL_inputStickerSetItem) Decode(b []byte) {}

//====inputPhoneCall#1e36fded====

type TL_inputPhoneCall struct {
	_id          int64
	_access_hash int64
}

func New_TL_inputPhoneCall() *TL_inputPhoneCall {
	return &TL_inputPhoneCall{}
}

func (t *TL_inputPhoneCall) Encode() []byte {
	return nil
}

func (t *TL_inputPhoneCall) Decode(b []byte) {}

//====phoneCallEmpty#5366c915====

type TL_phoneCallEmpty struct {
	_id int64
}

func New_TL_phoneCallEmpty() *TL_phoneCallEmpty {
	return &TL_phoneCallEmpty{}
}

func (t *TL_phoneCallEmpty) Encode() []byte {
	return nil
}

func (t *TL_phoneCallEmpty) Decode(b []byte) {}

//====phoneCallWaiting#1b8f4ad1====

type TL_phoneCallWaiting struct {
	_flags          []byte
	_id             int64
	_access_hash    int64
	_date           []byte
	_admin_id       []byte
	_participant_id []byte
	_protocol       []byte
	_receive_date   []byte
}

func New_TL_phoneCallWaiting() *TL_phoneCallWaiting {
	return &TL_phoneCallWaiting{}
}

func (t *TL_phoneCallWaiting) Encode() []byte {
	return nil
}

func (t *TL_phoneCallWaiting) Decode(b []byte) {}

//====phoneCallRequested#83761ce4====

type TL_phoneCallRequested struct {
	_id             int64
	_access_hash    int64
	_date           []byte
	_admin_id       []byte
	_participant_id []byte
	_g_a_hash       []byte
	_protocol       []byte
}

func New_TL_phoneCallRequested() *TL_phoneCallRequested {
	return &TL_phoneCallRequested{}
}

func (t *TL_phoneCallRequested) Encode() []byte {
	return nil
}

func (t *TL_phoneCallRequested) Decode(b []byte) {}

//====phoneCallAccepted#6d003d3f====

type TL_phoneCallAccepted struct {
	_id             int64
	_access_hash    int64
	_date           []byte
	_admin_id       []byte
	_participant_id []byte
	_g_b            []byte
	_protocol       []byte
}

func New_TL_phoneCallAccepted() *TL_phoneCallAccepted {
	return &TL_phoneCallAccepted{}
}

func (t *TL_phoneCallAccepted) Encode() []byte {
	return nil
}

func (t *TL_phoneCallAccepted) Decode(b []byte) {}

//====phoneCall#ffe6ab67====

type TL_phoneCall struct {
	_id                      int64
	_access_hash             int64
	_date                    []byte
	_admin_id                []byte
	_participant_id          []byte
	_g_a_or_b                []byte
	_key_fingerprint         int64
	_protocol                []byte
	_connection              []byte
	_alternative_connections []byte
	_start_date              []byte
}

func New_TL_phoneCall() *TL_phoneCall {
	return &TL_phoneCall{}
}

func (t *TL_phoneCall) Encode() []byte {
	return nil
}

func (t *TL_phoneCall) Decode(b []byte) {}

//====phoneCallDiscarded#50ca4de1====

type TL_phoneCallDiscarded struct {
	_flags       []byte
	_need_rating []byte
	_need_debug  []byte
	_id          int64
	_reason      []byte
	_duration    []byte
}

func New_TL_phoneCallDiscarded() *TL_phoneCallDiscarded {
	return &TL_phoneCallDiscarded{}
}

func (t *TL_phoneCallDiscarded) Encode() []byte {
	return nil
}

func (t *TL_phoneCallDiscarded) Decode(b []byte) {}

//====phoneConnection#9d4c17c0====

type TL_phoneConnection struct {
	_id       int64
	_ip       string
	_ipv6     string
	_port     []byte
	_peer_tag []byte
}

func New_TL_phoneConnection() *TL_phoneConnection {
	return &TL_phoneConnection{}
}

func (t *TL_phoneConnection) Encode() []byte {
	return nil
}

func (t *TL_phoneConnection) Decode(b []byte) {}

//====phoneCallProtocol#a2bb35cb====

type TL_phoneCallProtocol struct {
	_flags         []byte
	_udp_p2p       []byte
	_udp_reflector []byte
	_min_layer     []byte
	_max_layer     []byte
}

func New_TL_phoneCallProtocol() *TL_phoneCallProtocol {
	return &TL_phoneCallProtocol{}
}

func (t *TL_phoneCallProtocol) Encode() []byte {
	return nil
}

func (t *TL_phoneCallProtocol) Decode(b []byte) {}

//====phone_phoneCall#ec82e140====

type TL_phone_phoneCall struct {
	_phone_call []byte
	_users      []byte
}

func New_TL_phone_phoneCall() *TL_phone_phoneCall {
	return &TL_phone_phoneCall{}
}

func (t *TL_phone_phoneCall) Encode() []byte {
	return nil
}

func (t *TL_phone_phoneCall) Decode(b []byte) {}

//====upload_cdnFileReuploadNeeded#eea8e46e====

type TL_upload_cdnFileReuploadNeeded struct {
	_request_token []byte
}

func New_TL_upload_cdnFileReuploadNeeded() *TL_upload_cdnFileReuploadNeeded {
	return &TL_upload_cdnFileReuploadNeeded{}
}

func (t *TL_upload_cdnFileReuploadNeeded) Encode() []byte {
	return nil
}

func (t *TL_upload_cdnFileReuploadNeeded) Decode(b []byte) {}

//====upload_cdnFile#a99fca4f====

type TL_upload_cdnFile struct {
	_bytes []byte
}

func New_TL_upload_cdnFile() *TL_upload_cdnFile {
	return &TL_upload_cdnFile{}
}

func (t *TL_upload_cdnFile) Encode() []byte {
	return nil
}

func (t *TL_upload_cdnFile) Decode(b []byte) {}

//====cdnPublicKey#c982eaba====

type TL_cdnPublicKey struct {
	_dc_id      []byte
	_public_key string
}

func New_TL_cdnPublicKey() *TL_cdnPublicKey {
	return &TL_cdnPublicKey{}
}

func (t *TL_cdnPublicKey) Encode() []byte {
	return nil
}

func (t *TL_cdnPublicKey) Decode(b []byte) {}

//====cdnConfig#5725e40a====

type TL_cdnConfig struct {
	_public_keys []byte
}

func New_TL_cdnConfig() *TL_cdnConfig {
	return &TL_cdnConfig{}
}

func (t *TL_cdnConfig) Encode() []byte {
	return nil
}

func (t *TL_cdnConfig) Decode(b []byte) {}

//====langPackString#cad181f6====

type TL_langPackString struct {
	_key   string
	_value string
}

func New_TL_langPackString() *TL_langPackString {
	return &TL_langPackString{}
}

func (t *TL_langPackString) Encode() []byte {
	return nil
}

func (t *TL_langPackString) Decode(b []byte) {}

//====langPackStringPluralized#6c47ac9f====

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

func New_TL_langPackStringPluralized() *TL_langPackStringPluralized {
	return &TL_langPackStringPluralized{}
}

func (t *TL_langPackStringPluralized) Encode() []byte {
	return nil
}

func (t *TL_langPackStringPluralized) Decode(b []byte) {}

//====langPackStringDeleted#2979eeb2====

type TL_langPackStringDeleted struct {
	_key string
}

func New_TL_langPackStringDeleted() *TL_langPackStringDeleted {
	return &TL_langPackStringDeleted{}
}

func (t *TL_langPackStringDeleted) Encode() []byte {
	return nil
}

func (t *TL_langPackStringDeleted) Decode(b []byte) {}

//====langPackDifference#f385c1f6====

type TL_langPackDifference struct {
	_lang_code    string
	_from_version []byte
	_version      []byte
	_strings      []byte
}

func New_TL_langPackDifference() *TL_langPackDifference {
	return &TL_langPackDifference{}
}

func (t *TL_langPackDifference) Encode() []byte {
	return nil
}

func (t *TL_langPackDifference) Decode(b []byte) {}

//====langPackLanguage#117698f1====

type TL_langPackLanguage struct {
	_name        string
	_native_name string
	_lang_code   string
}

func New_TL_langPackLanguage() *TL_langPackLanguage {
	return &TL_langPackLanguage{}
}

func (t *TL_langPackLanguage) Encode() []byte {
	return nil
}

func (t *TL_langPackLanguage) Decode(b []byte) {}

//====channelAdminRights#5d7ceba5====

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

func New_TL_channelAdminRights() *TL_channelAdminRights {
	return &TL_channelAdminRights{}
}

func (t *TL_channelAdminRights) Encode() []byte {
	return nil
}

func (t *TL_channelAdminRights) Decode(b []byte) {}

//====channelBannedRights#58cf4249====

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
	_until_date    []byte
}

func New_TL_channelBannedRights() *TL_channelBannedRights {
	return &TL_channelBannedRights{}
}

func (t *TL_channelBannedRights) Encode() []byte {
	return nil
}

func (t *TL_channelBannedRights) Decode(b []byte) {}

//====channelAdminLogEventActionChangeTitle#e6dfb825====

type TL_channelAdminLogEventActionChangeTitle struct {
	_prev_value string
	_new_value  string
}

func New_TL_channelAdminLogEventActionChangeTitle() *TL_channelAdminLogEventActionChangeTitle {
	return &TL_channelAdminLogEventActionChangeTitle{}
}

func (t *TL_channelAdminLogEventActionChangeTitle) Encode() []byte {
	return nil
}

func (t *TL_channelAdminLogEventActionChangeTitle) Decode(b []byte) {}

//====channelAdminLogEventActionChangeAbout#55188a2e====

type TL_channelAdminLogEventActionChangeAbout struct {
	_prev_value string
	_new_value  string
}

func New_TL_channelAdminLogEventActionChangeAbout() *TL_channelAdminLogEventActionChangeAbout {
	return &TL_channelAdminLogEventActionChangeAbout{}
}

func (t *TL_channelAdminLogEventActionChangeAbout) Encode() []byte {
	return nil
}

func (t *TL_channelAdminLogEventActionChangeAbout) Decode(b []byte) {}

//====channelAdminLogEventActionChangeUsername#6a4afc38====

type TL_channelAdminLogEventActionChangeUsername struct {
	_prev_value string
	_new_value  string
}

func New_TL_channelAdminLogEventActionChangeUsername() *TL_channelAdminLogEventActionChangeUsername {
	return &TL_channelAdminLogEventActionChangeUsername{}
}

func (t *TL_channelAdminLogEventActionChangeUsername) Encode() []byte {
	return nil
}

func (t *TL_channelAdminLogEventActionChangeUsername) Decode(b []byte) {}

//====channelAdminLogEventActionChangePhoto#b82f55c3====

type TL_channelAdminLogEventActionChangePhoto struct {
	_prev_photo []byte
	_new_photo  []byte
}

func New_TL_channelAdminLogEventActionChangePhoto() *TL_channelAdminLogEventActionChangePhoto {
	return &TL_channelAdminLogEventActionChangePhoto{}
}

func (t *TL_channelAdminLogEventActionChangePhoto) Encode() []byte {
	return nil
}

func (t *TL_channelAdminLogEventActionChangePhoto) Decode(b []byte) {}

//====channelAdminLogEventActionToggleInvites#1b7907ae====

type TL_channelAdminLogEventActionToggleInvites struct {
	_new_value []byte
}

func New_TL_channelAdminLogEventActionToggleInvites() *TL_channelAdminLogEventActionToggleInvites {
	return &TL_channelAdminLogEventActionToggleInvites{}
}

func (t *TL_channelAdminLogEventActionToggleInvites) Encode() []byte {
	return nil
}

func (t *TL_channelAdminLogEventActionToggleInvites) Decode(b []byte) {}

//====channelAdminLogEventActionToggleSignatures#26ae0971====

type TL_channelAdminLogEventActionToggleSignatures struct {
	_new_value []byte
}

func New_TL_channelAdminLogEventActionToggleSignatures() *TL_channelAdminLogEventActionToggleSignatures {
	return &TL_channelAdminLogEventActionToggleSignatures{}
}

func (t *TL_channelAdminLogEventActionToggleSignatures) Encode() []byte {
	return nil
}

func (t *TL_channelAdminLogEventActionToggleSignatures) Decode(b []byte) {}

//====channelAdminLogEventActionUpdatePinned#e9e82c18====

type TL_channelAdminLogEventActionUpdatePinned struct {
	_message []byte
}

func New_TL_channelAdminLogEventActionUpdatePinned() *TL_channelAdminLogEventActionUpdatePinned {
	return &TL_channelAdminLogEventActionUpdatePinned{}
}

func (t *TL_channelAdminLogEventActionUpdatePinned) Encode() []byte {
	return nil
}

func (t *TL_channelAdminLogEventActionUpdatePinned) Decode(b []byte) {}

//====channelAdminLogEventActionEditMessage#709b2405====

type TL_channelAdminLogEventActionEditMessage struct {
	_prev_message []byte
	_new_message  []byte
}

func New_TL_channelAdminLogEventActionEditMessage() *TL_channelAdminLogEventActionEditMessage {
	return &TL_channelAdminLogEventActionEditMessage{}
}

func (t *TL_channelAdminLogEventActionEditMessage) Encode() []byte {
	return nil
}

func (t *TL_channelAdminLogEventActionEditMessage) Decode(b []byte) {}

//====channelAdminLogEventActionDeleteMessage#42e047bb====

type TL_channelAdminLogEventActionDeleteMessage struct {
	_message []byte
}

func New_TL_channelAdminLogEventActionDeleteMessage() *TL_channelAdminLogEventActionDeleteMessage {
	return &TL_channelAdminLogEventActionDeleteMessage{}
}

func (t *TL_channelAdminLogEventActionDeleteMessage) Encode() []byte {
	return nil
}

func (t *TL_channelAdminLogEventActionDeleteMessage) Decode(b []byte) {}

//====channelAdminLogEventActionParticipantJoin#183040d3====

type TL_channelAdminLogEventActionParticipantJoin struct {
}

func New_TL_channelAdminLogEventActionParticipantJoin() *TL_channelAdminLogEventActionParticipantJoin {
	return &TL_channelAdminLogEventActionParticipantJoin{}
}

func (t *TL_channelAdminLogEventActionParticipantJoin) Encode() []byte {
	return nil
}

func (t *TL_channelAdminLogEventActionParticipantJoin) Decode(b []byte) {}

//====channelAdminLogEventActionParticipantLeave#f89777f2====

type TL_channelAdminLogEventActionParticipantLeave struct {
}

func New_TL_channelAdminLogEventActionParticipantLeave() *TL_channelAdminLogEventActionParticipantLeave {
	return &TL_channelAdminLogEventActionParticipantLeave{}
}

func (t *TL_channelAdminLogEventActionParticipantLeave) Encode() []byte {
	return nil
}

func (t *TL_channelAdminLogEventActionParticipantLeave) Decode(b []byte) {}

//====channelAdminLogEventActionParticipantInvite#e31c34d8====

type TL_channelAdminLogEventActionParticipantInvite struct {
	_participant []byte
}

func New_TL_channelAdminLogEventActionParticipantInvite() *TL_channelAdminLogEventActionParticipantInvite {
	return &TL_channelAdminLogEventActionParticipantInvite{}
}

func (t *TL_channelAdminLogEventActionParticipantInvite) Encode() []byte {
	return nil
}

func (t *TL_channelAdminLogEventActionParticipantInvite) Decode(b []byte) {}

//====channelAdminLogEventActionParticipantToggleBan#e6d83d7e====

type TL_channelAdminLogEventActionParticipantToggleBan struct {
	_prev_participant []byte
	_new_participant  []byte
}

func New_TL_channelAdminLogEventActionParticipantToggleBan() *TL_channelAdminLogEventActionParticipantToggleBan {
	return &TL_channelAdminLogEventActionParticipantToggleBan{}
}

func (t *TL_channelAdminLogEventActionParticipantToggleBan) Encode() []byte {
	return nil
}

func (t *TL_channelAdminLogEventActionParticipantToggleBan) Decode(b []byte) {}

//====channelAdminLogEventActionParticipantToggleAdmin#d5676710====

type TL_channelAdminLogEventActionParticipantToggleAdmin struct {
	_prev_participant []byte
	_new_participant  []byte
}

func New_TL_channelAdminLogEventActionParticipantToggleAdmin() *TL_channelAdminLogEventActionParticipantToggleAdmin {
	return &TL_channelAdminLogEventActionParticipantToggleAdmin{}
}

func (t *TL_channelAdminLogEventActionParticipantToggleAdmin) Encode() []byte {
	return nil
}

func (t *TL_channelAdminLogEventActionParticipantToggleAdmin) Decode(b []byte) {}

//====channelAdminLogEventActionChangeStickerSet#b1c3caa7====

type TL_channelAdminLogEventActionChangeStickerSet struct {
	_prev_stickerset []byte
	_new_stickerset  []byte
}

func New_TL_channelAdminLogEventActionChangeStickerSet() *TL_channelAdminLogEventActionChangeStickerSet {
	return &TL_channelAdminLogEventActionChangeStickerSet{}
}

func (t *TL_channelAdminLogEventActionChangeStickerSet) Encode() []byte {
	return nil
}

func (t *TL_channelAdminLogEventActionChangeStickerSet) Decode(b []byte) {}

//====channelAdminLogEventActionTogglePreHistoryHidden#5f5c95f1====

type TL_channelAdminLogEventActionTogglePreHistoryHidden struct {
	_new_value []byte
}

func New_TL_channelAdminLogEventActionTogglePreHistoryHidden() *TL_channelAdminLogEventActionTogglePreHistoryHidden {
	return &TL_channelAdminLogEventActionTogglePreHistoryHidden{}
}

func (t *TL_channelAdminLogEventActionTogglePreHistoryHidden) Encode() []byte {
	return nil
}

func (t *TL_channelAdminLogEventActionTogglePreHistoryHidden) Decode(b []byte) {}

//====channelAdminLogEvent#3b5a3e40====

type TL_channelAdminLogEvent struct {
	_id      int64
	_date    []byte
	_user_id []byte
	_action  []byte
}

func New_TL_channelAdminLogEvent() *TL_channelAdminLogEvent {
	return &TL_channelAdminLogEvent{}
}

func (t *TL_channelAdminLogEvent) Encode() []byte {
	return nil
}

func (t *TL_channelAdminLogEvent) Decode(b []byte) {}

//====channels_adminLogResults#ed8af74d====

type TL_channels_adminLogResults struct {
	_events []byte
	_chats  []byte
	_users  []byte
}

func New_TL_channels_adminLogResults() *TL_channels_adminLogResults {
	return &TL_channels_adminLogResults{}
}

func (t *TL_channels_adminLogResults) Encode() []byte {
	return nil
}

func (t *TL_channels_adminLogResults) Decode(b []byte) {}

//====channelAdminLogEventsFilter#ea107ae4====

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

func New_TL_channelAdminLogEventsFilter() *TL_channelAdminLogEventsFilter {
	return &TL_channelAdminLogEventsFilter{}
}

func (t *TL_channelAdminLogEventsFilter) Encode() []byte {
	return nil
}

func (t *TL_channelAdminLogEventsFilter) Decode(b []byte) {}

//====popularContact#5ce14175====

type TL_popularContact struct {
	_client_id int64
	_importers []byte
}

func New_TL_popularContact() *TL_popularContact {
	return &TL_popularContact{}
}

func (t *TL_popularContact) Encode() []byte {
	return nil
}

func (t *TL_popularContact) Decode(b []byte) {}

//====cdnFileHash#77eec38f====

type TL_cdnFileHash struct {
	_offset []byte
	_limit  []byte
	_hash   []byte
}

func New_TL_cdnFileHash() *TL_cdnFileHash {
	return &TL_cdnFileHash{}
}

func (t *TL_cdnFileHash) Encode() []byte {
	return nil
}

func (t *TL_cdnFileHash) Decode(b []byte) {}

//====messages_favedStickersNotModified#9e8fa6d3====

type TL_messages_favedStickersNotModified struct {
}

func New_TL_messages_favedStickersNotModified() *TL_messages_favedStickersNotModified {
	return &TL_messages_favedStickersNotModified{}
}

func (t *TL_messages_favedStickersNotModified) Encode() []byte {
	return nil
}

func (t *TL_messages_favedStickersNotModified) Decode(b []byte) {}

//====messages_favedStickers#f37f2f16====

type TL_messages_favedStickers struct {
	_hash     []byte
	_packs    []byte
	_stickers []byte
}

func New_TL_messages_favedStickers() *TL_messages_favedStickers {
	return &TL_messages_favedStickers{}
}

func (t *TL_messages_favedStickers) Encode() []byte {
	return nil
}

func (t *TL_messages_favedStickers) Decode(b []byte) {}

//====recentMeUrlUnknown#46e1d13d====

type TL_recentMeUrlUnknown struct {
	_url string
}

func New_TL_recentMeUrlUnknown() *TL_recentMeUrlUnknown {
	return &TL_recentMeUrlUnknown{}
}

func (t *TL_recentMeUrlUnknown) Encode() []byte {
	return nil
}

func (t *TL_recentMeUrlUnknown) Decode(b []byte) {}

//====recentMeUrlUser#8dbc3336====

type TL_recentMeUrlUser struct {
	_url     string
	_user_id []byte
}

func New_TL_recentMeUrlUser() *TL_recentMeUrlUser {
	return &TL_recentMeUrlUser{}
}

func (t *TL_recentMeUrlUser) Encode() []byte {
	return nil
}

func (t *TL_recentMeUrlUser) Decode(b []byte) {}

//====recentMeUrlChat#a01b22f9====

type TL_recentMeUrlChat struct {
	_url     string
	_chat_id []byte
}

func New_TL_recentMeUrlChat() *TL_recentMeUrlChat {
	return &TL_recentMeUrlChat{}
}

func (t *TL_recentMeUrlChat) Encode() []byte {
	return nil
}

func (t *TL_recentMeUrlChat) Decode(b []byte) {}

//====recentMeUrlChatInvite#eb49081d====

type TL_recentMeUrlChatInvite struct {
	_url         string
	_chat_invite []byte
}

func New_TL_recentMeUrlChatInvite() *TL_recentMeUrlChatInvite {
	return &TL_recentMeUrlChatInvite{}
}

func (t *TL_recentMeUrlChatInvite) Encode() []byte {
	return nil
}

func (t *TL_recentMeUrlChatInvite) Decode(b []byte) {}

//====recentMeUrlStickerSet#bc0a57dc====

type TL_recentMeUrlStickerSet struct {
	_url string
	_set []byte
}

func New_TL_recentMeUrlStickerSet() *TL_recentMeUrlStickerSet {
	return &TL_recentMeUrlStickerSet{}
}

func (t *TL_recentMeUrlStickerSet) Encode() []byte {
	return nil
}

func (t *TL_recentMeUrlStickerSet) Decode(b []byte) {}

//====help_recentMeUrls#0e0310d7====

type TL_help_recentMeUrls struct {
	_urls  []byte
	_chats []byte
	_users []byte
}

func New_TL_help_recentMeUrls() *TL_help_recentMeUrls {
	return &TL_help_recentMeUrls{}
}

func (t *TL_help_recentMeUrls) Encode() []byte {
	return nil
}

func (t *TL_help_recentMeUrls) Decode(b []byte) {}

//====inputSingleMedia#5eaa7809====

type TL_inputSingleMedia struct {
	_media     []byte
	_random_id int64
}

func New_TL_inputSingleMedia() *TL_inputSingleMedia {
	return &TL_inputSingleMedia{}
}

func (t *TL_inputSingleMedia) Encode() []byte {
	return nil
}

func (t *TL_inputSingleMedia) Decode(b []byte) {}

//====invokeAfterMsg#cb9f372d====

type TL_invokeAfterMsg struct {
	_msg_id int64
	_query  []byte
}

func New_TL_invokeAfterMsg() *TL_invokeAfterMsg {
	return &TL_invokeAfterMsg{}
}

func (t *TL_invokeAfterMsg) Encode() []byte {
	return nil
}

func (t *TL_invokeAfterMsg) Decode(b []byte) {}

//====invokeAfterMsgs#3dc4b4f0====

type TL_invokeAfterMsgs struct {
	_msg_ids []byte
	_query   []byte
}

func New_TL_invokeAfterMsgs() *TL_invokeAfterMsgs {
	return &TL_invokeAfterMsgs{}
}

func (t *TL_invokeAfterMsgs) Encode() []byte {
	return nil
}

func (t *TL_invokeAfterMsgs) Decode(b []byte) {}

//====initConnection#c7481da6====

type TL_initConnection struct {
	_api_id           []byte
	_device_model     string
	_system_version   string
	_app_version      string
	_system_lang_code string
	_lang_pack        string
	_lang_code        string
	_query            []byte
}

func New_TL_initConnection() *TL_initConnection {
	return &TL_initConnection{}
}

func (t *TL_initConnection) Encode() []byte {
	return nil
}

func (t *TL_initConnection) Decode(b []byte) {}

//====invokeWithLayer#da9b0d0d====

type TL_invokeWithLayer struct {
	_layer []byte
	_query []byte
}

func New_TL_invokeWithLayer() *TL_invokeWithLayer {
	return &TL_invokeWithLayer{}
}

func (t *TL_invokeWithLayer) Encode() []byte {
	return nil
}

func (t *TL_invokeWithLayer) Decode(b []byte) {}

//====invokeWithoutUpdates#bf9459b7====

type TL_invokeWithoutUpdates struct {
	_query []byte
}

func New_TL_invokeWithoutUpdates() *TL_invokeWithoutUpdates {
	return &TL_invokeWithoutUpdates{}
}

func (t *TL_invokeWithoutUpdates) Encode() []byte {
	return nil
}

func (t *TL_invokeWithoutUpdates) Decode(b []byte) {}

//====auth_checkPhone#6fe51dfb====

type TL_auth_checkPhone struct {
	_phone_number string
}

func New_TL_auth_checkPhone() *TL_auth_checkPhone {
	return &TL_auth_checkPhone{}
}

func (t *TL_auth_checkPhone) Encode() []byte {
	return nil
}

func (t *TL_auth_checkPhone) Decode(b []byte) {}

//====auth_sendCode#86aef0ec====

type TL_auth_sendCode struct {
	_flags           []byte
	_allow_flashcall []byte
	_phone_number    string
	_current_number  []byte
	_api_id          []byte
	_api_hash        string
}

func New_TL_auth_sendCode() *TL_auth_sendCode {
	return &TL_auth_sendCode{}
}

func (t *TL_auth_sendCode) Encode() []byte {
	return nil
}

func (t *TL_auth_sendCode) Decode(b []byte) {}

//====auth_signUp#1b067634====

type TL_auth_signUp struct {
	_phone_number    string
	_phone_code_hash string
	_phone_code      string
	_first_name      string
	_last_name       string
}

func New_TL_auth_signUp() *TL_auth_signUp {
	return &TL_auth_signUp{}
}

func (t *TL_auth_signUp) Encode() []byte {
	return nil
}

func (t *TL_auth_signUp) Decode(b []byte) {}

//====auth_signIn#bcd51581====

type TL_auth_signIn struct {
	_phone_number    string
	_phone_code_hash string
	_phone_code      string
}

func New_TL_auth_signIn() *TL_auth_signIn {
	return &TL_auth_signIn{}
}

func (t *TL_auth_signIn) Encode() []byte {
	return nil
}

func (t *TL_auth_signIn) Decode(b []byte) {}

//====auth_logOut#5717da40====

type TL_auth_logOut struct {
}

func New_TL_auth_logOut() *TL_auth_logOut {
	return &TL_auth_logOut{}
}

func (t *TL_auth_logOut) Encode() []byte {
	return nil
}

func (t *TL_auth_logOut) Decode(b []byte) {}

//====auth_resetAuthorizations#9fab0d1a====

type TL_auth_resetAuthorizations struct {
}

func New_TL_auth_resetAuthorizations() *TL_auth_resetAuthorizations {
	return &TL_auth_resetAuthorizations{}
}

func (t *TL_auth_resetAuthorizations) Encode() []byte {
	return nil
}

func (t *TL_auth_resetAuthorizations) Decode(b []byte) {}

//====auth_sendInvites#771c1d97====

type TL_auth_sendInvites struct {
	_phone_numbers []byte
	_message       string
}

func New_TL_auth_sendInvites() *TL_auth_sendInvites {
	return &TL_auth_sendInvites{}
}

func (t *TL_auth_sendInvites) Encode() []byte {
	return nil
}

func (t *TL_auth_sendInvites) Decode(b []byte) {}

//====auth_exportAuthorization#e5bfffcd====

type TL_auth_exportAuthorization struct {
	_dc_id []byte
}

func New_TL_auth_exportAuthorization() *TL_auth_exportAuthorization {
	return &TL_auth_exportAuthorization{}
}

func (t *TL_auth_exportAuthorization) Encode() []byte {
	return nil
}

func (t *TL_auth_exportAuthorization) Decode(b []byte) {}

//====auth_importAuthorization#e3ef9613====

type TL_auth_importAuthorization struct {
	_id    []byte
	_bytes []byte
}

func New_TL_auth_importAuthorization() *TL_auth_importAuthorization {
	return &TL_auth_importAuthorization{}
}

func (t *TL_auth_importAuthorization) Encode() []byte {
	return nil
}

func (t *TL_auth_importAuthorization) Decode(b []byte) {}

//====auth_bindTempAuthKey#cdd42a05====

type TL_auth_bindTempAuthKey struct {
	_perm_auth_key_id  int64
	_nonce             int64
	_expires_at        []byte
	_encrypted_message []byte
}

func New_TL_auth_bindTempAuthKey() *TL_auth_bindTempAuthKey {
	return &TL_auth_bindTempAuthKey{}
}

func (t *TL_auth_bindTempAuthKey) Encode() []byte {
	return nil
}

func (t *TL_auth_bindTempAuthKey) Decode(b []byte) {}

//====auth_importBotAuthorization#67a3ff2c====

type TL_auth_importBotAuthorization struct {
	_flags          []byte
	_api_id         []byte
	_api_hash       string
	_bot_auth_token string
}

func New_TL_auth_importBotAuthorization() *TL_auth_importBotAuthorization {
	return &TL_auth_importBotAuthorization{}
}

func (t *TL_auth_importBotAuthorization) Encode() []byte {
	return nil
}

func (t *TL_auth_importBotAuthorization) Decode(b []byte) {}

//====auth_checkPassword#0a63011e====

type TL_auth_checkPassword struct {
	_password_hash []byte
}

func New_TL_auth_checkPassword() *TL_auth_checkPassword {
	return &TL_auth_checkPassword{}
}

func (t *TL_auth_checkPassword) Encode() []byte {
	return nil
}

func (t *TL_auth_checkPassword) Decode(b []byte) {}

//====auth_requestPasswordRecovery#d897bc66====

type TL_auth_requestPasswordRecovery struct {
}

func New_TL_auth_requestPasswordRecovery() *TL_auth_requestPasswordRecovery {
	return &TL_auth_requestPasswordRecovery{}
}

func (t *TL_auth_requestPasswordRecovery) Encode() []byte {
	return nil
}

func (t *TL_auth_requestPasswordRecovery) Decode(b []byte) {}

//====auth_recoverPassword#4ea56e92====

type TL_auth_recoverPassword struct {
	_code string
}

func New_TL_auth_recoverPassword() *TL_auth_recoverPassword {
	return &TL_auth_recoverPassword{}
}

func (t *TL_auth_recoverPassword) Encode() []byte {
	return nil
}

func (t *TL_auth_recoverPassword) Decode(b []byte) {}

//====auth_resendCode#3ef1a9bf====

type TL_auth_resendCode struct {
	_phone_number    string
	_phone_code_hash string
}

func New_TL_auth_resendCode() *TL_auth_resendCode {
	return &TL_auth_resendCode{}
}

func (t *TL_auth_resendCode) Encode() []byte {
	return nil
}

func (t *TL_auth_resendCode) Decode(b []byte) {}

//====auth_cancelCode#1f040578====

type TL_auth_cancelCode struct {
	_phone_number    string
	_phone_code_hash string
}

func New_TL_auth_cancelCode() *TL_auth_cancelCode {
	return &TL_auth_cancelCode{}
}

func (t *TL_auth_cancelCode) Encode() []byte {
	return nil
}

func (t *TL_auth_cancelCode) Decode(b []byte) {}

//====auth_dropTempAuthKeys#8e48a188====

type TL_auth_dropTempAuthKeys struct {
	_except_auth_keys []byte
}

func New_TL_auth_dropTempAuthKeys() *TL_auth_dropTempAuthKeys {
	return &TL_auth_dropTempAuthKeys{}
}

func (t *TL_auth_dropTempAuthKeys) Encode() []byte {
	return nil
}

func (t *TL_auth_dropTempAuthKeys) Decode(b []byte) {}

//====account_registerDevice#637ea878====

type TL_account_registerDevice struct {
	_token_type []byte
	_token      string
}

func New_TL_account_registerDevice() *TL_account_registerDevice {
	return &TL_account_registerDevice{}
}

func (t *TL_account_registerDevice) Encode() []byte {
	return nil
}

func (t *TL_account_registerDevice) Decode(b []byte) {}

//====account_unregisterDevice#65c55b40====

type TL_account_unregisterDevice struct {
	_token_type []byte
	_token      string
}

func New_TL_account_unregisterDevice() *TL_account_unregisterDevice {
	return &TL_account_unregisterDevice{}
}

func (t *TL_account_unregisterDevice) Encode() []byte {
	return nil
}

func (t *TL_account_unregisterDevice) Decode(b []byte) {}

//====account_updateNotifySettings#84be5b93====

type TL_account_updateNotifySettings struct {
	_peer     []byte
	_settings []byte
}

func New_TL_account_updateNotifySettings() *TL_account_updateNotifySettings {
	return &TL_account_updateNotifySettings{}
}

func (t *TL_account_updateNotifySettings) Encode() []byte {
	return nil
}

func (t *TL_account_updateNotifySettings) Decode(b []byte) {}

//====account_getNotifySettings#12b3ad31====

type TL_account_getNotifySettings struct {
	_peer []byte
}

func New_TL_account_getNotifySettings() *TL_account_getNotifySettings {
	return &TL_account_getNotifySettings{}
}

func (t *TL_account_getNotifySettings) Encode() []byte {
	return nil
}

func (t *TL_account_getNotifySettings) Decode(b []byte) {}

//====account_resetNotifySettings#db7e1747====

type TL_account_resetNotifySettings struct {
}

func New_TL_account_resetNotifySettings() *TL_account_resetNotifySettings {
	return &TL_account_resetNotifySettings{}
}

func (t *TL_account_resetNotifySettings) Encode() []byte {
	return nil
}

func (t *TL_account_resetNotifySettings) Decode(b []byte) {}

//====account_updateProfile#78515775====

type TL_account_updateProfile struct {
	_flags      []byte
	_first_name []byte
	_last_name  []byte
	_about      []byte
}

func New_TL_account_updateProfile() *TL_account_updateProfile {
	return &TL_account_updateProfile{}
}

func (t *TL_account_updateProfile) Encode() []byte {
	return nil
}

func (t *TL_account_updateProfile) Decode(b []byte) {}

//====account_updateStatus#6628562c====

type TL_account_updateStatus struct {
	_offline []byte
}

func New_TL_account_updateStatus() *TL_account_updateStatus {
	return &TL_account_updateStatus{}
}

func (t *TL_account_updateStatus) Encode() []byte {
	return nil
}

func (t *TL_account_updateStatus) Decode(b []byte) {}

//====account_getWallPapers#c04cfac2====

type TL_account_getWallPapers struct {
}

func New_TL_account_getWallPapers() *TL_account_getWallPapers {
	return &TL_account_getWallPapers{}
}

func (t *TL_account_getWallPapers) Encode() []byte {
	return nil
}

func (t *TL_account_getWallPapers) Decode(b []byte) {}

//====account_reportPeer#ae189d5f====

type TL_account_reportPeer struct {
	_peer   []byte
	_reason []byte
}

func New_TL_account_reportPeer() *TL_account_reportPeer {
	return &TL_account_reportPeer{}
}

func (t *TL_account_reportPeer) Encode() []byte {
	return nil
}

func (t *TL_account_reportPeer) Decode(b []byte) {}

//====account_checkUsername#2714d86c====

type TL_account_checkUsername struct {
	_username string
}

func New_TL_account_checkUsername() *TL_account_checkUsername {
	return &TL_account_checkUsername{}
}

func (t *TL_account_checkUsername) Encode() []byte {
	return nil
}

func (t *TL_account_checkUsername) Decode(b []byte) {}

//====account_updateUsername#3e0bdd7c====

type TL_account_updateUsername struct {
	_username string
}

func New_TL_account_updateUsername() *TL_account_updateUsername {
	return &TL_account_updateUsername{}
}

func (t *TL_account_updateUsername) Encode() []byte {
	return nil
}

func (t *TL_account_updateUsername) Decode(b []byte) {}

//====account_getPrivacy#dadbc950====

type TL_account_getPrivacy struct {
	_key []byte
}

func New_TL_account_getPrivacy() *TL_account_getPrivacy {
	return &TL_account_getPrivacy{}
}

func (t *TL_account_getPrivacy) Encode() []byte {
	return nil
}

func (t *TL_account_getPrivacy) Decode(b []byte) {}

//====account_setPrivacy#c9f81ce8====

type TL_account_setPrivacy struct {
	_key   []byte
	_rules []byte
}

func New_TL_account_setPrivacy() *TL_account_setPrivacy {
	return &TL_account_setPrivacy{}
}

func (t *TL_account_setPrivacy) Encode() []byte {
	return nil
}

func (t *TL_account_setPrivacy) Decode(b []byte) {}

//====account_deleteAccount#418d4e0b====

type TL_account_deleteAccount struct {
	_reason string
}

func New_TL_account_deleteAccount() *TL_account_deleteAccount {
	return &TL_account_deleteAccount{}
}

func (t *TL_account_deleteAccount) Encode() []byte {
	return nil
}

func (t *TL_account_deleteAccount) Decode(b []byte) {}

//====account_getAccountTTL#08fc711d====

type TL_account_getAccountTTL struct {
}

func New_TL_account_getAccountTTL() *TL_account_getAccountTTL {
	return &TL_account_getAccountTTL{}
}

func (t *TL_account_getAccountTTL) Encode() []byte {
	return nil
}

func (t *TL_account_getAccountTTL) Decode(b []byte) {}

//====account_setAccountTTL#2442485e====

type TL_account_setAccountTTL struct {
	_ttl []byte
}

func New_TL_account_setAccountTTL() *TL_account_setAccountTTL {
	return &TL_account_setAccountTTL{}
}

func (t *TL_account_setAccountTTL) Encode() []byte {
	return nil
}

func (t *TL_account_setAccountTTL) Decode(b []byte) {}

//====account_sendChangePhoneCode#08e57deb====

type TL_account_sendChangePhoneCode struct {
	_flags           []byte
	_allow_flashcall []byte
	_phone_number    string
	_current_number  []byte
}

func New_TL_account_sendChangePhoneCode() *TL_account_sendChangePhoneCode {
	return &TL_account_sendChangePhoneCode{}
}

func (t *TL_account_sendChangePhoneCode) Encode() []byte {
	return nil
}

func (t *TL_account_sendChangePhoneCode) Decode(b []byte) {}

//====account_changePhone#70c32edb====

type TL_account_changePhone struct {
	_phone_number    string
	_phone_code_hash string
	_phone_code      string
}

func New_TL_account_changePhone() *TL_account_changePhone {
	return &TL_account_changePhone{}
}

func (t *TL_account_changePhone) Encode() []byte {
	return nil
}

func (t *TL_account_changePhone) Decode(b []byte) {}

//====account_updateDeviceLocked#38df3532====

type TL_account_updateDeviceLocked struct {
	_period []byte
}

func New_TL_account_updateDeviceLocked() *TL_account_updateDeviceLocked {
	return &TL_account_updateDeviceLocked{}
}

func (t *TL_account_updateDeviceLocked) Encode() []byte {
	return nil
}

func (t *TL_account_updateDeviceLocked) Decode(b []byte) {}

//====account_getAuthorizations#e320c158====

type TL_account_getAuthorizations struct {
}

func New_TL_account_getAuthorizations() *TL_account_getAuthorizations {
	return &TL_account_getAuthorizations{}
}

func (t *TL_account_getAuthorizations) Encode() []byte {
	return nil
}

func (t *TL_account_getAuthorizations) Decode(b []byte) {}

//====account_resetAuthorization#df77f3bc====

type TL_account_resetAuthorization struct {
	_hash int64
}

func New_TL_account_resetAuthorization() *TL_account_resetAuthorization {
	return &TL_account_resetAuthorization{}
}

func (t *TL_account_resetAuthorization) Encode() []byte {
	return nil
}

func (t *TL_account_resetAuthorization) Decode(b []byte) {}

//====account_getPassword#548a30f5====

type TL_account_getPassword struct {
}

func New_TL_account_getPassword() *TL_account_getPassword {
	return &TL_account_getPassword{}
}

func (t *TL_account_getPassword) Encode() []byte {
	return nil
}

func (t *TL_account_getPassword) Decode(b []byte) {}

//====account_getPasswordSettings#bc8d11bb====

type TL_account_getPasswordSettings struct {
	_current_password_hash []byte
}

func New_TL_account_getPasswordSettings() *TL_account_getPasswordSettings {
	return &TL_account_getPasswordSettings{}
}

func (t *TL_account_getPasswordSettings) Encode() []byte {
	return nil
}

func (t *TL_account_getPasswordSettings) Decode(b []byte) {}

//====account_updatePasswordSettings#fa7c4b86====

type TL_account_updatePasswordSettings struct {
	_current_password_hash []byte
	_new_settings          []byte
}

func New_TL_account_updatePasswordSettings() *TL_account_updatePasswordSettings {
	return &TL_account_updatePasswordSettings{}
}

func (t *TL_account_updatePasswordSettings) Encode() []byte {
	return nil
}

func (t *TL_account_updatePasswordSettings) Decode(b []byte) {}

//====account_sendConfirmPhoneCode#1516d7bd====

type TL_account_sendConfirmPhoneCode struct {
	_flags           []byte
	_allow_flashcall []byte
	_hash            string
	_current_number  []byte
}

func New_TL_account_sendConfirmPhoneCode() *TL_account_sendConfirmPhoneCode {
	return &TL_account_sendConfirmPhoneCode{}
}

func (t *TL_account_sendConfirmPhoneCode) Encode() []byte {
	return nil
}

func (t *TL_account_sendConfirmPhoneCode) Decode(b []byte) {}

//====account_confirmPhone#5f2178c3====

type TL_account_confirmPhone struct {
	_phone_code_hash string
	_phone_code      string
}

func New_TL_account_confirmPhone() *TL_account_confirmPhone {
	return &TL_account_confirmPhone{}
}

func (t *TL_account_confirmPhone) Encode() []byte {
	return nil
}

func (t *TL_account_confirmPhone) Decode(b []byte) {}

//====account_getTmpPassword#4a82327e====

type TL_account_getTmpPassword struct {
	_password_hash []byte
	_period        []byte
}

func New_TL_account_getTmpPassword() *TL_account_getTmpPassword {
	return &TL_account_getTmpPassword{}
}

func (t *TL_account_getTmpPassword) Encode() []byte {
	return nil
}

func (t *TL_account_getTmpPassword) Decode(b []byte) {}

//====users_getUsers#0d91a548====

type TL_users_getUsers struct {
	_id []byte
}

func New_TL_users_getUsers() *TL_users_getUsers {
	return &TL_users_getUsers{}
}

func (t *TL_users_getUsers) Encode() []byte {
	return nil
}

func (t *TL_users_getUsers) Decode(b []byte) {}

//====users_getFullUser#ca30a5b1====

type TL_users_getFullUser struct {
	_id []byte
}

func New_TL_users_getFullUser() *TL_users_getFullUser {
	return &TL_users_getFullUser{}
}

func (t *TL_users_getFullUser) Encode() []byte {
	return nil
}

func (t *TL_users_getFullUser) Decode(b []byte) {}

//====contacts_getStatuses#c4a353ee====

type TL_contacts_getStatuses struct {
}

func New_TL_contacts_getStatuses() *TL_contacts_getStatuses {
	return &TL_contacts_getStatuses{}
}

func (t *TL_contacts_getStatuses) Encode() []byte {
	return nil
}

func (t *TL_contacts_getStatuses) Decode(b []byte) {}

//====contacts_getContacts#c023849f====

type TL_contacts_getContacts struct {
	_hash []byte
}

func New_TL_contacts_getContacts() *TL_contacts_getContacts {
	return &TL_contacts_getContacts{}
}

func (t *TL_contacts_getContacts) Encode() []byte {
	return nil
}

func (t *TL_contacts_getContacts) Decode(b []byte) {}

//====contacts_importContacts#2c800be5====

type TL_contacts_importContacts struct {
	_contacts []byte
}

func New_TL_contacts_importContacts() *TL_contacts_importContacts {
	return &TL_contacts_importContacts{}
}

func (t *TL_contacts_importContacts) Encode() []byte {
	return nil
}

func (t *TL_contacts_importContacts) Decode(b []byte) {}

//====contacts_deleteContact#8e953744====

type TL_contacts_deleteContact struct {
	_id []byte
}

func New_TL_contacts_deleteContact() *TL_contacts_deleteContact {
	return &TL_contacts_deleteContact{}
}

func (t *TL_contacts_deleteContact) Encode() []byte {
	return nil
}

func (t *TL_contacts_deleteContact) Decode(b []byte) {}

//====contacts_deleteContacts#59ab389e====

type TL_contacts_deleteContacts struct {
	_id []byte
}

func New_TL_contacts_deleteContacts() *TL_contacts_deleteContacts {
	return &TL_contacts_deleteContacts{}
}

func (t *TL_contacts_deleteContacts) Encode() []byte {
	return nil
}

func (t *TL_contacts_deleteContacts) Decode(b []byte) {}

//====contacts_block#332b49fc====

type TL_contacts_block struct {
	_id []byte
}

func New_TL_contacts_block() *TL_contacts_block {
	return &TL_contacts_block{}
}

func (t *TL_contacts_block) Encode() []byte {
	return nil
}

func (t *TL_contacts_block) Decode(b []byte) {}

//====contacts_unblock#e54100bd====

type TL_contacts_unblock struct {
	_id []byte
}

func New_TL_contacts_unblock() *TL_contacts_unblock {
	return &TL_contacts_unblock{}
}

func (t *TL_contacts_unblock) Encode() []byte {
	return nil
}

func (t *TL_contacts_unblock) Decode(b []byte) {}

//====contacts_getBlocked#f57c350f====

type TL_contacts_getBlocked struct {
	_offset []byte
	_limit  []byte
}

func New_TL_contacts_getBlocked() *TL_contacts_getBlocked {
	return &TL_contacts_getBlocked{}
}

func (t *TL_contacts_getBlocked) Encode() []byte {
	return nil
}

func (t *TL_contacts_getBlocked) Decode(b []byte) {}

//====contacts_exportCard#84e53737====

type TL_contacts_exportCard struct {
}

func New_TL_contacts_exportCard() *TL_contacts_exportCard {
	return &TL_contacts_exportCard{}
}

func (t *TL_contacts_exportCard) Encode() []byte {
	return nil
}

func (t *TL_contacts_exportCard) Decode(b []byte) {}

//====contacts_importCard#4fe196fe====

type TL_contacts_importCard struct {
	_export_card []byte
}

func New_TL_contacts_importCard() *TL_contacts_importCard {
	return &TL_contacts_importCard{}
}

func (t *TL_contacts_importCard) Encode() []byte {
	return nil
}

func (t *TL_contacts_importCard) Decode(b []byte) {}

//====contacts_search#11f812d8====

type TL_contacts_search struct {
	_q     string
	_limit []byte
}

func New_TL_contacts_search() *TL_contacts_search {
	return &TL_contacts_search{}
}

func (t *TL_contacts_search) Encode() []byte {
	return nil
}

func (t *TL_contacts_search) Decode(b []byte) {}

//====contacts_resolveUsername#f93ccba3====

type TL_contacts_resolveUsername struct {
	_username string
}

func New_TL_contacts_resolveUsername() *TL_contacts_resolveUsername {
	return &TL_contacts_resolveUsername{}
}

func (t *TL_contacts_resolveUsername) Encode() []byte {
	return nil
}

func (t *TL_contacts_resolveUsername) Decode(b []byte) {}

//====contacts_getTopPeers#d4982db5====

type TL_contacts_getTopPeers struct {
	_flags          []byte
	_correspondents []byte
	_bots_pm        []byte
	_bots_inline    []byte
	_phone_calls    []byte
	_groups         []byte
	_channels       []byte
	_offset         []byte
	_limit          []byte
	_hash           []byte
}

func New_TL_contacts_getTopPeers() *TL_contacts_getTopPeers {
	return &TL_contacts_getTopPeers{}
}

func (t *TL_contacts_getTopPeers) Encode() []byte {
	return nil
}

func (t *TL_contacts_getTopPeers) Decode(b []byte) {}

//====contacts_resetTopPeerRating#1ae373ac====

type TL_contacts_resetTopPeerRating struct {
	_category []byte
	_peer     []byte
}

func New_TL_contacts_resetTopPeerRating() *TL_contacts_resetTopPeerRating {
	return &TL_contacts_resetTopPeerRating{}
}

func (t *TL_contacts_resetTopPeerRating) Encode() []byte {
	return nil
}

func (t *TL_contacts_resetTopPeerRating) Decode(b []byte) {}

//====contacts_resetSaved#879537f1====

type TL_contacts_resetSaved struct {
}

func New_TL_contacts_resetSaved() *TL_contacts_resetSaved {
	return &TL_contacts_resetSaved{}
}

func (t *TL_contacts_resetSaved) Encode() []byte {
	return nil
}

func (t *TL_contacts_resetSaved) Decode(b []byte) {}

//====messages_getMessages#4222fa74====

type TL_messages_getMessages struct {
	_id []byte
}

func New_TL_messages_getMessages() *TL_messages_getMessages {
	return &TL_messages_getMessages{}
}

func (t *TL_messages_getMessages) Encode() []byte {
	return nil
}

func (t *TL_messages_getMessages) Decode(b []byte) {}

//====messages_getDialogs#191ba9c5====

type TL_messages_getDialogs struct {
	_flags          []byte
	_exclude_pinned []byte
	_offset_date    []byte
	_offset_id      []byte
	_offset_peer    []byte
	_limit          []byte
}

func New_TL_messages_getDialogs() *TL_messages_getDialogs {
	return &TL_messages_getDialogs{}
}

func (t *TL_messages_getDialogs) Encode() []byte {
	return nil
}

func (t *TL_messages_getDialogs) Decode(b []byte) {}

//====messages_getHistory#dcbb8260====

type TL_messages_getHistory struct {
	_peer        []byte
	_offset_id   []byte
	_offset_date []byte
	_add_offset  []byte
	_limit       []byte
	_max_id      []byte
	_min_id      []byte
	_hash        []byte
}

func New_TL_messages_getHistory() *TL_messages_getHistory {
	return &TL_messages_getHistory{}
}

func (t *TL_messages_getHistory) Encode() []byte {
	return nil
}

func (t *TL_messages_getHistory) Decode(b []byte) {}

//====messages_search#039e9ea0====

type TL_messages_search struct {
	_flags      []byte
	_peer       []byte
	_q          string
	_from_id    []byte
	_filter     []byte
	_min_date   []byte
	_max_date   []byte
	_offset_id  []byte
	_add_offset []byte
	_limit      []byte
	_max_id     []byte
	_min_id     []byte
}

func New_TL_messages_search() *TL_messages_search {
	return &TL_messages_search{}
}

func (t *TL_messages_search) Encode() []byte {
	return nil
}

func (t *TL_messages_search) Decode(b []byte) {}

//====messages_readHistory#0e306d3a====

type TL_messages_readHistory struct {
	_peer   []byte
	_max_id []byte
}

func New_TL_messages_readHistory() *TL_messages_readHistory {
	return &TL_messages_readHistory{}
}

func (t *TL_messages_readHistory) Encode() []byte {
	return nil
}

func (t *TL_messages_readHistory) Decode(b []byte) {}

//====messages_deleteHistory#1c015b09====

type TL_messages_deleteHistory struct {
	_flags      []byte
	_just_clear []byte
	_peer       []byte
	_max_id     []byte
}

func New_TL_messages_deleteHistory() *TL_messages_deleteHistory {
	return &TL_messages_deleteHistory{}
}

func (t *TL_messages_deleteHistory) Encode() []byte {
	return nil
}

func (t *TL_messages_deleteHistory) Decode(b []byte) {}

//====messages_deleteMessages#e58e95d2====

type TL_messages_deleteMessages struct {
	_flags  []byte
	_revoke []byte
	_id     []byte
}

func New_TL_messages_deleteMessages() *TL_messages_deleteMessages {
	return &TL_messages_deleteMessages{}
}

func (t *TL_messages_deleteMessages) Encode() []byte {
	return nil
}

func (t *TL_messages_deleteMessages) Decode(b []byte) {}

//====messages_receivedMessages#05a954c0====

type TL_messages_receivedMessages struct {
	_max_id []byte
}

func New_TL_messages_receivedMessages() *TL_messages_receivedMessages {
	return &TL_messages_receivedMessages{}
}

func (t *TL_messages_receivedMessages) Encode() []byte {
	return nil
}

func (t *TL_messages_receivedMessages) Decode(b []byte) {}

//====messages_setTyping#a3825e50====

type TL_messages_setTyping struct {
	_peer   []byte
	_action []byte
}

func New_TL_messages_setTyping() *TL_messages_setTyping {
	return &TL_messages_setTyping{}
}

func (t *TL_messages_setTyping) Encode() []byte {
	return nil
}

func (t *TL_messages_setTyping) Decode(b []byte) {}

//====messages_sendMessage#fa88427a====

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

func New_TL_messages_sendMessage() *TL_messages_sendMessage {
	return &TL_messages_sendMessage{}
}

func (t *TL_messages_sendMessage) Encode() []byte {
	return nil
}

func (t *TL_messages_sendMessage) Decode(b []byte) {}

//====messages_sendMedia#c8f16791====

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

func New_TL_messages_sendMedia() *TL_messages_sendMedia {
	return &TL_messages_sendMedia{}
}

func (t *TL_messages_sendMedia) Encode() []byte {
	return nil
}

func (t *TL_messages_sendMedia) Decode(b []byte) {}

//====messages_forwardMessages#708e0195====

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

func New_TL_messages_forwardMessages() *TL_messages_forwardMessages {
	return &TL_messages_forwardMessages{}
}

func (t *TL_messages_forwardMessages) Encode() []byte {
	return nil
}

func (t *TL_messages_forwardMessages) Decode(b []byte) {}

//====messages_reportSpam#cf1592db====

type TL_messages_reportSpam struct {
	_peer []byte
}

func New_TL_messages_reportSpam() *TL_messages_reportSpam {
	return &TL_messages_reportSpam{}
}

func (t *TL_messages_reportSpam) Encode() []byte {
	return nil
}

func (t *TL_messages_reportSpam) Decode(b []byte) {}

//====messages_hideReportSpam#a8f1709b====

type TL_messages_hideReportSpam struct {
	_peer []byte
}

func New_TL_messages_hideReportSpam() *TL_messages_hideReportSpam {
	return &TL_messages_hideReportSpam{}
}

func (t *TL_messages_hideReportSpam) Encode() []byte {
	return nil
}

func (t *TL_messages_hideReportSpam) Decode(b []byte) {}

//====messages_getPeerSettings#3672e09c====

type TL_messages_getPeerSettings struct {
	_peer []byte
}

func New_TL_messages_getPeerSettings() *TL_messages_getPeerSettings {
	return &TL_messages_getPeerSettings{}
}

func (t *TL_messages_getPeerSettings) Encode() []byte {
	return nil
}

func (t *TL_messages_getPeerSettings) Decode(b []byte) {}

//====messages_getChats#3c6aa187====

type TL_messages_getChats struct {
	_id []byte
}

func New_TL_messages_getChats() *TL_messages_getChats {
	return &TL_messages_getChats{}
}

func (t *TL_messages_getChats) Encode() []byte {
	return nil
}

func (t *TL_messages_getChats) Decode(b []byte) {}

//====messages_getFullChat#3b831c66====

type TL_messages_getFullChat struct {
	_chat_id []byte
}

func New_TL_messages_getFullChat() *TL_messages_getFullChat {
	return &TL_messages_getFullChat{}
}

func (t *TL_messages_getFullChat) Encode() []byte {
	return nil
}

func (t *TL_messages_getFullChat) Decode(b []byte) {}

//====messages_editChatTitle#dc452855====

type TL_messages_editChatTitle struct {
	_chat_id []byte
	_title   string
}

func New_TL_messages_editChatTitle() *TL_messages_editChatTitle {
	return &TL_messages_editChatTitle{}
}

func (t *TL_messages_editChatTitle) Encode() []byte {
	return nil
}

func (t *TL_messages_editChatTitle) Decode(b []byte) {}

//====messages_editChatPhoto#ca4c79d8====

type TL_messages_editChatPhoto struct {
	_chat_id []byte
	_photo   []byte
}

func New_TL_messages_editChatPhoto() *TL_messages_editChatPhoto {
	return &TL_messages_editChatPhoto{}
}

func (t *TL_messages_editChatPhoto) Encode() []byte {
	return nil
}

func (t *TL_messages_editChatPhoto) Decode(b []byte) {}

//====messages_addChatUser#f9a0aa09====

type TL_messages_addChatUser struct {
	_chat_id   []byte
	_user_id   []byte
	_fwd_limit []byte
}

func New_TL_messages_addChatUser() *TL_messages_addChatUser {
	return &TL_messages_addChatUser{}
}

func (t *TL_messages_addChatUser) Encode() []byte {
	return nil
}

func (t *TL_messages_addChatUser) Decode(b []byte) {}

//====messages_deleteChatUser#e0611f16====

type TL_messages_deleteChatUser struct {
	_chat_id []byte
	_user_id []byte
}

func New_TL_messages_deleteChatUser() *TL_messages_deleteChatUser {
	return &TL_messages_deleteChatUser{}
}

func (t *TL_messages_deleteChatUser) Encode() []byte {
	return nil
}

func (t *TL_messages_deleteChatUser) Decode(b []byte) {}

//====messages_createChat#09cb126e====

type TL_messages_createChat struct {
	_users []byte
	_title string
}

func New_TL_messages_createChat() *TL_messages_createChat {
	return &TL_messages_createChat{}
}

func (t *TL_messages_createChat) Encode() []byte {
	return nil
}

func (t *TL_messages_createChat) Decode(b []byte) {}

//====messages_forwardMessage#33963bf9====

type TL_messages_forwardMessage struct {
	_peer      []byte
	_id        []byte
	_random_id int64
}

func New_TL_messages_forwardMessage() *TL_messages_forwardMessage {
	return &TL_messages_forwardMessage{}
}

func (t *TL_messages_forwardMessage) Encode() []byte {
	return nil
}

func (t *TL_messages_forwardMessage) Decode(b []byte) {}

//====messages_getDhConfig#26cf8950====

type TL_messages_getDhConfig struct {
	_version       []byte
	_random_length []byte
}

func New_TL_messages_getDhConfig() *TL_messages_getDhConfig {
	return &TL_messages_getDhConfig{}
}

func (t *TL_messages_getDhConfig) Encode() []byte {
	return nil
}

func (t *TL_messages_getDhConfig) Decode(b []byte) {}

//====messages_requestEncryption#f64daf43====

type TL_messages_requestEncryption struct {
	_user_id   []byte
	_random_id []byte
	_g_a       []byte
}

func New_TL_messages_requestEncryption() *TL_messages_requestEncryption {
	return &TL_messages_requestEncryption{}
}

func (t *TL_messages_requestEncryption) Encode() []byte {
	return nil
}

func (t *TL_messages_requestEncryption) Decode(b []byte) {}

//====messages_acceptEncryption#3dbc0415====

type TL_messages_acceptEncryption struct {
	_peer            []byte
	_g_b             []byte
	_key_fingerprint int64
}

func New_TL_messages_acceptEncryption() *TL_messages_acceptEncryption {
	return &TL_messages_acceptEncryption{}
}

func (t *TL_messages_acceptEncryption) Encode() []byte {
	return nil
}

func (t *TL_messages_acceptEncryption) Decode(b []byte) {}

//====messages_discardEncryption#edd923c5====

type TL_messages_discardEncryption struct {
	_chat_id []byte
}

func New_TL_messages_discardEncryption() *TL_messages_discardEncryption {
	return &TL_messages_discardEncryption{}
}

func (t *TL_messages_discardEncryption) Encode() []byte {
	return nil
}

func (t *TL_messages_discardEncryption) Decode(b []byte) {}

//====messages_setEncryptedTyping#791451ed====

type TL_messages_setEncryptedTyping struct {
	_peer   []byte
	_typing []byte
}

func New_TL_messages_setEncryptedTyping() *TL_messages_setEncryptedTyping {
	return &TL_messages_setEncryptedTyping{}
}

func (t *TL_messages_setEncryptedTyping) Encode() []byte {
	return nil
}

func (t *TL_messages_setEncryptedTyping) Decode(b []byte) {}

//====messages_readEncryptedHistory#7f4b690a====

type TL_messages_readEncryptedHistory struct {
	_peer     []byte
	_max_date []byte
}

func New_TL_messages_readEncryptedHistory() *TL_messages_readEncryptedHistory {
	return &TL_messages_readEncryptedHistory{}
}

func (t *TL_messages_readEncryptedHistory) Encode() []byte {
	return nil
}

func (t *TL_messages_readEncryptedHistory) Decode(b []byte) {}

//====messages_sendEncrypted#a9776773====

type TL_messages_sendEncrypted struct {
	_peer      []byte
	_random_id int64
	_data      []byte
}

func New_TL_messages_sendEncrypted() *TL_messages_sendEncrypted {
	return &TL_messages_sendEncrypted{}
}

func (t *TL_messages_sendEncrypted) Encode() []byte {
	return nil
}

func (t *TL_messages_sendEncrypted) Decode(b []byte) {}

//====messages_sendEncryptedFile#9a901b66====

type TL_messages_sendEncryptedFile struct {
	_peer      []byte
	_random_id int64
	_data      []byte
	_file      []byte
}

func New_TL_messages_sendEncryptedFile() *TL_messages_sendEncryptedFile {
	return &TL_messages_sendEncryptedFile{}
}

func (t *TL_messages_sendEncryptedFile) Encode() []byte {
	return nil
}

func (t *TL_messages_sendEncryptedFile) Decode(b []byte) {}

//====messages_sendEncryptedService#32d439a4====

type TL_messages_sendEncryptedService struct {
	_peer      []byte
	_random_id int64
	_data      []byte
}

func New_TL_messages_sendEncryptedService() *TL_messages_sendEncryptedService {
	return &TL_messages_sendEncryptedService{}
}

func (t *TL_messages_sendEncryptedService) Encode() []byte {
	return nil
}

func (t *TL_messages_sendEncryptedService) Decode(b []byte) {}

//====messages_receivedQueue#55a5bb66====

type TL_messages_receivedQueue struct {
	_max_qts []byte
}

func New_TL_messages_receivedQueue() *TL_messages_receivedQueue {
	return &TL_messages_receivedQueue{}
}

func (t *TL_messages_receivedQueue) Encode() []byte {
	return nil
}

func (t *TL_messages_receivedQueue) Decode(b []byte) {}

//====messages_reportEncryptedSpam#4b0c8c0f====

type TL_messages_reportEncryptedSpam struct {
	_peer []byte
}

func New_TL_messages_reportEncryptedSpam() *TL_messages_reportEncryptedSpam {
	return &TL_messages_reportEncryptedSpam{}
}

func (t *TL_messages_reportEncryptedSpam) Encode() []byte {
	return nil
}

func (t *TL_messages_reportEncryptedSpam) Decode(b []byte) {}

//====messages_readMessageContents#36a73f77====

type TL_messages_readMessageContents struct {
	_id []byte
}

func New_TL_messages_readMessageContents() *TL_messages_readMessageContents {
	return &TL_messages_readMessageContents{}
}

func (t *TL_messages_readMessageContents) Encode() []byte {
	return nil
}

func (t *TL_messages_readMessageContents) Decode(b []byte) {}

//====messages_getAllStickers#1c9618b1====

type TL_messages_getAllStickers struct {
	_hash []byte
}

func New_TL_messages_getAllStickers() *TL_messages_getAllStickers {
	return &TL_messages_getAllStickers{}
}

func (t *TL_messages_getAllStickers) Encode() []byte {
	return nil
}

func (t *TL_messages_getAllStickers) Decode(b []byte) {}

//====messages_getWebPagePreview#25223e24====

type TL_messages_getWebPagePreview struct {
	_message string
}

func New_TL_messages_getWebPagePreview() *TL_messages_getWebPagePreview {
	return &TL_messages_getWebPagePreview{}
}

func (t *TL_messages_getWebPagePreview) Encode() []byte {
	return nil
}

func (t *TL_messages_getWebPagePreview) Decode(b []byte) {}

//====messages_exportChatInvite#7d885289====

type TL_messages_exportChatInvite struct {
	_chat_id []byte
}

func New_TL_messages_exportChatInvite() *TL_messages_exportChatInvite {
	return &TL_messages_exportChatInvite{}
}

func (t *TL_messages_exportChatInvite) Encode() []byte {
	return nil
}

func (t *TL_messages_exportChatInvite) Decode(b []byte) {}

//====messages_checkChatInvite#3eadb1bb====

type TL_messages_checkChatInvite struct {
	_hash string
}

func New_TL_messages_checkChatInvite() *TL_messages_checkChatInvite {
	return &TL_messages_checkChatInvite{}
}

func (t *TL_messages_checkChatInvite) Encode() []byte {
	return nil
}

func (t *TL_messages_checkChatInvite) Decode(b []byte) {}

//====messages_importChatInvite#6c50051c====

type TL_messages_importChatInvite struct {
	_hash string
}

func New_TL_messages_importChatInvite() *TL_messages_importChatInvite {
	return &TL_messages_importChatInvite{}
}

func (t *TL_messages_importChatInvite) Encode() []byte {
	return nil
}

func (t *TL_messages_importChatInvite) Decode(b []byte) {}

//====messages_getStickerSet#2619a90e====

type TL_messages_getStickerSet struct {
	_stickerset []byte
}

func New_TL_messages_getStickerSet() *TL_messages_getStickerSet {
	return &TL_messages_getStickerSet{}
}

func (t *TL_messages_getStickerSet) Encode() []byte {
	return nil
}

func (t *TL_messages_getStickerSet) Decode(b []byte) {}

//====messages_installStickerSet#c78fe460====

type TL_messages_installStickerSet struct {
	_stickerset []byte
	_archived   []byte
}

func New_TL_messages_installStickerSet() *TL_messages_installStickerSet {
	return &TL_messages_installStickerSet{}
}

func (t *TL_messages_installStickerSet) Encode() []byte {
	return nil
}

func (t *TL_messages_installStickerSet) Decode(b []byte) {}

//====messages_uninstallStickerSet#f96e55de====

type TL_messages_uninstallStickerSet struct {
	_stickerset []byte
}

func New_TL_messages_uninstallStickerSet() *TL_messages_uninstallStickerSet {
	return &TL_messages_uninstallStickerSet{}
}

func (t *TL_messages_uninstallStickerSet) Encode() []byte {
	return nil
}

func (t *TL_messages_uninstallStickerSet) Decode(b []byte) {}

//====messages_startBot#e6df7378====

type TL_messages_startBot struct {
	_bot         []byte
	_peer        []byte
	_random_id   int64
	_start_param string
}

func New_TL_messages_startBot() *TL_messages_startBot {
	return &TL_messages_startBot{}
}

func (t *TL_messages_startBot) Encode() []byte {
	return nil
}

func (t *TL_messages_startBot) Decode(b []byte) {}

//====messages_getMessagesViews#c4c8a55d====

type TL_messages_getMessagesViews struct {
	_peer      []byte
	_id        []byte
	_increment []byte
}

func New_TL_messages_getMessagesViews() *TL_messages_getMessagesViews {
	return &TL_messages_getMessagesViews{}
}

func (t *TL_messages_getMessagesViews) Encode() []byte {
	return nil
}

func (t *TL_messages_getMessagesViews) Decode(b []byte) {}

//====messages_toggleChatAdmins#ec8bd9e1====

type TL_messages_toggleChatAdmins struct {
	_chat_id []byte
	_enabled []byte
}

func New_TL_messages_toggleChatAdmins() *TL_messages_toggleChatAdmins {
	return &TL_messages_toggleChatAdmins{}
}

func (t *TL_messages_toggleChatAdmins) Encode() []byte {
	return nil
}

func (t *TL_messages_toggleChatAdmins) Decode(b []byte) {}

//====messages_editChatAdmin#a9e69f2e====

type TL_messages_editChatAdmin struct {
	_chat_id  []byte
	_user_id  []byte
	_is_admin []byte
}

func New_TL_messages_editChatAdmin() *TL_messages_editChatAdmin {
	return &TL_messages_editChatAdmin{}
}

func (t *TL_messages_editChatAdmin) Encode() []byte {
	return nil
}

func (t *TL_messages_editChatAdmin) Decode(b []byte) {}

//====messages_migrateChat#15a3b8e3====

type TL_messages_migrateChat struct {
	_chat_id []byte
}

func New_TL_messages_migrateChat() *TL_messages_migrateChat {
	return &TL_messages_migrateChat{}
}

func (t *TL_messages_migrateChat) Encode() []byte {
	return nil
}

func (t *TL_messages_migrateChat) Decode(b []byte) {}

//====messages_searchGlobal#9e3cacb0====

type TL_messages_searchGlobal struct {
	_q           string
	_offset_date []byte
	_offset_peer []byte
	_offset_id   []byte
	_limit       []byte
}

func New_TL_messages_searchGlobal() *TL_messages_searchGlobal {
	return &TL_messages_searchGlobal{}
}

func (t *TL_messages_searchGlobal) Encode() []byte {
	return nil
}

func (t *TL_messages_searchGlobal) Decode(b []byte) {}

//====messages_reorderStickerSets#78337739====

type TL_messages_reorderStickerSets struct {
	_flags []byte
	_masks []byte
	_order []byte
}

func New_TL_messages_reorderStickerSets() *TL_messages_reorderStickerSets {
	return &TL_messages_reorderStickerSets{}
}

func (t *TL_messages_reorderStickerSets) Encode() []byte {
	return nil
}

func (t *TL_messages_reorderStickerSets) Decode(b []byte) {}

//====messages_getDocumentByHash#338e2464====

type TL_messages_getDocumentByHash struct {
	_sha256    []byte
	_size      []byte
	_mime_type string
}

func New_TL_messages_getDocumentByHash() *TL_messages_getDocumentByHash {
	return &TL_messages_getDocumentByHash{}
}

func (t *TL_messages_getDocumentByHash) Encode() []byte {
	return nil
}

func (t *TL_messages_getDocumentByHash) Decode(b []byte) {}

//====messages_searchGifs#bf9a776b====

type TL_messages_searchGifs struct {
	_q      string
	_offset []byte
}

func New_TL_messages_searchGifs() *TL_messages_searchGifs {
	return &TL_messages_searchGifs{}
}

func (t *TL_messages_searchGifs) Encode() []byte {
	return nil
}

func (t *TL_messages_searchGifs) Decode(b []byte) {}

//====messages_getSavedGifs#83bf3d52====

type TL_messages_getSavedGifs struct {
	_hash []byte
}

func New_TL_messages_getSavedGifs() *TL_messages_getSavedGifs {
	return &TL_messages_getSavedGifs{}
}

func (t *TL_messages_getSavedGifs) Encode() []byte {
	return nil
}

func (t *TL_messages_getSavedGifs) Decode(b []byte) {}

//====messages_saveGif#327a30cb====

type TL_messages_saveGif struct {
	_id     []byte
	_unsave []byte
}

func New_TL_messages_saveGif() *TL_messages_saveGif {
	return &TL_messages_saveGif{}
}

func (t *TL_messages_saveGif) Encode() []byte {
	return nil
}

func (t *TL_messages_saveGif) Decode(b []byte) {}

//====messages_getInlineBotResults#514e999d====

type TL_messages_getInlineBotResults struct {
	_flags     []byte
	_bot       []byte
	_peer      []byte
	_geo_point []byte
	_query     string
	_offset    string
}

func New_TL_messages_getInlineBotResults() *TL_messages_getInlineBotResults {
	return &TL_messages_getInlineBotResults{}
}

func (t *TL_messages_getInlineBotResults) Encode() []byte {
	return nil
}

func (t *TL_messages_getInlineBotResults) Decode(b []byte) {}

//====messages_setInlineBotResults#eb5ea206====

type TL_messages_setInlineBotResults struct {
	_flags       []byte
	_gallery     []byte
	_private     []byte
	_query_id    int64
	_results     []byte
	_cache_time  []byte
	_next_offset []byte
	_switch_pm   []byte
}

func New_TL_messages_setInlineBotResults() *TL_messages_setInlineBotResults {
	return &TL_messages_setInlineBotResults{}
}

func (t *TL_messages_setInlineBotResults) Encode() []byte {
	return nil
}

func (t *TL_messages_setInlineBotResults) Decode(b []byte) {}

//====messages_sendInlineBotResult#b16e06fe====

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

func New_TL_messages_sendInlineBotResult() *TL_messages_sendInlineBotResult {
	return &TL_messages_sendInlineBotResult{}
}

func (t *TL_messages_sendInlineBotResult) Encode() []byte {
	return nil
}

func (t *TL_messages_sendInlineBotResult) Decode(b []byte) {}

//====messages_getMessageEditData#fda68d36====

type TL_messages_getMessageEditData struct {
	_peer []byte
	_id   []byte
}

func New_TL_messages_getMessageEditData() *TL_messages_getMessageEditData {
	return &TL_messages_getMessageEditData{}
}

func (t *TL_messages_getMessageEditData) Encode() []byte {
	return nil
}

func (t *TL_messages_getMessageEditData) Decode(b []byte) {}

//====messages_editMessage#05d1b8dd====

type TL_messages_editMessage struct {
	_flags         []byte
	_no_webpage    []byte
	_stop_geo_live []byte
	_peer          []byte
	_id            []byte
	_message       []byte
	_reply_markup  []byte
	_entities      []byte
	_geo_point     []byte
}

func New_TL_messages_editMessage() *TL_messages_editMessage {
	return &TL_messages_editMessage{}
}

func (t *TL_messages_editMessage) Encode() []byte {
	return nil
}

func (t *TL_messages_editMessage) Decode(b []byte) {}

//====messages_editInlineBotMessage#b0e08243====

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

func New_TL_messages_editInlineBotMessage() *TL_messages_editInlineBotMessage {
	return &TL_messages_editInlineBotMessage{}
}

func (t *TL_messages_editInlineBotMessage) Encode() []byte {
	return nil
}

func (t *TL_messages_editInlineBotMessage) Decode(b []byte) {}

//====messages_getBotCallbackAnswer#810a9fec====

type TL_messages_getBotCallbackAnswer struct {
	_flags  []byte
	_game   []byte
	_peer   []byte
	_msg_id []byte
	_data   []byte
}

func New_TL_messages_getBotCallbackAnswer() *TL_messages_getBotCallbackAnswer {
	return &TL_messages_getBotCallbackAnswer{}
}

func (t *TL_messages_getBotCallbackAnswer) Encode() []byte {
	return nil
}

func (t *TL_messages_getBotCallbackAnswer) Decode(b []byte) {}

//====messages_setBotCallbackAnswer#d58f130a====

type TL_messages_setBotCallbackAnswer struct {
	_flags      []byte
	_alert      []byte
	_query_id   int64
	_message    []byte
	_url        []byte
	_cache_time []byte
}

func New_TL_messages_setBotCallbackAnswer() *TL_messages_setBotCallbackAnswer {
	return &TL_messages_setBotCallbackAnswer{}
}

func (t *TL_messages_setBotCallbackAnswer) Encode() []byte {
	return nil
}

func (t *TL_messages_setBotCallbackAnswer) Decode(b []byte) {}

//====messages_getPeerDialogs#2d9776b9====

type TL_messages_getPeerDialogs struct {
	_peers []byte
}

func New_TL_messages_getPeerDialogs() *TL_messages_getPeerDialogs {
	return &TL_messages_getPeerDialogs{}
}

func (t *TL_messages_getPeerDialogs) Encode() []byte {
	return nil
}

func (t *TL_messages_getPeerDialogs) Decode(b []byte) {}

//====messages_saveDraft#bc39e14b====

type TL_messages_saveDraft struct {
	_flags           []byte
	_no_webpage      []byte
	_reply_to_msg_id []byte
	_peer            []byte
	_message         string
	_entities        []byte
}

func New_TL_messages_saveDraft() *TL_messages_saveDraft {
	return &TL_messages_saveDraft{}
}

func (t *TL_messages_saveDraft) Encode() []byte {
	return nil
}

func (t *TL_messages_saveDraft) Decode(b []byte) {}

//====messages_getAllDrafts#6a3f8d65====

type TL_messages_getAllDrafts struct {
}

func New_TL_messages_getAllDrafts() *TL_messages_getAllDrafts {
	return &TL_messages_getAllDrafts{}
}

func (t *TL_messages_getAllDrafts) Encode() []byte {
	return nil
}

func (t *TL_messages_getAllDrafts) Decode(b []byte) {}

//====messages_getFeaturedStickers#2dacca4f====

type TL_messages_getFeaturedStickers struct {
	_hash []byte
}

func New_TL_messages_getFeaturedStickers() *TL_messages_getFeaturedStickers {
	return &TL_messages_getFeaturedStickers{}
}

func (t *TL_messages_getFeaturedStickers) Encode() []byte {
	return nil
}

func (t *TL_messages_getFeaturedStickers) Decode(b []byte) {}

//====messages_readFeaturedStickers#5b118126====

type TL_messages_readFeaturedStickers struct {
	_id []byte
}

func New_TL_messages_readFeaturedStickers() *TL_messages_readFeaturedStickers {
	return &TL_messages_readFeaturedStickers{}
}

func (t *TL_messages_readFeaturedStickers) Encode() []byte {
	return nil
}

func (t *TL_messages_readFeaturedStickers) Decode(b []byte) {}

//====messages_getRecentStickers#5ea192c9====

type TL_messages_getRecentStickers struct {
	_flags    []byte
	_attached []byte
	_hash     []byte
}

func New_TL_messages_getRecentStickers() *TL_messages_getRecentStickers {
	return &TL_messages_getRecentStickers{}
}

func (t *TL_messages_getRecentStickers) Encode() []byte {
	return nil
}

func (t *TL_messages_getRecentStickers) Decode(b []byte) {}

//====messages_saveRecentSticker#392718f8====

type TL_messages_saveRecentSticker struct {
	_flags    []byte
	_attached []byte
	_id       []byte
	_unsave   []byte
}

func New_TL_messages_saveRecentSticker() *TL_messages_saveRecentSticker {
	return &TL_messages_saveRecentSticker{}
}

func (t *TL_messages_saveRecentSticker) Encode() []byte {
	return nil
}

func (t *TL_messages_saveRecentSticker) Decode(b []byte) {}

//====messages_clearRecentStickers#8999602d====

type TL_messages_clearRecentStickers struct {
	_flags    []byte
	_attached []byte
}

func New_TL_messages_clearRecentStickers() *TL_messages_clearRecentStickers {
	return &TL_messages_clearRecentStickers{}
}

func (t *TL_messages_clearRecentStickers) Encode() []byte {
	return nil
}

func (t *TL_messages_clearRecentStickers) Decode(b []byte) {}

//====messages_getArchivedStickers#57f17692====

type TL_messages_getArchivedStickers struct {
	_flags     []byte
	_masks     []byte
	_offset_id int64
	_limit     []byte
}

func New_TL_messages_getArchivedStickers() *TL_messages_getArchivedStickers {
	return &TL_messages_getArchivedStickers{}
}

func (t *TL_messages_getArchivedStickers) Encode() []byte {
	return nil
}

func (t *TL_messages_getArchivedStickers) Decode(b []byte) {}

//====messages_getMaskStickers#65b8c79f====

type TL_messages_getMaskStickers struct {
	_hash []byte
}

func New_TL_messages_getMaskStickers() *TL_messages_getMaskStickers {
	return &TL_messages_getMaskStickers{}
}

func (t *TL_messages_getMaskStickers) Encode() []byte {
	return nil
}

func (t *TL_messages_getMaskStickers) Decode(b []byte) {}

//====messages_getAttachedStickers#cc5b67cc====

type TL_messages_getAttachedStickers struct {
	_media []byte
}

func New_TL_messages_getAttachedStickers() *TL_messages_getAttachedStickers {
	return &TL_messages_getAttachedStickers{}
}

func (t *TL_messages_getAttachedStickers) Encode() []byte {
	return nil
}

func (t *TL_messages_getAttachedStickers) Decode(b []byte) {}

//====messages_setGameScore#8ef8ecc0====

type TL_messages_setGameScore struct {
	_flags        []byte
	_edit_message []byte
	_force        []byte
	_peer         []byte
	_id           []byte
	_user_id      []byte
	_score        []byte
}

func New_TL_messages_setGameScore() *TL_messages_setGameScore {
	return &TL_messages_setGameScore{}
}

func (t *TL_messages_setGameScore) Encode() []byte {
	return nil
}

func (t *TL_messages_setGameScore) Decode(b []byte) {}

//====messages_setInlineGameScore#15ad9f64====

type TL_messages_setInlineGameScore struct {
	_flags        []byte
	_edit_message []byte
	_force        []byte
	_id           []byte
	_user_id      []byte
	_score        []byte
}

func New_TL_messages_setInlineGameScore() *TL_messages_setInlineGameScore {
	return &TL_messages_setInlineGameScore{}
}

func (t *TL_messages_setInlineGameScore) Encode() []byte {
	return nil
}

func (t *TL_messages_setInlineGameScore) Decode(b []byte) {}

//====messages_getGameHighScores#e822649d====

type TL_messages_getGameHighScores struct {
	_peer    []byte
	_id      []byte
	_user_id []byte
}

func New_TL_messages_getGameHighScores() *TL_messages_getGameHighScores {
	return &TL_messages_getGameHighScores{}
}

func (t *TL_messages_getGameHighScores) Encode() []byte {
	return nil
}

func (t *TL_messages_getGameHighScores) Decode(b []byte) {}

//====messages_getInlineGameHighScores#0f635e1b====

type TL_messages_getInlineGameHighScores struct {
	_id      []byte
	_user_id []byte
}

func New_TL_messages_getInlineGameHighScores() *TL_messages_getInlineGameHighScores {
	return &TL_messages_getInlineGameHighScores{}
}

func (t *TL_messages_getInlineGameHighScores) Encode() []byte {
	return nil
}

func (t *TL_messages_getInlineGameHighScores) Decode(b []byte) {}

//====messages_getCommonChats#0d0a48c4====

type TL_messages_getCommonChats struct {
	_user_id []byte
	_max_id  []byte
	_limit   []byte
}

func New_TL_messages_getCommonChats() *TL_messages_getCommonChats {
	return &TL_messages_getCommonChats{}
}

func (t *TL_messages_getCommonChats) Encode() []byte {
	return nil
}

func (t *TL_messages_getCommonChats) Decode(b []byte) {}

//====messages_getAllChats#eba80ff0====

type TL_messages_getAllChats struct {
	_except_ids []byte
}

func New_TL_messages_getAllChats() *TL_messages_getAllChats {
	return &TL_messages_getAllChats{}
}

func (t *TL_messages_getAllChats) Encode() []byte {
	return nil
}

func (t *TL_messages_getAllChats) Decode(b []byte) {}

//====messages_getWebPage#32ca8f91====

type TL_messages_getWebPage struct {
	_url  string
	_hash []byte
}

func New_TL_messages_getWebPage() *TL_messages_getWebPage {
	return &TL_messages_getWebPage{}
}

func (t *TL_messages_getWebPage) Encode() []byte {
	return nil
}

func (t *TL_messages_getWebPage) Decode(b []byte) {}

//====messages_toggleDialogPin#3289be6a====

type TL_messages_toggleDialogPin struct {
	_flags  []byte
	_pinned []byte
	_peer   []byte
}

func New_TL_messages_toggleDialogPin() *TL_messages_toggleDialogPin {
	return &TL_messages_toggleDialogPin{}
}

func (t *TL_messages_toggleDialogPin) Encode() []byte {
	return nil
}

func (t *TL_messages_toggleDialogPin) Decode(b []byte) {}

//====messages_reorderPinnedDialogs#959ff644====

type TL_messages_reorderPinnedDialogs struct {
	_flags []byte
	_force []byte
	_order []byte
}

func New_TL_messages_reorderPinnedDialogs() *TL_messages_reorderPinnedDialogs {
	return &TL_messages_reorderPinnedDialogs{}
}

func (t *TL_messages_reorderPinnedDialogs) Encode() []byte {
	return nil
}

func (t *TL_messages_reorderPinnedDialogs) Decode(b []byte) {}

//====messages_getPinnedDialogs#e254d64e====

type TL_messages_getPinnedDialogs struct {
}

func New_TL_messages_getPinnedDialogs() *TL_messages_getPinnedDialogs {
	return &TL_messages_getPinnedDialogs{}
}

func (t *TL_messages_getPinnedDialogs) Encode() []byte {
	return nil
}

func (t *TL_messages_getPinnedDialogs) Decode(b []byte) {}

//====messages_setBotShippingResults#e5f672fa====

type TL_messages_setBotShippingResults struct {
	_flags            []byte
	_query_id         int64
	_error            []byte
	_shipping_options []byte
}

func New_TL_messages_setBotShippingResults() *TL_messages_setBotShippingResults {
	return &TL_messages_setBotShippingResults{}
}

func (t *TL_messages_setBotShippingResults) Encode() []byte {
	return nil
}

func (t *TL_messages_setBotShippingResults) Decode(b []byte) {}

//====messages_setBotPrecheckoutResults#09c2dd95====

type TL_messages_setBotPrecheckoutResults struct {
	_flags    []byte
	_success  []byte
	_query_id int64
	_error    []byte
}

func New_TL_messages_setBotPrecheckoutResults() *TL_messages_setBotPrecheckoutResults {
	return &TL_messages_setBotPrecheckoutResults{}
}

func (t *TL_messages_setBotPrecheckoutResults) Encode() []byte {
	return nil
}

func (t *TL_messages_setBotPrecheckoutResults) Decode(b []byte) {}

//====messages_uploadMedia#519bc2b1====

type TL_messages_uploadMedia struct {
	_peer  []byte
	_media []byte
}

func New_TL_messages_uploadMedia() *TL_messages_uploadMedia {
	return &TL_messages_uploadMedia{}
}

func (t *TL_messages_uploadMedia) Encode() []byte {
	return nil
}

func (t *TL_messages_uploadMedia) Decode(b []byte) {}

//====messages_sendScreenshotNotification#c97df020====

type TL_messages_sendScreenshotNotification struct {
	_peer            []byte
	_reply_to_msg_id []byte
	_random_id       int64
}

func New_TL_messages_sendScreenshotNotification() *TL_messages_sendScreenshotNotification {
	return &TL_messages_sendScreenshotNotification{}
}

func (t *TL_messages_sendScreenshotNotification) Encode() []byte {
	return nil
}

func (t *TL_messages_sendScreenshotNotification) Decode(b []byte) {}

//====messages_getFavedStickers#21ce0b0e====

type TL_messages_getFavedStickers struct {
	_hash []byte
}

func New_TL_messages_getFavedStickers() *TL_messages_getFavedStickers {
	return &TL_messages_getFavedStickers{}
}

func (t *TL_messages_getFavedStickers) Encode() []byte {
	return nil
}

func (t *TL_messages_getFavedStickers) Decode(b []byte) {}

//====messages_faveSticker#b9ffc55b====

type TL_messages_faveSticker struct {
	_id     []byte
	_unfave []byte
}

func New_TL_messages_faveSticker() *TL_messages_faveSticker {
	return &TL_messages_faveSticker{}
}

func (t *TL_messages_faveSticker) Encode() []byte {
	return nil
}

func (t *TL_messages_faveSticker) Decode(b []byte) {}

//====messages_getUnreadMentions#46578472====

type TL_messages_getUnreadMentions struct {
	_peer       []byte
	_offset_id  []byte
	_add_offset []byte
	_limit      []byte
	_max_id     []byte
	_min_id     []byte
}

func New_TL_messages_getUnreadMentions() *TL_messages_getUnreadMentions {
	return &TL_messages_getUnreadMentions{}
}

func (t *TL_messages_getUnreadMentions) Encode() []byte {
	return nil
}

func (t *TL_messages_getUnreadMentions) Decode(b []byte) {}

//====messages_readMentions#0f0189d3====

type TL_messages_readMentions struct {
	_peer []byte
}

func New_TL_messages_readMentions() *TL_messages_readMentions {
	return &TL_messages_readMentions{}
}

func (t *TL_messages_readMentions) Encode() []byte {
	return nil
}

func (t *TL_messages_readMentions) Decode(b []byte) {}

//====messages_getRecentLocations#249431e2====

type TL_messages_getRecentLocations struct {
	_peer  []byte
	_limit []byte
}

func New_TL_messages_getRecentLocations() *TL_messages_getRecentLocations {
	return &TL_messages_getRecentLocations{}
}

func (t *TL_messages_getRecentLocations) Encode() []byte {
	return nil
}

func (t *TL_messages_getRecentLocations) Decode(b []byte) {}

//====messages_sendMultiMedia#2095512f====

type TL_messages_sendMultiMedia struct {
	_flags           []byte
	_silent          []byte
	_background      []byte
	_clear_draft     []byte
	_peer            []byte
	_reply_to_msg_id []byte
	_multi_media     []byte
}

func New_TL_messages_sendMultiMedia() *TL_messages_sendMultiMedia {
	return &TL_messages_sendMultiMedia{}
}

func (t *TL_messages_sendMultiMedia) Encode() []byte {
	return nil
}

func (t *TL_messages_sendMultiMedia) Decode(b []byte) {}

//====messages_uploadEncryptedFile#5057c497====

type TL_messages_uploadEncryptedFile struct {
	_peer []byte
	_file []byte
}

func New_TL_messages_uploadEncryptedFile() *TL_messages_uploadEncryptedFile {
	return &TL_messages_uploadEncryptedFile{}
}

func (t *TL_messages_uploadEncryptedFile) Encode() []byte {
	return nil
}

func (t *TL_messages_uploadEncryptedFile) Decode(b []byte) {}

//====updates_getState#edd4882a====

type TL_updates_getState struct {
}

func New_TL_updates_getState() *TL_updates_getState {
	return &TL_updates_getState{}
}

func (t *TL_updates_getState) Encode() []byte {
	return nil
}

func (t *TL_updates_getState) Decode(b []byte) {}

//====updates_getDifference#25939651====

type TL_updates_getDifference struct {
	_flags           []byte
	_pts             []byte
	_pts_total_limit []byte
	_date            []byte
	_qts             []byte
}

func New_TL_updates_getDifference() *TL_updates_getDifference {
	return &TL_updates_getDifference{}
}

func (t *TL_updates_getDifference) Encode() []byte {
	return nil
}

func (t *TL_updates_getDifference) Decode(b []byte) {}

//====updates_getChannelDifference#03173d78====

type TL_updates_getChannelDifference struct {
	_flags   []byte
	_force   []byte
	_channel []byte
	_filter  []byte
	_pts     []byte
	_limit   []byte
}

func New_TL_updates_getChannelDifference() *TL_updates_getChannelDifference {
	return &TL_updates_getChannelDifference{}
}

func (t *TL_updates_getChannelDifference) Encode() []byte {
	return nil
}

func (t *TL_updates_getChannelDifference) Decode(b []byte) {}

//====photos_updateProfilePhoto#f0bb5152====

type TL_photos_updateProfilePhoto struct {
	_id []byte
}

func New_TL_photos_updateProfilePhoto() *TL_photos_updateProfilePhoto {
	return &TL_photos_updateProfilePhoto{}
}

func (t *TL_photos_updateProfilePhoto) Encode() []byte {
	return nil
}

func (t *TL_photos_updateProfilePhoto) Decode(b []byte) {}

//====photos_uploadProfilePhoto#4f32c098====

type TL_photos_uploadProfilePhoto struct {
	_file []byte
}

func New_TL_photos_uploadProfilePhoto() *TL_photos_uploadProfilePhoto {
	return &TL_photos_uploadProfilePhoto{}
}

func (t *TL_photos_uploadProfilePhoto) Encode() []byte {
	return nil
}

func (t *TL_photos_uploadProfilePhoto) Decode(b []byte) {}

//====photos_deletePhotos#87cf7f2f====

type TL_photos_deletePhotos struct {
	_id []byte
}

func New_TL_photos_deletePhotos() *TL_photos_deletePhotos {
	return &TL_photos_deletePhotos{}
}

func (t *TL_photos_deletePhotos) Encode() []byte {
	return nil
}

func (t *TL_photos_deletePhotos) Decode(b []byte) {}

//====photos_getUserPhotos#91cd32a8====

type TL_photos_getUserPhotos struct {
	_user_id []byte
	_offset  []byte
	_max_id  int64
	_limit   []byte
}

func New_TL_photos_getUserPhotos() *TL_photos_getUserPhotos {
	return &TL_photos_getUserPhotos{}
}

func (t *TL_photos_getUserPhotos) Encode() []byte {
	return nil
}

func (t *TL_photos_getUserPhotos) Decode(b []byte) {}

//====upload_saveFilePart#b304a621====

type TL_upload_saveFilePart struct {
	_file_id   int64
	_file_part []byte
	_bytes     []byte
}

func New_TL_upload_saveFilePart() *TL_upload_saveFilePart {
	return &TL_upload_saveFilePart{}
}

func (t *TL_upload_saveFilePart) Encode() []byte {
	return nil
}

func (t *TL_upload_saveFilePart) Decode(b []byte) {}

//====upload_getFile#e3a6cfb5====

type TL_upload_getFile struct {
	_location []byte
	_offset   []byte
	_limit    []byte
}

func New_TL_upload_getFile() *TL_upload_getFile {
	return &TL_upload_getFile{}
}

func (t *TL_upload_getFile) Encode() []byte {
	return nil
}

func (t *TL_upload_getFile) Decode(b []byte) {}

//====upload_saveBigFilePart#de7b673d====

type TL_upload_saveBigFilePart struct {
	_file_id          int64
	_file_part        []byte
	_file_total_parts []byte
	_bytes            []byte
}

func New_TL_upload_saveBigFilePart() *TL_upload_saveBigFilePart {
	return &TL_upload_saveBigFilePart{}
}

func (t *TL_upload_saveBigFilePart) Encode() []byte {
	return nil
}

func (t *TL_upload_saveBigFilePart) Decode(b []byte) {}

//====upload_getWebFile#24e6818d====

type TL_upload_getWebFile struct {
	_location []byte
	_offset   []byte
	_limit    []byte
}

func New_TL_upload_getWebFile() *TL_upload_getWebFile {
	return &TL_upload_getWebFile{}
}

func (t *TL_upload_getWebFile) Encode() []byte {
	return nil
}

func (t *TL_upload_getWebFile) Decode(b []byte) {}

//====upload_getCdnFile#2000bcc3====

type TL_upload_getCdnFile struct {
	_file_token []byte
	_offset     []byte
	_limit      []byte
}

func New_TL_upload_getCdnFile() *TL_upload_getCdnFile {
	return &TL_upload_getCdnFile{}
}

func (t *TL_upload_getCdnFile) Encode() []byte {
	return nil
}

func (t *TL_upload_getCdnFile) Decode(b []byte) {}

//====upload_reuploadCdnFile#1af91c09====

type TL_upload_reuploadCdnFile struct {
	_file_token    []byte
	_request_token []byte
}

func New_TL_upload_reuploadCdnFile() *TL_upload_reuploadCdnFile {
	return &TL_upload_reuploadCdnFile{}
}

func (t *TL_upload_reuploadCdnFile) Encode() []byte {
	return nil
}

func (t *TL_upload_reuploadCdnFile) Decode(b []byte) {}

//====upload_getCdnFileHashes#f715c87b====

type TL_upload_getCdnFileHashes struct {
	_file_token []byte
	_offset     []byte
}

func New_TL_upload_getCdnFileHashes() *TL_upload_getCdnFileHashes {
	return &TL_upload_getCdnFileHashes{}
}

func (t *TL_upload_getCdnFileHashes) Encode() []byte {
	return nil
}

func (t *TL_upload_getCdnFileHashes) Decode(b []byte) {}

//====help_getConfig#c4f9186b====

type TL_help_getConfig struct {
}

func New_TL_help_getConfig() *TL_help_getConfig {
	return &TL_help_getConfig{}
}

func (t *TL_help_getConfig) Encode() []byte {
	return nil
}

func (t *TL_help_getConfig) Decode(b []byte) {}

//====help_getNearestDc#1fb33026====

type TL_help_getNearestDc struct {
}

func New_TL_help_getNearestDc() *TL_help_getNearestDc {
	return &TL_help_getNearestDc{}
}

func (t *TL_help_getNearestDc) Encode() []byte {
	return nil
}

func (t *TL_help_getNearestDc) Decode(b []byte) {}

//====help_getAppUpdate#ae2de196====

type TL_help_getAppUpdate struct {
}

func New_TL_help_getAppUpdate() *TL_help_getAppUpdate {
	return &TL_help_getAppUpdate{}
}

func (t *TL_help_getAppUpdate) Encode() []byte {
	return nil
}

func (t *TL_help_getAppUpdate) Decode(b []byte) {}

//====help_saveAppLog#6f02f748====

type TL_help_saveAppLog struct {
	_events []byte
}

func New_TL_help_saveAppLog() *TL_help_saveAppLog {
	return &TL_help_saveAppLog{}
}

func (t *TL_help_saveAppLog) Encode() []byte {
	return nil
}

func (t *TL_help_saveAppLog) Decode(b []byte) {}

//====help_getInviteText#4d392343====

type TL_help_getInviteText struct {
}

func New_TL_help_getInviteText() *TL_help_getInviteText {
	return &TL_help_getInviteText{}
}

func (t *TL_help_getInviteText) Encode() []byte {
	return nil
}

func (t *TL_help_getInviteText) Decode(b []byte) {}

//====help_getSupport#9cdf08cd====

type TL_help_getSupport struct {
}

func New_TL_help_getSupport() *TL_help_getSupport {
	return &TL_help_getSupport{}
}

func (t *TL_help_getSupport) Encode() []byte {
	return nil
}

func (t *TL_help_getSupport) Decode(b []byte) {}

//====help_getAppChangelog#9010ef6f====

type TL_help_getAppChangelog struct {
	_prev_app_version string
}

func New_TL_help_getAppChangelog() *TL_help_getAppChangelog {
	return &TL_help_getAppChangelog{}
}

func (t *TL_help_getAppChangelog) Encode() []byte {
	return nil
}

func (t *TL_help_getAppChangelog) Decode(b []byte) {}

//====help_getTermsOfService#350170f3====

type TL_help_getTermsOfService struct {
}

func New_TL_help_getTermsOfService() *TL_help_getTermsOfService {
	return &TL_help_getTermsOfService{}
}

func (t *TL_help_getTermsOfService) Encode() []byte {
	return nil
}

func (t *TL_help_getTermsOfService) Decode(b []byte) {}

//====help_setBotUpdatesStatus#ec22cfcd====

type TL_help_setBotUpdatesStatus struct {
	_pending_updates_count []byte
	_message               string
}

func New_TL_help_setBotUpdatesStatus() *TL_help_setBotUpdatesStatus {
	return &TL_help_setBotUpdatesStatus{}
}

func (t *TL_help_setBotUpdatesStatus) Encode() []byte {
	return nil
}

func (t *TL_help_setBotUpdatesStatus) Decode(b []byte) {}

//====help_getCdnConfig#52029342====

type TL_help_getCdnConfig struct {
}

func New_TL_help_getCdnConfig() *TL_help_getCdnConfig {
	return &TL_help_getCdnConfig{}
}

func (t *TL_help_getCdnConfig) Encode() []byte {
	return nil
}

func (t *TL_help_getCdnConfig) Decode(b []byte) {}

//====help_getRecentMeUrls#3dc0f114====

type TL_help_getRecentMeUrls struct {
	_referer string
}

func New_TL_help_getRecentMeUrls() *TL_help_getRecentMeUrls {
	return &TL_help_getRecentMeUrls{}
}

func (t *TL_help_getRecentMeUrls) Encode() []byte {
	return nil
}

func (t *TL_help_getRecentMeUrls) Decode(b []byte) {}

//====channels_readHistory#cc104937====

type TL_channels_readHistory struct {
	_channel []byte
	_max_id  []byte
}

func New_TL_channels_readHistory() *TL_channels_readHistory {
	return &TL_channels_readHistory{}
}

func (t *TL_channels_readHistory) Encode() []byte {
	return nil
}

func (t *TL_channels_readHistory) Decode(b []byte) {}

//====channels_deleteMessages#84c1fd4e====

type TL_channels_deleteMessages struct {
	_channel []byte
	_id      []byte
}

func New_TL_channels_deleteMessages() *TL_channels_deleteMessages {
	return &TL_channels_deleteMessages{}
}

func (t *TL_channels_deleteMessages) Encode() []byte {
	return nil
}

func (t *TL_channels_deleteMessages) Decode(b []byte) {}

//====channels_deleteUserHistory#d10dd71b====

type TL_channels_deleteUserHistory struct {
	_channel []byte
	_user_id []byte
}

func New_TL_channels_deleteUserHistory() *TL_channels_deleteUserHistory {
	return &TL_channels_deleteUserHistory{}
}

func (t *TL_channels_deleteUserHistory) Encode() []byte {
	return nil
}

func (t *TL_channels_deleteUserHistory) Decode(b []byte) {}

//====channels_reportSpam#fe087810====

type TL_channels_reportSpam struct {
	_channel []byte
	_user_id []byte
	_id      []byte
}

func New_TL_channels_reportSpam() *TL_channels_reportSpam {
	return &TL_channels_reportSpam{}
}

func (t *TL_channels_reportSpam) Encode() []byte {
	return nil
}

func (t *TL_channels_reportSpam) Decode(b []byte) {}

//====channels_getMessages#93d7b347====

type TL_channels_getMessages struct {
	_channel []byte
	_id      []byte
}

func New_TL_channels_getMessages() *TL_channels_getMessages {
	return &TL_channels_getMessages{}
}

func (t *TL_channels_getMessages) Encode() []byte {
	return nil
}

func (t *TL_channels_getMessages) Decode(b []byte) {}

//====channels_getParticipants#123e05e9====

type TL_channels_getParticipants struct {
	_channel []byte
	_filter  []byte
	_offset  []byte
	_limit   []byte
	_hash    []byte
}

func New_TL_channels_getParticipants() *TL_channels_getParticipants {
	return &TL_channels_getParticipants{}
}

func (t *TL_channels_getParticipants) Encode() []byte {
	return nil
}

func (t *TL_channels_getParticipants) Decode(b []byte) {}

//====channels_getParticipant#546dd7a6====

type TL_channels_getParticipant struct {
	_channel []byte
	_user_id []byte
}

func New_TL_channels_getParticipant() *TL_channels_getParticipant {
	return &TL_channels_getParticipant{}
}

func (t *TL_channels_getParticipant) Encode() []byte {
	return nil
}

func (t *TL_channels_getParticipant) Decode(b []byte) {}

//====channels_getChannels#0a7f6bbb====

type TL_channels_getChannels struct {
	_id []byte
}

func New_TL_channels_getChannels() *TL_channels_getChannels {
	return &TL_channels_getChannels{}
}

func (t *TL_channels_getChannels) Encode() []byte {
	return nil
}

func (t *TL_channels_getChannels) Decode(b []byte) {}

//====channels_getFullChannel#08736a09====

type TL_channels_getFullChannel struct {
	_channel []byte
}

func New_TL_channels_getFullChannel() *TL_channels_getFullChannel {
	return &TL_channels_getFullChannel{}
}

func (t *TL_channels_getFullChannel) Encode() []byte {
	return nil
}

func (t *TL_channels_getFullChannel) Decode(b []byte) {}

//====channels_createChannel#f4893d7f====

type TL_channels_createChannel struct {
	_flags     []byte
	_broadcast []byte
	_megagroup []byte
	_title     string
	_about     string
}

func New_TL_channels_createChannel() *TL_channels_createChannel {
	return &TL_channels_createChannel{}
}

func (t *TL_channels_createChannel) Encode() []byte {
	return nil
}

func (t *TL_channels_createChannel) Decode(b []byte) {}

//====channels_editAbout#13e27f1e====

type TL_channels_editAbout struct {
	_channel []byte
	_about   string
}

func New_TL_channels_editAbout() *TL_channels_editAbout {
	return &TL_channels_editAbout{}
}

func (t *TL_channels_editAbout) Encode() []byte {
	return nil
}

func (t *TL_channels_editAbout) Decode(b []byte) {}

//====channels_editAdmin#20b88214====

type TL_channels_editAdmin struct {
	_channel      []byte
	_user_id      []byte
	_admin_rights []byte
}

func New_TL_channels_editAdmin() *TL_channels_editAdmin {
	return &TL_channels_editAdmin{}
}

func (t *TL_channels_editAdmin) Encode() []byte {
	return nil
}

func (t *TL_channels_editAdmin) Decode(b []byte) {}

//====channels_editTitle#566decd0====

type TL_channels_editTitle struct {
	_channel []byte
	_title   string
}

func New_TL_channels_editTitle() *TL_channels_editTitle {
	return &TL_channels_editTitle{}
}

func (t *TL_channels_editTitle) Encode() []byte {
	return nil
}

func (t *TL_channels_editTitle) Decode(b []byte) {}

//====channels_editPhoto#f12e57c9====

type TL_channels_editPhoto struct {
	_channel []byte
	_photo   []byte
}

func New_TL_channels_editPhoto() *TL_channels_editPhoto {
	return &TL_channels_editPhoto{}
}

func (t *TL_channels_editPhoto) Encode() []byte {
	return nil
}

func (t *TL_channels_editPhoto) Decode(b []byte) {}

//====channels_checkUsername#10e6bd2c====

type TL_channels_checkUsername struct {
	_channel  []byte
	_username string
}

func New_TL_channels_checkUsername() *TL_channels_checkUsername {
	return &TL_channels_checkUsername{}
}

func (t *TL_channels_checkUsername) Encode() []byte {
	return nil
}

func (t *TL_channels_checkUsername) Decode(b []byte) {}

//====channels_updateUsername#3514b3de====

type TL_channels_updateUsername struct {
	_channel  []byte
	_username string
}

func New_TL_channels_updateUsername() *TL_channels_updateUsername {
	return &TL_channels_updateUsername{}
}

func (t *TL_channels_updateUsername) Encode() []byte {
	return nil
}

func (t *TL_channels_updateUsername) Decode(b []byte) {}

//====channels_joinChannel#24b524c5====

type TL_channels_joinChannel struct {
	_channel []byte
}

func New_TL_channels_joinChannel() *TL_channels_joinChannel {
	return &TL_channels_joinChannel{}
}

func (t *TL_channels_joinChannel) Encode() []byte {
	return nil
}

func (t *TL_channels_joinChannel) Decode(b []byte) {}

//====channels_leaveChannel#f836aa95====

type TL_channels_leaveChannel struct {
	_channel []byte
}

func New_TL_channels_leaveChannel() *TL_channels_leaveChannel {
	return &TL_channels_leaveChannel{}
}

func (t *TL_channels_leaveChannel) Encode() []byte {
	return nil
}

func (t *TL_channels_leaveChannel) Decode(b []byte) {}

//====channels_inviteToChannel#199f3a6c====

type TL_channels_inviteToChannel struct {
	_channel []byte
	_users   []byte
}

func New_TL_channels_inviteToChannel() *TL_channels_inviteToChannel {
	return &TL_channels_inviteToChannel{}
}

func (t *TL_channels_inviteToChannel) Encode() []byte {
	return nil
}

func (t *TL_channels_inviteToChannel) Decode(b []byte) {}

//====channels_exportInvite#c7560885====

type TL_channels_exportInvite struct {
	_channel []byte
}

func New_TL_channels_exportInvite() *TL_channels_exportInvite {
	return &TL_channels_exportInvite{}
}

func (t *TL_channels_exportInvite) Encode() []byte {
	return nil
}

func (t *TL_channels_exportInvite) Decode(b []byte) {}

//====channels_deleteChannel#c0111fe3====

type TL_channels_deleteChannel struct {
	_channel []byte
}

func New_TL_channels_deleteChannel() *TL_channels_deleteChannel {
	return &TL_channels_deleteChannel{}
}

func (t *TL_channels_deleteChannel) Encode() []byte {
	return nil
}

func (t *TL_channels_deleteChannel) Decode(b []byte) {}

//====channels_toggleInvites#49609307====

type TL_channels_toggleInvites struct {
	_channel []byte
	_enabled []byte
}

func New_TL_channels_toggleInvites() *TL_channels_toggleInvites {
	return &TL_channels_toggleInvites{}
}

func (t *TL_channels_toggleInvites) Encode() []byte {
	return nil
}

func (t *TL_channels_toggleInvites) Decode(b []byte) {}

//====channels_exportMessageLink#c846d22d====

type TL_channels_exportMessageLink struct {
	_channel []byte
	_id      []byte
}

func New_TL_channels_exportMessageLink() *TL_channels_exportMessageLink {
	return &TL_channels_exportMessageLink{}
}

func (t *TL_channels_exportMessageLink) Encode() []byte {
	return nil
}

func (t *TL_channels_exportMessageLink) Decode(b []byte) {}

//====channels_toggleSignatures#1f69b606====

type TL_channels_toggleSignatures struct {
	_channel []byte
	_enabled []byte
}

func New_TL_channels_toggleSignatures() *TL_channels_toggleSignatures {
	return &TL_channels_toggleSignatures{}
}

func (t *TL_channels_toggleSignatures) Encode() []byte {
	return nil
}

func (t *TL_channels_toggleSignatures) Decode(b []byte) {}

//====channels_updatePinnedMessage#a72ded52====

type TL_channels_updatePinnedMessage struct {
	_flags   []byte
	_silent  []byte
	_channel []byte
	_id      []byte
}

func New_TL_channels_updatePinnedMessage() *TL_channels_updatePinnedMessage {
	return &TL_channels_updatePinnedMessage{}
}

func (t *TL_channels_updatePinnedMessage) Encode() []byte {
	return nil
}

func (t *TL_channels_updatePinnedMessage) Decode(b []byte) {}

//====channels_getAdminedPublicChannels#8d8d82d7====

type TL_channels_getAdminedPublicChannels struct {
}

func New_TL_channels_getAdminedPublicChannels() *TL_channels_getAdminedPublicChannels {
	return &TL_channels_getAdminedPublicChannels{}
}

func (t *TL_channels_getAdminedPublicChannels) Encode() []byte {
	return nil
}

func (t *TL_channels_getAdminedPublicChannels) Decode(b []byte) {}

//====channels_editBanned#bfd915cd====

type TL_channels_editBanned struct {
	_channel       []byte
	_user_id       []byte
	_banned_rights []byte
}

func New_TL_channels_editBanned() *TL_channels_editBanned {
	return &TL_channels_editBanned{}
}

func (t *TL_channels_editBanned) Encode() []byte {
	return nil
}

func (t *TL_channels_editBanned) Decode(b []byte) {}

//====channels_getAdminLog#33ddf480====

type TL_channels_getAdminLog struct {
	_flags         []byte
	_channel       []byte
	_q             string
	_events_filter []byte
	_admins        []byte
	_max_id        int64
	_min_id        int64
	_limit         []byte
}

func New_TL_channels_getAdminLog() *TL_channels_getAdminLog {
	return &TL_channels_getAdminLog{}
}

func (t *TL_channels_getAdminLog) Encode() []byte {
	return nil
}

func (t *TL_channels_getAdminLog) Decode(b []byte) {}

//====channels_setStickers#ea8ca4f9====

type TL_channels_setStickers struct {
	_channel    []byte
	_stickerset []byte
}

func New_TL_channels_setStickers() *TL_channels_setStickers {
	return &TL_channels_setStickers{}
}

func (t *TL_channels_setStickers) Encode() []byte {
	return nil
}

func (t *TL_channels_setStickers) Decode(b []byte) {}

//====channels_readMessageContents#eab5dc38====

type TL_channels_readMessageContents struct {
	_channel []byte
	_id      []byte
}

func New_TL_channels_readMessageContents() *TL_channels_readMessageContents {
	return &TL_channels_readMessageContents{}
}

func (t *TL_channels_readMessageContents) Encode() []byte {
	return nil
}

func (t *TL_channels_readMessageContents) Decode(b []byte) {}

//====channels_deleteHistory#af369d42====

type TL_channels_deleteHistory struct {
	_channel []byte
	_max_id  []byte
}

func New_TL_channels_deleteHistory() *TL_channels_deleteHistory {
	return &TL_channels_deleteHistory{}
}

func (t *TL_channels_deleteHistory) Encode() []byte {
	return nil
}

func (t *TL_channels_deleteHistory) Decode(b []byte) {}

//====channels_togglePreHistoryHidden#eabbb94c====

type TL_channels_togglePreHistoryHidden struct {
	_channel []byte
	_enabled []byte
}

func New_TL_channels_togglePreHistoryHidden() *TL_channels_togglePreHistoryHidden {
	return &TL_channels_togglePreHistoryHidden{}
}

func (t *TL_channels_togglePreHistoryHidden) Encode() []byte {
	return nil
}

func (t *TL_channels_togglePreHistoryHidden) Decode(b []byte) {}

//====bots_sendCustomRequest#aa2769ed====

type TL_bots_sendCustomRequest struct {
	_custom_method string
	_params        []byte
}

func New_TL_bots_sendCustomRequest() *TL_bots_sendCustomRequest {
	return &TL_bots_sendCustomRequest{}
}

func (t *TL_bots_sendCustomRequest) Encode() []byte {
	return nil
}

func (t *TL_bots_sendCustomRequest) Decode(b []byte) {}

//====bots_answerWebhookJSONQuery#e6213f4d====

type TL_bots_answerWebhookJSONQuery struct {
	_query_id int64
	_data     []byte
}

func New_TL_bots_answerWebhookJSONQuery() *TL_bots_answerWebhookJSONQuery {
	return &TL_bots_answerWebhookJSONQuery{}
}

func (t *TL_bots_answerWebhookJSONQuery) Encode() []byte {
	return nil
}

func (t *TL_bots_answerWebhookJSONQuery) Decode(b []byte) {}

//====payments_getPaymentForm#99f09745====

type TL_payments_getPaymentForm struct {
	_msg_id []byte
}

func New_TL_payments_getPaymentForm() *TL_payments_getPaymentForm {
	return &TL_payments_getPaymentForm{}
}

func (t *TL_payments_getPaymentForm) Encode() []byte {
	return nil
}

func (t *TL_payments_getPaymentForm) Decode(b []byte) {}

//====payments_getPaymentReceipt#a092a980====

type TL_payments_getPaymentReceipt struct {
	_msg_id []byte
}

func New_TL_payments_getPaymentReceipt() *TL_payments_getPaymentReceipt {
	return &TL_payments_getPaymentReceipt{}
}

func (t *TL_payments_getPaymentReceipt) Encode() []byte {
	return nil
}

func (t *TL_payments_getPaymentReceipt) Decode(b []byte) {}

//====payments_validateRequestedInfo#770a8e74====

type TL_payments_validateRequestedInfo struct {
	_flags  []byte
	_save   []byte
	_msg_id []byte
	_info   []byte
}

func New_TL_payments_validateRequestedInfo() *TL_payments_validateRequestedInfo {
	return &TL_payments_validateRequestedInfo{}
}

func (t *TL_payments_validateRequestedInfo) Encode() []byte {
	return nil
}

func (t *TL_payments_validateRequestedInfo) Decode(b []byte) {}

//====payments_sendPaymentForm#2b8879b3====

type TL_payments_sendPaymentForm struct {
	_flags              []byte
	_msg_id             []byte
	_requested_info_id  []byte
	_shipping_option_id []byte
	_credentials        []byte
}

func New_TL_payments_sendPaymentForm() *TL_payments_sendPaymentForm {
	return &TL_payments_sendPaymentForm{}
}

func (t *TL_payments_sendPaymentForm) Encode() []byte {
	return nil
}

func (t *TL_payments_sendPaymentForm) Decode(b []byte) {}

//====payments_getSavedInfo#227d824b====

type TL_payments_getSavedInfo struct {
}

func New_TL_payments_getSavedInfo() *TL_payments_getSavedInfo {
	return &TL_payments_getSavedInfo{}
}

func (t *TL_payments_getSavedInfo) Encode() []byte {
	return nil
}

func (t *TL_payments_getSavedInfo) Decode(b []byte) {}

//====payments_clearSavedInfo#d83d70c1====

type TL_payments_clearSavedInfo struct {
	_flags       []byte
	_credentials []byte
	_info        []byte
}

func New_TL_payments_clearSavedInfo() *TL_payments_clearSavedInfo {
	return &TL_payments_clearSavedInfo{}
}

func (t *TL_payments_clearSavedInfo) Encode() []byte {
	return nil
}

func (t *TL_payments_clearSavedInfo) Decode(b []byte) {}

//====stickers_createStickerSet#9bd86e6a====

type TL_stickers_createStickerSet struct {
	_flags      []byte
	_masks      []byte
	_user_id    []byte
	_title      string
	_short_name string
	_stickers   []byte
}

func New_TL_stickers_createStickerSet() *TL_stickers_createStickerSet {
	return &TL_stickers_createStickerSet{}
}

func (t *TL_stickers_createStickerSet) Encode() []byte {
	return nil
}

func (t *TL_stickers_createStickerSet) Decode(b []byte) {}

//====stickers_removeStickerFromSet#f7760f51====

type TL_stickers_removeStickerFromSet struct {
	_sticker []byte
}

func New_TL_stickers_removeStickerFromSet() *TL_stickers_removeStickerFromSet {
	return &TL_stickers_removeStickerFromSet{}
}

func (t *TL_stickers_removeStickerFromSet) Encode() []byte {
	return nil
}

func (t *TL_stickers_removeStickerFromSet) Decode(b []byte) {}

//====stickers_changeStickerPosition#ffb6d4ca====

type TL_stickers_changeStickerPosition struct {
	_sticker  []byte
	_position []byte
}

func New_TL_stickers_changeStickerPosition() *TL_stickers_changeStickerPosition {
	return &TL_stickers_changeStickerPosition{}
}

func (t *TL_stickers_changeStickerPosition) Encode() []byte {
	return nil
}

func (t *TL_stickers_changeStickerPosition) Decode(b []byte) {}

//====stickers_addStickerToSet#8653febe====

type TL_stickers_addStickerToSet struct {
	_stickerset []byte
	_sticker    []byte
}

func New_TL_stickers_addStickerToSet() *TL_stickers_addStickerToSet {
	return &TL_stickers_addStickerToSet{}
}

func (t *TL_stickers_addStickerToSet) Encode() []byte {
	return nil
}

func (t *TL_stickers_addStickerToSet) Decode(b []byte) {}

//====phone_getCallConfig#55451fa9====

type TL_phone_getCallConfig struct {
}

func New_TL_phone_getCallConfig() *TL_phone_getCallConfig {
	return &TL_phone_getCallConfig{}
}

func (t *TL_phone_getCallConfig) Encode() []byte {
	return nil
}

func (t *TL_phone_getCallConfig) Decode(b []byte) {}

//====phone_requestCall#5b95b3d4====

type TL_phone_requestCall struct {
	_user_id   []byte
	_random_id []byte
	_g_a_hash  []byte
	_protocol  []byte
}

func New_TL_phone_requestCall() *TL_phone_requestCall {
	return &TL_phone_requestCall{}
}

func (t *TL_phone_requestCall) Encode() []byte {
	return nil
}

func (t *TL_phone_requestCall) Decode(b []byte) {}

//====phone_acceptCall#3bd2b4a0====

type TL_phone_acceptCall struct {
	_peer     []byte
	_g_b      []byte
	_protocol []byte
}

func New_TL_phone_acceptCall() *TL_phone_acceptCall {
	return &TL_phone_acceptCall{}
}

func (t *TL_phone_acceptCall) Encode() []byte {
	return nil
}

func (t *TL_phone_acceptCall) Decode(b []byte) {}

//====phone_confirmCall#2efe1722====

type TL_phone_confirmCall struct {
	_peer            []byte
	_g_a             []byte
	_key_fingerprint int64
	_protocol        []byte
}

func New_TL_phone_confirmCall() *TL_phone_confirmCall {
	return &TL_phone_confirmCall{}
}

func (t *TL_phone_confirmCall) Encode() []byte {
	return nil
}

func (t *TL_phone_confirmCall) Decode(b []byte) {}

//====phone_receivedCall#17d54f61====

type TL_phone_receivedCall struct {
	_peer []byte
}

func New_TL_phone_receivedCall() *TL_phone_receivedCall {
	return &TL_phone_receivedCall{}
}

func (t *TL_phone_receivedCall) Encode() []byte {
	return nil
}

func (t *TL_phone_receivedCall) Decode(b []byte) {}

//====phone_discardCall#78d413a6====

type TL_phone_discardCall struct {
	_peer          []byte
	_duration      []byte
	_reason        []byte
	_connection_id int64
}

func New_TL_phone_discardCall() *TL_phone_discardCall {
	return &TL_phone_discardCall{}
}

func (t *TL_phone_discardCall) Encode() []byte {
	return nil
}

func (t *TL_phone_discardCall) Decode(b []byte) {}

//====phone_setCallRating#1c536a34====

type TL_phone_setCallRating struct {
	_peer    []byte
	_rating  []byte
	_comment string
}

func New_TL_phone_setCallRating() *TL_phone_setCallRating {
	return &TL_phone_setCallRating{}
}

func (t *TL_phone_setCallRating) Encode() []byte {
	return nil
}

func (t *TL_phone_setCallRating) Decode(b []byte) {}

//====phone_saveCallDebug#277add7e====

type TL_phone_saveCallDebug struct {
	_peer  []byte
	_debug []byte
}

func New_TL_phone_saveCallDebug() *TL_phone_saveCallDebug {
	return &TL_phone_saveCallDebug{}
}

func (t *TL_phone_saveCallDebug) Encode() []byte {
	return nil
}

func (t *TL_phone_saveCallDebug) Decode(b []byte) {}

//====langpack_getLangPack#9ab5c58e====

type TL_langpack_getLangPack struct {
	_lang_code string
}

func New_TL_langpack_getLangPack() *TL_langpack_getLangPack {
	return &TL_langpack_getLangPack{}
}

func (t *TL_langpack_getLangPack) Encode() []byte {
	return nil
}

func (t *TL_langpack_getLangPack) Decode(b []byte) {}

//====langpack_getStrings#2e1ee318====

type TL_langpack_getStrings struct {
	_lang_code string
	_keys      []byte
}

func New_TL_langpack_getStrings() *TL_langpack_getStrings {
	return &TL_langpack_getStrings{}
}

func (t *TL_langpack_getStrings) Encode() []byte {
	return nil
}

func (t *TL_langpack_getStrings) Decode(b []byte) {}

//====langpack_getDifference#0b2e4d7d====

type TL_langpack_getDifference struct {
	_from_version []byte
}

func New_TL_langpack_getDifference() *TL_langpack_getDifference {
	return &TL_langpack_getDifference{}
}

func (t *TL_langpack_getDifference) Encode() []byte {
	return nil
}

func (t *TL_langpack_getDifference) Decode(b []byte) {}

//====langpack_getLanguages#800fd57d====

type TL_langpack_getLanguages struct {
}

func New_TL_langpack_getLanguages() *TL_langpack_getLanguages {
	return &TL_langpack_getLanguages{}
}

func (t *TL_langpack_getLanguages) Encode() []byte {
	return nil
}

func (t *TL_langpack_getLanguages) Decode(b []byte) {}
