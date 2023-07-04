# Claude-Proxy

## 一键部署
```
s deploy
```

## 使用
```
curl --request POST \
     --url YOUR_PROXY_URL/v1/complete \
     --header 'accept: application/json' \
     --header 'anthropic-version: 2023-06-01' \
     --header 'content-type: application/json' \
     --header 'x-api-key: $ANTHROPIC_API_KEY' \
     --data '
{
  "model": "claude-1",
  "prompt": "\n\nHuman: Hello, world!\n\nAssistant:",
  "max_tokens_to_sample": 256
}
'
```
