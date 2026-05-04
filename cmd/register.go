package cmd

import (
	"sync"

	dcf "github.com/df-mc/dragonfly/server/cmd"

	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/we/edit"
)

var registerOnce sync.Once

// commandDefs is the single source of truth for Dragonfly cmd registration.
// Names starting with "/" support the Bedrock double-slash UX: "//set" strips
// one slash and resolves to command name "/set".
var commandDefs = []struct {
	name, desc string
	aliases    []string
	r          dcf.Runnable
}{
	{"/wand", "Varita de selección", []string{"wand"}, WandCommand{}},
	{"/pos1", "Definir primera posición", []string{"pos1"}, Pos1Command{}},
	{"/pos2", "Definir segunda posición", []string{"pos2"}, Pos2Command{}},
	{"/set", "Rellenar área seleccionada", []string{"set", "/fill", "fill"}, SetCommand{}},
	{"/copy", "Copiar área seleccionada", []string{"copy"}, CopyCommand{}},
	{"/paste", "Pegar portapapeles", []string{"paste"}, PasteCommand{}},
	{"/clearclipboard", "Limpiar portapapeles", []string{"clearclipboard"}, ClearClipboardCommand{}},
	{"/cut", "Cortar área seleccionada", []string{"cut"}, CutCommand{}},
	{"/schematic", "Gestionar esquemas", []string{"schematic", "/schem", "schem"}, SchematicCommand{}},
	{"/undo", "Deshacer cambio", []string{"undo"}, UndoCommand{}},
	{"/redo", "Rehacer cambio", []string{"redo"}, RedoCommand{}},
	{"/center", "Marcar centro de la selección", []string{"center"}, CenterCommand{}},
	{"/walls", "Construir muros de la selección", []string{"walls"}, WallsCommand{}},
	{"/drain", "Drenar fluidos", []string{"drain"}, DrainCommand{}},
	{"/biome", "Listar o definir biomas", []string{"biome"}, BiomeCommand{}},
	{"/replace", "Reemplazar bloques seleccionados", []string{"replace"}, ReplaceCommand{}},
	{"/replacenear", "Reemplazar bloques cercanos", []string{"replacenear"}, ReplaceNearCommand{}},
	{"/toplayer", "Reemplazar capa superior", []string{"toplayer"}, TopLayerCommand{}},
	{"/overlay", "Superponer capa superior", []string{"overlay", "/layer", "layer"}, OverlayCommand{}},
	{"/removeabove", "Eliminar bloques sobre el jugador", []string{"removeabove"}, RemoveAboveCommand{}},
	{"/removebelow", "Eliminar bloques bajo el jugador", []string{"removebelow"}, RemoveBelowCommand{}},
	{"/removenear", "Eliminar bloques cercanos coincidentes", []string{"removenear"}, RemoveNearCommand{}},
	{"/naturalize", "Naturalizar terreno seleccionado", []string{"naturalize"}, NaturalizeCommand{}},
	{"/move", "Mover selección", []string{"move"}, MoveCommand{}},
	{"/stack", "Apilar selección", []string{"stack"}, StackCommand{}},
	{"/rotate", "Rotar portapapeles", []string{"rotate"}, RotateCommand{}},
	{"/flip", "Voltear portapapeles", []string{"flip"}, FlipCommand{}},
	{"/line", "Dibujar línea de pos1 a pos2", []string{"line"}, LineCommand{}},
	{"/sphere", "Crear esfera", []string{"sphere"}, ShapeCommand{Kind: edit.ShapeSphere}},
	{"/cylinder", "Crear cilindro", []string{"cylinder"}, ShapeCommand{Kind: edit.ShapeCylinder}},
	{"/pyramid", "Crear pirámide", []string{"pyramid"}, ShapeCommand{Kind: edit.ShapePyramid}},
	{"/cone", "Crear cono", []string{"cone"}, ShapeCommand{Kind: edit.ShapeCone}},
	{"/cube", "Crear cubo", []string{"cube"}, ShapeCommand{Kind: edit.ShapeCube}},
	{"/brush", "Vincular brocha al ítem en mano", []string{"brush"}, BrushCommand{}},
	{"/searchitem", "Buscar ítems registrados", []string{"searchitem", "/search", "search", "/l", "l"}, SearchItemCommand{}},
}

func registerAll() {
	for _, e := range commandDefs {
		reg(e.name, e.desc, e.aliases, e.r)
	}
}

func init() {
	registerOnce.Do(registerAll)
}

// RegisterCommands is idempotent; commands are normally registered from init
// when this package is imported. Call this if you import this package indirectly
// and need to force registration before init ordering would run.
func RegisterCommands() {
	registerOnce.Do(registerAll)
}

func reg(name, desc string, aliases []string, r ...dcf.Runnable) {
	dcf.Register(dcf.New(name, desc, aliases, r...))
}

type playerCommand struct{}

func (playerCommand) Allow(src dcf.Source) bool {
	_, ok := src.(*player.Player)
	return ok
}
