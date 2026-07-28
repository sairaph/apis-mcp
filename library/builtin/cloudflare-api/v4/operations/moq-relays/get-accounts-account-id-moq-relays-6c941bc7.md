---
title: List relays
page_id: operation-get-accounts-account-id-moq-relays-a9ebb027
path: operations/moq-relays
description: |-
    Lists all MoQ relays for the account. Returns only metadata.
    Config, status, and tokens are omitted.

    Results are cursor-paginated (keyset on the `created` timestamp).
    Use `created_before` / `created_after` with the `created` value of the
    first/last item in a page to fetch the adjacent page. `result_info`
    reports the page `count` and the `total` matching the cursor filters.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/moq/relays
operation_ids:
    - moq-relays-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List relays

`GET /accounts/{account_id}/moq/relays`

Operation ID: `moq-relays-list`

Lists all MoQ relays for the account. Returns only metadata.
Config, status, and tokens are omitted.

Results are cursor-paginated (keyset on the `created` timestamp).
Use `created_before` / `created_after` with the `created` value of the
first/last item in a page to fetch the adjacent page. `result_info`
reports the page `count` and the `total` matching the cursor filters.

## Definition

```yaml
{"operationId": "moq-relays-list", "summary": "List relays", "description": "Lists all MoQ relays for the account. Returns only metadata.\nConfig, status, and tokens are omitted.\n\nResults are cursor-paginated (keyset on the `created` timestamp).\nUse `created_before` / `created_after` with the `created` value of the\nfirst/last item in a page to fetch the adjacent page. `result_info`\nreports the page `count` and the `total` matching the cursor filters.\n", "parameters": [{"$ref": "#/components/parameters/moq_account_id"}, {"$ref": "#/components/parameters/moq_created_before"}, {"$ref": "#/components/parameters/moq_created_after"}, {"$ref": "#/components/parameters/moq_per_page"}, {"$ref": "#/components/parameters/moq_asc"}], "responses": {"200": {"description": "Relay list retrieved successfully.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/moq_api-response-common"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/moq_relay_list_item"}}, "result_info": {"$ref": "#/components/schemas/moq_result_info"}}, "type": "object"}]}}}}, "500": {"description": "Error 21006: Unexpected server error.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/moq_api-response-error"}}}}}, "security": [{"api_token": []}], "tags": ["MoQ Relays"], "x-stability": "beta"}
```
