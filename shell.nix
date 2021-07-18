let
  # Pinned nixpkgs, deterministic. version 21.05 Last updated: 2021/07/18.
  pkgs = import (fetchTarball("https://github.com/NixOS/nixpkgs/archive/7e9b0dff974c89e070da1ad85713ff3c20b0ca97.tar.gz")) {};

  # Rolling updates, not deterministic.
  #pkgs = import (fetchTarball("channel:nixpkgs-unstable")) {};

  yggmail = pkgs.callPackage ./default.nix {};

  claws-wrapped = pkgs.symlinkJoin {
    name = "claws-wrapped";
    paths = [ pkgs.claws-mail ];
    buildInputs = [ pkgs.makeWrapper ];
    postBuild = ''
      wrapProgram $out/bin/claws-mail --add-flags "--alternate-config-dir ./claws-mail"
    '';
  };

in pkgs.mkShell {
  buildInputs = [ yggmail claws-wrapped ];
}
