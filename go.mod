module github.com/PretendoNetwork/pokemon-gen6

go 1.21

replace github.com/PretendoNetwork/nex-protocols-go/v2 => ./nex-protocols-go
replace github.com/PretendoNetwork/nex-protocols-common-go/v2 => ./nex-protocols-common-go

require (
	github.com/PretendoNetwork/grpc-go v1.0.2
	github.com/PretendoNetwork/nex-go/v2 v2.0.1
	github.com/PretendoNetwork/nex-protocols-common-go/v2 v2.0.2
	github.com/PretendoNetwork/nex-protocols-go/v2 v2.0.1
	github.com/PretendoNetwork/plogger-go v1.0.4
	github.com/joho/godotenv v1.5.1
	github.com/lib/pq v1.10.9
	google.golang.org/grpc v1.63.0
)

require (
	github.com/dolthub/maphash v0.1.0 // indirect
	github.com/fatih/color v1.16.0 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/jwalton/go-supportscolor v1.2.0 // indirect
	github.com/klauspost/compress v1.17.7 // indirect
	github.com/lxzan/gws v1.8.1 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/rasky/go-lzo v0.0.0-20200203143853-96a758eda86e // indirect
	github.com/superwhiskers/crunch/v3 v3.5.7 // indirect
	golang.org/x/exp v0.0.0-20240404231335-c0f41cb1a7a0 // indirect
	golang.org/x/mod v0.17.0 // indirect
	golang.org/x/net v0.24.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
	golang.org/x/term v0.19.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240401170217-c3f982113cda // indirect
	google.golang.org/protobuf v1.33.0 // indirect
)
