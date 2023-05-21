{ pkgs ? import <nixpkgs> {} }:

with pkgs;

mkShell {
  buildInputs = [
    xcaddy
    go
    just
    wrk
  ];
}
