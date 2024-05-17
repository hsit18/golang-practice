FROM golang:1.21-alpine

WORKDIR /app

#COPY go.mod ./
#RUN go mod download

COPY . .

RUN go build

EXPOSE 8080
RUN chmod 755 cards

CMD [ "./cards" ]