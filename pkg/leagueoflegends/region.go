package leagueoflegends

type PlatformRouting string
type RegionalRouting string

const (
	EUW1PlatformRouting PlatformRouting = "EUW1"
	BR1PlatformRouting  PlatformRouting = "BR1"

	EuropeRegionalRouting   RegionalRouting = "europe"
	AmericasRegionalRouting RegionalRouting = "americas"
)

func (s service) fromPlatformRoutingToRegionalRouting(platformRouting PlatformRouting) RegionalRouting {
	switch platformRouting {
	case EUW1PlatformRouting:
		return EuropeRegionalRouting
	case BR1PlatformRouting:
		return AmericasRegionalRouting
	default:
		return EuropeRegionalRouting
	}
}
