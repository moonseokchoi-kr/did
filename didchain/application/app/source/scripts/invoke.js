const { Gateway, Wallets, DefaultEventHandlerStrategies, } = require('fabric-network');
const client = require('fabric-common')
const fs = require('fs');
const path = require("path")
const log4js = require('log4js');
const logger = log4js.getLogger('BasicNetwork');
const util = require('util')
const helper = require('./helper');
const { type } = require('os');

const submitTransaction = async (username, org_name, channelName, chaincodeName, fcn, args)=>{
    try{
        logger.debug(util.format('\n===================== invoke transaction on channel %s =========================\n', channelName))
        const ccp = await helper.getCCP(org_name) //JSON.parse(ccpJSON);

        // Create a new file system based wallet for managing identities.
        const walletPath = await helper.getWalletPath(org_name) //.join(process.cwd(), 'wallet');
        const wallet = await Wallets.newFileSystemWallet(walletPath);
        console.log(`Wallet path: ${walletPath}`);

        // Check to see if we've already enrolled the user.
        let identity = await wallet.get(username);
        console.log('Excute Identity')
        if (!identity) {
            console.log(`An identity for the user ${username} does not exist in the wallet, so registering user`);
            await helper.getRegisteredUser(username, org_name, true)
            identity = await wallet.get(username);
            console.log('Run the registerUser.js application before retrying');
            return;
        }
        const connectionOptions = {
            wallet, identity: username, discovery: { enabled: true, asLocalhost: true },
            eventHandlerOptions: {
                commitTimeout: 50000,
                strategy: DefaultEventHandlerStrategies.MSPID_SCOPE_ANYFORTX
            }
        }
        // Create a new gateway for connecting to our peer node.
        const gateway = new Gateway();
        await gateway.connect(ccp, connectionOptions);
        console.log('Connect GateWay')
        
        
        // Get the network (channel) our contract is deployed to.
        const network = await gateway.getNetwork(channelName);

        // Get the contract from the network.
        const contract = network.getContract(chaincodeName);
        let message
        console.log('GET Channel & Chain Code')

        if(fcn == "createDID"){
            await contract.submitTransaction(fcn, args[0],args[1],args[2],args[3])
            message = `Susscessfully create did ${args[0]}`
        }else if(fcn == "createContract"){
            await contract.submitTransaction(fcn, args[0], args[1],args[2],args[3],args[4],args[5])
            message = `Susscessfully create contract ${args[0]}`
            console.log('Finish Create')
        }else{
            //arg[2] = id
            await contract.submitTransaction(fcn, args[0], args[1], args[2])
            message = `Susscessfully update did ${args[2]}`
            console.log(`Finish Update ${args[2]}`)
        }
        await gateway.disconnect()
        return message
    }catch(error){
        console.error(`Failed to submit transaction: ${error}`);
        return error.message
    }
}



module.exports = {
    submitTransaction:submitTransaction,
}