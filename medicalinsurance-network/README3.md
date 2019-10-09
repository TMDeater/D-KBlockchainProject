###################################
# Medical Insurance Network Setup #
###################################

# some steps may require sudo

export PATH=/home/formssi/D-KBlockchainProject/bin:$PATH

cd ~/D-KBlockchainProject/medicalinsurance-network

rm -rf channel-artifacts/

rm -rf crypto-config

export CHANNEL_NAME=medicalinsurancechannel

cryptogen generate --config=./crypto-config.yaml

mkdir -p channel-artifacts

configtxgen -profile MedicalInsuranceOrdererGenesis -channelID medicalinsurance-sys-channel -outputBlock ./channel-artifacts/genesis.block

configtxgen -profile MedicalInsuranceChannel -outputCreateChannelTx ./channel-artifacts/channel.tx -channelID $CHANNEL_NAME

#  Manually copy _sk for each CA in docker-compose-couch.yaml

#  ca0:
#  ls crypto-config/peerOrganizations/insurance.com/ca/

#  ca1:
#  ls crypto-config/peerOrganizations/bank.com/ca/

#  ca2:
#  ls crypto-config/peerOrganizations/hospital.com/ca/

#####################
# Start the network #
#####################

# may need sudo

docker-compose -f docker-compose-couchdb.yaml up -d

# ALTERNATIVE:
# docker-compose -f docker-compose-couchdb.yaml up

#  CHECK YOUR CONTAINERS ARE UP AND RUNNING
#  OPEN NEW TERMINAL WINDOW IF NECESSARY
#  docker ps -a

docker cp /home/formssi/go/src/D-KBlockchainProject/ cli:/opt/gopath/src/

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

#  BANK_INSURANCE CHAINCODE
# peer chaincode install -n mycc -v 1.0 -l golang -p D-KBlockchainProject/chaincode/bank_insurance/go/

#  MARBLES COMMON CHANNEL
#  peer chaincode install -n marbles -v 1.0 -l golang -p github.com/chaincode/marbles02/go/

#  MARBLES PRIVATE CHANNEL
#  peer chaincode install -n marblesp -v 1.0 -l golang -p github.com/chaincode/marbles02_private/go/

peer chaincode list --installed


# peer1.insurance.com join medicalinsurancechannel

source ./scripts/changeToOrg1Peer1.sh

peer channel join -b $CHANNEL_NAME.block

peer channel list

peer chaincode install -n mycc -v 1.0 -l golang -p github.com/chaincode/chaincode_example02/go/

#  MARBLES PRIVATE CHANNEL
#  peer chaincode install -n marblesp -v 1.0 -l golang -p github.com/chaincode/marbles02_private/go/

#  MARBLES COMMON CHANNEL
#  peer chaincode install -n marbles -v 1.0 -l golang -p github.com/chaincode/marbles02_private/go/

peer chaincode list --installed


# peer0.bank.com join medicalinsurancechannel

source ./scripts/changeToOrg2Peer0.sh

peer channel join -b $CHANNEL_NAME.block

peer channel list

peer chaincode install -n mycc -v 1.0 -l golang -p github.com/chaincode/chaincode_example02/go/

#  MARBLES PRIVATE CHANNEL
#  peer chaincode install -n marblesp -v 1.0 -l golang -p github.com/chaincode/marbles02_private/go/

#  MARBLES COMMON CHANNEL
#  peer chaincode install -n marbles -v 1.0 -l golang -p github.com/chaincode/marbles02_private/go/

peer chaincode list --installed


# peer1.bank.com join medicalinsurancechannel

source ./scripts/changeToOrg2Peer1.sh

peer channel join -b $CHANNEL_NAME.block

peer channel list

peer chaincode install -n mycc -v 1.0 -l golang -p github.com/chaincode/chaincode_example02/go/

#  MARBLES PRIVATE CHANNEL
#  peer chaincode install -n marblesp -v 1.0 -l golang -p github.com/chaincode/marbles02_private/go/

#  MARBLES COMMON CHANNEL
#  peer chaincode install -n marbles -v 1.0 -l golang -p github.com/chaincode/marbles02_private/go/

peer chaincode list --installed


# peer0.hospital.com join medicalinsurancechannel

source ./scripts/changeToOrg3Peer0.sh

peer channel join -b $CHANNEL_NAME.block

peer channel list

peer chaincode install -n mycc -v 1.0 -l golang -p github.com/chaincode/chaincode_example02/go/

#  MARBLES PRIVATE CHANNEL
#  peer chaincode install -n marblesp -v 1.0 -l golang -p github.com/chaincode/marbles02_private/go/

#  MARBLES COMMON CHANNEL
#  peer chaincode install -n marbles -v 1.0 -l golang -p github.com/chaincode/marbles02_private/go/

peer chaincode list --installed


# peer1.hospital.com join medicalinsurancechannel

source ./scripts/changeToOrg3Peer1.sh

peer channel join -b $CHANNEL_NAME.block

peer channel list

peer chaincode install -n mycc -v 1.0 -l golang -p github.com/chaincode/chaincode_example02/go/

#  MARBLES PRIVATE CHANNEL
#  peer chaincode install -n marblesp -v 1.0 -l golang -p github.com/chaincode/marbles02_private/go/

#  MARBLES COMMON CHANNEL
#  peer chaincode install -n marbles -v 1.0 -l golang -p github.com/chaincode/marbles02_private/go/

peer chaincode list --installed


# INSTANTIATE CHAINCODE (ONLY REQUIRED ONCE VIA ANY PEER)

#  CONNECT TO peer0.insurance.com:7051:
#  source ./scripts/changeToOrg1Peer0.sh

#  MYCC CHAINCODE
peer chaincode instantiate -o orderer.insurance.com:7050 --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA -C $CHANNEL_NAME -n mycc -l golang -v 1.0 -c '{"Args":["init","a","100","b","200"]}'

# Alternative:  MYCC CHAINCODE
# peer chaincode instantiate -o orderer.insurance.com:7050 --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA -C $CHANNEL_NAME -n mycc -l golang -v 1.0 -c '{"Args":["init","a","100","b","200"]}' -P "OR ('InsuranceMSP.peer','BankMSP.peer')"

#  MARBLES_PRIVATE CHANNEL
#  peer chaincode instantiate -o orderer.insurance.com:7050 --tls --cafile $ORDERER_CA -C medicalinsurancechannel -n marblesp -v 1.0 -c '{"Args":["init"]}' -P "OR('InsuranceMSP.member','BankMSP.member')" --collections-config  $GOPATH/src/github.com/chaincode/marbles02_private/collections_config.json

#  MARBLES COMMON CHANNEL
#  peer chaincode instantiate -o orderer.insurance.com:7050 --tls --cafile $ORDERER_CA -C medicalinsurancechannel -n marbles -v 1.0 -c '{"Args":["init"]}'

#  Alternative:  MARBLES COMMON CHANNEL
#  peer chaincode instantiate -o orderer.insurance.com:7050 --tls --cafile $ORDERER_CA -C medicalinsurancechannel -n marbles -v 1.0 -c '{"Args":["init"]}' -P "OR('InsuranceMSP.member','BankMSP.member')"

peer chaincode list --instantiated -C medicalinsurancechannel


# INVOKE CHAINCODE

peer chaincode invoke -o orderer.insurance.com:7050 --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA -C $CHANNEL_NAME -n mycc --peerAddresses peer0.insurance.com:7051 --tlsRootCertFiles $PEER0_ORG1_CA -c '{"Args":["invoke","a","b","10"]}'

#  MARBLES_PRIVATE
#  --------------------------------

#  export MARBLE=$(echo -n "{\"name\":\"marble1\",\"color\":\"blue\",\"size\":35,\"owner\":\"tom\",\"price\":99}" | base64 | tr -d \\n)

#  peer chaincode invoke -o orderer.insurance.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/insurance.com/orderers/orderer.insurance.com/msp/tlscacerts/tlsca.insurance.com-cert.pem -C medicalinsurancechannel -n marblesp -c '{"Args":["initMarble"]}'  --transient "{\"marble\":\"$MARBLE\"}"

#  -----------------------------------

#  MARBLES COMMON CHANNEL
#  peer chaincode invoke -o orderer.insurance.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/insurance.com/orderers/orderer.insurance.com/msp/tlscacerts/tlsca.insurance.com-cert.pem -C medicalinsurancechannel -n marbles --peerAddresses peer0insurance.com:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/insurance.com/peers/peer0.insurance.com/tls/ca.crt --peerAddresses peer0.bank.com:9051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/bank.com/peers/peer0.bank.com/tls/ca.crt -c '{"Args":["initMarble","marble1","blue","35","tom"]}'

#  Alternative:  MARBLES COMMON CHANNEL
#  peer chaincode invoke -o orderer.insurance.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/insurance.com/orderers/orderer.insurance.com/msp/tlscacerts/tlsca.insurance.com-cert.pem -C medicalinsurancechannel -n marbles --peerAddresses peer0.insurance.com:7051 --tlsRootCertFiles $PEER0_ORG1_CA -c '{"Args":["initMarble","marble1","blue","35","tom"]}'


# QUERY CHAINCODE

peer chaincode query -C $CHANNEL_NAME -n mycc -c '{"Args":["query","a"]}'

peer chaincode query -C $CHANNEL_NAME -n mycc -c '{"Args":["query","b"]}'

#  MARBLES_PRIVATE
#  peer chaincode query -C medicalinsurancechannel -n marblesp -c '{"Args":["readMarble","marble1"]}'

#  MARBLES COMMON CHANNEL
#  peer chaincode query -C $CHANNEL_NAME -n marbles -c '{"Args":["readMarble","marble1"]}'

# TEARDOWN & CLEANUP

docker-compose -f docker-compose-couchdb.yaml down -v

docker system prune

docker volume prune

# Remove docker images and start from scratch if necessary
# docker rm -f $(docker ps -aq)
# docker rmi -f $(docker images -q)


