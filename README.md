# Golang console notebook

---

#### TODO:
- use this lib for db connection https://github.com/mattn/go-sqlite3
- db connection interface
- Define Note Struct
- Catch flags ( + flags "combo" ) and call handlers
- Screen renderer ( as described below, handle arrow keys )
- Build 

**NOTE: Work in progress**

---

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

Saved. Size: 324| Date: 04.12.2021

# -l --list
$ notebook -l
> 1     this is text for new ...     04.12.2021
  2     old one note item            03.12.2021
  3     This is very old note...     29.11.2021
# Enter btn to select end view/edit note

# -d --delete
$ notebook -d 1
deleted

 # -s --search
$ notebook -s this is
> 1     this is text for new ...     04.12.2021
  3     This is very old note...     29.11.2021
```
