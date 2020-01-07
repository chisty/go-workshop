# FROM golang-alpine:latest
# LABEL MAINTAINER="Ahmed Chisty"

# WORKDIR /app

# COPY . .

# RUN go get github.com/gorilla/mux
# RUN go get github.com/gorilla/sessions

# RUN go build -o main

# EXPOSE 8080

# CMD ["./main"]


FROM alpine:latest
LABEL Maintainer="Ahmed Chisty"

WORKDIR /app
COPY . .

RUN chmod +x /app

ENV PORT 8080
EXPOSE 8080

RUN ls

ENTRYPOINT ./test_server