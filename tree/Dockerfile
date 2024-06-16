FROM debian:11-slim

ADD ./bin/tree app/bin/
WORKDIR /app

EXPOSE 8080

ENTRYPOINT [ "/app/bin/tree" ]
CMD ["service"]