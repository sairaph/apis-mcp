---
title: Get message details
page_id: operation-get-accounts-account-id-email-security-investigate-investigate-id-4fb517cb
path: operations/email-security
description: Retrieves comprehensive details for a specific email message including headers, recipients, sender information, and current quarantine status. Use the investigate_id from search results to fetch detailed information.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/email-security/investigate/{investigate_id}
operation_ids:
    - email_security_get_message
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get message details

`GET /accounts/{account_id}/email-security/investigate/{investigate_id}`

Operation ID: `email_security_get_message`

Retrieves comprehensive details for a specific email message including headers, recipients, sender information, and current quarantine status. Use the investigate_id from search results to fetch detailed information.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/email-security_account_id"}, {"name": "investigate_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email-security_InvestigateId"}}, {"$ref": "#/components/parameters/email-security_submission"}]
```

## Definition

```yaml
{"operationId": "email_security_get_message", "summary": "Get message details", "description": "Retrieves comprehensive details for a specific email message including headers, recipients, sender information, and current quarantine status. Use the investigate_id from search results to fetch detailed information.", "responses": {"200": {"description": "Email message details.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-security_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/email-security_MessageDetails"}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/email-security_ClientError"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Security"], "x-api-token-group": ["Cloud Email Security: Write", "Cloud Email Security: Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "email-security.investigate", "x-fern-sdk-method-name": "get", "x-forge-hidden": true, "x-stability": "beta"}
```
