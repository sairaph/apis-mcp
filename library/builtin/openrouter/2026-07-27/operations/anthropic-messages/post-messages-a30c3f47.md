---
title: Create a message
page_id: operation-post-messages-cfde5869
path: operations/anthropic-messages
description: Creates a message using the Anthropic Messages API format. Supports text, images, PDFs, tools, and extended thinking.
source: https://openrouter.ai/openapi.json
http_methods:
    - POST
api_endpoints:
    - /messages
operation_ids:
    - createMessages
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# Create a message

`POST /messages`

Operation ID: `createMessages`

Creates a message using the Anthropic Messages API format. Supports text, images, PDFs, tools, and extended thinking.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"description": "Creates a message using the Anthropic Messages API format. Supports text, images, PDFs, tools, and extended thinking.", "operationId": "createMessages", "parameters": [{"description": "Opt-in to surface routing metadata on the response under `openrouter_metadata`. Defaults to `disabled`. The legacy header `X-OpenRouter-Experimental-Metadata` is also accepted for backward compatibility.", "example": "enabled", "in": "header", "name": "X-OpenRouter-Metadata", "required": false, "schema": {"$ref": "#/components/schemas/MetadataLevel"}}], "requestBody": {"content": {"application/json": {"example": {"max_tokens": 1024, "messages": [{"content": "Hello, how are you?", "role": "user"}], "model": "anthropic/claude-sonnet-4"}, "schema": {"$ref": "#/components/schemas/MessagesRequest"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"example": {"container": null, "content": [{"citations": [], "text": "I'm doing well, thank you for asking! How can I help you today?", "type": "text"}], "id": "msg_abc123", "model": "anthropic/claude-sonnet-4", "role": "assistant", "stop_details": null, "stop_reason": "end_turn", "stop_sequence": null, "type": "message", "usage": {"cache_creation": null, "cache_creation_input_tokens": null, "cache_read_input_tokens": null, "inference_geo": null, "input_tokens": 12, "output_tokens": 18, "output_tokens_details": null, "server_tool_use": null, "service_tier": "standard"}}, "schema": {"$ref": "#/components/schemas/MessagesResult"}}, "text/event-stream": {"example": {"data": {"delta": {"text": "Hello", "type": "text_delta"}, "index": 0, "type": "content_block_delta"}, "event": "content_block_delta"}, "schema": {"$ref": "#/components/schemas/MessagesStreamingResponse"}, "x-speakeasy-sse-sentinel": "[DONE]"}}, "description": "Successful response"}, "400": {"content": {"application/json": {"example": {"error": {"message": "Invalid request: messages is required", "type": "invalid_request_error"}, "type": "error"}, "schema": {"$ref": "#/components/schemas/MessagesErrorResponse"}}}, "description": "Invalid request error"}, "401": {"content": {"application/json": {"example": {"error": {"message": "Invalid API key", "type": "authentication_error"}, "type": "error"}, "schema": {"$ref": "#/components/schemas/MessagesErrorResponse"}}}, "description": "Authentication error"}, "403": {"content": {"application/json": {"examples": {"guardrail-blocked": {"summary": "Guardrail blocked the request", "value": {"error": {"code": 403, "message": "Request blocked: prompt injection patterns detected", "metadata": {"patterns": ["ignore all previous instructions"]}}, "openrouter_metadata": {"attempt": 1, "endpoints": {"available": [{"model": "openai/gpt-4o", "provider": "OpenAI", "selected": false}], "total": 1}, "is_byok": false, "pipeline": [{"data": {"action": "blocked", "detected": true, "engines": ["regex"], "patterns": ["ignore all previous instructions"]}, "guardrail_id": "grd_abc123", "guardrail_scope": "api-key", "name": "regex_pi_detection", "summary": "Blocked: prompt injection detected (1 pattern matched)", "type": "guardrail"}], "region": "iad", "requested": "openai/gpt-4o", "strategy": "direct", "summary": "available=1"}}}, "insufficient-permissions": {"summary": "Insufficient permissions", "value": {"error": {"code": 403, "message": "Only management keys can perform this operation"}}}}, "schema": {"$ref": "#/components/schemas/ForbiddenResponse"}}}, "description": "Forbidden - Authentication successful but insufficient permissions, or a guardrail blocked the request. When guardrails block and the `X-OpenRouter-Metadata: enabled` header is present, the response includes `openrouter_metadata` with full routing context and a `pipeline` array containing guardrail stage details."}, "404": {"content": {"application/json": {"example": {"error": {"message": "Model not found", "type": "not_found_error"}, "type": "error"}, "schema": {"$ref": "#/components/schemas/MessagesErrorResponse"}}}, "description": "Not found error"}, "429": {"content": {"application/json": {"example": {"error": {"message": "Rate limit exceeded", "type": "rate_limit_error"}, "type": "error"}, "schema": {"$ref": "#/components/schemas/MessagesErrorResponse"}}}, "description": "Rate limit error"}, "500": {"content": {"application/json": {"example": {"error": {"message": "Internal server error", "type": "api_error"}, "type": "error"}, "schema": {"$ref": "#/components/schemas/MessagesErrorResponse"}}}, "description": "API error"}, "503": {"content": {"application/json": {"example": {"error": {"message": "Service temporarily overloaded", "type": "overloaded_error"}, "type": "error"}, "schema": {"$ref": "#/components/schemas/MessagesErrorResponse"}}}, "description": "Overloaded error"}, "529": {"content": {"application/json": {"example": {"error": {"message": "Provider is temporarily overloaded", "type": "overloaded_error"}, "type": "error"}, "schema": {"$ref": "#/components/schemas/MessagesErrorResponse"}}}, "description": "Overloaded error"}}, "summary": "Create a message", "tags": ["Anthropic Messages"], "x-speakeasy-ignore": true, "x-speakeasy-name-override": "create", "x-speakeasy-stream-request-field": "stream"}
```
