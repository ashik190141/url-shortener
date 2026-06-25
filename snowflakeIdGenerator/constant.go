package snowflakeIdGenerator

const (
	customEpoch     int64 = 1288834974657
	sequenceBits    uint8 = 12
	machineBits     uint8 = 5
	datacenterBits  uint8 = 5

	maxSequence int64 = -1 ^ (-1 << sequenceBits)

	machineShift    uint8 = sequenceBits
	datacenterShift uint8 = sequenceBits + machineBits
	timestampShift  uint8 = sequenceBits + machineBits + datacenterBits
)