package mysql

import (
	"context"
	"crypto/tls"
	"database/sql"
	mysql_ "github.com/go-sql-driver/mysql"
	models "github.com/lyleshaw/open-plugin/biz/model/orm_gen"
	"github.com/lyleshaw/open-plugin/biz/model/query"
	"log"
	"os"
	"testing"
)

func clearPluginDatabase() {
	_ = mysql_.RegisterTLSConfig("tidb", &tls.Config{
		MinVersion: tls.VersionTLS12,
		ServerName: "gateway01.us-east-1.prod.aws.tidbcloud.com",
	})
	testDB, err := sql.Open("mysql", os.Getenv("DSN_DEV"))

	_, err = testDB.Exec("TRUNCATE plugins;")
	if err != nil {
		log.Fatal(err)
	}
}

func TestBatchGetPlugins(t *testing.T) {
	Init()
	query.SetDefault(DB)
	ctx := context.Background()

	// create plugins
	plugin1 := &models.Plugin{
		PluginName:    "Plugin 1",
		PluginConfig:  "{}",
		PluginOpenapi: "{}",
	}
	err := query.Plugin.WithContext(ctx).Create(plugin1)
	if err != nil {
		t.Fatalf("Failed to create plugin: %v", err)
	}
	plugin2 := &models.Plugin{
		PluginName:    "Plugin 2",
		PluginConfig:  "{}",
		PluginOpenapi: "{}",
	}
	err = query.Plugin.WithContext(ctx).Create(plugin2)
	if err != nil {
		t.Fatalf("Failed to create plugin: %v", err)
	}

	plugins, err := BatchGetPlugins(ctx, []string{})
	if err != nil {
		t.Fatalf("BatchGetPlugins() failed: %v", err)
	}

	if len(plugins) != 2 {
		t.Errorf("Expected 2 plugins, got %d", len(plugins))
	}
	if plugins[0].PluginName != plugin1.PluginName {
		t.Errorf("Expected plugin name %q, got %q", plugin1.PluginName, plugins[0].PluginName)
	}
	if plugins[1].PluginName != plugin2.PluginName {
		t.Errorf("Expected plugin name %q, got %q", plugin2.PluginName, plugins[1].PluginName)
	}

	pluginNames := []string{plugin1.PluginName}
	plugins, err = BatchGetPlugins(ctx, pluginNames)
	if err != nil {
		t.Fatalf("BatchGetPlugins() failed: %v", err)
	}

	if len(plugins) != 1 {
		t.Errorf("Expected 1 plugins, got %d", len(plugins))
	}
	if plugins[0].PluginName != plugin1.PluginName {
		t.Errorf("Expected plugin name %q, got %q", plugin1.PluginName, plugins[0].PluginName)
	}

	clearPluginDatabase()
}

func TestGetPluginByID(t *testing.T) {
	Init()
	query.SetDefault(DB)
	ctx := context.Background()

	// create plugin
	plugin := &models.Plugin{
		PluginName:    "Plugin",
		PluginConfig:  "{}",
		PluginOpenapi: "{}",
	}
	err := query.Plugin.WithContext(ctx).Create(plugin)
	if err != nil {
		t.Fatalf("Failed to create plugin: %v", err)
	}

	retrievedPlugin, err := GetPluginByID(ctx, int(plugin.PluginID))
	if err != nil {
		t.Fatalf("GetPluginByID() failed: %v", err)
	}

	if retrievedPlugin.PluginName != plugin.PluginName {
		t.Errorf("Expected plugin name %q, got %q", plugin.PluginName, retrievedPlugin.PluginName)
	}

	clearPluginDatabase()
}

func TestCreateOrUpdatePlugin(t *testing.T) {
	Init()
	query.SetDefault(DB)
	ctx := context.Background()

	// create plugin
	plugin := &models.Plugin{
		PluginName:    "Plugin",
		PluginConfig:  "{}",
		PluginOpenapi: "{}",
	}
	createdPlugin, err := CreateOrUpdatePlugin(ctx, plugin)
	if err != nil {
		t.Fatalf("CreateOrUpdatePlugin() failed: %v", err)
	}
	if createdPlugin.PluginID == 0 {
		t.Error("Expected plugin ID to be non-zero")
	}

	// update plugin
	updatedPluginName := "Updated Plugin"
	createdPlugin.PluginName = updatedPluginName
	updatedPlugin, err := CreateOrUpdatePlugin(ctx, createdPlugin)
	if err != nil {
		t.Fatalf("CreateOrUpdatePlugin() failed: %v", err)
	}
	if updatedPlugin.PluginName != updatedPluginName {
		t.Errorf("Expected plugin name %q, got %q", updatedPluginName, updatedPlugin.PluginName)
	}

	clearPluginDatabase()
}

func TestSoftDeletePluginByID(t *testing.T) {
	Init()
	query.SetDefault(DB)
	ctx := context.Background()

	// create plugin
	plugin := &models.Plugin{
		PluginName:    "Plugin",
		PluginConfig:  "{}",
		PluginOpenapi: "{}",
	}
	err := query.Plugin.WithContext(ctx).Create(plugin)
	if err != nil {
		t.Fatalf("Failed to create plugin: %v", err)
	}

	err = SoftDeletePluginByID(ctx, int(plugin.PluginID))
	if err != nil {
		t.Fatalf("SoftDeletePluginByID() failed: %v", err)
	}

	retrievedPlugin, err := query.Plugin.WithContext(ctx).Where(query.Plugin.PluginID.Eq(plugin.PluginID)).First()
	if err != nil {
		t.Fatalf("GetPluginByID() failed: %v", err)
	}
	if retrievedPlugin.IsDeleted != true {
		t.Error("Expected deleted plugin to have IsDeleted = true")
	}

	clearPluginDatabase()
}
