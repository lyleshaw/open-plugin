package mysql

import (
	"context"
	"crypto/tls"
	"database/sql"
	mysql_ "github.com/go-sql-driver/mysql"
	"github.com/lyleshaw/open-plugin/biz/model/orm_gen"
	"github.com/lyleshaw/open-plugin/biz/model/query"
	"log"
	"os"
	"testing"
)

func clearChatDatabase() {
	_ = mysql_.RegisterTLSConfig("tidb", &tls.Config{
		MinVersion: tls.VersionTLS12,
		ServerName: "gateway01.us-east-1.prod.aws.tidbcloud.com",
	})
	testDB, err := sql.Open("mysql", os.Getenv("DSN_DEV"))

	_, err = testDB.Exec("TRUNCATE chats;")
	if err != nil {
		log.Fatal(err)
	}
}

func TestGetChatByID(t *testing.T) {
	Init()
	query.SetDefault(DB)
	ctx := context.Background()

	// create chat
	chat := &orm_gen.Chat{
		ChatContent: "Hello",
		PluginNames: "{}",
		ChatModel:   "{}",
	}
	err := query.Chat.WithContext(ctx).Create(chat)
	if err != nil {
		t.Fatalf("Failed to create chat: %v", err)
	}

	retrievedChat, err := GetChatByID(ctx, chat.ChatID)
	if err != nil {
		t.Fatalf("GetChatByConversationID() failed: %v", err)
	}
	if retrievedChat.ChatContent != chat.ChatContent {
		t.Errorf("Expected chat content %q, got %q", chat.ChatContent, retrievedChat.ChatContent)
	}
	clearChatDatabase()
}

func TestGetChatByConversationID(t *testing.T) {
	Init()
	query.SetDefault(DB)
	ctx := context.Background()

	// create chat
	chat := &orm_gen.Chat{
		ConversationID: "2aeb533e-0031-11ee-be56-0242ac120002",
		ChatContent:    "Hello",
		PluginNames:    "{}",
		ChatModel:      "{}",
	}
	err := query.Chat.WithContext(ctx).Create(chat)
	if err != nil {
		t.Fatalf("Failed to create chat: %v", err)
	}

	retrievedChat, err := GetChatByConversationID(ctx, chat.ConversationID)
	if err != nil {
		t.Fatalf("GetChatByConversationID() failed: %v", err)
	}
	if retrievedChat.ChatContent != chat.ChatContent {
		t.Errorf("Expected chat content %q, got %q", chat.ChatContent, retrievedChat.ChatContent)
	}
	clearChatDatabase()
}

func TestCreateChat(t *testing.T) {
	Init()
	query.SetDefault(DB)
	ctx := context.Background()

	chat := &orm_gen.Chat{
		ChatContent: "Hello",
		PluginNames: "{}",
		ChatModel:   "{}",
	}
	createdChat, err := CreateChat(ctx, chat)
	if err != nil {
		t.Fatalf("CreateChat() failed: %v", err)
	}
	if createdChat.ChatID == 0 {
		t.Error("Expected chat ID to be non-zero")
	}
	clearChatDatabase()
}

func TestUpdateChatContent(t *testing.T) {
	Init()
	query.SetDefault(DB)
	ctx := context.Background()

	// create chat
	chat := &orm_gen.Chat{
		ChatContent: "Hello",
		PluginNames: "{}",
		ChatModel:   "{}",
	}
	err := query.Chat.WithContext(ctx).Create(chat)
	if err != nil {
		t.Fatalf("Failed to create chat: %v", err)
	}

	updatedContent := "Updated content"
	updatedChat, err := UpdateChatContent(ctx, int(chat.ChatID), updatedContent)
	if err != nil {
		t.Fatalf("UpdateChatContent() failed: %v", err)
	}
	if updatedChat.ChatContent != updatedContent {
		t.Errorf("Expected chat content %q, got %q", updatedContent, updatedChat.ChatContent)
	}
	clearChatDatabase()
}
