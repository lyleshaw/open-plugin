-- Table Plugins
CREATE TABLE Plugins
(
    plugin_id          INT          NOT NULL AUTO_INCREMENT COMMENT 'Unique identifier for the plugin',
    plugin_name        VARCHAR(255) NOT NULL COMMENT 'Name of the plugin',
    plugin_config_url  VARCHAR(255) COMMENT 'URL for the plugin configuration',
    plugin_openapi_url VARCHAR(255) COMMENT 'URL for the plugin OpenAPI documentation',
    plugin_config      JSON COMMENT 'Configuration for the plugin',
    plugin_openapi     JSON COMMENT 'OpenAPI documentation for the plugin',
    is_deleted         BOOLEAN      NOT NULL DEFAULT FALSE COMMENT 'Flag indicating if the plugin has been deleted',
    created_at         TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Timestamp for when the plugin was created',
    updated_at         TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Timestamp for when the plugin was last updated',
    PRIMARY KEY (plugin_id)
);

-- Table Chats
CREATE TABLE Chats
(
    chat_id         INT         NOT NULL AUTO_INCREMENT COMMENT 'Unique identifier for the chat',
    conversation_id VARCHAR(36) NOT NULL UNIQUE COMMENT 'Unique identifier for the conversation that the chat belongs to',
    plugin_names    JSON        NOT NULL COMMENT 'Array of plugin names used in the chat',
    chat_content    TEXT COMMENT 'Content of the chat message',
    chat_model      JSON COMMENT 'Model for the chat message',
    user_id         INT         NOT NULL COMMENT 'Unique identifier for the user who sent the chat message',
    is_deleted      BOOLEAN     NOT NULL DEFAULT FALSE COMMENT 'Flag indicating if the chat message has been deleted',
    created_at      TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Timestamp for when the chat message was created',
    updated_at      TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Timestamp for when the chat message was last updated',
    PRIMARY KEY (chat_id)
);