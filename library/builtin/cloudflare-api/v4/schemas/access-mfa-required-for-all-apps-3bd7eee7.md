---
title: access_mfa_required_for_all_apps
page_id: schema-access-mfa-required-for-all-apps-3bd7eee7
path: schemas
description: 'Determines whether global MFA settings apply to applications by default. The organization must have MFA enabled with at least one authentication method and a session duration configured. Note: ''allowed_authenticators'' cannot only contain ''piv_key'' if the organization has any non-infrastructure applications because PIV keys are only compatible with infrastructure apps.'
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_mfa_required_for_all_apps

Determines whether global MFA settings apply to applications by default. The organization must have MFA enabled with at least one authentication method and a session duration configured. Note: 'allowed_authenticators' cannot only contain 'piv_key' if the organization has any non-infrastructure applications because PIV keys are only compatible with infrastructure apps.

```yaml
{"description": "Determines whether global MFA settings apply to applications by default. The organization must have MFA enabled with at least one authentication method and a session duration configured. Note: 'allowed_authenticators' cannot only contain 'piv_key' if the organization has any non-infrastructure applications because PIV keys are only compatible with infrastructure apps.", "type": "boolean", "example": false, "default": false, "x-auditable": true}
```
