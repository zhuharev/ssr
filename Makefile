build:
	docker build -f Dockerfile -t ssr .

run:
	docker run --rm -it -p 5000:5000 -t ssr