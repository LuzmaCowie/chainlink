package llo

import (
	"context"
	"maps"

	relayllo "github.com/smartcontractkit/chainlink-relay/pkg/reportingplugins/llo"
	"github.com/smartcontractkit/chainlink-relay/pkg/services"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
)

type ChannelDefinitionCache interface {
	relayllo.ChannelDefinitionCache
	services.Service
}

var _ ChannelDefinitionCache = &channelDefinitionCache{}

type channelDefinitionCache struct {
	services.StateMachine

	lggr        logger.Logger
	definitions relayllo.ChannelDefinitions
}

func NewChannelDefinitionCache() ChannelDefinitionCache {
	return &channelDefinitionCache{}
}

// TODO: Needs a way to subscribe/unsubscribe to contracts

func (c *channelDefinitionCache) Start(ctx context.Context) error {
	// TODO: Initial load, then poll
	// TODO: needs to be populated asynchronously from onchain ConfigurationStore
	return nil
}

func (c *channelDefinitionCache) Close() error {
	// TODO
	return nil
}

func (c *channelDefinitionCache) HealthReport() map[string]error {
	report := map[string]error{c.Name(): c.Healthy()}
	return report
}

func (c *channelDefinitionCache) Name() string { return c.lggr.Name() }

func (c *channelDefinitionCache) Definitions() relayllo.ChannelDefinitions {
	c.StateMachine.RLock()
	defer c.StateMachine.RUnlock()
	return maps.Clone(c.definitions)
}
