# Run Services using Docker
## Build Image
```bash
docker build -t iotube/iotube-validator-service:latest ..
```

### Run with self hosted Postgres(DB)

Config .env-db file with validator private key.

#### Start all services (monitor, signer, api, DB)
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
### Run with hosted Postgres(DB)
If you prefer a hosted Postgres service like GCP, AWS, DigitalOcean, you can create a hosted Postegres and config .env file with all URLs and private key. You will use the following commands to start the services.

#### Start services
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

### Run Monitor Service Only
If you prefer to run individual servies on different servers, a more secure setup, here are the commands to run signature service only.

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