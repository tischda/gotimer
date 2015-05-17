Building windows go programs on linux
=====================================

Instructions here:
* https://github.com/golang/go/wiki/WindowsCrossCompiling
* http://stackoverflow.com/questions/12168873/cross-compile-go-on-osx

On MacOS, commands need to be run with sudo.


sudo caveat
-----------
sudo has a default policy of resetting the Environment and
setting a secure path. you can setup sudo not to reset certain
environment variables by adding some explicit environment
settings to keep in /etc/sudoers:

Run `sudo visudo`

Add the following to the bottom:

Defaults env_keep += "GOROOT"


Install
-------
The instructions from the wiki do not seem to work:

    hansolo:cross-compile daniel$ sudo ./buildpkg.sh windows amd64
    pkg/runtime (windows/amd64)
    go tool dist: opendir /usr/local/go/src/pkg/runtime: No such file or directory
    runtime
    go build runtime: windows/amd64 must be bootstrapped using make.bash

So follow the instruction from StackOverflow:
~~~
cd /usr/local/go/src
sudo GOOS=windows CGO_ENABLED=0 ./make.bash --no-clean
~~~

Actual ompilation:
~~~
GOOS=windows CGO_ENABLED=0 go build -o timer.exe
~~~


Alternatives to consider
------------------------
* https://github.com/mitchellh/gox
