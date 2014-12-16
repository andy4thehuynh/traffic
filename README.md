Traffic
====

A CLI tool that checks the status of your program. Traffic will turn your Tmux status bar yellow when your process begins. It will emit a green status bar if your program successfully finishes and red if you get errors.

## Installation
```
$ cd ~/
$ git clone git@github.com:andy4thehuynh/traffic.git
$ go install
```

Run a Tmux session and try this:

```
$ traffic echo hello world
hello world
```

Enjoy!
