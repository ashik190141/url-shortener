package snowflakeIdGenerator

import (
	"fmt"
	"time"
)

func currentTimestamp() int64 {
	return time.Now().UnixMilli()
}

func waitNextMillis(lastTimestamp int64) int64 {
	timestamp := currentTimestamp() - customEpoch
	for timestamp <= lastTimestamp {
		timestamp = currentTimestamp() - customEpoch
	}
	return timestamp
}

func (s *Snowflake) Generate() (uint64, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	currentTimestamp := currentTimestamp() - customEpoch
	if currentTimestamp < s.lastTimestamp {
		return 0, fmt.Errorf("clock moved backwards. Refusing to generate id for %d milliseconds", s.lastTimestamp-currentTimestamp)
	}

	if currentTimestamp == s.lastTimestamp {
		s.sequence = (s.sequence + 1) & maxSequence
		if s.sequence == 0 {
			currentTimestamp = waitNextMillis(s.lastTimestamp)
			
		}
	}else{
		s.sequence = 0
	}
	s.lastTimestamp = currentTimestamp
	return uint64(currentTimestamp)<<timestampShift | uint64(s.datacenterID)<<datacenterShift | uint64(s.machineID)<<machineShift | uint64(s.sequence), nil
}
