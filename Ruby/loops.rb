# WHILE - Elementos de control de flujo.

x = 0
numero = 5

while x < 5
	puts "Hola estoy en la posición No. #{x}"
	x += 1
end
puts "Fin"

# UNTIL - No necesita la variable de referencia, sino que utiliza la misma
# variable de incremento para saber cuantas veces tiene que repetirse.

x = 0 
until x == 5
	puts "Hola estoy en until #{x}"
	x += 1
end
puts "Fin"

x = 0
loop do 
	puts "Hola, estoy dentro del loop #{x}"
	x += 1
	break if x == 5
end
puts "Fin"


# FOR CYCLE - No es necesario declarar la variable de referencia.

for x in (0..4)
	puts "Hola estoy en el for #{x}"
end