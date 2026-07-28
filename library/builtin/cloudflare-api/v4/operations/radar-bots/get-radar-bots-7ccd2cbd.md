---
title: List bots
page_id: operation-get-radar-bots-a2487976
path: operations/radar-bots
description: Retrieves a list of bots.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /radar/bots
operation_ids:
    - radar-get-bots
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List bots

`GET /radar/bots`

Operation ID: `radar-get-bots`

Retrieves a list of bots.

## Definition

```yaml
{"operationId": "radar-get-bots", "summary": "List bots", "description": "Retrieves a list of bots.", "parameters": [{"name": "limit", "in": "query", "description": "Limits the number of objects returned in the response.", "schema": {"description": "Limits the number of objects returned in the response.", "type": "integer", "example": 5, "default": 5, "exclusiveMinimum": true, "minimum": 0}}, {"name": "offset", "in": "query", "description": "Skips the specified number of objects before fetching the results.", "schema": {"description": "Skips the specified number of objects before fetching the results.", "type": "integer", "minimum": 0}}, {"name": "botCategory", "in": "query", "description": "Filters results by bot category.", "schema": {"description": "Filters results by bot category.", "type": "string", "enum": ["SEARCH_ENGINE_CRAWLER", "SEARCH_ENGINE_OPTIMIZATION", "MONITORING_AND_ANALYTICS", "ADVERTISING_AND_MARKETING", "SOCIAL_MEDIA_MARKETING", "PAGE_PREVIEW", "ACADEMIC_RESEARCH", "SECURITY", "ACCESSIBILITY", "WEBHOOKS", "FEED_FETCHER", "AI_CRAWLER", "AGGREGATOR", "AI_ASSISTANT", "AI_SEARCH", "ARCHIVER"]}}, {"name": "botOperator", "in": "query", "description": "Filters results by bot operator.", "schema": {"description": "Filters results by bot operator.", "type": "string", "maxLength": 100}}, {"name": "kind", "in": "query", "description": "Filters results by bot kind. Deprecated: the Verified Bot / Signed Agent distinction is being removed.", "schema": {"description": "Filters results by bot kind. Deprecated: the Verified Bot / Signed Agent distinction is being removed.", "type": "string", "enum": ["AGENT", "BOT"]}, "deprecated": true}, {"name": "botVerificationStatus", "in": "query", "description": "Filters results by bot verification status.", "schema": {"description": "Filters results by bot verification status.", "type": "string", "enum": ["VERIFIED"]}}, {"name": "format", "in": "query", "description": "Format in which results will be returned.", "schema": {"description": "Format in which results will be returned.", "type": "string", "example": "json", "enum": ["JSON", "CSV"]}}], "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"bots": {"type": "array", "items": {"properties": {"category": {"description": "The category of the bot.", "type": "string", "example": "AI_CRAWLER"}, "description": {"description": "A summary for the bot (e.g., purpose).", "type": "string", "example": "OpenAI/ChatGPT's web crawler"}, "kind": {"description": "The kind of the bot.", "type": "string", "example": "AGENT", "deprecated": true}, "name": {"description": "The name of the bot.", "type": "string", "example": "GPTBot"}, "operator": {"description": "The organization that owns and operates the bot.", "type": "string", "example": "OpenAI"}, "slug": {"description": "A kebab-case identifier derived from the bot name.", "type": "string", "example": "gptbot"}, "userAgentPatterns": {"type": "array", "items": {"description": "User agent patterns that identify the bot in web traffic.", "example": "GPTBot", "type": "string"}}}, "required": ["slug", "name", "description", "kind", "operator", "category", "userAgentPatterns"], "type": "object"}}}, "required": ["bots"]}, "success": {"type": "boolean", "example": true}}, "required": ["result", "success"]}}}}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean", "example": false}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Radar Bots"], "x-api-token-group": ["User Details Write", "User Details Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "radar.bots", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
