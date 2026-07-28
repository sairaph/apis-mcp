---
title: Create transcription
page_id: operation-post-audio-transcriptions-54e94493
path: operations/stt
description: Transcribes audio into text. Accepts base64-encoded audio input as JSON or an OpenAI-style multipart/form-data file upload, and returns the transcribed text.
source: https://openrouter.ai/openapi.json
http_methods:
    - POST
api_endpoints:
    - /audio/transcriptions
operation_ids:
    - createAudioTranscriptions
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# Create transcription

`POST /audio/transcriptions`

Operation ID: `createAudioTranscriptions`

Transcribes audio into text. Accepts base64-encoded audio input as JSON or an OpenAI-style multipart/form-data file upload, and returns the transcribed text.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"description": "Transcribes audio into text. Accepts base64-encoded audio input as JSON or an OpenAI-style multipart/form-data file upload, and returns the transcribed text.", "operationId": "createAudioTranscriptions", "requestBody": {"content": {"application/json": {"example": {"input_audio": {"data": "UklGRiQA...", "format": "wav"}, "language": "en", "model": "openai/whisper-large-v3"}, "schema": {"$ref": "#/components/schemas/STTRequest"}}, "multipart/form-data": {"example": {"file": "audio.wav", "language": "en", "model": "openai/whisper-large-v3"}, "schema": {"properties": {"file": {"description": "The audio file to transcribe. The format is derived from the filename extension or the file part content type. Max 25 MB; send larger files as base64 JSON via input_audio.", "format": "binary", "type": "string"}, "language": {"description": "The language of the input audio (ISO-639-1).", "type": "string"}, "model": {"description": "The model to use for transcription.", "type": "string"}, "response_format": {"description": "The response format. \"json\" (default) returns { text, usage }; \"verbose_json\" additionally returns task, language, duration, and segment-level timestamps (OpenAI-compatible providers only).", "enum": ["json", "verbose_json"], "type": "string", "x-speakeasy-unknown-values": "allow"}, "temperature": {"description": "The sampling temperature.", "type": "number"}, "timestamp_granularities[]": {"description": "Timestamp detail levels to include when response_format is \"verbose_json\". \"word\" additionally returns word-level timestamps in the words array.", "items": {"enum": ["word", "segment"], "type": "string", "x-speakeasy-unknown-values": "allow"}, "type": "array"}}, "required": ["file", "model"], "type": "object"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"example": {"text": "Hello, this is a test of OpenAI speech-to-text transcription.", "usage": {"cost": 0.000508, "input_tokens": 83, "output_tokens": 30, "seconds": 9.2, "total_tokens": 113}}, "schema": {"$ref": "#/components/schemas/STTResponse"}}}, "description": "Transcription result"}, "400": {"content": {"application/json": {"example": {"error": {"code": 400, "message": "Invalid request parameters"}}, "schema": {"$ref": "#/components/schemas/BadRequestResponse"}}}, "description": "Bad Request - Invalid request parameters or malformed input"}, "401": {"content": {"application/json": {"example": {"error": {"code": 401, "message": "Missing Authentication header"}}, "schema": {"$ref": "#/components/schemas/UnauthorizedResponse"}}}, "description": "Unauthorized - Authentication required or invalid credentials"}, "402": {"content": {"application/json": {"example": {"error": {"code": 402, "message": "Insufficient credits. Add more using https://openrouter.ai/credits"}}, "schema": {"$ref": "#/components/schemas/PaymentRequiredResponse"}}}, "description": "Payment Required - Insufficient credits or quota to complete request"}, "404": {"content": {"application/json": {"example": {"error": {"code": 404, "message": "Resource not found"}}, "schema": {"$ref": "#/components/schemas/NotFoundResponse"}}}, "description": "Not Found - Resource does not exist"}, "429": {"content": {"application/json": {"example": {"error": {"code": 429, "message": "Rate limit exceeded"}}, "schema": {"$ref": "#/components/schemas/TooManyRequestsResponse"}}}, "description": "Too Many Requests - Rate limit exceeded"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}, "502": {"content": {"application/json": {"example": {"error": {"code": 502, "message": "Provider returned error"}}, "schema": {"$ref": "#/components/schemas/BadGatewayResponse"}}}, "description": "Bad Gateway - Provider/upstream API failure"}, "503": {"content": {"application/json": {"example": {"error": {"code": 503, "message": "Service temporarily unavailable"}}, "schema": {"$ref": "#/components/schemas/ServiceUnavailableResponse"}}}, "description": "Service Unavailable - Service temporarily unavailable"}, "524": {"content": {"application/json": {"example": {"error": {"code": 524, "message": "Request timed out. Please try again later."}}, "schema": {"$ref": "#/components/schemas/EdgeNetworkTimeoutResponse"}}}, "description": "Infrastructure Timeout - Provider request timed out at edge network"}, "529": {"content": {"application/json": {"example": {"error": {"code": 529, "message": "Provider returned error"}}, "schema": {"$ref": "#/components/schemas/ProviderOverloadedResponse"}}}, "description": "Provider Overloaded - Provider is temporarily overloaded"}}, "summary": "Create transcription", "tags": ["STT"], "x-speakeasy-max-method-params": 1, "x-speakeasy-name-override": "createTranscription"}
```
