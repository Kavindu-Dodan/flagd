package runtime

import (
	"context"
	"fmt"

	"github.com/open-feature/flagd/core/pkg/evaluator"
	"github.com/open-feature/flagd/core/pkg/logger"
	"github.com/open-feature/flagd/core/pkg/service"
	"github.com/open-feature/flagd/core/pkg/store"
	"github.com/open-feature/flagd/core/pkg/sync"
	syncbuilder "github.com/open-feature/flagd/core/pkg/sync/builder"
	"github.com/open-feature/flagd/core/pkg/telemetry"
	flageval "github.com/open-feature/flagd/flagd/pkg/service/flag-evaluation"
	"github.com/open-feature/flagd/flagd/pkg/service/flag-evaluation/ofrep"
	flagsync "github.com/open-feature/flagd/flagd/pkg/service/flag-sync"
	"go.uber.org/zap"
)

// from_config is a collection of structures and parsers responsible for deriving flagd runtime

const svcName = "flagd"

// Config is the configuration structure derived from startup arguments.
type Config struct {
	MetricExporter    string
	ManagementPort    uint16
	OfrepServicePort  uint16
	OtelCollectorURI  string
	ServiceCertPath   string
	ServiceKeyPath    string
	ServicePort       uint16
	ServiceSocketPath string
	SyncServicePort   uint16

	SyncProviders []sync.SourceConfig
	CORS          []string
}

// FromConfig builds a runtime from startup configurations
// nolint: funlen
func FromConfig(logger *logger.Logger, version string, config Config) (*Runtime, error) {
	telCfg := telemetry.Config{
		MetricsExporter: config.MetricExporter,
		CollectorTarget: config.OtelCollectorURI,
	}

	// register error handling for OpenTelemetry
	telemetry.RegisterErrorHandling(logger)

	// register trace provider for the runtime
	err := telemetry.BuildTraceProvider(context.Background(), logger, svcName, version, telCfg)
	if err != nil {
		return nil, fmt.Errorf("error building trace provider: %w", err)
	}

	// build metrics recorder with startup configurations
	recorder, err := telemetry.BuildMetricsRecorder(context.Background(), svcName, version, telCfg)
	if err != nil {
		return nil, fmt.Errorf("error building metrics recorder: %w", err)
	}

	// build flag store, collect flag sources & fill sources details
	s := store.NewFlags()
	sources := []string{}

	for _, provider := range config.SyncProviders {
		s.FlagSources = append(s.FlagSources, provider.URI)
		s.SourceMetadata[provider.URI] = store.SourceDetails{
			Source:   provider.URI,
			Selector: provider.Selector,
		}
		sources = append(sources, provider.URI)
	}

	// derive evaluator
	evaluator := setupJSONEvaluator(logger, s)

	// derive services

	// connect service
	connectService := flageval.NewConnectService(
		logger.WithFields(zap.String("component", "service")),
		evaluator,
		recorder)

	// ofrep service
	ofrepService, err := ofrep.NewOfrepService(evaluator, ofrep.SvcConfiguration{
		Logger: logger.WithFields(zap.String("component", "OFREPService")),
		Port:   config.OfrepServicePort,
	})
	if err != nil {
		return nil, fmt.Errorf("error creating ofrep service")
	}

	// flag sync service
	flagSyncService, err := flagsync.NewSyncService(flagsync.SvcConfigurations{
		Logger:  logger.WithFields(zap.String("component", "FlagSyncService")),
		Port:    config.SyncServicePort,
		Sources: sources,
		Store:   s,
	})
	if err != nil {
		return nil, fmt.Errorf("error creating sync service: %w", err)
	}

	// build sync providers
	syncLogger := logger.WithFields(zap.String("component", "sync"))
	iSyncs, err := syncProvidersFromConfig(syncLogger, config.SyncProviders)
	if err != nil {
		return nil, err
	}

	options, err := telemetry.BuildConnectOptions(telCfg)
	if err != nil {
		return nil, fmt.Errorf("failed to build connect options, %w", err)
	}

	return &Runtime{
		Logger:       logger.WithFields(zap.String("component", "runtime")),
		Evaluator:    evaluator,
		FlagSync:     flagSyncService,
		OfrepService: ofrepService,
		Service:      connectService,
		ServiceConfig: service.Configuration{
			Port:           config.ServicePort,
			ManagementPort: config.ManagementPort,
			ServiceName:    svcName,
			KeyPath:        config.ServiceKeyPath,
			CertPath:       config.ServiceCertPath,
			SocketPath:     config.ServiceSocketPath,
			CORS:           config.CORS,
			Options:        options,
		},
		SyncImpl: iSyncs,
	}, nil
}

func setupJSONEvaluator(logger *logger.Logger, s *store.Flags) *evaluator.JSON {
	evaluator := evaluator.NewJSON(
		logger,
		s,
		evaluator.WithEvaluator(
			evaluator.FractionEvaluationName,
			evaluator.NewFractional(logger).Evaluate,
		),
		evaluator.WithEvaluator(
			evaluator.StartsWithEvaluationName,
			evaluator.NewStringComparisonEvaluator(logger).StartsWithEvaluation,
		),
		evaluator.WithEvaluator(
			evaluator.EndsWithEvaluationName,
			evaluator.NewStringComparisonEvaluator(logger).EndsWithEvaluation,
		),
		evaluator.WithEvaluator(
			evaluator.SemVerEvaluationName,
			evaluator.NewSemVerComparison(logger).SemVerEvaluation,
		),
		// deprecated: will be removed before v1!
		evaluator.WithEvaluator(
			evaluator.LegacyFractionEvaluationName,
			evaluator.NewLegacyFractional(logger).LegacyFractionalEvaluation,
		),
	)
	return evaluator
}

// syncProvidersFromConfig is a helper to build ISync implementations from SourceConfig
func syncProvidersFromConfig(logger *logger.Logger, sources []sync.SourceConfig) ([]sync.ISync, error) {
	builder := syncbuilder.NewSyncBuilder()
	syncs, err := builder.SyncsFromConfig(sources, logger)
	if err != nil {
		return nil, fmt.Errorf("could not create sync sources from config: %w", err)
	}

	return syncs, nil
}
