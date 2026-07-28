---
title: Add Prefix
page_id: operation-post-accounts-account-id-addressing-prefixes-f6a46788
path: operations/ip-address-management-prefixes
description: Add a new prefix under the account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/addressing/prefixes
operation_ids:
    - ip-address-management-prefixes-add-prefix
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Add Prefix

`POST /accounts/{account_id}/addressing/prefixes`

Operation ID: `ip-address-management-prefixes-add-prefix`

Add a new prefix under the account.

## Definition

```yaml
{"operationId": "ip-address-management-prefixes-add-prefix", "summary": "Add Prefix", "description": "Add a new prefix under the account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_account_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"properties": {"asn": {"$ref": "#/components/schemas/addressing_asn"}, "cidr": {"$ref": "#/components/schemas/addressing_cidr"}, "delegate_loa_creation": {"$ref": "#/components/schemas/addressing_delegate_loa_creation"}, "description": {"$ref": "#/components/schemas/addressing_description"}, "loa_document_id": {"$ref": "#/components/schemas/addressing_loa_document_identifier"}}, "required": ["cidr", "asn"]}}}}, "responses": {"201": {"description": "Add Prefix response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/addressing_single_response"}}}}, "4XX": {"description": "Add Prefix response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/addressing_single_response"}, {"$ref": "#/components/schemas/addressing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["IP Address Management Prefixes"], "x-api-token-group": null}
```
