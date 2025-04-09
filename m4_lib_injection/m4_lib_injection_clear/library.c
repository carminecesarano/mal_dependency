// malicious.c
#include <stdio.h>
#include <stdlib.h>

__attribute__((constructor))
void init() {
    //system("echo 'MALICIOUS PAYLOAD RAN' >> /tmp/hijack.log");
}
