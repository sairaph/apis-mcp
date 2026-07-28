---
title: one_ApplicationDetail
page_id: schema-one-applicationdetail-e170d4c6
path: schemas
description: Full application detail for onboarding UI.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# one_ApplicationDetail

Full application detail for onboarding UI.

```yaml
{"description": "Full application detail for onboarding UI.", "type": "object", "properties": {"auth_methods": {"description": "Available authentication methods.", "type": "array", "items": {"$ref": "#/components/schemas/one_AuthMethod"}}, "category": {"description": "Vendor category.", "type": "string"}, "description": {"description": "Brief description.", "type": "string"}, "display_name": {"description": "Human-readable vendor name.", "type": "string"}, "dlp_enabled": {"description": "Whether DLP scanning is supported.", "type": "boolean"}, "id": {"description": "Vendor identifier.\n\n* `ANTHROPIC` - ANTHROPIC\n* `BITBUCKET` - BITBUCKET\n* `BOX` - BOX\n* `CONFLUENCE` - CONFLUENCE\n* `DROPBOX` - DROPBOX\n* `GITHUB` - GITHUB\n* `GOOGLE_CLOUD_PLATFORM` - GOOGLE_CLOUD_PLATFORM\n* `GOOGLE_WORKSPACE` - GOOGLE_WORKSPACE\n* `JIRA` - JIRA\n* `MICROSOFT_INTERNAL` - MICROSOFT_INTERNAL\n* `OPENAI` - OPENAI\n* `SALESFORCE` - SALESFORCE\n* `SLACK` - SLACK", "type": "string", "enum": ["ANTHROPIC", "BITBUCKET", "BOX", "CONFLUENCE", "DROPBOX", "GITHUB", "GOOGLE_CLOUD_PLATFORM", "GOOGLE_WORKSPACE", "JIRA", "MICROSOFT_INTERNAL", "OPENAI", "SALESFORCE", "SLACK"]}, "instructions": {"description": "Setup instructions for the user.", "type": "string", "nullable": true}, "logo": {"description": "Logo path.", "type": "string", "nullable": true}, "use_cases": {"description": "Use cases with full scope details.", "type": "array", "items": {"$ref": "#/components/schemas/one_UseCaseDetail"}}}, "required": ["auth_methods", "category", "description", "display_name", "dlp_enabled", "id", "instructions", "logo", "use_cases"]}
```
