---
title: Update a user seat
page_id: operation-patch-accounts-account-id-access-seats-81d6863f
path: operations/zero-trust-seats
description: Removes a user from a Zero Trust seat when both `access_seat` and `gateway_seat` are set to false.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/access/seats
operation_ids:
    - zero-trust-seats-update-a-user-seat
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update a user seat

`PATCH /accounts/{account_id}/access/seats`

Operation ID: `zero-trust-seats-update-a-user-seat`

Removes a user from a Zero Trust seat when both `access_seat` and `gateway_seat` are set to false.

## Definition

```yaml
{"operationId": "zero-trust-seats-update-a-user-seat", "summary": "Update a user seat", "description": "Removes a user from a Zero Trust seat when both `access_seat` and `gateway_seat` are set to false.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier-2"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_seats_definition"}}}}, "responses": {"200": {"description": "Update a user seat response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_response_collection-13"}}}}, "4XX": {"description": "Update a user seat response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust seats"], "x-api-token-group": ["Zero Trust: Seats Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.seats", "x-fern-sdk-method-name": "edit", "x-forge-hidden": true}
```
