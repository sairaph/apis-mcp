---
title: Delete an email allow policy
page_id: operation-delete-accounts-account-id-email-security-settings-allow-policies-policy-912236f0
path: operations/email-security-settings
description: Removes an allow policy. After deletion, emails matching this pattern will be subject to normal security scanning and disposition actions.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/email-security/settings/allow_policies/{policy_id}
operation_ids:
    - email_security_delete_allow_policy
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete an email allow policy

`DELETE /accounts/{account_id}/email-security/settings/allow_policies/{policy_id}`

Operation ID: `email_security_delete_allow_policy`

Removes an allow policy. After deletion, emails matching this pattern will be subject to normal security scanning and disposition actions.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/email-security_account_id"}, {"name": "policy_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email-security_AllowPolicyId"}}]
```

## Definition

```yaml
{"operationId": "email_security_delete_allow_policy", "summary": "Delete an email allow policy", "description": "Removes an allow policy. After deletion, emails matching this pattern will be subject to normal security scanning and disposition actions.", "responses": {"200": {"description": "Deleted allow policy.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-security_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/email-security_DeletedAllowPolicy"}}, "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/email-security_ClientError"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Security Settings"], "x-api-token-group": ["Cloud Email Security: Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "email-security.settings.allow-policies", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true, "x-stability": "beta"}
```
