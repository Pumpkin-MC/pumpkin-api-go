package api

import (
	"github.com/Pumpkin-MC/pumpkin-api-go/pkg/pumpkin_plugin_command"
	"github.com/Pumpkin-MC/pumpkin-api-go/pkg/pumpkin_plugin_context"
	"github.com/Pumpkin-MC/pumpkin-api-go/pkg/pumpkin_plugin_event"
	"github.com/Pumpkin-MC/pumpkin-api-go/pkg/pumpkin_plugin_metadata"
	"github.com/Pumpkin-MC/pumpkin-api-go/pkg/pumpkin_plugin_server"
)

type Metadata = pumpkin_plugin_metadata.PluginMetadata

type Plugin interface {
	Metadata() Metadata
	OnLoad(ctx *pumpkin_plugin_context.Context)
	OnUnload(ctx *pumpkin_plugin_context.Context)
	HandleEvent(eventId uint32, server *pumpkin_plugin_server.Server, event pumpkin_plugin_event.Event) pumpkin_plugin_event.Event
	HandleCommand(commandId uint32, sender *pumpkin_plugin_command.CommandSender, server *pumpkin_plugin_server.Server, args *pumpkin_plugin_command.ConsumedArgs) int32
}

var instance Plugin

func RegisterPlugin(p Plugin) {
	instance = p
}

func GetInstance() Plugin {
	return instance
}

type DefaultPlugin struct{}

func (p *DefaultPlugin) OnLoad(ctx *pumpkin_plugin_context.Context)   {}
func (p *DefaultPlugin) OnUnload(ctx *pumpkin_plugin_context.Context) {}
func (p *DefaultPlugin) HandleEvent(eventId uint32, server *pumpkin_plugin_server.Server, event pumpkin_plugin_event.Event) pumpkin_plugin_event.Event {
	return event
}
func (p *DefaultPlugin) HandleCommand(commandId uint32, sender *pumpkin_plugin_command.CommandSender, server *pumpkin_plugin_server.Server, args *pumpkin_plugin_command.ConsumedArgs) int32 {
	return 0
}
