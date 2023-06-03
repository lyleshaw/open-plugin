namespace go service

struct Chat {
    1: i32 chat_id,
    2: string conversation_id,
    3: list<string> plugin_names,
    4: string chat_content,
    5: string chat_model,
    6: i32 user_id,
    7: bool is_deleted,
    8: string created_at,
    9: string updated_at
}

struct Plugin {
    1: i32 plugin_id,
    2: string plugin_name,
    3: string plugin_config_url,
    4: string plugin_openapi_url,
    5: string plugin_config,
    6: string plugin_openapi,
    7: bool is_deleted,
    8: string created_at,
    9: string updated_at
}

struct ChatResp {
    1: i32 code,
    2: string message,
    3: Chat data
}

struct SSE {
}

struct PluginsResp {
    1: i32 code,
    2: string message,
    3: list<Plugin> data
}

struct PluginResp {
    1: i32 code,
    2: string message,
    3: Plugin data
}

struct PluginReq {
    1: string plugin_url
}

service ChatService {
    ChatResp GetChat(1: GetChatReq req) (api.get="/api/chat");

    SSE PostChat(1: PostChatReq req) (api.post="/api/chat");
}

service PluginService {
    PluginsResp GetPlugins(1: GetPluginsReq req) (api.get="/api/plugins");

    PluginResp GetPlugin(1: GetPluginReq req) (api.get="/api/plugin");

    PluginResp PostPlugin(1: PluginReq req) (api.post="/api/plugin");

    SuccessResp DeletePlugin(1: DeletePluginReq req) (api.delete="/api/plugin");
}

struct GetChatReq {
    1: string conversation_id
}

struct PostChatReq {
    1: string conversation_id,
    2: list<string> plugin_names,
    3: string chat_model
    4: string prompt
}

struct GetPluginsReq {}

struct GetPluginReq {
    1: i32 plugin_id
}

struct DeletePluginReq {
    1: i32 plugin_id
}

struct SuccessResp {
    1: i32 code,
    2: string message
}
