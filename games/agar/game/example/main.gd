extends Node2D

# TODO (hub33k): add tests

const packets := preload("res://packets.gd")

@onready var icon = $Icon

func _init():
	print("init")

func _ready() -> void:
	print("ready")

	if OS.is_debug_build():
		print("debug build")
	else:
		print("release build")

	print()
	serialize_data()
	print()
	deserialize_data()

func serialize_data():
	var packet := packets.Packet.new()
	packet.set_sender_id(69)
	var chat_msg = packet.new_chat()
	chat_msg.set_msg("Hello, world!")
	# print(packet)

	var data := packet.to_bytes()

	print("serialize data")
	print(data) # bytes

func deserialize_data():
	var data := [8, 69, 18, 15, 10, 13, 72, 101, 108, 108, 111, 44, 32, 119, 111, 114, 108, 100, 33]
	var packet := packets.Packet.new()
	packet.from_bytes(data)

	print("deserialize data")
	print(packet)

func _process(delta: float) -> void:
	icon.position.x += 400 * delta

	var screen_size = get_viewport_rect().size
	var iconWidth = icon.texture.get_width()

	if icon.position.x > screen_size.x:
		icon.position.x = - iconWidth

func _input(event: InputEvent) -> void:
	if event is InputEventKey and event.pressed and event.keycode == KEY_ESCAPE:
		get_tree().quit()
