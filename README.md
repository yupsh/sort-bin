# yup-sort

```
NAME:
   sort - sort lines of text files

USAGE:
   sort [OPTIONS] [FILE...]

      Write sorted concatenation of all FILE(s) to standard output.
      With no FILE, or when FILE is -, read standard input.

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --reverse, -r                      reverse the result of comparisons (default: false)
   --numeric-sort, -n                 compare according to string numerical value (default: false)
   --unique, -u                       output only the first of an equal run (default: false)
   --ignore-case, -f                  fold lower case to upper case characters (default: false)
   --key value, -k value              sort via a key; KEYDEF gives location and type (default: 0)
   --field-separator value, -t value  use SEP instead of non-blank to blank transition
   --random-sort, -R                  shuffle, but group identical keys (default: false)
   --ignore-leading-blanks, -b        ignore leading blanks (default: false)
   --version-sort, -V                 natural sort of (version) numbers within text (default: false)
   --human-numeric-sort               compare human readable numbers (e.g., 2K 1G) (default: false)
   --month-sort, -M                   compare (unknown) < 'JAN' < ... < 'DEC' (default: false)
   --stable, -s                       stabilize sort by disabling last-resort comparison (default: false)
   --help, -h                         show help
```
