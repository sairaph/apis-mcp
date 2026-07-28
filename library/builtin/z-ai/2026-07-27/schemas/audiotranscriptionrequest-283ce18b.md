---
title: AudioTranscriptionRequest
page_id: schema-audiotranscriptionrequest-283ce18b
path: schemas
source: https://docs.z.ai/openapi.json
source_type: openapi
imported_from: https://docs.z.ai/openapi.json
---

# AudioTranscriptionRequest

```yaml
{"type": "object", "required": ["file", "model"], "properties": {"file": {"type": "string", "format": "binary", "description": "The audio file to be transcribed. Supported audio file formats: `.wav / .mp3`. Specifications: file size ≤ `25 MB`, audio duration ≤ `30 seconds`."}, "file_base64": {"type": "string", "description": "Base64 encoded audio file. Only one of file_base64 or file needs to be provided (if both are provided, file takes precedence)."}, "model": {"type": "string", "description": "The model ID to invoke.", "enum": ["glm-asr-2512"], "default": "glm-asr-2512"}, "prompt": {"type": "string", "description": "In long text scenarios, you can provide previous transcription results as context. Recommended to be less than 8000 characters."}, "hotwords": {"type": "array", "description": "Hotword list to improve recognition accuracy for domain-specific vocabulary. Format example: [\"person_name\",\"place_name\"]. Recommended not to exceed 100 items.", "items": {"type": "string"}, "maxItems": 100}, "stream": {"type": "boolean", "default": false, "description": "This parameter should be set to `false` or omitted when using synchronous calls. It indicates that the model returns all content at once after generating all content. Default is `false`. If set to `true`, the model will return generated content in chunks via standard `Event Stream`. When the `Event Stream` ends, a `data: [DONE]` message will be returned."}, "request_id": {"type": "string", "description": "Passed by the user side, needs to be unique; used to distinguish each request, 6–64 characters. If not provided by the user side, the platform will generate one by default.", "minLength": 6, "maxLength": 64}, "user_id": {"type": "string", "description": "Unique ID for the end user, 6–128 characters. Avoid using sensitive information.", "minLength": 6, "maxLength": 128}}}
```
