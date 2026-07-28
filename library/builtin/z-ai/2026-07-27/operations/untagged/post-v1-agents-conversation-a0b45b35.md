---
title: POST /v1/agents/conversation
page_id: operation-post-v1-agents-conversation-701b9705
path: operations/untagged
description: This endpoint is used to query the agent conversation history.Only support slides_glm_agent
source: https://docs.z.ai/openapi.json
http_methods:
    - POST
api_endpoints:
    - /v1/agents/conversation
source_type: openapi
imported_from: https://docs.z.ai/openapi.json
---

# POST /v1/agents/conversation

`POST /v1/agents/conversation`

This endpoint is used to query the agent conversation history.Only support slides_glm_agent

## Definition

```yaml
{"description": "This endpoint is used to query the agent conversation history.Only support slides_glm_agent", "parameters": [{"$ref": "#/components/parameters/AcceptLanguage"}], "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/GlmSlideAgentConversationRequest"}}}, "required": true}, "responses": {"200": {"description": "Processing successful", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/GlmSlideAgentConversationResponse"}}}}, "default": {"description": "The request has failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/Error"}}}}}}
```
