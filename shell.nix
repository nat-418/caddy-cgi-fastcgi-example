{ pkgs ? import <nixpkgs> {} }:

with pkgs;

mkShell {
  buildInputs = [
    caddy
    go
    just
    wrk
  ];
}
