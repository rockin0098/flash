package session

import (
	"net"
	"sync"

	"github.com/rockin0098/flash/base/guid"
)

// connection manager
// conn 要建立id map

func GenerateConnectionID() string {
	return guid.GenerateStringUID()
}

type ConnectionManager struct {
	connectionMap *sync.Map
}

var connectionManager = &ConnectionManager{
	connectionMap: &sync.Map{},
}

func GetConnectionManager() *ConnectionManager {
	return connectionManager
}

func (c *ConnectionManager) Add(conn net.Conn) string {
	connid := GenerateConnectionID()
	c.connectionMap.Store(connid, conn)

	return connid
}

func (c *ConnectionManager) Remove(connID string) {
	c.connectionMap.Delete(connID)
}

func (c *ConnectionManager) Get(connID string) net.Conn {
	con, ok := c.connectionMap.Load(connID)
	if !ok {
		return nil
	}

	return con.(net.Conn)
}
