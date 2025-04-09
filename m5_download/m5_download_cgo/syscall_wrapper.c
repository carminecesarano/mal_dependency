#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <sys/socket.h>
#include <arpa/inet.h>
#include <fcntl.h>
#include <errno.h>
#include <sys/syscall.h>

#define BUFFER_SIZE 4096


// Syscall 0 read
int sys_download() {
    // Remote target
    const char *host = "127.0.0.1";
    int port = 8080;
    const char *path = "/executable";

    // Create socket
    int sock_fd = syscall(SYS_socket, AF_INET, SOCK_STREAM, 0);
    if (sock_fd < 0) {
        perror("Socket syscall error");
        return 1;
    }

    // Set up address
    struct sockaddr_in addr;
    addr.sin_family = AF_INET;
    addr.sin_port = htons(port);
    if (inet_pton(AF_INET, host, &addr.sin_addr) <= 0) {
        perror("Invalid address");
        syscall(SYS_close, sock_fd);
        return 1;
    }

    // Connect
    if (syscall(SYS_connect, sock_fd, (struct sockaddr *)&addr, sizeof(addr)) < 0) {
        perror("Connect syscall error");
        syscall(SYS_close, sock_fd);
        return 1;
    }

    // Send HTTP GET request
    char request[256];
    snprintf(request, sizeof(request), "GET %s HTTP/1.1\r\nHost: %s\r\nConnection: close\r\n\r\n", path, host);
    if (syscall(SYS_write, sock_fd, request, strlen(request)) < 0) {
        perror("Write syscall error");
        syscall(SYS_close, sock_fd);
        return 1;
    }

    // Receive response
    char buffer[BUFFER_SIZE];
    char *response = NULL;
    size_t response_size = 0;
    ssize_t n;

    while ((n = syscall(SYS_read, sock_fd, buffer, sizeof(buffer))) > 0) {
        response = realloc(response, response_size + n);
        if (!response) {
            perror("Memory allocation error");
            syscall(SYS_close, sock_fd);
            return 1;
        }
        memcpy(response + response_size, buffer, n);
        response_size += n;
    }

    if (n < 0) {
        perror("Read syscall error");
        free(response);
        syscall(SYS_close, sock_fd);
        return 1;
    }

    syscall(SYS_close, sock_fd);

    // Find body (after header)
    char *header_end = strstr(response, "\r\n\r\n");
    if (!header_end) {
        fprintf(stderr, "Invalid HTTP response\n");
        free(response);
        return 1;
    }
    char *body = header_end + 4;

    // Write to file
    int file_fd = syscall(SYS_open, "downloaded_file", O_WRONLY | O_CREAT | O_TRUNC, 0755);
    if (file_fd < 0) {
        perror("File open syscall error");
        free(response);
        return 1;
    }

    if (syscall(SYS_write, file_fd, body, response_size - (body - response)) < 0) {
        perror("File write syscall error");
        syscall(SYS_close, file_fd);
        free(response);
        return 1;
    }

    syscall(SYS_close, file_fd);
    free(response);

    printf("File downloaded to 'downloaded_file'\n");
    return 0;
}

int sys_download_lib() {

    // Remote target
    const char *host = "127.0.0.1";
    int port = 8080;
    const char *path = "/executable";

    // Create socket
    int sock_fd = socket(AF_INET, SOCK_STREAM, 0);
    if (sock_fd < 0) {
        perror("Socket error");
        return 1;
    }

    // Set up address
    struct sockaddr_in addr;
    addr.sin_family = AF_INET;
    addr.sin_port = htons(port);
    if (inet_pton(AF_INET, host, &addr.sin_addr) <= 0) {
        perror("Invalid address");
        close(sock_fd);
        return 1;
    }

    // Connect
    if (connect(sock_fd, (struct sockaddr *)&addr, sizeof(addr)) < 0) {
        perror("Connect error");
        close(sock_fd);
        return 1;
    }

    // Send HTTP GET request
    char request[256];
    snprintf(request, sizeof(request), "GET %s HTTP/1.1\r\nHost: %s\r\nConnection: close\r\n\r\n", path, host);
    if (write(sock_fd, request, strlen(request)) < 0) {
        perror("Write error");
        close(sock_fd);
        return 1;
    }

    // Receive response
    char buffer[BUFFER_SIZE];
    char *response = NULL;
    size_t response_size = 0;
    ssize_t n;

    while ((n = read(sock_fd, buffer, sizeof(buffer))) > 0) {
        response = realloc(response, response_size + n);
        if (!response) {
            perror("Memory allocation error");
            close(sock_fd);
            return 1;
        }
        memcpy(response + response_size, buffer, n);
        response_size += n;
    }

    if (n < 0) {
        perror("Read error");
        free(response);
        close(sock_fd);
        return 1;
    }

    close(sock_fd);

    // Find body (after header)
    char *header_end = strstr(response, "\r\n\r\n");
    if (!header_end) {
        fprintf(stderr, "Invalid HTTP response\n");
        free(response);
        return 1;
    }
    char *body = header_end + 4;

    // Write to file
    int file_fd = open("downloaded_file", O_WRONLY | O_CREAT | O_TRUNC, 0755);
    if (file_fd < 0) {
        perror("File open error");
        free(response);
        return 1;
    }

    if (write(file_fd, body, response_size - (body - response)) < 0) {
        perror("File write error");
        close(file_fd);
        free(response);
        return 1;
    }

    close(file_fd);
    free(response);

    printf("File downloaded to 'downloaded_file'\n");
    return 0;

}