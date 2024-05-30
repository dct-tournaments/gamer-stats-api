package leagueoflegends

type PlatformRouting string
type RegionalRouting string

const (
	EUW1PlatformRouting PlatformRouting = "EUW1"

	EuropeRegionalRouting RegionalRouting = "europe"
)

func (s service) fromPlatformRoutingToRegionalRouting(platformRouting PlatformRouting) RegionalRouting {
	switch platformRouting {
	case EUW1PlatformRouting:
		return EuropeRegionalRouting
	default:
		return EuropeRegionalRouting
	}
}
