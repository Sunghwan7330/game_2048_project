#include <stdio.h>
#include "frame.h"

int main(void) {
    int numbers[4][4] = {
        {1, 2, 3, 4},{1, 2, 3, 4},
        {1, 2, 3, 4},{1, 2, 3, 4},
    };

    frame_draw(numbers, 4);
}