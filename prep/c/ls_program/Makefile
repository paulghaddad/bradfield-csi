CC = cc
CFLAGS = -Wall -std=c99
DEBUG = -g
DEPS = ls_c.h
OBJ = ls_c.o
PROGRAM_NAME = ls_c


%.o: %.c $(DEPS)
	$(CC) -o $@ $< $(CFLAGS)

ls: clean ls_c.c
	$(CC) $(CFLAGS) -o $(PROGRAM_NAME) $(PROGRAM_NAME).c

debug: clean ls_c.c
	$(CC) $(DEBUG) $(CFLAGS) -o $(PROGRAM_NAME) $(PROGRAM_NAME).c

clean:
	rm -rf ls_c ls_c.dSYM