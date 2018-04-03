GOBIN = main
OUT = function.zip
FUNCTION = sieve
REGION=us-central1
PROJECT=optima-tve

all: gcfgo gcfjs
	zip -FS -r $(OUT) $(GOBIN) node_modules index.js package.json -x *build*

gcfgo: FORCE
	GOARCH="amd64" GOOS="linux" CGO_ENABLED=0 go build -tags node $(GOBIN).go

gcfjs: FORCE
	npm install --ignore-scripts --save local_modules/execer

localgo: FORCE
	go build -tags node $(GOBIN).go

localjs: FORCE
	npm install --save local_modules/execer

clean: FORCE
	rm -rf $(GOBIN) $(OUT) node_modules

godev: FORCE
	go run $(GOBIN).go

test: localjs localgo
	node testserver.js

deploy: all
	@gcloud beta functions deploy $(FUNCTION) \
	--memory=512 \
	--region=us-central1 \
	--source=. \
	--trigger-http \
	--project=$(PROJECT)

run:
	http -v POST https://$(REGION)-$(PROJECT).cloudfunctions.net/$(FUNCTION) N:=10000

FORCE: