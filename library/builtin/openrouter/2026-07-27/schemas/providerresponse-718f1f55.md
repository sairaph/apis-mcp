---
title: ProviderResponse
page_id: schema-providerresponse-718f1f55
path: schemas
description: Details of a provider response for a generation attempt
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ProviderResponse

Details of a provider response for a generation attempt

```yaml
{"description": "Details of a provider response for a generation attempt", "example": {"endpoint_id": "ep_abc123", "id": "chatcmpl-abc123", "is_byok": false, "latency": 1200, "model_permaslug": "openai/gpt-4", "provider_name": "OpenAI", "status": 200}, "properties": {"endpoint_id": {"description": "Internal endpoint identifier", "example": "ep_abc123", "type": "string"}, "id": {"description": "Upstream provider response identifier", "example": "chatcmpl-abc123", "type": "string"}, "is_byok": {"description": "Whether the request used a bring-your-own-key", "example": false, "type": "boolean"}, "latency": {"description": "Response latency in milliseconds", "example": 1200, "format": "double", "type": "number"}, "model_permaslug": {"description": "Canonical model slug", "example": "openai/gpt-4", "type": "string"}, "provider_name": {"description": "Name of the provider", "enum": ["AnyScale", "Atoma", "Cent-ML", "CrofAI", "Enfer", "GoPomelo", "HuggingFace", "Hyperbolic", "Hyperbolic 2", "InoCloud", "Kluster", "Lambda", "Lepton", "Lynn 2", "Lynn", "Mancer", "Modal", "Nineteen", "OctoAI", "Recursal", "Reflection", "Replicate", "SambaNova 2", "SF Compute", "Targon", "Together 2", "Ubicloud", "01.AI", "AkashML", "AI21", "AionLabs", "Alibaba", "Ambient", "Baidu", "Amazon Bedrock", "Amazon Nova", "Anthropic", "Arcee AI", "AtlasCloud", "Avian", "Azure", "BaseTen", "BytePlus", "Black Forest Labs", "Cerebras", "Chutes", "Cirrascale", "Claude Platform on AWS", "Clarifai", "Cloudflare", "Cohere", "CoreWeave", "Crucible", "Crusoe", "Darkbloom", "Decart", "Deepgram", "DeepInfra", "DeepSeek", "DekaLLM", "DigitalOcean", "Featherless", "Fireworks", "Fish Audio", "Friendli", "GMICloud", "Google", "Google AI Studio", "Groq", "HeyGen", "Inception", "Inceptron", "InferenceNet", "Ionstream", "Infermatic", "Io Net", "Inferact vLLM", "Inflection", "Liquid", "Mara", "Mancer 2", "Meta", "Minimax", "ModelRun", "Mistral", "Modular", "Moonshot AI", "Morph", "VoyageAI by MongoDB", "NCompass", "Nebius", "Nex AGI", "NextBit", "Novita", "Nvidia", "OpenAI", "OpenInference", "Parasail", "Poolside", "Perceptron", "Perplexity", "Phala", "Recraft", "Reka", "Relace", "Sail Research", "Sakana AI", "SambaNova", "Seed", "SiliconFlow", "Sourceful", "StepFun", "Stealth", "StreamLake", "Switchpoint", "Tencent", "Tenstorrent", "Together", "Upstage", "Venice", "Wafer", "WandB", "Quiver", "Krea", "Runway", "Xiaomi", "xAI", "Z.AI", "FakeProvider"], "example": "OpenAI", "type": "string", "x-speakeasy-unknown-values": "allow"}, "routed_service_tier": {"description": "The service tier this request was routed to (e.g. flex, priority). The tier actually applied and billed is determined by the provider response and may differ.", "enum": ["flex", "priority"], "example": "priority", "type": "string", "x-speakeasy-unknown-values": "allow"}, "status": {"description": "HTTP status code from the provider", "example": 200, "type": ["integer", "null"]}}, "required": ["status"], "type": "object"}
```
