#include <stdio.h>
int main()
{
	int a = 5, b = 7;
	float x = 67.8564f, y = -789.124f;
	char c = 'A';
	long n = 1234567;
	unsigned u = 4294967295;
	printf("%d%d\n", a, b);
	printf("%3d%3d\n", a, b);
	printf(/*BLANK*/"%f,%f\n", x, y);
	printf(/*BLANK*/"%-6f,%-6f\n", x, y);
	printf(/*BLANK*/"%9.2f,%6.2f,%.4f,%.4f,%f,%-6f\n", x, y, x, y, x, y);
	printf(/*BLANK*/"%e,%4.2e\n", x, y);
	printf(/*BLANK*/"%c,%d,%o,%x\n", c, c, c, c);
	printf(/*BLANK*/"%ld,%lo,%x\n", n, n, n);
	printf(/*BLANK*/"%u,%o,%x,%d\n", u, u, u, u);
	printf(/*BLANK*/"%s,%5.3s\n", "COMPUTER", "COMPUTER");
	return 0;
}
r e t u r n   0  
 