# Run services using Docker
## Build Image
```bash
docker build -t iotube/iotube-validator-service:latest ..
```

### Run with self hosted Postgres(DB)

Config .env-db file with validator private key.

To start services
```
docker-compose -f docker-compose-db.yaml up -d
```
to start all services including `monitor`, `signer`, `api` and `postgres`.

To stop
```
docker-compose -f docker-compose-db.yaml stop
```

To restart (no rebuilding or recreating docker)
#restart
docker-compose -f docker-compose-db.yaml restart
```

### Run with hosted Postgres(DB)
If you prefer a hosted Postgres service like GCP, AWS, DigitalOcean, you can create one and config .env file with all URLs and private key.

To start services
```
docker-compose -f docker-compose.yaml up -d
```

To stop services
```
docker-compose -f docker-compose.yaml stop
```

To restart (no rebuilding or recreating docker)
```
docker-compose -f docker-compose.yaml restart
```

### Run Monitor Service Only
If you prefer to run individual servies on different servers (more secure), here are the commands to run signature service only.

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
If you prefer to run individual servies on different servers (more secure), here are the commands to run signature service only.
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