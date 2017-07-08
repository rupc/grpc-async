# grpc-peer
Basic assumption of gRPC communication model is traditional client/server model. But I know that there are some cases which uses gRPC as a peer communication framework. I feel that server(peer) should have ability to send a message to other connected (client)peers without a request. 

Each peer should maintain its own server and list of client endpoints.
