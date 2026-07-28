---
title: ListEndpointsResponse
page_id: schema-listendpointsresponse-69e16a5d
path: schemas
description: List of available endpoints for a model
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ListEndpointsResponse

List of available endpoints for a model

```yaml
{"description": "List of available endpoints for a model", "example": {"architecture": {"input_modalities": ["text"], "instruct_type": "chatml", "modality": "text->text", "output_modalities": ["text"], "tokenizer": "GPT"}, "created": 1692901234, "description": "GPT-4 is a large multimodal model that can solve difficult problems with greater accuracy.", "endpoints": [{"context_length": 8192, "latency_last_30m": {"p50": 0.25, "p75": 0.35, "p90": 0.48, "p99": 0.85}, "max_completion_tokens": 4096, "max_prompt_tokens": 8192, "model_id": "openai/gpt-4", "model_name": "GPT-4", "name": "OpenAI: GPT-4", "pricing": {"completion": "0.00006", "image": "0", "prompt": "0.00003", "request": "0"}, "provider_name": "OpenAI", "quantization": "fp16", "status": 0, "supported_parameters": ["temperature", "top_p", "max_tokens", "frequency_penalty", "presence_penalty"], "supports_implicit_caching": true, "tag": "openai", "throughput_last_30m": {"p50": 45.2, "p75": 38.5, "p90": 28.3, "p99": 15.1}, "uptime_last_1d": 99.8, "uptime_last_30m": 99.5, "uptime_last_5m": 100}], "id": "openai/gpt-4", "name": "GPT-4"}, "properties": {"architecture": {"allOf": [{"$ref": "#/components/schemas/ModelArchitecture"}, {"properties": {"input_modalities": {"description": "Supported input modalities", "items": {"$ref": "#/components/schemas/InputModality"}, "type": "array"}, "instruct_type": {"$ref": "#/components/schemas/InstructType"}, "modality": {"description": "Primary modality of the model", "example": "text", "type": ["string", "null"]}, "output_modalities": {"description": "Supported output modalities", "items": {"$ref": "#/components/schemas/OutputModality"}, "type": "array"}, "tokenizer": {"anyOf": [{"$ref": "#/components/schemas/ModelGroup"}, {"type": "null"}]}}, "required": ["tokenizer", "instruct_type", "modality", "input_modalities", "output_modalities"]}]}, "created": {"description": "Unix timestamp of when the model was created", "example": 1692901234, "type": "integer"}, "description": {"description": "Description of the model", "example": "GPT-4 is a large multimodal model that can solve difficult problems with greater accuracy.", "type": "string"}, "endpoints": {"description": "List of available endpoints for this model", "items": {"$ref": "#/components/schemas/PublicEndpoint"}, "type": "array"}, "id": {"description": "Unique identifier for the model", "example": "openai/gpt-4", "type": "string"}, "name": {"description": "Display name of the model", "example": "GPT-4", "type": "string"}}, "required": ["id", "name", "created", "description", "architecture", "endpoints"], "type": "object"}
```
