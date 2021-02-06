#include <dirent.h>
#include <sys/stat.h>
#include <stdio.h>
#include <stdlib.h>

int main(int argc, char *argv[]) {

  DIR* directory = opendir(".");
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
