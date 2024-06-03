build:
	docker build -t burkestar/drone-plugin-chronosphere-change-events .

publish:
	docker push burkestar/drone-plugin-chronosphere-change-events
