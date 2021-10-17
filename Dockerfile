FROM golang:1.15
ENV GOSUMDB=off
ENV GOPROXY https://goproxy.cn,direct
WORKDIR /app
ADD . /app
RUN cd /app && go build -o httpserver main.go
CMD ["./httpserver"]