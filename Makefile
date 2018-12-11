all: deploy

.PHONY: start
start:
	@go run cmd/start.go

deploy: build compress
	@echo "deploying..\n"
	@MAKE done

.PHONY: build
build:
	@echo "building.."
	@command go build -ldflags "-s -w" -o bin/go-coindrop-api
	# $(MAKE) compress

.PHONY: compress
compress:
	@echo "compressing.."
	@command upx bin/go-coindrop-api

.PHONY: done
done:
	@echo "done"
