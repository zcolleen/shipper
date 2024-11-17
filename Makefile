
install:
	go install go.uber.org/mock/mockgen
	go install github.com/vektra/mockery/v2
	go install github.com/gojuno/minimock/v3/cmd/minimock

generate:
	cd ./internal && go generate ./...
	mockery --config .mockery.yaml

generate-shipment-creator:
	cd ./shipment_creator && go generate ./...
	mockery --config shipper.mockery.yaml