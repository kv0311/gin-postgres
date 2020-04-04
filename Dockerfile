FROM golang:alpine
COPY . . 
CMD ["./gin-demo"]