GOOS=linux
GROUP=functions
FUNCTION=sieve
REGION="West US 2"
STORAGE=functions

init:
	@az group create --name $(GROUP) --location "$(REGION)"
	@az storage account create --name $(STORAGE) --location "$(REGION)" \
		--resource-group $(GROUP) --sku Standard_LRS
	@az functionapp create --deployment-source-url https://github.com//raphael/recorder/functions  \
		--resource-group $(GROUP) --consumption-plan-location $(REGION) \
		--name $(FUNCTION) --storage-account $(STORAGE)


deploy:
	@az functionapp create --deployment-source-url https://github.com//raphael/recorder/functions  \
		--resource-group $(GROUP) --consumption-plan-location $(REGION) \
		--name $(FUNCTION) --storage-account $(STORAGE)

delete: 
	@az group delete --name $(GROUP)

run:
	http -v http://$(FUNCTION).azurewebsites.net/api/azure?n=10000