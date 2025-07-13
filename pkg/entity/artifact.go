package entity

import "errors"

var (
	ErrInvalidArtifactID      = errors.New("artifact ID cannot be empty")
	ErrInvalidArtifactType    = errors.New("invalid artifact type")
	ErrInvalidArtifactSet     = errors.New("invalid artifact set")
	ErrInvalidPrimaryStatType = errors.New("invalid primary stat")
	ErrInvalidSubstatType     = errors.New("invalid substat type")
)

type ArtifactType string

const ARTIFACT_TYPE_FLOWER ArtifactType = "FLOWER"
const ARTIFACT_TYPE_PLUME ArtifactType = "PLUME"
const ARTIFACT_TYPE_SANDS ArtifactType = "SANDS"
const ARTIFACT_TYPE_GOBLET ArtifactType = "GOBLET"
const ARTIFACT_TYPE_CIRCLET ArtifactType = "CIRCLET"

type ArtifactSet string

const ARTIFACT_SET_GLADIATORS_FINALOFFERING ArtifactSet = "Gladiator"
const ARTIFACT_SET_WANDERERS_TROUPE ArtifactSet = "Wanderer"
const ARTIFACT_SET_NOBLESSE_OBLIGE ArtifactSet = "Noblesse"
const ARTIFACT_SET_BLOODSTAINED_CHIVALRY ArtifactSet = "Bloodstained"
const ARTIFACT_SET_MAIDENS_BELLSING ArtifactSet = "Maiden"
const ARTIFACT_SET_VIRIDESCENT_VENERER ArtifactSet = "Vermillion"

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

func NewPrimaryStat(statType string, value float64) (*PrimaryStat, error) {
	statTypeEnum := PrimaryStatType(statType)
	switch statTypeEnum {
	case ATK_PERCENT, HP_PERCENT, DEF_PERCENT, ELEMENTAL_MASTERY,
		CRIT_RATE, CRIT_DMG, ENERGY_RECHARGE, PHYSICAL_DMG_BONUS,
		ELEMENTAL_DMG_BONUS, HEALING_BONUS:
	default:
		return nil, ErrInvalidPrimaryStatType
	}

	return &PrimaryStat{
		Type:  statTypeEnum,
		Value: value,
	}, nil
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

func NewSubstat(substatType string, value float64) (*Substat, error) {
	substatTypeEnum := SubstatType(substatType)
	switch substatTypeEnum {
	case SUBSTAT_ATK_PERCENT, SUBSTAT_HP_PERCENT, SUBSTAT_DEF_PERCENT,
		SUBSTAT_ELEMENTAL_MASTERY, SUBSTAT_CRIT_RATE, SUBSTAT_CRIT_DMG,
		SUBSTAT_ENERGY_RECHARGE:
	default:
		return nil, ErrInvalidSubstatType
	}

	return &Substat{
		Type:  substatTypeEnum,
		Value: value,
	}, nil
}

type Artifact struct {
	ID          string
	ArtifactSet ArtifactSet
	Type        ArtifactType
	Level       int
	PrimaryStat PrimaryStat
	Substats    []Substat
}

func NewArtifact(id string, artifactSet, artifactType string, level int, primaryStat PrimaryStat, substats []Substat) (*Artifact, error) {
	if id == "" {
		return nil, ErrInvalidArtifactID
	}

	artifactSetEnum := ArtifactSet(artifactSet)
	switch artifactSetEnum {
	case ARTIFACT_SET_GLADIATORS_FINALOFFERING, ARTIFACT_SET_WANDERERS_TROUPE,
		ARTIFACT_SET_NOBLESSE_OBLIGE, ARTIFACT_SET_BLOODSTAINED_CHIVALRY,
		ARTIFACT_SET_MAIDENS_BELLSING, ARTIFACT_SET_VIRIDESCENT_VENERER:
	default:
		return nil, ErrInvalidArtifactSet
	}

	artifactTypeEnum := ArtifactType(artifactType)
	switch artifactTypeEnum {
	case ARTIFACT_TYPE_FLOWER, ARTIFACT_TYPE_PLUME, ARTIFACT_TYPE_SANDS,
		ARTIFACT_TYPE_GOBLET, ARTIFACT_TYPE_CIRCLET:
	default:
		return nil, ErrInvalidArtifactType
	}

	return &Artifact{
		ID:          id,
		ArtifactSet: artifactSetEnum,
		Type:        artifactTypeEnum,
		Level:       level,
		PrimaryStat: primaryStat,
		Substats:    substats,
	}, nil
}
