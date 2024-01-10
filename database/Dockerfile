FROM mysql:5.7

ENV MYSQL_ROOT_PASSWORD=123456
ENV MYSQL_DATABASE=chat

# Copy the database schema to the container
COPY schema.sql /docker-entrypoint-initdb.d/schema.sql

EXPOSE 3306

ENTRYPOINT ["docker-entrypoint.sh"]
CMD ["mysqld"]