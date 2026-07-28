---
title: iam_sso_connector
page_id: schema-iam-sso-connector-8596f974
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_sso_connector

```yaml
{"type": "object", "properties": {"created_on": {"description": "Timestamp for the creation of the SSO connector", "type": "string", "format": "date-time", "example": "2025-01-01T12:21:02.0000Z"}, "email_domain": {"type": "string", "example": "example.com"}, "enabled": {"type": "boolean", "example": false}, "id": {"$ref": "#/components/schemas/iam_sso_connector_identifier"}, "updated_on": {"description": "Timestamp for the last update of the SSO connector", "type": "string", "format": "date-time", "example": "2025-01-01T12:21:02.0000Z"}, "use_fedramp_language": {"$ref": "#/components/schemas/iam_use_fedramp_language"}, "verification": {"$ref": "#/components/schemas/iam_sso_connector_verification_info"}}}
```
