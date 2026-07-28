---
title: Change email classification
page_id: operation-post-accounts-account-id-email-security-investigate-investigate-id-recla-79ee3762
path: operations/email-security
description: Submits a request to reclassify an email's disposition. Use for reporting false positives or false negatives. Optionally provide the raw EML content for reanalysis. The reclassification is processed asynchronously.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/email-security/investigate/{investigate_id}/reclassify
operation_ids:
    - email_security_post_reclassify
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Change email classification

`POST /accounts/{account_id}/email-security/investigate/{investigate_id}/reclassify`

Operation ID: `email_security_post_reclassify`

Submits a request to reclassify an email's disposition. Use for reporting false positives or false negatives. Optionally provide the raw EML content for reanalysis. The reclassification is processed asynchronously.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/email-security_account_id"}, {"name": "investigate_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email-security_InvestigateId"}}]
```

## Definition

```yaml
{"operationId": "email_security_post_reclassify", "summary": "Change email classification", "description": "Submits a request to reclassify an email's disposition. Use for reporting false positives or false negatives. Optionally provide the raw EML content for reanalysis. The reclassification is processed asynchronously.", "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/email-security_ReclassifyRequest"}}}}, "responses": {"202": {"description": "Reclassification request accepted.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-security_api-response-common"}, {"properties": {"result": {"type": "object"}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/email-security_ClientError"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Security"], "x-api-token-group": ["Cloud Email Security: Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "email-security.investigate.reclassify", "x-fern-sdk-method-name": "create", "x-forge-hidden": true, "x-stability": "beta"}
```
