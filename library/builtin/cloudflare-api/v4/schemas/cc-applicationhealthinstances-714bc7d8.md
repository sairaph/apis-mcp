---
title: cc_ApplicationHealthInstances
page_id: schema-cc-applicationhealthinstances-714bc7d8
path: schemas
description: Shows a count of application instance states.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cc_ApplicationHealthInstances

Shows a count of application instance states.

```yaml
{"description": "Shows a count of application instance states.", "type": "object", "properties": {"active": {"description": "Number of instances with a running container.", "type": "integer"}, "assigned": {"description": "Number of instances assigned to a container, but the container is not yet running.", "type": "integer"}}, "required": ["active", "assigned"]}
```
