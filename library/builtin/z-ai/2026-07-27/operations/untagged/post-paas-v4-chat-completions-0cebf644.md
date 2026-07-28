---
title: POST /paas/v4/chat/completions
page_id: operation-post-paas-v4-chat-completions-a8fcab79
path: operations/untagged
description: Create a chat completion model that generates AI replies for given conversation messages. It supports multimodal inputs (text, images, audio, video, file), offers configurable parameters (like temperature, max tokens, tool use), and supports both streaming and non-streaming output modes.
source: https://docs.z.ai/openapi.json
http_methods:
    - POST
api_endpoints:
    - /paas/v4/chat/completions
source_type: openapi
imported_from: https://docs.z.ai/openapi.json
---

# POST /paas/v4/chat/completions

`POST /paas/v4/chat/completions`

Create a chat completion model that generates AI replies for given conversation messages. It supports multimodal inputs (text, images, audio, video, file), offers configurable parameters (like temperature, max tokens, tool use), and supports both streaming and non-streaming output modes.

## Definition

```yaml
{"description": "Create a chat completion model that generates AI replies for given conversation messages. It supports multimodal inputs (text, images, audio, video, file), offers configurable parameters (like temperature, max tokens, tool use), and supports both streaming and non-streaming output modes.", "parameters": [{"$ref": "#/components/parameters/AcceptLanguage"}], "requestBody": {"content": {"application/json": {"schema": {"oneOf": [{"title": "Text Model", "$ref": "#/components/schemas/ChatCompletionTextRequest"}, {"title": "Vision Model", "$ref": "#/components/schemas/ChatCompletionVisionRequest"}]}, "examples": {"Basic Example": {"value": {"model": "glm-5.2", "messages": [{"role": "system", "content": "You are a helpful coding assistant."}, {"role": "user", "content": "Write a Python function to calculate the factorial of a number."}], "temperature": 1, "stream": false}}, "Stream Example": {"value": {"model": "glm-5.2", "messages": [{"role": "system", "content": "You are a helpful coding assistant."}, {"role": "user", "content": "Write a Python function to calculate the factorial of a number."}], "temperature": 1, "stream": true}}, "Thinking Example": {"value": {"model": "glm-5.2", "messages": [{"role": "system", "content": "You are a helpful coding assistant."}, {"role": "user", "content": "Write a Python function to calculate the factorial of a number."}], "thinking": {"type": "enabled"}, "stream": true}}, "Multi Conversation": {"value": {"model": "glm-5.2", "messages": [{"role": "system", "content": "You are a professional programming assistant."}, {"role": "user", "content": "What is recursion?"}, {"role": "assistant", "content": "Recursion is a programming technique where a function calls itself to solve a problem... What is recursion"}, {"role": "user", "content": "Can you give me an example of Python recursion?"}], "stream": true}}, "Image Visual Example": {"value": {"model": "glm-5v-turbo", "messages": [{"role": "user", "content": [{"type": "image_url", "image_url": {"url": "https://cdn.bigmodel.cn/static/logo/register.png"}}, {"type": "image_url", "image_url": {"url": "https://cdn.bigmodel.cn/static/logo/api-key.png"}}, {"type": "text", "text": "What are the pics talk about?"}]}]}}, "Video Visual Example": {"value": {"model": "glm-5v-turbo", "messages": [{"role": "user", "content": [{"type": "video_url", "video_url": {"url": "https://cdn.bigmodel.cn/agent-demos/lark/113123.mov"}}, {"type": "text", "text": "What are the video show about?"}]}]}}, "File Visual Example": {"value": {"model": "glm-5v-turbo", "messages": [{"role": "user", "content": [{"type": "file_url", "file_url": {"url": "https://cdn.bigmodel.cn/static/demo/demo2.txt"}}, {"type": "file_url", "file_url": {"url": "https://cdn.bigmodel.cn/static/demo/demo1.pdf"}}, {"type": "text", "text": "What are the files show about?"}]}]}}, "Function Call Example": {"value": {"model": "glm-5.2", "messages": [{"role": "user", "content": "Is there an example of how the weather in Beijing is today?"}], "tools": [{"type": "function", "function": {"name": "get_weather", "description": "Get weather information for the specified city.", "parameters": {"type": "object", "properties": {"city": {"type": "string", "description": "City Name"}}, "required": ["city"]}}}], "tool_choice": "auto", "temperature": 0.3}}}}}, "required": true}, "responses": {"200": {"description": "Processing successful", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/ChatCompletionResponse"}}}}, "default": {"description": "The request has failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/Error"}}}}}}
```
