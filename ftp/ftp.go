package ftp

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func (c *FtpConn) Serve() {
	defer (*c.Conn).Close()
	input := bufio.NewScanner(*c.Conn)
	for input.Scan() {
		//echo(c, input.Text(), 1*time.Second)
		s := strings.Fields(input.Text())
		if len(s) == 0 {
			continue
		}

		command, args := s[0], s[1:]

		switch command {
		case "cd":
			c.CdCommand(args)
		case "ls":
			c.LsCommand()
		case "get":
			c.GetCommand(args)
		case "close":
			return
		}
	}
}

func (c *FtpConn) LsCommand() {
	files, err := ioutil.ReadDir(c.WorkDir)
	if err != nil {
		log.Print(err)
		_, _ = fmt.Fprintln(*c.Conn, err.Error())
	}

	fmt.Fprintln(*c.Conn, "Name\tIs_dir")

	for _, file := range files {
		fmt.Fprintln(*c.Conn, file.Name(), "\t", file.IsDir())
	}
}

func (c *FtpConn) CdCommand(args []string) {
	if len(args) == 0 {
		_, _ = fmt.Fprintln(*c.Conn, "Command CD need 1 argument-PATH. Try again.")
		return
	}

	fmt.Printf("New path for client: %v\n", args[0])

	absPath, err := filepath.Abs(c.WorkDir + args[0])
	_, err = ioutil.ReadDir(absPath)
	if err != nil {
		var err2 error
		absPath, err2 = filepath.Abs(args[0])
		_, err2 = ioutil.ReadDir(absPath)
		if err2 != nil {
			log.Print(err2)
			_, _ = fmt.Fprintln(*c.Conn, err2.Error())
			return
		}
	}
	c.WorkDir = absPath
	fmt.Fprintln(*c.Conn, "New work directory = ", absPath)
}

func (c *FtpConn) HelpCommand() {
	fmt.Fprintln(*c.Conn, "FtpSrv support next commands:")
	fmt.Fprintln(*c.Conn, "\t* help\treturn client-info about ftpSrv")
	fmt.Fprintln(*c.Conn, "\t* cd\tset new work directory for client, need 1 argument - new dir-path")
	fmt.Fprintln(*c.Conn, "\t* ls\treturn list of directories and files in work-dir")
	fmt.Fprintln(*c.Conn, "\t* get\treturn file from work-dir, need 1 argument - files name")
}

func (c *FtpConn) GetCommand(args []string) {
	if len(args) == 0 {
		_, _ = fmt.Fprintln(*c.Conn, "Command GET need 1 argument-FILENAME. Try again.")
		return
	}

	path := filepath.Join(c.WorkDir, args[0])
	file, err := os.Open(path)
	if err != nil {
	}

	_, err = io.Copy(*c.Conn, file)
	if err != nil {
		log.Print(err)
		return
	}
	io.WriteString(*c.Conn, "\r\n")
}
