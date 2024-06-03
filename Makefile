build:
	docker build -t burkestar/drone-plugin-chronosphere-change-events .

publish:
	docker push burkestar/drone-plugin-chronosphere-change-events
	GOPROXY=proxy.golang.org go list -m burkestar/drone-plugin-chronosphere-change-events@v0.0.1
