timer
=====

Utility for Windows written in GO to measure the time between two events.

~~~
Usage of timer:
  -clear=false: clear all timers
  -elapsed="REQUIRED": print elapsed time for timer
  -start="REQUIRED": start timer
  -stop="REQUIRED": stop timer and print elapsed time
~~~

Example usage
-------------

~~~
U:\src\timer>timer -start t1
Starting timer t1

U:\src\timer>timer -read t1
Elapsed time (t1): 5.9200225s

U:\src\timer>timer -start t2 -read t1
Starting timer t2
Elapsed time (t1): 56.3191111s

U:\src\timer>timer -start t3 -read t1 -stop t2
Starting timer t3
Elapsed time (t1): 1m30.6471884s
Elapsed time (t2): 34.3280773s

U:\src\timer>timer -clear
All timers deleted
~~~
