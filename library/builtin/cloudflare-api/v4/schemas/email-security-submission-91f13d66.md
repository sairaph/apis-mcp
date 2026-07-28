---
title: email-security_Submission
page_id: schema-email-security-submission-91f13d66
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email-security_Submission

```yaml
{"type": "object", "properties": {"customer_status": {"$ref": "#/components/schemas/email-security_CustomerStatus"}, "escalated_as": {"$ref": "#/components/schemas/email-security_OptionalSubmissionDisposition"}, "escalated_at": {"type": "string", "format": "date-time", "nullable": true}, "escalated_by": {"type": "string", "nullable": true}, "escalated_submission_id": {"type": "string", "nullable": true}, "original_disposition": {"$ref": "#/components/schemas/email-security_OptionalSubmissionDisposition"}, "original_edf_hash": {"type": "string", "nullable": true}, "original_postfix_id": {"description": "The postfix ID of the original message that was submitted.", "type": "string", "nullable": true}, "outcome": {"type": "string", "nullable": true}, "outcome_disposition": {"$ref": "#/components/schemas/email-security_OptionalSubmissionDisposition"}, "requested_at": {"description": "When the submission was requested (UTC).", "type": "string", "format": "date-time"}, "requested_by": {"type": "string", "nullable": true}, "requested_disposition": {"$ref": "#/components/schemas/email-security_OptionalSubmissionDisposition"}, "requested_ts": {"description": "Deprecated, use `requested_at` instead.", "type": "string", "deprecated": true, "readOnly": true, "x-stainless-deprecation-message": "Use `requested_at` instead."}, "status": {"type": "string", "nullable": true}, "subject": {"type": "string", "nullable": true}, "submission_id": {"type": "string", "x-auditable": true}, "type": {"description": "Indicates whether a team member or an end user created the submission.", "type": "string", "enum": ["Team", "User"], "nullable": true}}, "required": ["submission_id", "requested_at"]}
```
