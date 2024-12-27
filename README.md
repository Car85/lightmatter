# LightMatter

**LightMatter** is a lightweight and user-friendly application designed to simplify the process of managing and launching Singularity containers. Built using Go and Fyne, it provides a clean graphical interface to streamline container workflows.

## Technologies Used

| Technology | Logo |
|------------|------|
| Go         | ![Go Logo](https://go.dev/blog/go-brand/Go-Logo_Blue.png) |
| Fyne       | ![Fyne Logo](https://raw.githubusercontent.com/fyne-io/fyne/master/logo/fyne.png) |
| Debian     | ![Debian Logo](https://www.debian.org/logos/openlogo-nd.svg) |
| Linux      | <img src="https://upload.wikimedia.org/wikipedia/commons/a/af/Tux.png" alt="Linux Logo" width="50" />


## Features

- Simplified management of Singularity containers.
- Easy-to-use interface for selecting and launching containers.
- Support for multiple containers in the same session.
- Lightweight design with minimal dependencies.

## Installation

1. Download the `.deb` package:
   ```bash
   wget https://github.com/Car85/lightmatter/blob/main/lightmatter.deb

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

