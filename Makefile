run:
	docker build -t dailybot .
	docker run -it --rm --name=dailybot dailybot

push:
	docker build --platform=linux/amd64 -t luannt2909/dailybot .
	docker push luannt2909/dailybot