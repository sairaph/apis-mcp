---
title: posture-api_WebhookJobParameters
page_id: schema-posture-api-webhookjobparameters-9ffa4435
path: schemas
description: Parameters for a webhook job.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# posture-api_WebhookJobParameters

Parameters for a webhook job.

```yaml
{"description": "Parameters for a webhook job.", "type": "object", "properties": {"finding_instance_id": {"description": "ID of the finding instance.", "type": "string", "format": "uuid", "example": "3f7b8c9d-6e5a-4f3b-9c2d-1e0a8b7c6d5e"}}, "required": ["finding_instance_id"]}
```
