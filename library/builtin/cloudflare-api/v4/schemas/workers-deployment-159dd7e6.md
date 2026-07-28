---
title: workers_deployment
page_id: schema-workers-deployment-159dd7e6
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_deployment

```yaml
{"type": "object", "properties": {"annotations": {"type": "object", "properties": {"workers/message": {"description": "Human-readable message about the deployment. Truncated to 1000 bytes if longer.", "type": "string", "example": "Deploy bug fix.", "maxLength": 1000, "x-auditable": true}, "workers/triggered_by": {"description": "Operation that triggered the creation of the deployment.", "type": "string", "example": "deployment", "readOnly": true}}}, "author_email": {"type": "string", "format": "email", "readOnly": true, "x-auditable": true}, "created_on": {"type": "string", "format": "date-time", "readOnly": true, "x-auditable": true}, "id": {"type": "string", "format": "uuid", "readOnly": true, "x-auditable": true}, "source": {"type": "string", "example": "api", "readOnly": true, "x-auditable": true}, "strategy": {"type": "string", "enum": ["percentage"], "x-auditable": true}, "versions": {"type": "array", "items": {"properties": {"percentage": {"type": "number", "example": 100, "maximum": 100, "minimum": 0.01, "x-auditable": true}, "version_id": {"type": "string", "format": "uuid", "x-auditable": true}}, "required": ["version_id", "percentage"], "type": "object"}, "x-auditable": true}}, "required": ["id", "source", "strategy", "versions", "created_on"]}
```
