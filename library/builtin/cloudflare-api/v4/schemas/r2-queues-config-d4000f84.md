---
title: r2_queues-config
page_id: schema-r2-queues-config-d4000f84
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# r2_queues-config

```yaml
{"type": "object", "properties": {"queueId": {"description": "Queue ID.", "type": "string", "example": "11111aa1-11aa-111a-a1a1-a1a111a11a11", "x-auditable": true}, "queueName": {"description": "Name of the queue.", "type": "string", "example": "first-queue", "x-auditable": true}, "rules": {"type": "array", "items": {"allOf": [{"$ref": "#/components/schemas/r2_rule"}, {"properties": {"createdAt": {"description": "Timestamp when the rule was created.", "type": "string", "example": "2024-09-19T21:54:48.405Z", "x-auditable": true}, "description": {"description": "A description that can be used to identify the event notification rule after creation.", "type": "string", "example": "Notifications from source bucket to queue", "x-auditable": true}, "ruleId": {"description": "Rule ID.", "type": "string", "example": "11111aa1-11aa-111a-a1a1-a1a111a11a11", "x-auditable": true}}, "type": "object"}]}}}}
```
