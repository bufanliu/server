package auto

import (
	"bitbucket.org/funplus/server/excel"
	"bitbucket.org/funplus/server/utils"
	"github.com/mitchellh/mapstructure"
	"github.com/rs/zerolog/log"
)

var equipLevelupEntries *EquipLevelupEntries //EquipLevelup.xlsx全局变量

// EquipLevelup.xlsx属性表
type EquipLevelupEntry struct {
	Id           int32   `json:"Id,omitempty"`           // 主键
	Exp          []int32 `json:"Exp,omitempty"`          //不同品质升级所需经验值
	PromoteLimit int32   `json:"PromoteLimit,omitempty"` //突破次数限制
}

// EquipLevelup.xlsx属性表集合
type EquipLevelupEntries struct {
	Rows map[int32]*EquipLevelupEntry `json:"Rows,omitempty"` //
}

func init() {
	excel.AddEntryLoader("EquipLevelup.xlsx", (*EquipLevelupEntries)(nil))
}

func (e *EquipLevelupEntries) Load(excelFileRaw *excel.ExcelFileRaw) error {

	equipLevelupEntries = &EquipLevelupEntries{
		Rows: make(map[int32]*EquipLevelupEntry, 100),
	}

	for _, v := range excelFileRaw.CellData {
		entry := &EquipLevelupEntry{}
		err := mapstructure.Decode(v, entry)
		if !utils.ErrCheck(err, "decode excel data to struct failed", v) {
			return err
		}

		equipLevelupEntries.Rows[entry.Id] = entry
	}

	log.Info().Str("excel_file", excelFileRaw.Filename).Msg("excel load success")
	return nil

}

func GetEquipLevelupEntry(id int32) (*EquipLevelupEntry, bool) {
	entry, ok := equipLevelupEntries.Rows[id]
	return entry, ok
}

func GetEquipLevelupSize() int32 {
	return int32(len(equipLevelupEntries.Rows))
}

func GetEquipLevelupRows() map[int32]*EquipLevelupEntry {
	return equipLevelupEntries.Rows
}
