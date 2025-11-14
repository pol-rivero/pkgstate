# pkgstate

Declarative package management for Arch Linux, without the pain of using Nix. Define the desired state of installed packages in a set of YAML files, and `pkgstate` will ensure that
the system matches that state.

In addition to installed packages, `pkgstate` can also track enabled/disabled systemd units (services, timers, etc.) and which groups the current user
is a member of, as some packages need specific group memberships to function properly.

## Installation

Install `pkgstate-bin` from the [AUR](https://aur.archlinux.org/packages/pkgstate-bin/).  
You can also build it from source by installing `pkgstate` or `pkgstate-git`.

## How to use

1. `pkgstate` is meant to used with a dotfiles repository. If you don't have one yet, I recommend checking [the `doot` Getting Started guide](https://github.com/pol-rivero/doot/wiki/Getting-Started),
    but you can use any dotfiles manager you like.

2. Create a directory for your package definitions. By default, `~/.config/packages/` is used, but you can change this with the `PKGSTATE_DIR` environment variable.

3. Inside that directory, create one or more YAML files defining the desired state of your system. To use your current state as a starting point, you can run `pkgstate generate`.

```yaml
# Packages to be explicitly installed. All other packages will be uninstalled (or marked as dependencies).
packages:
  - base-devel
  - git
  - ssh
  - docker

# User groups that the current user must be a member of. Groups not listed here will be ignored (not modified).
groups:
  - wheel
  - docker

# Systemd units to be enabled/disabled (systemd --system). Units not listed here will be ignored (not modified).
systemd: 
  sshd.service: enabled
  docker.socket: enabled
  docker.service: disabled

# Systemd units to be enabled/disabled (systemd --user). Units not listed here will be ignored (not modified).
systemd_user:
  my-timer.timer: enabled
  another-service.service: disabled
```

> [!TIP]
> All `.yaml` files in the config directory will be merged together. This is known as ["drop-in" configuration](https://devoptimize.org/practices/drop-in-configuration/), and allows you to have separate files for specific use cases.  
> For example, `printing.yaml` would define all packages, user groups, and systemd units related to controlling a printer.

4. Run `pkgstate` to get the diff of the current state vs. the desired state.

5. Run `pkgstate fix` to interactively apply the changes, or `pkgstate fix --yes` to apply them without confirmation.


## FAQs

### Should I run `pkgstate` as root?

> No, run it as a normal user. `pkgstate` will invoke `sudo` when necessary.

### Which package managers are supported?

> The currently supported managers are `pacman`, `yay`, and `paru`.  
`pkgstate` will automatically detect and use the best available package manager on your system.  
Please note that, if `pacman` is the only available package manager, you won't be able to install AUR packages using `pkgstate`.

### Can I use `pkgstate` on non-Arch Linux distributions?

> Yes, as long as they use `pacman` and `systemd`. This includes most Arch-based distributions like Manjaro, EndeavourOS, etc.

### Does `pkgstate` support per-host configurations?

> Not directly, and that's by design.  
Per-host configs should be the responsibility of your dotfiles manager, not individual programs like `pkgstate`. This ensures that host-specific logic is centralized in a single place, making it easier to manage and debug. For example, if you rename your computer, you only need to change the name in the dotfiles manager configuration, rather than updating multiple programs' configs.  
>
> In its place, `pkgstate` supports drop-in configuration (see "Tip" above), which allows your dotfiles manager to install the appropriate config files for each host.
