---
title: Execute AI Model (Generic)
page_id: operation-post-accounts-account-id-ai-run-9a82de98
path: operations/workers-ai
description: |-
    Execute an AI model by specifying the model name in the request body.

    This endpoint provides a generic interface for running AI models where the model name is part of the request payload rather than the URL path. It supports all AI Gateway features including caching, custom headers, and request options.

    Model-specific inputs available in [Cloudflare Docs](https://developers.cloudflare.com/workers-ai/models/).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/ai/run
operation_ids:
    - workers-ai-post-run-generic
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Execute AI Model (Generic)

`POST /accounts/{account_id}/ai/run`

Operation ID: `workers-ai-post-run-generic`

Execute an AI model by specifying the model name in the request body.

This endpoint provides a generic interface for running AI models where the model name is part of the request payload rather than the URL path. It supports all AI Gateway features including caching, custom headers, and request options.

Model-specific inputs available in [Cloudflare Docs](https://developers.cloudflare.com/workers-ai/models/).

## Definition

```yaml
{"operationId": "workers-ai-post-run-generic", "summary": "Execute AI Model (Generic)", "description": "Execute an AI model by specifying the model name in the request body.\n\nThis endpoint provides a generic interface for running AI models where the model name is part of the request payload rather than the URL path. It supports all AI Gateway features including caching, custom headers, and request options.\n\nModel-specific inputs available in [Cloudflare Docs](https://developers.cloudflare.com/workers-ai/models/).", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "023e105f4ecef8ad9ca31a8372d0c353"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"input": {"description": "Model-specific input data. Format varies by model type.", "type": "object", "example": {"prompt": "Tell me a joke"}}, "model": {"description": "The AI model to execute (e.g., openai/gpt-5.5, anthropic/claude-opus-4.7)", "type": "string", "example": "openai/gpt-5.5"}, "options": {"type": "object", "properties": {"extraHeaders": {"description": "Additional headers to pass to the AI provider", "type": "object", "additionalProperties": {"type": "string"}}, "gateway": {"type": "object", "properties": {"cacheTtl": {"description": "Cache TTL in seconds", "type": "number", "example": 3600}, "id": {"description": "AI Gateway ID for caching and logging", "type": "string", "example": "my-gateway"}, "skipCache": {"description": "Skip cache lookup for this request", "type": "boolean"}}}}}}, "required": ["model", "input"]}}}}, "responses": {"200": {"description": "Model response", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"type": "object"}}, "messages": {"type": "array", "items": {"type": "object"}}, "result": {"description": "Model-specific output. Format varies by model type.", "type": "object"}, "success": {"type": "boolean"}}, "required": ["success", "result", "errors", "messages"]}}}}, "400": {"description": "Bad request - missing required fields or invalid input", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "success": {"type": "boolean"}}, "required": ["success", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Workers AI"], "x-api-token-group": ["Workers AI Write", "Workers AI Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.ai"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
