#include <stdio.h>
#include <sys/syscall.h>

// Syscall 0 read
void hello() {
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
void sys_write_c() {
    const char *message = "Hello via syscall!\n";
    syscall(SYS_write, 1, message, 19);
}

// Syscall 145 sched_getscheduler
void sys_sched_getscheduler() {
    int policy = syscall(SYS_sched_getscheduler, 0); 
    printf("Current scheduling policy: %d\n", policy);    
}


