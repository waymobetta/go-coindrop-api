all: deploy

.PHONY: vendor
vendor:
	@go mod vendor

.PHONY: test
test:
	@go test -v ./...

.PHONY: test/local
test/local:
	(. .env.local && go test -v controllers/* $(ARGS))

.PHONY: test/auth
test/auth:
	@go test -v auth/*.go $(ARGS)

.PHONY: test/controllers
test/controllers:
	@go test -v controllers/*.go $(ARGS)

.PHONY: start
start:
	@go run cmd/coindrop/main.go

start/local:
	@(. .env.local && go run cmd/coindrop/main.go)

start/staging:
	@(. .env.staging && go run cmd/coindrop/main.go)

start/prod:
	@(. .env.prod && go run cmd/.go)

.PHONY: start/docs
start/docs:
	@(cd web/documentation && python -m SimpleHTTPServer 8000)

deploy: build compress
	@echo "deploying..\n"
	@MAKE done

.PHONY: build
build:
	@echo "building.."
	@command go build -ldflags "-s -w" -o cmd/go-coindrop-api cmd/start.go
	# $(MAKE) compress
	#
.PHONY: build/docker
build/docker:
	@docker build -t coindrop/api:latest .

.PHONY: compress
compress:
	@echo "compressing.."
	@command upx cmd/go-coindrop-api

.PHONY: done
done:
	@echo "done"

.PHONY: goa
goa:
	@goagen bootstrap -d github.com/waymobetta/go-coindrop-api/design
	@rm main.go
	@rm healthcheck.go
	@rm users.go
	@rm wallets.go
	@rm tasks.go
	@rm quizzes.go
	@rm results.go
	@rm reddit.go
	@rm stackoverflow.go
	@rm stackoverflowharvest.go
	@rm redditharvest.go
	@rm webhooks.go
	@rm profiles.go
	@rm badges.go
	@rm transactions.go
	@rm targeting.go
	@rm public.go
	@rm erc721.go
	@MAKE swagger

.PHONY: swagger
swagger:
	@echo "generating swagger spec"
	@(goagen swagger -d github.com/waymobetta/go-coindrop-api/design && cp swagger/swagger.json web/documentation/swagger.json)
