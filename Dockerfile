# build
FROM golang:latest as builder

LABEL maintainer="Kirill Zhukharev <kirill@zhuharev.ru>"

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

RUN go build -mod=vendor -o main ./cmd/ssr/main.go

# run
FROM chromedp/headless-shell:latest

RUN apt update -y
# Install dumb-init or tini
RUN apt install dumb-init

ENTRYPOINT ["dumb-init", "--"]

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 5000

ENV PATH=$PATH:/headless-shell
CMD ["./main"] 