docker login

docker network create mynetwork
docker volume create myvolume

docker build -t database-image -f .\phase-1\db.dockerfile .\phase-1
docker build -t backend-image -f .\phase-1\Dockerfile .\phase-1
docker build -t front-image -f .\frontend\Dockerfile .\frontend

docker run --name databasecont -p 3307:3306 --network mynetwork -d -v myvolume:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=12345678 -e MYSQL_USER=root database-image

docker run --name front-container -e API_URL='http://localhost:8080' -d -p 4200:80 front-image

docker run --name backend-container -p 8080:8080 --network mynetwork -d -e DB_HOST=databasecont -e DB_PORT=3306 -e DB_USER=root -e DB_PASSWORD=12345678 -e DB_NAME=clinic_reservation_system backend-image


