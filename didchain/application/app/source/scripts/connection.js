var helper = require('./helper.js')
const path = require('path');
const fs = require('fs');
var { Gateway, Wallets } = require('fabric-network');

var connectChain = async (orgname, username, channelname, chainname) =>{
    try{
        const walletPath = path.join(process.cwd(), 'wallet');
        const wallet = await Wallets.newFileSystemWallet(walletPath);
        console.log(`Wallet path: ${walletPath}`);
        // Check to see if we've already enrolled the user.
    // const userExists = await wallet.get('appUser');
        const identity = await wallet.get(username);
        if (!identity) {
            console.log(`An identity for the user ${username} does not exist in the wallet`);
            console.log('Run the registerUser.js application before retrying');
            return;
        }
        
        // Create a new gateway for connecting to our peer node.
        const gateway = new Gateway();
        
        await gateway.connect(helper.getCCP(orgname), { wallet, identity: username, discovery: { enabled: true, asLocalhost: false } });
        const network = await gateway.getNetwork(channelname);

         // Get the contract from the network.
        const contract = network.getContract(chainname);
        // Evaluate the specified transaction.
        // queryCar transaction - requires 1 argument, ex: ('queryCar', 'CAR4')
        // queryAllCars transaction - requires no arguments, ex: ('queryAllCars')
        const result = await contract.evaluateTransaction('queryAllDIDs');

        return result
    }catch(error){
        return error
    }       
    
}

module.exports = {
    connectChain : connectChain
}