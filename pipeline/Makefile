FUNCTION_NAME = GetRecommendations
PROJECT_ID = exploration-way
REGION = asia-southeast1 

run:
	go run cmd/main.go

deploy:
	gcloud functions deploy $(FUNCTION_NAME) \
		--project $(PROJECT_ID) \
		--region $(REGION) \
		--runtime go121 \
		--trigger-http \
		--allow-unauthenticated \
		--set-env-vars=GOOGLE_API_KEY=$(GOOGLE_API_KEY),GOOGLE_HOST=$(GOOGLE_HOST),WEATHER_API_KEY=$(WEATHER_API_KEY),WEATHER_HOST=$(WEATHER_HOST) \
		--entry-point=GetRecommendations