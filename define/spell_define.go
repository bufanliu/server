package define

import "errors"

//-------------------------------------------------------------------------------
// 技能释放失败错误返回值
//-------------------------------------------------------------------------------
var (
	ErrSpellCooldown         = errors.New("spell cool down")
	ErrSpellCasterStateLimit = errors.New("spell caster state limit")
	ErrSpellCasterCostLimit  = errors.New("spell caster cost limit")
	ErrSpellTargetNotExist   = errors.New("spell target not exist")
	ErrSpellTargetInvalid    = errors.New("spell target invalid")
	ErrSpellTargetStateLimit = errors.New("spell target state limit")
)

const (
	SpellEffectNum      = 3 // 技能效果数量
	SpellEffectParamNum = 6
	AuraFlagNum         = 32
	MaxChargeNum        = 6
	MaxFormSpellNum     = 20
)

//------------------------------------------------------
// 伤害类型
//------------------------------------------------------
type ESchoolType int32

const (
	SchoolType_Begin   ESchoolType = iota
	SchoolType_Physics ESchoolType = iota - 1 // 0 物理伤害
	SchoolType_Magic                          // 1 魔法伤害
	SchoolType_End
)

//------------------------------------------------------
// 技能选中类型
//------------------------------------------------------
type ESelectTargetType int32

const (
	SelectTarget_Begin        ESelectTargetType = iota
	SelectTarget_Null         ESelectTargetType = iota - 1 // 0 无目标
	SelectTarget_Self                                      // 1 自己
	SelectTarget_FriendHPMax                               // 2 友方血量最多
	SelectTarget_FriendHPMin                               // 3 友方血量最少
	SelectTarget_FriendRandom                              // 4 随机友方
	SelectTarget_FriendAll                                 // 5 所有友方
	SelectTarget_EnemySingle                               // 6 单个敌方
	SelectTarget_EnemyRandom                               // 7 随机敌方
	SelectTarget_EnemyAll                                  // 8 所有敌方
	SelectTarget_EnemyHPMax                                // 9 敌方血量最多
	SelectTarget_EnemyHPMin                                // 10 敌方血量最少
	SelectTarget_End
)

//------------------------------------------------------
// 攻击类型
//------------------------------------------------------
type EAttackType int32

const (
	AttackType_Begin EAttackType = iota
	AttackType_Melee EAttackType = iota - 1 // 0 物理攻击
	AttackType_Magic                        // 1 魔法攻击
	AttackType_End
)

//------------------------------------------------------
// 技能类型
//------------------------------------------------------
type ESpellType int32

const (
	SpellType_Begin        ESpellType = iota
	SpellType_Null         ESpellType = iota - 1 // 0 无
	SpellType_Melee                              // 1 普通攻击
	SpellType_TriggerMelee                       // 2 普通攻击触发技能
	SpellType_TriggerAura                        // 3 aura触发技能
	SpellType_BeatBack                           // 4 反击
	SpellType_End
)

//------------------------------------------------------
// 技能结果
//------------------------------------------------------
type EAuraEventType int32

const (
	AuraEvent_Begin       EAuraEventType = iota
	AuraEvent_Null        EAuraEventType = iota - 1 // 0 无
	AuraEvent_NormalHit                             // 1 普通命中
	AuraEvent_CriticalHit                           // 2 暴击
	AuraEvent_Miss                                  // 3 未命中
	AuraEvent_Dodge                                 // 4 躲闪
	AuraEvent_Block                                 // 5 格挡
	AuraEvent_Immnne                                // 6 免疫
	AuraEvent_Absorb                                // 7 吸收
	AuraEvent_End
)

//------------------------------------------------------
// 伤害方式
//------------------------------------------------------
type EDmgInfoType int32

const (
	DmgInfo_Begin     EDmgInfoType = iota
	DmgInfo_Null      EDmgInfoType = iota - 1 // 0 无
	DmgInfo_Damage                            // 1 伤害
	DmgInfo_Heal                              // 2 治疗
	DmgInfo_Placate                           // 3 安抚
	DmgInfo_Enrage                            // 4 激怒
	DmgInfo_AverageHP                         // 5 平均血量
	DmgInfo_End
)

//------------------------------------------------------
// 技能施放方式
//------------------------------------------------------
type ESpellCastType int32

const (
	SpellCast_Begin   ESpellCastType = iota
	SpellCast_Melee   ESpellCastType = iota - 1 // 0 普通攻击
	SpellCast_Instant                           // 1 瞬发施放技能
	SpellCast_Channel                           // 2 引导技能
	SpellCast_End
)

//------------------------------------------------------
// 技能效果类型
//------------------------------------------------------
type ESpellEffectType int32

const (
	SpellEffectType_Begin           ESpellEffectType = iota
	SpellEffectType_Null            ESpellEffectType = iota - 1 // 0 无
	SpellEffectType_Damage                                      // 1 伤害
	SpellEffectType_Heal                                        // 2 治疗
	SpellEffectType_AddAura                                     // 3 添加aura
	SpellEffectType_Placate                                     // 4 安抚
	SpellEffectType_Enrage                                      // 5 激怒
	SpellEffectType_CastSpell                                   // 6 施放技能
	SpellEffectType_Dispel                                      // 7 驱散
	SpellEffectType_ModAuraDuration                             // 8 强化aura作用时间
	SpellEffectType_AverageHP                                   // 9 平均血量
	SpellEffectType_AuraNumDmg                                  // 10 根据buff数量计算伤害
	SpellEffectType_TargetAttDamage                             // 11 根据目标某一属性计算伤害
	SpellEffectType_CasterAttDamage                             // 12 根据施放者某一属性计算伤害
	SpellEffectType_AddLevelAura                                // 13 根据目标等级添加aura
	SpellEffectType_AddStateAura                                // 14 添加状态类aura，并计算状态抗性
	SpellEffectType_AddWrapAura                                 // 15 添加可叠加aura
	SpellEffectType_End
)

//-------------------------------------------------------------------------------
// 技能效果消耗类型
//-------------------------------------------------------------------------------
type ESpellCostType int32

const (
	SpellCostType_Begin ESpellCostType = iota
	SpellCostType_Null                 = iota - 1 // 0 无
	SpellCostType_HP                              // 1 生命
	SpellCostType_MP                              // 2 法力
	SpellCostType_End
)

//-------------------------------------------------------------------------------
// 客户端目标选取方式
//-------------------------------------------------------------------------------
type ETargetChooseType int32

const (
	TargetChooseType_Begin        = iota
	TargetChooseType_Null         = iota - 1 // 0 不需要选取
	TargetChooseType_SingleTarget            // 1 选取单个目标
	TargetChooseType_Area                    // 2 选取地面范围
	TargetChooseType_End
)

//-----------------------------------------------------------------------------
// 技能目的类型
//-----------------------------------------------------------------------------
type ESpellPurposeMaskType int32

const (
	SpellPurposeMask_Begin    = iota
	SpellPurposeMask_None     = iota - 1 // 0 无目的
	SpellPurposeMask_Positive            // 1 有益技能
	SpellPurposeMask_Negative            // 2 有害技能
)

//-----------------------------------------------------------------------------
// 技能进入战斗类型
//-----------------------------------------------------------------------------
type EEnterCombatType int32

const (
	EnterCombatType_Begin        = iota
	EnterCombatType_None         = iota - 1 // 0 不会进入战斗状态
	EnterCombatType_TargetHit               // 1 命中目标进入战斗状态
	EnterCombatType_TargetCombat            // 2 根据目标是否在战斗状态
	EnterCombatType_End
)

//-----------------------------------------------------------------------------
// 选择方式
//-----------------------------------------------------------------------------
type EChooseUnitType int32

const (
	ChooseUnit_Begin  = iota
	ChooseUnit_Caster = iota - 1 // 0 技能施放者
	ChooseUnit_Owner             // 1 技能拥有者
	ChooseUnit_Target            // 2 目标
	ChooseUnit_End
)

//-------------------------------------------------------------------------------
// 效果的目标限制
//-------------------------------------------------------------------------------
type EEffectTargetLimitType int32

const (
	EffectTargetLimit_Begin            = iota
	EffectTargetLimit_Null             = iota - 1 // 0 无限制
	EffectTargetLimit_Self                        // 1 效果只能作用于技能施放者
	EffectTargetLimit_UnSelf                      // 2 效果不会作用于技能施放者
	EffectTargetLimit_Caster_State                // 3 施放者只有处于某种状态效果才会作用
	EffectTargetLimit_Target_State                // 4 目标只有处于某种状态效果才会作用
	EffectTargetLimit_Caster_HP_Low               // 5 施放者当前血量低于最大血量百分比时触发
	EffectTargetLimit_Target_HP_Low               // 6 目标当前血量低于最大血量百分比时触发
	EffectTargetLimit_Target_HP_High              // 7 目标当前血量高于最大血量百分比时触发
	EffectTargetLimit_Pct                         // 8 百分比几率触发(1-10000）
	EffectTargetLimit_Target_AuraNot              // 9 目标处于某个aura，效果不会作用
	EffectTargetLimit_Target_Aura                 // 10 目标处于某个aura，效果会作用
	EffectTargetLimit_Target_Race                 // 11 效果目标种族类型限制
	EffectTargetLimit_Caster_AuraState            // 12 释放者AuraState限制
	EffectTargetLimit_Target_AuraState            // 13 目标AuraState限制
	EffectTargetLimit_Target_GT_Level             // 14 大于目标等级限制
	EffectTargetLimit_Target_LT_Level             // 15 小于等于目标等级限制
	EffectTargetLimit_Caster_AuraPN               // 16 释放者增益减益Buff限制
	EffectTargetLimit_Target_AuraPN               // 17 目标增益减益Buff限制
	EffectTargetLimit_End
)

//-------------------------------------------------------------------------------
// 技能效果状态类型
//-------------------------------------------------------------------------------
type ESpellStateType int32

const (
	SpellState_Begin   ESpellStateType = iota
	SpellState_Prepare ESpellStateType = iota - 1 // 0 吟唱
	SpellState_Casting                            // 1 施放
	SpellState_Finish                             // 2 结束
	SpellState_Idle                               // 3 空闲
	SpellState_End
)

//-------------------------------------------------------------------------------
// 被动技能或buff对主动技能的影响类型
//-------------------------------------------------------------------------------
type ESpellEffectModType int32

const (
	SpellEffectMod_Begin                  ESpellEffectModType = iota
	SpellEffectMod_DmgPct                 ESpellEffectModType = iota - 1 // 0 伤害百分比加成											// misc1 伤害
	SpellEffectMod_Duration                                              // 1 技能效果作用时间加成										// misc1 引导时间
	SpellEffectMod_Threat                                                // 2 仇恨加成													// misc1 仇恨
	SpellEffectMod_Cost                                                  // 3 消耗(需要同步客户端)
	SpellEffectMod_Range                                                 // 4 技能效果作用距离加成(需要同步客户端)						// misc1 距离
	SpellEffectMod_Radius                                                // 5 技能效果作用半径加成(需要同步客户端, 功能暂时未提供)		// misc1 半径
	SpellEffectMod_Crit_Pct                                              // 6 暴击率加成（1-10000）											// misc1 暴击率
	SpellEffectMod_Not_Delay                                             // 7 被攻击时起手时间不会被延长								// misc1==1 起手时间不会被延长
	SpellEffectMod_Prepare_Time                                          // 8 起手时间													// misc  吟唱时间
	SpellEffectMod_Cooldown                                              // 9 冷却时间													// misc 冷却时间
	SpellEffectMod_Hit                                                   // 10 命中加成													// 命中率
	SpellEffectMod_Charge                                                // 11 材料消耗(功能暂时未提供)									// misc1==1 不会消耗材料
	SpellEffectMod_EffectTimes                                           // 12 作用次数													// misc1 aura作用次数
	SpellEffectMod_PassiveDuration                                       // 13 被动持续时间加成											// misc1 aura持续时间
	SpellEffectMod_EffectPct                                             // 14 spell效果触发几率										// misc1 百分比增加量
	SpellEffectMod_CritDmgPct                                            // 15 暴击伤害百分比加成										// misc1 暴击伤害加成百分比
	SpellEffectMod_CategoryCooldown                                      // 16 组冷却时间												// misc 冷却时间
	SpellEffectMod_IgnoreCasterStateCheck                                // 17 无视释放者状态限制判断
	SpellEffectMod_End
)

//-------------------------------------------------------------------------------
// 被动技能或buff对主动技能的影响方式
//-------------------------------------------------------------------------------
type ESpellValueType int32

const (
	SpellValue_Begin ESpellValueType = iota
	SpellValue_FLAT                  = iota - 1 // 0 平值加成
	SpellValue_PCT                              // 1 百分比加成
	SpellValue_MASK                             // 2 掩码
)

//-------------------------------------------------------------------------------
// spell base静态属性
//-------------------------------------------------------------------------------
type SpellBase struct {
	ID         uint32        `json:"_id"`         // 技能id
	FamilyMask uint32        `json:"family_mask"` // 技能所属分组
	FamilyRace EHeroRaceType `json:"family_race"` // 技能对应英灵种族
	SchoolType ESchoolType   `json:"school_type"` // 伤害类型
	HaveVisual bool          `json:"have_visual"` // 是否有特效
	Passive    bool          `json:"passive"`

	MechanicFlags uint32 `json:"mechanic_flags"` // 技能效果的所属控制类型

	AmountEffects [SpellEffectNum]int32  `json:"amount_effects"` // 传递用值影响系数
	MiscID        [SpellEffectNum]uint32 `json:"misc_id"`        // ID参数
	MiscType1     [SpellEffectNum]int32  `json:"misc_type1"`     // 类型参数1
	MiscType2     [SpellEffectNum]int32  `json:"misc_type2"`     // 类型参数2
	MiscValue1    [SpellEffectNum]int32  `json:"misc_value1"`    // 参数值1
	MiscValue2    [SpellEffectNum]int32  `json:"misc_value2"`    // 参数值2
	MiscValue3    [SpellEffectNum]int32  `json:"misc_value3"`    // 参数值3
	BasePoint     [SpellEffectNum]int32  `json:"base_point"`     // 效果动态参数基础值
	LevelPoint    [SpellEffectNum]int32  `json:"level_point"`    // 效果动态参数随机范围基础值
	Multiple      [SpellEffectNum]int32  `json:"multiple"`       // 效果百分比加成(1-10000)
}

//-------------------------------------------------------------------------------
// spell静态属性
//-------------------------------------------------------------------------------
type SpellEntry struct {
	SpellBase `json:",inline"`

	SelectType           ESelectTargetType `json:"select_type"`             // 目标选取方式
	TargetNum            int32             `json:"target_num"`              // 目标数量
	IncludeSelf          bool              `json:"include_self"`            // 目标是否包括自己
	TargetRace           uint32            `json:"target_race"`             // 目标种族
	CasterStateCheckFlag uint32            `json:"caster_state_check_flag"` // 是否判断施放者状态标示(用来标示是否需要判断施放者某个状态的限制)
	CasterStateLimit     uint32            `json:"caster_state_limit"`      // 施放者状态限制
	TargetStateCheckFlag uint32            `json:"target_state_check_flag"` // 是否判断目标状态标示(用来标示是否需要判断某个状态的限制)
	TargetStateLimit     uint32            `json:"target_state_limit"`      // 目标状态限制
	TriggerSpellID       uint32            `json:"trigger_spell_id"`        // 技能释放完成后，触发的技能id
	TriggerSpellProp     int32             `json:"trigger_spell_prop"`      // 触发技能的几率[0-10000]
	PurposeTypeMask      uint32            `json:"purpose_type_mask"`       // 技能目的类型
	TargetAuraState      uint32            `json:"target_aura_state"`       // 目标拥有该aura状态才可使用
	TargetAuraStateNot   uint32            `json:"target_aura_state_not"`   // 目标不拥有该aura状态才可使用
	CasterAuraState      uint32            `json:"caster_aura_state"`       // 释放者拥有该aura状态才可使用
	CasterAuraStateNot   uint32            `json:"caster_aura_state_not"`   // 释放者不拥有该aura状态才可使用

	HitCertainly bool `json:"hit_certainly"` // 该技能是否必命中(如果不是则根据技能攻击方式计算命中率)
	NotCrit      bool `json:"not_crit"`      // 技能不暴击
	NotBlock     bool `json:"not_block"`     // 技能不被格挡
	CanNotArmor  bool `json:"can_not_armor"` // 该技能是否不计算伤害加成减免
	BeatBack     bool `json:"beat_back"`     // 该技能是否能被反击
	GroupDmg     bool `json:"group_dmg"`     // 是否是群体伤害

	SpellHit    int32 `json:"spell_hit"`    // 技能命中加成
	SpellCrit   int32 `json:"spell_crit"`   // 技能暴击加成
	SpellBroken int32 `json:"spell_broken"` // 技能破击加成

	Effects         [SpellEffectNum]ESpellEffectType `json:"effects"`          // 技能效果类型
	EffectsMechanic [SpellEffectNum]uint32           `json:"effects_mechanic"` // 效果所属控制类型

	EffectsTargetLimit    [2][SpellEffectNum]uint32 `json:"effects_target_limit"`     // 效果的目标限制
	EffectsValidMiscValue [2][SpellEffectNum]uint32 `json:"effects_valid_misc_value"` // 效果的目标限制所需参数
	EffectsRaceLimit      [SpellEffectNum]uint32    `json:"effects_race_limit"`       // 效果的目标限制对应种族
	EffectsLimitRaceMod   [SpellEffectNum]uint32    `json:"effects_limit_race_mod"`   // 效果的目标限制所需参数种族加成

	EnergyCost int32 `json:"energy_cost"` // 消耗能量
	RuneCD     int32 `json:"rune_cd"`     // 符文技能CD

	DecayLevel int32 `json:"decay_level"` // 技能概率衰减等级
	DecayRate  int32 `json:"decay_rate"`  // 技能等级衰减系数(1-10000)

	RoundCastMask uint32 `json:"round_cast_mask"` // 技能可释放回合数掩码(0-31回合)
}

//-------------------------------------------------------------------------------
// 冷却时间同步结构
//-------------------------------------------------------------------------------
type SpellCDData struct {
	Entry  *SpellEntry // 技能
	CDTime float32     // 冷却时间
}

//-------------------------------------------------------------------------------
// 伤害信息
//-------------------------------------------------------------------------------
type CalcDamageInfo struct {
	DmgInfoType EDmgInfoType // 伤害方式
	//tagHeroLocation			stCaster;
	//tagHeroLocation			stTarget;
	SchoolType ESchoolType // 伤害类型
	Damage     int32       // 伤害量
	SpellID    uint32      // 技能ID
	//DWORD					dwProcCaster;
	//DWORD					dwProcTarget;
	//DWORD					dwProcEx;							// 技能结果类型掩码

}