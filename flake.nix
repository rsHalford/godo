{
  description = "A command line todo list application";

  inputs.nixpkgs.url = "nixpkgs/nixos-unstable";

  outputs = { self, nixpkgs }:
    let
      version = "0.16.1";
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
              golangci-lint
              gopls
              pre-commit
            ];
          };
          godo = pkgs.buildGoModule {
            pname = "godo";
            inherit version;
            src = ./.;
            # vendorSha256 = pkgs.lib.fakeSha256;
            vendorSha256 = "sha256-EXVS17zcIB/wZeOqjn3KL9Pm1lstbI4AmDVlpTtPedQ=";
          };
        });
      defaultPackage = forAllSystems (system: self.packages.${system}.godo);
    };
}
