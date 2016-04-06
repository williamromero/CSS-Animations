# CREAR PROGRAMA PARA CONVERTIR DÓLARES A MONEDA LOCAL

# def namefunction (local variable)
# Se puede predefinir el valor del argumento-variable local para no tener que enviarlo
# al momento de llamar al método.
# puts dollar_to_currency(200) -> Sin error en consola.



def dollar_to_currency(dollars,tipo_cambio=7.80)
	return dollars * tipo_cambio
end

dol_convert = 200

puts dollar_to_currency(dol_convert, 10)
# Si se predefine el valor del argumento en el método, puede no otorgarselé un valor al llamar la función.
puts dollar_to_currency(dol_convert)
