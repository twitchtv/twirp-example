# twirp-example

This is an exmaple Twirp service for educational purposes. Learn more about
Twirp at its [website](https://twitchtv.github.io/twirp/docs/intro.html) or
[repo](https://github.com/twitchtv/twirp).

## Try it out

First, download this repo with the Go tool and cd into the repository:

```
go get github.com/twitchtv/twirp-example/...
```

Make sure that your machine has `Make` installed.

Next, try running the server and then client:

In one console window:

```
make server
```

And then run the client in another window:

```
make client
```

In the client, you should see something like this:

```
-> % ./client
size:12 color:"red" name:"baseball cap"
```

In the server, something like this:

```% ./server
received req svc="Haberdasher" method="MakeHat"
response sent svc="Haberdasher" method="MakeHat" time="109.01µs"
```

If you edit the `service.proto` file and want to regenerate the updated twirp files:

```
make gen-twirp
```

If you want to make a docker image and then serve the container locally:

```
make docker-image && make docker-container
```

## Code structure

The protobuf definition for the service lives in
`rpc/haberdasher/service.proto`. The `rpc` directory name is a good way to
signal where your service definitions reside.

The generated Twirp and Go protobuf code is in the same directory. This makes it
easy to import for both internal and external users - internally, we need to
import it to have the right types for our implmentation of the service
interface, and externally it needs to be available so clients can import it.

The implementation of the server is in `internal/server`. Putting it
in `internal` means that it can't be imported from outside this repository,
which is nice because we don't have to think about API stability nearly as much.

In addition, `internal/hooks/logging.go` has a file which provides
[`ServerHooks`](https://twitchtv.github.io/twirp/docs/hooks.html) which can log
requests. This is a good demo of how you can use hooks to extend Twirp's basic
functionality - you can use hooks to add instrumentation or even for
authentication.

`tools/` store the codegen packages that are used to generate the twirp go code. It is stored in such a way that go modules could recognise them and we could create the tools binaries using Makefile.

Finally, `cmd/server` and `cmd/client` wrap things together into executable main
packages.

## License

This library is licensed under the Apache 2.0 License.
