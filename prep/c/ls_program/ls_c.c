#include <dirent.h>
#include <sys/stat.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <pwd.h>
#include <grp.h>
#include <time.h>

struct flag_opts {
  bool long_format;
  bool all_files;
};

void printDirectoryContents(char* path, struct flag_opts *options);

int main(int argc, char *argv[]) {
  // default path is current working directory
  char* path = ".";

  struct flag_opts ls_options = {false, false};

  if (argc > 1) {
    // skip first arg, the program name
    argv++;

    while (*argv != NULL) {
      // Identify flags
      if (**argv == '-') {
        switch ((*argv)[1]) {
        case 'l':
          ls_options.long_format = true;
          break;
        case 'a':
          ls_options.all_files = true;
          break;
        default:
          printf("ls: illegal option -- %c\n", (*argv)[1]);
          printf("usage: ls [-@ABCFGHLOPRSTUWabcdefghiklmnopqrstuwx1] [file ...]\n");
          exit(1);
        }
      } else
        path = *argv;

      argv++;
    }
  }

  printDirectoryContents(path, &ls_options);

  return EXIT_SUCCESS;
}

void printDirectoryContents(char* path, struct flag_opts *flags) {
  DIR* directory = opendir(path);
  struct dirent *curdir;

  while ((curdir = readdir(directory)) != NULL) {
    struct stat file_stats;
    char* filename = curdir->d_name;
    lstat(filename, &file_stats);

    // If -a flag not on, skip hidden files
    if (filename[0] == '.' && flags->all_files)
      continue;

    char* formattedName;

    // if directory
    if (S_ISDIR(file_stats.st_mode))
      formattedName = strcat(filename, "/");
    // if regular file
    else if (S_ISREG(file_stats.st_mode))
      formattedName = filename;

    // Long format
    if (flags->long_format) {

      // File permissions
      char permission_buf[20];
      char dir = (S_ISDIR(file_stats.st_mode)) ? 'd' : '-';
      char owner_r = (file_stats.st_mode & S_IRUSR) ? 'r' : '-';
      char owner_w = (file_stats.st_mode & S_IWUSR) ? 'w' : '-';
      char owner_x = (file_stats.st_mode & S_IXUSR) ? 'x' : '-';
      char group_r = (file_stats.st_mode & S_IRGRP) ? 'r' : '-';
      char group_w = (file_stats.st_mode & S_IWGRP) ? 'w' : '-';
      char group_x = (file_stats.st_mode & S_IXGRP) ? 'x' : '-';
      char other_r = (file_stats.st_mode & S_IROTH) ? 'r' : '-';
      char other_w = (file_stats.st_mode & S_IWOTH) ? 'w' : '-';
      char other_x = (file_stats.st_mode & S_IXOTH) ? 'x' : '-';
      sprintf(permission_buf, "%c%c%c%c%c%c%c%c%c%c", dir, owner_r, owner_w, owner_x, group_r, group_w, group_x, other_r, other_w, other_x);

      // number of hardlinks
      int hardlinks = file_stats.st_nlink;

      // owner name
      int owner_id = file_stats.st_uid;
      struct passwd *pwd;
      pwd = getpwuid(owner_id);

      // group name
      int group_id = file_stats.st_gid;
      struct group *grp;
      grp = getgrgid(group_id);

      // file size
      int file_size = file_stats.st_size;

      // last modified time
      char ts_buffer[80];
      struct tm timestamp;
      timestamp = *localtime(&file_stats.st_mtime);
      strftime(ts_buffer, 80, "%m %d %H:%M", &timestamp);

      printf("%s %d %s %s %d %s %s\n", permission_buf, hardlinks, pwd->pw_name, grp->gr_name, file_size, ts_buffer, formattedName);
    } else
      printf("%s\n", formattedName);
  }

  closedir(directory);
}
