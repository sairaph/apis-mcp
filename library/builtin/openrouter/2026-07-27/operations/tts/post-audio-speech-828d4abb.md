---
title: Create speech
page_id: operation-post-audio-speech-5ec0b642
path: operations/tts
description: Synthesizes audio from the input text. Returns a raw audio bytestream in the requested format (e.g. mp3, pcm, wav).
source: https://openrouter.ai/openapi.json
http_methods:
    - POST
api_endpoints:
    - /audio/speech
operation_ids:
    - createAudioSpeech
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# Create speech

`POST /audio/speech`

Operation ID: `createAudioSpeech`

Synthesizes audio from the input text. Returns a raw audio bytestream in the requested format (e.g. mp3, pcm, wav).

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"description": "Synthesizes audio from the input text. Returns a raw audio bytestream in the requested format (e.g. mp3, pcm, wav).", "operationId": "createAudioSpeech", "requestBody": {"content": {"application/json": {"example": {"input": "Hello world", "model": "mistralai/voxtral-mini-tts-2603", "response_format": "pcm", "speed": 1, "voice": "en_paul_neutral"}, "schema": {"$ref": "#/components/schemas/SpeechRequest"}}}, "required": true}, "responses": {"200": {"content": {"audio/*": {"schema": {"description": "Raw audio bytestream. Content-Type varies by requested format (audio/mpeg for mp3, audio/pcm for pcm — 16-bit little-endian).", "example": "<binary audio data>", "format": "binary", "type": "string"}}}, "description": "Audio bytes stream"}, "400": {"content": {"application/json": {"example": {"error": {"code": 400, "message": "Invalid request parameters"}}, "schema": {"$ref": "#/components/schemas/BadRequestResponse"}}}, "description": "Bad Request - Invalid request parameters or malformed input"}, "401": {"content": {"application/json": {"example": {"error": {"code": 401, "message": "Missing Authentication header"}}, "schema": {"$ref": "#/components/schemas/UnauthorizedResponse"}}}, "description": "Unauthorized - Authentication required or invalid credentials"}, "402": {"content": {"application/json": {"example": {"error": {"code": 402, "message": "Insufficient credits. Add more using https://openrouter.ai/credits"}}, "schema": {"$ref": "#/components/schemas/PaymentRequiredResponse"}}}, "description": "Payment Required - Insufficient credits or quota to complete request"}, "404": {"content": {"application/json": {"example": {"error": {"code": 404, "message": "Resource not found"}}, "schema": {"$ref": "#/components/schemas/NotFoundResponse"}}}, "description": "Not Found - Resource does not exist"}, "429": {"content": {"application/json": {"example": {"error": {"code": 429, "message": "Rate limit exceeded"}}, "schema": {"$ref": "#/components/schemas/TooManyRequestsResponse"}}}, "description": "Too Many Requests - Rate limit exceeded"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}, "502": {"content": {"application/json": {"example": {"error": {"code": 502, "message": "Provider returned error"}}, "schema": {"$ref": "#/components/schemas/BadGatewayResponse"}}}, "description": "Bad Gateway - Provider/upstream API failure"}, "503": {"content": {"application/json": {"example": {"error": {"code": 503, "message": "Service temporarily unavailable"}}, "schema": {"$ref": "#/components/schemas/ServiceUnavailableResponse"}}}, "description": "Service Unavailable - Service temporarily unavailable"}, "524": {"content": {"application/json": {"example": {"error": {"code": 524, "message": "Request timed out. Please try again later."}}, "schema": {"$ref": "#/components/schemas/EdgeNetworkTimeoutResponse"}}}, "description": "Infrastructure Timeout - Provider request timed out at edge network"}, "529": {"content": {"application/json": {"example": {"error": {"code": 529, "message": "Provider returned error"}}, "schema": {"$ref": "#/components/schemas/ProviderOverloadedResponse"}}}, "description": "Provider Overloaded - Provider is temporarily overloaded"}}, "summary": "Create speech", "tags": ["TTS"], "x-speakeasy-max-method-params": 1, "x-speakeasy-name-override": "createSpeech"}
```
