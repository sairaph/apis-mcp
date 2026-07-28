---
title: Top apps by token usage
page_id: operation-get-datasets-app-rankings-ad81bcbd
path: operations/datasets
description: |-
    Returns the top public apps on OpenRouter ranked by token usage inside the requested
    date window, matching the public apps marketplace on openrouter.ai/apps. Token totals
    are `prompt_tokens + completion_tokens`; hidden and private apps are excluded and
    traffic from related app aliases is merged into the canonical visible app.

    `sort=popular` (default) ranks by total token volume inside the window.
    `sort=trending` ranks by absolute excess token growth: window volume minus the average
    volume of the three equal-length periods immediately preceding the window. Apps with
    no excess growth are omitted, so `trending` may return fewer than `limit` rows.

    Filter with `category` (marketplace category group, e.g. `coding`) or `subcategory`
    (e.g. `cli-agent`). Ranks are re-numbered 1..N after filtering. Page with `offset` —
    `rank` stays absolute, so the first row of `offset=50` is `rank: 51`.

    Authenticate with any valid OpenRouter API key (same key used for inference).
    Rate-limited to 30 requests/minute per key and 500 requests/day per account.

    When republishing or quoting this dataset, OpenRouter must be cited as:
    "Source: OpenRouter (openrouter.ai/apps), as of {as_of}."

    Token counts come from each upstream provider's own tokenizer, so a token attributed
    to one app is not directly comparable to a token attributed to another app whose
    traffic flows through a different provider.
source: https://openrouter.ai/openapi.json
http_methods:
    - GET
api_endpoints:
    - /datasets/app-rankings
operation_ids:
    - getAppRankings
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# Top apps by token usage

`GET /datasets/app-rankings`

Operation ID: `getAppRankings`

Returns the top public apps on OpenRouter ranked by token usage inside the requested
date window, matching the public apps marketplace on openrouter.ai/apps. Token totals
are `prompt_tokens + completion_tokens`; hidden and private apps are excluded and
traffic from related app aliases is merged into the canonical visible app.

`sort=popular` (default) ranks by total token volume inside the window.
`sort=trending` ranks by absolute excess token growth: window volume minus the average
volume of the three equal-length periods immediately preceding the window. Apps with
no excess growth are omitted, so `trending` may return fewer than `limit` rows.

Filter with `category` (marketplace category group, e.g. `coding`) or `subcategory`
(e.g. `cli-agent`). Ranks are re-numbered 1..N after filtering. Page with `offset` —
`rank` stays absolute, so the first row of `offset=50` is `rank: 51`.

Authenticate with any valid OpenRouter API key (same key used for inference).
Rate-limited to 30 requests/minute per key and 500 requests/day per account.

When republishing or quoting this dataset, OpenRouter must be cited as:
"Source: OpenRouter (openrouter.ai/apps), as of {as_of}."

Token counts come from each upstream provider's own tokenizer, so a token attributed
to one app is not directly comparable to a token attributed to another app whose
traffic flows through a different provider.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"description": "Returns the top public apps on OpenRouter ranked by token usage inside the requested\ndate window, matching the public apps marketplace on openrouter.ai/apps. Token totals\nare `prompt_tokens + completion_tokens`; hidden and private apps are excluded and\ntraffic from related app aliases is merged into the canonical visible app.\n\n`sort=popular` (default) ranks by total token volume inside the window.\n`sort=trending` ranks by absolute excess token growth: window volume minus the average\nvolume of the three equal-length periods immediately preceding the window. Apps with\nno excess growth are omitted, so `trending` may return fewer than `limit` rows.\n\nFilter with `category` (marketplace category group, e.g. `coding`) or `subcategory`\n(e.g. `cli-agent`). Ranks are re-numbered 1..N after filtering. Page with `offset` —\n`rank` stays absolute, so the first row of `offset=50` is `rank: 51`.\n\nAuthenticate with any valid OpenRouter API key (same key used for inference).\nRate-limited to 30 requests/minute per key and 500 requests/day per account.\n\nWhen republishing or quoting this dataset, OpenRouter must be cited as:\n\"Source: OpenRouter (openrouter.ai/apps), as of {as_of}.\"\n\nToken counts come from each upstream provider's own tokenizer, so a token attributed\nto one app is not directly comparable to a token attributed to another app whose\ntraffic flows through a different provider.", "operationId": "getAppRankings", "parameters": [{"description": "Marketplace category group to filter by (e.g. `coding`). Only apps tagged with a subcategory inside this group are returned. Mutually combinable with `subcategory` — when both are supplied the `subcategory` must belong to the `category` group.", "in": "query", "name": "category", "required": false, "schema": {"description": "Marketplace category group to filter by (e.g. `coding`). Only apps tagged with a subcategory inside this group are returned. Mutually combinable with `subcategory` — when both are supplied the `subcategory` must belong to the `category` group.", "enum": ["coding", "creative", "productivity", "entertainment"], "example": "coding", "type": "string", "x-speakeasy-unknown-values": "allow"}}, {"description": "Marketplace subcategory to filter by (e.g. `cli-agent`). Takes precedence over `category` for the actual filter; when `category` is also supplied the pair must be consistent.", "in": "query", "name": "subcategory", "required": false, "schema": {"description": "Marketplace subcategory to filter by (e.g. `cli-agent`). Takes precedence over `category` for the actual filter; when `category` is also supplied the pair must be consistent.", "enum": ["cli-agent", "ide-extension", "cloud-agent", "programming-app", "native-app-builder", "creative-writing", "video-gen", "image-gen", "audio-gen", "roleplay", "game", "writing-assistant", "general-chat", "personal-agent", "legal"], "example": "cli-agent", "type": "string", "x-speakeasy-unknown-values": "allow"}}, {"description": "`popular` ranks apps by total token volume inside the date window. `trending` ranks apps by absolute excess token growth: window volume minus the average volume of the three equal-length periods immediately preceding the window. Apps with no excess growth are omitted from `trending` results.", "in": "query", "name": "sort", "required": false, "schema": {"default": "popular", "description": "`popular` ranks apps by total token volume inside the date window. `trending` ranks apps by absolute excess token growth: window volume minus the average volume of the three equal-length periods immediately preceding the window. Apps with no excess growth are omitted from `trending` results.", "enum": ["popular", "trending"], "example": "popular", "type": "string", "x-speakeasy-unknown-values": "allow"}}, {"description": "Start of the date window in YYYY-MM-DD (UTC), inclusive. Defaults to 30 days before `end_date`. The dataset begins at 2025-01-01; earlier values are clamped forward to that floor and the resolved value is echoed in `meta.start_date`.", "in": "query", "name": "start_date", "required": false, "schema": {"description": "Start of the date window in YYYY-MM-DD (UTC), inclusive. Defaults to 30 days before `end_date`. The dataset begins at 2025-01-01; earlier values are clamped forward to that floor and the resolved value is echoed in `meta.start_date`.", "example": "2026-04-12", "pattern": "^\\d{4}-\\d{2}-\\d{2}$", "type": "string"}}, {"description": "End of the date window in YYYY-MM-DD (UTC), inclusive. Defaults to the most recent completed UTC day. Must be on or after 2025-01-01; earlier values are rejected with a 400.", "in": "query", "name": "end_date", "required": false, "schema": {"description": "End of the date window in YYYY-MM-DD (UTC), inclusive. Defaults to the most recent completed UTC day. Must be on or after 2025-01-01; earlier values are rejected with a 400.", "example": "2026-05-11", "pattern": "^\\d{4}-\\d{2}-\\d{2}$", "type": "string"}}, {"description": "Maximum number of apps to return (1-100). Defaults to 50.", "in": "query", "name": "limit", "required": false, "schema": {"default": 50, "description": "Maximum number of apps to return (1-100). Defaults to 50.", "example": 50, "maximum": 100, "minimum": 1, "type": "integer"}}, {"description": "Number of ranked apps to skip before the first returned row (0-100). Defaults to 0. `rank` stays absolute, so the first row of `offset=50` is `rank: 51`.", "in": "query", "name": "offset", "required": false, "schema": {"default": 0, "description": "Number of ranked apps to skip before the first returned row (0-100). Defaults to 0. `rank` stays absolute, so the first row of `offset=50` is `rank: 51`.", "example": 0, "maximum": 100, "minimum": 0, "type": ["integer", "null"]}}], "responses": {"200": {"content": {"application/json": {"example": {"data": [{"app_id": 12345, "app_name": "Cline", "rank": 1, "total_requests": 4321, "total_tokens": "12345678"}, {"app_id": 67890, "app_name": "Roo Code", "rank": 2, "total_requests": 2109, "total_tokens": "9876543"}], "meta": {"as_of": "2026-05-12T02:00:00Z", "end_date": "2026-05-11", "start_date": "2026-04-12", "version": "v1"}}, "schema": {"$ref": "#/components/schemas/AppRankingsResponse"}}}, "description": "Apps ranked per the requested `sort`, re-numbered 1..N. `popular` sorts by `total_tokens` descending; `trending` sorts by absolute excess token growth descending and may return fewer than `limit` rows."}, "400": {"content": {"application/json": {"example": {"error": {"code": 400, "message": "Invalid request parameters"}}, "schema": {"$ref": "#/components/schemas/BadRequestResponse"}}}, "description": "Bad Request - Invalid request parameters or malformed input"}, "401": {"content": {"application/json": {"example": {"error": {"code": 401, "message": "Missing Authentication header"}}, "schema": {"$ref": "#/components/schemas/UnauthorizedResponse"}}}, "description": "Unauthorized - Authentication required or invalid credentials"}, "429": {"content": {"application/json": {"example": {"error": {"code": 429, "message": "Rate limit exceeded"}}, "schema": {"$ref": "#/components/schemas/TooManyRequestsResponse"}}}, "description": "Too Many Requests - Rate limit exceeded"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}}, "summary": "Top apps by token usage", "tags": ["Datasets"], "x-speakeasy-pagination": {"inputs": [{"in": "parameters", "name": "offset", "type": "offset"}, {"in": "parameters", "name": "limit", "type": "limit"}], "outputs": {"results": "$.data"}, "type": "offsetLimit"}}
```
