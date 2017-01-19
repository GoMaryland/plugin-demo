# plugin-demo

To build the plugins, change to the plugin folder and run:

```
go build --buildmode=plugin -o add.so add.go
go build --buildmode=plugin -o subtract.so subtract.go
```

Then, go to the main directory and run `go build` and `./plugin-demo`
