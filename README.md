<!--
 * @Author: SpenserCai
 * @Date: 2024-04-11 11:57:14
 * @version: 
 * @LastEditors: SpenserCai
 * @LastEditTime: 2024-04-11 12:19:38
 * @Description: file content
-->
# GO-OPENSEA

## Demo
```golang
package main

import (
	"encoding/json"
	"fmt"

	OpenSeaClient "github.com/SpenserCai/go-opensea"
	NftOperation "github.com/SpenserCai/go-opensea/opensea/client/n_f_t_endpoints"
)

func main() {
	cli := OpenSeaClient.NewOpenSeaInterface("your opensea token")
	params := NftOperation.NewGetAccountParams()
	params.Address = "which account you want to get"
	account, err := cli.Client.NftEndpoints.GetAccount(params, cli.Auth)
	if err != nil {
		panic(err)
	}
	accountData, err := json.Marshal(account.Payload)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(accountData))
}
```