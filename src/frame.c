#include <stdio.h>
#include <stdint.h>
#include "frame.h"

void gotoxy(int x, int y) {
    printf("%c[%d;%df", 0x1B, y, x);
}

static void print_line(const size_t size) {
    for (int i = 0; i < 6 * size + 1; i++)
        printf("-");
}

void frame_draw(int numbers[][4], const size_t size) {
    print_line(size);
    printf("\n");
    for (int y = 0; y < size; y++) {
        for (int x = 0; x < size; x++) {
            printf("|%5d", numbers[y][x]);
        }
        printf("|\n");
        print_line(size);
        printf("\n");
    }
}