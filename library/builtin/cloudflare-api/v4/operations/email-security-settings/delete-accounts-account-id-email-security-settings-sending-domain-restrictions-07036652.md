---
title: Delete a sending domain restriction
page_id: operation-delete-accounts-account-id-email-security-settings-sending-domain-restri-d988ede5
path: operations/email-security-settings
description: Removes a sending domain restriction. After deletion, TLS will no longer be enforced for emails from this domain.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/email-security/settings/sending_domain_restrictions/{sending_domain_restriction_id}
operation_ids:
    - email_security_delete_sending_domain_restriction
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a sending domain restriction

`DELETE /accounts/{account_id}/email-security/settings/sending_domain_restrictions/{sending_domain_restriction_id}`

Operation ID: `email_security_delete_sending_domain_restriction`

Removes a sending domain restriction. After deletion, TLS will no longer be enforced for emails from this domain.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/email-security_account_id"}, {"name": "sending_domain_restriction_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/email-security_SendingDomainRestrictionId"}}]
```

## Definition

```yaml
{"operationId": "email_security_delete_sending_domain_restriction", "summary": "Delete a sending domain restriction", "description": "Removes a sending domain restriction. After deletion, TLS will no longer be enforced for emails from this domain.", "responses": {"200": {"description": "Deleted sending domain restriction.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/email-security_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/email-security_DeletedSendingDomainRestriction"}}, "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/email-security_ClientError"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Security Settings"], "x-api-token-group": ["Cloud Email Security: Write"], "x-stability": "beta"}
```
