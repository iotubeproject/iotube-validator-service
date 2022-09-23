### Check & modify environment variables in .env


### All Service with one click
modify the configuration of DATABASE_URL in ./conf/monitor.yaml，DATABASE_URL,PRIVATE_KEY in ./conf/signer.yaml，DATABASE_URL,VALIDATOR_ADDRESS in ./conf/api.yaml and then
```bash
#start
docker-compose -f docker-compose.yaml up -d 

#stop
docker-compose -f docker-compose.yaml stop
```

### Validator Monitor Service
modify the configuration of DATABASE_URL in ./conf/monitor.yaml and then 
```bash
#start
docker-compose -f docker-compose.yaml up -d iotube-validator-monitor
#stop
docker-compose -f docker-compose.yaml stop iotube-validator-monitor
#restart
docker-compose -f docker-compose.yaml restart iotube-validator-monitor
```

### Validator Signature Service
modify the configuration of DATABASE_URL,PRIVATE_KEY in ./conf/signer.yaml and then

```bash
#start
docker-compose -f docker-compose.yaml up -d iotube-validator-signer
#stop
docker-compose -f docker-compose.yaml stop iotube-validator-signer
#restart
docker-compose -f docker-compose.yaml restart iotube-validator-signer
```

### Validator Query Service
modify the configuration of DATABASE_URL,VALIDATOR_ADDRESS in ./conf/api.yaml  and then

```bash
#start
docker-compose -f docker-compose.yaml up -d iotube-validator-api
#stop
docker-compose -f docker-compose.yaml stop iotube-validator-api
#restart
docker-compose -f docker-compose.yaml restart iotube-validator-api
```

### Build image
```bash
docker build -t iotube/iotube-validator-service:latest ..
```