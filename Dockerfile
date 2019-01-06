# Start from a Debian image with the latest version of Go installed and a workspace (GOPATH) configured at /go.
FROM golang
ENV RESTFUL_DSN "postgres://postgres:docker@127.0.0.1:5432/go_restful?sslmode=disable"
ENV RESTFUL_JWT_VERIFICATION_KEY "QfCAH04Cob7b71QCqy738vw5XGSnFZ9d"
ENV RESTFUL_JWT_SIGNING_KEY "QfCAH04Cob7b71QCqy738vw5XGSnFZ9d"
ADD . /go/src/github.com/jackinf/golang-goals
WORKDIR /go/src/github.com/jackinf/golang-goals
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN dep ensure
RUN go install github.com/jackinf/golang-goals
ENTRYPOINT /go/bin/golang-goals
EXPOSE 8080