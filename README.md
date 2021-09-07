# Golang console notebook

**NOTE: Work in progress**

This is simple console notebook written on golang. Linux only.

Usage:

```bash
# symbol ">" is indicator to navigate notes with arrows up or down.
# "Enter" to select note to read/edit
# "Esc" to exit note
# Shift+Enter on open note - to save note

# just type "notebook" to start new note
$ notebook 
this is text for new note
/save
Saved. Size: 324| Date: 04.12.2021

# -l --list
$ notebook -l
> 1     this is text for new ...     04.12.2021
  2     old one note item            03.12.2021
  3     This is very old note...     29.11.2021

# read/edit note
$ notebook 1
old one note item _

# -d --delete
$ notebook -d 1
deleted

 # -s --search
$ notebook -s this is
> 1     this is text for new ...     04.12.2021
  3     This is very old note...     29.11.2021
```