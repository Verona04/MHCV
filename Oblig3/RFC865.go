package main
import (
	"fmt"
	"log"
	"net"
)
func startTCP() {
	l, err := net.Listen("tcp", ":17")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go func(c net.Conn) {
			c.Write([]byte("Quote of the day asd\n"))
			c.Close()
		}(conn)
	}
}
func startUDP() {
	s, err := net.ResolveUDPAddr("udp", ":17")
	l, err := net.ListenUDP("udp", s)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	buf := make([]byte, 1024)
	for {
		r, raddr, err := l.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error: ", err)
		}
		if r > 0 {
			l.WriteToUDP([]byte("Quote of the day asd\n"), raddr)
		}
	}
}
func main() {
	go startTCP()
	startUDP()
}
