---
title: access_org_mfa_config
page_id: schema-access-org-mfa-config-5d9c2e7f
path: schemas
description: Configures multi-factor authentication (MFA) settings for an organization.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_org_mfa_config

Configures multi-factor authentication (MFA) settings for an organization.

```yaml
{"description": "Configures multi-factor authentication (MFA) settings for an organization.", "type": "object", "properties": {"allowed_authenticators": {"description": "Lists the MFA methods that users can authenticate with.", "type": "array", "items": {"enum": ["totp", "biometrics", "security_key", "piv_key"], "type": "string"}, "example": ["totp", "biometrics", "security_key"]}, "amr_matching_session_duration": {"description": "Allows a user to skip MFA via Authentication Method Reference (AMR) matching when the AMR claim provided by the IdP the user used to authenticate contains \"mfa\". Must be in minutes (m) or hours (h). Minimum: 0m. Maximum: 720h (30 days).", "type": "string", "example": "12h"}, "required_aaguids": {"description": "Specifies a Cloudflare List of required FIDO2 authenticator device AAGUIDs.", "type": "string", "format": "uuid", "example": "2fc0579f-8113-47ea-b116-bb5a8db9202a"}, "session_duration": {"description": "Defines the duration of an MFA session. Must be in minutes (m) or hours (h). Minimum: 0m. Maximum: 720h (30 days). Examples:`5m` or `24h`.", "type": "string", "example": "24h"}}, "x-auditable": true}
```
