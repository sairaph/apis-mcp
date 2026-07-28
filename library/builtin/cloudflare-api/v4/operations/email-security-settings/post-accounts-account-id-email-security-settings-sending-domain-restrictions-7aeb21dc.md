---
title: Batch sending domain restrictions operations
page_id: operation-post-accounts-account-id-email-security-settings-sending-domain-restrict-0c504444
path: operations/email-security-settings
description: Executes multiple delete operations on sending domain restrictions atomically. All operations succeed or fail together as a transaction. Currently only supports batch deletion. Removes TLS enforcement requirements for the specified domains.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/email-security/settings/sending_domain_restrictions/batch
operation_ids:
    - email_security_batch_sending_domain_restrictions
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Batch sending domain restrictions operations

`POST /accounts/{account_id}/email-security/settings/sending_domain_restrictions/batch`

Operation ID: `email_security_batch_sending_domain_restrictions`

Executes multiple delete operations on sending domain restrictions atomically. All operations succeed or fail together as a transaction. Currently only supports batch deletion. Removes TLS enforcement requirements for the specified domains.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/email-security_account_id"}]
```

## Definition

```yaml
{"operationId": "email_security_batch_sending_domain_restrictions", "summary": "Batch sending domain restrictions operations", "description": "Executes multiple delete operations on sending domain restrictions atomically. All operations succeed or fail together as a transaction. Currently only supports batch deletion. Removes TLS enforcement requirements for the specified domains.", "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"deletes": {"type": "array", "items": {"properties": {"id": {"$ref": "#/components/schemas/email-security_SendingDomainRestrictionId"}}, "required": ["id"], "type": "object"}, "maxItems": 1000}}, "required": ["deletes"]}}}}, "responses": {"200": {"description": "Batch operation results.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-security_api-response-single"}, {"properties": {"result": {"type": "object", "properties": {"deletes": {"type": "array", "items": {"properties": {"id": {"$ref": "#/components/schemas/email-security_SendingDomainRestrictionId"}}, "required": ["id"], "type": "object"}}}, "required": ["deletes"]}}, "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/email-security_ClientError"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Security Settings"], "x-api-token-group": ["Cloud Email Security: Write"], "x-stability": "beta"}
```
