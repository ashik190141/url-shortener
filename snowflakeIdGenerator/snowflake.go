package snowflakeIdGenerator

import (
	"fmt"
	"sync"
)

type Snowflake struct {
	mutex sync.Mutex

	lastTimestamp int64
	sequence      int64

	datacenterID int64
	machineID    int64
}

func NewSnowflake(
	datacenterID,
	machineID int64,
) (*Snowflake, error) {

	if datacenterID < 0 || datacenterID > 31 {
		return nil, fmt.Errorf(
			"datacenterID must be between 0 and 31",
		)
	}

	if machineID < 0 || machineID > 31 {
		return nil, fmt.Errorf(
			"machineID must be between 0 and 31",
		)
	}

	return &Snowflake{
		lastTimestamp: -1,
		sequence:      0,
		datacenterID:  datacenterID,
		machineID:     machineID,
	}, nil
}