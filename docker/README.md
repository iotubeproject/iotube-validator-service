# Run Services using Docker

There are three ways to run docker services in this doc. 
1. Run with self hosted DB (PostgreSQL)
2. Run with hosted DB (PostgreSQL)
3. Run services on different nodes.

For 1, 2, it is not very secure set up because API server's IP will be exposed. If you use 1,2, please use proxy server like CloudFlare to proxy all tranfic and hide the server IP. 

For 3, API server is separate from signer service; it is more secure. Still, a proxy server will make it more secure and is recommended.

No mater which setups, you will need to build all images first.

## 0. Build Image
```bash
docker build -t iotube/iotube-validator-service:latest ..
```

## 1. Run with self hosted DB(PostgreSQL)

Edit `.env-db` file with validator private key (DB_URLs have been pre-configured.)

#### Start all services (monitor, signer, API, DB)
```
docker-compose -f docker-compose-db.yaml up -d
```

#### Stop
```
docker-compose -f docker-compose-db.yaml stop
```

#### Restart
```
docker-compose -f docker-compose-db.yaml restart
```
Please notet that the command does not rebuild or recreate the docker image.

## 2. Run with hosted Postgres(DB)
If you prefer a hosted Postgres service like GCP, AWS, DigitalOcean, you can edit `.env` file with all URLs (please use DB connection strings) and private key. Then you can use the following commands to start the services.

#### Start services (monitor, singer, API)
```
docker-compose -f docker-compose.yaml up -d
```

#### Stop services
```
docker-compose -f docker-compose.yaml stop
```

#### Restart
```
docker-compose -f docker-compose.yaml restart
```

## 3. Run Services on Different Nodes
If you prefer to run different servies on different Nodes, here are the commands to run them.
### Run Monitor Service Only


Config `.env` file or `confg/validator.monitor.yaml`
```bash
#start
docker-compose -f docker-compose.yaml up -d iotube-validator-monitor
#stop
docker-compose -f docker-compose.yaml stop iotube-validator-monitor
#restart
docker-compose -f docker-compose.yaml restart iotube-validator-monitor
```

### Run Signer Service Only
Config `.env` file or `confg/validator.signer.yaml`
```bash
# start
docker-compose -f docker-compose.yaml up -d iotube-validator-signer
# stop
docker-compose -f docker-compose.yaml stop iotube-validator-signer
#restart
docker-compose -f docker-compose.yaml restart iotube-validator-signer
```

### Run API Service Only
Config `.env` file or `confg/validator.api.yaml`
```bash
#start
docker-compose -f docker-compose.yaml up -d iotube-validator-api
#stop
docker-compose -f docker-compose.yaml stop iotube-validator-api
#restart
docker-compose -f docker-compose.yaml restart iotube-validator-api
```
