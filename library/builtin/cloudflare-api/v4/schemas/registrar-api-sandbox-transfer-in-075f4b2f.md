---
title: registrar-api-sandbox_transfer_in
page_id: schema-registrar-api-sandbox-transfer-in-075f4b2f
path: schemas
description: Statuses for domain transfers into Cloudflare Registrar.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# registrar-api-sandbox_transfer_in

Statuses for domain transfers into Cloudflare Registrar.

```yaml
{"description": "Statuses for domain transfers into Cloudflare Registrar.", "properties": {"accept_foa": {"description": "Form of authorization has been accepted by the registrant.", "type": "string", "example": "needed", "enum": ["needed", "ok"]}, "approve_transfer": {"description": "Shows transfer status with the registry.", "type": "string", "example": "unknown", "enum": ["needed", "ok", "pending", "trying", "rejected", "unknown"]}, "can_cancel_transfer": {"description": "Indicates if cancellation is still possible.", "type": "boolean", "example": true}, "disable_privacy": {"description": "Privacy guards are disabled at the foreign registrar.", "type": "string", "example": "ok", "enum": ["needed", "ok", "unknown"]}, "enter_auth_code": {"description": "Auth code has been entered and verified.", "type": "string", "example": "needed", "enum": ["needed", "ok", "pending", "trying", "rejected"]}, "unlock_domain": {"description": "Domain is unlocked at the foreign registrar.", "type": "string", "example": "ok", "enum": ["needed", "ok", "pending", "trying", "unknown"]}}}
```
