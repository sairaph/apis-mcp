---
title: Delete a relay
page_id: operation-delete-accounts-account-id-moq-relays-relay-id-2fdaa77c
path: operations/moq-relays
description: |-
    Soft-deletes a MoQ relay. The relay ID goes in the URL path —
    `DELETE /accounts/{account_id}/moq/relays/{relay_id}` — not the
    request body; there is no collection-level delete endpoint.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/moq/relays/{relay_id}
operation_ids:
    - moq-relays-delete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a relay

`DELETE /accounts/{account_id}/moq/relays/{relay_id}`

Operation ID: `moq-relays-delete`

Soft-deletes a MoQ relay. The relay ID goes in the URL path —
`DELETE /accounts/{account_id}/moq/relays/{relay_id}` — not the
request body; there is no collection-level delete endpoint.

## Definition

```yaml
{"operationId": "moq-relays-delete", "summary": "Delete a relay", "description": "Soft-deletes a MoQ relay. The relay ID goes in the URL path —\n`DELETE /accounts/{account_id}/moq/relays/{relay_id}` — not the\nrequest body; there is no collection-level delete endpoint.\n", "parameters": [{"$ref": "#/components/parameters/moq_account_id"}, {"$ref": "#/components/parameters/moq_relay_id"}], "responses": {"200": {"description": "Relay deleted successfully.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/moq_api-response-common"}, {"properties": {"result": {"type": "object", "nullable": true}}, "type": "object"}]}}}}, "400": {"description": "Error 21003: Relay ID should be 32 hex characters.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/moq_api-response-error"}}}}, "404": {"description": "Error 21007: Relay not found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/moq_api-response-error"}}}}, "500": {"description": "Error 21006: Unexpected server error.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/moq_api-response-error"}}}}}, "security": [{"api_token": []}], "tags": ["MoQ Relays"], "x-stability": "beta"}
```
