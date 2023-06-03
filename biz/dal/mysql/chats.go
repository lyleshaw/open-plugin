package mysql

import (
	"context"
	"github.com/bytedance/gopkg/util/logger"
	models "github.com/lyleshaw/open-plugin/biz/model/orm_gen"
	"github.com/lyleshaw/open-plugin/biz/model/query"
)

// GetChatByID Get Chat By chat_id
func GetChatByID(ctx context.Context, ChatID int32) (*models.Chat, error) {
	logger.Infof("Getting chat with ChatID: %d", ChatID)
	chat, err := query.Chat.WithContext(ctx).Where(query.Chat.ChatID.Eq(ChatID)).First()
	if err != nil {
		logger.Errorf("Failed to get chat with ChatID %d: %v", ChatID, err)
	}
	return chat, err
}

// GetChatByConversationID Get Chat By conversation_id
func GetChatByConversationID(ctx context.Context, ConversationID string) (*models.Chat, error) {
	logger.Infof("Getting chat with ConversationID: %d", ConversationID)
	chat, err := query.Chat.WithContext(ctx).Where(query.Chat.ConversationID.Eq(ConversationID)).First()
	if err != nil {
		logger.Errorf("Failed to get chat with ConversationID %d: %v", ConversationID, err)
	}
	return chat, err
}

// CreateChat Create Chat
func CreateChat(ctx context.Context, chat *models.Chat) (*models.Chat, error) {
	var err error
	err = query.Chat.WithContext(ctx).Create(chat)
	if err != nil {
		logger.Errorf("Failed to create chat: %v", err)
	}
	chat, err = GetChatByID(ctx, chat.ChatID)
	if err != nil {
		logger.Errorf("Failed to get chat with ID %d: %v in CreateChat", chat.ChatID, err)
		return nil, err
	}
	return chat, err
}

// UpdateChatContent Update Chat Content
func UpdateChatContent(ctx context.Context, chatID int, newContent string) (*models.Chat, error) {
	logger.Infof("Updating chat content with ID: %d", chatID)
	chat, err := GetChatByID(ctx, int32(chatID))
	if err != nil {
		return nil, err
	}
	chat.ChatContent = newContent
	err = query.Chat.WithContext(ctx).Save(chat)
	if err != nil {
		logger.Errorf("Failed to update chat content with ID %d: %v", chatID, err)
	}
	return chat, err
}
