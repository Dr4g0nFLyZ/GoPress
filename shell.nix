{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  buildInputs = with pkgs; [
    go
    gopls
    sqlite-interactive
    protobuf
  ];

  shellHook = ''
    echo "Entering Golang development shell with Go version: $(go version)"
    export GOPATH=$HOME/go
  '';
}
