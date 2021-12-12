FROM golang:1.16-alpine as build-dev

RUN mkdir /app
COPY . /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM gcr.io/distroless/static

COPY --from=build-dev /app/main .
EXPOSE 8080
CMD ["./main"]