FROM golang:1.17
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
ENV PORT 5000
RUN go build
CMD [ "./go-api" ]