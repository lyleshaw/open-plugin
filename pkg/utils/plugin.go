package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/lyleshaw/open-plugin/pkg/constants"
	"github.com/lyleshaw/open-plugin/pkg/llm_anthropic"
	"github.com/lyleshaw/open-plugin/pkg/llm_openai"
	"github.com/madebywelch/anthropic-go/pkg/anthropic"
	"github.com/sashabaranov/go-openai"
	"github.com/valyala/fasttemplate"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func PluginToAPIDetails(p *constants.Plugin) string {
	desc := ""
	for _, api := range p.Apis {
		desc += fmt.Sprintf("api: %s|%s\n", api.URL, api.Method)
		desc += fmt.Sprintf("desc: %s\n", api.Description)
		for _, r := range api.Requires {
			desc += fmt.Sprintf("parameter: %s:%s:%s\n", r.Name, r.Type, r.Description)
		}
	}
	return desc
}

func PluginToDescription(p *constants.Plugin) string {
	desc := fmt.Sprintf("-%s\n", p.Name)
	desc += fmt.Sprintf("%s\n", p.Description)
	return desc
}

func APIToResponse(baseURL string, parameterResponse *constants.ParameterResponse) (responseString string, err error) {
	// Send HTTP request
	var resp *http.Response
	if parameterResponse.Method == "GET" {
		v := url.Values{}
		for _, value := range parameterResponse.Parameters {
			switch value.Value.(type) {
			case string:
				v.Set(value.Key, value.Value.(string))
			case int:
				v.Set(value.Key, strconv.Itoa(value.Value.(int)))
			case float64:
				v.Set(value.Key, strconv.FormatFloat(value.Value.(float64), 'f', -1, 64))
			case bool:
				v.Set(value.Key, strconv.FormatBool(value.Value.(bool)))
			default:
				v.Set(value.Key, value.Value.(string))
			}
		}
		fmt.Printf(baseURL + parameterResponse.URL + "?" + v.Encode())

		resp, err = http.Get(baseURL + parameterResponse.URL + "?" + v.Encode())
		if err != nil {
			return "", err

		}
	} else if parameterResponse.Method == "POST" {
		var parameters map[string]interface{}
		parameters = make(map[string]interface{}, len(parameterResponse.Parameters))
		for _, value := range parameterResponse.Parameters {
			fmt.Printf("parameter: %v\n", value)
			fmt.Printf("parameter: %v\n", value.Key)
			fmt.Printf("parameter: %v\n", value.Value)
			parameters[value.Key] = value.Value
		}
		jsonParameters, err := json.Marshal(parameters)
		if err != nil {
			return "", err
		}
		resp, err = http.Post(baseURL+parameterResponse.URL, "application/json", bytes.NewBuffer(jsonParameters))
		if err != nil {
			return "", err
		}
	} else {
		return "", errors.New("APIToResponse Wrong")
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {

		}
	}(resp.Body)

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Convert response body to string
	responseString = string(body)

	return responseString, nil
}

func FetchData(pluginURL string) (*constants.PluginData, *openapi3.T, error) {
	var pluginData constants.PluginData

	resp, err := http.Get(pluginURL)
	if err != nil {
		return &pluginData, nil, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &pluginData, nil, err
	}

	err = json.Unmarshal(body, &pluginData)
	if err != nil {
		return &pluginData, nil, err
	}

	// Check if API URL is a valid URL
	apiURL, err := url.Parse(pluginData.API.URL)
	if err != nil {
		return &pluginData, nil, err
	}

	// Parse API spec
	ctx := context.Background()
	loader := &openapi3.Loader{Context: ctx, IsExternalRefsAllowed: true}
	doc, _ := loader.LoadFromURI(apiURL)
	// Validate document
	err = doc.Validate(ctx)
	if err != nil {
		return &pluginData, nil, err
	}

	return &pluginData, doc, err
}

func PluginDataAndDocToPlugin(pluginData *constants.PluginData, doc *openapi3.T) *constants.Plugin {
	var plugin constants.Plugin
	plugin.Name = *&pluginData.NameForModel
	plugin.Description = *&pluginData.DescriptionForModel
	var apis []*constants.API
	for apiURL, path := range doc.Paths {
		if path.Get != nil {
			parameters := make([]constants.Parameter, len(path.Get.Parameters))
			for i, parameter := range path.Get.Parameters {
				parameters[i] = constants.Parameter{
					Name:        strings.ToLower(parameter.Value.Schema.Value.Title),
					Description: parameter.Value.Schema.Value.Description,
					Type:        parameter.Value.Schema.Value.Type,
				}
			}
			apis = append(apis, &constants.API{
				URL:         apiURL,
				Requires:    parameters,
				Description: path.Get.Description,
				Method:      "GET",
			})
		}

		if path.Post != nil {
			properties := path.Post.RequestBody.Value.Content["application/json"].Schema.Value.Properties
			parameters := make([]constants.Parameter, len(properties))
			i := 0
			for name, parameter := range properties {
				parameters[i] = constants.Parameter{
					Name:        name,
					Description: parameter.Value.Description,
					Type:        parameter.Value.Type,
				}
				i++
			}
			apis = append(apis, &constants.API{
				URL:         apiURL,
				Requires:    parameters,
				Description: path.Post.Description,
				Method:      "POST",
			})
			fmt.Printf("apis: %v\n", apis)
		}

		if path.Patch != nil {
			parameters := make([]constants.Parameter, len(path.Patch.Parameters))
			for i, parameter := range path.Get.Parameters {
				parameters[i] = constants.Parameter{
					Name:        parameter.Value.Schema.Value.Title,
					Description: parameter.Value.Schema.Value.Description,
					Type:        parameter.Value.Schema.Value.Type,
				}
			}
			apis = append(apis, &constants.API{
				URL:         apiURL,
				Requires:    parameters,
				Description: path.Get.Description,
				Method:      "PATCH",
			})
		}

		if path.Put != nil {
			parameters := make([]constants.Parameter, len(path.Put.Parameters))
			for i, parameter := range path.Get.Parameters {
				parameters[i] = constants.Parameter{
					Name:        parameter.Value.Schema.Value.Title,
					Description: parameter.Value.Schema.Value.Description,
					Type:        parameter.Value.Schema.Value.Type,
				}
			}
			apis = append(apis, &constants.API{
				URL:         apiURL,
				Requires:    parameters,
				Description: path.Get.Description,
				Method:      "PUT",
			})
		}

		if path.Delete != nil {
			parameters := make([]constants.Parameter, len(path.Delete.Parameters))
			for i, parameter := range path.Get.Parameters {
				parameters[i] = constants.Parameter{
					Name:        parameter.Value.Schema.Value.Title,
					Description: parameter.Value.Schema.Value.Description,
					Type:        parameter.Value.Schema.Value.Type,
				}
			}
			apis = append(apis, &constants.API{
				URL:         apiURL,
				Requires:    parameters,
				Description: path.Get.Description,
				Method:      "DELETE",
			})
		}
	}

	plugin.Apis = apis
	plugin.BaseURL = doc.Servers[0].URL
	return &plugin
}

func SelectPluginsByOpenAI(ctx context.Context, client *openai.Client, userPrompt string, plugins []*constants.Plugin) (string, error) {
	var pluginNames, pluginDesc string
	var pluginNameArray = make([]string, len(plugins)+1)
	pluginNameArray[0] = "None"
	pluginNumber := strconv.Itoa(len(plugins))
	for i, p := range plugins {
		pluginNames += "'" + p.Name + "',"
		pluginDesc += PluginToDescription(p)
		pluginNameArray[i+1] = p.Name
	}

	t := fasttemplate.New(constants.SelectPlugin, "{{", "}}")
	systemPrompt := t.ExecuteString(map[string]interface{}{
		"pluginNumber": pluginNumber,
		"pluginNames":  pluginNames,
		"pluginDesc":   pluginDesc,
		"userPrompt":   userPrompt,
	})

	fmt.Printf("systemPrompt: %s\n", systemPrompt)

	resp, err := llm_openai.CompletionWithoutSession(ctx, client, systemPrompt, 0)
	if err != nil {
		return "", err
	}

	for _, v := range pluginNameArray {
		if resp == v {
			return resp, nil
		}
	}

	return resp, errors.New("selectPlugins Wrong")
}

func DetectParamsByOpenAI(ctx context.Context, client *openai.Client, userPrompt string, plugin *constants.Plugin) (*constants.ParameterResponse, error) {
	var pluginAPIDetails string
	pluginAPIDetails = PluginToAPIDetails(plugin)
	t := fasttemplate.New(constants.DetectParam, "{{", "}}")
	systemPrompt := t.ExecuteString(map[string]interface{}{
		"pluginName":       plugin.Name,
		"pluginDesc":       plugin.Description,
		"pluginAPIDetails": pluginAPIDetails,
		"userPrompt":       userPrompt,
	})
	fmt.Printf("systemPrompt: %s\n", systemPrompt)
	resp, err := llm_openai.CompletionWithoutSession(ctx, client, systemPrompt, 0)
	if err != nil {
		return nil, err
	}
	fmt.Printf("resp: %s\n", resp)
	var parameterResponse *constants.ParameterResponse
	err = json.Unmarshal([]byte(resp), &parameterResponse)
	fmt.Printf("parameterResponse: %v\n", parameterResponse)
	if err != nil {
		return nil, err
	}
	return parameterResponse, nil
}

func MixResponsesByOpenAI(ctx context.Context, client *openai.Client, pluginResponse string, userPrompt string, plugin *constants.Plugin) (string, error) {
	t := fasttemplate.New(constants.MixResponse, "{{", "}}")
	systemPrompt := t.ExecuteString(map[string]interface{}{
		"pluginName":  plugin.Name,
		"apiResponse": pluginResponse,
		"userPrompt":  userPrompt,
	})
	fmt.Printf("systemPrompt from mix: %s\n", systemPrompt)
	resp, err := llm_openai.CompletionWithoutSession(ctx, client, systemPrompt, 0)
	if err != nil {
		return "", err
	}
	return resp, nil
}

func SelectPluginsByAnthropic(client *anthropic.Client, userPrompt string, plugins []*constants.Plugin) (string, error) {
	var pluginNames, pluginDesc string
	var pluginNameArray = make([]string, len(plugins)+1)
	pluginNameArray[0] = "None"
	pluginNumber := strconv.Itoa(len(plugins))
	for i, p := range plugins {
		pluginNames += "'" + p.Name + "',"
		pluginDesc += PluginToDescription(p)
		pluginNameArray[i+1] = p.Name
	}

	t := fasttemplate.New(constants.SelectPlugin, "{{", "}}")
	systemPrompt := t.ExecuteString(map[string]interface{}{
		"pluginNumber": pluginNumber,
		"pluginNames":  pluginNames,
		"pluginDesc":   pluginDesc,
		"userPrompt":   userPrompt,
	})

	fmt.Printf("systemPrompt: %s\n", systemPrompt)

	resp, err := llm_anthropic.CompletionWithoutSessionByClaude(client, systemPrompt)
	if err != nil {
		return "", err
	}

	for _, v := range pluginNameArray {
		if resp == v {
			return resp, nil
		}
	}

	return resp, errors.New("selectPlugins Wrong")
}

func DetectParamsByAnthropic(client *anthropic.Client, userPrompt string, plugin *constants.Plugin) (*constants.ParameterResponse, error) {
	var pluginAPIDetails string
	pluginAPIDetails = PluginToAPIDetails(plugin)
	t := fasttemplate.New(constants.DetectParam, "{{", "}}")
	systemPrompt := t.ExecuteString(map[string]interface{}{
		"pluginName":       plugin.Name,
		"pluginDesc":       plugin.Description,
		"pluginAPIDetails": pluginAPIDetails,
		"userPrompt":       userPrompt,
	})
	fmt.Printf("systemPrompt: %s\n", systemPrompt)
	resp, err := llm_anthropic.CompletionWithoutSessionByClaude(client, systemPrompt)
	if err != nil {
		return nil, err
	}
	var parameterResponse *constants.ParameterResponse
	err = json.Unmarshal([]byte(resp), &parameterResponse)
	fmt.Printf("parameterResponse: %v\n", parameterResponse)
	if err != nil {
		return nil, err
	}
	return parameterResponse, nil
}

func MixResponsesByAnthropic(client *anthropic.Client, pluginResponse string, userPrompt string, plugin *constants.Plugin) (string, error) {
	t := fasttemplate.New(constants.MixResponse, "{{", "}}")
	systemPrompt := t.ExecuteString(map[string]interface{}{
		"pluginName":  plugin.Name,
		"apiResponse": pluginResponse,
		"userPrompt":  userPrompt,
	})
	fmt.Printf("systemPrompt from mix: %s\n", systemPrompt)
	resp, err := llm_anthropic.CompletionWithoutSessionByClaude(client, systemPrompt)
	if err != nil {
		return "", err
	}
	return resp, nil
}

func main() {
	ctx := context.Background()
	config := openai.DefaultConfig(llm_openai.API_KEY)
	config.BaseURL = llm_openai.BASE_URL
	client := openai.NewClientWithConfig(config)
	pluginURL := "https://search.aireview.tech/.well-known/ai-plugin.json"
	prompt := "杭州今天天气怎么样？"
	pluginData, doc, err := FetchData(pluginURL)
	if err != nil {
		fmt.Println(err)
		return
	}
	plugin := PluginDataAndDocToPlugin(pluginData, doc)

	//selectedPlugin, err := SelectPluginsByOpenAI(ctx, client, prompt, []*constants.Plugin{plugin})
	//fmt.Printf("selectedPlugin: %+v\n", selectedPlugin)
	//if selectedPlugin == plugin.Name {
	//	params, err := DetectParamsByOpenAI(ctx, client, prompt, plugin)
	//	if err != nil {
	//		fmt.Printf("err: %+v\n", err)
	//		return
	//	}
	//	fmt.Printf("params: %+v\n", params)
	//	apiResponse, err := APIToResponse(plugin.BaseURL, params)
	//	if err != nil {
	//		fmt.Printf("err: %+v\n", err)
	//		return
	//	}
	//	fmt.Printf("apiResponse: %+v\n", apiResponse)
	//	response, err := MixResponsesByOpenAI(ctx, client, apiResponse, prompt, plugin)
	//	if err != nil {
	//		fmt.Printf("err: %+v\n", err)
	//		return
	//	}
	//	fmt.Printf("response: %+v\n", response)
	//}

	fmt.Printf("--------------------\n")

	c, _ := llm_anthropic.GetClaudeClient()
	selectedPlugin, err := SelectPluginsByAnthropic(c, prompt, []*constants.Plugin{plugin})
	fmt.Printf("selectedPlugin: %+v\n", selectedPlugin)
	if selectedPlugin == " "+plugin.Name {
		params, err := DetectParamsByOpenAI(ctx, client, prompt, plugin)
		if err != nil {
			fmt.Printf("err: %+v\n", err)
			return
		}
		fmt.Printf("params: %+v\n", params)
		apiResponse, err := APIToResponse(plugin.BaseURL, params)
		if err != nil {
			fmt.Printf("err: %+v\n", err)
			return
		}
		fmt.Printf("apiResponse: %+v\n", apiResponse)
		response, err := MixResponsesByAnthropic(c, apiResponse, prompt, plugin)
		if err != nil {
			fmt.Printf("err: %+v\n", err)
			return
		}
		fmt.Printf("response: %+v\n", response)
	}
}
