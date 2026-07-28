---
title: cloudforce-one-requests_priority-edit
page_id: schema-cloudforce-one-requests-priority-edit-e575496d
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cloudforce-one-requests_priority-edit

```yaml
{"type": "object", "properties": {"labels": {"$ref": "#/components/schemas/cloudforce-one-requests_labels"}, "priority": {"description": "Priority.", "type": "integer", "example": 1, "x-auditable": true}, "requirement": {"description": "Requirement.", "type": "string", "example": "DoS attacks carried out by CVEs", "x-auditable": true}, "tlp": {"$ref": "#/components/schemas/cloudforce-one-requests_tlp"}}, "required": ["labels", "priority", "requirement", "tlp"], "title": "Priority Editable Attributes"}
```
