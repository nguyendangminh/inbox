# inbox
Inbox - Tool for downloading messages in Facebook Page


### Bug
- Download failed when token is expired
2016/10/25 11:08:53 Request failed: status code 400 - {"error":{"message":"Error validating access token: Session has expired on Monday, 24-Oct-16 21:00:00 PDT. The current time is Monday, 24-Oct-16 21:08:52 PDT.","type":"OAuthException","code":190,"error_subcode":463,"fbtrace_id":"FWyaAGXtoZg"}}
https://developers.facebook.com/docs/facebook-login/access-tokens#pagetokens
https://developers.facebook.com/docs/facebook-login/access-tokens/expiration-and-extension


### TODO
- All logs are showed on web page (use SSE)
- Allow user choose output directory
- Allow user pick one of his existing page
- Distribute this software to all platforms
- Allow user download messages in a time range
- Retry or log ID of messages that failed to download