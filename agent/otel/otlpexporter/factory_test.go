package otlpexporter_test

import (
	"context"
	"github.com/ns1labs/orb/agent/otel/otlpexporter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/config/configtest"
	"go.opentelemetry.io/collector/exporter/exporterhelper"

	"testing"
)

func TestCreateDefaultConfig(t *testing.T) {
	factory := otlpexporter.NewFactory()
	cfg := factory.CreateDefaultConfig()
	assert.NotNil(t, cfg, "failed to create default config")
	assert.NoError(t, configtest.CheckConfigStruct(cfg))

	ocfg, ok := factory.CreateDefaultConfig().(*otlpexporter.Config)
	assert.True(t, ok)
	assert.Equal(t, ocfg.RetrySettings, exporterhelper.DefaultRetrySettings())
	assert.Equal(t, ocfg.QueueSettings, exporterhelper.DefaultQueueSettings())
	assert.Equal(t, ocfg.TimeoutSettings, exporterhelper.DefaultTimeoutSettings())
}

func TestCreateMetricsExporter(t *testing.T) {
	factory := otlpexporter.NewFactory()
	cfg := factory.CreateDefaultConfig().(*otlpexporter.Config)
	cfg.GRPCClientSettings.Endpoint = "localhost:1234"

	set := componenttest.NewNopExporterCreateSettings()
	oexp, err := factory.CreateMetricsExporter(context.Background(), set, cfg)
	require.Nil(t, err)
	require.NotNil(t, oexp)
}
