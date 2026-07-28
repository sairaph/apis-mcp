---
title: access_mfa_config
page_id: schema-access-mfa-config-9c18ca92
path: schemas
description: Configures multi-factor authentication (MFA) settings.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_mfa_config

Configures multi-factor authentication (MFA) settings.

```yaml
{"description": "Configures multi-factor authentication (MFA) settings.", "type": "object", "properties": {"allowed_authenticators": {"description": "Lists the MFA methods that users can authenticate with.", "type": "array", "items": {"enum": ["totp", "biometrics", "security_key"], "type": "string"}, "example": ["totp", "biometrics", "security_key"]}, "mfa_disabled": {"description": "Indicates whether to disable MFA for this resource. This option is available at the application and policy level.", "type": "boolean", "example": false}, "session_duration": {"description": "Defines the duration of an MFA session. Must be in minutes (m) or hours (h). Minimum: 0m. Maximum: 720h (30 days). Examples:`5m` or `24h`.", "type": "string", "example": "24h"}}, "x-auditable": true}
```
