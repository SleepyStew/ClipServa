<div align="center">
  <h1>
    ClipServa
  </h1>
  <p><strong>Instantly serve your system clipboard content over HTTP!</strong></p>
  <p>Smartly serves as JSON if your clipboard holds valid JSON, otherwise serves plain text.</p>
</div>

---

ClipServa is a lightweight, cross-platform command-line tool that starts a local HTTP server to expose the current content of your system clipboard. It's perfect for quick data transfers, local development testing, frontend teams, or any scenario where you need to programmatically access your clipboard's content via HTTP.

## ✨ Features

* **Dynamic Content Type**: Automatically serves content as `application/json` if it's valid JSON, or `text/plain` otherwise.
* **Real-time**: Fetches the latest clipboard content on every API request.
* **Fast**: Starts instantly, parses and serves your clipboard in **~2 miliseconds**.
* **Configurable**: Set host, port, and API endpoint path via command-line flags.
* **CORS Compatible**: Includes `Access-Control-Allow-Origin: *` header for easy integration with web-based frontends and local development.
* **Lightweight & Simple**: Single Go binary with minimal dependencies.
* **Cross-Platform**: Works on Windows, macOS, and Linux (wherever Go and clipboard access are supported).

## 🤔 Why ClipServa?

* **Rapid Prototyping**: Quickly feed data from your clipboard to a local web app or script.
* **Testing API Clients**: Use your clipboard to provide mock JSON responses for HTTP clients you're developing.
* **Simple Data Transfer**: Easily get text or JSON from your desktop to a VM or another device on your local network.
* **Debugging**: Inspect complex JSON data copied from other applications by viewing it formatted in your browser or tools like `jq`.

## 🚀 Installation

1.  Ensure you have Go installed and your Go environment (e.g., `GOPATH`, `GOBIN`) is set up.
2.  Install ClipServa:
    ```bash
    go install github.com/SleepyStew/clipserva@latest
    ```
3.  This will build the `clipserva` binary and place it in your `$GOPATH/bin` or `$GOBIN` directory. Make sure this directory is in your system's `PATH`.

Alternatively, clone the repository and build manually:
```bash
git clone https://github.com/SleepyStew/clipserva.git
cd clipserva
go build -o clipserva .
````

Then you can run `./clipserva` or move it to a directory in your PATH.

## 🗃️ Download (Alternative)

[Latest build for Windows](https://github.com/SleepyStew/ClipServa/releases/)

## 🛠️ Usage

Run ClipServa from your terminal:

```bash
clipserva [flags]
```

Upon starting, ClipServa will log the address and API endpoint it's listening on.

```
   _____ _ _       _____
  / ____| (_)     / ____|
 | |    | |_ _ __| (___   ___ _ ____   ____ _
 | |    | | | '_ \\___ \ / _ \ '__\ \ / / _` |
 | |____| | | |_) |___) |  __/ |   \ V / (_| |
  \_____|_|_| .__/_____/ \___|_|    \_/ \__,_|
            | | By SleepyStew
            |_|

Starting ClipServa on http://localhost:6901/api
Copy some text (including JSON) to your clipboard!
Run with --help to see available options.
```

### Command-Line Flags

ClipServa provides the following command-line flags to customize its behavior:

| Flag        | Description                                        | Default     |
| :---------- | :------------------------------------------------- | :---------- |
| `--host`    | Host address to serve on                           | `localhost` |
| `--port`    | Port to serve on                                   | `6901`      |
| `--api-path`| API endpoint path for serving clipboard content    | `/api`      |
| `--help`    | Show help message                                  |             |

**Example:** To run ClipServa on all network interfaces, on port `8080`, with the API endpoint at `/clipboard`:

```bash
clipserva --host 0.0.0.0 -port 8080 --api-path /clipboard
```

The server would then be accessible at `http://<your-ip>:8080/clipboard`.

### Accessing Clipboard Content

Once ClipServa is running:

1.  **Copy Content**: Copy any text or valid JSON data to your system clipboard.
2.  **Make a GET Request**:
      * **Using a browser**: Navigate to `http://<host>:<port><api-path>` (e.g., `http://localhost:6901/api`).
          * If the content is JSON, your browser might display it formatted or offer to download it.
          * If it's plain text, it will be displayed as such.
      * **Using `curl`**:
        ```bash
        curl http://localhost:6901/api
        ```
        To pretty-print JSON output with `jq`:
        ```bash
        curl -s http://localhost:6901/api | jq
        ```
      * **From your application**: Make an HTTP GET request to the configured endpoint.

The root path (`http://localhost:6901/` by default) displays a simple status message:

```
ClipServa is running! Send a GET request to /api to get current clipboard content (JSON or plain text).
```

## ⚙️ How It Works

ClipServa uses the `github.com/atotto/clipboard` library to read from the system clipboard.
When a GET request is received at the configured API endpoint:

1.  It reads the current content from the clipboard.
2.  It attempts to parse the content as JSON.
      * If successful, it sets the `Content-Type` header to `application/json` and returns the JSON data.
      * If parsing fails, it assumes the content is plain text, sets `Content-Type` to `text/plain; charset=utf-8`, and returns the text.
3.  The `Access-Control-Allow-Origin: *` header is always set to allow cross-origin requests.

An initial check for clipboard accessibility is performed on startup, but the server will attempt to run even if this initial check fails (as clipboard content might become available later).

## 📄 License

Distributed under the MIT License. See `LICENSE` file for more information.

## 🙏 Acknowledgements

  * Relies on the excellent `github.com/atotto/clipboard` library.

-----

Made by SleepyStew.