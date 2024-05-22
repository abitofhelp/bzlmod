# bzlmod

... A work in progress.

## Branches
* Main branch builds and runs using Go and Buf; No Bazel.
  * buf generate proto 
  * go build ./... 
* Bazel branch generates the protobuf proto library but not the connectrpc-go proto library.  It is a work in progress... 
  * bazel build //proto/abitofhelp/helloworld/v1:abitofhelp_helloworld_v1_go_proto 

## Manually generating the protobuf and connect .go files from helloworld.proto.
* make buf_generate \
You will need to comment/uncomment the imports of these files in cmd/client and cmd/server. 

## Using Bazel to generate the protobuf and connect.go files from helloworld.proto.
* bazel build //proto/abitofhelp/helloworld/v1:abitofhelp_helloworld_v1_go_proto \
FIXME: At this time, the following command only generate the protobuf file, not the connect one.


## (Re)Generating BUILD.bazel files using Gazelle.
Here are two ways to generate/regenerate BUILD.bazel files.   

* Using Go/Gazelle's CLI \
  I have found using Gazelle's CLI will generate BUILD.bazel files when the Bazel command doesn't do it. 
  * go install github.com/bazelbuild/bazel-gazelle/cmd/gazelle@latest 
  * gazelle -go_prefix github.com/abitofhelp/bzlmod 
  

* Using Bazel 
  * bazel run //:gazelle 

  
