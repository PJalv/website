let
  unstableTarball = fetchTarball
    "https://github.com/NixOS/nixpkgs/archive/nixos-unstable.tar.gz";
  pkgs = import unstableTarball { };
in pkgs.mkShell {
  env = {

  };
  packages = with pkgs; [
    go
    gopls
    templ
    repomix
  ];
}

