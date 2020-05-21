package entries

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"strings"

	logger "github.com/sirupsen/logrus"
	"github.com/yokaiio/yokai_server/define"
	"github.com/yokaiio/yokai_server/utils"
)

type Entries struct {
	HeroEntries         map[int32]*define.HeroEntry
	ItemEntries         map[int32]*define.ItemEntry
	EquipEnchantEntries map[int32]*define.EquipEnchantEntry
	TokenEntries        map[int32]*define.TokenEntry
	TalentEntries       map[int32]*define.TalentEntry
	BladeEntries        map[int32]*define.BladeEntry
	RuneEntries         map[int32]*define.RuneEntry
	RuneSuitEntries     map[int32]*define.RuneSuitEntry
	CostLootEntries     map[int32]*define.CostLootEntry
	AttEntries          map[int32]*define.AttEntry
	SceneEntries        map[int32]*define.SceneEntry

	PlayerLevelupEntries map[int32]*define.PlayerLevelupEntry
}

var (
	DefaultEntries *Entries
)

func InitEntries() {
	DefaultEntries = newEntries()
}

func GetHeroEntry(id int32) *define.HeroEntry {
	return DefaultEntries.HeroEntries[id]
}

func GetItemEntry(id int32) *define.ItemEntry {
	return DefaultEntries.ItemEntries[id]
}

func GetEquipEnchantEntry(id int32) *define.EquipEnchantEntry {
	return DefaultEntries.EquipEnchantEntries[id]
}

func GetTokenEntry(id int32) *define.TokenEntry {
	return DefaultEntries.TokenEntries[id]
}

func GetTalentEntry(id int32) *define.TalentEntry {
	return DefaultEntries.TalentEntries[id]
}

func GetBladeEntry(id int32) *define.BladeEntry {
	return DefaultEntries.BladeEntries[id]
}

func GetRuneEntry(id int32) *define.RuneEntry {
	return DefaultEntries.RuneEntries[id]
}

func GetRuneSuitEntry(id int32) *define.RuneSuitEntry {
	return DefaultEntries.RuneSuitEntries[id]
}

func GetCostLootEntry(id int32) *define.CostLootEntry {
	return DefaultEntries.CostLootEntries[id]
}

func GetAttEntry(id int32) *define.AttEntry {
	return DefaultEntries.AttEntries[id]
}

func GetSceneEntry(id int32) *define.SceneEntry {
	return DefaultEntries.SceneEntries[id]
}

func GetPlayerLevelupEntry(id int32) *define.PlayerLevelupEntry {
	return DefaultEntries.PlayerLevelupEntries[id]
}

func newEntries() *Entries {
	var wg utils.WaitGroupWrapper

	m := &Entries{
		HeroEntries:         make(map[int32]*define.HeroEntry),
		ItemEntries:         make(map[int32]*define.ItemEntry),
		EquipEnchantEntries: make(map[int32]*define.EquipEnchantEntry),
		TokenEntries:        make(map[int32]*define.TokenEntry),
		TalentEntries:       make(map[int32]*define.TalentEntry),
		BladeEntries:        make(map[int32]*define.BladeEntry),
		RuneEntries:         make(map[int32]*define.RuneEntry),
		RuneSuitEntries:     make(map[int32]*define.RuneSuitEntry),
		CostLootEntries:     make(map[int32]*define.CostLootEntry),
		AttEntries:          make(map[int32]*define.AttEntry),
		SceneEntries:        make(map[int32]*define.SceneEntry),

		PlayerLevelupEntries: make(map[int32]*define.PlayerLevelupEntry),
	}

	// HeroConfig.json
	wg.Wrap(func() {
		entry := make([]*define.HeroEntry, 0)
		readEntry("HeroConfig.json", &entry, m.HeroEntries)
	})

	// ItemConfig.json
	wg.Wrap(func() {
		entry := make([]*define.ItemEntry, 0)
		readEntry("ItemConfig.json", &entry, m.ItemEntries)
	})

	// EquipEnchantConfig.json
	wg.Wrap(func() {
		entry := make([]*define.EquipEnchantEntry, 0)
		readEntry("EquipEnchantConfig.json", &entry, m.EquipEnchantEntries)
	})

	// TokenConfig.json
	wg.Wrap(func() {
		entry := make([]*define.TokenEntry, 0)
		readEntry("TokenConfig.json", &entry, m.TokenEntries)
	})

	// talent_entry.json
	wg.Wrap(func() {
		entry := make([]*define.TalentEntry, 0)
		readEntry("talent_entry.json", &entry, m.TalentEntries)
	})

	// blade_entry.json
	wg.Wrap(func() {
		entry := make([]*define.BladeEntry, 0)
		readEntry("blade_entry.json", &entry, m.BladeEntries)
	})

	// RuneConfig.json
	wg.Wrap(func() {
		entry := make([]*define.RuneEntry, 0)
		readEntry("RuneConfig.json", &entry, m.RuneEntries)
	})

	// RuneSuitConfig.json
	wg.Wrap(func() {
		entry := make([]*define.RuneSuitEntry, 0)
		readEntry("RuneSuitConfig.json", &entry, m.RuneSuitEntries)
	})

	// cost_loot_entry.json
	wg.Wrap(func() {
		entry := make([]*define.CostLootEntry, 0)
		readEntry("CostLootConfig.json", &entry, m.CostLootEntries)
	})

	// AttConfig.json
	wg.Wrap(func() {
		entry := make([]*define.AttEntry, 0)
		readEntry("AttConfig.json", &entry, m.AttEntries)
	})

	// SceneConfig.json
	wg.Wrap(func() {
		entry := make([]*define.SceneEntry, 0)
		readEntry("SceneConfig.json", &entry, m.SceneEntries)
	})

	// player_levelup_entry.json
	wg.Wrap(func() {
		entry := make([]*define.PlayerLevelupEntry, 0)
		readEntry("PlayerLevelupConfig.json", &entry, m.PlayerLevelupEntries)
	})

	wg.Wait()
	return m
}

// read entries(v) to map(m)
func readEntry(filePath string, v interface{}, m interface{}) {
	absPath := strings.Join([]string{"config/entry/", filePath}, "")
	data, err := ioutil.ReadFile(absPath)
	if err != nil {
		logger.Fatal(err)
	}

	err = json.Unmarshal(data, v)
	if err != nil {
		logger.Fatal(err)
	}

	tp := reflect.TypeOf(v)
	if tp.Kind() == reflect.Ptr || tp.Kind() == reflect.Struct {
		entryField := reflect.ValueOf(v).Elem()
		mapValue := reflect.ValueOf(m)

		for n := 0; n < entryField.Len(); n++ {
			elem := entryField.Index(n)
			key := elem.Elem().FieldByName("ID")

			// if key existed
			keyExist := mapValue.MapIndex(key)
			if keyExist.IsValid() {
				logger.WithFields(logger.Fields{
					"file": absPath,
					"id":   key.Int(),
				}).Fatal("error loading entry")
			}

			mapValue.SetMapIndex(key, elem)
		}
	} else {
		logger.WithFields(logger.Fields{
			"file": absPath,
		}).Fatal("skip reading entry")
	}
}