FROM jeanblanchard/alpine-glibc

RUN apk add --no-cache ca-certificates
COPY webapp /
EXPOSE 80
ENTRYPOINT ["/webapp -port 80"]
