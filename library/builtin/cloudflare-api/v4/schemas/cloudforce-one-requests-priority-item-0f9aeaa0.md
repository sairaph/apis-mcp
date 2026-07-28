---
title: cloudforce-one-requests_priority-item
page_id: schema-cloudforce-one-requests-priority-item-0f9aeaa0
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cloudforce-one-requests_priority-item

```yaml
{"type": "object", "properties": {"created": {"description": "Priority creation time.", "allOf": [{"$ref": "#/components/schemas/cloudforce-one-requests_time"}]}, "id": {"$ref": "#/components/schemas/cloudforce-one-requests_uuid"}, "labels": {"$ref": "#/components/schemas/cloudforce-one-requests_labels"}, "priority": {"description": "Priority.", "type": "integer", "example": 1, "x-auditable": true}, "requirement": {"description": "Requirement.", "type": "string", "example": "DoS attacks carried out by CVEs", "x-auditable": true}, "tlp": {"$ref": "#/components/schemas/cloudforce-one-requests_tlp"}, "updated": {"description": "Priority last updated time.", "allOf": [{"$ref": "#/components/schemas/cloudforce-one-requests_time"}]}}, "required": ["id", "created", "updated", "labels", "priority", "requirement", "tlp"], "title": "Priority Item"}
```
