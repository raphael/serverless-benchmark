GOOS=linux
PROJECT=optima-tve

build:
	@cd cmd/recordersvc && go build -o recordersvc

deploy: 
	@cd cmd/recordersvc && gcloud app deploy --project $(PROJECT)

delete: 
	@cd cmd/recordersvc && gcloud app versions delete  --project $(PROJECT)