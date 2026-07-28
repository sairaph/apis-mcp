---
title: Get message detection details
page_id: operation-get-accounts-account-id-email-security-investigate-investigate-id-detect-0dc95366
path: operations/email-security
description: Returns detection details such as threat categories and sender information for non-benign messages.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/email-security/investigate/{investigate_id}/detections
operation_ids:
    - email_security_get_message_detections
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get message detection details

`GET /accounts/{account_id}/email-security/investigate/{investigate_id}/detections`

Operation ID: `email_security_get_message_detections`

Returns detection details such as threat categories and sender information for non-benign messages.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/email-security_account_id"}, {"name": "investigate_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email-security_InvestigateId"}}]
```

## Definition

```yaml
{"operationId": "email_security_get_message_detections", "summary": "Get message detection details", "description": "Returns detection details such as threat categories and sender information for non-benign messages.", "responses": {"200": {"description": "Email message detection details.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-security_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/email-security_MessageDetectionDetails"}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/email-security_ClientError"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Security"], "x-api-token-group": ["Cloud Email Security: Write", "Cloud Email Security: Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "email-security.investigate.detections", "x-fern-sdk-method-name": "get", "x-forge-hidden": true, "x-stability": "beta"}
```
