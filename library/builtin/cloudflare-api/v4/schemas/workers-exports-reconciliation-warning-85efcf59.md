---
title: workers_exports_reconciliation_warning
page_id: schema-workers-exports-reconciliation-warning-85efcf59
path: schemas
description: |-
    A non-blocking reconciliation warning. Reserved: no scenario
    populates this array today (`code_class_not_in_exports` is
    surfaced as info and `provisioned_class_missing_from_config` is a
    hard error). Clients should still surface any entries that appear.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_exports_reconciliation_warning

A non-blocking reconciliation warning. Reserved: no scenario
populates this array today (`code_class_not_in_exports` is
surfaced as info and `provisioned_class_missing_from_config` is a
hard error). Clients should still surface any entries that appear.

```yaml
{"description": "A non-blocking reconciliation warning. Reserved: no scenario\npopulates this array today (`code_class_not_in_exports` is\nsurfaced as info and `provisioned_class_missing_from_config` is a\nhard error). Clients should still surface any entries that appear.\n", "type": "object", "properties": {"class": {"description": "The class name the warning is about.", "type": "string"}, "message": {"description": "Human-readable explanation of the warning.", "type": "string"}, "namespace_id": {"description": "The provisioned namespace the warning relates to, when applicable.", "type": "string", "format": "uuid"}, "scenario": {"$ref": "#/components/schemas/workers_exports_reconciliation_scenario"}}, "required": ["class", "scenario", "message"]}
```
