# LightMatter

![Go](https://img.shields.io/badge/Go-1.21-blue?style=flat-square&logo=go)
![Debian](https://img.shields.io/badge/Debian-11-red?style=flat-square&logo=debian)
![Linux](https://img.shields.io/badge/Linux-Universal-orange?style=flat-square&logo=linux)
![License](https://img.shields.io/badge/License-GPL--3.0-green?style=flat-square&logo=gnu)


**LightMatter** is a lightweight and user-friendly application designed to simplify the process of managing and launching Singularity containers. Built using Go and Fyne, it provides a clean graphical interface to streamline container workflows.


## Features

- Simplified management of Singularity containers.
- Easy-to-use interface for selecting and launching containers.
- Support for multiple containers in the same session.
- Lightweight design with minimal dependencies.

## Demo

[LightMatter Demo](https://github.com/user-attachments/assets/b75f77cf-a6e7-4c41-ad2f-4510c6f042e6)


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

---

## License

This project is licensed under the GNU General Public License v3.0. See the [LICENSE](LICENSE) file for details.

