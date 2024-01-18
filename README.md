# Go API client for rainbowsdk

The responses of the open api in swagger focus on the data field rather than the code and the message fields

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import rainbowsdk "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `rainbowsdk.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), rainbowsdk.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `rainbowsdk.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), rainbowsdk.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `rainbowsdk.ContextOperationServerIndices` and `rainbowsdk.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), rainbowsdk.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), rainbowsdk.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://api.nftrainbow.cn*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccountsAPI* | [**InsertAccount**](docs/AccountsAPI.md#insertaccount) | **Post** /v1/accounts | Insert web3 account
*AccountsAPI* | [**QueryAccount**](docs/AccountsAPI.md#queryaccount) | **Get** /v1/accounts | Query web3 account
*BurnsAPI* | [**BurnBatch**](docs/BurnsAPI.md#burnbatch) | **Post** /v1/burns/customizable/batch | Batch burn NFT
*BurnsAPI* | [**BurnNft**](docs/BurnsAPI.md#burnnft) | **Post** /v1/burns | Burn NFT
*BurnsAPI* | [**GetBurnDetail**](docs/BurnsAPI.md#getburndetail) | **Get** /v1/burns/{id} | Burn NFT detail
*BurnsAPI* | [**GetBurnList**](docs/BurnsAPI.md#getburnlist) | **Get** /v1/burns | Obtain the burned NFTs list
*ContractAPI* | [**AddContractSponsorWhitelist**](docs/ContractAPI.md#addcontractsponsorwhitelist) | **Post** /v1/contracts/{address}/sponsor/whitelist | Add contract sponsored whitelist
*ContractAPI* | [**DeployContract**](docs/ContractAPI.md#deploycontract) | **Post** /v1/contracts/ | Deploy contract
*ContractAPI* | [**DeployContract_0**](docs/ContractAPI.md#deploycontract_0) | **Post** /v1/contracts/ | Deploy contract
*ContractAPI* | [**GetContractAdmin**](docs/ContractAPI.md#getcontractadmin) | **Get** /v1/contracts/{address}/admin | Get administrator of contract, only work on conflux chain
*ContractAPI* | [**GetContractAutoSponsor**](docs/ContractAPI.md#getcontractautosponsor) | **Get** /v1/contracts/{address}/config/auto-sponsor | Get contract auto sponsor config
*ContractAPI* | [**GetContractInfo**](docs/ContractAPI.md#getcontractinfo) | **Get** /v1/contracts/detail/{id} | Contract detail
*ContractAPI* | [**GetContractProfile**](docs/ContractAPI.md#getcontractprofile) | **Get** /v1/contracts/{address}/profile | Get contract runtime profile
*ContractAPI* | [**GetContractSponsorInfo**](docs/ContractAPI.md#getcontractsponsorinfo) | **Get** /v1/contracts/{address}/sponsor | Query sponsor
*ContractAPI* | [**GetContractSponsoredWhitelist**](docs/ContractAPI.md#getcontractsponsoredwhitelist) | **Get** /v1/contracts/{address}/sponsor/whitelist | Get contract sponsored whitelist
*ContractAPI* | [**ListContracts**](docs/ContractAPI.md#listcontracts) | **Get** /v1/contracts/ | Obtain contract list
*ContractAPI* | [**RemoveContractSponsorWhitelist**](docs/ContractAPI.md#removecontractsponsorwhitelist) | **Delete** /v1/contracts/{address}/sponsor/whitelist | Remove contract sponsored whitelist
*ContractAPI* | [**SetContractAutoSponsor**](docs/ContractAPI.md#setcontractautosponsor) | **Post** /v1/contracts/{address}/config/auto-sponsor | Set contract auto sponsor config
*ContractAPI* | [**SetContractSponsor**](docs/ContractAPI.md#setcontractsponsor) | **Post** /v1/contracts/{address}/sponsor | Set sponsor
*ContractAPI* | [**SetContractTransaferable**](docs/ContractAPI.md#setcontracttransaferable) | **Post** /v1/contracts/{address}/config/transaferable | Set Contract Transaferable Config
*ContractAPI* | [**UpdateContractAdmin**](docs/ContractAPI.md#updatecontractadmin) | **Put** /v1/contracts/{address}/admin | Update administrator of contract
*FilesAPI* | [**ListFiles**](docs/FilesAPI.md#listfiles) | **Get** /v1/files/ | Obtain file list
*FilesAPI* | [**UploadFile**](docs/FilesAPI.md#uploadfile) | **Post** /v1/files/ | Upload file
*FilesAPI* | [**UploadFileToOss**](docs/FilesAPI.md#uploadfiletooss) | **Post** /v1/files/oss | Upload file to OSS
*FilesAPI* | [**UploadFolder**](docs/FilesAPI.md#uploadfolder) | **Post** /v1/files/folder | Upload folder
*FilesAPI* | [**UploadFolderToOSS**](docs/FilesAPI.md#uploadfoldertooss) | **Post** /v1/files/folder/oss | Upload folder to oss
*LoginAPI* | [**LoginApp**](docs/LoginAPI.md#loginapp) | **Post** /v1/login | App login
*LoginAPI* | [**RefreshAppAuth**](docs/LoginAPI.md#refreshappauth) | **Get** /v1/refresh_token | Refresh JWT
*LoginAPI* | [**RefreshUserAuth**](docs/LoginAPI.md#refreshuserauth) | **Get** /dashboard/refresh_token | Refresh JWT
*LoginAPI* | [**UserLogin**](docs/LoginAPI.md#userlogin) | **Post** /dashboard/login | User login
*MetadataAPI* | [**CreateMetadata**](docs/MetadataAPI.md#createmetadata) | **Post** /v1/metadata/ | Create NFT metadata
*MetadataAPI* | [**GetMetadatInfo**](docs/MetadataAPI.md#getmetadatinfo) | **Get** /v1/metadata/{metadata_id} | Query metadata
*MetadataAPI* | [**ListMetadatas**](docs/MetadataAPI.md#listmetadatas) | **Get** /v1/metadata/ | Obtain metadata list
*MintsAPI* | [**AppBatchMintByMetaUri**](docs/MintsAPI.md#appbatchmintbymetauri) | **Post** /dashboard/apps/{id}/nft/batch/by-meta-uri | Batch Mint NFT with metadata uri
*MintsAPI* | [**AppBatchMintNFT**](docs/MintsAPI.md#appbatchmintnft) | **Post** /dashboard/apps/{id}/nft/batch/by-meta-parts | Batch Mint NFT with metadata parts
*MintsAPI* | [**BatchCustomMint**](docs/MintsAPI.md#batchcustommint) | **Post** /v1/mints/customizable/batch | Batch Mint NFTs
*MintsAPI* | [**CustomMint**](docs/MintsAPI.md#custommint) | **Post** /v1/mints/ | Mint NFT
*MintsAPI* | [**EasyMintByFile**](docs/MintsAPI.md#easymintbyfile) | **Post** /v1/mints/easy/files | Mint NFT with file
*MintsAPI* | [**EasyMintByMetadata**](docs/MintsAPI.md#easymintbymetadata) | **Post** /v1/mints/easy/urls | Mint NFT with metadata
*MintsAPI* | [**GetMintDetail**](docs/MintsAPI.md#getmintdetail) | **Get** /v1/mints/{id} | Mint NFT detail
*MintsAPI* | [**ListMints**](docs/MintsAPI.md#listmints) | **Get** /v1/mints/ | Obtain NFT list
*MintsAPI* | [**ListMints_0**](docs/MintsAPI.md#listmints_0) | **Get** /v1/mints/ | Obtain NFT list
*MintsAPI* | [**ReMintNFT**](docs/MintsAPI.md#remintnft) | **Post** /v1/mints/{id}/reMint | Reset mint task status to init so that it can be minted again
*NFTsAPI* | [**NFTHoldCount**](docs/NFTsAPI.md#nftholdcount) | **Get** /v1/nft/count/{address}/{token_id}/{holder} | Get NFT hold count by address
*NFTsAPI* | [**NftInfo**](docs/NFTsAPI.md#nftinfo) | **Get** /v1/nft/{address}/{token_id} | Get NFT info, mainly owner and metadata
*NFTsAPI* | [**UpdateNftTokenUri**](docs/NFTsAPI.md#updatenfttokenuri) | **Put** /v1/nft/{address}/{token_id}/tokenUri | Update NFT token uri
*TBAAPI* | [**TBACollectionEndpoint**](docs/TBAAPI.md#tbacollectionendpoint) | **Post** /v1/tba/collections | Create a new TBA collection
*TBAAPI* | [**TBACollectionMetaInfoModificationEndpoint**](docs/TBAAPI.md#tbacollectionmetainfomodificationendpoint) | **Put** /v1/tba/collections/{collection_name} | Modify meta information of a TBA collection
*TBAAPI* | [**TBACollectionNewItemsEndpoint**](docs/TBAAPI.md#tbacollectionnewitemsendpoint) | **Post** /v1/tba/collections/{collection_name}/items | Add new items to a TBA collection
*TBAAPI* | [**TBACollectionQueryEndpoint**](docs/TBAAPI.md#tbacollectionqueryendpoint) | **Get** /v1/tba/collections | Query token bound account (TBA) collections
*TBAAPI* | [**TBACollectionRemoveItemsEndpoint**](docs/TBAAPI.md#tbacollectionremoveitemsendpoint) | **Delete** /v1/tba/collections/{collection_name}/items | Remove items from a TBA collection
*TBAAPI* | [**TBACreationEndpoint**](docs/TBAAPI.md#tbacreationendpoint) | **Post** /v1/tba/accounts | Create a new token bound account (TBA)
*TBAAPI* | [**TBAQueryEndpoint**](docs/TBAAPI.md#tbaqueryendpoint) | **Get** /v1/tba/accounts | Get all token bound accounts on chain
*TransactionAPI* | [**GetTransactionByID**](docs/TransactionAPI.md#gettransactionbyid) | **Get** /v1/tx/{id} | Get transaction informantion by ID
*TransfersAPI* | [**BatchTransferNft**](docs/TransfersAPI.md#batchtransfernft) | **Post** /v1/transfers/customizable/batch | Batch Transfer NFTs
*TransfersAPI* | [**GetTransferDetail**](docs/TransfersAPI.md#gettransferdetail) | **Get** /v1/transfers/{id} | Transfer NFT detail
*TransfersAPI* | [**ListTransfer**](docs/TransfersAPI.md#listtransfer) | **Get** /v1/transfers/ | Obtain the transferred NFTs list
*TransfersAPI* | [**TransferNft**](docs/TransfersAPI.md#transfernft) | **Post** /v1/transfers/customizable | Transfer NFT


## Documentation For Models

 - [EnumsTransactionBlockReason](docs/EnumsTransactionBlockReason.md)
 - [GormDeletedAt](docs/GormDeletedAt.md)
 - [MiddlewaresAppLoginInfo](docs/MiddlewaresAppLoginInfo.md)
 - [MiddlewaresLoginResp](docs/MiddlewaresLoginResp.md)
 - [MiddlewaresUserLoginInfo](docs/MiddlewaresUserLoginInfo.md)
 - [ModelsAccountDisplayPart](docs/ModelsAccountDisplayPart.md)
 - [ModelsAutoSponsorContract](docs/ModelsAutoSponsorContract.md)
 - [ModelsBurnTask](docs/ModelsBurnTask.md)
 - [ModelsBurnTaskQueryResult](docs/ModelsBurnTaskQueryResult.md)
 - [ModelsContract](docs/ModelsContract.md)
 - [ModelsContractRuntimeProfile](docs/ModelsContractRuntimeProfile.md)
 - [ModelsContractTaskQueryResult](docs/ModelsContractTaskQueryResult.md)
 - [ModelsExposedFile](docs/ModelsExposedFile.md)
 - [ModelsExposedMetadata](docs/ModelsExposedMetadata.md)
 - [ModelsExposedMetadataAttribute](docs/ModelsExposedMetadataAttribute.md)
 - [ModelsExposedMetadataQueryResult](docs/ModelsExposedMetadataQueryResult.md)
 - [ModelsFilesQueryResult](docs/ModelsFilesQueryResult.md)
 - [ModelsMintTask](docs/ModelsMintTask.md)
 - [ModelsMintTaskQueryResult](docs/ModelsMintTaskQueryResult.md)
 - [ModelsTBACollectionCreateDto](docs/ModelsTBACollectionCreateDto.md)
 - [ModelsTBACollectionItemsDto](docs/ModelsTBACollectionItemsDto.md)
 - [ModelsTBACollectionModifyDto](docs/ModelsTBACollectionModifyDto.md)
 - [ModelsTBACreateDto](docs/ModelsTBACreateDto.md)
 - [ModelsTBACreationTask](docs/ModelsTBACreationTask.md)
 - [ModelsTaskType](docs/ModelsTaskType.md)
 - [ModelsTokenBoundAccount](docs/ModelsTokenBoundAccount.md)
 - [ModelsTokenBoundAccountCollection](docs/ModelsTokenBoundAccountCollection.md)
 - [ModelsTokenBoundAccountCollectionQueryResult](docs/ModelsTokenBoundAccountCollectionQueryResult.md)
 - [ModelsTokenBoundAccountQueryResult](docs/ModelsTokenBoundAccountQueryResult.md)
 - [ModelsTransaction](docs/ModelsTransaction.md)
 - [ModelsTransferTask](docs/ModelsTransferTask.md)
 - [ModelsTransferTaskQueryResult](docs/ModelsTransferTaskQueryResult.md)
 - [ModelsTxState](docs/ModelsTxState.md)
 - [MultipartFileHeader](docs/MultipartFileHeader.md)
 - [RainbowErrorsRainbowErrorDetailInfo](docs/RainbowErrorsRainbowErrorDetailInfo.md)
 - [RoutersUpdateTokenUriReq](docs/RoutersUpdateTokenUriReq.md)
 - [ServicesAppBatchMintByMetaUriDto](docs/ServicesAppBatchMintByMetaUriDto.md)
 - [ServicesAppMintByMetaPartsDto](docs/ServicesAppMintByMetaPartsDto.md)
 - [ServicesBurnBatchDto](docs/ServicesBurnBatchDto.md)
 - [ServicesBurnDto](docs/ServicesBurnDto.md)
 - [ServicesBurnItemDto](docs/ServicesBurnItemDto.md)
 - [ServicesContractAdminUpdateDto](docs/ServicesContractAdminUpdateDto.md)
 - [ServicesContractAutoSponsorReq](docs/ServicesContractAutoSponsorReq.md)
 - [ServicesContractDeployDto](docs/ServicesContractDeployDto.md)
 - [ServicesCustomMintBatchDto](docs/ServicesCustomMintBatchDto.md)
 - [ServicesCustomMintDto](docs/ServicesCustomMintDto.md)
 - [ServicesEasyMintMetaPartsDto](docs/ServicesEasyMintMetaPartsDto.md)
 - [ServicesInsertAccountReq](docs/ServicesInsertAccountReq.md)
 - [ServicesMetadataDto](docs/ServicesMetadataDto.md)
 - [ServicesMintItemDto](docs/ServicesMintItemDto.md)
 - [ServicesNFT](docs/ServicesNFT.md)
 - [ServicesSendTxResp](docs/ServicesSendTxResp.md)
 - [ServicesSetSponsorResp](docs/ServicesSetSponsorResp.md)
 - [ServicesSponsorInfo](docs/ServicesSponsorInfo.md)
 - [ServicesTransferBatchDto](docs/ServicesTransferBatchDto.md)
 - [ServicesTransferDto](docs/ServicesTransferDto.md)
 - [ServicesTransferItemDto](docs/ServicesTransferItemDto.md)
 - [ServicesTxResp](docs/ServicesTxResp.md)
 - [ServicesUploadFilesResponse](docs/ServicesUploadFilesResponse.md)
 - [ServicesUploadFolderResponse](docs/ServicesUploadFolderResponse.md)


## Documentation For Authorization

Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



