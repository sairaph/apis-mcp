---
title: one_PatchedUpdateIntegrationV2Request
page_id: schema-one-patchedupdateintegrationv2request-a940b2d8
path: schemas
description: Serializer for v2 integration PATCH requests.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# one_PatchedUpdateIntegrationV2Request

Serializer for v2 integration PATCH requests.

```yaml
{"description": "Serializer for v2 integration PATCH requests.", "type": "object", "properties": {"credentials": {"description": "Partial credential fields to merge with existing.", "type": "object", "additionalProperties": {}}, "dlp_profiles": {"description": "List of DLP profile IDs to associate with the integration.", "type": "array", "items": {"format": "uuid", "type": "string"}, "maxItems": 20}, "name": {"description": "Name of the integration.", "type": "string", "maxLength": 256, "minLength": 1}, "permissions": {"description": "List of permission scopes granted to the integration.", "type": "array", "items": {"minLength": 1, "type": "string"}}, "use_cases": {"description": "List of use case or feature slugs to enroll (e.g., ['casb', 'ces', 'auto_remediation']).", "type": "array", "items": {"description": "* `casb` - casb\n* `ces` - ces\n* `auto_remediation` - auto_remediation", "enum": ["casb", "ces", "auto_remediation"], "type": "string"}}}}
```
