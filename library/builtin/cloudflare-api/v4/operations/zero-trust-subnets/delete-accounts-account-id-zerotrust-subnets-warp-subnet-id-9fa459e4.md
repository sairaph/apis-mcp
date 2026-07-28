---
title: Delete WARP IP subnet
page_id: operation-delete-accounts-account-id-zerotrust-subnets-warp-subnet-id-336b3dd7
path: operations/zero-trust-subnets
description: Delete a WARP IP assignment subnet. This operation is idempotent - deleting an already-deleted or non-existent subnet will return success with a null result.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/zerotrust/subnets/warp/{subnet_id}
operation_ids:
    - zero-trust-networks-subnet-delete-warp
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete WARP IP subnet

`DELETE /accounts/{account_id}/zerotrust/subnets/warp/{subnet_id}`

Operation ID: `zero-trust-networks-subnet-delete-warp`

Delete a WARP IP assignment subnet. This operation is idempotent - deleting an already-deleted or non-existent subnet will return success with a null result.

## Definition

```yaml
{"operationId": "zero-trust-networks-subnet-delete-warp", "summary": "Delete WARP IP subnet", "description": "Delete a WARP IP assignment subnet. This operation is idempotent - deleting an already-deleted or non-existent subnet will return success with a null result.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_account_id"}}, {"name": "subnet_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_subnet_id"}}], "responses": {"200": {"description": "Delete subnet response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tunnel_subnet_response_single_nullable"}}}}, "4XX": {"description": "Delete subnet response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tunnel_subnet_response_single_nullable"}, {"$ref": "#/components/schemas/tunnel_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["Zero Trust Subnets"], "x-api-token-group": ["Cloudflare One Networks Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.networks.subnets.warp", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
