---
title: Get Multiple Domain Details
page_id: operation-get-accounts-account-id-intel-domain-bulk-9e19c39f
path: operations/domain-intelligence
description: |-
    Returns security details and statistics about multiple domains in a
    single request.

    **Behavior change — domain ranking is becoming opt-in.** This endpoint
    previously included domain ranking data in every response and accepted
    a `skip_ranking=true` query parameter to opt out. That parameter is
    being deprecated and ranking will no longer be returned by default.
    Callers that want ranking data must pass `include_ranking=true`. The
    `skip_ranking` parameter will be silently ignored once the change ships.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/intel/domain/bulk
operation_ids:
    - domain-intelligence-get-multiple-domain-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Multiple Domain Details

`GET /accounts/{account_id}/intel/domain/bulk`

Operation ID: `domain-intelligence-get-multiple-domain-details`

Returns security details and statistics about multiple domains in a
single request.

**Behavior change — domain ranking is becoming opt-in.** This endpoint
previously included domain ranking data in every response and accepted
a `skip_ranking=true` query parameter to opt out. That parameter is
being deprecated and ranking will no longer be returned by default.
Callers that want ranking data must pass `include_ranking=true`. The
`skip_ranking` parameter will be silently ignored once the change ships.

## Definition

```yaml
{"operationId": "domain-intelligence-get-multiple-domain-details", "summary": "Get Multiple Domain Details", "description": "Returns security details and statistics about multiple domains in a\nsingle request.\n\n**Behavior change — domain ranking is becoming opt-in.** This endpoint\npreviously included domain ranking data in every response and accepted\na `skip_ranking=true` query parameter to opt out. That parameter is\nbeing deprecated and ranking will no longer be returned by default.\nCallers that want ranking data must pass `include_ranking=true`. The\n`skip_ranking` parameter will be silently ignored once the change ships.\n", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/intel_identifier"}}, {"name": "domain", "in": "query", "description": "Accepts multiple values like `?domain=cloudflare.com&domain=example.com`.", "schema": {"type": "array", "items": {"type": "string"}}, "explode": true, "style": "form"}, {"name": "include_ranking", "in": "query", "description": "Whether to include domain ranking data in the response. Defaults to\n`false` — ranking lookups are expensive at bulk scale and most\ncallers do not need them. Set to `true` to opt in. This parameter\nreplaces the deprecated `skip_ranking` (see below).\n", "schema": {"type": "boolean", "default": false}}, {"name": "skip_ranking", "in": "query", "description": "**Deprecated.** Previously controlled whether the ranking lookup\nwas skipped (defaulted to `false`, meaning ranking ran). The\nendpoint's default behavior is being flipped — ranking is now\nopt-in via `include_ranking=true` — and this parameter will be\nsilently ignored. Remove it from your callers and use\n`include_ranking` instead.\n", "schema": {"type": "boolean"}, "deprecated": true}], "responses": {"200": {"description": "Get Multiple Domain Details response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/intel_collection_response"}}}}, "4XX": {"description": "Get Multiple Domain Details response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/intel_collection_response"}, {"$ref": "#/components/schemas/intel_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Domain Intelligence"], "x-api-token-group": ["Intel Write", "Intel Read"]}
```
