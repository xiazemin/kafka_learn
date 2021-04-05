% protoc --go_out=plugins=grpc:exp2/proto/message.pb.go exp2/proto/*
2021/04/05 11:04:56 WARNING: Missing 'go_package' option in "exp2/proto/message.proto",
please specify it with the full Go package path as
a future release of protoc-gen-go will require this be specified.
See https://developers.google.com/protocol-buffers/docs/reference/go-generated#package for more information.

--go_out: protoc-gen-go: plugins are not supported; use 'protoc --go-grpc_out=...' to generate gRPC


go get github.com/golang/protobuf/protoc-gen-go 

 % protoc --go_out=plugins=grpc:exp2/proto/ exp2/proto/* 

 https://www.cnblogs.com/li-peng/p/14201079.html

 这个file_protos_model_students_proto_rawDesc是proto里数据的描述信息。如 proto的路径、包名称，message信息等等。
file_protos_model_students_proto_rawDesc描述信息有什么用呢？
当我们在执行proto.Marshal的时候，会对传入的参数Message进行验证，比如每个message字段的index、数据类型，是否和file_protos_model_students_proto_rawDesc一致。如果不一致就说明是有问题的。

% go run exp2/producer/main.go
Delivered message to test[0]@11

% go run exp2/consumer/main.go
Consumer error: Subscribed topic not available: ^aRegex.*[Tt]opic: Broker: Unknown topic or partition (<nil>) 
<nil>
Consumer error: <nil> ({{{} [] [] 0xc000104160} 0 [] 1234567 test proto kafka})  []byte(nil)
Message on <nil>,test[0]@16: 


