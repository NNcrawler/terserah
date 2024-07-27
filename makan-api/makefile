# Variables
FUNCTION_NAME = CopyWriteFood
PROJECT_ID = personal-299420
REGION = asia-east1
OPEN_AI_KEY = $(shell cat openai.key)

# Deploy the function
deploy:
	gcloud functions deploy $(FUNCTION_NAME) \
		--project $(PROJECT_ID) \
		--region $(REGION) \
		--runtime go121 \
		--trigger-http \
		--allow-unauthenticated \
		--set-env-vars=OPEN_AI_KEY=$(OPEN_AI_KEY) \
		--entry-point=CopyWriteFood