package ftp

import "net"

type FtpConn struct {
	Conn    *net.Conn
	WorkDir string
}
