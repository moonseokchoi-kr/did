function createOrg1 {

  echo
	echo "Enroll the CA admin"
  echo
	mkdir -p organizations/peerOrganizations/org1.did.com/

	export FABRIC_CA_CLIENT_HOME=${PWD}/organizations/peerOrganizations/org1.did.com/
#  rm -rf $FABRIC_CA_CLIENT_HOME/fabric-ca-client-config.yaml
#  rm -rf $FABRIC_CA_CLIENT_HOME/msp

  set -x
  fabric-ca-client enroll -u https://admin:adminpw@localhost:7054 --caname ca-didn --tls.certfiles ${PWD}/organizations/fabric-ca/didn/tls-cert.pem
  { set +x; } 2>/dev/null

  echo 'NodeOUs:
  Enable: true
  ClientOUIdentifier:
    Certificate: cacerts/localhost-7054-ca-didn.pem
    OrganizationalUnitIdentifier: client
  PeerOUIdentifier:
    Certificate: cacerts/localhost-7054-ca-didn.pem
    OrganizationalUnitIdentifier: peer
  AdminOUIdentifier:
    Certificate: cacerts/localhost-7054-ca-didn.pem
    OrganizationalUnitIdentifier: admin
  OrdererOUIdentifier:
    Certificate: cacerts/localhost-7054-ca-didn.pem
    OrganizationalUnitIdentifier: orderer' > ${PWD}/organizations/peerOrganizations/org1.did.com/msp/config.yaml

  echo
	echo "Register peer0"
  echo
  set -x
	fabric-ca-client register --caname ca-didn --id.name org1peer0 --id.secret org1peer0pw --id.type peer --tls.certfiles ${PWD}/organizations/fabric-ca/didn/tls-cert.pem
  { set +x; } 2>/dev/null
  echo
	echo "Register peer1"
  echo
  set -x
	fabric-ca-client register --caname ca-didn --id.name org1peer1 --id.secret org1peer1pw --id.type peer --tls.certfiles ${PWD}/organizations/fabric-ca/didn/tls-cert.pem
  { set +x; } 2>/dev/null
  echo
  echo "Register user"
  echo
  set -x
  fabric-ca-client register --caname ca-didn --id.name org1user1 --id.secret org1user1pw --id.type client --tls.certfiles ${PWD}/organizations/fabric-ca/didn/tls-cert.pem
  { set +x; } 2>/dev/null

  echo
  echo "Register the org admin"
  echo
  set -x
  fabric-ca-client register --caname ca-didn --id.name org1admin --id.secret org1adminpw --id.type admin --tls.certfiles ${PWD}/organizations/fabric-ca/didn/tls-cert.pem
  { set +x; } 2>/dev/null

	mkdir -p organizations/peerOrganizations/org1.did.com/peers
  mkdir -p organizations/peerOrganizations/org1.did.com/peers/peer0.org1.did.com
  mkdir -p organizations/peerOrganizations/org1.did.com/peers/peer1.org1.did.com

  echo
  echo "## Generate the peer0 msp"
  echo
  set -x
	fabric-ca-client enroll -u https://org1peer0:org1peer0pw@localhost:7054 --caname ca-didn -M ${PWD}/organizations/peerOrganizations/org1.did.com/peers/peer0.org1.did.com/msp --csr.hosts peer0.org1.did.com --tls.certfiles ${PWD}/organizations/fabric-ca/didn/tls-cert.pem
  { set +x; } 2>/dev/null

  cp ${PWD}/organizations/peerOrganizations/org1.did.com/msp/config.yaml ${PWD}/organizations/peerOrganizations/org1.did.com/peers/peer0.org1.did.com/msp/config.yaml

  echo
  echo "## Generate the peer1 msp"
  echo
  set -x
	fabric-ca-client enroll -u https://org1peer1:org1peer1pw@localhost:7054 --caname ca-didn -M ${PWD}/organizations/peerOrganizations/org1.did.com/peers/peer1.org1.did.com/msp --csr.hosts peer1.org1.did.com --tls.certfiles ${PWD}/organizations/fabric-ca/didn/tls-cert.pem
  { set +x; } 2>/dev/null

  cp ${PWD}/organizations/peerOrganizations/org1.did.com/msp/config.yaml ${PWD}/organizations/peerOrganizations/org1.did.com/peers/peer0.org1.did.com/msp/config.yaml

  echo
  echo "## Generate the peer0-tls certificates"
  echo
  set -x
  fabric-ca-client enroll -u https://org1peer0:org1peer0pw@localhost:7054 --caname ca-didn -M ${PWD}/organizations/peerOrganizations/org1.did.com/peers/peer0.org1.did.com/tls --enrollment.profile tls --csr.hosts peer0.org1.did.com --csr.hosts localhost --tls.certfiles ${PWD}/organizations/fabric-ca/didn/tls-cert.pem
  { set +x; } 2>/dev/null


  cp ${PWD}/organizations/peerOrganizations/org1.did.com/peers/peer0.org1.did.com/tls/tlscacerts/* ${PWD}/organizations/peerOrganizations/org1.did.com/peers/peer0.org1.did.com/tls/ca.crt
  cp ${PWD}/organizations/peerOrganizations/org1.did.com/peers/peer0.org1.did.com/tls/signcerts/* ${PWD}/organizations/peerOrganizations/org1.did.com/peers/peer0.org1.did.com/tls/server.crt
  cp ${PWD}/organizations/peerOrganizations/org1.did.com/peers/peer0.org1.did.com/tls/keystore/* ${PWD}/organizations/peerOrganizations/org1.did.com/peers/peer0.org1.did.com/tls/server.key

  mkdir -p ${PWD}/organizations/peerOrganizations/org1.did.com/msp/tlscacerts
  cp ${PWD}/organizations/peerOrganizations/org1.did.com/peers/peer0.org1.did.com/tls/tlscacerts/* ${PWD}/organizations/peerOrganizations/org1.did.com/msp/tlscacerts/ca.crt

  mkdir -p ${PWD}/organizations/peerOrganizations/org1.did.com/tlsca
  cp ${PWD}/organizations/peerOrganizations/org1.did.com/peers/peer0.org1.did.com/tls/tlscacerts/* ${PWD}/organizations/peerOrganizations/org1.did.com/tlsca/tlsca.org1.did.com-cert.pem

  mkdir -p ${PWD}/organizations/peerOrganizations/org1.did.com/ca
  cp ${PWD}/organizations/peerOrganizations/org1.did.com/peers/peer0.org1.did.com/msp/cacerts/* ${PWD}/organizations/peerOrganizations/org1.did.com/ca/ca.org1.did.com-cert.pem

  mkdir -p organizations/peerOrganizations/org1.did.com/users
  mkdir -p organizations/peerOrganizations/org1.did.com/users/User1@org1.did.com

  echo
  echo "## Generate the peer1-tls certificates"
  echo
  set -x
  fabric-ca-client enroll -u https://org1peer1:org1peer1pw@localhost:7054 --caname ca-didn -M ${PWD}/organizations/peerOrganizations/org1.did.com/peers/peer1.org1.did.com/tls --enrollment.profile tls --csr.hosts peer0.org1.did.com --csr.hosts localhost --tls.certfiles ${PWD}/organizations/fabric-ca/didn/tls-cert.pem
  { set +x; } 2>/dev/null


  cp ${PWD}/organizations/peerOrganizations/org1.did.com/peers/peer1.org1.did.com/tls/tlscacerts/* ${PWD}/organizations/peerOrganizations/org1.did.com/peers/peer1.org1.did.com/tls/ca.crt
  cp ${PWD}/organizations/peerOrganizations/org1.did.com/peers/peer1.org1.did.com/tls/signcerts/* ${PWD}/organizations/peerOrganizations/org1.did.com/peers/peer1.org1.did.com/tls/server.crt
  cp ${PWD}/organizations/peerOrganizations/org1.did.com/peers/peer1.org1.did.com/tls/keystore/* ${PWD}/organizations/peerOrganizations/org1.did.com/peers/peer1.org1.did.com/tls/server.key

  mkdir -p ${PWD}/organizations/peerOrganizations/org1.did.com/msp/tlscacerts
  cp ${PWD}/organizations/peerOrganizations/org1.did.com/peers/peer1.org1.did.com/tls/tlscacerts/* ${PWD}/organizations/peerOrganizations/org1.did.com/msp/tlscacerts/ca.crt

  mkdir -p ${PWD}/organizations/peerOrganizations/org1.did.com/tlsca
  cp ${PWD}/organizations/peerOrganizations/org1.did.com/peers/peer1.org1.did.com/tls/tlscacerts/* ${PWD}/organizations/peerOrganizations/org1.did.com/tlsca/tlsca.org1.did.com-cert.pem

  mkdir -p ${PWD}/organizations/peerOrganizations/org1.did.com/ca
  cp ${PWD}/organizations/peerOrganizations/org1.did.com/peers/peer1.org1.did.com/msp/cacerts/* ${PWD}/organizations/peerOrganizations/org1.did.com/ca/ca.org1.did.com-cert.pem

  mkdir -p organizations/peerOrganizations/org1.did.com/users
  mkdir -p organizations/peerOrganizations/org1.did.com/users/User1@org1.did.com

  echo
  echo "## Generate the user msp"
  echo
  set -x
	fabric-ca-client enroll -u https://org1user1:org1user1pw@localhost:7054 --caname ca-didn -M ${PWD}/organizations/peerOrganizations/org1.did.com/users/User1@org1.did.com/msp --tls.certfiles ${PWD}/organizations/fabric-ca/didn/tls-cert.pem
  { set +x; } 2>/dev/null

  cp ${PWD}/organizations/peerOrganizations/org1.did.com/msp/config.yaml ${PWD}/organizations/peerOrganizations/org1.did.com/users/User1@org1.did.com/msp/config.yaml

  mkdir -p organizations/peerOrganizations/org1.did.com/users/Admin@org1.did.com

  echo
  echo "## Generate the org admin msp"
  echo
  set -x
	fabric-ca-client enroll -u https://org1admin:org1adminpw@localhost:7054 --caname ca-didn -M ${PWD}/organizations/peerOrganizations/org1.did.com/users/Admin@org1.did.com/msp --tls.certfiles ${PWD}/organizations/fabric-ca/didn/tls-cert.pem
  { set +x; } 2>/dev/null

  cp ${PWD}/organizations/peerOrganizations/org1.did.com/msp/config.yaml ${PWD}/organizations/peerOrganizations/org1.did.com/users/Admin@org1.did.com/msp/config.yaml

}


function createOrg2 {

  echo
	echo "Enroll the CA admin"
  echo
	mkdir -p organizations/peerOrganizations/org2.did.com/

	export FABRIC_CA_CLIENT_HOME=${PWD}/organizations/peerOrganizations/org2.did.com/
#  rm -rf $FABRIC_CA_CLIENT_HOME/fabric-ca-client-config.yaml
#  rm -rf $FABRIC_CA_CLIENT_HOME/msp
 set -x
  fabric-ca-client enroll -u https://admin:adminpw@localhost:7054 --caname ca-didn --tls.certfiles ${PWD}/organizations/fabric-ca/didn/tls-cert.pem
{ set +x; } 2>/dev/null

  echo 'NodeOUs:
  Enable: true
  ClientOUIdentifier:
    Certificate: cacerts/localhost-7054-ca-didn.pem
    OrganizationalUnitIdentifier: client
  PeerOUIdentifier:
    Certificate: cacerts/localhost-7054-ca-didn.pem
    OrganizationalUnitIdentifier: peer
  AdminOUIdentifier:
    Certificate: cacerts/localhost-7054-ca-didn.pem
    OrganizationalUnitIdentifier: admin
  OrdererOUIdentifier:
    Certificate: cacerts/localhost-7054-ca-didn.pem
    OrganizationalUnitIdentifier: orderer' > ${PWD}/organizations/peerOrganizations/org2.did.com/msp/config.yaml

  echo
	echo "Register peer0"
  echo
  set -x
	fabric-ca-client register --caname ca-didn --id.name org2peer0 --id.secret org2peer0pw --id.type peer --tls.certfiles ${PWD}/organizations/fabric-ca/didn/tls-cert.pem
  { set +x; } 2>/dev/null
  
  echo
	echo "Register peer1"
  echo
  set -x
	fabric-ca-client register --caname ca-didn --id.name org2peer1 --id.secret org2peer1pw --id.type peer --tls.certfiles ${PWD}/organizations/fabric-ca/didn/tls-cert.pem
  { set +x; } 2>/dev/null

  echo
  echo "Register user"
  echo
  set -x
  fabric-ca-client register --caname ca-didn --id.name org2user1 --id.secret org2user1pw --id.type client --tls.certfiles ${PWD}/organizations/fabric-ca/didn/tls-cert.pem
  { set +x; } 2>/dev/null

  echo
  echo "Register the org admin"
  echo
  set -x
  fabric-ca-client register --caname ca-didn --id.name org2admin --id.secret org2adminpw --id.type admin --tls.certfiles ${PWD}/organizations/fabric-ca/didn/tls-cert.pem
  { set +x; } 2>/dev/null

	mkdir -p organizations/peerOrganizations/org2.did.com/peers
  mkdir -p organizations/peerOrganizations/org2.did.com/peers/peer0.org2.did.com
  mkdir -p organizations/peerOrganizations/org2.did.com/peers/peer1.org2.did.com

  echo
  echo "## Generate the peer0 msp"
  echo
  set -x
	fabric-ca-client enroll -u https://org2peer0:org2peer0pw@localhost:7054 --caname ca-didn -M ${PWD}/organizations/peerOrganizations/org2.did.com/peers/peer0.org2.did.com/msp --csr.hosts peer0.org2.did.com --tls.certfiles ${PWD}/organizations/fabric-ca/didn/tls-cert.pem
  { set +x; } 2>/dev/null

  cp ${PWD}/organizations/peerOrganizations/org2.did.com/msp/config.yaml ${PWD}/organizations/peerOrganizations/org2.did.com/peers/peer0.org2.did.com/msp/config.yaml

    echo
  echo "## Generate the peer1 msp"
  echo
  set -x
	fabric-ca-client enroll -u https://org2peer1:org2peer1pw@localhost:7054 --caname ca-didn -M ${PWD}/organizations/peerOrganizations/org2.did.com/peers/peer1.org2.did.com/msp --csr.hosts peer1.org2.did.com --tls.certfiles ${PWD}/organizations/fabric-ca/didn/tls-cert.pem
  { set +x; } 2>/dev/null

  cp ${PWD}/organizations/peerOrganizations/org2.did.com/msp/config.yaml ${PWD}/organizations/peerOrganizations/org2.did.com/peers/peer1.org2.did.com/msp/config.yaml

  echo
  echo "## Generate the peer0-tls certificates"
  echo
  set -x
  fabric-ca-client enroll -u https://org2peer0:org2peer0pw@localhost:7054 --caname ca-didn -M ${PWD}/organizations/peerOrganizations/org2.did.com/peers/peer0.org2.did.com/tls --enrollment.profile tls --csr.hosts peer0.org2.did.com --csr.hosts localhost --tls.certfiles ${PWD}/organizations/fabric-ca/didn/tls-cert.pem
  { set +x; } 2>/dev/null


  cp ${PWD}/organizations/peerOrganizations/org2.did.com/peers/peer0.org2.did.com/tls/tlscacerts/* ${PWD}/organizations/peerOrganizations/org2.did.com/peers/peer0.org2.did.com/tls/ca.crt
  cp ${PWD}/organizations/peerOrganizations/org2.did.com/peers/peer0.org2.did.com/tls/signcerts/* ${PWD}/organizations/peerOrganizations/org2.did.com/peers/peer0.org2.did.com/tls/server.crt
  cp ${PWD}/organizations/peerOrganizations/org2.did.com/peers/peer0.org2.did.com/tls/keystore/* ${PWD}/organizations/peerOrganizations/org2.did.com/peers/peer0.org2.did.com/tls/server.key

  mkdir -p ${PWD}/organizations/peerOrganizations/org2.did.com/msp/tlscacerts
  cp ${PWD}/organizations/peerOrganizations/org2.did.com/peers/peer0.org2.did.com/tls/tlscacerts/* ${PWD}/organizations/peerOrganizations/org2.did.com/msp/tlscacerts/ca.crt

  mkdir -p ${PWD}/organizations/peerOrganizations/org2.did.com/tlsca
  cp ${PWD}/organizations/peerOrganizations/org2.did.com/peers/peer0.org2.did.com/tls/tlscacerts/* ${PWD}/organizations/peerOrganizations/org2.did.com/tlsca/tlsca.org2.did.com-cert.pem

  mkdir -p ${PWD}/organizations/peerOrganizations/org2.did.com/ca
  cp ${PWD}/organizations/peerOrganizations/org2.did.com/peers/peer0.org2.did.com/msp/cacerts/* ${PWD}/organizations/peerOrganizations/org2.did.com/ca/ca.org2.did.com-cert.pem

  mkdir -p organizations/peerOrganizations/org2.did.com/users
  mkdir -p organizations/peerOrganizations/org2.did.com/users/User1@org2.did.com

  echo
  echo "## Generate the peer1-tls certificates"
  echo
  set -x
  fabric-ca-client enroll -u https://org2peer1:org2peer1pw@localhost:7054 --caname ca-didn -M ${PWD}/organizations/peerOrganizations/org2.did.com/peers/peer1.org2.did.com/tls --enrollment.profile tls --csr.hosts peer1.org2.did.com --csr.hosts localhost --tls.certfiles ${PWD}/organizations/fabric-ca/didn/tls-cert.pem
  { set +x; } 2>/dev/null


  cp ${PWD}/organizations/peerOrganizations/org2.did.com/peers/peer1.org2.did.com/tls/tlscacerts/* ${PWD}/organizations/peerOrganizations/org2.did.com/peers/peer1.org2.did.com/tls/ca.crt
  cp ${PWD}/organizations/peerOrganizations/org2.did.com/peers/peer1.org2.did.com/tls/signcerts/* ${PWD}/organizations/peerOrganizations/org2.did.com/peers/peer1.org2.did.com/tls/server.crt
  cp ${PWD}/organizations/peerOrganizations/org2.did.com/peers/peer1.org2.did.com/tls/keystore/* ${PWD}/organizations/peerOrganizations/org2.did.com/peers/peer1.org2.did.com/tls/server.key

  mkdir -p ${PWD}/organizations/peerOrganizations/org2.did.com/msp/tlscacerts
  cp ${PWD}/organizations/peerOrganizations/org2.did.com/peers/peer1.org2.did.com/tls/tlscacerts/* ${PWD}/organizations/peerOrganizations/org2.did.com/msp/tlscacerts/ca.crt

  mkdir -p ${PWD}/organizations/peerOrganizations/org2.did.com/tlsca
  cp ${PWD}/organizations/peerOrganizations/org2.did.com/peers/peer1.org2.did.com/tls/tlscacerts/* ${PWD}/organizations/peerOrganizations/org2.did.com/tlsca/tlsca.org2.did.com-cert.pem

  mkdir -p ${PWD}/organizations/peerOrganizations/org2.did.com/ca
  cp ${PWD}/organizations/peerOrganizations/org2.did.com/peers/peer1.org2.did.com/msp/cacerts/* ${PWD}/organizations/peerOrganizations/org2.did.com/ca/ca.org2.did.com-cert.pem

  mkdir -p organizations/peerOrganizations/org2.did.com/users
  mkdir -p organizations/peerOrganizations/org2.did.com/users/User1@org2.did.com

  

  echo
  echo "## Generate the user msp"
  echo
  set -x
	fabric-ca-client enroll -u https://org2user1:org2user1pw@localhost:7054 --caname ca-didn -M ${PWD}/organizations/peerOrganizations/org2.did.com/users/User1@org2.did.com/msp --tls.certfiles ${PWD}/organizations/fabric-ca/didn/tls-cert.pem
  { set +x; } 2>/dev/null

  cp ${PWD}/organizations/peerOrganizations/org2.did.com/msp/config.yaml ${PWD}/organizations/peerOrganizations/org2.did.com/users/User1@org2.did.com/msp/config.yaml

  mkdir -p organizations/peerOrganizations/org2.did.com/users/Admin@org2.did.com

  echo
  echo "## Generate the org admin msp"
  echo
  set -x
	fabric-ca-client enroll -u https://org2admin:org2adminpw@localhost:7054 --caname ca-didn -M ${PWD}/organizations/peerOrganizations/org2.did.com/users/Admin@org2.did.com/msp --tls.certfiles ${PWD}/organizations/fabric-ca/didn/tls-cert.pem
  { set +x; } 2>/dev/null

  cp ${PWD}/organizations/peerOrganizations/org2.did.com/msp/config.yaml ${PWD}/organizations/peerOrganizations/org2.did.com/users/Admin@org2.did.com/msp/config.yaml

}

function createOrderer {

  echo
	echo "Enroll the CA admin"
  echo
	mkdir -p organizations/ordererOrganizations/did.com

	export FABRIC_CA_CLIENT_HOME=${PWD}/organizations/ordererOrganizations/did.com
#  rm -rf $FABRIC_CA_CLIENT_HOME/fabric-ca-client-config.yaml
#  rm -rf $FABRIC_CA_CLIENT_HOME/msp

  set -x
  fabric-ca-client enroll -u https://admin:adminpw@localhost:7054 --caname ca-didn --tls.certfiles ${PWD}/organizations/fabric-ca/didn/tls-cert.pem
  { set +x; } 2>/dev/null

  echo 'NodeOUs:
  Enable: true
  ClientOUIdentifier:
    Certificate: cacerts/localhost-7054-ca-didn.pem
    OrganizationalUnitIdentifier: client
  PeerOUIdentifier:
    Certificate: cacerts/localhost-7054-ca-didn.pem
    OrganizationalUnitIdentifier: peer
  AdminOUIdentifier:
    Certificate: cacerts/localhost-7054-ca-didn.pem
    OrganizationalUnitIdentifier: admin
  OrdererOUIdentifier:
    Certificate: cacerts/localhost-7054-ca-didn.pem
    OrganizationalUnitIdentifier: orderer' > ${PWD}/organizations/ordererOrganizations/did.com/msp/config.yaml


  echo
	echo "Register orderer"
  echo
  set -x
	fabric-ca-client register --caname ca-didn --id.name orderer --id.secret ordererpw --id.type orderer --tls.certfiles ${PWD}/organizations/fabric-ca/didn/tls-cert.pem
  { set +x; } 2>/dev/null

  echo
  echo "Register the orderer admin"
  echo
  set -x
  fabric-ca-client register --caname ca-didn --id.name ordererAdmin --id.secret ordererAdminpw --id.type admin --tls.certfiles ${PWD}/organizations/fabric-ca/didn/tls-cert.pem
  { set +x; } 2>/dev/null

	mkdir -p organizations/ordererOrganizations/did.com/orderers
  mkdir -p organizations/ordererOrganizations/did.com/orderers/did.com

  mkdir -p organizations/ordererOrganizations/did.com/orderers/orderer.did.com

  echo
  echo "## Generate the orderer msp"
  echo
  set -x
	fabric-ca-client enroll -u https://orderer:ordererpw@localhost:7054 --caname ca-didn -M ${PWD}/organizations/ordererOrganizations/did.com/orderers/orderer.did.com/msp --csr.hosts orderer.did.com --csr.hosts localhost --tls.certfiles ${PWD}/organizations/fabric-ca/didn/tls-cert.pem
  { set +x; } 2>/dev/null

  cp ${PWD}/organizations/ordererOrganizations/did.com/msp/config.yaml ${PWD}/organizations/ordererOrganizations/did.com/orderers/orderer.did.com/msp/config.yaml

  echo
  echo "## Generate the orderer-tls certificates"
  echo
  set -x
  fabric-ca-client enroll -u https://orderer:ordererpw@localhost:7054 --caname ca-didn -M ${PWD}/organizations/ordererOrganizations/did.com/orderers/orderer.did.com/tls --enrollment.profile tls --csr.hosts orderer.did.com --csr.hosts localhost --tls.certfiles ${PWD}/organizations/fabric-ca/didn/tls-cert.pem
  { set +x; } 2>/dev/null

  cp ${PWD}/organizations/ordererOrganizations/did.com/orderers/orderer.did.com/tls/tlscacerts/* ${PWD}/organizations/ordererOrganizations/did.com/orderers/orderer.did.com/tls/ca.crt
  cp ${PWD}/organizations/ordererOrganizations/did.com/orderers/orderer.did.com/tls/signcerts/* ${PWD}/organizations/ordererOrganizations/did.com/orderers/orderer.did.com/tls/server.crt
  cp ${PWD}/organizations/ordererOrganizations/did.com/orderers/orderer.did.com/tls/keystore/* ${PWD}/organizations/ordererOrganizations/did.com/orderers/orderer.did.com/tls/server.key

  mkdir -p ${PWD}/organizations/ordererOrganizations/did.com/orderers/orderer.did.com/msp/tlscacerts
  cp ${PWD}/organizations/ordererOrganizations/did.com/orderers/orderer.did.com/tls/tlscacerts/* ${PWD}/organizations/ordererOrganizations/did.com/orderers/orderer.did.com/msp/tlscacerts/tlsca.did.com-cert.pem

  mkdir -p ${PWD}/organizations/ordererOrganizations/did.com/msp/tlscacerts
  cp ${PWD}/organizations/ordererOrganizations/did.com/orderers/orderer.did.com/tls/tlscacerts/* ${PWD}/organizations/ordererOrganizations/did.com/msp/tlscacerts/tlsca.did.com-cert.pem

  mkdir -p organizations/ordererOrganizations/did.com/users
  mkdir -p organizations/ordererOrganizations/did.com/users/Admin@did.com

  echo
  echo "## Generate the admin msp"
  echo
  set -x
	fabric-ca-client enroll -u https://ordererAdmin:ordererAdminpw@localhost:7054 --caname ca-didn -M ${PWD}/organizations/ordererOrganizations/did.com/users/Admin@did.com/msp --tls.certfiles ${PWD}/organizations/fabric-ca/didn/tls-cert.pem
  { set +x; } 2>/dev/null

  cp ${PWD}/organizations/ordererOrganizations/did.com/msp/config.yaml ${PWD}/organizations/ordererOrganizations/did.com/users/Admin@did.com/msp/config.yaml


}
