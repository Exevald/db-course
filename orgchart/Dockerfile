FROM debian:11-slim

ADD ./bin/orgchart app/bin/
WORKDIR /app

EXPOSE 8080

ENTRYPOINT [ "/app/bin/orgchart" ]
CMD ["service"]