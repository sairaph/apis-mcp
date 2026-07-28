---
title: posture-api_WebhookJobFailure
page_id: schema-posture-api-webhookjobfailure-cd65a283
path: schemas
description: Information about a failed webhook job creation.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# posture-api_WebhookJobFailure

Information about a failed webhook job creation.

```yaml
{"description": "Information about a failed webhook job creation.", "type": "object", "properties": {"error": {"description": "Error message describing the failure.", "type": "string", "example": "Failed to create webhook job"}, "finding_instance_id": {"description": "ID of the finding instance that failed to create a webhook job.", "type": "string", "format": "uuid", "example": "2e6b4c8a-9d1f-4e3b-8c7a-5f9e2d1a6b4c"}, "webhook_id": {"description": "ID of the webhook configuration.", "type": "string", "format": "uuid", "example": "a1b2c3d4-e5f6-7890-abcd-ef1234567890"}}, "required": ["webhook_id", "finding_instance_id", "error"]}
```
