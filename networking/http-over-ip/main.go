package main

import (
	"fmt"
	"log"
	"syscall"
)

func main(){
	// Create a raw socket file descriptor (fd)
	// fd: a unique id number assigned by OS to the program to represent an open file or network connection
	// AF_INET = IPv4
	// SOCK_RAW = Raw packet access
	// IP protocol number 253 is used for experimentation/testing purpose
	const experimentProtocol = 253


	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, experimentProtocol)

	if err != nil{
		log.Fatalf("Error: Failed to create a raw socket: %v",err)
	}
	defer syscall.Close(fd)

	fmt.Printf("Success: Raw socket created successfully (File Descriptor: %d)\n",fd)

	destination := &syscall.SockaddrInet4{
		Addr: [4]byte{127,0,0,1},
	}

	payload := []byte("HELLO FROM RAW IP")

	err = syscall.Sendto(fd, payload,0,destination)

	if err != nil {
		log.Fatalf("Error: failed to send the packet: %v",err)
	}
	fmt.Printf("Packet sent successfully\n")

	fmt.Println("Waiting to receive the packet back...")

	buf := make([]byte, 1024)

	n, from, err := syscall.Recvfrom(fd, buf, 0)

	if err != nil {
		log.Fatalf("Error: failed to reveive data: %v",err)
	}

	if addr, ok := from.(*syscall.SockaddrInet4); ok {
		fmt.Printf("Received %d bytes from %d.%d.%d.%d\n", n, addr.Addr[0],addr.Addr[1], addr.Addr[2], addr.Addr[3])
		fmt.Printf("Raw bytes: %x\n",buf[:n])
		fmt.Printf("As string: %q\n",buf[:n])
	}
}
