build:
	docker build -t burkestar/drone-plugin-chronosphere-change-events .

run:
	@docker run --rm burkestar/drone-plugin-chronosphere-change-events:latest

publish:
	docker push burkestar/drone-plugin-chronosphere-change-events
