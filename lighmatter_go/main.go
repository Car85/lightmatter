package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Crear la aplicación
	myApp := app.NewWithID("com.lightmatter.launcher")
	myApp.Settings().SetTheme(theme.DarkTheme()) // Aplicar tema oscuro para apariencia moderna
	myWindow := myApp.NewWindow("LightMatter - Singularity Launcher")

	var sifFiles []string         // Archivos .sif detectados
	var selectedDirectory string  // Carpeta seleccionada por el usuario
	var selectedIndex int = -1    // Índice del contenedor seleccionado
	var mu sync.Mutex             // Control de concurrencia

	// Etiqueta para mostrar la carpeta seleccionada
	directoryLabel := widget.NewLabelWithStyle(
		"No directory selected",
		fyne.TextAlignLeading,
		fyne.TextStyle{Italic: true},
	)

	// Lista para mostrar los archivos .sif
	list := widget.NewList(
		func() int { return len(sifFiles) },
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(id widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(sifFiles[id])
		},
	)
	list.OnSelected = func(id widget.ListItemID) {
		selectedIndex = id
	}

	// Botón para seleccionar carpeta
	selectDirButton := widget.NewButtonWithIcon("Select Directory", theme.FolderOpenIcon(), func() {
		dialog.ShowFolderOpen(func(uri fyne.ListableURI, err error) {
			if err != nil {
				dialog.ShowError(fmt.Errorf("error selecting folder: %v", err), myWindow)
				return
			}
			if uri == nil {
				return
			}
			selectedDirectory = uri.Path()
			directoryLabel.SetText("Selected Directory: " + selectedDirectory)
			sifFiles = loadSifFiles(selectedDirectory)
			list.Refresh()
		}, myWindow)
	})

	// Botón para ejecutar contenedor
	runButton := widget.NewButtonWithIcon("Run", theme.MediaPlayIcon(), func() {
		if selectedIndex < 0 || selectedIndex >= len(sifFiles) {
			dialog.ShowInformation("Info", "Please select a container", myWindow)
			return
		}

		container := sifFiles[selectedIndex]
		go func() {
			mu.Lock()
			defer mu.Unlock()
			runSingularity(selectedDirectory, container, myWindow)
		}()
	})

	// Layout para un diseño más moderno
	header := container.NewVBox(
		widget.NewLabelWithStyle("LightMatter - Modern Singularity Launcher", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		directoryLabel,
		selectDirButton,
	)
	content := container.NewBorder(
		header,                   // Top
		container.NewHBox(runButton), // Bottom: Botón de ejecución
		nil, nil,                 // Left, Right
		container.NewVScroll(list), // Center: Lista de contenedores
	)

	// Configuración de la ventana
	myWindow.SetContent(container.NewPadded(content)) // Agregar un padding para un diseño limpio
	myWindow.Resize(fyne.NewSize(450, 400))           // Tamaño inicial más compacto
	myWindow.ShowAndRun()
}

// loadSifFiles escanea el directorio en busca de archivos .sif
func loadSifFiles(dir string) []string {
	var sifFiles []string
	_ = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if strings.HasSuffix(info.Name(), ".sif") {
			sifFiles = append(sifFiles, info.Name())
		}
		return nil
	})
	return sifFiles
}

// runSingularity ejecuta un contenedor de Singularity
func runSingularity(dir, container string, parent fyne.Window) {
	fullPath := filepath.Join(dir, container)
	cmd := exec.Command("singularity", "run", fullPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		dialog.ShowError(fmt.Errorf("failed to run container: %v", err), parent)
		return
	}
	dialog.ShowInformation("Success", fmt.Sprintf("Container '%s' is running!", container), parent)
}

