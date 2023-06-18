package main

import (
	"flag"
	"fmt"
	"ftpSrv/ftp"
	"log"
	"net"
	"path/filepath"
)

var rootDir string
var port int

func init() {
	flag.IntVar(&port, "port", 8080, "Port of server")
	flag.StringVar(&rootDir, "rootDir", "C:/", "Root directory for clients")
}

func main() {
	flag.Parse()
	srvAddr := fmt.Sprintf("localhost:%d", port)
	fmt.Printf("Ftp server is running on %v\n", srvAddr)

	listener, err := net.Listen("tcp", srvAddr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // Например, обрыв соединения continue
		}
		go handleConn(conn) // Обработка подключения
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	absPath, err := filepath.Abs(rootDir)
	if err != nil {
		log.Fatal(err)
	}

	newConn := &ftp.FtpConn{
		Conn:    &c,
		WorkDir: absPath,
	}
	newConn.Serve()
}
