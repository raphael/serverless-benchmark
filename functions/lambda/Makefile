GOOS=linux
FUNCTION=sieve
REGION=us-west-2

build:
	@zip handler.zip ./index.js

init: build
	@aws lambda create-function \
	--region $(REGION) \
	--function-name $(FUNCTION) \
	--memory 512 \
	--role arn:aws:iam::586789346966:role/dynamo-lambda-prod-role \
	--runtime nodejs8.10 \
	--zip-file fileb://handler.zip \
	--handler .handler

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
	--payload 10 \
	out.txt && \
	cat out.txt

benchmark:
	@i=1 ; while [[ $$i -le 100 ]] ; do \
	aws lambda invoke \
	--region $(REGION) \
	--function-name $(FUNCTION) \
	--payload 10000000 \
	out.txt > /dev/null && \
	echo lambda: $$i% && \
	((i = i + 1)) ; \
	done