---
title: Text Tokenizer
page_id: operation-post-paas-v4-tokenizer-0c02ef62
path: operations/untagged
description: '`Tokenizer` is used to split text into `tokens` recognizable by the model and calculate the count. It receives user input text, processes it through the model for tokenization, and finally returns the corresponding `token` count. It is suitable for text length evaluation, model input estimation, dialogue context truncation, cost calculation, etc.'
source: https://docs.z.ai/openapi.json
http_methods:
    - POST
api_endpoints:
    - /paas/v4/tokenizer
source_type: openapi
imported_from: https://docs.z.ai/openapi.json
---

# Text Tokenizer

`POST /paas/v4/tokenizer`

`Tokenizer` is used to split text into `tokens` recognizable by the model and calculate the count. It receives user input text, processes it through the model for tokenization, and finally returns the corresponding `token` count. It is suitable for text length evaluation, model input estimation, dialogue context truncation, cost calculation, etc.

## Definition

```yaml
{"summary": "Text Tokenizer", "description": "`Tokenizer` is used to split text into `tokens` recognizable by the model and calculate the count. It receives user input text, processes it through the model for tokenization, and finally returns the corresponding `token` count. It is suitable for text length evaluation, model input estimation, dialogue context truncation, cost calculation, etc.", "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/TokenizerRequest"}, "examples": {"Text Tokenization Example": {"value": {"model": "glm-4.6", "messages": [{"role": "user", "content": "What opportunities and challenges will the Chinese large model industry face in 2025?"}]}}}}}, "required": true}, "responses": {"200": {"description": "Business processing successful", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/TokenizerResponse"}}}}, "default": {"description": "The request has failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/Error"}}}}}}
```
