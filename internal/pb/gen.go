package pb

//go:generate protoc --go_out=. --go_opt=paths=source_relative ./url.proto --experimental_allow_proto3_optional
//go:generate protoc --go_out=. --go_opt=paths=source_relative ./round_trip.proto --experimental_allow_proto3_optional
