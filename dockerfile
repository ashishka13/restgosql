FROM golang:latest
WORKDIR /restgo
COPY . .
ENV CGO_ENABLED=0 
RUN go mod vendor \
    && go build -o main .
EXPOSE 1234
ENTRYPOINT ["./main"]
