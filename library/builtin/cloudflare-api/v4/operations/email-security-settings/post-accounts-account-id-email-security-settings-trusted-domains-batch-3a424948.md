---
title: Batch trusted domains operations
page_id: operation-post-accounts-account-id-email-security-settings-trusted-domains-batch-b706f2b2
path: operations/email-security-settings
description: |-
    Executes multiple operations atomically. All four operation arrays
    (deletes, patches, puts, posts) are required and executed in order.
    Send empty arrays for unused operations.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/email-security/settings/trusted_domains/batch
operation_ids:
    - email_security_batch_trusted_domains
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Batch trusted domains operations

`POST /accounts/{account_id}/email-security/settings/trusted_domains/batch`

Operation ID: `email_security_batch_trusted_domains`

Executes multiple operations atomically. All four operation arrays
(deletes, patches, puts, posts) are required and executed in order.
Send empty arrays for unused operations.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/email-security_account_id"}]
```

## Definition

```yaml
{"operationId": "email_security_batch_trusted_domains", "summary": "Batch trusted domains operations", "description": "Executes multiple operations atomically. All four operation arrays\n(deletes, patches, puts, posts) are required and executed in order.\nSend empty arrays for unused operations.\n", "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"deletes": {"type": "array", "items": {"properties": {"id": {"$ref": "#/components/schemas/email-security_TrustedDomainId"}}, "required": ["id"], "type": "object"}, "maxItems": 1000}, "patches": {"type": "array", "items": {"allOf": [{"$ref": "#/components/schemas/email-security_UpdateTrustedDomain"}, {"properties": {"id": {"allOf": [{"$ref": "#/components/schemas/email-security_TrustedDomainId"}]}}, "required": ["id"], "type": "object"}]}, "maxItems": 1000}, "posts": {"type": "array", "items": {"$ref": "#/components/schemas/email-security_CreateTrustedDomain"}, "maxItems": 1000}, "puts": {"type": "array", "items": {"allOf": [{"$ref": "#/components/schemas/email-security_TrustedDomain"}, {"properties": {"id": {"allOf": [{"$ref": "#/components/schemas/email-security_TrustedDomainId"}]}}, "required": ["id"], "type": "object"}]}, "maxItems": 1000}}, "required": ["deletes", "patches", "puts", "posts"]}}}}, "responses": {"200": {"description": "Batch operation results.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-security_api-response-single"}, {"properties": {"result": {"type": "object", "properties": {"deletes": {"type": "array", "items": {"$ref": "#/components/schemas/email-security_DeletedTrustedDomain"}}, "patches": {"$ref": "#/components/schemas/email-security_TrustedDomainList"}, "posts": {"$ref": "#/components/schemas/email-security_TrustedDomainList"}, "puts": {"$ref": "#/components/schemas/email-security_TrustedDomainList"}}}}, "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/email-security_ClientError"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Security Settings"], "x-api-token-group": ["Cloud Email Security: Write"], "x-stability": "beta"}
```
