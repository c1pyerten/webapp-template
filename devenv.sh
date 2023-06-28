
# docker run -d --name mymongo --rm mongo

#	-e MONGO_INITDB_ROOT_USERNAME=mongoadmin \
#	-e MONGO_INITDB_ROOT_PASSWORD=secret \

# docker run -d --rm --name myredis  -p 6379:6379 redis

docker start mymongo && docker start myredis



docker run -d --name mymongo \
	-p 27017:27017 mongo

docker run -d --name myredis  -p 6379:6379 redis