package mtools

import (
	"errors"
	"fmt"
	"github.com/namejlt/gozen"
	"github.com/satori/go.uuid"
	"github.com/teris-io/shortid"
	"math/rand"
	"net"
	"strconv"
	"sync"
	"time"
)

// ========================================== func
// ========================================== func
// ========================================== func

func GetRandomUUId(ver uint8) (ret string, err error) {
	var (
		id uuid.UUID
	)
	switch ver {
	case uuid.V1:
		id = uuid.NewV1() //基于时间的UUID 适合应用于分布式计算环境下，具有高度的唯一性
	case uuid.V2:
		id = uuid.NewV2(uuid.DomainGroup) //DCE安全的UUID 适合应用于分布式计算环境下，具有高度的唯一性
	case uuid.V3:
		id = uuid.NewV3(uuid.NamespaceURL, "saas-module") //基于名字的UUID（MD5）适合于一定范围内名字唯一
	case uuid.V4:
		id = uuid.NewV4() //随机UUID
	case uuid.V5:
		id = uuid.NewV5(uuid.NamespaceURL, "saas-module") //基于名字的UUID（SHA1）适合于一定范围内名字唯一
	default:
		id = uuid.NewV1()
	}
	ret = id.String()
	return
}

func GetRandomShortUUId() (ret string, err error) {
	ret, err = shortid.Generate()
	return
}

func GetRandomId() (num uint64) {
	var err error
	num, err = MakeSnowflakeId()
	if err != nil {
		num = uint64(time.Now().UnixNano())
	}
	return
}

func MakeSnowflakeId() (num uint64, err error) {
	num, err = snowflake.NextID()
	return
}

// ========================================= init
// ========================================= init
// ========================================= init

var (
	snowflake      *Sonyflake
	shortSnowflake *ShortSonyflake
)

func init() {
	var (
		st  Settings
		sst ShortSettings
		now time.Time
	)
	now = time.Now()

	st.StartTime = now
	midStr := gozen.ConfigAppGetString("SonyflakeMachineID", "1")
	mid, err := strconv.Atoi(midStr)
	if err != nil {
		panic("sonyflake MachineID is error")
	}
	if mid != 0 {
		st.MachineID = func() (uint16, error) {
			return uint16(mid), nil
		}
	}
	snowflake = NewSonyflake(st)
	if snowflake == nil {
		panic("sonyflake not created")
	}

	sst.StartTime = now
	shortSnowflake = NewShortSonyflake(sst)
	if shortSnowflake == nil {
		panic("short sonyflake not created")
	}
}

// ==========================================Sonyflake
// ==========================================Sonyflake
// ==========================================Sonyflake

// These constants are the bit lengths of Sonyflake ID parts.
const (
	BitLenTime      = 39                               // bit length of time
	BitLenSequence  = 8                                // bit length of sequence number
	BitLenMachineID = 63 - BitLenTime - BitLenSequence // bit length of machine id
)

// Settings configures Sonyflake:
//
// StartTime is the time since which the Sonyflake time is defined as the elapsed time.
// If StartTime is 0, the start time of the Sonyflake is set to "2014-09-01 00:00:00 +0000 UTC".
// If StartTime is ahead of the current time, Sonyflake is not created.
//
// MachineID returns the unique ID of the Sonyflake instance.
// If MachineID returns an error, Sonyflake is not created.
// If MachineID is nil, default MachineID is used.
// Default MachineID returns the lower 16 bits of the private IP address.
//
// CheckMachineID validates the uniqueness of the machine ID.
// If CheckMachineID returns false, Sonyflake is not created.
// If CheckMachineID is nil, no validation is done.
type Settings struct {
	StartTime      time.Time
	MachineID      func() (uint16, error)
	CheckMachineID func(uint16) bool
}

// Sonyflake is a distributed unique ID generator.
type Sonyflake struct {
	mutex       *sync.Mutex
	startTime   int64
	elapsedTime int64
	sequence    uint16
	machineID   uint16
}

// NewSonyflake returns a new Sonyflake configured with the given Settings.
// NewSonyflake returns nil in the following cases:
// - Settings.StartTime is ahead of the current time.
// - Settings.MachineID returns an error.
// - Settings.CheckMachineID returns false.
func NewSonyflake(st Settings) *Sonyflake {
	sf := new(Sonyflake)
	sf.mutex = new(sync.Mutex)
	sf.sequence = uint16(1<<BitLenSequence - 1)

	if st.StartTime.After(time.Now()) {
		return nil
	}
	if st.StartTime.IsZero() {
		sf.startTime = toSonyflakeTime(time.Date(2014, 9, 1, 0, 0, 0, 0, time.UTC))
	} else {
		sf.startTime = toSonyflakeTime(st.StartTime)
	}

	var err error
	if st.MachineID == nil {
		sf.machineID, err = lower16BitPrivateIP()
	} else {
		sf.machineID, err = st.MachineID()
	}
	if err != nil || (st.CheckMachineID != nil && !st.CheckMachineID(sf.machineID)) {
		return nil
	}

	return sf
}

// NextID generates a next unique ID.
// After the Sonyflake time overflows, NextID returns an error.
func (sf *Sonyflake) NextID() (uint64, error) {
	const maskSequence = uint16(1<<BitLenSequence - 1)

	sf.mutex.Lock()
	defer sf.mutex.Unlock()

	current := currentElapsedTime(sf.startTime)
	if sf.elapsedTime < current {
		sf.elapsedTime = current
		sf.sequence = 0
	} else { // sf.elapsedTime >= current
		sf.sequence = (sf.sequence + 1) & maskSequence
		if sf.sequence == 0 {
			sf.elapsedTime++
			overtime := sf.elapsedTime - current
			time.Sleep(sleepTime((overtime)))
		}
	}

	return sf.toID()
}

const sonyflakeTimeUnit = 1e7 // nsec, i.e. 10 msec

func toSonyflakeTime(t time.Time) int64 {
	return t.UTC().UnixNano() / sonyflakeTimeUnit
}

func currentElapsedTime(startTime int64) int64 {
	return toSonyflakeTime(time.Now()) - startTime
}

func sleepTime(overtime int64) time.Duration {
	return time.Duration(overtime)*10*time.Millisecond -
		time.Duration(time.Now().UTC().UnixNano()%sonyflakeTimeUnit)*time.Nanosecond
}

func (sf *Sonyflake) toID() (uint64, error) {
	if sf.elapsedTime >= 1<<BitLenTime {
		return 0, errors.New("over the time limit")
	}

	return uint64(sf.elapsedTime)<<(BitLenSequence+BitLenMachineID) |
		uint64(sf.sequence)<<BitLenMachineID |
		uint64(sf.machineID), nil
}

func privateIPv4() (net.IP, error) {
	as, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	for _, a := range as {
		ipnet, ok := a.(*net.IPNet)
		if !ok || ipnet.IP.IsLoopback() {
			continue
		}

		ip := ipnet.IP.To4()
		if isPrivateIPv4(ip) {
			return ip, nil
		}
	}
	return nil, errors.New("no private ip address")
}

func isPrivateIPv4(ip net.IP) bool {
	return ip != nil &&
		(ip[0] == 10 || ip[0] == 172 && (ip[1] >= 16 && ip[1] < 32) || ip[0] == 192 && ip[1] == 168)
}

func lower16BitPrivateIP() (uint16, error) {
	ip, err := privateIPv4()
	if err != nil {
		return 0, err
	}

	return uint16(ip[2])<<8 + uint16(ip[3]), nil
}

// Decompose returns a set of Sonyflake ID parts.
func Decompose(id uint64) map[string]uint64 {
	const maskSequence = uint64((1<<BitLenSequence - 1) << BitLenMachineID)
	const maskMachineID = uint64(1<<BitLenMachineID - 1)

	msb := id >> 63
	times := id >> (BitLenSequence + BitLenMachineID)
	sequence := id & maskSequence >> BitLenMachineID
	machineID := id & maskMachineID
	return map[string]uint64{
		"id":         id,
		"msb":        msb,
		"times":      times,
		"sequence":   sequence,
		"machine-id": machineID,
	}
}

//================================ short
//================================ short
//================================ short

const (
	ShortBitLenSequence = 1 // bit length of sequence number
	MaxRandomNum        = 99
)

type ShortSettings struct {
	StartTime time.Time
}

type ShortSonyflake struct {
	mutex       *sync.Mutex
	startTime   int64
	elapsedTime int64
	sequence    uint16
	randomNum   uint16
}

func NewShortSonyflake(st ShortSettings) *ShortSonyflake {
	sf := new(ShortSonyflake)
	sf.mutex = new(sync.Mutex)
	sf.sequence = uint16(1<<ShortBitLenSequence - 1)

	if st.StartTime.After(time.Now()) {
		return nil
	}
	if st.StartTime.IsZero() {
		sf.startTime = toSonyflakeTime(time.Date(2014, 9, 1, 0, 0, 0, 0, time.UTC))
	} else {
		sf.startTime = toSonyflakeTime(st.StartTime)
	}
	sf.randomNum = uint16(getRandomInt(MaxRandomNum))

	return sf
}

func (sf *ShortSonyflake) SetRandom() *ShortSonyflake {
	sf.randomNum = uint16(getRandomInt(MaxRandomNum))
	return sf
}

func (sf *ShortSonyflake) ShortNextID() (uint64, error) {
	const maskSequence = uint16(1<<ShortBitLenSequence - 1)
	sf.mutex.Lock()
	defer sf.mutex.Unlock()
	current := currentElapsedTime(sf.startTime)
	if sf.elapsedTime < current {
		sf.elapsedTime = current
		sf.sequence = 0
	} else { // sf.elapsedTime >= current
		sf.sequence = (sf.sequence + 1) & maskSequence
		if sf.sequence == 0 {
			sf.elapsedTime++
			overtime := sf.elapsedTime - current
			time.Sleep(sleepTime((overtime)))
		}
	}

	return sf.shortToID()
}

func (sf *ShortSonyflake) shortToID() (uint64, error) {
	numStr := fmt.Sprintf("%03d%d%02d", sf.elapsedTime%1000, sf.sequence, sf.randomNum)
	return strconv.ParseUint(numStr, 10, 64)
}

func getRandomInt(num int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(num)
}
