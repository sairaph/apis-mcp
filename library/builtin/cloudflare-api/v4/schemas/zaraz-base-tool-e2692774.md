---
title: zaraz_base-tool
page_id: schema-zaraz-base-tool-e2692774
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zaraz_base-tool

```yaml
{"type": "object", "properties": {"blockingTriggers": {"description": "List of blocking trigger IDs.", "type": "array", "items": {"type": "string", "x-auditable": true}}, "defaultFields": {"description": "Default fields for tool's actions.", "type": "object", "additionalProperties": {"anyOf": [{"type": "string"}, {"type": "boolean"}], "x-auditable": true}}, "defaultPurpose": {"description": "Default consent purpose ID.", "type": "string", "x-auditable": true}, "enabled": {"description": "Whether tool is enabled.", "type": "boolean", "x-auditable": true}, "name": {"description": "Tool's name defined by the user.", "type": "string", "x-auditable": true}, "vendorName": {"description": "Vendor name for TCF compliant consent modal, required for Custom Managed Components and Custom HTML tool with a defaultPurpose assigned.", "type": "string", "x-auditable": true}, "vendorPolicyUrl": {"description": "Vendor's Privacy Policy URL for TCF compliant consent modal, required for Custom Managed Components and Custom HTML tool with a defaultPurpose assigned.", "type": "string", "x-auditable": true}}, "required": ["enabled", "blockingTriggers", "name", "defaultFields"]}
```
