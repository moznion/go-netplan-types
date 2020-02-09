package netplan

// BondMode represents netplan's bond mode as nillable.
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

// BalanceRRBondMode returns `balance-rr` bond mode.
func BalanceRRBondMode() *BondMode {
	return &BondMode{
		val:        balanceRRBondMode,
		isAssigned: true,
	}
}

// ActiveBackupBondMode returns `active-backup` bond mode.
func ActiveBackupBondMode() *BondMode {
	return &BondMode{
		val:        activeBackupBondMode,
		isAssigned: true,
	}
}

// BalanceXORBondMode returns `balance-xor` bond mode.
func BalanceXORBondMode() *BondMode {
	return &BondMode{
		val:        balanceXORBondMode,
		isAssigned: true,
	}
}

// BroadcastBondMode returns `broadcast` bond mode.
func BroadcastBondMode() *BondMode {
	return &BondMode{
		val:        broadcastBondMode,
		isAssigned: true,
	}
}

// IEEE8023adBondMode returns `802.3ad` bond mode.
func IEEE8023adBondMode() *BondMode {
	return &BondMode{
		val:        ieee8023adBondMode,
		isAssigned: true,
	}
}

// BalanceTLBBondMode returns `balance-tlb` bond mode.
func BalanceTLBBondMode() *BondMode {
	return &BondMode{
		val:        balanceTLBBondMode,
		isAssigned: true,
	}
}

// BalanceALBBondMode returns `balance-alb` bond mode.
func BalanceALBBondMode() *BondMode {
	return &BondMode{
		val:        balanceALBBondMode,
		isAssigned: true,
	}
}

// MarshalYAML marshals BondMode as YAML.
// This method used on marshaling YAML internally.
func (bm *BondMode) MarshalYAML() (interface{}, error) {
	if bm.isAssigned {
		return bm.val, nil
	}
	return nil, nil
}

// UnmarshalYAML unmarshals BondMode as YAML.
// This method used on unmarshaling YAML internally.
func (bm *BondMode) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val bondMode
	if err := unmarshal(&val); err != nil {
		return err
	}

	bm.val = val
	bm.isAssigned = true

	return nil
}
