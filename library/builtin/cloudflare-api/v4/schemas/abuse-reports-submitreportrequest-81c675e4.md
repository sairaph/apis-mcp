---
title: abuse-reports_SubmitReportRequest
page_id: schema-abuse-reports-submitreportrequest-81c675e4
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# abuse-reports_SubmitReportRequest

```yaml
{"discriminator": {"mapping": {"abuse_children": "#/components/schemas/abuse-reports_CSAMReport", "abuse_dmca": "#/components/schemas/abuse-reports_DMCAReport", "abuse_general": "#/components/schemas/abuse-reports_GeneralReport", "abuse_ncsei": "#/components/schemas/abuse-reports_NCSEIReport", "abuse_phishing": "#/components/schemas/abuse-reports_PhishingReport", "abuse_registrar_whois": "#/components/schemas/abuse-reports_RegistrarWhoisReport", "abuse_threat": "#/components/schemas/abuse-reports_ThreatReport", "abuse_trademark": "#/components/schemas/abuse-reports_TrademarkReport"}, "propertyName": "act"}, "oneOf": [{"$ref": "#/components/schemas/abuse-reports_DMCAReport"}, {"$ref": "#/components/schemas/abuse-reports_TrademarkReport"}, {"$ref": "#/components/schemas/abuse-reports_GeneralReport"}, {"$ref": "#/components/schemas/abuse-reports_PhishingReport"}, {"$ref": "#/components/schemas/abuse-reports_CSAMReport"}, {"$ref": "#/components/schemas/abuse-reports_ThreatReport"}, {"$ref": "#/components/schemas/abuse-reports_RegistrarWhoisReport"}, {"$ref": "#/components/schemas/abuse-reports_NCSEIReport"}]}
```
