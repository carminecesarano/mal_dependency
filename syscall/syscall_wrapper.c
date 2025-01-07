#include <stdio.h>
#include <unistd.h>
#include <string.h>
#include <sys/syscall.h>
#include <linux/ioprio.h>
#include <sched.h>

// Syscall 0 read
void sys_read() {
    // output /etc/passwd file
    FILE *f = fopen("/etc/passwd", "r");
    char c;
    while ((c = fgetc(f)) != EOF) {
        printf("%c", c);
    }
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




