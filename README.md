## Serverless cacheless chat room 

The idea is to play with Golang and push it's boundaries. Here, we build a 
binary that simulates a chatroom! 

* Everyone who has this binary running on their machines, effectively starts their own chat!
* Start the binary using `gochat -name Gautam -port 12345 -host 192.168.1.12` 
* It broadcasts on port 33333 to update any new chatroom about it's existance. Effectively, this is also a health check
* It listens on the port mentioned in the cmd line for any incoming requests for chatting.
* All communication is via gRPC so that we are very sure about the structure of the messages exchanges.
* Channels and Go-routines should be used to manage each chatroom for public and private messaging!

Innovate over this idea!

## Learning from master branch

There are 9 steps which are marked in the code `TODO-WORKSHOP-STEP-?` which need to be implemented in order. Implement them and have fun!


## Contributing back

Fork and submit PR

