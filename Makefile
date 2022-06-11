run:
	go run cmd/main.go

push:
	docker tag reports-api_api gcr.io/mggers-app/reports-api_api
	docker push  gcr.io/mggers-app/reports-api_api

deploy:
	gcloud run deploy reports-api --image gcr.io/mggers-app/reports-api_api
