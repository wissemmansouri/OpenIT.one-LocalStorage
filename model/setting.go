package model 

const (
        SINGLE = iota 
        SITE
        STYLE
        PREVIEW
        GLOBAL
        ARIA2
        INDEX
        GITHUB
)

const (
        PUBLIC = iota
        PRIVATE 
        READONLY
        DEPRECATED
)

type SettingItem struct {
        Key     string `json:"key" gorm:"primaryKey" binding:"required"` // unique key
        Value   string `json:"value"`                                    // value
	      Help    string `json:"help"`                                     // help message
	      Type    string `json:"type"`                                     // string, number, bool, select
	      Options string `json:"options"`                                  // values for select
	      Group   int    `json:"group"`                                    // use to group setting in frontend
	      Flag    int    `json:"flag"
}

func (s SettingItem) IsDeprecated() bool {
        return s.Flag == DEPRECATED
}
