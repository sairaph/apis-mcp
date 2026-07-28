---
title: Patch multiple Zero Trust Gateway rules
page_id: operation-patch-accounts-account-id-gateway-rules-e49cc750
path: operations/zero-trust-gateway-rules
description: Update select fields of multiple Zero Trust Gateway rules in a single request. This is commonly used to reorder rules by updating their precedence values. Only the fields provided for each rule are updated.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/gateway/rules
operation_ids:
    - zero-trust-gateway-rules-patch-multiple-zero-trust-gateway-rules
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Patch multiple Zero Trust Gateway rules

`PATCH /accounts/{account_id}/gateway/rules`

Operation ID: `zero-trust-gateway-rules-patch-multiple-zero-trust-gateway-rules`

Update select fields of multiple Zero Trust Gateway rules in a single request. This is commonly used to reorder rules by updating their precedence values. Only the fields provided for each rule are updated.

## Definition

```yaml
{"operationId": "zero-trust-gateway-rules-patch-multiple-zero-trust-gateway-rules", "summary": "Patch multiple Zero Trust Gateway rules", "description": "Update select fields of multiple Zero Trust Gateway rules in a single request. This is commonly used to reorder rules by updating their precedence values. Only the fields provided for each rule are updated.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_identifier-2"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "array", "items": {"properties": {"description": {"$ref": "#/components/schemas/zero-trust-gateway_description-2"}, "enabled": {"$ref": "#/components/schemas/zero-trust-gateway_enabled"}, "id": {"$ref": "#/components/schemas/zero-trust-gateway_uuid-2"}, "name": {"$ref": "#/components/schemas/zero-trust-gateway_name-3"}, "precedence": {"$ref": "#/components/schemas/zero-trust-gateway_precedence"}}, "required": ["id"], "type": "object"}}}}}, "responses": {"200": {"description": "Patch multiple Zero Trust Gateway rules response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zero-trust-gateway_response_collection-6"}}}}, "4XX": {"description": "Patch multiple Zero Trust Gateway rules response failure.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/zero-trust-gateway_response_collection-6"}, {"$ref": "#/components/schemas/zero-trust-gateway_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust Gateway rules"]}
```
