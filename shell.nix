# shell.nix
{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
   name = "go-dev";
  
   nativeBuildInputs = [
      pkgs.pkg-config
   ];

   buildInputs = with pkgs; [
      go
      gnumake 
  ];

   shellHook = ''
      echo "Entering Golang development shell ($(go version))"
      export GOPATH=$(pwd)/.go-packages
      mkdir -p $GOPATH
   '';
}
