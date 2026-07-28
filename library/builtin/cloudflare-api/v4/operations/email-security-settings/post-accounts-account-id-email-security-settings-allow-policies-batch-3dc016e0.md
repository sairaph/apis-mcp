---
title: Batch allow policies operations
page_id: operation-post-accounts-account-id-email-security-settings-allow-policies-batch-72df8f8a
path: operations/email-security-settings
description: |-
    Executes multiple operations atomically. All four operation arrays
    (deletes, patches, puts, posts) are required and executed in order.
    Send empty arrays for unused operations.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/email-security/settings/allow_policies/batch
operation_ids:
    - email_security_batch_allow_policies
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Batch allow policies operations

`POST /accounts/{account_id}/email-security/settings/allow_policies/batch`

Operation ID: `email_security_batch_allow_policies`

Executes multiple operations atomically. All four operation arrays
(deletes, patches, puts, posts) are required and executed in order.
Send empty arrays for unused operations.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/email-security_account_id"}]
```

## Definition

```yaml
{"operationId": "email_security_batch_allow_policies", "summary": "Batch allow policies operations", "description": "Executes multiple operations atomically. All four operation arrays\n(deletes, patches, puts, posts) are required and executed in order.\nSend empty arrays for unused operations.\n", "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"deletes": {"type": "array", "items": {"properties": {"id": {"$ref": "#/components/schemas/email-security_AllowPolicyId"}}, "required": ["id"], "type": "object"}, "maxItems": 1000}, "patches": {"type": "array", "items": {"allOf": [{"$ref": "#/components/schemas/email-security_UpdateAllowPolicy"}, {"properties": {"id": {"allOf": [{"$ref": "#/components/schemas/email-security_AllowPolicyId"}]}}, "required": ["id"], "type": "object"}]}, "maxItems": 1000}, "posts": {"type": "array", "items": {"$ref": "#/components/schemas/email-security_CreateAllowPolicy"}, "maxItems": 1000}, "puts": {"type": "array", "items": {"allOf": [{"$ref": "#/components/schemas/email-security_AllowPolicy"}, {"properties": {"id": {"allOf": [{"$ref": "#/components/schemas/email-security_AllowPolicyId"}]}}, "required": ["id"], "type": "object"}]}, "maxItems": 1000}}, "required": ["deletes", "patches", "puts", "posts"]}}}}, "responses": {"200": {"description": "Batch operation results.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-security_api-response-single"}, {"properties": {"result": {"type": "object", "properties": {"deletes": {"type": "array", "items": {"$ref": "#/components/schemas/email-security_DeletedAllowPolicy"}}, "patches": {"$ref": "#/components/schemas/email-security_AllowPolicyList"}, "posts": {"$ref": "#/components/schemas/email-security_AllowPolicyList"}, "puts": {"$ref": "#/components/schemas/email-security_AllowPolicyList"}}}}, "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/email-security_ClientError"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Security Settings"], "x-api-token-group": ["Cloud Email Security: Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "email-security.settings.batch.allow.policies.batch", "x-fern-sdk-method-name": "create", "x-forge-hidden": true, "x-stability": "beta"}
```
