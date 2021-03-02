package auto

import (
	"bitbucket.org/funplus/server/excel"
	"bitbucket.org/funplus/server/utils"
	"github.com/mitchellh/mapstructure"
	"github.com/rs/zerolog/log"
)

var	crystalAttRepoEntries	*CrystalAttRepoEntries	//CrystalAttRepo.xlsx全局变量

// CrystalAttRepo.xlsx属性表
type CrystalAttRepoEntry struct {
	Id             	int32               	`json:"Id,omitempty"`	// 主键       
	Pos            	int32               	`json:"Pos,omitempty"`	//晶石位置      
	Quality        	int32               	`json:"Quality,omitempty"`	//晶石品质      
	Type           	int32               	`json:"Type,omitempty"`	//属性库类型     
	AttType        	int32               	`json:"AttType,omitempty"`	//属性类型      
	AttValueMin    	int32               	`json:"AttValueMin,omitempty"`	//属性值最小值    
	AttValueMax    	int32               	`json:"AttValueMax,omitempty"`	//属性值最大值    
	AttGrowRatio   	int32               	`json:"AttGrowRatio,omitempty"`	//属性成长率     
	AttWeight      	int32               	`json:"AttWeight,omitempty"`	//属性权重      
}

// CrystalAttRepo.xlsx属性表集合
type CrystalAttRepoEntries struct {
	Rows           	map[int32]*CrystalAttRepoEntry	`json:"Rows,omitempty"`	//          
}

func  init()  {
	excel.AddEntryLoader("CrystalAttRepo.xlsx", (*CrystalAttRepoEntries)(nil))
}

func (e *CrystalAttRepoEntries) Load(excelFileRaw *excel.ExcelFileRaw) error {
	
	crystalAttRepoEntries = &CrystalAttRepoEntries{
		Rows: make(map[int32]*CrystalAttRepoEntry, 100),
	}

	for _, v := range excelFileRaw.CellData {
		entry := &CrystalAttRepoEntry{}
		err := mapstructure.Decode(v, entry)
		if !utils.ErrCheck(err, "decode excel data to struct failed", v) {
			return err
		}

	 	crystalAttRepoEntries.Rows[entry.Id] = entry
	}

	log.Info().Str("excel_file", excelFileRaw.Filename).Msg("excel load success")
	return nil
	
}

func  GetCrystalAttRepoEntry(id int32) (*CrystalAttRepoEntry, bool) {
	entry, ok := crystalAttRepoEntries.Rows[id]
	return entry, ok
}

func  GetCrystalAttRepoSize() int32 {
	return int32(len(crystalAttRepoEntries.Rows))
}

func  GetCrystalAttRepoRows() map[int32]*CrystalAttRepoEntry {
	return crystalAttRepoEntries.Rows
}

