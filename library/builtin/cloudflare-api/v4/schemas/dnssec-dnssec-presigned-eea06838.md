---
title: dnssec_dnssec_presigned
page_id: schema-dnssec-dnssec-presigned-eea06838
path: schemas
description: |-
    If true, allows Cloudflare to transfer in a DNSSEC-signed zone
    including signatures from an external provider, without requiring
    Cloudflare to sign any records on the fly.

    Note that this feature has some limitations.
    See [Cloudflare as Secondary](https://developers.cloudflare.com/dns/zone-setups/zone-transfers/cloudflare-as-secondary/setup/#dnssec) for details.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dnssec_dnssec_presigned

If true, allows Cloudflare to transfer in a DNSSEC-signed zone
including signatures from an external provider, without requiring
Cloudflare to sign any records on the fly.

Note that this feature has some limitations.
See [Cloudflare as Secondary](https://developers.cloudflare.com/dns/zone-setups/zone-transfers/cloudflare-as-secondary/setup/#dnssec) for details.

```yaml
{"description": "If true, allows Cloudflare to transfer in a DNSSEC-signed zone\nincluding signatures from an external provider, without requiring\nCloudflare to sign any records on the fly.\n\nNote that this feature has some limitations.\nSee [Cloudflare as Secondary](https://developers.cloudflare.com/dns/zone-setups/zone-transfers/cloudflare-as-secondary/setup/#dnssec) for details.", "type": "boolean", "example": true, "x-auditable": true}
```
