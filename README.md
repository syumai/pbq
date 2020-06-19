# pbq

* pbq is a tool to write a query to ProtocolBuffer's .proto file.
* This tool allows you to easily filter services, rpcs and messages from your proto file.

## Usage

* Show message
```proto
$ cat bookstore.proto | pbq Shelf
message Shelf {
  ...
}
```

* Filter RPC
```proto
$ cat bookstore.proto | pbq BookStore.*Shelf
rpc CreateShelf(CreateShelfRequest) returns (Shelf) {}
rpc GetShelf(GetShelfRequest) returns (Shelf) {}
rpc DeleteShelf(DeleteShelfRequest) returns (google.protobuf.Empty) {}
```

* Filter RPC with service definition (WIP)
```proto
$ cat bookstore.proto | pbq '{BookStore: BookStore.*Book}'
service Bookstore {
  rpc CreateBook(CreateBookRequest) returns (Book) {}
  rpc GetBook(GetBookRequest) returns (Book) {}
  rpc DeleteBook(DeleteBookRequest) returns (google.protobuf.Empty) {}
}
```

* Show message with detail (WIP)
```proto
$ cat bookstore.proto | pbq --detail Shelf
syntax = "proto3";

package endpoints.examples.bookstore;

option java_multiple_files = true;
option java_outer_classname = "BookstoreProto";
option java_package = "com.google.endpoints.examples.bookstore";

import "google/protobuf/empty.proto";

message Shelf {
  ...
}
```

* Show messages recursively from RPC (WIP)
```proto
$ cat bookstore.proto | pbq -I=./protobuf/src --recursive BookStore.ListShelves
rpc ListShelves(google.protobuf.Empty) returns (ListShelvesResponse) {}

message ListShelvesResponse {
  repeated Shelf shelves = 1;
}

message Shelf {
  ...
}

--- google.protobuf
message Empty {}
```

## Installation

```
go get -u github.com/syumai/pbq/cmd/pbq
```

## License

MIT

## Author

syumai