---
title: realtimekit_SummarizationConfig
page_id: schema-realtimekit-summarizationconfig-3bfd398a
path: schemas
description: Summary Config
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# realtimekit_SummarizationConfig

Summary Config

```yaml
{"description": "Summary Config", "type": "object", "properties": {"summary_type": {"description": "Defines the style of the summary, such as general, team meeting, or sales call.", "type": "string", "default": "general", "enum": ["general", "team_meeting", "sales_call", "client_check_in", "interview", "daily_standup", "one_on_one_meeting", "lecture", "code_review"]}, "text_format": {"description": "Determines the text format of the summary, such as plain text or markdown.", "type": "string", "default": "markdown", "enum": ["plain_text", "markdown"]}, "word_limit": {"description": "Sets the maximum number of words in the meeting summary.", "type": "integer", "default": 500, "maximum": 1000, "minimum": 150}}, "title": "SummarizationConfig"}
```
