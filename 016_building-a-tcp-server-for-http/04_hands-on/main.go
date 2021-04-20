package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8055")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn)  {
	defer func() {
		fmt.Println("Connection closing")
		conn.Close()
	}()

	request(conn)

	//response(conn)
}

func request(conn net.Conn) {
	i := 0
	method := ""
	uri := "/"
	fmt.Println("in request")
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			// request line
			method = strings.Fields(ln)[0] // method
			uri = strings.Fields(ln)[1] // uri
			fmt.Println("***METHOD", method)
			fmt.Println("***URI", uri)
		}
		if ln == "" {
			// headers are done
			break
		}
		i++
	}
	if uri == "/test" {
		response(conn, "test data")
	} else if uri == "/haha" {
		response(conn, "keep smiling :)")
	} else {
		response(conn, uri)
	}
}

func response(conn net.Conn, data string) {
	fmt.Println("in response")
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>` + data + `</strong></body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
