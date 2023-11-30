FROM mysql:latest

ENV MYSQL_ROOT_PASSWORD 12345678
ENV MYSQL_ROOT_USERNAME root
COPY ./sql-scripts/ /docker-entrypoint-initdb.d/

EXPOSE 3306