import argparse
import os
from http.server import BaseHTTPRequestHandler, ThreadingHTTPServer
from pathlib import Path


class Handler(BaseHTTPRequestHandler):
    def do_GET(self):
        if self.path == "/install.ps1":
            body = (self.server.fixture_root / "install.ps1").read_bytes()
            self.send_response(200)
            self.send_header("Content-Type", "text/plain; charset=utf-8")
            self.send_header("Content-Length", str(len(body)))
            self.end_headers()
            self.wfile.write(body)
            return
        if self.path == "/releases/latest":
            self.send_response(302)
            self.send_header("Location", "/releases/tag/v1.2.3")
            self.end_headers()
            return
        if self.path == "/releases/tag/v1.2.3":
            self.send_response(200)
            self.end_headers()
            self.wfile.write(b"fixture release")
            return

        prefix = "/releases/download/v1.2.3/"
        if self.path.startswith(prefix):
            name = self.path[len(prefix) :]
            if "/" not in name:
                source = self.server.fixture_root / name
                if source.is_file():
                    with self.server.event_log.open("a", encoding="utf-8") as log:
                        log.write("download\n")
                    body = source.read_bytes()
                    mode = self.server.download_mode_file.read_text(
                        encoding="utf-8"
                    ).strip()
                    if mode == "fail":
                        self.send_response(503)
                        self.end_headers()
                        return
                    if mode == "zero":
                        body = b""
                    self.send_response(200)
                    declared_length = len(body)
                    if mode == "truncated":
                        declared_length = len(body)
                        body = body[: max(1, len(body) // 2)]
                        self.close_connection = True
                    self.send_header("Content-Length", str(declared_length))
                    self.end_headers()
                    self.wfile.write(body)
                    return

        self.send_response(404)
        self.end_headers()

    def log_message(self, _format, *_args):
        pass


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("--root", required=True)
    parser.add_argument("--port", required=True, type=int)
    args = parser.parse_args()

    server = ThreadingHTTPServer(("127.0.0.1", args.port), Handler)
    server.fixture_root = Path(args.root)
    server.event_log = Path(os.environ["INSTALLER_EVENT_LOG"])
    server.download_mode_file = Path(
        os.environ["INSTALLER_DOWNLOAD_MODE_FILE"]
    )
    server.serve_forever()


if __name__ == "__main__":
    main()
