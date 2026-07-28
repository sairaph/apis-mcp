---
title: RetrievalObject
page_id: schema-retrievalobject-194eb52e
path: schemas
source: https://docs.z.ai/openapi.json
source_type: openapi
imported_from: https://docs.z.ai/openapi.json
---

# RetrievalObject

```yaml
{"type": "object", "properties": {"knowledge_id": {"type": "string", "description": "Knowledge base ID, created or obtained from the platform"}, "prompt_template": {"type": "string", "description": "Prompt template for requesting the model, a custom request template containing placeholders `{{ knowledge }}` and `{{ question }}`. Default template: Search for the answer to the question `{{question}}` in the document `{{ knowledge }}`. If an answer is found, respond only using statements from the document; if no answer is found, use your own knowledge to answer and inform the user that the information is not from the document. Do not repeat the question, start the answer directly."}}, "required": ["knowledge_id"]}
```
