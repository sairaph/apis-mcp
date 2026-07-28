---
title: CreateObservabilityDestinationRequest
page_id: schema-createobservabilitydestinationrequest-f98603f8
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# CreateObservabilityDestinationRequest

```yaml
{"example": {"config": {"baseUrl": "https://us.cloud.langfuse.com", "publicKey": "pk-l...EfGh", "secretKey": "sk-l...AbCd"}, "name": "Production Langfuse", "type": "langfuse"}, "properties": {"api_key_hashes": {"description": "Optional allowlist of OpenRouter API key hashes whose traffic is forwarded. `null` or omitted means all keys. Must contain at least one hash if provided.", "example": null, "items": {"type": "string"}, "minItems": 1, "type": ["array", "null"]}, "config": {"additionalProperties": {}, "description": "Provider-specific configuration. The shape depends on `type` and is validated server-side.", "example": {"baseUrl": "https://us.cloud.langfuse.com", "publicKey": "pk-l...EfGh", "secretKey": "sk-l...AbCd"}, "type": "object"}, "enabled": {"default": true, "description": "Whether this destination should be enabled immediately.", "example": true, "type": "boolean"}, "filter_rules": {"$ref": "#/components/schemas/ObservabilityFilterRulesConfigNullable"}, "name": {"description": "Human-readable name for the destination.", "example": "Production Langfuse", "type": "string"}, "privacy_mode": {"default": false, "description": "When true, request/response bodies are not forwarded — only metadata.", "example": false, "type": "boolean"}, "sampling_rate": {"description": "Sampling rate between 0.0001 and 1 (1 = 100%).", "example": 1, "format": "double", "type": "number"}, "type": {"description": "The destination type. Only stable destination types are accepted.", "enum": ["arize", "braintrust", "clickhouse", "datadog", "grafana", "langfuse", "langsmith", "newrelic", "opik", "otel-collector", "posthog", "ramp", "s3", "sentry", "snowflake", "weave", "webhook"], "example": "langfuse", "type": "string", "x-speakeasy-unknown-values": "allow"}, "workspace_id": {"description": "Optional workspace ID. Defaults to the authenticated entity's default workspace.", "example": "550e8400-e29b-41d4-a716-446655440000", "format": "uuid", "type": "string"}}, "required": ["type", "name", "config"], "type": "object"}
```
