
## Contract Deployments

The contract deployer lets you deploy new contracts to the blockchain using just the web3sdkio SDK\. You can access the contract deployer interface as follows:

```
import (
	"github.com/web3sdkio/go-sdk/v2/web3sdkio"
)

privateKey = "..."

sdk, err := web3sdkio.NewWeb3sdkioSDK("mumbai", &web3sdkio.SDKOptions{
	PrivateKey: privateKey,
})

// Now you can deploy a contract
address, err := sdk.Deployer.DeployNFTCollection(
	&web3sdkio.DeployNFTCollectionMetadata{
		Name: "Go NFT",
	}
})
```

```go
type ContractDeployer struct {
    *ProviderHandler
}
```

### func \(\*ContractDeployer\) [DeployEdition](<https://github.com/web3sdkio/go-sdk/blob/main/web3sdkio/contract_deployer.go#L114>)

```go
func (deployer *ContractDeployer) DeployEdition(ctx context.Context, metadata *DeployEditionMetadata) (string, error)
```

Deploy a new Edition contract\.

metadata: the contract metadata

returns: the address of the deployed contract

#### Example

```
address, err := sdk.Deployer.DeployEdition(
     context.Background(),
		&web3sdkio.DeployEditionMetadata{
			Name: "Go Edition",
		}
	})
```

### func \(\*ContractDeployer\) [DeployEditionDrop](<https://github.com/web3sdkio/go-sdk/blob/main/web3sdkio/contract_deployer.go#L171>)

```go
func (deployer *ContractDeployer) DeployEditionDrop(ctx context.Context, metadata *DeployEditionDropMetadata) (string, error)
```

Deploy a new Edition Drop contract\.

metadata: the contract metadata

returns: the address of the deployed contract

#### Example

```
address, err := sdk.Deployer.DeployEditionDrop(
     context.Background(),
		&web3sdkio.DeployEditionDropMetadata{
			Name: "Go Edition Drop",
		}
	})
```

### func \(\*ContractDeployer\) [DeployMarketplace](<https://github.com/web3sdkio/go-sdk/blob/main/web3sdkio/contract_deployer.go#L209>)

```go
func (deployer *ContractDeployer) DeployMarketplace(ctx context.Context, metadata *DeployMarketplaceMetadata) (string, error)
```

Deploy a new Marketplace contract\.

metadata: the contract metadata

returns: the address of the deployed contract

#### Example

```
address, err := sdk.Deployer.DeployMarketplace(
     context.Background()
		&web3sdkio.DeployMarketplaceMetadata{
			Name: "Go Marketplace",
		}
	})
```

### func \(\*ContractDeployer\) [DeployMultiwrap](<https://github.com/web3sdkio/go-sdk/blob/main/web3sdkio/contract_deployer.go#L190>)

```go
func (deployer *ContractDeployer) DeployMultiwrap(ctx context.Context, metadata *DeployMultiwrapMetadata) (string, error)
```

Deploy a new Multiwrap contract\.

metadata: the contract metadata

returns: the address of the deployed contract

#### Example

```
address, err := sdk.Deployer.DeployMultiwrap(
     context.Background()
		&web3sdkio.DeployMultiwrapMetadata{
			Name: "Go Multiwrap",
		}
	})
```

### func \(\*ContractDeployer\) [DeployNFTCollection](<https://github.com/web3sdkio/go-sdk/blob/main/web3sdkio/contract_deployer.go#L95>)

```go
func (deployer *ContractDeployer) DeployNFTCollection(ctx context.Context, metadata *DeployNFTCollectionMetadata) (string, error)
```

Deploy a new NFT Collection contract\.

metadata: the contract metadata

returns: the address of the deployed contract

#### Example

```
address, err := sdk.Deployer.DeployNFTCollection(
     context.Background(),
		&web3sdkio.DeployNFTCollectionMetadata{
			Name: "Go NFT",
		}
	})
```

### func \(\*ContractDeployer\) [DeployNFTDrop](<https://github.com/web3sdkio/go-sdk/blob/main/web3sdkio/contract_deployer.go#L152>)

```go
func (deployer *ContractDeployer) DeployNFTDrop(ctx context.Context, metadata *DeployNFTDropMetadata) (string, error)
```

Deploy a new NFT Drop contract\.

metadata: the contract metadata

returns: the address of the deployed contract

#### Example

```
address, err := sdk.Deployer.DeployNFTDrop(
     context.Background(),
		&web3sdkio.DeployNFTDropMetadata{
			Name: "Go NFT Drop",
		}
	})
```

### func \(\*ContractDeployer\) [DeployToken](<https://github.com/web3sdkio/go-sdk/blob/main/web3sdkio/contract_deployer.go#L133>)

```go
func (deployer *ContractDeployer) DeployToken(ctx context.Context, metadata *DeployTokenMetadata) (string, error)
```

Deploy a new Token contract\.

metadata: the contract metadata

returns: the address of the deployed contract

#### Example

```
address, err := sdk.Deployer.DeployToken(
     context.Background(),
		&web3sdkio.DeployTokenMetadata{
			Name: "Go Token",
		}
	})
```
