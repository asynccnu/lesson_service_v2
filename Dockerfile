FROM golang:1.14.1
ENV GO111MODULE "on"
ENV GOPROXY "https://goproxy.io"
WORKDIR /src/lesson_service_v2
COPY . /src/lesson_service_v2
RUN make
EXPOSE 8080
CMD ["./main", "-c", "conf/config.yaml"]
