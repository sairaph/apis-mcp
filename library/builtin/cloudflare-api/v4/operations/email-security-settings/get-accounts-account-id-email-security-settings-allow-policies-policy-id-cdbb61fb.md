---
title: Get an email allow policy
page_id: operation-get-accounts-account-id-email-security-settings-allow-policies-policy-id-fe24f158
path: operations/email-security-settings
description: Retrieves details for a specific allow policy including its pattern, dispositions that are exempted, and whether it applies to all detections.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/email-security/settings/allow_policies/{policy_id}
operation_ids:
    - email_security_get_allow_policy
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get an email allow policy

`GET /accounts/{account_id}/email-security/settings/allow_policies/{policy_id}`

Operation ID: `email_security_get_allow_policy`

Retrieves details for a specific allow policy including its pattern, dispositions that are exempted, and whether it applies to all detections.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/email-security_account_id"}, {"name": "policy_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email-security_AllowPolicyId"}}]
```

## Definition

```yaml
{"operationId": "email_security_get_allow_policy", "summary": "Get an email allow policy", "description": "Retrieves details for a specific allow policy including its pattern, dispositions that are exempted, and whether it applies to all detections.", "responses": {"200": {"description": "Allow policy details.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-security_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/email-security_AllowPolicy"}}, "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/email-security_ClientError"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Security Settings"], "x-api-token-group": ["Cloud Email Security: Write", "Cloud Email Security: Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "email-security.settings.allow-policies", "x-fern-sdk-method-name": "get", "x-forge-hidden": true, "x-stability": "beta"}
```
