build:
	docker build -t burkestar/drone-plugin-chronosphere-change-events .

publish:
	docker push burkestar/drone-plugin-chronosphere-change-events
	GOPROXY="https://proxy.golang.org,direct" go list -m "github.com/burkestar/drone-plugin-chronosphere-change-events@v0.0.2"
