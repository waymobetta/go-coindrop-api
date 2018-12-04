all: deploy

deploy: build compress
	@echo "deploying..\n"
	@MAKE done

# .PHONY: build
build:
	@echo "building.."
	@command go build -ldflags "-s -w" -o bin/go-coindrop-api
	# $(MAKE) compress

compress:
	@echo "compressing.."
	@command upx bin/go-coindrop-api

done:
	@echo "done"
