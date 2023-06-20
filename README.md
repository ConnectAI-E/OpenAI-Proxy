# Openai-Proxy

## 一键部署
```
s deploy
```

## 使用
```
curl --location 'YOUR_PROXY_URL/v1/chat/completions' \
--header 'Authorization: Bearer YOUR_AUTHORIZATION' \
--header 'Content-Type: application/json' \
--data '{
    "max_tokens": 250,
    "model": "gpt-3.5-turbo",
    "messages": [
        {
            "role": "user",
            "content": "Hello!"
        }
    ]
}'
```
