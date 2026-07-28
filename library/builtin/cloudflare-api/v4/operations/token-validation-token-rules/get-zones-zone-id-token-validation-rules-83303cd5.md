---
title: List token validation rules
page_id: operation-get-zones-zone-id-token-validation-rules-3e59eaf3
path: operations/token-validation-token-rules
description: List token validation rules
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/token_validation/rules
operation_ids:
    - token-validation-rules-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List token validation rules

`GET /zones/{zone_id}/token_validation/rules`

Operation ID: `token-validation-rules-list`

List token validation rules

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}]
```

## Definition

```yaml
{"operationId": "token-validation-rules-list", "summary": "List token validation rules", "description": "List token validation rules", "parameters": [{"$ref": "#/components/parameters/api-shield_per_page"}, {"$ref": "#/components/parameters/api-shield_page"}, {"name": "token_configuration", "in": "query", "description": "Select rules using any of these token configurations.", "schema": {"type": "array", "items": {"$ref": "#/components/schemas/api-shield_uuid-2"}}, "examples": {"multipleIDs": {"value": ["f0963ed4-f3a1-4599-8a61-9e5658c865e0,3aad66bb-059d-4b3d-87e1-cdf4d406f412"]}, "multipleIDsExploded": {"value": ["f0963ed4-f3a1-4599-8a61-9e5658c865e0", "3aad66bb-059d-4b3d-87e1-cdf4d406f412"]}}}, {"name": "action", "in": "query", "schema": {"$ref": "#/components/schemas/api-shield_action"}}, {"name": "enabled", "in": "query", "schema": {"$ref": "#/components/schemas/api-shield_enabled"}}, {"name": "id", "in": "query", "description": "Select rules with these IDs.", "schema": {"$ref": "#/components/schemas/api-shield_uuid-2"}}, {"name": "rule_id", "in": "query", "description": "Select rules with these IDs.", "schema": {"$ref": "#/components/schemas/api-shield_uuid-2"}}, {"name": "host", "in": "query", "description": "Select rules with this host in `include`.", "schema": {"$ref": "#/components/schemas/api-shield_host"}}, {"name": "hostname", "in": "query", "description": "Select rules with this host in `include`.", "schema": {"$ref": "#/components/schemas/api-shield_host"}}], "responses": {"200": {"$ref": "#/components/responses/api-shield_list-rules-success"}, "4XX": {"$ref": "#/components/responses/api-shield_generic_failure"}}, "externalDocs": {"description": "Learn more about JSON Web Tokens Validation.", "url": "https://developers.cloudflare.com/api-shield/security/jwt-validation/"}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Token Validation Token Rules"], "x-api-token-group": ["Account API Gateway", "Account API Gateway Read", "Domain API Gateway", "Domain API Gateway Read"]}
```
