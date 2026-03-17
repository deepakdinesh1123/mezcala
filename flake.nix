{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-25.11";
    flake-utils.url = "github:numtide/flake-utils";
  };

  description = "Mezcala";

  outputs = { self, nixpkgs, flake-utils, ... }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        migrate-overlay = import ./hack/migrate-overlay.nix;
        pkgs = import nixpkgs { inherit system; overlays = [ migrate-overlay ]; };
        arch = builtins.head (builtins.match "^([^-]+)-.*" system);
      in
      rec {

        # Common project wide dependencies
        mezDeps = with pkgs; [
          just
        ];

        backendInputs = mezDeps ++ import ./pkgs/backend/shell.nix { inherit pkgs; };
        frontendInputs = mezDeps ++ import ./pkgs/mezui/shell.nix { inherit pkgs; };

        devShells = {
          default = pkgs.mkShell {
            buildInputs = backendInputs;
          };
          backend = pkgs.mkShell {
            buildInputs = backendInputs;
          };
          frontend = pkgs.mkShell {
            buildInputs = frontendInputs;
          };
        };
      }
    );
}