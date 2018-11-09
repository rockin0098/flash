package mtproto

import (
	"fmt"

	"github.com/gogo/protobuf/proto"
	"github.com/golang/glog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TLRpcErrorCodes int32

const (
	// 	Error handling
	//
	// 	There will be errors when working with the API, and they must be correctly handled on the client.
	//
	// 	An error is characterized by several parameters:
	//
	// 	Error Code
	// 	Similar to HTTP status. Contains information on the type of error that occurred: for example,
	// 	a data input error, privacy error, or server error. This is a required parameter.
	//
	// 	Error Type
	// 	A string literal in the form of /[A-Z_0-9]+/, which summarizes the problem. For example, AUTH_KEY_UNREGISTERED.
	// 	This is an optional parameter.
	//
	// 	Error Description
	// 	May contain more detailed information on the error and how to resolve it,
	// 	for example: authorization required, use auth.* methods. Please note that the description text is subject to change,
	// 	one should avoid tying application logic to these messages. This is an optional parameter.
	//
	// 	Error Constructors
	// 	There should be a way to handle errors that are returned in rpc_error constructors.
	//
	// 	If an error constructor does not differentiate between type and description
	// 	but instead contains a single field called error_message (as in the example above),
	// 	it must be split into 2 components, for example, using the following regular expression: /^([A-Z_0-9]+)(: (.+))?/.
	//
	// Protobuf3 enum first is 0
	TLRpcErrorCodes_ERROR_CODE_OK TLRpcErrorCodes = 0
	// The request must be repeated, but directed to a different data center.
	//
	// FILE_MIGRATE_X: the file to be accessed is currently stored in a different data center.
	// PHONE_MIGRATE_X: the phone number a user is trying to use for authorization is associated with a different data center.
	// NETWORK_MIGRATE_X: the source IP address is associated with a different data center (for registration)
	// USER_MIGRATE_X: the user whose identity is being used to execute
	// 				   queries is associated with a different data center (for registration)
	//
	// In all these cases, the error descriptionâ€™s string literal
	// 		contains the number of the data center (instead of the X) to which the repeated query must be sent.
	// More information about redirects between data centers Â»
	//
	TLRpcErrorCodes_FILE_MIGRATE_X    TLRpcErrorCodes = 303000
	TLRpcErrorCodes_PHONE_MIGRATE_X   TLRpcErrorCodes = 303001
	TLRpcErrorCodes_NETWORK_MIGRATE_X TLRpcErrorCodes = 303002
	TLRpcErrorCodes_USER_MIGRATE_X    TLRpcErrorCodes = 303003
	TLRpcErrorCodes_ERROR_SEE_OTHER   TLRpcErrorCodes = 303
	// The query contains errors. In the event that a request was created using a form
	// and contains user generated data,
	// the user should be notified that the data must be corrected before the query is repeated.
	//
	//
	// Examples of Errors:
	// 	FIRSTNAME_INVALID: The first name is invalid
	// 	LASTNAME_INVALID: The last name is invalid
	// 	PHONE_NUMBER_INVALID: The phone number is invalid
	// 	PHONE_CODE_HASH_EMPTY: phone_code_hash is missing
	// 	PHONE_CODE_EMPTY: phone_code is missing
	// 	PHONE_CODE_EXPIRED: The confirmation code has expired
	// 	API_ID_INVALID: The api_id/api_hash combination is invalid
	// 	PHONE_NUMBER_OCCUPIED: The phone number is already in use
	// 	PHONE_NUMBER_UNOCCUPIED: The phone number is not yet being used
	// 	USERS_TOO_FEW: Not enough users (to create a chat, for example)
	// 	USERS_TOO_MUCH: The maximum number of users has been exceeded (to create a chat, for example)
	// 	TYPE_CONSTRUCTOR_INVALID: The type constructor is invalid
	// 	FILE_PART_INVALID: The file part number is invalid
	// 	FILE_PARTS_INVALID: The number of file parts is invalid
	// 	FILE_PART_Ð¥_MISSING: Part X (where X is a number) of the file is missing from storage
	// 	MD5_CHECKSUM_INVALID: The MD5 checksums do not match
	// 	PHOTO_INVALID_DIMENSIONS: The photo dimensions are invalid
	// 	FIELD_NAME_INVALID: The field with the name FIELD_NAME is invalid
	// 	FIELD_NAME_EMPTY: The field with the name FIELD_NAME is missing
	// 	MSG_WAIT_FAILED: A waiting call returned an error
	//
	TLRpcErrorCodes_FIRSTNAME_INVALID        TLRpcErrorCodes = 400000
	TLRpcErrorCodes_LASTNAME_INVALID         TLRpcErrorCodes = 400001
	TLRpcErrorCodes_PHONE_NUMBER_INVALID     TLRpcErrorCodes = 400002
	TLRpcErrorCodes_PHONE_CODE_HASH_EMPTY    TLRpcErrorCodes = 400003
	TLRpcErrorCodes_PHONE_CODE_EMPTY         TLRpcErrorCodes = 400004
	TLRpcErrorCodes_PHONE_CODE_EXPIRED       TLRpcErrorCodes = 400005
	TLRpcErrorCodes_API_ID_INVALID           TLRpcErrorCodes = 400006
	TLRpcErrorCodes_PHONE_NUMBER_OCCUPIED    TLRpcErrorCodes = 400007
	TLRpcErrorCodes_PHONE_NUMBER_UNOCCUPIED  TLRpcErrorCodes = 400008
	TLRpcErrorCodes_USERS_TOO_FEW            TLRpcErrorCodes = 400009
	TLRpcErrorCodes_USERS_TOO_MUCH           TLRpcErrorCodes = 400010
	TLRpcErrorCodes_TYPE_CONSTRUCTOR_INVALID TLRpcErrorCodes = 400011
	TLRpcErrorCodes_FILE_PART_INVALID        TLRpcErrorCodes = 400012
	TLRpcErrorCodes_FILE_PART_X_MISSING      TLRpcErrorCodes = 400013
	TLRpcErrorCodes_MD5_CHECKSUM_INVALID     TLRpcErrorCodes = 400014
	TLRpcErrorCodes_PHOTO_INVALID_DIMENSIONS TLRpcErrorCodes = 400015
	TLRpcErrorCodes_FIELD_NAME_INVALID       TLRpcErrorCodes = 400016
	TLRpcErrorCodes_FIELD_NAME_EMPTY         TLRpcErrorCodes = 400017
	TLRpcErrorCodes_MSG_WAIT_FAILED          TLRpcErrorCodes = 400018
	// android client code:
	//    if (error.code == 400 && "PARTICIPANT_VERSION_OUTDATED".equals(error.text)) {
	//        callFailed(VoIPController.ERROR_PEER_OUTDATED);
	TLRpcErrorCodes_PARTICIPANT_VERSION_OUTDATED TLRpcErrorCodes = 400019
	//
	TLRpcErrorCodes_PHONE_CODE_INVALID      TLRpcErrorCodes = 400020
	TLRpcErrorCodes_PHONE_NUMBER_BANNED     TLRpcErrorCodes = 400030
	TLRpcErrorCodes_SESSION_PASSWORD_NEEDED TLRpcErrorCodes = 400040
	// password
	TLRpcErrorCodes_CODE_INVALID          TLRpcErrorCodes = 400050
	TLRpcErrorCodes_PASSWORD_HASH_INVALID TLRpcErrorCodes = 400051
	TLRpcErrorCodes_NEW_PASSWORD_BAD      TLRpcErrorCodes = 400052
	TLRpcErrorCodes_NEW_SALT_INVALID      TLRpcErrorCodes = 400053
	TLRpcErrorCodes_EMAIL_INVALID         TLRpcErrorCodes = 400054
	TLRpcErrorCodes_EMAIL_UNCONFIRMED     TLRpcErrorCodes = 400055
	// username
	TLRpcErrorCodes_USERNAME_INVALID      TLRpcErrorCodes = 400060
	TLRpcErrorCodes_USERNAME_OCCUPIED     TLRpcErrorCodes = 400061
	TLRpcErrorCodes_USERNAMES_UNAVAILABLE TLRpcErrorCodes = 400062
	// chat
	TLRpcErrorCodes_CHAT_ID_INVALID         TLRpcErrorCodes = 400070
	TLRpcErrorCodes_CHAT_NOT_MODIFIED       TLRpcErrorCodes = 400071
	TLRpcErrorCodes_PARTICIPANT_NOT_EXISTS  TLRpcErrorCodes = 400072
	TLRpcErrorCodes_NO_EDIT_CHAT_PERMISSION TLRpcErrorCodes = 400073
	//
	TLRpcErrorCodes_BAD_REQUEST TLRpcErrorCodes = 400
	// There was an unauthorized attempt to use functionality available only to authorized users.
	//
	// Examples of Errors:
	// 	AUTH_KEY_UNREGISTERED: The key is not registered in the system
	// 	AUTH_KEY_INVALID: The key is invalid
	// 	USER_DEACTIVATED: The user has been deleted/deactivated
	// 	SESSION_REVOKED: The authorization has been invalidated, because of the user terminating all sessions
	// 	SESSION_EXPIRED: The authorization has expired
	// 	ACTIVE_USER_REQUIRED: The method is only available to already activated users
	// 	AUTH_KEY_PERM_EMPTY: The method is unavailble for temporary authorization key, not bound to permanent
	//
	TLRpcErrorCodes_AUTH_KEY_UNREGISTERED TLRpcErrorCodes = 401000
	TLRpcErrorCodes_AUTH_KEY_INVALID      TLRpcErrorCodes = 401001
	TLRpcErrorCodes_USER_DEACTIVATED      TLRpcErrorCodes = 401002
	TLRpcErrorCodes_SESSION_REVOKED       TLRpcErrorCodes = 401003
	TLRpcErrorCodes_SESSION_EXPIRED       TLRpcErrorCodes = 401004
	TLRpcErrorCodes_ACTIVE_USER_REQUIRED  TLRpcErrorCodes = 401005
	TLRpcErrorCodes_AUTH_KEY_PERM_EMPTY   TLRpcErrorCodes = 401006
	// Only a small portion of the API methods are available to unauthorized users:
	//
	// auth.sendCode
	// auth.sendCall
	// auth.checkPhone
	// auth.signUp
	// auth.signIn
	// auth.importAuthorization
	// help.getConfig
	// help.getNearestDc
	//
	// Other methods will result in an error: 401 UNAUTHORIZED.
	TLRpcErrorCodes_UNAUTHORIZED TLRpcErrorCodes = 401
	// Privacy violation. For example, an attempt to write a message to someone who has blacklisted the current user.
	//
	//
	// android client code:
	//    } else if(error.code==403 && "USER_PRIVACY_RESTRICTED".equals(error.text)){
	//        callFailed(VoIPController.ERROR_PRIVACY);
	TLRpcErrorCodes_USER_PRIVACY_RESTRICTED TLRpcErrorCodes = 403000
	TLRpcErrorCodes_FORBIDDEN               TLRpcErrorCodes = 403
	// 406
	// android client code:
	// }else if(error.code==406){
	//     callFailed(VoIPController.ERROR_LOCALIZED);
	TLRpcErrorCodes_ERROR_LOCALIZED TLRpcErrorCodes = 406000
	TLRpcErrorCodes_LOCALIZED       TLRpcErrorCodes = 406
	// The maximum allowed number of attempts to invoke the given method with the given input parameters has been exceeded.
	// For example, in an attempt to request a large number of text messages (SMS) for the same phone number.
	//
	// Error Example:
	// FLOOD_WAIT_X: A wait of X seconds is required (where X is a number)
	//
	TLRpcErrorCodes_FLOOD_WAIT_X TLRpcErrorCodes = 420000
	// PEER_FLOOD
	// FLOOD_WAIT
	TLRpcErrorCodes_FLOOD TLRpcErrorCodes = 420
	// An internal server error occurred while a request was being processed;
	// for example, there was a disruption while accessing a database or file storage.
	//
	// If a client receives a 500 error, or you believe this error should not have occurred,
	// please collect as much information as possible about the query and error and send it to the developers.
	TLRpcErrorCodes_INTERNAL              TLRpcErrorCodes = 500
	TLRpcErrorCodes_INTERNAL_SERVER_ERROR TLRpcErrorCodes = 500000
	// If a server returns an error with a code other than the ones listed above,
	// it may be considered the same as a 500 error and treated as an internal server error.
	//
	TLRpcErrorCodes_OTHER TLRpcErrorCodes = 501
	//    // OFFSET_INVALID
	//    // RETRY_LIMIT
	//    // FILE_TOKEN_INVALID
	//    // REQUEST_TOKEN_INVALID
	//
	//    // CHANNEL_PRIVATE
	//    // CHANNEL_PUBLIC_GROUP_NA
	//    // USER_BANNED_IN_CHANNEL
	//
	//
	//    // MESSAGE_NOT_MODIFIED
	//
	//    // USERS_TOO_MUCH
	//
	//    // -1000
	//
	//    /////////////////////////////////////////////////////////////
	//     // android client code:
	//       } else if (request instanceof TLRPC.TL_auth_resendCode) {
	//        if (error.text.contains("PHONE_NUMBER_INVALID")) {
	//            showSimpleAlert(fragment, LocaleController.getString("InvalidPhoneNumber", R.string.InvalidPhoneNumber));
	//        } else if (error.text.contains("PHONE_CODE_EMPTY") || error.text.contains("PHONE_CODE_INVALID")) {
	//            showSimpleAlert(fragment, LocaleController.getString("InvalidCode", R.string.InvalidCode));
	//        } else if (error.text.contains("PHONE_CODE_EXPIRED")) {
	//            showSimpleAlert(fragment, LocaleController.getString("CodeExpired", R.string.CodeExpired));
	//        } else if (error.text.startsWith("FLOOD_WAIT")) {
	//            showSimpleAlert(fragment, LocaleController.getString("FloodWait", R.string.FloodWait));
	//        } else if (error.code != -1000) {
	//            showSimpleAlert(fragment, LocaleController.getString("ErrorOccurred", R.string.ErrorOccurred) + "\n" + error.text);
	//        }
	//
	//     /////////////////////////////////////////////////////////////
	//        } else if (request instanceof TLRPC.TL_updateUserName) {
	//            switch (error.text) {
	//                case "USERNAME_INVALID":
	//                    showSimpleAlert(fragment, LocaleController.getString("UsernameInvalid", R.string.UsernameInvalid));
	//                    break;
	//                case "USERNAME_OCCUPIED":
	//                    showSimpleAlert(fragment, LocaleController.getString("UsernameInUse", R.string.UsernameInUse));
	//                    break;
	//                case "USERNAMES_UNAVAILABLE":
	//                    showSimpleAlert(fragment, LocaleController.getString("FeatureUnavailable", R.string.FeatureUnavailable));
	//                    break;
	//                default:
	//                    showSimpleAlert(fragment, LocaleController.getString("ErrorOccurred", R.string.ErrorOccurred));
	//                    break;
	//            }
	//
	//     /////////////////////////////////////////////////////////////
	//            } else if (request instanceof TLRPC.TL_payments_sendPaymentForm) {
	//            switch (error.text) {
	//                case "BOT_PRECHECKOUT_FAILED":
	//                    showSimpleToast(fragment, LocaleController.getString("PaymentPrecheckoutFailed", R.string.PaymentPrecheckoutFailed));
	//                    break;
	//                case "PAYMENT_FAILED":
	//                    showSimpleToast(fragment, LocaleController.getString("PaymentFailed", R.string.PaymentFailed));
	//                    break;
	//                default:
	//                    showSimpleToast(fragment, error.text);
	//                    break;
	//            }
	//        } else if (request instanceof TLRPC.TL_payments_validateRequestedInfo) {
	//            switch (error.text) {
	//                case "SHIPPING_NOT_AVAILABLE":
	//                    showSimpleToast(fragment, LocaleController.getString("PaymentNoShippingMethod", R.string.PaymentNoShippingMethod));
	//                    break;
	//                default:
	//                    showSimpleToast(fragment, error.text);
	//                    break;
	//            }
	//        }
	//
	//     /////////////////////////////////////////////////////////////
	//
	//        } else {
	//            if (error.text.equals("2FA_RECENT_CONFIRM")) {
	//                needShowAlert(LocaleController.getString("AppName", R.string.AppName), LocaleController.getString("ResetAccountCancelledAlert", R.string.ResetAccountCancelledAlert));
	//            } else if (error.text.startsWith("2FA_CONFIRM_WAIT_")) {
	//                Bundle params = new Bundle();
	//                params.putString("phoneFormated", requestPhone);
	//                params.putString("phoneHash", phoneHash);
	//                params.putString("code", phoneCode);
	//                params.putInt("startTime", ConnectionsManager.getInstance().getCurrentTime());
	//                params.putInt("waitTime", Utilities.parseInt(error.text.replace("2FA_CONFIRM_WAIT_", "")));
	//                setPage(8, true, params, false);
	//            } else {
	TLRpcErrorCodes_OTHER2 TLRpcErrorCodes = 502
	// db error
	TLRpcErrorCodes_DBERR      TLRpcErrorCodes = 600
	TLRpcErrorCodes_DBERR_SQL  TLRpcErrorCodes = 600000
	TLRpcErrorCodes_DBERR_CONN TLRpcErrorCodes = 600001
	// db error
	TLRpcErrorCodes_NOTRETURN_CLIENT TLRpcErrorCodes = 700
)

var TLRpcErrorCodes_name = map[int32]string{
	0:      "ERROR_CODE_OK",
	303000: "FILE_MIGRATE_X",
	303001: "PHONE_MIGRATE_X",
	303002: "NETWORK_MIGRATE_X",
	303003: "USER_MIGRATE_X",
	303:    "ERROR_SEE_OTHER",
	400000: "FIRSTNAME_INVALID",
	400001: "LASTNAME_INVALID",
	400002: "PHONE_NUMBER_INVALID",
	400003: "PHONE_CODE_HASH_EMPTY",
	400004: "PHONE_CODE_EMPTY",
	400005: "PHONE_CODE_EXPIRED",
	400006: "API_ID_INVALID",
	400007: "PHONE_NUMBER_OCCUPIED",
	400008: "PHONE_NUMBER_UNOCCUPIED",
	400009: "USERS_TOO_FEW",
	400010: "USERS_TOO_MUCH",
	400011: "TYPE_CONSTRUCTOR_INVALID",
	400012: "FILE_PART_INVALID",
	400013: "FILE_PART_X_MISSING",
	400014: "MD5_CHECKSUM_INVALID",
	400015: "PHOTO_INVALID_DIMENSIONS",
	400016: "FIELD_NAME_INVALID",
	400017: "FIELD_NAME_EMPTY",
	400018: "MSG_WAIT_FAILED",
	400019: "PARTICIPANT_VERSION_OUTDATED",
	400020: "PHONE_CODE_INVALID",
	400030: "PHONE_NUMBER_BANNED",
	400040: "SESSION_PASSWORD_NEEDED",
	400050: "CODE_INVALID",
	400051: "PASSWORD_HASH_INVALID",
	400052: "NEW_PASSWORD_BAD",
	400053: "NEW_SALT_INVALID",
	400054: "EMAIL_INVALID",
	400055: "EMAIL_UNCONFIRMED",
	400060: "USERNAME_INVALID",
	400061: "USERNAME_OCCUPIED",
	400062: "USERNAMES_UNAVAILABLE",
	400070: "CHAT_ID_INVALID",
	400071: "CHAT_NOT_MODIFIED",
	400072: "PARTICIPANT_NOT_EXISTS",
	400073: "NO_EDIT_CHAT_PERMISSION",
	400:    "BAD_REQUEST",
	401000: "AUTH_KEY_UNREGISTERED",
	401001: "AUTH_KEY_INVALID",
	401002: "USER_DEACTIVATED",
	401003: "SESSION_REVOKED",
	401004: "SESSION_EXPIRED",
	401005: "ACTIVE_USER_REQUIRED",
	401006: "AUTH_KEY_PERM_EMPTY",
	401:    "UNAUTHORIZED",
	403000: "USER_PRIVACY_RESTRICTED",
	403:    "FORBIDDEN",
	406000: "ERROR_LOCALIZED",
	406:    "LOCALIZED",
	420000: "FLOOD_WAIT_X",
	420:    "FLOOD",
	500:    "INTERNAL",
	500000: "INTERNAL_SERVER_ERROR",
	501:    "OTHER",
	502:    "OTHER2",
	600:    "DBERR",
	600000: "DBERR_SQL",
	600001: "DBERR_CONN",
	700:    "NOTRETURN_CLIENT",
}
var TLRpcErrorCodes_value = map[string]int32{
	"ERROR_CODE_OK":                0,
	"FILE_MIGRATE_X":               303000,
	"PHONE_MIGRATE_X":              303001,
	"NETWORK_MIGRATE_X":            303002,
	"USER_MIGRATE_X":               303003,
	"ERROR_SEE_OTHER":              303,
	"FIRSTNAME_INVALID":            400000,
	"LASTNAME_INVALID":             400001,
	"PHONE_NUMBER_INVALID":         400002,
	"PHONE_CODE_HASH_EMPTY":        400003,
	"PHONE_CODE_EMPTY":             400004,
	"PHONE_CODE_EXPIRED":           400005,
	"API_ID_INVALID":               400006,
	"PHONE_NUMBER_OCCUPIED":        400007,
	"PHONE_NUMBER_UNOCCUPIED":      400008,
	"USERS_TOO_FEW":                400009,
	"USERS_TOO_MUCH":               400010,
	"TYPE_CONSTRUCTOR_INVALID":     400011,
	"FILE_PART_INVALID":            400012,
	"FILE_PART_X_MISSING":          400013,
	"MD5_CHECKSUM_INVALID":         400014,
	"PHOTO_INVALID_DIMENSIONS":     400015,
	"FIELD_NAME_INVALID":           400016,
	"FIELD_NAME_EMPTY":             400017,
	"MSG_WAIT_FAILED":              400018,
	"PARTICIPANT_VERSION_OUTDATED": 400019,
	"PHONE_CODE_INVALID":           400020,
	"PHONE_NUMBER_BANNED":          400030,
	"SESSION_PASSWORD_NEEDED":      400040,
	"CODE_INVALID":                 400050,
	"PASSWORD_HASH_INVALID":        400051,
	"NEW_PASSWORD_BAD":             400052,
	"NEW_SALT_INVALID":             400053,
	"EMAIL_INVALID":                400054,
	"EMAIL_UNCONFIRMED":            400055,
	"USERNAME_INVALID":             400060,
	"USERNAME_OCCUPIED":            400061,
	"USERNAMES_UNAVAILABLE":        400062,
	"CHAT_ID_INVALID":              400070,
	"CHAT_NOT_MODIFIED":            400071,
	"PARTICIPANT_NOT_EXISTS":       400072,
	"NO_EDIT_CHAT_PERMISSION":      400073,
	"BAD_REQUEST":                  400,
	"AUTH_KEY_UNREGISTERED":        401000,
	"AUTH_KEY_INVALID":             401001,
	"USER_DEACTIVATED":             401002,
	"SESSION_REVOKED":              401003,
	"SESSION_EXPIRED":              401004,
	"ACTIVE_USER_REQUIRED":         401005,
	"AUTH_KEY_PERM_EMPTY":          401006,
	"UNAUTHORIZED":                 401,
	"USER_PRIVACY_RESTRICTED":      403000,
	"FORBIDDEN":                    403,
	"ERROR_LOCALIZED":              406000,
	"LOCALIZED":                    406,
	"FLOOD_WAIT_X":                 420000,
	"FLOOD":                        420,
	"INTERNAL":                     500,
	"INTERNAL_SERVER_ERROR":        500000,
	"OTHER":                        501,
	"OTHER2":                       502,
	"DBERR":                        600,
	"DBERR_SQL":                    600000,
	"DBERR_CONN":                   600001,
	"NOTRETURN_CLIENT":             700,
}

func (x TLRpcErrorCodes) String() string {
	return proto.EnumName(TLRpcErrorCodes_name, int32(x))
}

// FILE_MIGRATE_X = 303000;
// PHONE_MIGRATE_X = 303001;
// NETWORK_MIGRATE_X = 303002;
// USER_MIGRATE_X = 303003;
//
// ERROR_SEE_OTHER code has _X is dc number, We use custom NewXXXX()
func NewFileMigrateX(dc int32, message string) *TL_rpc_error {
	return &TL_rpc_error{
		M_error_code:    int32(TLRpcErrorCodes_ERROR_SEE_OTHER),
		M_error_message: fmt.Sprintf("FILE_MIGRATE_%d: %s", dc, message),
	}
}

func NewFileMigrateX2(dc int) *TL_rpc_error {
	return &TL_rpc_error{
		M_error_code:    int32(TLRpcErrorCodes_ERROR_SEE_OTHER),
		M_error_message: fmt.Sprintf("FILE_MIGRATE_%d", dc),
	}
}

func NewPhoneMigrateX(dc int32, message string) *TL_rpc_error {
	return &TL_rpc_error{
		M_error_code:    int32(TLRpcErrorCodes_ERROR_SEE_OTHER),
		M_error_message: fmt.Sprintf("PHONE_MIGRATE_%d: %s", dc, message),
	}
}

func NewPhoneMigrateX2(dc int) *TL_rpc_error {
	return &TL_rpc_error{
		M_error_code:    int32(TLRpcErrorCodes_ERROR_SEE_OTHER),
		M_error_message: fmt.Sprintf("PHONE_MIGRATE_%d", dc),
	}
}

func NewNetworkMigrateX(dc int32, message string) *TL_rpc_error {
	return &TL_rpc_error{
		M_error_code:    int32(TLRpcErrorCodes_ERROR_SEE_OTHER),
		M_error_message: fmt.Sprintf("NETWORK_MIGRATE_%d: %s", dc, message),
	}
}

func NewNetworkMigrateX2(dc int) *TL_rpc_error {
	return &TL_rpc_error{
		M_error_code:    int32(TLRpcErrorCodes_ERROR_SEE_OTHER),
		M_error_message: fmt.Sprintf("NETWORK_MIGRATE_%d", dc),
	}
}

func NewUserMigrateX(dc int32, message string) *TL_rpc_error {
	return &TL_rpc_error{
		M_error_code:    int32(TLRpcErrorCodes_ERROR_SEE_OTHER),
		M_error_message: fmt.Sprintf("USER_MIGRATE_%d: %s", dc, message),
	}
}

func NewUserMigrateX2(dc int32) *TL_rpc_error {
	return &TL_rpc_error{
		M_error_code:    int32(TLRpcErrorCodes_ERROR_SEE_OTHER),
		M_error_message: fmt.Sprintf("USER_MIGRATE_%d", dc),
	}
}

// FLOOD_WAIT_X: A wait of X seconds is required (where X is a number)
//
func NewFloodWaitX(second int32, message string) *TL_rpc_error {
	return &TL_rpc_error{
		M_error_code:    int32(TLRpcErrorCodes_FLOOD),
		M_error_message: fmt.Sprintf("FLOOD_WAIT_%d: %s", second, message),
	}
}

func NewFloodWaitX2(second int) *TL_rpc_error {
	return &TL_rpc_error{
		M_error_code:    int32(TLRpcErrorCodes_FLOOD),
		M_error_message: fmt.Sprintf("FLOOD_WAIT_%d", second),
	}
}

// normal code NewXXX
func NewRpcError(code int32, message string) (err *TL_rpc_error) {
	if name, ok := TLRpcErrorCodes_name[int32(code)]; ok {
		if code <= int32(TLRpcErrorCodes_OTHER2) {
			err = &TL_rpc_error{
				M_error_code:    code,
				M_error_message: fmt.Sprintf("%s: %s", name, message),
			}
		} else {
			switch code {
			// Not
			case int32(TLRpcErrorCodes_FILE_MIGRATE_X),
				int32(TLRpcErrorCodes_NETWORK_MIGRATE_X),
				int32(TLRpcErrorCodes_PHONE_MIGRATE_X),
				int32(TLRpcErrorCodes_USER_MIGRATE_X):
				err = &TL_rpc_error{
					M_error_code:    int32(TLRpcErrorCodes_OTHER2),
					M_error_message: fmt.Sprintf("INTERNAL_SERVER_ERROR: Not invoke NewRpcError(%s), please use New%s(dc, %s), ", name, name, message),
				}
				glog.Error(err)

			case int32(TLRpcErrorCodes_FLOOD_WAIT_X):
				err = &TL_rpc_error{
					M_error_code:    int32(TLRpcErrorCodes_FLOOD),
					M_error_message: fmt.Sprintf("FLOOD_WAIT_%s: %s", name, name),
				}
				glog.Error(err)
			default:
				code2 := code / 1000
				if code2 == 0 {
					code2 = code
				}

				err = &TL_rpc_error{
					// subcode = code * 1000 + i
					M_error_code:    int32(code2),
					M_error_message: name,
				}
			}
		}
	} else {
		err = &TL_rpc_error{
			// subcode = code * 10000 + i
			M_error_code:    int32(TLRpcErrorCodes_INTERNAL),
			M_error_message: fmt.Sprintf("INTERNAL_SERVER_ERROR: code = %d, message = %s", code, message),
		}
	}

	return
}

// normal code NewXXX
func NewRpcError2(code TLRpcErrorCodes) (err *TL_rpc_error) {
	if name, ok := TLRpcErrorCodes_name[int32(code)]; ok {
		if code <= TLRpcErrorCodes_OTHER2 {
			err = &TL_rpc_error{
				M_error_code:    int32(code),
				M_error_message: name,
			}
		} else {
			switch code {
			// Not
			case TLRpcErrorCodes_FILE_MIGRATE_X,
				TLRpcErrorCodes_NETWORK_MIGRATE_X,
				TLRpcErrorCodes_PHONE_MIGRATE_X,
				TLRpcErrorCodes_USER_MIGRATE_X:
				err = &TL_rpc_error{
					M_error_code:    int32(TLRpcErrorCodes_OTHER2),
					M_error_message: fmt.Sprintf("INTERNAL_SERVER_ERROR: Not invoke NewRpcError(%s), please use New%s(dc), ", name, name),
				}
				glog.Fatal(err)
			case TLRpcErrorCodes_FLOOD_WAIT_X:
				err = &TL_rpc_error{
					M_error_code:    int32(TLRpcErrorCodes_FLOOD),
					M_error_message: fmt.Sprintf("INTERNAL_SERVER_ERROR: Not invoke NewRpcError(%s), please use NewFloodWaitX2(seconds), ", name),
				}
				glog.Error(err)
			default:
				code2 := code / 1000
				if code2 == 0 {
					code2 = code
				}

				err = &TL_rpc_error{
					// subcode = code * 1000 + i
					M_error_code:    int32(code2),
					M_error_message: name,
				}
			}
		}
	} else {
		err = &TL_rpc_error{
			// subcode = code * 10000 + i
			M_error_code:    int32(TLRpcErrorCodes_INTERNAL),
			M_error_message: "INTERNAL_SERVER_ERROR",
		}
	}

	return
}

// Impl error interface
func (e *TL_rpc_error) Error() string {
	return fmt.Sprintf("rpc error: code = %d desc = %s", e.Get_error_code(), e.Get_error_message())
}

// Impl error interface
func (e *TL_rpc_error) ToGrpcStatus() *status.Status {
	return status.New(codes.Internal, e.Error())
}

/*
// Impl error interface
func (e *TL_rpc_error) ToMetadata() (metadata.MD) {
	// return status.New(codes.Internal, e.Error())
	if name2, ok := TLRpcErrorCodes_name[e.ErrorCode]; ok {
		return metadata.Pairs(
			"rpc_error_code", name2,
			"rpc_error_message", e.ErrorMessage)
	}

	return metadata.Pairs(
		"rpc_error_code", "OTHER2",
		"rpc_error_message", fmt.Sprintf("INTERNAL_SERVER_ERROR: %s", e.ErrorMessage))
}

func NewRpcErrorFromMetadata(md metadata.MD) (*TL_rpc_error, error) {
	e := &TL_rpc_error{}

	if v, ok := getFirstKeyVal(md, "rpc_error_code"); ok {
		if code, ok := TLRpcErrorCodes_value[v]; !ok {
			return nil, fmt.Errorf("Invalid rpc_error_code: %s", v)
		} else {
			e.ErrorCode = code
		}
	} else {
		return nil, fmt.Errorf("Not found metadata's key: rpc_error_code")
	}

	if v, ok := getFirstKeyVal(md, "rpc_error_message"); !ok {
		e.ErrorMessage = v
	} else {
		return nil, fmt.Errorf("Not found metadata's key: rpc_error_message")
	}

	return e, nil
}
*/
