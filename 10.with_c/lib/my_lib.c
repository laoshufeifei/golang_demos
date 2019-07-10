/*
gcc -c my_lib.c -o my_lib.o
ar -crs libmy_lib.a my_lib.o
*/

#include "my_lib.h"
int myAdd3(int a, int b) {
	return a + b;
}