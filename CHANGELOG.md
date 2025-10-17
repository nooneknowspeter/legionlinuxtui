## [0.1.2] - 2025-10-17

### Documentation

- Update CHANGELOG.md
- Add LICENSE
- Update CHANGELOG.md

### Miscellaneous Tasks

- _(nix)_ Update default devshell packages
- Go mod vendor
- Update treefmt config
- Go mod vendor
- Release 0.1.2

## [0.1.1] - 2025-10-17

### Features

- Add filtering
- Update sensor information function

### Refactor

- Update sensor information paths
- Fix punctuation in command exec error

### Documentation

- Update TODO.md

### Miscellaneous Tasks

- _(ci/cd)_ Update build and release workflow
- _(nix)_ Update default devshell packages
- _(ci/cd)_ Add changelog workflow
- _(ci/cd)_ Update build and release workflow
- Add git-cliff configuration
- Release 0.1.1

## [0.1.0] - 2025-10-16

### Features

- Read persistant files
- Write to persistant file
- Add paths to sysfs driver files
- System information
- Cpu information
- Gpu information
- Battery information
- Fan information
- Define driver module function struct type
- Implement reading and updating driver value methods
- Implement methods to cycle through driver modes
- Add command driver modes
- Add camera power driver toggle
- Add conservation mode driver toggle
- Add fnlock driver toggle
- Add overdrive driver toggle
- Add fan controller drivers toggle
- Add touchpad driver toggle
- Add usb charging driver toggle
- Add rapid charge driver toggle
- Add winkey driver toggle
- Add powermode driver toggle
- Add ui-ux figma mockup
- Update ui-ux figma mockup
- Add status component
- Add battery indicator component
- Add streamline temperature chart component
- Add global lipgloss styles
- Add status pane
- Setup bubbletea fullscreen tui canvas
- Add status pane to tui view
- Add tui update ticker function and command
- Run tui through main entrypoint
- Update global styles
- Cli

### Bug Fixes

- Clear screen after selection

### Other

- Add README.md

### Refactor

- Move drivers/driver toggle file -> system/sysfs
- Rename helpers/io -> helpers/rw
- Rename tests/driver_test -> system/tests/sysfs_test
- Rename tests/sysinfo_test -> system/tests/sysinfo_test
- Rename tests/io_test -> helpers/tests/rw_test
- Add symbols directly to print statement strings
- Update idea drivers path
- Update sysfs driver functions
- Remove status components output style
- Status pane output
- Remove terminalWidth parameter
- Remove unnecessary revive comments
- Add comments to rw helper functions
- Remove unnecessary type declaration; infer from right side
- Add comments to driver sys path constants
- Disabled toggle status stdout
- Update driver function display keys
- Update cpu and gpu status lines

### Documentation

- Add TODO.md
- Update TODO.md
- Update TODO.md
- Update TODO.md
- Update TODO.md
- Update README.md
- Update README.md
- Update TODO.md
- Update README.md
- Update TODO.md

### Styling

- Flake.nix
- Treefmt

### Testing

- Add camera power mode toggle test
- Add conservation mode toggle test
- Add fnlock toggle test
- Add powermode toggle test
- Add rapid charge toggle test
- Add touchpad toggle test
- Add usb charging toggle test
- Add winkey toggle test
- Add nerdfont symbols for stdout
- Add rw file test
- Add battery information test
- Add cpu information test
- Add gpu information test
- Add system information test
- Update sysfs tests

### Miscellaneous Tasks

- Update flake.nix description
- Treefmt --init
- Add gofmt to treefmt
- Add goimports to treefmt
- Add nixfmt-rfc-style to treefmt
- Add Makefile
- Add gotools to flake.nix default devshell packages
- _(git)_ Ignore go artifacts
- Update treefmt.toml
- Update flake.nix
- Go mod init
- _(git)_ Ignore test artifacts
- Add quick command to run tests in verbose mode
- Update dependencies
- Export tui design wireframes
- _(ci/cd)_ Add static analysis workflow
- Add treefmt lint config
- _(nix)_ Update default devshell packages
- Update Makefile
- _(nix)_ Update default devshell packages
- _(ci/cd)_ Add build and upload binary workflow
- Release 0.1.0
