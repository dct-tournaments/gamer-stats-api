package leagueoflegends

type PlatformRouting string
type RegionalRouting string

const (
	BR1PlatformRouting  PlatformRouting = "br1"
	EUN1PlatformRouting PlatformRouting = "eun1"
	EUW1PlatformRouting PlatformRouting = "euw1"
	JP1PlatformRouting  PlatformRouting = "jp1"
	KR1PlatformRouting  PlatformRouting = "kr1"
	LA1PlatformRouting  PlatformRouting = "la1"
	LA2PlatformRouting  PlatformRouting = "la2"
	NA1PlatformRouting  PlatformRouting = "na1"
	OC1PlatformRouting  PlatformRouting = "oc1"
	TR1PlatformRouting  PlatformRouting = "tr1"
	RU1PlatformRouting  PlatformRouting = "ru1"
	PH2PlatformRouting  PlatformRouting = "ph2"
	SG2PlatformRouting  PlatformRouting = "sg2"
	TH2PlatformRouting  PlatformRouting = "th2"
	TW2PlatformRouting  PlatformRouting = "tw2"
	VN2PlatformRouting  PlatformRouting = "vn2"

	EuropeRegionalRouting   RegionalRouting = "europe"
	AmericasRegionalRouting RegionalRouting = "americas"
	AsiaRegionalRouting     RegionalRouting = "asia"
	SeaRegionalRouting      RegionalRouting = "sea"
)

func (s service) fromPlatformRoutingToRegionalRouting(platformRouting PlatformRouting) RegionalRouting {
	switch platformRouting {
	case BR1PlatformRouting:
		return AmericasRegionalRouting
	case EUW1PlatformRouting, EUN1PlatformRouting, TR1PlatformRouting, RU1PlatformRouting:
		return EuropeRegionalRouting
	case JP1PlatformRouting, KR1PlatformRouting, TW2PlatformRouting:
		return AsiaRegionalRouting
	case LA1PlatformRouting, LA2PlatformRouting, NA1PlatformRouting:
		return AmericasRegionalRouting
	case OC1PlatformRouting, PH2PlatformRouting, SG2PlatformRouting, TH2PlatformRouting, VN2PlatformRouting:
		return SeaRegionalRouting
	default:
		return EuropeRegionalRouting
	}
}
