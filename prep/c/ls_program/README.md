# ls Clone

## Usage

This program is a clone of Unix's ls program. Provided a path and option flags, it will print out information for the contents in the path directory.

List non-hidden files in the current directory: ls_c
List non-hidden files in the path: ls_c {path}

## Options

It supports two options currently, with more to be added.

```
-a: List all files, including hidden ones
Example: ls_c -a {path}
 
-l: Print directory contents in long format
Example: ls_c -l {path}

Notes:
 * Multiple flag options can be used provided they are separated by a space-delimited list.
 * Example: `ls_c -l -a {path}`
```

## Things I need to fix

* Passing in a single filename produces a segfault at `opendir`.
* Add more robust error checking
