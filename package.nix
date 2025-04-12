{
  lib,
  buildGoModule,
  ...
}:
buildGoModule {
  name = "stallmanix";

  src = ./.;
  vendorHash = null;

  # removes debug info, making the binary smaller
  ldflags = ["-s" "-w"];

  meta = with lib; {
    mainProgram = "stallmanix";
    platforms = platforms.linux;
  };
}
