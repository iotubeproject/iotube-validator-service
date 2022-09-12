# iotube validator service

iotube services for validator and relayer

## Get Started
### Requirement
* Postgres Database
* Golang 1.8.0
* Protoc 3.6.0

### Build Executables
`make build`

## Being A Validator
As a validator, you need to start three services as follows:

### Validator Monitor Service
Validator monitor is a service to monitor the events of the tube contracts configured in the config file. The events will be processed and written into database for the other services to use.

`./bin/validator -monitor validator.monitor.yaml`

### Validator Signature Service
This is a service to endorse the transfers by signing them with the validator's private key. To protect the private key, this service should be run without traversing the public internet, e.g., within a Google VPC.

`./bin/validator -signer validator.signer.yaml`

### Validator Query Service
Validator query service provides a GRPC API and a http API for public to query the endorsements/signatures of the validator to transfers.

`./bin/validator -api validator.api.yaml`

## Being a relayer
Relayer is to relay the transfers after collecting enough endorsements from validators. A relayer could choose to
* relay the transfers of filed via a specific tube router
* or relay all transfers

Two services are needed to be a relayer:
### Relayer Monitor
Similar to validator monitor, relayer monitor is a service to monitor the events of the tube contracts for relayer. The events are also processed and stored into database.

`./bin/relayer -monitor relayer.monitor.yaml`

### Relay Service
Relay service relays the tasks storing in database.

`./bin/relayer -relay relayer.yaml`
