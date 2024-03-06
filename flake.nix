{
  description = "GoDo: Note taking with SSH";

  inputs.nixpkgs.url = "nixpkgs/nixos-unstable";

  outputs = { self, nixpkgs }:
    let
      version = "2.0.0";
      supportedSystems = [ "x86_64-linux" "x86_64-darwin" "aarch64-linux" "aarch64-darwin" ];
      forAllSystems = nixpkgs.lib.genAttrs supportedSystems;
      nixpkgsFor = forAllSystems (system: import nixpkgs { inherit system; });

    in
    {
      packages = forAllSystems (system:
        let
          pkgs = nixpkgsFor.${system};
        in
        {
          devShell = pkgs.mkShell {
            buildInputs = with pkgs; [
              git-chglog
              gopls
            ];
          };
          godo = pkgs.buildGoModule {
            pname = "godo";
            inherit version;
            src = ./.;
            # vendorHash = pkgs.lib.fakeSha256;
            vendorHash = "sha256-EPsJFg8mhAVrTYskENltpdE7bvDoBoGud25vzZ38DVU=";
          };
        });
      defaultPackage = forAllSystems (system: self.packages.${system}.godo);
    };
}
