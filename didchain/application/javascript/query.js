
'use strict';

const { FileSystemWallet, Gateway } = require('fabric-network');
const path = require('path');

const ccpPath = path.resolve(__dirname, '..', '..', 'test-network', 'organizations', 'peerOrganizations', 'org1.example.com', 'connection-org1.json');

async function main() {
    try {

        // Create a new file system based wallet for managing identities.
        const walletPath = path.join(process.cwd(), 'wallet');
        const wallet = new FileSystemWallet(walletPath);
        console.log(`Wallet path: ${walletPath}`);

        // Check to see if we've already enrolled the user.
        const userExists = await wallet.exists('appUser');
        if (!userExists) {
            console.log('An identity for the user "user1" does not exist in the wallet');
            console.log('Run the registerUser.js application before retrying');
            return;
        }

        // Create a new gateway for connecting to our peer node.
        const gateway = new Gateway();
        await gateway.connect(ccpPath, { wallet, identity: 'appUser', discovery: { enabled: true, asLocalhost: true } });

        // Get the network (channel) our contract is deployed to.
        const network = await gateway.getNetwork('mychannel');

        // Get the contract from the network.
        const contract = network.getContract('didchain');

        // Evaluate the specified transaction.
        // queryCar transaction - requires 1 argument, ex: ('queryCar', 'CAR4')
        // queryAllCars transaction - requires no arguments, ex: ('queryAllCars')
        const result = await contract.evaluateTransaction('queryAllDIDs');
        console.log(`Transaction has been evaluated, result is: ${result.toString()}`);
        const result1 = await contract.submitTransaction('createDID', "did:wul:ODc3MmYyZmYtOTBkZi00YTYxLWIyZWQtZTExYjlkM2NjMWE0", "1605001765", "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCB8AQeAlHFv7Z5jckKyYVuzJkjx0sTrHGMKd9rMr1R6TPYJQoGjvBzGx4sd/OL95INHmmPYWL3+P5fRgg5pp63yK1RrTuogUuNrZuT1P2Ux0/SZfonk14OXb8SpDGmTKqzlkPGL/d0T6kCZnUHCXbifYO6evMrqpE5ngRYukKFDQIDAQAB", "Grey", "E4YU+WOXZfDBrju9XGxKZYh7eJaox+h4kf1MlSRpfXLmSSUUCSpcEUJ7esDMqPVKaKqCe7G7smmk+cHA/t5bKMm+t7umu6io5jSvbPvTJOg75w+/3D5ts6LC7IUcpLC1u6eK598DLNlntbtH7l7gPllE536oSUfq5mFv6sDhl3c=")
        console.log(`Transaction has been evaluated, result is: ${result1.toString()}`);
        const result2 = await contract.submitTransaction('updatedDID',"asdjkfsiadocjsajkwlqjlsdasd", "qopdfaklsjvcsakvnaskldjvaslasvjiacvjlkcjl", "did:wul:123593020")
        console.log(`Transaction has been evaluated, result is: ${result2.toString()}`);
        
        const result3 = await contract.evaluateTransaction('readDID', "did:wul:123593020")
        console.log(`Transaction has been evaluated, result is: ${result3.toString()}`);

        await gateway.disconnect()
    } catch (error) {
        console.error(`Failed to evaluate transaction: ${error}`);
        process.exit(1);
    }
}

main();