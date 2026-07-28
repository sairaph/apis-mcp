---
title: one_IntegrationV2Detail
page_id: schema-one-integrationv2detail-a8aa4899
path: schemas
description: Serializer for v2 integration detail response with use cases.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# one_IntegrationV2Detail

Serializer for v2 integration detail response with use cases.

```yaml
{"description": "Serializer for v2 integration detail response with use cases.", "type": "object", "properties": {"application": {"type": "object", "additionalProperties": {"nullable": true, "type": "string"}, "readOnly": true}, "auth_method": {"description": "The integration's authentication method.", "type": "object", "additionalProperties": {"type": "string"}, "nullable": true, "readOnly": true}, "authorization_link": {"description": "Authorization link for the integration.", "type": "object", "nullable": true, "properties": {"components": {"additionalProperties": {}, "nullable": true, "type": "object"}, "link": {"type": "string", "nullable": true}}, "readOnly": true, "required": ["components", "link"]}, "created": {"description": "When the integration was created.", "type": "string", "format": "date-time", "readOnly": true}, "credentials_expiry": {"description": "Credentials expiry time.", "type": "string", "format": "date-time", "readOnly": true}, "dlp_profiles": {"description": "DLP Profiles enabled for the integration.", "type": "array", "items": {"format": "uuid", "type": "string"}, "readOnly": true}, "health_details": {"description": "Health details with remediation hints.", "type": "array", "items": {"additionalProperties": {}, "type": "object"}, "readOnly": true}, "id": {"description": "Integration ID.", "type": "string", "format": "uuid", "readOnly": true}, "is_paused": {"description": "Whether the user paused the integration.", "type": "boolean", "readOnly": true}, "last_hydrated": {"description": "Last time the integration was hydrated.", "type": "string", "format": "date-time", "readOnly": true}, "name": {"description": "Name of the integration.", "type": "string", "readOnly": true}, "organization_id": {"description": "Organization ID.", "type": "integer", "readOnly": true}, "status": {"description": "Integration status.", "type": "string"}, "updated": {"description": "When the integration was last updated.", "type": "string", "format": "date-time", "readOnly": true}, "use_cases": {"description": "Use cases enabled for the integration.", "type": "array", "items": {"additionalProperties": {}, "type": "object"}, "readOnly": true}}, "required": ["application", "auth_method", "authorization_link", "created", "credentials_expiry", "dlp_profiles", "health_details", "id", "is_paused", "last_hydrated", "name", "organization_id", "status", "updated", "use_cases"]}
```
