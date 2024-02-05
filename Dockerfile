FROM golang:alpine
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o /cash-grab
CMD [ "/cash-grab" ]
