puts "Ingresa tu edad:\s"
edad = gets.to_i
mayor_edad = 18

# puts (edad == mayor_edad)
# puts (edad >= mayor_edad)

# puts "mayor de edad" == "mayor de EDAD"
# puts "joven" == "joven"
# puts "joven" != "joveN"


# && == and | && tiene mayor precedencia que "and"

if edad > mayor_edad
	puts "Eres mayor de edad, puedes entrar"
elsif edad < 30
	puts "Tienes menos de 30"
else
	puts "Eres menor de edad, no puedes entrar"
end

# ESTRUCTURA AND

if edad > 13 && edad < 18
	puts "Eres un adolecente"
end

# ESTRUCTURA OR

if (edad == 15) || (edad == 16)
	puts "yay, felicidades princesa... LOL"
end
