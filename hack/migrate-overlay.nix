final: prev: {
  go-migrate = prev.go-migrate.overrideAttrs (old: {
    tags = ["postgres" "sqlite3"];
  });
}