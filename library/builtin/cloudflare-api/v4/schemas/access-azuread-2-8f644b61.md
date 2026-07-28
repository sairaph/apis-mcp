---
title: access_azureAD-2
page_id: schema-access-azuread-2-8f644b61
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_azureAD-2

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/access_identity-provider-2"}, {"properties": {"config": {"allOf": [{"$ref": "#/components/schemas/access_generic-oauth-config-2"}, {"properties": {"conditional_access_enabled": {"description": "Should Cloudflare try to load authentication contexts from your account", "type": "boolean"}, "directory_id": {"description": "Your Azure directory uuid", "type": "string", "example": "<your azure directory uuid>"}, "prompt": {"description": "Indicates the type of user interaction that is required. prompt=login forces the user to enter their credentials on that request, negating single-sign on. prompt=none is the opposite. It ensures that the user isn't presented with any interactive prompt. If the request can't be completed silently by using single-sign on, the Microsoft identity platform returns an interaction_required error. prompt=select_account interrupts single sign-on providing account selection experience listing all the accounts either in session or any remembered account or an option to choose to use a different account altogether.", "type": "string", "enum": ["login", "select_account", "none"]}, "support_groups": {"description": "Should Cloudflare try to load groups from your account", "type": "boolean"}}, "type": "object"}]}}, "type": "object"}], "title": "Azure AD"}
```
