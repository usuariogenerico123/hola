package funcs

import (
	
	"testing"
)


func TestVerifyDataType(t *testing.T){

	cases := []struct{
		name string
		dato1, dato2 any 
		esperado bool
	}{
		{"Datos iguales string" ,"hola", "otro dato", true},
		{"Datos distintos","perro", 23, false},
		{"Dato con nil",nil, "algo", false},
		{"Datos iguales enteros", 4, 65, true},

	}

	for _, c := range cases{
		t.Run(c.name, func(t *testing.T) {
			resp := VerifyDataType(c.dato1, c.dato2)
			if(resp != c.esperado){
				t.Fatal("error")
			}
		})

	}
}

func TestVerifyNumber(t *testing.T){
	cases := []struct{
		nombre string
		dato string 
		esperado bool
	}{
		{"Numero real en string", "28888", true},
		{"Letras en string", "asdhkj", false},
		{"Espacios", " ", false},
		
	}
	for _, v := range cases{
		t.Run(v.nombre, func(t *testing.T) {
			resp := VerifyNumber(v.dato)
			if(resp != v.esperado){
				t.Fatal("error")
			}
		})
	}

}
