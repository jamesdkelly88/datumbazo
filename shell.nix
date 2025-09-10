{ pkgs ? import <nixpkgs> {} }:pkgs.mkShell {
  packages = with pkgs; [
    gh
    go
    go-task
  ];
}