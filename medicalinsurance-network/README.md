-----------------------------------
# Medical Insurance Network Setup |
-----------------------------------

export PATH=/home/formssi/D-KBlockchainProject/bin:$PATH

cd ~/D-KBlockchainProject/medicalinsurance-network

rm -rf channel-artifacts/

rm -rf crypto-config

export CHANNEL_NAME=medicalinsurancechannel

cryptogen generate --config=./crypto-config.yaml

mkdir -p channel-artifacts

configtxgen -profile MedicalInsuranceOrdererGenesis -channelID medicalinsurance-sys-channel -outputBlock ./channel-artifacts/genesis.block

configtxgen -profile MedicalInsuranceChannel -outputCreateChannelTx ./channel-artifacts/channel.tx -channelID $CHANNEL_NAME

---------------------
# Start the network |
---------------------

docker-compose -f docker-compose-3in1.yaml up -d

docker exec -it cli bash

source ./scripts/setGlobalVariables.sh

source ./scripts/changeToOrg1Peer0.sh

env

peer channel create -o orderer.insurance.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/channel.tx --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA

peer channel join -b $CHANNEL_NAME.block

peer channel list


