package guid

import (
	"fmt"
	"sync/atomic"
	"time"
)

var gid = int64(10000)

func GenerateUID() int64 {
	newid := atomic.AddInt64(&gid, 1)

	return newid
}

func GenerateStringUID() string {
	s := fmt.Sprintf("%v", GenerateUID())
	return s
}

func GenerateSessionID() string {
	return GenerateStringUID()
}

func GenerateMessageID() int64 {
	const nano = 1000 * 1000 * 1000
	unixnano := time.Now().UnixNano()

	messageID := ((unixnano / nano) << 32) | ((unixnano % nano) & -4)
	for {
		//rpc_response
		if (messageID % 4) != 1 {
			messageID += 1
		} else {
			break
		}

		/****************************
		 * // rpc_request
		 * if (messageID % 4) != 3 {
		 * 	messageID += 1
		 * } else {
		 * 	break
		 * }
		 */
	}

	return messageID
}
