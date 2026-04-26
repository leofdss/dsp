## Container

```bash
mkdir ~/projects/proj-foundation-box -p
```

```bash
distrobox create -i ubuntu:24.04 \
	--name proj-foundation-box \
	--home ~/projects/proj-foundation-box/
```

```bash
distrobox enter proj-foundation-box
```

## Dependencies

```bash
sudo apt update
sudo apt install build-essential git golang
```

## Vscodium

```bash
wget -qO - https://gitlab.com/paulcarroty/vscodium-deb-rpm-repo/raw/master/pub.gpg \
    | gpg --dearmor \
    | sudo dd of=/usr/share/keyrings/vscodium-archive-keyring.gpg
```

```bash
echo -e 'Types: deb\nURIs: https://download.vscodium.com/debs\nSuites: vscodium\nComponents: main\nArchitectures: amd64 arm64\nSigned-by: /usr/share/keyrings/vscodium-archive-keyring.gpg' \
| sudo tee /etc/apt/sources.list.d/vscodium.sources
```

```bash
sudo apt update && sudo apt install codium
```

```bash
distrobox-export --app codium
```

### Alias

```bash
echo 'alias code="codium"' >> ~/.bashrc
echo 'alias podman="distrobox-host-exec podman"' >> ~/.bashrc
```