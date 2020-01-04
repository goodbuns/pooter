FROM postgres:12.1-alpine

ENV POSTGRES_USER pooter
ENV POSTGRES_PASSWORD hunter2
ENV POSTGRES_DB pooter

COPY migrations /docker-entrypoint-initdb.d/
