extends Node2D

@onready var icon = $Icon

func _init():
	print("init")

func _ready() -> void:
	print("ready")

	if OS.is_debug_build():
		print("debug build")
	else:
		print("release build")

func _process(delta: float) -> void:
	icon.position.x += 400 * delta

	var screen_size = get_viewport_rect().size
	var iconWidth = icon.texture.get_width()

	if icon.position.x > screen_size.x:
		icon.position.x = - iconWidth

func _input(event: InputEvent) -> void:
	if event is InputEventKey and event.pressed and event.keycode == KEY_ESCAPE:
		get_tree().quit()
