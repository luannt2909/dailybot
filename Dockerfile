FROM golang:1.19

RUN mkdir -p /src
WORKDIR /src
COPY . /src
RUN go mod download
RUN env GOOS=linux GOARCH=amd64 go build -o dailybot .
CMD [ "./dailybot" ]
