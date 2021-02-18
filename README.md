# kubectl_mux

A tool to run the same kubectl command against multiple contexts, and report the result back in quick succession

```
$ kmux get pods 

[list of avaialble contexts]

[context name
NAME                 READY   STATUS      RESTARTS   AGE
pod-name              1/1    Running       0         3d
...
]
```

The program first prints the contexts that match the filters.

## installation

```
$ make build
```
and add the resulting `kmux` binary to your path. 

## limiting the contexts against which the command is run

By default the command is run against any available context. You can use the `--filter` flag to run the command only against the contexts that match the filter (simple string match). You can specify multiple values, then the command will be run against all contexts that matches any of the filters. 

```
$ kmux describe pods --filter 001 --filter 002
```

## using with kubectl flags

If you tried running `kmux get pods -n default --filter 001`, then you would get an error message, saying that -n is not a flag known to `kmux`. The problem is that kmux can't tell apart the flags you intend to for it, and the flags you intend for `kubectl`. So you have to help it, by using the `--` flag separator. After that, none of the flags are parsed as flags, but as arguments. 
So
```
$ kmux --filter 001 get pods -n default
```

will work. 


## Authentication

The program doesn't do any authentication or setup, it simply executes the kubectl + the commands specified within the os. The `kubectl` command needs to be available where the binary is executed, and it needs to be able to authenticate (based on environmental variables or saved credentials).

