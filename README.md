Traffic
====

A CLI tool that changes your Tmux status bar to yellow when your process begins, red if it fails and green if it runs successfully. Simply prefix your executables with `traffic` and keep working while traffic keeps you posted.

## Installation
```
$ go get github.com/andy4thehuynh/traffic
```

In a Tmux session, try this:

```
$ traffic echo hello world
hello world
```

Enjoy!
