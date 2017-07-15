# grpc-async(peer)
Example code of asynchronous communication using gRPC with goroutine

```
go get google.golang.org/grpc
```

## Multiple request 
Client sends multiple messages asynchronously with go routine. This is useful when client don't need to wait response.

## Atomic counter
Client send a request asynchronously without immediate response. Client send another request to increase atomic counter on server side. When the counter has value of 2, original request is going to get a response.

## Relay server
It would be useful when a program has both client side and server side. An example of this is relay server. It waits for cilent's request and relay this to relay server. Picture below shows the communication process.
![Relay](https://github.com/rupc/grpc-async/blob/master/pic/relay.jpg)
