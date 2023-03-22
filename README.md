# **UTIL Send-Telegram**

Utility send-telegram - the utility is designed to send messages, files, and photos to telegram.

## **Installation**
- install [golang](https://go.dev/) 1.16+
- go get github.com/xxandev/util-send-telega
- cd ..../util-send-telega
- make [ build | arm6 | arm7 | arm8 | linux64 | linux32 | win64 | win32 | win64i | win32i ] or go build .

## **Run**
- create config file
```yaml
telegram:
  chat_id: <telegram chat id - int64>
  bot_token: <telegram bot token - string>
```
- run 
```bash
.../send-telegram -m 'test message'

.../send-telegram -d '<path>/document.doc' -m 'test document'

.../send-telegram -p '<path>/image.jpg' -m 'test image'
```
OR
```bash
.../send-telegram \
    -bot '987654321:xxxxxxxxx-xxx-xxxxxxxxxxxx' \
    -id 123456789 \
    -m 'test message'

.../send-telegram \
    -bot '987654321:xxxxxxxxx-xxx-xxxxxxxxxxxx' \
    -id 123456789 \
    -d '<path>/document.doc' \
    -m 'test document'
    
.../send-telegram \
    -bot '987654321:xxxxxxxxx-xxx-xxxxxxxxxxxx' \
    -id 123456789 \
    -p '<path>/image.jpg' \
    -m 'test image'
```