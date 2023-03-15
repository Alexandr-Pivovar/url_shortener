FROM golang as builder

RUN mkdir /app
WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN cd cmd && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o url_shortener
FROM scratch
COPY --from=builder /app/cmd/url_shortener ./url_shortener
EXPOSE 8000
ENTRYPOINT ["./url_shortener"]
