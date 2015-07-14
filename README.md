# timer [![Build status](https://ci.appveyor.com/api/projects/status/au8q12tabnam2t9a?svg=true)](https://ci.appveyor.com/project/tischda/timer)

Windows utility written in [Go](https://www.golang.org) to measure the time between two events.
Timers are persisted in the Windows registry:

`HKEY_CURRENT_USER\Software\Tischer\timers`

Name  | Type      | Data
----  | ----      | ----
t1    | REG_QWORD | 13de77095f0a6014

Data is the number of nanoseconds elapsed since January 1, 1970 UTC.

### Install

There are no dependencies.

~~~
go get github.com/tischda/timer
~~~

### Usage

~~~
Usage: ./timer [option] command name

 COMMANDS:
  start: start timer
  read: read timer (elapsed time)
  stop: read and then clear timer
  list: list timers
  clear: clear timer, name empty = uninstall
  exec: execute process and print elapsed time

 OPTIONS:
  -quiet=false: hide process output
  -version=false: print version and exit
~~~

Example:

~~~
C:\>timer start t1
C:\>timer read t1
Elapsed time (t1): 5.9200225s

C:\>timer start t2
C:\>timer list
[t1 t2]

C:\>timer stop t1
Elapsed time (t1): 1m30.6471884s

C:\>timer clear

C:\>timer -quiet exec "dir /s"
Total time: 91.2001ms
~~~

### Other timers

* [clTimer](http://www.cylog.org/tools/cmdline.jsp)
* [utime](http://www.rohitab.com/discuss/topic/38678-unix-time-on-windows/)
