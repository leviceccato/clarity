extends TextureButton

func _pressed() -> void:
	var code := get_tree().change_scene("res://scenes/Main.tscn")
	if code != OK:
		print("Unable to switch to Main scene %s" % code)
