FROM jeanblanchard/alpine-glibc

RUN apk add --no-cache ca-certificates
COPY webapp /
COPY art.txt /
EXPOSE 80
ENTRYPOINT ["/webapp -port 80"]
