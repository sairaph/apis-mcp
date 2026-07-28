---
title: posture-api_IntegrationSummary
page_id: schema-posture-api-integrationsummary-bd728936
path: schemas
description: Summary information about an integration.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# posture-api_IntegrationSummary

Summary information about an integration.

```yaml
{"description": "Summary information about an integration.", "type": "object", "properties": {"created": {"description": "When entity was created.", "type": "string", "format": "date-time", "example": "2021-08-10T20:16:11.851451Z", "readOnly": true}, "credential_health_status": {"$ref": "#/components/schemas/posture-api_CredentialHealthStatusEnum"}, "credentials_expiry": {"description": "The date and time when the integration credentials will expire.", "type": "string", "format": "date-time", "example": "2025-03-18T17:25:38.697902Z", "nullable": true}, "id": {"description": "Integration ID.", "type": "string", "format": "uuid", "example": "c416bc38-75dc-425f-ae25-c37b5df5c37f"}, "is_paused": {"description": "Whether the given integration is paused by the user.", "type": "boolean", "example": false, "default": false}, "last_hydrated": {"description": "When were the integration credentials last updated.", "type": "string", "format": "date-time", "example": "2025-03-18T17:25:38.697894Z", "readOnly": true}, "name": {"description": "Name of the integration.", "type": "string", "example": "Example integration", "maxLength": 256}, "permissions": {"description": "The vendor-specific permissions associated with the integration.", "type": "array", "items": {"type": "string"}, "example": ["GroupMember.Read.All", "Group.Read.All"], "readOnly": true}, "policy": {"$ref": "#/components/schemas/posture-api_IntegrationPolicy"}, "status": {"description": "Current status of the integration.", "type": "string", "example": "Healthy", "readOnly": true}, "updated": {"description": "Last entity was updated.", "type": "string", "format": "date-time", "example": "2021-08-10T20:16:11.851451Z", "readOnly": true}, "upgradable": {"description": "Whether the integrations permissions can be updated.", "type": "boolean", "example": false, "readOnly": true}, "upgrade_dismissed": {"description": "UI State as to whether a potential permissions upgrade has been dismissed.", "type": "boolean", "example": false, "default": false}, "vendor": {"$ref": "#/components/schemas/posture-api_Vendor"}, "zt_enrollments": {"description": "Zero Trust products associated with this integration.", "type": "array", "items": {"$ref": "#/components/schemas/posture-api_ZeroTrustProduct"}, "readOnly": true}}, "required": ["created", "last_hydrated", "name", "permissions", "policy", "status", "updated", "upgradable", "vendor", "zt_enrollments"]}
```
