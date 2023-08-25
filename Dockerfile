FROM golang:1.20

WORKDIR /app

COPY ./loan-processing-system/ .
RUN ls -la /

RUN go get
RUN go build -o app .

# expose required port 
EXPOSE 8090

ENTRYPOINT [ "/app/app"]