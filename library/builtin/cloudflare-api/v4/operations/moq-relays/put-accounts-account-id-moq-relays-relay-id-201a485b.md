---
title: Update a relay
page_id: operation-put-accounts-account-id-moq-relays-relay-id-746da3b8
path: operations/moq-relays
description: |-
    Updates a relay's name and/or configuration. The relay ID goes in
    the URL path — `PUT /accounts/{account_id}/moq/relays/{relay_id}` —
    not the request body; there is no collection-level update endpoint.
    This is also the only way to set a relay's config (config cannot be
    set at create time). Partial updates: omitted fields are preserved;
    config sub-objects replace as whole objects when present.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/moq/relays/{relay_id}
operation_ids:
    - moq-relays-update
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update a relay

`PUT /accounts/{account_id}/moq/relays/{relay_id}`

Operation ID: `moq-relays-update`

Updates a relay's name and/or configuration. The relay ID goes in
the URL path — `PUT /accounts/{account_id}/moq/relays/{relay_id}` —
not the request body; there is no collection-level update endpoint.
This is also the only way to set a relay's config (config cannot be
set at create time). Partial updates: omitted fields are preserved;
config sub-objects replace as whole objects when present.

## Definition

```yaml
{"operationId": "moq-relays-update", "summary": "Update a relay", "description": "Updates a relay's name and/or configuration. The relay ID goes in\nthe URL path — `PUT /accounts/{account_id}/moq/relays/{relay_id}` —\nnot the request body; there is no collection-level update endpoint.\nThis is also the only way to set a relay's config (config cannot be\nset at create time). Partial updates: omitted fields are preserved;\nconfig sub-objects replace as whole objects when present.\n", "parameters": [{"$ref": "#/components/parameters/moq_account_id"}, {"$ref": "#/components/parameters/moq_relay_id"}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"config": {"$ref": "#/components/schemas/moq_relay_config"}, "name": {"type": "string"}}}}}}, "responses": {"200": {"description": "Relay updated successfully.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/moq_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/moq_relay"}}, "type": "object"}]}}}}, "400": {"description": "Bad request. Possible errors: 21003 (relay ID should be 32 hex characters), 21004 (failed to decode body, invalid JSON), 21011 (name must not be empty), 21013 (invalid upstream URL — must be an absolute moqt:// or https:// URL with a host).", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/moq_api-response-error"}}}}, "404": {"description": "Error 21007: Relay not found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/moq_api-response-error"}}}}, "500": {"description": "Error 21006: Unexpected server error.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/moq_api-response-error"}}}}}, "security": [{"api_token": []}], "tags": ["MoQ Relays"], "x-stability": "beta"}
```
