# ABOUT
Very simple program that parses the first line from /proc/stat file.

# RUNNING
* without parameters it will read the /proc/stat from the current system
~~~
$ go run main.go
%Cpu(s): 9.6 us, 2.7 sy, 0.1 ni, 86.7 id, 0.0 wa, 0.6 hi, 0.3 si, 0.0 st
~~~

* it is also possible to pass a stat file gathered from another system. For example:
~~~
$ go run main sosreport-hostabc-12345/proc/stat
%Cpu(s): 5.4 us, 1.7 sy, 0.0 ni, 74.4 id, 17.8 wa, 0.3 hi, 0.3 si, 0.0 st
~~~
