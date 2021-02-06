#include <dirent.h>
#include <sys/stat.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

int main(int argc, char *argv[]) {
  char* path = ".";

  bool FLAG_LONG_FORMAT = false;
  bool FLAG_ALL = false;
  bool FLAG_SORT_BY_SIZE = false;

  if (argc > 1) {
    // skip first arg, the program name
    argv++;

    char * arg;
    while ((arg = *argv) != NULL) {
      // Flags
      if (arg[0] == '-') {
        if (arg[1] == 'l') {
          printf("Long format\n");
          FLAG_LONG_FORMAT = true;
        }
        if (arg[1] == 'a') {
          printf("All flag\n");
          FLAG_ALL = true;
        }
        if (arg[1] == 'S') {
          printf("Sort by size flag\n");
          FLAG_SORT_BY_SIZE = true;
        }

        printf("Flags: %s\n", arg);
      } else {
        printf("Arg: %s\n", arg);
      }
      argv++;
    }
  }

  DIR* directory = opendir(path);
  struct dirent *curdir;

  while ((curdir = readdir(directory)) != NULL) {
    struct stat file_stats;
    lstat(curdir->d_name, &file_stats);

    // if directory
    if ((file_stats.st_mode & S_IFMT) == S_IFDIR)
      printf("%s/\n", curdir->d_name);
    // if regular file
    else if ((file_stats.st_mode & S_IFMT) == S_IFREG)
      printf("%s\n", curdir->d_name);
  }

  closedir(directory);

  return EXIT_SUCCESS;
}

// Supported flags:
// -a: include files beginning with a dot
// -l: long format
// -S: Sort by size
// Add error checking
