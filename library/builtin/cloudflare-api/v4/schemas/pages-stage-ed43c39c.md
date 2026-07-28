---
title: pages_stage
page_id: schema-pages-stage-ed43c39c
path: schemas
description: The status of the deployment.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# pages_stage

The status of the deployment.

```yaml
{"description": "The status of the deployment.", "type": "object", "properties": {"ended_on": {"description": "When the stage ended.", "type": "string", "format": "date-time", "example": "2021-03-09T00:58:59.045655", "nullable": true, "readOnly": true, "x-auditable": true}, "name": {"description": "The current build stage.", "type": "string", "example": "deploy", "enum": ["queued", "initialize", "clone_repo", "build", "deploy"], "x-auditable": true}, "started_on": {"description": "When the stage started.", "type": "string", "format": "date-time", "example": "2021-03-09T00:55:03.923456Z", "nullable": true, "readOnly": true, "x-auditable": true}, "status": {"description": "State of the current stage.", "type": "string", "example": "success", "enum": ["success", "idle", "active", "failure", "canceled"], "readOnly": true, "x-auditable": true}}, "readOnly": true, "required": ["name", "started_on", "ended_on", "status"]}
```
