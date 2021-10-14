# ABOUT
Very simple program that parses the first line from /proc/stat file.

# RUNNING
* without parameters it will read the /proc/stat from the current system
~~~
go run main.go
~~~

* it is also possible to pass a stat file gathered from another system. For example:
~~~
go run main sosreport-hostabc-12345/proc/stat
~~~
