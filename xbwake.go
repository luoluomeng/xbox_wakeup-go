package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

//WakeUp Package format
/*
	package format:
	[3]byte 0xdd0200
	byte payloadlen
	[2]byte 0x0000
	{//payload
		byte 0x00
		byte idlen
		[]byte liveid
		byte 0x00
	}
*/
func main() {
	if len(os.Args) < 3 {
		fmt.Println("parameter error!")
		fmt.Println("usage:program ip liveID")
	}
	conn, err := connectConsole(os.Args[1])
	if err != nil {
		fmt.Println("connect fail!")
	}
	defer conn.Close()
	fmt.Printf("Start to wake up XBOX on IP:%v,liveID: %v \n", os.Args[1], os.Args[2])
	sendWakePackage(conn, os.Args[2])
	fmt.Printf("Check is XBOX waken up...\n")
	if checAwake(conn) {
		fmt.Println("XBOX wake up success!")
	} else {
		fmt.Println("XBOX wake up fail!")
	}
}

func connectConsole(ip string) (*net.UDPConn, error) {
	var host = flag.String("host", ip, "host")
	var port = flag.String("port", "5050", "port")

	addr, err := net.ResolveUDPAddr("udp", *host+":"+*port)
	if err != nil {
		fmt.Println("Can't resolve address: ", err)
		return nil, err
	}
	return net.DialUDP("udp", nil, addr)
}

func generateWakePackage(liveid string) []byte {
	//generate payload
	powerPayload := []byte{0x00}
	idlen := byte(len(liveid))
	powerPayload = append(powerPayload, idlen)
	powerPayload = append(powerPayload, liveid...)
	powerPayload = append(powerPayload, 0x00)

	//generate package
	powerpacket := []byte{0xdd, 0x02, 0x00}
	payloadlen := byte(len(powerPayload))
	powerpacket = append(powerpacket, payloadlen, 0x00, 0x00)
	powerpacket = append(powerpacket, powerPayload...)
	print("Generating power on packets... \n")
	return powerpacket
}

func sendWakePackage(conn *net.UDPConn, liveid string) (int, error) {
	return conn.Write(generateWakePackage(liveid))
}

//check is XBOX waken, retry 5 times
func checAwake(conn *net.UDPConn) bool {
	pingPackage := []byte{0xdd, 0x00, 0x00, 0x0a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x02}
	_, err := conn.Write(pingPackage)
	if err != nil {
		println("send ping fail,err=%v", err)
		return false
	}
	readBuf := make([]byte, 1024)
	i := 0
	for i < 5 {
		fmt.Printf("Checking is XBOX waken up,retry=%v\n", i+1)
		conn.SetReadDeadline(time.Now().Add(time.Second * 2))
		_, err := conn.Read(readBuf)
		if err == nil {
			break

			return true
		}
		i++
	}

	return false
}
