# LightMatter

**LightMatter** is a lightweight and user-friendly application designed to simplify the process of managing and launching Singularity containers. Built using Go and Fyne, it provides a clean graphical interface to streamline container workflows.

## Features

- Simplified management of Singularity containers.
- Easy-to-use interface for selecting and launching containers.
- Support for multiple containers in the same session.
- Lightweight design with minimal dependencies.

## Installation

1. Download the `.deb` package:
   ```bash
   wget https://github.com/Car85/lightmatter.git
   ```

2. Install the package:
   ```bash
   sudo dpkg -i lightmatter.deb
   ```

3. Resolve any missing dependencies:
   ```bash
   sudo apt-get install -f
   ```

## Usage

1. Launch the application:
   ```bash
   lightmatter
   ```

2. Select a directory containing `.sif` files (Singularity containers).
3. Browse and run containers with a single click.

## License

This project is licensed under the GNU General Public License v3.0. See the [LICENSE](LICENSE) file for details.

