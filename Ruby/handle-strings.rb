# CADENAS - Las cadenas también son objetos, por lo cual tiene métodos.
	msg = "ruby es divertido\n"
	puts msg.capitalize
	puts msg.downcase
	puts msg.swapcase

	x = "alumnos de Platzi".insert(0,"Hola ")
	puts x
	puts x.reverse


	# Para substituir caracteres de strings
	bad_string = "Ruby es divretido"
	puts "Ej. "+bad_string
	puts "Res. "+bad_string.gsub('divretido', 'divertido')
