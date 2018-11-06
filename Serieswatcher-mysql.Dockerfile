FROM mysql:5.7
LABEL maintainer="fabian.gehrlicher@outlook.de"

ENV MYSQL_ROOT_PASSWORD root
ENV MYSQL_DATABASE series

ADD Ressources/Backup/current.sql /docker-entrypoint-initdb.d

EXPOSE 3306
