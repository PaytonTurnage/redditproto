.PHONY: go-proto
go-proto:
	protoc --go_out=. --go_opt=paths=source_relative *.proto