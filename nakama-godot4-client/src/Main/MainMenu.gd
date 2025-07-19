extends Node

# Maximum number of times to retry a server request if the previous attempt failed.
const MAX_REQUEST_ATTEMPTS := 3

# Path to the scene to load after selecting a character.
# export (String, FILE) var next_scene_path := ""
@export_file("*.tscn") var next_scene_path: String = ""

var _server_request_attempts := 0

#onready var login_and_register := $CanvasLayer/LoginAndRegister
#onready var character_menu := $CanvasLayer/CharacterMenu

func authenticate_user_async(email: String, password: String, do_remember_email := false) -> int:
	pass
	
func create_character_async(name: String, color: Color) -> void:
	pass

func delete_character_async(index: int) -> void:
	pass

func join_game_world_async(player_name: String, player_color: Color) -> int:
	pass
	
func open_character_menu_async() -> void:
	pass
