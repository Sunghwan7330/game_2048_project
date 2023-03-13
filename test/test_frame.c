#ifdef TEST

#include <stdio.h>
#include <stdarg.h>

#include "unity.h"
#include "cmock.h"
#include "mock_stdio.h"

#include "frame.h"

#define TEST_STDOUT_BUFF  10240
char printf_mock_arg0_val[TEST_STDOUT_BUFF];

int printf_mock(const char *format, ...) {
    int ret;
    va_list args;
    va_start(args, format);
    ret = vsnprintf(printf_mock_arg0_val, TEST_STDOUT_BUFF, format, args);
    va_end(args);
    return ret;
}

void setUp(void)
{
}

void tearDown(void)
{

}

void test_frame_NeedToImplement(void)
{
    char expected_output[] = "-------------------------\n"
"|    1|    2|    3|    4|\n"
"-------------------------\n"
"|    1|    2|    3|    4|\n"
"-------------------------\n"
"|    1|    2|    3|    4|\n"
"-------------------------\n"
"|    1|    2|    3|    4|\n"
"-------------------------\n";

    int numbers[4][4] = {
        {1, 2, 3, 4},{1, 2, 3, 4},
        {1, 2, 3, 4},{1, 2, 3, 4},
    };

    TEST_ASSERT_EQUAL_INT(0, printf_mock_StubWithCallback(printf_mock));

    frame_draw(numbers, 4);

    //printf(expected_output);
    fflush(stdout);
    TEST_ASSERT_EQUAL_STRING_LEN(expected_output, printf_mock_arg0_val, strlen(expected_output));

    //TEST_ASSERT_EQUAL_STRING(expected_output, printf_mock_arg0_val);

}

#endif // TEST
