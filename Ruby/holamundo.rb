# PUTS keyboard para imprimir texto.
puts "1. hola alumnos de platzi"
puts "2. hola!"
puts "3. yay"

# ERRORS -> Si se invoca a un KEYBOARD ó METHOD saltará un alert.
# Buscar errores en la documentación.

nombre_del_curso = "Curso de Ruby"
puts nombre_del_curso

nombre = "Gabriel"

puts "Hola #{nombre}! Bienvenido al #{nombre_del_curso}"
puts "\n"

def welcome
	nombre_del_curso2 = "Curso de Ruby"
	nombre2 = "Gabriel"
	puts "Hola #{nombre2}! Bienvenido al #{nombre_del_curso2}"
end

welcome
