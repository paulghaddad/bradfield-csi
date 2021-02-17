#include <dirent.h>
#include <grp.h>
#include <stdbool.h>
#include <sys/stat.h>
#include <sys/types.h>

struct flag_opts {
  bool long_format;
  bool all_files;
};

void printDirectoryContents(char* path, struct flag_opts *options);
char* formatFilename(char* filename, mode_t filemode);
void printLongFormat(struct stat *file_stats, char* filename);
