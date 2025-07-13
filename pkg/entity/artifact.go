package entity

type ArtifactType string

const ARTIFACT_TYPE_FLOWER ArtifactType = "FLOWER"
const ARTIFACT_TYPE_PLUME ArtifactType = "PLUME"
const ARTIFACT_TYPE_SANDS ArtifactType = "SANDS"
const ARTIFACT_TYPE_GOBLET ArtifactType = "GOBLET"
const ARTIFACT_TYPE_CIRCLET ArtifactType = "CIRCLET"

type ArtifactSet string

const ARTIFACT_SET_GLADIATORS_FINALOFFERING ArtifactSet = "Gladiator's Finale"
const ARTIFACT_SET_WANDERERS_TROUPE ArtifactSet = "Wanderer's Troupe"
const ARTIFACT_SET_NOBLESSE_OBLIGE ArtifactSet = "Noblesse Oblige"
const ARTIFACT_SET_BLOODSTAINED_CHIVALRY ArtifactSet = "Bloodstained Chivalry"
const ARTIFACT_SET_MAIDENS_BELLSING ArtifactSet = "Maiden's Beloved"
const ARTIFACT_SET_VIRIDESCENT_VENERER ArtifactSet = "Vermillion Hereafter"

type PrimaryStatType string

const ATK_PERCENT PrimaryStatType = "ATK_PERCENT"
const HP_PERCENT PrimaryStatType = "HP_PERCENT"
const DEF_PERCENT PrimaryStatType = "DEF_PERCENT"
const ELEMENTAL_MASTERY PrimaryStatType = "ELEMENTAL_MASTERY"
const CRIT_RATE PrimaryStatType = "CRIT_RATE"
const CRIT_DMG PrimaryStatType = "CRIT_DMG"
const ENERGY_RECHARGE PrimaryStatType = "ENERGY_RECHARGE"
const PHYSICAL_DMG_BONUS PrimaryStatType = "PHYSICAL_DMG_BONUS"
const ELEMENTAL_DMG_BONUS PrimaryStatType = "ELEMENTAL_DMG_BONUS"
const HEALING_BONUS PrimaryStatType = "HEALING_BONUS"

type PrimaryStat struct {
	Type  PrimaryStatType
	Value float64
}

func NewPrimaryStat(statType PrimaryStatType, value float64) PrimaryStat {
	return PrimaryStat{
		Type:  statType,
		Value: value,
	}
}

type SubstatType string

const SUBSTAT_ATK_PERCENT SubstatType = "ATK_PERCENT"
const SUBSTAT_HP_PERCENT SubstatType = "HP_PERCENT"
const SUBSTAT_DEF_PERCENT SubstatType = "DEF_PERCENT"
const SUBSTAT_ELEMENTAL_MASTERY SubstatType = "ELEMENTAL_MASTERY"
const SUBSTAT_CRIT_RATE SubstatType = "CRIT_RATE"
const SUBSTAT_CRIT_DMG SubstatType = "CRIT_DMG"
const SUBSTAT_ENERGY_RECHARGE SubstatType = "ENERGY_RECHARGE"

type Substat struct {
	Type  SubstatType
	Value float64
}

func NewSubstat(substatType SubstatType, value float64) Substat {
	return Substat{
		Type:  substatType,
		Value: value,
	}
}

type Artifact struct {
	ID          string
	ArtifactSet ArtifactSet
	Type        ArtifactType
	Level       int
	PrimaryStat PrimaryStat
	Substats    []Substat
}

func NewArtifact(id string, artifactSet ArtifactSet, artifactType ArtifactType, level int, primaryStat PrimaryStat, substats []Substat) *Artifact {
	return &Artifact{
		ID:          id,
		ArtifactSet: artifactSet,
		Type:        artifactType,
		Level:       level,
		PrimaryStat: primaryStat,
		Substats:    substats,
	}
}
