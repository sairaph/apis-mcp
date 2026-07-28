---
title: access_policies
page_id: schema-access-policies-1bb8d41e
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_policies

```yaml
{"type": "object", "properties": {"approval_groups": {"$ref": "#/components/schemas/access_approval_groups-2"}, "approval_required": {"$ref": "#/components/schemas/access_approval_required-2"}, "created_at": {"$ref": "#/components/schemas/access_timestamp"}, "decision": {"$ref": "#/components/schemas/access_decision-2"}, "exclude": {"$ref": "#/components/schemas/access_exclude-3"}, "id": {"$ref": "#/components/schemas/access_uuid"}, "include": {"$ref": "#/components/schemas/access_include"}, "isolation_required": {"$ref": "#/components/schemas/access_isolation_required-2"}, "name": {"$ref": "#/components/schemas/access_name-9"}, "precedence": {"$ref": "#/components/schemas/access_precedence-2"}, "purpose_justification_prompt": {"$ref": "#/components/schemas/access_purpose_justification_prompt"}, "purpose_justification_required": {"$ref": "#/components/schemas/access_purpose_justification_required-2"}, "require": {"$ref": "#/components/schemas/access_require-3"}, "updated_at": {"$ref": "#/components/schemas/access_timestamp"}}}
```
