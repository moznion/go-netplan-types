package netplan

type LACPRate struct {
	val        lacpRate
	isAssigned bool
}

type lacpRate string

const (
	slowLACPRate lacpRate = "slow"
	fastLACPRate          = "fast"
)

func SlowLACPRate() *LACPRate {
	return &LACPRate{
		val:        slowLACPRate,
		isAssigned: true,
	}
}

func FastLACPRate() *LACPRate {
	return &LACPRate{
		val:        fastLACPRate,
		isAssigned: true,
	}
}

func (lr *LACPRate) MarshalYAML() (interface{}, error) {
	if lr.isAssigned {
		return lr.val, nil
	}
	return nil, nil
}

func (lr *LACPRate) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val lacpRate
	if err := unmarshal(&val); err != nil {
		return err
	}

	lr.val = val
	lr.isAssigned = true

	return nil
}
