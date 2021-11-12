brew install golang-migrate
install docker
docker pull postgres
docker run --name postgres --env-file .env -d postgres