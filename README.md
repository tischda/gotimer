﻿# gotimer [![Build status](https://ci.appveyor.com/api/projects/status/ybwsfvbfv5vdteqy?svg=true)](https://ci.appveyor.com/project/tischda/gotimer)

Windows utility written in [Go](https://www.golang.org) to measure the time between two events.
Timers are persisted in the Windows registry:

`HKEY_CURRENT_USER\Software\Tischer\timers`

Name  | Type      | Data
----  | ----      | ----
t1    | REG_QWORD | 13de77095f0a6014

Data is the number of nanoseconds elapsed since January 1, 1970 UTC.

### Install

~~~
go install github.com/tischda/gotimer
~~~

### Usage

~~~
Usage: gotimer [OPTION] exec task
       gotimer [OPTION] COMMAND timer-name

 COMMANDS:
  start: start named timer
  read: read timer (elapsed time)
  stop: read and then clear timer
  list: list timers
  clear: clear named timer, remove from registry
  exec: execute task and print elapsed time

 OPTIONS:
  -quiet
        hide process output
  -version
        print version and exit
~~~

Examples:

~~~
C:\>gotimer start t1
C:\>gotimer read t1
Elapsed time (t1): 5.9200225s

C:\>gotimer start t2
C:\>gotimer list
[t1 t2]

C:\>gotimer stop t1
Elapsed time (t1): 1m30.6471884s

C:\>gotimer clear

C:\>gotimer -quiet exec "dir /s"
Total time: 91.2001ms
~~~

### Other timers

* [clTimer](http://www.cylog.org/tools/cmdline.jsp)
* [utime](http://www.rohitab.com/discuss/topic/38678-unix-time-on-windows/)
