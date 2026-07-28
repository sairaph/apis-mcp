---
title: Bulk create token validation rules
page_id: operation-post-zones-zone-id-token-validation-rules-bulk-3e9f4f30
path: operations/token-validation-token-rules
description: |-
    Create zone token validation rules.

    A request can create multiple Token Validation Rules.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/token_validation/rules/bulk
operation_ids:
    - token-validation-rules-bulk-create
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Bulk create token validation rules

`POST /zones/{zone_id}/token_validation/rules/bulk`

Operation ID: `token-validation-rules-bulk-create`

Create zone token validation rules.

A request can create multiple Token Validation Rules.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}]
```

## Definition

```yaml
{"operationId": "token-validation-rules-bulk-create", "summary": "Bulk create token validation rules", "description": "Create zone token validation rules.\n\nA request can create multiple Token Validation Rules.\n", "requestBody": {"$ref": "#/components/requestBodies/api-shield_bulk-create-rules"}, "responses": {"200": {"$ref": "#/components/responses/api-shield_bulk-create-rules-success"}, "4XX": {"$ref": "#/components/responses/api-shield_generic_failure"}}, "externalDocs": {"description": "Learn more about JSON Web Tokens Validation.", "url": "https://developers.cloudflare.com/api-shield/security/jwt-validation/"}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Token Validation Token Rules"], "x-api-token-group": ["Account API Gateway", "Domain API Gateway"]}
```
