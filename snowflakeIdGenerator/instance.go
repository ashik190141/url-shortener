package snowflakeIdGenerator

var Generator *Snowflake

func Init() error {

	generator, err := NewSnowflake(1, 1)
	if err != nil {
		return err
	}

	Generator = generator

	return nil
}