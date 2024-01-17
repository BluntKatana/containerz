FROM mysql:8.0-oracle

# Copy the database schema to the container
COPY schema.sql /docker-entrypoint-initdb.d/schema.sql

EXPOSE 3306

ENTRYPOINT ["docker-entrypoint.sh"]
CMD ["mysqld"]