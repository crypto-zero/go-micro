package example

//go:generate go build ../
//go:generate protoc --plugin=protoc-gen-micro=./protoc-gen-micro --go_out=. --go_opt=paths=source_relative --micro_out=. --micro_opt=paths=source_relative person/person.proto
//go:generate protoc --plugin=protoc-gen-micro=./protoc-gen-micro --go_out=. --go_opt=paths=source_relative --micro_out=. --micro_opt=paths=source_relative greeter/greeter.proto
//go:generate rm protoc-gen-micro
