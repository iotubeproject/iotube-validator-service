package registry

type (
	Asset struct {
		BlockHeight uint64 `gorm:"unsigned;index;primaryKey;autoIncrement:false"`
		TxHash      string `gorm:"size:66;index"`
		LogIndex    uint   `gorm:"unsigned;index;primaryKey;autoIncrement:false"`
		ID          uint64 `gorm:"unsigned;index"`
		Deactive    bool   `gorm:"type:bool"`
	}

	Tube struct {
		BlockHeight uint64 `gorm:"unsigned;index;primaryKey;autoIncrement:false"`
		TxHash      string `gorm:"size:66;index"`
		LogIndex    uint   `gorm:"unsigned;index;primaryKey;autoIncrement:false"`
		ID          uint64 `gorm:"unsigned;index"`
		Deactive    bool   `gorm:"type:bool"`
	}

	AssetOnTube struct {
		BlockHeight  uint64 `gorm:"unsigned;index;primaryKey;autoIncrement:false"`
		TxHash       string `gorm:"size:66;index"`
		LogIndex     uint   `gorm:"unsigned;index;primaryKey;autoIncrement:false"`
		AssetID      uint64 `gorm:"unsigned;index"`
		TubeID       uint64 `gorm:"unsigned;index"`
		AssetAddress string `gorm:"size:42;index"`
		Active       bool   `gorm:"type:bool"`
	}

	Validator struct {
		BlockHeight uint64 `gorm:"unsigned;index;primaryKey;autoIncrement:false"`
		TxHash      string `gorm:"size:66;index"`
		LogIndex    uint   `gorm:"unsigned;index;primaryKey;autoIncrement:false"`
		Address     string `gorm:"size:42;index"`
		Genre       uint64 `gorm:"unsigned;index"`
		URI         string `gorm:""`
	}

	AssetUpperBound struct {
		BlockHeight uint64 `gorm:"unsigned;index;primaryKey;autoIncrement:false"`
		TxHash      string `gorm:"size:66;index"`
		LogIndex    uint   `gorm:"unsigned;index;primaryKey;autoIncrement:false"`
		AssetID     uint64 `gorm:"unsigned;index"`
		UpperBound  string `gorm:""`
	}

	Meta struct {
		Key    string `gorm:"primaryKey"`
		Height uint64 `gorm:"unsigned"`
	}
)

func (Validator) TableName() string {
	return "r_validators"
}

func (AssetUpperBound) TableName() string {
	return "r_asset_upper_bounds"
}

func (Asset) TableName() string {
	return "r_assets"
}

func (Tube) TableName() string {
	return "r_tubes"
}

func (AssetOnTube) TableName() string {
	return "r_asset_on_tube"
}

func (Meta) TableName() string {
	return "r_meta"
}
