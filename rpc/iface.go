package rpc

import (
	"encoding/json"
	"time"

	"github.com/go-resty/resty/v2"
)

// IFace is an interface mocking a GoTezos object.
type IFace interface {
	Block(blockID BlockID) (*resty.Response, *Block, error)
	EndorsingPower(input EndorsingPowerInput) (*resty.Response, int, error)
	Hash(blockID BlockID) (*resty.Response, string, error)
	Header(blockID BlockID) (*resty.Response, Header, error)
	HeaderRaw(blockID BlockID) (*resty.Response, string, error)
	HeaderShell(blockID BlockID) (*resty.Response, HeaderShell, error)
	HeaderProtocolData(blockID BlockID) (*resty.Response, ProtocolData, error)
	HeaderProtocolDataRaw(blockID BlockID) (*resty.Response, string, error)
	LiveBlocks(blockID BlockID) (*resty.Response, []string, error)
	Metadata(blockID BlockID) (*resty.Response, Metadata, error)
	MetadataHash(blockID BlockID) (*resty.Response, string, error)
	MinimalValidTime(input MinimalValidTimeInput) (*resty.Response, time.Time, error)
	OperationHashes(input OperationHashesInput) (*resty.Response, OperationHashes, error)
	OperationMetadataHashes(input OperationMetadataHashesInput) (*resty.Response, OperationMetadataHashes, error)
	Operations(input OperationsInput) (*resty.Response, FlattenedOperations, error)
	OperationsMetadataHash(blockID BlockID) (*resty.Response, string, error)
	Protocols(blockID BlockID) (*resty.Response, Protocols, error)
	RequiredEndorsements(input RequiredEndorsementsInput) (*resty.Response, int, error)
	BigMap(input BigMapInput) (*resty.Response, error)
	Constants(input ConstantsInput) (*resty.Response, Constants, error)
	Contracts(input ContractsInput) (*resty.Response, []string, error)
	Contract(input ContractInput) (*resty.Response, Contract, error)
	ContractBalance(input ContractBalanceInput) (*resty.Response, string, error)
	ContractCounter(input ContractCounterInput) (*resty.Response, int, error)
	ContractDelegate(input ContractDelegateInput) (*resty.Response, string, error)
	ContractEntrypoints(input ContractEntrypointsInput) (*resty.Response, map[string]*json.RawMessage, error)
	ContractEntrypoint(input ContractEntrypointInput) (*resty.Response, *json.RawMessage, error)
	ContractManagerKey(input ContractManagerKeyInput) (*resty.Response, string, error)
	ContractScript(input ContractScriptInput) (*resty.Response, error)
	ContractSaplingDiff(input ContractSaplingDiffInput) (*resty.Response, error)
	ContractStorage(input ContractStorageInput) (*resty.Response, error)
	Delegates(input DelegatesInput) (*resty.Response, []string, error)
	Delegate(input DelegateInput) (*resty.Response, Delegate, error)
	DelegateBalance(input DelegateBalanceInput) (*resty.Response, string, error)
	DelegateDeactivated(input DelegateDeactivatedInput) (*resty.Response, bool, error)
	DelegateDelegatedBalance(input DelegateDelegatedBalanceInput) (*resty.Response, string, error)
	DelegateDelegatedContracts(input DelegateDelegatedContractsInput) (*resty.Response, []string, error)
	DelegateFrozenBalance(input DelegateFrozenBalanceInput) (*resty.Response, string, error)
	DelegateFrozenBalanceByCycle(input DelegateFrozenBalanceByCycleInput) (*resty.Response, []FrozenBalanceByCycle, error)
	DelegateGracePeriod(input DelegateGracePeriodInput) (*resty.Response, int, error)
	DelegateStakingBalance(input DelegateStakingBalanceInput) (*resty.Response, string, error)
	DelegateVotingPower(input DelegateVotingPowerInput) (*resty.Response, int, error)
	Nonces(input NoncesInput) (*resty.Response, Nonces, error)
	RawBytes(input RawBytesInput) (*resty.Response, error)
	SaplingDiff(input SaplingDiffInput) (*resty.Response, error)
	Seed(input SeedInput) (*resty.Response, string, error)
	Cycle(cycle int) (*resty.Response, Cycle, error)
	BakingRights(input BakingRightsInput) (*resty.Response, []BakingRights, error)
	CompletePrefix(input CompletePrefixInput) (*resty.Response, []string, error)
	CurrentLevel(input CurrentLevelInput) (*resty.Response, CurrentLevel, error)
	EndorsingRights(input EndorsingRightsInput) (*resty.Response, []EndorsingRights, error)
	ForgeOperations(input ForgeOperationsInput) (*resty.Response, string, error)
	ForgeBlockHeader(input ForgeBlockHeaderInput) (*resty.Response, ForgeBlockHeader, error)
	LevelsInCurrentCycle(input LevelsInCurrentCycleInput) (*resty.Response, LevelsInCurrentCycle, error)
	ParseBlock(input ParseBlockInput) (*resty.Response, BlockHeaderSignedContents, error)
	ParseOperations(input ParseOperationsInput) (*resty.Response, []Operations, error)
	PreapplyBlock(input PreapplyBlockInput) (*resty.Response, PreappliedBlock, error)
	PreapplyOperations(input PreapplyOperationsInput) (*resty.Response, []Operations, error)
	Entrypoint(input EntrypointInput) (*resty.Response, Entrypoint, error)
	Entrypoints(input EntrypointsInput) (*resty.Response, Entrypoints, error)
	PackData(input PackDataInput) (*resty.Response, PackedData, error)
	RunCode(input RunCodeInput) (*resty.Response, RanCode, error)
	RunOperation(input RunOperationInput) (*resty.Response, Operations, error)
	TraceCode(input TraceCodeInput) (*resty.Response, TracedCode, error)
	TypecheckCode(input TypeCheckcodeInput) (*resty.Response, TypecheckedCode, error)
	TypecheckData(input TypecheckDataInput) (*resty.Response, TypecheckedData, error)
	BallotList(blockID BlockID) (*resty.Response, BallotList, error)
	Ballots(blockID BlockID) (*resty.Response, Ballots, error)
	CurrentPeriod(blockID BlockID) (*resty.Response, VotingPeriod, error)
	CurrentPeriodKind(blockID BlockID) (*resty.Response, string, error)
	CurrentProposal(blockID BlockID) (*resty.Response, string, error)
	CurrentQuorum(blockID BlockID) (*resty.Response, int, error)
	Listings(blockID BlockID) (*resty.Response, Listings, error)
	Proposals(blockID BlockID) (*resty.Response, Proposals, error)
	SuccessorPeriod(blockID BlockID) (*resty.Response, VotingPeriod, error)
	TotalVotingPower(blockID BlockID) (*resty.Response, int, error)
	GetFA12Balance(input GetFA12BalanceInput) (*resty.Response, string, error)
	GetFA12Supply(input GetFA12SupplyInput) (*resty.Response, string, error)
	GetFA12Allowance(input GetFA12AllowanceInput) (*resty.Response, string, error)
	InjectionOperation(input InjectionOperationInput) (*resty.Response, string, error)
	InjectionBlock(input InjectionBlockInput) (*resty.Response, error)
	Connections() (*resty.Response, Connections, error)
	ActiveChains() (*resty.Response, ActiveChains, error)
}