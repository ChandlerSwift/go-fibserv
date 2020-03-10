FROM debian

COPY go-fibserv /

CMD ["/go-fibserv"]
