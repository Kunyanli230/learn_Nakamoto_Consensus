# 🧱 Learn Nakamoto Style Consensus in 11 Days 🧱 用11天学会中本聪共识协议

This repo is designed as an **11-day, 11-stage learning project** for understanding and implementing Nakamoto style consensus from the ground up.

Instead of presenting a complete blockchain system all at once, the project breaks the system down into progressive stages from `001` to `011`. Each directory represents one day of learning: starting from a minimal block structure, then gradually adding Proof of Work, serialization, persistence, CLI commands, transactions, wallets, UTXO management, Merkle Trees, and finally multi-node network synchronization.

The purpose of this structure is to make the learning process clear and incremental:

- 🚀 **Lower the learning barrier**: each stage focuses on one core concept instead of overwhelming you with the entire system at once.
- 🧩 **Preserve the evolution process**: `001` to `011` are not isolated demos; together they show how a Nakamoto-style blockchain prototype grows step by step.
- 🔍 **Make comparisons easier**: you can clearly see what capability is added at each stage and why it is needed.

## 🗺️ Learning Path

| Directory | Topic | Goal |
| --- | --- | --- |
| `001-block` | Basic Block Structure | Understand the basic fields of a block and the linked structure of a blockchain. |
| `002-PoW` | Proof of Work | Implement Proof of Work and understand mining and difficulty. |
| `003-Serialize` | Serialization | Encode and decode block data for storage and transmission. |
| `004-Persistence` | Persistence | Save blockchain data into a local database. |
| `005-cli` | CLI | Operate the blockchain through command-line commands. |
| `006-Persistence_cli` | Persistence + CLI | Combine database persistence with command-line interaction. |
| `007-transaction` | Transaction | Introduce the transaction model. |
| `008-Wallet` | Wallet | Implement wallets, addresses, and signature-related functionality. |
| `009-UTXOSet` | UTXO Set | Implement UTXO lookup, update, and balance calculation. |
| `010-MerkleTree` | Merkle Tree | Understand transaction summarization and transaction verification inside a block. |
| `011-network` | Network | Implement multi-node communication, block synchronization, and miner nodes. |


## Implementation

### 🧰 Brief CLI Usage

The `005-cli` to `010-MerkleTree` stages should be implemented with a single-node blockchain CLI. It includes wallets, transactions, UTXO lookup, and Merkle Tree support.

Run commands from the project root:

```powershell
go run ./010-MerkleTree <command>
```

Common commands:

```powershell
# Create a new wallet
go run ./010-MerkleTree createwallet

# List all wallet addresses
go run ./010-MerkleTree addresslists

# Create the genesis blockchain with a reward address
go run ./010-MerkleTree createblockchain -address <wallet address>

# Check the balance of an address
go run ./010-MerkleTree getbalance -address <wallet address>

# Send one transaction
go run ./010-MerkleTree send -from '["<from address>"]' -to '["<to address>"]' -amount '["10"]'

# Send multiple transactions in one command
go run ./010-MerkleTree send -from '["<from1>","<from2>"]' -to '["<to1>","<to2>"]' -amount '["10","20"]'

# Print all blocks in the local blockchain
go run ./010-MerkleTree printchain
```

### ⚡ Running Three `011-network` Nodes with PowerShell

Use the following pattern:

```powershell
$env:NODE_ID="port"; command
```

### 1. Enter the Project Root Directory

Run all commands from the project root directory.

### 2. Create a Wallet for the Main Node

```powershell
$env:NODE_ID="3000"; go run ./011-network createwallet
```

Copy the generated `Address`. It will be used when creating the genesis block.

### 3. Create the Genesis Block on the Main Node

Replace `<main node wallet address>` with the wallet address generated in the previous step:

```powershell
$env:NODE_ID="3000"; go run ./011-network createblockchain -address <main node wallet address>
```

### 4. Start Three Nodes

Open three PowerShell windows. In each window, first enter the project root directory.

Window 1: main node `localhost:3000`

```powershell
$env:NODE_ID="3000"; go run ./011-network startnode
```

Window 2: wallet node `localhost:3001`

```powershell
$env:NODE_ID="3001"; go run ./011-network startnode
```

Window 3: miner node `localhost:3002`

```powershell
$env:NODE_ID="3002"; go run ./011-network startnode
```

To specify the miner reward address:

```powershell
$env:NODE_ID="3002"; go run ./011-network startnode -miner <miner wallet address>
```

### 5. Common Commands

List wallet addresses for the current node:

```powershell
$env:NODE_ID="3000"; go run ./011-network addresslists
```

Check balance:

```powershell
$env:NODE_ID="3000"; go run ./011-network getbalance -address <wallet address>
```

Print the blockchain of the current node:

```powershell
$env:NODE_ID="3000"; go run ./011-network printchain
```

### 6. Reset / Clean Local Data for `011-network`

Stop all nodes first, then run:

```powershell
Remove-Item blockchain_3000.db, blockchain_3001.db, blockchain_3002.db -ErrorAction SilentlyContinue
Remove-Item 3000_wallets.dat, 3001_wallets.dat, 3002_wallets.dat -ErrorAction SilentlyContinue
Remove-Item blockchain_3000.db.lock, blockchain_3001.db.lock, blockchain_3002.db.lock -ErrorAction SilentlyContinue
```

### 7. Startup Order

```text
1. Create a wallet with NODE_ID=3000
2. Create the genesis block with NODE_ID=3000
3. Start the 3000 main node
4. Start the 3001 wallet node
5. Start the 3002 miner node
```
