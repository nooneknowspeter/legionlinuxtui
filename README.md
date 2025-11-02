# legionlinuxtui

> [!NOTE]
>
> Ensure that the [Lenovo Legion Linux Drivers](https://github.com/johnfanv2/LenovoLegionLinux)
> and [`lm_sensors`](https://github.com/lm-sensors/lm-sensors) are installed properly on your system.

A terminal user interface for controlling Lenovo Legion laptops.

It utilizes the [Lenovo Legion Linux Drivers](https://github.com/johnfanv2/LenovoLegionLinux).

[![demo1](https://asciinema.org/a/S2jHeMt6bLGDgOlsIMJwZH7PU.svg)](https://asciinema.org/a/S2jHeMt6bLGDgOlsIMJwZH7PU)

![demo2](./public/legionlinuxtuidemo.gif)

## Features

- Sensor information display (100ms update)
- Capability toggles; convservative mode, power mode, ...
- Additional information

## Installation

### NixOS

#### NUR

Users who have the [NUR](https://github.com/nix-community/NUR) setup, you can install [legionlinuxtui](https://nur.nix-community.org/repos/nooneknowspeter/) using:

```nix
nur.repos.nooneknowspeter.legionlinuxtui
```

#### Flakes

You can install legionlinuxtui without having the [NUR](https://github.com/nix-community/NUR)
setup by using my [NUR flake](https://github.com/nooneknowspeter/nurpkgs) directly.

In your `flake.nix`, add the repo as an input:

```nix
inputs = {
    nooneknowspeter = {
      url = "github:nooneknowspeter/nurpkgs";
    };
}
```

> [!NOTE]
>
> Ensure that you [inherit the inputs](https://nixos.wiki/wiki/Flakes) in your NixOS or Home-Manager configurations

> [!NOTE]
>
> Either `inputs.nooneknowspeter.packages.x86_64-linux.legionlinuxtui`
> or `inputs.nooneknowspeter.packages.x86_64-linux.legionlinuxtui-git` can be used at a time.
>
> `legionlinuxtui-git` builds from the latest commit from the main branch (rolling release)
> and `legionlinux` builds from each tagged version, so pick your poison.

Add the package to your system environment packages:

```nix
environment.systemPackages = with pkgs; [
    inputs.nooneknowspeter.packages.x86_64-linux.legionlinuxtui
]
```

or in Home-Manager packages:

```nix
home.packages = with pkgs; [
    inputs.nooneknowspeter.packages.x86_64-linux.legionlinuxtui
]
```

> [!NOTE]
>
> Run the application with `sudo` privileges.

## Setup

> [!NOTE]
>
> Ensure Go >= 1.24.4 is installed.

#### Nix

> [!NOTE]
>
> Ensure [Nix](https://nixos.org/) is installed and
> [experimental features](https://nixos.wiki/wiki/Flakes) are on.

A Nix [flake](./flake.nix) development shell declaration is available in the repository.

Enter the shell using:

```sh
nix develop
```

> [!NOTE]
>
> A [`.envrc`](./.envrc) is also available in the repository for
> [`direnv`](https://github.com/direnv/direnv) users

## Build and Run

> [!NOTE]
>
> A [`Makefile`](./Makefile) is avaiable with quick commands.

> [!NOTE]
>
> The application needs `sudo` privileges for I/O; interacting with the drivers,
> changing values in `sysfs`

Run the application:

```sh
sudo go run .
```

or:

```
go build -o build/
sudo ./build/legionlinuxtui
```

## Contributing

Although this project is suited to my needs and wants;
add my own features and fix my own issues,
I will gladly appreciate and accept any PRs and issues that help the project.

### PRs

> [!NOTE]
>
> Ensure code is formatted with [`treefmt`](https://github.com/numtide/treefmt).
>
> A [`treefmt.toml`](./treefmt.toml) is available in the repository.

> [!NOTE]
>
> Use the `make lint` command in [`Makefile`](./Makefile) for linting.

CI/CD is available, PRs will get checked for any errors before merging.

### Feedback

Leaving issues; bugs and requested features are appreciated.
They will help the project making it stable and practical as possible for others.

### TODO

I keep track of the project tasks [here](./TODO.md) as well as in the codebase; `TODO:`.

## Resources

- https://github.com/johnfanv2/LenovoLegionLinux
