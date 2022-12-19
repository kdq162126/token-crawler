package models

type Token struct {
	BaseModelID
	Address     string  `gorm:"index;unique;not null"`
	Name        string  `json:"name"`
	Symbol      string  `json:"symbol"`
	IconUrl     *string `json:"icon_url"`
	SiteUrl     *string `json:"site_url"`
	ProjectCode string  `json:"project_code"`
	Chain       string  `json:"chain"`
	Type        string  `json:"type"`
	Decimals    uint64  `json:"decimals"`
	Price       float64 `json:"price"`
	TotalSupply uint64  `json:"total_supply"`
	IsCore      bool    `json:"is_core" gorm:"default:false"`
	IsVerified  bool    `json:"is_verified" gorm:"default:true"`
	IsActive    bool    `json:"is_active" gorm:"default:true"`
}
