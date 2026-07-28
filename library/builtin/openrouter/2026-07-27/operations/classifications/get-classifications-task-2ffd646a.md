---
title: Task classification market share
page_id: operation-get-classifications-task-94855cfb
path: operations/classifications
description: |-
    Returns the market-share breakdown of OpenRouter traffic by task classification
    (e.g. code generation, web search, summarization) over a trailing time window.

    Each classification reports its share of classified sampled requests (`usage_share`)
    and classified sampled token volume (`token_share`) as fractions between 0 and 1.
    The unclassified `other` bucket is excluded. Absolute volumes are not exposed
    because the underlying data is sampled.

    Each classification also includes a `models` array listing the top models by
    request volume within that classification, with their within-tag usage and token shares.

    Classifications are grouped into macro-categories (Code, Data, Agent, General)
    with aggregate shares provided for each.

    Authenticate with any valid OpenRouter API key (same key used for inference).
    Rate-limited to 30 requests/minute per key and 500 requests/day per account.

    When republishing or quoting this data, cite as:
    "Source: OpenRouter (openrouter.ai/rankings), as of {as_of}."
source: https://openrouter.ai/openapi.json
http_methods:
    - GET
api_endpoints:
    - /classifications/task
operation_ids:
    - getTaskClassifications
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# Task classification market share

`GET /classifications/task`

Operation ID: `getTaskClassifications`

Returns the market-share breakdown of OpenRouter traffic by task classification
(e.g. code generation, web search, summarization) over a trailing time window.

Each classification reports its share of classified sampled requests (`usage_share`)
and classified sampled token volume (`token_share`) as fractions between 0 and 1.
The unclassified `other` bucket is excluded. Absolute volumes are not exposed
because the underlying data is sampled.

Each classification also includes a `models` array listing the top models by
request volume within that classification, with their within-tag usage and token shares.

Classifications are grouped into macro-categories (Code, Data, Agent, General)
with aggregate shares provided for each.

Authenticate with any valid OpenRouter API key (same key used for inference).
Rate-limited to 30 requests/minute per key and 500 requests/day per account.

When republishing or quoting this data, cite as:
"Source: OpenRouter (openrouter.ai/rankings), as of {as_of}."

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"description": "Returns the market-share breakdown of OpenRouter traffic by task classification\n(e.g. code generation, web search, summarization) over a trailing time window.\n\nEach classification reports its share of classified sampled requests (`usage_share`)\nand classified sampled token volume (`token_share`) as fractions between 0 and 1.\nThe unclassified `other` bucket is excluded. Absolute volumes are not exposed\nbecause the underlying data is sampled.\n\nEach classification also includes a `models` array listing the top models by\nrequest volume within that classification, with their within-tag usage and token shares.\n\nClassifications are grouped into macro-categories (Code, Data, Agent, General)\nwith aggregate shares provided for each.\n\nAuthenticate with any valid OpenRouter API key (same key used for inference).\nRate-limited to 30 requests/minute per key and 500 requests/day per account.\n\nWhen republishing or quoting this data, cite as:\n\"Source: OpenRouter (openrouter.ai/rankings), as of {as_of}.\"", "operationId": "getTaskClassifications", "parameters": [{"description": "Trailing time window for the classification data. Currently only `7d` (trailing 7 days) is supported.", "in": "query", "name": "window", "required": false, "schema": {"default": "7d", "description": "Trailing time window for the classification data. Currently only `7d` (trailing 7 days) is supported.", "enum": ["7d"], "example": "7d", "type": "string"}}], "responses": {"200": {"content": {"application/json": {"example": {"data": {"as_of": "2026-06-17", "classifications": [{"category_token_share": 0.48, "category_usage_share": 0.51, "display_name": "Code Generation", "macro_category": "code", "models": [{"id": "openai/gpt-4.1-mini", "tag_token_share": 0.75, "tag_usage_share": 0.55}], "tag": "code:general_impl", "token_share": 0.31, "usage_share": 0.23}], "macro_categories": [{"key": "code", "label": "Code", "token_share": 0.52, "usage_share": 0.45}], "window_days": 7}}, "schema": {"$ref": "#/components/schemas/TaskClassificationResponse"}}}, "description": "Task classification market-share data for the requested trailing window."}, "400": {"content": {"application/json": {"example": {"error": {"code": 400, "message": "Invalid request parameters"}}, "schema": {"$ref": "#/components/schemas/BadRequestResponse"}}}, "description": "Bad Request - Invalid request parameters or malformed input"}, "401": {"content": {"application/json": {"example": {"error": {"code": 401, "message": "Missing Authentication header"}}, "schema": {"$ref": "#/components/schemas/UnauthorizedResponse"}}}, "description": "Unauthorized - Authentication required or invalid credentials"}, "429": {"content": {"application/json": {"example": {"error": {"code": 429, "message": "Rate limit exceeded"}}, "schema": {"$ref": "#/components/schemas/TooManyRequestsResponse"}}}, "description": "Too Many Requests - Rate limit exceeded"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}}, "summary": "Task classification market share", "tags": ["Classifications"]}
```
