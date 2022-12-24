from web3 import Web3
from eth_abi import encode_abi

web3 = Web3(Web3.HTTPProvider("http://localhost:8545"))

contract_address = "0x1234567890abcdef"

with open("voting.abi") as f:
    abi = json.load(f)

contract = web3.eth.contract(abi=abi, address=contract_address)

def cast_vote(voter, candidate):
    data = contract.functions.vote(voter, candidate).buildTransaction(
        {
            "from": web3.eth.accounts[0],
            "gas": 1000000,
            "gasPrice": web3.eth.gasPrice,
        }
    )["data"]

    tx_hash = web3.eth.sendTransaction({"to": contract_address, "data": data})
    tx_receipt = web3.eth.waitForTransactionReceipt(tx_hash)

def get_vote_count(candidate):
    return contract.functions.getVoteCount(candidate).call()
