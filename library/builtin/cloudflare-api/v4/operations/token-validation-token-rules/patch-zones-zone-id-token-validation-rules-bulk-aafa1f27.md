---
title: Bulk edit token validation rules
page_id: operation-patch-zones-zone-id-token-validation-rules-bulk-945f22f7
path: operations/token-validation-token-rules
description: |-
    Edit token validation rules.

    A request can update multiple Token Validation Rules.

    Rules can be re-ordered using the `position` field.

    Returns all updated rules.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /zones/{zone_id}/token_validation/rules/bulk
operation_ids:
    - token-validation-rules-bulk-edit
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Bulk edit token validation rules

`PATCH /zones/{zone_id}/token_validation/rules/bulk`

Operation ID: `token-validation-rules-bulk-edit`

Edit token validation rules.

A request can update multiple Token Validation Rules.

Rules can be re-ordered using the `position` field.

Returns all updated rules.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}]
```

## Definition

```yaml
{"operationId": "token-validation-rules-bulk-edit", "summary": "Bulk edit token validation rules", "description": "Edit token validation rules.\n\nA request can update multiple Token Validation Rules.\n\nRules can be re-ordered using the `position` field.\n\nReturns all updated rules.\n", "requestBody": {"$ref": "#/components/requestBodies/api-shield_bulk-edit-rules"}, "responses": {"200": {"$ref": "#/components/responses/api-shield_bulk-edit-rules-success"}, "4XX": {"$ref": "#/components/responses/api-shield_generic_failure"}}, "externalDocs": {"description": "Learn more about JSON Web Tokens Validation.", "url": "https://developers.cloudflare.com/api-shield/security/jwt-validation/"}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Token Validation Token Rules"], "x-api-token-group": ["Account API Gateway", "Domain API Gateway"]}
```
