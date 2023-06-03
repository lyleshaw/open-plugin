// Code generated by hertz generator.

package service

import (
	"context"
	"encoding/json"
	"github.com/lyleshaw/open-plugin/biz/dal/mysql"
	"github.com/lyleshaw/open-plugin/biz/model/orm_gen"
	"github.com/lyleshaw/open-plugin/pkg/constants"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/lyleshaw/open-plugin/biz/model/service"
	"github.com/lyleshaw/open-plugin/pkg/utils"
)

// GetPlugins .
// @router /api/plugins [GET]
func GetPlugins(ctx context.Context, c *app.RequestContext) {
	var err error
	var req service.GetPluginsReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	plugins, err := mysql.BatchGetPlugins(ctx, []string{})
	if err != nil {
		c.String(400, err.Error())
		return
	}

	pluginsResp := make([]*service.Plugin, len(plugins))
	for i, plugin := range plugins {
		pluginsResp[i] = &service.Plugin{
			PluginID:         plugin.PluginID,
			PluginName:       plugin.PluginName,
			PluginConfigURL:  plugin.PluginConfigURL,
			PluginOpenapiURL: plugin.PluginOpenapiURL,
			PluginConfig:     plugin.PluginConfig,
			PluginOpenapi:    plugin.PluginOpenapi,
			CreatedAt:        plugin.CreatedAt.Format(constants.TimeFormat),
			UpdatedAt:        plugin.UpdatedAt.Format(constants.TimeFormat),
		}
	}

	resp := service.PluginsResp{
		Code:    0,
		Message: "Success",
		Data:    pluginsResp,
	}

	c.JSON(200, resp)
}

// GetPlugin .
// @router /api/plugin [GET]
func GetPlugin(ctx context.Context, c *app.RequestContext) {
	var err error
	var req service.GetPluginReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	plugin, err := mysql.GetPluginByID(ctx, int(req.PluginID))
	if err != nil {
		c.String(400, err.Error())
		return
	}

	resp := service.PluginResp{
		Code:    0,
		Message: "Success",
		Data: &service.Plugin{
			PluginID:         plugin.PluginID,
			PluginName:       plugin.PluginName,
			PluginConfigURL:  plugin.PluginConfigURL,
			PluginOpenapiURL: plugin.PluginOpenapiURL,
			PluginConfig:     plugin.PluginConfig,
			PluginOpenapi:    plugin.PluginOpenapi,
			IsDeleted:        plugin.IsDeleted,
			CreatedAt:        plugin.CreatedAt.Format(constants.TimeFormat),
			UpdatedAt:        plugin.UpdatedAt.Format(constants.TimeFormat),
		},
	}

	c.JSON(200, resp)
}

// PostPlugin .
// @router /api/plugin [POST]
func PostPlugin(ctx context.Context, c *app.RequestContext) {
	var err error
	var req service.PluginReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	pluginData, openapi, err := utils.FetchData(req.PluginURL)
	if err != nil {
		c.String(400, err.Error())
	}

	plugin := utils.PluginDataAndDocToPlugin(pluginData, openapi)
	pluginJson, err := json.Marshal(plugin)
	if err != nil {
		c.String(400, err.Error())
	}

	openapiJson, err := json.Marshal(openapi)
	if err != nil {
		c.String(400, err.Error())
	}

	modelPlugin := orm_gen.Plugin{
		PluginName:       pluginData.NameForModel,
		PluginConfigURL:  req.PluginURL,
		PluginOpenapiURL: pluginData.API.URL,
		PluginConfig:     string(pluginJson),
		PluginOpenapi:    string(openapiJson),
	}

	createOrUpdatePlugin, err := mysql.CreateOrUpdatePlugin(ctx, &modelPlugin)
	if err != nil {
		return
	}

	resp := service.PluginResp{
		Code:    0,
		Message: "Success",
		Data: &service.Plugin{
			PluginID:         createOrUpdatePlugin.PluginID,
			PluginName:       createOrUpdatePlugin.PluginName,
			PluginConfigURL:  createOrUpdatePlugin.PluginConfigURL,
			PluginOpenapiURL: createOrUpdatePlugin.PluginOpenapiURL,
			PluginConfig:     createOrUpdatePlugin.PluginConfig,
			PluginOpenapi:    createOrUpdatePlugin.PluginOpenapi,
			IsDeleted:        createOrUpdatePlugin.IsDeleted,
			CreatedAt:        createOrUpdatePlugin.CreatedAt.Format(constants.TimeFormat),
			UpdatedAt:        createOrUpdatePlugin.UpdatedAt.Format(constants.TimeFormat),
		},
	}

	c.JSON(200, resp)
}

// DeletePlugin .
// @router /api/plugin [DELETE]
func DeletePlugin(ctx context.Context, c *app.RequestContext) {
	var err error
	var req service.DeletePluginReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	err = mysql.SoftDeletePluginByID(ctx, int(req.PluginID))

	resp := service.SuccessResp{
		Code:    0,
		Message: "Success",
	}

	c.JSON(200, resp)
}
