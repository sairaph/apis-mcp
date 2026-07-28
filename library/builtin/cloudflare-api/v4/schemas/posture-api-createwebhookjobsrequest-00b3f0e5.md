---
title: posture-api_CreateWebhookJobsRequest
page_id: schema-posture-api-createwebhookjobsrequest-00b3f0e5
path: schemas
description: Request body for creating webhook jobs
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# posture-api_CreateWebhookJobsRequest

Request body for creating webhook jobs

```yaml
{"description": "Request body for creating webhook jobs", "type": "object", "properties": {"finding_instance_ids": {"description": "Array of finding instance IDs to send to the webhooks", "type": "array", "items": {"format": "uuid", "type": "string"}, "example": ["770e8400-e29b-41d4-a716-446655440002", "660e8400-e29b-41d4-a716-446655440001"]}, "webhook_ids": {"description": "Array of webhook IDs to trigger jobs for", "type": "array", "items": {"format": "uuid", "type": "string"}, "example": ["550e8400-e29b-41d4-a716-446655440000", "660e8400-e29b-41d4-a716-446655440001"], "minItems": 1}}, "required": ["webhook_ids", "finding_instance_ids"]}
```
