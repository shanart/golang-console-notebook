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
$ notebook add
this is text for new note

Saved. Size: 324| Date: 04.12.2021

# -l --list
$ notebook list
> 1     this is text for new ...     04.12.2021
  2     old one note item            03.12.2021
  3     This is very old note...     29.11.2021

# read/edit note
$ notebook view 
note id: 2
old one note item _

# -d --delete
$ notebook delete 
note id: 1
Note deleted.

 # -s --search
$ notebook search 
search text: this is
> 1     this is text for new ...     04.12.2021
  3     This is very old note...     29.11.2021
```
