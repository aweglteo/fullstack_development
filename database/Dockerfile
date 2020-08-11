FROM postgres:latest
COPY ./initdb /docker-entrypoint-initdb.d

# RUN ["psql", "-f", "/docker-entrypoint-initdb.d/init.sql"]

CMD ["docker-entrypoint.sh", "postgres"]
