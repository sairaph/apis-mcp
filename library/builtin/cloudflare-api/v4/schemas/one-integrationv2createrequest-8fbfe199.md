---
title: one_IntegrationV2CreateRequest
page_id: schema-one-integrationv2createrequest-8fbfe199
path: schemas
description: Serializer for v2 integration create requests.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# one_IntegrationV2CreateRequest

Serializer for v2 integration create requests.

```yaml
{"description": "Serializer for v2 integration create requests.", "type": "object", "properties": {"application": {"description": "Vendor/application slug (e.g., GOOGLE_WORKSPACE).\n\n* `ANTHROPIC` - ANTHROPIC\n* `BITBUCKET` - BITBUCKET\n* `BOX` - BOX\n* `CONFLUENCE` - CONFLUENCE\n* `DROPBOX` - DROPBOX\n* `GITHUB` - GITHUB\n* `GOOGLE_CLOUD_PLATFORM` - GOOGLE_CLOUD_PLATFORM\n* `GOOGLE_WORKSPACE` - GOOGLE_WORKSPACE\n* `JIRA` - JIRA\n* `MICROSOFT_INTERNAL` - MICROSOFT_INTERNAL\n* `OPENAI` - OPENAI\n* `SALESFORCE` - SALESFORCE\n* `SLACK` - SLACK", "type": "string", "enum": ["ANTHROPIC", "BITBUCKET", "BOX", "CONFLUENCE", "DROPBOX", "GITHUB", "GOOGLE_CLOUD_PLATFORM", "GOOGLE_WORKSPACE", "JIRA", "MICROSOFT_INTERNAL", "OPENAI", "SALESFORCE", "SLACK"]}, "auth_method": {"description": "Authentication method slug (uses default if omitted).", "type": "string", "minLength": 1, "nullable": true}, "credentials": {"description": "Credentials for the integration.", "type": "object", "additionalProperties": {}}, "dlp_profiles": {"description": "List of DLP profile IDs to associate.", "type": "array", "items": {"format": "uuid", "type": "string"}, "maxItems": 20}, "name": {"description": "Name of the integration.", "type": "string", "maxLength": 256, "minLength": 1}, "permissions": {"description": "List of permission scopes (uses policy defaults if empty).", "type": "array", "items": {"minLength": 1, "type": "string"}}, "use_cases": {"description": "List of use case or feature slugs to enroll (e.g., ['casb', 'ces', 'auto_remediation']).", "type": "array", "items": {"description": "* `casb` - casb\n* `ces` - ces\n* `auto_remediation` - auto_remediation", "enum": ["casb", "ces", "auto_remediation"], "type": "string"}}}, "required": ["application", "credentials", "name"]}
```
