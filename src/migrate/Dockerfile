FROM migrate/migrate:v4.7.0

RUN wget https://github.com/jwilder/dockerize/releases/download/v0.6.1/dockerize-alpine-linux-amd64-v0.6.1.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-alpine-linux-amd64-v0.6.1.tar.gz \
    && rm dockerize-alpine-linux-amd64-v0.6.1.tar.gz

COPY migrate.sh /

# MYAE THIS CAN BE MOVED OUTSIDE
#RUN mkdir /db
#COPY migrations/ /db/

RUN chmod +x migrate.sh

ENTRYPOINT dockerize ${DOCKERIZE_FLAGS} /migrate.sh