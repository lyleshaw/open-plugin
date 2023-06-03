package mysql

import (
	"context"
	"github.com/bytedance/gopkg/util/logger"
	models "github.com/lyleshaw/open-plugin/biz/model/orm_gen"
	"github.com/lyleshaw/open-plugin/biz/model/query"
)

// BatchGetPlugins Batch Get Plugins
func BatchGetPlugins(ctx context.Context, pluginNames []string) ([]*models.Plugin, error) {
	logger.Info("Batch getting plugins")
	var plugins []*models.Plugin
	var err error
	if len(pluginNames) > 0 {
		plugins, err = query.Plugin.WithContext(ctx).Where(query.Plugin.PluginName.In(pluginNames...)).Where(query.Plugin.IsDeleted.Is(false)).Find()
		if err != nil {
			logger.Errorf("Failed to get plugins: %v", err)
		}
		return plugins, err
	}

	plugins, err = query.Plugin.WithContext(ctx).Where(query.Plugin.IsDeleted.Is(false)).Find()
	if err != nil {
		logger.Errorf("Failed to get plugins: %v", err)
	}
	return plugins, err
}

// GetPluginByID Get Plugin By plugin_id
func GetPluginByID(ctx context.Context, pluginID int) (*models.Plugin, error) {
	logger.Infof("Getting plugin with ID: %d", pluginID)
	plugin, err := query.Plugin.WithContext(ctx).Where(query.Plugin.PluginID.Eq(int32(pluginID))).Where(query.Plugin.IsDeleted.Is(false)).First()
	if err != nil {
		logger.Errorf("Failed to get plugin with ID %d: %v", pluginID, err)
	}
	return plugin, err
}

// CreateOrUpdatePlugin Create Or Update Plugin
func CreateOrUpdatePlugin(ctx context.Context, plugin *models.Plugin) (*models.Plugin, error) {
	var err error
	if plugin.PluginID == 0 { // create plugin
		err = query.Plugin.WithContext(ctx).Create(plugin)
		if err != nil {
			logger.Errorf("Failed to create plugin: %v", err)
		}
	} else { // update plugin
		err = query.Plugin.WithContext(ctx).Save(plugin)
		if err != nil {
			logger.Errorf("Failed to update plugin with ID %d: %v", plugin.PluginID, err)
		}
	}

	plugin, err = GetPluginByID(ctx, int(plugin.PluginID))
	if err != nil {
		logger.Errorf("Failed to get plugin with ID %d: %v in CreateOrUpdatePlugin", plugin.PluginID, err)
		return nil, err
	}
	return plugin, err
}

// SoftDeletePluginByID Soft Delete Plugin By plugin_id
func SoftDeletePluginByID(ctx context.Context, pluginID int) error {
	logger.Infof("Soft deleting plugin with ID: %d", pluginID)
	plugin, err := GetPluginByID(ctx, pluginID)
	if err != nil {
		return err
	}
	plugin.IsDeleted = true
	err = query.Plugin.WithContext(ctx).Save(plugin)
	if err != nil {
		logger.Errorf("Failed to soft delete plugin with ID %d: %v", pluginID, err)
	}
	return err
}
