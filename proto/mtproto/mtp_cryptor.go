package mtproto

import (
	"math/big"

	"github.com/rockin0098/flash/base/crypto"
)

const (
	SHA_DIGEST_LENGTH = 20
)

var (
	// headerRpcMetadata = "auth_key_metadata"

	// TODO(@benqi): 预先计算出fingerprint
	// 这里直接使用了0xc3b42b026ce86b21
	// 计算方法
	// https://github.com/nebulaim/telegramd/blob/27362f60d2b8db1b6a51c8ab1c9a74b786124001/access/auth_key/testing/fingerprint.go
	// fingerprint uint64 = 12240908862933197005
	fingerprint uint64 = 1871186702420581553

	// TODO(@benqi): 使用算法生成PQ
	// 这里直接指定了PQ值: {0x17, 0xED, 0x48, 0x94, 0x1A, 0x08, 0xF9, 0x81}
	pq = string([]byte{0x17, 0xED, 0x48, 0x94, 0x1A, 0x08, 0xF9, 0x81})

	// TODO(@benqi): 直接指定了p和q
	p = []byte{0x49, 0x4C, 0x55, 0x3B}
	q = []byte{0x53, 0x91, 0x10, 0x73}

	// TODO(@benqi): 直接指定了dh2048_p和dh2048_g!!!
	// andriod client 指定的good prime
	//
	// static const char *goodPrime = "
	//
	// c71caeb9c6b1c9048e6c522f
	// 70f13f73980d40238e3e21c1
	// 4934d037563d930f48198a0a
	// a7c14058229493d22530f4db
	// fa336f6e0ac925139543aed4
	// 4cce7c3720fd51f69458705a
	// c68cd4fe6b6b13abdc974651
	// 2969328454f18faf8c595f64
	// 2477fe96bb2a941d5bcd1d4a
	// c8cc49880708fa9b378e3c4f
	// 3a9060bee67cf9a4a4a69581
	// 1051907e162753b56b0f6b41
	// 0dba74d8a84b2a14b3144e0e
	// f1284754fd17ed950d5965b4
	// b9dd46582db1178d169c6bc4
	// 65b0d6ff9ca3928fef5b9ae4
	// e418fc15e83ebea0f87fa9ff
	// 5eed70050ded2849f47bf959
	// d956850ce929851f0d8115f6
	// 35b105ee2e4e15d04b2454bf
	// 6f4fadf034b10403119cd8e3
	// b92fcc5b";
	//
	dh2048_p = []byte{
		0xc7, 0x1c, 0xae, 0xb9, 0xc6, 0xb1, 0xc9, 0x04, 0x8e, 0x6c, 0x52, 0x2f,
		0x70, 0xf1, 0x3f, 0x73, 0x98, 0x0d, 0x40, 0x23, 0x8e, 0x3e, 0x21, 0xc1,
		0x49, 0x34, 0xd0, 0x37, 0x56, 0x3d, 0x93, 0x0f, 0x48, 0x19, 0x8a, 0x0a,
		0xa7, 0xc1, 0x40, 0x58, 0x22, 0x94, 0x93, 0xd2, 0x25, 0x30, 0xf4, 0xdb,
		0xfa, 0x33, 0x6f, 0x6e, 0x0a, 0xc9, 0x25, 0x13, 0x95, 0x43, 0xae, 0xd4,
		0x4c, 0xce, 0x7c, 0x37, 0x20, 0xfd, 0x51, 0xf6, 0x94, 0x58, 0x70, 0x5a,
		0xc6, 0x8c, 0xd4, 0xfe, 0x6b, 0x6b, 0x13, 0xab, 0xdc, 0x97, 0x46, 0x51,
		0x29, 0x69, 0x32, 0x84, 0x54, 0xf1, 0x8f, 0xaf, 0x8c, 0x59, 0x5f, 0x64,
		0x24, 0x77, 0xfe, 0x96, 0xbb, 0x2a, 0x94, 0x1d, 0x5b, 0xcd, 0x1d, 0x4a,
		0xc8, 0xcc, 0x49, 0x88, 0x07, 0x08, 0xfa, 0x9b, 0x37, 0x8e, 0x3c, 0x4f,
		0x3a, 0x90, 0x60, 0xbe, 0xe6, 0x7c, 0xf9, 0xa4, 0xa4, 0xa6, 0x95, 0x81,
		0x10, 0x51, 0x90, 0x7e, 0x16, 0x27, 0x53, 0xb5, 0x6b, 0x0f, 0x6b, 0x41,
		0x0d, 0xba, 0x74, 0xd8, 0xa8, 0x4b, 0x2a, 0x14, 0xb3, 0x14, 0x4e, 0x0e,
		0xf1, 0x28, 0x47, 0x54, 0xfd, 0x17, 0xed, 0x95, 0x0d, 0x59, 0x65, 0xb4,
		0xb9, 0xdd, 0x46, 0x58, 0x2d, 0xb1, 0x17, 0x8d, 0x16, 0x9c, 0x6b, 0xc4,
		0x65, 0xb0, 0xd6, 0xff, 0x9c, 0xa3, 0x92, 0x8f, 0xef, 0x5b, 0x9a, 0xe4,
		0xe4, 0x18, 0xfc, 0x15, 0xe8, 0x3e, 0xbe, 0xa0, 0xf8, 0x7f, 0xa9, 0xff,
		0x5e, 0xed, 0x70, 0x05, 0x0d, 0xed, 0x28, 0x49, 0xf4, 0x7b, 0xf9, 0x59,
		0xd9, 0x56, 0x85, 0x0c, 0xe9, 0x29, 0x85, 0x1f, 0x0d, 0x81, 0x15, 0xf6,
		0x35, 0xb1, 0x05, 0xee, 0x2e, 0x4e, 0x15, 0xd0, 0x4b, 0x24, 0x54, 0xbf,
		0x6f, 0x4f, 0xad, 0xf0, 0x34, 0xb1, 0x04, 0x03, 0x11, 0x9c, 0xd8, 0xe3,
		0xb9, 0x2f, 0xcc, 0x5b,
	}

	dh2048_g = []byte{0x02}
)

type MTProtoCryptor struct {
	Fingerprint    uint64
	PQ             string
	P              []byte
	Q              []byte
	DH2048_P       []byte
	DH2048_G       []byte
	BigIntDH2048_P *big.Int
	BigIntDH2048_G *big.Int
	RSA            *crypto.RSACryptor
}

func NewMTProtoCryptor() *MTProtoCryptor {
	return &MTProtoCryptor{
		Fingerprint:    fingerprint,
		PQ:             pq,
		P:              p,
		Q:              q,
		DH2048_P:       dh2048_p,
		DH2048_G:       dh2048_g,
		BigIntDH2048_P: new(big.Int).SetBytes(dh2048_p),
		BigIntDH2048_G: new(big.Int).SetBytes(dh2048_g),
		RSA:            crypto.NewRSACryptor(),
	}
}
