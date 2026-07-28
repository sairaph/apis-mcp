---
title: Get bot details
page_id: operation-get-radar-bots-bot-slug-55cac1ff
path: operations/radar-bots
description: Retrieves the requested bot information.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /radar/bots/{bot_slug}
operation_ids:
    - radar-get-bot-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get bot details

`GET /radar/bots/{bot_slug}`

Operation ID: `radar-get-bot-details`

Retrieves the requested bot information.

## Definition

```yaml
{"operationId": "radar-get-bot-details", "summary": "Get bot details", "description": "Retrieves the requested bot information.", "parameters": [{"name": "bot_slug", "in": "path", "description": "Bot slug.", "required": true, "schema": {"description": "Bot slug.", "type": "string", "example": "gptbot", "maxLength": 100}}, {"name": "format", "in": "query", "description": "Format in which results will be returned.", "schema": {"description": "Format in which results will be returned.", "type": "string", "example": "json", "enum": ["JSON", "CSV"]}}], "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"bot": {"type": "object", "properties": {"category": {"description": "The category of the bot.", "type": "string", "example": "AI_CRAWLER"}, "description": {"description": "A summary for the bot (e.g., purpose).", "type": "string", "example": "OpenAI/ChatGPT's web crawler"}, "kind": {"description": "The kind of the bot.", "type": "string", "example": "AGENT", "deprecated": true}, "name": {"description": "The name of the bot.", "type": "string", "example": "GPTBot"}, "operator": {"description": "The organization that owns and operates the bot.", "type": "string", "example": "OpenAI"}, "operatorUrl": {"description": "The link to the bot documentation.", "type": "string", "example": "https://platform.openai.com/docs/bots"}, "signatureAgentUrl": {"description": "The URL of the agent's [Web Bot Auth](https://blog.cloudflare.com/web-bot-auth/) resource. Null for bots not verified via request signature.", "type": "string", "example": "https://example.com/signature-agent", "nullable": true}, "slug": {"description": "A kebab-case identifier derived from the bot name.", "type": "string", "example": "gptbot"}, "userAgentPatterns": {"type": "array", "items": {"description": "User agent patterns that identify the bot in web traffic.", "example": "GPTBot", "type": "string"}}, "userAgents": {"type": "array", "items": {"description": "User agent strings used by this bot in HTTP requests.", "example": "GPTBot", "type": "string"}}}, "required": ["slug", "name", "description", "kind", "operator", "operatorUrl", "category", "userAgents", "userAgentPatterns"]}}, "required": ["bot"]}, "success": {"type": "boolean", "example": true}}, "required": ["result", "success"]}}}}, "404": {"description": "Not found.", "content": {"application/json": {"schema": {"type": "object", "properties": {"error": {"type": "string", "example": "Not Found."}}, "required": ["error"]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Radar Bots"], "x-api-token-group": ["User Details Write", "User Details Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "radar.bots", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
