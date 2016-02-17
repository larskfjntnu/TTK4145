package udp

/*
	This package acts as the network program for the phoenix program.
	It is a standardized udp package that sends a byte slice over the network
	using the udp protocol.
*/

import (
	"fmt"
	"net"
	"strconv"
)
type Message struct{
	Value int
}
func SendUdp(message chan int) {
	localaddress, err := net.ResolveUDPAddr("udp4", ":9001")
	broadcastaddress,err := net.ResolveUDPAddr("udp4", "255.255.255.255:9000")
	if err != nil {
		fmt.Println("UDP:\t Could not resolve UDPAddress.")
		return 
	}
	sock, err := net.ListenUDP("udp4", localaddress)
	defer func() {sock.Close()}()
	if err != nil {
		fmt.Printf("%s\n",err)
		return
	}
	for{
		select{
			case val :=<-message:
			tempstr := strconv.Itoa(val)
			tempslc := []byte(tempstr)
			sock.WriteToUDP(tempslc, broadcastaddress)
		}
	}
	fmt.Printf("Closing socket.\n")
	sock.Close()
}

/*
	This will be run as a goroutine
*/
func ReadUdp(readChannel chan int,  killChannel , killedChannel chan struct{}) {
	localaddress, err := net.ResolveUDPAddr("udp4", "0.0.0.0:9000")
	sock, err := net.ListenUDP("udp4", localaddress)
	
	defer func() {sock.Close()}()
	defer close(killedChannel)

	if err != nil {
		fmt.Printf("%s\n",err)
		return
	}
	var buf []byte = make([]byte, 16)
	for {
		select{
		case  _ , ok := <- killChannel:
			if !ok {
				return
			}
		default:
			n,_,err := sock.ReadFromUDP(buf)
			if err != nil {
				fmt.Println(err)
			}
			countstr := string(buf[0:n])
			count,_ := strconv.Atoi(countstr)
			readChannel<-count
		}	
	}		
}