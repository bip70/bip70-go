protobuf:
	protoc --go_out=. pkg/protobuf/payments/*.proto

test-bip70:
	go test github.com/bip70/bip70-go/pkg/bip70/... \
            -coverprofile=coverage/bip70.out -v \
	        $(TESTARGS)

test-protobuf:
	go test github.com/bip70/bip70-go/pkg/protobuf/payments/... \
	                        -coverprofile=coverage/bip70.out -v \
	                        $(TESTARGS)

# concat all coverage reports together
coverage-concat:
	echo "mode: set" > coverage/full && \
    grep -h -v "^mode:" coverage/*.out >> coverage/full

# full coverage report
coverage: coverage-concat
	go tool cover -func=coverage/full $(COVERAGEARGS)

minimum-coverage: coverage
	./tools/minimum-coverage.sh 60

# full coverage report
coverage-html: coverage-concat
	go tool cover -html=coverage/full $(COVERAGEARGS)

test: test-bip70 test-protobuf