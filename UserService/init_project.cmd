go mod init UserService
go get github.com/spf13/viper
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
go get golang.org/x/crypto/bcrypt
go get -u github.com/go-sql-driver/mysql