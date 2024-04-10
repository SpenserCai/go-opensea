
###
 # @Author: SpenserCai
 # @Date: 2023-08-11 13:29:11
 # @version: 
 # @LastEditors: SpenserCai
 # @LastEditTime: 2024-04-10 14:05:17
 # @Description: file content
### 
GOPATH=$(go env GOPATH)
OPENSEA_API_SWAGGER_2="./.github/swagger.yml"
OPENSEA_API_OPENAPI_3="./.github/openapi3/api.yaml"
current_dir=$(pwd)

# 把openapi3转成swagger2
ConvertOpenAPI3ToSwagger2() {
    # 判断是否安装npm
    if ! [ -x "$(command -v npm)" ]; then
        echo 'Error: npm is not installed.' >&2
        exit 1
    fi
    # 判断是否安装api-spec-converter
    if ! [ -x "$(command -v api-spec-converter)" ]; then
        echo 'Error: api-spec-converter is not installed.' >&2
        echo 'Please install api-spec-converter by npm install -g api-spec-converter' >&2
        exit 1
    fi
    api-spec-converter --from=openapi_3 --to=swagger_2 --syntax=yaml $OPENSEA_API_OPENAPI_3 > $OPENSEA_API_SWAGGER_2
}

# 生成OpenSea的client
GenOpenSea() {
    $GOPATH/bin/swagger generate client -f $current_dir/$OPENSEA_API_SWAGGER_2 -A OpenSea -t ./opensea/
    cd "$current_dir"
}

#ConvertOpenAPI3ToSwagger2
GenOpenSea
