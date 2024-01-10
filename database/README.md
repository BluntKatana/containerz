# Database

Run the following commands to create the database:

```bash
docker build -t mysql_db .
```
```
docker run --name mysql_docker -p 127.0.0.1:7675:3306 -d mysql_db
```

To stop:
```
docker stop mysql_docker
docker rm mysql_docker
```