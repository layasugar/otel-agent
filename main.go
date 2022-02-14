package main

import (
	"context"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/jaegerreceiver"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/service"
	"go.opentelemetry.io/collector/service/defaultcomponents"
	"log"
)

func main() {
	factories, err := components()
	if err != nil {
		log.Fatalf("failed to build components: %v", err)
	}

	info := component.BuildInfo{
		Command:     "otel-agent",
		Description: "Custom OpenTelemetry agent distribution",
		Version:     "1.0.0",
	}

	set := service.CollectorSettings{
		BuildInfo:      info,
		Factories:      factories,
		ConfigProvider: service.MustNewDefaultConfigProvider([]string{"otel-agent.yaml"}, []string{}),
	}

	app, err := service.New(set)
	if err != nil {
		log.Fatal("failed to construct the application: %w", err)
	}

	err = app.Run(context.Background())
	if err != nil {
		log.Fatal("application run finished with error: %w", err)
	}
}

func components() (component.Factories, error) {
	factories, err := defaultcomponents.Components()
	if err != nil {
		return component.Factories{}, err
	}

	// 接收器
	receivers := []component.ReceiverFactory{
		jaegerreceiver.NewFactory(),
	}
	for _, pr := range factories.Receivers {
		receivers = append(receivers, pr)
	}
	factories.Receivers, err = component.MakeReceiverFactoryMap(receivers...)
	if err != nil {
		return component.Factories{}, err
	}

	// 处理器
	//processors := []component.ProcessorFactory{
	//	tailsamplingprocessor.NewFactory(),
	//}
	//for _, pr := range factories.Processors {
	//	processors = append(processors, pr)
	//}
	//factories.Processors, err = component.MakeProcessorFactoryMap(processors...)
	//if err != nil {
	//	return component.Factories{}, err
	//}

	//// 导出器
	//exporters := []component.ExporterFactory{
	//	prometheusexporter.NewFactory(),
	//}
	//for _, pr := range factories.Exporters {
	//	exporters = append(exporters, pr)
	//}
	//factories.Exporters, err = component.MakeExporterFactoryMap(exporters...)
	//if err != nil {
	//	return component.Factories{}, err
	//}

	return factories, err
}
