// Auto-generated by avdl-compiler v1.3.13 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/keybase1/common.avdl

package keybase1

import (
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
)

type Time int64
type DurationSec float64
type StringKVPair struct {
	Key   string `codec:"key" json:"key"`
	Value string `codec:"value" json:"value"`
}

type Status struct {
	Code   int            `codec:"code" json:"code"`
	Name   string         `codec:"name" json:"name"`
	Desc   string         `codec:"desc" json:"desc"`
	Fields []StringKVPair `codec:"fields" json:"fields"`
}

type UID string
type DeviceID string
type SigID string
type KID string
type LinkID string
type BinaryKID []byte
type TLFID string
type TeamID string
type UserOrTeamID string
type Seqno int64
type Bytes32 [32]byte
type Text struct {
	Data   string `codec:"data" json:"data"`
	Markup bool   `codec:"markup" json:"markup"`
}

type PGPIdentity struct {
	Username string `codec:"username" json:"username"`
	Comment  string `codec:"comment" json:"comment"`
	Email    string `codec:"email" json:"email"`
}

type PublicKey struct {
	KID               KID           `codec:"KID" json:"KID"`
	PGPFingerprint    string        `codec:"PGPFingerprint" json:"PGPFingerprint"`
	PGPIdentities     []PGPIdentity `codec:"PGPIdentities" json:"PGPIdentities"`
	IsSibkey          bool          `codec:"isSibkey" json:"isSibkey"`
	IsEldest          bool          `codec:"isEldest" json:"isEldest"`
	ParentID          string        `codec:"parentID" json:"parentID"`
	DeviceID          DeviceID      `codec:"deviceID" json:"deviceID"`
	DeviceDescription string        `codec:"deviceDescription" json:"deviceDescription"`
	DeviceType        string        `codec:"deviceType" json:"deviceType"`
	CTime             Time          `codec:"cTime" json:"cTime"`
	ETime             Time          `codec:"eTime" json:"eTime"`
	IsRevoked         bool          `codec:"isRevoked" json:"isRevoked"`
}

type KeybaseTime struct {
	Unix  Time `codec:"unix" json:"unix"`
	Chain int  `codec:"chain" json:"chain"`
}

type RevokedKey struct {
	Key  PublicKey   `codec:"key" json:"key"`
	Time KeybaseTime `codec:"time" json:"time"`
	By   KID         `codec:"by" json:"by"`
}

type User struct {
	Uid      UID    `codec:"uid" json:"uid"`
	Username string `codec:"username" json:"username"`
}

type Device struct {
	Type         string   `codec:"type" json:"type"`
	Name         string   `codec:"name" json:"name"`
	DeviceID     DeviceID `codec:"deviceID" json:"deviceID"`
	CTime        Time     `codec:"cTime" json:"cTime"`
	MTime        Time     `codec:"mTime" json:"mTime"`
	LastUsedTime Time     `codec:"lastUsedTime" json:"lastUsedTime"`
	EncryptKey   KID      `codec:"encryptKey" json:"encryptKey"`
	VerifyKey    KID      `codec:"verifyKey" json:"verifyKey"`
	Status       int      `codec:"status" json:"status"`
}

type DeviceType int

const (
	DeviceType_DESKTOP DeviceType = 0
	DeviceType_MOBILE  DeviceType = 1
)

var DeviceTypeMap = map[string]DeviceType{
	"DESKTOP": 0,
	"MOBILE":  1,
}

var DeviceTypeRevMap = map[DeviceType]string{
	0: "DESKTOP",
	1: "MOBILE",
}

func (e DeviceType) String() string {
	if v, ok := DeviceTypeRevMap[e]; ok {
		return v
	}
	return ""
}

type Stream struct {
	Fd int `codec:"fd" json:"fd"`
}

type LogLevel int

const (
	LogLevel_NONE     LogLevel = 0
	LogLevel_DEBUG    LogLevel = 1
	LogLevel_INFO     LogLevel = 2
	LogLevel_NOTICE   LogLevel = 3
	LogLevel_WARN     LogLevel = 4
	LogLevel_ERROR    LogLevel = 5
	LogLevel_CRITICAL LogLevel = 6
	LogLevel_FATAL    LogLevel = 7
)

var LogLevelMap = map[string]LogLevel{
	"NONE":     0,
	"DEBUG":    1,
	"INFO":     2,
	"NOTICE":   3,
	"WARN":     4,
	"ERROR":    5,
	"CRITICAL": 6,
	"FATAL":    7,
}

var LogLevelRevMap = map[LogLevel]string{
	0: "NONE",
	1: "DEBUG",
	2: "INFO",
	3: "NOTICE",
	4: "WARN",
	5: "ERROR",
	6: "CRITICAL",
	7: "FATAL",
}

func (e LogLevel) String() string {
	if v, ok := LogLevelRevMap[e]; ok {
		return v
	}
	return ""
}

type ClientType int

const (
	ClientType_NONE       ClientType = 0
	ClientType_CLI        ClientType = 1
	ClientType_GUI_MAIN   ClientType = 2
	ClientType_KBFS       ClientType = 3
	ClientType_GUI_HELPER ClientType = 4
)

var ClientTypeMap = map[string]ClientType{
	"NONE":       0,
	"CLI":        1,
	"GUI_MAIN":   2,
	"KBFS":       3,
	"GUI_HELPER": 4,
}

var ClientTypeRevMap = map[ClientType]string{
	0: "NONE",
	1: "CLI",
	2: "GUI_MAIN",
	3: "KBFS",
	4: "GUI_HELPER",
}

type UserVersionVector struct {
	Id               int64 `codec:"id" json:"id"`
	SigHints         int   `codec:"sigHints" json:"sigHints"`
	SigChain         int64 `codec:"sigChain" json:"sigChain"`
	CachedAt         Time  `codec:"cachedAt" json:"cachedAt"`
	LastIdentifiedAt Time  `codec:"lastIdentifiedAt" json:"lastIdentifiedAt"`
}

type PerUserKey struct {
	Gen    int   `codec:"gen" json:"gen"`
	Seqno  Seqno `codec:"seqno" json:"seqno"`
	SigKID KID   `codec:"sigKID" json:"sigKID"`
	EncKID KID   `codec:"encKID" json:"encKID"`
}

type PerTeamKey struct {
	Gen    int   `codec:"gen" json:"gen"`
	Seqno  Seqno `codec:"seqno" json:"seqno"`
	SigKID KID   `codec:"sigKID" json:"sigKID"`
	EncKID KID   `codec:"encKID" json:"encKID"`
}

type UserPlusKeys struct {
	Uid               UID               `codec:"uid" json:"uid"`
	Username          string            `codec:"username" json:"username"`
	EldestSeqno       Seqno             `codec:"eldestSeqno" json:"eldestSeqno"`
	DeviceKeys        []PublicKey       `codec:"deviceKeys" json:"deviceKeys"`
	RevokedDeviceKeys []RevokedKey      `codec:"revokedDeviceKeys" json:"revokedDeviceKeys"`
	PGPKeyCount       int               `codec:"pgpKeyCount" json:"pgpKeyCount"`
	Uvv               UserVersionVector `codec:"uvv" json:"uvv"`
	DeletedDeviceKeys []PublicKey       `codec:"deletedDeviceKeys" json:"deletedDeviceKeys"`
	PerUserKeys       []PerUserKey      `codec:"perUserKeys" json:"perUserKeys"`
}

type UserOrTeamLite struct {
	Id   UserOrTeamID `codec:"id" json:"id"`
	Name string       `codec:"name" json:"name"`
}

type RemoteTrack struct {
	Username string `codec:"username" json:"username"`
	Uid      UID    `codec:"uid" json:"uid"`
	LinkID   LinkID `codec:"linkID" json:"linkID"`
}

type UserPlusAllKeys struct {
	Base         UserPlusKeys  `codec:"base" json:"base"`
	PGPKeys      []PublicKey   `codec:"pgpKeys" json:"pgpKeys"`
	RemoteTracks []RemoteTrack `codec:"remoteTracks" json:"remoteTracks"`
}

type MerkleTreeID int

const (
	MerkleTreeID_MASTER       MerkleTreeID = 0
	MerkleTreeID_KBFS_PUBLIC  MerkleTreeID = 1
	MerkleTreeID_KBFS_PRIVATE MerkleTreeID = 2
)

var MerkleTreeIDMap = map[string]MerkleTreeID{
	"MASTER":       0,
	"KBFS_PUBLIC":  1,
	"KBFS_PRIVATE": 2,
}

var MerkleTreeIDRevMap = map[MerkleTreeID]string{
	0: "MASTER",
	1: "KBFS_PUBLIC",
	2: "KBFS_PRIVATE",
}

// SocialAssertionService is a service that can be used to assert proofs for a
// user.
type SocialAssertionService string

// SocialAssertion contains a service and username for that service, that
// together form an assertion about a user. Resolving an assertion requires
// that the user posts a Keybase proof on the asserted service as the asserted
// user.
type SocialAssertion struct {
	User    string                 `codec:"user" json:"user"`
	Service SocialAssertionService `codec:"service" json:"service"`
}

// UserResolution maps how an unresolved user assertion has been resolved.
type UserResolution struct {
	Assertion SocialAssertion `codec:"assertion" json:"assertion"`
	UserID    UID             `codec:"userID" json:"userID"`
}

type TeamRole int

const (
	TeamRole_NONE   TeamRole = 0
	TeamRole_OWNER  TeamRole = 1
	TeamRole_ADMIN  TeamRole = 2
	TeamRole_WRITER TeamRole = 3
	TeamRole_READER TeamRole = 4
)

var TeamRoleMap = map[string]TeamRole{
	"NONE":   0,
	"OWNER":  1,
	"ADMIN":  2,
	"WRITER": 3,
	"READER": 4,
}

var TeamRoleRevMap = map[TeamRole]string{
	0: "NONE",
	1: "OWNER",
	2: "ADMIN",
	3: "WRITER",
	4: "READER",
}

func (e TeamRole) String() string {
	if v, ok := TeamRoleRevMap[e]; ok {
		return v
	}
	return ""
}

type CommonInterface interface {
}

func CommonProtocol(i CommonInterface) rpc.Protocol {
	return rpc.Protocol{
		Name:    "keybase.1.Common",
		Methods: map[string]rpc.ServeHandlerDescription{},
	}
}

type CommonClient struct {
	Cli rpc.GenericClient
}
