---
title: Update an email allow policy
page_id: operation-patch-accounts-account-id-email-security-settings-allow-policies-policy-03345943
path: operations/email-security-settings
description: Updates an existing allow policy. Only provided fields will be modified. Changes take effect for new emails matching the pattern.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/email-security/settings/allow_policies/{policy_id}
operation_ids:
    - email_security_update_allow_policy
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update an email allow policy

`PATCH /accounts/{account_id}/email-security/settings/allow_policies/{policy_id}`

Operation ID: `email_security_update_allow_policy`

Updates an existing allow policy. Only provided fields will be modified. Changes take effect for new emails matching the pattern.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/email-security_account_id"}, {"name": "policy_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email-security_AllowPolicyId"}}]
```

## Definition

```yaml
{"operationId": "email_security_update_allow_policy", "summary": "Update an email allow policy", "description": "Updates an existing allow policy. Only provided fields will be modified. Changes take effect for new emails matching the pattern.", "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/email-security_UpdateAllowPolicy"}}}}, "responses": {"200": {"description": "Updated allow policy.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-security_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/email-security_AllowPolicy"}}, "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/email-security_ClientError"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Security Settings"], "x-api-token-group": ["Cloud Email Security: Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "email-security.settings.allow-policies", "x-fern-sdk-method-name": "edit", "x-forge-hidden": true, "x-stability": "beta"}
```
