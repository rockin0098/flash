package service

import (
	"encoding/hex"
	"net"
	"testing"
	"time"

	. "github.com/rockin0098/flash/base/logger"
)

func TestMain(m *testing.M) {
	Log.Info("test starting...")
	m.Run()
	Log.Info("test ending...")
}

// TL_req_p_q onReqPq
// raw request
// 542afc30d59d4016b3d81213f00a045b5a6603048df4fd24d785328103d58d34132059bc3a57ad8160375b5f677c21a683db3fa0806869cc484565fd1d66a983110c9365e22742b13a113c891bac769cd0154eb994bd9290069f8d61768523985a15059cc5a753b86c
// biz request
// 0000000000000000100bc583386c08541400000078974660a426a32f080b15d83ac95f39657e0232
// nonce
// a426a32f080b15d83ac95f39657e0232
// server nonce
// d5f2e344b9f287be3d7be47e861f4c2f
// biz response
// 0000000000000000b9493b2c0db7b85b4000000063241605a426a32f080b15d83ac95f39657e0232d5f2e344b9f287be3d7be47e861f4c2f0817ed48941a08f98100000015c4b51c01000000cd601077c171e0a9
// raw response
// 69d7d0037bba59133377ea2cbdc4eac50f7777fe989816f98e5e2f296710c03159b710ff62dca372199ea30bec8e2102697a37b6b3e50f7e5566827ef6b6ed3c43223e60af3955e1752ae412eb9c2a909177b12845
func TestSend(t *testing.T) {

	conn, err := net.Dial("tcp", "localhost:5222")
	if err != nil {
		Log.Error(err)
		return
	}

	data, err := hex.DecodeString("542afc30d59d4016b3d81213f00a045b5a6603048df4fd24d785328103d58d34132059bc3a57ad8160375b5f677c21a683db3fa0806869cc484565fd1d66a983110c9365e22742b13a113c891bac769cd0154eb994bd9290069f8d61768523985a15059cc5a753b86c")
	if err != nil {
		Log.Error(err)
		return
	}

	n, err := conn.Write(data)
	if err != nil {
		Log.Error(err)
		return
	}

	Log.Infof("n = %v bytes sent", n)

	rb := make([]byte, 1024)
	n, err = conn.Read(rb)
	if err != nil {
		Log.Error(err)
		return
	}

	Log.Infof("n = %v bytes read", n)

	time.Sleep(time.Duration(5) * time.Second)
}
