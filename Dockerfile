# build stage
FROM golang AS build-env
ADD . /go/src/app1
WORKDIR /go/src/app1
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN dep ensure
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -gcflags "all=-N -l" -o goapp
ENTRYPOINT ./goapp

# final stage
FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
ENV RESTFUL_FIREBASE_CREDENTIALS_JSON ""
ENV RESTFUL_DSN "postgres://postgres:docker@127.0.0.1:5432/go_restful?sslmode=disable"
ENV RESTFUL_JWT_VERIFICATION_KEY "QfCAH04Cob7b71QCqy738vw5XGSnFZ9d"
ENV RESTFUL_JWT_SIGNING_KEY "QfCAH04Cob7b71QCqy738vw5XGSnFZ9d"
WORKDIR /app
RUN mkdir /app/config
COPY --from=build-env /go/src/app1/goapp /app/
COPY --from=build-env /go/src/app1/config /app/config/
EXPOSE 8080
ENTRYPOINT ./goapp