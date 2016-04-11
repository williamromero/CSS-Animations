puts "Introduzca el valor que desea cambiar:"
dolares = gets.to_f

puts "Monto del tipo de cambio"
cambio = gets.to_f

def dollar_to_currency(dollars,tipo_cambio=7.80)
	return dollars * tipo_cambio
end

puts dollar_to_currency(dolares, cambio).floor


