---
title: Create email allow policy
page_id: operation-post-accounts-account-id-email-security-settings-allow-policies-77090f75
path: operations/email-security-settings
description: Creates a new allow policy that exempts matching emails from security detections. Use with caution as this bypasses email security scanning. Policies can match on sender patterns and apply to specific detections or all detections.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/email-security/settings/allow_policies
operation_ids:
    - email_security_create_allow_policy
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create email allow policy

`POST /accounts/{account_id}/email-security/settings/allow_policies`

Operation ID: `email_security_create_allow_policy`

Creates a new allow policy that exempts matching emails from security detections. Use with caution as this bypasses email security scanning. Policies can match on sender patterns and apply to specific detections or all detections.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/email-security_account_id"}]
```

## Definition

```yaml
{"operationId": "email_security_create_allow_policy", "summary": "Create email allow policy", "description": "Creates a new allow policy that exempts matching emails from security detections. Use with caution as this bypasses email security scanning. Policies can match on sender patterns and apply to specific detections or all detections.", "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/email-security_CreateAllowPolicy"}}}}, "responses": {"201": {"description": "Created allow policy.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-security_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/email-security_AllowPolicy"}}, "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/email-security_ClientError"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Security Settings"], "x-api-token-group": ["Cloud Email Security: Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "email-security.settings.allow-policies", "x-fern-sdk-method-name": "create", "x-forge-hidden": true, "x-stability": "beta"}
```
