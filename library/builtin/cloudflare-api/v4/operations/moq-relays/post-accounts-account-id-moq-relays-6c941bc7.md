---
title: Create a relay
page_id: operation-post-accounts-account-id-moq-relays-e69890e7
path: operations/moq-relays
description: |-
    Provisions a new MoQ relay instance. Auto-creates a publish+subscribe
    token and a subscribe-only token. Token values are included in the
    response (shown once). Config is always set to defaults (upstreams
    off) and cannot be supplied here — sending a non-empty `config` is
    rejected (21014); `null` or `{}` is accepted as absent. Use PUT to
    configure the relay after it exists.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/moq/relays
operation_ids:
    - moq-relays-create
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a relay

`POST /accounts/{account_id}/moq/relays`

Operation ID: `moq-relays-create`

Provisions a new MoQ relay instance. Auto-creates a publish+subscribe
token and a subscribe-only token. Token values are included in the
response (shown once). Config is always set to defaults (upstreams
off) and cannot be supplied here — sending a non-empty `config` is
rejected (21014); `null` or `{}` is accepted as absent. Use PUT to
configure the relay after it exists.

## Definition

```yaml
{"operationId": "moq-relays-create", "summary": "Create a relay", "description": "Provisions a new MoQ relay instance. Auto-creates a publish+subscribe\ntoken and a subscribe-only token. Token values are included in the\nresponse (shown once). Config is always set to defaults (upstreams\noff) and cannot be supplied here — sending a non-empty `config` is\nrejected (21014); `null` or `{}` is accepted as absent. Use PUT to\nconfigure the relay after it exists.\n", "parameters": [{"$ref": "#/components/parameters/moq_account_id"}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"name": {"description": "Human-readable name for the relay.", "type": "string", "example": "Production Live Stream", "minLength": 1}}, "required": ["name"]}}}}, "responses": {"201": {"description": "Relay created successfully.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/moq_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/moq_relay_create_response"}}, "type": "object"}]}}}}, "400": {"description": "Bad request. Possible errors:\n- 21002: Request body too small\n- 21004: Failed to decode body (invalid JSON)\n- 21011: Invalid relay name (must not be empty)\n- 21014: Config cannot be set on create (set it via PUT after the relay exists)\n", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/moq_api-response-error"}}}}, "409": {"description": "Error 21008: Relay limit exceeded for this account.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/moq_api-response-error"}}}}, "413": {"description": "Error 21001: Request body too large.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/moq_api-response-error"}}}}, "500": {"description": "Error 21006: Unexpected server error.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/moq_api-response-error"}}}}}, "security": [{"api_token": []}], "tags": ["MoQ Relays"], "x-stability": "beta"}
```
