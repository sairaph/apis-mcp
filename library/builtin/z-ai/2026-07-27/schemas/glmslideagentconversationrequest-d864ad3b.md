---
title: GlmSlideAgentConversationRequest
page_id: schema-glmslideagentconversationrequest-d864ad3b
path: schemas
source: https://docs.z.ai/openapi.json
source_type: openapi
imported_from: https://docs.z.ai/openapi.json
---

# GlmSlideAgentConversationRequest

```yaml
{"type": "object", "properties": {"agent_id": {"type": "string", "description": "Agent ID"}, "conversation_id": {"type": "string", "description": "Conversation ID"}, "custom_variables": {"type": "object", "description": "Custom variables", "properties": {"include_pdf": {"type": "boolean", "description": "Is export the pdf file"}, "pages": {"type": "array", "description": "Slides Pages", "items": {"type": "object", "properties": {"position": {"type": "number", "description": "Slide Page Position"}, "width": {"type": "number", "description": "Slide Width, unit: pt"}, "height": {"type": "number", "description": "Slide Height, unit: pt"}}}}}}}}
```
