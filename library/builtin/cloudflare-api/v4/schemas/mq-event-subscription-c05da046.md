---
title: mq_event-subscription
page_id: schema-mq-event-subscription-c05da046
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mq_event-subscription

```yaml
{"type": "object", "properties": {"created_at": {"description": "When the subscription was created", "type": "string", "format": "date-time", "x-auditable": true}, "destination": {"$ref": "#/components/schemas/mq_event-destination"}, "enabled": {"description": "Whether the subscription is active", "type": "boolean", "x-auditable": true}, "events": {"description": "List of event types this subscription handles", "type": "array", "items": {"type": "string"}, "minItems": 1, "x-auditable": true}, "id": {"description": "Unique identifier for the subscription", "type": "string", "x-auditable": true}, "modified_at": {"description": "When the subscription was last modified", "type": "string", "format": "date-time", "x-auditable": true}, "name": {"description": "Name of the subscription", "type": "string", "x-auditable": true}, "source": {"$ref": "#/components/schemas/mq_event-source"}}, "required": ["id", "created_at", "modified_at", "name", "enabled", "source", "destination", "events"]}
```
