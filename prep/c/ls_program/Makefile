CC = cc
CFLAGS = -Wall -std=c99 -c
DEBUG = -g
DEPS = ls_c.h
OBJ = ls_c.o


%.o: %.c $(DEPS)
	$(CC) -c -o $@ $< $(CFLAGS)

ls: ls_c.c
	$(CC) $(CFLAGS) -o ls_c ls_c.c

debug: ls_c.c
	$(CC) $(DEBUG) $(CFLAGS) -o ls_c ls_c.c

clean:
	rm -rf ls_c ls_c.dSYM
