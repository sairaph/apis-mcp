---
title: Preview operations covered by a Token Validation rule
page_id: operation-post-zones-zone-id-token-validation-rules-preview-2f209b11
path: operations/token-validation-token-rules
description: |-
    Preview operations covered by a Token Validation rule.

    The API will return all operations on a zone annotated with an additional `state` field.
    Operations with an `included` `state` will be covered by a Token Validation Rule.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/token_validation/rules/preview
operation_ids:
    - token-validation-rules-preview
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Preview operations covered by a Token Validation rule

`POST /zones/{zone_id}/token_validation/rules/preview`

Operation ID: `token-validation-rules-preview`

Preview operations covered by a Token Validation rule.

The API will return all operations on a zone annotated with an additional `state` field.
Operations with an `included` `state` will be covered by a Token Validation Rule.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}]
```

## Definition

```yaml
{"operationId": "token-validation-rules-preview", "summary": "Preview operations covered by a Token Validation rule", "description": "Preview operations covered by a Token Validation rule.\n\nThe API will return all operations on a zone annotated with an additional `state` field.\nOperations with an `included` `state` will be covered by a Token Validation Rule.\n", "parameters": [{"$ref": "#/components/parameters/api-shield_per_page"}, {"$ref": "#/components/parameters/api-shield_page"}, {"name": "state", "in": "query", "schema": {"type": "array", "items": {"$ref": "#/components/schemas/api-shield_selector-operation-state"}}, "examples": {"multiple": {"value": ["included,excluded"]}, "multipleExploded": {"value": ["included", "excluded"]}}}, {"name": "host", "in": "query", "description": "Filter operations by host.", "schema": {"type": "array", "items": {"$ref": "#/components/schemas/api-shield_host"}}}, {"name": "hostname", "in": "query", "description": "Filter operations by host.", "schema": {"type": "array", "items": {"$ref": "#/components/schemas/api-shield_host"}}}, {"name": "method", "in": "query", "description": "Filter operations by method.", "schema": {"type": "array", "items": {"$ref": "#/components/schemas/api-shield_method"}}}, {"name": "endpoint", "in": "query", "description": "Filter operations by endpoint. Allows substring matching.", "schema": {"type": "array", "items": {"$ref": "#/components/schemas/api-shield_endpoint"}}}], "requestBody": {"$ref": "#/components/requestBodies/api-shield_preview-rules"}, "responses": {"200": {"$ref": "#/components/responses/api-shield_preview-rules-success"}, "4XX": {"$ref": "#/components/responses/api-shield_generic_failure"}}, "externalDocs": {"description": "Learn more about JSON Web Tokens Validation.", "url": "https://developers.cloudflare.com/api-shield/security/jwt-validation/"}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Token Validation Token Rules"], "x-api-token-group": ["Account API Gateway", "Account API Gateway Read", "Domain API Gateway", "Domain API Gateway Read"]}
```
