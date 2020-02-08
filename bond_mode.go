package netplan

type BondMode struct {
	val        bondMode
	isAssigned bool
}

type bondMode string

const (
	balanceRRBondMode    bondMode = "balance-rr"
	activeBackupBondMode          = "active-backup"
	balanceXORBondMode            = "balance-xor"
	broadcastBondMode             = "broadcast"
	ieee8023adBondMode            = "802.3ad"
	balanceTLBBondMode            = "balance-tlb"
	balanceALBBondMode            = "balance-alb"
)

func BalanceRRBondMode() *BondMode {
	return &BondMode{
		val:        balanceRRBondMode,
		isAssigned: true,
	}
}

func ActiveBackupBondMode() *BondMode {
	return &BondMode{
		val:        activeBackupBondMode,
		isAssigned: true,
	}
}

func BalanceXORBondMode() *BondMode {
	return &BondMode{
		val:        balanceXORBondMode,
		isAssigned: true,
	}
}

func BroadcastBondMode() *BondMode {
	return &BondMode{
		val:        broadcastBondMode,
		isAssigned: true,
	}
}

func IEEE8023adBondMode() *BondMode {
	return &BondMode{
		val:        ieee8023adBondMode,
		isAssigned: true,
	}
}

func BalanceTLBBondMode() *BondMode {
	return &BondMode{
		val:        balanceTLBBondMode,
		isAssigned: true,
	}
}

func BalanceALBBondMode() *BondMode {
	return &BondMode{
		val:        balanceALBBondMode,
		isAssigned: true,
	}
}

func (bm *BondMode) MarshalYAML() (interface{}, error) {
	if bm.isAssigned {
		return bm.val, nil
	}
	return nil, nil
}

func (bm *BondMode) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val bondMode
	if err := unmarshal(&val); err != nil {
		return err
	}

	bm.val = val
	bm.isAssigned = true

	return nil
}
