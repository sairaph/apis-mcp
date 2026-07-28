---
title: posture-api_FindingInstance
page_id: schema-posture-api-findinginstance-913319a1
path: schemas
description: A specific instance of a security finding. In the API interface, we refer to the 'finding' table in our DB as finding instances, optimized for the p99 use case.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# posture-api_FindingInstance

A specific instance of a security finding. In the API interface, we refer to the 'finding' table in our DB as finding instances, optimized for the p99 use case.

```yaml
{"description": "A specific instance of a security finding. In the API interface, we refer to the 'finding' table in our DB as finding instances, optimized for the p99 use case.", "type": "object", "properties": {"affliction_date": {"description": "When this specific instance was identified.", "type": "string", "format": "date-time", "example": "2025-03-18T17:25:38.700541Z"}, "asset": {"$ref": "#/components/schemas/posture-api_Asset"}, "dlp_contexts": {"description": "DLP context information if this is a content finding.", "type": "array", "items": {"$ref": "#/components/schemas/posture-api_DlpContext"}, "readOnly": true}, "id": {"description": "Unique identifier for the finding instance.", "type": "string", "format": "uuid", "example": "497f6eca-6276-4993-bfeb-53cbbbba6f08"}, "is_archived": {"description": "Whether this finding instance has been archived.", "type": "boolean", "example": false, "default": false}, "remediations": {"description": "A list of the 10 most recent remediation jobs for this finding instance, ordered by creation time (most recent first). The 'stale' field indicates whether the remediation job was created before the finding instance's affliction_date (true) or after it (false). If there has never been a remediation job for this finding instance, this field will be an empty array.", "type": "array", "items": {"$ref": "#/components/schemas/posture-api_RemediationJobSummary"}}, "webhooks": {"description": "The most recent webhook job invocation for each webhook configuration associated with this finding instance. Each entry represents the latest job (any status) per webhook config. The 'stale' field indicates whether the job was invoked before the finding instance's current affliction_date. If no webhook jobs have been created, this field will be an empty array.", "type": "array", "items": {"$ref": "#/components/schemas/posture-api_WebhookInvocationSummary"}}}, "required": ["affliction_date", "asset", "dlp_contexts", "remediations", "webhooks"]}
```
