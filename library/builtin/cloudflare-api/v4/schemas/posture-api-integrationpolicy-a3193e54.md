---
title: posture-api_IntegrationPolicy
page_id: schema-posture-api-integrationpolicy-a3193e54
path: schemas
description: Policy configuration for an integration.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# posture-api_IntegrationPolicy

Policy configuration for an integration.

```yaml
{"description": "Policy configuration for an integration.", "type": "object", "properties": {"client_id": {"description": "OAuth client ID for the policy.", "type": "string", "nullable": true}, "compliance_level": {"description": "Compliance level for the policy.", "type": "string", "example": "standard"}, "dlp_enabled": {"description": "Whether DLP is enabled for this policy.", "type": "boolean", "example": true}, "id": {"description": "Policy identifier.", "type": "string", "format": "uuid", "example": "d647642e-09ac-4b34-8acc-ac30f57adc2c"}, "link": {"description": "Link to policy documentation.", "type": "string", "format": "uri", "nullable": true}, "name": {"description": "Policy name.", "type": "string", "example": "Google Workspace Standard Policy"}, "permissions": {"description": "List of permissions included in the policy.", "type": "array", "items": {"type": "string"}, "example": ["https://www.googleapis.com/auth/admin.directory.domain.readonly", "https://www.googleapis.com/auth/admin.directory.user.readonly"]}}, "example": {"compliance_level": "standard", "dlp_enabled": true, "id": "d647642e-09ac-4b34-8acc-ac30f57adc2c", "name": "Google Workspace Standard Policy"}}
```
