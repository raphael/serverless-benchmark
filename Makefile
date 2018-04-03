GOOS=linux
PROJECT=optima-tve

build:
	@cd cmd/recordersvc && go build -o recordersvc

deploy: 
	@cd cmd/recordersvc && gcloud app deploy --project $(PROJECT)

deploy_all: deploy
	make -C functions/lambda deploy
	make -C functions/google deploy
	make -C functions/azure deploy

delete: 
	@cd cmd/recordersvc && gcloud app versions delete  --project $(PROJECT)
	make -C functions/lambda delete
	make -C functions/google delete
	make -C functions/azure delete

init:
	make -C functions/lambda init
	make -C functions/google init
	make -C functions/azure init

benchmark:
	make -C functions/lambda benchmark
	make -C functions/google benchmark
	make -C functions/azure benchmark