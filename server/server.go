package main

import (
	"encoding/binary"
	"log"
	"net"
)

func main() {
	go handle_workers()
}

func handle_workers() {
	ln, err := net.Listen("tcp", ":42069")
	checkErr(err)

	conn, err := ln.Accept()
	checkErr(err)
	msg := receive_string(conn)
	return

	for {
		conn, err := ln.Accept()
		checkErr(err)
		go connect_worker(conn)
	}
}

func receive_string(conn net.Conn) string {
	sz := receive_int64(conn)
	arr := make([]byte, sz)
	err := binary.Read(conn, binary.LittleEndian, arr)
	checkErr(err)
	return string(arr)
}

func connect_worker(conn net.Conn) {
	defer conn.Close()
}

func send_int64(con net.Conn, x int64) {
	err := binary.Write(con, binary.LittleEndian, x)
	checkErr(err)
}

func receive_int64(con net.Conn) int64 {
	var res int64
	err := binary.Read(con, binary.LittleEndian, &res)
	checkErr(err)
	return res
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
