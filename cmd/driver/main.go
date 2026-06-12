package main

import (
	"fmt"
	"math/rand/v2"
	"time"

	"github.com/Anthill-Software/GatherPipe/core"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
)

// Implémentation concrète du Driver
type DummyDriver struct {
	logger hclog.Logger
}

// Init implémente la méthode requise par l'interface core.Driver
func (d *DummyDriver) Init(config map[string]string) error {
	// Ici, on peut lire des valeurs comme config["unit"]
	d.logger.Info("Initialisé avec succès")
	return nil
}

func (p *DummyDriver) Name() (string, error) {
	return "DummyDriver", nil
}

func (d *DummyDriver) Fetch() ([]core.Metric, error) {
	var metrics []core.Metric
	now := time.Now()

	metrics = append(metrics, core.Metric{
		Timestamp: now,
		ID:        "sensor.temperature",
		Value:     fmt.Sprintf("%f", 20.0+rand.Float64()*10),
		Format:    "float",
		Unit:      "°C",
	})
	metrics = append(metrics, core.Metric{
		Timestamp: now,
		ID:        "sensor.humidity",
		Value:     fmt.Sprintf("%f", 40.0+rand.Float64()*20),
		Format:    "float",
		Unit:      "%",
	})

	d.logger.Debug("Collecte réussie", "count", len(metrics))
	return metrics, nil
}

func main() {
	pluginLogger := hclog.New(&hclog.LoggerOptions{
		Level:       hclog.Debug,
		DisableTime: true,
	})

	driver := &DummyDriver{
		logger: pluginLogger,
	}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: core.Handshake,
		Plugins: map[string]plugin.Plugin{
			"driver": &core.DriverPlugin{Impl: driver},
		},
		Logger: pluginLogger,
	})
}
