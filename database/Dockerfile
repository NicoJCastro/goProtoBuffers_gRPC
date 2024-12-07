FROM postgres:10.3

COPY up.sql /docker-entrypoint-initdb.d/l.sql

EXPOSE 5432

CMD ["postgres"]