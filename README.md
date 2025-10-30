# pkgstate

Declarative package management for Arch Linux. Define the desired state of installed packages in a set of YAML files, and `pkgstate` will ensure that
the system matches that state.

In addition to installed packages, `pkgstate` can also track enabled/disabled systemd units (services, timers, etc.) and which groups the current user
is a member of, as some packages need specific group memberships to function properly.


## How to use

1. `pkgstate` is meant to used in a dotfiles repository. If you don't have one yet, I recommend checking out my other project [doot](https://github.com/pol-rivero/doot),
    but you can use any dotfiles manager you like.

1. Create a directory for your package definitions. By default, `~/.config/packages/` is used, but you can change this with the `PKGSTATE_DIR` environment variable.

1. Inside that directory, create one or more YAML files defining the desired state of your system. If there are multiple files, they will be merged together.
    ```yaml
    # Packages to be installed. All other packages will be uninstalled.
    packages:
      - base-devel
      - git
      - ssh
      - docker

    # User groups the current user should be a member of. Groups not listed here will be ignored (not modified).
    groups:
      - wheel
      - docker

    # Systemd units to be enabled/disabled (systemd --system). Units not listed here will be ignored (not modified).
    systemd: 
      "sshd.service": "enabled"
      "docker.socket": "enabled"
      "docker.service": "disabled"

    # Systemd units to be enabled/disabled (systemd --user). Units not listed here will be ignored (not modified).
    systemd_user:
      "my-timer.timer": "enabled"
      "another-service.service": "disabled"
    ```

1. Run `pkgstate` to get the diff of the current state vs the desired state.

1. Run `pkgstate fix` to interactively apply the changes, or `pkgstate fix --yes` to apply them without confirmation.
