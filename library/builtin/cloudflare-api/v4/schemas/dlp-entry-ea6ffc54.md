---
title: dlp_Entry
page_id: schema-dlp-entry-ea6ffc54
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_Entry

```yaml
{"oneOf": [{"allOf": [{"$ref": "#/components/schemas/dlp_CustomEntry"}, {"properties": {"type": {"type": "string", "enum": ["custom"]}}, "required": ["type"], "type": "object"}], "title": "Custom Entry"}, {"allOf": [{"$ref": "#/components/schemas/dlp_CustomPromptTopicEntry"}, {"properties": {"type": {"type": "string", "enum": ["custom_prompt_topic"]}}, "required": ["type"], "type": "object"}], "title": "Custom Prompt Topic Entry"}, {"allOf": [{"$ref": "#/components/schemas/dlp_PredefinedEntry"}, {"properties": {"type": {"type": "string", "enum": ["predefined"]}}, "required": ["type"], "type": "object"}], "title": "Predefined Entry"}, {"allOf": [{"$ref": "#/components/schemas/dlp_IntegrationEntry"}, {"properties": {"type": {"type": "string", "enum": ["integration"]}}, "required": ["type"], "type": "object"}], "title": "Integration Entry"}, {"allOf": [{"$ref": "#/components/schemas/dlp_ExactDataEntry"}, {"properties": {"type": {"type": "string", "enum": ["exact_data"]}}, "required": ["type"], "type": "object"}], "title": "Exact Data Entry"}, {"allOf": [{"$ref": "#/components/schemas/dlp_DocumentFingerprintEntry"}, {"properties": {"type": {"type": "string", "enum": ["document_fingerprint"]}}, "required": ["type"], "type": "object"}], "title": "Document Fingerprint Entry"}, {"allOf": [{"$ref": "#/components/schemas/dlp_WordListEntry"}, {"properties": {"type": {"type": "string", "enum": ["word_list"]}}, "required": ["type"], "type": "object"}], "title": "Word List Entry"}]}
```
