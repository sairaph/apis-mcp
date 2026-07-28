import argparse
from http.server import BaseHTTPRequestHandler, ThreadingHTTPServer
from pathlib import Path


class Handler(BaseHTTPRequestHandler):
    def do_GET(self):
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
                    body = source.read_bytes()
                    self.send_response(200)
                    self.send_header("Content-Length", str(len(body)))
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
    server.serve_forever()


if __name__ == "__main__":
    main()
