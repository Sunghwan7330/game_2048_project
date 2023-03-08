#ifdef TEST

#include "unity.h"

#include "frame.h"

void setUp(void)
{
}

void tearDown(void)
{
}

void test_frame_NeedToImplement(void)
{
    const size_t SIZE = 4;
    int numbers[SIZE][SIZE] = {
        {1, 2, 3, 4},
        {5, 6, 7, 8},
        {9, 10, 11, 12},
        {13, 14, 15, 16}
    };

    char buffer[65536];
    // gotoxy 를 목으로 변경
    frame_draw(numbers, SIZE);

    const char *expected = "
    ______";
    TEST_ASSERT_EQUAL_MEMORY(expected, buffer);

}

#endif // TEST
