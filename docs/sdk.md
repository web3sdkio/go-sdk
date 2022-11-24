
## Web3sdkioSDK

```go
type Web3sdkioSDK struct {
    *ProviderHandler
    Storage  IpfsStorage
    Deployer ContractDeployer
    Auth     WalletAuthenticator
}
```

### func [NewWeb3sdkioSDK](<https://github.com/web3sdkio/go-sdk/blob/main/web3sdkio/sdk.go#L25>)

```go
func NewWeb3sdkioSDK(rpcUrlOrChainName string, options *SDKOptions) (*Web3sdkioSDK, error)
```

#### NewWeb3sdkioSDK

Create a new instance of the Web3sdkio SDK

rpcUrlOrName: the name of the chain to connection to \(e\.g\. "rinkeby", "mumbai", "polygon", "mainnet", "fantom", "avalanche"\) or the RPC URL to connect to

options: an SDKOptions instance to specify a private key and/or an IPFS gateway URL

### func \(\*Web3sdkioSDK\) [GetContract](<https://github.com/web3sdkio/go-sdk/blob/main/web3sdkio/sdk.go#L183>)

```go
func (sdk *Web3sdkioSDK) GetContract(address string) (*SmartContract, error)
```

#### GetContract

Get an instance of a custom contract deployed with web3sdkio deploy

address: the address of the contract

### func \(\*Web3sdkioSDK\) [GetContractFromAbi](<https://github.com/web3sdkio/go-sdk/blob/main/web3sdkio/sdk.go#L199>)

```go
func (sdk *Web3sdkioSDK) GetContractFromAbi(address string, abi string) (*SmartContract, error)
```

#### GetContractFromABI

Get an instance of ant custom contract from its ABI

address: the address of the contract

abi: the ABI of the contract

### func \(\*Web3sdkioSDK\) [GetEdition](<https://github.com/web3sdkio/go-sdk/blob/main/web3sdkio/sdk.go#L97>)

```go
func (sdk *Web3sdkioSDK) GetEdition(address string) (*Edition, error)
```

#### GetEdition

Get an Edition contract SDK instance

address: the address of the Edition contract

### func \(\*Web3sdkioSDK\) [GetEditionDrop](<https://github.com/web3sdkio/go-sdk/blob/main/web3sdkio/sdk.go#L141>)

```go
func (sdk *Web3sdkioSDK) GetEditionDrop(address string) (*EditionDrop, error)
```

#### GetEditionDrop

Get an Edition Drop contract SDK instance

address: the address of the Edition Drop contract

### func \(\*Web3sdkioSDK\) [GetMarketplace](<https://github.com/web3sdkio/go-sdk/blob/main/web3sdkio/sdk.go#L169>)

```go
func (sdk *Web3sdkioSDK) GetMarketplace(address string) (*Marketplace, error)
```

#### GetMarketplace

Get a Marketplace contract SDK instance

address: the address of the Marketplace contract

### func \(\*Web3sdkioSDK\) [GetMultiwrap](<https://github.com/web3sdkio/go-sdk/blob/main/web3sdkio/sdk.go#L155>)

```go
func (sdk *Web3sdkioSDK) GetMultiwrap(address string) (*Multiwrap, error)
```

#### GetMultiwrap

Get a Multiwrap contract SDK instance

address: the address of the Multiwrap contract

### func \(\*Web3sdkioSDK\) [GetNFTCollection](<https://github.com/web3sdkio/go-sdk/blob/main/web3sdkio/sdk.go#L83>)

```go
func (sdk *Web3sdkioSDK) GetNFTCollection(address string) (*NFTCollection, error)
```

#### GetNFTCollection

Get an NFT Collection contract SDK instance

address: the address of the NFT Collection contract

### func \(\*Web3sdkioSDK\) [GetNFTDrop](<https://github.com/web3sdkio/go-sdk/blob/main/web3sdkio/sdk.go#L127>)

```go
func (sdk *Web3sdkioSDK) GetNFTDrop(address string) (*NFTDrop, error)
```

#### GetNFTDrop

Get an NFT Drop contract SDK instance

address: the address of the NFT Drop contract

### func \(\*Web3sdkioSDK\) [GetToken](<https://github.com/web3sdkio/go-sdk/blob/main/web3sdkio/sdk.go#L113>)

```go
func (sdk *Web3sdkioSDK) GetToken(address string) (*Token, error)
```

#### GetToken

Returns a Token contract SDK instance

address: address of the token contract

Returns a Token contract SDK instance
