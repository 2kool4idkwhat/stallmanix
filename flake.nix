{
  inputs.nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";

  outputs = { self, nixpkgs }: let
    supportedSystems = ["x86_64-linux" "aarch64-linux"];

    forAllSystems = nixpkgs.lib.genAttrs supportedSystems;
    nixpkgsFor = forAllSystems (system: import nixpkgs {inherit system;});
  in {
    packages = forAllSystems (system: let
        pkgs = nixpkgsFor.${system};
      in {
        default = pkgs.callPackage ./package.nix {};
      }
    );

    devShells = forAllSystems (system: let
        pkgs = nixpkgsFor.${system};
      in {
        default = pkgs.mkShell {
          buildInputs = with pkgs; [
            go
            gopls
          ];
        };
      }
    );

  };
}
