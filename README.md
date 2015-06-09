# timer [![Build status](https://ci.appveyor.com/api/projects/status/au8q12tabnam2t9a?svg=true)](https://ci.appveyor.com/project/tischda/timer)

Windows utility written in [Go](https://www.golang.org) to measure the time between two events.
Timers are persisted in the Windows registry key
`HKCU\Software\Tischer\timers` as follows:

Name  | Type      | Data
----  | ----      | ----
t1    | REG_QWORD | 13de77095f0a6014

Data is the number of nanoseconds elapsed since January 1, 1970 UTC.

### Compile

Tested with GO 1.4.2. There are no dependencies.

~~~
go build
~~~

### Usage


~~~
Usage: timer [options] name
  name: name of the timer
  -C="REQUIRED": print elapsed time for command
  -clear=false: clear all timers
  -list=false: list timers
  -read="REQUIRED": read timer (elapsed time)
  -start="REQUIRED": start timer
  -stop="REQUIRED": stop timer and print elapsed time
  -verbose=false: verbose output
  -version=false: print version
~~~

### Example

~~~
U:\src\timer>timer -start t1 -verbose
Starting timer t1

U:\src\timer>timer -read t1
Elapsed time (t1): 5.9200225s

U:\src\timer>timer -start t2 -read t1
Elapsed time (t1): 56.3191111s

U:\src\timer>timer -start t3 -read t1 -stop t2
Elapsed time (t1): 1m30.6471884s
Elapsed time (t2): 34.3280773s

U:\src\timer>timer -clear -verbose
All timers deleted
~~~

### Other timers

* [clTimer](http://www.cylog.org/tools/cmdline.jsp)
* [utime](http://www.rohitab.com/discuss/topic/38678-unix-time-on-windows/)
