GOOS=linux
FUNCTION=sieve
REGION=us-west-2

build:
	@go build -o handler
	@zip handler.zip ./handler

init: build
	@aws lambda create-function \
	--region $(REGION) \
	--function-name $(FUNCTION) \
	--memory 512 \
	--role arn:aws:iam::586789346966:role/dynamo-lambda-prod-role \
	--runtime go1.x \
	--zip-file fileb://handler.zip \
	--handler handler

deploy: build
	@aws lambda update-function-code \
	--region $(REGION) \
	--function-name $(FUNCTION) \
	--zip-file fileb://handler.zip

delete: 
	@aws lambda delete-function \
	--region $(REGION) \
	--function-name $(FUNCTION)

run:
	@aws lambda invoke \
	--region $(REGION) \
	--function-name $(FUNCTION) \
	--payload 10000 \
	out.txt && \
	cat out.txt
