package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

type P struct {
	PType byte
}

//spawned function to take care of incoming packets
func rpak(con net.Conn) {
	fmt.Printf("connected to %s\n", con.RemoteAddr().String())

	for {
		gobDecode := gob.NewDecoder(con)
		var packetIn P
		err := gobDecode.Decode(&packetIn)
		fmt.Println(packetIn)
		packetType := 4

		if err != nil {
			fmt.Println(err)
			fmt.Println("fucl")
			return
		}

		//packets are identified by the first byte and handled accordingly

		//DATa packet
		if packetType == 0 {
			fmt.Println("recieved data packet \n")
			break
		}

		//Request to Send Data
		if packetType == 1 {
			fmt.Println("recieved a request to send data \n")
			break
		}

		//Action of Elevation Message
		if packetType == 2 {
			fmt.Println("recieved notice, becoming a cluster head \n")
			break
		}

		//Action of Elevation Confirmation
		if packetType == 3 {
			fmt.Println("recieved notice, requested new cluster head added to network \n")
			break
		}

		//INFo packet
		if packetType == 4 {
			fmt.Println("recieved information \n")
			break
		}

		//INfo packet Query (or INQuery, whatever you prefer)
		if packetType == 5 {
			fmt.Println("recieved a request for information \n")
			break
		}

		//Marco Polo Question
		if packetType == 6 {
			fmt.Println("recieved a marco-polo question \n")
			break
		}

		//Marco Polo Responce
		if packetType == 7 {
			fmt.Println("recieved a marco-polo solution \n")
			break
		}
	}
	fmt.Println("closed pipe")
	con.Close()
}

func main() {
	l, err := net.Listen("tcp4", ":5831")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer l.Close()

	for {
		con, err := l.Accept()

		if err != nil {
			fmt.Println(err)
			return
		}

		go rpak(con)

	}
}