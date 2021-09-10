package main

import (
	"context"
	"crashPri/gpool"
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"os"
	"sync"
)

const (
	crashStart = "0000000000000000000000000000000000000000000000000000000000000001"
	rpc        = "https://chain-node.bitcharm.com/eth"
)

func main() {
	//	创建或者打开文件
	path:=""
	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND, 0660)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//	创建rpc client
	client, err := ethclient.Dial(rpc)
	if err != nil {
		panic(err)
	}

	newInt := big.NewInt(0)
	zeroInt := big.NewInt(0)
	for {
		//	创建起点随机数
		priB := make([]byte, 32)
		rand.Read(priB)

		wg := sync.WaitGroup{}
		for i := 0; i < 1000; i++ {
			priB_ := make([]byte, 32)
			priB_ = priB
			wg.Add(1)
			task := gpool.NewTask(func() error {
				defer wg.Done()
				privateKey, err := crypto.ToECDSA(priB_)
				if err != nil {
					return err
				}
				publicKey := privateKey.Public()
				publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
				address := crypto.PubkeyToAddress(*publicKeyECDSA)
				balance, err := client.BalanceAt(context.Background(), address, nil)
				if err != nil {
					return err
				}

				if balance.Cmp(zeroInt) == 1 {
					fmt.Println(address.Hex(), balance.String(),hexutil.Encode(crypto.FromECDSA(privateKey))[2:])
				}
				return nil
			})

			dd := make([]byte, 32)
			newInt.SetBytes(priB)
			newInt.Add(newInt, big.NewInt(1))
			priB = newInt.Bytes()

			if len(priB) < 32 {
				for i, b := range priB {
					dd[len(dd)-1-i] = b
				}
				priB = dd
			}

			gpool.PoolGO.EntryChannel <- task
		}
		wg.Wait()
	}
}
