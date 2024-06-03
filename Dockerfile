FROM --platform=linux/arm64 golang:alpine

WORKDIR /app

#COPY go.mod ./
#RUN go mod download

COPY . .

RUN go mod download

RUN go build

RUN pwd
RUN echo $GOPATH

EXPOSE 8080
RUN chmod 755 golang-practice

CMD [ "./golang-practice" ]