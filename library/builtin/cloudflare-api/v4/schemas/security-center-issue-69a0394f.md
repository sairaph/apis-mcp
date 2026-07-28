---
title: security-center_issue
page_id: schema-security-center-issue-69a0394f
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# security-center_issue

```yaml
{"type": "object", "properties": {"dismissed": {"type": "boolean", "example": false, "x-auditable": true}, "has_extended_context": {"description": "Indicates whether the insight has a large payload that requires fetching via the context endpoint.", "type": "boolean", "example": false, "x-auditable": true}, "id": {"type": "string", "x-auditable": true}, "issue_class": {"$ref": "#/components/schemas/security-center_issueClass"}, "issue_type": {"$ref": "#/components/schemas/security-center_issueType"}, "payload": {"type": "object", "properties": {"detection_method": {"description": "Describes the method used to detect insight.", "type": "string", "example": "We detected security rules referencing multiple IP addresses directly in the rules.", "x-auditable": true}, "zone_tag": {"type": "string", "x-auditable": true}}}, "resolve_link": {"type": "string", "x-auditable": true}, "resolve_text": {"type": "string", "x-auditable": true}, "severity": {"type": "string", "enum": ["Low", "Moderate", "Critical"], "x-auditable": true}, "since": {"type": "string", "format": "date-time", "x-auditable": true}, "status": {"description": "The current status of the insight.", "type": "string", "example": "active", "enum": ["active", "resolved"], "x-auditable": true}, "subject": {"$ref": "#/components/schemas/security-center_subject"}, "timestamp": {"type": "string", "format": "date-time", "x-auditable": true}, "user_classification": {"$ref": "#/components/schemas/security-center_userClassification"}}}
```
