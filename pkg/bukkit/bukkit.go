package bukkit

import (
	"github.com/eujandergois/mc-control/pkg/config"
)

type BukkitSettings struct {
	AllowEnd            bool   `yaml:"allow-end"`
	WarnOnOverload      bool   `yaml:"warn-on-overload"`
	PermissionsFileName string `yaml:"permissions-file"`
	UpdateFolderName    string `yaml:"update-folder"`
	PluginProfiling     bool   `yaml:"plugin-profiling"`
	ConnectionThrottle  int    `yaml:"connection-throttle"`
	QueryPlugins        bool   `yaml:"query-plugins"`
	DeprecatedVerbose   string `yaml:"deprecated-verbose"`
	ShutdownMessage     string `yaml:"shutdown-message"`
	MinimumAPI          string `yaml:"minimum-api"`
}

type BukkitSpawnLimits struct {
	Monsters     int `yaml:"monsters"`
	Animals      int `yaml:"animals"`
	WaterAnimals int `yaml:"water-animals"`
	WaterAmbient int `yaml:"water-ambient"`
	Ambient      int `yaml:"ambient"`
}

type BukkitChunkGC struct {
	PeriodInTicks int `yaml:"period-in-ticks"`
}

type BukkitTicksPer struct {
	AnimalSpawns       int `yaml:"animal-spawns"`
	MonsterSpawns      int `yaml:"monster-spawns"`
	WaterSpawns        int `yaml:"water-spawns"`
	WaterAmbientSpawns int `yaml:"water-ambient-spawns"`
	AmbientSpawns      int `yaml:"ambient-spawns"`
	Autosave           int `yaml:"autosave"`
}

type BukkitYAML struct {
	Settings    BukkitSettings    `yaml:"settings"`
	SpawnLimits BukkitSpawnLimits `yaml:"spawn-limits"`
	ChunkGC     BukkitChunkGC     `yaml:"chunk-gc"`
	TicksPer    BukkitTicksPer    `yaml:"ticks-per"`
}

func GetBukkitSettings() (*BukkitYAML, error) {
	var bukkitSettings BukkitYAML
	err := config.LoadYAML("bukkit.yml", &bukkitSettings)
	if err != nil {
		return nil, err
	}

	return &bukkitSettings, nil
}
