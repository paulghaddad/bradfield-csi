#include <stdbool.h>
#include <sys/stat.h>

struct flag_opts {
  bool long_format;
  bool all_files;
};

void printDirectoryContents(char* path, struct flag_opts *options);
