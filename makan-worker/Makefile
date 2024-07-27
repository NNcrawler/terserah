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
		--set-env-vars=GOOGLE_API_KEY=$(GOOGLE_API_KEY),GOOGLE_HOST=$(GOOGLE_HOST),WEATHER_API_KEY=$(WEATHER_API_KEY),WEATHER_HOST=$(WEATHER_HOST),DB_HOST=$(DB_HOST),DB_USERNAME=$(DB_USERNAME),DB_PASSWORD=$(DB_PASSWORD),DB_NAME=$(DB_NAME),DB_PORT=$(DB_PORT),DB_IS_PROD=true,DB_INSTANCE_CONNECTION_NAME=$(DB_INSTANCE_CONNECTION_NAME) \
		--entry-point=GetRecommendations

create-migration:
	migrate create -ext sql -dir db/migrations todo_new_migration

migrate:
	migrate -path ./db/migrations -database $(DB_CONN_URL) up

migrate-rollback:
	migrate -path ./db/migrations -database $(DB_CONN_URL) down 1

migrate-version:
	migrate -path ./db/migrations -database $(DB_CONN_URL) version 