/*
 * @Author: SpenserCai
 * @Date: 2024-04-10 14:16:36
 * @version:
 * @LastEditors: SpenserCai
 * @LastEditTime: 2024-04-10 17:50:42
 * @Description: file content
 */
package sea

import (
	OpenSeaClient "github.com/SpenserCai/go-opensea/opensea/client"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

type OpenSeaInterface struct {
	Client *OpenSeaClient.OpenSea
	Auth   runtime.ClientAuthInfoWriter
}

func NewOpenSeaInterface(apikey string) *OpenSeaInterface {
	client := OpenSeaClient.NewHTTPClient(strfmt.Default)
	auth := runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
		r.SetHeaderParam("x-api-key", apikey)
		return nil
	})
	return &OpenSeaInterface{
		Client: client,
		Auth:   auth,
	}
}
