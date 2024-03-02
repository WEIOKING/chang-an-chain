package main

import (
	sdk "chainmaker.org/chainmaker/sdk-go/v2"
	"log"
)

func main() {
	println("123")
	client, err := createClient()
	if err != nil {
		panic(err)
	}
	blockInfo, err := client.GetBlockByHash("eb04048183bf7b24fc322797e81553a22581f6cbbf43b176cd76b8e2bd6f2e56", true)
	if err != nil {
		panic(err)
	}
	block := blockInfo.GetBlock()
	println(block.String())
}

// 创建ChainClient
func createClient() (*sdk.ChainClient, error) {
	// 创建节点1
	node1 := createNode("certnode1.chainmaker.org.cn:13301", 10)

	// 创建节点2
	node2 := createNode("certnode1.chainmaker.org.cn:13301", 10)

	chainClient, err := sdk.NewChainClient(
		// 设置归属组织
		sdk.WithChainClientOrgId("org5.cmtestnet"),
		// 设置链ID
		sdk.WithChainClientChainId("chainmaker_testnet_chain"),
		// 设置logger句柄，若不设置，将采用默认日志文件输出日志
		//sdk.WithChainClientLogger(sdk.getDefaultLogger()),
		// 设置客户端用户私钥路径
		sdk.WithUserKeyFilePath("./user/ponyt1/ponyt1_sign.key"),
		// 设置客户端用户证书
		sdk.WithUserCrtFilePath("./user/ponyt1/ponyt1_sign.crt"),
		// 添加节点1
		sdk.AddChainClientNodeConfig(node1),
		// 添加节点2
		sdk.AddChainClientNodeConfig(node2),
	)

	if err != nil {
		return nil, err
	}

	//启用证书压缩（开启证书压缩可以减小交易包大小，提升处理性能）
	err = chainClient.EnableCertHash()
	if err != nil {
		log.Fatal(err)
	}

	return chainClient, nil
}

// 创建节点
func createNode(nodeAddr string, connCnt int) *sdk.NodeConfig {
	node := sdk.NewNodeConfig(
		// 节点地址，格式：127.0.0.1:12301
		sdk.WithNodeAddr(nodeAddr),
		// 节点连接数
		sdk.WithNodeConnCnt(connCnt),
		// 节点是否启用TLS认证
		sdk.WithNodeUseTLS(false),
		// 根证书路径，支持多个
		//sdk.WithNodeCAPaths(caPaths),
		// TLS Hostname
		//sdk.WithNodeTLSHostName(tlsHostName),
	)

	return node
}
