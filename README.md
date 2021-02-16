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

## Authentication

The program doesn't do any authentication or setup, it simply executes the kubectl + the commands specified within the os. The `kubectl` command needs to be available where the binary is executed, and it needs to be able to authenticate (based on environmental variables or saved credentials).
