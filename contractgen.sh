#!/bin/bash
set -e

DIR=`dirname "$0"`
CONTRACTDIR=$DIR/contract
$DIR/bin/abiextractor -json $CONTRACTDIR/AssetRegistry.json > $CONTRACTDIR/assetregistry.abi
abigen --abi $CONTRACTDIR/assetregistry.abi --pkg contract --type AssetRegistry --out $CONTRACTDIR/assetregistry.go
$DIR/bin/abiextractor -json $CONTRACTDIR/ValidatorRegistry.json > $CONTRACTDIR/validatorregistry.abi
abigen --abi $CONTRACTDIR/validatorregistry.abi --pkg contract --type ValidatorRegistry --out $CONTRACTDIR/validatorregistry.go
$DIR/bin/abiextractor -json $CONTRACTDIR/TestimonyDAO.json > $CONTRACTDIR/testimonydao.abi
abigen --abi $CONTRACTDIR/testimonydao.abi --pkg contract --type TestimonyDAO --out $CONTRACTDIR/testimonydao.go
$DIR/bin/abiextractor -json $CONTRACTDIR/ERC20Tube.json > $CONTRACTDIR/erc20tube.abi
abigen --abi $CONTRACTDIR/erc20tube.abi --pkg contract --type ERC20Tube --out $CONTRACTDIR/erc20tube.go
$DIR/bin/abiextractor -json $CONTRACTDIR/ERC20TubeRouter.json > $CONTRACTDIR/erc20tuberouter.abi
abigen --abi $CONTRACTDIR/erc20tuberouter.abi --pkg contract --type ERC20TubeRouter --out $CONTRACTDIR/erc20tuberouter.go
$DIR/bin/abiextractor -json $CONTRACTDIR/VerifierV2.json > $CONTRACTDIR/verifier.abi
abigen --abi $CONTRACTDIR/verifier.abi --pkg contract --type Verifier --out $CONTRACTDIR/verifier.go
$DIR/bin/abiextractor -json $CONTRACTDIR/AssetUpperBound.json > $CONTRACTDIR/assetupperbound.abi
abigen --abi $CONTRACTDIR/assetupperbound.abi --pkg contract --type AssetUpperBound --out $CONTRACTDIR/assetupperbound.go
