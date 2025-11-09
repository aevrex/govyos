module github.com/aevrex/soren

go 1.25.3

retract (
    // Replace X.Y.Z with the version range that had the 'client' package.
    // If it was only present in version v0.1.1, use:
    v0.1.4

    // If it was present from v0.1.0 up to the latest known version v0.1.2, use:
    [v0.1.0, v0.1.4]
)