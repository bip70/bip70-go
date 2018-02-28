protobuf:
	protoc --go_out=. pkg/protobuf/payments/*.proto

test-bip70:
	go test github.com/bip70/bip70-go/pkg/bip70/... \
	$(TESTARGS)

test-protobuf:
	go test github.com/bip70/bip70-go/pkg/protobuf/payments/... \
	$(TESTARGS)

test: test-bip70 test-protobuf