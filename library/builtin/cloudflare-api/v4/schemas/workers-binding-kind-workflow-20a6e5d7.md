---
title: workers_binding_kind_workflow
page_id: schema-workers-binding-kind-workflow-20a6e5d7
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_binding_kind_workflow

```yaml
{"type": "object", "properties": {"class_name": {"description": "Class name of the Workflow. Should only be provided if the Workflow belongs to this script.", "type": "string", "example": "my-workflow"}, "name": {"$ref": "#/components/schemas/workers_binding_name"}, "script_name": {"description": "Script name that contains the Workflow. If not provided, defaults to this script name.", "type": "string", "example": "my-workflow", "x-auditable": true}, "type": {"description": "The kind of resource that the binding provides.", "type": "string", "enum": ["workflow"]}, "workflow_name": {"description": "Name of the Workflow to bind to.", "type": "string", "example": "my-workflow"}}, "required": ["name", "type", "workflow_name"]}
```
