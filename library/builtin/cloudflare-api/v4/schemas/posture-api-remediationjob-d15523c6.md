---
title: posture-api_RemediationJob
page_id: schema-posture-api-remediationjob-d15523c6
path: schemas
description: Information about a remediation job.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# posture-api_RemediationJob

Information about a remediation job.

```yaml
{"description": "Information about a remediation job.", "type": "object", "properties": {"asset": {"$ref": "#/components/schemas/posture-api_RemediationJobAsset"}, "created_at": {"description": "When the remediation job was created.", "type": "string", "format": "date-time", "example": "2025-07-07T18:39:13.123456Z"}, "finding_id": {"description": "Encoded finding ID.", "type": "string", "example": "MDAwMDAwMDAtMDAwMC0wMDAwLTAwMDAtMDAwMDAwMDAwMDAxOjAwMDAwMDAwLTAwMDAtMDAwMC0wMDAwLTAwMDAwMDAwMDAwMgo="}, "finding_instance_id": {"description": "ID of the finding instance being remediated.", "type": "string", "format": "uuid", "example": "3f7b8c9d-6e5a-4f3b-9c2d-1e0a8b7c6d5e"}, "finding_type_id": {"description": "ID of the finding type.", "type": "string", "format": "uuid", "example": "775c5f38-efcf-4b2b-93db-8428979eb6a2"}, "finding_type_name": {"description": "Name of the finding type.", "type": "string", "example": "Microsoft: File publicly accessible with edit access"}, "id": {"description": "Unique identifier for the remediation job.", "type": "string", "format": "uuid", "example": "c416bc38-75db-425f-ae25-c37b5df5c37f"}, "integration_name": {"description": "Name of the integration.", "type": "string", "example": "Microsoft"}, "last_updated": {"description": "When the remediation job was last updated.", "type": "string", "format": "date-time", "example": "2025-07-07T18:39:13.123456Z"}, "remediation_type": {"description": "Type of remediation being performed.", "type": "string", "example": "Remove publicly accessible edit url"}, "status": {"$ref": "#/components/schemas/posture-api_RemediationJobStatusEnum"}, "triggered_by_actor": {"description": "Type of actor that triggered the remediation job. Null on legacy rows created before this column was populated.", "type": "string", "example": "user", "enum": ["user", "account_token", null], "nullable": true}, "triggered_by_id": {"description": "ID of the actor that triggered the job. Meaning depends on triggered_by_actor. Null on legacy rows.", "type": "string", "example": "0123456789abcdef0123456789abcdef", "nullable": true}, "triggered_by_user": {"description": "Email of the user who triggered the remediation. For account-token actors this is the literal \"Account API Token\"; for policy actors this is empty.", "type": "string", "example": "user@example.com"}}, "required": ["id", "status", "finding_instance_id", "finding_type_id", "finding_id", "finding_type_name", "remediation_type", "integration_name", "triggered_by_user", "created_at", "last_updated", "asset"]}
```
