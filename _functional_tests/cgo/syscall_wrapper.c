#include <stdio.h>
#include <unistd.h>
#include <string.h>
#include <sys/syscall.h>
#include <linux/ioprio.h>
#include <sched.h>
#include <fcntl.h>

// Syscall 0 read
void sys_read() {
    char buffer[100];
    int fd;
    long no_bytes_read;    

    fd = syscall(SYS_open, "/etc/passwd", O_RDONLY);    // ID = 2      
    no_bytes_read = syscall(SYS_read, fd, buffer, 100); // ID = 0
    syscall(SYS_write, 1, buffer, no_bytes_read);       // ID = 1
}

// Syscall 0 read via clib
void sys_read_clib() {
    char buffer[100]; 
    FILE *f;
    size_t no_bytes_read;

    f = fopen("/etc/passwd", "r");              // SYS_openat ID=257     [missing stack traces]
    no_bytes_read = fread(buffer, 1, 100, f);   // SYS_read ID=0   [missing stack traces]
    fwrite(buffer, 1, no_bytes_read, stdout);   // SYS_write ID=1  [missing stack traces]
}

// Syscall 1 write via assembly
void sys_write_assembly() {
    int ret;
    asm volatile(
        "syscall"
        : "=a"(ret)
        : "a"(SYS_write),
        "D"(1),
        "S"("Hello via assembly!\n"),
        "d"(19)
        : "rcx", "r11", "memory"
    );
}

// Syscall 1 write via syscall
void sys_write() {
    const char *message = "Hello via syscall!\n";
    syscall(SYS_write, 1, message, 19);
}

// Syscall 145 sched_getscheduler
void sys_sched_getscheduler() {
    int policy = syscall(SYS_sched_getscheduler, 0); 
    printf("Current scheduling policy: %d\n", policy);    
}

// Syscall 170 sethostname
void sys_sethostname() {
    char new_hostname[] = "new_hostname";
    syscall(SYS_sethostname, new_hostname, strlen(new_hostname));
}

// Syscall 204 sched_getaffinity
void sys_sched_getaffinity() {
    unsigned long mask = 0;
    syscall(SYS_sched_getaffinity, 0, sizeof(mask), &mask);
    printf("CPU affinity mask: %lx\n", mask);
}

// Syscall 252 ioprio_get
void sys_ioprio_get() {
    int ioprio = syscall(SYS_ioprio_get, IOPRIO_WHO_PROCESS, getpid());
    printf("IO priority: %d\n", ioprio);
}




