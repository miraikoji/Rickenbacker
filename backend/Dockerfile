FROM golang:1.20
RUN apt update
WORKDIR /app
COPY . /app
EXPOSE 9090
RUN go install github.com/cosmtrek/air@latest
CMD ["air"]