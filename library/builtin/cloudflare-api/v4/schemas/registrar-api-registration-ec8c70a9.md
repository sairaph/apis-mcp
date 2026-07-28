---
title: registrar-api_registration
page_id: schema-registrar-api-registration-ec8c70a9
path: schemas
description: A domain registration resource representing the current state of a registered domain.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# registrar-api_registration

A domain registration resource representing the current state of a registered domain.

```yaml
{"description": "A domain registration resource representing the current state of a registered domain.", "type": "object", "properties": {"auto_renew": {"description": "Whether the domain will be automatically renewed before expiration.", "type": "boolean", "example": true}, "created_at": {"description": "When the domain was registered. Present when the registration resource exists.", "type": "string", "format": "date-time", "example": "2025-01-15T10:00:00Z"}, "domain_name": {"$ref": "#/components/schemas/registrar-api_domain_name"}, "expires_at": {"description": "When the domain registration expires. Present when the registration is ready; may be null only while `status` is `registration_pending`.", "type": "string", "format": "date-time", "example": "2026-01-15T10:00:00Z", "nullable": true}, "locked": {"description": "Whether the domain is locked for transfer.", "type": "boolean", "example": true}, "privacy_mode": {"description": "Current WHOIS privacy mode for the registration.", "type": "string", "example": "redaction", "enum": [false, "redaction"]}, "status": {"description": "Current registration status.\n- `active`: Domain is registered and operational\n- `registration_pending`: Registration is in progress\n- `expired`: Domain has expired\n- `suspended`: Domain is suspended by the registry\n- `redemption_period`: Domain is in the redemption grace period\n- `pending_delete`: Domain is pending deletion by the registry\n", "type": "string", "example": "active", "enum": ["active", "registration_pending", "expired", "suspended", "redemption_period", "pending_delete"]}}, "required": ["domain_name", "status", "created_at", "expires_at", "auto_renew", "privacy_mode", "locked"]}
```
