---
title: posture-api_WebhookInvocationSummary
page_id: schema-posture-api-webhookinvocationsummary-31c302ee
path: schemas
description: Summary of the most recent webhook job invocation for a specific webhook configuration.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# posture-api_WebhookInvocationSummary

Summary of the most recent webhook job invocation for a specific webhook configuration.

```yaml
{"description": "Summary of the most recent webhook job invocation for a specific webhook configuration.", "type": "object", "properties": {"latest_job": {"description": "The most recent webhook job for this webhook configuration.", "type": "object", "properties": {"created_at": {"description": "When the webhook job was created.", "type": "string", "format": "date-time", "example": "2025-03-18T18:30:15.123456Z"}, "id": {"description": "Unique identifier for the webhook job.", "type": "string", "format": "uuid", "example": "123e4567-e89b-12d3-a456-426614174000"}, "stale": {"description": "Whether this webhook job is stale (created before the finding instance's current affliction_date).", "type": "boolean", "example": false}, "status": {"description": "Current status of the webhook job.", "type": "string", "example": "pending", "enum": ["pending", "processing", "completed"]}}, "required": ["id", "status", "created_at", "stale"]}, "webhook_id": {"description": "Unique identifier for the webhook configuration.", "type": "string", "format": "uuid", "example": "550e8400-e29b-41d4-a716-446655440000"}, "webhook_label": {"description": "Account-specified display label for the webhook configuration.", "type": "string", "example": "Send to Gmail"}}, "required": ["webhook_id", "webhook_label", "latest_job"]}
```
