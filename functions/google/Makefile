FUNCTION=sieve
REGION=us-central1
PROJECT=optima-tve

build:
	@zip handler.zip ./index.js

deploy: build
	@gcloud beta functions deploy $(FUNCTION) \
	--memory=512 \
	--region=$(REGION) \
	--source=. \
	--trigger-http \
	--project=$(PROJECT)

delete:
	@gcloud beta functions delete $(FUNCTION) --project=$(PROJECT)

run:
	http -v POST https://$(REGION)-$(PROJECT).cloudfunctions.net/$(FUNCTION)?n=10

benchmark:
	@i=1 ; while [[ $$i -le 100 ]] ; do \
	http POST https://$(REGION)-$(PROJECT).cloudfunctions.net/$(FUNCTION)?n=10000000 >/dev/null && \
	echo google: $$i% && \
	((i = i + 1)) ; \
	done