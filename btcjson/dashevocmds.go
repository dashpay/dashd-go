// Copyright (c) 2014-2017 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

// NOTE: This file is intended to house the RPC commands that are supported by
// a chain server.

package btcjson

import (
	"errors"
	"fmt"
)

func init() {
	// No special flags for commands in this file.
	flags := UsageFlag(0)

	MustRegisterCmd("quorum", (*QuorumCmd)(nil), flags)
	MustRegisterCmd("bls", (*BLSCmd)(nil), flags)
	MustRegisterCmd("protx", (*ProTxCmd)(nil), flags)
}

type BLSSubCmd string

const (
	BLSGenerate   BLSSubCmd = "generate"
	BLSFromSecret BLSSubCmd = "fromsecret"
)

type BLSCmd struct {
	SubCmd BLSSubCmd `jsonrpcusage:"\"generate|fromsecret\""`
	Secret *string   `json:",omitempty"`
}

type ProTxSubCmd string

const (
	ProTxRegister        ProTxSubCmd = "register"
	ProTxRegisterFund    ProTxSubCmd = "register_fund"
	ProTxRegisterPrepare ProTxSubCmd = "register_prepare"
	ProTxRegisterSubmit  ProTxSubCmd = "register_submit"
	ProTxList            ProTxSubCmd = "list"
	ProTxInfo            ProTxSubCmd = "info"
	ProTxUpdateService   ProTxSubCmd = "update_service"
	ProTxUpdateRegistrar ProTxSubCmd = "update_registrar"
	ProTxRevoke          ProTxSubCmd = "revoke"
	ProTxDiff            ProTxSubCmd = "diff"
)

type ProTxCmd struct {
	SubCmd ProTxSubCmd `jsonrpcusage:"\"register|register_fund|register_prepare|register_submit|list|info|update_service|update_registrar|revoke|diff\""`

	ProTxHash *string `json:",omitempty"`

	Type     *ProTxListType `json:",omitempty"`
	Detailed *bool          `json:",omitempty"`
	Height   *int           `json:",omitempty"`

	BaseBlock *int `json:",omitempty"`
	Block     *int `json:",omitempty"`

	CollateralHash        *string  `json:",omitempty"`
	CollateralIndex       *int     `json:",omitempty"`
	CollateralAddress     *string  `json:",omitempty"`
	IPAndPort             *string  `json:",omitempty"`
	OwnerAddress          *string  `json:",omitempty"`
	OperatorPubKey        *string  `json:",omitempty"`
	OperatorPrivateKey    *string  `json:",omitempty"`
	OperatorPayoutAddress *string  `json:",omitempty"`
	VotingAddress         *string  `json:",omitempty"`
	OperatorReward        *float64 `json:",omitempty"`
	PayoutAddress         *string  `json:",omitempty"`
	FundAddress           *string  `json:",omitempty"`
	Reason                *int     `json:",omitempty"`
	FeeSourceAddress      *string  `json:",omitempty"`
	Submit                *bool    `json:",omitempty"`

	Tx  *string `json:",omitempty"`
	Sig *string `json:",omitempty"`
}

type ProTxListType string

const (
	ProTxListTypeRegistered ProTxListType = "registered"
	ProTxListTypeValid      ProTxListType = "valid"
	ProTxListTypeWallet     ProTxListType = "wallet"
)

// QuorumCmdSubCmd defines the sub command used in the quorum JSON-RPC command.
type QuorumCmdSubCmd string

// Quorum commands https://dashcore.readme.io/docs/core-api-ref-remote-procedure-calls-evo#quorum
const (
	QuorumSign         QuorumCmdSubCmd = "sign"
	QuorumSignPlatform QuorumCmdSubCmd = "platformsign"
	QuorumVerify       QuorumCmdSubCmd = "verify"
	QuorumInfo         QuorumCmdSubCmd = "info"

	// QuorumList lists all quorums
	QuorumList QuorumCmdSubCmd = "list"

	QuorumSelectQuorum  QuorumCmdSubCmd = "selectquorum"
	QuorumDKGStatus     QuorumCmdSubCmd = "dkgstatus"
	QuorumMemberOf      QuorumCmdSubCmd = "memberof"
	QuorumGetRecSig     QuorumCmdSubCmd = "getrecsig"
	QuorumHasRecSig     QuorumCmdSubCmd = "hasrecsig"
	QuorumIsConflicting QuorumCmdSubCmd = "isconflicting"
)

// DetailLevel is the level of detail used in dkgstatus
type DetailLevel int

// Detail Levels for dkgstatsu
const (
	DetailLevelCounts             DetailLevel = 0
	DetailLevelIndexes            DetailLevel = 1
	DetailLevelMembersProTxHashes DetailLevel = 2
)

// LLMQType is the type of quorum
type LLMQType int

// Enum of LLMQTypes
// See https://github.com/dashpay/dips/blob/master/dip-0006.md#current-llmq-types and
// https://github.com/dashpay/dash/blob/master/src/llmq/params.h
const (
	LLMQType_50_60            LLMQType = 1   // 50 members, 30 (60%) threshold, one per hour
	LLMQType_400_60           LLMQType = 2   // 400 members, 240 (60%) threshold, one every 12 hours
	LLMQType_400_85           LLMQType = 3   // 400 members, 340 (85%) threshold, one every 24 hours
	LLMQType_100_67           LLMQType = 4   // 100 members, 67 (67%) threshold, one per hour
	LLMQType_60_75            LLMQType = 5   // 60 members, 45 (75%) threshold, one every 12 hours
	LLMQType_25_67            LLMQType = 6   // 25 members, 17 (67%) threshold, one per hour
	LLMQType_TEST             LLMQType = 100 // 3 members, 2 (66%) threshold, one per hour
	LLMQType_DEVNET           LLMQType = 101 // 12 members, 6 (50%) threshold, one per hour
	LLMQType_TEST_V17         LLMQType = 102 // 3 members, 2 (66%) threshold, one per hour
	LLMQType_TEST_DIP0024     LLMQType = 103 // 4 members, 2 (66%) threshold, one per hour
	LLMQType_TEST_INSTANTSEND LLMQType = 104 // 3 members, 2 (66%) threshold, one per hour
	LLMQType_DEVNET_DIP0024   LLMQType = 105 // 8 members, 4 (50%) threshold, one per hour
	LLMQType_TEST_PLATFORM    LLMQType = 106 // 3 members, 2 (66%) threshold, one per hour
	LLMQType_DEVNET_PLATFORM  LLMQType = 107 // 12 members, 8 (67%) threshold, one per hour
	LLMQType_SINGLE_NODE      LLMQType = 111 // 1 memeber, 1 threshold, one per hour.

	// LLMQType_5_60 is replaced with LLMQType_TEST to adhere to DIP-0006 naming
	LLMQType_5_60 LLMQType = LLMQType_TEST
)

var (
	errWrongSizeOfArgs           = errors.New("wrong size of arguments")
	errQuorumUnmarshalerNotFound = errors.New("quorum unmarshaler not found")

	llmqTypes = map[string]LLMQType{
		"llmq_50_60":            LLMQType_50_60,
		"llmq_400_60":           LLMQType_400_60,
		"llmq_400_85":           LLMQType_400_85,
		"llmq_100_67":           LLMQType_100_67,
		"llmq_60_75":            LLMQType_60_75,
		"llmq_25_67":            LLMQType_25_67,
		"llmq_test":             LLMQType_TEST,
		"llmq_devnet":           LLMQType_DEVNET,
		"llmq_test_v17":         LLMQType_TEST_V17,
		"llmq_test_dip0024":     LLMQType_TEST_DIP0024,
		"llmq_test_instantsend": LLMQType_TEST_INSTANTSEND,
		"llmq_devnet_dip0024":   LLMQType_DEVNET_DIP0024,
		"llmq_test_platform":    LLMQType_TEST_PLATFORM,
		"llmq_devnet_platform":  LLMQType_DEVNET_PLATFORM,
		"llmq_single_node":      LLMQType_SINGLE_NODE,
	}
)

// GetLLMQType returns LLMQ type for the given name.
// Returns 0 when the name is not supported.
func GetLLMQType(name string) LLMQType {
	return llmqTypes[name]
}

// Name returns name of the LLMQType.
// Returns empty string when the type is invalid.
// See https://github.com/dashpay/dash/blob/master/src/llmq/params.h
func (t LLMQType) Name() string {
	for name, item := range llmqTypes {
		if t == item {
			return name
		}
	}
	return ""
}

// Validate checks if provided LLMQ type is valid, eg. if it's one of LLMQ types
// defined in accordance with DIP-0006.
// See https://github.com/dashpay/dips/blob/master/dip-0006/llmq-types.md
func (t LLMQType) Validate() error {
	if (t >= LLMQType_50_60 && t <= LLMQType_25_67) || (t >= LLMQType_TEST && t <= LLMQType_DEVNET_PLATFORM) || t == LLMQType_SINGLE_NODE {
		return nil
	}

	return fmt.Errorf("unsupported quorum type %d", t)
}

// QuorumCmd defines the quorum JSON-RPC command.
type QuorumCmd struct {
	SubCmd QuorumCmdSubCmd `jsonrpcusage:"\"list|info|dkgstatus|sign|platformsign|getrecsig|hasrecsig|isconflicting|memberof|selectquorum\""`

	LLMQType    *LLMQType `json:",omitempty"`
	RequestID   *string   `json:",omitempty"`
	MessageHash *string   `json:",omitempty"`
	Signature   *string   `json:",omitempty"`
	QuorumHash  *string   `json:",omitempty"`

	Submit               *bool        `json:",omitempty"`
	IncludeSkShare       *bool        `json:",omitempty"`
	DKGStatusDetailLevel *DetailLevel `json:",omitempty"`
	ProTxHash            *string      `json:",omitempty"`
	ScanQuorumsCount     *int         `json:",omitempty"`
}

// NewQuorumSignCmd returns a new instance which can be used to issue a quorum
// JSON-RPC command.
func NewQuorumSignCmd(quorumType LLMQType, requestID, messageHash, quorumHash string, submit bool) *QuorumCmd {
	cmd := &QuorumCmd{
		SubCmd:      QuorumSign,
		LLMQType:    &quorumType,
		RequestID:   &requestID,
		MessageHash: &messageHash,
	}
	if quorumHash == "" {
		return cmd
	}
	cmd.QuorumHash = &quorumHash
	cmd.Submit = &submit
	return cmd
}

// NewQuorumPlatformSignCmd returns a new instance which can be used to issue a quorum
// JSON-RPC command.
func NewQuorumPlatformSignCmd(requestID, messageHash, quorumHash string, submit bool) *QuorumCmd {
	cmd := &QuorumCmd{
		SubCmd:      QuorumSignPlatform,
		RequestID:   &requestID,
		MessageHash: &messageHash,
	}
	if quorumHash == "" {
		return cmd
	}
	cmd.QuorumHash = &quorumHash
	cmd.Submit = &submit
	return cmd
}

// NewQuorumVerifyCmd returns a new instance which can be used to issue a quorum
// JSON-RPC command.
func NewQuorumVerifyCmd(quorumType LLMQType, requestID string, messageHash string, signature string, quorumHash string) *QuorumCmd {
	cmd := &QuorumCmd{
		SubCmd:      QuorumVerify,
		LLMQType:    &quorumType,
		RequestID:   &requestID,
		MessageHash: &messageHash,
		Signature:   &signature,
		QuorumHash:  &quorumHash,
	}
	return cmd
}

// NewQuorumInfoCmd returns a new instance which can be used to issue a quorum
// JSON-RPC command.
func NewQuorumInfoCmd(quorumType LLMQType, quorumHash string, includeSkShare bool) *QuorumCmd {
	return &QuorumCmd{
		SubCmd:         QuorumInfo,
		LLMQType:       &quorumType,
		QuorumHash:     &quorumHash,
		IncludeSkShare: &includeSkShare,
	}
}

// NewQuorumListCmd returns a list of quorums
// JSON-RPC command.
func NewQuorumListCmd() *QuorumCmd {
	return &QuorumCmd{
		SubCmd: QuorumList,
	}
}

// NewQuorumSelectQuorumCmd returns the selected quorum
func NewQuorumSelectQuorumCmd(quorumType LLMQType, requestID string) *QuorumCmd {
	return &QuorumCmd{
		SubCmd:    QuorumSelectQuorum,
		LLMQType:  &quorumType,
		RequestID: &requestID,
	}
}

// NewQuorumDKGStatusCmd returns the result from quorum dkgstatus
func NewQuorumDKGStatusCmd(detailLevel DetailLevel) *QuorumCmd {
	return &QuorumCmd{
		SubCmd:               QuorumDKGStatus,
		DKGStatusDetailLevel: &detailLevel,
	}
}

// NewQuorumMemberOfCmd returns the result from quorum memberof
func NewQuorumMemberOfCmd(proTxHash string, scanQuorumsCount int) *QuorumCmd {
	cmd := &QuorumCmd{
		SubCmd:    QuorumMemberOf,
		ProTxHash: &proTxHash,
	}
	if scanQuorumsCount != 0 {
		cmd.ScanQuorumsCount = &scanQuorumsCount
	}
	return cmd
}

// NewQuorumGetRecSig returns the result from quorum getrecsig
func NewQuorumGetRecSig(quorumType LLMQType, requestID, messageHash string) *QuorumCmd {
	return &QuorumCmd{
		SubCmd:      QuorumGetRecSig,
		LLMQType:    &quorumType,
		RequestID:   &requestID,
		MessageHash: &messageHash,
	}
}

// NewQuorumHasRecSig returns the result from quorum hasrecsig
func NewQuorumHasRecSig(quorumType LLMQType, requestID, messageHash string) *QuorumCmd {
	return &QuorumCmd{
		SubCmd:      QuorumHasRecSig,
		LLMQType:    &quorumType,
		RequestID:   &requestID,
		MessageHash: &messageHash,
	}
}

// NewQuorumIsConflicting returns the result from quorum isconflicting
func NewQuorumIsConflicting(quorumType LLMQType, requestID, messageHash string) *QuorumCmd {
	return &QuorumCmd{
		SubCmd:      QuorumIsConflicting,
		LLMQType:    &quorumType,
		RequestID:   &requestID,
		MessageHash: &messageHash,
	}
}

func NewBLSGenerate() *BLSCmd {
	return &BLSCmd{SubCmd: BLSGenerate}
}

func NewBLSFromSecret(secret string) *BLSCmd {
	return &BLSCmd{
		SubCmd: BLSFromSecret,
		Secret: &secret,
	}
}

// NewProTxRegisterCmd returns a new instance which can be used to issue a protx register
// JSON-RPC command.
func NewProTxRegisterCmd(collateralHash string, collateralIndex int, ipAndPort, ownerAddress, operatorPubKey, votingAddress string, operatorReward float64, payoutAddress, feeSourceAddress string, submit bool) *ProTxCmd {
	r := &ProTxCmd{
		SubCmd:          ProTxRegister,
		CollateralHash:  &collateralHash,
		CollateralIndex: &collateralIndex,
		IPAndPort:       &ipAndPort,
		OwnerAddress:    &ownerAddress,
		OperatorPubKey:  &operatorPubKey,
		VotingAddress:   &votingAddress,
		OperatorReward:  &operatorReward,
		PayoutAddress:   &payoutAddress,
	}
	if feeSourceAddress == "" {
		return r
	}
	r.FeeSourceAddress = &feeSourceAddress
	r.Submit = &submit
	return r
}

// NewProTxRegisterFundCmd returns a new instance which can be used to issue a protx register_fund
// JSON-RPC command.
func NewProTxRegisterFundCmd(collateralAddress, ipAndPort, ownerAddress, operatorPubKey, votingAddress string, operatorReward float64, payoutAddress, fundAddress string, submit bool) *ProTxCmd {
	r := &ProTxCmd{
		SubCmd:            ProTxRegisterFund,
		CollateralAddress: &collateralAddress,
		IPAndPort:         &ipAndPort,
		OwnerAddress:      &ownerAddress,
		OperatorPubKey:    &operatorPubKey,
		VotingAddress:     &votingAddress,
		OperatorReward:    &operatorReward,
		PayoutAddress:     &payoutAddress,
	}
	if fundAddress == "" {
		return r
	}
	r.FundAddress = &fundAddress
	r.Submit = &submit
	return r
}

// NewProTxRegisterPrepareCmd returns a new instance which can be used to issue a protx register_prepare
// JSON-RPC command.
func NewProTxRegisterPrepareCmd(collateralHash string, collateralIndex int, ipAndPort, ownerAddress, operatorPubKey, votingAddress string, operatorReward float64, payoutAddress, feeSourceAddress string) *ProTxCmd {
	r := &ProTxCmd{
		SubCmd:          ProTxRegisterPrepare,
		CollateralHash:  &collateralHash,
		CollateralIndex: &collateralIndex,
		IPAndPort:       &ipAndPort,
		OwnerAddress:    &ownerAddress,
		OperatorPubKey:  &operatorPubKey,
		VotingAddress:   &votingAddress,
		OperatorReward:  &operatorReward,
		PayoutAddress:   &payoutAddress,
	}
	if feeSourceAddress == "" {
		return r
	}
	r.FeeSourceAddress = &feeSourceAddress
	return r
}

// NewProTxInfoCmd returns a new instance which can be used to issue a protx info
// JSON-RPC command.
func NewProTxInfoCmd(proTxHash string) *ProTxCmd {
	return &ProTxCmd{
		SubCmd:    ProTxInfo,
		ProTxHash: &proTxHash,
	}
}

// NewProTxListCmd returns a new instance which can be used to issue a protx list
// JSON-RPC command.
func NewProTxListCmd(cmdType ProTxListType, detailed bool, height int) *ProTxCmd {
	r := &ProTxCmd{
		SubCmd: ProTxList,
	}
	if cmdType == "" {
		return r
	}
	r.Type = &cmdType
	r.Detailed = &detailed
	if height == 0 {
		return r
	}
	r.Height = &height
	return r
}

// NewProTxRegisterSubmitCmd returns a new instance which can be used to issue a protx register_submit
// JSON-RPC command.
func NewProTxRegisterSubmitCmd(tx, sig string) *ProTxCmd {
	return &ProTxCmd{
		SubCmd: ProTxRegisterSubmit,
		Tx:     &tx,
		Sig:    &sig,
	}
}

// NewProTxDiffCmd returns a new instance which can be used to issue a protx diff
// JSON-RPC command.
func NewProTxDiffCmd(baseBlock, block int) *ProTxCmd {
	return &ProTxCmd{
		SubCmd:    ProTxDiff,
		BaseBlock: &baseBlock,
		Block:     &block,
	}
}

// NewProTxUpdateServiceCmd returns a new instance which can be used to issue a protx update_service
// JSON-RPC command.
func NewProTxUpdateServiceCmd(proTxHash, ipAndPort, operatorPubKey, operatorPayoutAddress, feeSourceAddress string) *ProTxCmd {
	r := &ProTxCmd{
		SubCmd:         ProTxUpdateService,
		ProTxHash:      &proTxHash,
		IPAndPort:      &ipAndPort,
		OperatorPubKey: &operatorPubKey,
	}
	if operatorPayoutAddress == "" {
		return r
	}
	r.OperatorPayoutAddress = &operatorPayoutAddress
	if feeSourceAddress == "" {
		return r
	}
	r.FeeSourceAddress = &feeSourceAddress
	return r
}

// NewProTxUpdateRegistrarCmd returns a new instance which can be used to issue a protx update_registrar
// JSON-RPC command.
func NewProTxUpdateRegistrarCmd(proTxHash, operatorPubKey, votingAddress, payoutAddress, feeSourceAddress string) *ProTxCmd {
	r := &ProTxCmd{
		SubCmd:         ProTxUpdateRegistrar,
		ProTxHash:      &proTxHash,
		OperatorPubKey: &operatorPubKey,
		VotingAddress:  &votingAddress,
		PayoutAddress:  &payoutAddress,
	}
	if feeSourceAddress == "" {
		return r
	}
	r.FeeSourceAddress = &feeSourceAddress
	return r
}

// NewProTxRevokeCmd returns a new instance which can be used to issue a protx revoke
// JSON-RPC command.
func NewProTxRevokeCmd(proTxHash, operatorPrivateKey string, reason int, feeSourceAddress string) *ProTxCmd {
	r := &ProTxCmd{
		SubCmd:             ProTxRevoke,
		ProTxHash:          &proTxHash,
		OperatorPrivateKey: &operatorPrivateKey,
	}
	if reason == 0 {
		return r
	}
	r.Reason = &reason

	if feeSourceAddress == "" {
		return r
	}
	r.FeeSourceAddress = &feeSourceAddress
	return r
}

// UnmarshalArgs maps a list of arguments to quorum struct
func (q *QuorumCmd) UnmarshalArgs(args []interface{}) error {
	if len(args) == 0 {
		return errWrongSizeOfArgs
	}
	subCmd := args[0].(string)
	q.SubCmd = QuorumCmdSubCmd(subCmd)
	unmarshaler, ok := quorumCmdUnmarshalers[string(q.SubCmd)]
	if !ok {
		return errQuorumUnmarshalerNotFound
	}
	return unmarshaler(q, args[1:])
}

type unmarshalQuorumCmdFunc func(*QuorumCmd, []interface{}) error

var quorumCmdUnmarshalers = map[string]unmarshalQuorumCmdFunc{
	string(QuorumInfo):         withQuorumUnmarshaler(quorumInfoUnmarshaler, validateQuorumArgs(3), unmarshalQuorumLLMQType),
	string(QuorumSign):         withQuorumUnmarshaler(quorumSignUnmarshaler, validateQuorumArgs(5), unmarshalQuorumLLMQType),
	string(QuorumSignPlatform): withQuorumUnmarshaler(quorumPlatformSignUnmarshaler, validateQuorumArgs(4)),
	string(QuorumVerify):       withQuorumUnmarshaler(quorumVerifyUnmarshaler, validateQuorumArgs(5), unmarshalQuorumLLMQType),
}

func unmarshalLLMQType(val interface{}) (LLMQType, error) {
	var vInt int
	switch tv := val.(type) {
	case float64:
		vInt = int(tv)
	case float32:
		vInt = int(tv)
	case int:
		vInt = tv
	case LLMQType:
		return tv, nil
	}
	return LLMQType(vInt), nil
}

func quorumInfoUnmarshaler(q *QuorumCmd, args []interface{}) error {
	q.QuorumHash = strPtr(args[1].(string))
	q.IncludeSkShare = boolPtr(args[2].(bool))
	return nil
}

func quorumSignUnmarshaler(q *QuorumCmd, args []interface{}) error {
	q.RequestID = strPtr(args[1].(string))
	q.MessageHash = strPtr(args[2].(string))
	q.QuorumHash = strPtr(args[3].(string))
	q.Submit = boolPtr(args[4].(bool))
	return nil
}

func quorumPlatformSignUnmarshaler(q *QuorumCmd, args []interface{}) error {
	q.RequestID = strPtr(args[0].(string))
	q.MessageHash = strPtr(args[1].(string))
	q.QuorumHash = strPtr(args[2].(string))
	q.Submit = boolPtr(args[3].(bool))
	return nil
}

func unmarshalQuorumLLMQType(next unmarshalQuorumCmdFunc) unmarshalQuorumCmdFunc {
	return func(q *QuorumCmd, args []interface{}) error {
		val, err := unmarshalLLMQType(args[0])
		if err != nil {
			return err
		}
		q.LLMQType = llmqTypePtr(val)
		return next(q, args)
	}
}

func quorumVerifyUnmarshaler(q *QuorumCmd, args []interface{}) error {
	q.RequestID = strPtr(args[1].(string))
	q.MessageHash = strPtr(args[2].(string))
	q.Signature = strPtr(args[3].(string))
	q.QuorumHash = strPtr(args[4].(string))
	return nil
}

func validateQuorumArgs(n int) func(unmarshalQuorumCmdFunc) unmarshalQuorumCmdFunc {
	return func(next unmarshalQuorumCmdFunc) unmarshalQuorumCmdFunc {
		return func(q *QuorumCmd, args []interface{}) error {
			if n > len(args) {
				return errWrongSizeOfArgs
			}
			return next(q, args)
		}
	}
}

func withQuorumUnmarshaler(
	unmarshaler unmarshalQuorumCmdFunc,
	fns ...func(unmarshalQuorumCmdFunc) unmarshalQuorumCmdFunc,
) unmarshalQuorumCmdFunc {
	for _, fn := range fns {
		unmarshaler = fn(unmarshaler)
	}
	return unmarshaler
}
