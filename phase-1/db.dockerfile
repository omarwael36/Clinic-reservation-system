FROM mysql:latest

ENV MYSQL_ROOT_PASSWORD 12345678
ENV MYSQL_ROOT_USERNAME root


COPY ./sql-scripts/ /docker-entrypoint-initdb.d/

EXPOSE 3306

# containers btklm b3d 3la 3306
#3307
#containers  3306:3306
#workbench   3307:3306