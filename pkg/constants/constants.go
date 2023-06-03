package constants

// TimeFormat Constants for time format
const (
	TimeFormat = "2006-01-02T15:04:05Z" // RFC3339 时间格式
)

const SelectPlugin = `
you are a professional bot that can detect the intention of a sentence.
user's sentence is "{{userPrompt}}"
here are {{pluginNumber}} plugins following:
{{pluginDesc}}
or:
- None
this sentence has no intention to use any other plugins
please only response {{pluginNames}}, 'None' which exactly match the user's intention.
`

const DetectParam = `
user's sentence is "{{userPrompt}}"
description: {{pluginDesc}}
the api schemas are following:
{{pluginAPIDetails}}
here is a format you should follow
DO NOT CHANGE THE FORMAT ANT THE URL SCHEMAS OFFERED
---
{"url": "/todo", "method": "POST", parameters:[{ "key": "username", "value": "小明" },  { "key": "todo", "value": "17 点提醒我看书" }]}
---
`

const MixResponse = `
you are a professional bot that good at analyzing documents and answer user's questions.
here are a response from {{pluginName}} plugin.
response is {{apiResponse}}
user's sentence is {{userPrompt}}
`

type ParameterResponse struct {
	URL        string `json:"url"`
	Method     string `json:"method"`
	Parameters []struct {
		Key   string      `json:"key"`
		Value interface{} `json:"value"`
	} `json:"parameters"`
}

type Parameter struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

type API struct {
	URL         string      `json:"url"`
	Method      string      `json:"method"`
	Description string      `json:"description"`
	Requires    []Parameter `json:"requires"`
	Returns     string      `json:"returns"`
}

type Plugin struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Apis        []*API `json:"apis"`
	BaseURL     string `json:"base_url"`
}

type PluginData struct {
	SchemaVersion       string `json:"schema_version"`
	NameForHuman        string `json:"name_for_human"`
	NameForModel        string `json:"name_for_model"`
	DescriptionForHuman string `json:"description_for_human"`
	DescriptionForModel string `json:"description_for_model"`
	Auth                struct {
		Type                     string `json:"type"`
		ClientURL                string `json:"client_url"`
		Scope                    string `json:"scope"`
		AuthorizationURL         string `json:"authorization_url"`
		AuthorizationContentType string `json:"authorization_content_type"`
		VerificationTokens       struct {
			Openai string `json:"openai"`
		} `json:"verification_tokens"`
	} `json:"auth"`
	API struct {
		Type string `json:"type"`
		URL  string `json:"url"`
	} `json:"api"`
	LogoURL      string `json:"logo_url"`
	ContactEmail string `json:"contact_email"`
	LegalInfoURL string `json:"legal_info_url"`
}
