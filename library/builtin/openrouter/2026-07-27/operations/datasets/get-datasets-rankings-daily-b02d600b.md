---
title: Daily token totals for top 50 models
page_id: operation-get-datasets-rankings-daily-ed35e699
path: operations/datasets
description: |-
    Returns the top 50 public models per day by total token usage on OpenRouter, plus a
    single aggregated `other` row per day that sums every model outside that top 50.
    Token totals are `prompt_tokens + completion_tokens`, matching the public rankings
    chart on openrouter.ai/rankings.

    Each row is a distinct `(date, model_permaslug)` pair. The `other` row uses the
    reserved permaslug `other` and is always returned last within its date, so callers
    can compute `top-50 traffic / total daily traffic` without a second request.

    Optional filters slice the dataset. `period` (`day`/`week`/`month`) sets the time
    grain. `modality` and `context_bucket` narrow the exact dataset by output/input
    modality (or tool-calling activity) and request context length. `category` and
    `language_type` instead read a sampled, upsampled dataset whose `total_tokens` are
    weekly-grain estimates — they cannot be combined with each other or with the exact
    filters, and reject `period=day` with a 400.

    Authenticate with any valid OpenRouter API key (same key used for inference).
    Rate-limited to 30 requests/minute per key and 500 requests/day per account.

    When republishing or quoting this dataset, OpenRouter must be cited as:
    "Source: OpenRouter (openrouter.ai/rankings), as of {as_of}."

    Token counts come from each upstream provider's own tokenizer (Anthropic counts
    are as reported by Anthropic, OpenAI counts are as reported by OpenAI, etc.), so
    a token in one row is not directly comparable to a token in another row from a
    different provider.
source: https://openrouter.ai/openapi.json
http_methods:
    - GET
api_endpoints:
    - /datasets/rankings-daily
operation_ids:
    - getRankingsDaily
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# Daily token totals for top 50 models

`GET /datasets/rankings-daily`

Operation ID: `getRankingsDaily`

Returns the top 50 public models per day by total token usage on OpenRouter, plus a
single aggregated `other` row per day that sums every model outside that top 50.
Token totals are `prompt_tokens + completion_tokens`, matching the public rankings
chart on openrouter.ai/rankings.

Each row is a distinct `(date, model_permaslug)` pair. The `other` row uses the
reserved permaslug `other` and is always returned last within its date, so callers
can compute `top-50 traffic / total daily traffic` without a second request.

Optional filters slice the dataset. `period` (`day`/`week`/`month`) sets the time
grain. `modality` and `context_bucket` narrow the exact dataset by output/input
modality (or tool-calling activity) and request context length. `category` and
`language_type` instead read a sampled, upsampled dataset whose `total_tokens` are
weekly-grain estimates — they cannot be combined with each other or with the exact
filters, and reject `period=day` with a 400.

Authenticate with any valid OpenRouter API key (same key used for inference).
Rate-limited to 30 requests/minute per key and 500 requests/day per account.

When republishing or quoting this dataset, OpenRouter must be cited as:
"Source: OpenRouter (openrouter.ai/rankings), as of {as_of}."

Token counts come from each upstream provider's own tokenizer (Anthropic counts
are as reported by Anthropic, OpenAI counts are as reported by OpenAI, etc.), so
a token in one row is not directly comparable to a token in another row from a
different provider.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"description": "Returns the top 50 public models per day by total token usage on OpenRouter, plus a\nsingle aggregated `other` row per day that sums every model outside that top 50.\nToken totals are `prompt_tokens + completion_tokens`, matching the public rankings\nchart on openrouter.ai/rankings.\n\nEach row is a distinct `(date, model_permaslug)` pair. The `other` row uses the\nreserved permaslug `other` and is always returned last within its date, so callers\ncan compute `top-50 traffic / total daily traffic` without a second request.\n\nOptional filters slice the dataset. `period` (`day`/`week`/`month`) sets the time\ngrain. `modality` and `context_bucket` narrow the exact dataset by output/input\nmodality (or tool-calling activity) and request context length. `category` and\n`language_type` instead read a sampled, upsampled dataset whose `total_tokens` are\nweekly-grain estimates — they cannot be combined with each other or with the exact\nfilters, and reject `period=day` with a 400.\n\nAuthenticate with any valid OpenRouter API key (same key used for inference).\nRate-limited to 30 requests/minute per key and 500 requests/day per account.\n\nWhen republishing or quoting this dataset, OpenRouter must be cited as:\n\"Source: OpenRouter (openrouter.ai/rankings), as of {as_of}.\"\n\nToken counts come from each upstream provider's own tokenizer (Anthropic counts\nare as reported by Anthropic, OpenAI counts are as reported by OpenAI, etc.), so\na token in one row is not directly comparable to a token in another row from a\ndifferent provider.", "operationId": "getRankingsDaily", "parameters": [{"description": "Start of the date window in YYYY-MM-DD (UTC), inclusive. Defaults to 30 days before `end_date`. The dataset begins at 2025-01-01; earlier values are clamped forward to that floor and the resolved value is echoed in `meta.start_date`.", "in": "query", "name": "start_date", "required": false, "schema": {"description": "Start of the date window in YYYY-MM-DD (UTC), inclusive. Defaults to 30 days before `end_date`. The dataset begins at 2025-01-01; earlier values are clamped forward to that floor and the resolved value is echoed in `meta.start_date`.", "example": "2026-04-12", "pattern": "^\\d{4}-\\d{2}-\\d{2}$", "type": "string"}}, {"description": "End of the date window in YYYY-MM-DD (UTC), inclusive. Defaults to the most recent completed UTC day. Must be on or after 2025-01-01; earlier values are rejected with a 400.", "in": "query", "name": "end_date", "required": false, "schema": {"description": "End of the date window in YYYY-MM-DD (UTC), inclusive. Defaults to the most recent completed UTC day. Must be on or after 2025-01-01; earlier values are rejected with a 400.", "example": "2026-05-11", "pattern": "^\\d{4}-\\d{2}-\\d{2}$", "type": "string"}}, {"description": "Time grain of each row. `day` (default) returns the per-UTC-day series; `week` buckets by ISO week start; `month` buckets by month start. With `category` or `language_type` only `week` (default) and `month` are available — `day` is rejected with a 400 because those datasets are aggregated weekly. For those sampled datasets `period=month` buckets each week by its week-start month, so totals are approximate at month boundaries.", "in": "query", "name": "period", "required": false, "schema": {"description": "Time grain of each row. `day` (default) returns the per-UTC-day series; `week` buckets by ISO week start; `month` buckets by month start. With `category` or `language_type` only `week` (default) and `month` are available — `day` is rejected with a 400 because those datasets are aggregated weekly. For those sampled datasets `period=month` buckets each week by its week-start month, so totals are approximate at month boundaries.", "enum": ["day", "week", "month"], "example": "day", "type": "string", "x-speakeasy-unknown-values": "allow"}}, {"description": "Restrict to models for a modality surface: `text` / `image_output` match output modality, `image` / `audio` match input modality, and `tool_calling` keeps only rows that recorded at least one tool call. Exact dataset — cannot be combined with `category` or `language_type`.", "in": "query", "name": "modality", "required": false, "schema": {"description": "Restrict to models for a modality surface: `text` / `image_output` match output modality, `image` / `audio` match input modality, and `tool_calling` keeps only rows that recorded at least one tool call. Exact dataset — cannot be combined with `category` or `language_type`.", "enum": ["text", "image", "image_output", "audio", "tool_calling"], "example": "text", "type": "string", "x-speakeasy-unknown-values": "allow"}}, {"description": "Restrict to requests whose context length falls in this bucket (`1K`, `10K`, `100K`, `1M`, or `10M`). Exact dataset — cannot be combined with `category` or `language_type`.", "in": "query", "name": "context_bucket", "required": false, "schema": {"description": "Restrict to requests whose context length falls in this bucket (`1K`, `10K`, `100K`, `1M`, or `10M`). Exact dataset — cannot be combined with `category` or `language_type`.", "enum": ["1K", "10K", "100K", "1M", "10M"], "example": "100K", "type": "string", "x-speakeasy-unknown-values": "allow"}}, {"description": "Restrict to a use-case category (e.g. `programming`, `roleplay`). Sourced from a sampled, upsampled dataset, so `total_tokens` is an estimate and is aggregated weekly (the trailing weekly bucket may include traffic past `end_date`). Cannot be combined with `modality`, `context_bucket`, or `language_type`.", "in": "query", "name": "category", "required": false, "schema": {"description": "Restrict to a use-case category (e.g. `programming`, `roleplay`). Sourced from a sampled, upsampled dataset, so `total_tokens` is an estimate and is aggregated weekly (the trailing weekly bucket may include traffic past `end_date`). Cannot be combined with `modality`, `context_bucket`, or `language_type`.", "enum": ["programming", "roleplay", "marketing", "marketing/seo", "technology", "science", "translation", "legal", "finance", "health", "trivia", "academia"], "example": "programming", "type": "string", "x-speakeasy-unknown-values": "allow"}}, {"description": "Restrict to natural-language or programming-language tagged activity. Sourced from a sampled, upsampled dataset, so `total_tokens` is an estimate and is aggregated weekly (the trailing weekly bucket may include traffic past `end_date`). Cannot be combined with `modality`, `context_bucket`, or `category`.", "in": "query", "name": "language_type", "required": false, "schema": {"description": "Restrict to natural-language or programming-language tagged activity. Sourced from a sampled, upsampled dataset, so `total_tokens` is an estimate and is aggregated weekly (the trailing weekly bucket may include traffic past `end_date`). Cannot be combined with `modality`, `context_bucket`, or `category`.", "enum": ["natural", "programming"], "example": "natural", "type": "string", "x-speakeasy-unknown-values": "allow"}}], "responses": {"200": {"content": {"application/json": {"example": {"data": [{"date": "2026-05-11", "model_permaslug": "openai/gpt-4o-2024-05-13", "total_tokens": "12345678"}, {"date": "2026-05-11", "model_permaslug": "anthropic/claude-3.5-sonnet-20241022", "total_tokens": "9876543"}, {"date": "2026-05-11", "model_permaslug": "other", "total_tokens": "4321098"}], "meta": {"as_of": "2026-05-12T02:00:00Z", "end_date": "2026-05-11", "start_date": "2026-04-12", "version": "v1"}}, "schema": {"$ref": "#/components/schemas/RankingsDailyResponse"}}}, "description": "Up to 51 rows per day — the top 50 public models by `total_tokens` plus a single aggregated `other` row covering every model outside that top 50. Sorted by `date` ascending, then by `total_tokens` descending, with `other` pinned last within its date."}, "400": {"content": {"application/json": {"example": {"error": {"code": 400, "message": "Invalid request parameters"}}, "schema": {"$ref": "#/components/schemas/BadRequestResponse"}}}, "description": "Bad Request - Invalid request parameters or malformed input"}, "401": {"content": {"application/json": {"example": {"error": {"code": 401, "message": "Missing Authentication header"}}, "schema": {"$ref": "#/components/schemas/UnauthorizedResponse"}}}, "description": "Unauthorized - Authentication required or invalid credentials"}, "429": {"content": {"application/json": {"example": {"error": {"code": 429, "message": "Rate limit exceeded"}}, "schema": {"$ref": "#/components/schemas/TooManyRequestsResponse"}}}, "description": "Too Many Requests - Rate limit exceeded"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}}, "summary": "Daily token totals for top 50 models", "tags": ["Datasets"]}
```
