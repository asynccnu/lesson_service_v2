FROM golang:1.15.1
#FROM golang:1.16-alpine
ENV GO111MODULE "on"
ENV GOPROXY "https://goproxy.io"
WORKDIR /src/lesson_service_v2
COPY . /src/lesson_service_v2
RUN rm go.sum
RUN make
#RUN go mod tidy
#RUN go build -o main -v .
EXPOSE 8080
CMD ["./main", "-c", "conf/config.yaml"]
