package main

import (
	"github.com/dct-tournaments/gamer-stats-api/pkg/leagueoflegends"
	"github.com/ilyakaznacheev/cleanenv"
)

type config struct {
	LeagueOfLegends leagueoflegends.Config
}

func NewConfig() (config, error) {
	c := config{}

	err := cleanenv.ReadEnv(&c)

	return c, err
}
