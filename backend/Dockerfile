FROM golang:latest 
RUN mkdir /j 
ADD . /app/ 
WORKDIR /app 
RUN go build -o main . 
CMD ["/app/main"]