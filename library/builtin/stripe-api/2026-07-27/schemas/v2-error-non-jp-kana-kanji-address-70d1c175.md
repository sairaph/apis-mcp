---
title: v2.error.non_jp_kana_kanji_address
page_id: schema-v2-error-non-jp-kana-kanji-address-70d1c175
path: schemas
description: Kana Kanji script addresses must have JP country.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# v2.error.non_jp_kana_kanji_address

Kana Kanji script addresses must have JP country.

```yaml
{"required": ["error"], "type": "object", "properties": {"error": {"required": ["code", "message", "user_message"], "type": "object", "properties": {"code": {"type": "string", "description": "Short code to identify the error, should not be handled programmatically", "enum": ["non_jp_kana_kanji_address"]}, "message": {"type": "string", "description": "A human-readable message providing more details about the error"}, "user_message": {"type": "string", "description": "A user-friendly message that can be shown to end-users"}}, "description": "Information about the error that occurred"}}, "description": "Kana Kanji script addresses must have JP country."}
```
