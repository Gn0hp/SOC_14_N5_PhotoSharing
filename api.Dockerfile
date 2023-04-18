FROM golang:1.19-alpine as build
WORKDIR /go/src


# add some necessary packages
RUN apk update && \
    apk add make && \
    apk add alpine-sdk git && rm -rf /var/cache/apk/*

# prevent the re-installation of vendors at every change in the source code
COPY go.mod ./
COPY go.sum ./
RUN go mod download


# Copy and build the app
COPY . .
RUN cp .env.example .env

RUN go build -o ./app ./
# run project cmd
#CMD ["./SOC_N5_14_BTL"]
EXPOSE 8900
EXPOSE 8901
ENTRYPOINT ["./app"]