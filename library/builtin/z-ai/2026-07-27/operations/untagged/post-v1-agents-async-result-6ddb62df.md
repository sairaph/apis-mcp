---
title: POST /v1/agents/async-result
page_id: operation-post-v1-agents-async-result-17f9323a
path: operations/untagged
description: This endpoint is used to query the result of an asynchronous request.
source: https://docs.z.ai/openapi.json
http_methods:
    - POST
api_endpoints:
    - /v1/agents/async-result
source_type: openapi
imported_from: https://docs.z.ai/openapi.json
---

# POST /v1/agents/async-result

`POST /v1/agents/async-result`

This endpoint is used to query the result of an asynchronous request.

## Definition

```yaml
{"description": "This endpoint is used to query the result of an asynchronous request.", "parameters": [{"$ref": "#/components/parameters/AcceptLanguage"}], "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/CommonAgentResultRequest"}}}, "required": true}, "responses": {"200": {"description": "Processing successful", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/CommonAgentResultResponse"}}}}, "default": {"description": "The request has failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/SpecialEffectsVideosAgentError"}}}}}}
```
