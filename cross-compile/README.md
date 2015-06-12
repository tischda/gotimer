Building windows GO programs on linux
=====================================

Instructions (not working):
* https://github.com/golang/go/wiki/WindowsCrossCompiling

Use this instead:
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

`Defaults env_keep += "GOROOT"`


Install
-------
~~~
cd /usr/local/go/src
sudo GOOS=windows CGO_ENABLED=0 ./make.bash --no-clean
~~~

Actual compilation:
~~~
GOOS=windows CGO_ENABLED=0 go build -o timer.exe
~~~


Alternatives to consider
------------------------
* https://github.com/mitchellh/gox
* http://dave.cheney.net/2015/03/03/cross-compilation-just-got-a-whole-lot-better-in-go-1-5

