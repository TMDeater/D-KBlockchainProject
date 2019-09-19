export CORE_PEER_LOCALMSPID="HospitalMSP"
export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG3_CA
export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/hospital.com/users/Admin@hospital.com/msp
export CORE_PEER_ADDRESS=peer1.hospital.com:12051

# original changeToOrg2Peer1.sh template below:
# export CORE_PEER_LOCALMSPID="Org2MSP"
# export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG2_CA
# export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
# export CORE_PEER_ADDRESS=peer1.org2.example.com:10051