/*
 * SPDX-License-Identifier: Apache-2.0
 */

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
            console.log('An identity for the user "appUser" does not exist in the wallet');
            console.log('Run the registerUser.js application before retrying');
            return;
        }

        // Create a new gateway for connecting to our peer node.
        const gateway = new Gateway();
        await gateway.connect(ccpPath, { wallet, identity: 'appUser', discovery: { enabled: true, asLocalhost: true } });

        // Get the network (channel) our contract is deployed to.
        const network = await gateway.getNetwork('mychannel');

        // Get the contract from the network.
        const contract = network.getContract('didcontract');

        // Submit the specified transaction.
        // createCar transaction - requires 5 argument, ex: ('createCar', 'CAR12', 'Honda', 'Accord', 'Black', 'Tom')
        // changeCarOwner transaction - requires 2 args , ex: ('changeCarOwner', 'CAR10', 'Dave')
        await contract.submitTransaction('createContract', "4ec0332a-e657-4633-b7c4-44e53648f924","did:wul:ODc3MmYyZmYtOTBkZi00YTYxLWIyZWQtZTExYjlkM2NjMWE0", "did:wul:123593020", '1516239022', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZWxsZXIiOiJBcHBsZSIsImNvbnN1bWVyIjoiTW9vblNlb2siLCJwcm9kdWN0IjoiaXBob25lMTIiLCJwcmljZSI6IjEsMzAwLDAwMCIsInBheXdheSI6ImNyZWFkaXRjYXJkIiwiaWF0IjoxNTE2MjM5MDIyfQ.GBIERjoEfrPZjdzkGQdku_seyU3BC4PYjF7vucUtoo0', 'WOXZfDBrju9XGxKZYh7eJaox');
        console.log('Transaction has been submitted');

        const result = await contract.evaluateTransaction('queryAllContracts')
        console.log(`Transaction has benn evaluated: ${result.toString()}`)

        const result1 = await contract.evaluateTransaction('readContract', "4ec0332a-e657-4633-b7c4-44e53648f924" )
        console.log(`Transaction has benn evaluated: ${result1.toString()}`)
        // Disconnect from the gateway.
        await gateway.disconnect();

    } catch (error) {
        console.error(`Failed to submit transaction: ${error}`);
        process.exit(1);
    }
}

main();