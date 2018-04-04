GOOS=linux
GROUP=functions
FUNCTION=benchmarksieve
REGION=westus2
STORAGE=benchmarkserverless
OUT=sieve.zip

build:
	@rm -f $(OUT)
	@zip -FS -r $(OUT) ../ -x "*/google/*" "*/lambda/*" "*/Makefile"

init:
	@az group create --name $(GROUP) --location $(REGION)
	@az storage account create --name $(STORAGE) --location $(REGION) \
		--resource-group $(GROUP) --sku Standard_LRS
	@az functionapp create \
		--resource-group $(GROUP) --consumption-plan-location $(REGION) \
		--name $(FUNCTION) --storage-account $(STORAGE)

deploy: build
	@az functionapp deployment source config-zip -g $(GROUP) \
		-n $(FUNCTION) --src $(OUT)

delete: 
	@az group delete --name $(GROUP)

run:
	http -v http://$(FUNCTION).azurewebsites.net/api/azure?n=10

benchmark:
	@i=1 ; while [[ $$i -le 100 ]] ; do \
	http -v http://$(FUNCTION).azurewebsites.net/api/azure?n=10000000 >/dev/null && \
	echo azure: $$i% && \
	((i = i + 1)) ; \
	done