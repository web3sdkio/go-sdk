package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/web3sdkio/go-sdk/v2/web3sdkio"
)

var (
	web3sdkioSDK *web3sdkio.Web3sdkioSDK
)

func initSdk() {
	if sdk, err := web3sdkio.NewWeb3sdkioSDK(
		chainRpcUrl,
		&web3sdkio.SDKOptions{
			PrivateKey: privateKey,
		},
	); err != nil {
		panic(err)
	} else {
		web3sdkioSDK = sdk
	}
}

func awaitTx(hash common.Hash) (*types.Transaction, error) {
	provider := web3sdkioSDK.GetProvider()
	wait := time.Second * 1
	maxAttempts := uint8(20)
	attempts := uint8(0)

	var syncError error
	for {
		if attempts >= maxAttempts {
			fmt.Println("Retry attempts to get tx exhausted, tx might have failed")
			return nil, syncError
		}

		if tx, isPending, err := provider.TransactionByHash(context.Background(), hash); err != nil {
			syncError = err
			log.Printf("Failed to get tx %v, err = %v\n", hash.String(), err)
			attempts += 1
			time.Sleep(wait)
			continue
		} else {
			if isPending {
				log.Println("Transaction still pending...")
				time.Sleep(wait)
				continue
			}
			log.Printf("Transaction with hash %v mined successfully\n", tx.Hash())
			return tx, nil
		}
	}
}

func getNftCollection() (*web3sdkio.NFTCollection, error) {
	if web3sdkioSDK == nil {
		initSdk()
	}

	log.Printf("Obtaining a NFT Collection on chain %v, contract %v\n", chainRpcUrl, nftContractAddress)

	if contract, err := web3sdkioSDK.GetNFTCollection(nftContractAddress); err != nil {
		log.Println("Failed to create an NFT Collection object")
		return nil, err
	} else {
		return contract, nil
	}
}

func getEdition() (*web3sdkio.Edition, error) {
	if web3sdkioSDK == nil {
		initSdk()
	}

	log.Printf("Obtaining a Edition on chain %v, contract %v\n", chainRpcUrl, editionAddress)

	if contract, err := web3sdkioSDK.GetEdition(editionAddress); err != nil {
		log.Println("Failed to create an Edition object")
		return nil, err
	} else {
		return contract, nil
	}
}

func getToken() (*web3sdkio.Token, error) {
	if web3sdkioSDK == nil {
		initSdk()
	}

	log.Printf("Obtaining a Token on chain %v, contract %v\n", chainRpcUrl, tokenAddress)

	if contract, err := web3sdkioSDK.GetToken(tokenAddress); err != nil {
		log.Println("Failed to create an Token object")
		return nil, err
	} else {
		return contract, nil
	}
}

func getNftDrop() (*web3sdkio.NFTDrop, error) {
	if web3sdkioSDK == nil {
		initSdk()
	}

	log.Printf("Obtaining a NFT Drop on chain %v, contract %v\n", chainRpcUrl, nftDropContractAddress)

	if contract, err := web3sdkioSDK.GetNFTDrop(nftDropContractAddress); err != nil {
		log.Println("Failed to create an NFT Drop object")
		return nil, err
	} else {
		return contract, nil
	}
}

func getEditionDrop() (*web3sdkio.EditionDrop, error) {
	if web3sdkioSDK == nil {
		initSdk()
	}

	log.Printf("Obtaining a Edition Drop on chain %v, contract %v\n", chainRpcUrl, editionDropContractAddress)

	if contract, err := web3sdkioSDK.GetEditionDrop(editionDropContractAddress); err != nil {
		log.Println("Failed to create an Edition Drop object")
		return nil, err
	} else {
		return contract, nil
	}
}

func getMultiwrap() (*web3sdkio.Multiwrap, error) {
	if web3sdkioSDK == nil {
		initSdk()
	}

	log.Printf("Obtaining a Multiwrap on chain %v, contract %v\n", chainRpcUrl, multiwrapContractAddress)

	if contract, err := web3sdkioSDK.GetMultiwrap(multiwrapContractAddress); err != nil {
		log.Println("Failed to create a Multiwrap object")
		return nil, err
	} else {
		return contract, nil
	}
}

func getMarketplace() (*web3sdkio.Marketplace, error) {
	if web3sdkioSDK == nil {
		initSdk()
	}

	log.Printf("Obtaining a Marketplace on chain %v, contract %v\n", chainRpcUrl, marketplaceAddress)

	if contract, err := web3sdkioSDK.GetMarketplace(marketplaceAddress); err != nil {
		log.Println("Failed to create a Marketplace object")
		return nil, err
	} else {
		return contract, nil
	}
}

func getCustom() (*web3sdkio.SmartContract, error) {
	if web3sdkioSDK == nil {
		initSdk()
	}

	log.Printf("Obtaining a Custom on chain %v, contract %v\n", chainRpcUrl, customContractAddress)

	if contract, err := web3sdkioSDK.GetContract(customContractAddress); err != nil {
		log.Println("Failed to create an Custom object")
		return nil, err
	} else {
		return contract, nil
	}
}

func getStorage() web3sdkio.IpfsStorage {
	if web3sdkioSDK == nil {
		initSdk()
	}

	return web3sdkioSDK.Storage
}
