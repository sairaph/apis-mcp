---
title: List BYOK provider credentials
page_id: operation-get-byok-c5d1f812
path: operations/byok
description: List the bring-your-own-key (BYOK) provider credentials for the authenticated entity's default workspace. Use the `workspace_id` query parameter to scope the result to a different workspace, or the `provider` query parameter to filter by upstream provider. [Management key](/docs/guides/overview/auth/management-api-keys) required.
source: https://openrouter.ai/openapi.json
http_methods:
    - GET
api_endpoints:
    - /byok
operation_ids:
    - listBYOKKeys
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# List BYOK provider credentials

`GET /byok`

Operation ID: `listBYOKKeys`

List the bring-your-own-key (BYOK) provider credentials for the authenticated entity's default workspace. Use the `workspace_id` query parameter to scope the result to a different workspace, or the `provider` query parameter to filter by upstream provider. [Management key](/docs/guides/overview/auth/management-api-keys) required.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"description": "List the bring-your-own-key (BYOK) provider credentials for the authenticated entity's default workspace. Use the `workspace_id` query parameter to scope the result to a different workspace, or the `provider` query parameter to filter by upstream provider. [Management key](/docs/guides/overview/auth/management-api-keys) required.", "operationId": "listBYOKKeys", "parameters": [{"description": "Number of records to skip for pagination", "in": "query", "name": "offset", "required": false, "schema": {"default": 0, "description": "Number of records to skip for pagination", "example": 0, "minimum": 0, "type": ["integer", "null"]}}, {"description": "Maximum number of records to return (max 100)", "in": "query", "name": "limit", "required": false, "schema": {"default": 50, "description": "Maximum number of records to return (max 100)", "example": 50, "maximum": 100, "minimum": 1, "type": "integer"}}, {"description": "Optional workspace ID to filter by. Defaults to the authenticated entity's default workspace.", "in": "query", "name": "workspace_id", "required": false, "schema": {"description": "Optional workspace ID to filter by. Defaults to the authenticated entity's default workspace.", "example": "550e8400-e29b-41d4-a716-446655440000", "format": "uuid", "type": "string"}}, {"description": "Optional provider slug to filter by (e.g. `openai`, `anthropic`, `amazon-bedrock`).", "in": "query", "name": "provider", "required": false, "schema": {"description": "Optional provider slug to filter by (e.g. `openai`, `anthropic`, `amazon-bedrock`).", "enum": ["ai21", "aion-labs", "akashml", "alibaba", "amazon-bedrock", "amazon-nova", "ambient", "anthropic", "arcee-ai", "atlas-cloud", "avian", "azure", "baidu", "baseten", "black-forest-labs", "byteplus", "cerebras", "chutes", "cirrascale", "clarifai", "cloudflare", "cohere", "coreweave", "crusoe", "darkbloom", "decart", "deepgram", "deepinfra", "deepseek", "dekallm", "digitalocean", "featherless", "fireworks", "fish-audio", "friendli", "gmicloud", "google-ai-studio", "google-vertex", "groq", "heygen", "inception", "inceptron", "inferact-vllm", "inference-net", "infermatic", "inflection", "io-net", "ionstream", "krea", "liquid", "mancer", "mara", "meta", "minimax", "mistral", "modal", "modelrun", "modular", "moonshotai", "morph", "ncompass", "nebius", "nex-agi", "nextbit", "novita", "nvidia", "open-inference", "openai", "parasail", "perceptron", "perplexity", "phala", "poolside", "quiver", "recraft", "reka", "relace", "runway", "sail-research", "sakana", "sakana-ai", "sambanova", "seed", "siliconflow", "sourceful", "stepfun", "streamlake", "switchpoint", "tencent", "tenstorrent", "together", "upstage", "venice", "wafer", "wandb", "wandb-legacy", "xai", "xiaomi", "z-ai"], "example": "openai", "type": "string", "x-speakeasy-unknown-values": "allow"}}], "responses": {"200": {"content": {"application/json": {"example": {"data": [{"allowed_api_key_hashes": null, "allowed_models": null, "allowed_user_ids": null, "created_at": "2025-08-24T10:30:00Z", "disabled": false, "id": "11111111-2222-3333-4444-555555555555", "is_fallback": false, "label": "sk-...AbCd", "name": "Production OpenAI Key", "provider": "openai", "sort_order": 0, "workspace_id": "550e8400-e29b-41d4-a716-446655440000"}], "total_count": 1}, "schema": {"$ref": "#/components/schemas/ListBYOKKeysResponse"}}}, "description": "List of BYOK credentials"}, "400": {"content": {"application/json": {"example": {"error": {"code": 400, "message": "Invalid request parameters"}}, "schema": {"$ref": "#/components/schemas/BadRequestResponse"}}}, "description": "Bad Request - Invalid request parameters or malformed input"}, "401": {"content": {"application/json": {"example": {"error": {"code": 401, "message": "Missing Authentication header"}}, "schema": {"$ref": "#/components/schemas/UnauthorizedResponse"}}}, "description": "Unauthorized - Authentication required or invalid credentials"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}}, "summary": "List BYOK provider credentials", "tags": ["BYOK"], "x-speakeasy-name-override": "list", "x-speakeasy-pagination": {"inputs": [{"in": "parameters", "name": "offset", "type": "offset"}, {"in": "parameters", "name": "limit", "type": "limit"}], "outputs": {"results": "$.data"}, "type": "offsetLimit"}}
```
