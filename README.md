percol
======

(VERY ALPHA) Port of percol to Go

Installation
============

```
go get github.com/lestrrat/percol
```

(Go by default compiles binaries based on the directory name -- that's why
the repo name is the same as the original percol version :/)

Notes
=====

Much code stolen from https://github.com/mattn/gof
Currently can only search by entering characters, and then when you have 1 entry, you can press enter to print the selected line, which you can feed to another command

Example:

```
ps aux | percol
```

Then try typing the PID of some running process, then press enter.
