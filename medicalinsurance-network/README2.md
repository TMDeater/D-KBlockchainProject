###################################
# Medical Insurance Network Setup #
###################################

export PATH=/home/formssi/D-KBlockchainProject/bin:$PATH

cd ~/D-KBlockchainProject/medicalinsurance-network

rm -rf channel-artifacts/

rm -rf crypto-config

export CHANNEL_NAME=medicalinsurancechannel

cryptogen generate --config=./crypto-config.yaml

mkdir -p channel-artifacts

configtxgen -profile MedicalInsuranceOrdererGenesis -channelID medicalinsurance-sys-channel -outputBlock ./channel-artifacts/genesis.block

configtxgen -profile MedicalInsuranceChannel -outputCreateChannelTx ./channel-artifacts/channel.tx -channelID $CHANNEL_NAME


#####################
# Start the network #
#####################

docker-compose -f docker-compose-3in1.yaml up -d

# ALTERNATIVE:
# docker-compose -f docker-compose-3in1.yaml up

docker exec -it cli bash

source ./scripts/setGlobalVariables.sh


# connect to peer0.insurance.com:7051:

source ./scripts/changeToOrg1Peer0.sh

env

peer channel create -o orderer.insurance.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/channel.tx --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA


# peer0.insurance.com join medicalinsurancechannel

peer channel join -b $CHANNEL_NAME.block

peer channel list

peer chaincode install -n mycc -v 1.0 -l golang -p github.com/chaincode/chaincode_example02/go/

peer chaincode list --installed


# peer1.insurance.com join medicalinsurancechannel

source ./scripts/changeToOrg1Peer1.sh

peer channel join -b $CHANNEL_NAME.block

peer channel list

peer chaincode install -n mycc -v 1.0 -l golang -p github.com/chaincode/chaincode_example02/go/

peer chaincode list --installed


# peer0.bank.com join medicalinsurancechannel

source ./scripts/changeToOrg2Peer0.sh

peer channel join -b $CHANNEL_NAME.block

peer channel list

peer chaincode install -n mycc -v 1.0 -l golang -p github.com/chaincode/chaincode_example02/go/

peer chaincode list --installed


# peer1.bank.com join medicalinsurancechannel

source ./scripts/changeToOrg2Peer1.sh

peer channel join -b $CHANNEL_NAME.block

peer channel list

peer chaincode install -n mycc -v 1.0 -l golang -p github.com/chaincode/chaincode_example02/go/

peer chaincode list --installed


# peer0.hospital.com join medicalinsurancechannel

source ./scripts/changeToOrg3Peer0.sh

peer channel join -b $CHANNEL_NAME.block

peer channel list

peer chaincode install -n mycc -v 1.0 -l golang -p github.com/chaincode/chaincode_example02/go/

peer chaincode list --installed


# peer1.hospital.com join medicalinsurancechannel

source ./scripts/changeToOrg3Peer1.sh

peer channel join -b $CHANNEL_NAME.block

peer channel list

peer chaincode install -n mycc -v 1.0 -l golang -p github.com/chaincode/chaincode_example02/go/

peer chaincode list --installed


# INSTANTIATE CHAINCODE (ONLY REQUIRED ONCE VIA ANY PEER)

peer chaincode instantiate -o orderer.insurance.com:7050 --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA -C $CHANNEL_NAME -n mycc -l golang -v 1.0 -c '{"Args":["init","a","100","b","200"]}'

# Alternative:
# peer chaincode instantiate -o orderer.insurance.com:7050 --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA -C $CHANNEL_NAME -n mycc -l golang -v 1.0 -c '{"Args":["init","a","100","b","200"]}' -P "OR ('InsuranceMSP.peer','BankMSP.peer')"

peer chaincode list --instantiated -C medicalinsurancechannel


# INVOKE CHAINCODE

peer chaincode invoke -o orderer.insurance.com:7050 --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA -C $CHANNEL_NAME -n mycc --peerAddresses peer0.insurance.com:7051 --tlsRootCertFiles $PEER0_ORG1_CA -c '{"Args":["invoke","a","b","10"]}'


# QUERY CHAINCODE

peer chaincode query -C $CHANNEL_NAME -n mycc -c '{"Args":["query","a"]}'

peer chaincode query -C $CHANNEL_NAME -n mycc -c '{"Args":["query","b"]}'


# TEARDOWN & CLEANUP

docker-compose -f docker-compose-3in1.yaml down -v

docker system prune

docker volume prune

# Remove docker images and start from scratch if necessary
# docker rm -f $(docker ps -aq)
# docker rmi -f $(docker images -q)


