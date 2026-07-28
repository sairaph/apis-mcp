---
title: one_ApplicationList
page_id: schema-one-applicationlist-b2b30d52
path: schemas
description: Application item in list response.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# one_ApplicationList

Application item in list response.

```yaml
{"description": "Application item in list response.", "type": "object", "properties": {"auth_methods": {"description": "Available auth methods.", "type": "array", "items": {"$ref": "#/components/schemas/one_AuthMethodSummary"}}, "category": {"description": "Vendor category (e.g. Productivity, AI).", "type": "string"}, "description": {"description": "Brief description of the integration.", "type": "string"}, "display_name": {"description": "Human-readable vendor name.", "type": "string"}, "dlp_enabled": {"description": "Whether DLP scanning is supported.", "type": "boolean"}, "id": {"description": "Vendor identifier (e.g. microsoft_internal, google_workspace).\n\n* `ANTHROPIC` - ANTHROPIC\n* `BITBUCKET` - BITBUCKET\n* `BOX` - BOX\n* `CONFLUENCE` - CONFLUENCE\n* `DROPBOX` - DROPBOX\n* `GITHUB` - GITHUB\n* `GOOGLE_CLOUD_PLATFORM` - GOOGLE_CLOUD_PLATFORM\n* `GOOGLE_WORKSPACE` - GOOGLE_WORKSPACE\n* `JIRA` - JIRA\n* `MICROSOFT_INTERNAL` - MICROSOFT_INTERNAL\n* `OPENAI` - OPENAI\n* `SALESFORCE` - SALESFORCE\n* `SLACK` - SLACK", "type": "string", "enum": ["ANTHROPIC", "BITBUCKET", "BOX", "CONFLUENCE", "DROPBOX", "GITHUB", "GOOGLE_CLOUD_PLATFORM", "GOOGLE_WORKSPACE", "JIRA", "MICROSOFT_INTERNAL", "OPENAI", "SALESFORCE", "SLACK"]}, "logo": {"description": "Logo path.", "type": "string", "nullable": true}, "permissions": {"description": "All permissions with severity.", "type": "array", "items": {"$ref": "#/components/schemas/one_Permission"}}, "supported_environments": {"description": "Environments this vendor supports (standard, fedramp).", "type": "array", "items": {"type": "string"}}, "use_cases": {"description": "Supported use cases.", "type": "array", "items": {"$ref": "#/components/schemas/one_UseCaseSummary"}}}, "required": ["auth_methods", "category", "description", "display_name", "dlp_enabled", "id", "logo", "permissions", "supported_environments", "use_cases"]}
```
