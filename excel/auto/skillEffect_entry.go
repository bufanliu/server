package auto

import (
	"bitbucket.org/funplus/server/excel"
	"bitbucket.org/funplus/server/utils"
	"github.com/mitchellh/mapstructure"
	"github.com/rs/zerolog/log"
	"github.com/shopspring/decimal"
)

var skillEffectEntries *SkillEffectEntries //SkillEffect.xlsx全局变量

// SkillEffect.xlsx属性表
type SkillEffectEntry struct {
	Id              int32             `json:"Id,omitempty"`              // 主键
	Showid          string            `json:"Showid,omitempty"`          //效果表演
	IsEffectHit     int32             `json:"IsEffectHit,omitempty"`     //判定类型
	Prob            int32             `json:"Prob,omitempty"`            //触发概率
	EffectType      int32             `json:"EffectType,omitempty"`      //效果类型
	ParameterA      decimal.Decimal   `json:"ParameterA,omitempty"`      //参数1
	ParameterB      decimal.Decimal   `json:"ParameterB,omitempty"`      //参数2
	ParameterC      decimal.Decimal   `json:"ParameterC,omitempty"`      //参数3
	ParameterD      decimal.Decimal   `json:"ParameterD,omitempty"`      //参数4
	ParameterE      decimal.Decimal   `json:"ParameterE,omitempty"`      //参数5
	ArryA           []decimal.Decimal `json:"ArryA,omitempty"`           //数组1
	ArryB           []decimal.Decimal `json:"ArryB,omitempty"`           //数组2
	ArryC           []decimal.Decimal `json:"ArryC,omitempty"`           //数组3
	RigidityTime    decimal.Decimal   `json:"RigidityTime,omitempty"`    //僵直时间
	HitBackDistance decimal.Decimal   `json:"HitBackDistance,omitempty"` //击退距离
	RangeType       int32             `json:"RangeType,omitempty"`       //目标范围
	TargetType      int32             `json:"TargetType,omitempty"`      //目标类型
	SkillLaunch     int32             `json:"SkillLaunch,omitempty"`     //发起类型
	TargetLength    int32             `json:"TargetLength,omitempty"`    //范围长
	TargetWide      int32             `json:"TargetWide,omitempty"`      //范围宽
	Scope           int32             `json:"Scope,omitempty"`           //作用对象
}

// SkillEffect.xlsx属性表集合
type SkillEffectEntries struct {
	Rows map[int32]*SkillEffectEntry `json:"Rows,omitempty"` //
}

func init() {
	excel.AddEntryLoader("SkillEffect.xlsx", (*SkillEffectEntries)(nil))
}

func (e *SkillEffectEntries) Load(excelFileRaw *excel.ExcelFileRaw) error {

	skillEffectEntries = &SkillEffectEntries{
		Rows: make(map[int32]*SkillEffectEntry, 100),
	}

	for _, v := range excelFileRaw.CellData {
		entry := &SkillEffectEntry{}
		err := mapstructure.Decode(v, entry)
		if !utils.ErrCheck(err, "decode excel data to struct failed", v) {
			return err
		}

		skillEffectEntries.Rows[entry.Id] = entry
	}

	log.Info().Str("excel_file", excelFileRaw.Filename).Msg("excel load success")
	return nil

}

func GetSkillEffectEntry(id int32) (*SkillEffectEntry, bool) {
	entry, ok := skillEffectEntries.Rows[id]
	return entry, ok
}

func GetSkillEffectSize() int32 {
	return int32(len(skillEffectEntries.Rows))
}

func GetSkillEffectRows() map[int32]*SkillEffectEntry {
	return skillEffectEntries.Rows
}
