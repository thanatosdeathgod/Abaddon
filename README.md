# Abaddon

**Abaddon** is a powerful raid bot developed in Go, designed for automating mass actions on Discord servers. It includes features like channel deletion, user removal, and server takeover. **Use with caution and only in environments where you have permission.**

## Features

- **Channel Deletion:** Removes all channels in the server.
- **User Removal:** Kicks all members from the server.
- **Server Takeover:** Creates new channels and sends messages.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/abaddon.git
   cd abaddon
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Build the bot:

   ```bash
   go build -o abaddon main.go
   ```

## Usage

1. Replace the placeholder `BOT_TOKEN_HERE` in `main.go` with your Discord bot token.

2. Run the bot:

   ```bash
   ./abaddon
   ```

3. Invite the bot to your Discord server.

4. Use the commands to initiate actions:
   - `/fuckserver` - Deletes all channels.
   - `/conquer` - Takes over the server by creating channels and sending messages.
   - `/kickall` - Kicks all members from the server.

## Contributing

Feel free to fork this project and submit pull requests. Contributions are welcome!
