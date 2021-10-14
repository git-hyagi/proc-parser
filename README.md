# ABOUT
Very simple program that parses the first line from /proc/stat file.
This is a workaround for situations in which we need to get a better understanding about each CPU time and we don't have the output of "top" command (like, when we have only the data from sosreport available).

# RUNNING
* without parameters it will read the /proc/stat from the current system
~~~
$ go run main.go
%Cpu(s): 9.1 us, 2.6 sy, 0.1 ni, 87.3 id, 0.0 wa, 0.6 hi, 0.3 si, 0.0 st
%cpu0: 6.9 us, 1.9 sy, 0.0 ni, 90.3 id, 0.1 wa, 0.4 hi, 0.4 si, 0.0 st
%cpu1: 9.7 us, 2.7 sy, 0.1 ni, 86.4 id, 0.0 wa, 0.7 hi, 0.4 si, 0.0 st
%cpu2: 9.8 us, 2.7 sy, 0.1 ni, 86.6 id, 0.0 wa, 0.6 hi, 0.2 si, 0.0 st
%cpu3: 9.7 us, 2.8 sy, 0.1 ni, 86.5 id, 0.1 wa, 0.6 hi, 0.3 si, 0.0 st
%cpu4: 9.6 us, 2.7 sy, 0.1 ni, 86.8 id, 0.0 wa, 0.6 hi, 0.2 si, 0.0 st
%cpu5: 9.6 us, 2.7 sy, 0.1 ni, 86.8 id, 0.0 wa, 0.6 hi, 0.2 si, 0.0 st
%cpu6: 9.1 us, 2.8 sy, 0.1 ni, 86.8 id, 0.0 wa, 0.9 hi, 0.3 si, 0.0 st
%cpu7: 9.6 us, 2.8 sy, 0.1 ni, 86.6 id, 0.0 wa, 0.6 hi, 0.2 si, 0.0 st
~~~

* it is also possible to pass a stat file gathered from another system (through sosreport, for example):
~~~
$ go run main sosreport-hostabc-12345/proc/stat
%Cpu(s): 5.2 us, 1.4 sy, 0.0 ni, 91.8 id, 1.0 wa, 0.3 hi, 0.3 si, 0.0 st
%cpu0: 5.4 us, 1.6 sy, 0.0 ni, 91.8 id, 0.8 wa, 0.3 hi, 0.1 si, 0.0 st
%cpu1: 6.1 us, 1.8 sy, 0.0 ni, 90.5 id, 0.8 wa, 0.4 hi, 0.5 si, 0.0 st
%cpu2: 5.8 us, 1.7 sy, 0.0 ni, 90.8 id, 1.2 wa, 0.3 hi, 0.1 si, 0.0 st
%cpu3: 5.7 us, 1.6 sy, 0.0 ni, 91.5 id, 0.8 wa, 0.3 hi, 0.1 si, 0.0 st
%cpu4: 5.4 us, 1.5 sy, 0.0 ni, 91.8 id, 0.9 wa, 0.4 hi, 0.1 si, 0.0 st
%cpu5: 5.4 us, 1.5 sy, 0.0 ni, 91.8 id, 0.7 wa, 0.3 hi, 0.3 si, 0.0 st
%cpu6: 5.5 us, 1.5 sy, 0.0 ni, 91.4 id, 0.8 wa, 0.3 hi, 0.4 si, 0.0 st
%cpu7: 6.4 us, 1.7 sy, 0.0 ni, 90.0 id, 1.2 wa, 0.3 hi, 0.3 si, 0.0 st
%cpu8: 4.9 us, 1.6 sy, 0.0 ni, 92.1 id, 1.1 wa, 0.2 hi, 0.1 si, 0.0 st
%cpu9: 4.2 us, 1.4 sy, 0.0 ni, 91.9 id, 1.4 wa, 0.3 hi, 0.8 si, 0.0 st
%cpu10: 4.9 us, 1.2 sy, 0.0 ni, 92.6 id, 1.0 wa, 0.2 hi, 0.1 si, 0.0 st
%cpu11: 4.6 us, 1.4 sy, 0.0 ni, 92.2 id, 1.0 wa, 0.2 hi, 0.5 si, 0.0 st
%cpu12: 5.4 us, 1.2 sy, 0.0 ni, 91.8 id, 1.0 wa, 0.2 hi, 0.5 si, 0.0 st
%cpu13: 4.2 us, 1.0 sy, 0.0 ni, 93.6 id, 0.9 wa, 0.2 hi, 0.1 si, 0.0 st
%cpu14: 4.2 us, 1.2 sy, 0.0 ni, 93.1 id, 0.9 wa, 0.2 hi, 0.5 si, 0.0 st
%cpu15: 4.8 us, 1.3 sy, 0.0 ni, 92.4 id, 1.2 wa, 0.2 hi, 0.1 si, 0.0 st
~~~
