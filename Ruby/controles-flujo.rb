# REQUERIMIENTO:
# 	- Clasificación de usuarios de acuerdo a su edad

# age = gets.to_i

# if age >= 0 && age <=2
# 	puts "Bebé"
# elsif age >= 3 && age <= 12
# 	puts "Niño"
# elsif age >= 13 && age <=17
# 	puts "Adolescente"
# elsif age <= 60
# 	puts "Adulto"
# else
# 	puts "Anciano"
# end	


# case age
# when 0..2		# Es lo mismo que 0,1,2
# 	puts "Bebe"
# when 3..12
# 	puts "Niño"
# when 13..17
# 	puts "Adolescente"
# when 18..60
# 	puts "Adulto"
# else
# 	puts "Anciano"
# end


def clasificador(age)
	case age
	when 0..2
		puts "RFC = Bebe"
	when 3..12
		puts "RFC = Niño"
	when 13..17
		puts "RFC = Adolescente"
	when 18..60
		puts "RFC = Adulto"
	when 61..100
		puts "RFC = Anciano"
	end
end

puts "Ingrese su edad: "
edad = gets.to_i
puts clasificador(edad)
