all: deploy

.PHONY: start
start:
	@go run cmd/start.go

start/local:
	@(. .env.local && go run cmd/start.go)

start/staging:
	@(. .env.staging && go run cmd/start.go)

start/prod:
	@(. .env.prod && go run cmd/start.go)

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

.PHONY: prep
prep:
	@echo "prepping..\n"
	@command rm -rf ../go-coindrop-api-EB; mkdir ../go-coindrop-api-EB; cp -r . ../go-coindrop-api-EB; rm -rf ../go-coindrop-api-EB/.git; cd ../go-coindrop-api-EB; zip ../go-coindrop-api-EB.zip -r * .[^.]*; mv ../go-coindrop-api-EB.zip ~/Desktop; rm -rf ~/go/src/github.com/waymobetta/go-coindrop-api-EB

.PHONY: goa
goa:
	@goagen bootstrap -d github.com/waymobetta/go-coindrop-api/design

.PHONY: docs/swagger
docs/swagger:
	@cp swagger/swagger.json web/docs/swagger.json