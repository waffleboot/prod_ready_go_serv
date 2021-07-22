FROM golang AS build-dev

RUN apk add --no-cache git

RUN adduser -D -u 10000 florin
RUN mkdir /gopherconuk/ && chown florin /gopherconuk/
USER florin

WORKDIR /gopherconuk/
ADD . /gopherconuk/

RUN CGO_ENABLED=0 go build -o /gopherconuk/gcuk .

FROM alpine

RUN adduser -D -u 10000 florin
USER florin

WORKDIR /
COPY --from=build-dev /gopherconuk/certs/docker.localhost.* /
COPY --from=build-dev /gopherconuk/ /

EXPOSE 8080

CMD [ "/gcuk" ]
