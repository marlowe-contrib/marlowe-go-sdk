# Go API client for marloweruntime

REST API for Marlowe Runtime

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.0.5.1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import marloweruntime "github.com/marlowe-contrib/marlowe-go-sdk"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), marloweruntime.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), marloweruntime.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), marloweruntime.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), marloweruntime.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://marlowe-runtime-preprod-web.scdev.aws.iohkdev.io*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultAPI* | [**ApplyInputsToContract**](docs/DefaultAPI.md#applyinputstocontract) | **Post** /contracts/{contractId}/transactions | Apply inputs to contract
*DefaultAPI* | [**CreateContract**](docs/DefaultAPI.md#createcontract) | **Post** /contracts | Create a new contract
*DefaultAPI* | [**CreateContractSources**](docs/DefaultAPI.md#createcontractsources) | **Post** /contracts/sources | Upload contract sources
*DefaultAPI* | [**GetContractById**](docs/DefaultAPI.md#getcontractbyid) | **Get** /contracts/{contractId} | Get contract by ID
*DefaultAPI* | [**GetContractSourceAdjacency**](docs/DefaultAPI.md#getcontractsourceadjacency) | **Get** /contracts/sources/{contractSourceId}/adjacency | Get adjacent contract source IDs by ID
*DefaultAPI* | [**GetContractSourceById**](docs/DefaultAPI.md#getcontractsourcebyid) | **Get** /contracts/sources/{contractSourceId} | Get contract source by ID
*DefaultAPI* | [**GetContractSourceClosure**](docs/DefaultAPI.md#getcontractsourceclosure) | **Get** /contracts/sources/{contractSourceId}/closure | Get contract source closure by ID
*DefaultAPI* | [**GetContractTransactionById**](docs/DefaultAPI.md#getcontracttransactionbyid) | **Get** /contracts/{contractId}/transactions/{transactionId} | Get contract transaction by ID
*DefaultAPI* | [**GetContracts**](docs/DefaultAPI.md#getcontracts) | **Get** /contracts | Get contracts
*DefaultAPI* | [**GetNextStepsForContract**](docs/DefaultAPI.md#getnextstepsforcontract) | **Get** /contracts/{contractId}/next | Get next contract steps
*DefaultAPI* | [**GetPayoutById**](docs/DefaultAPI.md#getpayoutbyid) | **Get** /payouts/{payoutId} | Get payout by ID
*DefaultAPI* | [**GetPayouts**](docs/DefaultAPI.md#getpayouts) | **Get** /payouts | Get role payouts
*DefaultAPI* | [**GetTransactionsForContract**](docs/DefaultAPI.md#gettransactionsforcontract) | **Get** /contracts/{contractId}/transactions | Get transactions for contract
*DefaultAPI* | [**GetWithdrawalById**](docs/DefaultAPI.md#getwithdrawalbyid) | **Get** /withdrawals/{withdrawalId} | Get withdrawal by ID
*DefaultAPI* | [**GetWithdrawals**](docs/DefaultAPI.md#getwithdrawals) | **Get** /withdrawals | Get withdrawals
*DefaultAPI* | [**Healthcheck**](docs/DefaultAPI.md#healthcheck) | **Get** /healthcheck | Test server status
*DefaultAPI* | [**SubmitContract**](docs/DefaultAPI.md#submitcontract) | **Put** /contracts/{contractId} | Submit contract to chain
*DefaultAPI* | [**SubmitContractTransaction**](docs/DefaultAPI.md#submitcontracttransaction) | **Put** /contracts/{contractId}/transactions/{transactionId} | Submit contract input application
*DefaultAPI* | [**SubmitWithdrawal**](docs/DefaultAPI.md#submitwithdrawal) | **Put** /withdrawals/{withdrawalId} | Submit payout withdrawal
*DefaultAPI* | [**WithdrawPayouts**](docs/DefaultAPI.md#withdrawpayouts) | **Post** /withdrawals | Withdraw payouts


## Documentation For Models

 - [AccountTokenTupleInner](docs/AccountTokenTupleInner.md)
 - [Action](docs/Action.md)
 - [ActionObject](docs/ActionObject.md)
 - [Add](docs/Add.md)
 - [AddObject](docs/AddObject.md)
 - [AddressAndMetadata](docs/AddressAndMetadata.md)
 - [And](docs/And.md)
 - [AndObject](docs/AndObject.md)
 - [ApplicableInputs](docs/ApplicableInputs.md)
 - [ApplyInputsResponse](docs/ApplyInputsResponse.md)
 - [ApplyInputsResponseLinks](docs/ApplyInputsResponseLinks.md)
 - [ApplyInputsTxEnvelope](docs/ApplyInputsTxEnvelope.md)
 - [Assert](docs/Assert.md)
 - [AssertFail](docs/AssertFail.md)
 - [AssertObject](docs/AssertObject.md)
 - [AssetId](docs/AssetId.md)
 - [Assets](docs/Assets.md)
 - [BlockHeader](docs/BlockHeader.md)
 - [Bound](docs/Bound.md)
 - [CanChoose](docs/CanChoose.md)
 - [CanDeposit](docs/CanDeposit.md)
 - [CanNotify](docs/CanNotify.md)
 - [Case](docs/Case.md)
 - [CaseMerkleizedThen](docs/CaseMerkleizedThen.md)
 - [CaseMerkleizedThenObject](docs/CaseMerkleizedThenObject.md)
 - [CaseObject](docs/CaseObject.md)
 - [CaseThen](docs/CaseThen.md)
 - [CaseThenObject](docs/CaseThenObject.md)
 - [ChoiceAction](docs/ChoiceAction.md)
 - [ChoiceActionObject](docs/ChoiceActionObject.md)
 - [ChoiceContinuationInput](docs/ChoiceContinuationInput.md)
 - [ChoiceId](docs/ChoiceId.md)
 - [ChoiceIdObject](docs/ChoiceIdObject.md)
 - [ChoiceInput](docs/ChoiceInput.md)
 - [ChooseFor](docs/ChooseFor.md)
 - [ChooseForObject](docs/ChooseForObject.md)
 - [Close](docs/Close.md)
 - [CloseObject](docs/CloseObject.md)
 - [ContinuationInput](docs/ContinuationInput.md)
 - [Contract](docs/Contract.md)
 - [ContractHeader](docs/ContractHeader.md)
 - [ContractObject](docs/ContractObject.md)
 - [ContractSourceIds](docs/ContractSourceIds.md)
 - [ContractState](docs/ContractState.md)
 - [CreateContractResponse](docs/CreateContractResponse.md)
 - [CreateContractResponseLinks](docs/CreateContractResponseLinks.md)
 - [CreateTxBodyEnvelope](docs/CreateTxBodyEnvelope.md)
 - [CreateTxEnvelope](docs/CreateTxEnvelope.md)
 - [DepositAction](docs/DepositAction.md)
 - [DepositActionObject](docs/DepositActionObject.md)
 - [DepositContinuationInput](docs/DepositContinuationInput.md)
 - [DepositInput](docs/DepositInput.md)
 - [Divide](docs/Divide.md)
 - [DivideObject](docs/DivideObject.md)
 - [Equal](docs/Equal.md)
 - [EqualObject](docs/EqualObject.md)
 - [ExBudget](docs/ExBudget.md)
 - [GetContractResponse](docs/GetContractResponse.md)
 - [GetContractResponseLinks](docs/GetContractResponseLinks.md)
 - [GetContractsResponse](docs/GetContractsResponse.md)
 - [GetContractsResponseResultsInner](docs/GetContractsResponseResultsInner.md)
 - [GetContractsResponseResultsInnerLinks](docs/GetContractsResponseResultsInnerLinks.md)
 - [GetPayoutResponse](docs/GetPayoutResponse.md)
 - [GetPayoutResponseLinks](docs/GetPayoutResponseLinks.md)
 - [GetPayoutsResponse](docs/GetPayoutsResponse.md)
 - [GetPayoutsResponseResultsInner](docs/GetPayoutsResponseResultsInner.md)
 - [GetPayoutsResponseResultsInnerLinks](docs/GetPayoutsResponseResultsInnerLinks.md)
 - [GetTransactionResponse](docs/GetTransactionResponse.md)
 - [GetTransactionResponseLinks](docs/GetTransactionResponseLinks.md)
 - [GetTransactionsResponse](docs/GetTransactionsResponse.md)
 - [GetTransactionsResponseResultsInner](docs/GetTransactionsResponseResultsInner.md)
 - [GetWithdrawalsResponse](docs/GetWithdrawalsResponse.md)
 - [GetWithdrawalsResponseResultsInner](docs/GetWithdrawalsResponseResultsInner.md)
 - [GetWithdrawalsResponseResultsInnerLinks](docs/GetWithdrawalsResponseResultsInnerLinks.md)
 - [Greater](docs/Greater.md)
 - [GreaterObject](docs/GreaterObject.md)
 - [GreaterOrEqual](docs/GreaterOrEqual.md)
 - [GreaterOrEqualObject](docs/GreaterOrEqualObject.md)
 - [If](docs/If.md)
 - [IfObject](docs/IfObject.md)
 - [IfValue](docs/IfValue.md)
 - [IfValueObject](docs/IfValueObject.md)
 - [Input](docs/Input.md)
 - [IntervalError](docs/IntervalError.md)
 - [IntervalInPast](docs/IntervalInPast.md)
 - [IntervalInPastIntervalInPastError](docs/IntervalInPastIntervalInPastError.md)
 - [InvalidInterval](docs/InvalidInterval.md)
 - [InvalidIntervalInvalidInterval](docs/InvalidIntervalInvalidInterval.md)
 - [LabelRef](docs/LabelRef.md)
 - [LabelledObject](docs/LabelledObject.md)
 - [LabelledObjectValue](docs/LabelledObjectValue.md)
 - [Lesser](docs/Lesser.md)
 - [LesserObject](docs/LesserObject.md)
 - [LesserOrEqual](docs/LesserOrEqual.md)
 - [LesserOrEqualObject](docs/LesserOrEqualObject.md)
 - [Let](docs/Let.md)
 - [LetObject](docs/LetObject.md)
 - [MarloweState](docs/MarloweState.md)
 - [MarloweStateAccountsInnerInner](docs/MarloweStateAccountsInnerInner.md)
 - [MarloweStateBoundValuesInnerInner](docs/MarloweStateBoundValuesInnerInner.md)
 - [MarloweStateChoicesInnerInner](docs/MarloweStateChoicesInnerInner.md)
 - [MarloweVersion](docs/MarloweVersion.md)
 - [Metadata](docs/Metadata.md)
 - [MetadataAndRecipients](docs/MetadataAndRecipients.md)
 - [MetadataAndScript](docs/MetadataAndScript.md)
 - [Minus](docs/Minus.md)
 - [MinusObject](docs/MinusObject.md)
 - [Multiply](docs/Multiply.md)
 - [MultiplyObject](docs/MultiplyObject.md)
 - [Negate](docs/Negate.md)
 - [NegateObject](docs/NegateObject.md)
 - [Next](docs/Next.md)
 - [NonPositiveDeposit](docs/NonPositiveDeposit.md)
 - [NonPositivePayment](docs/NonPositivePayment.md)
 - [Not](docs/Not.md)
 - [NotObject](docs/NotObject.md)
 - [NotifyAction](docs/NotifyAction.md)
 - [NotifyActionObject](docs/NotifyActionObject.md)
 - [NotifyInput](docs/NotifyInput.md)
 - [Observation](docs/Observation.md)
 - [ObservationObject](docs/ObservationObject.md)
 - [Or](docs/Or.md)
 - [OrObject](docs/OrObject.md)
 - [PartialPayment](docs/PartialPayment.md)
 - [Party](docs/Party.md)
 - [PartyAddress](docs/PartyAddress.md)
 - [PartyObject](docs/PartyObject.md)
 - [PartyRoleName](docs/PartyRoleName.md)
 - [Pay](docs/Pay.md)
 - [PayObject](docs/PayObject.md)
 - [PayToAccount](docs/PayToAccount.md)
 - [PayToAccountObject](docs/PayToAccountObject.md)
 - [PayToParty](docs/PayToParty.md)
 - [PayToPartyObject](docs/PayToPartyObject.md)
 - [Payee](docs/Payee.md)
 - [PayeeObject](docs/PayeeObject.md)
 - [Payment](docs/Payment.md)
 - [Payout](docs/Payout.md)
 - [PayoutHeader](docs/PayoutHeader.md)
 - [PayoutState](docs/PayoutState.md)
 - [PayoutStatus](docs/PayoutStatus.md)
 - [PlutusAddress](docs/PlutusAddress.md)
 - [PlutusCredential](docs/PlutusCredential.md)
 - [PlutusStakingCredential](docs/PlutusStakingCredential.md)
 - [PostContractSourceResponse](docs/PostContractSourceResponse.md)
 - [PostContractsRequest](docs/PostContractsRequest.md)
 - [PostContractsRequestContract](docs/PostContractsRequestContract.md)
 - [PostTransactionsRequest](docs/PostTransactionsRequest.md)
 - [PostWithdrawalsRequest](docs/PostWithdrawalsRequest.md)
 - [PubKeyCredential](docs/PubKeyCredential.md)
 - [RoleTokenConfig](docs/RoleTokenConfig.md)
 - [RolesConfig](docs/RolesConfig.md)
 - [SafetyError](docs/SafetyError.md)
 - [ScriptCredential](docs/ScriptCredential.md)
 - [StakingHash](docs/StakingHash.md)
 - [StakingPointer](docs/StakingPointer.md)
 - [TextEnvelope](docs/TextEnvelope.md)
 - [TimeInterval](docs/TimeInterval.md)
 - [Token](docs/Token.md)
 - [TokenInAccount](docs/TokenInAccount.md)
 - [TokenInAccountObject](docs/TokenInAccountObject.md)
 - [TokenMetadata](docs/TokenMetadata.md)
 - [TokenMetadataFile](docs/TokenMetadataFile.md)
 - [TokenObject](docs/TokenObject.md)
 - [Transaction](docs/Transaction.md)
 - [TransactionError](docs/TransactionError.md)
 - [TransactionErrorOneOf](docs/TransactionErrorOneOf.md)
 - [TransactionInput](docs/TransactionInput.md)
 - [TransactionInputTxInterval](docs/TransactionInputTxInterval.md)
 - [TransactionOutput](docs/TransactionOutput.md)
 - [TransactionWarning](docs/TransactionWarning.md)
 - [Tx](docs/Tx.md)
 - [TxHeader](docs/TxHeader.md)
 - [TxOutputError](docs/TxOutputError.md)
 - [TxOutputSuccess](docs/TxOutputSuccess.md)
 - [TxStatus](docs/TxStatus.md)
 - [UseValue](docs/UseValue.md)
 - [Value](docs/Value.md)
 - [ValueObject](docs/ValueObject.md)
 - [ValueOfChoice](docs/ValueOfChoice.md)
 - [ValueOfChoiceObject](docs/ValueOfChoiceObject.md)
 - [VariableShadowing](docs/VariableShadowing.md)
 - [When](docs/When.md)
 - [WhenObject](docs/WhenObject.md)
 - [WithdrawPayoutsResponse](docs/WithdrawPayoutsResponse.md)
 - [WithdrawTxBodyEnvelope](docs/WithdrawTxBodyEnvelope.md)
 - [WithdrawTxEnvelope](docs/WithdrawTxEnvelope.md)
 - [Withdrawal](docs/Withdrawal.md)
 - [WithdrawalHeader](docs/WithdrawalHeader.md)


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



